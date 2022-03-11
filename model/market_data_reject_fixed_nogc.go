package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDataRejectFixedPointer The MarketDataReject message is sent by the Server to the Client to reject
// a MarketDataRequest message for any reason.
type MarketDataRejectFixedPointer struct {
	message.FixedPointer
}

// MarketDataRejectFixedPointerBuilder The MarketDataReject message is sent by the Server to the Client to reject
// a MarketDataRequest message for any reason.
type MarketDataRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataRejectFixed() MarketDataRejectFixedPointerBuilder {
	a := MarketDataRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _MarketDataRejectFixedDefault)
	return a
}

func AllocMarketDataRejectFixedFrom(b []byte) MarketDataRejectFixedPointer {
	return MarketDataRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = MarketDataRejectFixedSize  (104)
//     Type       = MARKET_DATA_REJECT  (103)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketDataRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataRejectFixedDefault)
}

// ToBuilder
func (m MarketDataRejectFixedPointer) ToBuilder() MarketDataRejectFixedPointerBuilder {
	return MarketDataRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataRejectFixedPointerBuilder) Finish() MarketDataRejectFixedPointer {
	return MarketDataRejectFixedPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectFixedPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectFixedPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectFixedPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m MarketDataRejectFixedPointer) Copy(to MarketDataRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketDataRejectFixedPointerBuilder) Copy(to MarketDataRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDataRejectFixedPointer) CopyPointer(to MarketDataRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDataRejectFixedPointerBuilder) CopyPointer(to MarketDataRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDataRejectFixedPointer) CopyTo(to MarketDataRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDataRejectFixedPointerBuilder) CopyTo(to MarketDataRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDataRejectFixedPointer) CopyToPointer(to MarketDataRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDataRejectFixedPointerBuilder) CopyToPointer(to MarketDataRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
