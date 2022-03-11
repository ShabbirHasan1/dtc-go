package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const IndividualTradePositionSize = 54

// {
//     Size                    = IndividualTradePositionSize  (54)
//     Type                    = INDIVIDUAL_TRADE_POSITION  (10112)
//     BaseSize                = IndividualTradePositionSize  (54)
//     Symbol                  = ""
//     IsSimulated             = false
//     TradeAccount            = ""
//     Quantity                = 0
//     PositionPrice           = 0
//     OpenProfitLoss          = 0
//     TradeDateTime           = 0
//     FillExecutionIdentifier = ""
//     IsSnapshot              = false
//     IsFirstMessageInBatch   = false
//     IsLastMessageInBatch    = false
// }
var _IndividualTradePositionDefault = []byte{54, 0, 128, 39, 54, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type IndividualTradePosition struct {
	message.VLS
}

type IndividualTradePositionBuilder struct {
	message.VLSBuilder
}

type IndividualTradePositionPointer struct {
	message.VLSPointer
}

type IndividualTradePositionPointerBuilder struct {
	message.VLSPointerBuilder
}

func NewIndividualTradePositionFrom(b []byte) IndividualTradePosition {
	return IndividualTradePosition{VLS: message.NewVLSFrom(b)}
}

func WrapIndividualTradePosition(b []byte) IndividualTradePosition {
	return IndividualTradePosition{VLS: message.WrapVLS(b)}
}

func NewIndividualTradePosition() IndividualTradePositionBuilder {
	a := IndividualTradePositionBuilder{message.NewVLSBuilder(54)}
	a.Unsafe().SetBytes(0, _IndividualTradePositionDefault)
	return a
}

func AllocIndividualTradePosition() IndividualTradePositionPointerBuilder {
	a := IndividualTradePositionPointerBuilder{message.AllocVLSBuilder(54)}
	a.Ptr.SetBytes(0, _IndividualTradePositionDefault)
	return a
}

func AllocIndividualTradePositionFrom(b []byte) IndividualTradePositionPointer {
	return IndividualTradePositionPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                    = IndividualTradePositionSize  (54)
//     Type                    = INDIVIDUAL_TRADE_POSITION  (10112)
//     BaseSize                = IndividualTradePositionSize  (54)
//     Symbol                  = ""
//     IsSimulated             = false
//     TradeAccount            = ""
//     Quantity                = 0
//     PositionPrice           = 0
//     OpenProfitLoss          = 0
//     TradeDateTime           = 0
//     FillExecutionIdentifier = ""
//     IsSnapshot              = false
//     IsFirstMessageInBatch   = false
//     IsLastMessageInBatch    = false
// }
func (m IndividualTradePositionBuilder) Clear() {
	m.Unsafe().SetBytes(0, _IndividualTradePositionDefault)
}

// Clear
// {
//     Size                    = IndividualTradePositionSize  (54)
//     Type                    = INDIVIDUAL_TRADE_POSITION  (10112)
//     BaseSize                = IndividualTradePositionSize  (54)
//     Symbol                  = ""
//     IsSimulated             = false
//     TradeAccount            = ""
//     Quantity                = 0
//     PositionPrice           = 0
//     OpenProfitLoss          = 0
//     TradeDateTime           = 0
//     FillExecutionIdentifier = ""
//     IsSnapshot              = false
//     IsFirstMessageInBatch   = false
//     IsLastMessageInBatch    = false
// }
func (m IndividualTradePositionPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _IndividualTradePositionDefault)
}

// ToBuilder
func (m IndividualTradePosition) ToBuilder() IndividualTradePositionBuilder {
	return IndividualTradePositionBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m IndividualTradePositionPointer) ToBuilder() IndividualTradePositionPointerBuilder {
	return IndividualTradePositionPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m IndividualTradePositionBuilder) Finish() IndividualTradePosition {
	return IndividualTradePosition{m.VLSBuilder.Finish()}
}

// Finish
func (m *IndividualTradePositionPointerBuilder) Finish() IndividualTradePositionPointer {
	return IndividualTradePositionPointer{m.VLSPointerBuilder.Finish()}
}

// Symbol
func (m IndividualTradePosition) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol
func (m IndividualTradePositionBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol
func (m IndividualTradePositionPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Symbol
func (m IndividualTradePositionPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// IsSimulated
func (m IndividualTradePosition) IsSimulated() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsSimulated
func (m IndividualTradePositionBuilder) IsSimulated() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsSimulated
func (m IndividualTradePositionPointer) IsSimulated() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// IsSimulated
func (m IndividualTradePositionPointerBuilder) IsSimulated() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// TradeAccount
func (m IndividualTradePosition) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// TradeAccount
func (m IndividualTradePositionBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// TradeAccount
func (m IndividualTradePositionPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// TradeAccount
func (m IndividualTradePositionPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// Quantity
func (m IndividualTradePosition) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 23, 15)
}

// Quantity
func (m IndividualTradePositionBuilder) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 23, 15)
}

// Quantity
func (m IndividualTradePositionPointer) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 23, 15)
}

// Quantity
func (m IndividualTradePositionPointerBuilder) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 23, 15)
}

