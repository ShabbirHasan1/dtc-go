package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeAccountDataUpdateSize = 225

// {
//     Size                                                                    = TradeAccountDataUpdateSize  (225)
//     Type                                                                    = TRADE_ACCOUNT_DATA_UPDATE  (10117)
//     BaseSize                                                                = TradeAccountDataUpdateSize  (225)
//     RequestID                                                               = 0
//     IsNewAccount                                                            = false
//     NewAccountAuthorizedUsername                                            = ""
//     TradeAccount                                                            = ""
//     CurrencyCodeIsSet                                                       = false
//     CurrencyCode                                                            = ""
//     DailyNetLossLimitInAccountCurrencyIsSet                                 = false
//     DailyNetLossLimitInAccountCurrency                                      = 0
//     PercentOfCashBalanceForDailyNetLossLimitIsSet                           = false
//     PercentOfCashBalanceForDailyNetLossLimit                                = 0
//     UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet               = false
//     UseTrailingAccountValueToNotAllowIncreaseInPositions                    = false
//     DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet                      = false
//     DoNotAllowIncreaseInPositionsAtDailyLossLimit                           = 0
//     FlattenPositionsAtDailyLossLimitIsSet                                   = false
//     FlattenPositionsAtDailyLossLimit                                        = 0
//     ClosePositionsAtEndOfDayIsSet                                           = false
//     ClosePositionsAtEndOfDay                                                = 0
//     FlattenPositionsWhenUnderMarginIntradayIsSet                            = false
//     FlattenPositionsWhenUnderMarginIntraday                                 = 0
//     FlattenPositionsWhenUnderMarginAtEndOfDayIsSet                          = false
//     FlattenPositionsWhenUnderMarginAtEndOfDay                               = 0
//     SenderSubIDIsSet                                                        = false
//     SenderSubID                                                             = ""
//     SenderLocationIdIsSet                                                   = false
//     SenderLocationId                                                        = ""
//     SelfMatchPreventionIDIsSet                                              = false
//     SelfMatchPreventionID                                                   = ""
//     CTICodeIsSet                                                            = false
//     CTICode                                                                 = 0
//     TradeAccountIsReadOnlyIsSet                                             = false
//     TradeAccountIsReadOnly                                                  = false
//     MaximumGlobalPositionQuantityIsSet                                      = false
//     MaximumGlobalPositionQuantity                                           = 0
//     TradeFeePerContractIsSet                                                = false
//     TradeFeePerContract                                                     = 0
//     TradeFeePerShareIsSet                                                   = false
//     TradeFeePerShare                                                        = 0
//     RequireSufficientMarginForNewPositionsIsSet                             = false
//     RequireSufficientMarginForNewPositions                                  = 1
//     UsePercentOfMarginIsSet                                                 = false
//     UsePercentOfMargin                                                      = 100
//     UsePercentOfMarginForDayTradingIsSet                                    = false
//     UsePercentOfMarginForDayTrading                                         = 100
//     MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet               = false
//     MaximumAllowedAccountBalanceForPositionsAsPercentage                    = 100
//     FirmIDIsSet                                                             = false
//     FirmID                                                                  = ""
//     TradingIsDisabledIsSet                                                  = false
//     TradingIsDisabled                                                       = false
//     DescriptiveNameIsSet                                                    = false
//     DescriptiveName                                                         = ""
//     IsMasterFirmControlAccountIsSet                                         = false
//     IsMasterFirmControlAccount                                              = false
//     MinimumRequiredAccountValueIsSet                                        = false
//     MinimumRequiredAccountValue                                             = 0
//     BeginTimeForDayMarginIsSet                                              = false
//     BeginTimeForDayMargin                                                   = 0
//     EndTimeForDayMarginIsSet                                                = false
//     EndTimeForDayMargin                                                     = 0
//     DayMarginTimeZoneIsSet                                                  = false
//     DayMarginTimeZone                                                       = ""
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet              = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday                   = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet            = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay                 = false
//     UseMasterFirm_SymbolLimitsArrayIsSet                                    = false
//     UseMasterFirm_SymbolLimitsArray                                         = false
//     UseMasterFirm_TradeFeesIsSet                                            = false
//     UseMasterFirm_TradeFees                                                 = false
//     UseMasterFirm_TradeFeePerShareIsSet                                     = false
//     UseMasterFirm_TradeFeePerShare                                          = false
//     UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet               = false
//     UseMasterFirm_RequireSufficientMarginForNewPositions                    = false
//     UseMasterFirm_UsePercentOfMarginIsSet                                   = false
//     UseMasterFirm_UsePercentOfMargin                                        = false
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet = false
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage      = false
//     UseMasterFirm_MinimumRequiredAccountValueIsSet                          = false
//     UseMasterFirm_MinimumRequiredAccountValue                               = false
//     UseMasterFirm_MarginTimeSettingsIsSet                                   = false
//     UseMasterFirm_MarginTimeSettings                                        = false
//     UseMasterFirm_TradingIsDisabledIsSet                                    = false
//     UseMasterFirm_TradingIsDisabled                                         = false
//     IsTradeStatisticsPublicallySharedIsSet                                  = false
//     IsTradeStatisticsPublicallyShared                                       = false
//     IsReadOnlyFollowingRequestsAllowedIsSet                                 = false
//     IsReadOnlyFollowingRequestsAllowed                                      = false
//     IsTradeAccountSharingAllowedIsSet                                       = false
//     IsTradeAccountSharingAllowed                                            = false
//     ReadOnlyShareToAllSCUsernamesIsSet                                      = false
//     ReadOnlyShareToAllSCUsernames                                           = 0
//     UseMasterFirm_SymbolCommissionsArrayIsSet                               = false
//     UseMasterFirm_SymbolCommissionsArray                                    = false
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet        = false
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit             = false
//     UseMasterFirm_UsePercentOfMarginForDayTradingIsSet                      = false
//     UseMasterFirm_UsePercentOfMarginForDayTrading                           = false
//     UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet                   = false
//     UseMasterFirm_SymbolCommissionsArrayFullOverride                        = false
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet         = false
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders              = false
//     UseMasterFirm_UsePercentOfMarginFullOverrideIsSet                       = false
//     UseMasterFirm_UsePercentOfMarginFullOverride                            = false
//     UseMasterFirm_TradeFeesFullOverrideIsSet                                = false
//     UseMasterFirm_TradeFeesFullOverride                                     = false
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet          = false
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride               = false
//     LiquidationOnlyModeIsSet                                                = false
//     LiquidationOnlyMode                                                     = 0
//     CustomerOrFirmIsSet                                                     = false
//     CustomerOrFirm                                                          = 0
//     MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet               = false
//     MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet                    = 0
//     MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet               = false
//     MasterFirm_FlattenCancelWhenUnderMinimumAccountValue                    = 0
//     MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet                    = false
//     MasterFirm_FlattenCancelWhenUnderMarginIntraday                         = 0
//     MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet                    = false
//     MasterFirm_FlattenCancelWhenUnderMarginEndOfDay                         = 0
//     MasterFirm_MaximumOrderQuantityIsSet                                    = 0
//     MasterFirm_MaximumOrderQuantity                                         = 0
//     ExchangeTraderIdIsSet                                                   = false
//     ExchangeTraderId                                                        = ""
// }
var _TradeAccountDataUpdateDefault = []byte{225, 0, 133, 39, 225, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 100, 0, 0, 0, 0, 100, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataUpdate struct {
	message.VLS
}

type TradeAccountDataUpdateBuilder struct {
	message.VLSBuilder
}

type TradeAccountDataUpdatePointer struct {
	message.VLSPointer
}

type TradeAccountDataUpdatePointerBuilder struct {
	message.VLSPointerBuilder
}

func NewTradeAccountDataUpdateFrom(b []byte) TradeAccountDataUpdate {
	return TradeAccountDataUpdate{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataUpdate(b []byte) TradeAccountDataUpdate {
	return TradeAccountDataUpdate{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataUpdate() TradeAccountDataUpdateBuilder {
	a := TradeAccountDataUpdateBuilder{message.NewVLSBuilder(225)}
	a.Unsafe().SetBytes(0, _TradeAccountDataUpdateDefault)
	return a
}

func AllocTradeAccountDataUpdate() TradeAccountDataUpdatePointerBuilder {
	a := TradeAccountDataUpdatePointerBuilder{message.AllocVLSBuilder(225)}
	a.Ptr.SetBytes(0, _TradeAccountDataUpdateDefault)
	return a
}

func AllocTradeAccountDataUpdateFrom(b []byte) TradeAccountDataUpdatePointer {
	return TradeAccountDataUpdatePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                                                                    = TradeAccountDataUpdateSize  (225)
//     Type                                                                    = TRADE_ACCOUNT_DATA_UPDATE  (10117)
//     BaseSize                                                                = TradeAccountDataUpdateSize  (225)
//     RequestID                                                               = 0
//     IsNewAccount                                                            = false
//     NewAccountAuthorizedUsername                                            = ""
//     TradeAccount                                                            = ""
//     CurrencyCodeIsSet                                                       = false
//     CurrencyCode                                                            = ""
//     DailyNetLossLimitInAccountCurrencyIsSet                                 = false
//     DailyNetLossLimitInAccountCurrency                                      = 0
//     PercentOfCashBalanceForDailyNetLossLimitIsSet                           = false
//     PercentOfCashBalanceForDailyNetLossLimit                                = 0
//     UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet               = false
//     UseTrailingAccountValueToNotAllowIncreaseInPositions                    = false
//     DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet                      = false
//     DoNotAllowIncreaseInPositionsAtDailyLossLimit                           = 0
//     FlattenPositionsAtDailyLossLimitIsSet                                   = false
//     FlattenPositionsAtDailyLossLimit                                        = 0
//     ClosePositionsAtEndOfDayIsSet                                           = false
//     ClosePositionsAtEndOfDay                                                = 0
//     FlattenPositionsWhenUnderMarginIntradayIsSet                            = false
//     FlattenPositionsWhenUnderMarginIntraday                                 = 0
//     FlattenPositionsWhenUnderMarginAtEndOfDayIsSet                          = false
//     FlattenPositionsWhenUnderMarginAtEndOfDay                               = 0
//     SenderSubIDIsSet                                                        = false
//     SenderSubID                                                             = ""
//     SenderLocationIdIsSet                                                   = false
//     SenderLocationId                                                        = ""
//     SelfMatchPreventionIDIsSet                                              = false
//     SelfMatchPreventionID                                                   = ""
//     CTICodeIsSet                                                            = false
//     CTICode                                                                 = 0
//     TradeAccountIsReadOnlyIsSet                                             = false
//     TradeAccountIsReadOnly                                                  = false
//     MaximumGlobalPositionQuantityIsSet                                      = false
//     MaximumGlobalPositionQuantity                                           = 0
//     TradeFeePerContractIsSet                                                = false
//     TradeFeePerContract                                                     = 0
//     TradeFeePerShareIsSet                                                   = false
//     TradeFeePerShare                                                        = 0
//     RequireSufficientMarginForNewPositionsIsSet                             = false
//     RequireSufficientMarginForNewPositions                                  = 1
//     UsePercentOfMarginIsSet                                                 = false
//     UsePercentOfMargin                                                      = 100
//     UsePercentOfMarginForDayTradingIsSet                                    = false
//     UsePercentOfMarginForDayTrading                                         = 100
//     MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet               = false
//     MaximumAllowedAccountBalanceForPositionsAsPercentage                    = 100
//     FirmIDIsSet                                                             = false
//     FirmID                                                                  = ""
//     TradingIsDisabledIsSet                                                  = false
//     TradingIsDisabled                                                       = false
//     DescriptiveNameIsSet                                                    = false
//     DescriptiveName                                                         = ""
//     IsMasterFirmControlAccountIsSet                                         = false
//     IsMasterFirmControlAccount                                              = false
//     MinimumRequiredAccountValueIsSet                                        = false
//     MinimumRequiredAccountValue                                             = 0
//     BeginTimeForDayMarginIsSet                                              = false
//     BeginTimeForDayMargin                                                   = 0
//     EndTimeForDayMarginIsSet                                                = false
//     EndTimeForDayMargin                                                     = 0
//     DayMarginTimeZoneIsSet                                                  = false
//     DayMarginTimeZone                                                       = ""
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet              = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday                   = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet            = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay                 = false
//     UseMasterFirm_SymbolLimitsArrayIsSet                                    = false
//     UseMasterFirm_SymbolLimitsArray                                         = false
//     UseMasterFirm_TradeFeesIsSet                                            = false
//     UseMasterFirm_TradeFees                                                 = false
//     UseMasterFirm_TradeFeePerShareIsSet                                     = false
//     UseMasterFirm_TradeFeePerShare                                          = false
//     UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet               = false
//     UseMasterFirm_RequireSufficientMarginForNewPositions                    = false
//     UseMasterFirm_UsePercentOfMarginIsSet                                   = false
//     UseMasterFirm_UsePercentOfMargin                                        = false
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet = false
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage      = false
//     UseMasterFirm_MinimumRequiredAccountValueIsSet                          = false
//     UseMasterFirm_MinimumRequiredAccountValue                               = false
//     UseMasterFirm_MarginTimeSettingsIsSet                                   = false
//     UseMasterFirm_MarginTimeSettings                                        = false
//     UseMasterFirm_TradingIsDisabledIsSet                                    = false
//     UseMasterFirm_TradingIsDisabled                                         = false
//     IsTradeStatisticsPublicallySharedIsSet                                  = false
//     IsTradeStatisticsPublicallyShared                                       = false
//     IsReadOnlyFollowingRequestsAllowedIsSet                                 = false
//     IsReadOnlyFollowingRequestsAllowed                                      = false
//     IsTradeAccountSharingAllowedIsSet                                       = false
//     IsTradeAccountSharingAllowed                                            = false
//     ReadOnlyShareToAllSCUsernamesIsSet                                      = false
//     ReadOnlyShareToAllSCUsernames                                           = 0
//     UseMasterFirm_SymbolCommissionsArrayIsSet                               = false
//     UseMasterFirm_SymbolCommissionsArray                                    = false
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet        = false
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit             = false
//     UseMasterFirm_UsePercentOfMarginForDayTradingIsSet                      = false
//     UseMasterFirm_UsePercentOfMarginForDayTrading                           = false
//     UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet                   = false
//     UseMasterFirm_SymbolCommissionsArrayFullOverride                        = false
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet         = false
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders              = false
//     UseMasterFirm_UsePercentOfMarginFullOverrideIsSet                       = false
//     UseMasterFirm_UsePercentOfMarginFullOverride                            = false
//     UseMasterFirm_TradeFeesFullOverrideIsSet                                = false
//     UseMasterFirm_TradeFeesFullOverride                                     = false
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet          = false
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride               = false
//     LiquidationOnlyModeIsSet                                                = false
//     LiquidationOnlyMode                                                     = 0
//     CustomerOrFirmIsSet                                                     = false
//     CustomerOrFirm                                                          = 0
//     MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet               = false
//     MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet                    = 0
//     MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet               = false
//     MasterFirm_FlattenCancelWhenUnderMinimumAccountValue                    = 0
//     MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet                    = false
//     MasterFirm_FlattenCancelWhenUnderMarginIntraday                         = 0
//     MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet                    = false
//     MasterFirm_FlattenCancelWhenUnderMarginEndOfDay                         = 0
//     MasterFirm_MaximumOrderQuantityIsSet                                    = 0
//     MasterFirm_MaximumOrderQuantity                                         = 0
//     ExchangeTraderIdIsSet                                                   = false
//     ExchangeTraderId                                                        = ""
// }
func (m TradeAccountDataUpdateBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataUpdateDefault)
}

// Clear
// {
//     Size                                                                    = TradeAccountDataUpdateSize  (225)
//     Type                                                                    = TRADE_ACCOUNT_DATA_UPDATE  (10117)
//     BaseSize                                                                = TradeAccountDataUpdateSize  (225)
//     RequestID                                                               = 0
//     IsNewAccount                                                            = false
//     NewAccountAuthorizedUsername                                            = ""
//     TradeAccount                                                            = ""
//     CurrencyCodeIsSet                                                       = false
//     CurrencyCode                                                            = ""
//     DailyNetLossLimitInAccountCurrencyIsSet                                 = false
//     DailyNetLossLimitInAccountCurrency                                      = 0
//     PercentOfCashBalanceForDailyNetLossLimitIsSet                           = false
//     PercentOfCashBalanceForDailyNetLossLimit                                = 0
//     UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet               = false
//     UseTrailingAccountValueToNotAllowIncreaseInPositions                    = false
//     DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet                      = false
//     DoNotAllowIncreaseInPositionsAtDailyLossLimit                           = 0
//     FlattenPositionsAtDailyLossLimitIsSet                                   = false
//     FlattenPositionsAtDailyLossLimit                                        = 0
//     ClosePositionsAtEndOfDayIsSet                                           = false
//     ClosePositionsAtEndOfDay                                                = 0
//     FlattenPositionsWhenUnderMarginIntradayIsSet                            = false
//     FlattenPositionsWhenUnderMarginIntraday                                 = 0
//     FlattenPositionsWhenUnderMarginAtEndOfDayIsSet                          = false
//     FlattenPositionsWhenUnderMarginAtEndOfDay                               = 0
//     SenderSubIDIsSet                                                        = false
//     SenderSubID                                                             = ""
//     SenderLocationIdIsSet                                                   = false
//     SenderLocationId                                                        = ""
//     SelfMatchPreventionIDIsSet                                              = false
//     SelfMatchPreventionID                                                   = ""
//     CTICodeIsSet                                                            = false
//     CTICode                                                                 = 0
//     TradeAccountIsReadOnlyIsSet                                             = false
//     TradeAccountIsReadOnly                                                  = false
//     MaximumGlobalPositionQuantityIsSet                                      = false
//     MaximumGlobalPositionQuantity                                           = 0
//     TradeFeePerContractIsSet                                                = false
//     TradeFeePerContract                                                     = 0
//     TradeFeePerShareIsSet                                                   = false
//     TradeFeePerShare                                                        = 0
//     RequireSufficientMarginForNewPositionsIsSet                             = false
//     RequireSufficientMarginForNewPositions                                  = 1
//     UsePercentOfMarginIsSet                                                 = false
//     UsePercentOfMargin                                                      = 100
//     UsePercentOfMarginForDayTradingIsSet                                    = false
//     UsePercentOfMarginForDayTrading                                         = 100
//     MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet               = false
//     MaximumAllowedAccountBalanceForPositionsAsPercentage                    = 100
//     FirmIDIsSet                                                             = false
//     FirmID                                                                  = ""
//     TradingIsDisabledIsSet                                                  = false
//     TradingIsDisabled                                                       = false
//     DescriptiveNameIsSet                                                    = false
//     DescriptiveName                                                         = ""
//     IsMasterFirmControlAccountIsSet                                         = false
//     IsMasterFirmControlAccount                                              = false
//     MinimumRequiredAccountValueIsSet                                        = false
//     MinimumRequiredAccountValue                                             = 0
//     BeginTimeForDayMarginIsSet                                              = false
//     BeginTimeForDayMargin                                                   = 0
//     EndTimeForDayMarginIsSet                                                = false
//     EndTimeForDayMargin                                                     = 0
//     DayMarginTimeZoneIsSet                                                  = false
//     DayMarginTimeZone                                                       = ""
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet              = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday                   = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet            = false
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay                 = false
//     UseMasterFirm_SymbolLimitsArrayIsSet                                    = false
//     UseMasterFirm_SymbolLimitsArray                                         = false
//     UseMasterFirm_TradeFeesIsSet                                            = false
//     UseMasterFirm_TradeFees                                                 = false
//     UseMasterFirm_TradeFeePerShareIsSet                                     = false
//     UseMasterFirm_TradeFeePerShare                                          = false
//     UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet               = false
//     UseMasterFirm_RequireSufficientMarginForNewPositions                    = false
//     UseMasterFirm_UsePercentOfMarginIsSet                                   = false
//     UseMasterFirm_UsePercentOfMargin                                        = false
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet = false
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage      = false
//     UseMasterFirm_MinimumRequiredAccountValueIsSet                          = false
//     UseMasterFirm_MinimumRequiredAccountValue                               = false
//     UseMasterFirm_MarginTimeSettingsIsSet                                   = false
//     UseMasterFirm_MarginTimeSettings                                        = false
//     UseMasterFirm_TradingIsDisabledIsSet                                    = false
//     UseMasterFirm_TradingIsDisabled                                         = false
//     IsTradeStatisticsPublicallySharedIsSet                                  = false
//     IsTradeStatisticsPublicallyShared                                       = false
//     IsReadOnlyFollowingRequestsAllowedIsSet                                 = false
//     IsReadOnlyFollowingRequestsAllowed                                      = false
//     IsTradeAccountSharingAllowedIsSet                                       = false
//     IsTradeAccountSharingAllowed                                            = false
//     ReadOnlyShareToAllSCUsernamesIsSet                                      = false
//     ReadOnlyShareToAllSCUsernames                                           = 0
//     UseMasterFirm_SymbolCommissionsArrayIsSet                               = false
//     UseMasterFirm_SymbolCommissionsArray                                    = false
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet        = false
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit             = false
//     UseMasterFirm_UsePercentOfMarginForDayTradingIsSet                      = false
//     UseMasterFirm_UsePercentOfMarginForDayTrading                           = false
//     UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet                   = false
//     UseMasterFirm_SymbolCommissionsArrayFullOverride                        = false
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet         = false
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders              = false
//     UseMasterFirm_UsePercentOfMarginFullOverrideIsSet                       = false
//     UseMasterFirm_UsePercentOfMarginFullOverride                            = false
//     UseMasterFirm_TradeFeesFullOverrideIsSet                                = false
//     UseMasterFirm_TradeFeesFullOverride                                     = false
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet          = false
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride               = false
//     LiquidationOnlyModeIsSet                                                = false
//     LiquidationOnlyMode                                                     = 0
//     CustomerOrFirmIsSet                                                     = false
//     CustomerOrFirm                                                          = 0
//     MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet               = false
//     MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet                    = 0
//     MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet               = false
//     MasterFirm_FlattenCancelWhenUnderMinimumAccountValue                    = 0
//     MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet                    = false
//     MasterFirm_FlattenCancelWhenUnderMarginIntraday                         = 0
//     MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet                    = false
//     MasterFirm_FlattenCancelWhenUnderMarginEndOfDay                         = 0
//     MasterFirm_MaximumOrderQuantityIsSet                                    = 0
//     MasterFirm_MaximumOrderQuantity                                         = 0
//     ExchangeTraderIdIsSet                                                   = false
//     ExchangeTraderId                                                        = ""
// }
func (m TradeAccountDataUpdatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataUpdateDefault)
}

// ToBuilder
func (m TradeAccountDataUpdate) ToBuilder() TradeAccountDataUpdateBuilder {
	return TradeAccountDataUpdateBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradeAccountDataUpdatePointer) ToBuilder() TradeAccountDataUpdatePointerBuilder {
	return TradeAccountDataUpdatePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m TradeAccountDataUpdateBuilder) Finish() TradeAccountDataUpdate {
	return TradeAccountDataUpdate{m.VLSBuilder.Finish()}
}

// Finish
func (m *TradeAccountDataUpdatePointerBuilder) Finish() TradeAccountDataUpdatePointer {
	return TradeAccountDataUpdatePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataUpdate) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataUpdateBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataUpdatePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataUpdatePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// IsNewAccount
func (m TradeAccountDataUpdate) IsNewAccount() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsNewAccount
func (m TradeAccountDataUpdateBuilder) IsNewAccount() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsNewAccount
func (m TradeAccountDataUpdatePointer) IsNewAccount() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// IsNewAccount
func (m TradeAccountDataUpdatePointerBuilder) IsNewAccount() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// NewAccountAuthorizedUsername
func (m TradeAccountDataUpdate) NewAccountAuthorizedUsername() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// NewAccountAuthorizedUsername
func (m TradeAccountDataUpdateBuilder) NewAccountAuthorizedUsername() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// NewAccountAuthorizedUsername
func (m TradeAccountDataUpdatePointer) NewAccountAuthorizedUsername() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// NewAccountAuthorizedUsername
func (m TradeAccountDataUpdatePointerBuilder) NewAccountAuthorizedUsername() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// TradeAccount
func (m TradeAccountDataUpdate) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
}

// TradeAccount
func (m TradeAccountDataUpdateBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
}

// TradeAccount
func (m TradeAccountDataUpdatePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 19, 15)
}

// TradeAccount
func (m TradeAccountDataUpdatePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 19, 15)
}

// CurrencyCodeIsSet
func (m TradeAccountDataUpdate) CurrencyCodeIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 20, 19)
}

