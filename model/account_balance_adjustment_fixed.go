package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _AccountBalanceAdjustmentFixedDefault = []byte{160, 0, 95, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AccountBalanceAdjustmentFixedSize = 160

type AccountBalanceAdjustmentFixed struct {
	message.Fixed
}

type AccountBalanceAdjustmentFixedBuilder struct {
	message.Fixed
}

func NewAccountBalanceAdjustmentFixedFrom(b []byte) AccountBalanceAdjustmentFixed {
	return AccountBalanceAdjustmentFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapAccountBalanceAdjustmentFixed(b []byte) AccountBalanceAdjustmentFixed {
	return AccountBalanceAdjustmentFixed{Fixed: message.WrapFixed(b)}
}

func NewAccountBalanceAdjustmentFixed() AccountBalanceAdjustmentFixedBuilder {
	a := AccountBalanceAdjustmentFixedBuilder{message.NewFixed(160)}
	a.Unsafe().SetBytes(0, _AccountBalanceAdjustmentFixedDefault)
	return a
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
func (m AccountBalanceAdjustmentFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceAdjustmentFixedDefault)
}

// ToBuilder
func (m AccountBalanceAdjustmentFixed) ToBuilder() AccountBalanceAdjustmentFixedBuilder {
	return AccountBalanceAdjustmentFixedBuilder{m.Fixed}
}

// Finish
func (m AccountBalanceAdjustmentFixedBuilder) Finish() AccountBalanceAdjustmentFixed {
	return AccountBalanceAdjustmentFixed{m.Fixed}
}

// RequestID
func (m AccountBalanceAdjustmentFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m AccountBalanceAdjustmentFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// TradeAccount
func (m AccountBalanceAdjustmentFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// TradeAccount
func (m AccountBalanceAdjustmentFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// CreditAmount
func (m AccountBalanceAdjustmentFixed) CreditAmount() float64 {
	return message.Float64Fixed(m.Unsafe(), 48, 40)
}

// CreditAmount
func (m AccountBalanceAdjustmentFixedBuilder) CreditAmount() float64 {
	return message.Float64Fixed(m.Unsafe(), 48, 40)
}

// DebitAmount
func (m AccountBalanceAdjustmentFixed) DebitAmount() float64 {
	return message.Float64Fixed(m.Unsafe(), 56, 48)
}

// DebitAmount
func (m AccountBalanceAdjustmentFixedBuilder) DebitAmount() float64 {
	return message.Float64Fixed(m.Unsafe(), 56, 48)
}

// Currency
func (m AccountBalanceAdjustmentFixed) Currency() string {
	return message.StringFixed(m.Unsafe(), 64, 56)
}

// Currency
func (m AccountBalanceAdjustmentFixedBuilder) Currency() string {
	return message.StringFixed(m.Unsafe(), 64, 56)
}

// Reason
func (m AccountBalanceAdjustmentFixed) Reason() string {
	return message.StringFixed(m.Unsafe(), 160, 64)
}

// Reason
func (m AccountBalanceAdjustmentFixedBuilder) Reason() string {
	return message.StringFixed(m.Unsafe(), 160, 64)
}

// SetRequestID
func (m AccountBalanceAdjustmentFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetTradeAccount
func (m AccountBalanceAdjustmentFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetCreditAmount
func (m AccountBalanceAdjustmentFixedBuilder) SetCreditAmount(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 48, 40, value)
}

// SetDebitAmount
func (m AccountBalanceAdjustmentFixedBuilder) SetDebitAmount(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 56, 48, value)
}

// SetCurrency
func (m AccountBalanceAdjustmentFixedBuilder) SetCurrency(value string) {
	message.SetStringFixed(m.Unsafe(), 64, 56, value)
}

// SetReason
func (m AccountBalanceAdjustmentFixedBuilder) SetReason(value string) {
	message.SetStringFixed(m.Unsafe(), 160, 64, value)
}

// Copy
func (m AccountBalanceAdjustmentFixed) Copy(to AccountBalanceAdjustmentFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// Copy
func (m AccountBalanceAdjustmentFixedBuilder) Copy(to AccountBalanceAdjustmentFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyPointer
func (m AccountBalanceAdjustmentFixed) CopyPointer(to AccountBalanceAdjustmentFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyPointer
func (m AccountBalanceAdjustmentFixedBuilder) CopyPointer(to AccountBalanceAdjustmentFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyToPointer
func (m AccountBalanceAdjustmentFixed) CopyToPointer(to AccountBalanceAdjustmentPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyToPointer
func (m AccountBalanceAdjustmentFixedBuilder) CopyToPointer(to AccountBalanceAdjustmentPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyTo
func (m AccountBalanceAdjustmentFixed) CopyTo(to AccountBalanceAdjustmentBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyTo
func (m AccountBalanceAdjustmentFixedBuilder) CopyTo(to AccountBalanceAdjustmentBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}
