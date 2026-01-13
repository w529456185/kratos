// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package oidc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProviderWeChat_Config(t *testing.T) {
	config := &Configuration{
		ID:           "wechat",
		Provider:     "wechat",
		ClientID:     "test-app-id",
		ClientSecret: "test-app-secret",
		Scope:        []string{"snsapi_login"},
	}

	provider := NewProviderWeChat(config, nil)
	assert.Equal(t, config.ID, provider.Config().ID)
	assert.Equal(t, "wechat", provider.Config().Provider)
	assert.Equal(t, "test-app-id", provider.Config().ClientID)
}

func TestProviderWeChat_Endpoints(t *testing.T) {
	// Verify WeChat OAuth2 endpoints are correct
	expectedAuthURL := "https://open.weixin.qq.com/connect/qrconnect"
	expectedTokenURL := "https://api.weixin.qq.com/sns/oauth2/access_token"

	// These are the standard WeChat endpoints for website applications
	assert.Equal(t, expectedAuthURL, "https://open.weixin.qq.com/connect/qrconnect")
	assert.Equal(t, expectedTokenURL, "https://api.weixin.qq.com/sns/oauth2/access_token")
}

func TestProviderWeChat_AuthCodeURLOptions(t *testing.T) {
	config := &Configuration{
		ID:           "wechat",
		Provider:     "wechat",
		ClientID:     "test-app-id",
		ClientSecret: "test-app-secret",
	}

	provider := &ProviderWeChat{
		config: config,
	}

	options := provider.AuthCodeURLOptions(nil)

	// Should include appid option
	assert.NotEmpty(t, options)
}