// CurrencyCodeIsSet
func (m TradeAccountDataUpdateBuilder) CurrencyCodeIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 20, 19)
}

// CurrencyCodeIsSet
func (m TradeAccountDataUpdatePointer) CurrencyCodeIsSet() bool {
	return message.BoolVLS(m.Ptr, 20, 19)
}

// CurrencyCodeIsSet
func (m TradeAccountDataUpdatePointerBuilder) CurrencyCodeIsSet() bool {
	return message.BoolVLS(m.Ptr, 20, 19)
}

// CurrencyCode
func (m TradeAccountDataUpdate) CurrencyCode() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// CurrencyCode
func (m TradeAccountDataUpdateBuilder) CurrencyCode() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// CurrencyCode
func (m TradeAccountDataUpdatePointer) CurrencyCode() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// CurrencyCode
func (m TradeAccountDataUpdatePointerBuilder) CurrencyCode() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// DailyNetLossLimitInAccountCurrencyIsSet
func (m TradeAccountDataUpdate) DailyNetLossLimitInAccountCurrencyIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 25, 24)
}

// DailyNetLossLimitInAccountCurrencyIsSet
func (m TradeAccountDataUpdateBuilder) DailyNetLossLimitInAccountCurrencyIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 25, 24)
}

// DailyNetLossLimitInAccountCurrencyIsSet
func (m TradeAccountDataUpdatePointer) DailyNetLossLimitInAccountCurrencyIsSet() bool {
	return message.BoolVLS(m.Ptr, 25, 24)
}

// DailyNetLossLimitInAccountCurrencyIsSet
func (m TradeAccountDataUpdatePointerBuilder) DailyNetLossLimitInAccountCurrencyIsSet() bool {
	return message.BoolVLS(m.Ptr, 25, 24)
}

// DailyNetLossLimitInAccountCurrency
func (m TradeAccountDataUpdate) DailyNetLossLimitInAccountCurrency() float32 {
	return message.Float32VLS(m.Unsafe(), 29, 25)
}

// DailyNetLossLimitInAccountCurrency
func (m TradeAccountDataUpdateBuilder) DailyNetLossLimitInAccountCurrency() float32 {
	return message.Float32VLS(m.Unsafe(), 29, 25)
}

// DailyNetLossLimitInAccountCurrency
func (m TradeAccountDataUpdatePointer) DailyNetLossLimitInAccountCurrency() float32 {
	return message.Float32VLS(m.Ptr, 29, 25)
}

// DailyNetLossLimitInAccountCurrency
func (m TradeAccountDataUpdatePointerBuilder) DailyNetLossLimitInAccountCurrency() float32 {
	return message.Float32VLS(m.Ptr, 29, 25)
}

// PercentOfCashBalanceForDailyNetLossLimitIsSet
func (m TradeAccountDataUpdate) PercentOfCashBalanceForDailyNetLossLimitIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 30, 29)
}

// PercentOfCashBalanceForDailyNetLossLimitIsSet
func (m TradeAccountDataUpdateBuilder) PercentOfCashBalanceForDailyNetLossLimitIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 30, 29)
}

// PercentOfCashBalanceForDailyNetLossLimitIsSet
func (m TradeAccountDataUpdatePointer) PercentOfCashBalanceForDailyNetLossLimitIsSet() bool {
	return message.BoolVLS(m.Ptr, 30, 29)
}

// PercentOfCashBalanceForDailyNetLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) PercentOfCashBalanceForDailyNetLossLimitIsSet() bool {
	return message.BoolVLS(m.Ptr, 30, 29)
}

// PercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataUpdate) PercentOfCashBalanceForDailyNetLossLimit() int32 {
	return message.Int32VLS(m.Unsafe(), 34, 30)
}

// PercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataUpdateBuilder) PercentOfCashBalanceForDailyNetLossLimit() int32 {
	return message.Int32VLS(m.Unsafe(), 34, 30)
}

// PercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataUpdatePointer) PercentOfCashBalanceForDailyNetLossLimit() int32 {
	return message.Int32VLS(m.Ptr, 34, 30)
}

// PercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataUpdatePointerBuilder) PercentOfCashBalanceForDailyNetLossLimit() int32 {
	return message.Int32VLS(m.Ptr, 34, 30)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet
func (m TradeAccountDataUpdate) UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 35, 34)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet
func (m TradeAccountDataUpdateBuilder) UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 35, 34)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet
func (m TradeAccountDataUpdatePointer) UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet() bool {
	return message.BoolVLS(m.Ptr, 35, 34)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet() bool {
	return message.BoolVLS(m.Ptr, 35, 34)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataUpdate) UseTrailingAccountValueToNotAllowIncreaseInPositions() bool {
	return message.BoolVLS(m.Unsafe(), 36, 35)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataUpdateBuilder) UseTrailingAccountValueToNotAllowIncreaseInPositions() bool {
	return message.BoolVLS(m.Unsafe(), 36, 35)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataUpdatePointer) UseTrailingAccountValueToNotAllowIncreaseInPositions() bool {
	return message.BoolVLS(m.Ptr, 36, 35)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataUpdatePointerBuilder) UseTrailingAccountValueToNotAllowIncreaseInPositions() bool {
	return message.BoolVLS(m.Ptr, 36, 35)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdate) DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 37, 36)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdateBuilder) DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 37, 36)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointer) DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() bool {
	return message.BoolVLS(m.Ptr, 37, 36)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() bool {
	return message.BoolVLS(m.Ptr, 37, 36)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdate) DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 38, 37)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdateBuilder) DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 38, 37)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointer) DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 38, 37)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointerBuilder) DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 38, 37)
}

// FlattenPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdate) FlattenPositionsAtDailyLossLimitIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 39, 38)
}

// FlattenPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdateBuilder) FlattenPositionsAtDailyLossLimitIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 39, 38)
}

// FlattenPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointer) FlattenPositionsAtDailyLossLimitIsSet() bool {
	return message.BoolVLS(m.Ptr, 39, 38)
}

// FlattenPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) FlattenPositionsAtDailyLossLimitIsSet() bool {
	return message.BoolVLS(m.Ptr, 39, 38)
}

// FlattenPositionsAtDailyLossLimit
func (m TradeAccountDataUpdate) FlattenPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 40, 39)
}

// FlattenPositionsAtDailyLossLimit
func (m TradeAccountDataUpdateBuilder) FlattenPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 40, 39)
}

// FlattenPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointer) FlattenPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 40, 39)
}

// FlattenPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointerBuilder) FlattenPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 40, 39)
}

// ClosePositionsAtEndOfDayIsSet
func (m TradeAccountDataUpdate) ClosePositionsAtEndOfDayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 41, 40)
}

// ClosePositionsAtEndOfDayIsSet
func (m TradeAccountDataUpdateBuilder) ClosePositionsAtEndOfDayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 41, 40)
}

// ClosePositionsAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointer) ClosePositionsAtEndOfDayIsSet() bool {
	return message.BoolVLS(m.Ptr, 41, 40)
}

// ClosePositionsAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) ClosePositionsAtEndOfDayIsSet() bool {
	return message.BoolVLS(m.Ptr, 41, 40)
}

// ClosePositionsAtEndOfDay
func (m TradeAccountDataUpdate) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// ClosePositionsAtEndOfDay
func (m TradeAccountDataUpdateBuilder) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// ClosePositionsAtEndOfDay
func (m TradeAccountDataUpdatePointer) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 42, 41)
}

// ClosePositionsAtEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 42, 41)
}

// FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdate) FlattenPositionsWhenUnderMarginIntradayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 43, 42)
}

// FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdateBuilder) FlattenPositionsWhenUnderMarginIntradayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 43, 42)
}

// FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointer) FlattenPositionsWhenUnderMarginIntradayIsSet() bool {
	return message.BoolVLS(m.Ptr, 43, 42)
}

// FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointerBuilder) FlattenPositionsWhenUnderMarginIntradayIsSet() bool {
	return message.BoolVLS(m.Ptr, 43, 42)
}

// FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdate) FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 44, 43)
}

// FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdateBuilder) FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 44, 43)
}

// FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointer) FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 44, 43)
}

// FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointerBuilder) FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 44, 43)
}

// FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdate) FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 45, 44)
}

// FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdateBuilder) FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 45, 44)
}

// FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointer) FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() bool {
	return message.BoolVLS(m.Ptr, 45, 44)
}

// FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() bool {
	return message.BoolVLS(m.Ptr, 45, 44)
}

// FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdate) FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 46, 45)
}

// FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdateBuilder) FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 46, 45)
}

// FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdatePointer) FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 46, 45)
}

// FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 46, 45)
}

// SenderSubIDIsSet
func (m TradeAccountDataUpdate) SenderSubIDIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 47, 46)
}

// SenderSubIDIsSet
func (m TradeAccountDataUpdateBuilder) SenderSubIDIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 47, 46)
}

// SenderSubIDIsSet
func (m TradeAccountDataUpdatePointer) SenderSubIDIsSet() bool {
	return message.BoolVLS(m.Ptr, 47, 46)
}

// SenderSubIDIsSet
func (m TradeAccountDataUpdatePointerBuilder) SenderSubIDIsSet() bool {
	return message.BoolVLS(m.Ptr, 47, 46)
}

// SenderSubID
func (m TradeAccountDataUpdate) SenderSubID() string {
	return message.StringVLS(m.Unsafe(), 51, 47)
}

// SenderSubID
func (m TradeAccountDataUpdateBuilder) SenderSubID() string {
	return message.StringVLS(m.Unsafe(), 51, 47)
}

// SenderSubID
func (m TradeAccountDataUpdatePointer) SenderSubID() string {
	return message.StringVLS(m.Ptr, 51, 47)
}

// SenderSubID
func (m TradeAccountDataUpdatePointerBuilder) SenderSubID() string {
	return message.StringVLS(m.Ptr, 51, 47)
}

// SenderLocationIdIsSet
func (m TradeAccountDataUpdate) SenderLocationIdIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 52, 51)
}

// SenderLocationIdIsSet
func (m TradeAccountDataUpdateBuilder) SenderLocationIdIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 52, 51)
}

// SenderLocationIdIsSet
func (m TradeAccountDataUpdatePointer) SenderLocationIdIsSet() bool {
	return message.BoolVLS(m.Ptr, 52, 51)
}

// SenderLocationIdIsSet
func (m TradeAccountDataUpdatePointerBuilder) SenderLocationIdIsSet() bool {
	return message.BoolVLS(m.Ptr, 52, 51)
}

// SenderLocationId
func (m TradeAccountDataUpdate) SenderLocationId() string {
	return message.StringVLS(m.Unsafe(), 56, 52)
}

// SenderLocationId
func (m TradeAccountDataUpdateBuilder) SenderLocationId() string {
	return message.StringVLS(m.Unsafe(), 56, 52)
}

// SenderLocationId
func (m TradeAccountDataUpdatePointer) SenderLocationId() string {
	return message.StringVLS(m.Ptr, 56, 52)
}

// SenderLocationId
func (m TradeAccountDataUpdatePointerBuilder) SenderLocationId() string {
	return message.StringVLS(m.Ptr, 56, 52)
}

// SelfMatchPreventionIDIsSet
func (m TradeAccountDataUpdate) SelfMatchPreventionIDIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 57, 56)
}

// SelfMatchPreventionIDIsSet
func (m TradeAccountDataUpdateBuilder) SelfMatchPreventionIDIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 57, 56)
}

// SelfMatchPreventionIDIsSet
func (m TradeAccountDataUpdatePointer) SelfMatchPreventionIDIsSet() bool {
	return message.BoolVLS(m.Ptr, 57, 56)
}

// SelfMatchPreventionIDIsSet
func (m TradeAccountDataUpdatePointerBuilder) SelfMatchPreventionIDIsSet() bool {
	return message.BoolVLS(m.Ptr, 57, 56)
}

// SelfMatchPreventionID
func (m TradeAccountDataUpdate) SelfMatchPreventionID() string {
	return message.StringVLS(m.Unsafe(), 61, 57)
}

