package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalOrderFillsRejectSize = 16

const HistoricalOrderFillsRejectFixedSize = 104

// {
//     Size       = HistoricalOrderFillsRejectSize  (16)
//     Type       = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     BaseSize   = HistoricalOrderFillsRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _HistoricalOrderFillsRejectDefault = []byte{16, 0, 52, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size       = HistoricalOrderFillsRejectFixedSize  (104)
//     Type       = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     RequestID  = 0
//     RejectText = ""
// }
var _HistoricalOrderFillsRejectFixedDefault = []byte{104, 0, 52, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// HistoricalOrderFillsReject If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsReject struct {
	message.VLS
}

// HistoricalOrderFillsRejectBuilder If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsRejectBuilder struct {
	message.VLSBuilder
}

// HistoricalOrderFillsRejectFixed If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsRejectFixed struct {
	message.Fixed
}

// HistoricalOrderFillsRejectFixedBuilder If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsRejectFixedBuilder struct {
	message.Fixed
}

// HistoricalOrderFillsRejectPointer If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsRejectPointer struct {
	message.VLSPointer
}

// HistoricalOrderFillsRejectPointerBuilder If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

// HistoricalOrderFillsRejectFixedPointer If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsRejectFixedPointer struct {
	message.FixedPointer
}

// HistoricalOrderFillsRejectFixedPointerBuilder If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalOrderFillsRejectFrom(b []byte) HistoricalOrderFillsReject {
	return HistoricalOrderFillsReject{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalOrderFillsReject(b []byte) HistoricalOrderFillsReject {
	return HistoricalOrderFillsReject{VLS: message.WrapVLS(b)}
}

func NewHistoricalOrderFillsReject() HistoricalOrderFillsRejectBuilder {
	a := HistoricalOrderFillsRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _HistoricalOrderFillsRejectDefault)
	return a
}

func NewHistoricalOrderFillsRejectFixedFrom(b []byte) HistoricalOrderFillsRejectFixed {
	return HistoricalOrderFillsRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalOrderFillsRejectFixed(b []byte) HistoricalOrderFillsRejectFixed {
	return HistoricalOrderFillsRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalOrderFillsRejectFixed() HistoricalOrderFillsRejectFixedBuilder {
	a := HistoricalOrderFillsRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _HistoricalOrderFillsRejectFixedDefault)
	return a
}

func AllocHistoricalOrderFillsReject() HistoricalOrderFillsRejectPointerBuilder {
	a := HistoricalOrderFillsRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _HistoricalOrderFillsRejectDefault)
	return a
}

func AllocHistoricalOrderFillsRejectFrom(b []byte) HistoricalOrderFillsRejectPointer {
	return HistoricalOrderFillsRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocHistoricalOrderFillsRejectFixed() HistoricalOrderFillsRejectFixedPointerBuilder {
	a := HistoricalOrderFillsRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _HistoricalOrderFillsRejectFixedDefault)
	return a
}

func AllocHistoricalOrderFillsRejectFixedFrom(b []byte) HistoricalOrderFillsRejectFixedPointer {
	return HistoricalOrderFillsRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = HistoricalOrderFillsRejectSize  (16)
//     Type       = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     BaseSize   = HistoricalOrderFillsRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalOrderFillsRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalOrderFillsRejectDefault)
}

// Clear
// {
//     Size       = HistoricalOrderFillsRejectFixedSize  (104)
//     Type       = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalOrderFillsRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalOrderFillsRejectFixedDefault)
}

// Clear
// {
//     Size       = HistoricalOrderFillsRejectSize  (16)
//     Type       = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     BaseSize   = HistoricalOrderFillsRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalOrderFillsRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalOrderFillsRejectDefault)
}

// Clear
// {
//     Size       = HistoricalOrderFillsRejectFixedSize  (104)
//     Type       = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalOrderFillsRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalOrderFillsRejectFixedDefault)
}

// ToBuilder
func (m HistoricalOrderFillsReject) ToBuilder() HistoricalOrderFillsRejectBuilder {
	return HistoricalOrderFillsRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalOrderFillsRejectFixed) ToBuilder() HistoricalOrderFillsRejectFixedBuilder {
	return HistoricalOrderFillsRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalOrderFillsRejectPointer) ToBuilder() HistoricalOrderFillsRejectPointerBuilder {
	return HistoricalOrderFillsRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalOrderFillsRejectFixedPointer) ToBuilder() HistoricalOrderFillsRejectFixedPointerBuilder {
	return HistoricalOrderFillsRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalOrderFillsRejectBuilder) Finish() HistoricalOrderFillsReject {
	return HistoricalOrderFillsReject{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalOrderFillsRejectFixedBuilder) Finish() HistoricalOrderFillsRejectFixed {
	return HistoricalOrderFillsRejectFixed{m.Fixed}
}

// Finish
func (m *HistoricalOrderFillsRejectPointerBuilder) Finish() HistoricalOrderFillsRejectPointer {
	return HistoricalOrderFillsRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalOrderFillsRejectFixedPointerBuilder) Finish() HistoricalOrderFillsRejectFixedPointer {
	return HistoricalOrderFillsRejectFixedPointer{m.FixedPointer}
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m *HistoricalOrderFillsRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m *HistoricalOrderFillsRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m HistoricalOrderFillsReject) Copy(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalOrderFillsRejectBuilder) Copy(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsReject) CopyPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectBuilder) CopyPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsReject) CopyToPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectBuilder) CopyToPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsReject) CopyTo(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectBuilder) CopyTo(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalOrderFillsRejectFixed) Copy(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalOrderFillsRejectFixedBuilder) Copy(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectFixed) CopyPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectFixedBuilder) CopyPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectFixed) CopyToPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectFixedBuilder) CopyToPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectFixed) CopyTo(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectFixedBuilder) CopyTo(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalOrderFillsRejectPointer) Copy(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalOrderFillsRejectPointerBuilder) Copy(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectPointer) CopyPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectPointerBuilder) CopyPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectPointer) CopyTo(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectPointerBuilder) CopyTo(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectPointer) CopyToPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectPointerBuilder) CopyToPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalOrderFillsRejectFixedPointer) Copy(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalOrderFillsRejectFixedPointerBuilder) Copy(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectFixedPointer) CopyPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectFixedPointerBuilder) CopyPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectFixedPointer) CopyTo(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectFixedPointerBuilder) CopyTo(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectFixedPointer) CopyToPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectFixedPointerBuilder) CopyToPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
