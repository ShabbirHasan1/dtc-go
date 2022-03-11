package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataUpdatePointer struct {
	message.VLSPointer
}

type TradeAccountDataUpdatePointerBuilder struct {
	message.VLSPointerBuilder
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
//     IsNewAccount                                                            = 0
//     NewAccountAuthorizedUsername                                            = ""
//     TradeAccount                                                            = ""
//     CurrencyCodeIsSet                                                       = 0
//     CurrencyCode                                                            = ""
//     DailyNetLossLimitInAccountCurrencyIsSet                                 = 0
//     DailyNetLossLimitInAccountCurrency                                      = 0
//     PercentOfCashBalanceForDailyNetLossLimitIsSet                           = 0
//     PercentOfCashBalanceForDailyNetLossLimit                                = 0
//     UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet               = 0
//     UseTrailingAccountValueToNotAllowIncreaseInPositions                    = 0
//     DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet                      = 0
//     DoNotAllowIncreaseInPositionsAtDailyLossLimit                           = 0
//     FlattenPositionsAtDailyLossLimitIsSet                                   = 0
//     FlattenPositionsAtDailyLossLimit                                        = 0
//     ClosePositionsAtEndOfDayIsSet                                           = 0
//     ClosePositionsAtEndOfDay                                                = 0
//     FlattenPositionsWhenUnderMarginIntradayIsSet                            = 0
//     FlattenPositionsWhenUnderMarginIntraday                                 = 0
//     FlattenPositionsWhenUnderMarginAtEndOfDayIsSet                          = 0
//     FlattenPositionsWhenUnderMarginAtEndOfDay                               = 0
//     SenderSubIDIsSet                                                        = 0
//     SenderSubID                                                             = ""
//     SenderLocationIdIsSet                                                   = 0
//     SenderLocationId                                                        = ""
//     SelfMatchPreventionIDIsSet                                              = 0
//     SelfMatchPreventionID                                                   = ""
//     CTICodeIsSet                                                            = 0
//     CTICode                                                                 = 0
//     TradeAccountIsReadOnlyIsSet                                             = 0
//     TradeAccountIsReadOnly                                                  = 0
//     MaximumGlobalPositionQuantityIsSet                                      = 0
//     MaximumGlobalPositionQuantity                                           = 0
//     TradeFeePerContractIsSet                                                = 0
//     TradeFeePerContract                                                     = 0
//     TradeFeePerShareIsSet                                                   = 0
//     TradeFeePerShare                                                        = 0
//     RequireSufficientMarginForNewPositionsIsSet                             = 0
//     RequireSufficientMarginForNewPositions                                  = 1
//     UsePercentOfMarginIsSet                                                 = 0
//     UsePercentOfMargin                                                      = 100
//     UsePercentOfMarginForDayTradingIsSet                                    = 0
//     UsePercentOfMarginForDayTrading                                         = 100
//     MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet               = 0
//     MaximumAllowedAccountBalanceForPositionsAsPercentage                    = 100
//     FirmIDIsSet                                                             = 0
//     FirmID                                                                  = ""
//     TradingIsDisabledIsSet                                                  = 0
//     TradingIsDisabled                                                       = 0
//     DescriptiveNameIsSet                                                    = 0
//     DescriptiveName                                                         = ""
//     IsMasterFirmControlAccountIsSet                                         = 0
//     IsMasterFirmControlAccount                                              = 0
//     MinimumRequiredAccountValueIsSet                                        = 0
//     MinimumRequiredAccountValue                                             = 0
//     BeginTimeForDayMarginIsSet                                              = 0
//     BeginTimeForDayMargin                                                   = 0
//     EndTimeForDayMarginIsSet                                                = 0
//     EndTimeForDayMargin                                                     = 0
//     DayMarginTimeZoneIsSet                                                  = 0
//     DayMarginTimeZone                                                       = ""
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet              = 0
//     UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday                   = 0
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet            = 0
//     UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay                 = 0
//     UseMasterFirm_SymbolLimitsArrayIsSet                                    = 0
//     UseMasterFirm_SymbolLimitsArray                                         = 0
//     UseMasterFirm_TradeFeesIsSet                                            = 0
//     UseMasterFirm_TradeFees                                                 = 0
//     UseMasterFirm_TradeFeePerShareIsSet                                     = 0
//     UseMasterFirm_TradeFeePerShare                                          = 0
//     UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet               = 0
//     UseMasterFirm_RequireSufficientMarginForNewPositions                    = 0
//     UseMasterFirm_UsePercentOfMarginIsSet                                   = 0
//     UseMasterFirm_UsePercentOfMargin                                        = 0
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet = 0
//     UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage      = 0
//     UseMasterFirm_MinimumRequiredAccountValueIsSet                          = 0
//     UseMasterFirm_MinimumRequiredAccountValue                               = 0
//     UseMasterFirm_MarginTimeSettingsIsSet                                   = 0
//     UseMasterFirm_MarginTimeSettings                                        = 0
//     UseMasterFirm_TradingIsDisabledIsSet                                    = 0
//     UseMasterFirm_TradingIsDisabled                                         = 0
//     IsTradeStatisticsPublicallySharedIsSet                                  = 0
//     IsTradeStatisticsPublicallyShared                                       = 0
//     IsReadOnlyFollowingRequestsAllowedIsSet                                 = 0
//     IsReadOnlyFollowingRequestsAllowed                                      = 0
//     IsTradeAccountSharingAllowedIsSet                                       = 0
//     IsTradeAccountSharingAllowed                                            = 0
//     ReadOnlyShareToAllSCUsernamesIsSet                                      = 0
//     ReadOnlyShareToAllSCUsernames                                           = 0
//     UseMasterFirm_SymbolCommissionsArrayIsSet                               = 0
//     UseMasterFirm_SymbolCommissionsArray                                    = 0
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet        = 0
//     UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit             = 0
//     UseMasterFirm_UsePercentOfMarginForDayTradingIsSet                      = 0
//     UseMasterFirm_UsePercentOfMarginForDayTrading                           = 0
//     UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet                   = 0
//     UseMasterFirm_SymbolCommissionsArrayFullOverride                        = 0
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet         = 0
//     UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders              = 0
//     UseMasterFirm_UsePercentOfMarginFullOverrideIsSet                       = 0
//     UseMasterFirm_UsePercentOfMarginFullOverride                            = 0
//     UseMasterFirm_TradeFeesFullOverrideIsSet                                = 0
//     UseMasterFirm_TradeFeesFullOverride                                     = 0
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet          = 0
//     UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride               = 0
//     LiquidationOnlyModeIsSet                                                = 0
//     LiquidationOnlyMode                                                     = 0
//     CustomerOrFirmIsSet                                                     = 0
//     CustomerOrFirm                                                          = 0
//     MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet               = 0
//     MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet                    = 0
//     MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet               = 0
//     MasterFirm_FlattenCancelWhenUnderMinimumAccountValue                    = 0
//     MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet                    = 0
//     MasterFirm_FlattenCancelWhenUnderMarginIntraday                         = 0
//     MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet                    = 0
//     MasterFirm_FlattenCancelWhenUnderMarginEndOfDay                         = 0
//     MasterFirm_MaximumOrderQuantityIsSet                                    = 0
//     MasterFirm_MaximumOrderQuantity                                         = 0
//     ExchangeTraderIdIsSet                                                   = 0
//     ExchangeTraderId                                                        = ""
// }
func (m TradeAccountDataUpdatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataUpdateDefault)
}

