package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = TradeAccountDataUsernameToShareWithAddSize  (19)
//     Type         = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD  (10128)
//     BaseSize     = TradeAccountDataUsernameToShareWithAddSize  (19)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
//     IsReadOnly   = 0
// }
var _TradeAccountDataUsernameToShareWithAddDefault = []byte{19, 0, 144, 39, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataUsernameToShareWithAddSize = 19

type TradeAccountDataUsernameToShareWithAdd struct {
	message.VLS
}

type TradeAccountDataUsernameToShareWithAddBuilder struct {
	message.VLSBuilder
}

func NewTradeAccountDataUsernameToShareWithAddFrom(b []byte) TradeAccountDataUsernameToShareWithAdd {
	return TradeAccountDataUsernameToShareWithAdd{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataUsernameToShareWithAdd(b []byte) TradeAccountDataUsernameToShareWithAdd {
	return TradeAccountDataUsernameToShareWithAdd{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataUsernameToShareWithAdd() TradeAccountDataUsernameToShareWithAddBuilder {
	a := TradeAccountDataUsernameToShareWithAddBuilder{message.NewVLSBuilder(19)}
	a.Unsafe().SetBytes(0, _TradeAccountDataUsernameToShareWithAddDefault)
	return a
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
func (m TradeAccountDataUsernameToShareWithAddBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataUsernameToShareWithAddDefault)
}

// ToBuilder
func (m TradeAccountDataUsernameToShareWithAdd) ToBuilder() TradeAccountDataUsernameToShareWithAddBuilder {
	return TradeAccountDataUsernameToShareWithAddBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataUsernameToShareWithAddBuilder) Finish() TradeAccountDataUsernameToShareWithAdd {
	return TradeAccountDataUsernameToShareWithAdd{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataUsernameToShareWithAdd) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithAddBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithAdd) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithAddBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Username
func (m TradeAccountDataUsernameToShareWithAdd) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Username
func (m TradeAccountDataUsernameToShareWithAddBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithAdd) IsReadOnly() uint8 {
	return message.Uint8VLS(m.Unsafe(), 19, 18)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithAddBuilder) IsReadOnly() uint8 {
	return message.Uint8VLS(m.Unsafe(), 19, 18)
}

// SetRequestID
func (m TradeAccountDataUsernameToShareWithAddBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataUsernameToShareWithAddBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataUsernameToShareWithAddBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetIsReadOnly
func (m TradeAccountDataUsernameToShareWithAddBuilder) SetIsReadOnly(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 19, 18, value)
}

// Copy
func (m TradeAccountDataUsernameToShareWithAdd) Copy(to TradeAccountDataUsernameToShareWithAddBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// Copy
func (m TradeAccountDataUsernameToShareWithAddBuilder) Copy(to TradeAccountDataUsernameToShareWithAddBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithAdd) CopyPointer(to TradeAccountDataUsernameToShareWithAddPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithAddBuilder) CopyPointer(to TradeAccountDataUsernameToShareWithAddPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}
