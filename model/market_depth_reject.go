package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = MarketDepthRejectSize  (16)
//     Type       = MARKET_DEPTH_REJECT  (121)
//     BaseSize   = MarketDepthRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
var _MarketDepthRejectDefault = []byte{16, 0, 121, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDepthRejectSize = 16

// MarketDepthReject The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthReject struct {
	message.VLS
}

// MarketDepthRejectBuilder The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectBuilder struct {
	message.VLSBuilder
}

func NewMarketDepthRejectFrom(b []byte) MarketDepthReject {
	return MarketDepthReject{VLS: message.NewVLSFrom(b)}
}

func WrapMarketDepthReject(b []byte) MarketDepthReject {
	return MarketDepthReject{VLS: message.WrapVLS(b)}
}

func NewMarketDepthReject() MarketDepthRejectBuilder {
	a := MarketDepthRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _MarketDepthRejectDefault)
	return a
}

// Clear
// {
//     Size       = MarketDepthRejectSize  (16)
//     Type       = MARKET_DEPTH_REJECT  (121)
//     BaseSize   = MarketDepthRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketDepthRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthRejectDefault)
}

// ToBuilder
func (m MarketDepthReject) ToBuilder() MarketDepthRejectBuilder {
	return MarketDepthRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m MarketDepthRejectBuilder) Finish() MarketDepthReject {
	return MarketDepthReject{m.VLSBuilder.Finish()}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthReject) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 12, 8)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 12, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m *MarketDepthRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// Copy
func (m MarketDepthReject) Copy(to MarketDepthRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketDepthRejectBuilder) Copy(to MarketDepthRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDepthReject) CopyPointer(to MarketDepthRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDepthRejectBuilder) CopyPointer(to MarketDepthRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDepthReject) CopyToPointer(to MarketDepthRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDepthRejectBuilder) CopyToPointer(to MarketDepthRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDepthReject) CopyTo(to MarketDepthRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDepthRejectBuilder) CopyTo(to MarketDepthRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
