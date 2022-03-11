package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = MarketDataRejectFixedSize  (104)
//     Type       = MARKET_DATA_REJECT  (103)
//     SymbolID   = 0
//     RejectText = ""
// }
var _MarketDataRejectFixedDefault = []byte{104, 0, 103, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataRejectFixedSize = 104

// MarketDataRejectFixed The MarketDataReject message is sent by the Server to the Client to reject
// a MarketDataRequest message for any reason.
type MarketDataRejectFixed struct {
	message.Fixed
}

// MarketDataRejectFixedBuilder The MarketDataReject message is sent by the Server to the Client to reject
// a MarketDataRequest message for any reason.
type MarketDataRejectFixedBuilder struct {
	message.Fixed
}

func NewMarketDataRejectFixedFrom(b []byte) MarketDataRejectFixed {
	return MarketDataRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataRejectFixed(b []byte) MarketDataRejectFixed {
	return MarketDataRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewMarketDataRejectFixed() MarketDataRejectFixedBuilder {
	a := MarketDataRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _MarketDataRejectFixedDefault)
	return a
}

// Clear
// {
//     Size       = MarketDataRejectFixedSize  (104)
//     Type       = MARKET_DATA_REJECT  (103)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketDataRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataRejectFixedDefault)
}

// ToBuilder
func (m MarketDataRejectFixed) ToBuilder() MarketDataRejectFixedBuilder {
	return MarketDataRejectFixedBuilder{m.Fixed}
}

// Finish
func (m MarketDataRejectFixedBuilder) Finish() MarketDataRejectFixed {
	return MarketDataRejectFixed{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectFixed) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectFixedBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// Copy
func (m MarketDataRejectFixed) Copy(to MarketDataRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketDataRejectFixedBuilder) Copy(to MarketDataRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDataRejectFixed) CopyPointer(to MarketDataRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDataRejectFixedBuilder) CopyPointer(to MarketDataRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDataRejectFixed) CopyToPointer(to MarketDataRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDataRejectFixedBuilder) CopyToPointer(to MarketDataRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDataRejectFixed) CopyTo(to MarketDataRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDataRejectFixedBuilder) CopyTo(to MarketDataRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