// ToBuilder
func (m TradeAccountDataUpdatePointer) ToBuilder() TradeAccountDataUpdatePointerBuilder {
	return TradeAccountDataUpdatePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataUpdatePointerBuilder) Finish() TradeAccountDataUpdatePointer {
	return TradeAccountDataUpdatePointer{m.VLSPointerBuilder.Finish()}
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
func (m TradeAccountDataUpdatePointer) IsNewAccount() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// IsNewAccount
func (m TradeAccountDataUpdatePointerBuilder) IsNewAccount() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
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
func (m TradeAccountDataUpdatePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 19, 15)
}

// TradeAccount
func (m TradeAccountDataUpdatePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 19, 15)
}

// CurrencyCodeIsSet
func (m TradeAccountDataUpdatePointer) CurrencyCodeIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 20, 19)
}

// CurrencyCodeIsSet
func (m TradeAccountDataUpdatePointerBuilder) CurrencyCodeIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 20, 19)
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
func (m TradeAccountDataUpdatePointer) DailyNetLossLimitInAccountCurrencyIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 25, 24)
}

// DailyNetLossLimitInAccountCurrencyIsSet
func (m TradeAccountDataUpdatePointerBuilder) DailyNetLossLimitInAccountCurrencyIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 25, 24)
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
func (m TradeAccountDataUpdatePointer) PercentOfCashBalanceForDailyNetLossLimitIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 30, 29)
}

// PercentOfCashBalanceForDailyNetLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) PercentOfCashBalanceForDailyNetLossLimitIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 30, 29)
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
func (m TradeAccountDataUpdatePointer) UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 35, 34)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 35, 34)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataUpdatePointer) UseTrailingAccountValueToNotAllowIncreaseInPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 36, 35)
}

// UseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataUpdatePointerBuilder) UseTrailingAccountValueToNotAllowIncreaseInPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 36, 35)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointer) DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 37, 36)
}

// DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 37, 36)
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
func (m TradeAccountDataUpdatePointer) FlattenPositionsAtDailyLossLimitIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 39, 38)
}

// FlattenPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) FlattenPositionsAtDailyLossLimitIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 39, 38)
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
func (m TradeAccountDataUpdatePointer) ClosePositionsAtEndOfDayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 41, 40)
}

// ClosePositionsAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) ClosePositionsAtEndOfDayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 41, 40)
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
func (m TradeAccountDataUpdatePointer) FlattenPositionsWhenUnderMarginIntradayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 43, 42)
}

// FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointerBuilder) FlattenPositionsWhenUnderMarginIntradayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 43, 42)
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
func (m TradeAccountDataUpdatePointer) FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 45, 44)
}

// FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 45, 44)
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
func (m TradeAccountDataUpdatePointer) SenderSubIDIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 47, 46)
}

// SenderSubIDIsSet
func (m TradeAccountDataUpdatePointerBuilder) SenderSubIDIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 47, 46)
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
func (m TradeAccountDataUpdatePointer) SenderLocationIdIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 52, 51)
}

