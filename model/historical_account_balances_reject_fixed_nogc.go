package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HistoricalAccountBalancesRejectFixedPointer This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectFixedPointer struct {
	message.FixedPointer
}

// HistoricalAccountBalancesRejectFixedPointerBuilder This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalAccountBalancesRejectFixed() HistoricalAccountBalancesRejectFixedPointerBuilder {
	a := HistoricalAccountBalancesRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _HistoricalAccountBalancesRejectFixedDefault)
	return a
}

func AllocHistoricalAccountBalancesRejectFixedFrom(b []byte) HistoricalAccountBalancesRejectFixedPointer {
	return HistoricalAccountBalancesRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = HistoricalAccountBalancesRejectFixedSize  (104)
//     Type       = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalAccountBalancesRejectFixedDefault)
}

// ToBuilder
func (m HistoricalAccountBalancesRejectFixedPointer) ToBuilder() HistoricalAccountBalancesRejectFixedPointerBuilder {
	return HistoricalAccountBalancesRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalAccountBalancesRejectFixedPointerBuilder) Finish() HistoricalAccountBalancesRejectFixedPointer {
	return HistoricalAccountBalancesRejectFixedPointer{m.FixedPointer}
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m HistoricalAccountBalancesRejectFixedPointer) Copy(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) Copy(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectFixedPointer) CopyPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) CopyPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectFixedPointer) CopyTo(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) CopyTo(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectFixedPointer) CopyToPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectFixedPointerBuilder) CopyToPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
