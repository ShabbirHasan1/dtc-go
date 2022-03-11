package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalTradesResponseFixedPointer struct {
	message.FixedPointer
}

type HistoricalTradesResponseFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalTradesResponseFixed() HistoricalTradesResponseFixedPointerBuilder {
	a := HistoricalTradesResponseFixedPointerBuilder{message.AllocFixed(317)}
	a.Ptr.SetBytes(0, _HistoricalTradesResponseFixedDefault)
	return a
}

func AllocHistoricalTradesResponseFixedFrom(b []byte) HistoricalTradesResponseFixedPointer {
	return HistoricalTradesResponseFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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
func (m HistoricalTradesResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalTradesResponseFixedDefault)
}

// ToBuilder
func (m HistoricalTradesResponseFixedPointer) ToBuilder() HistoricalTradesResponseFixedPointerBuilder {
	return HistoricalTradesResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalTradesResponseFixedPointerBuilder) Finish() HistoricalTradesResponseFixedPointer {
	return HistoricalTradesResponseFixedPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalTradesResponseFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalTradesResponseFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// IsFinalMessage
func (m HistoricalTradesResponseFixedPointer) IsFinalMessage() uint8 {
	return message.Uint8Fixed(m.Ptr, 9, 8)
}

// IsFinalMessage
func (m HistoricalTradesResponseFixedPointerBuilder) IsFinalMessage() uint8 {
	return message.Uint8Fixed(m.Ptr, 9, 8)
}

// Symbol
func (m HistoricalTradesResponseFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 73, 9)
}

// Symbol
func (m HistoricalTradesResponseFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 73, 9)
}

// TradeAccount
func (m HistoricalTradesResponseFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 105, 73)
}

// TradeAccount
func (m HistoricalTradesResponseFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 105, 73)
}

// EntryDateTime
func (m HistoricalTradesResponseFixedPointer) EntryDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 113, 105))
}

// EntryDateTime
func (m HistoricalTradesResponseFixedPointerBuilder) EntryDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 113, 105))
}

// ExitDateTime
func (m HistoricalTradesResponseFixedPointer) ExitDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 121, 113))
}

// ExitDateTime
func (m HistoricalTradesResponseFixedPointerBuilder) ExitDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 121, 113))
}

// EntryPrice
func (m HistoricalTradesResponseFixedPointer) EntryPrice() float64 {
	return message.Float64Fixed(m.Ptr, 129, 121)
}

// EntryPrice
func (m HistoricalTradesResponseFixedPointerBuilder) EntryPrice() float64 {
	return message.Float64Fixed(m.Ptr, 129, 121)
}

// ExitPrice
func (m HistoricalTradesResponseFixedPointer) ExitPrice() float64 {
	return message.Float64Fixed(m.Ptr, 137, 129)
}

// ExitPrice
func (m HistoricalTradesResponseFixedPointerBuilder) ExitPrice() float64 {
	return message.Float64Fixed(m.Ptr, 137, 129)
}

// TradeType
func (m HistoricalTradesResponseFixedPointer) TradeType() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 141, 137))
}

// TradeType
func (m HistoricalTradesResponseFixedPointerBuilder) TradeType() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 141, 137))
}

// EntryQuantity
func (m HistoricalTradesResponseFixedPointer) EntryQuantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 145, 141)
}

// EntryQuantity
func (m HistoricalTradesResponseFixedPointerBuilder) EntryQuantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 145, 141)
}

// ExitQuantity
func (m HistoricalTradesResponseFixedPointer) ExitQuantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 149, 145)
}

// ExitQuantity
func (m HistoricalTradesResponseFixedPointerBuilder) ExitQuantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 149, 145)
}

// MaxOpenQuantity
func (m HistoricalTradesResponseFixedPointer) MaxOpenQuantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 153, 149)
}

// MaxOpenQuantity
func (m HistoricalTradesResponseFixedPointerBuilder) MaxOpenQuantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 153, 149)
}

// ClosedProfitLoss
func (m HistoricalTradesResponseFixedPointer) ClosedProfitLoss() float64 {
	return message.Float64Fixed(m.Ptr, 161, 153)
}

// ClosedProfitLoss
func (m HistoricalTradesResponseFixedPointerBuilder) ClosedProfitLoss() float64 {
	return message.Float64Fixed(m.Ptr, 161, 153)
}

// MaximumOpenPositionLoss
func (m HistoricalTradesResponseFixedPointer) MaximumOpenPositionLoss() float64 {
	return message.Float64Fixed(m.Ptr, 169, 161)
}

// MaximumOpenPositionLoss
func (m HistoricalTradesResponseFixedPointerBuilder) MaximumOpenPositionLoss() float64 {
	return message.Float64Fixed(m.Ptr, 169, 161)
}

// MaximumOpenPositionProfit
func (m HistoricalTradesResponseFixedPointer) MaximumOpenPositionProfit() float64 {
	return message.Float64Fixed(m.Ptr, 177, 169)
}

// MaximumOpenPositionProfit
func (m HistoricalTradesResponseFixedPointerBuilder) MaximumOpenPositionProfit() float64 {
	return message.Float64Fixed(m.Ptr, 177, 169)
}

