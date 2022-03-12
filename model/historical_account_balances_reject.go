package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalAccountBalancesRejectSize = 16

const HistoricalAccountBalancesRejectFixedSize = 104

// {
//     Size       = HistoricalAccountBalancesRejectSize  (16)
//     Type       = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     BaseSize   = HistoricalAccountBalancesRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _HistoricalAccountBalancesRejectDefault = []byte{16, 0, 92, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size       = HistoricalAccountBalancesRejectFixedSize  (104)
//     Type       = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     RequestID  = 0
//     RejectText = ""
// }
var _HistoricalAccountBalancesRejectFixedDefault = []byte{104, 0, 92, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// HistoricalAccountBalancesReject This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesReject struct {
	message.VLS
}

// HistoricalAccountBalancesRejectBuilder This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectBuilder struct {
	message.VLSBuilder
}

// HistoricalAccountBalancesRejectFixed This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectFixed struct {
	message.Fixed
}

// HistoricalAccountBalancesRejectFixedBuilder This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectFixedBuilder struct {
	message.Fixed
}

// HistoricalAccountBalancesRejectPointer This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectPointer struct {
	message.VLSPointer
}

// HistoricalAccountBalancesRejectPointerBuilder This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

// HistoricalAccountBalancesRejectFixedPointer This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectFixedPointer struct {
	message.FixedPointer
}

// HistoricalAccountBalancesRejectFixedPointerBuilder This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalAccountBalancesRejectFrom(b []byte) HistoricalAccountBalancesReject {
	return HistoricalAccountBalancesReject{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalAccountBalancesReject(b []byte) HistoricalAccountBalancesReject {
	return HistoricalAccountBalancesReject{VLS: message.WrapVLS(b)}
}

func NewHistoricalAccountBalancesReject() HistoricalAccountBalancesRejectBuilder {
	a := HistoricalAccountBalancesRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _HistoricalAccountBalancesRejectDefault)
	return a
}

func NewHistoricalAccountBalancesRejectFixedFrom(b []byte) HistoricalAccountBalancesRejectFixed {
	return HistoricalAccountBalancesRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalAccountBalancesRejectFixed(b []byte) HistoricalAccountBalancesRejectFixed {
	return HistoricalAccountBalancesRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalAccountBalancesRejectFixed() HistoricalAccountBalancesRejectFixedBuilder {
	a := HistoricalAccountBalancesRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _HistoricalAccountBalancesRejectFixedDefault)
	return a
}

func AllocHistoricalAccountBalancesReject() HistoricalAccountBalancesRejectPointerBuilder {
	a := HistoricalAccountBalancesRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _HistoricalAccountBalancesRejectDefault)
	return a
}

func AllocHistoricalAccountBalancesRejectFrom(b []byte) HistoricalAccountBalancesRejectPointer {
	return HistoricalAccountBalancesRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocHistoricalAccountBalancesRejectFixed() HistoricalAccountBalancesRejectFixedPointerBuilder {
	a := HistoricalAccountBalancesRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _HistoricalAccountBalancesRejectFixedDefault)
	return a
}

func AllocHistoricalAccountBalancesRejectFixedFrom(b []byte) HistoricalAccountBalancesRejectFixedPointer {
	return HistoricalAccountBalancesRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = HistoricalAccountBalancesRejectSize  (16)
//     Type       = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     BaseSize   = HistoricalAccountBalancesRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalAccountBalancesRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalAccountBalancesRejectDefault)
}

// Clear
// {
//     Size       = HistoricalAccountBalancesRejectFixedSize  (104)
//     Type       = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalAccountBalancesRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalAccountBalancesRejectFixedDefault)
}

// Clear
// {
//     Size       = HistoricalAccountBalancesRejectSize  (16)
//     Type       = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     BaseSize   = HistoricalAccountBalancesRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalAccountBalancesRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalAccountBalancesRejectDefault)
}

// Clear
// {
//     Size       = HistoricalAccountBalancesRejectFixedSize  (104)
//     Type       = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalAccountBalancesRejectFixedDefault)
}

// ToBuilder
func (m HistoricalAccountBalancesReject) ToBuilder() HistoricalAccountBalancesRejectBuilder {
	return HistoricalAccountBalancesRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalAccountBalancesRejectFixed) ToBuilder() HistoricalAccountBalancesRejectFixedBuilder {
	return HistoricalAccountBalancesRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalAccountBalancesRejectPointer) ToBuilder() HistoricalAccountBalancesRejectPointerBuilder {
	return HistoricalAccountBalancesRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalAccountBalancesRejectFixedPointer) ToBuilder() HistoricalAccountBalancesRejectFixedPointerBuilder {
	return HistoricalAccountBalancesRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalAccountBalancesRejectBuilder) Finish() HistoricalAccountBalancesReject {
	return HistoricalAccountBalancesReject{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalAccountBalancesRejectFixedBuilder) Finish() HistoricalAccountBalancesRejectFixed {
	return HistoricalAccountBalancesRejectFixed{m.Fixed}
}

// Finish
func (m *HistoricalAccountBalancesRejectPointerBuilder) Finish() HistoricalAccountBalancesRejectPointer {
	return HistoricalAccountBalancesRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalAccountBalancesRejectFixedPointerBuilder) Finish() HistoricalAccountBalancesRejectFixedPointer {
	return HistoricalAccountBalancesRejectFixedPointer{m.FixedPointer}
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m *HistoricalAccountBalancesRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m *HistoricalAccountBalancesRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m HistoricalAccountBalancesReject) Copy(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalAccountBalancesRejectBuilder) Copy(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesReject) CopyPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectBuilder) CopyPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesReject) CopyToPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectBuilder) CopyToPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesReject) CopyTo(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectBuilder) CopyTo(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalAccountBalancesRejectFixed) Copy(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalAccountBalancesRejectFixedBuilder) Copy(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectFixed) CopyPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectFixedBuilder) CopyPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectFixed) CopyToPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectFixedBuilder) CopyToPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectFixed) CopyTo(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectFixedBuilder) CopyTo(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalAccountBalancesRejectPointer) Copy(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalAccountBalancesRejectPointerBuilder) Copy(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectPointer) CopyPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectPointerBuilder) CopyPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectPointer) CopyTo(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectPointerBuilder) CopyTo(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectPointer) CopyToPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectPointerBuilder) CopyToPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalAccountBalancesRejectFixedPointer) Copy(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) Copy(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectFixedPointer) CopyPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) CopyPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectFixedPointer) CopyTo(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) CopyTo(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectFixedPointer) CopyToPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) CopyToPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
