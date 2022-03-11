package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const TradePositionConsolidatedSize = 153

// {
//     Size                                  = TradePositionConsolidatedSize  (153)
//     Type                                  = TRADE_POSITION_CONSOLIDATED  (10113)
//     BaseSize                              = TradePositionConsolidatedSize  (153)
//     IsDeleted                             = false
//     Symbol                                = ""
//     IsSimulated                           = false
//     TradeAccount                          = ""
//     CurrencyCode                          = ""
//     Quantity                              = 0
//     AveragePrice                          = 0
//     OpenProfitLoss                        = 0
//     DailyProfitLoss                       = 0
//     LastDailyProfitLossResetDateTimeUTC   = 0
//     ServicePositionQuantity               = 0
//     PositionHasBeenUpdatedByService       = 0
//     PriceHighDuringPosition               = 0
//     PriceLowDuringPosition                = 0
//     PriceLastDuringPosition               = 0
//     LastProcessedTimeAndSalesSequence     = 0
//     TotalMarginRequirement                = 0.000000
//     InitialEntryDateTimeUTC               = 0
//     IsFromDTCServerReplay                 = false
//     MostRecentPositionIncreaseDateTimeUTC = 0
//     IsSnapshot                            = false
//     IsFirstMessageInBatch                 = false
//     IsLastMessageInBatch                  = false
//     MarginRequirementFull                 = 0
//     MarginRequirementFullPositionsOnly    = 0
//     MaxPotentialPositionQuantity          = 0
// }
var _TradePositionConsolidatedDefault = []byte{153, 0, 129, 39, 153, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradePositionConsolidated struct {
	message.VLS
}

type TradePositionConsolidatedBuilder struct {
	message.VLSBuilder
}

type TradePositionConsolidatedPointer struct {
	message.VLSPointer
}

type TradePositionConsolidatedPointerBuilder struct {
	message.VLSPointerBuilder
}

func NewTradePositionConsolidatedFrom(b []byte) TradePositionConsolidated {
	return TradePositionConsolidated{VLS: message.NewVLSFrom(b)}
}

func WrapTradePositionConsolidated(b []byte) TradePositionConsolidated {
	return TradePositionConsolidated{VLS: message.WrapVLS(b)}
}

func NewTradePositionConsolidated() TradePositionConsolidatedBuilder {
	a := TradePositionConsolidatedBuilder{message.NewVLSBuilder(153)}
	a.Unsafe().SetBytes(0, _TradePositionConsolidatedDefault)
	return a
}

func AllocTradePositionConsolidated() TradePositionConsolidatedPointerBuilder {
	a := TradePositionConsolidatedPointerBuilder{message.AllocVLSBuilder(153)}
	a.Ptr.SetBytes(0, _TradePositionConsolidatedDefault)
	return a
}

func AllocTradePositionConsolidatedFrom(b []byte) TradePositionConsolidatedPointer {
	return TradePositionConsolidatedPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                                  = TradePositionConsolidatedSize  (153)
//     Type                                  = TRADE_POSITION_CONSOLIDATED  (10113)
//     BaseSize                              = TradePositionConsolidatedSize  (153)
//     IsDeleted                             = false
//     Symbol                                = ""
//     IsSimulated                           = false
//     TradeAccount                          = ""
//     CurrencyCode                          = ""
//     Quantity                              = 0
//     AveragePrice                          = 0
//     OpenProfitLoss                        = 0
//     DailyProfitLoss                       = 0
//     LastDailyProfitLossResetDateTimeUTC   = 0
//     ServicePositionQuantity               = 0
//     PositionHasBeenUpdatedByService       = 0
//     PriceHighDuringPosition               = 0
//     PriceLowDuringPosition                = 0
//     PriceLastDuringPosition               = 0
//     LastProcessedTimeAndSalesSequence     = 0
//     TotalMarginRequirement                = 0.000000
//     InitialEntryDateTimeUTC               = 0
//     IsFromDTCServerReplay                 = false
//     MostRecentPositionIncreaseDateTimeUTC = 0
//     IsSnapshot                            = false
//     IsFirstMessageInBatch                 = false
//     IsLastMessageInBatch                  = false
//     MarginRequirementFull                 = 0
//     MarginRequirementFullPositionsOnly    = 0
//     MaxPotentialPositionQuantity          = 0
// }
func (m TradePositionConsolidatedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradePositionConsolidatedDefault)
}

// Clear
// {
//     Size                                  = TradePositionConsolidatedSize  (153)
//     Type                                  = TRADE_POSITION_CONSOLIDATED  (10113)
//     BaseSize                              = TradePositionConsolidatedSize  (153)
//     IsDeleted                             = false
//     Symbol                                = ""
//     IsSimulated                           = false
//     TradeAccount                          = ""
//     CurrencyCode                          = ""
//     Quantity                              = 0
//     AveragePrice                          = 0
//     OpenProfitLoss                        = 0
//     DailyProfitLoss                       = 0
//     LastDailyProfitLossResetDateTimeUTC   = 0
//     ServicePositionQuantity               = 0
//     PositionHasBeenUpdatedByService       = 0
//     PriceHighDuringPosition               = 0
//     PriceLowDuringPosition                = 0
//     PriceLastDuringPosition               = 0
//     LastProcessedTimeAndSalesSequence     = 0
//     TotalMarginRequirement                = 0.000000
//     InitialEntryDateTimeUTC               = 0
//     IsFromDTCServerReplay                 = false
//     MostRecentPositionIncreaseDateTimeUTC = 0
//     IsSnapshot                            = false
//     IsFirstMessageInBatch                 = false
//     IsLastMessageInBatch                  = false
//     MarginRequirementFull                 = 0
//     MarginRequirementFullPositionsOnly    = 0
//     MaxPotentialPositionQuantity          = 0
// }
func (m TradePositionConsolidatedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradePositionConsolidatedDefault)
}

