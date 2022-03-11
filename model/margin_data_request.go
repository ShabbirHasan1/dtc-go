package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = MarginDataRequestSize  (18)
//     Type         = MARGIN_DATA_REQUEST  (10141)
//     BaseSize     = MarginDataRequestSize  (18)
//     RequestID    = 0
//     TradeAccount = ""
//     Symbol       = ""
// }
var _MarginDataRequestDefault = []byte{18, 0, 157, 39, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarginDataRequestSize = 18

type MarginDataRequest struct {
	message.VLS
}

type MarginDataRequestBuilder struct {
	message.VLSBuilder
}

func NewMarginDataRequestFrom(b []byte) MarginDataRequest {
	return MarginDataRequest{VLS: message.NewVLSFrom(b)}
}

func WrapMarginDataRequest(b []byte) MarginDataRequest {
	return MarginDataRequest{VLS: message.WrapVLS(b)}
}

func NewMarginDataRequest() MarginDataRequestBuilder {
	a := MarginDataRequestBuilder{message.NewVLSBuilder(18)}
	a.Unsafe().SetBytes(0, _MarginDataRequestDefault)
	return a
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
func (m MarginDataRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarginDataRequestDefault)
}

// ToBuilder
func (m MarginDataRequest) ToBuilder() MarginDataRequestBuilder {
	return MarginDataRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m MarginDataRequestBuilder) Finish() MarginDataRequest {
	return MarginDataRequest{m.VLSBuilder.Finish()}
}

// RequestID
func (m MarginDataRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m MarginDataRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m MarginDataRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m MarginDataRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Symbol
func (m MarginDataRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Symbol
func (m MarginDataRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// SetRequestID
func (m MarginDataRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 10, 6, value)
}

// SetTradeAccount
func (m *MarginDataRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetSymbol
func (m *MarginDataRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// Copy
func (m MarginDataRequest) Copy(to MarginDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
}

// Copy
func (m MarginDataRequestBuilder) Copy(to MarginDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m MarginDataRequest) CopyPointer(to MarginDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m MarginDataRequestBuilder) CopyPointer(to MarginDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
}
