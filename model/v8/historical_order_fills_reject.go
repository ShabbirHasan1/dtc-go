// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-16 08:00:54.519198 +0800 WITA m=+0.055446543

package v8

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalOrderFillsRejectSize = 16

const HistoricalOrderFillsRejectFixedSize = 104

//     Size        uint16  = HistoricalOrderFillsRejectSize  (16)
//     Type        uint16  = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     BaseSize    uint16  = HistoricalOrderFillsRejectSize  (16)
//     RequestID   int32   = 0
//     RejectText  string  = ""
var _HistoricalOrderFillsRejectDefault = []byte{16, 0, 52, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size        uint16      = HistoricalOrderFillsRejectFixedSize  (104)
//     Type        uint16      = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     RequestID   int32       = 0
//     RejectText  string[96]  = ""
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
	return HistoricalOrderFillsReject{VLS: message.NewVLS(b)}
}

func WrapHistoricalOrderFillsReject(b []byte) HistoricalOrderFillsReject {
	return HistoricalOrderFillsReject{VLS: message.WrapVLS(b)}
}

func NewHistoricalOrderFillsReject() HistoricalOrderFillsRejectBuilder {
	return HistoricalOrderFillsRejectBuilder{message.NewVLSBuilder(_HistoricalOrderFillsRejectDefault)}
}

func NewHistoricalOrderFillsRejectFixedFrom(b []byte) HistoricalOrderFillsRejectFixed {
	return HistoricalOrderFillsRejectFixed{Fixed: message.NewFixed(b)}
}

func WrapHistoricalOrderFillsRejectFixed(b []byte) HistoricalOrderFillsRejectFixed {
	return HistoricalOrderFillsRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalOrderFillsRejectFixed() HistoricalOrderFillsRejectFixedBuilder {
	return HistoricalOrderFillsRejectFixedBuilder{message.NewFixed(_HistoricalOrderFillsRejectFixedDefault)}
}

func AllocHistoricalOrderFillsReject() HistoricalOrderFillsRejectPointerBuilder {
	return HistoricalOrderFillsRejectPointerBuilder{message.AllocVLSBuilder(_HistoricalOrderFillsRejectDefault)}
}

func AllocHistoricalOrderFillsRejectFrom(b []byte) HistoricalOrderFillsRejectPointer {
	return HistoricalOrderFillsRejectPointer{VLSPointer: message.AllocVLS(b)}
}

func AllocHistoricalOrderFillsRejectFixed() HistoricalOrderFillsRejectFixedPointerBuilder {
	return HistoricalOrderFillsRejectFixedPointerBuilder{message.AllocFixed(_HistoricalOrderFillsRejectFixedDefault)}
}

func AllocHistoricalOrderFillsRejectFixedFrom(b []byte) HistoricalOrderFillsRejectFixedPointer {
	return HistoricalOrderFillsRejectFixedPointer{FixedPointer: message.AllocFixed(b)}
}

// Clear
//     Size        uint16  = HistoricalOrderFillsRejectSize  (16)
//     Type        uint16  = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     BaseSize    uint16  = HistoricalOrderFillsRejectSize  (16)
//     RequestID   int32   = 0
//     RejectText  string  = ""
func (m HistoricalOrderFillsRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalOrderFillsRejectDefault)
}

// Clear
//     Size        uint16      = HistoricalOrderFillsRejectFixedSize  (104)
//     Type        uint16      = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     RequestID   int32       = 0
//     RejectText  string[96]  = ""
func (m HistoricalOrderFillsRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalOrderFillsRejectFixedDefault)
}

// Clear
//     Size        uint16  = HistoricalOrderFillsRejectSize  (16)
//     Type        uint16  = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     BaseSize    uint16  = HistoricalOrderFillsRejectSize  (16)
//     RequestID   int32   = 0
//     RejectText  string  = ""
func (m HistoricalOrderFillsRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalOrderFillsRejectDefault)
}

// Clear
//     Size        uint16      = HistoricalOrderFillsRejectFixedSize  (104)
//     Type        uint16      = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     RequestID   int32       = 0
//     RejectText  string[96]  = ""
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