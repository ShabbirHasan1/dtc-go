package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                                                               = TradeAccountDataResponseSize  (311)
//     Type                                                               = TRADE_ACCOUNT_DATA_RESPONSE  (10116)
//     BaseSize                                                           = TradeAccountDataResponseSize  (311)
//     RequestID                                                          = 0
//     TradeAccountNotExist                                               = 0
//     TradeAccount                                                       = ""
//     IsSimulated                                                        = 0
//     CurrencyCode                                                       = ""
//     CurrentCashBalance                                                 = 0
//     AvailableFundsForNewPositions                                      = 0
//     MarginRequirement                                                  = 0
//     AccountValue                                                       = 0
//     OpenPositionsProfitLoss                                            = 0
//     DailyProfitLoss                                                    = 0
//     TransactionIdentifierForCashBalanceAdjustment                      = 0
//     LastTransactionDateTime                                            = 0
//     TrailingAccountValueAtWhichToNotAllowNewPositions                  = 0
//     CalculatedDailyNetLossLimitInAccountCurrency                       = 0
//     DailyNetLossLimitHasBeenReached                                    = 0
//     LastResetDailyNetLossManagementVariablesDateTimeUTC                = 0
//     IsUnderRequiredMargin                                              = 0
//     DailyNetLossLimitInAccountCurrency                                 = 0
//     PercentOfCashBalanceForDailyNetLossLimit                           = 0
//     UseTrailingAccountValueToNotAllowIncreaseInPositions               = 0
//     DoNotAllowIncreaseInPositionsAtDailyLossLimit                      = 0
//     FlattenPositionsAtDailyLossLimit                                   = 0
//     ClosePositionsAtEndOfDay                                           = 0
//     FlattenPositionsWhenUnderMarginIntraday                            = 1
//     FlattenPositionsWhenUnderMarginAtEndOfDay                          = 0
//     SenderSubID                                                        = ""
//     SenderLocationId                                                   = ""
//     SelfMatchPreventionID                                              = ""
//     CTICode                                                            = 0
//     TradeAccountIsReadOnly                                             = 0
//     MaximumGlobalPositionQuantity                                      = 0
//     TradeFeePerContract                                                = 0
//     TradeFeePerShare                                                   = 0
//     RequireSufficientMarginForNewPositions                             = 1
//     UsePercentOfMargin                                                 = 100
//     UsePercentOfMarginForDayTrading                                    = 100
//     MaximumAllowedAccountBalanceForPositionsAsPercentage               = 100
//     FirmID                                                             = ""
//     TradingIsDisabled                                                  = 0
//     DescriptiveName                                                    = ""
//     IsMasterFirmControlAccount                                         = 0
//     MinimumRequiredAccountValue                                        = 0
//     BeginTimeForDayMargin                                              = 0
//     EndTimeForDayMargin                                                = 0
//     DayMarginTimeZone                                                  = ""
//     IsSnapshot                                                         = 0
//     IsFirstMessageInBatch                                              = 0
//     IsLastMessageInBatch                                               = 0
//     IsDeleted                                                          = 0
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday              = 0
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay            = 0
//     UseMasterFirm_SymbolLimitsArray                                    = 0
//     UseMasterFirm_TradeFees                                            = 0
//     UseMasterFirm_TradeFeePerShare                                     = 0
//     UseMasterFirm_RequireSufficientMarginForNewPositions               = 0
//     UseMasterFirm_UsePercentOfMargin                                   = 0
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage = 0
//     UseMasterFirm_MinimumRequiredAccountValue                          = 0
//     UseMasterFirm_MarginTimeSettings                                   = 0
//     UseMasterFirm_TradingIsDisabled                                    = 0
//     IsTradeStatisticsPublicallyShared                                  = 0
//     IsReadOnlyFollowingRequestsAllowed                                 = 0
//     IsTradeAccountSharingAllowed                                       = 0
//     ReadOnlyShareToAllSCUsernames                                      = 0
//     UseMasterFirm_SymbolCommissionsArray                               = 0
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit        = 0
//     UseMasterFirm_UsePercentOfMarginForDayTrading                      = 0
//     OpenPositionsProfitLossBasedOnSettlement                           = 0
//     MarginRequirementFull                                              = 0
//     MarginRequirementFullPositionsOnly                                 = 0
//     UseMasterFirm_TradeFeesFullOverride                                = 0
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders         = 0
//     UseMasterFirm_UsePercentOfMarginFullOverride                       = 0
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride          = 0
//     PeakMarginRequirement                                              = 0
//     LiquidationOnlyMode                                                = 0
//     FlattenPositionsWhenUnderMarginIntradayTriggered                   = 0
//     FlattenPositionsWhenUnderMinimumAccountValueTriggered              = 0
//     AccountValueAtEndOfDayCaptureTime                                  = 0
//     EndOfDayCaptureTime                                                = 0
//     CustomerOrFirm                                                     = 0
//     MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet               = 0
//     MasterFirm_FlattenCancelWhenUnderMinimumAccountValue               = 0
//     MasterFirm_FlattenCancelWhenUnderMarginIntraday                    = 0
//     MasterFirm_FlattenCancelWhenUnderMarginEndOfDay                    = 0
//     MasterFirm_MaximumOrderQuantity                                    = 0
//     LastTriggerDateTimeUTCForDailyLossLimit                            = 0
//     OpenPositionsProfitLossIsDelayed                                   = 0
//     ExchangeTraderId                                                   = ""
// }
var _TradeAccountDataResponseDefault = []byte{55, 1, 132, 39, 55, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 100, 0, 0, 0, 100, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataResponseSize = 311

type TradeAccountDataResponse struct {
	message.VLS
}

type TradeAccountDataResponseBuilder struct {
	message.VLSBuilder
}

func NewTradeAccountDataResponseFrom(b []byte) TradeAccountDataResponse {
	return TradeAccountDataResponse{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataResponse(b []byte) TradeAccountDataResponse {
	return TradeAccountDataResponse{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataResponse() TradeAccountDataResponseBuilder {
	a := TradeAccountDataResponseBuilder{message.NewVLSBuilder(311)}
	a.Unsafe().SetBytes(0, _TradeAccountDataResponseDefault)
	return a
}

// Clear
// {
//     Size                                                               = TradeAccountDataResponseSize  (311)
//     Type                                                               = TRADE_ACCOUNT_DATA_RESPONSE  (10116)
//     BaseSize                                                           = TradeAccountDataResponseSize  (311)
//     RequestID                                                          = 0
//     TradeAccountNotExist                                               = 0
//     TradeAccount                                                       = ""
//     IsSimulated                                                        = 0
//     CurrencyCode                                                       = ""
//     CurrentCashBalance                                                 = 0
//     AvailableFundsForNewPositions                                      = 0
//     MarginRequirement                                                  = 0
//     AccountValue                                                       = 0
//     OpenPositionsProfitLoss                                            = 0
//     DailyProfitLoss                                                    = 0
//     TransactionIdentifierForCashBalanceAdjustment                      = 0
//     LastTransactionDateTime                                            = 0
//     TrailingAccountValueAtWhichToNotAllowNewPositions                  = 0
//     CalculatedDailyNetLossLimitInAccountCurrency                       = 0
//     DailyNetLossLimitHasBeenReached                                    = 0
//     LastResetDailyNetLossManagementVariablesDateTimeUTC                = 0
//     IsUnderRequiredMargin                                              = 0
//     DailyNetLossLimitInAccountCurrency                                 = 0
//     PercentOfCashBalanceForDailyNetLossLimit                           = 0
//     UseTrailingAccountValueToNotAllowIncreaseInPositions               = 0
//     DoNotAllowIncreaseInPositionsAtDailyLossLimit                      = 0
//     FlattenPositionsAtDailyLossLimit                                   = 0
//     ClosePositionsAtEndOfDay                                           = 0
//     FlattenPositionsWhenUnderMarginIntraday                            = 1
//     FlattenPositionsWhenUnderMarginAtEndOfDay                          = 0
//     SenderSubID                                                        = ""
//     SenderLocationId                                                   = ""
//     SelfMatchPreventionID                                              = ""
//     CTICode                                                            = 0
//     TradeAccountIsReadOnly                                             = 0
//     MaximumGlobalPositionQuantity                                      = 0
//     TradeFeePerContract                                                = 0
//     TradeFeePerShare                                                   = 0
//     RequireSufficientMarginForNewPositions                             = 1
//     UsePercentOfMargin                                                 = 100
//     UsePercentOfMarginForDayTrading                                    = 100
//     MaximumAllowedAccountBalanceForPositionsAsPercentage               = 100
//     FirmID                                                             = ""
//     TradingIsDisabled                                                  = 0
//     DescriptiveName                                                    = ""
//     IsMasterFirmControlAccount                                         = 0
//     MinimumRequiredAccountValue                                        = 0
//     BeginTimeForDayMargin                                              = 0
//     EndTimeForDayMargin                                                = 0
//     DayMarginTimeZone                                                  = ""
//     IsSnapshot                                                         = 0
//     IsFirstMessageInBatch                                              = 0
//     IsLastMessageInBatch                                               = 0
//     IsDeleted                                                          = 0
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday              = 0
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay            = 0
//     UseMasterFirm_SymbolLimitsArray                                    = 0
//     UseMasterFirm_TradeFees                                            = 0
//     UseMasterFirm_TradeFeePerShare                                     = 0
//     UseMasterFirm_RequireSufficientMarginForNewPositions               = 0
//     UseMasterFirm_UsePercentOfMargin                                   = 0
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage = 0
//     UseMasterFirm_MinimumRequiredAccountValue                          = 0
//     UseMasterFirm_MarginTimeSettings                                   = 0
//     UseMasterFirm_TradingIsDisabled                                    = 0
//     IsTradeStatisticsPublicallyShared                                  = 0
//     IsReadOnlyFollowingRequestsAllowed                                 = 0
//     IsTradeAccountSharingAllowed                                       = 0
//     ReadOnlyShareToAllSCUsernames                                      = 0
//     UseMasterFirm_SymbolCommissionsArray                               = 0
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit        = 0
//     UseMasterFirm_UsePercentOfMarginForDayTrading                      = 0
//     OpenPositionsProfitLossBasedOnSettlement                           = 0
//     MarginRequirementFull                                              = 0
//     MarginRequirementFullPositionsOnly                                 = 0
//     UseMasterFirm_TradeFeesFullOverride                                = 0
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders         = 0
//     UseMasterFirm_UsePercentOfMarginFullOverride                       = 0
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride          = 0
//     PeakMarginRequirement                                              = 0
//     LiquidationOnlyMode                                                = 0
//     FlattenPositionsWhenUnderMarginIntradayTriggered                   = 0
//     FlattenPositionsWhenUnderMinimumAccountValueTriggered              = 0
//     AccountValueAtEndOfDayCaptureTime                                  = 0
//     EndOfDayCaptureTime                                                = 0
//     CustomerOrFirm                                                     = 0
//     MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet               = 0
//     MasterFirm_FlattenCancelWhenUnderMinimumAccountValue               = 0
//     MasterFirm_FlattenCancelWhenUnderMarginIntraday                    = 0
//     MasterFirm_FlattenCancelWhenUnderMarginEndOfDay                    = 0
//     MasterFirm_MaximumOrderQuantity                                    = 0
//     LastTriggerDateTimeUTCForDailyLossLimit                            = 0
//     OpenPositionsProfitLossIsDelayed                                   = 0
//     ExchangeTraderId                                                   = ""
// }
func (m TradeAccountDataResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataResponseDefault)
}

// ToBuilder
func (m TradeAccountDataResponse) ToBuilder() TradeAccountDataResponseBuilder {
	return TradeAccountDataResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataResponseBuilder) Finish() TradeAccountDataResponse {
	return TradeAccountDataResponse{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataResponse) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataResponseBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// TradeAccountNotExist
func (m TradeAccountDataResponse) TradeAccountNotExist() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// TradeAccountNotExist
func (m TradeAccountDataResponseBuilder) TradeAccountNotExist() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// TradeAccount
func (m TradeAccountDataResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// TradeAccount
func (m TradeAccountDataResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// IsSimulated
func (m TradeAccountDataResponse) IsSimulated() uint8 {
	return message.Uint8VLS(m.Unsafe(), 16, 15)
}

// IsSimulated
func (m TradeAccountDataResponseBuilder) IsSimulated() uint8 {
	return message.Uint8VLS(m.Unsafe(), 16, 15)
}

// CurrencyCode
func (m TradeAccountDataResponse) CurrencyCode() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// CurrencyCode
func (m TradeAccountDataResponseBuilder) CurrencyCode() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// CurrentCashBalance
func (m TradeAccountDataResponse) CurrentCashBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 28, 20)
}

// CurrentCashBalance
func (m TradeAccountDataResponseBuilder) CurrentCashBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 28, 20)
}

// AvailableFundsForNewPositions
func (m TradeAccountDataResponse) AvailableFundsForNewPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 36, 28)
}

// AvailableFundsForNewPositions
func (m TradeAccountDataResponseBuilder) AvailableFundsForNewPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 36, 28)
}

