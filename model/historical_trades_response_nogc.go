package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalTradesResponsePointer struct {
	message.VLSPointer
}

type HistoricalTradesResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocHistoricalTradesResponse() HistoricalTradesResponsePointerBuilder {
	a := HistoricalTradesResponsePointerBuilder{message.AllocVLSBuilder(111)}
	a.Ptr.SetBytes(0, _HistoricalTradesResponseDefault)
	return a
}

func AllocHistoricalTradesResponseFrom(b []byte) HistoricalTradesResponsePointer {
	return HistoricalTradesResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                      = HistoricalTradesResponseSize  (111)
//     Type                      = HISTORICAL_TRADES_RESPONSE  (10102)
//     BaseSize                  = HistoricalTradesResponseSize  (111)
//     RequestID                 = 0
//     IsFinalMessage            = 0
//     Symbol                    = ""
//     TradeAccount              = ""
//     EntryDateTime             = 0
//     ExitDateTime              = 0
//     EntryPrice                = 0
//     ExitPrice                 = 0
//     TradeType                 = BUY_SELL_UNSET  (0)
//     EntryQuantity             = 0
//     ExitQuantity              = 0
//     MaxOpenQuantity           = 0
//     ClosedProfitLoss          = 0
//     MaximumOpenPositionLoss   = 0
//     MaximumOpenPositionProfit = 0
//     Commission                = 0
//     OpenFillExecutionID       = ""
//     CloseFillExecutionID      = ""
//     SubAccountIdentifier      = 0
// }
func (m HistoricalTradesResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalTradesResponseDefault)
}

// ToBuilder
func (m HistoricalTradesResponsePointer) ToBuilder() HistoricalTradesResponsePointerBuilder {
	return HistoricalTradesResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalTradesResponsePointerBuilder) Finish() HistoricalTradesResponsePointer {
	return HistoricalTradesResponsePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m HistoricalTradesResponsePointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m HistoricalTradesResponsePointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// IsFinalMessage
func (m HistoricalTradesResponsePointer) IsFinalMessage() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// IsFinalMessage
func (m HistoricalTradesResponsePointerBuilder) IsFinalMessage() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// Symbol
func (m HistoricalTradesResponsePointer) Symbol() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// Symbol
func (m HistoricalTradesResponsePointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// TradeAccount
func (m HistoricalTradesResponsePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 19, 15)
}

// TradeAccount
func (m HistoricalTradesResponsePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 19, 15)
}

// EntryDateTime
func (m HistoricalTradesResponsePointer) EntryDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Ptr, 27, 19))
}

// EntryDateTime
func (m HistoricalTradesResponsePointerBuilder) EntryDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Ptr, 27, 19))
}

// ExitDateTime
func (m HistoricalTradesResponsePointer) ExitDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Ptr, 35, 27))
}

// ExitDateTime
func (m HistoricalTradesResponsePointerBuilder) ExitDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Ptr, 35, 27))
}

// EntryPrice
func (m HistoricalTradesResponsePointer) EntryPrice() float64 {
	return message.Float64VLS(m.Ptr, 43, 35)
}

// EntryPrice
func (m HistoricalTradesResponsePointerBuilder) EntryPrice() float64 {
	return message.Float64VLS(m.Ptr, 43, 35)
}

// ExitPrice
func (m HistoricalTradesResponsePointer) ExitPrice() float64 {
	return message.Float64VLS(m.Ptr, 51, 43)
}

// ExitPrice
func (m HistoricalTradesResponsePointerBuilder) ExitPrice() float64 {
	return message.Float64VLS(m.Ptr, 51, 43)
}

// TradeType
func (m HistoricalTradesResponsePointer) TradeType() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 55, 51))
}

// TradeType
func (m HistoricalTradesResponsePointerBuilder) TradeType() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 55, 51))
}

// EntryQuantity
func (m HistoricalTradesResponsePointer) EntryQuantity() uint32 {
	return message.Uint32VLS(m.Ptr, 59, 55)
}

// EntryQuantity
func (m HistoricalTradesResponsePointerBuilder) EntryQuantity() uint32 {
	return message.Uint32VLS(m.Ptr, 59, 55)
}

// ExitQuantity
func (m HistoricalTradesResponsePointer) ExitQuantity() uint32 {
	return message.Uint32VLS(m.Ptr, 63, 59)
}

