package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalTradesRejectFixedPointer struct {
	message.FixedPointer
}

type HistoricalTradesRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalTradesRejectFixed() HistoricalTradesRejectFixedPointerBuilder {
	a := HistoricalTradesRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _HistoricalTradesRejectFixedDefault)
	return a
}

func AllocHistoricalTradesRejectFixedFrom(b []byte) HistoricalTradesRejectFixedPointer {
	return HistoricalTradesRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = HistoricalTradesRejectFixedSize  (104)
//     Type       = HISTORICAL_TRADES_REJECT  (10101)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalTradesRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalTradesRejectFixedDefault)
}

// ToBuilder
func (m HistoricalTradesRejectFixedPointer) ToBuilder() HistoricalTradesRejectFixedPointerBuilder {
	return HistoricalTradesRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalTradesRejectFixedPointerBuilder) Finish() HistoricalTradesRejectFixedPointer {
	return HistoricalTradesRejectFixedPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalTradesRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalTradesRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText
func (m HistoricalTradesRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText
func (m HistoricalTradesRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID
func (m HistoricalTradesRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText
func (m HistoricalTradesRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m HistoricalTradesRejectFixedPointer) Copy(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalTradesRejectFixedPointerBuilder) Copy(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectFixedPointer) CopyPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectFixedPointerBuilder) CopyPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectFixedPointer) CopyTo(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectFixedPointerBuilder) CopyTo(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectFixedPointer) CopyToPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectFixedPointerBuilder) CopyToPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