// SelfMatchPreventionID
func (m TradeAccountDataUpdateBuilder) SelfMatchPreventionID() string {
	return message.StringVLS(m.Unsafe(), 61, 57)
}

// SelfMatchPreventionID
func (m TradeAccountDataUpdatePointer) SelfMatchPreventionID() string {
	return message.StringVLS(m.Ptr, 61, 57)
}

// SelfMatchPreventionID
func (m TradeAccountDataUpdatePointerBuilder) SelfMatchPreventionID() string {
	return message.StringVLS(m.Ptr, 61, 57)
}

// CTICodeIsSet
func (m TradeAccountDataUpdate) CTICodeIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 62, 61)
}

// CTICodeIsSet
func (m TradeAccountDataUpdateBuilder) CTICodeIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 62, 61)
}

// CTICodeIsSet
func (m TradeAccountDataUpdatePointer) CTICodeIsSet() bool {
	return message.BoolVLS(m.Ptr, 62, 61)
}

// CTICodeIsSet
func (m TradeAccountDataUpdatePointerBuilder) CTICodeIsSet() bool {
	return message.BoolVLS(m.Ptr, 62, 61)
}

// CTICode
func (m TradeAccountDataUpdate) CTICode() int32 {
	return message.Int32VLS(m.Unsafe(), 66, 62)
}

// CTICode
func (m TradeAccountDataUpdateBuilder) CTICode() int32 {
	return message.Int32VLS(m.Unsafe(), 66, 62)
}

// CTICode
func (m TradeAccountDataUpdatePointer) CTICode() int32 {
	return message.Int32VLS(m.Ptr, 66, 62)
}

// CTICode
func (m TradeAccountDataUpdatePointerBuilder) CTICode() int32 {
	return message.Int32VLS(m.Ptr, 66, 62)
}

// TradeAccountIsReadOnlyIsSet
func (m TradeAccountDataUpdate) TradeAccountIsReadOnlyIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 67, 66)
}

// TradeAccountIsReadOnlyIsSet
func (m TradeAccountDataUpdateBuilder) TradeAccountIsReadOnlyIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 67, 66)
}

// TradeAccountIsReadOnlyIsSet
func (m TradeAccountDataUpdatePointer) TradeAccountIsReadOnlyIsSet() bool {
	return message.BoolVLS(m.Ptr, 67, 66)
}

// TradeAccountIsReadOnlyIsSet
func (m TradeAccountDataUpdatePointerBuilder) TradeAccountIsReadOnlyIsSet() bool {
	return message.BoolVLS(m.Ptr, 67, 66)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataUpdate) TradeAccountIsReadOnly() bool {
	return message.BoolVLS(m.Unsafe(), 68, 67)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataUpdateBuilder) TradeAccountIsReadOnly() bool {
	return message.BoolVLS(m.Unsafe(), 68, 67)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataUpdatePointer) TradeAccountIsReadOnly() bool {
	return message.BoolVLS(m.Ptr, 68, 67)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataUpdatePointerBuilder) TradeAccountIsReadOnly() bool {
	return message.BoolVLS(m.Ptr, 68, 67)
}

// MaximumGlobalPositionQuantityIsSet
func (m TradeAccountDataUpdate) MaximumGlobalPositionQuantityIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 69, 68)
}

// MaximumGlobalPositionQuantityIsSet
func (m TradeAccountDataUpdateBuilder) MaximumGlobalPositionQuantityIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 69, 68)
}

// MaximumGlobalPositionQuantityIsSet
func (m TradeAccountDataUpdatePointer) MaximumGlobalPositionQuantityIsSet() bool {
	return message.BoolVLS(m.Ptr, 69, 68)
}

// MaximumGlobalPositionQuantityIsSet
func (m TradeAccountDataUpdatePointerBuilder) MaximumGlobalPositionQuantityIsSet() bool {
	return message.BoolVLS(m.Ptr, 69, 68)
}

// MaximumGlobalPositionQuantity
func (m TradeAccountDataUpdate) MaximumGlobalPositionQuantity() int32 {
	return message.Int32VLS(m.Unsafe(), 73, 69)
}

// MaximumGlobalPositionQuantity
func (m TradeAccountDataUpdateBuilder) MaximumGlobalPositionQuantity() int32 {
	return message.Int32VLS(m.Unsafe(), 73, 69)
}

// MaximumGlobalPositionQuantity
func (m TradeAccountDataUpdatePointer) MaximumGlobalPositionQuantity() int32 {
	return message.Int32VLS(m.Ptr, 73, 69)
}

// MaximumGlobalPositionQuantity
func (m TradeAccountDataUpdatePointerBuilder) MaximumGlobalPositionQuantity() int32 {
	return message.Int32VLS(m.Ptr, 73, 69)
}

// TradeFeePerContractIsSet
func (m TradeAccountDataUpdate) TradeFeePerContractIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 74, 73)
}

// TradeFeePerContractIsSet
func (m TradeAccountDataUpdateBuilder) TradeFeePerContractIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 74, 73)
}

// TradeFeePerContractIsSet
func (m TradeAccountDataUpdatePointer) TradeFeePerContractIsSet() bool {
	return message.BoolVLS(m.Ptr, 74, 73)
}

// TradeFeePerContractIsSet
func (m TradeAccountDataUpdatePointerBuilder) TradeFeePerContractIsSet() bool {
	return message.BoolVLS(m.Ptr, 74, 73)
}

// TradeFeePerContract
func (m TradeAccountDataUpdate) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Unsafe(), 82, 74)
}

// TradeFeePerContract
func (m TradeAccountDataUpdateBuilder) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Unsafe(), 82, 74)
}

// TradeFeePerContract
func (m TradeAccountDataUpdatePointer) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Ptr, 82, 74)
}

// TradeFeePerContract
func (m TradeAccountDataUpdatePointerBuilder) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Ptr, 82, 74)
}

// TradeFeePerShareIsSet
func (m TradeAccountDataUpdate) TradeFeePerShareIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 83, 82)
}

// TradeFeePerShareIsSet
func (m TradeAccountDataUpdateBuilder) TradeFeePerShareIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 83, 82)
}

// TradeFeePerShareIsSet
func (m TradeAccountDataUpdatePointer) TradeFeePerShareIsSet() bool {
	return message.BoolVLS(m.Ptr, 83, 82)
}

// TradeFeePerShareIsSet
func (m TradeAccountDataUpdatePointerBuilder) TradeFeePerShareIsSet() bool {
	return message.BoolVLS(m.Ptr, 83, 82)
}

// TradeFeePerShare
func (m TradeAccountDataUpdate) TradeFeePerShare() float64 {
	return message.Float64VLS(m.Unsafe(), 91, 83)
}

// TradeFeePerShare
func (m TradeAccountDataUpdateBuilder) TradeFeePerShare() float64 {
	return message.Float64VLS(m.Unsafe(), 91, 83)
}

// TradeFeePerShare
func (m TradeAccountDataUpdatePointer) TradeFeePerShare() float64 {
	return message.Float64VLS(m.Ptr, 91, 83)
}

// TradeFeePerShare
func (m TradeAccountDataUpdatePointerBuilder) TradeFeePerShare() float64 {
	return message.Float64VLS(m.Ptr, 91, 83)
}

// RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdate) RequireSufficientMarginForNewPositionsIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 92, 91)
}

// RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdateBuilder) RequireSufficientMarginForNewPositionsIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 92, 91)
}

// RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdatePointer) RequireSufficientMarginForNewPositionsIsSet() bool {
	return message.BoolVLS(m.Ptr, 92, 91)
}

// RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdatePointerBuilder) RequireSufficientMarginForNewPositionsIsSet() bool {
	return message.BoolVLS(m.Ptr, 92, 91)
}

// RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdate) RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 93, 92)
}

// RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdateBuilder) RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 93, 92)
}

// RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdatePointer) RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 93, 92)
}

// RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdatePointerBuilder) RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 93, 92)
}

// UsePercentOfMarginIsSet
func (m TradeAccountDataUpdate) UsePercentOfMarginIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 94, 93)
}

// UsePercentOfMarginIsSet
func (m TradeAccountDataUpdateBuilder) UsePercentOfMarginIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 94, 93)
}

// UsePercentOfMarginIsSet
func (m TradeAccountDataUpdatePointer) UsePercentOfMarginIsSet() bool {
	return message.BoolVLS(m.Ptr, 94, 93)
}

// UsePercentOfMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) UsePercentOfMarginIsSet() bool {
	return message.BoolVLS(m.Ptr, 94, 93)
}

// UsePercentOfMargin
func (m TradeAccountDataUpdate) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Unsafe(), 98, 94)
}

// UsePercentOfMargin
func (m TradeAccountDataUpdateBuilder) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Unsafe(), 98, 94)
}

// UsePercentOfMargin
func (m TradeAccountDataUpdatePointer) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Ptr, 98, 94)
}

// UsePercentOfMargin
func (m TradeAccountDataUpdatePointerBuilder) UsePercentOfMargin() int32 {
	return message.Int32VLS(m.Ptr, 98, 94)
}

// UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdate) UsePercentOfMarginForDayTradingIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 99, 98)
}

// UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdateBuilder) UsePercentOfMarginForDayTradingIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 99, 98)
}

// UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdatePointer) UsePercentOfMarginForDayTradingIsSet() bool {
	return message.BoolVLS(m.Ptr, 99, 98)
}

// UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdatePointerBuilder) UsePercentOfMarginForDayTradingIsSet() bool {
	return message.BoolVLS(m.Ptr, 99, 98)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdate) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Unsafe(), 103, 99)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdateBuilder) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Unsafe(), 103, 99)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdatePointer) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Ptr, 103, 99)
}

// UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdatePointerBuilder) UsePercentOfMarginForDayTrading() int32 {
	return message.Int32VLS(m.Ptr, 103, 99)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdate) MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 104, 103)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdateBuilder) MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 104, 103)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdatePointer) MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() bool {
	return message.BoolVLS(m.Ptr, 104, 103)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdatePointerBuilder) MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() bool {
	return message.BoolVLS(m.Ptr, 104, 103)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdate) MaximumAllowedAccountBalanceForPositionsAsPercentage() int32 {
	return message.Int32VLS(m.Unsafe(), 108, 104)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdateBuilder) MaximumAllowedAccountBalanceForPositionsAsPercentage() int32 {
	return message.Int32VLS(m.Unsafe(), 108, 104)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdatePointer) MaximumAllowedAccountBalanceForPositionsAsPercentage() int32 {
	return message.Int32VLS(m.Ptr, 108, 104)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdatePointerBuilder) MaximumAllowedAccountBalanceForPositionsAsPercentage() int32 {
	return message.Int32VLS(m.Ptr, 108, 104)
}

// FirmIDIsSet
func (m TradeAccountDataUpdate) FirmIDIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 109, 108)
}

// FirmIDIsSet
func (m TradeAccountDataUpdateBuilder) FirmIDIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 109, 108)
}

// FirmIDIsSet
func (m TradeAccountDataUpdatePointer) FirmIDIsSet() bool {
	return message.BoolVLS(m.Ptr, 109, 108)
}

// FirmIDIsSet
func (m TradeAccountDataUpdatePointerBuilder) FirmIDIsSet() bool {
	return message.BoolVLS(m.Ptr, 109, 108)
}

// FirmID
func (m TradeAccountDataUpdate) FirmID() string {
	return message.StringVLS(m.Unsafe(), 113, 109)
}

// FirmID
func (m TradeAccountDataUpdateBuilder) FirmID() string {
	return message.StringVLS(m.Unsafe(), 113, 109)
}

// FirmID
func (m TradeAccountDataUpdatePointer) FirmID() string {
	return message.StringVLS(m.Ptr, 113, 109)
}

// FirmID
func (m TradeAccountDataUpdatePointerBuilder) FirmID() string {
	return message.StringVLS(m.Ptr, 113, 109)
}

// TradingIsDisabledIsSet
func (m TradeAccountDataUpdate) TradingIsDisabledIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 114, 113)
}

// TradingIsDisabledIsSet
func (m TradeAccountDataUpdateBuilder) TradingIsDisabledIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 114, 113)
}

// TradingIsDisabledIsSet
func (m TradeAccountDataUpdatePointer) TradingIsDisabledIsSet() bool {
	return message.BoolVLS(m.Ptr, 114, 113)
}

// TradingIsDisabledIsSet
func (m TradeAccountDataUpdatePointerBuilder) TradingIsDisabledIsSet() bool {
	return message.BoolVLS(m.Ptr, 114, 113)
}

// TradingIsDisabled
func (m TradeAccountDataUpdate) TradingIsDisabled() bool {
	return message.BoolVLS(m.Unsafe(), 115, 114)
}

// TradingIsDisabled
func (m TradeAccountDataUpdateBuilder) TradingIsDisabled() bool {
	return message.BoolVLS(m.Unsafe(), 115, 114)
}

// TradingIsDisabled
func (m TradeAccountDataUpdatePointer) TradingIsDisabled() bool {
	return message.BoolVLS(m.Ptr, 115, 114)
}

// TradingIsDisabled
func (m TradeAccountDataUpdatePointerBuilder) TradingIsDisabled() bool {
	return message.BoolVLS(m.Ptr, 115, 114)
}

// DescriptiveNameIsSet
func (m TradeAccountDataUpdate) DescriptiveNameIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 116, 115)
}

// DescriptiveNameIsSet
func (m TradeAccountDataUpdateBuilder) DescriptiveNameIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 116, 115)
}

// DescriptiveNameIsSet
func (m TradeAccountDataUpdatePointer) DescriptiveNameIsSet() bool {
	return message.BoolVLS(m.Ptr, 116, 115)
}

// DescriptiveNameIsSet
func (m TradeAccountDataUpdatePointerBuilder) DescriptiveNameIsSet() bool {
	return message.BoolVLS(m.Ptr, 116, 115)
}

// DescriptiveName
func (m TradeAccountDataUpdate) DescriptiveName() string {
	return message.StringVLS(m.Unsafe(), 120, 116)
}

// DescriptiveName
func (m TradeAccountDataUpdateBuilder) DescriptiveName() string {
	return message.StringVLS(m.Unsafe(), 120, 116)
}

// DescriptiveName
func (m TradeAccountDataUpdatePointer) DescriptiveName() string {
	return message.StringVLS(m.Ptr, 120, 116)
}

// DescriptiveName
func (m TradeAccountDataUpdatePointerBuilder) DescriptiveName() string {
	return message.StringVLS(m.Ptr, 120, 116)
}

// IsMasterFirmControlAccountIsSet
func (m TradeAccountDataUpdate) IsMasterFirmControlAccountIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 121, 120)
}

// IsMasterFirmControlAccountIsSet
func (m TradeAccountDataUpdateBuilder) IsMasterFirmControlAccountIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 121, 120)
}

// IsMasterFirmControlAccountIsSet
func (m TradeAccountDataUpdatePointer) IsMasterFirmControlAccountIsSet() bool {
	return message.BoolVLS(m.Ptr, 121, 120)
}

// IsMasterFirmControlAccountIsSet
func (m TradeAccountDataUpdatePointerBuilder) IsMasterFirmControlAccountIsSet() bool {
	return message.BoolVLS(m.Ptr, 121, 120)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataUpdate) IsMasterFirmControlAccount() bool {
	return message.BoolVLS(m.Unsafe(), 122, 121)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataUpdateBuilder) IsMasterFirmControlAccount() bool {
	return message.BoolVLS(m.Unsafe(), 122, 121)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataUpdatePointer) IsMasterFirmControlAccount() bool {
	return message.BoolVLS(m.Ptr, 122, 121)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataUpdatePointerBuilder) IsMasterFirmControlAccount() bool {
	return message.BoolVLS(m.Ptr, 122, 121)
}

// MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdate) MinimumRequiredAccountValueIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 123, 122)
}

// MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdateBuilder) MinimumRequiredAccountValueIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 123, 122)
}

// MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdatePointer) MinimumRequiredAccountValueIsSet() bool {
	return message.BoolVLS(m.Ptr, 123, 122)
}

// MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdatePointerBuilder) MinimumRequiredAccountValueIsSet() bool {
	return message.BoolVLS(m.Ptr, 123, 122)
}

// MinimumRequiredAccountValue
func (m TradeAccountDataUpdate) MinimumRequiredAccountValue() float64 {
	return message.Float64VLS(m.Unsafe(), 131, 123)
}

// MinimumRequiredAccountValue
func (m TradeAccountDataUpdateBuilder) MinimumRequiredAccountValue() float64 {
	return message.Float64VLS(m.Unsafe(), 131, 123)
}

// MinimumRequiredAccountValue
func (m TradeAccountDataUpdatePointer) MinimumRequiredAccountValue() float64 {
	return message.Float64VLS(m.Ptr, 131, 123)
}

// MinimumRequiredAccountValue
func (m TradeAccountDataUpdatePointerBuilder) MinimumRequiredAccountValue() float64 {
	return message.Float64VLS(m.Ptr, 131, 123)
}

// BeginTimeForDayMarginIsSet
func (m TradeAccountDataUpdate) BeginTimeForDayMarginIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 132, 131)
}

// BeginTimeForDayMarginIsSet
func (m TradeAccountDataUpdateBuilder) BeginTimeForDayMarginIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 132, 131)
}

// BeginTimeForDayMarginIsSet
func (m TradeAccountDataUpdatePointer) BeginTimeForDayMarginIsSet() bool {
	return message.BoolVLS(m.Ptr, 132, 131)
}

// BeginTimeForDayMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) BeginTimeForDayMarginIsSet() bool {
	return message.BoolVLS(m.Ptr, 132, 131)
}

// BeginTimeForDayMargin
func (m TradeAccountDataUpdate) BeginTimeForDayMargin() int64 {
	return message.Int64VLS(m.Unsafe(), 140, 132)
}

// BeginTimeForDayMargin
func (m TradeAccountDataUpdateBuilder) BeginTimeForDayMargin() int64 {
	return message.Int64VLS(m.Unsafe(), 140, 132)
}

// BeginTimeForDayMargin
func (m TradeAccountDataUpdatePointer) BeginTimeForDayMargin() int64 {
	return message.Int64VLS(m.Ptr, 140, 132)
}

// BeginTimeForDayMargin
func (m TradeAccountDataUpdatePointerBuilder) BeginTimeForDayMargin() int64 {
	return message.Int64VLS(m.Ptr, 140, 132)
}

// EndTimeForDayMarginIsSet
func (m TradeAccountDataUpdate) EndTimeForDayMarginIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 141, 140)
}

