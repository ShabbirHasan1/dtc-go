package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                      = HistoricalTradesResponseFixedSize  (317)
//     Type                      = HISTORICAL_TRADES_RESPONSE  (10102)
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
var _HistoricalTradesResponseFixedDefault = []byte{61, 1, 118, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalTradesResponseFixedSize = 317

type HistoricalTradesResponseFixed struct {
	message.Fixed
}

type HistoricalTradesResponseFixedBuilder struct {
	message.Fixed
}

func NewHistoricalTradesResponseFixedFrom(b []byte) HistoricalTradesResponseFixed {
	return HistoricalTradesResponseFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalTradesResponseFixed(b []byte) HistoricalTradesResponseFixed {
	return HistoricalTradesResponseFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalTradesResponseFixed() HistoricalTradesResponseFixedBuilder {
	a := HistoricalTradesResponseFixedBuilder{message.NewFixed(317)}
	a.Unsafe().SetBytes(0, _HistoricalTradesResponseFixedDefault)
	return a
}

// Clear
// {
//     Size                      = HistoricalTradesResponseFixedSize  (317)
//     Type                      = HISTORICAL_TRADES_RESPONSE  (10102)
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
func (m HistoricalTradesResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalTradesResponseFixedDefault)
}

// ToBuilder
func (m HistoricalTradesResponseFixed) ToBuilder() HistoricalTradesResponseFixedBuilder {
	return HistoricalTradesResponseFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalTradesResponseFixedBuilder) Finish() HistoricalTradesResponseFixed {
	return HistoricalTradesResponseFixed{m.Fixed}
}

// RequestID
func (m HistoricalTradesResponseFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalTradesResponseFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// IsFinalMessage
func (m HistoricalTradesResponseFixed) IsFinalMessage() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 9, 8)
}

// IsFinalMessage
func (m HistoricalTradesResponseFixedBuilder) IsFinalMessage() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 9, 8)
}

// Symbol
func (m HistoricalTradesResponseFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 73, 9)
}

// Symbol
func (m HistoricalTradesResponseFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 73, 9)
}

// TradeAccount
func (m HistoricalTradesResponseFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 105, 73)
}

// TradeAccount
func (m HistoricalTradesResponseFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 105, 73)
}

// EntryDateTime
func (m HistoricalTradesResponseFixed) EntryDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 113, 105))
}

// EntryDateTime
func (m HistoricalTradesResponseFixedBuilder) EntryDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 113, 105))
}

// ExitDateTime
func (m HistoricalTradesResponseFixed) ExitDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 121, 113))
}

// ExitDateTime
func (m HistoricalTradesResponseFixedBuilder) ExitDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 121, 113))
}

// EntryPrice
func (m HistoricalTradesResponseFixed) EntryPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 129, 121)
}

// EntryPrice
func (m HistoricalTradesResponseFixedBuilder) EntryPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 129, 121)
}

// ExitPrice
func (m HistoricalTradesResponseFixed) ExitPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 137, 129)
}

// ExitPrice
func (m HistoricalTradesResponseFixedBuilder) ExitPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 137, 129)
}

// TradeType
func (m HistoricalTradesResponseFixed) TradeType() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 141, 137))
}

// TradeType
func (m HistoricalTradesResponseFixedBuilder) TradeType() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 141, 137))
}

// EntryQuantity
func (m HistoricalTradesResponseFixed) EntryQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 145, 141)
}

// EntryQuantity
func (m HistoricalTradesResponseFixedBuilder) EntryQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 145, 141)
}

// ExitQuantity
func (m HistoricalTradesResponseFixed) ExitQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 149, 145)
}

// ExitQuantity
func (m HistoricalTradesResponseFixedBuilder) ExitQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 149, 145)
}

// MaxOpenQuantity
func (m HistoricalTradesResponseFixed) MaxOpenQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 153, 149)
}

// MaxOpenQuantity
func (m HistoricalTradesResponseFixedBuilder) MaxOpenQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 153, 149)
}

// ClosedProfitLoss
func (m HistoricalTradesResponseFixed) ClosedProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 161, 153)
}

// ClosedProfitLoss
func (m HistoricalTradesResponseFixedBuilder) ClosedProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 161, 153)
}

// MaximumOpenPositionLoss
func (m HistoricalTradesResponseFixed) MaximumOpenPositionLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 169, 161)
}

// MaximumOpenPositionLoss
func (m HistoricalTradesResponseFixedBuilder) MaximumOpenPositionLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 169, 161)
}

// MaximumOpenPositionProfit
func (m HistoricalTradesResponseFixed) MaximumOpenPositionProfit() float64 {
	return message.Float64Fixed(m.Unsafe(), 177, 169)
}

// MaximumOpenPositionProfit
func (m HistoricalTradesResponseFixedBuilder) MaximumOpenPositionProfit() float64 {
	return message.Float64Fixed(m.Unsafe(), 177, 169)
}

