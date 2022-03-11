package model

import (
	"github.com/moontrade/dtc-go/message"
)

// CurrentPositionsRejectFixedPointer If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectFixedPointer struct {
	message.FixedPointer
}

// CurrentPositionsRejectFixedPointerBuilder If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocCurrentPositionsRejectFixed() CurrentPositionsRejectFixedPointerBuilder {
	a := CurrentPositionsRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _CurrentPositionsRejectFixedDefault)
	return a
}

func AllocCurrentPositionsRejectFixedFrom(b []byte) CurrentPositionsRejectFixedPointer {
	return CurrentPositionsRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = CurrentPositionsRejectFixedSize  (104)
//     Type       = CURRENT_POSITIONS_REJECT  (307)
//     RequestID  = 0
//     RejectText = ""
// }
func (m CurrentPositionsRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CurrentPositionsRejectFixedDefault)
}

// ToBuilder
func (m CurrentPositionsRejectFixedPointer) ToBuilder() CurrentPositionsRejectFixedPointerBuilder {
	return CurrentPositionsRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *CurrentPositionsRejectFixedPointerBuilder) Finish() CurrentPositionsRejectFixedPointer {
	return CurrentPositionsRejectFixedPointer{m.FixedPointer}
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m CurrentPositionsRejectFixedPointer) Copy(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m CurrentPositionsRejectFixedPointerBuilder) Copy(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectFixedPointer) CopyPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectFixedPointerBuilder) CopyPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectFixedPointer) CopyTo(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectFixedPointerBuilder) CopyTo(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectFixedPointer) CopyToPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectFixedPointerBuilder) CopyToPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
