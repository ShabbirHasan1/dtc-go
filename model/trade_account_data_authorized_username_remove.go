package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = TradeAccountDataAuthorizedUsernameRemoveSize  (18)
//     Type         = TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE  (10126)
//     BaseSize     = TradeAccountDataAuthorizedUsernameRemoveSize  (18)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
// }
var _TradeAccountDataAuthorizedUsernameRemoveDefault = []byte{18, 0, 142, 39, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataAuthorizedUsernameRemoveSize = 18

type TradeAccountDataAuthorizedUsernameRemove struct {
	message.VLS
}

type TradeAccountDataAuthorizedUsernameRemoveBuilder struct {
	message.VLSBuilder
}

func NewTradeAccountDataAuthorizedUsernameRemoveFrom(b []byte) TradeAccountDataAuthorizedUsernameRemove {
	return TradeAccountDataAuthorizedUsernameRemove{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataAuthorizedUsernameRemove(b []byte) TradeAccountDataAuthorizedUsernameRemove {
	return TradeAccountDataAuthorizedUsernameRemove{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataAuthorizedUsernameRemove() TradeAccountDataAuthorizedUsernameRemoveBuilder {
	a := TradeAccountDataAuthorizedUsernameRemoveBuilder{message.NewVLSBuilder(18)}
	a.Unsafe().SetBytes(0, _TradeAccountDataAuthorizedUsernameRemoveDefault)
	return a
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
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataAuthorizedUsernameRemoveDefault)
}

// ToBuilder
func (m TradeAccountDataAuthorizedUsernameRemove) ToBuilder() TradeAccountDataAuthorizedUsernameRemoveBuilder {
	return TradeAccountDataAuthorizedUsernameRemoveBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) Finish() TradeAccountDataAuthorizedUsernameRemove {
	return TradeAccountDataAuthorizedUsernameRemove{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameRemove) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameRemove) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Username
func (m TradeAccountDataAuthorizedUsernameRemove) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Username
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// SetRequestID
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataAuthorizedUsernameRemoveBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataAuthorizedUsernameRemoveBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// Copy
func (m TradeAccountDataAuthorizedUsernameRemove) Copy(to TradeAccountDataAuthorizedUsernameRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// Copy
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) Copy(to TradeAccountDataAuthorizedUsernameRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameRemove) CopyPointer(to TradeAccountDataAuthorizedUsernameRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) CopyPointer(to TradeAccountDataAuthorizedUsernameRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}