// PositionPrice
func (m IndividualTradePosition) PositionPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 31, 23)
}

// PositionPrice
func (m IndividualTradePositionBuilder) PositionPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 31, 23)
}

// PositionPrice
func (m IndividualTradePositionPointer) PositionPrice() float64 {
	return message.Float64VLS(m.Ptr, 31, 23)
}

// PositionPrice
func (m IndividualTradePositionPointerBuilder) PositionPrice() float64 {
	return message.Float64VLS(m.Ptr, 31, 23)
}

// OpenProfitLoss
func (m IndividualTradePosition) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 39, 31)
}

// OpenProfitLoss
func (m IndividualTradePositionBuilder) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 39, 31)
}

// OpenProfitLoss
func (m IndividualTradePositionPointer) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 39, 31)
}

// OpenProfitLoss
func (m IndividualTradePositionPointerBuilder) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 39, 31)
}

// TradeDateTime
func (m IndividualTradePosition) TradeDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 47, 39)
}

// TradeDateTime
func (m IndividualTradePositionBuilder) TradeDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 47, 39)
}

// TradeDateTime
func (m IndividualTradePositionPointer) TradeDateTime() int64 {
	return message.Int64VLS(m.Ptr, 47, 39)
}

// TradeDateTime
func (m IndividualTradePositionPointerBuilder) TradeDateTime() int64 {
	return message.Int64VLS(m.Ptr, 47, 39)
}

// FillExecutionIdentifier
func (m IndividualTradePosition) FillExecutionIdentifier() string {
	return message.StringVLS(m.Unsafe(), 51, 47)
}

// FillExecutionIdentifier
func (m IndividualTradePositionBuilder) FillExecutionIdentifier() string {
	return message.StringVLS(m.Unsafe(), 51, 47)
}

// FillExecutionIdentifier
func (m IndividualTradePositionPointer) FillExecutionIdentifier() string {
	return message.StringVLS(m.Ptr, 51, 47)
}

// FillExecutionIdentifier
func (m IndividualTradePositionPointerBuilder) FillExecutionIdentifier() string {
	return message.StringVLS(m.Ptr, 51, 47)
}

// IsSnapshot
func (m IndividualTradePosition) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 52, 51)
}

// IsSnapshot
func (m IndividualTradePositionBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 52, 51)
}

// IsSnapshot
func (m IndividualTradePositionPointer) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 52, 51)
}

// IsSnapshot
func (m IndividualTradePositionPointerBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 52, 51)
}

// IsFirstMessageInBatch
func (m IndividualTradePosition) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 53, 52)
}

// IsFirstMessageInBatch
func (m IndividualTradePositionBuilder) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 53, 52)
}

// IsFirstMessageInBatch
func (m IndividualTradePositionPointer) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 53, 52)
}

// IsFirstMessageInBatch
func (m IndividualTradePositionPointerBuilder) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 53, 52)
}

// IsLastMessageInBatch
func (m IndividualTradePosition) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 54, 53)
}

// IsLastMessageInBatch
func (m IndividualTradePositionBuilder) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 54, 53)
}

// IsLastMessageInBatch
func (m IndividualTradePositionPointer) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 54, 53)
}

// IsLastMessageInBatch
func (m IndividualTradePositionPointerBuilder) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 54, 53)
}

// SetSymbol
func (m *IndividualTradePositionBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetSymbol
func (m *IndividualTradePositionPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetIsSimulated
func (m IndividualTradePositionBuilder) SetIsSimulated(value bool) {
	message.SetBoolVLS(m.Unsafe(), 11, 10, value)
}

// SetIsSimulated
func (m IndividualTradePositionPointerBuilder) SetIsSimulated(value bool) {
	message.SetBoolVLS(m.Ptr, 11, 10, value)
}

// SetTradeAccount
func (m *IndividualTradePositionBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 15, 11, value)
}

// SetTradeAccount
func (m *IndividualTradePositionPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetQuantity
func (m IndividualTradePositionBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 23, 15, value)
}

// SetQuantity
func (m IndividualTradePositionPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 23, 15, value)
}

