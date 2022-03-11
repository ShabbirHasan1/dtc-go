package model

import (
	"github.com/moontrade/dtc-go/message"
)

const AccountBalanceAdjustmentSize = 40

const AccountBalanceAdjustmentFixedSize = 160

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

type AccountBalanceAdjustment struct {
	message.VLS
}

type AccountBalanceAdjustmentBuilder struct {
	message.VLSBuilder
}

type AccountBalanceAdjustmentFixed struct {
	message.Fixed
}

type AccountBalanceAdjustmentFixedBuilder struct {
	message.Fixed
}

type AccountBalanceAdjustmentPointer struct {
	message.VLSPointer
}

type AccountBalanceAdjustmentPointerBuilder struct {
	message.VLSPointerBuilder
}

type AccountBalanceAdjustmentFixedPointer struct {
	message.FixedPointer
}

type AccountBalanceAdjustmentFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocAccountBalanceAdjustment() AccountBalanceAdjustmentPointerBuilder {
	a := AccountBalanceAdjustmentPointerBuilder{message.AllocVLSBuilder(40)}
	a.Ptr.SetBytes(0, _AccountBalanceAdjustmentDefault)
	return a
}

func AllocAccountBalanceAdjustmentFrom(b []byte) AccountBalanceAdjustmentPointer {
	return AccountBalanceAdjustmentPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m AccountBalanceAdjustment) ToBuilder() AccountBalanceAdjustmentBuilder {
	return AccountBalanceAdjustmentBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m AccountBalanceAdjustmentFixed) ToBuilder() AccountBalanceAdjustmentFixedBuilder {
	return AccountBalanceAdjustmentFixedBuilder{m.Fixed}
}

// ToBuilder
func (m AccountBalanceAdjustmentPointer) ToBuilder() AccountBalanceAdjustmentPointerBuilder {
	return AccountBalanceAdjustmentPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m AccountBalanceAdjustmentFixedPointer) ToBuilder() AccountBalanceAdjustmentFixedPointerBuilder {
	return AccountBalanceAdjustmentFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m AccountBalanceAdjustmentBuilder) Finish() AccountBalanceAdjustment {
	return AccountBalanceAdjustment{m.VLSBuilder.Finish()}
}

// Finish
func (m AccountBalanceAdjustmentFixedBuilder) Finish() AccountBalanceAdjustmentFixed {
	return AccountBalanceAdjustmentFixed{m.Fixed}
}

// Finish
func (m *AccountBalanceAdjustmentPointerBuilder) Finish() AccountBalanceAdjustmentPointer {
	return AccountBalanceAdjustmentPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *AccountBalanceAdjustmentFixedPointerBuilder) Finish() AccountBalanceAdjustmentFixedPointer {
	return AccountBalanceAdjustmentFixedPointer{m.FixedPointer}
}

// RequestID
func (m AccountBalanceAdjustment) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m AccountBalanceAdjustmentBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
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
func (m AccountBalanceAdjustment) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// TradeAccount
func (m AccountBalanceAdjustmentBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
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
func (m AccountBalanceAdjustment) CreditAmount() float64 {
	return message.Float64VLS(m.Unsafe(), 24, 16)
}

// CreditAmount
func (m AccountBalanceAdjustmentBuilder) CreditAmount() float64 {
	return message.Float64VLS(m.Unsafe(), 24, 16)
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
func (m AccountBalanceAdjustment) DebitAmount() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
}

// DebitAmount
func (m AccountBalanceAdjustmentBuilder) DebitAmount() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
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
func (m AccountBalanceAdjustment) Currency() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
}

// Currency
func (m AccountBalanceAdjustmentBuilder) Currency() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
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
func (m AccountBalanceAdjustment) Reason() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// Reason
func (m AccountBalanceAdjustmentBuilder) Reason() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// Reason
func (m AccountBalanceAdjustmentPointer) Reason() string {
	return message.StringVLS(m.Ptr, 40, 36)
}

// Reason
func (m AccountBalanceAdjustmentPointerBuilder) Reason() string {
	return message.StringVLS(m.Ptr, 40, 36)
}

