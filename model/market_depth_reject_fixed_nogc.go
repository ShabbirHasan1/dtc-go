package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDepthRejectFixedPointer The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectFixedPointer struct {
	message.FixedPointer
}

// MarketDepthRejectFixedPointerBuilder The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDepthRejectFixed() MarketDepthRejectFixedPointerBuilder {
	a := MarketDepthRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _MarketDepthRejectFixedDefault)
	return a
}

func AllocMarketDepthRejectFixedFrom(b []byte) MarketDepthRejectFixedPointer {
	return MarketDepthRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = MarketDepthRejectFixedSize  (104)
//     Type       = MARKET_DEPTH_REJECT  (121)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketDepthRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthRejectFixedDefault)
}

// ToBuilder
func (m MarketDepthRejectFixedPointer) ToBuilder() MarketDepthRejectFixedPointerBuilder {
	return MarketDepthRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDepthRejectFixedPointerBuilder) Finish() MarketDepthRejectFixedPointer {
	return MarketDepthRejectFixedPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixedPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixedPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixedPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m MarketDepthRejectFixedPointer) Copy(to MarketDepthRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketDepthRejectFixedPointerBuilder) Copy(to MarketDepthRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDepthRejectFixedPointer) CopyPointer(to MarketDepthRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDepthRejectFixedPointerBuilder) CopyPointer(to MarketDepthRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDepthRejectFixedPointer) CopyTo(to MarketDepthRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDepthRejectFixedPointerBuilder) CopyTo(to MarketDepthRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDepthRejectFixedPointer) CopyToPointer(to MarketDepthRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDepthRejectFixedPointerBuilder) CopyToPointer(to MarketDepthRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
