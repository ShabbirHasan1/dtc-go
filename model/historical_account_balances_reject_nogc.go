package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HistoricalAccountBalancesRejectPointer This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectPointer struct {
	message.VLSPointer
}

// HistoricalAccountBalancesRejectPointerBuilder This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocHistoricalAccountBalancesReject() HistoricalAccountBalancesRejectPointerBuilder {
	a := HistoricalAccountBalancesRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _HistoricalAccountBalancesRejectDefault)
	return a
}

func AllocHistoricalAccountBalancesRejectFrom(b []byte) HistoricalAccountBalancesRejectPointer {
	return HistoricalAccountBalancesRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size       = HistoricalAccountBalancesRejectSize  (16)
//     Type       = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     BaseSize   = HistoricalAccountBalancesRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalAccountBalancesRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalAccountBalancesRejectDefault)
}

// ToBuilder
func (m HistoricalAccountBalancesRejectPointer) ToBuilder() HistoricalAccountBalancesRejectPointerBuilder {
	return HistoricalAccountBalancesRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalAccountBalancesRejectPointerBuilder) Finish() HistoricalAccountBalancesRejectPointer {
	return HistoricalAccountBalancesRejectPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetRequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m *HistoricalAccountBalancesRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// Copy
func (m HistoricalAccountBalancesRejectPointer) Copy(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalAccountBalancesRejectPointerBuilder) Copy(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectPointer) CopyPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectPointerBuilder) CopyPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectPointer) CopyTo(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectPointerBuilder) CopyTo(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectPointer) CopyToPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectPointerBuilder) CopyToPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