// SenderLocationIdIsSet
func (m TradeAccountDataUpdatePointerBuilder) SenderLocationIdIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 52, 51)
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
func (m TradeAccountDataUpdatePointer) SelfMatchPreventionIDIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 57, 56)
}

// SelfMatchPreventionIDIsSet
func (m TradeAccountDataUpdatePointerBuilder) SelfMatchPreventionIDIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 57, 56)
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
func (m TradeAccountDataUpdatePointer) CTICodeIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 62, 61)
}

// CTICodeIsSet
func (m TradeAccountDataUpdatePointerBuilder) CTICodeIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 62, 61)
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
func (m TradeAccountDataUpdatePointer) TradeAccountIsReadOnlyIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 67, 66)
}

// TradeAccountIsReadOnlyIsSet
func (m TradeAccountDataUpdatePointerBuilder) TradeAccountIsReadOnlyIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 67, 66)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataUpdatePointer) TradeAccountIsReadOnly() uint8 {
	return message.Uint8VLS(m.Ptr, 68, 67)
}

// TradeAccountIsReadOnly
func (m TradeAccountDataUpdatePointerBuilder) TradeAccountIsReadOnly() uint8 {
	return message.Uint8VLS(m.Ptr, 68, 67)
}

// MaximumGlobalPositionQuantityIsSet
func (m TradeAccountDataUpdatePointer) MaximumGlobalPositionQuantityIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 69, 68)
}

// MaximumGlobalPositionQuantityIsSet
func (m TradeAccountDataUpdatePointerBuilder) MaximumGlobalPositionQuantityIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 69, 68)
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
func (m TradeAccountDataUpdatePointer) TradeFeePerContractIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 74, 73)
}

// TradeFeePerContractIsSet
func (m TradeAccountDataUpdatePointerBuilder) TradeFeePerContractIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 74, 73)
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
func (m TradeAccountDataUpdatePointer) TradeFeePerShareIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 83, 82)
}

// TradeFeePerShareIsSet
func (m TradeAccountDataUpdatePointerBuilder) TradeFeePerShareIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 83, 82)
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
func (m TradeAccountDataUpdatePointer) RequireSufficientMarginForNewPositionsIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 92, 91)
}

// RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdatePointerBuilder) RequireSufficientMarginForNewPositionsIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 92, 91)
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
func (m TradeAccountDataUpdatePointer) UsePercentOfMarginIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 94, 93)
}

// UsePercentOfMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) UsePercentOfMarginIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 94, 93)
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
func (m TradeAccountDataUpdatePointer) UsePercentOfMarginForDayTradingIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 99, 98)
}

// UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdatePointerBuilder) UsePercentOfMarginForDayTradingIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 99, 98)
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
func (m TradeAccountDataUpdatePointer) MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 104, 103)
}

// MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdatePointerBuilder) MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 104, 103)
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
func (m TradeAccountDataUpdatePointer) FirmIDIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 109, 108)
}

// FirmIDIsSet
func (m TradeAccountDataUpdatePointerBuilder) FirmIDIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 109, 108)
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
func (m TradeAccountDataUpdatePointer) TradingIsDisabledIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 114, 113)
}

// TradingIsDisabledIsSet
func (m TradeAccountDataUpdatePointerBuilder) TradingIsDisabledIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 114, 113)
}

// TradingIsDisabled
func (m TradeAccountDataUpdatePointer) TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Ptr, 115, 114)
}

// TradingIsDisabled
func (m TradeAccountDataUpdatePointerBuilder) TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Ptr, 115, 114)
}

// DescriptiveNameIsSet
func (m TradeAccountDataUpdatePointer) DescriptiveNameIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 116, 115)
}

// DescriptiveNameIsSet
func (m TradeAccountDataUpdatePointerBuilder) DescriptiveNameIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 116, 115)
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
func (m TradeAccountDataUpdatePointer) IsMasterFirmControlAccountIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 121, 120)
}

// IsMasterFirmControlAccountIsSet
func (m TradeAccountDataUpdatePointerBuilder) IsMasterFirmControlAccountIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 121, 120)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataUpdatePointer) IsMasterFirmControlAccount() uint8 {
	return message.Uint8VLS(m.Ptr, 122, 121)
}

// IsMasterFirmControlAccount
func (m TradeAccountDataUpdatePointerBuilder) IsMasterFirmControlAccount() uint8 {
	return message.Uint8VLS(m.Ptr, 122, 121)
}

// MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdatePointer) MinimumRequiredAccountValueIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 123, 122)
}

// MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdatePointerBuilder) MinimumRequiredAccountValueIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 123, 122)
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
func (m TradeAccountDataUpdatePointer) BeginTimeForDayMarginIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 132, 131)
}

// BeginTimeForDayMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) BeginTimeForDayMarginIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 132, 131)
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
func (m TradeAccountDataUpdatePointer) EndTimeForDayMarginIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 141, 140)
}

// EndTimeForDayMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) EndTimeForDayMarginIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 141, 140)
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
func (m TradeAccountDataUpdatePointer) DayMarginTimeZoneIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 150, 149)
}

