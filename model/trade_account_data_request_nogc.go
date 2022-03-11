package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataRequestPointer struct {
	message.VLSPointer
}

type TradeAccountDataRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeAccountDataRequest() TradeAccountDataRequestPointerBuilder {
	a := TradeAccountDataRequestPointerBuilder{message.AllocVLSBuilder(14)}
	a.Ptr.SetBytes(0, _TradeAccountDataRequestDefault)
	return a
}

func AllocTradeAccountDataRequestFrom(b []byte) TradeAccountDataRequestPointer {
	return TradeAccountDataRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = TradeAccountDataRequestSize  (14)
//     Type         = TRADE_ACCOUNT_DATA_REQUEST  (10115)
//     BaseSize     = TradeAccountDataRequestSize  (14)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m TradeAccountDataRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataRequestDefault)
}

// ToBuilder
func (m TradeAccountDataRequestPointer) ToBuilder() TradeAccountDataRequestPointerBuilder {
	return TradeAccountDataRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataRequestPointerBuilder) Finish() TradeAccountDataRequestPointer {
	return TradeAccountDataRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataRequestPointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataRequestPointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// SetRequestID
func (m TradeAccountDataRequestPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// Copy
func (m TradeAccountDataRequestPointer) Copy(to TradeAccountDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m TradeAccountDataRequestPointerBuilder) Copy(to TradeAccountDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataRequestPointer) CopyPointer(to TradeAccountDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataRequestPointerBuilder) CopyPointer(to TradeAccountDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}