// MarginRequirement
func (m TradeAccountDataResponse) MarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 44, 36)
}

// MarginRequirement
func (m TradeAccountDataResponseBuilder) MarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 44, 36)
}

// AccountValue
func (m TradeAccountDataResponse) AccountValue() float64 {
	return message.Float64VLS(m.Unsafe(), 52, 44)
}

// AccountValue
func (m TradeAccountDataResponseBuilder) AccountValue() float64 {
	return message.Float64VLS(m.Unsafe(), 52, 44)
}

// OpenPositionsProfitLoss
func (m TradeAccountDataResponse) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 60, 52)
}

// OpenPositionsProfitLoss
func (m TradeAccountDataResponseBuilder) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 60, 52)
}

// DailyProfitLoss
func (m TradeAccountDataResponse) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 68, 60)
}

// DailyProfitLoss
func (m TradeAccountDataResponseBuilder) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 68, 60)
}

// TransactionIdentifierForCashBalanceAdjustment
func (m TradeAccountDataResponse) TransactionIdentifierForCashBalanceAdjustment() uint64 {
	return message.Uint64VLS(m.Unsafe(), 76, 68)
}

// TransactionIdentifierForCashBalanceAdjustment
func (m TradeAccountDataResponseBuilder) TransactionIdentifierForCashBalanceAdjustment() uint64 {
	return message.Uint64VLS(m.Unsafe(), 76, 68)
}

