package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalAccountBalanceResponseSize = 56

const HistoricalAccountBalanceResponseFixedSize = 232

// {
//     Size              = HistoricalAccountBalanceResponseSize  (56)
//     Type              = HISTORICAL_ACCOUNT_BALANCE_RESPONSE  (606)
//     BaseSize          = HistoricalAccountBalanceResponseSize  (56)
//     RequestID         = 0
//     DateTime          = 0
//     CashBalance       = 0
//     AccountCurrency   = ""
//     TradeAccount      = ""
//     IsFinalResponse   = false
//     NoAccountBalances = 0
//     InfoText          = ""
//     TransactionId     = ""
// }
var _HistoricalAccountBalanceResponseDefault = []byte{56, 0, 94, 2, 56, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size              = HistoricalAccountBalanceResponseFixedSize  (232)
//     Type              = HISTORICAL_ACCOUNT_BALANCE_RESPONSE  (606)
//     RequestID         = 0
//     DateTime          = 0
//     CashBalance       = 0
//     AccountCurrency   = ""
//     TradeAccount      = ""
//     IsFinalResponse   = false
//     NoAccountBalances = 0
//     InfoText          = ""
//     TransactionId     = ""
// }
var _HistoricalAccountBalanceResponseFixedDefault = []byte{232, 0, 94, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HistoricalAccountBalanceResponse struct {
	message.VLS
}

type HistoricalAccountBalanceResponseBuilder struct {
	message.VLSBuilder
}

type HistoricalAccountBalanceResponseFixed struct {
	message.Fixed
}

type HistoricalAccountBalanceResponseFixedBuilder struct {
	message.Fixed
}

type HistoricalAccountBalanceResponsePointer struct {
	message.VLSPointer
}

type HistoricalAccountBalanceResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

type HistoricalAccountBalanceResponseFixedPointer struct {
	message.FixedPointer
}

type HistoricalAccountBalanceResponseFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocHistoricalAccountBalanceResponse() HistoricalAccountBalanceResponsePointerBuilder {
	a := HistoricalAccountBalanceResponsePointerBuilder{message.AllocVLSBuilder(56)}
	a.Ptr.SetBytes(0, _HistoricalAccountBalanceResponseDefault)
	return a
}

func AllocHistoricalAccountBalanceResponseFrom(b []byte) HistoricalAccountBalanceResponsePointer {
	return HistoricalAccountBalanceResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size              = HistoricalAccountBalanceResponseSize  (56)
//     Type              = HISTORICAL_ACCOUNT_BALANCE_RESPONSE  (606)
//     BaseSize          = HistoricalAccountBalanceResponseSize  (56)
//     RequestID         = 0
//     DateTime          = 0
//     CashBalance       = 0
//     AccountCurrency   = ""
//     TradeAccount      = ""
//     IsFinalResponse   = false
//     NoAccountBalances = 0
//     InfoText          = ""
//     TransactionId     = ""
// }
func (m HistoricalAccountBalanceResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalAccountBalanceResponseDefault)
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
//     IsFinalResponse   = false
//     NoAccountBalances = 0
//     InfoText          = ""
//     TransactionId     = ""
// }
func (m HistoricalAccountBalanceResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalAccountBalanceResponseFixedDefault)
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
//     IsFinalResponse   = false
//     NoAccountBalances = 0
//     InfoText          = ""
//     TransactionId     = ""
// }
func (m HistoricalAccountBalanceResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalAccountBalanceResponseDefault)
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
//     IsFinalResponse   = false
//     NoAccountBalances = 0
//     InfoText          = ""
//     TransactionId     = ""
// }
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalAccountBalanceResponseFixedDefault)
}

// ToBuilder
func (m HistoricalAccountBalanceResponse) ToBuilder() HistoricalAccountBalanceResponseBuilder {
	return HistoricalAccountBalanceResponseBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalAccountBalanceResponseFixed) ToBuilder() HistoricalAccountBalanceResponseFixedBuilder {
	return HistoricalAccountBalanceResponseFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalAccountBalanceResponsePointer) ToBuilder() HistoricalAccountBalanceResponsePointerBuilder {
	return HistoricalAccountBalanceResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalAccountBalanceResponseFixedPointer) ToBuilder() HistoricalAccountBalanceResponseFixedPointerBuilder {
	return HistoricalAccountBalanceResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalAccountBalanceResponseBuilder) Finish() HistoricalAccountBalanceResponse {
	return HistoricalAccountBalanceResponse{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalAccountBalanceResponseFixedBuilder) Finish() HistoricalAccountBalanceResponseFixed {
	return HistoricalAccountBalanceResponseFixed{m.Fixed}
}

// Finish
func (m *HistoricalAccountBalanceResponsePointerBuilder) Finish() HistoricalAccountBalanceResponsePointer {
	return HistoricalAccountBalanceResponsePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalAccountBalanceResponseFixedPointerBuilder) Finish() HistoricalAccountBalanceResponseFixedPointer {
	return HistoricalAccountBalanceResponseFixedPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalAccountBalanceResponse) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m HistoricalAccountBalanceResponseBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
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
func (m HistoricalAccountBalanceResponse) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 24, 16))
}