// ExitQuantity
func (m HistoricalTradesResponsePointerBuilder) ExitQuantity() uint32 {
	return message.Uint32VLS(m.Ptr, 63, 59)
}

// MaxOpenQuantity
func (m HistoricalTradesResponsePointer) MaxOpenQuantity() uint32 {
	return message.Uint32VLS(m.Ptr, 67, 63)
}

// MaxOpenQuantity
func (m HistoricalTradesResponsePointerBuilder) MaxOpenQuantity() uint32 {
	return message.Uint32VLS(m.Ptr, 67, 63)
}

// ClosedProfitLoss
func (m HistoricalTradesResponsePointer) ClosedProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 75, 67)
}

// ClosedProfitLoss
func (m HistoricalTradesResponsePointerBuilder) ClosedProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 75, 67)
}

// MaximumOpenPositionLoss
func (m HistoricalTradesResponsePointer) MaximumOpenPositionLoss() float64 {
	return message.Float64VLS(m.Ptr, 83, 75)
}

// MaximumOpenPositionLoss
func (m HistoricalTradesResponsePointerBuilder) MaximumOpenPositionLoss() float64 {
	return message.Float64VLS(m.Ptr, 83, 75)
}

// MaximumOpenPositionProfit
func (m HistoricalTradesResponsePointer) MaximumOpenPositionProfit() float64 {
	return message.Float64VLS(m.Ptr, 91, 83)
}

// MaximumOpenPositionProfit
func (m HistoricalTradesResponsePointerBuilder) MaximumOpenPositionProfit() float64 {
	return message.Float64VLS(m.Ptr, 91, 83)
}

// Commission
func (m HistoricalTradesResponsePointer) Commission() float64 {
	return message.Float64VLS(m.Ptr, 99, 91)
}

// Commission
func (m HistoricalTradesResponsePointerBuilder) Commission() float64 {
	return message.Float64VLS(m.Ptr, 99, 91)
}

// OpenFillExecutionID
func (m HistoricalTradesResponsePointer) OpenFillExecutionID() string {
	return message.StringVLS(m.Ptr, 103, 99)
}

// OpenFillExecutionID
func (m HistoricalTradesResponsePointerBuilder) OpenFillExecutionID() string {
	return message.StringVLS(m.Ptr, 103, 99)
}

// CloseFillExecutionID
func (m HistoricalTradesResponsePointer) CloseFillExecutionID() string {
	return message.StringVLS(m.Ptr, 107, 103)
}

// CloseFillExecutionID
func (m HistoricalTradesResponsePointerBuilder) CloseFillExecutionID() string {
	return message.StringVLS(m.Ptr, 107, 103)
}

// SubAccountIdentifier
func (m HistoricalTradesResponsePointer) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 111, 107)
}

// SubAccountIdentifier
func (m HistoricalTradesResponsePointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 111, 107)
}

// SetRequestID
func (m HistoricalTradesResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 10, 6, value)
}

// SetIsFinalMessage
func (m HistoricalTradesResponsePointerBuilder) SetIsFinalMessage(value uint8) {
	message.SetUint8VLS(m.Ptr, 11, 10, value)
}

// SetSymbol
func (m *HistoricalTradesResponsePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetTradeAccount
func (m *HistoricalTradesResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 19, 15, value)
}

// SetEntryDateTime
func (m HistoricalTradesResponsePointerBuilder) SetEntryDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Ptr, 27, 19, float64(value))
}

// SetExitDateTime
func (m HistoricalTradesResponsePointerBuilder) SetExitDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Ptr, 35, 27, float64(value))
}

// SetEntryPrice
func (m HistoricalTradesResponsePointerBuilder) SetEntryPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 43, 35, value)
}

// SetExitPrice
func (m HistoricalTradesResponsePointerBuilder) SetExitPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 51, 43, value)
}

// SetTradeType
func (m HistoricalTradesResponsePointerBuilder) SetTradeType(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 55, 51, int32(value))
}

// SetEntryQuantity
func (m HistoricalTradesResponsePointerBuilder) SetEntryQuantity(value uint32) {
	message.SetUint32VLS(m.Ptr, 59, 55, value)
}

// SetExitQuantity
func (m HistoricalTradesResponsePointerBuilder) SetExitQuantity(value uint32) {
	message.SetUint32VLS(m.Ptr, 63, 59, value)
}

// SetMaxOpenQuantity
func (m HistoricalTradesResponsePointerBuilder) SetMaxOpenQuantity(value uint32) {
	message.SetUint32VLS(m.Ptr, 67, 63, value)
}

