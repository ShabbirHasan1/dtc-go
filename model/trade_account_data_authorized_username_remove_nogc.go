package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataAuthorizedUsernameRemovePointer struct {
	message.VLSPointer
}

type TradeAccountDataAuthorizedUsernameRemovePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeAccountDataAuthorizedUsernameRemove() TradeAccountDataAuthorizedUsernameRemovePointerBuilder {
	a := TradeAccountDataAuthorizedUsernameRemovePointerBuilder{message.AllocVLSBuilder(18)}
	a.Ptr.SetBytes(0, _TradeAccountDataAuthorizedUsernameRemoveDefault)
	return a
}

func AllocTradeAccountDataAuthorizedUsernameRemoveFrom(b []byte) TradeAccountDataAuthorizedUsernameRemovePointer {
	return TradeAccountDataAuthorizedUsernameRemovePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = TradeAccountDataAuthorizedUsernameRemoveSize  (18)
//     Type         = TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE  (10126)
//     BaseSize     = TradeAccountDataAuthorizedUsernameRemoveSize  (18)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
// }
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataAuthorizedUsernameRemoveDefault)
}

// ToBuilder
func (m TradeAccountDataAuthorizedUsernameRemovePointer) ToBuilder() TradeAccountDataAuthorizedUsernameRemovePointerBuilder {
	return TradeAccountDataAuthorizedUsernameRemovePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataAuthorizedUsernameRemovePointerBuilder) Finish() TradeAccountDataAuthorizedUsernameRemovePointer {
	return TradeAccountDataAuthorizedUsernameRemovePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameRemovePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameRemovePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m TradeAccountDataAuthorizedUsernameRemovePointer) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Username
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// SetRequestID
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataAuthorizedUsernameRemovePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataAuthorizedUsernameRemovePointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// Copy
func (m TradeAccountDataAuthorizedUsernameRemovePointer) Copy(to TradeAccountDataAuthorizedUsernameRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// Copy
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) Copy(to TradeAccountDataAuthorizedUsernameRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameRemovePointer) CopyPointer(to TradeAccountDataAuthorizedUsernameRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) CopyPointer(to TradeAccountDataAuthorizedUsernameRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}