// ToBuilder
func (m TradePositionConsolidated) ToBuilder() TradePositionConsolidatedBuilder {
	return TradePositionConsolidatedBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradePositionConsolidatedPointer) ToBuilder() TradePositionConsolidatedPointerBuilder {
	return TradePositionConsolidatedPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m TradePositionConsolidatedBuilder) Finish() TradePositionConsolidated {
	return TradePositionConsolidated{m.VLSBuilder.Finish()}
}

// Finish
func (m *TradePositionConsolidatedPointerBuilder) Finish() TradePositionConsolidatedPointer {
	return TradePositionConsolidatedPointer{m.VLSPointerBuilder.Finish()}
}

// IsDeleted
func (m TradePositionConsolidated) IsDeleted() bool {
	return message.BoolVLS(m.Unsafe(), 7, 6)
}

// IsDeleted
func (m TradePositionConsolidatedBuilder) IsDeleted() bool {
	return message.BoolVLS(m.Unsafe(), 7, 6)
}

// IsDeleted
func (m TradePositionConsolidatedPointer) IsDeleted() bool {
	return message.BoolVLS(m.Ptr, 7, 6)
}

// IsDeleted
func (m TradePositionConsolidatedPointerBuilder) IsDeleted() bool {
	return message.BoolVLS(m.Ptr, 7, 6)
}

// Symbol
func (m TradePositionConsolidated) Symbol() string {
	return message.StringVLS(m.Unsafe(), 11, 7)
}

// Symbol
func (m TradePositionConsolidatedBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 11, 7)
}

// Symbol
func (m TradePositionConsolidatedPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 11, 7)
}

// Symbol
func (m TradePositionConsolidatedPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 11, 7)
}

// IsSimulated
func (m TradePositionConsolidated) IsSimulated() bool {
	return message.BoolVLS(m.Unsafe(), 12, 11)
}

// IsSimulated
func (m TradePositionConsolidatedBuilder) IsSimulated() bool {
	return message.BoolVLS(m.Unsafe(), 12, 11)
}

// IsSimulated
func (m TradePositionConsolidatedPointer) IsSimulated() bool {
	return message.BoolVLS(m.Ptr, 12, 11)
}

// IsSimulated
func (m TradePositionConsolidatedPointerBuilder) IsSimulated() bool {
	return message.BoolVLS(m.Ptr, 12, 11)
}

// TradeAccount
func (m TradePositionConsolidated) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// TradeAccount
func (m TradePositionConsolidatedBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// TradeAccount
func (m TradePositionConsolidatedPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// TradeAccount
func (m TradePositionConsolidatedPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// CurrencyCode
func (m TradePositionConsolidated) CurrencyCode() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// CurrencyCode
func (m TradePositionConsolidatedBuilder) CurrencyCode() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// CurrencyCode
func (m TradePositionConsolidatedPointer) CurrencyCode() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// CurrencyCode
func (m TradePositionConsolidatedPointerBuilder) CurrencyCode() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Quantity
func (m TradePositionConsolidated) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 28, 20)
}

// Quantity
func (m TradePositionConsolidatedBuilder) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 28, 20)
}

// Quantity
func (m TradePositionConsolidatedPointer) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 28, 20)
}

// Quantity
func (m TradePositionConsolidatedPointerBuilder) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 28, 20)
}

// AveragePrice
func (m TradePositionConsolidated) AveragePrice() float64 {
	return message.Float64VLS(m.Unsafe(), 36, 28)
}

// AveragePrice
func (m TradePositionConsolidatedBuilder) AveragePrice() float64 {
	return message.Float64VLS(m.Unsafe(), 36, 28)
}

