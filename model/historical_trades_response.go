package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _HistoricalTradesResponseDefault = []byte{111, 0, 118, 39, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalTradesResponseSize = 111

type HistoricalTradesResponse struct {
	message.VLS
}

type HistoricalTradesResponseBuilder struct {
	message.VLSBuilder
}

func NewHistoricalTradesResponseFrom(b []byte) HistoricalTradesResponse {
	return HistoricalTradesResponse{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalTradesResponse(b []byte) HistoricalTradesResponse {
	return HistoricalTradesResponse{VLS: message.WrapVLS(b)}
}

func NewHistoricalTradesResponse() HistoricalTradesResponseBuilder {
	a := HistoricalTradesResponseBuilder{message.NewVLSBuilder(111)}
	a.Unsafe().SetBytes(0, _HistoricalTradesResponseDefault)
	return a
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
func (m HistoricalTradesResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalTradesResponseDefault)
}

// ToBuilder
func (m HistoricalTradesResponse) ToBuilder() HistoricalTradesResponseBuilder {
	return HistoricalTradesResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m HistoricalTradesResponseBuilder) Finish() HistoricalTradesResponse {
	return HistoricalTradesResponse{m.VLSBuilder.Finish()}
}

// RequestID
func (m HistoricalTradesResponse) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m HistoricalTradesResponseBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// IsFinalMessage
func (m HistoricalTradesResponse) IsFinalMessage() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// IsFinalMessage
func (m HistoricalTradesResponseBuilder) IsFinalMessage() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// Symbol
func (m HistoricalTradesResponse) Symbol() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// Symbol
func (m HistoricalTradesResponseBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// TradeAccount
func (m HistoricalTradesResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
}

// TradeAccount
func (m HistoricalTradesResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
}

// EntryDateTime
func (m HistoricalTradesResponse) EntryDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 27, 19))
}

// EntryDateTime
func (m HistoricalTradesResponseBuilder) EntryDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 27, 19))
}

// ExitDateTime
func (m HistoricalTradesResponse) ExitDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 35, 27))
}

// ExitDateTime
func (m HistoricalTradesResponseBuilder) ExitDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 35, 27))
}

// EntryPrice
func (m HistoricalTradesResponse) EntryPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 43, 35)
}

// EntryPrice
func (m HistoricalTradesResponseBuilder) EntryPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 43, 35)
}

// ExitPrice
func (m HistoricalTradesResponse) ExitPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 51, 43)
}

// ExitPrice
func (m HistoricalTradesResponseBuilder) ExitPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 51, 43)
}

// TradeType
func (m HistoricalTradesResponse) TradeType() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 55, 51))
}

// TradeType
func (m HistoricalTradesResponseBuilder) TradeType() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 55, 51))
}

// EntryQuantity
func (m HistoricalTradesResponse) EntryQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 59, 55)
}

// EntryQuantity
func (m HistoricalTradesResponseBuilder) EntryQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 59, 55)
}

// ExitQuantity
func (m HistoricalTradesResponse) ExitQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 63, 59)
}

// ExitQuantity
func (m HistoricalTradesResponseBuilder) ExitQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 63, 59)
}

// MaxOpenQuantity
func (m HistoricalTradesResponse) MaxOpenQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 67, 63)
}

// MaxOpenQuantity
func (m HistoricalTradesResponseBuilder) MaxOpenQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 67, 63)
}

// ClosedProfitLoss
func (m HistoricalTradesResponse) ClosedProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 75, 67)
}

// ClosedProfitLoss
func (m HistoricalTradesResponseBuilder) ClosedProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 75, 67)
}

// MaximumOpenPositionLoss
func (m HistoricalTradesResponse) MaximumOpenPositionLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 83, 75)
}

// MaximumOpenPositionLoss
func (m HistoricalTradesResponseBuilder) MaximumOpenPositionLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 83, 75)
}

// MaximumOpenPositionProfit
func (m HistoricalTradesResponse) MaximumOpenPositionProfit() float64 {
	return message.Float64VLS(m.Unsafe(), 91, 83)
}

// MaximumOpenPositionProfit
func (m HistoricalTradesResponseBuilder) MaximumOpenPositionProfit() float64 {
	return message.Float64VLS(m.Unsafe(), 91, 83)
}

// Commission
func (m HistoricalTradesResponse) Commission() float64 {
	return message.Float64VLS(m.Unsafe(), 99, 91)
}

