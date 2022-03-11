package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataAuthorizedUsernameAddPointer struct {
	message.VLSPointer
}

type TradeAccountDataAuthorizedUsernameAddPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeAccountDataAuthorizedUsernameAdd() TradeAccountDataAuthorizedUsernameAddPointerBuilder {
	a := TradeAccountDataAuthorizedUsernameAddPointerBuilder{message.AllocVLSBuilder(18)}
	a.Ptr.SetBytes(0, _TradeAccountDataAuthorizedUsernameAddDefault)
	return a
}

func AllocTradeAccountDataAuthorizedUsernameAddFrom(b []byte) TradeAccountDataAuthorizedUsernameAddPointer {
	return TradeAccountDataAuthorizedUsernameAddPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m TradeAccountDataAuthorizedUsernameAddPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataAuthorizedUsernameAddDefault)
}

// ToBuilder
func (m TradeAccountDataAuthorizedUsernameAddPointer) ToBuilder() TradeAccountDataAuthorizedUsernameAddPointerBuilder {
	return TradeAccountDataAuthorizedUsernameAddPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataAuthorizedUsernameAddPointerBuilder) Finish() TradeAccountDataAuthorizedUsernameAddPointer {
	return TradeAccountDataAuthorizedUsernameAddPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameAddPointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameAddPointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameAddPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameAddPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m TradeAccountDataAuthorizedUsernameAddPointer) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Username
func (m TradeAccountDataAuthorizedUsernameAddPointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// SetRequestID
func (m TradeAccountDataAuthorizedUsernameAddPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataAuthorizedUsernameAddPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataAuthorizedUsernameAddPointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// Copy
func (m TradeAccountDataAuthorizedUsernameAddPointer) Copy(to TradeAccountDataAuthorizedUsernameAddBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// Copy
func (m TradeAccountDataAuthorizedUsernameAddPointerBuilder) Copy(to TradeAccountDataAuthorizedUsernameAddBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameAddPointer) CopyPointer(to TradeAccountDataAuthorizedUsernameAddPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameAddPointerBuilder) CopyPointer(to TradeAccountDataAuthorizedUsernameAddPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}