// AveragePrice
func (m TradePositionConsolidatedPointer) AveragePrice() float64 {
	return message.Float64VLS(m.Ptr, 36, 28)
}

// AveragePrice
func (m TradePositionConsolidatedPointerBuilder) AveragePrice() float64 {
	return message.Float64VLS(m.Ptr, 36, 28)
}

// OpenProfitLoss
func (m TradePositionConsolidated) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 44, 36)
}

// OpenProfitLoss
func (m TradePositionConsolidatedBuilder) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 44, 36)
}

// OpenProfitLoss
func (m TradePositionConsolidatedPointer) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 44, 36)
}

// OpenProfitLoss
func (m TradePositionConsolidatedPointerBuilder) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 44, 36)
}

// DailyProfitLoss
func (m TradePositionConsolidated) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 52, 44)
}

// DailyProfitLoss
func (m TradePositionConsolidatedBuilder) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 52, 44)
}

// DailyProfitLoss
func (m TradePositionConsolidatedPointer) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 52, 44)
}

// DailyProfitLoss
func (m TradePositionConsolidatedPointerBuilder) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 52, 44)
}

// LastDailyProfitLossResetDateTimeUTC
func (m TradePositionConsolidated) LastDailyProfitLossResetDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 60, 52)
}

// LastDailyProfitLossResetDateTimeUTC
func (m TradePositionConsolidatedBuilder) LastDailyProfitLossResetDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 60, 52)
}

// LastDailyProfitLossResetDateTimeUTC
func (m TradePositionConsolidatedPointer) LastDailyProfitLossResetDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 60, 52)
}

// LastDailyProfitLossResetDateTimeUTC
func (m TradePositionConsolidatedPointerBuilder) LastDailyProfitLossResetDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 60, 52)
}

// ServicePositionQuantity
func (m TradePositionConsolidated) ServicePositionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 68, 60)
}

// ServicePositionQuantity
func (m TradePositionConsolidatedBuilder) ServicePositionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 68, 60)
}

// ServicePositionQuantity
func (m TradePositionConsolidatedPointer) ServicePositionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 68, 60)
}

// ServicePositionQuantity
func (m TradePositionConsolidatedPointerBuilder) ServicePositionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 68, 60)
}

// PositionHasBeenUpdatedByService
func (m TradePositionConsolidated) PositionHasBeenUpdatedByService() uint8 {
	return message.Uint8VLS(m.Unsafe(), 69, 68)
}

// PositionHasBeenUpdatedByService
func (m TradePositionConsolidatedBuilder) PositionHasBeenUpdatedByService() uint8 {
	return message.Uint8VLS(m.Unsafe(), 69, 68)
}

// PositionHasBeenUpdatedByService
func (m TradePositionConsolidatedPointer) PositionHasBeenUpdatedByService() uint8 {
	return message.Uint8VLS(m.Ptr, 69, 68)
}

// PositionHasBeenUpdatedByService
func (m TradePositionConsolidatedPointerBuilder) PositionHasBeenUpdatedByService() uint8 {
	return message.Uint8VLS(m.Ptr, 69, 68)
}

// PriceHighDuringPosition
func (m TradePositionConsolidated) PriceHighDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 77, 69)
}

// PriceHighDuringPosition
func (m TradePositionConsolidatedBuilder) PriceHighDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 77, 69)
}

// PriceHighDuringPosition
func (m TradePositionConsolidatedPointer) PriceHighDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 77, 69)
}

// PriceHighDuringPosition
func (m TradePositionConsolidatedPointerBuilder) PriceHighDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 77, 69)
}

// PriceLowDuringPosition
func (m TradePositionConsolidated) PriceLowDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 85, 77)
}

// PriceLowDuringPosition
func (m TradePositionConsolidatedBuilder) PriceLowDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 85, 77)
}

// PriceLowDuringPosition
func (m TradePositionConsolidatedPointer) PriceLowDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 85, 77)
}

// PriceLowDuringPosition
func (m TradePositionConsolidatedPointerBuilder) PriceLowDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 85, 77)
}

// PriceLastDuringPosition
func (m TradePositionConsolidated) PriceLastDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 93, 85)
}

// PriceLastDuringPosition
func (m TradePositionConsolidatedBuilder) PriceLastDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 93, 85)
}

// PriceLastDuringPosition
func (m TradePositionConsolidatedPointer) PriceLastDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 93, 85)
}

// PriceLastDuringPosition
func (m TradePositionConsolidatedPointerBuilder) PriceLastDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 93, 85)
}

// LastProcessedTimeAndSalesSequence
func (m TradePositionConsolidated) LastProcessedTimeAndSalesSequence() int64 {
	return message.Int64VLS(m.Unsafe(), 101, 93)
}

