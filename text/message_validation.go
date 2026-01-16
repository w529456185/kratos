// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package text

import (
	"fmt"
	"strings"

	"github.com/ory/x/stringslice"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NewValidationErrorGeneric(reason string) *Message {
	return &Message{
		ID:   ErrorValidationGeneric,
		Text: reason,
		Type: Error,
		Context: context(map[string]any{
			"reason": reason,
		}),
	}
}

func NewValidationErrorRequired(missing string) *Message {
	return &Message{
		ID:   ErrorValidationRequired,
		Text: fmt.Sprintf("缺少属性 %s。", missing),
		Type: Error,
		Context: context(map[string]any{
			"property": missing,
		}),
	}
}

func NewErrorValidationMinLength(minLength, actualLength int) *Message {
	return &Message{
		ID:   ErrorValidationMinLength,
		Text: fmt.Sprintf("长度必须 >= %d，但实际为 %d。", minLength, actualLength),
		Type: Error,
		Context: context(map[string]any{
			"min_length":    minLength,
			"actual_length": actualLength,
		}),
	}
}

func NewErrorValidationMaxLength(maxLength, actualLength int) *Message {
	return &Message{
		ID:   ErrorValidationMaxLength,
		Text: fmt.Sprintf("长度必须 <= %d，但实际为 %d。", maxLength, actualLength),
		Type: Error,
		Context: context(map[string]any{
			"max_length":    maxLength,
			"actual_length": actualLength,
		}),
	}
}

func NewErrorValidationInvalidFormat(pattern string) *Message {
	return &Message{
		ID:   ErrorValidationInvalidFormat,
		Text: fmt.Sprintf("不符合模式 %q。", pattern),
		Type: Error,
		Context: context(map[string]any{
			"pattern": pattern,
		}),
	}
}

func NewErrorValidationMinimum(minimum, actual float64) *Message {
	return &Message{
		ID:   ErrorValidationMinimum,
		Text: fmt.Sprintf("必须 >= %v，但实际为 %v。", minimum, actual),
		Type: Error,
		Context: context(map[string]any{
			"minimum": minimum,
			"actual":  actual,
		}),
	}
}

func NewErrorValidationExclusiveMinimum(minimum, actual float64) *Message {
	return &Message{
		ID:   ErrorValidationExclusiveMinimum,
		Text: fmt.Sprintf("必须 > %v，但实际为 %v。", minimum, actual),
		Type: Error,
		Context: context(map[string]any{
			"minimum": minimum,
			"actual":  actual,
		}),
	}
}

func NewErrorValidationMaximum(maximum, actual float64) *Message {
	return &Message{
		ID:   ErrorValidationMaximum,
		Text: fmt.Sprintf("必须 <= %v，但实际为 %v。", maximum, actual),
		Type: Error,
		Context: context(map[string]any{
			"maximum": maximum,
			"actual":  actual,
		}),
	}
}

func NewErrorValidationExclusiveMaximum(maximum, actual float64) *Message {
	return &Message{
		ID:   ErrorValidationExclusiveMaximum,
		Text: fmt.Sprintf("必须 < %v，但实际为 %v。", maximum, actual),
		Type: Error,
		Context: context(map[string]any{
			"maximum": maximum,
			"actual":  actual,
		}),
	}
}

func NewErrorValidationMultipleOf(base, actual float64) *Message {
	return &Message{
		ID:   ErrorValidationMultipleOf,
		Text: fmt.Sprintf("%v 不是 %v 的倍数。", actual, base),
		Type: Error,
		Context: context(map[string]any{
			"base":   base,
			"actual": actual,
		}),
	}
}

func NewErrorValidationMaxItems(maxItems, actualItems int) *Message {
	return &Message{
		ID:   ErrorValidationMaxItems,
		Text: fmt.Sprintf("最多允许 %d 项，但实际为 %d 项。", maxItems, actualItems),
		Type: Error,
		Context: context(map[string]any{
			"max_items":    maxItems,
			"actual_items": actualItems,
		}),
	}
}

func NewErrorValidationMinItems(minItems, actualItems int) *Message {
	return &Message{
		ID:   ErrorValidationMinItems,
		Text: fmt.Sprintf("至少需要 %d 项，但实际为 %d 项。", minItems, actualItems),
		Type: Error,
		Context: context(map[string]any{
			"min_items":    minItems,
			"actual_items": actualItems,
		}),
	}
}

func NewErrorValidationUniqueItems(indexA, indexB int) *Message {
	return &Message{
		ID:   ErrorValidationUniqueItems,
		Text: fmt.Sprintf("索引 %d 和 %d 处的项相同。", indexA, indexB),
		Type: Error,
		Context: context(map[string]any{
			"index_a": indexA,
			"index_b": indexB,
		}),
	}
}

func NewErrorValidationWrongType(allowedTypes []string, actualType string) *Message {
	return &Message{
		ID:   ErrorValidationWrongType,
		Text: fmt.Sprintf("期望类型为 %s，但实际为 %s。", strings.Join(allowedTypes, " 或 "), actualType),
		Type: Error,
		Context: context(map[string]any{
			"allowed_types": allowedTypes,
			"actual_type":   actualType,
		}),
	}
}

func NewErrorValidationConst(expected any) *Message {
	return &Message{
		ID:   ErrorValidationConst,
		Text: fmt.Sprintf("必须等于常量 %v。", expected),
		Type: Error,
		Context: context(map[string]any{
			"expected": expected,
		}),
	}
}

func NewErrorValidationConstGeneric() *Message {
	return &Message{
		ID:   ErrorValidationConstGeneric,
		Text: "常量验证失败。",
		Type: Error,
	}
}

func NewErrorValidationPasswordPolicyViolationGeneric(reason string) *Message {
	return &Message{
		ID:   ErrorValidationPasswordPolicyViolationGeneric,
		Text: fmt.Sprintf("密码无法使用，因为 %s。", reason),
		Type: Error,
		Context: context(map[string]any{
			"reason": reason,
		}),
	}
}

func NewErrorValidationPasswordIdentifierTooSimilar() *Message {
	return &Message{
		ID:   ErrorValidationPasswordIdentifierTooSimilar,
		Text: "密码无法使用，因为它与标识符过于相似。",
		Type: Error,
	}
}

func NewErrorValidationPasswordMinLength(minLength, actualLength int) *Message {
	return &Message{
		ID:   ErrorValidationPasswordMinLength,
		Text: fmt.Sprintf("密码长度必须至少为 %d 个字符，但实际为 %d 个字符。", minLength, actualLength),
		Type: Error,
		Context: context(map[string]any{
			"min_length":    minLength,
			"actual_length": actualLength,
		}),
	}
}

func NewErrorValidationPasswordMaxLength(maxLength, actualLength int) *Message {
	return &Message{
		ID:   ErrorValidationPasswordMaxLength,
		Text: fmt.Sprintf("密码长度最多为 %d 个字符，但实际为 %d 个字符。", maxLength, actualLength),
		Type: Error,
		Context: context(map[string]any{
			"max_length":    maxLength,
			"actual_length": actualLength,
		}),
	}
}

func NewErrorValidationPasswordNewSameAsOld() *Message {
	return &Message{
		ID:   ErrorValidationPasswordNewSameAsOld,
		Text: "新密码必须与旧密码不同。",
		Type: Error,
	}
}

func NewErrorValidationPasswordTooManyBreaches(breaches int64) *Message {
	return &Message{
		ID:   ErrorValidationPasswordTooManyBreaches,
		Text: "密码已在数据泄露中被发现，不能再使用。",
		Type: Error,
		Context: context(map[string]any{
			"breaches": breaches,
		}),
	}
}

func NewErrorValidationInvalidCredentials() *Message {
	return &Message{
		ID:   ErrorValidationInvalidCredentials,
		Text: "提供的登录信息无效，请检查密码、邮箱地址或手机号是否正确。",
		Type: Error,
	}
}

func NewErrorValidationAccountNotFound() *Message {
	return &Message{
		ID:   ErrorValidationAccountNotFound,
		Text: "此账户不存在或未配置登录方法。",
		Type: Error,
	}
}

func NewErrorValidationDuplicateCredentials() *Message {
	return &Message{
		ID:   ErrorValidationDuplicateCredentials,
		Text: "已存在相同身份（邮箱、电话、用户名等）的账户。",
		Type: Error,
	}
}

func NewErrorValidationDuplicateCredentialsWithHints(availableCredentialTypes []string, availableOIDCProviders []string, credentialIdentifierHint string) *Message {
	identifier := credentialIdentifierHint
	if identifier == "" {
		identifier = "邮箱、电话或用户名"
	}
	oidcProviders := make([]string, 0, len(availableOIDCProviders))
	for _, provider := range availableOIDCProviders {
		oidcProviders = append(oidcProviders, cases.Title(language.English).String(provider))
	}

	reason := fmt.Sprintf("You tried signing in with %s which is already in use by another account.", identifier)
	if len(availableCredentialTypes) > 0 {
		humanReadable := make([]string, 0, len(availableCredentialTypes))
		for _, cred := range availableCredentialTypes {
			switch cred {
			case "password":
				humanReadable = append(humanReadable, "您的密码")
			case "oidc", "saml":
				humanReadable = append(humanReadable, "社交登录")
			case "webauthn":
				humanReadable = append(humanReadable, "您的 Passkey 或安全密钥")
			case "passkey":
				humanReadable = append(humanReadable, "您的 Passkey")
			}
		}
		if len(humanReadable) == 0 {
			// show at least some hint
			// also our example message generation tool runs into this case
			humanReadable = append(humanReadable, availableCredentialTypes...)
		}

		humanReadable = stringslice.Unique(humanReadable)

		// Final format: "You can sign in using foo, bar, or baz."
		if len(humanReadable) > 1 {
			humanReadable[len(humanReadable)-1] = "或 " + humanReadable[len(humanReadable)-1]
		}
		if len(humanReadable) > 0 {
			reason += fmt.Sprintf(" 您可以使用 %s 登录。", strings.Join(humanReadable, "、"))
		}
	}
	if len(oidcProviders) > 0 {
		reason += fmt.Sprintf(" 您可以使用以下社交登录提供商之一进行登录：%s。", strings.Join(oidcProviders, "、"))
	}

	return &Message{
		ID:   ErrorValidationDuplicateCredentialsWithHints,
		Text: reason,
		Type: Error,
		Context: context(map[string]any{
			"available_credential_types": availableCredentialTypes,
			"available_oidc_providers":   availableOIDCProviders,
			"credential_identifier_hint": credentialIdentifierHint,
		}),
	}
}

func NewErrorValidationDuplicateCredentialsOnOIDCLink() *Message {
	return &Message{
		ID: ErrorValidationDuplicateCredentialsOnOIDCLink,
		Text: "已存在相同身份（邮箱、电话、用户名等）的账户。" +
			"请登录现有账户以关联您的社交资料。",
		Type: Error,
	}
}

func NewErrorValidationTOTPVerifierWrong() *Message {
	return &Message{
		ID:   ErrorValidationTOTPVerifierWrong,
		Text: "提供的验证码无效，请重试。",
		Type: Error,
	}
}

func NewErrorValidationLookupAlreadyUsed() *Message {
	return &Message{
		ID:   ErrorValidationLookupAlreadyUsed,
		Text: "此备用恢复码已被使用。",
		Type: Error,
	}
}

func NewErrorValidationLookupInvalid() *Message {
	return &Message{
		ID:   ErrorValidationLookupInvalid,
		Text: "备用恢复码无效。",
		Type: Error,
	}
}

func NewErrorValidationIdentifierMissing() *Message {
	return &Message{
		ID:   ErrorValidationIdentifierMissing,
		Text: "未找到任何登录标识符。您是否忘记设置？也可能是服务器配置错误。",
		Type: Error,
	}
}

func NewErrorValidationAddressNotVerified() *Message {
	return &Message{
		ID:   ErrorValidationAddressNotVerified,
		Text: "账户尚未激活。您是否忘记验证邮箱地址？",
		Type: Error,
	}
}

func NewErrorValidationNoTOTPDevice() *Message {
	return &Message{
		ID:   ErrorValidationNoTOTPDevice,
		Text: "您未设置 TOTP 设备。",
		Type: Error,
	}
}

func NewErrorValidationNoLookup() *Message {
	return &Message{
		ID:   ErrorValidationNoLookup,
		Text: "您未设置备用恢复码。",
		Type: Error,
	}
}

func NewErrorValidationNoWebAuthnDevice() *Message {
	return &Message{
		ID:   ErrorValidationNoWebAuthnDevice,
		Text: "您未设置 WebAuthn 设备。",
		Type: Error,
	}
}

func NewErrorValidationSuchNoWebAuthnUser() *Message {
	return &Message{
		ID:   ErrorValidationSuchNoWebAuthnUser,
		Text: "此账户不存在或未设置安全密钥。",
		Type: Error,
	}
}

func NewErrorValidationNoCodeUser() *Message {
	return &Message{
		ID:   ErrorValidationNoCodeUser,
		Text: "此账户不存在或未设置验证码登录。",
		Type: Error,
	}
}

func NewErrorValidationTraitsMismatch() *Message {
	return &Message{
		ID:   ErrorValidationTraitsMismatch,
		Text: "提供的特征与此流程先前关联的特征不匹配。",
		Type: Error,
	}
}

func NewErrorCaptchaFailed() *Message {
	return &Message{
		ID:   ErrorValidationCaptchaError,
		Text: "验证码验证失败，请重试。",
		Type: Error,
	}
}