// EndTimeForDayMarginIsSet
func (m TradeAccountDataUpdateBuilder) EndTimeForDayMarginIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 141, 140)
}

// EndTimeForDayMarginIsSet
func (m TradeAccountDataUpdatePointer) EndTimeForDayMarginIsSet() bool {
	return message.BoolVLS(m.Ptr, 141, 140)
}

// EndTimeForDayMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) EndTimeForDayMarginIsSet() bool {
	return message.BoolVLS(m.Ptr, 141, 140)
}

// EndTimeForDayMargin
func (m TradeAccountDataUpdate) EndTimeForDayMargin() int64 {
	return message.Int64VLS(m.Unsafe(), 149, 141)
}

// EndTimeForDayMargin
func (m TradeAccountDataUpdateBuilder) EndTimeForDayMargin() int64 {
	return message.Int64VLS(m.Unsafe(), 149, 141)
}

// EndTimeForDayMargin
func (m TradeAccountDataUpdatePointer) EndTimeForDayMargin() int64 {
	return message.Int64VLS(m.Ptr, 149, 141)
}

// EndTimeForDayMargin
func (m TradeAccountDataUpdatePointerBuilder) EndTimeForDayMargin() int64 {
	return message.Int64VLS(m.Ptr, 149, 141)
}

// DayMarginTimeZoneIsSet
func (m TradeAccountDataUpdate) DayMarginTimeZoneIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 150, 149)
}

// DayMarginTimeZoneIsSet
func (m TradeAccountDataUpdateBuilder) DayMarginTimeZoneIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 150, 149)
}

// DayMarginTimeZoneIsSet
func (m TradeAccountDataUpdatePointer) DayMarginTimeZoneIsSet() bool {
	return message.BoolVLS(m.Ptr, 150, 149)
}

// DayMarginTimeZoneIsSet
func (m TradeAccountDataUpdatePointerBuilder) DayMarginTimeZoneIsSet() bool {
	return message.BoolVLS(m.Ptr, 150, 149)
}

// DayMarginTimeZone
func (m TradeAccountDataUpdate) DayMarginTimeZone() string {
	return message.StringVLS(m.Unsafe(), 154, 150)
}

// DayMarginTimeZone
func (m TradeAccountDataUpdateBuilder) DayMarginTimeZone() string {
	return message.StringVLS(m.Unsafe(), 154, 150)
}

// DayMarginTimeZone
func (m TradeAccountDataUpdatePointer) DayMarginTimeZone() string {
	return message.StringVLS(m.Ptr, 154, 150)
}

// DayMarginTimeZone
func (m TradeAccountDataUpdatePointerBuilder) DayMarginTimeZone() string {
	return message.StringVLS(m.Ptr, 154, 150)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 155, 154)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 155, 154)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet() bool {
	return message.BoolVLS(m.Ptr, 155, 154)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet() bool {
	return message.BoolVLS(m.Ptr, 155, 154)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdate) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() bool {
	return message.BoolVLS(m.Unsafe(), 156, 155)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() bool {
	return message.BoolVLS(m.Unsafe(), 156, 155)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointer) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() bool {
	return message.BoolVLS(m.Ptr, 156, 155)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() bool {
	return message.BoolVLS(m.Ptr, 156, 155)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 157, 156)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 157, 156)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() bool {
	return message.BoolVLS(m.Ptr, 157, 156)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() bool {
	return message.BoolVLS(m.Ptr, 157, 156)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdate) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() bool {
	return message.BoolVLS(m.Unsafe(), 158, 157)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() bool {
	return message.BoolVLS(m.Unsafe(), 158, 157)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdatePointer) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() bool {
	return message.BoolVLS(m.Ptr, 158, 157)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() bool {
	return message.BoolVLS(m.Ptr, 158, 157)
}

// UseMasterFirm_SymbolLimitsArrayIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_SymbolLimitsArrayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 159, 158)
}

// UseMasterFirm_SymbolLimitsArrayIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_SymbolLimitsArrayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 159, 158)
}

// UseMasterFirm_SymbolLimitsArrayIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_SymbolLimitsArrayIsSet() bool {
	return message.BoolVLS(m.Ptr, 159, 158)
}

// UseMasterFirm_SymbolLimitsArrayIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_SymbolLimitsArrayIsSet() bool {
	return message.BoolVLS(m.Ptr, 159, 158)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataUpdate) UseMasterFirm_SymbolLimitsArray() bool {
	return message.BoolVLS(m.Unsafe(), 160, 159)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_SymbolLimitsArray() bool {
	return message.BoolVLS(m.Unsafe(), 160, 159)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataUpdatePointer) UseMasterFirm_SymbolLimitsArray() bool {
	return message.BoolVLS(m.Ptr, 160, 159)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_SymbolLimitsArray() bool {
	return message.BoolVLS(m.Ptr, 160, 159)
}

// UseMasterFirm_TradeFeesIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_TradeFeesIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 161, 160)
}

// UseMasterFirm_TradeFeesIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_TradeFeesIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 161, 160)
}

// UseMasterFirm_TradeFeesIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradeFeesIsSet() bool {
	return message.BoolVLS(m.Ptr, 161, 160)
}

// UseMasterFirm_TradeFeesIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradeFeesIsSet() bool {
	return message.BoolVLS(m.Ptr, 161, 160)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataUpdate) UseMasterFirm_TradeFees() bool {
	return message.BoolVLS(m.Unsafe(), 162, 161)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_TradeFees() bool {
	return message.BoolVLS(m.Unsafe(), 162, 161)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradeFees() bool {
	return message.BoolVLS(m.Ptr, 162, 161)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradeFees() bool {
	return message.BoolVLS(m.Ptr, 162, 161)
}

// UseMasterFirm_TradeFeePerShareIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_TradeFeePerShareIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 163, 162)
}

// UseMasterFirm_TradeFeePerShareIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_TradeFeePerShareIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 163, 162)
}

// UseMasterFirm_TradeFeePerShareIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradeFeePerShareIsSet() bool {
	return message.BoolVLS(m.Ptr, 163, 162)
}

// UseMasterFirm_TradeFeePerShareIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradeFeePerShareIsSet() bool {
	return message.BoolVLS(m.Ptr, 163, 162)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataUpdate) UseMasterFirm_TradeFeePerShare() bool {
	return message.BoolVLS(m.Unsafe(), 164, 163)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_TradeFeePerShare() bool {
	return message.BoolVLS(m.Unsafe(), 164, 163)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradeFeePerShare() bool {
	return message.BoolVLS(m.Ptr, 164, 163)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradeFeePerShare() bool {
	return message.BoolVLS(m.Ptr, 164, 163)
}

// UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 165, 164)
}

// UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 165, 164)
}

// UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet() bool {
	return message.BoolVLS(m.Ptr, 165, 164)
}

// UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet() bool {
	return message.BoolVLS(m.Ptr, 165, 164)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdate) UseMasterFirm_RequireSufficientMarginForNewPositions() bool {
	return message.BoolVLS(m.Unsafe(), 166, 165)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_RequireSufficientMarginForNewPositions() bool {
	return message.BoolVLS(m.Unsafe(), 166, 165)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdatePointer) UseMasterFirm_RequireSufficientMarginForNewPositions() bool {
	return message.BoolVLS(m.Ptr, 166, 165)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_RequireSufficientMarginForNewPositions() bool {
	return message.BoolVLS(m.Ptr, 166, 165)
}

// UseMasterFirm_UsePercentOfMarginIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_UsePercentOfMarginIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 167, 166)
}

// UseMasterFirm_UsePercentOfMarginIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_UsePercentOfMarginIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 167, 166)
}

// UseMasterFirm_UsePercentOfMarginIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginIsSet() bool {
	return message.BoolVLS(m.Ptr, 167, 166)
}

// UseMasterFirm_UsePercentOfMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginIsSet() bool {
	return message.BoolVLS(m.Ptr, 167, 166)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataUpdate) UseMasterFirm_UsePercentOfMargin() bool {
	return message.BoolVLS(m.Unsafe(), 168, 167)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_UsePercentOfMargin() bool {
	return message.BoolVLS(m.Unsafe(), 168, 167)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMargin() bool {
	return message.BoolVLS(m.Ptr, 168, 167)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMargin() bool {
	return message.BoolVLS(m.Ptr, 168, 167)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 169, 168)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 169, 168)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() bool {
	return message.BoolVLS(m.Ptr, 169, 168)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() bool {
	return message.BoolVLS(m.Ptr, 169, 168)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdate) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() bool {
	return message.BoolVLS(m.Unsafe(), 170, 169)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() bool {
	return message.BoolVLS(m.Unsafe(), 170, 169)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdatePointer) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() bool {
	return message.BoolVLS(m.Ptr, 170, 169)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() bool {
	return message.BoolVLS(m.Ptr, 170, 169)
}

// UseMasterFirm_MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_MinimumRequiredAccountValueIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 171, 170)
}

// UseMasterFirm_MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_MinimumRequiredAccountValueIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 171, 170)
}

// UseMasterFirm_MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_MinimumRequiredAccountValueIsSet() bool {
	return message.BoolVLS(m.Ptr, 171, 170)
}

// UseMasterFirm_MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_MinimumRequiredAccountValueIsSet() bool {
	return message.BoolVLS(m.Ptr, 171, 170)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataUpdate) UseMasterFirm_MinimumRequiredAccountValue() bool {
	return message.BoolVLS(m.Unsafe(), 172, 171)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_MinimumRequiredAccountValue() bool {
	return message.BoolVLS(m.Unsafe(), 172, 171)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataUpdatePointer) UseMasterFirm_MinimumRequiredAccountValue() bool {
	return message.BoolVLS(m.Ptr, 172, 171)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_MinimumRequiredAccountValue() bool {
	return message.BoolVLS(m.Ptr, 172, 171)
}

// UseMasterFirm_MarginTimeSettingsIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_MarginTimeSettingsIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 173, 172)
}

// UseMasterFirm_MarginTimeSettingsIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_MarginTimeSettingsIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 173, 172)
}

// UseMasterFirm_MarginTimeSettingsIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_MarginTimeSettingsIsSet() bool {
	return message.BoolVLS(m.Ptr, 173, 172)
}

// UseMasterFirm_MarginTimeSettingsIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_MarginTimeSettingsIsSet() bool {
	return message.BoolVLS(m.Ptr, 173, 172)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataUpdate) UseMasterFirm_MarginTimeSettings() bool {
	return message.BoolVLS(m.Unsafe(), 174, 173)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_MarginTimeSettings() bool {
	return message.BoolVLS(m.Unsafe(), 174, 173)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataUpdatePointer) UseMasterFirm_MarginTimeSettings() bool {
	return message.BoolVLS(m.Ptr, 174, 173)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_MarginTimeSettings() bool {
	return message.BoolVLS(m.Ptr, 174, 173)
}

// UseMasterFirm_TradingIsDisabledIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_TradingIsDisabledIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 175, 174)
}

// UseMasterFirm_TradingIsDisabledIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_TradingIsDisabledIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 175, 174)
}

// UseMasterFirm_TradingIsDisabledIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradingIsDisabledIsSet() bool {
	return message.BoolVLS(m.Ptr, 175, 174)
}

// UseMasterFirm_TradingIsDisabledIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradingIsDisabledIsSet() bool {
	return message.BoolVLS(m.Ptr, 175, 174)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataUpdate) UseMasterFirm_TradingIsDisabled() bool {
	return message.BoolVLS(m.Unsafe(), 176, 175)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_TradingIsDisabled() bool {
	return message.BoolVLS(m.Unsafe(), 176, 175)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradingIsDisabled() bool {
	return message.BoolVLS(m.Ptr, 176, 175)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradingIsDisabled() bool {
	return message.BoolVLS(m.Ptr, 176, 175)
}

// IsTradeStatisticsPublicallySharedIsSet
func (m TradeAccountDataUpdate) IsTradeStatisticsPublicallySharedIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 177, 176)
}

// IsTradeStatisticsPublicallySharedIsSet
func (m TradeAccountDataUpdateBuilder) IsTradeStatisticsPublicallySharedIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 177, 176)
}

// IsTradeStatisticsPublicallySharedIsSet
func (m TradeAccountDataUpdatePointer) IsTradeStatisticsPublicallySharedIsSet() bool {
	return message.BoolVLS(m.Ptr, 177, 176)
}

// IsTradeStatisticsPublicallySharedIsSet
func (m TradeAccountDataUpdatePointerBuilder) IsTradeStatisticsPublicallySharedIsSet() bool {
	return message.BoolVLS(m.Ptr, 177, 176)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataUpdate) IsTradeStatisticsPublicallyShared() bool {
	return message.BoolVLS(m.Unsafe(), 178, 177)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataUpdateBuilder) IsTradeStatisticsPublicallyShared() bool {
	return message.BoolVLS(m.Unsafe(), 178, 177)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataUpdatePointer) IsTradeStatisticsPublicallyShared() bool {
	return message.BoolVLS(m.Ptr, 178, 177)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataUpdatePointerBuilder) IsTradeStatisticsPublicallyShared() bool {
	return message.BoolVLS(m.Ptr, 178, 177)
}

// IsReadOnlyFollowingRequestsAllowedIsSet
func (m TradeAccountDataUpdate) IsReadOnlyFollowingRequestsAllowedIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 179, 178)
}

// IsReadOnlyFollowingRequestsAllowedIsSet
func (m TradeAccountDataUpdateBuilder) IsReadOnlyFollowingRequestsAllowedIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 179, 178)
}

// IsReadOnlyFollowingRequestsAllowedIsSet
func (m TradeAccountDataUpdatePointer) IsReadOnlyFollowingRequestsAllowedIsSet() bool {
	return message.BoolVLS(m.Ptr, 179, 178)
}

// IsReadOnlyFollowingRequestsAllowedIsSet
func (m TradeAccountDataUpdatePointerBuilder) IsReadOnlyFollowingRequestsAllowedIsSet() bool {
	return message.BoolVLS(m.Ptr, 179, 178)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataUpdate) IsReadOnlyFollowingRequestsAllowed() bool {
	return message.BoolVLS(m.Unsafe(), 180, 179)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataUpdateBuilder) IsReadOnlyFollowingRequestsAllowed() bool {
	return message.BoolVLS(m.Unsafe(), 180, 179)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataUpdatePointer) IsReadOnlyFollowingRequestsAllowed() bool {
	return message.BoolVLS(m.Ptr, 180, 179)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataUpdatePointerBuilder) IsReadOnlyFollowingRequestsAllowed() bool {
	return message.BoolVLS(m.Ptr, 180, 179)
}

// IsTradeAccountSharingAllowedIsSet
func (m TradeAccountDataUpdate) IsTradeAccountSharingAllowedIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 181, 180)
}

// IsTradeAccountSharingAllowedIsSet
func (m TradeAccountDataUpdateBuilder) IsTradeAccountSharingAllowedIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 181, 180)
}

// IsTradeAccountSharingAllowedIsSet
func (m TradeAccountDataUpdatePointer) IsTradeAccountSharingAllowedIsSet() bool {
	return message.BoolVLS(m.Ptr, 181, 180)
}

// IsTradeAccountSharingAllowedIsSet
func (m TradeAccountDataUpdatePointerBuilder) IsTradeAccountSharingAllowedIsSet() bool {
	return message.BoolVLS(m.Ptr, 181, 180)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataUpdate) IsTradeAccountSharingAllowed() bool {
	return message.BoolVLS(m.Unsafe(), 182, 181)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataUpdateBuilder) IsTradeAccountSharingAllowed() bool {
	return message.BoolVLS(m.Unsafe(), 182, 181)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataUpdatePointer) IsTradeAccountSharingAllowed() bool {
	return message.BoolVLS(m.Ptr, 182, 181)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataUpdatePointerBuilder) IsTradeAccountSharingAllowed() bool {
	return message.BoolVLS(m.Ptr, 182, 181)
}

// ReadOnlyShareToAllSCUsernamesIsSet
func (m TradeAccountDataUpdate) ReadOnlyShareToAllSCUsernamesIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 183, 182)
}

// ReadOnlyShareToAllSCUsernamesIsSet
func (m TradeAccountDataUpdateBuilder) ReadOnlyShareToAllSCUsernamesIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 183, 182)
}

// ReadOnlyShareToAllSCUsernamesIsSet
func (m TradeAccountDataUpdatePointer) ReadOnlyShareToAllSCUsernamesIsSet() bool {
	return message.BoolVLS(m.Ptr, 183, 182)
}

// ReadOnlyShareToAllSCUsernamesIsSet
func (m TradeAccountDataUpdatePointerBuilder) ReadOnlyShareToAllSCUsernamesIsSet() bool {
	return message.BoolVLS(m.Ptr, 183, 182)
}

// ReadOnlyShareToAllSCUsernames
func (m TradeAccountDataUpdate) ReadOnlyShareToAllSCUsernames() uint8 {
	return message.Uint8VLS(m.Unsafe(), 184, 183)
}

// ReadOnlyShareToAllSCUsernames
func (m TradeAccountDataUpdateBuilder) ReadOnlyShareToAllSCUsernames() uint8 {
	return message.Uint8VLS(m.Unsafe(), 184, 183)
}

// ReadOnlyShareToAllSCUsernames
func (m TradeAccountDataUpdatePointer) ReadOnlyShareToAllSCUsernames() uint8 {
	return message.Uint8VLS(m.Ptr, 184, 183)
}

// ReadOnlyShareToAllSCUsernames
func (m TradeAccountDataUpdatePointerBuilder) ReadOnlyShareToAllSCUsernames() uint8 {
	return message.Uint8VLS(m.Ptr, 184, 183)
}

// UseMasterFirm_SymbolCommissionsArrayIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_SymbolCommissionsArrayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 185, 184)
}

// UseMasterFirm_SymbolCommissionsArrayIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_SymbolCommissionsArrayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 185, 184)
}

// UseMasterFirm_SymbolCommissionsArrayIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_SymbolCommissionsArrayIsSet() bool {
	return message.BoolVLS(m.Ptr, 185, 184)
}

