package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeAccountDataSymbolLimitsUpdateSize = 48

// {
//     Size                                              = TradeAccountDataSymbolLimitsUpdateSize  (48)
//     Type                                              = TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_UPDATE  (10122)
//     BaseSize                                          = TradeAccountDataSymbolLimitsUpdateSize  (48)
//     RequestID                                         = 0
//     IsDeleted                                         = false
//     TradeAccount                                      = ""
//     Symbol                                            = ""
//     TradePositionLimit                                = 0
//     OrderQuantityLimit                                = 0
//     UsePercentOfMargin                                = 100
//     OverrideMarginOtherAccounts                       = 0
//     UsePercentOfMarginForDayTrading                   = 0
//     NumberOfDaysBeforeLastTradingDateToDisallowOrders = 0
// }
var _TradeAccountDataSymbolLimitsUpdateDefault = []byte{48, 0, 138, 39, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataSymbolLimitsUpdate struct {
	message.VLS
}

type TradeAccountDataSymbolLimitsUpdateBuilder struct {
	message.VLSBuilder
}

type TradeAccountDataSymbolLimitsUpdatePointer struct {
	message.VLSPointer
}

type TradeAccountDataSymbolLimitsUpdatePointerBuilder struct {
	message.VLSPointerBuilder
}

func NewTradeAccountDataSymbolLimitsUpdateFrom(b []byte) TradeAccountDataSymbolLimitsUpdate {
	return TradeAccountDataSymbolLimitsUpdate{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataSymbolLimitsUpdate(b []byte) TradeAccountDataSymbolLimitsUpdate {
	return TradeAccountDataSymbolLimitsUpdate{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataSymbolLimitsUpdate() TradeAccountDataSymbolLimitsUpdateBuilder {
	a := TradeAccountDataSymbolLimitsUpdateBuilder{message.NewVLSBuilder(48)}
	a.Unsafe().SetBytes(0, _TradeAccountDataSymbolLimitsUpdateDefault)
	return a
}

func AllocTradeAccountDataSymbolLimitsUpdate() TradeAccountDataSymbolLimitsUpdatePointerBuilder {
	a := TradeAccountDataSymbolLimitsUpdatePointerBuilder{message.AllocVLSBuilder(48)}
	a.Ptr.SetBytes(0, _TradeAccountDataSymbolLimitsUpdateDefault)
	return a
}

func AllocTradeAccountDataSymbolLimitsUpdateFrom(b []byte) TradeAccountDataSymbolLimitsUpdatePointer {
	return TradeAccountDataSymbolLimitsUpdatePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                                              = TradeAccountDataSymbolLimitsUpdateSize  (48)
//     Type                                              = TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_UPDATE  (10122)
//     BaseSize                                          = TradeAccountDataSymbolLimitsUpdateSize  (48)
//     RequestID                                         = 0
//     IsDeleted                                         = false
//     TradeAccount                                      = ""
//     Symbol                                            = ""
//     TradePositionLimit                                = 0
//     OrderQuantityLimit                                = 0
//     UsePercentOfMargin                                = 100
//     OverrideMarginOtherAccounts                       = 0
//     UsePercentOfMarginForDayTrading                   = 0
//     NumberOfDaysBeforeLastTradingDateToDisallowOrders = 0
// }
func (m TradeAccountDataSymbolLimitsUpdateBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataSymbolLimitsUpdateDefault)
}

// Clear
// {
//     Size                                              = TradeAccountDataSymbolLimitsUpdateSize  (48)
//     Type                                              = TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_UPDATE  (10122)
//     BaseSize                                          = TradeAccountDataSymbolLimitsUpdateSize  (48)
//     RequestID                                         = 0
//     IsDeleted                                         = false
//     TradeAccount                                      = ""
//     Symbol                                            = ""
//     TradePositionLimit                                = 0
//     OrderQuantityLimit                                = 0
//     UsePercentOfMargin                                = 100
//     OverrideMarginOtherAccounts                       = 0
//     UsePercentOfMarginForDayTrading                   = 0
//     NumberOfDaysBeforeLastTradingDateToDisallowOrders = 0
// }
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataSymbolLimitsUpdateDefault)
}

// ToBuilder
func (m TradeAccountDataSymbolLimitsUpdate) ToBuilder() TradeAccountDataSymbolLimitsUpdateBuilder {
	return TradeAccountDataSymbolLimitsUpdateBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradeAccountDataSymbolLimitsUpdatePointer) ToBuilder() TradeAccountDataSymbolLimitsUpdatePointerBuilder {
	return TradeAccountDataSymbolLimitsUpdatePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m TradeAccountDataSymbolLimitsUpdateBuilder) Finish() TradeAccountDataSymbolLimitsUpdate {
	return TradeAccountDataSymbolLimitsUpdate{m.VLSBuilder.Finish()}
}

// Finish
func (m *TradeAccountDataSymbolLimitsUpdatePointerBuilder) Finish() TradeAccountDataSymbolLimitsUpdatePointer {
	return TradeAccountDataSymbolLimitsUpdatePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataSymbolLimitsUpdate) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataSymbolLimitsUpdateBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataSymbolLimitsUpdatePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// IsDeleted
func (m TradeAccountDataSymbolLimitsUpdate) IsDeleted() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsDeleted
func (m TradeAccountDataSymbolLimitsUpdateBuilder) IsDeleted() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsDeleted
func (m TradeAccountDataSymbolLimitsUpdatePointer) IsDeleted() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// IsDeleted
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) IsDeleted() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// TradeAccount
func (m TradeAccountDataSymbolLimitsUpdate) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// TradeAccount
func (m TradeAccountDataSymbolLimitsUpdateBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// TradeAccount
func (m TradeAccountDataSymbolLimitsUpdatePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// TradeAccount
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// Symbol
func (m TradeAccountDataSymbolLimitsUpdate) Symbol() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
}

// Symbol
func (m TradeAccountDataSymbolLimitsUpdateBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
}

// Symbol
func (m TradeAccountDataSymbolLimitsUpdatePointer) Symbol() string {
	return message.StringVLS(m.Ptr, 19, 15)
}

// Symbol
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 19, 15)
}

// TradePositionLimit
func (m TradeAccountDataSymbolLimitsUpdate) TradePositionLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 27, 19)
}

