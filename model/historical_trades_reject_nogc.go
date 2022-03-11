package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalTradesRejectPointer struct {
	message.VLSPointer
}

type HistoricalTradesRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocHistoricalTradesReject() HistoricalTradesRejectPointerBuilder {
	a := HistoricalTradesRejectPointerBuilder{message.AllocVLSBuilder(14)}
	a.Ptr.SetBytes(0, _HistoricalTradesRejectDefault)
	return a
}

func AllocHistoricalTradesRejectFrom(b []byte) HistoricalTradesRejectPointer {
	return HistoricalTradesRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size       = HistoricalTradesRejectSize  (14)
//     Type       = HISTORICAL_TRADES_REJECT  (10101)
//     BaseSize   = HistoricalTradesRejectSize  (14)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalTradesRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalTradesRejectDefault)
}

// ToBuilder
func (m HistoricalTradesRejectPointer) ToBuilder() HistoricalTradesRejectPointerBuilder {
	return HistoricalTradesRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalTradesRejectPointerBuilder) Finish() HistoricalTradesRejectPointer {
	return HistoricalTradesRejectPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m HistoricalTradesRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m HistoricalTradesRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// RejectText
func (m HistoricalTradesRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// RejectText
func (m HistoricalTradesRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// SetRequestID
func (m HistoricalTradesRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 10, 6, value)
}

// SetRejectText
func (m *HistoricalTradesRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// Copy
func (m HistoricalTradesRejectPointer) Copy(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalTradesRejectPointerBuilder) Copy(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectPointer) CopyPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectPointerBuilder) CopyPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectPointer) CopyTo(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectPointerBuilder) CopyTo(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectPointer) CopyToPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectPointerBuilder) CopyToPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
