package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = HistoricalAccountBalancesRequestFixedSize  (48)
//     Type          = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     RequestID     = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
var _HistoricalAccountBalancesRequestFixedDefault = []byte{48, 0, 91, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalAccountBalancesRequestFixedSize = 48

// HistoricalAccountBalancesRequestFixed This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestFixed struct {
	message.Fixed
}

// HistoricalAccountBalancesRequestFixedBuilder This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestFixedBuilder struct {
	message.Fixed
}

func NewHistoricalAccountBalancesRequestFixedFrom(b []byte) HistoricalAccountBalancesRequestFixed {
	return HistoricalAccountBalancesRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalAccountBalancesRequestFixed(b []byte) HistoricalAccountBalancesRequestFixed {
	return HistoricalAccountBalancesRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalAccountBalancesRequestFixed() HistoricalAccountBalancesRequestFixedBuilder {
	a := HistoricalAccountBalancesRequestFixedBuilder{message.NewFixed(48)}
	a.Unsafe().SetBytes(0, _HistoricalAccountBalancesRequestFixedDefault)
	return a
}

// Clear
// {
//     Size          = HistoricalAccountBalancesRequestFixedSize  (48)
//     Type          = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     RequestID     = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
func (m HistoricalAccountBalancesRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalAccountBalancesRequestFixedDefault)
}

// ToBuilder
func (m HistoricalAccountBalancesRequestFixed) ToBuilder() HistoricalAccountBalancesRequestFixedBuilder {
	return HistoricalAccountBalancesRequestFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalAccountBalancesRequestFixedBuilder) Finish() HistoricalAccountBalancesRequestFixed {
	return HistoricalAccountBalancesRequestFixed{m.Fixed}
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixed) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 48, 40))
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixedBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 48, 40))
}

// SetRequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetTradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetStartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixedBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 48, 40, int64(value))
}

// Copy
func (m HistoricalAccountBalancesRequestFixed) Copy(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalAccountBalancesRequestFixedBuilder) Copy(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestFixed) CopyPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestFixedBuilder) CopyPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestFixed) CopyToPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestFixedBuilder) CopyToPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestFixed) CopyTo(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestFixedBuilder) CopyTo(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}
