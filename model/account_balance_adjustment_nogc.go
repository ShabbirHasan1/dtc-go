package model

import (
	"github.com/moontrade/dtc-go/message"
)

type AccountBalanceAdjustmentPointer struct {
	message.VLSPointer
}

type AccountBalanceAdjustmentPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocAccountBalanceAdjustment() AccountBalanceAdjustmentPointerBuilder {
	a := AccountBalanceAdjustmentPointerBuilder{message.AllocVLSBuilder(40)}
	a.Ptr.SetBytes(0, _AccountBalanceAdjustmentDefault)
	return a
}

func AllocAccountBalanceAdjustmentFrom(b []byte) AccountBalanceAdjustmentPointer {
	return AccountBalanceAdjustmentPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m AccountBalanceAdjustmentPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceAdjustmentDefault)
}

// ToBuilder
func (m AccountBalanceAdjustmentPointer) ToBuilder() AccountBalanceAdjustmentPointerBuilder {
	return AccountBalanceAdjustmentPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *AccountBalanceAdjustmentPointerBuilder) Finish() AccountBalanceAdjustmentPointer {
	return AccountBalanceAdjustmentPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m AccountBalanceAdjustmentPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID
func (m AccountBalanceAdjustmentPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// TradeAccount
func (m AccountBalanceAdjustmentPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// TradeAccount
func (m AccountBalanceAdjustmentPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// CreditAmount
func (m AccountBalanceAdjustmentPointer) CreditAmount() float64 {
	return message.Float64VLS(m.Ptr, 24, 16)
}

// CreditAmount
func (m AccountBalanceAdjustmentPointerBuilder) CreditAmount() float64 {
	return message.Float64VLS(m.Ptr, 24, 16)
}

// DebitAmount
func (m AccountBalanceAdjustmentPointer) DebitAmount() float64 {
	return message.Float64VLS(m.Ptr, 32, 24)
}

// DebitAmount
func (m AccountBalanceAdjustmentPointerBuilder) DebitAmount() float64 {
	return message.Float64VLS(m.Ptr, 32, 24)
}

// Currency
func (m AccountBalanceAdjustmentPointer) Currency() string {
	return message.StringVLS(m.Ptr, 36, 32)
}

// Currency
func (m AccountBalanceAdjustmentPointerBuilder) Currency() string {
	return message.StringVLS(m.Ptr, 36, 32)
}

// Reason
func (m AccountBalanceAdjustmentPointer) Reason() string {
	return message.StringVLS(m.Ptr, 40, 36)
}

// Reason
func (m AccountBalanceAdjustmentPointerBuilder) Reason() string {
	return message.StringVLS(m.Ptr, 40, 36)
}

// SetRequestID
func (m AccountBalanceAdjustmentPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetTradeAccount
func (m *AccountBalanceAdjustmentPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetCreditAmount
func (m AccountBalanceAdjustmentPointerBuilder) SetCreditAmount(value float64) {
	message.SetFloat64VLS(m.Ptr, 24, 16, value)
}

// SetDebitAmount
func (m AccountBalanceAdjustmentPointerBuilder) SetDebitAmount(value float64) {
	message.SetFloat64VLS(m.Ptr, 32, 24, value)
}

// SetCurrency
func (m *AccountBalanceAdjustmentPointerBuilder) SetCurrency(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 36, 32, value)
}

// SetReason
func (m *AccountBalanceAdjustmentPointerBuilder) SetReason(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 40, 36, value)
}

// Copy
func (m AccountBalanceAdjustmentPointer) Copy(to AccountBalanceAdjustmentBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// Copy
func (m AccountBalanceAdjustmentPointerBuilder) Copy(to AccountBalanceAdjustmentBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyPointer
func (m AccountBalanceAdjustmentPointer) CopyPointer(to AccountBalanceAdjustmentPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyPointer
func (m AccountBalanceAdjustmentPointerBuilder) CopyPointer(to AccountBalanceAdjustmentPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyTo
func (m AccountBalanceAdjustmentPointer) CopyTo(to AccountBalanceAdjustmentFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyTo
func (m AccountBalanceAdjustmentPointerBuilder) CopyTo(to AccountBalanceAdjustmentFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyToPointer
func (m AccountBalanceAdjustmentPointer) CopyToPointer(to AccountBalanceAdjustmentFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyToPointer
func (m AccountBalanceAdjustmentPointerBuilder) CopyToPointer(to AccountBalanceAdjustmentFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}
