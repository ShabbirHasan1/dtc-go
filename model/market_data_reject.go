package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataRejectSize = 16

const MarketDataRejectFixedSize = 104

// {
//     Size       = MarketDataRejectSize  (16)
//     Type       = MARKET_DATA_REJECT  (103)
//     BaseSize   = MarketDataRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
var _MarketDataRejectDefault = []byte{16, 0, 103, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size       = MarketDataRejectFixedSize  (104)
//     Type       = MARKET_DATA_REJECT  (103)
//     SymbolID   = 0
//     RejectText = ""
// }
var _MarketDataRejectFixedDefault = []byte{104, 0, 103, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocMarketDataReject() MarketDataRejectPointerBuilder {
	a := MarketDataRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _MarketDataRejectDefault)
	return a
}

func AllocMarketDataRejectFrom(b []byte) MarketDataRejectPointer {
	return MarketDataRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size       = MarketDataRejectSize  (16)
//     Type       = MARKET_DATA_REJECT  (103)
//     BaseSize   = MarketDataRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketDataRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataRejectDefault)
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
func (m MarketDataReject) ToBuilder() MarketDataRejectBuilder {
	return MarketDataRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m MarketDataRejectFixed) ToBuilder() MarketDataRejectFixedBuilder {
	return MarketDataRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataRejectPointer) ToBuilder() MarketDataRejectPointerBuilder {
	return MarketDataRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m MarketDataRejectFixedPointer) ToBuilder() MarketDataRejectFixedPointerBuilder {
	return MarketDataRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataRejectBuilder) Finish() MarketDataReject {
	return MarketDataReject{m.VLSBuilder.Finish()}
}

// Finish
func (m MarketDataRejectFixedBuilder) Finish() MarketDataRejectFixed {
	return MarketDataRejectFixed{m.Fixed}
}

// Finish
func (m *MarketDataRejectPointerBuilder) Finish() MarketDataRejectPointer {
	return MarketDataRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *MarketDataRejectFixedPointerBuilder) Finish() MarketDataRejectFixedPointer {
	return MarketDataRejectFixedPointer{m.FixedPointer}
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
func (m MarketDataReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
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
func (m MarketDataRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
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
func (m MarketDataRejectBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 12, 8, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m *MarketDataRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m *MarketDataRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataRejectFixedPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m MarketDataRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
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
