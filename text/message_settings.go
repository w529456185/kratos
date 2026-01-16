// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package text

import (
	"fmt"
	"strings"
	"time"
)

func NewErrorValidationSettingsFlowExpired(expiredAt time.Time) *Message {
	return &Message{
		ID:   ErrorValidationSettingsFlowExpired,
		Text: fmt.Sprintf("设置流程已于 %.2f 分钟前过期，请重试。", Since(expiredAt).Minutes()),
		Type: Error,
		Context: context(map[string]any{
			"expired_at":      expiredAt,
			"expired_at_unix": expiredAt.Unix(),
		}),
	}
}

func NewInfoSelfServiceSettingsTOTPQRCode() *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsTOTPQRCode,
		Text: "验证器应用二维码",
		Type: Info,
	}
}

func NewInfoSelfServiceSettingsTOTPSecret(secret string) *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsTOTPSecret,
		Text: secret,
		Type: Info,
		Context: context(map[string]any{
			"secret": secret,
		}),
	}
}
func NewInfoSelfServiceSettingsTOTPSecretLabel() *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsTOTPSecretLabel,
		Text: "这是您的验证器应用密钥。如果无法扫描二维码，请使用此密钥。",
		Type: Info,
	}
}

func NewInfoSelfServiceSettingsUpdateSuccess() *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsUpdateSuccess,
		Text: "您的更改已保存！",
		Type: Success,
	}
}

func NewInfoSelfServiceSettingsUpdateUnlinkTOTP() *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsUpdateUnlinkTOTP,
		Text: "取消关联 TOTP 验证器应用",
		Type: Info,
	}
}

func NewInfoSelfServiceSettingsRevealLookup() *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsRevealLookup,
		Text: "显示备用恢复码",
		Type: Info,
	}
}

func NewInfoSelfServiceSettingsRegenerateLookup() *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsRegenerateLookup,
		Text: "生成新的备用恢复码",
		Type: Info,
	}
}

func NewInfoSelfServiceSettingsDisableLookup() *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsDisableLookup,
		Text: "禁用此方法",
		Type: Info,
	}
}

func NewInfoSelfServiceSettingsLookupConfirm() *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsLookupConfirm,
		Text: "确认备用恢复码",
		Type: Info,
	}
}

func NewInfoSelfServiceSettingsLookupSecretList(secrets []string, raw any) *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsLookupSecretList,
		Text: strings.Join(secrets, ", "),
		Type: Info,
		Context: context(map[string]any{
			"secrets": raw,
		}),
	}
}
func NewInfoSelfServiceSettingsLookupSecret(secret string) *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsLookupSecret,
		Text: secret,
		Type: Info,
		Context: context(map[string]any{
			"secret": secret,
		}),
	}
}

func NewInfoSelfServiceSettingsLookupSecretUsed(usedAt time.Time) *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsLookupSecretUsed,
		Text: fmt.Sprintf("密钥已于 %s 使用", usedAt),
		Type: Info,
		Context: context(map[string]any{
			"used_at":      usedAt,
			"used_at_unix": usedAt.Unix(),
		}),
	}
}

func NewInfoSelfServiceSettingsLookupSecretsLabel() *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsLookupSecretLabel,
		Text: "这是您的备用恢复码，请妥善保管！",
		Type: Info,
	}
}

func NewInfoSelfServiceSettingsUpdateLinkOIDC(provider string) *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsUpdateLinkOidc,
		Text: fmt.Sprintf("关联 %s", provider),
		Type: Info,
		Context: context(map[string]any{
			"provider": provider,
		}),
	}
}

func NewInfoSelfServiceSettingsUpdateUnlinkOIDC(provider string) *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsUpdateUnlinkOidc,
		Text: fmt.Sprintf("取消关联 %s", provider),
		Type: Info,
		Context: context(map[string]any{
			"provider": provider,
		}),
	}
}

func NewInfoSelfServiceSettingsRegisterWebAuthn() *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsRegisterWebAuthn,
		Text: "添加安全密钥",
		Type: Info,
	}
}

func NewInfoSelfServiceSettingsRegisterPasskey() *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsRegisterPasskey,
		Text: "添加 Passkey",
		Type: Info,
	}
}

func NewInfoSelfServiceRegisterWebAuthnDisplayName() *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsRegisterWebAuthnDisplayName,
		Text: "安全密钥名称",
		Type: Info,
	}
}

func NewInfoSelfServiceRemoveWebAuthn(name string, createdAt time.Time) *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsRemoveWebAuthn,
		Text: fmt.Sprintf("移除安全密钥 \"%s\"", name),
		Type: Info,
		Context: context(map[string]any{
			"display_name":  name,
			"added_at":      createdAt,
			"added_at_unix": createdAt.Unix(),
		}),
	}
}

func NewInfoSelfServiceRemovePasskey(name string, createdAt time.Time) *Message {
	return &Message{
		ID:   InfoSelfServiceSettingsRemovePasskey,
		Text: fmt.Sprintf("移除 Passkey \"%s\"", name),
		Type: Info,
		Context: context(map[string]any{
			"display_name":  name,
			"added_at":      createdAt,
			"added_at_unix": createdAt.Unix(),
		}),
	}
}
