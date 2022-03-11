package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalAccountBalanceResponseFixedPointer struct {
	message.FixedPointer
}

type HistoricalAccountBalanceResponseFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalAccountBalanceResponseFixed() HistoricalAccountBalanceResponseFixedPointerBuilder {
	a := HistoricalAccountBalanceResponseFixedPointerBuilder{message.AllocFixed(232)}
	a.Ptr.SetBytes(0, _HistoricalAccountBalanceResponseFixedDefault)
	return a
}

func AllocHistoricalAccountBalanceResponseFixedFrom(b []byte) HistoricalAccountBalanceResponseFixedPointer {
	return HistoricalAccountBalanceResponseFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size              = HistoricalAccountBalanceResponseFixedSize  (232)
//     Type              = HISTORICAL_ACCOUNT_BALANCE_RESPONSE  (606)
//     RequestID         = 0
//     DateTime          = 0
//     CashBalance       = 0
//     AccountCurrency   = ""
//     TradeAccount      = ""
//     IsFinalResponse   = 0
//     NoAccountBalances = 0
//     InfoText          = ""
//     TransactionId     = ""
// }
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalAccountBalanceResponseFixedDefault)
}

// ToBuilder
func (m HistoricalAccountBalanceResponseFixedPointer) ToBuilder() HistoricalAccountBalanceResponseFixedPointerBuilder {
	return HistoricalAccountBalanceResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalAccountBalanceResponseFixedPointerBuilder) Finish() HistoricalAccountBalanceResponseFixedPointer {
	return HistoricalAccountBalanceResponseFixedPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalAccountBalanceResponseFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// DateTime
func (m HistoricalAccountBalanceResponseFixedPointer) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 16, 8))
}

// DateTime
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 16, 8))
}

// CashBalance
func (m HistoricalAccountBalanceResponseFixedPointer) CashBalance() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// CashBalance
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) CashBalance() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// AccountCurrency
func (m HistoricalAccountBalanceResponseFixedPointer) AccountCurrency() string {
	return message.StringFixed(m.Ptr, 32, 24)
}

// AccountCurrency
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) AccountCurrency() string {
	return message.StringFixed(m.Ptr, 32, 24)
}

// TradeAccount
func (m HistoricalAccountBalanceResponseFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 64, 32)
}

// TradeAccount
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 64, 32)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponseFixedPointer) IsFinalResponse() uint8 {
	return message.Uint8Fixed(m.Ptr, 65, 64)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) IsFinalResponse() uint8 {
	return message.Uint8Fixed(m.Ptr, 65, 64)
}

// NoAccountBalances
func (m HistoricalAccountBalanceResponseFixedPointer) NoAccountBalances() uint8 {
	return message.Uint8Fixed(m.Ptr, 66, 65)
}

// NoAccountBalances
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) NoAccountBalances() uint8 {
	return message.Uint8Fixed(m.Ptr, 66, 65)
}

// InfoText
func (m HistoricalAccountBalanceResponseFixedPointer) InfoText() string {
	return message.StringFixed(m.Ptr, 162, 66)
}

// InfoText
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) InfoText() string {
	return message.StringFixed(m.Ptr, 162, 66)
}

// TransactionId
func (m HistoricalAccountBalanceResponseFixedPointer) TransactionId() string {
	return message.StringFixed(m.Ptr, 226, 162)
}

// TransactionId
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) TransactionId() string {
	return message.StringFixed(m.Ptr, 226, 162)
}

// SetRequestID
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetDateTime
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, float64(value))
}

// SetCashBalance
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetCashBalance(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetAccountCurrency
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetAccountCurrency(value string) {
	message.SetStringFixed(m.Ptr, 32, 24, value)
}

// SetTradeAccount
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 64, 32, value)
}

// SetIsFinalResponse
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetIsFinalResponse(value uint8) {
	message.SetUint8Fixed(m.Ptr, 65, 64, value)
}

// SetNoAccountBalances
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8Fixed(m.Ptr, 66, 65, value)
}

// SetInfoText
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Ptr, 162, 66, value)
}

// SetTransactionId
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetTransactionId(value string) {
	message.SetStringFixed(m.Ptr, 226, 162, value)
}

// Copy
func (m HistoricalAccountBalanceResponseFixedPointer) Copy(to HistoricalAccountBalanceResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetCashBalance(m.CashBalance())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsFinalResponse(m.IsFinalResponse())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetInfoText(m.InfoText())
	to.SetTransactionId(m.TransactionId())
}

// Copy
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) Copy(to HistoricalAccountBalanceResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetCashBalance(m.CashBalance())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsFinalResponse(m.IsFinalResponse())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetInfoText(m.InfoText())
	to.SetTransactionId(m.TransactionId())
}

// CopyPointer
func (m HistoricalAccountBalanceResponseFixedPointer) CopyPointer(to HistoricalAccountBalanceResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetCashBalance(m.CashBalance())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsFinalResponse(m.IsFinalResponse())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetInfoText(m.InfoText())
	to.SetTransactionId(m.TransactionId())
}

// CopyPointer
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) CopyPointer(to HistoricalAccountBalanceResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetCashBalance(m.CashBalance())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsFinalResponse(m.IsFinalResponse())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetInfoText(m.InfoText())
	to.SetTransactionId(m.TransactionId())
}

// CopyTo
func (m HistoricalAccountBalanceResponseFixedPointer) CopyTo(to HistoricalAccountBalanceResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetCashBalance(m.CashBalance())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsFinalResponse(m.IsFinalResponse())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetInfoText(m.InfoText())
	to.SetTransactionId(m.TransactionId())
}

// CopyTo
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) CopyTo(to HistoricalAccountBalanceResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetCashBalance(m.CashBalance())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsFinalResponse(m.IsFinalResponse())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetInfoText(m.InfoText())
	to.SetTransactionId(m.TransactionId())
}

// CopyToPointer
func (m HistoricalAccountBalanceResponseFixedPointer) CopyToPointer(to HistoricalAccountBalanceResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetCashBalance(m.CashBalance())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsFinalResponse(m.IsFinalResponse())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetInfoText(m.InfoText())
	to.SetTransactionId(m.TransactionId())
}

// CopyToPointer
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) CopyToPointer(to HistoricalAccountBalanceResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetCashBalance(m.CashBalance())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsFinalResponse(m.IsFinalResponse())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetInfoText(m.InfoText())
	to.SetTransactionId(m.TransactionId())
}
