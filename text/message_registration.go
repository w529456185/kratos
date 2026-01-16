// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package text

import (
	"fmt"
	"time"
)

func NewInfoRegistration() *Message {
	return &Message{
		ID:   InfoSelfServiceRegistration,
		Text: "注册",
		Type: Info,
	}
}

func NewInfoRegistrationWith(provider string, providerID string) *Message {
	return &Message{
		ID:   InfoSelfServiceRegistrationWith,
		Text: fmt.Sprintf("使用 %s 注册", provider),
		Type: Info,
		Context: context(map[string]any{
			"provider":    provider,
			"provider_id": providerID,
		}),
	}
}

func NewInfoRegistrationContinue() *Message {
	return &Message{
		ID:   InfoSelfServiceRegistrationContinue,
		Text: "继续",
		Type: Info,
	}
}

func NewInfoRegistrationBack() *Message {
	return &Message{
		ID:   InfoSelfServiceRegistrationBack,
		Text: "返回",
		Type: Info,
	}
}

func NewInfoSelfServiceChooseCredentials() *Message {
	return &Message{
		ID:   InfoSelfServiceRegistrationChooseCredentials,
		Text: "请选择一种凭证进行身份验证。",
		Type: Info,
	}
}

func NewErrorValidationRegistrationFlowExpired(expiredAt time.Time) *Message {
	return &Message{
		ID:   ErrorValidationRegistrationFlowExpired,
		Text: fmt.Sprintf("注册流程已于 %.2f 分钟前过期，请重试。", Since(expiredAt).Minutes()),
		Type: Error,
		Context: context(map[string]any{
			"expired_at":      expiredAt,
			"expired_at_unix": expiredAt.Unix(),
		}),
	}
}

func NewInfoSelfServiceRegistrationRegisterWebAuthn() *Message {
	return &Message{
		ID:   InfoSelfServiceRegistrationRegisterWebAuthn,
		Text: "使用安全密钥注册",
		Type: Info,
	}
}

func NewInfoSelfServiceRegistrationRegisterPasskey() *Message {
	return &Message{
		ID:   InfoSelfServiceRegistrationRegisterPasskey,
		Text: "使用 Passkey 注册",
		Type: Info,
	}
}

func NewRegistrationEmailWithCodeSent() *Message {
	return &Message{
		ID:   InfoSelfServiceRegistrationEmailWithCodeSent,
		Type: Info,
		Text: "验证码已发送至您提供的地址。如果未收到，请检查地址拼写并重试注册。",
	}
}

func NewErrorValidationRegistrationCodeInvalidOrAlreadyUsed() *Message {
	return &Message{
		ID:   ErrorValidationRegistrationCodeInvalidOrAlreadyUsed,
		Text: "注册码无效或已被使用，请重试。",
		Type: Error,
	}
}

func NewErrorValidationRegistrationRetrySuccessful() *Message {
	return &Message{
		ID:   ErrorValidateionRegistrationRetrySuccess,
		Type: Error,
		Text: "请求已完成，无法重试。",
	}
}

func NewInfoSelfServiceRegistrationRegisterCode() *Message {
	return &Message{
		ID:   InfoSelfServiceRegistrationRegisterCode,
		Text: "发送注册验证码",
		Type: Info,
	}
}
