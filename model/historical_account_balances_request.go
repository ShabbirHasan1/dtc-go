package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = HistoricalAccountBalancesRequestSize  (24)
//     Type          = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     BaseSize      = HistoricalAccountBalancesRequestSize  (24)
//     RequestID     = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
var _HistoricalAccountBalancesRequestDefault = []byte{24, 0, 91, 2, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalAccountBalancesRequestSize = 24

// HistoricalAccountBalancesRequest This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequest struct {
	message.VLS
}

// HistoricalAccountBalancesRequestBuilder This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestBuilder struct {
	message.VLSBuilder
}

func NewHistoricalAccountBalancesRequestFrom(b []byte) HistoricalAccountBalancesRequest {
	return HistoricalAccountBalancesRequest{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalAccountBalancesRequest(b []byte) HistoricalAccountBalancesRequest {
	return HistoricalAccountBalancesRequest{VLS: message.WrapVLS(b)}
}

func NewHistoricalAccountBalancesRequest() HistoricalAccountBalancesRequestBuilder {
	a := HistoricalAccountBalancesRequestBuilder{message.NewVLSBuilder(24)}
	a.Unsafe().SetBytes(0, _HistoricalAccountBalancesRequestDefault)
	return a
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
func (m HistoricalAccountBalancesRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalAccountBalancesRequestDefault)
}

// ToBuilder
func (m HistoricalAccountBalancesRequest) ToBuilder() HistoricalAccountBalancesRequestBuilder {
	return HistoricalAccountBalancesRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m HistoricalAccountBalancesRequestBuilder) Finish() HistoricalAccountBalancesRequest {
	return HistoricalAccountBalancesRequest{m.VLSBuilder.Finish()}
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequest) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 24, 16))
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 24, 16))
}

// SetRequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetTradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m *HistoricalAccountBalancesRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetStartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 24, 16, int64(value))
}

// Copy
func (m HistoricalAccountBalancesRequest) Copy(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalAccountBalancesRequestBuilder) Copy(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequest) CopyPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestBuilder) CopyPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequest) CopyToPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestBuilder) CopyToPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequest) CopyTo(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestBuilder) CopyTo(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}