// Commission
func (m HistoricalTradesResponseFixed) Commission() float64 {
	return message.Float64Fixed(m.Unsafe(), 185, 177)
}

// Commission
func (m HistoricalTradesResponseFixedBuilder) Commission() float64 {
	return message.Float64Fixed(m.Unsafe(), 185, 177)
}

// OpenFillExecutionID
func (m HistoricalTradesResponseFixed) OpenFillExecutionID() string {
	return message.StringFixed(m.Unsafe(), 249, 185)
}

// OpenFillExecutionID
func (m HistoricalTradesResponseFixedBuilder) OpenFillExecutionID() string {
	return message.StringFixed(m.Unsafe(), 249, 185)
}

// CloseFillExecutionID
func (m HistoricalTradesResponseFixed) CloseFillExecutionID() string {
	return message.StringFixed(m.Unsafe(), 313, 249)
}

// CloseFillExecutionID
func (m HistoricalTradesResponseFixedBuilder) CloseFillExecutionID() string {
	return message.StringFixed(m.Unsafe(), 313, 249)
}

// SubAccountIdentifier
func (m HistoricalTradesResponseFixed) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 317, 313)
}

// SubAccountIdentifier
func (m HistoricalTradesResponseFixedBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 317, 313)
}

// SetRequestID
func (m HistoricalTradesResponseFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetIsFinalMessage
func (m HistoricalTradesResponseFixedBuilder) SetIsFinalMessage(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 9, 8, value)
}

// SetSymbol
func (m HistoricalTradesResponseFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 73, 9, value)
}

// SetTradeAccount
func (m HistoricalTradesResponseFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 105, 73, value)
}

// SetEntryDateTime
func (m HistoricalTradesResponseFixedBuilder) SetEntryDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 113, 105, float64(value))
}

// SetExitDateTime
func (m HistoricalTradesResponseFixedBuilder) SetExitDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 121, 113, float64(value))
}

// SetEntryPrice
func (m HistoricalTradesResponseFixedBuilder) SetEntryPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 129, 121, value)
}

// SetExitPrice
func (m HistoricalTradesResponseFixedBuilder) SetExitPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 137, 129, value)
}

// SetTradeType
func (m HistoricalTradesResponseFixedBuilder) SetTradeType(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 141, 137, int32(value))
}

// SetEntryQuantity
func (m HistoricalTradesResponseFixedBuilder) SetEntryQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 145, 141, value)
}

// SetExitQuantity
func (m HistoricalTradesResponseFixedBuilder) SetExitQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 149, 145, value)
}

// SetMaxOpenQuantity
func (m HistoricalTradesResponseFixedBuilder) SetMaxOpenQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 153, 149, value)
}

// SetClosedProfitLoss
func (m HistoricalTradesResponseFixedBuilder) SetClosedProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 161, 153, value)
}

// SetMaximumOpenPositionLoss
func (m HistoricalTradesResponseFixedBuilder) SetMaximumOpenPositionLoss(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 169, 161, value)
}

// SetMaximumOpenPositionProfit
func (m HistoricalTradesResponseFixedBuilder) SetMaximumOpenPositionProfit(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 177, 169, value)
}

// SetCommission
func (m HistoricalTradesResponseFixedBuilder) SetCommission(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 185, 177, value)
}

// SetOpenFillExecutionID
func (m HistoricalTradesResponseFixedBuilder) SetOpenFillExecutionID(value string) {
	message.SetStringFixed(m.Unsafe(), 249, 185, value)
}

// SetCloseFillExecutionID
func (m HistoricalTradesResponseFixedBuilder) SetCloseFillExecutionID(value string) {
	message.SetStringFixed(m.Unsafe(), 313, 249, value)
}

// SetSubAccountIdentifier
func (m HistoricalTradesResponseFixedBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 317, 313, value)
}

// Copy
func (m HistoricalTradesResponseFixed) Copy(to HistoricalTradesResponseFixedBuilder) {
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
func (m HistoricalTradesResponseFixedBuilder) Copy(to HistoricalTradesResponseFixedBuilder) {
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
func (m HistoricalTradesResponseFixed) CopyPointer(to HistoricalTradesResponseFixedPointerBuilder) {
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
func (m HistoricalTradesResponseFixedBuilder) CopyPointer(to HistoricalTradesResponseFixedPointerBuilder) {
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
func (m HistoricalTradesResponseFixed) CopyToPointer(to HistoricalTradesResponsePointerBuilder) {
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
func (m HistoricalTradesResponseFixedBuilder) CopyToPointer(to HistoricalTradesResponsePointerBuilder) {
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
func (m HistoricalTradesResponseFixed) CopyTo(to HistoricalTradesResponseBuilder) {
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
func (m HistoricalTradesResponseFixedBuilder) CopyTo(to HistoricalTradesResponseBuilder) {
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