// LastTransactionDateTime
func (m TradeAccountDataResponse) LastTransactionDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 84, 76)
}

// LastTransactionDateTime
func (m TradeAccountDataResponseBuilder) LastTransactionDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 84, 76)
}

// TrailingAccountValueAtWhichToNotAllowNewPositions
func (m TradeAccountDataResponse) TrailingAccountValueAtWhichToNotAllowNewPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 92, 84)
}

// TrailingAccountValueAtWhichToNotAllowNewPositions
func (m TradeAccountDataResponseBuilder) TrailingAccountValueAtWhichToNotAllowNewPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 92, 84)
}

// CalculatedDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponse) CalculatedDailyNetLossLimitInAccountCurrency() float64 {
	return message.Float64VLS(m.Unsafe(), 100, 92)
}

// CalculatedDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponseBuilder) CalculatedDailyNetLossLimitInAccountCurrency() float64 {
	return message.Float64VLS(m.Unsafe(), 100, 92)
}

// DailyNetLossLimitHasBeenReached
func (m TradeAccountDataResponse) DailyNetLossLimitHasBeenReached() uint8 {
	return message.Uint8VLS(m.Unsafe(), 101, 100)
}

// DailyNetLossLimitHasBeenReached
func (m TradeAccountDataResponseBuilder) DailyNetLossLimitHasBeenReached() uint8 {
	return message.Uint8VLS(m.Unsafe(), 101, 100)
}

// LastResetDailyNetLossManagementVariablesDateTimeUTC
func (m TradeAccountDataResponse) LastResetDailyNetLossManagementVariablesDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 109, 101)
}

// LastResetDailyNetLossManagementVariablesDateTimeUTC
func (m TradeAccountDataResponseBuilder) LastResetDailyNetLossManagementVariablesDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 109, 101)
}

// IsUnderRequiredMargin
func (m TradeAccountDataResponse) IsUnderRequiredMargin() uint8 {
	return message.Uint8VLS(m.Unsafe(), 110, 109)
}

// IsUnderRequiredMargin
func (m TradeAccountDataResponseBuilder) IsUnderRequiredMargin() uint8 {
	return message.Uint8VLS(m.Unsafe(), 110, 109)
}

// DailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponse) DailyNetLossLimitInAccountCurrency() float32 {
	return message.Float32VLS(m.Unsafe(), 114, 110)
}

// DailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponseBuilder) DailyNetLossLimitInAccountCurrency() float32 {
	return message.Float32VLS(m.Unsafe(), 114, 110)
}

// PercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataResponse) PercentOfCashBalanceForDailyNetLossLimit() int32 {
	return message.Int32VLS(m.Unsafe(), 118, 114)
}

// PercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataResponseBuilder) PercentOfCashBalanceForDailyNetLossLimit() int32 {
	return message.Int32VLS(m.Unsafe(), 118, 114)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataResponse) UseTrailingAccountValueToNotAllowIncreaseInPositions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 119, 118)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataResponseBuilder) UseTrailingAccountValueToNotAllowIncreaseInPositions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 119, 118)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponse) DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 120, 119)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponseBuilder) DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 120, 119)
}

// FlattenPositionsAtDailyLossLimit
func (m TradeAccountDataResponse) FlattenPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 121, 120)
}

// FlattenPositionsAtDailyLossLimit
func (m TradeAccountDataResponseBuilder) FlattenPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 121, 120)
}

// ClosePositionsAtEndOfDay
func (m TradeAccountDataResponse) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 122, 121)
}

// ClosePositionsAtEndOfDay
func (m TradeAccountDataResponseBuilder) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 122, 121)
}

// FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponse) FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 123, 122)
}

// FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponseBuilder) FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 123, 122)
}

// FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponse) FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 124, 123)
}

// FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponseBuilder) FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 124, 123)
}

// SenderSubID
func (m TradeAccountDataResponse) SenderSubID() string {
	return message.StringVLS(m.Unsafe(), 128, 124)
}

// SenderSubID
func (m TradeAccountDataResponseBuilder) SenderSubID() string {
	return message.StringVLS(m.Unsafe(), 128, 124)
}

// SenderLocationId
func (m TradeAccountDataResponse) SenderLocationId() string {
	return message.StringVLS(m.Unsafe(), 132, 128)
}

// SenderLocationId
func (m TradeAccountDataResponseBuilder) SenderLocationId() string {
	return message.StringVLS(m.Unsafe(), 132, 128)
}

// SelfMatchPreventionID
func (m TradeAccountDataResponse) SelfMatchPreventionID() string {
	return message.StringVLS(m.Unsafe(), 136, 132)
}

// SelfMatchPreventionID
func (m TradeAccountDataResponseBuilder) SelfMatchPreventionID() string {
	return message.StringVLS(m.Unsafe(), 136, 132)
}

// CTICode
func (m TradeAccountDataResponse) CTICode() int32 {
	return message.Int32VLS(m.Unsafe(), 140, 136)
}

// CTICode
func (m TradeAccountDataResponseBuilder) CTICode() int32 {
	return message.Int32VLS(m.Unsafe(), 140, 136)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataResponse) TradeAccountIsReadOnly() uint8 {
	return message.Uint8VLS(m.Unsafe(), 141, 140)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataResponseBuilder) TradeAccountIsReadOnly() uint8 {
	return message.Uint8VLS(m.Unsafe(), 141, 140)
}

// MaximumGlobalPositionQuantity
func (m TradeAccountDataResponse) MaximumGlobalPositionQuantity() int32 {
	return message.Int32VLS(m.Unsafe(), 145, 141)
}

// MaximumGlobalPositionQuantity
func (m TradeAccountDataResponseBuilder) MaximumGlobalPositionQuantity() int32 {
	return message.Int32VLS(m.Unsafe(), 145, 141)
}

// TradeFeePerContract
func (m TradeAccountDataResponse) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Unsafe(), 153, 145)
}

// TradeFeePerContract
func (m TradeAccountDataResponseBuilder) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Unsafe(), 153, 145)
}

// TradeFeePerShare
func (m TradeAccountDataResponse) TradeFeePerShare() float64 {
	return message.Float64VLS(m.Unsafe(), 161, 153)
}

// TradeFeePerShare
func (m TradeAccountDataResponseBuilder) TradeFeePerShare() float64 {
	return message.Float64VLS(m.Unsafe(), 161, 153)
}

// RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponse) RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 162, 161)
}

// RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponseBuilder) RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 162, 161)
}

// UsePercentOfMargin
func (m TradeAccountDataResponse) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Unsafe(), 166, 162)
}

