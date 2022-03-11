package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalAccountBalanceResponsePointer struct {
	message.VLSPointer
}

type HistoricalAccountBalanceResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocHistoricalAccountBalanceResponse() HistoricalAccountBalanceResponsePointerBuilder {
	a := HistoricalAccountBalanceResponsePointerBuilder{message.AllocVLSBuilder(56)}
	a.Ptr.SetBytes(0, _HistoricalAccountBalanceResponseDefault)
	return a
}

func AllocHistoricalAccountBalanceResponseFrom(b []byte) HistoricalAccountBalanceResponsePointer {
	return HistoricalAccountBalanceResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m HistoricalAccountBalanceResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalAccountBalanceResponseDefault)
}

// ToBuilder
func (m HistoricalAccountBalanceResponsePointer) ToBuilder() HistoricalAccountBalanceResponsePointerBuilder {
	return HistoricalAccountBalanceResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalAccountBalanceResponsePointerBuilder) Finish() HistoricalAccountBalanceResponsePointer {
	return HistoricalAccountBalanceResponsePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m HistoricalAccountBalanceResponsePointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID
func (m HistoricalAccountBalanceResponsePointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// DateTime
func (m HistoricalAccountBalanceResponsePointer) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Ptr, 24, 16))
}

// DateTime
func (m HistoricalAccountBalanceResponsePointerBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Ptr, 24, 16))
}

// CashBalance
func (m HistoricalAccountBalanceResponsePointer) CashBalance() float64 {
	return message.Float64VLS(m.Ptr, 32, 24)
}

// CashBalance
func (m HistoricalAccountBalanceResponsePointerBuilder) CashBalance() float64 {
	return message.Float64VLS(m.Ptr, 32, 24)
}

// AccountCurrency
func (m HistoricalAccountBalanceResponsePointer) AccountCurrency() string {
	return message.StringVLS(m.Ptr, 36, 32)
}

// AccountCurrency
func (m HistoricalAccountBalanceResponsePointerBuilder) AccountCurrency() string {
	return message.StringVLS(m.Ptr, 36, 32)
}

// TradeAccount
func (m HistoricalAccountBalanceResponsePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 40, 36)
}

// TradeAccount
func (m HistoricalAccountBalanceResponsePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 40, 36)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponsePointer) IsFinalResponse() uint8 {
	return message.Uint8VLS(m.Ptr, 41, 40)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponsePointerBuilder) IsFinalResponse() uint8 {
	return message.Uint8VLS(m.Ptr, 41, 40)
}

// NoAccountBalances
func (m HistoricalAccountBalanceResponsePointer) NoAccountBalances() uint8 {
	return message.Uint8VLS(m.Ptr, 42, 41)
}

// NoAccountBalances
func (m HistoricalAccountBalanceResponsePointerBuilder) NoAccountBalances() uint8 {
	return message.Uint8VLS(m.Ptr, 42, 41)
}

// InfoText
func (m HistoricalAccountBalanceResponsePointer) InfoText() string {
	return message.StringVLS(m.Ptr, 46, 42)
}

// InfoText
func (m HistoricalAccountBalanceResponsePointerBuilder) InfoText() string {
	return message.StringVLS(m.Ptr, 46, 42)
}

// TransactionId
func (m HistoricalAccountBalanceResponsePointer) TransactionId() string {
	return message.StringVLS(m.Ptr, 50, 46)
}

// TransactionId
func (m HistoricalAccountBalanceResponsePointerBuilder) TransactionId() string {
	return message.StringVLS(m.Ptr, 50, 46)
}

// SetRequestID
func (m HistoricalAccountBalanceResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetDateTime
func (m HistoricalAccountBalanceResponsePointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Ptr, 24, 16, float64(value))
}

// SetCashBalance
func (m HistoricalAccountBalanceResponsePointerBuilder) SetCashBalance(value float64) {
	message.SetFloat64VLS(m.Ptr, 32, 24, value)
}

// SetAccountCurrency
func (m *HistoricalAccountBalanceResponsePointerBuilder) SetAccountCurrency(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 36, 32, value)
}

// SetTradeAccount
func (m *HistoricalAccountBalanceResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 40, 36, value)
}

// SetIsFinalResponse
func (m HistoricalAccountBalanceResponsePointerBuilder) SetIsFinalResponse(value uint8) {
	message.SetUint8VLS(m.Ptr, 41, 40, value)
}

// SetNoAccountBalances
func (m HistoricalAccountBalanceResponsePointerBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8VLS(m.Ptr, 42, 41, value)
}

// SetInfoText
func (m *HistoricalAccountBalanceResponsePointerBuilder) SetInfoText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 46, 42, value)
}

// SetTransactionId
func (m *HistoricalAccountBalanceResponsePointerBuilder) SetTransactionId(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 50, 46, value)
}

// Copy
func (m HistoricalAccountBalanceResponsePointer) Copy(to HistoricalAccountBalanceResponseBuilder) {
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
func (m HistoricalAccountBalanceResponsePointerBuilder) Copy(to HistoricalAccountBalanceResponseBuilder) {
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
func (m HistoricalAccountBalanceResponsePointer) CopyPointer(to HistoricalAccountBalanceResponsePointerBuilder) {
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
func (m HistoricalAccountBalanceResponsePointerBuilder) CopyPointer(to HistoricalAccountBalanceResponsePointerBuilder) {
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
func (m HistoricalAccountBalanceResponsePointer) CopyTo(to HistoricalAccountBalanceResponseFixedBuilder) {
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
func (m HistoricalAccountBalanceResponsePointerBuilder) CopyTo(to HistoricalAccountBalanceResponseFixedBuilder) {
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
func (m HistoricalAccountBalanceResponsePointer) CopyToPointer(to HistoricalAccountBalanceResponseFixedPointerBuilder) {
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
func (m HistoricalAccountBalanceResponsePointerBuilder) CopyToPointer(to HistoricalAccountBalanceResponseFixedPointerBuilder) {
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
