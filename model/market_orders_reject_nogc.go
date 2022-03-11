package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketOrdersRejectPointer struct {
	message.VLSPointer
}

type MarketOrdersRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocMarketOrdersReject() MarketOrdersRejectPointerBuilder {
	a := MarketOrdersRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _MarketOrdersRejectDefault)
	return a
}

func AllocMarketOrdersRejectFrom(b []byte) MarketOrdersRejectPointer {
	return MarketOrdersRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size       = MarketOrdersRejectSize  (16)
//     Type       = MARKET_ORDERS_REJECT  (151)
//     BaseSize   = MarketOrdersRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketOrdersRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersRejectDefault)
}

// ToBuilder
func (m MarketOrdersRejectPointer) ToBuilder() MarketOrdersRejectPointerBuilder {
	return MarketOrdersRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *MarketOrdersRejectPointerBuilder) Finish() MarketOrdersRejectPointer {
	return MarketOrdersRejectPointer{m.VLSPointerBuilder.Finish()}
}

// SymbolID
func (m MarketOrdersRejectPointer) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 12, 8)
}

// SymbolID
func (m MarketOrdersRejectPointerBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 12, 8)
}

// RejectText
func (m MarketOrdersRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText
func (m MarketOrdersRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetSymbolID
func (m MarketOrdersRejectPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText
func (m *MarketOrdersRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// Copy
func (m MarketOrdersRejectPointer) Copy(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketOrdersRejectPointerBuilder) Copy(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectPointer) CopyPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectPointerBuilder) CopyPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectPointer) CopyTo(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectPointerBuilder) CopyTo(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectPointer) CopyToPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectPointerBuilder) CopyToPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
