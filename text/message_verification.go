// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package text

import (
	"fmt"
	"time"
)

func NewErrorValidationVerificationFlowExpired(expiredAt time.Time) *Message {
	return &Message{
		ID:   ErrorValidationVerificationFlowExpired,
		Text: fmt.Sprintf("验证流程已于 %.2f 分钟前过期，请重试。", Since(expiredAt).Minutes()),
		Type: Error,
		Context: context(map[string]any{
			"expired_at":      expiredAt,
			"expired_at_unix": expiredAt.Unix(),
		}),
	}
}

func NewInfoSelfServiceVerificationSuccessful() *Message {
	return &Message{
		ID:   InfoSelfServiceVerificationSuccessful,
		Type: Success,
		Text: "您已成功验证邮箱地址。",
	}
}

func NewVerificationEmailSent() *Message {
	return &Message{
		ID:   InfoSelfServiceVerificationEmailSent,
		Type: Info,
		Text: "验证链接已发送至您提供的邮箱。如果未收到，请检查地址拼写并确保使用注册邮箱。",
	}
}

func NewErrorValidationVerificationTokenInvalidOrAlreadyUsed() *Message {
	return &Message{
		ID:   ErrorValidationVerificationTokenInvalidOrAlreadyUsed,
		Text: "验证令牌无效或已被使用，请重试流程。",
		Type: Error,
	}
}

func NewErrorValidationVerificationRetrySuccess() *Message {
	return &Message{
		ID:   ErrorValidationVerificationRetrySuccess,
		Text: "请求已完成，无法重试。",
		Type: Error,
	}
}

func NewErrorValidationVerificationStateFailure() *Message {
	return &Message{
		ID:   ErrorValidationVerificationStateFailure,
		Text: "验证流程已达到失败状态，必须重试。",
		Type: Error,
	}
}

func NewErrorValidationVerificationCodeInvalidOrAlreadyUsed() *Message {
	return &Message{
		ID:   ErrorValidationVerificationCodeInvalidOrAlreadyUsed,
		Text: "验证码无效或已被使用，请重试。",
		Type: Error,
	}
}

func NewVerificationEmailWithCodeSent() *Message {
	return &Message{
		ID:   InfoSelfServiceVerificationEmailWithCodeSent,
		Type: Info,
		Text: "验证码已发送至您提供的邮箱。如果未收到，请检查地址拼写并确保使用注册邮箱。",
	}
}
