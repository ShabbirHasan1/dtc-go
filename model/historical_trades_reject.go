package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = HistoricalTradesRejectSize  (14)
//     Type       = HISTORICAL_TRADES_REJECT  (10101)
//     BaseSize   = HistoricalTradesRejectSize  (14)
//     RequestID  = 0
//     RejectText = ""
// }
var _HistoricalTradesRejectDefault = []byte{14, 0, 117, 39, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalTradesRejectSize = 14

type HistoricalTradesReject struct {
	message.VLS
}

type HistoricalTradesRejectBuilder struct {
	message.VLSBuilder
}

func NewHistoricalTradesRejectFrom(b []byte) HistoricalTradesReject {
	return HistoricalTradesReject{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalTradesReject(b []byte) HistoricalTradesReject {
	return HistoricalTradesReject{VLS: message.WrapVLS(b)}
}

func NewHistoricalTradesReject() HistoricalTradesRejectBuilder {
	a := HistoricalTradesRejectBuilder{message.NewVLSBuilder(14)}
	a.Unsafe().SetBytes(0, _HistoricalTradesRejectDefault)
	return a
}

// Clear
// {
//     Size       = HistoricalTradesRejectSize  (14)
//     Type       = HISTORICAL_TRADES_REJECT  (10101)
//     BaseSize   = HistoricalTradesRejectSize  (14)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalTradesRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalTradesRejectDefault)
}

// ToBuilder
func (m HistoricalTradesReject) ToBuilder() HistoricalTradesRejectBuilder {
	return HistoricalTradesRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m HistoricalTradesRejectBuilder) Finish() HistoricalTradesReject {
	return HistoricalTradesReject{m.VLSBuilder.Finish()}
}

// RequestID
func (m HistoricalTradesReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m HistoricalTradesRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RejectText
func (m HistoricalTradesReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// RejectText
func (m HistoricalTradesRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// SetRequestID
func (m HistoricalTradesRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 10, 6, value)
}

// SetRejectText
func (m *HistoricalTradesRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// Copy
func (m HistoricalTradesReject) Copy(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalTradesRejectBuilder) Copy(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesReject) CopyPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectBuilder) CopyPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesReject) CopyToPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectBuilder) CopyToPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesReject) CopyTo(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectBuilder) CopyTo(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
