package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = MarketDepthRejectFixedSize  (104)
//     Type       = MARKET_DEPTH_REJECT  (121)
//     SymbolID   = 0
//     RejectText = ""
// }
var _MarketDepthRejectFixedDefault = []byte{104, 0, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDepthRejectFixedSize = 104

// MarketDepthRejectFixed The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectFixed struct {
	message.Fixed
}

// MarketDepthRejectFixedBuilder The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectFixedBuilder struct {
	message.Fixed
}

func NewMarketDepthRejectFixedFrom(b []byte) MarketDepthRejectFixed {
	return MarketDepthRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDepthRejectFixed(b []byte) MarketDepthRejectFixed {
	return MarketDepthRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewMarketDepthRejectFixed() MarketDepthRejectFixedBuilder {
	a := MarketDepthRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _MarketDepthRejectFixedDefault)
	return a
}

// Clear
// {
//     Size       = MarketDepthRejectFixedSize  (104)
//     Type       = MARKET_DEPTH_REJECT  (121)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketDepthRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthRejectFixedDefault)
}

// ToBuilder
func (m MarketDepthRejectFixed) ToBuilder() MarketDepthRejectFixedBuilder {
	return MarketDepthRejectFixedBuilder{m.Fixed}
}

// Finish
func (m MarketDepthRejectFixedBuilder) Finish() MarketDepthRejectFixed {
	return MarketDepthRejectFixed{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixed) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixedBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// Copy
func (m MarketDepthRejectFixed) Copy(to MarketDepthRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketDepthRejectFixedBuilder) Copy(to MarketDepthRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDepthRejectFixed) CopyPointer(to MarketDepthRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDepthRejectFixedBuilder) CopyPointer(to MarketDepthRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDepthRejectFixed) CopyToPointer(to MarketDepthRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDepthRejectFixedBuilder) CopyToPointer(to MarketDepthRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDepthRejectFixed) CopyTo(to MarketDepthRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDepthRejectFixedBuilder) CopyTo(to MarketDepthRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
