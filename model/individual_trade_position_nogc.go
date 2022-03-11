package model

import (
	"github.com/moontrade/dtc-go/message"
)

type IndividualTradePositionPointer struct {
	message.VLSPointer
}

type IndividualTradePositionPointerBuilder struct {
	message.VLSPointerBuilder
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
func (m IndividualTradePositionPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _IndividualTradePositionDefault)
}

// ToBuilder
func (m IndividualTradePositionPointer) ToBuilder() IndividualTradePositionPointerBuilder {
	return IndividualTradePositionPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *IndividualTradePositionPointerBuilder) Finish() IndividualTradePositionPointer {
	return IndividualTradePositionPointer{m.VLSPointerBuilder.Finish()}
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
func (m IndividualTradePositionPointer) IsSimulated() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// IsSimulated
func (m IndividualTradePositionPointerBuilder) IsSimulated() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
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
func (m IndividualTradePositionPointer) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 23, 15)
}

// Quantity
func (m IndividualTradePositionPointerBuilder) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 23, 15)
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
func (m IndividualTradePositionPointer) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 39, 31)
}

// OpenProfitLoss
func (m IndividualTradePositionPointerBuilder) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 39, 31)
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
func (m IndividualTradePositionPointer) FillExecutionIdentifier() string {
	return message.StringVLS(m.Ptr, 51, 47)
}

// FillExecutionIdentifier
func (m IndividualTradePositionPointerBuilder) FillExecutionIdentifier() string {
	return message.StringVLS(m.Ptr, 51, 47)
}

// IsSnapshot
func (m IndividualTradePositionPointer) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Ptr, 52, 51)
}

// IsSnapshot
func (m IndividualTradePositionPointerBuilder) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Ptr, 52, 51)
}

// IsFirstMessageInBatch
func (m IndividualTradePositionPointer) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 53, 52)
}

// IsFirstMessageInBatch
func (m IndividualTradePositionPointerBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 53, 52)
}

// IsLastMessageInBatch
func (m IndividualTradePositionPointer) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 54, 53)
}

// IsLastMessageInBatch
func (m IndividualTradePositionPointerBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 54, 53)
}

// SetSymbol
func (m *IndividualTradePositionPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetIsSimulated
func (m IndividualTradePositionPointerBuilder) SetIsSimulated(value uint8) {
	message.SetUint8VLS(m.Ptr, 11, 10, value)
}

// SetTradeAccount
func (m *IndividualTradePositionPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetQuantity
func (m IndividualTradePositionPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 23, 15, value)
}

// SetPositionPrice
func (m IndividualTradePositionPointerBuilder) SetPositionPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 31, 23, value)
}

// SetOpenProfitLoss
func (m IndividualTradePositionPointerBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 39, 31, value)
}

// SetTradeDateTime
func (m IndividualTradePositionPointerBuilder) SetTradeDateTime(value int64) {
	message.SetInt64VLS(m.Ptr, 47, 39, value)
}

// SetFillExecutionIdentifier
func (m *IndividualTradePositionPointerBuilder) SetFillExecutionIdentifier(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 51, 47, value)
}

// SetIsSnapshot
func (m IndividualTradePositionPointerBuilder) SetIsSnapshot(value uint8) {
	message.SetUint8VLS(m.Ptr, 52, 51, value)
}

// SetIsFirstMessageInBatch
func (m IndividualTradePositionPointerBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Ptr, 53, 52, value)
}

// SetIsLastMessageInBatch
func (m IndividualTradePositionPointerBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Ptr, 54, 53, value)
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
