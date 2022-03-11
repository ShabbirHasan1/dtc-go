package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDepthRejectPointer The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectPointer struct {
	message.VLSPointer
}

// MarketDepthRejectPointerBuilder The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocMarketDepthReject() MarketDepthRejectPointerBuilder {
	a := MarketDepthRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _MarketDepthRejectDefault)
	return a
}

func AllocMarketDepthRejectFrom(b []byte) MarketDepthRejectPointer {
	return MarketDepthRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size       = MarketDepthRejectSize  (16)
//     Type       = MARKET_DEPTH_REJECT  (121)
//     BaseSize   = MarketDepthRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketDepthRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthRejectDefault)
}

// ToBuilder
func (m MarketDepthRejectPointer) ToBuilder() MarketDepthRejectPointerBuilder {
	return MarketDepthRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *MarketDepthRejectPointerBuilder) Finish() MarketDepthRejectPointer {
	return MarketDepthRejectPointer{m.VLSPointerBuilder.Finish()}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectPointer) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 12, 8)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectPointerBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 12, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m *MarketDepthRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// Copy
func (m MarketDepthRejectPointer) Copy(to MarketDepthRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketDepthRejectPointerBuilder) Copy(to MarketDepthRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDepthRejectPointer) CopyPointer(to MarketDepthRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDepthRejectPointerBuilder) CopyPointer(to MarketDepthRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDepthRejectPointer) CopyTo(to MarketDepthRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDepthRejectPointerBuilder) CopyTo(to MarketDepthRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDepthRejectPointer) CopyToPointer(to MarketDepthRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDepthRejectPointerBuilder) CopyToPointer(to MarketDepthRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
