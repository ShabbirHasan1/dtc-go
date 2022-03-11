package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarginDataRequestPointer struct {
	message.VLSPointer
}

type MarginDataRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocMarginDataRequest() MarginDataRequestPointerBuilder {
	a := MarginDataRequestPointerBuilder{message.AllocVLSBuilder(18)}
	a.Ptr.SetBytes(0, _MarginDataRequestDefault)
	return a
}

func AllocMarginDataRequestFrom(b []byte) MarginDataRequestPointer {
	return MarginDataRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = MarginDataRequestSize  (18)
//     Type         = MARGIN_DATA_REQUEST  (10141)
//     BaseSize     = MarginDataRequestSize  (18)
//     RequestID    = 0
//     TradeAccount = ""
//     Symbol       = ""
// }
func (m MarginDataRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarginDataRequestDefault)
}

// ToBuilder
func (m MarginDataRequestPointer) ToBuilder() MarginDataRequestPointerBuilder {
	return MarginDataRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *MarginDataRequestPointerBuilder) Finish() MarginDataRequestPointer {
	return MarginDataRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m MarginDataRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m MarginDataRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m MarginDataRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m MarginDataRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Symbol
func (m MarginDataRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Symbol
func (m MarginDataRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// SetRequestID
func (m MarginDataRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *MarginDataRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetSymbol
func (m *MarginDataRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// Copy
func (m MarginDataRequestPointer) Copy(to MarginDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
}

// Copy
func (m MarginDataRequestPointerBuilder) Copy(to MarginDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m MarginDataRequestPointer) CopyPointer(to MarginDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m MarginDataRequestPointerBuilder) CopyPointer(to MarginDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
}