// DateTime
func (m HistoricalAccountBalanceResponseBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 24, 16))
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
func (m HistoricalAccountBalanceResponse) CashBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
}

// CashBalance
func (m HistoricalAccountBalanceResponseBuilder) CashBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
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
func (m HistoricalAccountBalanceResponse) AccountCurrency() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
}

// AccountCurrency
func (m HistoricalAccountBalanceResponseBuilder) AccountCurrency() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
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
func (m HistoricalAccountBalanceResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// TradeAccount
func (m HistoricalAccountBalanceResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
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
func (m HistoricalAccountBalanceResponse) IsFinalResponse() bool {
	return message.BoolVLS(m.Unsafe(), 41, 40)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponseBuilder) IsFinalResponse() bool {
	return message.BoolVLS(m.Unsafe(), 41, 40)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponsePointer) IsFinalResponse() bool {
	return message.BoolVLS(m.Ptr, 41, 40)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponsePointerBuilder) IsFinalResponse() bool {
	return message.BoolVLS(m.Ptr, 41, 40)
}

// NoAccountBalances
func (m HistoricalAccountBalanceResponse) NoAccountBalances() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// NoAccountBalances
func (m HistoricalAccountBalanceResponseBuilder) NoAccountBalances() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
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
func (m HistoricalAccountBalanceResponse) InfoText() string {
	return message.StringVLS(m.Unsafe(), 46, 42)
}

// InfoText
func (m HistoricalAccountBalanceResponseBuilder) InfoText() string {
	return message.StringVLS(m.Unsafe(), 46, 42)
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
func (m HistoricalAccountBalanceResponse) TransactionId() string {
	return message.StringVLS(m.Unsafe(), 50, 46)
}

// TransactionId
func (m HistoricalAccountBalanceResponseBuilder) TransactionId() string {
	return message.StringVLS(m.Unsafe(), 50, 46)
}

// TransactionId
func (m HistoricalAccountBalanceResponsePointer) TransactionId() string {
	return message.StringVLS(m.Ptr, 50, 46)
}

// TransactionId
func (m HistoricalAccountBalanceResponsePointerBuilder) TransactionId() string {
	return message.StringVLS(m.Ptr, 50, 46)
}

// RequestID
func (m HistoricalAccountBalanceResponseFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalAccountBalanceResponseFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m HistoricalAccountBalanceResponseFixed) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 16, 8))
}

// DateTime
func (m HistoricalAccountBalanceResponseFixedBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 16, 8))
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
func (m HistoricalAccountBalanceResponseFixed) CashBalance() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// CashBalance
func (m HistoricalAccountBalanceResponseFixedBuilder) CashBalance() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
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
func (m HistoricalAccountBalanceResponseFixed) AccountCurrency() string {
	return message.StringFixed(m.Unsafe(), 32, 24)
}

// AccountCurrency
func (m HistoricalAccountBalanceResponseFixedBuilder) AccountCurrency() string {
	return message.StringFixed(m.Unsafe(), 32, 24)
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
func (m HistoricalAccountBalanceResponseFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 64, 32)
}

// TradeAccount
func (m HistoricalAccountBalanceResponseFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 64, 32)
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
func (m HistoricalAccountBalanceResponseFixed) IsFinalResponse() bool {
	return message.BoolFixed(m.Unsafe(), 65, 64)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponseFixedBuilder) IsFinalResponse() bool {
	return message.BoolFixed(m.Unsafe(), 65, 64)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponseFixedPointer) IsFinalResponse() bool {
	return message.BoolFixed(m.Ptr, 65, 64)
}

// IsFinalResponse
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) IsFinalResponse() bool {
	return message.BoolFixed(m.Ptr, 65, 64)
}

// NoAccountBalances
func (m HistoricalAccountBalanceResponseFixed) NoAccountBalances() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 66, 65)
}

// NoAccountBalances
func (m HistoricalAccountBalanceResponseFixedBuilder) NoAccountBalances() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 66, 65)
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
func (m HistoricalAccountBalanceResponseFixed) InfoText() string {
	return message.StringFixed(m.Unsafe(), 162, 66)
}