// SetClosedProfitLoss
func (m HistoricalTradesResponsePointerBuilder) SetClosedProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 75, 67, value)
}

// SetMaximumOpenPositionLoss
func (m HistoricalTradesResponsePointerBuilder) SetMaximumOpenPositionLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 83, 75, value)
}

// SetMaximumOpenPositionProfit
func (m HistoricalTradesResponsePointerBuilder) SetMaximumOpenPositionProfit(value float64) {
	message.SetFloat64VLS(m.Ptr, 91, 83, value)
}

// SetCommission
func (m HistoricalTradesResponsePointerBuilder) SetCommission(value float64) {
	message.SetFloat64VLS(m.Ptr, 99, 91, value)
}

// SetOpenFillExecutionID
func (m *HistoricalTradesResponsePointerBuilder) SetOpenFillExecutionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 103, 99, value)
}

// SetCloseFillExecutionID
func (m *HistoricalTradesResponsePointerBuilder) SetCloseFillExecutionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 107, 103, value)
}

// SetSubAccountIdentifier
func (m HistoricalTradesResponsePointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Ptr, 111, 107, value)
}

// Copy
func (m HistoricalTradesResponsePointer) Copy(to HistoricalTradesResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetExitDateTime(m.ExitDateTime())
	to.SetEntryPrice(m.EntryPrice())
	to.SetExitPrice(m.ExitPrice())
	to.SetTradeType(m.TradeType())
	to.SetEntryQuantity(m.EntryQuantity())
	to.SetExitQuantity(m.ExitQuantity())
	to.SetMaxOpenQuantity(m.MaxOpenQuantity())
	to.SetClosedProfitLoss(m.ClosedProfitLoss())
	to.SetMaximumOpenPositionLoss(m.MaximumOpenPositionLoss())
	to.SetMaximumOpenPositionProfit(m.MaximumOpenPositionProfit())
	to.SetCommission(m.Commission())
	to.SetOpenFillExecutionID(m.OpenFillExecutionID())
	to.SetCloseFillExecutionID(m.CloseFillExecutionID())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// Copy
func (m HistoricalTradesResponsePointerBuilder) Copy(to HistoricalTradesResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetExitDateTime(m.ExitDateTime())
	to.SetEntryPrice(m.EntryPrice())
	to.SetExitPrice(m.ExitPrice())
	to.SetTradeType(m.TradeType())
	to.SetEntryQuantity(m.EntryQuantity())
	to.SetExitQuantity(m.ExitQuantity())
	to.SetMaxOpenQuantity(m.MaxOpenQuantity())
	to.SetClosedProfitLoss(m.ClosedProfitLoss())
	to.SetMaximumOpenPositionLoss(m.MaximumOpenPositionLoss())
	to.SetMaximumOpenPositionProfit(m.MaximumOpenPositionProfit())
	to.SetCommission(m.Commission())
	to.SetOpenFillExecutionID(m.OpenFillExecutionID())
	to.SetCloseFillExecutionID(m.CloseFillExecutionID())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyPointer
func (m HistoricalTradesResponsePointer) CopyPointer(to HistoricalTradesResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetExitDateTime(m.ExitDateTime())
	to.SetEntryPrice(m.EntryPrice())
	to.SetExitPrice(m.ExitPrice())
	to.SetTradeType(m.TradeType())
	to.SetEntryQuantity(m.EntryQuantity())
	to.SetExitQuantity(m.ExitQuantity())
	to.SetMaxOpenQuantity(m.MaxOpenQuantity())
	to.SetClosedProfitLoss(m.ClosedProfitLoss())
	to.SetMaximumOpenPositionLoss(m.MaximumOpenPositionLoss())
	to.SetMaximumOpenPositionProfit(m.MaximumOpenPositionProfit())
	to.SetCommission(m.Commission())
	to.SetOpenFillExecutionID(m.OpenFillExecutionID())
	to.SetCloseFillExecutionID(m.CloseFillExecutionID())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyPointer
func (m HistoricalTradesResponsePointerBuilder) CopyPointer(to HistoricalTradesResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetExitDateTime(m.ExitDateTime())
	to.SetEntryPrice(m.EntryPrice())
	to.SetExitPrice(m.ExitPrice())
	to.SetTradeType(m.TradeType())
	to.SetEntryQuantity(m.EntryQuantity())
	to.SetExitQuantity(m.ExitQuantity())
	to.SetMaxOpenQuantity(m.MaxOpenQuantity())
	to.SetClosedProfitLoss(m.ClosedProfitLoss())
	to.SetMaximumOpenPositionLoss(m.MaximumOpenPositionLoss())
	to.SetMaximumOpenPositionProfit(m.MaximumOpenPositionProfit())
	to.SetCommission(m.Commission())
	to.SetOpenFillExecutionID(m.OpenFillExecutionID())
	to.SetCloseFillExecutionID(m.CloseFillExecutionID())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyTo
func (m HistoricalTradesResponsePointer) CopyTo(to HistoricalTradesResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetExitDateTime(m.ExitDateTime())
	to.SetEntryPrice(m.EntryPrice())
	to.SetExitPrice(m.ExitPrice())
	to.SetTradeType(m.TradeType())
	to.SetEntryQuantity(m.EntryQuantity())
	to.SetExitQuantity(m.ExitQuantity())
	to.SetMaxOpenQuantity(m.MaxOpenQuantity())
	to.SetClosedProfitLoss(m.ClosedProfitLoss())
	to.SetMaximumOpenPositionLoss(m.MaximumOpenPositionLoss())
	to.SetMaximumOpenPositionProfit(m.MaximumOpenPositionProfit())
	to.SetCommission(m.Commission())
	to.SetOpenFillExecutionID(m.OpenFillExecutionID())
	to.SetCloseFillExecutionID(m.CloseFillExecutionID())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyTo
func (m HistoricalTradesResponsePointerBuilder) CopyTo(to HistoricalTradesResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetExitDateTime(m.ExitDateTime())
	to.SetEntryPrice(m.EntryPrice())
	to.SetExitPrice(m.ExitPrice())
	to.SetTradeType(m.TradeType())
	to.SetEntryQuantity(m.EntryQuantity())
	to.SetExitQuantity(m.ExitQuantity())
	to.SetMaxOpenQuantity(m.MaxOpenQuantity())
	to.SetClosedProfitLoss(m.ClosedProfitLoss())
	to.SetMaximumOpenPositionLoss(m.MaximumOpenPositionLoss())
	to.SetMaximumOpenPositionProfit(m.MaximumOpenPositionProfit())
	to.SetCommission(m.Commission())
	to.SetOpenFillExecutionID(m.OpenFillExecutionID())
	to.SetCloseFillExecutionID(m.CloseFillExecutionID())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyToPointer
func (m HistoricalTradesResponsePointer) CopyToPointer(to HistoricalTradesResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetExitDateTime(m.ExitDateTime())
	to.SetEntryPrice(m.EntryPrice())
	to.SetExitPrice(m.ExitPrice())
	to.SetTradeType(m.TradeType())
	to.SetEntryQuantity(m.EntryQuantity())
	to.SetExitQuantity(m.ExitQuantity())
	to.SetMaxOpenQuantity(m.MaxOpenQuantity())
	to.SetClosedProfitLoss(m.ClosedProfitLoss())
	to.SetMaximumOpenPositionLoss(m.MaximumOpenPositionLoss())
	to.SetMaximumOpenPositionProfit(m.MaximumOpenPositionProfit())
	to.SetCommission(m.Commission())
	to.SetOpenFillExecutionID(m.OpenFillExecutionID())
	to.SetCloseFillExecutionID(m.CloseFillExecutionID())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyToPointer
func (m HistoricalTradesResponsePointerBuilder) CopyToPointer(to HistoricalTradesResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetExitDateTime(m.ExitDateTime())
	to.SetEntryPrice(m.EntryPrice())
	to.SetExitPrice(m.ExitPrice())
	to.SetTradeType(m.TradeType())
	to.SetEntryQuantity(m.EntryQuantity())
	to.SetExitQuantity(m.ExitQuantity())
	to.SetMaxOpenQuantity(m.MaxOpenQuantity())
	to.SetClosedProfitLoss(m.ClosedProfitLoss())
	to.SetMaximumOpenPositionLoss(m.MaximumOpenPositionLoss())
	to.SetMaximumOpenPositionProfit(m.MaximumOpenPositionProfit())
	to.SetCommission(m.Commission())
	to.SetOpenFillExecutionID(m.OpenFillExecutionID())
	to.SetCloseFillExecutionID(m.CloseFillExecutionID())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}
