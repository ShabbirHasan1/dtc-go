package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalTradesResponseSize = 111

const HistoricalTradesResponseFixedSize = 317

// {
//     Size                      = HistoricalTradesResponseSize  (111)
//     Type                      = HISTORICAL_TRADES_RESPONSE  (10102)
//     BaseSize                  = HistoricalTradesResponseSize  (111)
//     RequestID                 = 0
//     IsFinalMessage            = false
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

// {
//     Size                      = HistoricalTradesResponseFixedSize  (317)
//     Type                      = HISTORICAL_TRADES_RESPONSE  (10102)
//     RequestID                 = 0
//     IsFinalMessage            = false
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

type HistoricalTradesResponse struct {
	message.VLS
}

type HistoricalTradesResponseBuilder struct {
	message.VLSBuilder
}

type HistoricalTradesResponseFixed struct {
	message.Fixed
}

type HistoricalTradesResponseFixedBuilder struct {
	message.Fixed
}

type HistoricalTradesResponsePointer struct {
	message.VLSPointer
}

type HistoricalTradesResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

type HistoricalTradesResponseFixedPointer struct {
	message.FixedPointer
}

type HistoricalTradesResponseFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocHistoricalTradesResponse() HistoricalTradesResponsePointerBuilder {
	a := HistoricalTradesResponsePointerBuilder{message.AllocVLSBuilder(111)}
	a.Ptr.SetBytes(0, _HistoricalTradesResponseDefault)
	return a
}

func AllocHistoricalTradesResponseFrom(b []byte) HistoricalTradesResponsePointer {
	return HistoricalTradesResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size                      = HistoricalTradesResponseSize  (111)
//     Type                      = HISTORICAL_TRADES_RESPONSE  (10102)
//     BaseSize                  = HistoricalTradesResponseSize  (111)
//     RequestID                 = 0
//     IsFinalMessage            = false
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

// Clear
// {
//     Size                      = HistoricalTradesResponseFixedSize  (317)
//     Type                      = HISTORICAL_TRADES_RESPONSE  (10102)
//     RequestID                 = 0
//     IsFinalMessage            = false
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

// Clear
// {
//     Size                      = HistoricalTradesResponseSize  (111)
//     Type                      = HISTORICAL_TRADES_RESPONSE  (10102)
//     BaseSize                  = HistoricalTradesResponseSize  (111)
//     RequestID                 = 0
//     IsFinalMessage            = false
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

// Clear
// {
//     Size                      = HistoricalTradesResponseFixedSize  (317)
//     Type                      = HISTORICAL_TRADES_RESPONSE  (10102)
//     RequestID                 = 0
//     IsFinalMessage            = false
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
func (m HistoricalTradesResponse) ToBuilder() HistoricalTradesResponseBuilder {
	return HistoricalTradesResponseBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalTradesResponseFixed) ToBuilder() HistoricalTradesResponseFixedBuilder {
	return HistoricalTradesResponseFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalTradesResponsePointer) ToBuilder() HistoricalTradesResponsePointerBuilder {
	return HistoricalTradesResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalTradesResponseFixedPointer) ToBuilder() HistoricalTradesResponseFixedPointerBuilder {
	return HistoricalTradesResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalTradesResponseBuilder) Finish() HistoricalTradesResponse {
	return HistoricalTradesResponse{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalTradesResponseFixedBuilder) Finish() HistoricalTradesResponseFixed {
	return HistoricalTradesResponseFixed{m.Fixed}
}

// Finish
func (m *HistoricalTradesResponsePointerBuilder) Finish() HistoricalTradesResponsePointer {
	return HistoricalTradesResponsePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalTradesResponseFixedPointerBuilder) Finish() HistoricalTradesResponseFixedPointer {
	return HistoricalTradesResponseFixedPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalTradesResponse) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m HistoricalTradesResponseBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
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
func (m HistoricalTradesResponse) IsFinalMessage() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsFinalMessage
func (m HistoricalTradesResponseBuilder) IsFinalMessage() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsFinalMessage
func (m HistoricalTradesResponsePointer) IsFinalMessage() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// IsFinalMessage
func (m HistoricalTradesResponsePointerBuilder) IsFinalMessage() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// Symbol
func (m HistoricalTradesResponse) Symbol() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// Symbol
func (m HistoricalTradesResponseBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
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
func (m HistoricalTradesResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
}

// TradeAccount
func (m HistoricalTradesResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
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
func (m HistoricalTradesResponse) EntryDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 27, 19))
}

// EntryDateTime
func (m HistoricalTradesResponseBuilder) EntryDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 27, 19))
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
func (m HistoricalTradesResponse) ExitDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 35, 27))
}

