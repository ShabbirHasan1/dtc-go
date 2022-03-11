package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HistoricalOrderFillsRejectPointer If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsRejectPointer struct {
	message.VLSPointer
}

// HistoricalOrderFillsRejectPointerBuilder If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocHistoricalOrderFillsReject() HistoricalOrderFillsRejectPointerBuilder {
	a := HistoricalOrderFillsRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _HistoricalOrderFillsRejectDefault)
	return a
}

func AllocHistoricalOrderFillsRejectFrom(b []byte) HistoricalOrderFillsRejectPointer {
	return HistoricalOrderFillsRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size       = HistoricalOrderFillsRejectSize  (16)
//     Type       = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     BaseSize   = HistoricalOrderFillsRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalOrderFillsRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalOrderFillsRejectDefault)
}

// ToBuilder
func (m HistoricalOrderFillsRejectPointer) ToBuilder() HistoricalOrderFillsRejectPointerBuilder {
	return HistoricalOrderFillsRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalOrderFillsRejectPointerBuilder) Finish() HistoricalOrderFillsRejectPointer {
	return HistoricalOrderFillsRejectPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetRequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m *HistoricalOrderFillsRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// Copy
func (m HistoricalOrderFillsRejectPointer) Copy(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalOrderFillsRejectPointerBuilder) Copy(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectPointer) CopyPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectPointerBuilder) CopyPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectPointer) CopyTo(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectPointerBuilder) CopyTo(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectPointer) CopyToPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectPointerBuilder) CopyToPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