// LastProcessedTimeAndSalesSequence
func (m TradePositionConsolidatedBuilder) LastProcessedTimeAndSalesSequence() int64 {
	return message.Int64VLS(m.Unsafe(), 101, 93)
}

// LastProcessedTimeAndSalesSequence
func (m TradePositionConsolidatedPointer) LastProcessedTimeAndSalesSequence() int64 {
	return message.Int64VLS(m.Ptr, 101, 93)
}

// LastProcessedTimeAndSalesSequence
func (m TradePositionConsolidatedPointerBuilder) LastProcessedTimeAndSalesSequence() int64 {
	return message.Int64VLS(m.Ptr, 101, 93)
}

// TotalMarginRequirement
func (m TradePositionConsolidated) TotalMarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 109, 101)
}

// TotalMarginRequirement
func (m TradePositionConsolidatedBuilder) TotalMarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 109, 101)
}

// TotalMarginRequirement
func (m TradePositionConsolidatedPointer) TotalMarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 109, 101)
}

// TotalMarginRequirement
func (m TradePositionConsolidatedPointerBuilder) TotalMarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 109, 101)
}

// InitialEntryDateTimeUTC
func (m TradePositionConsolidated) InitialEntryDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 117, 109)
}

// InitialEntryDateTimeUTC
func (m TradePositionConsolidatedBuilder) InitialEntryDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 117, 109)
}

// InitialEntryDateTimeUTC
func (m TradePositionConsolidatedPointer) InitialEntryDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 117, 109)
}

// InitialEntryDateTimeUTC
func (m TradePositionConsolidatedPointerBuilder) InitialEntryDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 117, 109)
}

// IsFromDTCServerReplay
func (m TradePositionConsolidated) IsFromDTCServerReplay() bool {
	return message.BoolVLS(m.Unsafe(), 118, 117)
}

// IsFromDTCServerReplay
func (m TradePositionConsolidatedBuilder) IsFromDTCServerReplay() bool {
	return message.BoolVLS(m.Unsafe(), 118, 117)
}

// IsFromDTCServerReplay
func (m TradePositionConsolidatedPointer) IsFromDTCServerReplay() bool {
	return message.BoolVLS(m.Ptr, 118, 117)
}

// IsFromDTCServerReplay
func (m TradePositionConsolidatedPointerBuilder) IsFromDTCServerReplay() bool {
	return message.BoolVLS(m.Ptr, 118, 117)
}

// MostRecentPositionIncreaseDateTimeUTC
func (m TradePositionConsolidated) MostRecentPositionIncreaseDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 126, 118)
}

// MostRecentPositionIncreaseDateTimeUTC
func (m TradePositionConsolidatedBuilder) MostRecentPositionIncreaseDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 126, 118)
}

// MostRecentPositionIncreaseDateTimeUTC
func (m TradePositionConsolidatedPointer) MostRecentPositionIncreaseDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 126, 118)
}

// MostRecentPositionIncreaseDateTimeUTC
func (m TradePositionConsolidatedPointerBuilder) MostRecentPositionIncreaseDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 126, 118)
}

// IsSnapshot
func (m TradePositionConsolidated) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 127, 126)
}

// IsSnapshot
func (m TradePositionConsolidatedBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 127, 126)
}

// IsSnapshot
func (m TradePositionConsolidatedPointer) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 127, 126)
}

// IsSnapshot
func (m TradePositionConsolidatedPointerBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 127, 126)
}

// IsFirstMessageInBatch
func (m TradePositionConsolidated) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 128, 127)
}

// IsFirstMessageInBatch
func (m TradePositionConsolidatedBuilder) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 128, 127)
}

// IsFirstMessageInBatch
func (m TradePositionConsolidatedPointer) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 128, 127)
}

// IsFirstMessageInBatch
func (m TradePositionConsolidatedPointerBuilder) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 128, 127)
}

// IsLastMessageInBatch
func (m TradePositionConsolidated) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 129, 128)
}

// IsLastMessageInBatch
func (m TradePositionConsolidatedBuilder) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 129, 128)
}

// IsLastMessageInBatch
func (m TradePositionConsolidatedPointer) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 129, 128)
}

// IsLastMessageInBatch
func (m TradePositionConsolidatedPointerBuilder) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 129, 128)
}

// MarginRequirementFull
func (m TradePositionConsolidated) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Unsafe(), 137, 129)
}

// MarginRequirementFull
func (m TradePositionConsolidatedBuilder) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Unsafe(), 137, 129)
}

// MarginRequirementFull
func (m TradePositionConsolidatedPointer) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Ptr, 137, 129)
}

// MarginRequirementFull
func (m TradePositionConsolidatedPointerBuilder) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Ptr, 137, 129)
}

// MarginRequirementFullPositionsOnly
func (m TradePositionConsolidated) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Unsafe(), 145, 137)
}