// DayMarginTimeZoneIsSet
func (m TradeAccountDataUpdatePointerBuilder) DayMarginTimeZoneIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 150, 149)
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
func (m TradeAccountDataUpdatePointer) UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 155, 154)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 155, 154)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointer) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 156, 155)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() uint8 {
	return message.Uint8VLS(m.Ptr, 156, 155)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 157, 156)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 157, 156)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdatePointer) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 158, 157)
}

// UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 158, 157)
}

// UseMasterFirm_SymbolLimitsArrayIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_SymbolLimitsArrayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 159, 158)
}

// UseMasterFirm_SymbolLimitsArrayIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_SymbolLimitsArrayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 159, 158)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataUpdatePointer) UseMasterFirm_SymbolLimitsArray() uint8 {
	return message.Uint8VLS(m.Ptr, 160, 159)
}

// UseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_SymbolLimitsArray() uint8 {
	return message.Uint8VLS(m.Ptr, 160, 159)
}

// UseMasterFirm_TradeFeesIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradeFeesIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 161, 160)
}

// UseMasterFirm_TradeFeesIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradeFeesIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 161, 160)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradeFees() uint8 {
	return message.Uint8VLS(m.Ptr, 162, 161)
}

// UseMasterFirm_TradeFees
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradeFees() uint8 {
	return message.Uint8VLS(m.Ptr, 162, 161)
}

// UseMasterFirm_TradeFeePerShareIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradeFeePerShareIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 163, 162)
}

// UseMasterFirm_TradeFeePerShareIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradeFeePerShareIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 163, 162)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradeFeePerShare() uint8 {
	return message.Uint8VLS(m.Ptr, 164, 163)
}

// UseMasterFirm_TradeFeePerShare
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradeFeePerShare() uint8 {
	return message.Uint8VLS(m.Ptr, 164, 163)
}

// UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 165, 164)
}

// UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 165, 164)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdatePointer) UseMasterFirm_RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 166, 165)
}

// UseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_RequireSufficientMarginForNewPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 166, 165)
}

// UseMasterFirm_UsePercentOfMarginIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 167, 166)
}

// UseMasterFirm_UsePercentOfMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 167, 166)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMargin() uint8 {
	return message.Uint8VLS(m.Ptr, 168, 167)
}

// UseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMargin() uint8 {
	return message.Uint8VLS(m.Ptr, 168, 167)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 169, 168)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 169, 168)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdatePointer) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() uint8 {
	return message.Uint8VLS(m.Ptr, 170, 169)
}

// UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() uint8 {
	return message.Uint8VLS(m.Ptr, 170, 169)
}

// UseMasterFirm_MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_MinimumRequiredAccountValueIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 171, 170)
}

// UseMasterFirm_MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_MinimumRequiredAccountValueIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 171, 170)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataUpdatePointer) UseMasterFirm_MinimumRequiredAccountValue() uint8 {
	return message.Uint8VLS(m.Ptr, 172, 171)
}

// UseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_MinimumRequiredAccountValue() uint8 {
	return message.Uint8VLS(m.Ptr, 172, 171)
}

// UseMasterFirm_MarginTimeSettingsIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_MarginTimeSettingsIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 173, 172)
}

// UseMasterFirm_MarginTimeSettingsIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_MarginTimeSettingsIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 173, 172)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataUpdatePointer) UseMasterFirm_MarginTimeSettings() uint8 {
	return message.Uint8VLS(m.Ptr, 174, 173)
}

// UseMasterFirm_MarginTimeSettings
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_MarginTimeSettings() uint8 {
	return message.Uint8VLS(m.Ptr, 174, 173)
}

// UseMasterFirm_TradingIsDisabledIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradingIsDisabledIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 175, 174)
}

// UseMasterFirm_TradingIsDisabledIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradingIsDisabledIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 175, 174)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Ptr, 176, 175)
}

// UseMasterFirm_TradingIsDisabled
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Ptr, 176, 175)
}

// IsTradeStatisticsPublicallySharedIsSet
func (m TradeAccountDataUpdatePointer) IsTradeStatisticsPublicallySharedIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 177, 176)
}

// IsTradeStatisticsPublicallySharedIsSet
func (m TradeAccountDataUpdatePointerBuilder) IsTradeStatisticsPublicallySharedIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 177, 176)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataUpdatePointer) IsTradeStatisticsPublicallyShared() uint8 {
	return message.Uint8VLS(m.Ptr, 178, 177)
}

// IsTradeStatisticsPublicallyShared
func (m TradeAccountDataUpdatePointerBuilder) IsTradeStatisticsPublicallyShared() uint8 {
	return message.Uint8VLS(m.Ptr, 178, 177)
}

// IsReadOnlyFollowingRequestsAllowedIsSet
func (m TradeAccountDataUpdatePointer) IsReadOnlyFollowingRequestsAllowedIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 179, 178)
}

// IsReadOnlyFollowingRequestsAllowedIsSet
func (m TradeAccountDataUpdatePointerBuilder) IsReadOnlyFollowingRequestsAllowedIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 179, 178)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataUpdatePointer) IsReadOnlyFollowingRequestsAllowed() uint8 {
	return message.Uint8VLS(m.Ptr, 180, 179)
}

// IsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataUpdatePointerBuilder) IsReadOnlyFollowingRequestsAllowed() uint8 {
	return message.Uint8VLS(m.Ptr, 180, 179)
}

// IsTradeAccountSharingAllowedIsSet
func (m TradeAccountDataUpdatePointer) IsTradeAccountSharingAllowedIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 181, 180)
}

// IsTradeAccountSharingAllowedIsSet
func (m TradeAccountDataUpdatePointerBuilder) IsTradeAccountSharingAllowedIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 181, 180)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataUpdatePointer) IsTradeAccountSharingAllowed() uint8 {
	return message.Uint8VLS(m.Ptr, 182, 181)
}

// IsTradeAccountSharingAllowed
func (m TradeAccountDataUpdatePointerBuilder) IsTradeAccountSharingAllowed() uint8 {
	return message.Uint8VLS(m.Ptr, 182, 181)
}

// ReadOnlyShareToAllSCUsernamesIsSet
func (m TradeAccountDataUpdatePointer) ReadOnlyShareToAllSCUsernamesIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 183, 182)
}

// ReadOnlyShareToAllSCUsernamesIsSet
func (m TradeAccountDataUpdatePointerBuilder) ReadOnlyShareToAllSCUsernamesIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 183, 182)
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
func (m TradeAccountDataUpdatePointer) UseMasterFirm_SymbolCommissionsArrayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 185, 184)
}

// UseMasterFirm_SymbolCommissionsArrayIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_SymbolCommissionsArrayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 185, 184)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataUpdatePointer) UseMasterFirm_SymbolCommissionsArray() uint8 {
	return message.Uint8VLS(m.Ptr, 186, 185)
}

// UseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_SymbolCommissionsArray() uint8 {
	return message.Uint8VLS(m.Ptr, 186, 185)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 187, 186)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 187, 186)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointer) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 188, 187)
}

// UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() uint8 {
	return message.Uint8VLS(m.Ptr, 188, 187)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginForDayTradingIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 189, 188)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 189, 188)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginForDayTrading() uint8 {
	return message.Uint8VLS(m.Ptr, 190, 189)
}

// UseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginForDayTrading() uint8 {
	return message.Uint8VLS(m.Ptr, 190, 189)
}

// UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 191, 190)
}

// UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 191, 190)
}

// UseMasterFirm_SymbolCommissionsArrayFullOverride
func (m TradeAccountDataUpdatePointer) UseMasterFirm_SymbolCommissionsArrayFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 192, 191)
}

// UseMasterFirm_SymbolCommissionsArrayFullOverride
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_SymbolCommissionsArrayFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 192, 191)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 193, 192)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 193, 192)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataUpdatePointer) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() uint8 {
	return message.Uint8VLS(m.Ptr, 194, 193)
}

// UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() uint8 {
	return message.Uint8VLS(m.Ptr, 194, 193)
}

// UseMasterFirm_UsePercentOfMarginFullOverrideIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginFullOverrideIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 195, 194)
}

// UseMasterFirm_UsePercentOfMarginFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginFullOverrideIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 195, 194)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 196, 195)
}

// UseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 196, 195)
}

// UseMasterFirm_TradeFeesFullOverrideIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradeFeesFullOverrideIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 197, 196)
}

// UseMasterFirm_TradeFeesFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradeFeesFullOverrideIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 197, 196)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataUpdatePointer) UseMasterFirm_TradeFeesFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 198, 197)
}

// UseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_TradeFeesFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 198, 197)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 199, 198)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 199, 198)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataUpdatePointer) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 200, 199)
}

// UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataUpdatePointerBuilder) UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() uint8 {
	return message.Uint8VLS(m.Ptr, 200, 199)
}

// LiquidationOnlyModeIsSet
func (m TradeAccountDataUpdatePointer) LiquidationOnlyModeIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 201, 200)
}

// LiquidationOnlyModeIsSet
func (m TradeAccountDataUpdatePointerBuilder) LiquidationOnlyModeIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 201, 200)
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
func (m TradeAccountDataUpdatePointer) CustomerOrFirmIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 203, 202)
}

// CustomerOrFirmIsSet
func (m TradeAccountDataUpdatePointerBuilder) CustomerOrFirmIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 203, 202)
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
func (m TradeAccountDataUpdatePointer) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 205, 204)
}

// MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 205, 204)
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
func (m TradeAccountDataUpdatePointer) MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 207, 206)
}

// MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 207, 206)
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
func (m TradeAccountDataUpdatePointer) MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 209, 208)
}

// MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 209, 208)
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
func (m TradeAccountDataUpdatePointer) MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 211, 210)
}

// MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 211, 210)
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
func (m TradeAccountDataUpdatePointer) MasterFirm_MaximumOrderQuantityIsSet() uint32 {
	return message.Uint32VLS(m.Ptr, 216, 212)
}

// MasterFirm_MaximumOrderQuantityIsSet
func (m TradeAccountDataUpdatePointerBuilder) MasterFirm_MaximumOrderQuantityIsSet() uint32 {
	return message.Uint32VLS(m.Ptr, 216, 212)
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
func (m TradeAccountDataUpdatePointer) ExchangeTraderIdIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 221, 220)
}

