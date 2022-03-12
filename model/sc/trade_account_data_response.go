package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeAccountDataResponseSize = 311

// {
//     Size                                                               = TradeAccountDataResponseSize  (311)
//     Type                                                               = TRADE_ACCOUNT_DATA_RESPONSE  (10116)
//     BaseSize                                                           = TradeAccountDataResponseSize  (311)
//     RequestID                                                          = 0
//     TradeAccountNotExist                                               = 0
//     TradeAccount                                                       = ""
//     IsSimulated                                                        = false
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
//     IsUnderRequiredMargin                                              = false
//     DailyNetLossLimitInAccountCurrency                                 = 0
//     PercentOfCashBalanceForDailyNetLossLimit                           = 0
//     UseTrailingAccountValueToNotAllowIncreaseInPositions               = false
//     DoNotAllowIncreaseInPositionsAtDailyLossLimit                      = false
//     FlattenPositionsAtDailyLossLimit                                   = false
//     ClosePositionsAtEndOfDay                                           = false
//     FlattenPositionsWhenUnderMarginIntraday                            = true
//     FlattenPositionsWhenUnderMarginAtEndOfDay                          = false
//     SenderSubID                                                        = ""
//     SenderLocationId                                                   = ""
//     SelfMatchPreventionID                                              = ""
//     CTICode                                                            = 0
//     TradeAccountIsReadOnly                                             = false
//     MaximumGlobalPositionQuantity                                      = 0
//     TradeFeePerContract                                                = 0
//     TradeFeePerShare                                                   = 0
//     RequireSufficientMarginForNewPositions                             = true
//     UsePercentOfMargin                                                 = 100
//     UsePercentOfMarginForDayTrading                                    = 100
//     MaximumAllowedAccountBalanceForPositionsAsPercentage               = 100
//     FirmID                                                             = ""
//     TradingIsDisabled                                                  = false
//     DescriptiveName                                                    = ""
//     IsMasterFirmControlAccount                                         = false
//     MinimumRequiredAccountValue                                        = 0
//     BeginTimeForDayMargin                                              = 0
//     EndTimeForDayMargin                                                = 0
//     DayMarginTimeZone                                                  = ""
//     IsSnapshot                                                         = false
//     IsFirstMessageInBatch                                              = false
//     IsLastMessageInBatch                                               = false
//     IsDeleted                                                          = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday              = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay            = false
//     UseMasterFirm_SymbolLimitsArray                                    = false
//     UseMasterFirm_TradeFees                                            = false
//     UseMasterFirm_TradeFeePerShare                                     = false
//     UseMasterFirm_RequireSufficientMarginForNewPositions               = false
//     UseMasterFirm_UsePercentOfMargin                                   = false
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage = false
//     UseMasterFirm_MinimumRequiredAccountValue                          = false
//     UseMasterFirm_MarginTimeSettings                                   = false
//     UseMasterFirm_TradingIsDisabled                                    = false
//     IsTradeStatisticsPublicallyShared                                  = false
//     IsReadOnlyFollowingRequestsAllowed                                 = false
//     IsTradeAccountSharingAllowed                                       = false
//     ReadOnlyShareToAllSCUsernames                                      = 0
//     UseMasterFirm_SymbolCommissionsArray                               = false
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit        = false
//     UseMasterFirm_UsePercentOfMarginForDayTrading                      = false
//     OpenPositionsProfitLossBasedOnSettlement                           = 0
//     MarginRequirementFull                                              = 0
//     MarginRequirementFullPositionsOnly                                 = 0
//     UseMasterFirm_TradeFeesFullOverride                                = false
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders         = false
//     UseMasterFirm_UsePercentOfMarginFullOverride                       = false
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride          = false
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
//     OpenPositionsProfitLossIsDelayed                                   = false
//     ExchangeTraderId                                                   = ""
// }
var _TradeAccountDataResponseDefault = []byte{55, 1, 132, 39, 55, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 100, 0, 0, 0, 100, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataResponse struct {
	message.VLS
}

type TradeAccountDataResponseBuilder struct {
	message.VLSBuilder
}

type TradeAccountDataResponsePointer struct {
	message.VLSPointer
}

type TradeAccountDataResponsePointerBuilder struct {
	message.VLSPointerBuilder
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

func AllocTradeAccountDataResponse() TradeAccountDataResponsePointerBuilder {
	a := TradeAccountDataResponsePointerBuilder{message.AllocVLSBuilder(311)}
	a.Ptr.SetBytes(0, _TradeAccountDataResponseDefault)
	return a
}

func AllocTradeAccountDataResponseFrom(b []byte) TradeAccountDataResponsePointer {
	return TradeAccountDataResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                                                               = TradeAccountDataResponseSize  (311)
//     Type                                                               = TRADE_ACCOUNT_DATA_RESPONSE  (10116)
//     BaseSize                                                           = TradeAccountDataResponseSize  (311)
//     RequestID                                                          = 0
//     TradeAccountNotExist                                               = 0
//     TradeAccount                                                       = ""
//     IsSimulated                                                        = false
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
//     IsUnderRequiredMargin                                              = false
//     DailyNetLossLimitInAccountCurrency                                 = 0
//     PercentOfCashBalanceForDailyNetLossLimit                           = 0
//     UseTrailingAccountValueToNotAllowIncreaseInPositions               = false
//     DoNotAllowIncreaseInPositionsAtDailyLossLimit                      = false
//     FlattenPositionsAtDailyLossLimit                                   = false
//     ClosePositionsAtEndOfDay                                           = false
//     FlattenPositionsWhenUnderMarginIntraday                            = true
//     FlattenPositionsWhenUnderMarginAtEndOfDay                          = false
//     SenderSubID                                                        = ""
//     SenderLocationId                                                   = ""
//     SelfMatchPreventionID                                              = ""
//     CTICode                                                            = 0
//     TradeAccountIsReadOnly                                             = false
//     MaximumGlobalPositionQuantity                                      = 0
//     TradeFeePerContract                                                = 0
//     TradeFeePerShare                                                   = 0
//     RequireSufficientMarginForNewPositions                             = true
//     UsePercentOfMargin                                                 = 100
//     UsePercentOfMarginForDayTrading                                    = 100
//     MaximumAllowedAccountBalanceForPositionsAsPercentage               = 100
//     FirmID                                                             = ""
//     TradingIsDisabled                                                  = false
//     DescriptiveName                                                    = ""
//     IsMasterFirmControlAccount                                         = false
//     MinimumRequiredAccountValue                                        = 0
//     BeginTimeForDayMargin                                              = 0
//     EndTimeForDayMargin                                                = 0
//     DayMarginTimeZone                                                  = ""
//     IsSnapshot                                                         = false
//     IsFirstMessageInBatch                                              = false
//     IsLastMessageInBatch                                               = false
//     IsDeleted                                                          = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday              = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay            = false
//     UseMasterFirm_SymbolLimitsArray                                    = false
//     UseMasterFirm_TradeFees                                            = false
//     UseMasterFirm_TradeFeePerShare                                     = false
//     UseMasterFirm_RequireSufficientMarginForNewPositions               = false
//     UseMasterFirm_UsePercentOfMargin                                   = false
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage = false
//     UseMasterFirm_MinimumRequiredAccountValue                          = false
//     UseMasterFirm_MarginTimeSettings                                   = false
//     UseMasterFirm_TradingIsDisabled                                    = false
//     IsTradeStatisticsPublicallyShared                                  = false
//     IsReadOnlyFollowingRequestsAllowed                                 = false
//     IsTradeAccountSharingAllowed                                       = false
//     ReadOnlyShareToAllSCUsernames                                      = 0
//     UseMasterFirm_SymbolCommissionsArray                               = false
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit        = false
//     UseMasterFirm_UsePercentOfMarginForDayTrading                      = false
//     OpenPositionsProfitLossBasedOnSettlement                           = 0
//     MarginRequirementFull                                              = 0
//     MarginRequirementFullPositionsOnly                                 = 0
//     UseMasterFirm_TradeFeesFullOverride                                = false
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders         = false
//     UseMasterFirm_UsePercentOfMarginFullOverride                       = false
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride          = false
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
//     OpenPositionsProfitLossIsDelayed                                   = false
//     ExchangeTraderId                                                   = ""
// }
func (m TradeAccountDataResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataResponseDefault)
}

// Clear
// {
//     Size                                                               = TradeAccountDataResponseSize  (311)
//     Type                                                               = TRADE_ACCOUNT_DATA_RESPONSE  (10116)
//     BaseSize                                                           = TradeAccountDataResponseSize  (311)
//     RequestID                                                          = 0
//     TradeAccountNotExist                                               = 0
//     TradeAccount                                                       = ""
//     IsSimulated                                                        = false
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
//     IsUnderRequiredMargin                                              = false
//     DailyNetLossLimitInAccountCurrency                                 = 0
//     PercentOfCashBalanceForDailyNetLossLimit                           = 0
//     UseTrailingAccountValueToNotAllowIncreaseInPositions               = false
//     DoNotAllowIncreaseInPositionsAtDailyLossLimit                      = false
//     FlattenPositionsAtDailyLossLimit                                   = false
//     ClosePositionsAtEndOfDay                                           = false
//     FlattenPositionsWhenUnderMarginIntraday                            = true
//     FlattenPositionsWhenUnderMarginAtEndOfDay                          = false
//     SenderSubID                                                        = ""
//     SenderLocationId                                                   = ""
//     SelfMatchPreventionID                                              = ""
//     CTICode                                                            = 0
//     TradeAccountIsReadOnly                                             = false
//     MaximumGlobalPositionQuantity                                      = 0
//     TradeFeePerContract                                                = 0
//     TradeFeePerShare                                                   = 0
//     RequireSufficientMarginForNewPositions                             = true
//     UsePercentOfMargin                                                 = 100
//     UsePercentOfMarginForDayTrading                                    = 100
//     MaximumAllowedAccountBalanceForPositionsAsPercentage               = 100
//     FirmID                                                             = ""
//     TradingIsDisabled                                                  = false
//     DescriptiveName                                                    = ""
//     IsMasterFirmControlAccount                                         = false
//     MinimumRequiredAccountValue                                        = 0
//     BeginTimeForDayMargin                                              = 0
//     EndTimeForDayMargin                                                = 0
//     DayMarginTimeZone                                                  = ""
//     IsSnapshot                                                         = false
//     IsFirstMessageInBatch                                              = false
//     IsLastMessageInBatch                                               = false
//     IsDeleted                                                          = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday              = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay            = false
//     UseMasterFirm_SymbolLimitsArray                                    = false
//     UseMasterFirm_TradeFees                                            = false
//     UseMasterFirm_TradeFeePerShare                                     = false
//     UseMasterFirm_RequireSufficientMarginForNewPositions               = false
//     UseMasterFirm_UsePercentOfMargin                                   = false
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage = false
//     UseMasterFirm_MinimumRequiredAccountValue                          = false
//     UseMasterFirm_MarginTimeSettings                                   = false
//     UseMasterFirm_TradingIsDisabled                                    = false
//     IsTradeStatisticsPublicallyShared                                  = false
//     IsReadOnlyFollowingRequestsAllowed                                 = false
//     IsTradeAccountSharingAllowed                                       = false
//     ReadOnlyShareToAllSCUsernames                                      = 0
//     UseMasterFirm_SymbolCommissionsArray                               = false
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit        = false
//     UseMasterFirm_UsePercentOfMarginForDayTrading                      = false
//     OpenPositionsProfitLossBasedOnSettlement                           = 0
//     MarginRequirementFull                                              = 0
//     MarginRequirementFullPositionsOnly                                 = 0
//     UseMasterFirm_TradeFeesFullOverride                                = false
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders         = false
//     UseMasterFirm_UsePercentOfMarginFullOverride                       = false
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride          = false
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
//     OpenPositionsProfitLossIsDelayed                                   = false
//     ExchangeTraderId                                                   = ""
// }
func (m TradeAccountDataResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataResponseDefault)
}

// ToBuilder
func (m TradeAccountDataResponse) ToBuilder() TradeAccountDataResponseBuilder {
	return TradeAccountDataResponseBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradeAccountDataResponsePointer) ToBuilder() TradeAccountDataResponsePointerBuilder {
	return TradeAccountDataResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m TradeAccountDataResponseBuilder) Finish() TradeAccountDataResponse {
	return TradeAccountDataResponse{m.VLSBuilder.Finish()}
}

// Finish
func (m *TradeAccountDataResponsePointerBuilder) Finish() TradeAccountDataResponsePointer {
	return TradeAccountDataResponsePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataResponse) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataResponseBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataResponsePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataResponsePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccountNotExist
func (m TradeAccountDataResponse) TradeAccountNotExist() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// TradeAccountNotExist
func (m TradeAccountDataResponseBuilder) TradeAccountNotExist() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// TradeAccountNotExist
func (m TradeAccountDataResponsePointer) TradeAccountNotExist() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// TradeAccountNotExist
func (m TradeAccountDataResponsePointerBuilder) TradeAccountNotExist() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// TradeAccount
func (m TradeAccountDataResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// TradeAccount
func (m TradeAccountDataResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// TradeAccount
func (m TradeAccountDataResponsePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// TradeAccount
func (m TradeAccountDataResponsePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// IsSimulated
func (m TradeAccountDataResponse) IsSimulated() bool {
	return message.BoolVLS(m.Unsafe(), 16, 15)
}

// IsSimulated
func (m TradeAccountDataResponseBuilder) IsSimulated() bool {
	return message.BoolVLS(m.Unsafe(), 16, 15)
}

// IsSimulated
func (m TradeAccountDataResponsePointer) IsSimulated() bool {
	return message.BoolVLS(m.Ptr, 16, 15)
}

// IsSimulated
func (m TradeAccountDataResponsePointerBuilder) IsSimulated() bool {
	return message.BoolVLS(m.Ptr, 16, 15)
}

// CurrencyCode
func (m TradeAccountDataResponse) CurrencyCode() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// CurrencyCode
func (m TradeAccountDataResponseBuilder) CurrencyCode() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// CurrencyCode
func (m TradeAccountDataResponsePointer) CurrencyCode() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// CurrencyCode
func (m TradeAccountDataResponsePointerBuilder) CurrencyCode() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// CurrentCashBalance
func (m TradeAccountDataResponse) CurrentCashBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 28, 20)
}

// CurrentCashBalance
func (m TradeAccountDataResponseBuilder) CurrentCashBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 28, 20)
}

// CurrentCashBalance
func (m TradeAccountDataResponsePointer) CurrentCashBalance() float64 {
	return message.Float64VLS(m.Ptr, 28, 20)
}

// CurrentCashBalance
func (m TradeAccountDataResponsePointerBuilder) CurrentCashBalance() float64 {
	return message.Float64VLS(m.Ptr, 28, 20)
}

// AvailableFundsForNewPositions
func (m TradeAccountDataResponse) AvailableFundsForNewPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 36, 28)
}

