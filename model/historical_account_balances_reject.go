package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = HistoricalAccountBalancesRejectSize  (16)
//     Type       = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     BaseSize   = HistoricalAccountBalancesRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _HistoricalAccountBalancesRejectDefault = []byte{16, 0, 92, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalAccountBalancesRejectSize = 16

// HistoricalAccountBalancesReject This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesReject struct {
	message.VLS
}

// HistoricalAccountBalancesRejectBuilder This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectBuilder struct {
	message.VLSBuilder
}

func NewHistoricalAccountBalancesRejectFrom(b []byte) HistoricalAccountBalancesReject {
	return HistoricalAccountBalancesReject{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalAccountBalancesReject(b []byte) HistoricalAccountBalancesReject {
	return HistoricalAccountBalancesReject{VLS: message.WrapVLS(b)}
}

func NewHistoricalAccountBalancesReject() HistoricalAccountBalancesRejectBuilder {
	a := HistoricalAccountBalancesRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _HistoricalAccountBalancesRejectDefault)
	return a
}

// Clear
// {
//     Size       = HistoricalAccountBalancesRejectSize  (16)
//     Type       = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     BaseSize   = HistoricalAccountBalancesRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m HistoricalAccountBalancesRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalAccountBalancesRejectDefault)
}

// ToBuilder
func (m HistoricalAccountBalancesReject) ToBuilder() HistoricalAccountBalancesRejectBuilder {
	return HistoricalAccountBalancesRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m HistoricalAccountBalancesRejectBuilder) Finish() HistoricalAccountBalancesReject {
	return HistoricalAccountBalancesReject{m.VLSBuilder.Finish()}
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetRequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m *HistoricalAccountBalancesRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// Copy
func (m HistoricalAccountBalancesReject) Copy(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalAccountBalancesRejectBuilder) Copy(to HistoricalAccountBalancesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesReject) CopyPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m HistoricalAccountBalancesRejectBuilder) CopyPointer(to HistoricalAccountBalancesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesReject) CopyToPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m HistoricalAccountBalancesRejectBuilder) CopyToPointer(to HistoricalAccountBalancesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesReject) CopyTo(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectBuilder) CopyTo(to HistoricalAccountBalancesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
