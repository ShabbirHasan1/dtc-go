package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = HistoricalTradesRejectFixedSize  (104)
//     Type       = HISTORICAL_TRADES_REJECT  (10101)
//     RequestID  = 0
//     RejectText = ""
// }
var _HistoricalTradesRejectFixedDefault = []byte{104, 0, 117, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalTradesRejectFixedSize = 104

type HistoricalTradesRejectFixed struct {
	message.Fixed
}

type HistoricalTradesRejectFixedBuilder struct {
	message.Fixed
}

func NewHistoricalTradesRejectFixedFrom(b []byte) HistoricalTradesRejectFixed {
	return HistoricalTradesRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalTradesRejectFixed(b []byte) HistoricalTradesRejectFixed {
	return HistoricalTradesRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalTradesRejectFixed() HistoricalTradesRejectFixedBuilder {
	a := HistoricalTradesRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _HistoricalTradesRejectFixedDefault)
	return a
}

// Clear
// {
//     Size       = HistoricalTradesRejectFixedSize  (104)
//     Type       = HISTORICAL_TRADES_REJECT  (10101)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalTradesRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalTradesRejectFixedDefault)
}

// ToBuilder
func (m HistoricalTradesRejectFixed) ToBuilder() HistoricalTradesRejectFixedBuilder {
	return HistoricalTradesRejectFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalTradesRejectFixedBuilder) Finish() HistoricalTradesRejectFixed {
	return HistoricalTradesRejectFixed{m.Fixed}
}

// RequestID
func (m HistoricalTradesRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalTradesRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RejectText
func (m HistoricalTradesRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m HistoricalTradesRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// SetRequestID
func (m HistoricalTradesRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText
func (m HistoricalTradesRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// Copy
func (m HistoricalTradesRejectFixed) Copy(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalTradesRejectFixedBuilder) Copy(to HistoricalTradesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectFixed) CopyPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalTradesRejectFixedBuilder) CopyPointer(to HistoricalTradesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectFixed) CopyToPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalTradesRejectFixedBuilder) CopyToPointer(to HistoricalTradesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectFixed) CopyTo(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalTradesRejectFixedBuilder) CopyTo(to HistoricalTradesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