// MarginRequirementFullPositionsOnly
func (m TradePositionConsolidatedBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Unsafe(), 145, 137)
}

// MarginRequirementFullPositionsOnly
func (m TradePositionConsolidatedPointer) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Ptr, 145, 137)
}

// MarginRequirementFullPositionsOnly
func (m TradePositionConsolidatedPointerBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Ptr, 145, 137)
}

// MaxPotentialPositionQuantity
func (m TradePositionConsolidated) MaxPotentialPositionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 153, 145)
}

// MaxPotentialPositionQuantity
func (m TradePositionConsolidatedBuilder) MaxPotentialPositionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 153, 145)
}

// MaxPotentialPositionQuantity
func (m TradePositionConsolidatedPointer) MaxPotentialPositionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 153, 145)
}

// MaxPotentialPositionQuantity
func (m TradePositionConsolidatedPointerBuilder) MaxPotentialPositionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 153, 145)
}

// SetIsDeleted
func (m TradePositionConsolidatedBuilder) SetIsDeleted(value bool) {
	message.SetBoolVLS(m.Unsafe(), 7, 6, value)
}

// SetIsDeleted
func (m TradePositionConsolidatedPointerBuilder) SetIsDeleted(value bool) {
	message.SetBoolVLS(m.Ptr, 7, 6, value)
}

// SetSymbol
func (m *TradePositionConsolidatedBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 11, 7, value)
}

// SetSymbol
func (m *TradePositionConsolidatedPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 11, 7, value)
}

// SetIsSimulated
func (m TradePositionConsolidatedBuilder) SetIsSimulated(value bool) {
	message.SetBoolVLS(m.Unsafe(), 12, 11, value)
}

// SetIsSimulated
func (m TradePositionConsolidatedPointerBuilder) SetIsSimulated(value bool) {
	message.SetBoolVLS(m.Ptr, 12, 11, value)
}

// SetTradeAccount
func (m *TradePositionConsolidatedBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetTradeAccount
func (m *TradePositionConsolidatedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetCurrencyCode
func (m *TradePositionConsolidatedBuilder) SetCurrencyCode(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetCurrencyCode
func (m *TradePositionConsolidatedPointerBuilder) SetCurrencyCode(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetQuantity
func (m TradePositionConsolidatedBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 28, 20, value)
}

// SetQuantity
func (m TradePositionConsolidatedPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 28, 20, value)
}

// SetAveragePrice
func (m TradePositionConsolidatedBuilder) SetAveragePrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 36, 28, value)
}

// SetAveragePrice
func (m TradePositionConsolidatedPointerBuilder) SetAveragePrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 36, 28, value)
}

// SetOpenProfitLoss
func (m TradePositionConsolidatedBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 44, 36, value)
}

// SetOpenProfitLoss
func (m TradePositionConsolidatedPointerBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 44, 36, value)
}

// SetDailyProfitLoss
func (m TradePositionConsolidatedBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 52, 44, value)
}

// SetDailyProfitLoss
func (m TradePositionConsolidatedPointerBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 52, 44, value)
}

// SetLastDailyProfitLossResetDateTimeUTC
func (m TradePositionConsolidatedBuilder) SetLastDailyProfitLossResetDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Unsafe(), 60, 52, value)
}

// SetLastDailyProfitLossResetDateTimeUTC
func (m TradePositionConsolidatedPointerBuilder) SetLastDailyProfitLossResetDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 60, 52, value)
}

// SetServicePositionQuantity
func (m TradePositionConsolidatedBuilder) SetServicePositionQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 68, 60, value)
}

// SetServicePositionQuantity
func (m TradePositionConsolidatedPointerBuilder) SetServicePositionQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 68, 60, value)
}

// SetPositionHasBeenUpdatedByService
func (m TradePositionConsolidatedBuilder) SetPositionHasBeenUpdatedByService(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 69, 68, value)
}

// SetPositionHasBeenUpdatedByService
func (m TradePositionConsolidatedPointerBuilder) SetPositionHasBeenUpdatedByService(value uint8) {
	message.SetUint8VLS(m.Ptr, 69, 68, value)
}

// SetPriceHighDuringPosition
func (m TradePositionConsolidatedBuilder) SetPriceHighDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 77, 69, value)
}

// SetPriceHighDuringPosition
func (m TradePositionConsolidatedPointerBuilder) SetPriceHighDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 77, 69, value)
}

// SetPriceLowDuringPosition
func (m TradePositionConsolidatedBuilder) SetPriceLowDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 85, 77, value)
}

// SetPriceLowDuringPosition
func (m TradePositionConsolidatedPointerBuilder) SetPriceLowDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 85, 77, value)
}

