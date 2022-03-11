package model

import (
	"github.com/moontrade/dtc-go/message"
)

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

func AllocOpenOrdersReject() OpenOrdersRejectPointerBuilder {
	a := OpenOrdersRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _OpenOrdersRejectDefault)
	return a
}

func AllocOpenOrdersRejectFrom(b []byte) OpenOrdersRejectPointer {
	return OpenOrdersRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
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

// ToBuilder
func (m OpenOrdersRejectPointer) ToBuilder() OpenOrdersRejectPointerBuilder {
	return OpenOrdersRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *OpenOrdersRejectPointerBuilder) Finish() OpenOrdersRejectPointer {
	return OpenOrdersRejectPointer{m.VLSPointerBuilder.Finish()}
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
func (m OpenOrdersRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetRequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m *OpenOrdersRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
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
