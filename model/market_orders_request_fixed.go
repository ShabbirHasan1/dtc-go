package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                           = MarketOrdersRequestFixedSize  (96)
//     Type                           = MARKET_ORDERS_REQUEST  (150)
//     RequestAction                  = SUBSCRIBE  (1)
//     SymbolID                       = 0
//     Symbol                         = ""
//     Exchange                       = ""
//     SendQuantitiesGreaterOrEqualTo = 0
// }
var _MarketOrdersRequestFixedDefault = []byte{96, 0, 150, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketOrdersRequestFixedSize = 96

type MarketOrdersRequestFixed struct {
	message.Fixed
}

type MarketOrdersRequestFixedBuilder struct {
	message.Fixed
}

func NewMarketOrdersRequestFixedFrom(b []byte) MarketOrdersRequestFixed {
	return MarketOrdersRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketOrdersRequestFixed(b []byte) MarketOrdersRequestFixed {
	return MarketOrdersRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewMarketOrdersRequestFixed() MarketOrdersRequestFixedBuilder {
	a := MarketOrdersRequestFixedBuilder{message.NewFixed(96)}
	a.Unsafe().SetBytes(0, _MarketOrdersRequestFixedDefault)
	return a
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
func (m MarketOrdersRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersRequestFixedDefault)
}

// ToBuilder
func (m MarketOrdersRequestFixed) ToBuilder() MarketOrdersRequestFixedBuilder {
	return MarketOrdersRequestFixedBuilder{m.Fixed}
}

// Finish
func (m MarketOrdersRequestFixedBuilder) Finish() MarketOrdersRequestFixed {
	return MarketOrdersRequestFixed{m.Fixed}
}

// RequestAction
func (m MarketOrdersRequestFixed) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Unsafe(), 8, 4))
}

// RequestAction
func (m MarketOrdersRequestFixedBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Unsafe(), 8, 4))
}

// SymbolID
func (m MarketOrdersRequestFixed) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// SymbolID
func (m MarketOrdersRequestFixedBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// Symbol
func (m MarketOrdersRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 76, 12)
}

// Symbol
func (m MarketOrdersRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 76, 12)
}

// Exchange
func (m MarketOrdersRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 92, 76)
}

// Exchange
func (m MarketOrdersRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 92, 76)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixed) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 96, 92)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixedBuilder) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 96, 92)
}

// SetRequestAction
func (m MarketOrdersRequestFixedBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, int32(value))
}

// SetSymbolID
func (m MarketOrdersRequestFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 12, 8, value)
}

// SetSymbol
func (m MarketOrdersRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 76, 12, value)
}

// SetExchange
func (m MarketOrdersRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 92, 76, value)
}

// SetSendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixedBuilder) SetSendQuantitiesGreaterOrEqualTo(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 96, 92, value)
}

// Copy
func (m MarketOrdersRequestFixed) Copy(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// Copy
func (m MarketOrdersRequestFixedBuilder) Copy(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestFixed) CopyPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestFixedBuilder) CopyPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestFixed) CopyToPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestFixedBuilder) CopyToPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestFixed) CopyTo(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestFixedBuilder) CopyTo(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}
