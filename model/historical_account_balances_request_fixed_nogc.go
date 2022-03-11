package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HistoricalAccountBalancesRequestFixedPointer This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestFixedPointer struct {
	message.FixedPointer
}

// HistoricalAccountBalancesRequestFixedPointerBuilder This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalAccountBalancesRequestFixed() HistoricalAccountBalancesRequestFixedPointerBuilder {
	a := HistoricalAccountBalancesRequestFixedPointerBuilder{message.AllocFixed(48)}
	a.Ptr.SetBytes(0, _HistoricalAccountBalancesRequestFixedDefault)
	return a
}

func AllocHistoricalAccountBalancesRequestFixedFrom(b []byte) HistoricalAccountBalancesRequestFixedPointer {
	return HistoricalAccountBalancesRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = HistoricalAccountBalancesRequestFixedSize  (48)
//     Type          = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     RequestID     = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalAccountBalancesRequestFixedDefault)
}

// ToBuilder
func (m HistoricalAccountBalancesRequestFixedPointer) ToBuilder() HistoricalAccountBalancesRequestFixedPointerBuilder {
	return HistoricalAccountBalancesRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalAccountBalancesRequestFixedPointerBuilder) Finish() HistoricalAccountBalancesRequestFixedPointer {
	return HistoricalAccountBalancesRequestFixedPointer{m.FixedPointer}
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixedPointer) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 48, 40))
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 48, 40))
}

// SetRequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetStartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 48, 40, int64(value))
}

// Copy
func (m HistoricalAccountBalancesRequestFixedPointer) Copy(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) Copy(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestFixedPointer) CopyPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) CopyPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestFixedPointer) CopyTo(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) CopyTo(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestFixedPointer) CopyToPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) CopyToPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}
