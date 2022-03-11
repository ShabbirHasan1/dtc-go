package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = HistoricalOrderFillsRejectSize  (16)
//     Type       = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     BaseSize   = HistoricalOrderFillsRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _HistoricalOrderFillsRejectDefault = []byte{16, 0, 52, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalOrderFillsRejectSize = 16

// HistoricalOrderFillsReject If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsReject struct {
	message.VLS
}

// HistoricalOrderFillsRejectBuilder If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsRejectBuilder struct {
	message.VLSBuilder
}

func NewHistoricalOrderFillsRejectFrom(b []byte) HistoricalOrderFillsReject {
	return HistoricalOrderFillsReject{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalOrderFillsReject(b []byte) HistoricalOrderFillsReject {
	return HistoricalOrderFillsReject{VLS: message.WrapVLS(b)}
}

func NewHistoricalOrderFillsReject() HistoricalOrderFillsRejectBuilder {
	a := HistoricalOrderFillsRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _HistoricalOrderFillsRejectDefault)
	return a
}

// Clear
// {
//     Size       = HistoricalOrderFillsRejectSize  (16)
//     Type       = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     BaseSize   = HistoricalOrderFillsRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalOrderFillsRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalOrderFillsRejectDefault)
}

// ToBuilder
func (m HistoricalOrderFillsReject) ToBuilder() HistoricalOrderFillsRejectBuilder {
	return HistoricalOrderFillsRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m HistoricalOrderFillsRejectBuilder) Finish() HistoricalOrderFillsReject {
	return HistoricalOrderFillsReject{m.VLSBuilder.Finish()}
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetRequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m *HistoricalOrderFillsRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// Copy
func (m HistoricalOrderFillsReject) Copy(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalOrderFillsRejectBuilder) Copy(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsReject) CopyPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectBuilder) CopyPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsReject) CopyToPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectBuilder) CopyToPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsReject) CopyTo(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectBuilder) CopyTo(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
