package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataUsernameToShareWithRemovePointer struct {
	message.VLSPointer
}

type TradeAccountDataUsernameToShareWithRemovePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeAccountDataUsernameToShareWithRemove() TradeAccountDataUsernameToShareWithRemovePointerBuilder {
	a := TradeAccountDataUsernameToShareWithRemovePointerBuilder{message.AllocVLSBuilder(19)}
	a.Ptr.SetBytes(0, _TradeAccountDataUsernameToShareWithRemoveDefault)
	return a
}

func AllocTradeAccountDataUsernameToShareWithRemoveFrom(b []byte) TradeAccountDataUsernameToShareWithRemovePointer {
	return TradeAccountDataUsernameToShareWithRemovePointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataUsernameToShareWithRemoveDefault)
}

// ToBuilder
func (m TradeAccountDataUsernameToShareWithRemovePointer) ToBuilder() TradeAccountDataUsernameToShareWithRemovePointerBuilder {
	return TradeAccountDataUsernameToShareWithRemovePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataUsernameToShareWithRemovePointerBuilder) Finish() TradeAccountDataUsernameToShareWithRemovePointer {
	return TradeAccountDataUsernameToShareWithRemovePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataUsernameToShareWithRemovePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithRemovePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m TradeAccountDataUsernameToShareWithRemovePointer) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Username
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithRemovePointer) IsReadOnly() uint8 {
	return message.Uint8VLS(m.Ptr, 19, 18)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) IsReadOnly() uint8 {
	return message.Uint8VLS(m.Ptr, 19, 18)
}

// SetRequestID
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataUsernameToShareWithRemovePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataUsernameToShareWithRemovePointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetIsReadOnly
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) SetIsReadOnly(value uint8) {
	message.SetUint8VLS(m.Ptr, 19, 18, value)
}

// Copy
func (m TradeAccountDataUsernameToShareWithRemovePointer) Copy(to TradeAccountDataUsernameToShareWithRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// Copy
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) Copy(to TradeAccountDataUsernameToShareWithRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithRemovePointer) CopyPointer(to TradeAccountDataUsernameToShareWithRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) CopyPointer(to TradeAccountDataUsernameToShareWithRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}