// UseMasterFirm_SymbolCommissionsArrayIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_SymbolCommissionsArrayIsSet() bool {
	return message.BoolVLS(m.Ptr, 185, 184)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataUpdate) UseMasterFirm_SymbolCommissionsArray() bool {
	return message.BoolVLS(m.Unsafe(), 186, 185)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_SymbolCommissionsArray() bool {
	return message.BoolVLS(m.Unsafe(), 186, 185)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataUpdatePointer) UseMasterFirm_SymbolCommissionsArray() bool {
	return message.BoolVLS(m.Ptr, 186, 185)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_SymbolCommissionsArray() bool {
	return message.BoolVLS(m.Ptr, 186, 185)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 187, 186)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 187, 186)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() bool {
	return message.BoolVLS(m.Ptr, 187, 186)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() bool {
	return message.BoolVLS(m.Ptr, 187, 186)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdate) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() bool {
	return message.BoolVLS(m.Unsafe(), 188, 187)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() bool {
	return message.BoolVLS(m.Unsafe(), 188, 187)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointer) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() bool {
	return message.BoolVLS(m.Ptr, 188, 187)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() bool {
	return message.BoolVLS(m.Ptr, 188, 187)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_UsePercentOfMarginForDayTradingIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 189, 188)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 189, 188)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginForDayTradingIsSet() bool {
	return message.BoolVLS(m.Ptr, 189, 188)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingIsSet() bool {
	return message.BoolVLS(m.Ptr, 189, 188)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdate) UseMasterFirm_UsePercentOfMarginForDayTrading() bool {
	return message.BoolVLS(m.Unsafe(), 190, 189)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_UsePercentOfMarginForDayTrading() bool {
	return message.BoolVLS(m.Unsafe(), 190, 189)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginForDayTrading() bool {
	return message.BoolVLS(m.Ptr, 190, 189)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginForDayTrading() bool {
	return message.BoolVLS(m.Ptr, 190, 189)
}

// UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 191, 190)
}

// UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 191, 190)
}

// UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet() bool {
	return message.BoolVLS(m.Ptr, 191, 190)
}

// UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet() bool {
	return message.BoolVLS(m.Ptr, 191, 190)
}

// UseMasterFirm_SymbolCommissionsArrayFullOverride
func (m TradeAccountDataUpdate) UseMasterFirm_SymbolCommissionsArrayFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 192, 191)
}

// UseMasterFirm_SymbolCommissionsArrayFullOverride
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_SymbolCommissionsArrayFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 192, 191)
}

// UseMasterFirm_SymbolCommissionsArrayFullOverride
func (m TradeAccountDataUpdatePointer) UseMasterFirm_SymbolCommissionsArrayFullOverride() bool {
	return message.BoolVLS(m.Ptr, 192, 191)
}

// UseMasterFirm_SymbolCommissionsArrayFullOverride
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_SymbolCommissionsArrayFullOverride() bool {
	return message.BoolVLS(m.Ptr, 192, 191)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 193, 192)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 193, 192)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet() bool {
	return message.BoolVLS(m.Ptr, 193, 192)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet() bool {
	return message.BoolVLS(m.Ptr, 193, 192)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataUpdate) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() bool {
	return message.BoolVLS(m.Unsafe(), 194, 193)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() bool {
	return message.BoolVLS(m.Unsafe(), 194, 193)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataUpdatePointer) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() bool {
	return message.BoolVLS(m.Ptr, 194, 193)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() bool {
	return message.BoolVLS(m.Ptr, 194, 193)
}

// UseMasterFirm_UsePercentOfMarginFullOverrideIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_UsePercentOfMarginFullOverrideIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 195, 194)
}

// UseMasterFirm_UsePercentOfMarginFullOverrideIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_UsePercentOfMarginFullOverrideIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 195, 194)
}

// UseMasterFirm_UsePercentOfMarginFullOverrideIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginFullOverrideIsSet() bool {
	return message.BoolVLS(m.Ptr, 195, 194)
}

// UseMasterFirm_UsePercentOfMarginFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginFullOverrideIsSet() bool {
	return message.BoolVLS(m.Ptr, 195, 194)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataUpdate) UseMasterFirm_UsePercentOfMarginFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 196, 195)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_UsePercentOfMarginFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 196, 195)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginFullOverride() bool {
	return message.BoolVLS(m.Ptr, 196, 195)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginFullOverride() bool {
	return message.BoolVLS(m.Ptr, 196, 195)
}

// UseMasterFirm_TradeFeesFullOverrideIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_TradeFeesFullOverrideIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 197, 196)
}

// UseMasterFirm_TradeFeesFullOverrideIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_TradeFeesFullOverrideIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 197, 196)
}

// UseMasterFirm_TradeFeesFullOverrideIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradeFeesFullOverrideIsSet() bool {
	return message.BoolVLS(m.Ptr, 197, 196)
}

// UseMasterFirm_TradeFeesFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradeFeesFullOverrideIsSet() bool {
	return message.BoolVLS(m.Ptr, 197, 196)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataUpdate) UseMasterFirm_TradeFeesFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 198, 197)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_TradeFeesFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 198, 197)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradeFeesFullOverride() bool {
	return message.BoolVLS(m.Ptr, 198, 197)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradeFeesFullOverride() bool {
	return message.BoolVLS(m.Ptr, 198, 197)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet
func (m TradeAccountDataUpdate) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 199, 198)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 199, 198)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet() bool {
	return message.BoolVLS(m.Ptr, 199, 198)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet() bool {
	return message.BoolVLS(m.Ptr, 199, 198)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataUpdate) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 200, 199)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataUpdateBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() bool {
	return message.BoolVLS(m.Unsafe(), 200, 199)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() bool {
	return message.BoolVLS(m.Ptr, 200, 199)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() bool {
	return message.BoolVLS(m.Ptr, 200, 199)
}

// LiquidationOnlyModeIsSet
func (m TradeAccountDataUpdate) LiquidationOnlyModeIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 201, 200)
}

// LiquidationOnlyModeIsSet
func (m TradeAccountDataUpdateBuilder) LiquidationOnlyModeIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 201, 200)
}

// LiquidationOnlyModeIsSet
func (m TradeAccountDataUpdatePointer) LiquidationOnlyModeIsSet() bool {
	return message.BoolVLS(m.Ptr, 201, 200)
}

// LiquidationOnlyModeIsSet
func (m TradeAccountDataUpdatePointerBuilder) LiquidationOnlyModeIsSet() bool {
	return message.BoolVLS(m.Ptr, 201, 200)
}

// LiquidationOnlyMode
func (m TradeAccountDataUpdate) LiquidationOnlyMode() uint8 {
	return message.Uint8VLS(m.Unsafe(), 202, 201)
}

// LiquidationOnlyMode
func (m TradeAccountDataUpdateBuilder) LiquidationOnlyMode() uint8 {
	return message.Uint8VLS(m.Unsafe(), 202, 201)
}

// LiquidationOnlyMode
func (m TradeAccountDataUpdatePointer) LiquidationOnlyMode() uint8 {
	return message.Uint8VLS(m.Ptr, 202, 201)
}

// LiquidationOnlyMode
func (m TradeAccountDataUpdatePointerBuilder) LiquidationOnlyMode() uint8 {
	return message.Uint8VLS(m.Ptr, 202, 201)
}

// CustomerOrFirmIsSet
func (m TradeAccountDataUpdate) CustomerOrFirmIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 203, 202)
}

// CustomerOrFirmIsSet
func (m TradeAccountDataUpdateBuilder) CustomerOrFirmIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 203, 202)
}

// CustomerOrFirmIsSet
func (m TradeAccountDataUpdatePointer) CustomerOrFirmIsSet() bool {
	return message.BoolVLS(m.Ptr, 203, 202)
}

// CustomerOrFirmIsSet
func (m TradeAccountDataUpdatePointerBuilder) CustomerOrFirmIsSet() bool {
	return message.BoolVLS(m.Ptr, 203, 202)
}

// CustomerOrFirm
func (m TradeAccountDataUpdate) CustomerOrFirm() uint8 {
	return message.Uint8VLS(m.Unsafe(), 204, 203)
}

// CustomerOrFirm
func (m TradeAccountDataUpdateBuilder) CustomerOrFirm() uint8 {
	return message.Uint8VLS(m.Unsafe(), 204, 203)
}

// CustomerOrFirm
func (m TradeAccountDataUpdatePointer) CustomerOrFirm() uint8 {
	return message.Uint8VLS(m.Ptr, 204, 203)
}

// CustomerOrFirm
func (m TradeAccountDataUpdatePointerBuilder) CustomerOrFirm() uint8 {
	return message.Uint8VLS(m.Ptr, 204, 203)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet
func (m TradeAccountDataUpdate) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 205, 204)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet
func (m TradeAccountDataUpdateBuilder) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 205, 204)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet
func (m TradeAccountDataUpdatePointer) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet() bool {
	return message.BoolVLS(m.Ptr, 205, 204)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet() bool {
	return message.BoolVLS(m.Ptr, 205, 204)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataUpdate) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 206, 205)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataUpdateBuilder) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 206, 205)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataUpdatePointer) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() uint8 {
	return message.Uint8VLS(m.Ptr, 206, 205)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() uint8 {
	return message.Uint8VLS(m.Ptr, 206, 205)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet
func (m TradeAccountDataUpdate) MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 207, 206)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet
func (m TradeAccountDataUpdateBuilder) MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 207, 206)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet
func (m TradeAccountDataUpdatePointer) MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet() bool {
	return message.BoolVLS(m.Ptr, 207, 206)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet() bool {
	return message.BoolVLS(m.Ptr, 207, 206)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataUpdate) MasterFirm_FlattenCancelWhenUnderMinimumAccountValue() uint8 {
	return message.Uint8VLS(m.Unsafe(), 208, 207)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataUpdateBuilder) MasterFirm_FlattenCancelWhenUnderMinimumAccountValue() uint8 {
	return message.Uint8VLS(m.Unsafe(), 208, 207)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataUpdatePointer) MasterFirm_FlattenCancelWhenUnderMinimumAccountValue() uint8 {
	return message.Uint8VLS(m.Ptr, 208, 207)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_FlattenCancelWhenUnderMinimumAccountValue() uint8 {
	return message.Uint8VLS(m.Ptr, 208, 207)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdate) MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 209, 208)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdateBuilder) MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 209, 208)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointer) MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet() bool {
	return message.BoolVLS(m.Ptr, 209, 208)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet() bool {
	return message.BoolVLS(m.Ptr, 209, 208)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataUpdate) MasterFirm_FlattenCancelWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 210, 209)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataUpdateBuilder) MasterFirm_FlattenCancelWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Unsafe(), 210, 209)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointer) MasterFirm_FlattenCancelWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 210, 209)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_FlattenCancelWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 210, 209)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet
func (m TradeAccountDataUpdate) MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 211, 210)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet
func (m TradeAccountDataUpdateBuilder) MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 211, 210)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet
func (m TradeAccountDataUpdatePointer) MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet() bool {
	return message.BoolVLS(m.Ptr, 211, 210)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet() bool {
	return message.BoolVLS(m.Ptr, 211, 210)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataUpdate) MasterFirm_FlattenCancelWhenUnderMarginEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 212, 211)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataUpdateBuilder) MasterFirm_FlattenCancelWhenUnderMarginEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 212, 211)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataUpdatePointer) MasterFirm_FlattenCancelWhenUnderMarginEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 212, 211)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_FlattenCancelWhenUnderMarginEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 212, 211)
}

// MasterFirm_MaximumOrderQuantityIsSet
func (m TradeAccountDataUpdate) MasterFirm_MaximumOrderQuantityIsSet() uint32 {
	return message.Uint32VLS(m.Unsafe(), 216, 212)
}

// MasterFirm_MaximumOrderQuantityIsSet
func (m TradeAccountDataUpdateBuilder) MasterFirm_MaximumOrderQuantityIsSet() uint32 {
	return message.Uint32VLS(m.Unsafe(), 216, 212)
}

// MasterFirm_MaximumOrderQuantityIsSet
func (m TradeAccountDataUpdatePointer) MasterFirm_MaximumOrderQuantityIsSet() uint32 {
	return message.Uint32VLS(m.Ptr, 216, 212)
}

// MasterFirm_MaximumOrderQuantityIsSet
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_MaximumOrderQuantityIsSet() uint32 {
	return message.Uint32VLS(m.Ptr, 216, 212)
}

// MasterFirm_MaximumOrderQuantity
func (m TradeAccountDataUpdate) MasterFirm_MaximumOrderQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 220, 216)
}

// MasterFirm_MaximumOrderQuantity
func (m TradeAccountDataUpdateBuilder) MasterFirm_MaximumOrderQuantity() uint32 {
	return message.Uint32VLS(m.Unsafe(), 220, 216)
}

// MasterFirm_MaximumOrderQuantity
func (m TradeAccountDataUpdatePointer) MasterFirm_MaximumOrderQuantity() uint32 {
	return message.Uint32VLS(m.Ptr, 220, 216)
}

// MasterFirm_MaximumOrderQuantity
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_MaximumOrderQuantity() uint32 {
	return message.Uint32VLS(m.Ptr, 220, 216)
}

// ExchangeTraderIdIsSet
func (m TradeAccountDataUpdate) ExchangeTraderIdIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 221, 220)
}

// ExchangeTraderIdIsSet
func (m TradeAccountDataUpdateBuilder) ExchangeTraderIdIsSet() bool {
	return message.BoolVLS(m.Unsafe(), 221, 220)
}

// ExchangeTraderIdIsSet
func (m TradeAccountDataUpdatePointer) ExchangeTraderIdIsSet() bool {
	return message.BoolVLS(m.Ptr, 221, 220)
}

// ExchangeTraderIdIsSet
func (m TradeAccountDataUpdatePointerBuilder) ExchangeTraderIdIsSet() bool {
	return message.BoolVLS(m.Ptr, 221, 220)
}

// ExchangeTraderId
func (m TradeAccountDataUpdate) ExchangeTraderId() string {
	return message.StringVLS(m.Unsafe(), 225, 221)
}

// ExchangeTraderId
func (m TradeAccountDataUpdateBuilder) ExchangeTraderId() string {
	return message.StringVLS(m.Unsafe(), 225, 221)
}

// ExchangeTraderId
func (m TradeAccountDataUpdatePointer) ExchangeTraderId() string {
	return message.StringVLS(m.Ptr, 225, 221)
}

// ExchangeTraderId
func (m TradeAccountDataUpdatePointerBuilder) ExchangeTraderId() string {
	return message.StringVLS(m.Ptr, 225, 221)
}

// SetRequestID
func (m TradeAccountDataUpdateBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m TradeAccountDataUpdatePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetIsNewAccount
func (m TradeAccountDataUpdateBuilder) SetIsNewAccount(value bool) {
	message.SetBoolVLS(m.Unsafe(), 11, 10, value)
}

// SetIsNewAccount
func (m TradeAccountDataUpdatePointerBuilder) SetIsNewAccount(value bool) {
	message.SetBoolVLS(m.Ptr, 11, 10, value)
}

// SetNewAccountAuthorizedUsername
func (m *TradeAccountDataUpdateBuilder) SetNewAccountAuthorizedUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 15, 11, value)
}

// SetNewAccountAuthorizedUsername
func (m *TradeAccountDataUpdatePointerBuilder) SetNewAccountAuthorizedUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetTradeAccount
func (m *TradeAccountDataUpdateBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 19, 15, value)
}

// SetTradeAccount
func (m *TradeAccountDataUpdatePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 19, 15, value)
}

// SetCurrencyCodeIsSet
func (m TradeAccountDataUpdateBuilder) SetCurrencyCodeIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 20, 19, value)
}

// SetCurrencyCodeIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetCurrencyCodeIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 20, 19, value)
}

// SetCurrencyCode
func (m *TradeAccountDataUpdateBuilder) SetCurrencyCode(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetCurrencyCode
func (m *TradeAccountDataUpdatePointerBuilder) SetCurrencyCode(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetDailyNetLossLimitInAccountCurrencyIsSet
func (m TradeAccountDataUpdateBuilder) SetDailyNetLossLimitInAccountCurrencyIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 25, 24, value)
}

// SetDailyNetLossLimitInAccountCurrencyIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetDailyNetLossLimitInAccountCurrencyIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 25, 24, value)
}

// SetDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataUpdateBuilder) SetDailyNetLossLimitInAccountCurrency(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 29, 25, value)
}

// SetDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataUpdatePointerBuilder) SetDailyNetLossLimitInAccountCurrency(value float32) {
	message.SetFloat32VLS(m.Ptr, 29, 25, value)
}

// SetPercentOfCashBalanceForDailyNetLossLimitIsSet
func (m TradeAccountDataUpdateBuilder) SetPercentOfCashBalanceForDailyNetLossLimitIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 30, 29, value)
}

// SetPercentOfCashBalanceForDailyNetLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetPercentOfCashBalanceForDailyNetLossLimitIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 30, 29, value)
}

// SetPercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataUpdateBuilder) SetPercentOfCashBalanceForDailyNetLossLimit(value int32) {
	message.SetInt32VLS(m.Unsafe(), 34, 30, value)
}

// SetPercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataUpdatePointerBuilder) SetPercentOfCashBalanceForDailyNetLossLimit(value int32) {
	message.SetInt32VLS(m.Ptr, 34, 30, value)
}

// SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet
func (m TradeAccountDataUpdateBuilder) SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 35, 34, value)
}

// SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 35, 34, value)
}

// SetUseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataUpdateBuilder) SetUseTrailingAccountValueToNotAllowIncreaseInPositions(value bool) {
	message.SetBoolVLS(m.Unsafe(), 36, 35, value)
}

// SetUseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataUpdatePointerBuilder) SetUseTrailingAccountValueToNotAllowIncreaseInPositions(value bool) {
	message.SetBoolVLS(m.Ptr, 36, 35, value)
}

// SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdateBuilder) SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 37, 36, value)
}

// SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 37, 36, value)
}

// SetDoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdateBuilder) SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 38, 37, value)
}

// SetDoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointerBuilder) SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Ptr, 38, 37, value)
}

// SetFlattenPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdateBuilder) SetFlattenPositionsAtDailyLossLimitIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 39, 38, value)
}

// SetFlattenPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetFlattenPositionsAtDailyLossLimitIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 39, 38, value)
}

// SetFlattenPositionsAtDailyLossLimit
func (m TradeAccountDataUpdateBuilder) SetFlattenPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 40, 39, value)
}

// SetFlattenPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointerBuilder) SetFlattenPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Ptr, 40, 39, value)
}

