package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDataRejectPointer The MarketDataReject message is sent by the Server to the Client to reject
// a MarketDataRequest message for any reason.
type MarketDataRejectPointer struct {
	message.VLSPointer
}

// MarketDataRejectPointerBuilder The MarketDataReject message is sent by the Server to the Client to reject
// a MarketDataRequest message for any reason.
type MarketDataRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocMarketDataReject() MarketDataRejectPointerBuilder {
	a := MarketDataRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _MarketDataRejectDefault)
	return a
}

func AllocMarketDataRejectFrom(b []byte) MarketDataRejectPointer {
	return MarketDataRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size       = MarketDataRejectSize  (16)
//     Type       = MARKET_DATA_REJECT  (103)
//     BaseSize   = MarketDataRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketDataRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataRejectDefault)
}

// ToBuilder
func (m MarketDataRejectPointer) ToBuilder() MarketDataRejectPointerBuilder {
	return MarketDataRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *MarketDataRejectPointerBuilder) Finish() MarketDataRejectPointer {
	return MarketDataRejectPointer{m.VLSPointerBuilder.Finish()}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectPointer) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 12, 8)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectPointerBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 12, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m *MarketDataRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// Copy
func (m MarketDataRejectPointer) Copy(to MarketDataRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketDataRejectPointerBuilder) Copy(to MarketDataRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDataRejectPointer) CopyPointer(to MarketDataRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDataRejectPointerBuilder) CopyPointer(to MarketDataRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDataRejectPointer) CopyTo(to MarketDataRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDataRejectPointerBuilder) CopyTo(to MarketDataRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDataRejectPointer) CopyToPointer(to MarketDataRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDataRejectPointerBuilder) CopyToPointer(to MarketDataRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
