package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HistoricalAccountBalancesRequestPointer This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestPointer struct {
	message.VLSPointer
}

// HistoricalAccountBalancesRequestPointerBuilder This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocHistoricalAccountBalancesRequest() HistoricalAccountBalancesRequestPointerBuilder {
	a := HistoricalAccountBalancesRequestPointerBuilder{message.AllocVLSBuilder(24)}
	a.Ptr.SetBytes(0, _HistoricalAccountBalancesRequestDefault)
	return a
}

func AllocHistoricalAccountBalancesRequestFrom(b []byte) HistoricalAccountBalancesRequestPointer {
	return HistoricalAccountBalancesRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size          = HistoricalAccountBalancesRequestSize  (24)
//     Type          = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     BaseSize      = HistoricalAccountBalancesRequestSize  (24)
//     RequestID     = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
func (m HistoricalAccountBalancesRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalAccountBalancesRequestDefault)
}

// ToBuilder
func (m HistoricalAccountBalancesRequestPointer) ToBuilder() HistoricalAccountBalancesRequestPointerBuilder {
	return HistoricalAccountBalancesRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalAccountBalancesRequestPointerBuilder) Finish() HistoricalAccountBalancesRequestPointer {
	return HistoricalAccountBalancesRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestPointer) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 24, 16))
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 24, 16))
}

// SetRequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetTradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m *HistoricalAccountBalancesRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetStartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 24, 16, int64(value))
}

// Copy
func (m HistoricalAccountBalancesRequestPointer) Copy(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalAccountBalancesRequestPointerBuilder) Copy(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestPointer) CopyPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestPointerBuilder) CopyPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestPointer) CopyTo(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestPointerBuilder) CopyTo(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestPointer) CopyToPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestPointerBuilder) CopyToPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}