// UsePercentOfMargin
func (m TradeAccountDataResponseBuilder) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Unsafe(), 166, 162)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponse) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Unsafe(), 170, 166)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponseBuilder) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Unsafe(), 170, 166)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponse) MaximumAllowedAccountBalanceForPositionsAsPercentage() int32 {
	return message.Int32VLS(m.Unsafe(), 174, 170)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponseBuilder) MaximumAllowedAccountBalanceForPositionsAsPercentage() int32 {
	return message.Int32VLS(m.Unsafe(), 174, 170)
}

// FirmID
func (m TradeAccountDataResponse) FirmID() string {
	return message.StringVLS(m.Unsafe(), 178, 174)
}

// FirmID
func (m TradeAccountDataResponseBuilder) FirmID() string {
	return message.StringVLS(m.Unsafe(), 178, 174)
}

// TradingIsDisabled
func (m TradeAccountDataResponse) TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Unsafe(), 179, 178)
}

// TradingIsDisabled
func (m TradeAccountDataResponseBuilder) TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Unsafe(), 179, 178)
}

// DescriptiveName
func (m TradeAccountDataResponse) DescriptiveName() string {
	return message.StringVLS(m.Unsafe(), 183, 179)
}

// DescriptiveName
func (m TradeAccountDataResponseBuilder) DescriptiveName() string {
	return message.StringVLS(m.Unsafe(), 183, 179)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataResponse) IsMasterFirmControlAccount() uint8 {
	return message.Uint8VLS(m.Unsafe(), 184, 183)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataResponseBuilder) IsMasterFirmControlAccount() uint8 {
	return message.Uint8VLS(m.Unsafe(), 184, 183)
}

// MinimumRequiredAccountValue
func (m TradeAccountDataResponse) MinimumRequiredAccountValue() float64 {
	return message.Float64VLS(m.Unsafe(), 192, 184)
}

// MinimumRequiredAccountValue
func (m TradeAccountDataResponseBuilder) MinimumRequiredAccountValue() float64 {
	return message.Float64VLS(m.Unsafe(), 192, 184)
}

// BeginTimeForDayMargin
func (m TradeAccountDataResponse) BeginTimeForDayMargin() int64 {
	return message.Int64VLS(m.Unsafe(), 200, 192)
}

// BeginTimeForDayMargin
func (m TradeAccountDataResponseBuilder) BeginTimeForDayMargin() int64 {
	return message.Int64VLS(m.Unsafe(), 200, 192)
}

// EndTimeForDayMargin
func (m TradeAccountDataResponse) EndTimeForDayMargin() int64 {
	return message.Int64VLS(m.Unsafe(), 208, 200)
}

// EndTimeForDayMargin
func (m TradeAccountDataResponseBuilder) EndTimeForDayMargin() int64 {
	return message.Int64VLS(m.Unsafe(), 208, 200)
}

// DayMarginTimeZone
func (m TradeAccountDataResponse) DayMarginTimeZone() string {
	return message.StringVLS(m.Unsafe(), 212, 208)
}

// DayMarginTimeZone
func (m TradeAccountDataResponseBuilder) DayMarginTimeZone() string {
	return message.StringVLS(m.Unsafe(), 212, 208)
}

// IsSnapshot
func (m TradeAccountDataResponse) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Unsafe(), 213, 212)
}

// IsSnapshot
func (m TradeAccountDataResponseBuilder) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Unsafe(), 213, 212)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponse) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 214, 213)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponseBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 214, 213)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponse) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 215, 214)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponseBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 215, 214)
}

// IsDeleted
func (m TradeAccountDataResponse) IsDeleted() uint8 {
	return message.Uint8VLS(m.Unsafe(), 216, 215)
}

// IsDeleted
func (m TradeAccountDataResponseBuilder) IsDeleted() uint8 {
	return message.Uint8VLS(m.Unsafe(), 216, 215)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponse) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 217, 216)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponseBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 217, 216)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponse) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 218, 217)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponseBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 218, 217)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataResponse) UseMasterFirm_SymbolLimitsArray() uint8 {
	return message.Uint8VLS(m.Unsafe(), 219, 218)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataResponseBuilder) UseMasterFirm_SymbolLimitsArray() uint8 {
	return message.Uint8VLS(m.Unsafe(), 219, 218)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataResponse) UseMasterFirm_TradeFees() uint8 {
	return message.Uint8VLS(m.Unsafe(), 220, 219)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataResponseBuilder) UseMasterFirm_TradeFees() uint8 {
	return message.Uint8VLS(m.Unsafe(), 220, 219)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataResponse) UseMasterFirm_TradeFeePerShare() uint8 {
	return message.Uint8VLS(m.Unsafe(), 221, 220)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataResponseBuilder) UseMasterFirm_TradeFeePerShare() uint8 {
	return message.Uint8VLS(m.Unsafe(), 221, 220)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponse) UseMasterFirm_RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 222, 221)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponseBuilder) UseMasterFirm_RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 222, 221)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataResponse) UseMasterFirm_UsePercentOfMargin() uint8 {
	return message.Uint8VLS(m.Unsafe(), 223, 222)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataResponseBuilder) UseMasterFirm_UsePercentOfMargin() uint8 {
	return message.Uint8VLS(m.Unsafe(), 223, 222)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponse) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() uint8 {
	return message.Uint8VLS(m.Unsafe(), 224, 223)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponseBuilder) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() uint8 {
	return message.Uint8VLS(m.Unsafe(), 224, 223)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataResponse) UseMasterFirm_MinimumRequiredAccountValue() uint8 {
	return message.Uint8VLS(m.Unsafe(), 225, 224)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataResponseBuilder) UseMasterFirm_MinimumRequiredAccountValue() uint8 {
	return message.Uint8VLS(m.Unsafe(), 225, 224)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataResponse) UseMasterFirm_MarginTimeSettings() uint8 {
	return message.Uint8VLS(m.Unsafe(), 226, 225)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataResponseBuilder) UseMasterFirm_MarginTimeSettings() uint8 {
	return message.Uint8VLS(m.Unsafe(), 226, 225)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataResponse) UseMasterFirm_TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Unsafe(), 227, 226)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataResponseBuilder) UseMasterFirm_TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Unsafe(), 227, 226)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataResponse) IsTradeStatisticsPublicallyShared() uint8 {
	return message.Uint8VLS(m.Unsafe(), 228, 227)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataResponseBuilder) IsTradeStatisticsPublicallyShared() uint8 {
	return message.Uint8VLS(m.Unsafe(), 228, 227)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataResponse) IsReadOnlyFollowingRequestsAllowed() uint8 {
	return message.Uint8VLS(m.Unsafe(), 229, 228)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataResponseBuilder) IsReadOnlyFollowingRequestsAllowed() uint8 {
	return message.Uint8VLS(m.Unsafe(), 229, 228)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataResponse) IsTradeAccountSharingAllowed() uint8 {
	return message.Uint8VLS(m.Unsafe(), 230, 229)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataResponseBuilder) IsTradeAccountSharingAllowed() uint8 {
	return message.Uint8VLS(m.Unsafe(), 230, 229)
}

