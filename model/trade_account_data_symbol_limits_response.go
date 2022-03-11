package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _TradeAccountDataSymbolLimitsResponseDefault = []byte{47, 0, 137, 39, 47, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataSymbolLimitsResponseSize = 47

type TradeAccountDataSymbolLimitsResponse struct {
	message.VLS
}

type TradeAccountDataSymbolLimitsResponseBuilder struct {
	message.VLSBuilder
}

func NewTradeAccountDataSymbolLimitsResponseFrom(b []byte) TradeAccountDataSymbolLimitsResponse {
	return TradeAccountDataSymbolLimitsResponse{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataSymbolLimitsResponse(b []byte) TradeAccountDataSymbolLimitsResponse {
	return TradeAccountDataSymbolLimitsResponse{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataSymbolLimitsResponse() TradeAccountDataSymbolLimitsResponseBuilder {
	a := TradeAccountDataSymbolLimitsResponseBuilder{message.NewVLSBuilder(47)}
	a.Unsafe().SetBytes(0, _TradeAccountDataSymbolLimitsResponseDefault)
	return a
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
func (m TradeAccountDataSymbolLimitsResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataSymbolLimitsResponseDefault)
}

// ToBuilder
func (m TradeAccountDataSymbolLimitsResponse) ToBuilder() TradeAccountDataSymbolLimitsResponseBuilder {
	return TradeAccountDataSymbolLimitsResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataSymbolLimitsResponseBuilder) Finish() TradeAccountDataSymbolLimitsResponse {
	return TradeAccountDataSymbolLimitsResponse{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataSymbolLimitsResponse) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataSymbolLimitsResponseBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m TradeAccountDataSymbolLimitsResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataSymbolLimitsResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Symbol
func (m TradeAccountDataSymbolLimitsResponse) Symbol() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Symbol
func (m TradeAccountDataSymbolLimitsResponseBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradePositionLimit
func (m TradeAccountDataSymbolLimitsResponse) TradePositionLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 26, 18)
}

// TradePositionLimit
func (m TradeAccountDataSymbolLimitsResponseBuilder) TradePositionLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 26, 18)
}

// OrderQuantityLimit
func (m TradeAccountDataSymbolLimitsResponse) OrderQuantityLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 34, 26)
}

// OrderQuantityLimit
func (m TradeAccountDataSymbolLimitsResponseBuilder) OrderQuantityLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 34, 26)
}

// UsePercentOfMargin
func (m TradeAccountDataSymbolLimitsResponse) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Unsafe(), 38, 34)
}

// UsePercentOfMargin
func (m TradeAccountDataSymbolLimitsResponseBuilder) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Unsafe(), 38, 34)
}

// OverrideMarginOtherAccounts
func (m TradeAccountDataSymbolLimitsResponse) OverrideMarginOtherAccounts() uint8 {
	return message.Uint8VLS(m.Unsafe(), 39, 38)
}

// OverrideMarginOtherAccounts
func (m TradeAccountDataSymbolLimitsResponseBuilder) OverrideMarginOtherAccounts() uint8 {
	return message.Uint8VLS(m.Unsafe(), 39, 38)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataSymbolLimitsResponse) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Unsafe(), 43, 39)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataSymbolLimitsResponseBuilder) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Unsafe(), 43, 39)
}

// NumberOfDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataSymbolLimitsResponse) NumberOfDaysBeforeLastTradingDateToDisallowOrders() int32 {
	return message.Int32VLS(m.Unsafe(), 47, 43)
}

// NumberOfDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataSymbolLimitsResponseBuilder) NumberOfDaysBeforeLastTradingDateToDisallowOrders() int32 {
	return message.Int32VLS(m.Unsafe(), 47, 43)
}

// SetRequestID
func (m TradeAccountDataSymbolLimitsResponseBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataSymbolLimitsResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetSymbol
func (m *TradeAccountDataSymbolLimitsResponseBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetTradePositionLimit
func (m TradeAccountDataSymbolLimitsResponseBuilder) SetTradePositionLimit(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 26, 18, value)
}

// SetOrderQuantityLimit
func (m TradeAccountDataSymbolLimitsResponseBuilder) SetOrderQuantityLimit(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 34, 26, value)
}

// SetUsePercentOfMargin
func (m TradeAccountDataSymbolLimitsResponseBuilder) SetUsePercentOfMargin(value int32) {
	message.SetInt32VLS(m.Unsafe(), 38, 34, value)
}

// SetOverrideMarginOtherAccounts
func (m TradeAccountDataSymbolLimitsResponseBuilder) SetOverrideMarginOtherAccounts(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 39, 38, value)
}

// SetUsePercentOfMarginForDayTrading
func (m TradeAccountDataSymbolLimitsResponseBuilder) SetUsePercentOfMarginForDayTrading(value int32) {
	message.SetInt32VLS(m.Unsafe(), 43, 39, value)
}

// SetNumberOfDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataSymbolLimitsResponseBuilder) SetNumberOfDaysBeforeLastTradingDateToDisallowOrders(value int32) {
	message.SetInt32VLS(m.Unsafe(), 47, 43, value)
}

// Copy
func (m TradeAccountDataSymbolLimitsResponse) Copy(to TradeAccountDataSymbolLimitsResponseBuilder) {
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
func (m TradeAccountDataSymbolLimitsResponseBuilder) Copy(to TradeAccountDataSymbolLimitsResponseBuilder) {
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
func (m TradeAccountDataSymbolLimitsResponse) CopyPointer(to TradeAccountDataSymbolLimitsResponsePointerBuilder) {
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
func (m TradeAccountDataSymbolLimitsResponseBuilder) CopyPointer(to TradeAccountDataSymbolLimitsResponsePointerBuilder) {
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
