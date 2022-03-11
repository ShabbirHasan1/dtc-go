package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                           = MarketOrdersRequestSize  (28)
//     Type                           = MARKET_ORDERS_REQUEST  (150)
//     BaseSize                       = MarketOrdersRequestSize  (28)
//     RequestAction                  = 0
//     SymbolID                       = 0
//     Symbol                         = ""
//     Exchange                       = ""
//     SendQuantitiesGreaterOrEqualTo = 0
// }
var _MarketOrdersRequestDefault = []byte{28, 0, 150, 0, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketOrdersRequestSize = 28

type MarketOrdersRequest struct {
	message.VLS
}

type MarketOrdersRequestBuilder struct {
	message.VLSBuilder
}

func NewMarketOrdersRequestFrom(b []byte) MarketOrdersRequest {
	return MarketOrdersRequest{VLS: message.NewVLSFrom(b)}
}

func WrapMarketOrdersRequest(b []byte) MarketOrdersRequest {
	return MarketOrdersRequest{VLS: message.WrapVLS(b)}
}

func NewMarketOrdersRequest() MarketOrdersRequestBuilder {
	a := MarketOrdersRequestBuilder{message.NewVLSBuilder(28)}
	a.Unsafe().SetBytes(0, _MarketOrdersRequestDefault)
	return a
}

// Clear
// {
//     Size                           = MarketOrdersRequestSize  (28)
//     Type                           = MARKET_ORDERS_REQUEST  (150)
//     BaseSize                       = MarketOrdersRequestSize  (28)
//     RequestAction                  = 0
//     SymbolID                       = 0
//     Symbol                         = ""
//     Exchange                       = ""
//     SendQuantitiesGreaterOrEqualTo = 0
// }
func (m MarketOrdersRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersRequestDefault)
}

// ToBuilder
func (m MarketOrdersRequest) ToBuilder() MarketOrdersRequestBuilder {
	return MarketOrdersRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m MarketOrdersRequestBuilder) Finish() MarketOrdersRequest {
	return MarketOrdersRequest{m.VLSBuilder.Finish()}
}

// RequestAction
func (m MarketOrdersRequest) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Unsafe(), 12, 8))
}

// RequestAction
func (m MarketOrdersRequestBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Unsafe(), 12, 8))
}

// SymbolID
func (m MarketOrdersRequest) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 16, 12)
}

// SymbolID
func (m MarketOrdersRequestBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 16, 12)
}

// Symbol
func (m MarketOrdersRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Symbol
func (m MarketOrdersRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange
func (m MarketOrdersRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Exchange
func (m MarketOrdersRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequest) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32VLS(m.Unsafe(), 28, 24)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestBuilder) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32VLS(m.Unsafe(), 28, 24)
}

// SetRequestAction
func (m MarketOrdersRequestBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, int32(value))
}

// SetSymbolID
func (m MarketOrdersRequestBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 16, 12, value)
}

// SetSymbol
func (m *MarketOrdersRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetExchange
func (m *MarketOrdersRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetSendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestBuilder) SetSendQuantitiesGreaterOrEqualTo(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 28, 24, value)
}

// Copy
func (m MarketOrdersRequest) Copy(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// Copy
func (m MarketOrdersRequestBuilder) Copy(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequest) CopyPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestBuilder) CopyPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequest) CopyToPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestBuilder) CopyToPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequest) CopyTo(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestBuilder) CopyTo(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}
