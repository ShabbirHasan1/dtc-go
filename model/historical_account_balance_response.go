package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size              = HistoricalAccountBalanceResponseSize  (56)
//     Type              = HISTORICAL_ACCOUNT_BALANCE_RESPONSE  (606)
//     BaseSize          = HistoricalAccountBalanceResponseSize  (56)
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
var _HistoricalAccountBalanceResponseDefault = []byte{56, 0, 94, 2, 56, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalAccountBalanceResponseSize = 56

type HistoricalAccountBalanceResponse struct {
	message.VLS
}

type HistoricalAccountBalanceResponseBuilder struct {
	message.VLSBuilder
}

func NewHistoricalAccountBalanceResponseFrom(b []byte) HistoricalAccountBalanceResponse {
	return HistoricalAccountBalanceResponse{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalAccountBalanceResponse(b []byte) HistoricalAccountBalanceResponse {
	return HistoricalAccountBalanceResponse{VLS: message.WrapVLS(b)}
}

func NewHistoricalAccountBalanceResponse() HistoricalAccountBalanceResponseBuilder {
	a := HistoricalAccountBalanceResponseBuilder{message.NewVLSBuilder(56)}
	a.Unsafe().SetBytes(0, _HistoricalAccountBalanceResponseDefault)
	return a
}

// Clear
// {
//     Size              = HistoricalAccountBalanceResponseSize  (56)
//     Type              = HISTORICAL_ACCOUNT_BALANCE_RESPONSE  (606)
//     BaseSize          = HistoricalAccountBalanceResponseSize  (56)
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
func (m HistoricalAccountBalanceResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalAccountBalanceResponseDefault)
}

// ToBuilder
func (m HistoricalAccountBalanceResponse) ToBuilder() HistoricalAccountBalanceResponseBuilder {
	return HistoricalAccountBalanceResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m HistoricalAccountBalanceResponseBuilder) Finish() HistoricalAccountBalanceResponse {
	return HistoricalAccountBalanceResponse{m.VLSBuilder.Finish()}
}

// RequestID
func (m HistoricalAccountBalanceResponse) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m HistoricalAccountBalanceResponseBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// DateTime
func (m HistoricalAccountBalanceResponse) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 24, 16))
}

// DateTime
func (m HistoricalAccountBalanceResponseBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 24, 16))
}

// CashBalance
func (m HistoricalAccountBalanceResponse) CashBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
}

// CashBalance
func (m HistoricalAccountBalanceResponseBuilder) CashBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
}

// AccountCurrency
func (m HistoricalAccountBalanceResponse) AccountCurrency() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
}

// AccountCurrency
func (m HistoricalAccountBalanceResponseBuilder) AccountCurrency() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
}

// TradeAccount
func (m HistoricalAccountBalanceResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// TradeAccount
func (m HistoricalAccountBalanceResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponse) IsFinalResponse() uint8 {
	return message.Uint8VLS(m.Unsafe(), 41, 40)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponseBuilder) IsFinalResponse() uint8 {
	return message.Uint8VLS(m.Unsafe(), 41, 40)
}

// NoAccountBalances
func (m HistoricalAccountBalanceResponse) NoAccountBalances() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// NoAccountBalances
func (m HistoricalAccountBalanceResponseBuilder) NoAccountBalances() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// InfoText
func (m HistoricalAccountBalanceResponse) InfoText() string {
	return message.StringVLS(m.Unsafe(), 46, 42)
}

// InfoText
func (m HistoricalAccountBalanceResponseBuilder) InfoText() string {
	return message.StringVLS(m.Unsafe(), 46, 42)
}

// TransactionId
func (m HistoricalAccountBalanceResponse) TransactionId() string {
	return message.StringVLS(m.Unsafe(), 50, 46)
}

// TransactionId
func (m HistoricalAccountBalanceResponseBuilder) TransactionId() string {
	return message.StringVLS(m.Unsafe(), 50, 46)
}

// SetRequestID
func (m HistoricalAccountBalanceResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetDateTime
func (m HistoricalAccountBalanceResponseBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Unsafe(), 24, 16, float64(value))
}

// SetCashBalance
func (m HistoricalAccountBalanceResponseBuilder) SetCashBalance(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 32, 24, value)
}

// SetAccountCurrency
func (m *HistoricalAccountBalanceResponseBuilder) SetAccountCurrency(value string) {
	message.SetStringVLS(&m.VLSBuilder, 36, 32, value)
}

// SetTradeAccount
func (m *HistoricalAccountBalanceResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 40, 36, value)
}

// SetIsFinalResponse
func (m HistoricalAccountBalanceResponseBuilder) SetIsFinalResponse(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 41, 40, value)
}

// SetNoAccountBalances
func (m HistoricalAccountBalanceResponseBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 42, 41, value)
}

// SetInfoText
func (m *HistoricalAccountBalanceResponseBuilder) SetInfoText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 46, 42, value)
}

// SetTransactionId
func (m *HistoricalAccountBalanceResponseBuilder) SetTransactionId(value string) {
	message.SetStringVLS(&m.VLSBuilder, 50, 46, value)
}

// Copy
func (m HistoricalAccountBalanceResponse) Copy(to HistoricalAccountBalanceResponseBuilder) {
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
func (m HistoricalAccountBalanceResponseBuilder) Copy(to HistoricalAccountBalanceResponseBuilder) {
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
func (m HistoricalAccountBalanceResponse) CopyPointer(to HistoricalAccountBalanceResponsePointerBuilder) {
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
func (m HistoricalAccountBalanceResponseBuilder) CopyPointer(to HistoricalAccountBalanceResponsePointerBuilder) {
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
func (m HistoricalAccountBalanceResponse) CopyToPointer(to HistoricalAccountBalanceResponseFixedPointerBuilder) {
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
func (m HistoricalAccountBalanceResponseBuilder) CopyToPointer(to HistoricalAccountBalanceResponseFixedPointerBuilder) {
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
func (m HistoricalAccountBalanceResponse) CopyTo(to HistoricalAccountBalanceResponseFixedBuilder) {
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
func (m HistoricalAccountBalanceResponseBuilder) CopyTo(to HistoricalAccountBalanceResponseFixedBuilder) {
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