// RequestID
func (m AccountBalanceAdjustmentFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m AccountBalanceAdjustmentFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m AccountBalanceAdjustmentFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// TradeAccount
func (m AccountBalanceAdjustmentFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
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
func (m AccountBalanceAdjustmentFixed) CreditAmount() float64 {
	return message.Float64Fixed(m.Unsafe(), 48, 40)
}

// CreditAmount
func (m AccountBalanceAdjustmentFixedBuilder) CreditAmount() float64 {
	return message.Float64Fixed(m.Unsafe(), 48, 40)
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
func (m AccountBalanceAdjustmentFixed) DebitAmount() float64 {
	return message.Float64Fixed(m.Unsafe(), 56, 48)
}

// DebitAmount
func (m AccountBalanceAdjustmentFixedBuilder) DebitAmount() float64 {
	return message.Float64Fixed(m.Unsafe(), 56, 48)
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
func (m AccountBalanceAdjustmentFixed) Currency() string {
	return message.StringFixed(m.Unsafe(), 64, 56)
}

// Currency
func (m AccountBalanceAdjustmentFixedBuilder) Currency() string {
	return message.StringFixed(m.Unsafe(), 64, 56)
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
func (m AccountBalanceAdjustmentFixed) Reason() string {
	return message.StringFixed(m.Unsafe(), 160, 64)
}

// Reason
func (m AccountBalanceAdjustmentFixedBuilder) Reason() string {
	return message.StringFixed(m.Unsafe(), 160, 64)
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
func (m AccountBalanceAdjustmentBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID
func (m AccountBalanceAdjustmentPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetTradeAccount
func (m *AccountBalanceAdjustmentBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetTradeAccount
func (m *AccountBalanceAdjustmentPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetCreditAmount
func (m AccountBalanceAdjustmentBuilder) SetCreditAmount(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 24, 16, value)
}

// SetCreditAmount
func (m AccountBalanceAdjustmentPointerBuilder) SetCreditAmount(value float64) {
	message.SetFloat64VLS(m.Ptr, 24, 16, value)
}

// SetDebitAmount
func (m AccountBalanceAdjustmentBuilder) SetDebitAmount(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 32, 24, value)
}

// SetDebitAmount
func (m AccountBalanceAdjustmentPointerBuilder) SetDebitAmount(value float64) {
	message.SetFloat64VLS(m.Ptr, 32, 24, value)
}

// SetCurrency
func (m *AccountBalanceAdjustmentBuilder) SetCurrency(value string) {
	message.SetStringVLS(&m.VLSBuilder, 36, 32, value)
}

// SetCurrency
func (m *AccountBalanceAdjustmentPointerBuilder) SetCurrency(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 36, 32, value)
}

// SetReason
func (m *AccountBalanceAdjustmentBuilder) SetReason(value string) {
	message.SetStringVLS(&m.VLSBuilder, 40, 36, value)
}

// SetReason
func (m *AccountBalanceAdjustmentPointerBuilder) SetReason(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 40, 36, value)
}

// SetRequestID
func (m AccountBalanceAdjustmentFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m AccountBalanceAdjustmentFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTradeAccount
func (m AccountBalanceAdjustmentFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetTradeAccount
func (m AccountBalanceAdjustmentFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetCreditAmount
func (m AccountBalanceAdjustmentFixedBuilder) SetCreditAmount(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 48, 40, value)
}

// SetCreditAmount
func (m AccountBalanceAdjustmentFixedPointerBuilder) SetCreditAmount(value float64) {
	message.SetFloat64Fixed(m.Ptr, 48, 40, value)
}

// SetDebitAmount
func (m AccountBalanceAdjustmentFixedBuilder) SetDebitAmount(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 56, 48, value)
}

// SetDebitAmount
func (m AccountBalanceAdjustmentFixedPointerBuilder) SetDebitAmount(value float64) {
	message.SetFloat64Fixed(m.Ptr, 56, 48, value)
}

// SetCurrency
func (m AccountBalanceAdjustmentFixedBuilder) SetCurrency(value string) {
	message.SetStringFixed(m.Unsafe(), 64, 56, value)
}

// SetCurrency
func (m AccountBalanceAdjustmentFixedPointerBuilder) SetCurrency(value string) {
	message.SetStringFixed(m.Ptr, 64, 56, value)
}

// SetReason
func (m AccountBalanceAdjustmentFixedBuilder) SetReason(value string) {
	message.SetStringFixed(m.Unsafe(), 160, 64, value)
}

// SetReason
func (m AccountBalanceAdjustmentFixedPointerBuilder) SetReason(value string) {
	message.SetStringFixed(m.Ptr, 160, 64, value)
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