// SetClosePositionsAtEndOfDayIsSet
func (m TradeAccountDataUpdateBuilder) SetClosePositionsAtEndOfDayIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 41, 40, value)
}

// SetClosePositionsAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetClosePositionsAtEndOfDayIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 41, 40, value)
}

// SetClosePositionsAtEndOfDay
func (m TradeAccountDataUpdateBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 42, 41, value)
}

// SetClosePositionsAtEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 42, 41, value)
}

// SetFlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdateBuilder) SetFlattenPositionsWhenUnderMarginIntradayIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 43, 42, value)
}

// SetFlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetFlattenPositionsWhenUnderMarginIntradayIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 43, 42, value)
}

// SetFlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdateBuilder) SetFlattenPositionsWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 44, 43, value)
}

// SetFlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointerBuilder) SetFlattenPositionsWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Ptr, 44, 43, value)
}

// SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdateBuilder) SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 45, 44, value)
}

// SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 45, 44, value)
}

// SetFlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdateBuilder) SetFlattenPositionsWhenUnderMarginAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 46, 45, value)
}

// SetFlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) SetFlattenPositionsWhenUnderMarginAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 46, 45, value)
}

// SetSenderSubIDIsSet
func (m TradeAccountDataUpdateBuilder) SetSenderSubIDIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 47, 46, value)
}

// SetSenderSubIDIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetSenderSubIDIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 47, 46, value)
}

// SetSenderSubID
func (m *TradeAccountDataUpdateBuilder) SetSenderSubID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 51, 47, value)
}

// SetSenderSubID
func (m *TradeAccountDataUpdatePointerBuilder) SetSenderSubID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 51, 47, value)
}

// SetSenderLocationIdIsSet
func (m TradeAccountDataUpdateBuilder) SetSenderLocationIdIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 52, 51, value)
}

// SetSenderLocationIdIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetSenderLocationIdIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 52, 51, value)
}

// SetSenderLocationId
func (m *TradeAccountDataUpdateBuilder) SetSenderLocationId(value string) {
	message.SetStringVLS(&m.VLSBuilder, 56, 52, value)
}

// SetSenderLocationId
func (m *TradeAccountDataUpdatePointerBuilder) SetSenderLocationId(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 56, 52, value)
}

// SetSelfMatchPreventionIDIsSet
func (m TradeAccountDataUpdateBuilder) SetSelfMatchPreventionIDIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 57, 56, value)
}

// SetSelfMatchPreventionIDIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetSelfMatchPreventionIDIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 57, 56, value)
}

// SetSelfMatchPreventionID
func (m *TradeAccountDataUpdateBuilder) SetSelfMatchPreventionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 61, 57, value)
}

// SetSelfMatchPreventionID
func (m *TradeAccountDataUpdatePointerBuilder) SetSelfMatchPreventionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 61, 57, value)
}

// SetCTICodeIsSet
func (m TradeAccountDataUpdateBuilder) SetCTICodeIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 62, 61, value)
}

// SetCTICodeIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetCTICodeIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 62, 61, value)
}

// SetCTICode
func (m TradeAccountDataUpdateBuilder) SetCTICode(value int32) {
	message.SetInt32VLS(m.Unsafe(), 66, 62, value)
}

// SetCTICode
func (m TradeAccountDataUpdatePointerBuilder) SetCTICode(value int32) {
	message.SetInt32VLS(m.Ptr, 66, 62, value)
}

// SetTradeAccountIsReadOnlyIsSet
func (m TradeAccountDataUpdateBuilder) SetTradeAccountIsReadOnlyIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 67, 66, value)
}

// SetTradeAccountIsReadOnlyIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetTradeAccountIsReadOnlyIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 67, 66, value)
}

// SetTradeAccountIsReadOnly
func (m TradeAccountDataUpdateBuilder) SetTradeAccountIsReadOnly(value bool) {
	message.SetBoolVLS(m.Unsafe(), 68, 67, value)
}

// SetTradeAccountIsReadOnly
func (m TradeAccountDataUpdatePointerBuilder) SetTradeAccountIsReadOnly(value bool) {
	message.SetBoolVLS(m.Ptr, 68, 67, value)
}

// SetMaximumGlobalPositionQuantityIsSet
func (m TradeAccountDataUpdateBuilder) SetMaximumGlobalPositionQuantityIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 69, 68, value)
}

// SetMaximumGlobalPositionQuantityIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMaximumGlobalPositionQuantityIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 69, 68, value)
}

// SetMaximumGlobalPositionQuantity
func (m TradeAccountDataUpdateBuilder) SetMaximumGlobalPositionQuantity(value int32) {
	message.SetInt32VLS(m.Unsafe(), 73, 69, value)
}

// SetMaximumGlobalPositionQuantity
func (m TradeAccountDataUpdatePointerBuilder) SetMaximumGlobalPositionQuantity(value int32) {
	message.SetInt32VLS(m.Ptr, 73, 69, value)
}

// SetTradeFeePerContractIsSet
func (m TradeAccountDataUpdateBuilder) SetTradeFeePerContractIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 74, 73, value)
}

// SetTradeFeePerContractIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetTradeFeePerContractIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 74, 73, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataUpdateBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 82, 74, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataUpdatePointerBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Ptr, 82, 74, value)
}

// SetTradeFeePerShareIsSet
func (m TradeAccountDataUpdateBuilder) SetTradeFeePerShareIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 83, 82, value)
}

// SetTradeFeePerShareIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetTradeFeePerShareIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 83, 82, value)
}

// SetTradeFeePerShare
func (m TradeAccountDataUpdateBuilder) SetTradeFeePerShare(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 91, 83, value)
}

// SetTradeFeePerShare
func (m TradeAccountDataUpdatePointerBuilder) SetTradeFeePerShare(value float64) {
	message.SetFloat64VLS(m.Ptr, 91, 83, value)
}

// SetRequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdateBuilder) SetRequireSufficientMarginForNewPositionsIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 92, 91, value)
}

// SetRequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetRequireSufficientMarginForNewPositionsIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 92, 91, value)
}

// SetRequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdateBuilder) SetRequireSufficientMarginForNewPositions(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 93, 92, value)
}

// SetRequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdatePointerBuilder) SetRequireSufficientMarginForNewPositions(value uint8) {
	message.SetUint8VLS(m.Ptr, 93, 92, value)
}

// SetUsePercentOfMarginIsSet
func (m TradeAccountDataUpdateBuilder) SetUsePercentOfMarginIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 94, 93, value)
}

// SetUsePercentOfMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUsePercentOfMarginIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 94, 93, value)
}

// SetUsePercentOfMargin
func (m TradeAccountDataUpdateBuilder) SetUsePercentOfMargin(value int32) {
	message.SetInt32VLS(m.Unsafe(), 98, 94, value)
}

// SetUsePercentOfMargin
func (m TradeAccountDataUpdatePointerBuilder) SetUsePercentOfMargin(value int32) {
	message.SetInt32VLS(m.Ptr, 98, 94, value)
}

// SetUsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdateBuilder) SetUsePercentOfMarginForDayTradingIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 99, 98, value)
}

// SetUsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUsePercentOfMarginForDayTradingIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 99, 98, value)
}

// SetUsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdateBuilder) SetUsePercentOfMarginForDayTrading(value int32) {
	message.SetInt32VLS(m.Unsafe(), 103, 99, value)
}

// SetUsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdatePointerBuilder) SetUsePercentOfMarginForDayTrading(value int32) {
	message.SetInt32VLS(m.Ptr, 103, 99, value)
}

// SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdateBuilder) SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 104, 103, value)
}

// SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 104, 103, value)
}

// SetMaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdateBuilder) SetMaximumAllowedAccountBalanceForPositionsAsPercentage(value int32) {
	message.SetInt32VLS(m.Unsafe(), 108, 104, value)
}

// SetMaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdatePointerBuilder) SetMaximumAllowedAccountBalanceForPositionsAsPercentage(value int32) {
	message.SetInt32VLS(m.Ptr, 108, 104, value)
}

// SetFirmIDIsSet
func (m TradeAccountDataUpdateBuilder) SetFirmIDIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 109, 108, value)
}

// SetFirmIDIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetFirmIDIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 109, 108, value)
}

// SetFirmID
func (m *TradeAccountDataUpdateBuilder) SetFirmID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 113, 109, value)
}

// SetFirmID
func (m *TradeAccountDataUpdatePointerBuilder) SetFirmID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 113, 109, value)
}

// SetTradingIsDisabledIsSet
func (m TradeAccountDataUpdateBuilder) SetTradingIsDisabledIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 114, 113, value)
}

// SetTradingIsDisabledIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetTradingIsDisabledIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 114, 113, value)
}

// SetTradingIsDisabled
func (m TradeAccountDataUpdateBuilder) SetTradingIsDisabled(value bool) {
	message.SetBoolVLS(m.Unsafe(), 115, 114, value)
}

// SetTradingIsDisabled
func (m TradeAccountDataUpdatePointerBuilder) SetTradingIsDisabled(value bool) {
	message.SetBoolVLS(m.Ptr, 115, 114, value)
}

// SetDescriptiveNameIsSet
func (m TradeAccountDataUpdateBuilder) SetDescriptiveNameIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 116, 115, value)
}

// SetDescriptiveNameIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetDescriptiveNameIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 116, 115, value)
}

// SetDescriptiveName
func (m *TradeAccountDataUpdateBuilder) SetDescriptiveName(value string) {
	message.SetStringVLS(&m.VLSBuilder, 120, 116, value)
}

// SetDescriptiveName
func (m *TradeAccountDataUpdatePointerBuilder) SetDescriptiveName(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 120, 116, value)
}

// SetIsMasterFirmControlAccountIsSet
func (m TradeAccountDataUpdateBuilder) SetIsMasterFirmControlAccountIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 121, 120, value)
}

// SetIsMasterFirmControlAccountIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetIsMasterFirmControlAccountIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 121, 120, value)
}

// SetIsMasterFirmControlAccount
func (m TradeAccountDataUpdateBuilder) SetIsMasterFirmControlAccount(value bool) {
	message.SetBoolVLS(m.Unsafe(), 122, 121, value)
}

// SetIsMasterFirmControlAccount
func (m TradeAccountDataUpdatePointerBuilder) SetIsMasterFirmControlAccount(value bool) {
	message.SetBoolVLS(m.Ptr, 122, 121, value)
}

// SetMinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdateBuilder) SetMinimumRequiredAccountValueIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 123, 122, value)
}

// SetMinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMinimumRequiredAccountValueIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 123, 122, value)
}

// SetMinimumRequiredAccountValue
func (m TradeAccountDataUpdateBuilder) SetMinimumRequiredAccountValue(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 131, 123, value)
}

// SetMinimumRequiredAccountValue
func (m TradeAccountDataUpdatePointerBuilder) SetMinimumRequiredAccountValue(value float64) {
	message.SetFloat64VLS(m.Ptr, 131, 123, value)
}

// SetBeginTimeForDayMarginIsSet
func (m TradeAccountDataUpdateBuilder) SetBeginTimeForDayMarginIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 132, 131, value)
}

// SetBeginTimeForDayMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetBeginTimeForDayMarginIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 132, 131, value)
}

// SetBeginTimeForDayMargin
func (m TradeAccountDataUpdateBuilder) SetBeginTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Unsafe(), 140, 132, value)
}

// SetBeginTimeForDayMargin
func (m TradeAccountDataUpdatePointerBuilder) SetBeginTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Ptr, 140, 132, value)
}

// SetEndTimeForDayMarginIsSet
func (m TradeAccountDataUpdateBuilder) SetEndTimeForDayMarginIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 141, 140, value)
}

// SetEndTimeForDayMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetEndTimeForDayMarginIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 141, 140, value)
}

// SetEndTimeForDayMargin
func (m TradeAccountDataUpdateBuilder) SetEndTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Unsafe(), 149, 141, value)
}

// SetEndTimeForDayMargin
func (m TradeAccountDataUpdatePointerBuilder) SetEndTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Ptr, 149, 141, value)
}

// SetDayMarginTimeZoneIsSet
func (m TradeAccountDataUpdateBuilder) SetDayMarginTimeZoneIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 150, 149, value)
}

// SetDayMarginTimeZoneIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetDayMarginTimeZoneIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 150, 149, value)
}

// SetDayMarginTimeZone
func (m *TradeAccountDataUpdateBuilder) SetDayMarginTimeZone(value string) {
	message.SetStringVLS(&m.VLSBuilder, 154, 150, value)
}

// SetDayMarginTimeZone
func (m *TradeAccountDataUpdatePointerBuilder) SetDayMarginTimeZone(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 154, 150, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 155, 154, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 155, 154, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(value bool) {
	message.SetBoolVLS(m.Unsafe(), 156, 155, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(value bool) {
	message.SetBoolVLS(m.Ptr, 156, 155, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 157, 156, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 157, 156, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(value bool) {
	message.SetBoolVLS(m.Unsafe(), 158, 157, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(value bool) {
	message.SetBoolVLS(m.Ptr, 158, 157, value)
}

// SetUseMasterFirm_SymbolLimitsArrayIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_SymbolLimitsArrayIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 159, 158, value)
}

// SetUseMasterFirm_SymbolLimitsArrayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_SymbolLimitsArrayIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 159, 158, value)
}

// SetUseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_SymbolLimitsArray(value bool) {
	message.SetBoolVLS(m.Unsafe(), 160, 159, value)
}

// SetUseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_SymbolLimitsArray(value bool) {
	message.SetBoolVLS(m.Ptr, 160, 159, value)
}

// SetUseMasterFirm_TradeFeesIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_TradeFeesIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 161, 160, value)
}

// SetUseMasterFirm_TradeFeesIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradeFeesIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 161, 160, value)
}

// SetUseMasterFirm_TradeFees
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_TradeFees(value bool) {
	message.SetBoolVLS(m.Unsafe(), 162, 161, value)
}

// SetUseMasterFirm_TradeFees
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradeFees(value bool) {
	message.SetBoolVLS(m.Ptr, 162, 161, value)
}

// SetUseMasterFirm_TradeFeePerShareIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_TradeFeePerShareIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 163, 162, value)
}

// SetUseMasterFirm_TradeFeePerShareIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradeFeePerShareIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 163, 162, value)
}

// SetUseMasterFirm_TradeFeePerShare
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_TradeFeePerShare(value bool) {
	message.SetBoolVLS(m.Unsafe(), 164, 163, value)
}

// SetUseMasterFirm_TradeFeePerShare
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradeFeePerShare(value bool) {
	message.SetBoolVLS(m.Ptr, 164, 163, value)
}

// SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 165, 164, value)
}

// SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 165, 164, value)
}

// SetUseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_RequireSufficientMarginForNewPositions(value bool) {
	message.SetBoolVLS(m.Unsafe(), 166, 165, value)
}

// SetUseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_RequireSufficientMarginForNewPositions(value bool) {
	message.SetBoolVLS(m.Ptr, 166, 165, value)
}

// SetUseMasterFirm_UsePercentOfMarginIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_UsePercentOfMarginIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 167, 166, value)
}

// SetUseMasterFirm_UsePercentOfMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 167, 166, value)
}

// SetUseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_UsePercentOfMargin(value bool) {
	message.SetBoolVLS(m.Unsafe(), 168, 167, value)
}

// SetUseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMargin(value bool) {
	message.SetBoolVLS(m.Ptr, 168, 167, value)
}

// SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 169, 168, value)
}

// SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 169, 168, value)
}

// SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(value bool) {
	message.SetBoolVLS(m.Unsafe(), 170, 169, value)
}

// SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(value bool) {
	message.SetBoolVLS(m.Ptr, 170, 169, value)
}

// SetUseMasterFirm_MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_MinimumRequiredAccountValueIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 171, 170, value)
}

// SetUseMasterFirm_MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_MinimumRequiredAccountValueIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 171, 170, value)
}

// SetUseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_MinimumRequiredAccountValue(value bool) {
	message.SetBoolVLS(m.Unsafe(), 172, 171, value)
}

// SetUseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_MinimumRequiredAccountValue(value bool) {
	message.SetBoolVLS(m.Ptr, 172, 171, value)
}

// SetUseMasterFirm_MarginTimeSettingsIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_MarginTimeSettingsIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 173, 172, value)
}

// SetUseMasterFirm_MarginTimeSettingsIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_MarginTimeSettingsIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 173, 172, value)
}

// SetUseMasterFirm_MarginTimeSettings
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_MarginTimeSettings(value bool) {
	message.SetBoolVLS(m.Unsafe(), 174, 173, value)
}

// SetUseMasterFirm_MarginTimeSettings
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_MarginTimeSettings(value bool) {
	message.SetBoolVLS(m.Ptr, 174, 173, value)
}

// SetUseMasterFirm_TradingIsDisabledIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_TradingIsDisabledIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 175, 174, value)
}

// SetUseMasterFirm_TradingIsDisabledIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradingIsDisabledIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 175, 174, value)
}

// SetUseMasterFirm_TradingIsDisabled
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_TradingIsDisabled(value bool) {
	message.SetBoolVLS(m.Unsafe(), 176, 175, value)
}

// SetUseMasterFirm_TradingIsDisabled
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradingIsDisabled(value bool) {
	message.SetBoolVLS(m.Ptr, 176, 175, value)
}

// SetIsTradeStatisticsPublicallySharedIsSet
func (m TradeAccountDataUpdateBuilder) SetIsTradeStatisticsPublicallySharedIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 177, 176, value)
}

// SetIsTradeStatisticsPublicallySharedIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetIsTradeStatisticsPublicallySharedIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 177, 176, value)
}

// SetIsTradeStatisticsPublicallyShared
func (m TradeAccountDataUpdateBuilder) SetIsTradeStatisticsPublicallyShared(value bool) {
	message.SetBoolVLS(m.Unsafe(), 178, 177, value)
}

// SetIsTradeStatisticsPublicallyShared
func (m TradeAccountDataUpdatePointerBuilder) SetIsTradeStatisticsPublicallyShared(value bool) {
	message.SetBoolVLS(m.Ptr, 178, 177, value)
}

// SetIsReadOnlyFollowingRequestsAllowedIsSet
func (m TradeAccountDataUpdateBuilder) SetIsReadOnlyFollowingRequestsAllowedIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 179, 178, value)
}

// SetIsReadOnlyFollowingRequestsAllowedIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetIsReadOnlyFollowingRequestsAllowedIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 179, 178, value)
}

// SetIsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataUpdateBuilder) SetIsReadOnlyFollowingRequestsAllowed(value bool) {
	message.SetBoolVLS(m.Unsafe(), 180, 179, value)
}

// SetIsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataUpdatePointerBuilder) SetIsReadOnlyFollowingRequestsAllowed(value bool) {
	message.SetBoolVLS(m.Ptr, 180, 179, value)
}

// SetIsTradeAccountSharingAllowedIsSet
func (m TradeAccountDataUpdateBuilder) SetIsTradeAccountSharingAllowedIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 181, 180, value)
}

// SetIsTradeAccountSharingAllowedIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetIsTradeAccountSharingAllowedIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 181, 180, value)
}

// SetIsTradeAccountSharingAllowed
func (m TradeAccountDataUpdateBuilder) SetIsTradeAccountSharingAllowed(value bool) {
	message.SetBoolVLS(m.Unsafe(), 182, 181, value)
}

// SetIsTradeAccountSharingAllowed
func (m TradeAccountDataUpdatePointerBuilder) SetIsTradeAccountSharingAllowed(value bool) {
	message.SetBoolVLS(m.Ptr, 182, 181, value)
}

// SetReadOnlyShareToAllSCUsernamesIsSet
func (m TradeAccountDataUpdateBuilder) SetReadOnlyShareToAllSCUsernamesIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 183, 182, value)
}

// SetReadOnlyShareToAllSCUsernamesIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetReadOnlyShareToAllSCUsernamesIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 183, 182, value)
}

// SetReadOnlyShareToAllSCUsernames
func (m TradeAccountDataUpdateBuilder) SetReadOnlyShareToAllSCUsernames(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 184, 183, value)
}

// SetReadOnlyShareToAllSCUsernames
func (m TradeAccountDataUpdatePointerBuilder) SetReadOnlyShareToAllSCUsernames(value uint8) {
	message.SetUint8VLS(m.Ptr, 184, 183, value)
}

// SetUseMasterFirm_SymbolCommissionsArrayIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_SymbolCommissionsArrayIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 185, 184, value)
}

// SetUseMasterFirm_SymbolCommissionsArrayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_SymbolCommissionsArrayIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 185, 184, value)
}

// SetUseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_SymbolCommissionsArray(value bool) {
	message.SetBoolVLS(m.Unsafe(), 186, 185, value)
}

// SetUseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_SymbolCommissionsArray(value bool) {
	message.SetBoolVLS(m.Ptr, 186, 185, value)
}

// SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 187, 186, value)
}

// SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 187, 186, value)
}

// SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(value bool) {
	message.SetBoolVLS(m.Unsafe(), 188, 187, value)
}

// SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(value bool) {
	message.SetBoolVLS(m.Ptr, 188, 187, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 189, 188, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 189, 188, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTrading(value bool) {
	message.SetBoolVLS(m.Unsafe(), 190, 189, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTrading(value bool) {
	message.SetBoolVLS(m.Ptr, 190, 189, value)
}

// SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 191, 190, value)
}

// SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 191, 190, value)
}

// SetUseMasterFirm_SymbolCommissionsArrayFullOverride
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_SymbolCommissionsArrayFullOverride(value bool) {
	message.SetBoolVLS(m.Unsafe(), 192, 191, value)
}

// SetUseMasterFirm_SymbolCommissionsArrayFullOverride
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_SymbolCommissionsArrayFullOverride(value bool) {
	message.SetBoolVLS(m.Ptr, 192, 191, value)
}

// SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 193, 192, value)
}

// SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 193, 192, value)
}

// SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(value bool) {
	message.SetBoolVLS(m.Unsafe(), 194, 193, value)
}

// SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(value bool) {
	message.SetBoolVLS(m.Ptr, 194, 193, value)
}

// SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 195, 194, value)
}

// SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 195, 194, value)
}

// SetUseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_UsePercentOfMarginFullOverride(value bool) {
	message.SetBoolVLS(m.Unsafe(), 196, 195, value)
}

// SetUseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginFullOverride(value bool) {
	message.SetBoolVLS(m.Ptr, 196, 195, value)
}

// SetUseMasterFirm_TradeFeesFullOverrideIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_TradeFeesFullOverrideIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 197, 196, value)
}

// SetUseMasterFirm_TradeFeesFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradeFeesFullOverrideIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 197, 196, value)
}

// SetUseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_TradeFeesFullOverride(value bool) {
	message.SetBoolVLS(m.Unsafe(), 198, 197, value)
}

// SetUseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradeFeesFullOverride(value bool) {
	message.SetBoolVLS(m.Ptr, 198, 197, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 199, 198, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 199, 198, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataUpdateBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(value bool) {
	message.SetBoolVLS(m.Unsafe(), 200, 199, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(value bool) {
	message.SetBoolVLS(m.Ptr, 200, 199, value)
}

// SetLiquidationOnlyModeIsSet
func (m TradeAccountDataUpdateBuilder) SetLiquidationOnlyModeIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 201, 200, value)
}

// SetLiquidationOnlyModeIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetLiquidationOnlyModeIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 201, 200, value)
}

// SetLiquidationOnlyMode
func (m TradeAccountDataUpdateBuilder) SetLiquidationOnlyMode(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 202, 201, value)
}

// SetLiquidationOnlyMode
func (m TradeAccountDataUpdatePointerBuilder) SetLiquidationOnlyMode(value uint8) {
	message.SetUint8VLS(m.Ptr, 202, 201, value)
}

// SetCustomerOrFirmIsSet
func (m TradeAccountDataUpdateBuilder) SetCustomerOrFirmIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 203, 202, value)
}

// SetCustomerOrFirmIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetCustomerOrFirmIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 203, 202, value)
}

// SetCustomerOrFirm
func (m TradeAccountDataUpdateBuilder) SetCustomerOrFirm(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 204, 203, value)
}

// SetCustomerOrFirm
func (m TradeAccountDataUpdatePointerBuilder) SetCustomerOrFirm(value uint8) {
	message.SetUint8VLS(m.Ptr, 204, 203, value)
}

// SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet
func (m TradeAccountDataUpdateBuilder) SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 205, 204, value)
}

// SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 205, 204, value)
}

// SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataUpdateBuilder) SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 206, 205, value)
}

// SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(value uint8) {
	message.SetUint8VLS(m.Ptr, 206, 205, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet
func (m TradeAccountDataUpdateBuilder) SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 207, 206, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 207, 206, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataUpdateBuilder) SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 208, 207, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(value uint8) {
	message.SetUint8VLS(m.Ptr, 208, 207, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdateBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 209, 208, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 209, 208, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataUpdateBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 210, 209, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Ptr, 210, 209, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet
func (m TradeAccountDataUpdateBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 211, 210, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 211, 210, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataUpdateBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 212, 211, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 212, 211, value)
}

// SetMasterFirm_MaximumOrderQuantityIsSet
func (m TradeAccountDataUpdateBuilder) SetMasterFirm_MaximumOrderQuantityIsSet(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 216, 212, value)
}

// SetMasterFirm_MaximumOrderQuantityIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_MaximumOrderQuantityIsSet(value uint32) {
	message.SetUint32VLS(m.Ptr, 216, 212, value)
}

// SetMasterFirm_MaximumOrderQuantity
func (m TradeAccountDataUpdateBuilder) SetMasterFirm_MaximumOrderQuantity(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 220, 216, value)
}

// SetMasterFirm_MaximumOrderQuantity
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_MaximumOrderQuantity(value uint32) {
	message.SetUint32VLS(m.Ptr, 220, 216, value)
}

// SetExchangeTraderIdIsSet
func (m TradeAccountDataUpdateBuilder) SetExchangeTraderIdIsSet(value bool) {
	message.SetBoolVLS(m.Unsafe(), 221, 220, value)
}

// SetExchangeTraderIdIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetExchangeTraderIdIsSet(value bool) {
	message.SetBoolVLS(m.Ptr, 221, 220, value)
}

// SetExchangeTraderId
func (m *TradeAccountDataUpdateBuilder) SetExchangeTraderId(value string) {
	message.SetStringVLS(&m.VLSBuilder, 225, 221, value)
}

// SetExchangeTraderId
func (m *TradeAccountDataUpdatePointerBuilder) SetExchangeTraderId(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 225, 221, value)
}

