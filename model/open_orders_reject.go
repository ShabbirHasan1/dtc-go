package model

import (
	"github.com/moontrade/dtc-go/message"
)

const OpenOrdersRejectSize = 16

const OpenOrdersRejectFixedSize = 104

// {
//     Size       = OpenOrdersRejectSize  (16)
//     Type       = OPEN_ORDERS_REJECT  (302)
//     BaseSize   = OpenOrdersRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _OpenOrdersRejectDefault = []byte{16, 0, 46, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size       = OpenOrdersRejectFixedSize  (104)
//     Type       = OPEN_ORDERS_REJECT  (302)
//     RequestID  = 0
//     RejectText = ""
// }
var _OpenOrdersRejectFixedDefault = []byte{104, 0, 46, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// OpenOrdersReject If the Server is unable to serve the request for an OpenOrdersRequest
// message received, for a reason other than there not being any open orders,
// then send this message to the Client.
type OpenOrdersReject struct {
	message.VLS
}

// OpenOrdersRejectBuilder If the Server is unable to serve the request for an OpenOrdersRequest
// message received, for a reason other than there not being any open orders,
// then send this message to the Client.
type OpenOrdersRejectBuilder struct {
	message.VLSBuilder
}

// OpenOrdersRejectFixed If the Server is unable to serve the request for an OpenOrdersRequest
// message received, for a reason other than there not being any open orders,
// then send this message to the Client.
type OpenOrdersRejectFixed struct {
	message.Fixed
}

// OpenOrdersRejectFixedBuilder If the Server is unable to serve the request for an OpenOrdersRequest
// message received, for a reason other than there not being any open orders,
// then send this message to the Client.
type OpenOrdersRejectFixedBuilder struct {
	message.Fixed
}

// OpenOrdersRejectPointer If the Server is unable to serve the request for an OpenOrdersRequest
// message received, for a reason other than there not being any open orders,
// then send this message to the Client.
type OpenOrdersRejectPointer struct {
	message.VLSPointer
}

// OpenOrdersRejectPointerBuilder If the Server is unable to serve the request for an OpenOrdersRequest
// message received, for a reason other than there not being any open orders,
// then send this message to the Client.
type OpenOrdersRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

// OpenOrdersRejectFixedPointer If the Server is unable to serve the request for an OpenOrdersRequest
// message received, for a reason other than there not being any open orders,
// then send this message to the Client.
type OpenOrdersRejectFixedPointer struct {
	message.FixedPointer
}

// OpenOrdersRejectFixedPointerBuilder If the Server is unable to serve the request for an OpenOrdersRequest
// message received, for a reason other than there not being any open orders,
// then send this message to the Client.
type OpenOrdersRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func NewOpenOrdersRejectFrom(b []byte) OpenOrdersReject {
	return OpenOrdersReject{VLS: message.NewVLSFrom(b)}
}

func WrapOpenOrdersReject(b []byte) OpenOrdersReject {
	return OpenOrdersReject{VLS: message.WrapVLS(b)}
}

func NewOpenOrdersReject() OpenOrdersRejectBuilder {
	a := OpenOrdersRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _OpenOrdersRejectDefault)
	return a
}

func NewOpenOrdersRejectFixedFrom(b []byte) OpenOrdersRejectFixed {
	return OpenOrdersRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapOpenOrdersRejectFixed(b []byte) OpenOrdersRejectFixed {
	return OpenOrdersRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewOpenOrdersRejectFixed() OpenOrdersRejectFixedBuilder {
	a := OpenOrdersRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _OpenOrdersRejectFixedDefault)
	return a
}

func AllocOpenOrdersReject() OpenOrdersRejectPointerBuilder {
	a := OpenOrdersRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _OpenOrdersRejectDefault)
	return a
}

func AllocOpenOrdersRejectFrom(b []byte) OpenOrdersRejectPointer {
	return OpenOrdersRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocOpenOrdersRejectFixed() OpenOrdersRejectFixedPointerBuilder {
	a := OpenOrdersRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _OpenOrdersRejectFixedDefault)
	return a
}

func AllocOpenOrdersRejectFixedFrom(b []byte) OpenOrdersRejectFixedPointer {
	return OpenOrdersRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = OpenOrdersRejectSize  (16)
//     Type       = OPEN_ORDERS_REJECT  (302)
//     BaseSize   = OpenOrdersRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m OpenOrdersRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _OpenOrdersRejectDefault)
}

// Clear
// {
//     Size       = OpenOrdersRejectFixedSize  (104)
//     Type       = OPEN_ORDERS_REJECT  (302)
//     RequestID  = 0
//     RejectText = ""
// }
func (m OpenOrdersRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _OpenOrdersRejectFixedDefault)
}

