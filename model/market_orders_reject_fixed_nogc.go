package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketOrdersRejectFixedPointer struct {
	message.FixedPointer
}

type MarketOrdersRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketOrdersRejectFixed() MarketOrdersRejectFixedPointerBuilder {
	a := MarketOrdersRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _MarketOrdersRejectFixedDefault)
	return a
}

func AllocMarketOrdersRejectFixedFrom(b []byte) MarketOrdersRejectFixedPointer {
	return MarketOrdersRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = MarketOrdersRejectFixedSize  (104)
//     Type       = MARKET_ORDERS_REJECT  (151)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketOrdersRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersRejectFixedDefault)
}

// ToBuilder
func (m MarketOrdersRejectFixedPointer) ToBuilder() MarketOrdersRejectFixedPointerBuilder {
	return MarketOrdersRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketOrdersRejectFixedPointerBuilder) Finish() MarketOrdersRejectFixedPointer {
	return MarketOrdersRejectFixedPointer{m.FixedPointer}
}

// SymbolID
func (m MarketOrdersRejectFixedPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketOrdersRejectFixedPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RejectText
func (m MarketOrdersRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText
func (m MarketOrdersRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetSymbolID
func (m MarketOrdersRejectFixedPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText
func (m MarketOrdersRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m MarketOrdersRejectFixedPointer) Copy(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketOrdersRejectFixedPointerBuilder) Copy(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectFixedPointer) CopyPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectFixedPointerBuilder) CopyPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectFixedPointer) CopyTo(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectFixedPointerBuilder) CopyTo(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectFixedPointer) CopyToPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectFixedPointerBuilder) CopyToPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
