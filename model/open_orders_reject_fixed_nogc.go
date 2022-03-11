package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
//     Size       = OpenOrdersRejectFixedSize  (104)
//     Type       = OPEN_ORDERS_REJECT  (302)
//     RequestID  = 0
//     RejectText = ""
// }
func (m OpenOrdersRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _OpenOrdersRejectFixedDefault)
}

// ToBuilder
func (m OpenOrdersRejectFixedPointer) ToBuilder() OpenOrdersRejectFixedPointerBuilder {
	return OpenOrdersRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *OpenOrdersRejectFixedPointerBuilder) Finish() OpenOrdersRejectFixedPointer {
	return OpenOrdersRejectFixedPointer{m.FixedPointer}
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
func (m OpenOrdersRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
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
