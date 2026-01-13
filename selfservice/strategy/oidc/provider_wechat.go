// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package oidc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"

	"github.com/ory/herodot"
	"github.com/ory/x/httpx"

	"github.com/hashicorp/go-retryablehttp"
)

// ProviderWeChat implements the OAuth2 protocol for WeChat (微信) login.
// This supports WeChat Open Platform website applications (网站应用).
//
// WeChat OAuth2 Flow:
// 1. Redirect user to: https://open.weixin.qq.com/connect/qrconnect
// 2. User scans QR code and authorizes
// 3. WeChat redirects back with authorization code
// 4. Exchange code for access_token (which also contains openid and unionid)
// 5. Use access_token + openid to fetch user info
//
// Documentation: https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
type ProviderWeChat struct {
	config *Configuration
	reg    Dependencies
}

var _ OAuth2Provider = (*ProviderWeChat)(nil)

func NewProviderWeChat(
	config *Configuration,
	reg Dependencies,
) Provider {
	return &ProviderWeChat{
		config: config,
		reg:    reg,
	}
}

func (g *ProviderWeChat) Config() *Configuration {
	return g.config
}

func (g *ProviderWeChat) oauth2(ctx context.Context) *oauth2.Config {
	// WeChat uses non-standard endpoints
	endpoint := oauth2.Endpoint{
		// PC website QR code login
		AuthURL: "https://open.weixin.qq.com/connect/qrconnect",
		// Token exchange endpoint
		TokenURL: "https://api.weixin.qq.com/sns/oauth2/access_token",
	}

	// WeChat requires scope: snsapi_login for website applications
	scopes := g.config.Scope
	if len(scopes) == 0 {
		scopes = []string{"snsapi_login"}
	}

	return &oauth2.Config{
		ClientID:     g.config.ClientID,
		ClientSecret: g.config.ClientSecret,
		Endpoint:     endpoint,
		Scopes:       scopes,
		RedirectURL:  g.config.Redir(g.reg.Config().OIDCRedirectURIBase(ctx)),
	}
}

func (g *ProviderWeChat) AuthCodeURLOptions(r ider) []oauth2.AuthCodeOption {
	// WeChat requires appid instead of client_id in auth URL
	// and requires #wechat_redirect at the end of URL
	return []oauth2.AuthCodeOption{
		// WeChat uses 'appid' instead of 'client_id'
		oauth2.SetAuthURLParam("appid", g.config.ClientID),
	}
}

func (g *ProviderWeChat) OAuth2(ctx context.Context) (*oauth2.Config, error) {
	return g.oauth2(ctx), nil
}

// weChatTokenResponse represents the response from WeChat's token endpoint.
// WeChat returns openid and unionid along with access_token.
type weChatTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionID      string `json:"unionid,omitempty"` // Only available if bound to WeChat Open Platform

	// Error fields
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

// weChatUserInfo represents the user info response from WeChat.
type weChatUserInfo struct {
	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"` // 1=male, 2=female, 0=unknown
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionID    string   `json:"unionid,omitempty"`

	// Error fields
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

// ExchangeOAuth2Token exchanges the authorization code for an access token.
// WeChat has a non-standard token endpoint that requires query parameters.
func (g *ProviderWeChat) ExchangeOAuth2Token(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	conf, err := g.OAuth2(ctx)
	if err != nil {
		return nil, err
	}

	// WeChat token endpoint uses GET with query parameters (non-standard)
	tokenURL := fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		conf.Endpoint.TokenURL,
		url.QueryEscape(conf.ClientID),
		url.QueryEscape(conf.ClientSecret),
		url.QueryEscape(code),
	)

	client := g.reg.HTTPClient(ctx, httpx.ResilientClientDisallowInternalIPs())
	req, err := retryablehttp.NewRequest("GET", tokenURL, nil)
	if err != nil {
		return nil, errors.WithStack(herodot.ErrInternalServerError.WithWrap(err).WithReasonf("failed to create token request: %s", err))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(herodot.ErrUpstreamError.WithWrap(err).WithReasonf("failed to exchange token: %s", err))
	}
	defer func() { _ = resp.Body.Close() }()

	if err := logUpstreamError(g.reg.Logger(), resp); err != nil {
		return nil, err
	}

	var tokenResp weChatTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, errors.WithStack(herodot.ErrUpstreamError.WithWrap(err).WithReasonf("failed to decode token response: %s", err))
	}

	if tokenResp.ErrCode != 0 {
		return nil, errors.WithStack(herodot.ErrUpstreamError.WithReasonf("WeChat token error: code=%d, msg=%s", tokenResp.ErrCode, tokenResp.ErrMsg))
	}

	if tokenResp.AccessToken == "" || tokenResp.OpenID == "" {
		return nil, errors.WithStack(herodot.ErrUpstreamError.WithReasonf("WeChat returned empty access_token or openid"))
	}

	token := &oauth2.Token{
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
		TokenType:    "Bearer",
	}

	// Store openid and unionid in token extras for later use in Claims()
	token = token.WithExtra(map[string]interface{}{
		"openid":  tokenResp.OpenID,
		"unionid": tokenResp.UnionID,
	})

	return token, nil
}

