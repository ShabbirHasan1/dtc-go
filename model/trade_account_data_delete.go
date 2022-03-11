package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = TradeAccountDataDeleteSize  (14)
//     Type         = TRADE_ACCOUNT_DATA_DELETE  (10118)
//     BaseSize     = TradeAccountDataDeleteSize  (14)
//     RequestID    = 0
//     TradeAccount = ""
// }
var _TradeAccountDataDeleteDefault = []byte{14, 0, 134, 39, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataDeleteSize = 14

type TradeAccountDataDelete struct {
	message.VLS
}

type TradeAccountDataDeleteBuilder struct {
	message.VLSBuilder
}

func NewTradeAccountDataDeleteFrom(b []byte) TradeAccountDataDelete {
	return TradeAccountDataDelete{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataDelete(b []byte) TradeAccountDataDelete {
	return TradeAccountDataDelete{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataDelete() TradeAccountDataDeleteBuilder {
	a := TradeAccountDataDeleteBuilder{message.NewVLSBuilder(14)}
	a.Unsafe().SetBytes(0, _TradeAccountDataDeleteDefault)
	return a
}

// Clear
// {
//     Size         = TradeAccountDataDeleteSize  (14)
//     Type         = TRADE_ACCOUNT_DATA_DELETE  (10118)
//     BaseSize     = TradeAccountDataDeleteSize  (14)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m TradeAccountDataDeleteBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataDeleteDefault)
}

// ToBuilder
func (m TradeAccountDataDelete) ToBuilder() TradeAccountDataDeleteBuilder {
	return TradeAccountDataDeleteBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataDeleteBuilder) Finish() TradeAccountDataDelete {
	return TradeAccountDataDelete{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataDelete) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataDeleteBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m TradeAccountDataDelete) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataDeleteBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// SetRequestID
func (m TradeAccountDataDeleteBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataDeleteBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// Copy
func (m TradeAccountDataDelete) Copy(to TradeAccountDataDeleteBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m TradeAccountDataDeleteBuilder) Copy(to TradeAccountDataDeleteBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataDelete) CopyPointer(to TradeAccountDataDeletePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataDeleteBuilder) CopyPointer(to TradeAccountDataDeletePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}
