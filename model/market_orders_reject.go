package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = MarketOrdersRejectSize  (16)
//     Type       = MARKET_ORDERS_REJECT  (151)
//     BaseSize   = MarketOrdersRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
var _MarketOrdersRejectDefault = []byte{16, 0, 151, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketOrdersRejectSize = 16

type MarketOrdersReject struct {
	message.VLS
}

type MarketOrdersRejectBuilder struct {
	message.VLSBuilder
}

func NewMarketOrdersRejectFrom(b []byte) MarketOrdersReject {
	return MarketOrdersReject{VLS: message.NewVLSFrom(b)}
}

func WrapMarketOrdersReject(b []byte) MarketOrdersReject {
	return MarketOrdersReject{VLS: message.WrapVLS(b)}
}

func NewMarketOrdersReject() MarketOrdersRejectBuilder {
	a := MarketOrdersRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _MarketOrdersRejectDefault)
	return a
}

// Clear
// {
//     Size       = MarketOrdersRejectSize  (16)
//     Type       = MARKET_ORDERS_REJECT  (151)
//     BaseSize   = MarketOrdersRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketOrdersRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersRejectDefault)
}

// ToBuilder
func (m MarketOrdersReject) ToBuilder() MarketOrdersRejectBuilder {
	return MarketOrdersRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m MarketOrdersRejectBuilder) Finish() MarketOrdersReject {
	return MarketOrdersReject{m.VLSBuilder.Finish()}
}

// SymbolID
func (m MarketOrdersReject) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 12, 8)
}

// SymbolID
func (m MarketOrdersRejectBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 12, 8)
}

// RejectText
func (m MarketOrdersReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText
func (m MarketOrdersRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetSymbolID
func (m MarketOrdersRejectBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText
func (m *MarketOrdersRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// Copy
func (m MarketOrdersReject) Copy(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketOrdersRejectBuilder) Copy(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersReject) CopyPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectBuilder) CopyPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersReject) CopyToPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectBuilder) CopyToPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersReject) CopyTo(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectBuilder) CopyTo(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