// TradePositionLimit
func (m TradeAccountDataSymbolLimitsUpdateBuilder) TradePositionLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 27, 19)
}

// TradePositionLimit
func (m TradeAccountDataSymbolLimitsUpdatePointer) TradePositionLimit() float64 {
	return message.Float64VLS(m.Ptr, 27, 19)
}

// TradePositionLimit
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) TradePositionLimit() float64 {
	return message.Float64VLS(m.Ptr, 27, 19)
}

// OrderQuantityLimit
func (m TradeAccountDataSymbolLimitsUpdate) OrderQuantityLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 35, 27)
}

// OrderQuantityLimit
func (m TradeAccountDataSymbolLimitsUpdateBuilder) OrderQuantityLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 35, 27)
}

// OrderQuantityLimit
func (m TradeAccountDataSymbolLimitsUpdatePointer) OrderQuantityLimit() float64 {
	return message.Float64VLS(m.Ptr, 35, 27)
}

// OrderQuantityLimit
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) OrderQuantityLimit() float64 {
	return message.Float64VLS(m.Ptr, 35, 27)
}

// UsePercentOfMargin
func (m TradeAccountDataSymbolLimitsUpdate) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Unsafe(), 39, 35)
}

// UsePercentOfMargin
func (m TradeAccountDataSymbolLimitsUpdateBuilder) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Unsafe(), 39, 35)
}

// UsePercentOfMargin
func (m TradeAccountDataSymbolLimitsUpdatePointer) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Ptr, 39, 35)
}

// UsePercentOfMargin
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Ptr, 39, 35)
}

// OverrideMarginOtherAccounts
func (m TradeAccountDataSymbolLimitsUpdate) OverrideMarginOtherAccounts() uint8 {
	return message.Uint8VLS(m.Unsafe(), 40, 39)
}

// OverrideMarginOtherAccounts
func (m TradeAccountDataSymbolLimitsUpdateBuilder) OverrideMarginOtherAccounts() uint8 {
	return message.Uint8VLS(m.Unsafe(), 40, 39)
}

// OverrideMarginOtherAccounts
func (m TradeAccountDataSymbolLimitsUpdatePointer) OverrideMarginOtherAccounts() uint8 {
	return message.Uint8VLS(m.Ptr, 40, 39)
}

// OverrideMarginOtherAccounts
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) OverrideMarginOtherAccounts() uint8 {
	return message.Uint8VLS(m.Ptr, 40, 39)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataSymbolLimitsUpdate) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Unsafe(), 44, 40)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataSymbolLimitsUpdateBuilder) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Unsafe(), 44, 40)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataSymbolLimitsUpdatePointer) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Ptr, 44, 40)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Ptr, 44, 40)
}

// NumberOfDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataSymbolLimitsUpdate) NumberOfDaysBeforeLastTradingDateToDisallowOrders() int32 {
	return message.Int32VLS(m.Unsafe(), 48, 44)
}

// NumberOfDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataSymbolLimitsUpdateBuilder) NumberOfDaysBeforeLastTradingDateToDisallowOrders() int32 {
	return message.Int32VLS(m.Unsafe(), 48, 44)
}

// NumberOfDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataSymbolLimitsUpdatePointer) NumberOfDaysBeforeLastTradingDateToDisallowOrders() int32 {
	return message.Int32VLS(m.Ptr, 48, 44)
}

// NumberOfDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) NumberOfDaysBeforeLastTradingDateToDisallowOrders() int32 {
	return message.Int32VLS(m.Ptr, 48, 44)
}

// SetRequestID
func (m TradeAccountDataSymbolLimitsUpdateBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetIsDeleted
func (m TradeAccountDataSymbolLimitsUpdateBuilder) SetIsDeleted(value bool) {
	message.SetBoolVLS(m.Unsafe(), 11, 10, value)
}

// SetIsDeleted
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) SetIsDeleted(value bool) {
	message.SetBoolVLS(m.Ptr, 11, 10, value)
}

// SetTradeAccount
func (m *TradeAccountDataSymbolLimitsUpdateBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 15, 11, value)
}

// SetTradeAccount
func (m *TradeAccountDataSymbolLimitsUpdatePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetSymbol
func (m *TradeAccountDataSymbolLimitsUpdateBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 19, 15, value)
}

// SetSymbol
func (m *TradeAccountDataSymbolLimitsUpdatePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 19, 15, value)
}

// SetTradePositionLimit
func (m TradeAccountDataSymbolLimitsUpdateBuilder) SetTradePositionLimit(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 27, 19, value)
}

// SetTradePositionLimit
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) SetTradePositionLimit(value float64) {
	message.SetFloat64VLS(m.Ptr, 27, 19, value)
}

// SetOrderQuantityLimit
func (m TradeAccountDataSymbolLimitsUpdateBuilder) SetOrderQuantityLimit(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 35, 27, value)
}

// SetOrderQuantityLimit
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) SetOrderQuantityLimit(value float64) {
	message.SetFloat64VLS(m.Ptr, 35, 27, value)
}

// SetUsePercentOfMargin
func (m TradeAccountDataSymbolLimitsUpdateBuilder) SetUsePercentOfMargin(value int32) {
	message.SetInt32VLS(m.Unsafe(), 39, 35, value)
}

// SetUsePercentOfMargin
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) SetUsePercentOfMargin(value int32) {
	message.SetInt32VLS(m.Ptr, 39, 35, value)
}

// SetOverrideMarginOtherAccounts
func (m TradeAccountDataSymbolLimitsUpdateBuilder) SetOverrideMarginOtherAccounts(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 40, 39, value)
}

// SetOverrideMarginOtherAccounts
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) SetOverrideMarginOtherAccounts(value uint8) {
	message.SetUint8VLS(m.Ptr, 40, 39, value)
}

// SetUsePercentOfMarginForDayTrading
func (m TradeAccountDataSymbolLimitsUpdateBuilder) SetUsePercentOfMarginForDayTrading(value int32) {
	message.SetInt32VLS(m.Unsafe(), 44, 40, value)
}

// SetUsePercentOfMarginForDayTrading
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) SetUsePercentOfMarginForDayTrading(value int32) {
	message.SetInt32VLS(m.Ptr, 44, 40, value)
}

// SetNumberOfDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataSymbolLimitsUpdateBuilder) SetNumberOfDaysBeforeLastTradingDateToDisallowOrders(value int32) {
	message.SetInt32VLS(m.Unsafe(), 48, 44, value)
}

// SetNumberOfDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) SetNumberOfDaysBeforeLastTradingDateToDisallowOrders(value int32) {
	message.SetInt32VLS(m.Ptr, 48, 44, value)
}

// Copy
func (m TradeAccountDataSymbolLimitsUpdate) Copy(to TradeAccountDataSymbolLimitsUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
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
func (m TradeAccountDataSymbolLimitsUpdateBuilder) Copy(to TradeAccountDataSymbolLimitsUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
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
func (m TradeAccountDataSymbolLimitsUpdate) CopyPointer(to TradeAccountDataSymbolLimitsUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
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
func (m TradeAccountDataSymbolLimitsUpdateBuilder) CopyPointer(to TradeAccountDataSymbolLimitsUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
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
func (m TradeAccountDataSymbolLimitsUpdatePointer) Copy(to TradeAccountDataSymbolLimitsUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
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
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) Copy(to TradeAccountDataSymbolLimitsUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
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
func (m TradeAccountDataSymbolLimitsUpdatePointer) CopyPointer(to TradeAccountDataSymbolLimitsUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
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
func (m TradeAccountDataSymbolLimitsUpdatePointerBuilder) CopyPointer(to TradeAccountDataSymbolLimitsUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradePositionLimit(m.TradePositionLimit())
	to.SetOrderQuantityLimit(m.OrderQuantityLimit())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetOverrideMarginOtherAccounts(m.OverrideMarginOtherAccounts())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetNumberOfDaysBeforeLastTradingDateToDisallowOrders(m.NumberOfDaysBeforeLastTradingDateToDisallowOrders())
}