package model

import (
	"github.com/moontrade/dtc-go/message"
)

// CurrentPositionsRejectPointer If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectPointer struct {
	message.VLSPointer
}

// CurrentPositionsRejectPointerBuilder If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocCurrentPositionsReject() CurrentPositionsRejectPointerBuilder {
	a := CurrentPositionsRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _CurrentPositionsRejectDefault)
	return a
}

func AllocCurrentPositionsRejectFrom(b []byte) CurrentPositionsRejectPointer {
	return CurrentPositionsRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size       = CurrentPositionsRejectSize  (16)
//     Type       = CURRENT_POSITIONS_REJECT  (307)
//     BaseSize   = CurrentPositionsRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m CurrentPositionsRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CurrentPositionsRejectDefault)
}

// ToBuilder
func (m CurrentPositionsRejectPointer) ToBuilder() CurrentPositionsRejectPointerBuilder {
	return CurrentPositionsRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *CurrentPositionsRejectPointerBuilder) Finish() CurrentPositionsRejectPointer {
	return CurrentPositionsRejectPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetRequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Free-form text indicating the reason for the rejection.
//
func (m *CurrentPositionsRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// Copy
func (m CurrentPositionsRejectPointer) Copy(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m CurrentPositionsRejectPointerBuilder) Copy(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectPointer) CopyPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectPointerBuilder) CopyPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectPointer) CopyTo(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectPointerBuilder) CopyTo(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectPointer) CopyToPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectPointerBuilder) CopyToPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
