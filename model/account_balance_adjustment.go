package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = AccountBalanceAdjustmentSize  (40)
//     Type         = ACCOUNT_BALANCE_ADJUSTMENT  (607)
//     BaseSize     = AccountBalanceAdjustmentSize  (40)
//     RequestID    = 0
//     TradeAccount = ""
//     CreditAmount = 0
//     DebitAmount  = 0
//     Currency     = ""
//     Reason       = ""
// }
var _AccountBalanceAdjustmentDefault = []byte{40, 0, 95, 2, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AccountBalanceAdjustmentSize = 40

type AccountBalanceAdjustment struct {
	message.VLS
}

type AccountBalanceAdjustmentBuilder struct {
	message.VLSBuilder
}

func NewAccountBalanceAdjustmentFrom(b []byte) AccountBalanceAdjustment {
	return AccountBalanceAdjustment{VLS: message.NewVLSFrom(b)}
}

func WrapAccountBalanceAdjustment(b []byte) AccountBalanceAdjustment {
	return AccountBalanceAdjustment{VLS: message.WrapVLS(b)}
}

func NewAccountBalanceAdjustment() AccountBalanceAdjustmentBuilder {
	a := AccountBalanceAdjustmentBuilder{message.NewVLSBuilder(40)}
	a.Unsafe().SetBytes(0, _AccountBalanceAdjustmentDefault)
	return a
}

// Clear
// {
//     Size         = AccountBalanceAdjustmentSize  (40)
//     Type         = ACCOUNT_BALANCE_ADJUSTMENT  (607)
//     BaseSize     = AccountBalanceAdjustmentSize  (40)
//     RequestID    = 0
//     TradeAccount = ""
//     CreditAmount = 0
//     DebitAmount  = 0
//     Currency     = ""
//     Reason       = ""
// }
func (m AccountBalanceAdjustmentBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceAdjustmentDefault)
}

// ToBuilder
func (m AccountBalanceAdjustment) ToBuilder() AccountBalanceAdjustmentBuilder {
	return AccountBalanceAdjustmentBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m AccountBalanceAdjustmentBuilder) Finish() AccountBalanceAdjustment {
	return AccountBalanceAdjustment{m.VLSBuilder.Finish()}
}

// RequestID
func (m AccountBalanceAdjustment) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m AccountBalanceAdjustmentBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// TradeAccount
func (m AccountBalanceAdjustment) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// TradeAccount
func (m AccountBalanceAdjustmentBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// CreditAmount
func (m AccountBalanceAdjustment) CreditAmount() float64 {
	return message.Float64VLS(m.Unsafe(), 24, 16)
}

// CreditAmount
func (m AccountBalanceAdjustmentBuilder) CreditAmount() float64 {
	return message.Float64VLS(m.Unsafe(), 24, 16)
}

// DebitAmount
func (m AccountBalanceAdjustment) DebitAmount() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
}

// DebitAmount
func (m AccountBalanceAdjustmentBuilder) DebitAmount() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
}

// Currency
func (m AccountBalanceAdjustment) Currency() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
}

// Currency
func (m AccountBalanceAdjustmentBuilder) Currency() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
}

// Reason
func (m AccountBalanceAdjustment) Reason() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// Reason
func (m AccountBalanceAdjustmentBuilder) Reason() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// SetRequestID
func (m AccountBalanceAdjustmentBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetTradeAccount
func (m *AccountBalanceAdjustmentBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetCreditAmount
func (m AccountBalanceAdjustmentBuilder) SetCreditAmount(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 24, 16, value)
}

// SetDebitAmount
func (m AccountBalanceAdjustmentBuilder) SetDebitAmount(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 32, 24, value)
}

// SetCurrency
func (m *AccountBalanceAdjustmentBuilder) SetCurrency(value string) {
	message.SetStringVLS(&m.VLSBuilder, 36, 32, value)
}

// SetReason
func (m *AccountBalanceAdjustmentBuilder) SetReason(value string) {
	message.SetStringVLS(&m.VLSBuilder, 40, 36, value)
}

// Copy
func (m AccountBalanceAdjustment) Copy(to AccountBalanceAdjustmentBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// Copy
func (m AccountBalanceAdjustmentBuilder) Copy(to AccountBalanceAdjustmentBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyPointer
func (m AccountBalanceAdjustment) CopyPointer(to AccountBalanceAdjustmentPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyPointer
func (m AccountBalanceAdjustmentBuilder) CopyPointer(to AccountBalanceAdjustmentPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyToPointer
func (m AccountBalanceAdjustment) CopyToPointer(to AccountBalanceAdjustmentFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyToPointer
func (m AccountBalanceAdjustmentBuilder) CopyToPointer(to AccountBalanceAdjustmentFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyTo
func (m AccountBalanceAdjustment) CopyTo(to AccountBalanceAdjustmentFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyTo
func (m AccountBalanceAdjustmentBuilder) CopyTo(to AccountBalanceAdjustmentFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}