// SetPriceLastDuringPosition
func (m TradePositionConsolidatedBuilder) SetPriceLastDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 93, 85, value)
}

// SetPriceLastDuringPosition
func (m TradePositionConsolidatedPointerBuilder) SetPriceLastDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 93, 85, value)
}

// SetLastProcessedTimeAndSalesSequence
func (m TradePositionConsolidatedBuilder) SetLastProcessedTimeAndSalesSequence(value int64) {
	message.SetInt64VLS(m.Unsafe(), 101, 93, value)
}

// SetLastProcessedTimeAndSalesSequence
func (m TradePositionConsolidatedPointerBuilder) SetLastProcessedTimeAndSalesSequence(value int64) {
	message.SetInt64VLS(m.Ptr, 101, 93, value)
}

// SetTotalMarginRequirement
func (m TradePositionConsolidatedBuilder) SetTotalMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 109, 101, value)
}

// SetTotalMarginRequirement
func (m TradePositionConsolidatedPointerBuilder) SetTotalMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Ptr, 109, 101, value)
}

// SetInitialEntryDateTimeUTC
func (m TradePositionConsolidatedBuilder) SetInitialEntryDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Unsafe(), 117, 109, value)
}

// SetInitialEntryDateTimeUTC
func (m TradePositionConsolidatedPointerBuilder) SetInitialEntryDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 117, 109, value)
}

// SetIsFromDTCServerReplay
func (m TradePositionConsolidatedBuilder) SetIsFromDTCServerReplay(value bool) {
	message.SetBoolVLS(m.Unsafe(), 118, 117, value)
}

// SetIsFromDTCServerReplay
func (m TradePositionConsolidatedPointerBuilder) SetIsFromDTCServerReplay(value bool) {
	message.SetBoolVLS(m.Ptr, 118, 117, value)
}

// SetMostRecentPositionIncreaseDateTimeUTC
func (m TradePositionConsolidatedBuilder) SetMostRecentPositionIncreaseDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Unsafe(), 126, 118, value)
}

// SetMostRecentPositionIncreaseDateTimeUTC
func (m TradePositionConsolidatedPointerBuilder) SetMostRecentPositionIncreaseDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 126, 118, value)
}

// SetIsSnapshot
func (m TradePositionConsolidatedBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Unsafe(), 127, 126, value)
}

// SetIsSnapshot
func (m TradePositionConsolidatedPointerBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Ptr, 127, 126, value)
}

// SetIsFirstMessageInBatch
func (m TradePositionConsolidatedBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolVLS(m.Unsafe(), 128, 127, value)
}

// SetIsFirstMessageInBatch
func (m TradePositionConsolidatedPointerBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolVLS(m.Ptr, 128, 127, value)
}

// SetIsLastMessageInBatch
func (m TradePositionConsolidatedBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolVLS(m.Unsafe(), 129, 128, value)
}

// SetIsLastMessageInBatch
func (m TradePositionConsolidatedPointerBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolVLS(m.Ptr, 129, 128, value)
}

// SetMarginRequirementFull
func (m TradePositionConsolidatedBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 137, 129, value)
}

// SetMarginRequirementFull
func (m TradePositionConsolidatedPointerBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64VLS(m.Ptr, 137, 129, value)
}

// SetMarginRequirementFullPositionsOnly
func (m TradePositionConsolidatedBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 145, 137, value)
}

// SetMarginRequirementFullPositionsOnly
func (m TradePositionConsolidatedPointerBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64VLS(m.Ptr, 145, 137, value)
}

// SetMaxPotentialPositionQuantity
func (m TradePositionConsolidatedBuilder) SetMaxPotentialPositionQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 153, 145, value)
}

// SetMaxPotentialPositionQuantity
func (m TradePositionConsolidatedPointerBuilder) SetMaxPotentialPositionQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 153, 145, value)
}

// Copy
func (m TradePositionConsolidated) Copy(to TradePositionConsolidatedBuilder) {
	to.SetIsDeleted(m.IsDeleted())
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetLastDailyProfitLossResetDateTimeUTC(m.LastDailyProfitLossResetDateTimeUTC())
	to.SetServicePositionQuantity(m.ServicePositionQuantity())
	to.SetPositionHasBeenUpdatedByService(m.PositionHasBeenUpdatedByService())
	to.SetPriceHighDuringPosition(m.PriceHighDuringPosition())
	to.SetPriceLowDuringPosition(m.PriceLowDuringPosition())
	to.SetPriceLastDuringPosition(m.PriceLastDuringPosition())
	to.SetLastProcessedTimeAndSalesSequence(m.LastProcessedTimeAndSalesSequence())
	to.SetTotalMarginRequirement(m.TotalMarginRequirement())
	to.SetInitialEntryDateTimeUTC(m.InitialEntryDateTimeUTC())
	to.SetIsFromDTCServerReplay(m.IsFromDTCServerReplay())
	to.SetMostRecentPositionIncreaseDateTimeUTC(m.MostRecentPositionIncreaseDateTimeUTC())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetMaxPotentialPositionQuantity(m.MaxPotentialPositionQuantity())
}

