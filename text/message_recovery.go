// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package text

import (
	"fmt"
	"time"
)

func NewErrorValidationRecoveryFlowExpired(expiredAt time.Time) *Message {
	return &Message{
		ID:   ErrorValidationRecoveryFlowExpired,
		Text: fmt.Sprintf("恢复流程已于 %.2f 分钟前过期，请重试。", Since(expiredAt).Minutes()),
		Type: Error,
		Context: context(map[string]any{
			"expired_at":      expiredAt,
			"expired_at_unix": expiredAt.Unix(),
		}),
	}
}

func NewRecoverySuccessful(privilegedSessionExpiresAt time.Time) *Message {
	hasLeft := Until(privilegedSessionExpiresAt)
	return &Message{
		ID:   InfoSelfServiceRecoverySuccessful,
		Type: Success,
		Text: fmt.Sprintf("您已成功恢复账户。请在 %.2f 分钟内更改密码或设置替代登录方式（如社交登录）。", hasLeft.Minutes()),
		Context: context(map[string]any{
			"privilegedSessionExpiresAt":         privilegedSessionExpiresAt,
			"privileged_session_expires_at":      privilegedSessionExpiresAt,
			"privileged_session_expires_at_unix": privilegedSessionExpiresAt.Unix(),
		}),
	}
}

func NewRecoveryEmailSent() *Message {
	return &Message{
		ID:   InfoSelfServiceRecoveryEmailSent,
		Type: Info,
		Text: "恢复链接已发送至您提供的邮箱。如果未收到，请检查地址拼写并确保使用注册邮箱。",
	}
}

func NewRecoveryEmailWithCodeSent() *Message {
	return &Message{
		ID:   InfoSelfServiceRecoveryEmailWithCodeSent,
		Type: Info,
		Text: "恢复码已发送至您提供的邮箱。如果未收到，请检查地址拼写并确保使用注册邮箱。",
	}
}

func NewRecoveryAskAnyRecoveryAddress() *Message {
	return &Message{
		ID:   InfoNodeLabelRecoveryAddress,
		Text: "恢复地址",
		Type: Info,
	}
}

func NewRecoveryCodeRecoverySelectAddressSent(maskedAddress string) *Message {
	return &Message{
		ID:   InfoSelfServiceRecoveryMessageMaskedWithCodeSent,
		Type: Info,
		Text: fmt.Sprintf("恢复码已发送至 %s。如果未收到，请检查地址拼写并确保使用注册邮箱。", maskedAddress),
		Context: context(map[string]any{
			"masked_address": maskedAddress,
		}),
	}
}

func NewRecoveryAskForFullAddress() *Message {
	return &Message{
		ID:   InfoSelfServiceRecoveryAskForFullAddress,
		Type: Info,
		Text: "通过提供完整的恢复地址来恢复对账户的访问。",
	}
}

func NewRecoveryAskToChooseAddress() *Message {
	return &Message{
		ID:   InfoSelfServiceRecoveryAskToChooseAddress,
		Type: Info,
		Text: "您想如何恢复您的账户？",
	}
}

func NewRecoveryBack() *Message {
	return &Message{
		ID:   InfoSelfServiceRecoveryBack,
		Type: Info,
		Text: "返回",
	}
}

func NewErrorValidationRecoveryTokenInvalidOrAlreadyUsed() *Message {
	return &Message{
		ID:   ErrorValidationRecoveryTokenInvalidOrAlreadyUsed,
		Text: "恢复令牌无效或已被使用，请重试流程。",
		Type: Error,
	}
}

func NewErrorValidationRecoveryCodeInvalidOrAlreadyUsed() *Message {
	return &Message{
		ID:   ErrorValidationRecoveryCodeInvalidOrAlreadyUsed,
		Text: "恢复码无效或已被使用，请重试。",
		Type: Error,
	}
}

func NewErrorValidationRecoveryRetrySuccess() *Message {
	return &Message{
		ID:   ErrorValidationRecoveryRetrySuccess,
		Text: "请求已完成，无法重试。",
		Type: Error,
	}
}

func NewErrorValidationRecoveryStateFailure() *Message {
	return &Message{
		ID:   ErrorValidationRecoveryStateFailure,
		Text: "恢复流程已达到失败状态，必须重试。",
		Type: Error,
	}
}