// InfoText
func (m HistoricalAccountBalanceResponseFixedBuilder) InfoText() string {
	return message.StringFixed(m.Unsafe(), 162, 66)
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
func (m HistoricalAccountBalanceResponseFixed) TransactionId() string {
	return message.StringFixed(m.Unsafe(), 226, 162)
}

// TransactionId
func (m HistoricalAccountBalanceResponseFixedBuilder) TransactionId() string {
	return message.StringFixed(m.Unsafe(), 226, 162)
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
func (m HistoricalAccountBalanceResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID
func (m HistoricalAccountBalanceResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetDateTime
func (m HistoricalAccountBalanceResponseBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Unsafe(), 24, 16, float64(value))
}

// SetDateTime
func (m HistoricalAccountBalanceResponsePointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Ptr, 24, 16, float64(value))
}

// SetCashBalance
func (m HistoricalAccountBalanceResponseBuilder) SetCashBalance(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 32, 24, value)
}

// SetCashBalance
func (m HistoricalAccountBalanceResponsePointerBuilder) SetCashBalance(value float64) {
	message.SetFloat64VLS(m.Ptr, 32, 24, value)
}

// SetAccountCurrency
func (m *HistoricalAccountBalanceResponseBuilder) SetAccountCurrency(value string) {
	message.SetStringVLS(&m.VLSBuilder, 36, 32, value)
}

// SetAccountCurrency
func (m *HistoricalAccountBalanceResponsePointerBuilder) SetAccountCurrency(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 36, 32, value)
}

// SetTradeAccount
func (m *HistoricalAccountBalanceResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 40, 36, value)
}

// SetTradeAccount
func (m *HistoricalAccountBalanceResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 40, 36, value)
}

// SetIsFinalResponse
func (m HistoricalAccountBalanceResponseBuilder) SetIsFinalResponse(value bool) {
	message.SetBoolVLS(m.Unsafe(), 41, 40, value)
}

// SetIsFinalResponse
func (m HistoricalAccountBalanceResponsePointerBuilder) SetIsFinalResponse(value bool) {
	message.SetBoolVLS(m.Ptr, 41, 40, value)
}

// SetNoAccountBalances
func (m HistoricalAccountBalanceResponseBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 42, 41, value)
}

// SetNoAccountBalances
func (m HistoricalAccountBalanceResponsePointerBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8VLS(m.Ptr, 42, 41, value)
}

// SetInfoText
func (m *HistoricalAccountBalanceResponseBuilder) SetInfoText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 46, 42, value)
}

// SetInfoText
func (m *HistoricalAccountBalanceResponsePointerBuilder) SetInfoText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 46, 42, value)
}

// SetTransactionId
func (m *HistoricalAccountBalanceResponseBuilder) SetTransactionId(value string) {
	message.SetStringVLS(&m.VLSBuilder, 50, 46, value)
}

// SetTransactionId
func (m *HistoricalAccountBalanceResponsePointerBuilder) SetTransactionId(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 50, 46, value)
}

// SetRequestID
func (m HistoricalAccountBalanceResponseFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetDateTime
func (m HistoricalAccountBalanceResponseFixedBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, float64(value))
}

// SetDateTime
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, float64(value))
}

// SetCashBalance
func (m HistoricalAccountBalanceResponseFixedBuilder) SetCashBalance(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetCashBalance
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetCashBalance(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetAccountCurrency
func (m HistoricalAccountBalanceResponseFixedBuilder) SetAccountCurrency(value string) {
	message.SetStringFixed(m.Unsafe(), 32, 24, value)
}

// SetAccountCurrency
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetAccountCurrency(value string) {
	message.SetStringFixed(m.Ptr, 32, 24, value)
}

// SetTradeAccount
func (m HistoricalAccountBalanceResponseFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 64, 32, value)
}

// SetTradeAccount
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 64, 32, value)
}

// SetIsFinalResponse
func (m HistoricalAccountBalanceResponseFixedBuilder) SetIsFinalResponse(value bool) {
	message.SetBoolFixed(m.Unsafe(), 65, 64, value)
}

// SetIsFinalResponse
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetIsFinalResponse(value bool) {
	message.SetBoolFixed(m.Ptr, 65, 64, value)
}

// SetNoAccountBalances
func (m HistoricalAccountBalanceResponseFixedBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 66, 65, value)
}

// SetNoAccountBalances
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8Fixed(m.Ptr, 66, 65, value)
}

// SetInfoText
func (m HistoricalAccountBalanceResponseFixedBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Unsafe(), 162, 66, value)
}

// SetInfoText
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Ptr, 162, 66, value)
}

// SetTransactionId
func (m HistoricalAccountBalanceResponseFixedBuilder) SetTransactionId(value string) {
	message.SetStringFixed(m.Unsafe(), 226, 162, value)
}

// SetTransactionId
func (m HistoricalAccountBalanceResponseFixedPointerBuilder) SetTransactionId(value string) {
	message.SetStringFixed(m.Ptr, 226, 162, value)
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
