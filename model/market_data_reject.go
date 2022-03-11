package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = MarketDataRejectSize  (16)
//     Type       = MARKET_DATA_REJECT  (103)
//     BaseSize   = MarketDataRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
var _MarketDataRejectDefault = []byte{16, 0, 103, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataRejectSize = 16

// MarketDataReject The MarketDataReject message is sent by the Server to the Client to reject
// a MarketDataRequest message for any reason.
type MarketDataReject struct {
	message.VLS
}

// MarketDataRejectBuilder The MarketDataReject message is sent by the Server to the Client to reject
// a MarketDataRequest message for any reason.
type MarketDataRejectBuilder struct {
	message.VLSBuilder
}

func NewMarketDataRejectFrom(b []byte) MarketDataReject {
	return MarketDataReject{VLS: message.NewVLSFrom(b)}
}

func WrapMarketDataReject(b []byte) MarketDataReject {
	return MarketDataReject{VLS: message.WrapVLS(b)}
}

func NewMarketDataReject() MarketDataRejectBuilder {
	a := MarketDataRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _MarketDataRejectDefault)
	return a
}

// Clear
// {
//     Size       = MarketDataRejectSize  (16)
//     Type       = MARKET_DATA_REJECT  (103)
//     BaseSize   = MarketDataRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketDataRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataRejectDefault)
}

// ToBuilder
func (m MarketDataReject) ToBuilder() MarketDataRejectBuilder {
	return MarketDataRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m MarketDataRejectBuilder) Finish() MarketDataReject {
	return MarketDataReject{m.VLSBuilder.Finish()}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataReject) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 12, 8)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 12, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m *MarketDataRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// Copy
func (m MarketDataReject) Copy(to MarketDataRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketDataRejectBuilder) Copy(to MarketDataRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDataReject) CopyPointer(to MarketDataRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketDataRejectBuilder) CopyPointer(to MarketDataRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDataReject) CopyToPointer(to MarketDataRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketDataRejectBuilder) CopyToPointer(to MarketDataRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDataReject) CopyTo(to MarketDataRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketDataRejectBuilder) CopyTo(to MarketDataRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