// ExitDateTime
func (m HistoricalTradesResponseBuilder) ExitDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 35, 27))
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
func (m HistoricalTradesResponse) EntryPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 43, 35)
}

// EntryPrice
func (m HistoricalTradesResponseBuilder) EntryPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 43, 35)
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
func (m HistoricalTradesResponse) ExitPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 51, 43)
}

// ExitPrice
func (m HistoricalTradesResponseBuilder) ExitPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 51, 43)
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
func (m HistoricalTradesResponse) TradeType() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 55, 51))
}

// TradeType
func (m HistoricalTradesResponseBuilder) TradeType() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 55, 51))
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
func (m HistoricalTradesResponse) EntryQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 59, 55)
}

// EntryQuantity
func (m HistoricalTradesResponseBuilder) EntryQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 59, 55)
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
func (m HistoricalTradesResponse) ExitQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 63, 59)
}

// ExitQuantity
func (m HistoricalTradesResponseBuilder) ExitQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 63, 59)
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
func (m HistoricalTradesResponse) MaxOpenQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 67, 63)
}

// MaxOpenQuantity
func (m HistoricalTradesResponseBuilder) MaxOpenQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 67, 63)
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
func (m HistoricalTradesResponse) ClosedProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 75, 67)
}

// ClosedProfitLoss
func (m HistoricalTradesResponseBuilder) ClosedProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 75, 67)
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
func (m HistoricalTradesResponse) MaximumOpenPositionLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 83, 75)
}

// MaximumOpenPositionLoss
func (m HistoricalTradesResponseBuilder) MaximumOpenPositionLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 83, 75)
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
func (m HistoricalTradesResponse) MaximumOpenPositionProfit() float64 {
	return message.Float64VLS(m.Unsafe(), 91, 83)
}

// MaximumOpenPositionProfit
func (m HistoricalTradesResponseBuilder) MaximumOpenPositionProfit() float64 {
	return message.Float64VLS(m.Unsafe(), 91, 83)
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
func (m HistoricalTradesResponse) Commission() float64 {
	return message.Float64VLS(m.Unsafe(), 99, 91)
}

// Commission
func (m HistoricalTradesResponseBuilder) Commission() float64 {
	return message.Float64VLS(m.Unsafe(), 99, 91)
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
func (m HistoricalTradesResponse) OpenFillExecutionID() string {
	return message.StringVLS(m.Unsafe(), 103, 99)
}

// OpenFillExecutionID
func (m HistoricalTradesResponseBuilder) OpenFillExecutionID() string {
	return message.StringVLS(m.Unsafe(), 103, 99)
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
func (m HistoricalTradesResponse) CloseFillExecutionID() string {
	return message.StringVLS(m.Unsafe(), 107, 103)
}

// CloseFillExecutionID
func (m HistoricalTradesResponseBuilder) CloseFillExecutionID() string {
	return message.StringVLS(m.Unsafe(), 107, 103)
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
func (m HistoricalTradesResponse) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 111, 107)
}

// SubAccountIdentifier
func (m HistoricalTradesResponseBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 111, 107)
}

// SubAccountIdentifier
func (m HistoricalTradesResponsePointer) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 111, 107)
}

