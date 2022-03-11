package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataUsernameToShareWithAddPointer struct {
	message.VLSPointer
}

type TradeAccountDataUsernameToShareWithAddPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeAccountDataUsernameToShareWithAdd() TradeAccountDataUsernameToShareWithAddPointerBuilder {
	a := TradeAccountDataUsernameToShareWithAddPointerBuilder{message.AllocVLSBuilder(19)}
	a.Ptr.SetBytes(0, _TradeAccountDataUsernameToShareWithAddDefault)
	return a
}

func AllocTradeAccountDataUsernameToShareWithAddFrom(b []byte) TradeAccountDataUsernameToShareWithAddPointer {
	return TradeAccountDataUsernameToShareWithAddPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = TradeAccountDataUsernameToShareWithAddSize  (19)
//     Type         = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD  (10128)
//     BaseSize     = TradeAccountDataUsernameToShareWithAddSize  (19)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
//     IsReadOnly   = 0
// }
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataUsernameToShareWithAddDefault)
}

// ToBuilder
func (m TradeAccountDataUsernameToShareWithAddPointer) ToBuilder() TradeAccountDataUsernameToShareWithAddPointerBuilder {
	return TradeAccountDataUsernameToShareWithAddPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataUsernameToShareWithAddPointerBuilder) Finish() TradeAccountDataUsernameToShareWithAddPointer {
	return TradeAccountDataUsernameToShareWithAddPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataUsernameToShareWithAddPointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithAddPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m TradeAccountDataUsernameToShareWithAddPointer) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Username
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithAddPointer) IsReadOnly() uint8 {
	return message.Uint8VLS(m.Ptr, 19, 18)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) IsReadOnly() uint8 {
	return message.Uint8VLS(m.Ptr, 19, 18)
}

// SetRequestID
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataUsernameToShareWithAddPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataUsernameToShareWithAddPointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetIsReadOnly
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) SetIsReadOnly(value uint8) {
	message.SetUint8VLS(m.Ptr, 19, 18, value)
}

// Copy
func (m TradeAccountDataUsernameToShareWithAddPointer) Copy(to TradeAccountDataUsernameToShareWithAddBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// Copy
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) Copy(to TradeAccountDataUsernameToShareWithAddBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithAddPointer) CopyPointer(to TradeAccountDataUsernameToShareWithAddPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithAddPointerBuilder) CopyPointer(to TradeAccountDataUsernameToShareWithAddPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}