// AvailableFundsForNewPositions
func (m TradeAccountDataResponseBuilder) AvailableFundsForNewPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 36, 28)
}

// AvailableFundsForNewPositions
func (m TradeAccountDataResponsePointer) AvailableFundsForNewPositions() float64 {
	return message.Float64VLS(m.Ptr, 36, 28)
}

// AvailableFundsForNewPositions
func (m TradeAccountDataResponsePointerBuilder) AvailableFundsForNewPositions() float64 {
	return message.Float64VLS(m.Ptr, 36, 28)
}

// MarginRequirement
func (m TradeAccountDataResponse) MarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 44, 36)
}

// MarginRequirement
func (m TradeAccountDataResponseBuilder) MarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 44, 36)
}

// MarginRequirement
func (m TradeAccountDataResponsePointer) MarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 44, 36)
}

// MarginRequirement
func (m TradeAccountDataResponsePointerBuilder) MarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 44, 36)
}

// AccountValue
func (m TradeAccountDataResponse) AccountValue() float64 {
	return message.Float64VLS(m.Unsafe(), 52, 44)
}

// AccountValue
func (m TradeAccountDataResponseBuilder) AccountValue() float64 {
	return message.Float64VLS(m.Unsafe(), 52, 44)
}

// AccountValue
func (m TradeAccountDataResponsePointer) AccountValue() float64 {
	return message.Float64VLS(m.Ptr, 52, 44)
}

// AccountValue
func (m TradeAccountDataResponsePointerBuilder) AccountValue() float64 {
	return message.Float64VLS(m.Ptr, 52, 44)
}

// OpenPositionsProfitLoss
func (m TradeAccountDataResponse) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 60, 52)
}

// OpenPositionsProfitLoss
func (m TradeAccountDataResponseBuilder) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 60, 52)
}

// OpenPositionsProfitLoss
func (m TradeAccountDataResponsePointer) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 60, 52)
}

// OpenPositionsProfitLoss
func (m TradeAccountDataResponsePointerBuilder) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 60, 52)
}

// DailyProfitLoss
func (m TradeAccountDataResponse) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 68, 60)
}

// DailyProfitLoss
func (m TradeAccountDataResponseBuilder) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 68, 60)
}

// DailyProfitLoss
func (m TradeAccountDataResponsePointer) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 68, 60)
}

// DailyProfitLoss
func (m TradeAccountDataResponsePointerBuilder) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 68, 60)
}

// TransactionIdentifierForCashBalanceAdjustment
func (m TradeAccountDataResponse) TransactionIdentifierForCashBalanceAdjustment() uint64 {
	return message.Uint64VLS(m.Unsafe(), 76, 68)
}

// TransactionIdentifierForCashBalanceAdjustment
func (m TradeAccountDataResponseBuilder) TransactionIdentifierForCashBalanceAdjustment() uint64 {
	return message.Uint64VLS(m.Unsafe(), 76, 68)
}

// TransactionIdentifierForCashBalanceAdjustment
func (m TradeAccountDataResponsePointer) TransactionIdentifierForCashBalanceAdjustment() uint64 {
	return message.Uint64VLS(m.Ptr, 76, 68)
}

// TransactionIdentifierForCashBalanceAdjustment
func (m TradeAccountDataResponsePointerBuilder) TransactionIdentifierForCashBalanceAdjustment() uint64 {
	return message.Uint64VLS(m.Ptr, 76, 68)
}

// LastTransactionDateTime
func (m TradeAccountDataResponse) LastTransactionDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 84, 76)
}

// LastTransactionDateTime
func (m TradeAccountDataResponseBuilder) LastTransactionDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 84, 76)
}

// LastTransactionDateTime
func (m TradeAccountDataResponsePointer) LastTransactionDateTime() int64 {
	return message.Int64VLS(m.Ptr, 84, 76)
}

// LastTransactionDateTime
func (m TradeAccountDataResponsePointerBuilder) LastTransactionDateTime() int64 {
	return message.Int64VLS(m.Ptr, 84, 76)
}

// TrailingAccountValueAtWhichToNotAllowNewPositions
func (m TradeAccountDataResponse) TrailingAccountValueAtWhichToNotAllowNewPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 92, 84)
}

// TrailingAccountValueAtWhichToNotAllowNewPositions
func (m TradeAccountDataResponseBuilder) TrailingAccountValueAtWhichToNotAllowNewPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 92, 84)
}

// TrailingAccountValueAtWhichToNotAllowNewPositions
func (m TradeAccountDataResponsePointer) TrailingAccountValueAtWhichToNotAllowNewPositions() float64 {
	return message.Float64VLS(m.Ptr, 92, 84)
}

// TrailingAccountValueAtWhichToNotAllowNewPositions
func (m TradeAccountDataResponsePointerBuilder) TrailingAccountValueAtWhichToNotAllowNewPositions() float64 {
	return message.Float64VLS(m.Ptr, 92, 84)
}

// CalculatedDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponse) CalculatedDailyNetLossLimitInAccountCurrency() float64 {
	return message.Float64VLS(m.Unsafe(), 100, 92)
}

// CalculatedDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponseBuilder) CalculatedDailyNetLossLimitInAccountCurrency() float64 {
	return message.Float64VLS(m.Unsafe(), 100, 92)
}

// CalculatedDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponsePointer) CalculatedDailyNetLossLimitInAccountCurrency() float64 {
	return message.Float64VLS(m.Ptr, 100, 92)
}

// CalculatedDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponsePointerBuilder) CalculatedDailyNetLossLimitInAccountCurrency() float64 {
	return message.Float64VLS(m.Ptr, 100, 92)
}

// DailyNetLossLimitHasBeenReached
func (m TradeAccountDataResponse) DailyNetLossLimitHasBeenReached() uint8 {
	return message.Uint8VLS(m.Unsafe(), 101, 100)
}

// DailyNetLossLimitHasBeenReached
func (m TradeAccountDataResponseBuilder) DailyNetLossLimitHasBeenReached() uint8 {
	return message.Uint8VLS(m.Unsafe(), 101, 100)
}

// DailyNetLossLimitHasBeenReached
func (m TradeAccountDataResponsePointer) DailyNetLossLimitHasBeenReached() uint8 {
	return message.Uint8VLS(m.Ptr, 101, 100)
}

// DailyNetLossLimitHasBeenReached
func (m TradeAccountDataResponsePointerBuilder) DailyNetLossLimitHasBeenReached() uint8 {
	return message.Uint8VLS(m.Ptr, 101, 100)
}

// LastResetDailyNetLossManagementVariablesDateTimeUTC
func (m TradeAccountDataResponse) LastResetDailyNetLossManagementVariablesDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 109, 101)
}

// LastResetDailyNetLossManagementVariablesDateTimeUTC
func (m TradeAccountDataResponseBuilder) LastResetDailyNetLossManagementVariablesDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 109, 101)
}

// LastResetDailyNetLossManagementVariablesDateTimeUTC
func (m TradeAccountDataResponsePointer) LastResetDailyNetLossManagementVariablesDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 109, 101)
}

// LastResetDailyNetLossManagementVariablesDateTimeUTC
func (m TradeAccountDataResponsePointerBuilder) LastResetDailyNetLossManagementVariablesDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 109, 101)
}

// IsUnderRequiredMargin
func (m TradeAccountDataResponse) IsUnderRequiredMargin() bool {
	return message.BoolVLS(m.Unsafe(), 110, 109)
}

// IsUnderRequiredMargin
func (m TradeAccountDataResponseBuilder) IsUnderRequiredMargin() bool {
	return message.BoolVLS(m.Unsafe(), 110, 109)
}

// IsUnderRequiredMargin
func (m TradeAccountDataResponsePointer) IsUnderRequiredMargin() bool {
	return message.BoolVLS(m.Ptr, 110, 109)
}

// IsUnderRequiredMargin
func (m TradeAccountDataResponsePointerBuilder) IsUnderRequiredMargin() bool {
	return message.BoolVLS(m.Ptr, 110, 109)
}

// DailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponse) DailyNetLossLimitInAccountCurrency() float32 {
	return message.Float32VLS(m.Unsafe(), 114, 110)
}

// DailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponseBuilder) DailyNetLossLimitInAccountCurrency() float32 {
	return message.Float32VLS(m.Unsafe(), 114, 110)
}

// DailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponsePointer) DailyNetLossLimitInAccountCurrency() float32 {
	return message.Float32VLS(m.Ptr, 114, 110)
}

// DailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponsePointerBuilder) DailyNetLossLimitInAccountCurrency() float32 {
	return message.Float32VLS(m.Ptr, 114, 110)
}

// PercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataResponse) PercentOfCashBalanceForDailyNetLossLimit() int32 {
	return message.Int32VLS(m.Unsafe(), 118, 114)
}

// PercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataResponseBuilder) PercentOfCashBalanceForDailyNetLossLimit() int32 {
	return message.Int32VLS(m.Unsafe(), 118, 114)
}

// PercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataResponsePointer) PercentOfCashBalanceForDailyNetLossLimit() int32 {
	return message.Int32VLS(m.Ptr, 118, 114)
}

// PercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataResponsePointerBuilder) PercentOfCashBalanceForDailyNetLossLimit() int32 {
	return message.Int32VLS(m.Ptr, 118, 114)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataResponse) UseTrailingAccountValueToNotAllowIncreaseInPositions() bool {
	return message.BoolVLS(m.Unsafe(), 119, 118)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataResponseBuilder) UseTrailingAccountValueToNotAllowIncreaseInPositions() bool {
	return message.BoolVLS(m.Unsafe(), 119, 118)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataResponsePointer) UseTrailingAccountValueToNotAllowIncreaseInPositions() bool {
	return message.BoolVLS(m.Ptr, 119, 118)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataResponsePointerBuilder) UseTrailingAccountValueToNotAllowIncreaseInPositions() bool {
	return message.BoolVLS(m.Ptr, 119, 118)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponse) DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 120, 119)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponseBuilder) DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 120, 119)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointer) DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 120, 119)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 120, 119)
}

// FlattenPositionsAtDailyLossLimit
func (m TradeAccountDataResponse) FlattenPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 121, 120)
}

// FlattenPositionsAtDailyLossLimit
func (m TradeAccountDataResponseBuilder) FlattenPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 121, 120)
}

// FlattenPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointer) FlattenPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 121, 120)
}

// FlattenPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) FlattenPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 121, 120)
}

// ClosePositionsAtEndOfDay
func (m TradeAccountDataResponse) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 122, 121)
}

// ClosePositionsAtEndOfDay
func (m TradeAccountDataResponseBuilder) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 122, 121)
}

// ClosePositionsAtEndOfDay
func (m TradeAccountDataResponsePointer) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 122, 121)
}

// ClosePositionsAtEndOfDay
func (m TradeAccountDataResponsePointerBuilder) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 122, 121)
}

// FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponse) FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 123, 122)
}

// FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponseBuilder) FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 123, 122)
}

// FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointer) FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 123, 122)
}

// FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointerBuilder) FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 123, 122)
}

// FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponse) FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 124, 123)
}

// FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponseBuilder) FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 124, 123)
}

// FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponsePointer) FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 124, 123)
}

// FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponsePointerBuilder) FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 124, 123)
}

// SenderSubID
func (m TradeAccountDataResponse) SenderSubID() string {
	return message.StringVLS(m.Unsafe(), 128, 124)
}

// SenderSubID
func (m TradeAccountDataResponseBuilder) SenderSubID() string {
	return message.StringVLS(m.Unsafe(), 128, 124)
}

// SenderSubID
func (m TradeAccountDataResponsePointer) SenderSubID() string {
	return message.StringVLS(m.Ptr, 128, 124)
}

// SenderSubID
func (m TradeAccountDataResponsePointerBuilder) SenderSubID() string {
	return message.StringVLS(m.Ptr, 128, 124)
}

// SenderLocationId
func (m TradeAccountDataResponse) SenderLocationId() string {
	return message.StringVLS(m.Unsafe(), 132, 128)
}

// SenderLocationId
func (m TradeAccountDataResponseBuilder) SenderLocationId() string {
	return message.StringVLS(m.Unsafe(), 132, 128)
}

// SenderLocationId
func (m TradeAccountDataResponsePointer) SenderLocationId() string {
	return message.StringVLS(m.Ptr, 132, 128)
}

// SenderLocationId
func (m TradeAccountDataResponsePointerBuilder) SenderLocationId() string {
	return message.StringVLS(m.Ptr, 132, 128)
}

// SelfMatchPreventionID
func (m TradeAccountDataResponse) SelfMatchPreventionID() string {
	return message.StringVLS(m.Unsafe(), 136, 132)
}

// SelfMatchPreventionID
func (m TradeAccountDataResponseBuilder) SelfMatchPreventionID() string {
	return message.StringVLS(m.Unsafe(), 136, 132)
}

// SelfMatchPreventionID
func (m TradeAccountDataResponsePointer) SelfMatchPreventionID() string {
	return message.StringVLS(m.Ptr, 136, 132)
}

// SelfMatchPreventionID
func (m TradeAccountDataResponsePointerBuilder) SelfMatchPreventionID() string {
	return message.StringVLS(m.Ptr, 136, 132)
}

// CTICode
func (m TradeAccountDataResponse) CTICode() int32 {
	return message.Int32VLS(m.Unsafe(), 140, 136)
}

// CTICode
func (m TradeAccountDataResponseBuilder) CTICode() int32 {
	return message.Int32VLS(m.Unsafe(), 140, 136)
}

// CTICode
func (m TradeAccountDataResponsePointer) CTICode() int32 {
	return message.Int32VLS(m.Ptr, 140, 136)
}

// CTICode
func (m TradeAccountDataResponsePointerBuilder) CTICode() int32 {
	return message.Int32VLS(m.Ptr, 140, 136)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataResponse) TradeAccountIsReadOnly() bool {
	return message.BoolVLS(m.Unsafe(), 141, 140)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataResponseBuilder) TradeAccountIsReadOnly() bool {
	return message.BoolVLS(m.Unsafe(), 141, 140)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataResponsePointer) TradeAccountIsReadOnly() bool {
	return message.BoolVLS(m.Ptr, 141, 140)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataResponsePointerBuilder) TradeAccountIsReadOnly() bool {
	return message.BoolVLS(m.Ptr, 141, 140)
}

// MaximumGlobalPositionQuantity
func (m TradeAccountDataResponse) MaximumGlobalPositionQuantity() int32 {
	return message.Int32VLS(m.Unsafe(), 145, 141)
}

// MaximumGlobalPositionQuantity
func (m TradeAccountDataResponseBuilder) MaximumGlobalPositionQuantity() int32 {
	return message.Int32VLS(m.Unsafe(), 145, 141)
}

// MaximumGlobalPositionQuantity
func (m TradeAccountDataResponsePointer) MaximumGlobalPositionQuantity() int32 {
	return message.Int32VLS(m.Ptr, 145, 141)
}

// MaximumGlobalPositionQuantity
func (m TradeAccountDataResponsePointerBuilder) MaximumGlobalPositionQuantity() int32 {
	return message.Int32VLS(m.Ptr, 145, 141)
}

// TradeFeePerContract
func (m TradeAccountDataResponse) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Unsafe(), 153, 145)
}

// TradeFeePerContract
func (m TradeAccountDataResponseBuilder) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Unsafe(), 153, 145)
}

// TradeFeePerContract
func (m TradeAccountDataResponsePointer) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Ptr, 153, 145)
}

// TradeFeePerContract
func (m TradeAccountDataResponsePointerBuilder) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Ptr, 153, 145)
}

// TradeFeePerShare
func (m TradeAccountDataResponse) TradeFeePerShare() float64 {
	return message.Float64VLS(m.Unsafe(), 161, 153)
}

// TradeFeePerShare
func (m TradeAccountDataResponseBuilder) TradeFeePerShare() float64 {
	return message.Float64VLS(m.Unsafe(), 161, 153)
}

// TradeFeePerShare
func (m TradeAccountDataResponsePointer) TradeFeePerShare() float64 {
	return message.Float64VLS(m.Ptr, 161, 153)
}

// TradeFeePerShare
func (m TradeAccountDataResponsePointerBuilder) TradeFeePerShare() float64 {
	return message.Float64VLS(m.Ptr, 161, 153)
}

// RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponse) RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 162, 161)
}

// RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponseBuilder) RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 162, 161)
}

// RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponsePointer) RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 162, 161)
}

// RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponsePointerBuilder) RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 162, 161)
}

// UsePercentOfMargin
func (m TradeAccountDataResponse) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Unsafe(), 166, 162)
}

// UsePercentOfMargin
func (m TradeAccountDataResponseBuilder) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Unsafe(), 166, 162)
}

// UsePercentOfMargin
func (m TradeAccountDataResponsePointer) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Ptr, 166, 162)
}

// UsePercentOfMargin
func (m TradeAccountDataResponsePointerBuilder) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Ptr, 166, 162)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponse) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Unsafe(), 170, 166)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponseBuilder) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Unsafe(), 170, 166)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponsePointer) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Ptr, 170, 166)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponsePointerBuilder) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Ptr, 170, 166)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponse) MaximumAllowedAccountBalanceForPositionsAsPercentage() int32 {
	return message.Int32VLS(m.Unsafe(), 174, 170)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponseBuilder) MaximumAllowedAccountBalanceForPositionsAsPercentage() int32 {
	return message.Int32VLS(m.Unsafe(), 174, 170)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponsePointer) MaximumAllowedAccountBalanceForPositionsAsPercentage() int32 {
	return message.Int32VLS(m.Ptr, 174, 170)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponsePointerBuilder) MaximumAllowedAccountBalanceForPositionsAsPercentage() int32 {
	return message.Int32VLS(m.Ptr, 174, 170)
}

// FirmID
func (m TradeAccountDataResponse) FirmID() string {
	return message.StringVLS(m.Unsafe(), 178, 174)
}

// FirmID
func (m TradeAccountDataResponseBuilder) FirmID() string {
	return message.StringVLS(m.Unsafe(), 178, 174)
}

// FirmID
func (m TradeAccountDataResponsePointer) FirmID() string {
	return message.StringVLS(m.Ptr, 178, 174)
}

// FirmID
func (m TradeAccountDataResponsePointerBuilder) FirmID() string {
	return message.StringVLS(m.Ptr, 178, 174)
}

// TradingIsDisabled
func (m TradeAccountDataResponse) TradingIsDisabled() bool {
	return message.BoolVLS(m.Unsafe(), 179, 178)
}

// TradingIsDisabled
func (m TradeAccountDataResponseBuilder) TradingIsDisabled() bool {
	return message.BoolVLS(m.Unsafe(), 179, 178)
}

// TradingIsDisabled
func (m TradeAccountDataResponsePointer) TradingIsDisabled() bool {
	return message.BoolVLS(m.Ptr, 179, 178)
}

// TradingIsDisabled
func (m TradeAccountDataResponsePointerBuilder) TradingIsDisabled() bool {
	return message.BoolVLS(m.Ptr, 179, 178)
}

// DescriptiveName
func (m TradeAccountDataResponse) DescriptiveName() string {
	return message.StringVLS(m.Unsafe(), 183, 179)
}

// DescriptiveName
func (m TradeAccountDataResponseBuilder) DescriptiveName() string {
	return message.StringVLS(m.Unsafe(), 183, 179)
}

