package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalTradesRejectSize = 14

const HistoricalTradesRejectFixedSize = 104

// {
//     Size       = HistoricalTradesRejectSize  (14)
//     Type       = HISTORICAL_TRADES_REJECT  (10101)
//     BaseSize   = HistoricalTradesRejectSize  (14)
//     RequestID  = 0
//     RejectText = ""
// }
var _HistoricalTradesRejectDefault = []byte{14, 0, 117, 39, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size       = HistoricalTradesRejectFixedSize  (104)
//     Type       = HISTORICAL_TRADES_REJECT  (10101)
//     RequestID  = 0
//     RejectText = ""
// }
var _HistoricalTradesRejectFixedDefault = []byte{104, 0, 117, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HistoricalTradesReject struct {
	message.VLS
}

type HistoricalTradesRejectBuilder struct {
	message.VLSBuilder
}

type HistoricalTradesRejectFixed struct {
	message.Fixed
}

type HistoricalTradesRejectFixedBuilder struct {
	message.Fixed
}

type HistoricalTradesRejectPointer struct {
	message.VLSPointer
}

type HistoricalTradesRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

type HistoricalTradesRejectFixedPointer struct {
	message.FixedPointer
}

type HistoricalTradesRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalTradesRejectFrom(b []byte) HistoricalTradesReject {
	return HistoricalTradesReject{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalTradesReject(b []byte) HistoricalTradesReject {
	return HistoricalTradesReject{VLS: message.WrapVLS(b)}
}

func NewHistoricalTradesReject() HistoricalTradesRejectBuilder {
	a := HistoricalTradesRejectBuilder{message.NewVLSBuilder(14)}
	a.Unsafe().SetBytes(0, _HistoricalTradesRejectDefault)
	return a
}

func NewHistoricalTradesRejectFixedFrom(b []byte) HistoricalTradesRejectFixed {
	return HistoricalTradesRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalTradesRejectFixed(b []byte) HistoricalTradesRejectFixed {
	return HistoricalTradesRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalTradesRejectFixed() HistoricalTradesRejectFixedBuilder {
	a := HistoricalTradesRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _HistoricalTradesRejectFixedDefault)
	return a
}

func AllocHistoricalTradesReject() HistoricalTradesRejectPointerBuilder {
	a := HistoricalTradesRejectPointerBuilder{message.AllocVLSBuilder(14)}
	a.Ptr.SetBytes(0, _HistoricalTradesRejectDefault)
	return a
}

func AllocHistoricalTradesRejectFrom(b []byte) HistoricalTradesRejectPointer {
	return HistoricalTradesRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocHistoricalTradesRejectFixed() HistoricalTradesRejectFixedPointerBuilder {
	a := HistoricalTradesRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _HistoricalTradesRejectFixedDefault)
	return a
}

func AllocHistoricalTradesRejectFixedFrom(b []byte) HistoricalTradesRejectFixedPointer {
	return HistoricalTradesRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = HistoricalTradesRejectSize  (14)
//     Type       = HISTORICAL_TRADES_REJECT  (10101)
//     BaseSize   = HistoricalTradesRejectSize  (14)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalTradesRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalTradesRejectDefault)
}

// Clear
// {
//     Size       = HistoricalTradesRejectFixedSize  (104)
//     Type       = HISTORICAL_TRADES_REJECT  (10101)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalTradesRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalTradesRejectFixedDefault)
}

// Clear
// {
//     Size       = HistoricalTradesRejectSize  (14)
//     Type       = HISTORICAL_TRADES_REJECT  (10101)
//     BaseSize   = HistoricalTradesRejectSize  (14)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalTradesRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalTradesRejectDefault)
}

// Clear
// {
//     Size       = HistoricalTradesRejectFixedSize  (104)
//     Type       = HISTORICAL_TRADES_REJECT  (10101)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalTradesRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalTradesRejectFixedDefault)
}

// ToBuilder
func (m HistoricalTradesReject) ToBuilder() HistoricalTradesRejectBuilder {
	return HistoricalTradesRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalTradesRejectFixed) ToBuilder() HistoricalTradesRejectFixedBuilder {
	return HistoricalTradesRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalTradesRejectPointer) ToBuilder() HistoricalTradesRejectPointerBuilder {
	return HistoricalTradesRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalTradesRejectFixedPointer) ToBuilder() HistoricalTradesRejectFixedPointerBuilder {
	return HistoricalTradesRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalTradesRejectBuilder) Finish() HistoricalTradesReject {
	return HistoricalTradesReject{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalTradesRejectFixedBuilder) Finish() HistoricalTradesRejectFixed {
	return HistoricalTradesRejectFixed{m.Fixed}
}

// Finish
func (m *HistoricalTradesRejectPointerBuilder) Finish() HistoricalTradesRejectPointer {
	return HistoricalTradesRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalTradesRejectFixedPointerBuilder) Finish() HistoricalTradesRejectFixedPointer {
	return HistoricalTradesRejectFixedPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalTradesReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m HistoricalTradesRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m HistoricalTradesRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m HistoricalTradesRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// RejectText
func (m HistoricalTradesReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// RejectText
func (m HistoricalTradesRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// RejectText
func (m HistoricalTradesRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// RejectText
func (m HistoricalTradesRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// RequestID
func (m HistoricalTradesRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalTradesRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalTradesRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalTradesRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText
func (m HistoricalTradesRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m HistoricalTradesRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m HistoricalTradesRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText
func (m HistoricalTradesRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID
func (m HistoricalTradesRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m HistoricalTradesRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 10, 6, value)
}

// SetRejectText
func (m *HistoricalTradesRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetRejectText
func (m *HistoricalTradesRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetRequestID
func (m HistoricalTradesRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m HistoricalTradesRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText
func (m HistoricalTradesRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText
func (m HistoricalTradesRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m HistoricalTradesReject) Copy(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalTradesRejectBuilder) Copy(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesReject) CopyPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectBuilder) CopyPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesReject) CopyToPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectBuilder) CopyToPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesReject) CopyTo(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectBuilder) CopyTo(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalTradesRejectFixed) Copy(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalTradesRejectFixedBuilder) Copy(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectFixed) CopyPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectFixedBuilder) CopyPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectFixed) CopyToPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectFixedBuilder) CopyToPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectFixed) CopyTo(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectFixedBuilder) CopyTo(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalTradesRejectPointer) Copy(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalTradesRejectPointerBuilder) Copy(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectPointer) CopyPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectPointerBuilder) CopyPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectPointer) CopyTo(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectPointerBuilder) CopyTo(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectPointer) CopyToPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectPointerBuilder) CopyToPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalTradesRejectFixedPointer) Copy(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalTradesRejectFixedPointerBuilder) Copy(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectFixedPointer) CopyPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectFixedPointerBuilder) CopyPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectFixedPointer) CopyTo(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectFixedPointerBuilder) CopyTo(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectFixedPointer) CopyToPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectFixedPointerBuilder) CopyToPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