// Copy
func (m TradePositionConsolidatedBuilder) Copy(to TradePositionConsolidatedBuilder) {
	to.SetIsDeleted(m.IsDeleted())
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetLastDailyProfitLossResetDateTimeUTC(m.LastDailyProfitLossResetDateTimeUTC())
	to.SetServicePositionQuantity(m.ServicePositionQuantity())
	to.SetPositionHasBeenUpdatedByService(m.PositionHasBeenUpdatedByService())
	to.SetPriceHighDuringPosition(m.PriceHighDuringPosition())
	to.SetPriceLowDuringPosition(m.PriceLowDuringPosition())
	to.SetPriceLastDuringPosition(m.PriceLastDuringPosition())
	to.SetLastProcessedTimeAndSalesSequence(m.LastProcessedTimeAndSalesSequence())
	to.SetTotalMarginRequirement(m.TotalMarginRequirement())
	to.SetInitialEntryDateTimeUTC(m.InitialEntryDateTimeUTC())
	to.SetIsFromDTCServerReplay(m.IsFromDTCServerReplay())
	to.SetMostRecentPositionIncreaseDateTimeUTC(m.MostRecentPositionIncreaseDateTimeUTC())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetMaxPotentialPositionQuantity(m.MaxPotentialPositionQuantity())
}

// CopyPointer
func (m TradePositionConsolidated) CopyPointer(to TradePositionConsolidatedPointerBuilder) {
	to.SetIsDeleted(m.IsDeleted())
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetLastDailyProfitLossResetDateTimeUTC(m.LastDailyProfitLossResetDateTimeUTC())
	to.SetServicePositionQuantity(m.ServicePositionQuantity())
	to.SetPositionHasBeenUpdatedByService(m.PositionHasBeenUpdatedByService())
	to.SetPriceHighDuringPosition(m.PriceHighDuringPosition())
	to.SetPriceLowDuringPosition(m.PriceLowDuringPosition())
	to.SetPriceLastDuringPosition(m.PriceLastDuringPosition())
	to.SetLastProcessedTimeAndSalesSequence(m.LastProcessedTimeAndSalesSequence())
	to.SetTotalMarginRequirement(m.TotalMarginRequirement())
	to.SetInitialEntryDateTimeUTC(m.InitialEntryDateTimeUTC())
	to.SetIsFromDTCServerReplay(m.IsFromDTCServerReplay())
	to.SetMostRecentPositionIncreaseDateTimeUTC(m.MostRecentPositionIncreaseDateTimeUTC())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetMaxPotentialPositionQuantity(m.MaxPotentialPositionQuantity())
}

// CopyPointer
func (m TradePositionConsolidatedBuilder) CopyPointer(to TradePositionConsolidatedPointerBuilder) {
	to.SetIsDeleted(m.IsDeleted())
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetLastDailyProfitLossResetDateTimeUTC(m.LastDailyProfitLossResetDateTimeUTC())
	to.SetServicePositionQuantity(m.ServicePositionQuantity())
	to.SetPositionHasBeenUpdatedByService(m.PositionHasBeenUpdatedByService())
	to.SetPriceHighDuringPosition(m.PriceHighDuringPosition())
	to.SetPriceLowDuringPosition(m.PriceLowDuringPosition())
	to.SetPriceLastDuringPosition(m.PriceLastDuringPosition())
	to.SetLastProcessedTimeAndSalesSequence(m.LastProcessedTimeAndSalesSequence())
	to.SetTotalMarginRequirement(m.TotalMarginRequirement())
	to.SetInitialEntryDateTimeUTC(m.InitialEntryDateTimeUTC())
	to.SetIsFromDTCServerReplay(m.IsFromDTCServerReplay())
	to.SetMostRecentPositionIncreaseDateTimeUTC(m.MostRecentPositionIncreaseDateTimeUTC())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetMaxPotentialPositionQuantity(m.MaxPotentialPositionQuantity())
}

