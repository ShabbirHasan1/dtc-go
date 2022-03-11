//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUpdatePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10117,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsNewAccount())
	w.Data = append(w.Data, ',')
	w.String(m.NewAccountAuthorizedUsername())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CurrencyCodeIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.CurrencyCode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DailyNetLossLimitInAccountCurrencyIsSet())
	w.Data = append(w.Data, ',')
	w.Float32(m.DailyNetLossLimitInAccountCurrency())
	w.Data = append(w.Data, ',')
	w.Uint8(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.PercentOfCashBalanceForDailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SenderSubIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SenderSubID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SenderLocationIdIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SenderLocationId())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SelfMatchPreventionIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CTICodeIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.CTICode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeAccountIsReadOnlyIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeAccountIsReadOnly())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MaximumGlobalPositionQuantityIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumGlobalPositionQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeFeePerContractIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeFeePerShareIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequireSufficientMarginForNewPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsePercentOfMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsePercentOfMarginForDayTradingIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FirmIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.FirmID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradingIsDisabledIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DescriptiveNameIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.DescriptiveName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsMasterFirmControlAccountIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsMasterFirmControlAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MinimumRequiredAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BeginTimeForDayMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int64(m.BeginTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.EndTimeForDayMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int64(m.EndTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DayMarginTimeZoneIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.DayMarginTimeZone())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolLimitsArray())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeesIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFees())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeePerShareIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MarginTimeSettingsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MarginTimeSettings())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradingIsDisabledIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsTradeStatisticsPublicallySharedIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsTradeStatisticsPublicallyShared())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsReadOnlyFollowingRequestsAllowed())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsTradeAccountSharingAllowedIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsTradeAccountSharingAllowed())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReadOnlyShareToAllSCUsernamesIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReadOnlyShareToAllSCUsernames())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArray())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeesFullOverride())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.Data = append(w.Data, ',')
	w.Uint8(m.LiquidationOnlyModeIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.LiquidationOnlyMode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CustomerOrFirmIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CustomerOrFirm())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantityIsSet())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ExchangeTraderIdIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeTraderId())
	return w.Finish(), nil
}

func (m TradeAccountDataUpdatePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10117,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsNewAccount())
	w.Data = append(w.Data, ',')
	w.String(m.NewAccountAuthorizedUsername())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CurrencyCodeIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.CurrencyCode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DailyNetLossLimitInAccountCurrencyIsSet())
	w.Data = append(w.Data, ',')
	w.Float32(m.DailyNetLossLimitInAccountCurrency())
	w.Data = append(w.Data, ',')
	w.Uint8(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.PercentOfCashBalanceForDailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SenderSubIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SenderSubID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SenderLocationIdIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SenderLocationId())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SelfMatchPreventionIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CTICodeIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.CTICode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeAccountIsReadOnlyIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeAccountIsReadOnly())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MaximumGlobalPositionQuantityIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumGlobalPositionQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeFeePerContractIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeFeePerShareIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequireSufficientMarginForNewPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsePercentOfMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsePercentOfMarginForDayTradingIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FirmIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.FirmID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradingIsDisabledIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DescriptiveNameIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.DescriptiveName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsMasterFirmControlAccountIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsMasterFirmControlAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MinimumRequiredAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BeginTimeForDayMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int64(m.BeginTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.EndTimeForDayMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int64(m.EndTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DayMarginTimeZoneIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.DayMarginTimeZone())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolLimitsArray())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeesIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFees())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeePerShareIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MarginTimeSettingsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MarginTimeSettings())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradingIsDisabledIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsTradeStatisticsPublicallySharedIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsTradeStatisticsPublicallyShared())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsReadOnlyFollowingRequestsAllowed())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsTradeAccountSharingAllowedIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsTradeAccountSharingAllowed())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReadOnlyShareToAllSCUsernamesIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReadOnlyShareToAllSCUsernames())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArray())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeesFullOverride())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.Data = append(w.Data, ',')
	w.Uint8(m.LiquidationOnlyModeIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.LiquidationOnlyMode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CustomerOrFirmIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CustomerOrFirm())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantityIsSet())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ExchangeTraderIdIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeTraderId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUpdatePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10117)
	w.Uint32Field("RequestID", m.RequestID())
	w.Uint8Field("IsNewAccount", m.IsNewAccount())
	w.StringField("NewAccountAuthorizedUsername", m.NewAccountAuthorizedUsername())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("CurrencyCodeIsSet", m.CurrencyCodeIsSet())
	w.StringField("CurrencyCode", m.CurrencyCode())
	w.Uint8Field("DailyNetLossLimitInAccountCurrencyIsSet", m.DailyNetLossLimitInAccountCurrencyIsSet())
	w.Float32Field("DailyNetLossLimitInAccountCurrency", m.DailyNetLossLimitInAccountCurrency())
	w.Uint8Field("PercentOfCashBalanceForDailyNetLossLimitIsSet", m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	w.Int32Field("PercentOfCashBalanceForDailyNetLossLimit", m.PercentOfCashBalanceForDailyNetLossLimit())
	w.Uint8Field("UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet", m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	w.Uint8Field("UseTrailingAccountValueToNotAllowIncreaseInPositions", m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.Uint8Field("DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet", m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Uint8Field("FlattenPositionsAtDailyLossLimitIsSet", m.FlattenPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("FlattenPositionsAtDailyLossLimit", m.FlattenPositionsAtDailyLossLimit())
	w.Uint8Field("ClosePositionsAtEndOfDayIsSet", m.ClosePositionsAtEndOfDayIsSet())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.Uint8Field("FlattenPositionsWhenUnderMarginIntradayIsSet", m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Uint8Field("FlattenPositionsWhenUnderMarginIntraday", m.FlattenPositionsWhenUnderMarginIntraday())
	w.Uint8Field("FlattenPositionsWhenUnderMarginAtEndOfDayIsSet", m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Uint8Field("FlattenPositionsWhenUnderMarginAtEndOfDay", m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Uint8Field("SenderSubIDIsSet", m.SenderSubIDIsSet())
	w.StringField("SenderSubID", m.SenderSubID())
	w.Uint8Field("SenderLocationIdIsSet", m.SenderLocationIdIsSet())
	w.StringField("SenderLocationId", m.SenderLocationId())
	w.Uint8Field("SelfMatchPreventionIDIsSet", m.SelfMatchPreventionIDIsSet())
	w.StringField("SelfMatchPreventionID", m.SelfMatchPreventionID())
	w.Uint8Field("CTICodeIsSet", m.CTICodeIsSet())
	w.Int32Field("CTICode", m.CTICode())
	w.Uint8Field("TradeAccountIsReadOnlyIsSet", m.TradeAccountIsReadOnlyIsSet())
	w.Uint8Field("TradeAccountIsReadOnly", m.TradeAccountIsReadOnly())
	w.Uint8Field("MaximumGlobalPositionQuantityIsSet", m.MaximumGlobalPositionQuantityIsSet())
	w.Int32Field("MaximumGlobalPositionQuantity", m.MaximumGlobalPositionQuantity())
	w.Uint8Field("TradeFeePerContractIsSet", m.TradeFeePerContractIsSet())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	w.Uint8Field("TradeFeePerShareIsSet", m.TradeFeePerShareIsSet())
	w.Float64Field("TradeFeePerShare", m.TradeFeePerShare())
	w.Uint8Field("RequireSufficientMarginForNewPositionsIsSet", m.RequireSufficientMarginForNewPositionsIsSet())
	w.Uint8Field("RequireSufficientMarginForNewPositions", m.RequireSufficientMarginForNewPositions())
	w.Uint8Field("UsePercentOfMarginIsSet", m.UsePercentOfMarginIsSet())
	w.Int32Field("UsePercentOfMargin", m.UsePercentOfMargin())
	w.Uint8Field("UsePercentOfMarginForDayTradingIsSet", m.UsePercentOfMarginForDayTradingIsSet())
	w.Int32Field("UsePercentOfMarginForDayTrading", m.UsePercentOfMarginForDayTrading())
	w.Uint8Field("MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet", m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Int32Field("MaximumAllowedAccountBalanceForPositionsAsPercentage", m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Uint8Field("FirmIDIsSet", m.FirmIDIsSet())
	w.StringField("FirmID", m.FirmID())
	w.Uint8Field("TradingIsDisabledIsSet", m.TradingIsDisabledIsSet())
	w.Uint8Field("TradingIsDisabled", m.TradingIsDisabled())
	w.Uint8Field("DescriptiveNameIsSet", m.DescriptiveNameIsSet())
	w.StringField("DescriptiveName", m.DescriptiveName())
	w.Uint8Field("IsMasterFirmControlAccountIsSet", m.IsMasterFirmControlAccountIsSet())
	w.Uint8Field("IsMasterFirmControlAccount", m.IsMasterFirmControlAccount())
	w.Uint8Field("MinimumRequiredAccountValueIsSet", m.MinimumRequiredAccountValueIsSet())
	w.Float64Field("MinimumRequiredAccountValue", m.MinimumRequiredAccountValue())
	w.Uint8Field("BeginTimeForDayMarginIsSet", m.BeginTimeForDayMarginIsSet())
	w.Int64Field("BeginTimeForDayMargin", m.BeginTimeForDayMargin())
	w.Uint8Field("EndTimeForDayMarginIsSet", m.EndTimeForDayMarginIsSet())
	w.Int64Field("EndTimeForDayMargin", m.EndTimeForDayMargin())
	w.Uint8Field("DayMarginTimeZoneIsSet", m.DayMarginTimeZoneIsSet())
	w.StringField("DayMarginTimeZone", m.DayMarginTimeZone())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Uint8Field("m_UseMasterFirm_SymbolLimitsArrayIsSet", m.UseMasterFirm_SymbolLimitsArrayIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolLimitsArray", m.UseMasterFirm_SymbolLimitsArray())
	w.Uint8Field("m_UseMasterFirm_TradeFeesIsSet", m.UseMasterFirm_TradeFeesIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFees", m.UseMasterFirm_TradeFees())
	w.Uint8Field("m_UseMasterFirm_TradeFeePerShareIsSet", m.UseMasterFirm_TradeFeePerShareIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFeePerShare", m.UseMasterFirm_TradeFeePerShare())
	w.Uint8Field("m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet", m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	w.Uint8Field("m_UseMasterFirm_RequireSufficientMarginForNewPositions", m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginIsSet", m.UseMasterFirm_UsePercentOfMarginIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMargin", m.UseMasterFirm_UsePercentOfMargin())
	w.Uint8Field("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Uint8Field("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Uint8Field("m_UseMasterFirm_MinimumRequiredAccountValueIsSet", m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	w.Uint8Field("m_UseMasterFirm_MinimumRequiredAccountValue", m.UseMasterFirm_MinimumRequiredAccountValue())
	w.Uint8Field("m_UseMasterFirm_MarginTimeSettingsIsSet", m.UseMasterFirm_MarginTimeSettingsIsSet())
	w.Uint8Field("m_UseMasterFirm_MarginTimeSettings", m.UseMasterFirm_MarginTimeSettings())
	w.Uint8Field("m_UseMasterFirm_TradingIsDisabledIsSet", m.UseMasterFirm_TradingIsDisabledIsSet())
	w.Uint8Field("m_UseMasterFirm_TradingIsDisabled", m.UseMasterFirm_TradingIsDisabled())
	w.Uint8Field("IsTradeStatisticsPublicallySharedIsSet", m.IsTradeStatisticsPublicallySharedIsSet())
	w.Uint8Field("IsTradeStatisticsPublicallyShared", m.IsTradeStatisticsPublicallyShared())
	w.Uint8Field("IsReadOnlyFollowingRequestsAllowedIsSet", m.IsReadOnlyFollowingRequestsAllowedIsSet())
	w.Uint8Field("IsReadOnlyFollowingRequestsAllowed", m.IsReadOnlyFollowingRequestsAllowed())
	w.Uint8Field("IsTradeAccountSharingAllowedIsSet", m.IsTradeAccountSharingAllowedIsSet())
	w.Uint8Field("IsTradeAccountSharingAllowed", m.IsTradeAccountSharingAllowed())
	w.Uint8Field("ReadOnlyShareToAllSCUsernamesIsSet", m.ReadOnlyShareToAllSCUsernamesIsSet())
	w.Uint8Field("ReadOnlyShareToAllSCUsernames", m.ReadOnlyShareToAllSCUsernames())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArrayIsSet", m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArray", m.UseMasterFirm_SymbolCommissionsArray())
	w.Uint8Field("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet", m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTrading", m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet", m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArrayFullOverride", m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	w.Uint8Field("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	w.Uint8Field("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet", m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginFullOverride", m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.Uint8Field("m_UseMasterFirm_TradeFeesFullOverrideIsSet", m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFeesFullOverride", m.UseMasterFirm_TradeFeesFullOverride())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.Uint8Field("m_LiquidationOnlyModeIsSet", m.LiquidationOnlyModeIsSet())
	w.Uint8Field("m_LiquidationOnlyMode", m.LiquidationOnlyMode())
	w.Uint8Field("m_CustomerOrFirmIsSet", m.CustomerOrFirmIsSet())
	w.Uint8Field("m_CustomerOrFirm", m.CustomerOrFirm())
	w.Uint8Field("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet", m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginIntraday", m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantityIsSet", m.MasterFirm_MaximumOrderQuantityIsSet())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantity", m.MasterFirm_MaximumOrderQuantity())
	w.Uint8Field("m_ExchangeTraderIdIsSet", m.ExchangeTraderIdIsSet())
	w.StringField("m_ExchangeTraderId", m.ExchangeTraderId())
	return w.Finish(), nil
}

func (m TradeAccountDataUpdatePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10117)
	w.Uint32Field("RequestID", m.RequestID())
	w.Uint8Field("IsNewAccount", m.IsNewAccount())
	w.StringField("NewAccountAuthorizedUsername", m.NewAccountAuthorizedUsername())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("CurrencyCodeIsSet", m.CurrencyCodeIsSet())
	w.StringField("CurrencyCode", m.CurrencyCode())
	w.Uint8Field("DailyNetLossLimitInAccountCurrencyIsSet", m.DailyNetLossLimitInAccountCurrencyIsSet())
	w.Float32Field("DailyNetLossLimitInAccountCurrency", m.DailyNetLossLimitInAccountCurrency())
	w.Uint8Field("PercentOfCashBalanceForDailyNetLossLimitIsSet", m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	w.Int32Field("PercentOfCashBalanceForDailyNetLossLimit", m.PercentOfCashBalanceForDailyNetLossLimit())
	w.Uint8Field("UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet", m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	w.Uint8Field("UseTrailingAccountValueToNotAllowIncreaseInPositions", m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.Uint8Field("DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet", m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Uint8Field("FlattenPositionsAtDailyLossLimitIsSet", m.FlattenPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("FlattenPositionsAtDailyLossLimit", m.FlattenPositionsAtDailyLossLimit())
	w.Uint8Field("ClosePositionsAtEndOfDayIsSet", m.ClosePositionsAtEndOfDayIsSet())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.Uint8Field("FlattenPositionsWhenUnderMarginIntradayIsSet", m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Uint8Field("FlattenPositionsWhenUnderMarginIntraday", m.FlattenPositionsWhenUnderMarginIntraday())
	w.Uint8Field("FlattenPositionsWhenUnderMarginAtEndOfDayIsSet", m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Uint8Field("FlattenPositionsWhenUnderMarginAtEndOfDay", m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Uint8Field("SenderSubIDIsSet", m.SenderSubIDIsSet())
	w.StringField("SenderSubID", m.SenderSubID())
	w.Uint8Field("SenderLocationIdIsSet", m.SenderLocationIdIsSet())
	w.StringField("SenderLocationId", m.SenderLocationId())
	w.Uint8Field("SelfMatchPreventionIDIsSet", m.SelfMatchPreventionIDIsSet())
	w.StringField("SelfMatchPreventionID", m.SelfMatchPreventionID())
	w.Uint8Field("CTICodeIsSet", m.CTICodeIsSet())
	w.Int32Field("CTICode", m.CTICode())
	w.Uint8Field("TradeAccountIsReadOnlyIsSet", m.TradeAccountIsReadOnlyIsSet())
	w.Uint8Field("TradeAccountIsReadOnly", m.TradeAccountIsReadOnly())
	w.Uint8Field("MaximumGlobalPositionQuantityIsSet", m.MaximumGlobalPositionQuantityIsSet())
	w.Int32Field("MaximumGlobalPositionQuantity", m.MaximumGlobalPositionQuantity())
	w.Uint8Field("TradeFeePerContractIsSet", m.TradeFeePerContractIsSet())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	w.Uint8Field("TradeFeePerShareIsSet", m.TradeFeePerShareIsSet())
	w.Float64Field("TradeFeePerShare", m.TradeFeePerShare())
	w.Uint8Field("RequireSufficientMarginForNewPositionsIsSet", m.RequireSufficientMarginForNewPositionsIsSet())
	w.Uint8Field("RequireSufficientMarginForNewPositions", m.RequireSufficientMarginForNewPositions())
	w.Uint8Field("UsePercentOfMarginIsSet", m.UsePercentOfMarginIsSet())
	w.Int32Field("UsePercentOfMargin", m.UsePercentOfMargin())
	w.Uint8Field("UsePercentOfMarginForDayTradingIsSet", m.UsePercentOfMarginForDayTradingIsSet())
	w.Int32Field("UsePercentOfMarginForDayTrading", m.UsePercentOfMarginForDayTrading())
	w.Uint8Field("MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet", m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Int32Field("MaximumAllowedAccountBalanceForPositionsAsPercentage", m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Uint8Field("FirmIDIsSet", m.FirmIDIsSet())
	w.StringField("FirmID", m.FirmID())
	w.Uint8Field("TradingIsDisabledIsSet", m.TradingIsDisabledIsSet())
	w.Uint8Field("TradingIsDisabled", m.TradingIsDisabled())
	w.Uint8Field("DescriptiveNameIsSet", m.DescriptiveNameIsSet())
	w.StringField("DescriptiveName", m.DescriptiveName())
	w.Uint8Field("IsMasterFirmControlAccountIsSet", m.IsMasterFirmControlAccountIsSet())
	w.Uint8Field("IsMasterFirmControlAccount", m.IsMasterFirmControlAccount())
	w.Uint8Field("MinimumRequiredAccountValueIsSet", m.MinimumRequiredAccountValueIsSet())
	w.Float64Field("MinimumRequiredAccountValue", m.MinimumRequiredAccountValue())
	w.Uint8Field("BeginTimeForDayMarginIsSet", m.BeginTimeForDayMarginIsSet())
	w.Int64Field("BeginTimeForDayMargin", m.BeginTimeForDayMargin())
	w.Uint8Field("EndTimeForDayMarginIsSet", m.EndTimeForDayMarginIsSet())
	w.Int64Field("EndTimeForDayMargin", m.EndTimeForDayMargin())
	w.Uint8Field("DayMarginTimeZoneIsSet", m.DayMarginTimeZoneIsSet())
	w.StringField("DayMarginTimeZone", m.DayMarginTimeZone())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Uint8Field("m_UseMasterFirm_SymbolLimitsArrayIsSet", m.UseMasterFirm_SymbolLimitsArrayIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolLimitsArray", m.UseMasterFirm_SymbolLimitsArray())
	w.Uint8Field("m_UseMasterFirm_TradeFeesIsSet", m.UseMasterFirm_TradeFeesIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFees", m.UseMasterFirm_TradeFees())
	w.Uint8Field("m_UseMasterFirm_TradeFeePerShareIsSet", m.UseMasterFirm_TradeFeePerShareIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFeePerShare", m.UseMasterFirm_TradeFeePerShare())
	w.Uint8Field("m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet", m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	w.Uint8Field("m_UseMasterFirm_RequireSufficientMarginForNewPositions", m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginIsSet", m.UseMasterFirm_UsePercentOfMarginIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMargin", m.UseMasterFirm_UsePercentOfMargin())
	w.Uint8Field("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Uint8Field("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Uint8Field("m_UseMasterFirm_MinimumRequiredAccountValueIsSet", m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	w.Uint8Field("m_UseMasterFirm_MinimumRequiredAccountValue", m.UseMasterFirm_MinimumRequiredAccountValue())
	w.Uint8Field("m_UseMasterFirm_MarginTimeSettingsIsSet", m.UseMasterFirm_MarginTimeSettingsIsSet())
	w.Uint8Field("m_UseMasterFirm_MarginTimeSettings", m.UseMasterFirm_MarginTimeSettings())
	w.Uint8Field("m_UseMasterFirm_TradingIsDisabledIsSet", m.UseMasterFirm_TradingIsDisabledIsSet())
	w.Uint8Field("m_UseMasterFirm_TradingIsDisabled", m.UseMasterFirm_TradingIsDisabled())
	w.Uint8Field("IsTradeStatisticsPublicallySharedIsSet", m.IsTradeStatisticsPublicallySharedIsSet())
	w.Uint8Field("IsTradeStatisticsPublicallyShared", m.IsTradeStatisticsPublicallyShared())
	w.Uint8Field("IsReadOnlyFollowingRequestsAllowedIsSet", m.IsReadOnlyFollowingRequestsAllowedIsSet())
	w.Uint8Field("IsReadOnlyFollowingRequestsAllowed", m.IsReadOnlyFollowingRequestsAllowed())
	w.Uint8Field("IsTradeAccountSharingAllowedIsSet", m.IsTradeAccountSharingAllowedIsSet())
	w.Uint8Field("IsTradeAccountSharingAllowed", m.IsTradeAccountSharingAllowed())
	w.Uint8Field("ReadOnlyShareToAllSCUsernamesIsSet", m.ReadOnlyShareToAllSCUsernamesIsSet())
	w.Uint8Field("ReadOnlyShareToAllSCUsernames", m.ReadOnlyShareToAllSCUsernames())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArrayIsSet", m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArray", m.UseMasterFirm_SymbolCommissionsArray())
	w.Uint8Field("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet", m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTrading", m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet", m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArrayFullOverride", m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	w.Uint8Field("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	w.Uint8Field("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet", m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginFullOverride", m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.Uint8Field("m_UseMasterFirm_TradeFeesFullOverrideIsSet", m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFeesFullOverride", m.UseMasterFirm_TradeFeesFullOverride())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.Uint8Field("m_LiquidationOnlyModeIsSet", m.LiquidationOnlyModeIsSet())
	w.Uint8Field("m_LiquidationOnlyMode", m.LiquidationOnlyMode())
	w.Uint8Field("m_CustomerOrFirmIsSet", m.CustomerOrFirmIsSet())
	w.Uint8Field("m_CustomerOrFirm", m.CustomerOrFirm())
	w.Uint8Field("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet", m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginIntraday", m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantityIsSet", m.MasterFirm_MaximumOrderQuantityIsSet())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantity", m.MasterFirm_MaximumOrderQuantity())
	w.Uint8Field("m_ExchangeTraderIdIsSet", m.ExchangeTraderIdIsSet())
	w.StringField("m_ExchangeTraderId", m.ExchangeTraderId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUpdatePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10117 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsNewAccount(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetNewAccountAuthorizedUsername(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrencyCodeIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrencyCode(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimitInAccountCurrencyIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimitInAccountCurrency(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetPercentOfCashBalanceForDailyNetLossLimit(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsAtDailyLossLimitIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsAtDailyLossLimit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosePositionsAtEndOfDayIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosePositionsAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginIntradayIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginIntraday(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderSubIDIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderSubID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderLocationIdIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderLocationId(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSelfMatchPreventionIDIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSelfMatchPreventionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCTICodeIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetCTICode(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccountIsReadOnlyIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccountIsReadOnly(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumGlobalPositionQuantityIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumGlobalPositionQuantity(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerContractIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerContract(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerShareIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerShare(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequireSufficientMarginForNewPositionsIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequireSufficientMarginForNewPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMarginIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMargin(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMarginForDayTradingIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMarginForDayTrading(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetFirmIDIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFirmID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsDisabledIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsDisabled(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescriptiveNameIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescriptiveName(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsMasterFirmControlAccountIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsMasterFirmControlAccount(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMinimumRequiredAccountValueIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMinimumRequiredAccountValue(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBeginTimeForDayMarginIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetBeginTimeForDayMargin(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetEndTimeForDayMarginIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetEndTimeForDayMargin(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDayMarginTimeZoneIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDayMarginTimeZone(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolLimitsArrayIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolLimitsArray(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeesIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFees(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeePerShareIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeePerShare(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_RequireSufficientMarginForNewPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMargin(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MinimumRequiredAccountValue(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MarginTimeSettingsIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MarginTimeSettings(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradingIsDisabledIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradingIsDisabled(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeStatisticsPublicallySharedIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeStatisticsPublicallyShared(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsReadOnlyFollowingRequestsAllowedIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsReadOnlyFollowingRequestsAllowed(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeAccountSharingAllowedIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeAccountSharingAllowed(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetReadOnlyShareToAllSCUsernamesIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetReadOnlyShareToAllSCUsernames(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArrayIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArray(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTrading(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeesFullOverrideIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeesFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetLiquidationOnlyModeIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetLiquidationOnlyMode(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetCustomerOrFirmIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetCustomerOrFirm(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_MaximumOrderQuantityIsSet(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_MaximumOrderQuantity(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchangeTraderIdIsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchangeTraderId(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUpdatePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10117 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "RequestID":
			m.SetRequestID(r.Uint32())
		case "IsNewAccount":
			m.SetIsNewAccount(r.Uint8())
		case "NewAccountAuthorizedUsername":
			m.SetNewAccountAuthorizedUsername(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "CurrencyCodeIsSet":
			m.SetCurrencyCodeIsSet(r.Uint8())
		case "CurrencyCode":
			m.SetCurrencyCode(r.String())
		case "DailyNetLossLimitInAccountCurrencyIsSet":
			m.SetDailyNetLossLimitInAccountCurrencyIsSet(r.Uint8())
		case "DailyNetLossLimitInAccountCurrency":
			m.SetDailyNetLossLimitInAccountCurrency(r.Float32())
		case "PercentOfCashBalanceForDailyNetLossLimitIsSet":
			m.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(r.Uint8())
		case "PercentOfCashBalanceForDailyNetLossLimit":
			m.SetPercentOfCashBalanceForDailyNetLossLimit(r.Int32())
		case "UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet":
			m.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(r.Uint8())
		case "UseTrailingAccountValueToNotAllowIncreaseInPositions":
			m.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(r.Uint8())
		case "DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet":
			m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(r.Uint8())
		case "DoNotAllowIncreaseInPositionsAtDailyLossLimit":
			m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
		case "FlattenPositionsAtDailyLossLimitIsSet":
			m.SetFlattenPositionsAtDailyLossLimitIsSet(r.Uint8())
		case "FlattenPositionsAtDailyLossLimit":
			m.SetFlattenPositionsAtDailyLossLimit(r.Uint8())
		case "ClosePositionsAtEndOfDayIsSet":
			m.SetClosePositionsAtEndOfDayIsSet(r.Uint8())
		case "ClosePositionsAtEndOfDay":
			m.SetClosePositionsAtEndOfDay(r.Uint8())
		case "FlattenPositionsWhenUnderMarginIntradayIsSet":
			m.SetFlattenPositionsWhenUnderMarginIntradayIsSet(r.Uint8())
		case "FlattenPositionsWhenUnderMarginIntraday":
			m.SetFlattenPositionsWhenUnderMarginIntraday(r.Uint8())
		case "FlattenPositionsWhenUnderMarginAtEndOfDayIsSet":
			m.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(r.Uint8())
		case "FlattenPositionsWhenUnderMarginAtEndOfDay":
			m.SetFlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
		case "SenderSubIDIsSet":
			m.SetSenderSubIDIsSet(r.Uint8())
		case "SenderSubID":
			m.SetSenderSubID(r.String())
		case "SenderLocationIdIsSet":
			m.SetSenderLocationIdIsSet(r.Uint8())
		case "SenderLocationId":
			m.SetSenderLocationId(r.String())
		case "SelfMatchPreventionIDIsSet":
			m.SetSelfMatchPreventionIDIsSet(r.Uint8())
		case "SelfMatchPreventionID":
			m.SetSelfMatchPreventionID(r.String())
		case "CTICodeIsSet":
			m.SetCTICodeIsSet(r.Uint8())
		case "CTICode":
			m.SetCTICode(r.Int32())
		case "TradeAccountIsReadOnlyIsSet":
			m.SetTradeAccountIsReadOnlyIsSet(r.Uint8())
		case "TradeAccountIsReadOnly":
			m.SetTradeAccountIsReadOnly(r.Uint8())
		case "MaximumGlobalPositionQuantityIsSet":
			m.SetMaximumGlobalPositionQuantityIsSet(r.Uint8())
		case "MaximumGlobalPositionQuantity":
			m.SetMaximumGlobalPositionQuantity(r.Int32())
		case "TradeFeePerContractIsSet":
			m.SetTradeFeePerContractIsSet(r.Uint8())
		case "TradeFeePerContract":
			m.SetTradeFeePerContract(r.Float64())
		case "TradeFeePerShareIsSet":
			m.SetTradeFeePerShareIsSet(r.Uint8())
		case "TradeFeePerShare":
			m.SetTradeFeePerShare(r.Float64())
		case "RequireSufficientMarginForNewPositionsIsSet":
			m.SetRequireSufficientMarginForNewPositionsIsSet(r.Uint8())
		case "RequireSufficientMarginForNewPositions":
			m.SetRequireSufficientMarginForNewPositions(r.Uint8())
		case "UsePercentOfMarginIsSet":
			m.SetUsePercentOfMarginIsSet(r.Uint8())
		case "UsePercentOfMargin":
			m.SetUsePercentOfMargin(r.Int32())
		case "UsePercentOfMarginForDayTradingIsSet":
			m.SetUsePercentOfMarginForDayTradingIsSet(r.Uint8())
		case "UsePercentOfMarginForDayTrading":
			m.SetUsePercentOfMarginForDayTrading(r.Int32())
		case "MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet":
			m.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(r.Uint8())
		case "MaximumAllowedAccountBalanceForPositionsAsPercentage":
			m.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(r.Int32())
		case "FirmIDIsSet":
			m.SetFirmIDIsSet(r.Uint8())
		case "FirmID":
			m.SetFirmID(r.String())
		case "TradingIsDisabledIsSet":
			m.SetTradingIsDisabledIsSet(r.Uint8())
		case "TradingIsDisabled":
			m.SetTradingIsDisabled(r.Uint8())
		case "DescriptiveNameIsSet":
			m.SetDescriptiveNameIsSet(r.Uint8())
		case "DescriptiveName":
			m.SetDescriptiveName(r.String())
		case "IsMasterFirmControlAccountIsSet":
			m.SetIsMasterFirmControlAccountIsSet(r.Uint8())
		case "IsMasterFirmControlAccount":
			m.SetIsMasterFirmControlAccount(r.Uint8())
		case "MinimumRequiredAccountValueIsSet":
			m.SetMinimumRequiredAccountValueIsSet(r.Uint8())
		case "MinimumRequiredAccountValue":
			m.SetMinimumRequiredAccountValue(r.Float64())
		case "BeginTimeForDayMarginIsSet":
			m.SetBeginTimeForDayMarginIsSet(r.Uint8())
		case "BeginTimeForDayMargin":
			m.SetBeginTimeForDayMargin(r.Int64())
		case "EndTimeForDayMarginIsSet":
			m.SetEndTimeForDayMarginIsSet(r.Uint8())
		case "EndTimeForDayMargin":
			m.SetEndTimeForDayMargin(r.Int64())
		case "DayMarginTimeZoneIsSet":
			m.SetDayMarginTimeZoneIsSet(r.Uint8())
		case "DayMarginTimeZone":
			m.SetDayMarginTimeZone(r.String())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(r.Uint8())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(r.Uint8())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(r.Uint8())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
		case "m_UseMasterFirm_SymbolLimitsArrayIsSet":
			m.SetUseMasterFirm_SymbolLimitsArrayIsSet(r.Uint8())
		case "m_UseMasterFirm_SymbolLimitsArray":
			m.SetUseMasterFirm_SymbolLimitsArray(r.Uint8())
		case "m_UseMasterFirm_TradeFeesIsSet":
			m.SetUseMasterFirm_TradeFeesIsSet(r.Uint8())
		case "m_UseMasterFirm_TradeFees":
			m.SetUseMasterFirm_TradeFees(r.Uint8())
		case "m_UseMasterFirm_TradeFeePerShareIsSet":
			m.SetUseMasterFirm_TradeFeePerShareIsSet(r.Uint8())
		case "m_UseMasterFirm_TradeFeePerShare":
			m.SetUseMasterFirm_TradeFeePerShare(r.Uint8())
		case "m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet":
			m.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(r.Uint8())
		case "m_UseMasterFirm_RequireSufficientMarginForNewPositions":
			m.SetUseMasterFirm_RequireSufficientMarginForNewPositions(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginIsSet":
			m.SetUseMasterFirm_UsePercentOfMarginIsSet(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMargin":
			m.SetUseMasterFirm_UsePercentOfMargin(r.Uint8())
		case "m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet":
			m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(r.Uint8())
		case "m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage":
			m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(r.Uint8())
		case "m_UseMasterFirm_MinimumRequiredAccountValueIsSet":
			m.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(r.Uint8())
		case "m_UseMasterFirm_MinimumRequiredAccountValue":
			m.SetUseMasterFirm_MinimumRequiredAccountValue(r.Uint8())
		case "m_UseMasterFirm_MarginTimeSettingsIsSet":
			m.SetUseMasterFirm_MarginTimeSettingsIsSet(r.Uint8())
		case "m_UseMasterFirm_MarginTimeSettings":
			m.SetUseMasterFirm_MarginTimeSettings(r.Uint8())
		case "m_UseMasterFirm_TradingIsDisabledIsSet":
			m.SetUseMasterFirm_TradingIsDisabledIsSet(r.Uint8())
		case "m_UseMasterFirm_TradingIsDisabled":
			m.SetUseMasterFirm_TradingIsDisabled(r.Uint8())
		case "IsTradeStatisticsPublicallySharedIsSet":
			m.SetIsTradeStatisticsPublicallySharedIsSet(r.Uint8())
		case "IsTradeStatisticsPublicallyShared":
			m.SetIsTradeStatisticsPublicallyShared(r.Uint8())
		case "IsReadOnlyFollowingRequestsAllowedIsSet":
			m.SetIsReadOnlyFollowingRequestsAllowedIsSet(r.Uint8())
		case "IsReadOnlyFollowingRequestsAllowed":
			m.SetIsReadOnlyFollowingRequestsAllowed(r.Uint8())
		case "IsTradeAccountSharingAllowedIsSet":
			m.SetIsTradeAccountSharingAllowedIsSet(r.Uint8())
		case "IsTradeAccountSharingAllowed":
			m.SetIsTradeAccountSharingAllowed(r.Uint8())
		case "ReadOnlyShareToAllSCUsernamesIsSet":
			m.SetReadOnlyShareToAllSCUsernamesIsSet(r.Uint8())
		case "ReadOnlyShareToAllSCUsernames":
			m.SetReadOnlyShareToAllSCUsernames(r.Uint8())
		case "m_UseMasterFirm_SymbolCommissionsArrayIsSet":
			m.SetUseMasterFirm_SymbolCommissionsArrayIsSet(r.Uint8())
		case "m_UseMasterFirm_SymbolCommissionsArray":
			m.SetUseMasterFirm_SymbolCommissionsArray(r.Uint8())
		case "m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet":
			m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(r.Uint8())
		case "m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit":
			m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTrading":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTrading(r.Uint8())
		case "m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet":
			m.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(r.Uint8())
		case "m_UseMasterFirm_SymbolCommissionsArrayFullOverride":
			m.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(r.Uint8())
		case "m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet":
			m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(r.Uint8())
		case "m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders":
			m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet":
			m.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginFullOverride":
			m.SetUseMasterFirm_UsePercentOfMarginFullOverride(r.Uint8())
		case "m_UseMasterFirm_TradeFeesFullOverrideIsSet":
			m.SetUseMasterFirm_TradeFeesFullOverrideIsSet(r.Uint8())
		case "m_UseMasterFirm_TradeFeesFullOverride":
			m.SetUseMasterFirm_TradeFeesFullOverride(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(r.Uint8())
		case "m_LiquidationOnlyModeIsSet":
			m.SetLiquidationOnlyModeIsSet(r.Uint8())
		case "m_LiquidationOnlyMode":
			m.SetLiquidationOnlyMode(r.Uint8())
		case "m_CustomerOrFirmIsSet":
			m.SetCustomerOrFirmIsSet(r.Uint8())
		case "m_CustomerOrFirm":
			m.SetCustomerOrFirm(r.Uint8())
		case "m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet":
			m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(r.Uint8())
		case "m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet":
			m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet":
			m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue":
			m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginIntraday":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(r.Uint8())
		case "m_MasterFirm_MaximumOrderQuantityIsSet":
			m.SetMasterFirm_MaximumOrderQuantityIsSet(r.Uint32())
		case "m_MasterFirm_MaximumOrderQuantity":
			m.SetMasterFirm_MaximumOrderQuantity(r.Uint32())
		case "m_ExchangeTraderIdIsSet":
			m.SetExchangeTraderIdIsSet(r.Uint8())
		case "m_ExchangeTraderId":
			m.SetExchangeTraderId(r.String())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}