// ExchangeTraderIdIsSet
func (m TradeAccountDataUpdatePointerBuilder) ExchangeTraderIdIsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 221, 220)
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
func (m TradeAccountDataUpdatePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetIsNewAccount
func (m TradeAccountDataUpdatePointerBuilder) SetIsNewAccount(value uint8) {
	message.SetUint8VLS(m.Ptr, 11, 10, value)
}

// SetNewAccountAuthorizedUsername
func (m *TradeAccountDataUpdatePointerBuilder) SetNewAccountAuthorizedUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetTradeAccount
func (m *TradeAccountDataUpdatePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 19, 15, value)
}

// SetCurrencyCodeIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetCurrencyCodeIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 20, 19, value)
}

// SetCurrencyCode
func (m *TradeAccountDataUpdatePointerBuilder) SetCurrencyCode(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetDailyNetLossLimitInAccountCurrencyIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetDailyNetLossLimitInAccountCurrencyIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 25, 24, value)
}

// SetDailyNetLossLimitInAccountCurrency
func (m TradeAccountDataUpdatePointerBuilder) SetDailyNetLossLimitInAccountCurrency(value float32) {
	message.SetFloat32VLS(m.Ptr, 29, 25, value)
}

// SetPercentOfCashBalanceForDailyNetLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetPercentOfCashBalanceForDailyNetLossLimitIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 30, 29, value)
}

// SetPercentOfCashBalanceForDailyNetLossLimit
func (m TradeAccountDataUpdatePointerBuilder) SetPercentOfCashBalanceForDailyNetLossLimit(value int32) {
	message.SetInt32VLS(m.Ptr, 34, 30, value)
}

// SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 35, 34, value)
}

// SetUseTrailingAccountValueToNotAllowIncreaseInPositions
func (m TradeAccountDataUpdatePointerBuilder) SetUseTrailingAccountValueToNotAllowIncreaseInPositions(value uint8) {
	message.SetUint8VLS(m.Ptr, 36, 35, value)
}

// SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 37, 36, value)
}

// SetDoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointerBuilder) SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Ptr, 38, 37, value)
}

// SetFlattenPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetFlattenPositionsAtDailyLossLimitIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 39, 38, value)
}

// SetFlattenPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointerBuilder) SetFlattenPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Ptr, 40, 39, value)
}

// SetClosePositionsAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetClosePositionsAtEndOfDayIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 41, 40, value)
}

// SetClosePositionsAtEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 42, 41, value)
}

// SetFlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetFlattenPositionsWhenUnderMarginIntradayIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 43, 42, value)
}

// SetFlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointerBuilder) SetFlattenPositionsWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Ptr, 44, 43, value)
}

// SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 45, 44, value)
}

// SetFlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) SetFlattenPositionsWhenUnderMarginAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 46, 45, value)
}

// SetSenderSubIDIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetSenderSubIDIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 47, 46, value)
}

// SetSenderSubID
func (m *TradeAccountDataUpdatePointerBuilder) SetSenderSubID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 51, 47, value)
}

// SetSenderLocationIdIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetSenderLocationIdIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 52, 51, value)
}

// SetSenderLocationId
func (m *TradeAccountDataUpdatePointerBuilder) SetSenderLocationId(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 56, 52, value)
}

// SetSelfMatchPreventionIDIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetSelfMatchPreventionIDIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 57, 56, value)
}

// SetSelfMatchPreventionID
func (m *TradeAccountDataUpdatePointerBuilder) SetSelfMatchPreventionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 61, 57, value)
}

// SetCTICodeIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetCTICodeIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 62, 61, value)
}

// SetCTICode
func (m TradeAccountDataUpdatePointerBuilder) SetCTICode(value int32) {
	message.SetInt32VLS(m.Ptr, 66, 62, value)
}

// SetTradeAccountIsReadOnlyIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetTradeAccountIsReadOnlyIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 67, 66, value)
}

// SetTradeAccountIsReadOnly
func (m TradeAccountDataUpdatePointerBuilder) SetTradeAccountIsReadOnly(value uint8) {
	message.SetUint8VLS(m.Ptr, 68, 67, value)
}

// SetMaximumGlobalPositionQuantityIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMaximumGlobalPositionQuantityIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 69, 68, value)
}

// SetMaximumGlobalPositionQuantity
func (m TradeAccountDataUpdatePointerBuilder) SetMaximumGlobalPositionQuantity(value int32) {
	message.SetInt32VLS(m.Ptr, 73, 69, value)
}

// SetTradeFeePerContractIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetTradeFeePerContractIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 74, 73, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataUpdatePointerBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Ptr, 82, 74, value)
}

// SetTradeFeePerShareIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetTradeFeePerShareIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 83, 82, value)
}

// SetTradeFeePerShare
func (m TradeAccountDataUpdatePointerBuilder) SetTradeFeePerShare(value float64) {
	message.SetFloat64VLS(m.Ptr, 91, 83, value)
}

// SetRequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetRequireSufficientMarginForNewPositionsIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 92, 91, value)
}

// SetRequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdatePointerBuilder) SetRequireSufficientMarginForNewPositions(value uint8) {
	message.SetUint8VLS(m.Ptr, 93, 92, value)
}

// SetUsePercentOfMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUsePercentOfMarginIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 94, 93, value)
}

