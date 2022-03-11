package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = CurrentPositionsRejectSize  (16)
//     Type       = CURRENT_POSITIONS_REJECT  (307)
//     BaseSize   = CurrentPositionsRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _CurrentPositionsRejectDefault = []byte{16, 0, 51, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const CurrentPositionsRejectSize = 16

// CurrentPositionsReject If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsReject struct {
	message.VLS
}

// CurrentPositionsRejectBuilder If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectBuilder struct {
	message.VLSBuilder
}

func NewCurrentPositionsRejectFrom(b []byte) CurrentPositionsReject {
	return CurrentPositionsReject{VLS: message.NewVLSFrom(b)}
}

func WrapCurrentPositionsReject(b []byte) CurrentPositionsReject {
	return CurrentPositionsReject{VLS: message.WrapVLS(b)}
}

func NewCurrentPositionsReject() CurrentPositionsRejectBuilder {
	a := CurrentPositionsRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _CurrentPositionsRejectDefault)
	return a
}

// Clear
// {
//     Size       = CurrentPositionsRejectSize  (16)
//     Type       = CURRENT_POSITIONS_REJECT  (307)
//     BaseSize   = CurrentPositionsRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m CurrentPositionsRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CurrentPositionsRejectDefault)
}

// ToBuilder
func (m CurrentPositionsReject) ToBuilder() CurrentPositionsRejectBuilder {
	return CurrentPositionsRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m CurrentPositionsRejectBuilder) Finish() CurrentPositionsReject {
	return CurrentPositionsReject{m.VLSBuilder.Finish()}
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetRequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText Free-form text indicating the reason for the rejection.
//
func (m *CurrentPositionsRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// Copy
func (m CurrentPositionsReject) Copy(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m CurrentPositionsRejectBuilder) Copy(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsReject) CopyPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectBuilder) CopyPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsReject) CopyToPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectBuilder) CopyToPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsReject) CopyTo(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectBuilder) CopyTo(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
