package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = OpenOrdersRejectSize  (16)
//     Type       = OPEN_ORDERS_REJECT  (302)
//     BaseSize   = OpenOrdersRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _OpenOrdersRejectDefault = []byte{16, 0, 46, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const OpenOrdersRejectSize = 16

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

// ToBuilder
func (m OpenOrdersReject) ToBuilder() OpenOrdersRejectBuilder {
	return OpenOrdersRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m OpenOrdersRejectBuilder) Finish() OpenOrdersReject {
	return OpenOrdersReject{m.VLSBuilder.Finish()}
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

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetRequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m *OpenOrdersRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
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