// SubAccountIdentifier
func (m HistoricalTradesResponsePointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 111, 107)
}

// RequestID
func (m HistoricalTradesResponseFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalTradesResponseFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m HistoricalTradesResponseFixed) IsFinalMessage() bool {
	return message.BoolFixed(m.Unsafe(), 9, 8)
}

// IsFinalMessage
func (m HistoricalTradesResponseFixedBuilder) IsFinalMessage() bool {
	return message.BoolFixed(m.Unsafe(), 9, 8)
}

// IsFinalMessage
func (m HistoricalTradesResponseFixedPointer) IsFinalMessage() bool {
	return message.BoolFixed(m.Ptr, 9, 8)
}

// IsFinalMessage
func (m HistoricalTradesResponseFixedPointerBuilder) IsFinalMessage() bool {
	return message.BoolFixed(m.Ptr, 9, 8)
}

// Symbol
func (m HistoricalTradesResponseFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 73, 9)
}

// Symbol
func (m HistoricalTradesResponseFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 73, 9)
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
func (m HistoricalTradesResponseFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 105, 73)
}

// TradeAccount
func (m HistoricalTradesResponseFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 105, 73)
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
func (m HistoricalTradesResponseFixed) EntryDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 113, 105))
}

// EntryDateTime
func (m HistoricalTradesResponseFixedBuilder) EntryDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 113, 105))
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
func (m HistoricalTradesResponseFixed) ExitDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 121, 113))
}

// ExitDateTime
func (m HistoricalTradesResponseFixedBuilder) ExitDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 121, 113))
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
func (m HistoricalTradesResponseFixed) EntryPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 129, 121)
}

// EntryPrice
func (m HistoricalTradesResponseFixedBuilder) EntryPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 129, 121)
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
func (m HistoricalTradesResponseFixed) ExitPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 137, 129)
}

// ExitPrice
func (m HistoricalTradesResponseFixedBuilder) ExitPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 137, 129)
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
func (m HistoricalTradesResponseFixed) TradeType() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 141, 137))
}

// TradeType
func (m HistoricalTradesResponseFixedBuilder) TradeType() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 141, 137))
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
func (m HistoricalTradesResponseFixed) EntryQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 145, 141)
}

// EntryQuantity
func (m HistoricalTradesResponseFixedBuilder) EntryQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 145, 141)
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
func (m HistoricalTradesResponseFixed) ExitQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 149, 145)
}

// ExitQuantity
func (m HistoricalTradesResponseFixedBuilder) ExitQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 149, 145)
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
func (m HistoricalTradesResponseFixed) MaxOpenQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 153, 149)
}

// MaxOpenQuantity
func (m HistoricalTradesResponseFixedBuilder) MaxOpenQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 153, 149)
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
func (m HistoricalTradesResponseFixed) ClosedProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 161, 153)
}

// ClosedProfitLoss
func (m HistoricalTradesResponseFixedBuilder) ClosedProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 161, 153)
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
func (m HistoricalTradesResponseFixed) MaximumOpenPositionLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 169, 161)
}

// MaximumOpenPositionLoss
func (m HistoricalTradesResponseFixedBuilder) MaximumOpenPositionLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 169, 161)
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
func (m HistoricalTradesResponseFixed) MaximumOpenPositionProfit() float64 {
	return message.Float64Fixed(m.Unsafe(), 177, 169)
}

