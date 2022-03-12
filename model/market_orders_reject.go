package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketOrdersRejectSize = 16

const MarketOrdersRejectFixedSize = 104

// {
//     Size       = MarketOrdersRejectSize  (16)
//     Type       = MARKET_ORDERS_REJECT  (151)
//     BaseSize   = MarketOrdersRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
var _MarketOrdersRejectDefault = []byte{16, 0, 151, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size       = MarketOrdersRejectFixedSize  (104)
//     Type       = MARKET_ORDERS_REJECT  (151)
//     SymbolID   = 0
//     RejectText = ""
// }
var _MarketOrdersRejectFixedDefault = []byte{104, 0, 151, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketOrdersReject struct {
	message.VLS
}

type MarketOrdersRejectBuilder struct {
	message.VLSBuilder
}

type MarketOrdersRejectFixed struct {
	message.Fixed
}

type MarketOrdersRejectFixedBuilder struct {
	message.Fixed
}

type MarketOrdersRejectPointer struct {
	message.VLSPointer
}

type MarketOrdersRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

type MarketOrdersRejectFixedPointer struct {
	message.FixedPointer
}

type MarketOrdersRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func NewMarketOrdersRejectFrom(b []byte) MarketOrdersReject {
	return MarketOrdersReject{VLS: message.NewVLSFrom(b)}
}

func WrapMarketOrdersReject(b []byte) MarketOrdersReject {
	return MarketOrdersReject{VLS: message.WrapVLS(b)}
}

func NewMarketOrdersReject() MarketOrdersRejectBuilder {
	a := MarketOrdersRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _MarketOrdersRejectDefault)
	return a
}

func NewMarketOrdersRejectFixedFrom(b []byte) MarketOrdersRejectFixed {
	return MarketOrdersRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketOrdersRejectFixed(b []byte) MarketOrdersRejectFixed {
	return MarketOrdersRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewMarketOrdersRejectFixed() MarketOrdersRejectFixedBuilder {
	a := MarketOrdersRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _MarketOrdersRejectFixedDefault)
	return a
}

func AllocMarketOrdersReject() MarketOrdersRejectPointerBuilder {
	a := MarketOrdersRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _MarketOrdersRejectDefault)
	return a
}

func AllocMarketOrdersRejectFrom(b []byte) MarketOrdersRejectPointer {
	return MarketOrdersRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocMarketOrdersRejectFixed() MarketOrdersRejectFixedPointerBuilder {
	a := MarketOrdersRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _MarketOrdersRejectFixedDefault)
	return a
}

func AllocMarketOrdersRejectFixedFrom(b []byte) MarketOrdersRejectFixedPointer {
	return MarketOrdersRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = MarketOrdersRejectSize  (16)
//     Type       = MARKET_ORDERS_REJECT  (151)
//     BaseSize   = MarketOrdersRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketOrdersRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersRejectDefault)
}

// Clear
// {
//     Size       = MarketOrdersRejectFixedSize  (104)
//     Type       = MARKET_ORDERS_REJECT  (151)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketOrdersRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersRejectFixedDefault)
}

// Clear
// {
//     Size       = MarketOrdersRejectSize  (16)
//     Type       = MARKET_ORDERS_REJECT  (151)
//     BaseSize   = MarketOrdersRejectSize  (16)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketOrdersRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersRejectDefault)
}

// Clear
// {
//     Size       = MarketOrdersRejectFixedSize  (104)
//     Type       = MARKET_ORDERS_REJECT  (151)
//     SymbolID   = 0
//     RejectText = ""
// }
func (m MarketOrdersRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersRejectFixedDefault)
}

// ToBuilder
func (m MarketOrdersReject) ToBuilder() MarketOrdersRejectBuilder {
	return MarketOrdersRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m MarketOrdersRejectFixed) ToBuilder() MarketOrdersRejectFixedBuilder {
	return MarketOrdersRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m MarketOrdersRejectPointer) ToBuilder() MarketOrdersRejectPointerBuilder {
	return MarketOrdersRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m MarketOrdersRejectFixedPointer) ToBuilder() MarketOrdersRejectFixedPointerBuilder {
	return MarketOrdersRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketOrdersRejectBuilder) Finish() MarketOrdersReject {
	return MarketOrdersReject{m.VLSBuilder.Finish()}
}

// Finish
func (m MarketOrdersRejectFixedBuilder) Finish() MarketOrdersRejectFixed {
	return MarketOrdersRejectFixed{m.Fixed}
}

// Finish
func (m *MarketOrdersRejectPointerBuilder) Finish() MarketOrdersRejectPointer {
	return MarketOrdersRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *MarketOrdersRejectFixedPointerBuilder) Finish() MarketOrdersRejectFixedPointer {
	return MarketOrdersRejectFixedPointer{m.FixedPointer}
}

// SymbolID
func (m MarketOrdersReject) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 12, 8)
}

// SymbolID
func (m MarketOrdersRejectBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 12, 8)
}

// SymbolID
func (m MarketOrdersRejectPointer) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 12, 8)
}

// SymbolID
func (m MarketOrdersRejectPointerBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 12, 8)
}

// RejectText
func (m MarketOrdersReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText
func (m MarketOrdersRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText
func (m MarketOrdersRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText
func (m MarketOrdersRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SymbolID
func (m MarketOrdersRejectFixed) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketOrdersRejectFixedBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketOrdersRejectFixedPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketOrdersRejectFixedPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RejectText
func (m MarketOrdersRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m MarketOrdersRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m MarketOrdersRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText
func (m MarketOrdersRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetSymbolID
func (m MarketOrdersRejectBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 12, 8, value)
}

// SetSymbolID
func (m MarketOrdersRejectPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText
func (m *MarketOrdersRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText
func (m *MarketOrdersRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetSymbolID
func (m MarketOrdersRejectFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID
func (m MarketOrdersRejectFixedPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText
func (m MarketOrdersRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText
func (m MarketOrdersRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m MarketOrdersReject) Copy(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketOrdersRejectBuilder) Copy(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersReject) CopyPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectBuilder) CopyPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersReject) CopyToPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectBuilder) CopyToPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersReject) CopyTo(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectBuilder) CopyTo(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketOrdersRejectFixed) Copy(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketOrdersRejectFixedBuilder) Copy(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectFixed) CopyPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectFixedBuilder) CopyPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectFixed) CopyToPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectFixedBuilder) CopyToPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectFixed) CopyTo(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectFixedBuilder) CopyTo(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketOrdersRejectPointer) Copy(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketOrdersRejectPointerBuilder) Copy(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectPointer) CopyPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectPointerBuilder) CopyPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectPointer) CopyTo(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectPointerBuilder) CopyTo(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectPointer) CopyToPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectPointerBuilder) CopyToPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketOrdersRejectFixedPointer) Copy(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketOrdersRejectFixedPointerBuilder) Copy(to MarketOrdersRejectFixedBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectFixedPointer) CopyPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m MarketOrdersRejectFixedPointerBuilder) CopyPointer(to MarketOrdersRejectFixedPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectFixedPointer) CopyTo(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectFixedPointerBuilder) CopyTo(to MarketOrdersRejectBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectFixedPointer) CopyToPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m MarketOrdersRejectFixedPointerBuilder) CopyToPointer(to MarketOrdersRejectPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}
