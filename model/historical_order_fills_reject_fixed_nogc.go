package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HistoricalOrderFillsRejectFixedPointer If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsRejectFixedPointer struct {
	message.FixedPointer
}

// HistoricalOrderFillsRejectFixedPointerBuilder If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
type HistoricalOrderFillsRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalOrderFillsRejectFixed() HistoricalOrderFillsRejectFixedPointerBuilder {
	a := HistoricalOrderFillsRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _HistoricalOrderFillsRejectFixedDefault)
	return a
}

func AllocHistoricalOrderFillsRejectFixedFrom(b []byte) HistoricalOrderFillsRejectFixedPointer {
	return HistoricalOrderFillsRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = HistoricalOrderFillsRejectFixedSize  (104)
//     Type       = HISTORICAL_ORDER_FILLS_REJECT  (308)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalOrderFillsRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalOrderFillsRejectFixedDefault)
}

// ToBuilder
func (m HistoricalOrderFillsRejectFixedPointer) ToBuilder() HistoricalOrderFillsRejectFixedPointerBuilder {
	return HistoricalOrderFillsRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalOrderFillsRejectFixedPointerBuilder) Finish() HistoricalOrderFillsRejectFixedPointer {
	return HistoricalOrderFillsRejectFixedPointer{m.FixedPointer}
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID This is set to the RequestID field sent in the HistoricalOrderFillsRequest
// message.
func (m HistoricalOrderFillsRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Free-form text indicating the reason for rejection.
func (m HistoricalOrderFillsRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m HistoricalOrderFillsRejectFixedPointer) Copy(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalOrderFillsRejectFixedPointerBuilder) Copy(to HistoricalOrderFillsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectFixedPointer) CopyPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalOrderFillsRejectFixedPointerBuilder) CopyPointer(to HistoricalOrderFillsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectFixedPointer) CopyTo(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalOrderFillsRejectFixedPointerBuilder) CopyTo(to HistoricalOrderFillsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectFixedPointer) CopyToPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalOrderFillsRejectFixedPointerBuilder) CopyToPointer(to HistoricalOrderFillsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
