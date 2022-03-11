package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = OpenOrdersRejectFixedSize  (104)
//     Type       = OPEN_ORDERS_REJECT  (302)
//     RequestID  = 0
//     RejectText = ""
// }
var _OpenOrdersRejectFixedDefault = []byte{104, 0, 46, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const OpenOrdersRejectFixedSize = 104

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

// ToBuilder
func (m OpenOrdersRejectFixed) ToBuilder() OpenOrdersRejectFixedBuilder {
	return OpenOrdersRejectFixedBuilder{m.Fixed}
}

// Finish
func (m OpenOrdersRejectFixedBuilder) Finish() OpenOrdersRejectFixed {
	return OpenOrdersRejectFixed{m.Fixed}
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

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// SetRequestID This is set to the RequestID field sent in the OpenOrdersRequest message
// from the Client.
func (m OpenOrdersRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m OpenOrdersRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
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
