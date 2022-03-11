package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketOrdersRequestFixedPointer struct {
	message.FixedPointer
}

type MarketOrdersRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketOrdersRequestFixed() MarketOrdersRequestFixedPointerBuilder {
	a := MarketOrdersRequestFixedPointerBuilder{message.AllocFixed(96)}
	a.Ptr.SetBytes(0, _MarketOrdersRequestFixedDefault)
	return a
}

func AllocMarketOrdersRequestFixedFrom(b []byte) MarketOrdersRequestFixedPointer {
	return MarketOrdersRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                           = MarketOrdersRequestFixedSize  (96)
//     Type                           = MARKET_ORDERS_REQUEST  (150)
//     RequestAction                  = SUBSCRIBE  (1)
//     SymbolID                       = 0
//     Symbol                         = ""
//     Exchange                       = ""
//     SendQuantitiesGreaterOrEqualTo = 0
// }
func (m MarketOrdersRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersRequestFixedDefault)
}

// ToBuilder
func (m MarketOrdersRequestFixedPointer) ToBuilder() MarketOrdersRequestFixedPointerBuilder {
	return MarketOrdersRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketOrdersRequestFixedPointerBuilder) Finish() MarketOrdersRequestFixedPointer {
	return MarketOrdersRequestFixedPointer{m.FixedPointer}
}

// RequestAction
func (m MarketOrdersRequestFixedPointer) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Ptr, 8, 4))
}

// RequestAction
func (m MarketOrdersRequestFixedPointerBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Ptr, 8, 4))
}

// SymbolID
func (m MarketOrdersRequestFixedPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// SymbolID
func (m MarketOrdersRequestFixedPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// Symbol
func (m MarketOrdersRequestFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 76, 12)
}

// Symbol
func (m MarketOrdersRequestFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 76, 12)
}

// Exchange
func (m MarketOrdersRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 92, 76)
}

// Exchange
func (m MarketOrdersRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 92, 76)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixedPointer) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32Fixed(m.Ptr, 96, 92)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixedPointerBuilder) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32Fixed(m.Ptr, 96, 92)
}

// SetRequestAction
func (m MarketOrdersRequestFixedPointerBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32Fixed(m.Ptr, 8, 4, int32(value))
}

// SetSymbolID
func (m MarketOrdersRequestFixedPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 12, 8, value)
}

// SetSymbol
func (m MarketOrdersRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 76, 12, value)
}

// SetExchange
func (m MarketOrdersRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 92, 76, value)
}

// SetSendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixedPointerBuilder) SetSendQuantitiesGreaterOrEqualTo(value uint32) {
	message.SetUint32Fixed(m.Ptr, 96, 92, value)
}

// Copy
func (m MarketOrdersRequestFixedPointer) Copy(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// Copy
func (m MarketOrdersRequestFixedPointerBuilder) Copy(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestFixedPointer) CopyPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestFixedPointerBuilder) CopyPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestFixedPointer) CopyTo(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestFixedPointerBuilder) CopyTo(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestFixedPointer) CopyToPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestFixedPointerBuilder) CopyToPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}