// ReadOnlyShareToAllSCUsernames
func (m TradeAccountDataResponse) ReadOnlyShareToAllSCUsernames() uint8 {
	return message.Uint8VLS(m.Unsafe(), 231, 230)
}

// ReadOnlyShareToAllSCUsernames
func (m TradeAccountDataResponseBuilder) ReadOnlyShareToAllSCUsernames() uint8 {
	return message.Uint8VLS(m.Unsafe(), 231, 230)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataResponse) UseMasterFirm_SymbolCommissionsArray() uint8 {
	return message.Uint8VLS(m.Unsafe(), 232, 231)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataResponseBuilder) UseMasterFirm_SymbolCommissionsArray() uint8 {
	return message.Uint8VLS(m.Unsafe(), 232, 231)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponse) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 233, 232)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponseBuilder) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 233, 232)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponse) UseMasterFirm_UsePercentOfMarginForDayTrading() uint8 {
	return message.Uint8VLS(m.Unsafe(), 234, 233)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponseBuilder) UseMasterFirm_UsePercentOfMarginForDayTrading() uint8 {
	return message.Uint8VLS(m.Unsafe(), 234, 233)
}

// OpenPositionsProfitLossBasedOnSettlement
func (m TradeAccountDataResponse) OpenPositionsProfitLossBasedOnSettlement() float64 {
	return message.Float64VLS(m.Unsafe(), 242, 234)
}

// OpenPositionsProfitLossBasedOnSettlement
func (m TradeAccountDataResponseBuilder) OpenPositionsProfitLossBasedOnSettlement() float64 {
	return message.Float64VLS(m.Unsafe(), 242, 234)
}

// MarginRequirementFull
func (m TradeAccountDataResponse) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Unsafe(), 250, 242)
}

// MarginRequirementFull
func (m TradeAccountDataResponseBuilder) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Unsafe(), 250, 242)
}

// MarginRequirementFullPositionsOnly
func (m TradeAccountDataResponse) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Unsafe(), 258, 250)
}

// MarginRequirementFullPositionsOnly
func (m TradeAccountDataResponseBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Unsafe(), 258, 250)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataResponse) UseMasterFirm_TradeFeesFullOverride() uint8 {
	return message.Uint8VLS(m.Unsafe(), 259, 258)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataResponseBuilder) UseMasterFirm_TradeFeesFullOverride() uint8 {
	return message.Uint8VLS(m.Unsafe(), 259, 258)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataResponse) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() uint8 {
	return message.Uint8VLS(m.Unsafe(), 260, 259)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataResponseBuilder) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() uint8 {
	return message.Uint8VLS(m.Unsafe(), 260, 259)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataResponse) UseMasterFirm_UsePercentOfMarginFullOverride() uint8 {
	return message.Uint8VLS(m.Unsafe(), 261, 260)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataResponseBuilder) UseMasterFirm_UsePercentOfMarginFullOverride() uint8 {
	return message.Uint8VLS(m.Unsafe(), 261, 260)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataResponse) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() uint8 {
	return message.Uint8VLS(m.Unsafe(), 262, 261)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataResponseBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() uint8 {
	return message.Uint8VLS(m.Unsafe(), 262, 261)
}

// PeakMarginRequirement
func (m TradeAccountDataResponse) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 270, 262)
}

// PeakMarginRequirement
func (m TradeAccountDataResponseBuilder) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 270, 262)
}

// LiquidationOnlyMode
func (m TradeAccountDataResponse) LiquidationOnlyMode() uint8 {
	return message.Uint8VLS(m.Unsafe(), 271, 270)
}

// LiquidationOnlyMode
func (m TradeAccountDataResponseBuilder) LiquidationOnlyMode() uint8 {
	return message.Uint8VLS(m.Unsafe(), 271, 270)
}

// FlattenPositionsWhenUnderMarginIntradayTriggered
func (m TradeAccountDataResponse) FlattenPositionsWhenUnderMarginIntradayTriggered() uint8 {
	return message.Uint8VLS(m.Unsafe(), 272, 271)
}

// FlattenPositionsWhenUnderMarginIntradayTriggered
func (m TradeAccountDataResponseBuilder) FlattenPositionsWhenUnderMarginIntradayTriggered() uint8 {
	return message.Uint8VLS(m.Unsafe(), 272, 271)
}

// FlattenPositionsWhenUnderMinimumAccountValueTriggered
func (m TradeAccountDataResponse) FlattenPositionsWhenUnderMinimumAccountValueTriggered() uint8 {
	return message.Uint8VLS(m.Unsafe(), 273, 272)
}

// FlattenPositionsWhenUnderMinimumAccountValueTriggered
func (m TradeAccountDataResponseBuilder) FlattenPositionsWhenUnderMinimumAccountValueTriggered() uint8 {
	return message.Uint8VLS(m.Unsafe(), 273, 272)
}

// AccountValueAtEndOfDayCaptureTime
func (m TradeAccountDataResponse) AccountValueAtEndOfDayCaptureTime() float64 {
	return message.Float64VLS(m.Unsafe(), 281, 273)
}

// AccountValueAtEndOfDayCaptureTime
func (m TradeAccountDataResponseBuilder) AccountValueAtEndOfDayCaptureTime() float64 {
	return message.Float64VLS(m.Unsafe(), 281, 273)
}

// EndOfDayCaptureTime
func (m TradeAccountDataResponse) EndOfDayCaptureTime() int64 {
	return message.Int64VLS(m.Unsafe(), 289, 281)
}

// EndOfDayCaptureTime
func (m TradeAccountDataResponseBuilder) EndOfDayCaptureTime() int64 {
	return message.Int64VLS(m.Unsafe(), 289, 281)
}

// CustomerOrFirm
func (m TradeAccountDataResponse) CustomerOrFirm() uint8 {
	return message.Uint8VLS(m.Unsafe(), 290, 289)
}

// CustomerOrFirm
func (m TradeAccountDataResponseBuilder) CustomerOrFirm() uint8 {
	return message.Uint8VLS(m.Unsafe(), 290, 289)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataResponse) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 291, 290)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataResponseBuilder) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 291, 290)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataResponse) MasterFirm_FlattenCancelWhenUnderMinimumAccountValue() uint8 {
	return message.Uint8VLS(m.Unsafe(), 292, 291)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataResponseBuilder) MasterFirm_FlattenCancelWhenUnderMinimumAccountValue() uint8 {
	return message.Uint8VLS(m.Unsafe(), 292, 291)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataResponse) MasterFirm_FlattenCancelWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 293, 292)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataResponseBuilder) MasterFirm_FlattenCancelWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 293, 292)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataResponse) MasterFirm_FlattenCancelWhenUnderMarginEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 294, 293)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataResponseBuilder) MasterFirm_FlattenCancelWhenUnderMarginEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 294, 293)
}

// MasterFirm_MaximumOrderQuantity
func (m TradeAccountDataResponse) MasterFirm_MaximumOrderQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 298, 294)
}

