package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = TradeAccountDataAuthorizedUsernameAddSize  (18)
//     Type         = TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_ADD  (10125)
//     BaseSize     = TradeAccountDataAuthorizedUsernameAddSize  (18)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
// }
var _TradeAccountDataAuthorizedUsernameAddDefault = []byte{18, 0, 141, 39, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataAuthorizedUsernameAddSize = 18

type TradeAccountDataAuthorizedUsernameAdd struct {
	message.VLS
}

type TradeAccountDataAuthorizedUsernameAddBuilder struct {
	message.VLSBuilder
}

func NewTradeAccountDataAuthorizedUsernameAddFrom(b []byte) TradeAccountDataAuthorizedUsernameAdd {
	return TradeAccountDataAuthorizedUsernameAdd{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataAuthorizedUsernameAdd(b []byte) TradeAccountDataAuthorizedUsernameAdd {
	return TradeAccountDataAuthorizedUsernameAdd{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataAuthorizedUsernameAdd() TradeAccountDataAuthorizedUsernameAddBuilder {
	a := TradeAccountDataAuthorizedUsernameAddBuilder{message.NewVLSBuilder(18)}
	a.Unsafe().SetBytes(0, _TradeAccountDataAuthorizedUsernameAddDefault)
	return a
}

// Clear
// {
//     Size         = TradeAccountDataAuthorizedUsernameAddSize  (18)
//     Type         = TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_ADD  (10125)
//     BaseSize     = TradeAccountDataAuthorizedUsernameAddSize  (18)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
// }
func (m TradeAccountDataAuthorizedUsernameAddBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataAuthorizedUsernameAddDefault)
}

// ToBuilder
func (m TradeAccountDataAuthorizedUsernameAdd) ToBuilder() TradeAccountDataAuthorizedUsernameAddBuilder {
	return TradeAccountDataAuthorizedUsernameAddBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataAuthorizedUsernameAddBuilder) Finish() TradeAccountDataAuthorizedUsernameAdd {
	return TradeAccountDataAuthorizedUsernameAdd{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameAdd) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameAddBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameAdd) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameAddBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Username
func (m TradeAccountDataAuthorizedUsernameAdd) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Username
func (m TradeAccountDataAuthorizedUsernameAddBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// SetRequestID
func (m TradeAccountDataAuthorizedUsernameAddBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataAuthorizedUsernameAddBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataAuthorizedUsernameAddBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// Copy
func (m TradeAccountDataAuthorizedUsernameAdd) Copy(to TradeAccountDataAuthorizedUsernameAddBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// Copy
func (m TradeAccountDataAuthorizedUsernameAddBuilder) Copy(to TradeAccountDataAuthorizedUsernameAddBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameAdd) CopyPointer(to TradeAccountDataAuthorizedUsernameAddPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameAddBuilder) CopyPointer(to TradeAccountDataAuthorizedUsernameAddPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}
