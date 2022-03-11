package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = CurrentPositionsRejectFixedSize  (104)
//     Type       = CURRENT_POSITIONS_REJECT  (307)
//     RequestID  = 0
//     RejectText = ""
// }
var _CurrentPositionsRejectFixedDefault = []byte{104, 0, 51, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const CurrentPositionsRejectFixedSize = 104

// CurrentPositionsRejectFixed If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectFixed struct {
	message.Fixed
}

// CurrentPositionsRejectFixedBuilder If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectFixedBuilder struct {
	message.Fixed
}

func NewCurrentPositionsRejectFixedFrom(b []byte) CurrentPositionsRejectFixed {
	return CurrentPositionsRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapCurrentPositionsRejectFixed(b []byte) CurrentPositionsRejectFixed {
	return CurrentPositionsRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewCurrentPositionsRejectFixed() CurrentPositionsRejectFixedBuilder {
	a := CurrentPositionsRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _CurrentPositionsRejectFixedDefault)
	return a
}

// Clear
// {
//     Size       = CurrentPositionsRejectFixedSize  (104)
//     Type       = CURRENT_POSITIONS_REJECT  (307)
//     RequestID  = 0
//     RejectText = ""
// }
func (m CurrentPositionsRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CurrentPositionsRejectFixedDefault)
}

// ToBuilder
func (m CurrentPositionsRejectFixed) ToBuilder() CurrentPositionsRejectFixedBuilder {
	return CurrentPositionsRejectFixedBuilder{m.Fixed}
}

// Finish
func (m CurrentPositionsRejectFixedBuilder) Finish() CurrentPositionsRejectFixed {
	return CurrentPositionsRejectFixed{m.Fixed}
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// SetRequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// Copy
func (m CurrentPositionsRejectFixed) Copy(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m CurrentPositionsRejectFixedBuilder) Copy(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectFixed) CopyPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectFixedBuilder) CopyPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectFixed) CopyToPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectFixedBuilder) CopyToPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectFixed) CopyTo(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectFixedBuilder) CopyTo(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
