package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataSymbolLimitsResponsePointer struct {
	message.VLSPointer
}

type TradeAccountDataSymbolLimitsResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeAccountDataSymbolLimitsResponse() TradeAccountDataSymbolLimitsResponsePointerBuilder {
	a := TradeAccountDataSymbolLimitsResponsePointerBuilder{message.AllocVLSBuilder(47)}
	a.Ptr.SetBytes(0, _TradeAccountDataSymbolLimitsResponseDefault)
	return a
}

func AllocTradeAccountDataSymbolLimitsResponseFrom(b []byte) TradeAccountDataSymbolLimitsResponsePointer {
	return TradeAccountDataSymbolLimitsResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                                              = TradeAccountDataSymbolLimitsResponseSize  (47)
//     Type                                              = TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_RESPONSE  (10121)
//     BaseSize                                          = TradeAccountDataSymbolLimitsResponseSize  (47)
//     RequestID                                         = 0
//     TradeAccount                                      = ""
//     Symbol                                            = ""
//     TradePositionLimit                                = 0
//     OrderQuantityLimit                                = 0
//     UsePercentOfMargin                                = 0
//     OverrideMarginOtherAccounts                       = 0
//     UsePercentOfMarginForDayTrading                   = 0
//     NumberOfDaysBeforeLastTradingDateToDisallowOrders = 0
// }
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataSymbolLimitsResponseDefault)
}

// ToBuilder
func (m TradeAccountDataSymbolLimitsResponsePointer) ToBuilder() TradeAccountDataSymbolLimitsResponsePointerBuilder {
	return TradeAccountDataSymbolLimitsResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataSymbolLimitsResponsePointerBuilder) Finish() TradeAccountDataSymbolLimitsResponsePointer {
	return TradeAccountDataSymbolLimitsResponsePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataSymbolLimitsResponsePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataSymbolLimitsResponsePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Symbol
func (m TradeAccountDataSymbolLimitsResponsePointer) Symbol() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Symbol
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TradePositionLimit
func (m TradeAccountDataSymbolLimitsResponsePointer) TradePositionLimit() float64 {
	return message.Float64VLS(m.Ptr, 26, 18)
}

// TradePositionLimit
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) TradePositionLimit() float64 {
	return message.Float64VLS(m.Ptr, 26, 18)
}

// OrderQuantityLimit
func (m TradeAccountDataSymbolLimitsResponsePointer) OrderQuantityLimit() float64 {
	return message.Float64VLS(m.Ptr, 34, 26)
}

// OrderQuantityLimit
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) OrderQuantityLimit() float64 {
	return message.Float64VLS(m.Ptr, 34, 26)
}

// UsePercentOfMargin
func (m TradeAccountDataSymbolLimitsResponsePointer) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Ptr, 38, 34)
}

// UsePercentOfMargin
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Ptr, 38, 34)
}

// OverrideMarginOtherAccounts
func (m TradeAccountDataSymbolLimitsResponsePointer) OverrideMarginOtherAccounts() uint8 {
	return message.Uint8VLS(m.Ptr, 39, 38)
}

// OverrideMarginOtherAccounts
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) OverrideMarginOtherAccounts() uint8 {
	return message.Uint8VLS(m.Ptr, 39, 38)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataSymbolLimitsResponsePointer) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Ptr, 43, 39)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Ptr, 43, 39)
}

// NumberOfDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataSymbolLimitsResponsePointer) NumberOfDaysBeforeLastTradingDateToDisallowOrders() int32 {
	return message.Int32VLS(m.Ptr, 47, 43)
}

// NumberOfDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) NumberOfDaysBeforeLastTradingDateToDisallowOrders() int32 {
	return message.Int32VLS(m.Ptr, 47, 43)
}

// SetRequestID
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataSymbolLimitsResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetSymbol
func (m *TradeAccountDataSymbolLimitsResponsePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetTradePositionLimit
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) SetTradePositionLimit(value float64) {
	message.SetFloat64VLS(m.Ptr, 26, 18, value)
}

// SetOrderQuantityLimit
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) SetOrderQuantityLimit(value float64) {
	message.SetFloat64VLS(m.Ptr, 34, 26, value)
}

// SetUsePercentOfMargin
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) SetUsePercentOfMargin(value int32) {
	message.SetInt32VLS(m.Ptr, 38, 34, value)
}

// SetOverrideMarginOtherAccounts
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) SetOverrideMarginOtherAccounts(value uint8) {
	message.SetUint8VLS(m.Ptr, 39, 38, value)
}

// SetUsePercentOfMarginForDayTrading
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) SetUsePercentOfMarginForDayTrading(value int32) {
	message.SetInt32VLS(m.Ptr, 43, 39, value)
}

// SetNumberOfDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) SetNumberOfDaysBeforeLastTradingDateToDisallowOrders(value int32) {
	message.SetInt32VLS(m.Ptr, 47, 43, value)
}

// Copy
func (m TradeAccountDataSymbolLimitsResponsePointer) Copy(to TradeAccountDataSymbolLimitsResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradePositionLimit(m.TradePositionLimit())
	to.SetOrderQuantityLimit(m.OrderQuantityLimit())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetOverrideMarginOtherAccounts(m.OverrideMarginOtherAccounts())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetNumberOfDaysBeforeLastTradingDateToDisallowOrders(m.NumberOfDaysBeforeLastTradingDateToDisallowOrders())
}

// Copy
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) Copy(to TradeAccountDataSymbolLimitsResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradePositionLimit(m.TradePositionLimit())
	to.SetOrderQuantityLimit(m.OrderQuantityLimit())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetOverrideMarginOtherAccounts(m.OverrideMarginOtherAccounts())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetNumberOfDaysBeforeLastTradingDateToDisallowOrders(m.NumberOfDaysBeforeLastTradingDateToDisallowOrders())
}

// CopyPointer
func (m TradeAccountDataSymbolLimitsResponsePointer) CopyPointer(to TradeAccountDataSymbolLimitsResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradePositionLimit(m.TradePositionLimit())
	to.SetOrderQuantityLimit(m.OrderQuantityLimit())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetOverrideMarginOtherAccounts(m.OverrideMarginOtherAccounts())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetNumberOfDaysBeforeLastTradingDateToDisallowOrders(m.NumberOfDaysBeforeLastTradingDateToDisallowOrders())
}

// CopyPointer
func (m TradeAccountDataSymbolLimitsResponsePointerBuilder) CopyPointer(to TradeAccountDataSymbolLimitsResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradePositionLimit(m.TradePositionLimit())
	to.SetOrderQuantityLimit(m.OrderQuantityLimit())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetOverrideMarginOtherAccounts(m.OverrideMarginOtherAccounts())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetNumberOfDaysBeforeLastTradingDateToDisallowOrders(m.NumberOfDaysBeforeLastTradingDateToDisallowOrders())
}