// MasterFirm_MaximumOrderQuantity
func (m TradeAccountDataResponseBuilder) MasterFirm_MaximumOrderQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 298, 294)
}

// LastTriggerDateTimeUTCForDailyLossLimit
func (m TradeAccountDataResponse) LastTriggerDateTimeUTCForDailyLossLimit() int64 {
	return message.Int64VLS(m.Unsafe(), 306, 298)
}

// LastTriggerDateTimeUTCForDailyLossLimit
func (m TradeAccountDataResponseBuilder) LastTriggerDateTimeUTCForDailyLossLimit() int64 {
	return message.Int64VLS(m.Unsafe(), 306, 298)
}

// OpenPositionsProfitLossIsDelayed
func (m TradeAccountDataResponse) OpenPositionsProfitLossIsDelayed() uint8 {
	return message.Uint8VLS(m.Unsafe(), 307, 306)
}

// OpenPositionsProfitLossIsDelayed
func (m TradeAccountDataResponseBuilder) OpenPositionsProfitLossIsDelayed() uint8 {
	return message.Uint8VLS(m.Unsafe(), 307, 306)
}

// ExchangeTraderId
func (m TradeAccountDataResponse) ExchangeTraderId() string {
	return message.StringVLS(m.Unsafe(), 311, 307)
}

// ExchangeTraderId
func (m TradeAccountDataResponseBuilder) ExchangeTraderId() string {
	return message.StringVLS(m.Unsafe(), 311, 307)
}

// SetRequestID
func (m TradeAccountDataResponseBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetTradeAccountNotExist
func (m TradeAccountDataResponseBuilder) SetTradeAccountNotExist(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 11, 10, value)
}

// SetTradeAccount
func (m *TradeAccountDataResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 15, 11, value)
}

// SetIsSimulated
func (m TradeAccountDataResponseBuilder) SetIsSimulated(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 16, 15, value)
}

// SetCurrencyCode
func (m *TradeAccountDataResponseBuilder) SetCurrencyCode(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetCurrentCashBalance
func (m TradeAccountDataResponseBuilder) SetCurrentCashBalance(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 28, 20, value)
}

// SetAvailableFundsForNewPositions
func (m TradeAccountDataResponseBuilder) SetAvailableFundsForNewPositions(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 36, 28, value)
}

// SetMarginRequirement
func (m TradeAccountDataResponseBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 44, 36, value)
}

// SetAccountValue
func (m TradeAccountDataResponseBuilder) SetAccountValue(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 52, 44, value)
}

// SetOpenPositionsProfitLoss
func (m TradeAccountDataResponseBuilder) SetOpenPositionsProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 60, 52, value)
}

// SetDailyProfitLoss
func (m TradeAccountDataResponseBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 68, 60, value)
}

// SetTransactionIdentifierForCashBalanceAdjustment
func (m TradeAccountDataResponseBuilder) SetTransactionIdentifierForCashBalanceAdjustment(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 76, 68, value)
}

// SetLastTransactionDateTime
func (m TradeAccountDataResponseBuilder) SetLastTransactionDateTime(value int64) {
	message.SetInt64VLS(m.Unsafe(), 84, 76, value)
}

// SetTrailingAccountValueAtWhichToNotAllowNewPositions
func (m TradeAccountDataResponseBuilder) SetTrailingAccountValueAtWhichToNotAllowNewPositions(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 92, 84, value)
}

// SetCalculatedDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponseBuilder) SetCalculatedDailyNetLossLimitInAccountCurrency(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 100, 92, value)
}

// SetDailyNetLossLimitHasBeenReached
func (m TradeAccountDataResponseBuilder) SetDailyNetLossLimitHasBeenReached(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 101, 100, value)
}

// SetLastResetDailyNetLossManagementVariablesDateTimeUTC
func (m TradeAccountDataResponseBuilder) SetLastResetDailyNetLossManagementVariablesDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Unsafe(), 109, 101, value)
}

// SetIsUnderRequiredMargin
func (m TradeAccountDataResponseBuilder) SetIsUnderRequiredMargin(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 110, 109, value)
}

// SetDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponseBuilder) SetDailyNetLossLimitInAccountCurrency(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 114, 110, value)
}

// SetPercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataResponseBuilder) SetPercentOfCashBalanceForDailyNetLossLimit(value int32) {
	message.SetInt32VLS(m.Unsafe(), 118, 114, value)
}

// SetUseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataResponseBuilder) SetUseTrailingAccountValueToNotAllowIncreaseInPositions(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 119, 118, value)
}

// SetDoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponseBuilder) SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 120, 119, value)
}

// SetFlattenPositionsAtDailyLossLimit
func (m TradeAccountDataResponseBuilder) SetFlattenPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 121, 120, value)
}

// SetClosePositionsAtEndOfDay
func (m TradeAccountDataResponseBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 122, 121, value)
}

// SetFlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponseBuilder) SetFlattenPositionsWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 123, 122, value)
}

// SetFlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponseBuilder) SetFlattenPositionsWhenUnderMarginAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 124, 123, value)
}

// SetSenderSubID
func (m *TradeAccountDataResponseBuilder) SetSenderSubID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 128, 124, value)
}

// SetSenderLocationId
func (m *TradeAccountDataResponseBuilder) SetSenderLocationId(value string) {
	message.SetStringVLS(&m.VLSBuilder, 132, 128, value)
}

// SetSelfMatchPreventionID
func (m *TradeAccountDataResponseBuilder) SetSelfMatchPreventionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 136, 132, value)
}

// SetCTICode
func (m TradeAccountDataResponseBuilder) SetCTICode(value int32) {
	message.SetInt32VLS(m.Unsafe(), 140, 136, value)
}

// SetTradeAccountIsReadOnly
func (m TradeAccountDataResponseBuilder) SetTradeAccountIsReadOnly(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 141, 140, value)
}

// SetMaximumGlobalPositionQuantity
func (m TradeAccountDataResponseBuilder) SetMaximumGlobalPositionQuantity(value int32) {
	message.SetInt32VLS(m.Unsafe(), 145, 141, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataResponseBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 153, 145, value)
}

// SetTradeFeePerShare
func (m TradeAccountDataResponseBuilder) SetTradeFeePerShare(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 161, 153, value)
}

// SetRequireSufficientMarginForNewPositions
func (m TradeAccountDataResponseBuilder) SetRequireSufficientMarginForNewPositions(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 162, 161, value)
}

// SetUsePercentOfMargin
func (m TradeAccountDataResponseBuilder) SetUsePercentOfMargin(value int32) {
	message.SetInt32VLS(m.Unsafe(), 166, 162, value)
}

// SetUsePercentOfMarginForDayTrading
func (m TradeAccountDataResponseBuilder) SetUsePercentOfMarginForDayTrading(value int32) {
	message.SetInt32VLS(m.Unsafe(), 170, 166, value)
}

// SetMaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponseBuilder) SetMaximumAllowedAccountBalanceForPositionsAsPercentage(value int32) {
	message.SetInt32VLS(m.Unsafe(), 174, 170, value)
}