// Copy
func (m TradePositionConsolidatedPointer) Copy(to TradePositionConsolidatedBuilder) {
	to.SetIsDeleted(m.IsDeleted())
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetLastDailyProfitLossResetDateTimeUTC(m.LastDailyProfitLossResetDateTimeUTC())
	to.SetServicePositionQuantity(m.ServicePositionQuantity())
	to.SetPositionHasBeenUpdatedByService(m.PositionHasBeenUpdatedByService())
	to.SetPriceHighDuringPosition(m.PriceHighDuringPosition())
	to.SetPriceLowDuringPosition(m.PriceLowDuringPosition())
	to.SetPriceLastDuringPosition(m.PriceLastDuringPosition())
	to.SetLastProcessedTimeAndSalesSequence(m.LastProcessedTimeAndSalesSequence())
	to.SetTotalMarginRequirement(m.TotalMarginRequirement())
	to.SetInitialEntryDateTimeUTC(m.InitialEntryDateTimeUTC())
	to.SetIsFromDTCServerReplay(m.IsFromDTCServerReplay())
	to.SetMostRecentPositionIncreaseDateTimeUTC(m.MostRecentPositionIncreaseDateTimeUTC())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetMaxPotentialPositionQuantity(m.MaxPotentialPositionQuantity())
}

// Copy
func (m TradePositionConsolidatedPointerBuilder) Copy(to TradePositionConsolidatedBuilder) {
	to.SetIsDeleted(m.IsDeleted())
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetLastDailyProfitLossResetDateTimeUTC(m.LastDailyProfitLossResetDateTimeUTC())
	to.SetServicePositionQuantity(m.ServicePositionQuantity())
	to.SetPositionHasBeenUpdatedByService(m.PositionHasBeenUpdatedByService())
	to.SetPriceHighDuringPosition(m.PriceHighDuringPosition())
	to.SetPriceLowDuringPosition(m.PriceLowDuringPosition())
	to.SetPriceLastDuringPosition(m.PriceLastDuringPosition())
	to.SetLastProcessedTimeAndSalesSequence(m.LastProcessedTimeAndSalesSequence())
	to.SetTotalMarginRequirement(m.TotalMarginRequirement())
	to.SetInitialEntryDateTimeUTC(m.InitialEntryDateTimeUTC())
	to.SetIsFromDTCServerReplay(m.IsFromDTCServerReplay())
	to.SetMostRecentPositionIncreaseDateTimeUTC(m.MostRecentPositionIncreaseDateTimeUTC())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetMaxPotentialPositionQuantity(m.MaxPotentialPositionQuantity())
}

// CopyPointer
func (m TradePositionConsolidatedPointer) CopyPointer(to TradePositionConsolidatedPointerBuilder) {
	to.SetIsDeleted(m.IsDeleted())
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetLastDailyProfitLossResetDateTimeUTC(m.LastDailyProfitLossResetDateTimeUTC())
	to.SetServicePositionQuantity(m.ServicePositionQuantity())
	to.SetPositionHasBeenUpdatedByService(m.PositionHasBeenUpdatedByService())
	to.SetPriceHighDuringPosition(m.PriceHighDuringPosition())
	to.SetPriceLowDuringPosition(m.PriceLowDuringPosition())
	to.SetPriceLastDuringPosition(m.PriceLastDuringPosition())
	to.SetLastProcessedTimeAndSalesSequence(m.LastProcessedTimeAndSalesSequence())
	to.SetTotalMarginRequirement(m.TotalMarginRequirement())
	to.SetInitialEntryDateTimeUTC(m.InitialEntryDateTimeUTC())
	to.SetIsFromDTCServerReplay(m.IsFromDTCServerReplay())
	to.SetMostRecentPositionIncreaseDateTimeUTC(m.MostRecentPositionIncreaseDateTimeUTC())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetMaxPotentialPositionQuantity(m.MaxPotentialPositionQuantity())
}

// CopyPointer
func (m TradePositionConsolidatedPointerBuilder) CopyPointer(to TradePositionConsolidatedPointerBuilder) {
	to.SetIsDeleted(m.IsDeleted())
	to.SetSymbol(m.Symbol())
	to.SetIsSimulated(m.IsSimulated())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetLastDailyProfitLossResetDateTimeUTC(m.LastDailyProfitLossResetDateTimeUTC())
	to.SetServicePositionQuantity(m.ServicePositionQuantity())
	to.SetPositionHasBeenUpdatedByService(m.PositionHasBeenUpdatedByService())
	to.SetPriceHighDuringPosition(m.PriceHighDuringPosition())
	to.SetPriceLowDuringPosition(m.PriceLowDuringPosition())
	to.SetPriceLastDuringPosition(m.PriceLastDuringPosition())
	to.SetLastProcessedTimeAndSalesSequence(m.LastProcessedTimeAndSalesSequence())
	to.SetTotalMarginRequirement(m.TotalMarginRequirement())
	to.SetInitialEntryDateTimeUTC(m.InitialEntryDateTimeUTC())
	to.SetIsFromDTCServerReplay(m.IsFromDTCServerReplay())
	to.SetMostRecentPositionIncreaseDateTimeUTC(m.MostRecentPositionIncreaseDateTimeUTC())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetMaxPotentialPositionQuantity(m.MaxPotentialPositionQuantity())
}
