package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = MarketOrdersRejectFixedSize  (104)
//     Type       = MARKET_ORDERS_REJECT  (151)
//     SymbolID   = 0
//     RejectText = ""
// }
var _MarketOrdersRejectFixedDefault = []byte{104, 0, 151, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketOrdersRejectFixedSize = 104

type MarketOrdersRejectFixed struct {
	message.Fixed
}

type MarketOrdersRejectFixedBuilder struct {
	message.Fixed
}

func NewMarketOrdersRejectFixedFrom(b []byte) MarketOrdersRejectFixed {
	return MarketOrdersRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketOrdersRejectFixed(b []byte) MarketOrdersRejectFixed {
	return MarketOrdersRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewMarketOrdersRejectFixed() MarketOrdersRejectFixedBuilder {
	a := MarketOrdersRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _MarketOrdersRejectFixedDefault)
	return a
}

// Clear
// {
//     Size       = MarketOrdersRejectFixedSize  (104)
//     Type       = MARKET_ORDERS_REJECT  (151)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketOrdersRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersRejectFixedDefault)
}

// ToBuilder
func (m MarketOrdersRejectFixed) ToBuilder() MarketOrdersRejectFixedBuilder {
	return MarketOrdersRejectFixedBuilder{m.Fixed}
}

// Finish
func (m MarketOrdersRejectFixedBuilder) Finish() MarketOrdersRejectFixed {
	return MarketOrdersRejectFixed{m.Fixed}
}

// SymbolID
func (m MarketOrdersRejectFixed) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketOrdersRejectFixedBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RejectText
func (m MarketOrdersRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m MarketOrdersRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// SetSymbolID
func (m MarketOrdersRejectFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText
func (m MarketOrdersRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// Copy
func (m MarketOrdersRejectFixed) Copy(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketOrdersRejectFixedBuilder) Copy(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectFixed) CopyPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectFixedBuilder) CopyPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectFixed) CopyToPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectFixedBuilder) CopyToPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectFixed) CopyTo(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectFixedBuilder) CopyTo(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
