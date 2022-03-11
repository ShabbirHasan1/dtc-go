package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataDeletePointer struct {
	message.VLSPointer
}

type TradeAccountDataDeletePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeAccountDataDelete() TradeAccountDataDeletePointerBuilder {
	a := TradeAccountDataDeletePointerBuilder{message.AllocVLSBuilder(14)}
	a.Ptr.SetBytes(0, _TradeAccountDataDeleteDefault)
	return a
}

func AllocTradeAccountDataDeleteFrom(b []byte) TradeAccountDataDeletePointer {
	return TradeAccountDataDeletePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = TradeAccountDataDeleteSize  (14)
//     Type         = TRADE_ACCOUNT_DATA_DELETE  (10118)
//     BaseSize     = TradeAccountDataDeleteSize  (14)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m TradeAccountDataDeletePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataDeleteDefault)
}

// ToBuilder
func (m TradeAccountDataDeletePointer) ToBuilder() TradeAccountDataDeletePointerBuilder {
	return TradeAccountDataDeletePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataDeletePointerBuilder) Finish() TradeAccountDataDeletePointer {
	return TradeAccountDataDeletePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataDeletePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataDeletePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataDeletePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataDeletePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// SetRequestID
func (m TradeAccountDataDeletePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataDeletePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// Copy
func (m TradeAccountDataDeletePointer) Copy(to TradeAccountDataDeleteBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m TradeAccountDataDeletePointerBuilder) Copy(to TradeAccountDataDeleteBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataDeletePointer) CopyPointer(to TradeAccountDataDeletePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataDeletePointerBuilder) CopyPointer(to TradeAccountDataDeletePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}