// MaximumOpenPositionProfit
func (m HistoricalTradesResponseFixedBuilder) MaximumOpenPositionProfit() float64 {
	return message.Float64Fixed(m.Unsafe(), 177, 169)
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
func (m HistoricalTradesResponseFixed) Commission() float64 {
	return message.Float64Fixed(m.Unsafe(), 185, 177)
}

// Commission
func (m HistoricalTradesResponseFixedBuilder) Commission() float64 {
	return message.Float64Fixed(m.Unsafe(), 185, 177)
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
func (m HistoricalTradesResponseFixed) OpenFillExecutionID() string {
	return message.StringFixed(m.Unsafe(), 249, 185)
}

// OpenFillExecutionID
func (m HistoricalTradesResponseFixedBuilder) OpenFillExecutionID() string {
	return message.StringFixed(m.Unsafe(), 249, 185)
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
func (m HistoricalTradesResponseFixed) CloseFillExecutionID() string {
	return message.StringFixed(m.Unsafe(), 313, 249)
}

// CloseFillExecutionID
func (m HistoricalTradesResponseFixedBuilder) CloseFillExecutionID() string {
	return message.StringFixed(m.Unsafe(), 313, 249)
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
func (m HistoricalTradesResponseFixed) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 317, 313)
}

// SubAccountIdentifier
func (m HistoricalTradesResponseFixedBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 317, 313)
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
func (m HistoricalTradesResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m HistoricalTradesResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 10, 6, value)
}

// SetIsFinalMessage
func (m HistoricalTradesResponseBuilder) SetIsFinalMessage(value bool) {
	message.SetBoolVLS(m.Unsafe(), 11, 10, value)
}

// SetIsFinalMessage
func (m HistoricalTradesResponsePointerBuilder) SetIsFinalMessage(value bool) {
	message.SetBoolVLS(m.Ptr, 11, 10, value)
}

// SetSymbol
func (m *HistoricalTradesResponseBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 15, 11, value)
}

// SetSymbol
func (m *HistoricalTradesResponsePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetTradeAccount
func (m *HistoricalTradesResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 19, 15, value)
}

// SetTradeAccount
func (m *HistoricalTradesResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 19, 15, value)
}

// SetEntryDateTime
func (m HistoricalTradesResponseBuilder) SetEntryDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Unsafe(), 27, 19, float64(value))
}

// SetEntryDateTime
func (m HistoricalTradesResponsePointerBuilder) SetEntryDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Ptr, 27, 19, float64(value))
}

// SetExitDateTime
func (m HistoricalTradesResponseBuilder) SetExitDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Unsafe(), 35, 27, float64(value))
}

// SetExitDateTime
func (m HistoricalTradesResponsePointerBuilder) SetExitDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Ptr, 35, 27, float64(value))
}

// SetEntryPrice
func (m HistoricalTradesResponseBuilder) SetEntryPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 43, 35, value)
}

// SetEntryPrice
func (m HistoricalTradesResponsePointerBuilder) SetEntryPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 43, 35, value)
}

// SetExitPrice
func (m HistoricalTradesResponseBuilder) SetExitPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 51, 43, value)
}

// SetExitPrice
func (m HistoricalTradesResponsePointerBuilder) SetExitPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 51, 43, value)
}

// SetTradeType
func (m HistoricalTradesResponseBuilder) SetTradeType(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 55, 51, int32(value))
}

// SetTradeType
func (m HistoricalTradesResponsePointerBuilder) SetTradeType(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 55, 51, int32(value))
}

// SetEntryQuantity
func (m HistoricalTradesResponseBuilder) SetEntryQuantity(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 59, 55, value)
}

// SetEntryQuantity
func (m HistoricalTradesResponsePointerBuilder) SetEntryQuantity(value uint32) {
	message.SetUint32VLS(m.Ptr, 59, 55, value)
}

// SetExitQuantity
func (m HistoricalTradesResponseBuilder) SetExitQuantity(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 63, 59, value)
}

// SetExitQuantity
func (m HistoricalTradesResponsePointerBuilder) SetExitQuantity(value uint32) {
	message.SetUint32VLS(m.Ptr, 63, 59, value)
}

// SetMaxOpenQuantity
func (m HistoricalTradesResponseBuilder) SetMaxOpenQuantity(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 67, 63, value)
}

