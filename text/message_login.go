// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package text

import (
	"fmt"
	"time"
)

func NewInfoLoginReAuth() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginReAuth,
		Type: Info,
		Text: "请通过验证确认此操作。",
	}
}

func NewInfoLoginMFA() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginMFA,
		Type: Info,
		Text: "请完成第二重身份验证。",
	}
}

func NewInfoLoginWebAuthnPasswordless() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginWebAuthnPasswordless,
		Type: Info,
		Text: "准备好您的 WebAuthn 设备（如安全密钥、生物识别扫描仪等），然后点击继续。",
	}
}

func NewInfoLoginTOTPLabel() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginTOTPLabel,
		Type: Info,
		Text: "验证码",
	}
}

func NewInfoLoginLookupLabel() *Message {
	return &Message{
		ID:   InfoLoginLookupLabel,
		Type: Info,
		Text: "备用恢复码",
	}
}

func NewInfoLogin() *Message {
	return &Message{
		ID:   InfoSelfServiceLogin,
		Text: "登录",
		Type: Info,
	}
}

func NewInfoLoginLinkMessage(dupIdentifier, provider, newLoginURL string, availableCredentials, availableProviders []string) *Message {
	return &Message{
		ID:   InfoSelfServiceLoginLink,
		Type: Info,
		Text: fmt.Sprintf(
			"您尝试使用 %q 登录，但该邮箱已被其他账户使用。请使用以下选项登录您的账户，以将 %[1]q 添加为另一种登录方式。",
			dupIdentifier,
		),
		Context: context(map[string]any{
			"duplicateIdentifier":        dupIdentifier,
			"provider":                   provider,
			"newLoginUrl":                newLoginURL,
			"duplicate_identifier":       dupIdentifier,
			"new_login_url":              newLoginURL,
			"available_credential_types": availableCredentials,
			"available_providers":        availableProviders,
		}),
	}
}

func NewInfoLoginAndLink() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginAndLink,
		Text: "登录并关联",
		Type: Info,
	}
}

func NewInfoLoginTOTP() *Message {
	return &Message{
		ID:   InfoLoginTOTP,
		Text: "使用验证器",
		Type: Info,
	}
}

func NewInfoLoginPassword() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginPassword,
		Text: "使用密码登录",
		Type: Info,
	}
}

func NewInfoLoginLookup() *Message {
	return &Message{
		ID:   InfoLoginLookup,
		Text: "使用备用恢复码",
		Type: Info,
	}
}

func NewInfoLoginVerify() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginVerify,
		Text: "验证",
		Type: Info,
	}
}

func NewInfoLoginWith(provider string, providerId string) *Message {
	return &Message{
		ID:   InfoSelfServiceLoginWith,
		Text: fmt.Sprintf("使用 %s 登录", provider),
		Type: Info,
		Context: context(map[string]any{
			"provider":    provider,
			"provider_id": providerId,
		}),
	}
}

func NewInfoLoginWithAndLink(provider string) *Message {
	return &Message{
		ID:   InfoSelfServiceLoginWithAndLink,
		Text: fmt.Sprintf("使用 %s 确认", provider),
		Type: Info,
		Context: context(map[string]any{
			"provider": provider,
		}),
	}
}

func NewErrorValidationLoginFlowExpired(expiredAt time.Time) *Message {
	return &Message{
		ID:   ErrorValidationLoginFlowExpired,
		Text: fmt.Sprintf("登录流程已于 %.2f 分钟前过期，请重试。", Since(expiredAt).Minutes()),
		Type: Error,
		Context: context(map[string]any{
			"expired_at":      expiredAt,
			"expired_at_unix": expiredAt.Unix(),
		}),
	}
}

func NewErrorValidationLoginNoStrategyFound() *Message {
	return &Message{
		ID:   ErrorValidationLoginNoStrategyFound,
		Text: "未找到登录策略。您是否填写了正确的表单？",
		Type: Error,
	}
}

func NewErrorValidationRegistrationNoStrategyFound() *Message {
	return &Message{
		ID:   ErrorValidationRegistrationNoStrategyFound,
		Text: "未找到注册策略。您是否填写了正确的表单？",
		Type: Error,
	}
}

func NewErrorValidationSettingsNoStrategyFound() *Message {
	return &Message{
		ID:   ErrorValidationSettingsNoStrategyFound,
		Text: "未找到更新设置的策略。您是否填写了正确的表单？",
		Type: Error,
	}
}

func NewErrorValidationRecoveryNoStrategyFound() *Message {
	return &Message{
		ID:   ErrorValidationRecoveryNoStrategyFound,
		Text: "未找到恢复账户的策略。您是否填写了正确的表单？",
		Type: Error,
	}
}

func NewErrorValidationVerificationNoStrategyFound() *Message {
	return &Message{
		ID:   ErrorValidationVerificationNoStrategyFound,
		Text: "未找到验证账户的策略。您是否填写了正确的表单？",
		Type: Error,
	}
}

func NewInfoSelfServiceLoginWebAuthn() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginWebAuthn,
		Text: "使用硬件密钥登录",
		Type: Info,
	}
}

func NewInfoSelfServiceLoginPasskey() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginPasskey,
		Text: "使用 Passkey 登录",
		Type: Info,
	}
}

func NewInfoSelfServiceContinueLoginWebAuthn() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginContinueWebAuthn,
		Text: "使用硬件密钥登录",
		Type: Info,
	}
}

func NewInfoSelfServiceLoginContinue() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginContinue,
		Text: "继续",
		Type: Info,
	}
}

func NewLoginCodeSent() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginCodeSent,
		Type: Info,
		Text: "验证码已发送至您提供的地址。如果未收到，请检查地址拼写并重试登录。",
	}
}

func NewErrorValidationLoginCodeInvalidOrAlreadyUsed() *Message {
	return &Message{
		ID:   ErrorValidationLoginCodeInvalidOrAlreadyUsed,
		Text: "登录码无效或已被使用，请重试。",
		Type: Error,
	}
}

func NewErrorValidationLoginRetrySuccessful() *Message {
	return &Message{
		ID:   ErrorValidationLoginRetrySuccess,
		Type: Error,
		Text: "请求已完成，无法重试。",
	}
}

func NewInfoSelfServiceLoginCode() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginCode,
		Type: Info,
		Text: "发送登录验证码",
	}
}

func NewErrorValidationLoginLinkedCredentialsDoNotMatch() *Message {
	return &Message{
		ID:   ErrorValidationLoginLinkedCredentialsDoNotMatch,
		Text: "关联的凭证不匹配。",
		Type: Error,
	}
}

func NewErrorValidationAddressUnknown() *Message {
	return &Message{
		ID:   ErrorValidationLoginAddressUnknown,
		Text: "您输入的地址与当前账户中的任何已知地址不匹配。",
		Type: Error,
	}
}

func NewInfoSelfServiceLoginCodeMFA() *Message {
	return &Message{
		ID:   InfoSelfServiceLoginCodeMFA,
		Type: Info,
		Text: "请求验证码以继续",
	}
}

func NewInfoSelfServiceLoginAAL2CodeAddress(channel string, to string) *Message {
	return &Message{
		ID:   InfoSelfServiceLoginAAL2CodeAddress,
		Type: Info,
		Text: fmt.Sprintf("发送验证码至 %s", to),
		Context: context(map[string]any{
			"address": to,
			"channel": channel,
		}),
	}
}