func (g *ProviderWeChat) Claims(ctx context.Context, exchange *oauth2.Token, _ url.Values) (*Claims, error) {
	// Get openid from token extras (set during token exchange)
	openID, ok := exchange.Extra("openid").(string)
	if !ok || openID == "" {
		return nil, errors.WithStack(herodot.ErrUpstreamError.WithReasonf("WeChat openid not found in token"))
	}

	// Fetch user info from WeChat API
	userInfoURL := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN",
		url.QueryEscape(exchange.AccessToken),
		url.QueryEscape(openID),
	)

	client := g.reg.HTTPClient(ctx, httpx.ResilientClientDisallowInternalIPs())
	req, err := retryablehttp.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		return nil, errors.WithStack(herodot.ErrInternalServerError.WithWrap(err).WithReasonf("failed to create userinfo request: %s", err))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(herodot.ErrUpstreamError.WithWrap(err).WithReasonf("failed to fetch user info: %s", err))
	}
	defer func() { _ = resp.Body.Close() }()

	if err := logUpstreamError(g.reg.Logger(), resp); err != nil {
		return nil, err
	}

	var user weChatUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, errors.WithStack(herodot.ErrInternalServerError.WithWrap(err).WithReasonf("failed to decode user info: %s", err))
	}

	if user.ErrCode != 0 {
		return nil, errors.WithStack(herodot.ErrUpstreamError.WithReasonf("WeChat userinfo error: code=%d, msg=%s", user.ErrCode, user.ErrMsg))
	}

	// Get unionid - prefer the one from userinfo, fallback to token
	unionID := user.UnionID
	if unionID == "" {
		if uid, ok := exchange.Extra("unionid").(string); ok {
			unionID = uid
		}
	}

	// Use UnionID as Subject if available (recommended for cross-app identity)
	// Otherwise fallback to OpenID
	subject := unionID
	if subject == "" {
		subject = user.OpenID
	}

	// Map WeChat gender to standard format
	var gender string
	switch user.Sex {
	case 1:
		gender = "male"
	case 2:
		gender = "female"
	default:
		gender = ""
	}

	// Build locale from country info
	var locale string
	if user.Country != "" {
		locale = user.Country
	}

	return &Claims{
		Issuer:   "https://api.weixin.qq.com",
		Subject:  subject,
		Nickname: user.Nickname,
		Name:     user.Nickname,
		Picture:  user.HeadImgURL,
		Gender:   gender,
		Locale:   Locale(locale),
		// Store raw openid and unionid in raw claims for JSONNet mapper access
		RawClaims: map[string]interface{}{
			"openid":     user.OpenID,
			"unionid":    unionID,
			"nickname":   user.Nickname,
			"sex":        user.Sex,
			"province":   user.Province,
			"city":       user.City,
			"country":    user.Country,
			"headimgurl": user.HeadImgURL,
			"privilege":  user.Privilege,
		},
	}, nil
}
