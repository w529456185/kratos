// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package text

func NewInfoNodeLabelVerifyOTP() *Message {
	return &Message{
		ID:   InfoNodeLabelVerifyOTP,
		Text: "验证码",
		Type: Info,
	}
}

func NewInfoNodeLabelVerificationCode() *Message {
	return &Message{
		ID:   InfoNodeLabelVerificationCode,
		Text: "验证码",
		Type: Info,
	}
}

func NewInfoNodeLabelRecoveryCode() *Message {
	return &Message{
		ID:   InfoNodeLabelRecoveryCode,
		Text: "恢复码",
		Type: Info,
	}
}

func NewInfoNodeLabelRegistrationCode() *Message {
	return &Message{
		ID:   InfoNodeLabelRegistrationCode,
		Text: "注册码",
		Type: Info,
	}
}

func NewInfoNodeLabelLoginCode() *Message {
	return &Message{
		ID:   InfoNodeLabelLoginCode,
		Text: "登录码",
		Type: Info,
	}
}

func NewInfoNodeInputPassword() *Message {
	return &Message{
		ID:   InfoNodeLabelInputPassword,
		Text: "密码",
		Type: Info,
	}
}

func NewInfoNodeLabelGenerated(title string, name string) *Message {
	return &Message{
		ID:   InfoNodeLabelGenerated,
		Text: title,
		Type: Info,
		Context: context(map[string]any{
			"title": title,
			"name":  name,
		}),
	}
}

func NewInfoNodeLabelSave() *Message {
	return &Message{
		ID:   InfoNodeLabelSave,
		Text: "保存",
		Type: Info,
	}
}

func NewInfoNodeLabelSubmit() *Message {
	return &Message{
		ID:   InfoNodeLabelSubmit,
		Text: "提交",
		Type: Info,
	}
}

func NewInfoNodeLabelContinue() *Message {
	return &Message{
		ID:   InfoNodeLabelContinue,
		Text: "继续",
		Type: Info,
	}
}

func NewInfoNodeLabelID() *Message {
	return &Message{
		ID:   InfoNodeLabelID,
		Text: "ID",
		Type: Info,
	}
}

func NewInfoNodeInputEmail() *Message {
	return &Message{
		ID:   InfoNodeLabelEmail,
		Text: "邮箱",
		Type: Info,
	}
}

func NewInfoNodeInputPhoneNumber() *Message {
	return &Message{
		ID:   InfoNodeLabelPhoneNumber,
		Text: "手机号",
		Type: Info,
	}
}

func NewInfoNodeResendOTP() *Message {
	return &Message{
		ID:   InfoNodeLabelResendOTP,
		Text: "重新发送验证码",
		Type: Info,
	}
}

func NewInfoNodeLoginAndLinkCredential() *Message {
	return &Message{
		ID:   InfoNodeLabelLoginAndLinkCredential,
		Text: "登录并关联凭证",
		Type: Info,
	}
}