// SetUsePercentOfMargin
func (m TradeAccountDataUpdatePointerBuilder) SetUsePercentOfMargin(value int32) {
	message.SetInt32VLS(m.Ptr, 98, 94, value)
}

// SetUsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUsePercentOfMarginForDayTradingIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 99, 98, value)
}

// SetUsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdatePointerBuilder) SetUsePercentOfMarginForDayTrading(value int32) {
	message.SetInt32VLS(m.Ptr, 103, 99, value)
}

// SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 104, 103, value)
}

// SetMaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdatePointerBuilder) SetMaximumAllowedAccountBalanceForPositionsAsPercentage(value int32) {
	message.SetInt32VLS(m.Ptr, 108, 104, value)
}

// SetFirmIDIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetFirmIDIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 109, 108, value)
}

// SetFirmID
func (m *TradeAccountDataUpdatePointerBuilder) SetFirmID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 113, 109, value)
}

// SetTradingIsDisabledIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetTradingIsDisabledIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 114, 113, value)
}

// SetTradingIsDisabled
func (m TradeAccountDataUpdatePointerBuilder) SetTradingIsDisabled(value uint8) {
	message.SetUint8VLS(m.Ptr, 115, 114, value)
}

// SetDescriptiveNameIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetDescriptiveNameIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 116, 115, value)
}

// SetDescriptiveName
func (m *TradeAccountDataUpdatePointerBuilder) SetDescriptiveName(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 120, 116, value)
}

// SetIsMasterFirmControlAccountIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetIsMasterFirmControlAccountIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 121, 120, value)
}

// SetIsMasterFirmControlAccount
func (m TradeAccountDataUpdatePointerBuilder) SetIsMasterFirmControlAccount(value uint8) {
	message.SetUint8VLS(m.Ptr, 122, 121, value)
}

// SetMinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMinimumRequiredAccountValueIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 123, 122, value)
}

// SetMinimumRequiredAccountValue
func (m TradeAccountDataUpdatePointerBuilder) SetMinimumRequiredAccountValue(value float64) {
	message.SetFloat64VLS(m.Ptr, 131, 123, value)
}

// SetBeginTimeForDayMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetBeginTimeForDayMarginIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 132, 131, value)
}

// SetBeginTimeForDayMargin
func (m TradeAccountDataUpdatePointerBuilder) SetBeginTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Ptr, 140, 132, value)
}

// SetEndTimeForDayMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetEndTimeForDayMarginIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 141, 140, value)
}

// SetEndTimeForDayMargin
func (m TradeAccountDataUpdatePointerBuilder) SetEndTimeForDayMargin(value int64) {
	message.SetInt64VLS(m.Ptr, 149, 141, value)
}

// SetDayMarginTimeZoneIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetDayMarginTimeZoneIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 150, 149, value)
}

// SetDayMarginTimeZone
func (m *TradeAccountDataUpdatePointerBuilder) SetDayMarginTimeZone(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 154, 150, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 155, 154, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Ptr, 156, 155, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 157, 156, value)
}

// SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 158, 157, value)
}

// SetUseMasterFirm_SymbolLimitsArrayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_SymbolLimitsArrayIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 159, 158, value)
}

// SetUseMasterFirm_SymbolLimitsArray
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_SymbolLimitsArray(value uint8) {
	message.SetUint8VLS(m.Ptr, 160, 159, value)
}

// SetUseMasterFirm_TradeFeesIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradeFeesIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 161, 160, value)
}

// SetUseMasterFirm_TradeFees
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradeFees(value uint8) {
	message.SetUint8VLS(m.Ptr, 162, 161, value)
}

// SetUseMasterFirm_TradeFeePerShareIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradeFeePerShareIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 163, 162, value)
}

// SetUseMasterFirm_TradeFeePerShare
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradeFeePerShare(value uint8) {
	message.SetUint8VLS(m.Ptr, 164, 163, value)
}

// SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 165, 164, value)
}

// SetUseMasterFirm_RequireSufficientMarginForNewPositions
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_RequireSufficientMarginForNewPositions(value uint8) {
	message.SetUint8VLS(m.Ptr, 166, 165, value)
}

// SetUseMasterFirm_UsePercentOfMarginIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 167, 166, value)
}

// SetUseMasterFirm_UsePercentOfMargin
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMargin(value uint8) {
	message.SetUint8VLS(m.Ptr, 168, 167, value)
}

// SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 169, 168, value)
}

// SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(value uint8) {
	message.SetUint8VLS(m.Ptr, 170, 169, value)
}

// SetUseMasterFirm_MinimumRequiredAccountValueIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_MinimumRequiredAccountValueIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 171, 170, value)
}

// SetUseMasterFirm_MinimumRequiredAccountValue
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_MinimumRequiredAccountValue(value uint8) {
	message.SetUint8VLS(m.Ptr, 172, 171, value)
}

// SetUseMasterFirm_MarginTimeSettingsIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_MarginTimeSettingsIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 173, 172, value)
}

// SetUseMasterFirm_MarginTimeSettings
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_MarginTimeSettings(value uint8) {
	message.SetUint8VLS(m.Ptr, 174, 173, value)
}

// SetUseMasterFirm_TradingIsDisabledIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradingIsDisabledIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 175, 174, value)
}

