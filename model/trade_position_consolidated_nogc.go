package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradePositionConsolidatedPointer struct {
	message.VLSPointer
}

type TradePositionConsolidatedPointerBuilder struct {
	message.VLSPointerBuilder
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
//     IsDeleted                             = 0
//     Symbol                                = ""
//     IsSimulated                           = 0
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
//     IsFromDTCServerReplay                 = 0
//     MostRecentPositionIncreaseDateTimeUTC = 0
//     IsSnapshot                            = 0
//     IsFirstMessageInBatch                 = 0
//     IsLastMessageInBatch                  = 0
//     MarginRequirementFull                 = 0
//     MarginRequirementFullPositionsOnly    = 0
//     MaxPotentialPositionQuantity          = 0
// }
func (m TradePositionConsolidatedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradePositionConsolidatedDefault)
}

// ToBuilder
func (m TradePositionConsolidatedPointer) ToBuilder() TradePositionConsolidatedPointerBuilder {
	return TradePositionConsolidatedPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradePositionConsolidatedPointerBuilder) Finish() TradePositionConsolidatedPointer {
	return TradePositionConsolidatedPointer{m.VLSPointerBuilder.Finish()}
}

// IsDeleted
func (m TradePositionConsolidatedPointer) IsDeleted() uint8 {
	return message.Uint8VLS(m.Ptr, 7, 6)
}

// IsDeleted
func (m TradePositionConsolidatedPointerBuilder) IsDeleted() uint8 {
	return message.Uint8VLS(m.Ptr, 7, 6)
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
func (m TradePositionConsolidatedPointer) IsSimulated() uint8 {
	return message.Uint8VLS(m.Ptr, 12, 11)
}

// IsSimulated
func (m TradePositionConsolidatedPointerBuilder) IsSimulated() uint8 {
	return message.Uint8VLS(m.Ptr, 12, 11)
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
func (m TradePositionConsolidatedPointer) CurrencyCode() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// CurrencyCode
func (m TradePositionConsolidatedPointerBuilder) CurrencyCode() string {
	return message.StringVLS(m.Ptr, 20, 16)
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
func (m TradePositionConsolidatedPointer) AveragePrice() float64 {
	return message.Float64VLS(m.Ptr, 36, 28)
}

// AveragePrice
func (m TradePositionConsolidatedPointerBuilder) AveragePrice() float64 {
	return message.Float64VLS(m.Ptr, 36, 28)
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
func (m TradePositionConsolidatedPointer) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 52, 44)
}

// DailyProfitLoss
func (m TradePositionConsolidatedPointerBuilder) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 52, 44)
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
func (m TradePositionConsolidatedPointer) ServicePositionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 68, 60)
}

// ServicePositionQuantity
func (m TradePositionConsolidatedPointerBuilder) ServicePositionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 68, 60)
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
func (m TradePositionConsolidatedPointer) PriceHighDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 77, 69)
}

// PriceHighDuringPosition
func (m TradePositionConsolidatedPointerBuilder) PriceHighDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 77, 69)
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
func (m TradePositionConsolidatedPointer) PriceLastDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 93, 85)
}

// PriceLastDuringPosition
func (m TradePositionConsolidatedPointerBuilder) PriceLastDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 93, 85)
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
func (m TradePositionConsolidatedPointer) TotalMarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 109, 101)
}

// TotalMarginRequirement
func (m TradePositionConsolidatedPointerBuilder) TotalMarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 109, 101)
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
func (m TradePositionConsolidatedPointer) IsFromDTCServerReplay() uint8 {
	return message.Uint8VLS(m.Ptr, 118, 117)
}

// IsFromDTCServerReplay
func (m TradePositionConsolidatedPointerBuilder) IsFromDTCServerReplay() uint8 {
	return message.Uint8VLS(m.Ptr, 118, 117)
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
func (m TradePositionConsolidatedPointer) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Ptr, 127, 126)
}

// IsSnapshot
func (m TradePositionConsolidatedPointerBuilder) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Ptr, 127, 126)
}

// IsFirstMessageInBatch
func (m TradePositionConsolidatedPointer) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 128, 127)
}

// IsFirstMessageInBatch
func (m TradePositionConsolidatedPointerBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 128, 127)
}

// IsLastMessageInBatch
func (m TradePositionConsolidatedPointer) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 129, 128)
}