// DescriptiveName
func (m TradeAccountDataResponsePointer) DescriptiveName() string {
	return message.StringVLS(m.Ptr, 183, 179)
}

// DescriptiveName
func (m TradeAccountDataResponsePointerBuilder) DescriptiveName() string {
	return message.StringVLS(m.Ptr, 183, 179)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataResponse) IsMasterFirmControlAccount() bool {
	return message.BoolVLS(m.Unsafe(), 184, 183)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataResponseBuilder) IsMasterFirmControlAccount() bool {
	return message.BoolVLS(m.Unsafe(), 184, 183)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataResponsePointer) IsMasterFirmControlAccount() bool {
	return message.BoolVLS(m.Ptr, 184, 183)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataResponsePointerBuilder) IsMasterFirmControlAccount() bool {
	return message.BoolVLS(m.Ptr, 184, 183)
}

// MinimumRequiredAccountValue
func (m TradeAccountDataResponse) MinimumRequiredAccountValue() float64 {
	return message.Float64VLS(m.Unsafe(), 192, 184)
}

// MinimumRequiredAccountValue
func (m TradeAccountDataResponseBuilder) MinimumRequiredAccountValue() float64 {
	return message.Float64VLS(m.Unsafe(), 192, 184)
}

// MinimumRequiredAccountValue
func (m TradeAccountDataResponsePointer) MinimumRequiredAccountValue() float64 {
	return message.Float64VLS(m.Ptr, 192, 184)
}

// MinimumRequiredAccountValue
func (m TradeAccountDataResponsePointerBuilder) MinimumRequiredAccountValue() float64 {
	return message.Float64VLS(m.Ptr, 192, 184)
}

// BeginTimeForDayMargin
func (m TradeAccountDataResponse) BeginTimeForDayMargin() int64 {
	return message.Int64VLS(m.Unsafe(), 200, 192)
}

// BeginTimeForDayMargin
func (m TradeAccountDataResponseBuilder) BeginTimeForDayMargin() int64 {
	return message.Int64VLS(m.Unsafe(), 200, 192)
}

// BeginTimeForDayMargin
func (m TradeAccountDataResponsePointer) BeginTimeForDayMargin() int64 {
	return message.Int64VLS(m.Ptr, 200, 192)
}

// BeginTimeForDayMargin
func (m TradeAccountDataResponsePointerBuilder) BeginTimeForDayMargin() int64 {
	return message.Int64VLS(m.Ptr, 200, 192)
}

// EndTimeForDayMargin
func (m TradeAccountDataResponse) EndTimeForDayMargin() int64 {
	return message.Int64VLS(m.Unsafe(), 208, 200)
}

// EndTimeForDayMargin
func (m TradeAccountDataResponseBuilder) EndTimeForDayMargin() int64 {
	return message.Int64VLS(m.Unsafe(), 208, 200)
}

// EndTimeForDayMargin
func (m TradeAccountDataResponsePointer) EndTimeForDayMargin() int64 {
	return message.Int64VLS(m.Ptr, 208, 200)
}

// EndTimeForDayMargin
func (m TradeAccountDataResponsePointerBuilder) EndTimeForDayMargin() int64 {
	return message.Int64VLS(m.Ptr, 208, 200)
}

// DayMarginTimeZone
func (m TradeAccountDataResponse) DayMarginTimeZone() string {
	return message.StringVLS(m.Unsafe(), 212, 208)
}

// DayMarginTimeZone
func (m TradeAccountDataResponseBuilder) DayMarginTimeZone() string {
	return message.StringVLS(m.Unsafe(), 212, 208)
}

// DayMarginTimeZone
func (m TradeAccountDataResponsePointer) DayMarginTimeZone() string {
	return message.StringVLS(m.Ptr, 212, 208)
}

// DayMarginTimeZone
func (m TradeAccountDataResponsePointerBuilder) DayMarginTimeZone() string {
	return message.StringVLS(m.Ptr, 212, 208)
}

// IsSnapshot
func (m TradeAccountDataResponse) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 213, 212)
}

// IsSnapshot
func (m TradeAccountDataResponseBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 213, 212)
}

// IsSnapshot
func (m TradeAccountDataResponsePointer) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 213, 212)
}

// IsSnapshot
func (m TradeAccountDataResponsePointerBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 213, 212)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponse) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 214, 213)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponseBuilder) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 214, 213)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponsePointer) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 214, 213)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponsePointerBuilder) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 214, 213)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponse) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 215, 214)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponseBuilder) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 215, 214)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponsePointer) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 215, 214)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponsePointerBuilder) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 215, 214)
}

// IsDeleted
func (m TradeAccountDataResponse) IsDeleted() bool {
	return message.BoolVLS(m.Unsafe(), 216, 215)
}

// IsDeleted
func (m TradeAccountDataResponseBuilder) IsDeleted() bool {
	return message.BoolVLS(m.Unsafe(), 216, 215)
}

// IsDeleted
func (m TradeAccountDataResponsePointer) IsDeleted() bool {
	return message.BoolVLS(m.Ptr, 216, 215)
}

// IsDeleted
func (m TradeAccountDataResponsePointerBuilder) IsDeleted() bool {
	return message.BoolVLS(m.Ptr, 216, 215)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponse) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() bool {
	return message.BoolVLS(m.Unsafe(), 217, 216)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponseBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() bool {
	return message.BoolVLS(m.Unsafe(), 217, 216)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointer) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() bool {
	return message.BoolVLS(m.Ptr, 217, 216)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() bool {
	return message.BoolVLS(m.Ptr, 217, 216)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponse) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() bool {
	return message.BoolVLS(m.Unsafe(), 218, 217)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponseBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() bool {
	return message.BoolVLS(m.Unsafe(), 218, 217)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponsePointer) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() bool {
	return message.BoolVLS(m.Ptr, 218, 217)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() bool {
	return message.BoolVLS(m.Ptr, 218, 217)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataResponse) UseMasterFirm_SymbolLimitsArray() bool {
	return message.BoolVLS(m.Unsafe(), 219, 218)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataResponseBuilder) UseMasterFirm_SymbolLimitsArray() bool {
	return message.BoolVLS(m.Unsafe(), 219, 218)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataResponsePointer) UseMasterFirm_SymbolLimitsArray() bool {
	return message.BoolVLS(m.Ptr, 219, 218)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_SymbolLimitsArray() bool {
	return message.BoolVLS(m.Ptr, 219, 218)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataResponse) UseMasterFirm_TradeFees() bool {
	return message.BoolVLS(m.Unsafe(), 220, 219)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataResponseBuilder) UseMasterFirm_TradeFees() bool {
	return message.BoolVLS(m.Unsafe(), 220, 219)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataResponsePointer) UseMasterFirm_TradeFees() bool {
	return message.BoolVLS(m.Ptr, 220, 219)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_TradeFees() bool {
	return message.BoolVLS(m.Ptr, 220, 219)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataResponse) UseMasterFirm_TradeFeePerShare() bool {
	return message.BoolVLS(m.Unsafe(), 221, 220)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataResponseBuilder) UseMasterFirm_TradeFeePerShare() bool {
	return message.BoolVLS(m.Unsafe(), 221, 220)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataResponsePointer) UseMasterFirm_TradeFeePerShare() bool {
	return message.BoolVLS(m.Ptr, 221, 220)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_TradeFeePerShare() bool {
	return message.BoolVLS(m.Ptr, 221, 220)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponse) UseMasterFirm_RequireSufficientMarginForNewPositions() bool {
	return message.BoolVLS(m.Unsafe(), 222, 221)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponseBuilder) UseMasterFirm_RequireSufficientMarginForNewPositions() bool {
	return message.BoolVLS(m.Unsafe(), 222, 221)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponsePointer) UseMasterFirm_RequireSufficientMarginForNewPositions() bool {
	return message.BoolVLS(m.Ptr, 222, 221)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_RequireSufficientMarginForNewPositions() bool {
	return message.BoolVLS(m.Ptr, 222, 221)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataResponse) UseMasterFirm_UsePercentOfMargin() bool {
	return message.BoolVLS(m.Unsafe(), 223, 222)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataResponseBuilder) UseMasterFirm_UsePercentOfMargin() bool {
	return message.BoolVLS(m.Unsafe(), 223, 222)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataResponsePointer) UseMasterFirm_UsePercentOfMargin() bool {
	return message.BoolVLS(m.Ptr, 223, 222)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_UsePercentOfMargin() bool {
	return message.BoolVLS(m.Ptr, 223, 222)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponse) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() bool {
	return message.BoolVLS(m.Unsafe(), 224, 223)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponseBuilder) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() bool {
	return message.BoolVLS(m.Unsafe(), 224, 223)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponsePointer) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() bool {
	return message.BoolVLS(m.Ptr, 224, 223)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() bool {
	return message.BoolVLS(m.Ptr, 224, 223)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataResponse) UseMasterFirm_MinimumRequiredAccountValue() bool {
	return message.BoolVLS(m.Unsafe(), 225, 224)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataResponseBuilder) UseMasterFirm_MinimumRequiredAccountValue() bool {
	return message.BoolVLS(m.Unsafe(), 225, 224)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataResponsePointer) UseMasterFirm_MinimumRequiredAccountValue() bool {
	return message.BoolVLS(m.Ptr, 225, 224)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_MinimumRequiredAccountValue() bool {
	return message.BoolVLS(m.Ptr, 225, 224)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataResponse) UseMasterFirm_MarginTimeSettings() bool {
	return message.BoolVLS(m.Unsafe(), 226, 225)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataResponseBuilder) UseMasterFirm_MarginTimeSettings() bool {
	return message.BoolVLS(m.Unsafe(), 226, 225)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataResponsePointer) UseMasterFirm_MarginTimeSettings() bool {
	return message.BoolVLS(m.Ptr, 226, 225)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_MarginTimeSettings() bool {
	return message.BoolVLS(m.Ptr, 226, 225)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataResponse) UseMasterFirm_TradingIsDisabled() bool {
	return message.BoolVLS(m.Unsafe(), 227, 226)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataResponseBuilder) UseMasterFirm_TradingIsDisabled() bool {
	return message.BoolVLS(m.Unsafe(), 227, 226)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataResponsePointer) UseMasterFirm_TradingIsDisabled() bool {
	return message.BoolVLS(m.Ptr, 227, 226)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_TradingIsDisabled() bool {
	return message.BoolVLS(m.Ptr, 227, 226)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataResponse) IsTradeStatisticsPublicallyShared() bool {
	return message.BoolVLS(m.Unsafe(), 228, 227)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataResponseBuilder) IsTradeStatisticsPublicallyShared() bool {
	return message.BoolVLS(m.Unsafe(), 228, 227)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataResponsePointer) IsTradeStatisticsPublicallyShared() bool {
	return message.BoolVLS(m.Ptr, 228, 227)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataResponsePointerBuilder) IsTradeStatisticsPublicallyShared() bool {
	return message.BoolVLS(m.Ptr, 228, 227)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataResponse) IsReadOnlyFollowingRequestsAllowed() bool {
	return message.BoolVLS(m.Unsafe(), 229, 228)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataResponseBuilder) IsReadOnlyFollowingRequestsAllowed() bool {
	return message.BoolVLS(m.Unsafe(), 229, 228)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataResponsePointer) IsReadOnlyFollowingRequestsAllowed() bool {
	return message.BoolVLS(m.Ptr, 229, 228)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataResponsePointerBuilder) IsReadOnlyFollowingRequestsAllowed() bool {
	return message.BoolVLS(m.Ptr, 229, 228)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataResponse) IsTradeAccountSharingAllowed() bool {
	return message.BoolVLS(m.Unsafe(), 230, 229)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataResponseBuilder) IsTradeAccountSharingAllowed() bool {
	return message.BoolVLS(m.Unsafe(), 230, 229)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataResponsePointer) IsTradeAccountSharingAllowed() bool {
	return message.BoolVLS(m.Ptr, 230, 229)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataResponsePointerBuilder) IsTradeAccountSharingAllowed() bool {
	return message.BoolVLS(m.Ptr, 230, 229)
}

// ReadOnlyShareToAllSCUsernames
func (m TradeAccountDataResponse) ReadOnlyShareToAllSCUsernames() uint8 {
	return message.Uint8VLS(m.Unsafe(), 231, 230)
}

// ReadOnlyShareToAllSCUsernames
func (m TradeAccountDataResponseBuilder) ReadOnlyShareToAllSCUsernames() uint8 {
	return message.Uint8VLS(m.Unsafe(), 231, 230)
}

// ReadOnlyShareToAllSCUsernames
func (m TradeAccountDataResponsePointer) ReadOnlyShareToAllSCUsernames() uint8 {
	return message.Uint8VLS(m.Ptr, 231, 230)
}

// ReadOnlyShareToAllSCUsernames
func (m TradeAccountDataResponsePointerBuilder) ReadOnlyShareToAllSCUsernames() uint8 {
	return message.Uint8VLS(m.Ptr, 231, 230)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataResponse) UseMasterFirm_SymbolCommissionsArray() bool {
	return message.BoolVLS(m.Unsafe(), 232, 231)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataResponseBuilder) UseMasterFirm_SymbolCommissionsArray() bool {
	return message.BoolVLS(m.Unsafe(), 232, 231)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataResponsePointer) UseMasterFirm_SymbolCommissionsArray() bool {
	return message.BoolVLS(m.Ptr, 232, 231)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_SymbolCommissionsArray() bool {
	return message.BoolVLS(m.Ptr, 232, 231)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponse) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() bool {
	return message.BoolVLS(m.Unsafe(), 233, 232)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponseBuilder) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() bool {
	return message.BoolVLS(m.Unsafe(), 233, 232)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointer) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() bool {
	return message.BoolVLS(m.Ptr, 233, 232)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() bool {
	return message.BoolVLS(m.Ptr, 233, 232)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponse) UseMasterFirm_UsePercentOfMarginForDayTrading() bool {
	return message.BoolVLS(m.Unsafe(), 234, 233)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponseBuilder) UseMasterFirm_UsePercentOfMarginForDayTrading() bool {
	return message.BoolVLS(m.Unsafe(), 234, 233)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponsePointer) UseMasterFirm_UsePercentOfMarginForDayTrading() bool {
	return message.BoolVLS(m.Ptr, 234, 233)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_UsePercentOfMarginForDayTrading() bool {
	return message.BoolVLS(m.Ptr, 234, 233)
}

// OpenPositionsProfitLossBasedOnSettlement
func (m TradeAccountDataResponse) OpenPositionsProfitLossBasedOnSettlement() float64 {
	return message.Float64VLS(m.Unsafe(), 242, 234)
}

// OpenPositionsProfitLossBasedOnSettlement
func (m TradeAccountDataResponseBuilder) OpenPositionsProfitLossBasedOnSettlement() float64 {
	return message.Float64VLS(m.Unsafe(), 242, 234)
}

// OpenPositionsProfitLossBasedOnSettlement
func (m TradeAccountDataResponsePointer) OpenPositionsProfitLossBasedOnSettlement() float64 {
	return message.Float64VLS(m.Ptr, 242, 234)
}

// OpenPositionsProfitLossBasedOnSettlement
func (m TradeAccountDataResponsePointerBuilder) OpenPositionsProfitLossBasedOnSettlement() float64 {
	return message.Float64VLS(m.Ptr, 242, 234)
}

// MarginRequirementFull
func (m TradeAccountDataResponse) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Unsafe(), 250, 242)
}

// MarginRequirementFull
func (m TradeAccountDataResponseBuilder) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Unsafe(), 250, 242)
}

// MarginRequirementFull
func (m TradeAccountDataResponsePointer) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Ptr, 250, 242)
}

// MarginRequirementFull
func (m TradeAccountDataResponsePointerBuilder) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Ptr, 250, 242)
}

// MarginRequirementFullPositionsOnly
func (m TradeAccountDataResponse) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Unsafe(), 258, 250)
}

// MarginRequirementFullPositionsOnly
func (m TradeAccountDataResponseBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Unsafe(), 258, 250)
}

// MarginRequirementFullPositionsOnly
func (m TradeAccountDataResponsePointer) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Ptr, 258, 250)
}

// MarginRequirementFullPositionsOnly
func (m TradeAccountDataResponsePointerBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Ptr, 258, 250)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataResponse) UseMasterFirm_TradeFeesFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 259, 258)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataResponseBuilder) UseMasterFirm_TradeFeesFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 259, 258)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataResponsePointer) UseMasterFirm_TradeFeesFullOverride() bool {
	return message.BoolVLS(m.Ptr, 259, 258)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_TradeFeesFullOverride() bool {
	return message.BoolVLS(m.Ptr, 259, 258)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataResponse) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() bool {
	return message.BoolVLS(m.Unsafe(), 260, 259)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataResponseBuilder) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() bool {
	return message.BoolVLS(m.Unsafe(), 260, 259)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataResponsePointer) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() bool {
	return message.BoolVLS(m.Ptr, 260, 259)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() bool {
	return message.BoolVLS(m.Ptr, 260, 259)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataResponse) UseMasterFirm_UsePercentOfMarginFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 261, 260)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataResponseBuilder) UseMasterFirm_UsePercentOfMarginFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 261, 260)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataResponsePointer) UseMasterFirm_UsePercentOfMarginFullOverride() bool {
	return message.BoolVLS(m.Ptr, 261, 260)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_UsePercentOfMarginFullOverride() bool {
	return message.BoolVLS(m.Ptr, 261, 260)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataResponse) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 262, 261)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataResponseBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 262, 261)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataResponsePointer) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() bool {
	return message.BoolVLS(m.Ptr, 262, 261)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataResponsePointerBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() bool {
	return message.BoolVLS(m.Ptr, 262, 261)
}

// PeakMarginRequirement
func (m TradeAccountDataResponse) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 270, 262)
}