// SetFirmID
func (m *TradeAccountDataResponseBuilder) SetFirmID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 178, 174, value)
}

// SetTradingIsDisabled
func (m TradeAccountDataResponseBuilder) SetTradingIsDisabled(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 179, 178, value)
}

// SetDescriptiveName
func (m *TradeAccountDataResponseBuilder) SetDescriptiveName(value string) {
	message.SetStringVLS(&m.VLSBuilder, 183, 179, value)
}

// SetIsMasterFirmControlAccount
func (m TradeAccountDataResponseBuilder) SetIsMasterFirmControlAccount(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 184, 183, value)
}

// SetMinimumRequiredAccountValue
func (m TradeAccountDataResponseBuilder) SetMinimumRequiredAccountValue(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 192, 184, value)
}

// SetBeginTimeForDayMargin
func (m TradeAccountDataResponseBuilder) SetBeginTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Unsafe(), 200, 192, value)
}

// SetEndTimeForDayMargin
func (m TradeAccountDataResponseBuilder) SetEndTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Unsafe(), 208, 200, value)
}

// SetDayMarginTimeZone
func (m *TradeAccountDataResponseBuilder) SetDayMarginTimeZone(value string) {
	message.SetStringVLS(&m.VLSBuilder, 212, 208, value)
}

// SetIsSnapshot
func (m TradeAccountDataResponseBuilder) SetIsSnapshot(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 213, 212, value)
}

// SetIsFirstMessageInBatch
func (m TradeAccountDataResponseBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 214, 213, value)
}

// SetIsLastMessageInBatch
func (m TradeAccountDataResponseBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 215, 214, value)
}

// SetIsDeleted
func (m TradeAccountDataResponseBuilder) SetIsDeleted(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 216, 215, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 217, 216, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 218, 217, value)
}

// SetUseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_SymbolLimitsArray(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 219, 218, value)
}

// SetUseMasterFirm_TradeFees
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_TradeFees(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 220, 219, value)
}

// SetUseMasterFirm_TradeFeePerShare
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_TradeFeePerShare(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 221, 220, value)
}

// SetUseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_RequireSufficientMarginForNewPositions(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 222, 221, value)
}

// SetUseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_UsePercentOfMargin(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 223, 222, value)
}

// SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 224, 223, value)
}

// SetUseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_MinimumRequiredAccountValue(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 225, 224, value)
}

// SetUseMasterFirm_MarginTimeSettings
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_MarginTimeSettings(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 226, 225, value)
}

// SetUseMasterFirm_TradingIsDisabled
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_TradingIsDisabled(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 227, 226, value)
}

// SetIsTradeStatisticsPublicallyShared
func (m TradeAccountDataResponseBuilder) SetIsTradeStatisticsPublicallyShared(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 228, 227, value)
}

// SetIsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataResponseBuilder) SetIsReadOnlyFollowingRequestsAllowed(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 229, 228, value)
}

// SetIsTradeAccountSharingAllowed
func (m TradeAccountDataResponseBuilder) SetIsTradeAccountSharingAllowed(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 230, 229, value)
}

// SetReadOnlyShareToAllSCUsernames
func (m TradeAccountDataResponseBuilder) SetReadOnlyShareToAllSCUsernames(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 231, 230, value)
}

// SetUseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_SymbolCommissionsArray(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 232, 231, value)
}

// SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 233, 232, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTrading(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 234, 233, value)
}

// SetOpenPositionsProfitLossBasedOnSettlement
func (m TradeAccountDataResponseBuilder) SetOpenPositionsProfitLossBasedOnSettlement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 242, 234, value)
}

// SetMarginRequirementFull
func (m TradeAccountDataResponseBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 250, 242, value)
}

// SetMarginRequirementFullPositionsOnly
func (m TradeAccountDataResponseBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 258, 250, value)
}

// SetUseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_TradeFeesFullOverride(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 259, 258, value)
}

// SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 260, 259, value)
}

// SetUseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_UsePercentOfMarginFullOverride(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 261, 260, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 262, 261, value)
}

// SetPeakMarginRequirement
func (m TradeAccountDataResponseBuilder) SetPeakMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 270, 262, value)
}

// SetLiquidationOnlyMode
func (m TradeAccountDataResponseBuilder) SetLiquidationOnlyMode(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 271, 270, value)
}

// SetFlattenPositionsWhenUnderMarginIntradayTriggered
func (m TradeAccountDataResponseBuilder) SetFlattenPositionsWhenUnderMarginIntradayTriggered(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 272, 271, value)
}

// SetFlattenPositionsWhenUnderMinimumAccountValueTriggered
func (m TradeAccountDataResponseBuilder) SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 273, 272, value)
}

// SetAccountValueAtEndOfDayCaptureTime
func (m TradeAccountDataResponseBuilder) SetAccountValueAtEndOfDayCaptureTime(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 281, 273, value)
}

// SetEndOfDayCaptureTime
func (m TradeAccountDataResponseBuilder) SetEndOfDayCaptureTime(value int64) {
	message.SetInt64VLS(m.Unsafe(), 289, 281, value)
}

// SetCustomerOrFirm
func (m TradeAccountDataResponseBuilder) SetCustomerOrFirm(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 290, 289, value)
}

// SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataResponseBuilder) SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 291, 290, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataResponseBuilder) SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 292, 291, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataResponseBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 293, 292, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataResponseBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 294, 293, value)
}

// SetMasterFirm_MaximumOrderQuantity
func (m TradeAccountDataResponseBuilder) SetMasterFirm_MaximumOrderQuantity(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 298, 294, value)
}

// SetLastTriggerDateTimeUTCForDailyLossLimit
func (m TradeAccountDataResponseBuilder) SetLastTriggerDateTimeUTCForDailyLossLimit(value int64) {
	message.SetInt64VLS(m.Unsafe(), 306, 298, value)
}

// SetOpenPositionsProfitLossIsDelayed
func (m TradeAccountDataResponseBuilder) SetOpenPositionsProfitLossIsDelayed(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 307, 306, value)
}

// SetExchangeTraderId
func (m *TradeAccountDataResponseBuilder) SetExchangeTraderId(value string) {
	message.SetStringVLS(&m.VLSBuilder, 311, 307, value)
}