// IsLastMessageInBatch
func (m TradePositionConsolidatedPointerBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 129, 128)
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
func (m TradePositionConsolidatedPointer) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Ptr, 145, 137)
}

// MarginRequirementFullPositionsOnly
func (m TradePositionConsolidatedPointerBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Ptr, 145, 137)
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
func (m TradePositionConsolidatedPointerBuilder) SetIsDeleted(value uint8) {
	message.SetUint8VLS(m.Ptr, 7, 6, value)
}

// SetSymbol
func (m *TradePositionConsolidatedPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 11, 7, value)
}

// SetIsSimulated
func (m TradePositionConsolidatedPointerBuilder) SetIsSimulated(value uint8) {
	message.SetUint8VLS(m.Ptr, 12, 11, value)
}

// SetTradeAccount
func (m *TradePositionConsolidatedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetCurrencyCode
func (m *TradePositionConsolidatedPointerBuilder) SetCurrencyCode(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetQuantity
func (m TradePositionConsolidatedPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 28, 20, value)
}

// SetAveragePrice
func (m TradePositionConsolidatedPointerBuilder) SetAveragePrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 36, 28, value)
}

// SetOpenProfitLoss
func (m TradePositionConsolidatedPointerBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 44, 36, value)
}

// SetDailyProfitLoss
func (m TradePositionConsolidatedPointerBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 52, 44, value)
}

// SetLastDailyProfitLossResetDateTimeUTC
func (m TradePositionConsolidatedPointerBuilder) SetLastDailyProfitLossResetDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 60, 52, value)
}

// SetServicePositionQuantity
func (m TradePositionConsolidatedPointerBuilder) SetServicePositionQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 68, 60, value)
}

// SetPositionHasBeenUpdatedByService
func (m TradePositionConsolidatedPointerBuilder) SetPositionHasBeenUpdatedByService(value uint8) {
	message.SetUint8VLS(m.Ptr, 69, 68, value)
}

// SetPriceHighDuringPosition
func (m TradePositionConsolidatedPointerBuilder) SetPriceHighDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 77, 69, value)
}

// SetPriceLowDuringPosition
func (m TradePositionConsolidatedPointerBuilder) SetPriceLowDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 85, 77, value)
}

// SetPriceLastDuringPosition
func (m TradePositionConsolidatedPointerBuilder) SetPriceLastDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 93, 85, value)
}

// SetLastProcessedTimeAndSalesSequence
func (m TradePositionConsolidatedPointerBuilder) SetLastProcessedTimeAndSalesSequence(value int64) {
	message.SetInt64VLS(m.Ptr, 101, 93, value)
}

// SetTotalMarginRequirement
func (m TradePositionConsolidatedPointerBuilder) SetTotalMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Ptr, 109, 101, value)
}

// SetInitialEntryDateTimeUTC
func (m TradePositionConsolidatedPointerBuilder) SetInitialEntryDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 117, 109, value)
}

// SetIsFromDTCServerReplay
func (m TradePositionConsolidatedPointerBuilder) SetIsFromDTCServerReplay(value uint8) {
	message.SetUint8VLS(m.Ptr, 118, 117, value)
}

// SetMostRecentPositionIncreaseDateTimeUTC
func (m TradePositionConsolidatedPointerBuilder) SetMostRecentPositionIncreaseDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 126, 118, value)
}

// SetIsSnapshot
func (m TradePositionConsolidatedPointerBuilder) SetIsSnapshot(value uint8) {
	message.SetUint8VLS(m.Ptr, 127, 126, value)
}

// SetIsFirstMessageInBatch
func (m TradePositionConsolidatedPointerBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Ptr, 128, 127, value)
}

// SetIsLastMessageInBatch
func (m TradePositionConsolidatedPointerBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Ptr, 129, 128, value)
}

// SetMarginRequirementFull
func (m TradePositionConsolidatedPointerBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64VLS(m.Ptr, 137, 129, value)
}

// SetMarginRequirementFullPositionsOnly
func (m TradePositionConsolidatedPointerBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64VLS(m.Ptr, 145, 137, value)
}

// SetMaxPotentialPositionQuantity
func (m TradePositionConsolidatedPointerBuilder) SetMaxPotentialPositionQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 153, 145, value)
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