// SetMaxOpenQuantity
func (m HistoricalTradesResponsePointerBuilder) SetMaxOpenQuantity(value uint32) {
	message.SetUint32VLS(m.Ptr, 67, 63, value)
}

// SetClosedProfitLoss
func (m HistoricalTradesResponseBuilder) SetClosedProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 75, 67, value)
}

// SetClosedProfitLoss
func (m HistoricalTradesResponsePointerBuilder) SetClosedProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 75, 67, value)
}

// SetMaximumOpenPositionLoss
func (m HistoricalTradesResponseBuilder) SetMaximumOpenPositionLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 83, 75, value)
}

// SetMaximumOpenPositionLoss
func (m HistoricalTradesResponsePointerBuilder) SetMaximumOpenPositionLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 83, 75, value)
}

// SetMaximumOpenPositionProfit
func (m HistoricalTradesResponseBuilder) SetMaximumOpenPositionProfit(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 91, 83, value)
}

// SetMaximumOpenPositionProfit
func (m HistoricalTradesResponsePointerBuilder) SetMaximumOpenPositionProfit(value float64) {
	message.SetFloat64VLS(m.Ptr, 91, 83, value)
}

// SetCommission
func (m HistoricalTradesResponseBuilder) SetCommission(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 99, 91, value)
}

// SetCommission
func (m HistoricalTradesResponsePointerBuilder) SetCommission(value float64) {
	message.SetFloat64VLS(m.Ptr, 99, 91, value)
}

// SetOpenFillExecutionID
func (m *HistoricalTradesResponseBuilder) SetOpenFillExecutionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 103, 99, value)
}

// SetOpenFillExecutionID
func (m *HistoricalTradesResponsePointerBuilder) SetOpenFillExecutionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 103, 99, value)
}

// SetCloseFillExecutionID
func (m *HistoricalTradesResponseBuilder) SetCloseFillExecutionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 107, 103, value)
}

// SetCloseFillExecutionID
func (m *HistoricalTradesResponsePointerBuilder) SetCloseFillExecutionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 107, 103, value)
}

// SetSubAccountIdentifier
func (m HistoricalTradesResponseBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 111, 107, value)
}

// SetSubAccountIdentifier
func (m HistoricalTradesResponsePointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Ptr, 111, 107, value)
}

// SetRequestID
func (m HistoricalTradesResponseFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m HistoricalTradesResponseFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetIsFinalMessage
func (m HistoricalTradesResponseFixedBuilder) SetIsFinalMessage(value bool) {
	message.SetBoolFixed(m.Unsafe(), 9, 8, value)
}

// SetIsFinalMessage
func (m HistoricalTradesResponseFixedPointerBuilder) SetIsFinalMessage(value bool) {
	message.SetBoolFixed(m.Ptr, 9, 8, value)
}

// SetSymbol
func (m HistoricalTradesResponseFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 73, 9, value)
}

// SetSymbol
func (m HistoricalTradesResponseFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 73, 9, value)
}

// SetTradeAccount
func (m HistoricalTradesResponseFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 105, 73, value)
}

// SetTradeAccount
func (m HistoricalTradesResponseFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 105, 73, value)
}

// SetEntryDateTime
func (m HistoricalTradesResponseFixedBuilder) SetEntryDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 113, 105, float64(value))
}

// SetEntryDateTime
func (m HistoricalTradesResponseFixedPointerBuilder) SetEntryDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 113, 105, float64(value))
}

// SetExitDateTime
func (m HistoricalTradesResponseFixedBuilder) SetExitDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 121, 113, float64(value))
}

// SetExitDateTime
func (m HistoricalTradesResponseFixedPointerBuilder) SetExitDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 121, 113, float64(value))
}

// SetEntryPrice
func (m HistoricalTradesResponseFixedBuilder) SetEntryPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 129, 121, value)
}

// SetEntryPrice
func (m HistoricalTradesResponseFixedPointerBuilder) SetEntryPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 129, 121, value)
}

