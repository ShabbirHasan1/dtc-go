package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                    = IndividualTradePositionSize  (54)
//     Type                    = INDIVIDUAL_TRADE_POSITION  (10112)
//     BaseSize                = IndividualTradePositionSize  (54)
//     Symbol                  = ""
//     IsSimulated             = 0
//     TradeAccount            = ""
//     Quantity                = 0
//     PositionPrice           = 0
//     OpenProfitLoss          = 0
//     TradeDateTime           = 0
//     FillExecutionIdentifier = ""
//     IsSnapshot              = 0
//     IsFirstMessageInBatch   = 0
//     IsLastMessageInBatch    = 0
// }
var _IndividualTradePositionDefault = []byte{54, 0, 128, 39, 54, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const IndividualTradePositionSize = 54

type IndividualTradePosition struct {
	message.VLS
}

type IndividualTradePositionBuilder struct {
	message.VLSBuilder
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

// Clear
// {
//     Size                    = IndividualTradePositionSize  (54)
//     Type                    = INDIVIDUAL_TRADE_POSITION  (10112)
//     BaseSize                = IndividualTradePositionSize  (54)
//     Symbol                  = ""
//     IsSimulated             = 0
//     TradeAccount            = ""
//     Quantity                = 0
//     PositionPrice           = 0
//     OpenProfitLoss          = 0
//     TradeDateTime           = 0
//     FillExecutionIdentifier = ""
//     IsSnapshot              = 0
//     IsFirstMessageInBatch   = 0
//     IsLastMessageInBatch    = 0
// }
func (m IndividualTradePositionBuilder) Clear() {
	m.Unsafe().SetBytes(0, _IndividualTradePositionDefault)
}

// ToBuilder
func (m IndividualTradePosition) ToBuilder() IndividualTradePositionBuilder {
	return IndividualTradePositionBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m IndividualTradePositionBuilder) Finish() IndividualTradePosition {
	return IndividualTradePosition{m.VLSBuilder.Finish()}
}

// Symbol
func (m IndividualTradePosition) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol
func (m IndividualTradePositionBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// IsSimulated
func (m IndividualTradePosition) IsSimulated() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// IsSimulated
func (m IndividualTradePositionBuilder) IsSimulated() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// TradeAccount
func (m IndividualTradePosition) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// TradeAccount
func (m IndividualTradePositionBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// Quantity
func (m IndividualTradePosition) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 23, 15)
}

// Quantity
func (m IndividualTradePositionBuilder) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 23, 15)
}

// PositionPrice
func (m IndividualTradePosition) PositionPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 31, 23)
}

// PositionPrice
func (m IndividualTradePositionBuilder) PositionPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 31, 23)
}

// OpenProfitLoss
func (m IndividualTradePosition) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 39, 31)
}

// OpenProfitLoss
func (m IndividualTradePositionBuilder) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 39, 31)
}

// TradeDateTime
func (m IndividualTradePosition) TradeDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 47, 39)
}

// TradeDateTime
func (m IndividualTradePositionBuilder) TradeDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 47, 39)
}

// FillExecutionIdentifier
func (m IndividualTradePosition) FillExecutionIdentifier() string {
	return message.StringVLS(m.Unsafe(), 51, 47)
}

// FillExecutionIdentifier
func (m IndividualTradePositionBuilder) FillExecutionIdentifier() string {
	return message.StringVLS(m.Unsafe(), 51, 47)
}

// IsSnapshot
func (m IndividualTradePosition) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Unsafe(), 52, 51)
}

// IsSnapshot
func (m IndividualTradePositionBuilder) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Unsafe(), 52, 51)
}

// IsFirstMessageInBatch
func (m IndividualTradePosition) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 53, 52)
}

// IsFirstMessageInBatch
func (m IndividualTradePositionBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 53, 52)
}

// IsLastMessageInBatch
func (m IndividualTradePosition) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 54, 53)
}

// IsLastMessageInBatch
func (m IndividualTradePositionBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 54, 53)
}

// SetSymbol
func (m *IndividualTradePositionBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetIsSimulated
func (m IndividualTradePositionBuilder) SetIsSimulated(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 11, 10, value)
}

// SetTradeAccount
func (m *IndividualTradePositionBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 15, 11, value)
}

// SetQuantity
func (m IndividualTradePositionBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 23, 15, value)
}

// SetPositionPrice
func (m IndividualTradePositionBuilder) SetPositionPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 31, 23, value)
}

// SetOpenProfitLoss
func (m IndividualTradePositionBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 39, 31, value)
}

// SetTradeDateTime
func (m IndividualTradePositionBuilder) SetTradeDateTime(value int64) {
	message.SetInt64VLS(m.Unsafe(), 47, 39, value)
}

// SetFillExecutionIdentifier
func (m *IndividualTradePositionBuilder) SetFillExecutionIdentifier(value string) {
	message.SetStringVLS(&m.VLSBuilder, 51, 47, value)
}

// SetIsSnapshot
func (m IndividualTradePositionBuilder) SetIsSnapshot(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 52, 51, value)
}

// SetIsFirstMessageInBatch
func (m IndividualTradePositionBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 53, 52, value)
}

// SetIsLastMessageInBatch
func (m IndividualTradePositionBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 54, 53, value)
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