// Commission
func (m HistoricalTradesResponseFixedPointer) Commission() float64 {
	return message.Float64Fixed(m.Ptr, 185, 177)
}

// Commission
func (m HistoricalTradesResponseFixedPointerBuilder) Commission() float64 {
	return message.Float64Fixed(m.Ptr, 185, 177)
}

// OpenFillExecutionID
func (m HistoricalTradesResponseFixedPointer) OpenFillExecutionID() string {
	return message.StringFixed(m.Ptr, 249, 185)
}

// OpenFillExecutionID
func (m HistoricalTradesResponseFixedPointerBuilder) OpenFillExecutionID() string {
	return message.StringFixed(m.Ptr, 249, 185)
}

// CloseFillExecutionID
func (m HistoricalTradesResponseFixedPointer) CloseFillExecutionID() string {
	return message.StringFixed(m.Ptr, 313, 249)
}

// CloseFillExecutionID
func (m HistoricalTradesResponseFixedPointerBuilder) CloseFillExecutionID() string {
	return message.StringFixed(m.Ptr, 313, 249)
}

// SubAccountIdentifier
func (m HistoricalTradesResponseFixedPointer) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Ptr, 317, 313)
}

// SubAccountIdentifier
func (m HistoricalTradesResponseFixedPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Ptr, 317, 313)
}

// SetRequestID
func (m HistoricalTradesResponseFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetIsFinalMessage
func (m HistoricalTradesResponseFixedPointerBuilder) SetIsFinalMessage(value uint8) {
	message.SetUint8Fixed(m.Ptr, 9, 8, value)
}

// SetSymbol
func (m HistoricalTradesResponseFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 73, 9, value)
}

// SetTradeAccount
func (m HistoricalTradesResponseFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 105, 73, value)
}

// SetEntryDateTime
func (m HistoricalTradesResponseFixedPointerBuilder) SetEntryDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 113, 105, float64(value))
}

// SetExitDateTime
func (m HistoricalTradesResponseFixedPointerBuilder) SetExitDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 121, 113, float64(value))
}

// SetEntryPrice
func (m HistoricalTradesResponseFixedPointerBuilder) SetEntryPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 129, 121, value)
}

// SetExitPrice
func (m HistoricalTradesResponseFixedPointerBuilder) SetExitPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 137, 129, value)
}

// SetTradeType
func (m HistoricalTradesResponseFixedPointerBuilder) SetTradeType(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 141, 137, int32(value))
}

// SetEntryQuantity
func (m HistoricalTradesResponseFixedPointerBuilder) SetEntryQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 145, 141, value)
}

// SetExitQuantity
func (m HistoricalTradesResponseFixedPointerBuilder) SetExitQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 149, 145, value)
}

// SetMaxOpenQuantity
func (m HistoricalTradesResponseFixedPointerBuilder) SetMaxOpenQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 153, 149, value)
}

// SetClosedProfitLoss
func (m HistoricalTradesResponseFixedPointerBuilder) SetClosedProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Ptr, 161, 153, value)
}

// SetMaximumOpenPositionLoss
func (m HistoricalTradesResponseFixedPointerBuilder) SetMaximumOpenPositionLoss(value float64) {
	message.SetFloat64Fixed(m.Ptr, 169, 161, value)
}

// SetMaximumOpenPositionProfit
func (m HistoricalTradesResponseFixedPointerBuilder) SetMaximumOpenPositionProfit(value float64) {
	message.SetFloat64Fixed(m.Ptr, 177, 169, value)
}

// SetCommission
func (m HistoricalTradesResponseFixedPointerBuilder) SetCommission(value float64) {
	message.SetFloat64Fixed(m.Ptr, 185, 177, value)
}

// SetOpenFillExecutionID
func (m HistoricalTradesResponseFixedPointerBuilder) SetOpenFillExecutionID(value string) {
	message.SetStringFixed(m.Ptr, 249, 185, value)
}

// SetCloseFillExecutionID
func (m HistoricalTradesResponseFixedPointerBuilder) SetCloseFillExecutionID(value string) {
	message.SetStringFixed(m.Ptr, 313, 249, value)
}

// SetSubAccountIdentifier
func (m HistoricalTradesResponseFixedPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Ptr, 317, 313, value)
}

// Copy
func (m HistoricalTradesResponseFixedPointer) Copy(to HistoricalTradesResponseFixedBuilder) {
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
func (m HistoricalTradesResponseFixedPointerBuilder) Copy(to HistoricalTradesResponseFixedBuilder) {
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
func (m HistoricalTradesResponseFixedPointer) CopyPointer(to HistoricalTradesResponseFixedPointerBuilder) {
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
func (m HistoricalTradesResponseFixedPointerBuilder) CopyPointer(to HistoricalTradesResponseFixedPointerBuilder) {
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
func (m HistoricalTradesResponseFixedPointer) CopyTo(to HistoricalTradesResponseBuilder) {
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
func (m HistoricalTradesResponseFixedPointerBuilder) CopyTo(to HistoricalTradesResponseBuilder) {
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
func (m HistoricalTradesResponseFixedPointer) CopyToPointer(to HistoricalTradesResponsePointerBuilder) {
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
func (m HistoricalTradesResponseFixedPointerBuilder) CopyToPointer(to HistoricalTradesResponsePointerBuilder) {
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