// PeakMarginRequirement
func (m TradeAccountDataResponseBuilder) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 270, 262)
}

// PeakMarginRequirement
func (m TradeAccountDataResponsePointer) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 270, 262)
}

// PeakMarginRequirement
func (m TradeAccountDataResponsePointerBuilder) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 270, 262)
}

// LiquidationOnlyMode
func (m TradeAccountDataResponse) LiquidationOnlyMode() uint8 {
	return message.Uint8VLS(m.Unsafe(), 271, 270)
}

// LiquidationOnlyMode
func (m TradeAccountDataResponseBuilder) LiquidationOnlyMode() uint8 {
	return message.Uint8VLS(m.Unsafe(), 271, 270)
}

// LiquidationOnlyMode
func (m TradeAccountDataResponsePointer) LiquidationOnlyMode() uint8 {
	return message.Uint8VLS(m.Ptr, 271, 270)
}

// LiquidationOnlyMode
func (m TradeAccountDataResponsePointerBuilder) LiquidationOnlyMode() uint8 {
	return message.Uint8VLS(m.Ptr, 271, 270)
}

// FlattenPositionsWhenUnderMarginIntradayTriggered
func (m TradeAccountDataResponse) FlattenPositionsWhenUnderMarginIntradayTriggered() uint8 {
	return message.Uint8VLS(m.Unsafe(), 272, 271)
}

// FlattenPositionsWhenUnderMarginIntradayTriggered
func (m TradeAccountDataResponseBuilder) FlattenPositionsWhenUnderMarginIntradayTriggered() uint8 {
	return message.Uint8VLS(m.Unsafe(), 272, 271)
}

// FlattenPositionsWhenUnderMarginIntradayTriggered
func (m TradeAccountDataResponsePointer) FlattenPositionsWhenUnderMarginIntradayTriggered() uint8 {
	return message.Uint8VLS(m.Ptr, 272, 271)
}

// FlattenPositionsWhenUnderMarginIntradayTriggered
func (m TradeAccountDataResponsePointerBuilder) FlattenPositionsWhenUnderMarginIntradayTriggered() uint8 {
	return message.Uint8VLS(m.Ptr, 272, 271)
}

// FlattenPositionsWhenUnderMinimumAccountValueTriggered
func (m TradeAccountDataResponse) FlattenPositionsWhenUnderMinimumAccountValueTriggered() uint8 {
	return message.Uint8VLS(m.Unsafe(), 273, 272)
}

// FlattenPositionsWhenUnderMinimumAccountValueTriggered
func (m TradeAccountDataResponseBuilder) FlattenPositionsWhenUnderMinimumAccountValueTriggered() uint8 {
	return message.Uint8VLS(m.Unsafe(), 273, 272)
}

// FlattenPositionsWhenUnderMinimumAccountValueTriggered
func (m TradeAccountDataResponsePointer) FlattenPositionsWhenUnderMinimumAccountValueTriggered() uint8 {
	return message.Uint8VLS(m.Ptr, 273, 272)
}

// FlattenPositionsWhenUnderMinimumAccountValueTriggered
func (m TradeAccountDataResponsePointerBuilder) FlattenPositionsWhenUnderMinimumAccountValueTriggered() uint8 {
	return message.Uint8VLS(m.Ptr, 273, 272)
}

// AccountValueAtEndOfDayCaptureTime
func (m TradeAccountDataResponse) AccountValueAtEndOfDayCaptureTime() float64 {
	return message.Float64VLS(m.Unsafe(), 281, 273)
}

// AccountValueAtEndOfDayCaptureTime
func (m TradeAccountDataResponseBuilder) AccountValueAtEndOfDayCaptureTime() float64 {
	return message.Float64VLS(m.Unsafe(), 281, 273)
}

// AccountValueAtEndOfDayCaptureTime
func (m TradeAccountDataResponsePointer) AccountValueAtEndOfDayCaptureTime() float64 {
	return message.Float64VLS(m.Ptr, 281, 273)
}

// AccountValueAtEndOfDayCaptureTime
func (m TradeAccountDataResponsePointerBuilder) AccountValueAtEndOfDayCaptureTime() float64 {
	return message.Float64VLS(m.Ptr, 281, 273)
}

// EndOfDayCaptureTime
func (m TradeAccountDataResponse) EndOfDayCaptureTime() int64 {
	return message.Int64VLS(m.Unsafe(), 289, 281)
}

// EndOfDayCaptureTime
func (m TradeAccountDataResponseBuilder) EndOfDayCaptureTime() int64 {
	return message.Int64VLS(m.Unsafe(), 289, 281)
}

// EndOfDayCaptureTime
func (m TradeAccountDataResponsePointer) EndOfDayCaptureTime() int64 {
	return message.Int64VLS(m.Ptr, 289, 281)
}

// EndOfDayCaptureTime
func (m TradeAccountDataResponsePointerBuilder) EndOfDayCaptureTime() int64 {
	return message.Int64VLS(m.Ptr, 289, 281)
}

// CustomerOrFirm
func (m TradeAccountDataResponse) CustomerOrFirm() uint8 {
	return message.Uint8VLS(m.Unsafe(), 290, 289)
}