// SetExitPrice
func (m HistoricalTradesResponseFixedBuilder) SetExitPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 137, 129, value)
}

// SetExitPrice
func (m HistoricalTradesResponseFixedPointerBuilder) SetExitPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 137, 129, value)
}

// SetTradeType
func (m HistoricalTradesResponseFixedBuilder) SetTradeType(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 141, 137, int32(value))
}

// SetTradeType
func (m HistoricalTradesResponseFixedPointerBuilder) SetTradeType(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 141, 137, int32(value))
}

// SetEntryQuantity
func (m HistoricalTradesResponseFixedBuilder) SetEntryQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 145, 141, value)
}

// SetEntryQuantity
func (m HistoricalTradesResponseFixedPointerBuilder) SetEntryQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 145, 141, value)
}

// SetExitQuantity
func (m HistoricalTradesResponseFixedBuilder) SetExitQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 149, 145, value)
}

// SetExitQuantity
func (m HistoricalTradesResponseFixedPointerBuilder) SetExitQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 149, 145, value)
}

// SetMaxOpenQuantity
func (m HistoricalTradesResponseFixedBuilder) SetMaxOpenQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 153, 149, value)
}

// SetMaxOpenQuantity
func (m HistoricalTradesResponseFixedPointerBuilder) SetMaxOpenQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 153, 149, value)
}

// SetClosedProfitLoss
func (m HistoricalTradesResponseFixedBuilder) SetClosedProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 161, 153, value)
}

// SetClosedProfitLoss
func (m HistoricalTradesResponseFixedPointerBuilder) SetClosedProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Ptr, 161, 153, value)
}

// SetMaximumOpenPositionLoss
func (m HistoricalTradesResponseFixedBuilder) SetMaximumOpenPositionLoss(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 169, 161, value)
}

// SetMaximumOpenPositionLoss
func (m HistoricalTradesResponseFixedPointerBuilder) SetMaximumOpenPositionLoss(value float64) {
	message.SetFloat64Fixed(m.Ptr, 169, 161, value)
}

// SetMaximumOpenPositionProfit
func (m HistoricalTradesResponseFixedBuilder) SetMaximumOpenPositionProfit(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 177, 169, value)
}

// SetMaximumOpenPositionProfit
func (m HistoricalTradesResponseFixedPointerBuilder) SetMaximumOpenPositionProfit(value float64) {
	message.SetFloat64Fixed(m.Ptr, 177, 169, value)
}

// SetCommission
func (m HistoricalTradesResponseFixedBuilder) SetCommission(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 185, 177, value)
}

// SetCommission
func (m HistoricalTradesResponseFixedPointerBuilder) SetCommission(value float64) {
	message.SetFloat64Fixed(m.Ptr, 185, 177, value)
}

// SetOpenFillExecutionID
func (m HistoricalTradesResponseFixedBuilder) SetOpenFillExecutionID(value string) {
	message.SetStringFixed(m.Unsafe(), 249, 185, value)
}

// SetOpenFillExecutionID
func (m HistoricalTradesResponseFixedPointerBuilder) SetOpenFillExecutionID(value string) {
	message.SetStringFixed(m.Ptr, 249, 185, value)
}

// SetCloseFillExecutionID
func (m HistoricalTradesResponseFixedBuilder) SetCloseFillExecutionID(value string) {
	message.SetStringFixed(m.Unsafe(), 313, 249, value)
}

// SetCloseFillExecutionID
func (m HistoricalTradesResponseFixedPointerBuilder) SetCloseFillExecutionID(value string) {
	message.SetStringFixed(m.Ptr, 313, 249, value)
}

// SetSubAccountIdentifier
func (m HistoricalTradesResponseFixedBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 317, 313, value)
}

// SetSubAccountIdentifier
func (m HistoricalTradesResponseFixedPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Ptr, 317, 313, value)
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
