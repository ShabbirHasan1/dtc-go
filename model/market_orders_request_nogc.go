package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketOrdersRequestPointer struct {
	message.VLSPointer
}

type MarketOrdersRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocMarketOrdersRequest() MarketOrdersRequestPointerBuilder {
	a := MarketOrdersRequestPointerBuilder{message.AllocVLSBuilder(28)}
	a.Ptr.SetBytes(0, _MarketOrdersRequestDefault)
	return a
}

func AllocMarketOrdersRequestFrom(b []byte) MarketOrdersRequestPointer {
	return MarketOrdersRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m MarketOrdersRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersRequestDefault)
}

// ToBuilder
func (m MarketOrdersRequestPointer) ToBuilder() MarketOrdersRequestPointerBuilder {
	return MarketOrdersRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *MarketOrdersRequestPointerBuilder) Finish() MarketOrdersRequestPointer {
	return MarketOrdersRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestAction
func (m MarketOrdersRequestPointer) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Ptr, 12, 8))
}

// RequestAction
func (m MarketOrdersRequestPointerBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Ptr, 12, 8))
}

// SymbolID
func (m MarketOrdersRequestPointer) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 16, 12)
}

// SymbolID
func (m MarketOrdersRequestPointerBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 16, 12)
}

// Symbol
func (m MarketOrdersRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Symbol
func (m MarketOrdersRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Exchange
func (m MarketOrdersRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Exchange
func (m MarketOrdersRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestPointer) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32VLS(m.Ptr, 28, 24)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestPointerBuilder) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32VLS(m.Ptr, 28, 24)
}

// SetRequestAction
func (m MarketOrdersRequestPointerBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Ptr, 12, 8, int32(value))
}

// SetSymbolID
func (m MarketOrdersRequestPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Ptr, 16, 12, value)
}

// SetSymbol
func (m *MarketOrdersRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetExchange
func (m *MarketOrdersRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetSendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestPointerBuilder) SetSendQuantitiesGreaterOrEqualTo(value uint32) {
	message.SetUint32VLS(m.Ptr, 28, 24, value)
}

// Copy
func (m MarketOrdersRequestPointer) Copy(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// Copy
func (m MarketOrdersRequestPointerBuilder) Copy(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestPointer) CopyPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestPointerBuilder) CopyPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestPointer) CopyTo(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestPointerBuilder) CopyTo(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestPointer) CopyToPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestPointerBuilder) CopyToPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}