// Copy
func (m TradeAccountDataUpdate) Copy(to TradeAccountDataUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsNewAccount(m.IsNewAccount())
	to.SetNewAccountAuthorizedUsername(m.NewAccountAuthorizedUsername())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCodeIsSet(m.CurrencyCodeIsSet())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetDailyNetLossLimitInAccountCurrencyIsSet(m.DailyNetLossLimitInAccountCurrencyIsSet())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimitIsSet(m.FlattenPositionsAtDailyLossLimitIsSet())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDayIsSet(m.ClosePositionsAtEndOfDayIsSet())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntradayIsSet(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubIDIsSet(m.SenderSubIDIsSet())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationIdIsSet(m.SenderLocationIdIsSet())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionIDIsSet(m.SelfMatchPreventionIDIsSet())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICodeIsSet(m.CTICodeIsSet())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnlyIsSet(m.TradeAccountIsReadOnlyIsSet())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantityIsSet(m.MaximumGlobalPositionQuantityIsSet())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContractIsSet(m.TradeFeePerContractIsSet())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShareIsSet(m.TradeFeePerShareIsSet())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositionsIsSet(m.RequireSufficientMarginForNewPositionsIsSet())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMarginIsSet(m.UsePercentOfMarginIsSet())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTradingIsSet(m.UsePercentOfMarginForDayTradingIsSet())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmIDIsSet(m.FirmIDIsSet())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabledIsSet(m.TradingIsDisabledIsSet())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveNameIsSet(m.DescriptiveNameIsSet())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccountIsSet(m.IsMasterFirmControlAccountIsSet())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValueIsSet(m.MinimumRequiredAccountValueIsSet())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMarginIsSet(m.BeginTimeForDayMarginIsSet())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMarginIsSet(m.EndTimeForDayMarginIsSet())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZoneIsSet(m.DayMarginTimeZoneIsSet())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArrayIsSet(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFeesIsSet(m.UseMasterFirm_TradeFeesIsSet())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShareIsSet(m.UseMasterFirm_TradeFeePerShareIsSet())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMarginIsSet(m.UseMasterFirm_UsePercentOfMarginIsSet())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettingsIsSet(m.UseMasterFirm_MarginTimeSettingsIsSet())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabledIsSet(m.UseMasterFirm_TradingIsDisabledIsSet())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallySharedIsSet(m.IsTradeStatisticsPublicallySharedIsSet())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowedIsSet(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowedIsSet(m.IsTradeAccountSharingAllowedIsSet())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernamesIsSet(m.ReadOnlyShareToAllSCUsernamesIsSet())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArrayIsSet(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_TradeFeesFullOverrideIsSet(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetLiquidationOnlyModeIsSet(m.LiquidationOnlyModeIsSet())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetCustomerOrFirmIsSet(m.CustomerOrFirmIsSet())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantityIsSet(m.MasterFirm_MaximumOrderQuantityIsSet())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetExchangeTraderIdIsSet(m.ExchangeTraderIdIsSet())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// Copy
func (m TradeAccountDataUpdateBuilder) Copy(to TradeAccountDataUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsNewAccount(m.IsNewAccount())
	to.SetNewAccountAuthorizedUsername(m.NewAccountAuthorizedUsername())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCodeIsSet(m.CurrencyCodeIsSet())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetDailyNetLossLimitInAccountCurrencyIsSet(m.DailyNetLossLimitInAccountCurrencyIsSet())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimitIsSet(m.FlattenPositionsAtDailyLossLimitIsSet())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDayIsSet(m.ClosePositionsAtEndOfDayIsSet())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntradayIsSet(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubIDIsSet(m.SenderSubIDIsSet())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationIdIsSet(m.SenderLocationIdIsSet())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionIDIsSet(m.SelfMatchPreventionIDIsSet())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICodeIsSet(m.CTICodeIsSet())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnlyIsSet(m.TradeAccountIsReadOnlyIsSet())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantityIsSet(m.MaximumGlobalPositionQuantityIsSet())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContractIsSet(m.TradeFeePerContractIsSet())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShareIsSet(m.TradeFeePerShareIsSet())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositionsIsSet(m.RequireSufficientMarginForNewPositionsIsSet())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMarginIsSet(m.UsePercentOfMarginIsSet())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTradingIsSet(m.UsePercentOfMarginForDayTradingIsSet())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmIDIsSet(m.FirmIDIsSet())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabledIsSet(m.TradingIsDisabledIsSet())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveNameIsSet(m.DescriptiveNameIsSet())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccountIsSet(m.IsMasterFirmControlAccountIsSet())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValueIsSet(m.MinimumRequiredAccountValueIsSet())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMarginIsSet(m.BeginTimeForDayMarginIsSet())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMarginIsSet(m.EndTimeForDayMarginIsSet())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZoneIsSet(m.DayMarginTimeZoneIsSet())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArrayIsSet(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFeesIsSet(m.UseMasterFirm_TradeFeesIsSet())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShareIsSet(m.UseMasterFirm_TradeFeePerShareIsSet())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMarginIsSet(m.UseMasterFirm_UsePercentOfMarginIsSet())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettingsIsSet(m.UseMasterFirm_MarginTimeSettingsIsSet())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabledIsSet(m.UseMasterFirm_TradingIsDisabledIsSet())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallySharedIsSet(m.IsTradeStatisticsPublicallySharedIsSet())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowedIsSet(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowedIsSet(m.IsTradeAccountSharingAllowedIsSet())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernamesIsSet(m.ReadOnlyShareToAllSCUsernamesIsSet())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArrayIsSet(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_TradeFeesFullOverrideIsSet(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetLiquidationOnlyModeIsSet(m.LiquidationOnlyModeIsSet())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetCustomerOrFirmIsSet(m.CustomerOrFirmIsSet())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantityIsSet(m.MasterFirm_MaximumOrderQuantityIsSet())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetExchangeTraderIdIsSet(m.ExchangeTraderIdIsSet())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// CopyPointer
func (m TradeAccountDataUpdate) CopyPointer(to TradeAccountDataUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsNewAccount(m.IsNewAccount())
	to.SetNewAccountAuthorizedUsername(m.NewAccountAuthorizedUsername())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCodeIsSet(m.CurrencyCodeIsSet())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetDailyNetLossLimitInAccountCurrencyIsSet(m.DailyNetLossLimitInAccountCurrencyIsSet())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimitIsSet(m.FlattenPositionsAtDailyLossLimitIsSet())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDayIsSet(m.ClosePositionsAtEndOfDayIsSet())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntradayIsSet(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubIDIsSet(m.SenderSubIDIsSet())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationIdIsSet(m.SenderLocationIdIsSet())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionIDIsSet(m.SelfMatchPreventionIDIsSet())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICodeIsSet(m.CTICodeIsSet())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnlyIsSet(m.TradeAccountIsReadOnlyIsSet())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantityIsSet(m.MaximumGlobalPositionQuantityIsSet())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContractIsSet(m.TradeFeePerContractIsSet())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShareIsSet(m.TradeFeePerShareIsSet())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositionsIsSet(m.RequireSufficientMarginForNewPositionsIsSet())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMarginIsSet(m.UsePercentOfMarginIsSet())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTradingIsSet(m.UsePercentOfMarginForDayTradingIsSet())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmIDIsSet(m.FirmIDIsSet())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabledIsSet(m.TradingIsDisabledIsSet())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveNameIsSet(m.DescriptiveNameIsSet())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccountIsSet(m.IsMasterFirmControlAccountIsSet())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValueIsSet(m.MinimumRequiredAccountValueIsSet())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMarginIsSet(m.BeginTimeForDayMarginIsSet())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMarginIsSet(m.EndTimeForDayMarginIsSet())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZoneIsSet(m.DayMarginTimeZoneIsSet())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArrayIsSet(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFeesIsSet(m.UseMasterFirm_TradeFeesIsSet())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShareIsSet(m.UseMasterFirm_TradeFeePerShareIsSet())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMarginIsSet(m.UseMasterFirm_UsePercentOfMarginIsSet())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettingsIsSet(m.UseMasterFirm_MarginTimeSettingsIsSet())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabledIsSet(m.UseMasterFirm_TradingIsDisabledIsSet())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallySharedIsSet(m.IsTradeStatisticsPublicallySharedIsSet())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowedIsSet(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowedIsSet(m.IsTradeAccountSharingAllowedIsSet())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernamesIsSet(m.ReadOnlyShareToAllSCUsernamesIsSet())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArrayIsSet(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_TradeFeesFullOverrideIsSet(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetLiquidationOnlyModeIsSet(m.LiquidationOnlyModeIsSet())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetCustomerOrFirmIsSet(m.CustomerOrFirmIsSet())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantityIsSet(m.MasterFirm_MaximumOrderQuantityIsSet())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetExchangeTraderIdIsSet(m.ExchangeTraderIdIsSet())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// CopyPointer
func (m TradeAccountDataUpdateBuilder) CopyPointer(to TradeAccountDataUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsNewAccount(m.IsNewAccount())
	to.SetNewAccountAuthorizedUsername(m.NewAccountAuthorizedUsername())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCodeIsSet(m.CurrencyCodeIsSet())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetDailyNetLossLimitInAccountCurrencyIsSet(m.DailyNetLossLimitInAccountCurrencyIsSet())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimitIsSet(m.FlattenPositionsAtDailyLossLimitIsSet())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDayIsSet(m.ClosePositionsAtEndOfDayIsSet())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntradayIsSet(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubIDIsSet(m.SenderSubIDIsSet())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationIdIsSet(m.SenderLocationIdIsSet())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionIDIsSet(m.SelfMatchPreventionIDIsSet())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICodeIsSet(m.CTICodeIsSet())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnlyIsSet(m.TradeAccountIsReadOnlyIsSet())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantityIsSet(m.MaximumGlobalPositionQuantityIsSet())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContractIsSet(m.TradeFeePerContractIsSet())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShareIsSet(m.TradeFeePerShareIsSet())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositionsIsSet(m.RequireSufficientMarginForNewPositionsIsSet())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMarginIsSet(m.UsePercentOfMarginIsSet())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTradingIsSet(m.UsePercentOfMarginForDayTradingIsSet())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmIDIsSet(m.FirmIDIsSet())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabledIsSet(m.TradingIsDisabledIsSet())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveNameIsSet(m.DescriptiveNameIsSet())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccountIsSet(m.IsMasterFirmControlAccountIsSet())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValueIsSet(m.MinimumRequiredAccountValueIsSet())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMarginIsSet(m.BeginTimeForDayMarginIsSet())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMarginIsSet(m.EndTimeForDayMarginIsSet())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZoneIsSet(m.DayMarginTimeZoneIsSet())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArrayIsSet(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFeesIsSet(m.UseMasterFirm_TradeFeesIsSet())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShareIsSet(m.UseMasterFirm_TradeFeePerShareIsSet())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMarginIsSet(m.UseMasterFirm_UsePercentOfMarginIsSet())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettingsIsSet(m.UseMasterFirm_MarginTimeSettingsIsSet())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabledIsSet(m.UseMasterFirm_TradingIsDisabledIsSet())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallySharedIsSet(m.IsTradeStatisticsPublicallySharedIsSet())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowedIsSet(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowedIsSet(m.IsTradeAccountSharingAllowedIsSet())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernamesIsSet(m.ReadOnlyShareToAllSCUsernamesIsSet())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArrayIsSet(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_TradeFeesFullOverrideIsSet(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetLiquidationOnlyModeIsSet(m.LiquidationOnlyModeIsSet())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetCustomerOrFirmIsSet(m.CustomerOrFirmIsSet())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantityIsSet(m.MasterFirm_MaximumOrderQuantityIsSet())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetExchangeTraderIdIsSet(m.ExchangeTraderIdIsSet())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// Copy
func (m TradeAccountDataUpdatePointer) Copy(to TradeAccountDataUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsNewAccount(m.IsNewAccount())
	to.SetNewAccountAuthorizedUsername(m.NewAccountAuthorizedUsername())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCodeIsSet(m.CurrencyCodeIsSet())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetDailyNetLossLimitInAccountCurrencyIsSet(m.DailyNetLossLimitInAccountCurrencyIsSet())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimitIsSet(m.FlattenPositionsAtDailyLossLimitIsSet())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDayIsSet(m.ClosePositionsAtEndOfDayIsSet())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntradayIsSet(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubIDIsSet(m.SenderSubIDIsSet())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationIdIsSet(m.SenderLocationIdIsSet())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionIDIsSet(m.SelfMatchPreventionIDIsSet())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICodeIsSet(m.CTICodeIsSet())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnlyIsSet(m.TradeAccountIsReadOnlyIsSet())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantityIsSet(m.MaximumGlobalPositionQuantityIsSet())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContractIsSet(m.TradeFeePerContractIsSet())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShareIsSet(m.TradeFeePerShareIsSet())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositionsIsSet(m.RequireSufficientMarginForNewPositionsIsSet())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMarginIsSet(m.UsePercentOfMarginIsSet())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTradingIsSet(m.UsePercentOfMarginForDayTradingIsSet())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmIDIsSet(m.FirmIDIsSet())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabledIsSet(m.TradingIsDisabledIsSet())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveNameIsSet(m.DescriptiveNameIsSet())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccountIsSet(m.IsMasterFirmControlAccountIsSet())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValueIsSet(m.MinimumRequiredAccountValueIsSet())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMarginIsSet(m.BeginTimeForDayMarginIsSet())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMarginIsSet(m.EndTimeForDayMarginIsSet())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZoneIsSet(m.DayMarginTimeZoneIsSet())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArrayIsSet(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFeesIsSet(m.UseMasterFirm_TradeFeesIsSet())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShareIsSet(m.UseMasterFirm_TradeFeePerShareIsSet())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMarginIsSet(m.UseMasterFirm_UsePercentOfMarginIsSet())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettingsIsSet(m.UseMasterFirm_MarginTimeSettingsIsSet())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabledIsSet(m.UseMasterFirm_TradingIsDisabledIsSet())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallySharedIsSet(m.IsTradeStatisticsPublicallySharedIsSet())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowedIsSet(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowedIsSet(m.IsTradeAccountSharingAllowedIsSet())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernamesIsSet(m.ReadOnlyShareToAllSCUsernamesIsSet())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArrayIsSet(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_TradeFeesFullOverrideIsSet(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetLiquidationOnlyModeIsSet(m.LiquidationOnlyModeIsSet())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetCustomerOrFirmIsSet(m.CustomerOrFirmIsSet())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantityIsSet(m.MasterFirm_MaximumOrderQuantityIsSet())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetExchangeTraderIdIsSet(m.ExchangeTraderIdIsSet())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// Copy
func (m TradeAccountDataUpdatePointerBuilder) Copy(to TradeAccountDataUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsNewAccount(m.IsNewAccount())
	to.SetNewAccountAuthorizedUsername(m.NewAccountAuthorizedUsername())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCodeIsSet(m.CurrencyCodeIsSet())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetDailyNetLossLimitInAccountCurrencyIsSet(m.DailyNetLossLimitInAccountCurrencyIsSet())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimitIsSet(m.FlattenPositionsAtDailyLossLimitIsSet())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDayIsSet(m.ClosePositionsAtEndOfDayIsSet())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntradayIsSet(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubIDIsSet(m.SenderSubIDIsSet())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationIdIsSet(m.SenderLocationIdIsSet())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionIDIsSet(m.SelfMatchPreventionIDIsSet())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICodeIsSet(m.CTICodeIsSet())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnlyIsSet(m.TradeAccountIsReadOnlyIsSet())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantityIsSet(m.MaximumGlobalPositionQuantityIsSet())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContractIsSet(m.TradeFeePerContractIsSet())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShareIsSet(m.TradeFeePerShareIsSet())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositionsIsSet(m.RequireSufficientMarginForNewPositionsIsSet())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMarginIsSet(m.UsePercentOfMarginIsSet())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTradingIsSet(m.UsePercentOfMarginForDayTradingIsSet())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmIDIsSet(m.FirmIDIsSet())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabledIsSet(m.TradingIsDisabledIsSet())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveNameIsSet(m.DescriptiveNameIsSet())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccountIsSet(m.IsMasterFirmControlAccountIsSet())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValueIsSet(m.MinimumRequiredAccountValueIsSet())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMarginIsSet(m.BeginTimeForDayMarginIsSet())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMarginIsSet(m.EndTimeForDayMarginIsSet())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZoneIsSet(m.DayMarginTimeZoneIsSet())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArrayIsSet(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFeesIsSet(m.UseMasterFirm_TradeFeesIsSet())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShareIsSet(m.UseMasterFirm_TradeFeePerShareIsSet())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMarginIsSet(m.UseMasterFirm_UsePercentOfMarginIsSet())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettingsIsSet(m.UseMasterFirm_MarginTimeSettingsIsSet())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabledIsSet(m.UseMasterFirm_TradingIsDisabledIsSet())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallySharedIsSet(m.IsTradeStatisticsPublicallySharedIsSet())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowedIsSet(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowedIsSet(m.IsTradeAccountSharingAllowedIsSet())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernamesIsSet(m.ReadOnlyShareToAllSCUsernamesIsSet())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArrayIsSet(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_TradeFeesFullOverrideIsSet(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetLiquidationOnlyModeIsSet(m.LiquidationOnlyModeIsSet())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetCustomerOrFirmIsSet(m.CustomerOrFirmIsSet())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantityIsSet(m.MasterFirm_MaximumOrderQuantityIsSet())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetExchangeTraderIdIsSet(m.ExchangeTraderIdIsSet())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// CopyPointer
func (m TradeAccountDataUpdatePointer) CopyPointer(to TradeAccountDataUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsNewAccount(m.IsNewAccount())
	to.SetNewAccountAuthorizedUsername(m.NewAccountAuthorizedUsername())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCodeIsSet(m.CurrencyCodeIsSet())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetDailyNetLossLimitInAccountCurrencyIsSet(m.DailyNetLossLimitInAccountCurrencyIsSet())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimitIsSet(m.FlattenPositionsAtDailyLossLimitIsSet())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDayIsSet(m.ClosePositionsAtEndOfDayIsSet())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntradayIsSet(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubIDIsSet(m.SenderSubIDIsSet())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationIdIsSet(m.SenderLocationIdIsSet())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionIDIsSet(m.SelfMatchPreventionIDIsSet())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICodeIsSet(m.CTICodeIsSet())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnlyIsSet(m.TradeAccountIsReadOnlyIsSet())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantityIsSet(m.MaximumGlobalPositionQuantityIsSet())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContractIsSet(m.TradeFeePerContractIsSet())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShareIsSet(m.TradeFeePerShareIsSet())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositionsIsSet(m.RequireSufficientMarginForNewPositionsIsSet())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMarginIsSet(m.UsePercentOfMarginIsSet())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTradingIsSet(m.UsePercentOfMarginForDayTradingIsSet())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmIDIsSet(m.FirmIDIsSet())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabledIsSet(m.TradingIsDisabledIsSet())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveNameIsSet(m.DescriptiveNameIsSet())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccountIsSet(m.IsMasterFirmControlAccountIsSet())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValueIsSet(m.MinimumRequiredAccountValueIsSet())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMarginIsSet(m.BeginTimeForDayMarginIsSet())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMarginIsSet(m.EndTimeForDayMarginIsSet())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZoneIsSet(m.DayMarginTimeZoneIsSet())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArrayIsSet(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFeesIsSet(m.UseMasterFirm_TradeFeesIsSet())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShareIsSet(m.UseMasterFirm_TradeFeePerShareIsSet())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMarginIsSet(m.UseMasterFirm_UsePercentOfMarginIsSet())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettingsIsSet(m.UseMasterFirm_MarginTimeSettingsIsSet())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabledIsSet(m.UseMasterFirm_TradingIsDisabledIsSet())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallySharedIsSet(m.IsTradeStatisticsPublicallySharedIsSet())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowedIsSet(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowedIsSet(m.IsTradeAccountSharingAllowedIsSet())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernamesIsSet(m.ReadOnlyShareToAllSCUsernamesIsSet())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArrayIsSet(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_TradeFeesFullOverrideIsSet(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetLiquidationOnlyModeIsSet(m.LiquidationOnlyModeIsSet())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetCustomerOrFirmIsSet(m.CustomerOrFirmIsSet())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantityIsSet(m.MasterFirm_MaximumOrderQuantityIsSet())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetExchangeTraderIdIsSet(m.ExchangeTraderIdIsSet())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}

// CopyPointer
func (m TradeAccountDataUpdatePointerBuilder) CopyPointer(to TradeAccountDataUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsNewAccount(m.IsNewAccount())
	to.SetNewAccountAuthorizedUsername(m.NewAccountAuthorizedUsername())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCurrencyCodeIsSet(m.CurrencyCodeIsSet())
	to.SetCurrencyCode(m.CurrencyCode())
	to.SetDailyNetLossLimitInAccountCurrencyIsSet(m.DailyNetLossLimitInAccountCurrencyIsSet())
	to.SetDailyNetLossLimitInAccountCurrency(m.DailyNetLossLimitInAccountCurrency())
	to.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	to.SetPercentOfCashBalanceForDailyNetLossLimit(m.PercentOfCashBalanceForDailyNetLossLimit())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	to.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetFlattenPositionsAtDailyLossLimitIsSet(m.FlattenPositionsAtDailyLossLimitIsSet())
	to.SetFlattenPositionsAtDailyLossLimit(m.FlattenPositionsAtDailyLossLimit())
	to.SetClosePositionsAtEndOfDayIsSet(m.ClosePositionsAtEndOfDayIsSet())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetFlattenPositionsWhenUnderMarginIntradayIsSet(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetFlattenPositionsWhenUnderMarginIntraday(m.FlattenPositionsWhenUnderMarginIntraday())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetFlattenPositionsWhenUnderMarginAtEndOfDay(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetSenderSubIDIsSet(m.SenderSubIDIsSet())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationIdIsSet(m.SenderLocationIdIsSet())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionIDIsSet(m.SelfMatchPreventionIDIsSet())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICodeIsSet(m.CTICodeIsSet())
	to.SetCTICode(m.CTICode())
	to.SetTradeAccountIsReadOnlyIsSet(m.TradeAccountIsReadOnlyIsSet())
	to.SetTradeAccountIsReadOnly(m.TradeAccountIsReadOnly())
	to.SetMaximumGlobalPositionQuantityIsSet(m.MaximumGlobalPositionQuantityIsSet())
	to.SetMaximumGlobalPositionQuantity(m.MaximumGlobalPositionQuantity())
	to.SetTradeFeePerContractIsSet(m.TradeFeePerContractIsSet())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
	to.SetTradeFeePerShareIsSet(m.TradeFeePerShareIsSet())
	to.SetTradeFeePerShare(m.TradeFeePerShare())
	to.SetRequireSufficientMarginForNewPositionsIsSet(m.RequireSufficientMarginForNewPositionsIsSet())
	to.SetRequireSufficientMarginForNewPositions(m.RequireSufficientMarginForNewPositions())
	to.SetUsePercentOfMarginIsSet(m.UsePercentOfMarginIsSet())
	to.SetUsePercentOfMargin(m.UsePercentOfMargin())
	to.SetUsePercentOfMarginForDayTradingIsSet(m.UsePercentOfMarginForDayTradingIsSet())
	to.SetUsePercentOfMarginForDayTrading(m.UsePercentOfMarginForDayTrading())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetFirmIDIsSet(m.FirmIDIsSet())
	to.SetFirmID(m.FirmID())
	to.SetTradingIsDisabledIsSet(m.TradingIsDisabledIsSet())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescriptiveNameIsSet(m.DescriptiveNameIsSet())
	to.SetDescriptiveName(m.DescriptiveName())
	to.SetIsMasterFirmControlAccountIsSet(m.IsMasterFirmControlAccountIsSet())
	to.SetIsMasterFirmControlAccount(m.IsMasterFirmControlAccount())
	to.SetMinimumRequiredAccountValueIsSet(m.MinimumRequiredAccountValueIsSet())
	to.SetMinimumRequiredAccountValue(m.MinimumRequiredAccountValue())
	to.SetBeginTimeForDayMarginIsSet(m.BeginTimeForDayMarginIsSet())
	to.SetBeginTimeForDayMargin(m.BeginTimeForDayMargin())
	to.SetEndTimeForDayMarginIsSet(m.EndTimeForDayMarginIsSet())
	to.SetEndTimeForDayMargin(m.EndTimeForDayMargin())
	to.SetDayMarginTimeZoneIsSet(m.DayMarginTimeZoneIsSet())
	to.SetDayMarginTimeZone(m.DayMarginTimeZone())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	to.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	to.SetUseMasterFirm_SymbolLimitsArrayIsSet(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	to.SetUseMasterFirm_SymbolLimitsArray(m.UseMasterFirm_SymbolLimitsArray())
	to.SetUseMasterFirm_TradeFeesIsSet(m.UseMasterFirm_TradeFeesIsSet())
	to.SetUseMasterFirm_TradeFees(m.UseMasterFirm_TradeFees())
	to.SetUseMasterFirm_TradeFeePerShareIsSet(m.UseMasterFirm_TradeFeePerShareIsSet())
	to.SetUseMasterFirm_TradeFeePerShare(m.UseMasterFirm_TradeFeePerShare())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	to.SetUseMasterFirm_RequireSufficientMarginForNewPositions(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	to.SetUseMasterFirm_UsePercentOfMarginIsSet(m.UseMasterFirm_UsePercentOfMarginIsSet())
	to.SetUseMasterFirm_UsePercentOfMargin(m.UseMasterFirm_UsePercentOfMargin())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	to.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	to.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	to.SetUseMasterFirm_MinimumRequiredAccountValue(m.UseMasterFirm_MinimumRequiredAccountValue())
	to.SetUseMasterFirm_MarginTimeSettingsIsSet(m.UseMasterFirm_MarginTimeSettingsIsSet())
	to.SetUseMasterFirm_MarginTimeSettings(m.UseMasterFirm_MarginTimeSettings())
	to.SetUseMasterFirm_TradingIsDisabledIsSet(m.UseMasterFirm_TradingIsDisabledIsSet())
	to.SetUseMasterFirm_TradingIsDisabled(m.UseMasterFirm_TradingIsDisabled())
	to.SetIsTradeStatisticsPublicallySharedIsSet(m.IsTradeStatisticsPublicallySharedIsSet())
	to.SetIsTradeStatisticsPublicallyShared(m.IsTradeStatisticsPublicallyShared())
	to.SetIsReadOnlyFollowingRequestsAllowedIsSet(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	to.SetIsReadOnlyFollowingRequestsAllowed(m.IsReadOnlyFollowingRequestsAllowed())
	to.SetIsTradeAccountSharingAllowedIsSet(m.IsTradeAccountSharingAllowedIsSet())
	to.SetIsTradeAccountSharingAllowed(m.IsTradeAccountSharingAllowed())
	to.SetReadOnlyShareToAllSCUsernamesIsSet(m.ReadOnlyShareToAllSCUsernamesIsSet())
	to.SetReadOnlyShareToAllSCUsernames(m.ReadOnlyShareToAllSCUsernames())
	to.SetUseMasterFirm_SymbolCommissionsArrayIsSet(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArray(m.UseMasterFirm_SymbolCommissionsArray())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	to.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTrading(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	to.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	to.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginFullOverride(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	to.SetUseMasterFirm_TradeFeesFullOverrideIsSet(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	to.SetUseMasterFirm_TradeFeesFullOverride(m.UseMasterFirm_TradeFeesFullOverride())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	to.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	to.SetLiquidationOnlyModeIsSet(m.LiquidationOnlyModeIsSet())
	to.SetLiquidationOnlyMode(m.LiquidationOnlyMode())
	to.SetCustomerOrFirmIsSet(m.CustomerOrFirmIsSet())
	to.SetCustomerOrFirm(m.CustomerOrFirm())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	to.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	to.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	to.SetMasterFirm_MaximumOrderQuantityIsSet(m.MasterFirm_MaximumOrderQuantityIsSet())
	to.SetMasterFirm_MaximumOrderQuantity(m.MasterFirm_MaximumOrderQuantity())
	to.SetExchangeTraderIdIsSet(m.ExchangeTraderIdIsSet())
	to.SetExchangeTraderId(m.ExchangeTraderId())
}