// SetUseMasterFirm_TradingIsDisabled
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradingIsDisabled(value uint8) {
	message.SetUint8VLS(m.Ptr, 176, 175, value)
}

// SetIsTradeStatisticsPublicallySharedIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetIsTradeStatisticsPublicallySharedIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 177, 176, value)
}

// SetIsTradeStatisticsPublicallyShared
func (m TradeAccountDataUpdatePointerBuilder) SetIsTradeStatisticsPublicallyShared(value uint8) {
	message.SetUint8VLS(m.Ptr, 178, 177, value)
}

// SetIsReadOnlyFollowingRequestsAllowedIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetIsReadOnlyFollowingRequestsAllowedIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 179, 178, value)
}

// SetIsReadOnlyFollowingRequestsAllowed
func (m TradeAccountDataUpdatePointerBuilder) SetIsReadOnlyFollowingRequestsAllowed(value uint8) {
	message.SetUint8VLS(m.Ptr, 180, 179, value)
}

// SetIsTradeAccountSharingAllowedIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetIsTradeAccountSharingAllowedIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 181, 180, value)
}

// SetIsTradeAccountSharingAllowed
func (m TradeAccountDataUpdatePointerBuilder) SetIsTradeAccountSharingAllowed(value uint8) {
	message.SetUint8VLS(m.Ptr, 182, 181, value)
}

// SetReadOnlyShareToAllSCUsernamesIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetReadOnlyShareToAllSCUsernamesIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 183, 182, value)
}

// SetReadOnlyShareToAllSCUsernames
func (m TradeAccountDataUpdatePointerBuilder) SetReadOnlyShareToAllSCUsernames(value uint8) {
	message.SetUint8VLS(m.Ptr, 184, 183, value)
}

// SetUseMasterFirm_SymbolCommissionsArrayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_SymbolCommissionsArrayIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 185, 184, value)
}

// SetUseMasterFirm_SymbolCommissionsArray
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_SymbolCommissionsArray(value uint8) {
	message.SetUint8VLS(m.Ptr, 186, 185, value)
}

// SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 187, 186, value)
}

// SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(value uint8) {
	message.SetUint8VLS(m.Ptr, 188, 187, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 189, 188, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTrading
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTrading(value uint8) {
	message.SetUint8VLS(m.Ptr, 190, 189, value)
}

// SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 191, 190, value)
}

// SetUseMasterFirm_SymbolCommissionsArrayFullOverride
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_SymbolCommissionsArrayFullOverride(value uint8) {
	message.SetUint8VLS(m.Ptr, 192, 191, value)
}

// SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 193, 192, value)
}

// SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(value uint8) {
	message.SetUint8VLS(m.Ptr, 194, 193, value)
}

// SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 195, 194, value)
}

// SetUseMasterFirm_UsePercentOfMarginFullOverride
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginFullOverride(value uint8) {
	message.SetUint8VLS(m.Ptr, 196, 195, value)
}

// SetUseMasterFirm_TradeFeesFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradeFeesFullOverrideIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 197, 196, value)
}

// SetUseMasterFirm_TradeFeesFullOverride
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_TradeFeesFullOverride(value uint8) {
	message.SetUint8VLS(m.Ptr, 198, 197, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 199, 198, value)
}

// SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride
func (m TradeAccountDataUpdatePointerBuilder) SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(value uint8) {
	message.SetUint8VLS(m.Ptr, 200, 199, value)
}

// SetLiquidationOnlyModeIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetLiquidationOnlyModeIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 201, 200, value)
}

// SetLiquidationOnlyMode
func (m TradeAccountDataUpdatePointerBuilder) SetLiquidationOnlyMode(value uint8) {
	message.SetUint8VLS(m.Ptr, 202, 201, value)
}

// SetCustomerOrFirmIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetCustomerOrFirmIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 203, 202, value)
}

// SetCustomerOrFirm
func (m TradeAccountDataUpdatePointerBuilder) SetCustomerOrFirm(value uint8) {
	message.SetUint8VLS(m.Ptr, 204, 203, value)
}

// SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 205, 204, value)
}

// SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(value uint8) {
	message.SetUint8VLS(m.Ptr, 206, 205, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 207, 206, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(value uint8) {
	message.SetUint8VLS(m.Ptr, 208, 207, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 209, 208, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginIntraday
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(value uint8) {
	message.SetUint8VLS(m.Ptr, 210, 209, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 211, 210, value)
}

// SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 212, 211, value)
}

// SetMasterFirm_MaximumOrderQuantityIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_MaximumOrderQuantityIsSet(value uint32) {
	message.SetUint32VLS(m.Ptr, 216, 212, value)
}

// SetMasterFirm_MaximumOrderQuantity
func (m TradeAccountDataUpdatePointerBuilder) SetMasterFirm_MaximumOrderQuantity(value uint32) {
	message.SetUint32VLS(m.Ptr, 220, 216, value)
}

// SetExchangeTraderIdIsSet
func (m TradeAccountDataUpdatePointerBuilder) SetExchangeTraderIdIsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 221, 220, value)
}

// SetExchangeTraderId
func (m *TradeAccountDataUpdatePointerBuilder) SetExchangeTraderId(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 225, 221, value)
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
