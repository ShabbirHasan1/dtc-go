package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = TradeAccountDataUsernameToShareWithRemoveSize  (19)
//     Type         = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_REMOVE  (10129)
//     BaseSize     = TradeAccountDataUsernameToShareWithRemoveSize  (19)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
//     IsReadOnly   = 0
// }
var _TradeAccountDataUsernameToShareWithRemoveDefault = []byte{19, 0, 145, 39, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataUsernameToShareWithRemoveSize = 19

type TradeAccountDataUsernameToShareWithRemove struct {
	message.VLS
}

type TradeAccountDataUsernameToShareWithRemoveBuilder struct {
	message.VLSBuilder
}

func NewTradeAccountDataUsernameToShareWithRemoveFrom(b []byte) TradeAccountDataUsernameToShareWithRemove {
	return TradeAccountDataUsernameToShareWithRemove{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataUsernameToShareWithRemove(b []byte) TradeAccountDataUsernameToShareWithRemove {
	return TradeAccountDataUsernameToShareWithRemove{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataUsernameToShareWithRemove() TradeAccountDataUsernameToShareWithRemoveBuilder {
	a := TradeAccountDataUsernameToShareWithRemoveBuilder{message.NewVLSBuilder(19)}
	a.Unsafe().SetBytes(0, _TradeAccountDataUsernameToShareWithRemoveDefault)
	return a
}

// Clear
// {
//     Size         = TradeAccountDataUsernameToShareWithRemoveSize  (19)
//     Type         = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_REMOVE  (10129)
//     BaseSize     = TradeAccountDataUsernameToShareWithRemoveSize  (19)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
//     IsReadOnly   = 0
// }
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataUsernameToShareWithRemoveDefault)
}

// ToBuilder
func (m TradeAccountDataUsernameToShareWithRemove) ToBuilder() TradeAccountDataUsernameToShareWithRemoveBuilder {
	return TradeAccountDataUsernameToShareWithRemoveBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) Finish() TradeAccountDataUsernameToShareWithRemove {
	return TradeAccountDataUsernameToShareWithRemove{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataUsernameToShareWithRemove) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithRemove) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Username
func (m TradeAccountDataUsernameToShareWithRemove) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Username
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithRemove) IsReadOnly() uint8 {
	return message.Uint8VLS(m.Unsafe(), 19, 18)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) IsReadOnly() uint8 {
	return message.Uint8VLS(m.Unsafe(), 19, 18)
}

// SetRequestID
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataUsernameToShareWithRemoveBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataUsernameToShareWithRemoveBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetIsReadOnly
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) SetIsReadOnly(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 19, 18, value)
}

// Copy
func (m TradeAccountDataUsernameToShareWithRemove) Copy(to TradeAccountDataUsernameToShareWithRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// Copy
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) Copy(to TradeAccountDataUsernameToShareWithRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithRemove) CopyPointer(to TradeAccountDataUsernameToShareWithRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) CopyPointer(to TradeAccountDataUsernameToShareWithRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}