// Copy
func (m TradeAccountDataResponse) Copy(to TradeAccountDataResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccountNotExist(m.TradeAccountNotExist())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsSimulated(m.IsSimulated())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetCurrentCashBalance(m.CurrentCashBalance())
	to.SetAvailableFundsForNewPositions(m.AvailableFundsForNewPositions())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetAccountValue(m.AccountValue())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetTransactionIdentifierForCashBalanceAdjustment(m.TransactionIdentifierForCashBalanceAdjustment())
	to.SetLastTransactionDateTime(m.LastTransactionDateTime())
	to.SetTrailingAccountValueAtWhichToNotAllowNewPositions(m.TrailingAccountValueAtWhichToNotAllowNewPositions())
	to.SetCalculatedDailyNetLossLimitInAccountCurrency(m.CalculatedDailyNetLossLimitInAccountCurrency())
	to.SetDailyNetLossLimitHasBeenReached(m.DailyNetLossLimitHasBeenReached())
	to.SetLastResetDailyNetLossManagementVariablesDateTimeUTC(m.LastResetDailyNetLossManagementVariablesDateTimeUTC())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetIsDeleted(m.IsDeleted())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetOpenPositionsProfitLossBasedOnSettlement(m.OpenPositionsProfitLossBasedOnSettlement())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetFlattenPositionsWhenUnderMarginIntradayTriggered(m.FlattenPositionsWhenUnderMarginIntradayTriggered())
	to.SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(m.FlattenPositionsWhenUnderMinimumAccountValueTriggered())
	to.SetAccountValueAtEndOfDayCaptureTime(m.AccountValueAtEndOfDayCaptureTime())
	to.SetEndOfDayCaptureTime(m.EndOfDayCaptureTime())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetLastTriggerDateTimeUTCForDailyLossLimit(m.LastTriggerDateTimeUTCForDailyLossLimit())
	to.SetOpenPositionsProfitLossIsDelayed(m.OpenPositionsProfitLossIsDelayed())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// Copy
func (m TradeAccountDataResponseBuilder) Copy(to TradeAccountDataResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccountNotExist(m.TradeAccountNotExist())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsSimulated(m.IsSimulated())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetCurrentCashBalance(m.CurrentCashBalance())
	to.SetAvailableFundsForNewPositions(m.AvailableFundsForNewPositions())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetAccountValue(m.AccountValue())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetTransactionIdentifierForCashBalanceAdjustment(m.TransactionIdentifierForCashBalanceAdjustment())
	to.SetLastTransactionDateTime(m.LastTransactionDateTime())
	to.SetTrailingAccountValueAtWhichToNotAllowNewPositions(m.TrailingAccountValueAtWhichToNotAllowNewPositions())
	to.SetCalculatedDailyNetLossLimitInAccountCurrency(m.CalculatedDailyNetLossLimitInAccountCurrency())
	to.SetDailyNetLossLimitHasBeenReached(m.DailyNetLossLimitHasBeenReached())
	to.SetLastResetDailyNetLossManagementVariablesDateTimeUTC(m.LastResetDailyNetLossManagementVariablesDateTimeUTC())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetIsDeleted(m.IsDeleted())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetOpenPositionsProfitLossBasedOnSettlement(m.OpenPositionsProfitLossBasedOnSettlement())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetFlattenPositionsWhenUnderMarginIntradayTriggered(m.FlattenPositionsWhenUnderMarginIntradayTriggered())
	to.SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(m.FlattenPositionsWhenUnderMinimumAccountValueTriggered())
	to.SetAccountValueAtEndOfDayCaptureTime(m.AccountValueAtEndOfDayCaptureTime())
	to.SetEndOfDayCaptureTime(m.EndOfDayCaptureTime())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetLastTriggerDateTimeUTCForDailyLossLimit(m.LastTriggerDateTimeUTCForDailyLossLimit())
	to.SetOpenPositionsProfitLossIsDelayed(m.OpenPositionsProfitLossIsDelayed())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// CopyPointer
func (m TradeAccountDataResponse) CopyPointer(to TradeAccountDataResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccountNotExist(m.TradeAccountNotExist())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsSimulated(m.IsSimulated())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetCurrentCashBalance(m.CurrentCashBalance())
	to.SetAvailableFundsForNewPositions(m.AvailableFundsForNewPositions())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetAccountValue(m.AccountValue())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetTransactionIdentifierForCashBalanceAdjustment(m.TransactionIdentifierForCashBalanceAdjustment())
	to.SetLastTransactionDateTime(m.LastTransactionDateTime())
	to.SetTrailingAccountValueAtWhichToNotAllowNewPositions(m.TrailingAccountValueAtWhichToNotAllowNewPositions())
	to.SetCalculatedDailyNetLossLimitInAccountCurrency(m.CalculatedDailyNetLossLimitInAccountCurrency())
	to.SetDailyNetLossLimitHasBeenReached(m.DailyNetLossLimitHasBeenReached())
	to.SetLastResetDailyNetLossManagementVariablesDateTimeUTC(m.LastResetDailyNetLossManagementVariablesDateTimeUTC())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetIsDeleted(m.IsDeleted())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetOpenPositionsProfitLossBasedOnSettlement(m.OpenPositionsProfitLossBasedOnSettlement())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetFlattenPositionsWhenUnderMarginIntradayTriggered(m.FlattenPositionsWhenUnderMarginIntradayTriggered())
	to.SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(m.FlattenPositionsWhenUnderMinimumAccountValueTriggered())
	to.SetAccountValueAtEndOfDayCaptureTime(m.AccountValueAtEndOfDayCaptureTime())
	to.SetEndOfDayCaptureTime(m.EndOfDayCaptureTime())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetLastTriggerDateTimeUTCForDailyLossLimit(m.LastTriggerDateTimeUTCForDailyLossLimit())
	to.SetOpenPositionsProfitLossIsDelayed(m.OpenPositionsProfitLossIsDelayed())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// CopyPointer
func (m TradeAccountDataResponseBuilder) CopyPointer(to TradeAccountDataResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccountNotExist(m.TradeAccountNotExist())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsSimulated(m.IsSimulated())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetCurrentCashBalance(m.CurrentCashBalance())
	to.SetAvailableFundsForNewPositions(m.AvailableFundsForNewPositions())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetAccountValue(m.AccountValue())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetTransactionIdentifierForCashBalanceAdjustment(m.TransactionIdentifierForCashBalanceAdjustment())
	to.SetLastTransactionDateTime(m.LastTransactionDateTime())
	to.SetTrailingAccountValueAtWhichToNotAllowNewPositions(m.TrailingAccountValueAtWhichToNotAllowNewPositions())
	to.SetCalculatedDailyNetLossLimitInAccountCurrency(m.CalculatedDailyNetLossLimitInAccountCurrency())
	to.SetDailyNetLossLimitHasBeenReached(m.DailyNetLossLimitHasBeenReached())
	to.SetLastResetDailyNetLossManagementVariablesDateTimeUTC(m.LastResetDailyNetLossManagementVariablesDateTimeUTC())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetIsDeleted(m.IsDeleted())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetOpenPositionsProfitLossBasedOnSettlement(m.OpenPositionsProfitLossBasedOnSettlement())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetFlattenPositionsWhenUnderMarginIntradayTriggered(m.FlattenPositionsWhenUnderMarginIntradayTriggered())
	to.SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(m.FlattenPositionsWhenUnderMinimumAccountValueTriggered())
	to.SetAccountValueAtEndOfDayCaptureTime(m.AccountValueAtEndOfDayCaptureTime())
	to.SetEndOfDayCaptureTime(m.EndOfDayCaptureTime())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetLastTriggerDateTimeUTCForDailyLossLimit(m.LastTriggerDateTimeUTCForDailyLossLimit())
	to.SetOpenPositionsProfitLossIsDelayed(m.OpenPositionsProfitLossIsDelayed())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}