// Clear
// {
//     Size       = OpenOrdersRejectSize  (16)
//     Type       = OPEN_ORDERS_REJECT  (302)
//     BaseSize   = OpenOrdersRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m OpenOrdersRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _OpenOrdersRejectDefault)
}

// Clear
// {
//     Size       = OpenOrdersRejectFixedSize  (104)
//     Type       = OPEN_ORDERS_REJECT  (302)
//     RequestID  = 0
//     RejectText = ""
// }
func (m OpenOrdersRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _OpenOrdersRejectFixedDefault)
}

// ToBuilder
func (m OpenOrdersReject) ToBuilder() OpenOrdersRejectBuilder {
	return OpenOrdersRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m OpenOrdersRejectFixed) ToBuilder() OpenOrdersRejectFixedBuilder {
	return OpenOrdersRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m OpenOrdersRejectPointer) ToBuilder() OpenOrdersRejectPointerBuilder {
	return OpenOrdersRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m OpenOrdersRejectFixedPointer) ToBuilder() OpenOrdersRejectFixedPointerBuilder {
	return OpenOrdersRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m OpenOrdersRejectBuilder) Finish() OpenOrdersReject {
	return OpenOrdersReject{m.VLSBuilder.Finish()}
}

// Finish
func (m OpenOrdersRejectFixedBuilder) Finish() OpenOrdersRejectFixed {
	return OpenOrdersRejectFixed{m.Fixed}
}

// Finish
func (m *OpenOrdersRejectPointerBuilder) Finish() OpenOrdersRejectPointer {
	return OpenOrdersRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *OpenOrdersRejectFixedPointerBuilder) Finish() OpenOrdersRejectFixedPointer {
	return OpenOrdersRejectFixedPointer{m.FixedPointer}
}

// RequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m *OpenOrdersRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m *OpenOrdersRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m OpenOrdersReject) Copy(to OpenOrdersRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m OpenOrdersRejectBuilder) Copy(to OpenOrdersRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m OpenOrdersReject) CopyPointer(to OpenOrdersRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m OpenOrdersRejectBuilder) CopyPointer(to OpenOrdersRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m OpenOrdersReject) CopyToPointer(to OpenOrdersRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m OpenOrdersRejectBuilder) CopyToPointer(to OpenOrdersRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m OpenOrdersReject) CopyTo(to OpenOrdersRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m OpenOrdersRejectBuilder) CopyTo(to OpenOrdersRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m OpenOrdersRejectFixed) Copy(to OpenOrdersRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m OpenOrdersRejectFixedBuilder) Copy(to OpenOrdersRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m OpenOrdersRejectFixed) CopyPointer(to OpenOrdersRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m OpenOrdersRejectFixedBuilder) CopyPointer(to OpenOrdersRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m OpenOrdersRejectFixed) CopyToPointer(to OpenOrdersRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m OpenOrdersRejectFixedBuilder) CopyToPointer(to OpenOrdersRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m OpenOrdersRejectFixed) CopyTo(to OpenOrdersRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m OpenOrdersRejectFixedBuilder) CopyTo(to OpenOrdersRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m OpenOrdersRejectPointer) Copy(to OpenOrdersRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m OpenOrdersRejectPointerBuilder) Copy(to OpenOrdersRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m OpenOrdersRejectPointer) CopyPointer(to OpenOrdersRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m OpenOrdersRejectPointerBuilder) CopyPointer(to OpenOrdersRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m OpenOrdersRejectPointer) CopyTo(to OpenOrdersRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m OpenOrdersRejectPointerBuilder) CopyTo(to OpenOrdersRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m OpenOrdersRejectPointer) CopyToPointer(to OpenOrdersRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m OpenOrdersRejectPointerBuilder) CopyToPointer(to OpenOrdersRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m OpenOrdersRejectFixedPointer) Copy(to OpenOrdersRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m OpenOrdersRejectFixedPointerBuilder) Copy(to OpenOrdersRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m OpenOrdersRejectFixedPointer) CopyPointer(to OpenOrdersRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m OpenOrdersRejectFixedPointerBuilder) CopyPointer(to OpenOrdersRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m OpenOrdersRejectFixedPointer) CopyTo(to OpenOrdersRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m OpenOrdersRejectFixedPointerBuilder) CopyTo(to OpenOrdersRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m OpenOrdersRejectFixedPointer) CopyToPointer(to OpenOrdersRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m OpenOrdersRejectFixedPointerBuilder) CopyToPointer(to OpenOrdersRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
