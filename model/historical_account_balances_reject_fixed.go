package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = HistoricalAccountBalancesRejectFixedSize  (104)
//     Type       = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     RequestID  = 0
//     RejectText = ""
// }
var _HistoricalAccountBalancesRejectFixedDefault = []byte{104, 0, 92, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalAccountBalancesRejectFixedSize = 104

// HistoricalAccountBalancesRejectFixed This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectFixed struct {
	message.Fixed
}

// HistoricalAccountBalancesRejectFixedBuilder This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectFixedBuilder struct {
	message.Fixed
}

func NewHistoricalAccountBalancesRejectFixedFrom(b []byte) HistoricalAccountBalancesRejectFixed {
	return HistoricalAccountBalancesRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalAccountBalancesRejectFixed(b []byte) HistoricalAccountBalancesRejectFixed {
	return HistoricalAccountBalancesRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalAccountBalancesRejectFixed() HistoricalAccountBalancesRejectFixedBuilder {
	a := HistoricalAccountBalancesRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _HistoricalAccountBalancesRejectFixedDefault)
	return a
}

// Clear
// {
//     Size       = HistoricalAccountBalancesRejectFixedSize  (104)
//     Type       = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalAccountBalancesRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalAccountBalancesRejectFixedDefault)
}

// ToBuilder
func (m HistoricalAccountBalancesRejectFixed) ToBuilder() HistoricalAccountBalancesRejectFixedBuilder {
	return HistoricalAccountBalancesRejectFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalAccountBalancesRejectFixedBuilder) Finish() HistoricalAccountBalancesRejectFixed {
	return HistoricalAccountBalancesRejectFixed{m.Fixed}
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// SetRequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// Copy
func (m HistoricalAccountBalancesRejectFixed) Copy(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalAccountBalancesRejectFixedBuilder) Copy(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectFixed) CopyPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectFixedBuilder) CopyPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectFixed) CopyToPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectFixedBuilder) CopyToPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectFixed) CopyTo(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectFixedBuilder) CopyTo(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