// SetPositionPrice
func (m IndividualTradePositionBuilder) SetPositionPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 31, 23, value)
}

// SetPositionPrice
func (m IndividualTradePositionPointerBuilder) SetPositionPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 31, 23, value)
}

// SetOpenProfitLoss
func (m IndividualTradePositionBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 39, 31, value)
}

// SetOpenProfitLoss
func (m IndividualTradePositionPointerBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 39, 31, value)
}

// SetTradeDateTime
func (m IndividualTradePositionBuilder) SetTradeDateTime(value int64) {
	message.SetInt64VLS(m.Unsafe(), 47, 39, value)
}

// SetTradeDateTime
func (m IndividualTradePositionPointerBuilder) SetTradeDateTime(value int64) {
	message.SetInt64VLS(m.Ptr, 47, 39, value)
}

// SetFillExecutionIdentifier
func (m *IndividualTradePositionBuilder) SetFillExecutionIdentifier(value string) {
	message.SetStringVLS(&m.VLSBuilder, 51, 47, value)
}

// SetFillExecutionIdentifier
func (m *IndividualTradePositionPointerBuilder) SetFillExecutionIdentifier(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 51, 47, value)
}

// SetIsSnapshot
func (m IndividualTradePositionBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Unsafe(), 52, 51, value)
}

// SetIsSnapshot
func (m IndividualTradePositionPointerBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Ptr, 52, 51, value)
}

// SetIsFirstMessageInBatch
func (m IndividualTradePositionBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolVLS(m.Unsafe(), 53, 52, value)
}

// SetIsFirstMessageInBatch
func (m IndividualTradePositionPointerBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolVLS(m.Ptr, 53, 52, value)
}

// SetIsLastMessageInBatch
func (m IndividualTradePositionBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolVLS(m.Unsafe(), 54, 53, value)
}

// SetIsLastMessageInBatch
func (m IndividualTradePositionPointerBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolVLS(m.Ptr, 54, 53, value)
}

// Copy
func (m IndividualTradePosition) Copy(to IndividualTradePositionBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetQuantity(m.Quantity())
	to.SetPositionPrice(m.PositionPrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetTradeDateTime(m.TradeDateTime())
	to.SetFillExecutionIdentifier(m.FillExecutionIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// Copy
func (m IndividualTradePositionBuilder) Copy(to IndividualTradePositionBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetQuantity(m.Quantity())
	to.SetPositionPrice(m.PositionPrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetTradeDateTime(m.TradeDateTime())
	to.SetFillExecutionIdentifier(m.FillExecutionIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// CopyPointer
func (m IndividualTradePosition) CopyPointer(to IndividualTradePositionPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetQuantity(m.Quantity())
	to.SetPositionPrice(m.PositionPrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetTradeDateTime(m.TradeDateTime())
	to.SetFillExecutionIdentifier(m.FillExecutionIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// CopyPointer
func (m IndividualTradePositionBuilder) CopyPointer(to IndividualTradePositionPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetQuantity(m.Quantity())
	to.SetPositionPrice(m.PositionPrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetTradeDateTime(m.TradeDateTime())
	to.SetFillExecutionIdentifier(m.FillExecutionIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// Copy
func (m IndividualTradePositionPointer) Copy(to IndividualTradePositionBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetQuantity(m.Quantity())
	to.SetPositionPrice(m.PositionPrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetTradeDateTime(m.TradeDateTime())
	to.SetFillExecutionIdentifier(m.FillExecutionIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// Copy
func (m IndividualTradePositionPointerBuilder) Copy(to IndividualTradePositionBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetQuantity(m.Quantity())
	to.SetPositionPrice(m.PositionPrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetTradeDateTime(m.TradeDateTime())
	to.SetFillExecutionIdentifier(m.FillExecutionIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// CopyPointer
func (m IndividualTradePositionPointer) CopyPointer(to IndividualTradePositionPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetQuantity(m.Quantity())
	to.SetPositionPrice(m.PositionPrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetTradeDateTime(m.TradeDateTime())
	to.SetFillExecutionIdentifier(m.FillExecutionIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// CopyPointer
func (m IndividualTradePositionPointerBuilder) CopyPointer(to IndividualTradePositionPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetQuantity(m.Quantity())
	to.SetPositionPrice(m.PositionPrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetTradeDateTime(m.TradeDateTime())
	to.SetFillExecutionIdentifier(m.FillExecutionIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}
