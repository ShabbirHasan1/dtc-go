package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDepthRejectSize = 16

const MarketDepthRejectFixedSize = 104

// {
//     Size       = MarketDepthRejectSize  (16)
//     Type       = MARKET_DEPTH_REJECT  (121)
//     BaseSize   = MarketDepthRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
var _MarketDepthRejectDefault = []byte{16, 0, 121, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size       = MarketDepthRejectFixedSize  (104)
//     Type       = MARKET_DEPTH_REJECT  (121)
//     SymbolID   = 0
//     RejectText = ""
// }
var _MarketDepthRejectFixedDefault = []byte{104, 0, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocMarketDepthReject() MarketDepthRejectPointerBuilder {
	a := MarketDepthRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _MarketDepthRejectDefault)
	return a
}

func AllocMarketDepthRejectFrom(b []byte) MarketDepthRejectPointer {
	return MarketDepthRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size       = MarketDepthRejectSize  (16)
//     Type       = MARKET_DEPTH_REJECT  (121)
//     BaseSize   = MarketDepthRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketDepthRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthRejectDefault)
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
func (m MarketDepthReject) ToBuilder() MarketDepthRejectBuilder {
	return MarketDepthRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m MarketDepthRejectFixed) ToBuilder() MarketDepthRejectFixedBuilder {
	return MarketDepthRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDepthRejectPointer) ToBuilder() MarketDepthRejectPointerBuilder {
	return MarketDepthRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m MarketDepthRejectFixedPointer) ToBuilder() MarketDepthRejectFixedPointerBuilder {
	return MarketDepthRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDepthRejectBuilder) Finish() MarketDepthReject {
	return MarketDepthReject{m.VLSBuilder.Finish()}
}

// Finish
func (m MarketDepthRejectFixedBuilder) Finish() MarketDepthRejectFixed {
	return MarketDepthRejectFixed{m.Fixed}
}

// Finish
func (m *MarketDepthRejectPointerBuilder) Finish() MarketDepthRejectPointer {
	return MarketDepthRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *MarketDepthRejectFixedPointerBuilder) Finish() MarketDepthRejectFixedPointer {
	return MarketDepthRejectFixedPointer{m.FixedPointer}
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
func (m MarketDepthReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
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
func (m MarketDepthRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
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
func (m MarketDepthRejectBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 12, 8, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m *MarketDepthRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m *MarketDepthRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixedPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
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
