package model

import (
	"github.com/moontrade/dtc-go/message"
)

type AccountBalanceAdjustmentFixedPointer struct {
	message.FixedPointer
}

type AccountBalanceAdjustmentFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocAccountBalanceAdjustmentFixed() AccountBalanceAdjustmentFixedPointerBuilder {
	a := AccountBalanceAdjustmentFixedPointerBuilder{message.AllocFixed(160)}
	a.Ptr.SetBytes(0, _AccountBalanceAdjustmentFixedDefault)
	return a
}

func AllocAccountBalanceAdjustmentFixedFrom(b []byte) AccountBalanceAdjustmentFixedPointer {
	return AccountBalanceAdjustmentFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size         = AccountBalanceAdjustmentFixedSize  (160)
//     Type         = ACCOUNT_BALANCE_ADJUSTMENT  (607)
//     RequestID    = 0
//     TradeAccount = ""
//     CreditAmount = 0
//     DebitAmount  = 0
//     Currency     = ""
//     Reason       = ""
// }
func (m AccountBalanceAdjustmentFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceAdjustmentFixedDefault)
}

// ToBuilder
func (m AccountBalanceAdjustmentFixedPointer) ToBuilder() AccountBalanceAdjustmentFixedPointerBuilder {
	return AccountBalanceAdjustmentFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *AccountBalanceAdjustmentFixedPointerBuilder) Finish() AccountBalanceAdjustmentFixedPointer {
	return AccountBalanceAdjustmentFixedPointer{m.FixedPointer}
}

// RequestID
func (m AccountBalanceAdjustmentFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m AccountBalanceAdjustmentFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// TradeAccount
func (m AccountBalanceAdjustmentFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// TradeAccount
func (m AccountBalanceAdjustmentFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// CreditAmount
func (m AccountBalanceAdjustmentFixedPointer) CreditAmount() float64 {
	return message.Float64Fixed(m.Ptr, 48, 40)
}

// CreditAmount
func (m AccountBalanceAdjustmentFixedPointerBuilder) CreditAmount() float64 {
	return message.Float64Fixed(m.Ptr, 48, 40)
}

// DebitAmount
func (m AccountBalanceAdjustmentFixedPointer) DebitAmount() float64 {
	return message.Float64Fixed(m.Ptr, 56, 48)
}

// DebitAmount
func (m AccountBalanceAdjustmentFixedPointerBuilder) DebitAmount() float64 {
	return message.Float64Fixed(m.Ptr, 56, 48)
}

// Currency
func (m AccountBalanceAdjustmentFixedPointer) Currency() string {
	return message.StringFixed(m.Ptr, 64, 56)
}

// Currency
func (m AccountBalanceAdjustmentFixedPointerBuilder) Currency() string {
	return message.StringFixed(m.Ptr, 64, 56)
}

// Reason
func (m AccountBalanceAdjustmentFixedPointer) Reason() string {
	return message.StringFixed(m.Ptr, 160, 64)
}

// Reason
func (m AccountBalanceAdjustmentFixedPointerBuilder) Reason() string {
	return message.StringFixed(m.Ptr, 160, 64)
}

// SetRequestID
func (m AccountBalanceAdjustmentFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTradeAccount
func (m AccountBalanceAdjustmentFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetCreditAmount
func (m AccountBalanceAdjustmentFixedPointerBuilder) SetCreditAmount(value float64) {
	message.SetFloat64Fixed(m.Ptr, 48, 40, value)
}

// SetDebitAmount
func (m AccountBalanceAdjustmentFixedPointerBuilder) SetDebitAmount(value float64) {
	message.SetFloat64Fixed(m.Ptr, 56, 48, value)
}

// SetCurrency
func (m AccountBalanceAdjustmentFixedPointerBuilder) SetCurrency(value string) {
	message.SetStringFixed(m.Ptr, 64, 56, value)
}

// SetReason
func (m AccountBalanceAdjustmentFixedPointerBuilder) SetReason(value string) {
	message.SetStringFixed(m.Ptr, 160, 64, value)
}

// Copy
func (m AccountBalanceAdjustmentFixedPointer) Copy(to AccountBalanceAdjustmentFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// Copy
func (m AccountBalanceAdjustmentFixedPointerBuilder) Copy(to AccountBalanceAdjustmentFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyPointer
func (m AccountBalanceAdjustmentFixedPointer) CopyPointer(to AccountBalanceAdjustmentFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyPointer
func (m AccountBalanceAdjustmentFixedPointerBuilder) CopyPointer(to AccountBalanceAdjustmentFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyTo
func (m AccountBalanceAdjustmentFixedPointer) CopyTo(to AccountBalanceAdjustmentBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyTo
func (m AccountBalanceAdjustmentFixedPointerBuilder) CopyTo(to AccountBalanceAdjustmentBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyToPointer
func (m AccountBalanceAdjustmentFixedPointer) CopyToPointer(to AccountBalanceAdjustmentPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyToPointer
func (m AccountBalanceAdjustmentFixedPointerBuilder) CopyToPointer(to AccountBalanceAdjustmentPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}
