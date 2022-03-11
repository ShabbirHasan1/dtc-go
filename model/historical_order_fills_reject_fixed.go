package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = HistoricalOrderFillsRejectFixedSize  (104)
//     Type       = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     RequestID  = 0
//     RejectText = ""
// }
var _HistoricalOrderFillsRejectFixedDefault = []byte{104, 0, 52, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalOrderFillsRejectFixedSize = 104

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

func NewHistoricalOrderFillsRejectFixedFrom(b []byte) HistoricalOrderFillsRejectFixed {
	return HistoricalOrderFillsRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalOrderFillsRejectFixed(b []byte) HistoricalOrderFillsRejectFixed {
	return HistoricalOrderFillsRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalOrderFillsRejectFixed() HistoricalOrderFillsRejectFixedBuilder {
	a := HistoricalOrderFillsRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _HistoricalOrderFillsRejectFixedDefault)
	return a
}

// Clear
// {
//     Size       = HistoricalOrderFillsRejectFixedSize  (104)
//     Type       = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalOrderFillsRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalOrderFillsRejectFixedDefault)
}

// ToBuilder
func (m HistoricalOrderFillsRejectFixed) ToBuilder() HistoricalOrderFillsRejectFixedBuilder {
	return HistoricalOrderFillsRejectFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalOrderFillsRejectFixedBuilder) Finish() HistoricalOrderFillsRejectFixed {
	return HistoricalOrderFillsRejectFixed{m.Fixed}
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

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// SetRequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// Copy
func (m HistoricalOrderFillsRejectFixed) Copy(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalOrderFillsRejectFixedBuilder) Copy(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectFixed) CopyPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectFixedBuilder) CopyPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectFixed) CopyToPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectFixedBuilder) CopyToPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectFixed) CopyTo(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectFixedBuilder) CopyTo(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