// CustomerOrFirm
func (m TradeAccountDataResponseBuilder) CustomerOrFirm() uint8 {
	return message.Uint8VLS(m.Unsafe(), 290, 289)
}

// CustomerOrFirm
func (m TradeAccountDataResponsePointer) CustomerOrFirm() uint8 {
	return message.Uint8VLS(m.Ptr, 290, 289)
}

// CustomerOrFirm
func (m TradeAccountDataResponsePointerBuilder) CustomerOrFirm() uint8 {
	return message.Uint8VLS(m.Ptr, 290, 289)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataResponse) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 291, 290)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataResponseBuilder) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 291, 290)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataResponsePointer) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() uint8 {
	return message.Uint8VLS(m.Ptr, 291, 290)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataResponsePointerBuilder) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() uint8 {
	return message.Uint8VLS(m.Ptr, 291, 290)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataResponse) MasterFirm_FlattenCancelWhenUnderMinimumAccountValue() uint8 {
	return message.Uint8VLS(m.Unsafe(), 292, 291)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataResponseBuilder) MasterFirm_FlattenCancelWhenUnderMinimumAccountValue() uint8 {
	return message.Uint8VLS(m.Unsafe(), 292, 291)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataResponsePointer) MasterFirm_FlattenCancelWhenUnderMinimumAccountValue() uint8 {
	return message.Uint8VLS(m.Ptr, 292, 291)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataResponsePointerBuilder) MasterFirm_FlattenCancelWhenUnderMinimumAccountValue() uint8 {
	return message.Uint8VLS(m.Ptr, 292, 291)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataResponse) MasterFirm_FlattenCancelWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 293, 292)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataResponseBuilder) MasterFirm_FlattenCancelWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 293, 292)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointer) MasterFirm_FlattenCancelWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 293, 292)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointerBuilder) MasterFirm_FlattenCancelWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 293, 292)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataResponse) MasterFirm_FlattenCancelWhenUnderMarginEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 294, 293)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataResponseBuilder) MasterFirm_FlattenCancelWhenUnderMarginEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 294, 293)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataResponsePointer) MasterFirm_FlattenCancelWhenUnderMarginEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 294, 293)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataResponsePointerBuilder) MasterFirm_FlattenCancelWhenUnderMarginEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 294, 293)
}

// MasterFirm_MaximumOrderQuantity
func (m TradeAccountDataResponse) MasterFirm_MaximumOrderQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 298, 294)
}

// MasterFirm_MaximumOrderQuantity
func (m TradeAccountDataResponseBuilder) MasterFirm_MaximumOrderQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 298, 294)
}

// MasterFirm_MaximumOrderQuantity
func (m TradeAccountDataResponsePointer) MasterFirm_MaximumOrderQuantity() uint32 {
	return message.Uint32VLS(m.Ptr, 298, 294)
}

// MasterFirm_MaximumOrderQuantity
func (m TradeAccountDataResponsePointerBuilder) MasterFirm_MaximumOrderQuantity() uint32 {
	return message.Uint32VLS(m.Ptr, 298, 294)
}

// LastTriggerDateTimeUTCForDailyLossLimit
func (m TradeAccountDataResponse) LastTriggerDateTimeUTCForDailyLossLimit() int64 {
	return message.Int64VLS(m.Unsafe(), 306, 298)
}

// LastTriggerDateTimeUTCForDailyLossLimit
func (m TradeAccountDataResponseBuilder) LastTriggerDateTimeUTCForDailyLossLimit() int64 {
	return message.Int64VLS(m.Unsafe(), 306, 298)
}

// LastTriggerDateTimeUTCForDailyLossLimit
func (m TradeAccountDataResponsePointer) LastTriggerDateTimeUTCForDailyLossLimit() int64 {
	return message.Int64VLS(m.Ptr, 306, 298)
}

// LastTriggerDateTimeUTCForDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) LastTriggerDateTimeUTCForDailyLossLimit() int64 {
	return message.Int64VLS(m.Ptr, 306, 298)
}

// OpenPositionsProfitLossIsDelayed
func (m TradeAccountDataResponse) OpenPositionsProfitLossIsDelayed() bool {
	return message.BoolVLS(m.Unsafe(), 307, 306)
}

// OpenPositionsProfitLossIsDelayed
func (m TradeAccountDataResponseBuilder) OpenPositionsProfitLossIsDelayed() bool {
	return message.BoolVLS(m.Unsafe(), 307, 306)
}

// OpenPositionsProfitLossIsDelayed
func (m TradeAccountDataResponsePointer) OpenPositionsProfitLossIsDelayed() bool {
	return message.BoolVLS(m.Ptr, 307, 306)
}

// OpenPositionsProfitLossIsDelayed
func (m TradeAccountDataResponsePointerBuilder) OpenPositionsProfitLossIsDelayed() bool {
	return message.BoolVLS(m.Ptr, 307, 306)
}

// ExchangeTraderId
func (m TradeAccountDataResponse) ExchangeTraderId() string {
	return message.StringVLS(m.Unsafe(), 311, 307)
}

// ExchangeTraderId
func (m TradeAccountDataResponseBuilder) ExchangeTraderId() string {
	return message.StringVLS(m.Unsafe(), 311, 307)
}

// ExchangeTraderId
func (m TradeAccountDataResponsePointer) ExchangeTraderId() string {
	return message.StringVLS(m.Ptr, 311, 307)
}

// ExchangeTraderId
func (m TradeAccountDataResponsePointerBuilder) ExchangeTraderId() string {
	return message.StringVLS(m.Ptr, 311, 307)
}

// SetRequestID
func (m TradeAccountDataResponseBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m TradeAccountDataResponsePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccountNotExist
func (m TradeAccountDataResponseBuilder) SetTradeAccountNotExist(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 11, 10, value)
}

// SetTradeAccountNotExist
func (m TradeAccountDataResponsePointerBuilder) SetTradeAccountNotExist(value uint8) {
	message.SetUint8VLS(m.Ptr, 11, 10, value)
}

// SetTradeAccount
func (m *TradeAccountDataResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 15, 11, value)
}

// SetTradeAccount
func (m *TradeAccountDataResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetIsSimulated
func (m TradeAccountDataResponseBuilder) SetIsSimulated(value bool) {
	message.SetBoolVLS(m.Unsafe(), 16, 15, value)
}

// SetIsSimulated
func (m TradeAccountDataResponsePointerBuilder) SetIsSimulated(value bool) {
	message.SetBoolVLS(m.Ptr, 16, 15, value)
}

// SetCurrencyCode
func (m *TradeAccountDataResponseBuilder) SetCurrencyCode(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetCurrencyCode
func (m *TradeAccountDataResponsePointerBuilder) SetCurrencyCode(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetCurrentCashBalance
func (m TradeAccountDataResponseBuilder) SetCurrentCashBalance(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 28, 20, value)
}

// SetCurrentCashBalance
func (m TradeAccountDataResponsePointerBuilder) SetCurrentCashBalance(value float64) {
	message.SetFloat64VLS(m.Ptr, 28, 20, value)
}

// SetAvailableFundsForNewPositions
func (m TradeAccountDataResponseBuilder) SetAvailableFundsForNewPositions(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 36, 28, value)
}

// SetAvailableFundsForNewPositions
func (m TradeAccountDataResponsePointerBuilder) SetAvailableFundsForNewPositions(value float64) {
	message.SetFloat64VLS(m.Ptr, 36, 28, value)
}

// SetMarginRequirement
func (m TradeAccountDataResponseBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 44, 36, value)
}

// SetMarginRequirement
func (m TradeAccountDataResponsePointerBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Ptr, 44, 36, value)
}

// SetAccountValue
func (m TradeAccountDataResponseBuilder) SetAccountValue(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 52, 44, value)
}

// SetAccountValue
func (m TradeAccountDataResponsePointerBuilder) SetAccountValue(value float64) {
	message.SetFloat64VLS(m.Ptr, 52, 44, value)
}

// SetOpenPositionsProfitLoss
func (m TradeAccountDataResponseBuilder) SetOpenPositionsProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 60, 52, value)
}

// SetOpenPositionsProfitLoss
func (m TradeAccountDataResponsePointerBuilder) SetOpenPositionsProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 60, 52, value)
}

// SetDailyProfitLoss
func (m TradeAccountDataResponseBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 68, 60, value)
}

// SetDailyProfitLoss
func (m TradeAccountDataResponsePointerBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 68, 60, value)
}

// SetTransactionIdentifierForCashBalanceAdjustment
func (m TradeAccountDataResponseBuilder) SetTransactionIdentifierForCashBalanceAdjustment(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 76, 68, value)
}

// SetTransactionIdentifierForCashBalanceAdjustment
func (m TradeAccountDataResponsePointerBuilder) SetTransactionIdentifierForCashBalanceAdjustment(value uint64) {
	message.SetUint64VLS(m.Ptr, 76, 68, value)
}

// SetLastTransactionDateTime
func (m TradeAccountDataResponseBuilder) SetLastTransactionDateTime(value int64) {
	message.SetInt64VLS(m.Unsafe(), 84, 76, value)
}

// SetLastTransactionDateTime
func (m TradeAccountDataResponsePointerBuilder) SetLastTransactionDateTime(value int64) {
	message.SetInt64VLS(m.Ptr, 84, 76, value)
}

// SetTrailingAccountValueAtWhichToNotAllowNewPositions
func (m TradeAccountDataResponseBuilder) SetTrailingAccountValueAtWhichToNotAllowNewPositions(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 92, 84, value)
}

// SetTrailingAccountValueAtWhichToNotAllowNewPositions
func (m TradeAccountDataResponsePointerBuilder) SetTrailingAccountValueAtWhichToNotAllowNewPositions(value float64) {
	message.SetFloat64VLS(m.Ptr, 92, 84, value)
}

// SetCalculatedDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponseBuilder) SetCalculatedDailyNetLossLimitInAccountCurrency(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 100, 92, value)
}

// SetCalculatedDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponsePointerBuilder) SetCalculatedDailyNetLossLimitInAccountCurrency(value float64) {
	message.SetFloat64VLS(m.Ptr, 100, 92, value)
}

// SetDailyNetLossLimitHasBeenReached
func (m TradeAccountDataResponseBuilder) SetDailyNetLossLimitHasBeenReached(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 101, 100, value)
}

// SetDailyNetLossLimitHasBeenReached
func (m TradeAccountDataResponsePointerBuilder) SetDailyNetLossLimitHasBeenReached(value uint8) {
	message.SetUint8VLS(m.Ptr, 101, 100, value)
}

// SetLastResetDailyNetLossManagementVariablesDateTimeUTC
func (m TradeAccountDataResponseBuilder) SetLastResetDailyNetLossManagementVariablesDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Unsafe(), 109, 101, value)
}

// SetLastResetDailyNetLossManagementVariablesDateTimeUTC
func (m TradeAccountDataResponsePointerBuilder) SetLastResetDailyNetLossManagementVariablesDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 109, 101, value)
}

// SetIsUnderRequiredMargin
func (m TradeAccountDataResponseBuilder) SetIsUnderRequiredMargin(value bool) {
	message.SetBoolVLS(m.Unsafe(), 110, 109, value)
}

// SetIsUnderRequiredMargin
func (m TradeAccountDataResponsePointerBuilder) SetIsUnderRequiredMargin(value bool) {
	message.SetBoolVLS(m.Ptr, 110, 109, value)
}

// SetDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponseBuilder) SetDailyNetLossLimitInAccountCurrency(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 114, 110, value)
}

// SetDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataResponsePointerBuilder) SetDailyNetLossLimitInAccountCurrency(value float32) {
	message.SetFloat32VLS(m.Ptr, 114, 110, value)
}

// SetPercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataResponseBuilder) SetPercentOfCashBalanceForDailyNetLossLimit(value int32) {
	message.SetInt32VLS(m.Unsafe(), 118, 114, value)
}

// SetPercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataResponsePointerBuilder) SetPercentOfCashBalanceForDailyNetLossLimit(value int32) {
	message.SetInt32VLS(m.Ptr, 118, 114, value)
}

// SetUseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataResponseBuilder) SetUseTrailingAccountValueToNotAllowIncreaseInPositions(value bool) {
	message.SetBoolVLS(m.Unsafe(), 119, 118, value)
}

// SetUseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataResponsePointerBuilder) SetUseTrailingAccountValueToNotAllowIncreaseInPositions(value bool) {
	message.SetBoolVLS(m.Ptr, 119, 118, value)
}

// SetDoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponseBuilder) SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 120, 119, value)
}

// SetDoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Ptr, 120, 119, value)
}

// SetFlattenPositionsAtDailyLossLimit
func (m TradeAccountDataResponseBuilder) SetFlattenPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 121, 120, value)
}

// SetFlattenPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) SetFlattenPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Ptr, 121, 120, value)
}

// SetClosePositionsAtEndOfDay
func (m TradeAccountDataResponseBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 122, 121, value)
}

// SetClosePositionsAtEndOfDay
func (m TradeAccountDataResponsePointerBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 122, 121, value)
}

// SetFlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponseBuilder) SetFlattenPositionsWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 123, 122, value)
}

// SetFlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointerBuilder) SetFlattenPositionsWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Ptr, 123, 122, value)
}

// SetFlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponseBuilder) SetFlattenPositionsWhenUnderMarginAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 124, 123, value)
}

// SetFlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponsePointerBuilder) SetFlattenPositionsWhenUnderMarginAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 124, 123, value)
}

// SetSenderSubID
func (m *TradeAccountDataResponseBuilder) SetSenderSubID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 128, 124, value)
}

// SetSenderSubID
func (m *TradeAccountDataResponsePointerBuilder) SetSenderSubID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 128, 124, value)
}

// SetSenderLocationId
func (m *TradeAccountDataResponseBuilder) SetSenderLocationId(value string) {
	message.SetStringVLS(&m.VLSBuilder, 132, 128, value)
}

// SetSenderLocationId
func (m *TradeAccountDataResponsePointerBuilder) SetSenderLocationId(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 132, 128, value)
}

// SetSelfMatchPreventionID
func (m *TradeAccountDataResponseBuilder) SetSelfMatchPreventionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 136, 132, value)
}

// SetSelfMatchPreventionID
func (m *TradeAccountDataResponsePointerBuilder) SetSelfMatchPreventionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 136, 132, value)
}

// SetCTICode
func (m TradeAccountDataResponseBuilder) SetCTICode(value int32) {
	message.SetInt32VLS(m.Unsafe(), 140, 136, value)
}

// SetCTICode
func (m TradeAccountDataResponsePointerBuilder) SetCTICode(value int32) {
	message.SetInt32VLS(m.Ptr, 140, 136, value)
}

// SetTradeAccountIsReadOnly
func (m TradeAccountDataResponseBuilder) SetTradeAccountIsReadOnly(value bool) {
	message.SetBoolVLS(m.Unsafe(), 141, 140, value)
}

// SetTradeAccountIsReadOnly
func (m TradeAccountDataResponsePointerBuilder) SetTradeAccountIsReadOnly(value bool) {
	message.SetBoolVLS(m.Ptr, 141, 140, value)
}

// SetMaximumGlobalPositionQuantity
func (m TradeAccountDataResponseBuilder) SetMaximumGlobalPositionQuantity(value int32) {
	message.SetInt32VLS(m.Unsafe(), 145, 141, value)
}

// SetMaximumGlobalPositionQuantity
func (m TradeAccountDataResponsePointerBuilder) SetMaximumGlobalPositionQuantity(value int32) {
	message.SetInt32VLS(m.Ptr, 145, 141, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataResponseBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 153, 145, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataResponsePointerBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Ptr, 153, 145, value)
}

// SetTradeFeePerShare
func (m TradeAccountDataResponseBuilder) SetTradeFeePerShare(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 161, 153, value)
}

// SetTradeFeePerShare
func (m TradeAccountDataResponsePointerBuilder) SetTradeFeePerShare(value float64) {
	message.SetFloat64VLS(m.Ptr, 161, 153, value)
}

// SetRequireSufficientMarginForNewPositions
func (m TradeAccountDataResponseBuilder) SetRequireSufficientMarginForNewPositions(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 162, 161, value)
}

// SetRequireSufficientMarginForNewPositions
func (m TradeAccountDataResponsePointerBuilder) SetRequireSufficientMarginForNewPositions(value uint8) {
	message.SetUint8VLS(m.Ptr, 162, 161, value)
}

// SetUsePercentOfMargin
func (m TradeAccountDataResponseBuilder) SetUsePercentOfMargin(value int32) {
	message.SetInt32VLS(m.Unsafe(), 166, 162, value)
}

// SetUsePercentOfMargin
func (m TradeAccountDataResponsePointerBuilder) SetUsePercentOfMargin(value int32) {
	message.SetInt32VLS(m.Ptr, 166, 162, value)
}

// SetUsePercentOfMarginForDayTrading
func (m TradeAccountDataResponseBuilder) SetUsePercentOfMarginForDayTrading(value int32) {
	message.SetInt32VLS(m.Unsafe(), 170, 166, value)
}

// SetUsePercentOfMarginForDayTrading
func (m TradeAccountDataResponsePointerBuilder) SetUsePercentOfMarginForDayTrading(value int32) {
	message.SetInt32VLS(m.Ptr, 170, 166, value)
}

// SetMaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponseBuilder) SetMaximumAllowedAccountBalanceForPositionsAsPercentage(value int32) {
	message.SetInt32VLS(m.Unsafe(), 174, 170, value)
}

// SetMaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponsePointerBuilder) SetMaximumAllowedAccountBalanceForPositionsAsPercentage(value int32) {
	message.SetInt32VLS(m.Ptr, 174, 170, value)
}

// SetFirmID
func (m *TradeAccountDataResponseBuilder) SetFirmID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 178, 174, value)
}

// SetFirmID
func (m *TradeAccountDataResponsePointerBuilder) SetFirmID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 178, 174, value)
}

// SetTradingIsDisabled
func (m TradeAccountDataResponseBuilder) SetTradingIsDisabled(value bool) {
	message.SetBoolVLS(m.Unsafe(), 179, 178, value)
}

// SetTradingIsDisabled
func (m TradeAccountDataResponsePointerBuilder) SetTradingIsDisabled(value bool) {
	message.SetBoolVLS(m.Ptr, 179, 178, value)
}

// SetDescriptiveName
func (m *TradeAccountDataResponseBuilder) SetDescriptiveName(value string) {
	message.SetStringVLS(&m.VLSBuilder, 183, 179, value)
}

// SetDescriptiveName
func (m *TradeAccountDataResponsePointerBuilder) SetDescriptiveName(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 183, 179, value)
}

// SetIsMasterFirmControlAccount
func (m TradeAccountDataResponseBuilder) SetIsMasterFirmControlAccount(value bool) {
	message.SetBoolVLS(m.Unsafe(), 184, 183, value)
}

// SetIsMasterFirmControlAccount
func (m TradeAccountDataResponsePointerBuilder) SetIsMasterFirmControlAccount(value bool) {
	message.SetBoolVLS(m.Ptr, 184, 183, value)
}

// SetMinimumRequiredAccountValue
func (m TradeAccountDataResponseBuilder) SetMinimumRequiredAccountValue(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 192, 184, value)
}

// SetMinimumRequiredAccountValue
func (m TradeAccountDataResponsePointerBuilder) SetMinimumRequiredAccountValue(value float64) {
	message.SetFloat64VLS(m.Ptr, 192, 184, value)
}

// SetBeginTimeForDayMargin
func (m TradeAccountDataResponseBuilder) SetBeginTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Unsafe(), 200, 192, value)
}

// SetBeginTimeForDayMargin
func (m TradeAccountDataResponsePointerBuilder) SetBeginTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Ptr, 200, 192, value)
}

// SetEndTimeForDayMargin
func (m TradeAccountDataResponseBuilder) SetEndTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Unsafe(), 208, 200, value)
}

// SetEndTimeForDayMargin
func (m TradeAccountDataResponsePointerBuilder) SetEndTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Ptr, 208, 200, value)
}

// SetDayMarginTimeZone
func (m *TradeAccountDataResponseBuilder) SetDayMarginTimeZone(value string) {
	message.SetStringVLS(&m.VLSBuilder, 212, 208, value)
}

// SetDayMarginTimeZone
func (m *TradeAccountDataResponsePointerBuilder) SetDayMarginTimeZone(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 212, 208, value)
}

// SetIsSnapshot
func (m TradeAccountDataResponseBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Unsafe(), 213, 212, value)
}

// SetIsSnapshot
func (m TradeAccountDataResponsePointerBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Ptr, 213, 212, value)
}

// SetIsFirstMessageInBatch
func (m TradeAccountDataResponseBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolVLS(m.Unsafe(), 214, 213, value)
}

// SetIsFirstMessageInBatch
func (m TradeAccountDataResponsePointerBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolVLS(m.Ptr, 214, 213, value)
}

// SetIsLastMessageInBatch
func (m TradeAccountDataResponseBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolVLS(m.Unsafe(), 215, 214, value)
}

// SetIsLastMessageInBatch
func (m TradeAccountDataResponsePointerBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolVLS(m.Ptr, 215, 214, value)
}

// SetIsDeleted
func (m TradeAccountDataResponseBuilder) SetIsDeleted(value bool) {
	message.SetBoolVLS(m.Unsafe(), 216, 215, value)
}

// SetIsDeleted
func (m TradeAccountDataResponsePointerBuilder) SetIsDeleted(value bool) {
	message.SetBoolVLS(m.Ptr, 216, 215, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(value bool) {
	message.SetBoolVLS(m.Unsafe(), 217, 216, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(value bool) {
	message.SetBoolVLS(m.Ptr, 217, 216, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(value bool) {
	message.SetBoolVLS(m.Unsafe(), 218, 217, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(value bool) {
	message.SetBoolVLS(m.Ptr, 218, 217, value)
}

// SetUseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_SymbolLimitsArray(value bool) {
	message.SetBoolVLS(m.Unsafe(), 219, 218, value)
}

// SetUseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_SymbolLimitsArray(value bool) {
	message.SetBoolVLS(m.Ptr, 219, 218, value)
}

// SetUseMasterFirm_TradeFees
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_TradeFees(value bool) {
	message.SetBoolVLS(m.Unsafe(), 220, 219, value)
}

// SetUseMasterFirm_TradeFees
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_TradeFees(value bool) {
	message.SetBoolVLS(m.Ptr, 220, 219, value)
}

// SetUseMasterFirm_TradeFeePerShare
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_TradeFeePerShare(value bool) {
	message.SetBoolVLS(m.Unsafe(), 221, 220, value)
}

// SetUseMasterFirm_TradeFeePerShare
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_TradeFeePerShare(value bool) {
	message.SetBoolVLS(m.Ptr, 221, 220, value)
}

// SetUseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_RequireSufficientMarginForNewPositions(value bool) {
	message.SetBoolVLS(m.Unsafe(), 222, 221, value)
}

// SetUseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_RequireSufficientMarginForNewPositions(value bool) {
	message.SetBoolVLS(m.Ptr, 222, 221, value)
}

// SetUseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_UsePercentOfMargin(value bool) {
	message.SetBoolVLS(m.Unsafe(), 223, 222, value)
}

// SetUseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_UsePercentOfMargin(value bool) {
	message.SetBoolVLS(m.Ptr, 223, 222, value)
}

// SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(value bool) {
	message.SetBoolVLS(m.Unsafe(), 224, 223, value)
}

// SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(value bool) {
	message.SetBoolVLS(m.Ptr, 224, 223, value)
}

// SetUseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_MinimumRequiredAccountValue(value bool) {
	message.SetBoolVLS(m.Unsafe(), 225, 224, value)
}

// SetUseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_MinimumRequiredAccountValue(value bool) {
	message.SetBoolVLS(m.Ptr, 225, 224, value)
}

// SetUseMasterFirm_MarginTimeSettings
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_MarginTimeSettings(value bool) {
	message.SetBoolVLS(m.Unsafe(), 226, 225, value)
}

// SetUseMasterFirm_MarginTimeSettings
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_MarginTimeSettings(value bool) {
	message.SetBoolVLS(m.Ptr, 226, 225, value)
}

// SetUseMasterFirm_TradingIsDisabled
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_TradingIsDisabled(value bool) {
	message.SetBoolVLS(m.Unsafe(), 227, 226, value)
}

// SetUseMasterFirm_TradingIsDisabled
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_TradingIsDisabled(value bool) {
	message.SetBoolVLS(m.Ptr, 227, 226, value)
}

// SetIsTradeStatisticsPublicallyShared
func (m TradeAccountDataResponseBuilder) SetIsTradeStatisticsPublicallyShared(value bool) {
	message.SetBoolVLS(m.Unsafe(), 228, 227, value)
}

// SetIsTradeStatisticsPublicallyShared
func (m TradeAccountDataResponsePointerBuilder) SetIsTradeStatisticsPublicallyShared(value bool) {
	message.SetBoolVLS(m.Ptr, 228, 227, value)
}

// SetIsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataResponseBuilder) SetIsReadOnlyFollowingRequestsAllowed(value bool) {
	message.SetBoolVLS(m.Unsafe(), 229, 228, value)
}

// SetIsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataResponsePointerBuilder) SetIsReadOnlyFollowingRequestsAllowed(value bool) {
	message.SetBoolVLS(m.Ptr, 229, 228, value)
}

// SetIsTradeAccountSharingAllowed
func (m TradeAccountDataResponseBuilder) SetIsTradeAccountSharingAllowed(value bool) {
	message.SetBoolVLS(m.Unsafe(), 230, 229, value)
}

// SetIsTradeAccountSharingAllowed
func (m TradeAccountDataResponsePointerBuilder) SetIsTradeAccountSharingAllowed(value bool) {
	message.SetBoolVLS(m.Ptr, 230, 229, value)
}

// SetReadOnlyShareToAllSCUsernames
func (m TradeAccountDataResponseBuilder) SetReadOnlyShareToAllSCUsernames(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 231, 230, value)
}

// SetReadOnlyShareToAllSCUsernames
func (m TradeAccountDataResponsePointerBuilder) SetReadOnlyShareToAllSCUsernames(value uint8) {
	message.SetUint8VLS(m.Ptr, 231, 230, value)
}

// SetUseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_SymbolCommissionsArray(value bool) {
	message.SetBoolVLS(m.Unsafe(), 232, 231, value)
}

// SetUseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_SymbolCommissionsArray(value bool) {
	message.SetBoolVLS(m.Ptr, 232, 231, value)
}

// SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(value bool) {
	message.SetBoolVLS(m.Unsafe(), 233, 232, value)
}

// SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(value bool) {
	message.SetBoolVLS(m.Ptr, 233, 232, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTrading(value bool) {
	message.SetBoolVLS(m.Unsafe(), 234, 233, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTrading(value bool) {
	message.SetBoolVLS(m.Ptr, 234, 233, value)
}

// SetOpenPositionsProfitLossBasedOnSettlement
func (m TradeAccountDataResponseBuilder) SetOpenPositionsProfitLossBasedOnSettlement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 242, 234, value)
}

// SetOpenPositionsProfitLossBasedOnSettlement
func (m TradeAccountDataResponsePointerBuilder) SetOpenPositionsProfitLossBasedOnSettlement(value float64) {
	message.SetFloat64VLS(m.Ptr, 242, 234, value)
}

// SetMarginRequirementFull
func (m TradeAccountDataResponseBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 250, 242, value)
}

// SetMarginRequirementFull
func (m TradeAccountDataResponsePointerBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64VLS(m.Ptr, 250, 242, value)
}

// SetMarginRequirementFullPositionsOnly
func (m TradeAccountDataResponseBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 258, 250, value)
}

// SetMarginRequirementFullPositionsOnly
func (m TradeAccountDataResponsePointerBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64VLS(m.Ptr, 258, 250, value)
}

// SetUseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_TradeFeesFullOverride(value bool) {
	message.SetBoolVLS(m.Unsafe(), 259, 258, value)
}

// SetUseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_TradeFeesFullOverride(value bool) {
	message.SetBoolVLS(m.Ptr, 259, 258, value)
}

// SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(value bool) {
	message.SetBoolVLS(m.Unsafe(), 260, 259, value)
}

// SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(value bool) {
	message.SetBoolVLS(m.Ptr, 260, 259, value)
}

// SetUseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_UsePercentOfMarginFullOverride(value bool) {
	message.SetBoolVLS(m.Unsafe(), 261, 260, value)
}

// SetUseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_UsePercentOfMarginFullOverride(value bool) {
	message.SetBoolVLS(m.Ptr, 261, 260, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataResponseBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(value bool) {
	message.SetBoolVLS(m.Unsafe(), 262, 261, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataResponsePointerBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(value bool) {
	message.SetBoolVLS(m.Ptr, 262, 261, value)
}

// SetPeakMarginRequirement
func (m TradeAccountDataResponseBuilder) SetPeakMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 270, 262, value)
}

// SetPeakMarginRequirement
func (m TradeAccountDataResponsePointerBuilder) SetPeakMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Ptr, 270, 262, value)
}

// SetLiquidationOnlyMode
func (m TradeAccountDataResponseBuilder) SetLiquidationOnlyMode(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 271, 270, value)
}

// SetLiquidationOnlyMode
func (m TradeAccountDataResponsePointerBuilder) SetLiquidationOnlyMode(value uint8) {
	message.SetUint8VLS(m.Ptr, 271, 270, value)
}

// SetFlattenPositionsWhenUnderMarginIntradayTriggered
func (m TradeAccountDataResponseBuilder) SetFlattenPositionsWhenUnderMarginIntradayTriggered(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 272, 271, value)
}

// SetFlattenPositionsWhenUnderMarginIntradayTriggered
func (m TradeAccountDataResponsePointerBuilder) SetFlattenPositionsWhenUnderMarginIntradayTriggered(value uint8) {
	message.SetUint8VLS(m.Ptr, 272, 271, value)
}

// SetFlattenPositionsWhenUnderMinimumAccountValueTriggered
func (m TradeAccountDataResponseBuilder) SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 273, 272, value)
}

// SetFlattenPositionsWhenUnderMinimumAccountValueTriggered
func (m TradeAccountDataResponsePointerBuilder) SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(value uint8) {
	message.SetUint8VLS(m.Ptr, 273, 272, value)
}

// SetAccountValueAtEndOfDayCaptureTime
func (m TradeAccountDataResponseBuilder) SetAccountValueAtEndOfDayCaptureTime(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 281, 273, value)
}

// SetAccountValueAtEndOfDayCaptureTime
func (m TradeAccountDataResponsePointerBuilder) SetAccountValueAtEndOfDayCaptureTime(value float64) {
	message.SetFloat64VLS(m.Ptr, 281, 273, value)
}

// SetEndOfDayCaptureTime
func (m TradeAccountDataResponseBuilder) SetEndOfDayCaptureTime(value int64) {
	message.SetInt64VLS(m.Unsafe(), 289, 281, value)
}

// SetEndOfDayCaptureTime
func (m TradeAccountDataResponsePointerBuilder) SetEndOfDayCaptureTime(value int64) {
	message.SetInt64VLS(m.Ptr, 289, 281, value)
}

// SetCustomerOrFirm
func (m TradeAccountDataResponseBuilder) SetCustomerOrFirm(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 290, 289, value)
}

// SetCustomerOrFirm
func (m TradeAccountDataResponsePointerBuilder) SetCustomerOrFirm(value uint8) {
	message.SetUint8VLS(m.Ptr, 290, 289, value)
}

// SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataResponseBuilder) SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 291, 290, value)
}

// SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataResponsePointerBuilder) SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(value uint8) {
	message.SetUint8VLS(m.Ptr, 291, 290, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataResponseBuilder) SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 292, 291, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataResponsePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(value uint8) {
	message.SetUint8VLS(m.Ptr, 292, 291, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataResponseBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 293, 292, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataResponsePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Ptr, 293, 292, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataResponseBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 294, 293, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataResponsePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 294, 293, value)
}

// SetMasterFirm_MaximumOrderQuantity
func (m TradeAccountDataResponseBuilder) SetMasterFirm_MaximumOrderQuantity(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 298, 294, value)
}

// SetMasterFirm_MaximumOrderQuantity
func (m TradeAccountDataResponsePointerBuilder) SetMasterFirm_MaximumOrderQuantity(value uint32) {
	message.SetUint32VLS(m.Ptr, 298, 294, value)
}

// SetLastTriggerDateTimeUTCForDailyLossLimit
func (m TradeAccountDataResponseBuilder) SetLastTriggerDateTimeUTCForDailyLossLimit(value int64) {
	message.SetInt64VLS(m.Unsafe(), 306, 298, value)
}

// SetLastTriggerDateTimeUTCForDailyLossLimit
func (m TradeAccountDataResponsePointerBuilder) SetLastTriggerDateTimeUTCForDailyLossLimit(value int64) {
	message.SetInt64VLS(m.Ptr, 306, 298, value)
}

// SetOpenPositionsProfitLossIsDelayed
func (m TradeAccountDataResponseBuilder) SetOpenPositionsProfitLossIsDelayed(value bool) {
	message.SetBoolVLS(m.Unsafe(), 307, 306, value)
}

// SetOpenPositionsProfitLossIsDelayed
func (m TradeAccountDataResponsePointerBuilder) SetOpenPositionsProfitLossIsDelayed(value bool) {
	message.SetBoolVLS(m.Ptr, 307, 306, value)
}

// SetExchangeTraderId
func (m *TradeAccountDataResponseBuilder) SetExchangeTraderId(value string) {
	message.SetStringVLS(&m.VLSBuilder, 311, 307, value)
}

// SetExchangeTraderId
func (m *TradeAccountDataResponsePointerBuilder) SetExchangeTraderId(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 311, 307, value)
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

// Copy
func (m TradeAccountDataResponsePointer) Copy(to TradeAccountDataResponseBuilder) {
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
func (m TradeAccountDataResponsePointerBuilder) Copy(to TradeAccountDataResponseBuilder) {
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
func (m TradeAccountDataResponsePointer) CopyPointer(to TradeAccountDataResponsePointerBuilder) {
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
func (m TradeAccountDataResponsePointerBuilder) CopyPointer(to TradeAccountDataResponsePointerBuilder) {
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