// Commission
func (m HistoricalTradesResponseBuilder) Commission() float64 {
	return message.Float64VLS(m.Unsafe(), 99, 91)
}

// OpenFillExecutionID
func (m HistoricalTradesResponse) OpenFillExecutionID() string {
	return message.StringVLS(m.Unsafe(), 103, 99)
}

// OpenFillExecutionID
func (m HistoricalTradesResponseBuilder) OpenFillExecutionID() string {
	return message.StringVLS(m.Unsafe(), 103, 99)
}

// CloseFillExecutionID
func (m HistoricalTradesResponse) CloseFillExecutionID() string {
	return message.StringVLS(m.Unsafe(), 107, 103)
}

// CloseFillExecutionID
func (m HistoricalTradesResponseBuilder) CloseFillExecutionID() string {
	return message.StringVLS(m.Unsafe(), 107, 103)
}

// SubAccountIdentifier
func (m HistoricalTradesResponse) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 111, 107)
}

// SubAccountIdentifier
func (m HistoricalTradesResponseBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 111, 107)
}

// SetRequestID
func (m HistoricalTradesResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 10, 6, value)
}

// SetIsFinalMessage
func (m HistoricalTradesResponseBuilder) SetIsFinalMessage(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 11, 10, value)
}

// SetSymbol
func (m *HistoricalTradesResponseBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 15, 11, value)
}

// SetTradeAccount
func (m *HistoricalTradesResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 19, 15, value)
}

// SetEntryDateTime
func (m HistoricalTradesResponseBuilder) SetEntryDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Unsafe(), 27, 19, float64(value))
}

// SetExitDateTime
func (m HistoricalTradesResponseBuilder) SetExitDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Unsafe(), 35, 27, float64(value))
}

// SetEntryPrice
func (m HistoricalTradesResponseBuilder) SetEntryPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 43, 35, value)
}

// SetExitPrice
func (m HistoricalTradesResponseBuilder) SetExitPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 51, 43, value)
}

// SetTradeType
func (m HistoricalTradesResponseBuilder) SetTradeType(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 55, 51, int32(value))
}

// SetEntryQuantity
func (m HistoricalTradesResponseBuilder) SetEntryQuantity(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 59, 55, value)
}

// SetExitQuantity
func (m HistoricalTradesResponseBuilder) SetExitQuantity(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 63, 59, value)
}

// SetMaxOpenQuantity
func (m HistoricalTradesResponseBuilder) SetMaxOpenQuantity(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 67, 63, value)
}

// SetClosedProfitLoss
func (m HistoricalTradesResponseBuilder) SetClosedProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 75, 67, value)
}

// SetMaximumOpenPositionLoss
func (m HistoricalTradesResponseBuilder) SetMaximumOpenPositionLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 83, 75, value)
}

// SetMaximumOpenPositionProfit
func (m HistoricalTradesResponseBuilder) SetMaximumOpenPositionProfit(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 91, 83, value)
}

// SetCommission
func (m HistoricalTradesResponseBuilder) SetCommission(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 99, 91, value)
}

// SetOpenFillExecutionID
func (m *HistoricalTradesResponseBuilder) SetOpenFillExecutionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 103, 99, value)
}

// SetCloseFillExecutionID
func (m *HistoricalTradesResponseBuilder) SetCloseFillExecutionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 107, 103, value)
}

// SetSubAccountIdentifier
func (m HistoricalTradesResponseBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 111, 107, value)
}

// Copy
func (m HistoricalTradesResponse) Copy(to HistoricalTradesResponseBuilder) {
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
func (m HistoricalTradesResponseBuilder) Copy(to HistoricalTradesResponseBuilder) {
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
func (m HistoricalTradesResponse) CopyPointer(to HistoricalTradesResponsePointerBuilder) {
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
func (m HistoricalTradesResponseBuilder) CopyPointer(to HistoricalTradesResponsePointerBuilder) {
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
func (m HistoricalTradesResponse) CopyToPointer(to HistoricalTradesResponseFixedPointerBuilder) {
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
func (m HistoricalTradesResponseBuilder) CopyToPointer(to HistoricalTradesResponseFixedPointerBuilder) {
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
func (m HistoricalTradesResponse) CopyTo(to HistoricalTradesResponseFixedBuilder) {
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
func (m HistoricalTradesResponseBuilder) CopyTo(to HistoricalTradesResponseFixedBuilder) {
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
