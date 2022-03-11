package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _HistoricalAccountBalanceResponseFixedDefault = []byte{232, 0, 94, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalAccountBalanceResponseFixedSize = 232

type HistoricalAccountBalanceResponseFixed struct {
	message.Fixed
}

type HistoricalAccountBalanceResponseFixedBuilder struct {
	message.Fixed
}

func NewHistoricalAccountBalanceResponseFixedFrom(b []byte) HistoricalAccountBalanceResponseFixed {
	return HistoricalAccountBalanceResponseFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalAccountBalanceResponseFixed(b []byte) HistoricalAccountBalanceResponseFixed {
	return HistoricalAccountBalanceResponseFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalAccountBalanceResponseFixed() HistoricalAccountBalanceResponseFixedBuilder {
	a := HistoricalAccountBalanceResponseFixedBuilder{message.NewFixed(232)}
	a.Unsafe().SetBytes(0, _HistoricalAccountBalanceResponseFixedDefault)
	return a
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
func (m HistoricalAccountBalanceResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalAccountBalanceResponseFixedDefault)
}

// ToBuilder
func (m HistoricalAccountBalanceResponseFixed) ToBuilder() HistoricalAccountBalanceResponseFixedBuilder {
	return HistoricalAccountBalanceResponseFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalAccountBalanceResponseFixedBuilder) Finish() HistoricalAccountBalanceResponseFixed {
	return HistoricalAccountBalanceResponseFixed{m.Fixed}
}

// RequestID
func (m HistoricalAccountBalanceResponseFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalAccountBalanceResponseFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// DateTime
func (m HistoricalAccountBalanceResponseFixed) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 16, 8))
}

// DateTime
func (m HistoricalAccountBalanceResponseFixedBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 16, 8))
}

// CashBalance
func (m HistoricalAccountBalanceResponseFixed) CashBalance() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// CashBalance
func (m HistoricalAccountBalanceResponseFixedBuilder) CashBalance() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// AccountCurrency
func (m HistoricalAccountBalanceResponseFixed) AccountCurrency() string {
	return message.StringFixed(m.Unsafe(), 32, 24)
}

// AccountCurrency
func (m HistoricalAccountBalanceResponseFixedBuilder) AccountCurrency() string {
	return message.StringFixed(m.Unsafe(), 32, 24)
}

// TradeAccount
func (m HistoricalAccountBalanceResponseFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 64, 32)
}

// TradeAccount
func (m HistoricalAccountBalanceResponseFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 64, 32)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponseFixed) IsFinalResponse() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 65, 64)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponseFixedBuilder) IsFinalResponse() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 65, 64)
}

// NoAccountBalances
func (m HistoricalAccountBalanceResponseFixed) NoAccountBalances() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 66, 65)
}

// NoAccountBalances
func (m HistoricalAccountBalanceResponseFixedBuilder) NoAccountBalances() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 66, 65)
}

// InfoText
func (m HistoricalAccountBalanceResponseFixed) InfoText() string {
	return message.StringFixed(m.Unsafe(), 162, 66)
}

// InfoText
func (m HistoricalAccountBalanceResponseFixedBuilder) InfoText() string {
	return message.StringFixed(m.Unsafe(), 162, 66)
}

// TransactionId
func (m HistoricalAccountBalanceResponseFixed) TransactionId() string {
	return message.StringFixed(m.Unsafe(), 226, 162)
}

// TransactionId
func (m HistoricalAccountBalanceResponseFixedBuilder) TransactionId() string {
	return message.StringFixed(m.Unsafe(), 226, 162)
}

// SetRequestID
func (m HistoricalAccountBalanceResponseFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetDateTime
func (m HistoricalAccountBalanceResponseFixedBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, float64(value))
}

// SetCashBalance
func (m HistoricalAccountBalanceResponseFixedBuilder) SetCashBalance(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetAccountCurrency
func (m HistoricalAccountBalanceResponseFixedBuilder) SetAccountCurrency(value string) {
	message.SetStringFixed(m.Unsafe(), 32, 24, value)
}

// SetTradeAccount
func (m HistoricalAccountBalanceResponseFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 64, 32, value)
}

// SetIsFinalResponse
func (m HistoricalAccountBalanceResponseFixedBuilder) SetIsFinalResponse(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 65, 64, value)
}

// SetNoAccountBalances
func (m HistoricalAccountBalanceResponseFixedBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 66, 65, value)
}

// SetInfoText
func (m HistoricalAccountBalanceResponseFixedBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Unsafe(), 162, 66, value)
}

// SetTransactionId
func (m HistoricalAccountBalanceResponseFixedBuilder) SetTransactionId(value string) {
	message.SetStringFixed(m.Unsafe(), 226, 162, value)
}

// Copy
func (m HistoricalAccountBalanceResponseFixed) Copy(to HistoricalAccountBalanceResponseFixedBuilder) {
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
func (m HistoricalAccountBalanceResponseFixedBuilder) Copy(to HistoricalAccountBalanceResponseFixedBuilder) {
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
func (m HistoricalAccountBalanceResponseFixed) CopyPointer(to HistoricalAccountBalanceResponseFixedPointerBuilder) {
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
func (m HistoricalAccountBalanceResponseFixedBuilder) CopyPointer(to HistoricalAccountBalanceResponseFixedPointerBuilder) {
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
func (m HistoricalAccountBalanceResponseFixed) CopyToPointer(to HistoricalAccountBalanceResponsePointerBuilder) {
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
func (m HistoricalAccountBalanceResponseFixedBuilder) CopyToPointer(to HistoricalAccountBalanceResponsePointerBuilder) {
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
func (m HistoricalAccountBalanceResponseFixed) CopyTo(to HistoricalAccountBalanceResponseBuilder) {
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
func (m HistoricalAccountBalanceResponseFixedBuilder) CopyTo(to HistoricalAccountBalanceResponseBuilder) {
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
