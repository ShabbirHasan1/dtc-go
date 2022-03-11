//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUpdate) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10117,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsNewAccount())
	w.Data = append(w.Data, ',')
	w.String(m.NewAccountAuthorizedUsername())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.CurrencyCodeIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.CurrencyCode())
	w.Data = append(w.Data, ',')
	w.Bool(m.DailyNetLossLimitInAccountCurrencyIsSet())
	w.Data = append(w.Data, ',')
	w.Float32(m.DailyNetLossLimitInAccountCurrency())
	w.Data = append(w.Data, ',')
	w.Bool(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.PercentOfCashBalanceForDailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.Data = append(w.Data, ',')
	w.Bool(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.FlattenPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.ClosePositionsAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Bool(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.SenderSubIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SenderSubID())
	w.Data = append(w.Data, ',')
	w.Bool(m.SenderLocationIdIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SenderLocationId())
	w.Data = append(w.Data, ',')
	w.Bool(m.SelfMatchPreventionIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.Bool(m.CTICodeIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.CTICode())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeAccountIsReadOnlyIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeAccountIsReadOnly())
	w.Data = append(w.Data, ',')
	w.Bool(m.MaximumGlobalPositionQuantityIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumGlobalPositionQuantity())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeFeePerContractIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeFeePerShareIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Bool(m.RequireSufficientMarginForNewPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Bool(m.UsePercentOfMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.UsePercentOfMarginForDayTradingIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Bool(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Bool(m.FirmIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.FirmID())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabledIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Bool(m.DescriptiveNameIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.DescriptiveName())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsMasterFirmControlAccountIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsMasterFirmControlAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.MinimumRequiredAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Bool(m.BeginTimeForDayMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int64(m.BeginTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.EndTimeForDayMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int64(m.EndTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.DayMarginTimeZoneIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.DayMarginTimeZone())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolLimitsArray())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradeFeesIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFees())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradeFeePerShareIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_MarginTimeSettingsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MarginTimeSettings())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradingIsDisabledIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeStatisticsPublicallySharedIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeStatisticsPublicallyShared())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsReadOnlyFollowingRequestsAllowed())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeAccountSharingAllowedIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeAccountSharingAllowed())
	w.Data = append(w.Data, ',')
	w.Bool(m.ReadOnlyShareToAllSCUsernamesIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReadOnlyShareToAllSCUsernames())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArray())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeesFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.LiquidationOnlyModeIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.LiquidationOnlyMode())
	w.Data = append(w.Data, ',')
	w.Bool(m.CustomerOrFirmIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CustomerOrFirm())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantityIsSet())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantity())
	w.Data = append(w.Data, ',')
	w.Bool(m.ExchangeTraderIdIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeTraderId())
	return w.Finish(), nil
}

func (m TradeAccountDataUpdateBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10117,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsNewAccount())
	w.Data = append(w.Data, ',')
	w.String(m.NewAccountAuthorizedUsername())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.CurrencyCodeIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.CurrencyCode())
	w.Data = append(w.Data, ',')
	w.Bool(m.DailyNetLossLimitInAccountCurrencyIsSet())
	w.Data = append(w.Data, ',')
	w.Float32(m.DailyNetLossLimitInAccountCurrency())
	w.Data = append(w.Data, ',')
	w.Bool(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.PercentOfCashBalanceForDailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.Data = append(w.Data, ',')
	w.Bool(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.FlattenPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.ClosePositionsAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Bool(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.SenderSubIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SenderSubID())
	w.Data = append(w.Data, ',')
	w.Bool(m.SenderLocationIdIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SenderLocationId())
	w.Data = append(w.Data, ',')
	w.Bool(m.SelfMatchPreventionIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.Bool(m.CTICodeIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.CTICode())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeAccountIsReadOnlyIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeAccountIsReadOnly())
	w.Data = append(w.Data, ',')
	w.Bool(m.MaximumGlobalPositionQuantityIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumGlobalPositionQuantity())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeFeePerContractIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeFeePerShareIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Bool(m.RequireSufficientMarginForNewPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Bool(m.UsePercentOfMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.UsePercentOfMarginForDayTradingIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Bool(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Bool(m.FirmIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.FirmID())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabledIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Bool(m.DescriptiveNameIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.DescriptiveName())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsMasterFirmControlAccountIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsMasterFirmControlAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.MinimumRequiredAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Bool(m.BeginTimeForDayMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int64(m.BeginTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.EndTimeForDayMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int64(m.EndTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.DayMarginTimeZoneIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.DayMarginTimeZone())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolLimitsArray())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradeFeesIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFees())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradeFeePerShareIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_MarginTimeSettingsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MarginTimeSettings())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradingIsDisabledIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeStatisticsPublicallySharedIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeStatisticsPublicallyShared())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsReadOnlyFollowingRequestsAllowed())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeAccountSharingAllowedIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeAccountSharingAllowed())
	w.Data = append(w.Data, ',')
	w.Bool(m.ReadOnlyShareToAllSCUsernamesIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReadOnlyShareToAllSCUsernames())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArray())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeesFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.LiquidationOnlyModeIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.LiquidationOnlyMode())
	w.Data = append(w.Data, ',')
	w.Bool(m.CustomerOrFirmIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CustomerOrFirm())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantityIsSet())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantity())
	w.Data = append(w.Data, ',')
	w.Bool(m.ExchangeTraderIdIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeTraderId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUpdatePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10117,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsNewAccount())
	w.Data = append(w.Data, ',')
	w.String(m.NewAccountAuthorizedUsername())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.CurrencyCodeIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.CurrencyCode())
	w.Data = append(w.Data, ',')
	w.Bool(m.DailyNetLossLimitInAccountCurrencyIsSet())
	w.Data = append(w.Data, ',')
	w.Float32(m.DailyNetLossLimitInAccountCurrency())
	w.Data = append(w.Data, ',')
	w.Bool(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.PercentOfCashBalanceForDailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.Data = append(w.Data, ',')
	w.Bool(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.FlattenPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.ClosePositionsAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Bool(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.SenderSubIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SenderSubID())
	w.Data = append(w.Data, ',')
	w.Bool(m.SenderLocationIdIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SenderLocationId())
	w.Data = append(w.Data, ',')
	w.Bool(m.SelfMatchPreventionIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.Bool(m.CTICodeIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.CTICode())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeAccountIsReadOnlyIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeAccountIsReadOnly())
	w.Data = append(w.Data, ',')
	w.Bool(m.MaximumGlobalPositionQuantityIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumGlobalPositionQuantity())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeFeePerContractIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeFeePerShareIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Bool(m.RequireSufficientMarginForNewPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Bool(m.UsePercentOfMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.UsePercentOfMarginForDayTradingIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Bool(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Bool(m.FirmIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.FirmID())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabledIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Bool(m.DescriptiveNameIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.DescriptiveName())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsMasterFirmControlAccountIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsMasterFirmControlAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.MinimumRequiredAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Bool(m.BeginTimeForDayMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int64(m.BeginTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.EndTimeForDayMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int64(m.EndTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.DayMarginTimeZoneIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.DayMarginTimeZone())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolLimitsArray())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradeFeesIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFees())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradeFeePerShareIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_MarginTimeSettingsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MarginTimeSettings())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradingIsDisabledIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeStatisticsPublicallySharedIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeStatisticsPublicallyShared())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsReadOnlyFollowingRequestsAllowed())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeAccountSharingAllowedIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeAccountSharingAllowed())
	w.Data = append(w.Data, ',')
	w.Bool(m.ReadOnlyShareToAllSCUsernamesIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReadOnlyShareToAllSCUsernames())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArray())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeesFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.LiquidationOnlyModeIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.LiquidationOnlyMode())
	w.Data = append(w.Data, ',')
	w.Bool(m.CustomerOrFirmIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CustomerOrFirm())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantityIsSet())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantity())
	w.Data = append(w.Data, ',')
	w.Bool(m.ExchangeTraderIdIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeTraderId())
	return w.Finish(), nil
}

func (m TradeAccountDataUpdatePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10117,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsNewAccount())
	w.Data = append(w.Data, ',')
	w.String(m.NewAccountAuthorizedUsername())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.CurrencyCodeIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.CurrencyCode())
	w.Data = append(w.Data, ',')
	w.Bool(m.DailyNetLossLimitInAccountCurrencyIsSet())
	w.Data = append(w.Data, ',')
	w.Float32(m.DailyNetLossLimitInAccountCurrency())
	w.Data = append(w.Data, ',')
	w.Bool(m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.PercentOfCashBalanceForDailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.Data = append(w.Data, ',')
	w.Bool(m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.FlattenPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.ClosePositionsAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Bool(m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.SenderSubIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SenderSubID())
	w.Data = append(w.Data, ',')
	w.Bool(m.SenderLocationIdIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SenderLocationId())
	w.Data = append(w.Data, ',')
	w.Bool(m.SelfMatchPreventionIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.Bool(m.CTICodeIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.CTICode())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeAccountIsReadOnlyIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeAccountIsReadOnly())
	w.Data = append(w.Data, ',')
	w.Bool(m.MaximumGlobalPositionQuantityIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumGlobalPositionQuantity())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeFeePerContractIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradeFeePerShareIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Bool(m.RequireSufficientMarginForNewPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Bool(m.UsePercentOfMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.UsePercentOfMarginForDayTradingIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Bool(m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Bool(m.FirmIDIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.FirmID())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabledIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Bool(m.DescriptiveNameIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.DescriptiveName())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsMasterFirmControlAccountIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsMasterFirmControlAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.MinimumRequiredAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Float64(m.MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Bool(m.BeginTimeForDayMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int64(m.BeginTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.EndTimeForDayMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Int64(m.EndTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.DayMarginTimeZoneIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.DayMarginTimeZone())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_SymbolLimitsArrayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolLimitsArray())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradeFeesIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFees())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradeFeePerShareIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_MarginTimeSettingsIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MarginTimeSettings())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradingIsDisabledIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeStatisticsPublicallySharedIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeStatisticsPublicallyShared())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsReadOnlyFollowingRequestsAllowedIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsReadOnlyFollowingRequestsAllowed())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeAccountSharingAllowedIsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsTradeAccountSharingAllowed())
	w.Data = append(w.Data, ',')
	w.Bool(m.ReadOnlyShareToAllSCUsernamesIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReadOnlyShareToAllSCUsernames())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArray())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeesFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.Data = append(w.Data, ',')
	w.Bool(m.LiquidationOnlyModeIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.LiquidationOnlyMode())
	w.Data = append(w.Data, ',')
	w.Bool(m.CustomerOrFirmIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CustomerOrFirm())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Bool(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantityIsSet())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantity())
	w.Data = append(w.Data, ',')
	w.Bool(m.ExchangeTraderIdIsSet())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeTraderId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUpdate) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10117)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsNewAccount", m.IsNewAccount())
	w.StringField("NewAccountAuthorizedUsername", m.NewAccountAuthorizedUsername())
	w.StringField("TradeAccount", m.TradeAccount())
	w.BoolField("CurrencyCodeIsSet", m.CurrencyCodeIsSet())
	w.StringField("CurrencyCode", m.CurrencyCode())
	w.BoolField("DailyNetLossLimitInAccountCurrencyIsSet", m.DailyNetLossLimitInAccountCurrencyIsSet())
	w.Float32Field("DailyNetLossLimitInAccountCurrency", m.DailyNetLossLimitInAccountCurrency())
	w.BoolField("PercentOfCashBalanceForDailyNetLossLimitIsSet", m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	w.Int32Field("PercentOfCashBalanceForDailyNetLossLimit", m.PercentOfCashBalanceForDailyNetLossLimit())
	w.BoolField("UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet", m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	w.Uint8Field("UseTrailingAccountValueToNotAllowIncreaseInPositions", m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.BoolField("DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet", m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.BoolField("FlattenPositionsAtDailyLossLimitIsSet", m.FlattenPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("FlattenPositionsAtDailyLossLimit", m.FlattenPositionsAtDailyLossLimit())
	w.BoolField("ClosePositionsAtEndOfDayIsSet", m.ClosePositionsAtEndOfDayIsSet())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.BoolField("FlattenPositionsWhenUnderMarginIntradayIsSet", m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Uint8Field("FlattenPositionsWhenUnderMarginIntraday", m.FlattenPositionsWhenUnderMarginIntraday())
	w.BoolField("FlattenPositionsWhenUnderMarginAtEndOfDayIsSet", m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Uint8Field("FlattenPositionsWhenUnderMarginAtEndOfDay", m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.BoolField("SenderSubIDIsSet", m.SenderSubIDIsSet())
	w.StringField("SenderSubID", m.SenderSubID())
	w.BoolField("SenderLocationIdIsSet", m.SenderLocationIdIsSet())
	w.StringField("SenderLocationId", m.SenderLocationId())
	w.BoolField("SelfMatchPreventionIDIsSet", m.SelfMatchPreventionIDIsSet())
	w.StringField("SelfMatchPreventionID", m.SelfMatchPreventionID())
	w.BoolField("CTICodeIsSet", m.CTICodeIsSet())
	w.Int32Field("CTICode", m.CTICode())
	w.BoolField("TradeAccountIsReadOnlyIsSet", m.TradeAccountIsReadOnlyIsSet())
	w.BoolField("TradeAccountIsReadOnly", m.TradeAccountIsReadOnly())
	w.BoolField("MaximumGlobalPositionQuantityIsSet", m.MaximumGlobalPositionQuantityIsSet())
	w.Int32Field("MaximumGlobalPositionQuantity", m.MaximumGlobalPositionQuantity())
	w.BoolField("TradeFeePerContractIsSet", m.TradeFeePerContractIsSet())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	w.BoolField("TradeFeePerShareIsSet", m.TradeFeePerShareIsSet())
	w.Float64Field("TradeFeePerShare", m.TradeFeePerShare())
	w.BoolField("RequireSufficientMarginForNewPositionsIsSet", m.RequireSufficientMarginForNewPositionsIsSet())
	w.Uint8Field("RequireSufficientMarginForNewPositions", m.RequireSufficientMarginForNewPositions())
	w.BoolField("UsePercentOfMarginIsSet", m.UsePercentOfMarginIsSet())
	w.Int32Field("UsePercentOfMargin", m.UsePercentOfMargin())
	w.BoolField("UsePercentOfMarginForDayTradingIsSet", m.UsePercentOfMarginForDayTradingIsSet())
	w.Int32Field("UsePercentOfMarginForDayTrading", m.UsePercentOfMarginForDayTrading())
	w.BoolField("MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet", m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Int32Field("MaximumAllowedAccountBalanceForPositionsAsPercentage", m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.BoolField("FirmIDIsSet", m.FirmIDIsSet())
	w.StringField("FirmID", m.FirmID())
	w.BoolField("TradingIsDisabledIsSet", m.TradingIsDisabledIsSet())
	w.BoolField("TradingIsDisabled", m.TradingIsDisabled())
	w.BoolField("DescriptiveNameIsSet", m.DescriptiveNameIsSet())
	w.StringField("DescriptiveName", m.DescriptiveName())
	w.BoolField("IsMasterFirmControlAccountIsSet", m.IsMasterFirmControlAccountIsSet())
	w.BoolField("IsMasterFirmControlAccount", m.IsMasterFirmControlAccount())
	w.BoolField("MinimumRequiredAccountValueIsSet", m.MinimumRequiredAccountValueIsSet())
	w.Float64Field("MinimumRequiredAccountValue", m.MinimumRequiredAccountValue())
	w.BoolField("BeginTimeForDayMarginIsSet", m.BeginTimeForDayMarginIsSet())
	w.Int64Field("BeginTimeForDayMargin", m.BeginTimeForDayMargin())
	w.BoolField("EndTimeForDayMarginIsSet", m.EndTimeForDayMarginIsSet())
	w.Int64Field("EndTimeForDayMargin", m.EndTimeForDayMargin())
	w.BoolField("DayMarginTimeZoneIsSet", m.DayMarginTimeZoneIsSet())
	w.StringField("DayMarginTimeZone", m.DayMarginTimeZone())
	w.BoolField("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.BoolField("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.BoolField("m_UseMasterFirm_SymbolLimitsArrayIsSet", m.UseMasterFirm_SymbolLimitsArrayIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolLimitsArray", m.UseMasterFirm_SymbolLimitsArray())
	w.BoolField("m_UseMasterFirm_TradeFeesIsSet", m.UseMasterFirm_TradeFeesIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFees", m.UseMasterFirm_TradeFees())
	w.BoolField("m_UseMasterFirm_TradeFeePerShareIsSet", m.UseMasterFirm_TradeFeePerShareIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFeePerShare", m.UseMasterFirm_TradeFeePerShare())
	w.BoolField("m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet", m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	w.Uint8Field("m_UseMasterFirm_RequireSufficientMarginForNewPositions", m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginIsSet", m.UseMasterFirm_UsePercentOfMarginIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMargin", m.UseMasterFirm_UsePercentOfMargin())
	w.BoolField("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Uint8Field("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.BoolField("m_UseMasterFirm_MinimumRequiredAccountValueIsSet", m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	w.Uint8Field("m_UseMasterFirm_MinimumRequiredAccountValue", m.UseMasterFirm_MinimumRequiredAccountValue())
	w.BoolField("m_UseMasterFirm_MarginTimeSettingsIsSet", m.UseMasterFirm_MarginTimeSettingsIsSet())
	w.Uint8Field("m_UseMasterFirm_MarginTimeSettings", m.UseMasterFirm_MarginTimeSettings())
	w.BoolField("m_UseMasterFirm_TradingIsDisabledIsSet", m.UseMasterFirm_TradingIsDisabledIsSet())
	w.BoolField("m_UseMasterFirm_TradingIsDisabled", m.UseMasterFirm_TradingIsDisabled())
	w.BoolField("IsTradeStatisticsPublicallySharedIsSet", m.IsTradeStatisticsPublicallySharedIsSet())
	w.BoolField("IsTradeStatisticsPublicallyShared", m.IsTradeStatisticsPublicallyShared())
	w.BoolField("IsReadOnlyFollowingRequestsAllowedIsSet", m.IsReadOnlyFollowingRequestsAllowedIsSet())
	w.BoolField("IsReadOnlyFollowingRequestsAllowed", m.IsReadOnlyFollowingRequestsAllowed())
	w.BoolField("IsTradeAccountSharingAllowedIsSet", m.IsTradeAccountSharingAllowedIsSet())
	w.BoolField("IsTradeAccountSharingAllowed", m.IsTradeAccountSharingAllowed())
	w.BoolField("ReadOnlyShareToAllSCUsernamesIsSet", m.ReadOnlyShareToAllSCUsernamesIsSet())
	w.Uint8Field("ReadOnlyShareToAllSCUsernames", m.ReadOnlyShareToAllSCUsernames())
	w.BoolField("m_UseMasterFirm_SymbolCommissionsArrayIsSet", m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArray", m.UseMasterFirm_SymbolCommissionsArray())
	w.BoolField("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet", m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTrading", m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.BoolField("m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet", m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArrayFullOverride", m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	w.BoolField("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	w.Uint8Field("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet", m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginFullOverride", m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.BoolField("m_UseMasterFirm_TradeFeesFullOverrideIsSet", m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFeesFullOverride", m.UseMasterFirm_TradeFeesFullOverride())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.BoolField("m_LiquidationOnlyModeIsSet", m.LiquidationOnlyModeIsSet())
	w.Uint8Field("m_LiquidationOnlyMode", m.LiquidationOnlyMode())
	w.BoolField("m_CustomerOrFirmIsSet", m.CustomerOrFirmIsSet())
	w.Uint8Field("m_CustomerOrFirm", m.CustomerOrFirm())
	w.BoolField("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.BoolField("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.BoolField("m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet", m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginIntraday", m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.BoolField("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantityIsSet", m.MasterFirm_MaximumOrderQuantityIsSet())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantity", m.MasterFirm_MaximumOrderQuantity())
	w.BoolField("m_ExchangeTraderIdIsSet", m.ExchangeTraderIdIsSet())
	w.StringField("m_ExchangeTraderId", m.ExchangeTraderId())
	return w.Finish(), nil
}

func (m TradeAccountDataUpdateBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10117)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsNewAccount", m.IsNewAccount())
	w.StringField("NewAccountAuthorizedUsername", m.NewAccountAuthorizedUsername())
	w.StringField("TradeAccount", m.TradeAccount())
	w.BoolField("CurrencyCodeIsSet", m.CurrencyCodeIsSet())
	w.StringField("CurrencyCode", m.CurrencyCode())
	w.BoolField("DailyNetLossLimitInAccountCurrencyIsSet", m.DailyNetLossLimitInAccountCurrencyIsSet())
	w.Float32Field("DailyNetLossLimitInAccountCurrency", m.DailyNetLossLimitInAccountCurrency())
	w.BoolField("PercentOfCashBalanceForDailyNetLossLimitIsSet", m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	w.Int32Field("PercentOfCashBalanceForDailyNetLossLimit", m.PercentOfCashBalanceForDailyNetLossLimit())
	w.BoolField("UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet", m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	w.Uint8Field("UseTrailingAccountValueToNotAllowIncreaseInPositions", m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.BoolField("DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet", m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.BoolField("FlattenPositionsAtDailyLossLimitIsSet", m.FlattenPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("FlattenPositionsAtDailyLossLimit", m.FlattenPositionsAtDailyLossLimit())
	w.BoolField("ClosePositionsAtEndOfDayIsSet", m.ClosePositionsAtEndOfDayIsSet())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.BoolField("FlattenPositionsWhenUnderMarginIntradayIsSet", m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Uint8Field("FlattenPositionsWhenUnderMarginIntraday", m.FlattenPositionsWhenUnderMarginIntraday())
	w.BoolField("FlattenPositionsWhenUnderMarginAtEndOfDayIsSet", m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Uint8Field("FlattenPositionsWhenUnderMarginAtEndOfDay", m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.BoolField("SenderSubIDIsSet", m.SenderSubIDIsSet())
	w.StringField("SenderSubID", m.SenderSubID())
	w.BoolField("SenderLocationIdIsSet", m.SenderLocationIdIsSet())
	w.StringField("SenderLocationId", m.SenderLocationId())
	w.BoolField("SelfMatchPreventionIDIsSet", m.SelfMatchPreventionIDIsSet())
	w.StringField("SelfMatchPreventionID", m.SelfMatchPreventionID())
	w.BoolField("CTICodeIsSet", m.CTICodeIsSet())
	w.Int32Field("CTICode", m.CTICode())
	w.BoolField("TradeAccountIsReadOnlyIsSet", m.TradeAccountIsReadOnlyIsSet())
	w.BoolField("TradeAccountIsReadOnly", m.TradeAccountIsReadOnly())
	w.BoolField("MaximumGlobalPositionQuantityIsSet", m.MaximumGlobalPositionQuantityIsSet())
	w.Int32Field("MaximumGlobalPositionQuantity", m.MaximumGlobalPositionQuantity())
	w.BoolField("TradeFeePerContractIsSet", m.TradeFeePerContractIsSet())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	w.BoolField("TradeFeePerShareIsSet", m.TradeFeePerShareIsSet())
	w.Float64Field("TradeFeePerShare", m.TradeFeePerShare())
	w.BoolField("RequireSufficientMarginForNewPositionsIsSet", m.RequireSufficientMarginForNewPositionsIsSet())
	w.Uint8Field("RequireSufficientMarginForNewPositions", m.RequireSufficientMarginForNewPositions())
	w.BoolField("UsePercentOfMarginIsSet", m.UsePercentOfMarginIsSet())
	w.Int32Field("UsePercentOfMargin", m.UsePercentOfMargin())
	w.BoolField("UsePercentOfMarginForDayTradingIsSet", m.UsePercentOfMarginForDayTradingIsSet())
	w.Int32Field("UsePercentOfMarginForDayTrading", m.UsePercentOfMarginForDayTrading())
	w.BoolField("MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet", m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Int32Field("MaximumAllowedAccountBalanceForPositionsAsPercentage", m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.BoolField("FirmIDIsSet", m.FirmIDIsSet())
	w.StringField("FirmID", m.FirmID())
	w.BoolField("TradingIsDisabledIsSet", m.TradingIsDisabledIsSet())
	w.BoolField("TradingIsDisabled", m.TradingIsDisabled())
	w.BoolField("DescriptiveNameIsSet", m.DescriptiveNameIsSet())
	w.StringField("DescriptiveName", m.DescriptiveName())
	w.BoolField("IsMasterFirmControlAccountIsSet", m.IsMasterFirmControlAccountIsSet())
	w.BoolField("IsMasterFirmControlAccount", m.IsMasterFirmControlAccount())
	w.BoolField("MinimumRequiredAccountValueIsSet", m.MinimumRequiredAccountValueIsSet())
	w.Float64Field("MinimumRequiredAccountValue", m.MinimumRequiredAccountValue())
	w.BoolField("BeginTimeForDayMarginIsSet", m.BeginTimeForDayMarginIsSet())
	w.Int64Field("BeginTimeForDayMargin", m.BeginTimeForDayMargin())
	w.BoolField("EndTimeForDayMarginIsSet", m.EndTimeForDayMarginIsSet())
	w.Int64Field("EndTimeForDayMargin", m.EndTimeForDayMargin())
	w.BoolField("DayMarginTimeZoneIsSet", m.DayMarginTimeZoneIsSet())
	w.StringField("DayMarginTimeZone", m.DayMarginTimeZone())
	w.BoolField("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.BoolField("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.BoolField("m_UseMasterFirm_SymbolLimitsArrayIsSet", m.UseMasterFirm_SymbolLimitsArrayIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolLimitsArray", m.UseMasterFirm_SymbolLimitsArray())
	w.BoolField("m_UseMasterFirm_TradeFeesIsSet", m.UseMasterFirm_TradeFeesIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFees", m.UseMasterFirm_TradeFees())
	w.BoolField("m_UseMasterFirm_TradeFeePerShareIsSet", m.UseMasterFirm_TradeFeePerShareIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFeePerShare", m.UseMasterFirm_TradeFeePerShare())
	w.BoolField("m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet", m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	w.Uint8Field("m_UseMasterFirm_RequireSufficientMarginForNewPositions", m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginIsSet", m.UseMasterFirm_UsePercentOfMarginIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMargin", m.UseMasterFirm_UsePercentOfMargin())
	w.BoolField("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Uint8Field("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.BoolField("m_UseMasterFirm_MinimumRequiredAccountValueIsSet", m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	w.Uint8Field("m_UseMasterFirm_MinimumRequiredAccountValue", m.UseMasterFirm_MinimumRequiredAccountValue())
	w.BoolField("m_UseMasterFirm_MarginTimeSettingsIsSet", m.UseMasterFirm_MarginTimeSettingsIsSet())
	w.Uint8Field("m_UseMasterFirm_MarginTimeSettings", m.UseMasterFirm_MarginTimeSettings())
	w.BoolField("m_UseMasterFirm_TradingIsDisabledIsSet", m.UseMasterFirm_TradingIsDisabledIsSet())
	w.BoolField("m_UseMasterFirm_TradingIsDisabled", m.UseMasterFirm_TradingIsDisabled())
	w.BoolField("IsTradeStatisticsPublicallySharedIsSet", m.IsTradeStatisticsPublicallySharedIsSet())
	w.BoolField("IsTradeStatisticsPublicallyShared", m.IsTradeStatisticsPublicallyShared())
	w.BoolField("IsReadOnlyFollowingRequestsAllowedIsSet", m.IsReadOnlyFollowingRequestsAllowedIsSet())
	w.BoolField("IsReadOnlyFollowingRequestsAllowed", m.IsReadOnlyFollowingRequestsAllowed())
	w.BoolField("IsTradeAccountSharingAllowedIsSet", m.IsTradeAccountSharingAllowedIsSet())
	w.BoolField("IsTradeAccountSharingAllowed", m.IsTradeAccountSharingAllowed())
	w.BoolField("ReadOnlyShareToAllSCUsernamesIsSet", m.ReadOnlyShareToAllSCUsernamesIsSet())
	w.Uint8Field("ReadOnlyShareToAllSCUsernames", m.ReadOnlyShareToAllSCUsernames())
	w.BoolField("m_UseMasterFirm_SymbolCommissionsArrayIsSet", m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArray", m.UseMasterFirm_SymbolCommissionsArray())
	w.BoolField("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet", m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTrading", m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.BoolField("m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet", m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArrayFullOverride", m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	w.BoolField("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	w.Uint8Field("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet", m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginFullOverride", m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.BoolField("m_UseMasterFirm_TradeFeesFullOverrideIsSet", m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFeesFullOverride", m.UseMasterFirm_TradeFeesFullOverride())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.BoolField("m_LiquidationOnlyModeIsSet", m.LiquidationOnlyModeIsSet())
	w.Uint8Field("m_LiquidationOnlyMode", m.LiquidationOnlyMode())
	w.BoolField("m_CustomerOrFirmIsSet", m.CustomerOrFirmIsSet())
	w.Uint8Field("m_CustomerOrFirm", m.CustomerOrFirm())
	w.BoolField("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.BoolField("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.BoolField("m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet", m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginIntraday", m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.BoolField("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantityIsSet", m.MasterFirm_MaximumOrderQuantityIsSet())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantity", m.MasterFirm_MaximumOrderQuantity())
	w.BoolField("m_ExchangeTraderIdIsSet", m.ExchangeTraderIdIsSet())
	w.StringField("m_ExchangeTraderId", m.ExchangeTraderId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUpdatePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10117)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsNewAccount", m.IsNewAccount())
	w.StringField("NewAccountAuthorizedUsername", m.NewAccountAuthorizedUsername())
	w.StringField("TradeAccount", m.TradeAccount())
	w.BoolField("CurrencyCodeIsSet", m.CurrencyCodeIsSet())
	w.StringField("CurrencyCode", m.CurrencyCode())
	w.BoolField("DailyNetLossLimitInAccountCurrencyIsSet", m.DailyNetLossLimitInAccountCurrencyIsSet())
	w.Float32Field("DailyNetLossLimitInAccountCurrency", m.DailyNetLossLimitInAccountCurrency())
	w.BoolField("PercentOfCashBalanceForDailyNetLossLimitIsSet", m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	w.Int32Field("PercentOfCashBalanceForDailyNetLossLimit", m.PercentOfCashBalanceForDailyNetLossLimit())
	w.BoolField("UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet", m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	w.Uint8Field("UseTrailingAccountValueToNotAllowIncreaseInPositions", m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.BoolField("DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet", m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.BoolField("FlattenPositionsAtDailyLossLimitIsSet", m.FlattenPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("FlattenPositionsAtDailyLossLimit", m.FlattenPositionsAtDailyLossLimit())
	w.BoolField("ClosePositionsAtEndOfDayIsSet", m.ClosePositionsAtEndOfDayIsSet())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.BoolField("FlattenPositionsWhenUnderMarginIntradayIsSet", m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Uint8Field("FlattenPositionsWhenUnderMarginIntraday", m.FlattenPositionsWhenUnderMarginIntraday())
	w.BoolField("FlattenPositionsWhenUnderMarginAtEndOfDayIsSet", m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Uint8Field("FlattenPositionsWhenUnderMarginAtEndOfDay", m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.BoolField("SenderSubIDIsSet", m.SenderSubIDIsSet())
	w.StringField("SenderSubID", m.SenderSubID())
	w.BoolField("SenderLocationIdIsSet", m.SenderLocationIdIsSet())
	w.StringField("SenderLocationId", m.SenderLocationId())
	w.BoolField("SelfMatchPreventionIDIsSet", m.SelfMatchPreventionIDIsSet())
	w.StringField("SelfMatchPreventionID", m.SelfMatchPreventionID())
	w.BoolField("CTICodeIsSet", m.CTICodeIsSet())
	w.Int32Field("CTICode", m.CTICode())
	w.BoolField("TradeAccountIsReadOnlyIsSet", m.TradeAccountIsReadOnlyIsSet())
	w.BoolField("TradeAccountIsReadOnly", m.TradeAccountIsReadOnly())
	w.BoolField("MaximumGlobalPositionQuantityIsSet", m.MaximumGlobalPositionQuantityIsSet())
	w.Int32Field("MaximumGlobalPositionQuantity", m.MaximumGlobalPositionQuantity())
	w.BoolField("TradeFeePerContractIsSet", m.TradeFeePerContractIsSet())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	w.BoolField("TradeFeePerShareIsSet", m.TradeFeePerShareIsSet())
	w.Float64Field("TradeFeePerShare", m.TradeFeePerShare())
	w.BoolField("RequireSufficientMarginForNewPositionsIsSet", m.RequireSufficientMarginForNewPositionsIsSet())
	w.Uint8Field("RequireSufficientMarginForNewPositions", m.RequireSufficientMarginForNewPositions())
	w.BoolField("UsePercentOfMarginIsSet", m.UsePercentOfMarginIsSet())
	w.Int32Field("UsePercentOfMargin", m.UsePercentOfMargin())
	w.BoolField("UsePercentOfMarginForDayTradingIsSet", m.UsePercentOfMarginForDayTradingIsSet())
	w.Int32Field("UsePercentOfMarginForDayTrading", m.UsePercentOfMarginForDayTrading())
	w.BoolField("MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet", m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Int32Field("MaximumAllowedAccountBalanceForPositionsAsPercentage", m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.BoolField("FirmIDIsSet", m.FirmIDIsSet())
	w.StringField("FirmID", m.FirmID())
	w.BoolField("TradingIsDisabledIsSet", m.TradingIsDisabledIsSet())
	w.BoolField("TradingIsDisabled", m.TradingIsDisabled())
	w.BoolField("DescriptiveNameIsSet", m.DescriptiveNameIsSet())
	w.StringField("DescriptiveName", m.DescriptiveName())
	w.BoolField("IsMasterFirmControlAccountIsSet", m.IsMasterFirmControlAccountIsSet())
	w.BoolField("IsMasterFirmControlAccount", m.IsMasterFirmControlAccount())
	w.BoolField("MinimumRequiredAccountValueIsSet", m.MinimumRequiredAccountValueIsSet())
	w.Float64Field("MinimumRequiredAccountValue", m.MinimumRequiredAccountValue())
	w.BoolField("BeginTimeForDayMarginIsSet", m.BeginTimeForDayMarginIsSet())
	w.Int64Field("BeginTimeForDayMargin", m.BeginTimeForDayMargin())
	w.BoolField("EndTimeForDayMarginIsSet", m.EndTimeForDayMarginIsSet())
	w.Int64Field("EndTimeForDayMargin", m.EndTimeForDayMargin())
	w.BoolField("DayMarginTimeZoneIsSet", m.DayMarginTimeZoneIsSet())
	w.StringField("DayMarginTimeZone", m.DayMarginTimeZone())
	w.BoolField("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.BoolField("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.BoolField("m_UseMasterFirm_SymbolLimitsArrayIsSet", m.UseMasterFirm_SymbolLimitsArrayIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolLimitsArray", m.UseMasterFirm_SymbolLimitsArray())
	w.BoolField("m_UseMasterFirm_TradeFeesIsSet", m.UseMasterFirm_TradeFeesIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFees", m.UseMasterFirm_TradeFees())
	w.BoolField("m_UseMasterFirm_TradeFeePerShareIsSet", m.UseMasterFirm_TradeFeePerShareIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFeePerShare", m.UseMasterFirm_TradeFeePerShare())
	w.BoolField("m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet", m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	w.Uint8Field("m_UseMasterFirm_RequireSufficientMarginForNewPositions", m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginIsSet", m.UseMasterFirm_UsePercentOfMarginIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMargin", m.UseMasterFirm_UsePercentOfMargin())
	w.BoolField("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Uint8Field("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.BoolField("m_UseMasterFirm_MinimumRequiredAccountValueIsSet", m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	w.Uint8Field("m_UseMasterFirm_MinimumRequiredAccountValue", m.UseMasterFirm_MinimumRequiredAccountValue())
	w.BoolField("m_UseMasterFirm_MarginTimeSettingsIsSet", m.UseMasterFirm_MarginTimeSettingsIsSet())
	w.Uint8Field("m_UseMasterFirm_MarginTimeSettings", m.UseMasterFirm_MarginTimeSettings())
	w.BoolField("m_UseMasterFirm_TradingIsDisabledIsSet", m.UseMasterFirm_TradingIsDisabledIsSet())
	w.BoolField("m_UseMasterFirm_TradingIsDisabled", m.UseMasterFirm_TradingIsDisabled())
	w.BoolField("IsTradeStatisticsPublicallySharedIsSet", m.IsTradeStatisticsPublicallySharedIsSet())
	w.BoolField("IsTradeStatisticsPublicallyShared", m.IsTradeStatisticsPublicallyShared())
	w.BoolField("IsReadOnlyFollowingRequestsAllowedIsSet", m.IsReadOnlyFollowingRequestsAllowedIsSet())
	w.BoolField("IsReadOnlyFollowingRequestsAllowed", m.IsReadOnlyFollowingRequestsAllowed())
	w.BoolField("IsTradeAccountSharingAllowedIsSet", m.IsTradeAccountSharingAllowedIsSet())
	w.BoolField("IsTradeAccountSharingAllowed", m.IsTradeAccountSharingAllowed())
	w.BoolField("ReadOnlyShareToAllSCUsernamesIsSet", m.ReadOnlyShareToAllSCUsernamesIsSet())
	w.Uint8Field("ReadOnlyShareToAllSCUsernames", m.ReadOnlyShareToAllSCUsernames())
	w.BoolField("m_UseMasterFirm_SymbolCommissionsArrayIsSet", m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArray", m.UseMasterFirm_SymbolCommissionsArray())
	w.BoolField("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet", m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTrading", m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.BoolField("m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet", m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArrayFullOverride", m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	w.BoolField("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	w.Uint8Field("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet", m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginFullOverride", m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.BoolField("m_UseMasterFirm_TradeFeesFullOverrideIsSet", m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFeesFullOverride", m.UseMasterFirm_TradeFeesFullOverride())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.BoolField("m_LiquidationOnlyModeIsSet", m.LiquidationOnlyModeIsSet())
	w.Uint8Field("m_LiquidationOnlyMode", m.LiquidationOnlyMode())
	w.BoolField("m_CustomerOrFirmIsSet", m.CustomerOrFirmIsSet())
	w.Uint8Field("m_CustomerOrFirm", m.CustomerOrFirm())
	w.BoolField("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.BoolField("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.BoolField("m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet", m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginIntraday", m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.BoolField("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantityIsSet", m.MasterFirm_MaximumOrderQuantityIsSet())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantity", m.MasterFirm_MaximumOrderQuantity())
	w.BoolField("m_ExchangeTraderIdIsSet", m.ExchangeTraderIdIsSet())
	w.StringField("m_ExchangeTraderId", m.ExchangeTraderId())
	return w.Finish(), nil
}

func (m TradeAccountDataUpdatePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10117)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsNewAccount", m.IsNewAccount())
	w.StringField("NewAccountAuthorizedUsername", m.NewAccountAuthorizedUsername())
	w.StringField("TradeAccount", m.TradeAccount())
	w.BoolField("CurrencyCodeIsSet", m.CurrencyCodeIsSet())
	w.StringField("CurrencyCode", m.CurrencyCode())
	w.BoolField("DailyNetLossLimitInAccountCurrencyIsSet", m.DailyNetLossLimitInAccountCurrencyIsSet())
	w.Float32Field("DailyNetLossLimitInAccountCurrency", m.DailyNetLossLimitInAccountCurrency())
	w.BoolField("PercentOfCashBalanceForDailyNetLossLimitIsSet", m.PercentOfCashBalanceForDailyNetLossLimitIsSet())
	w.Int32Field("PercentOfCashBalanceForDailyNetLossLimit", m.PercentOfCashBalanceForDailyNetLossLimit())
	w.BoolField("UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet", m.UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet())
	w.Uint8Field("UseTrailingAccountValueToNotAllowIncreaseInPositions", m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.BoolField("DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet", m.DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.BoolField("FlattenPositionsAtDailyLossLimitIsSet", m.FlattenPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("FlattenPositionsAtDailyLossLimit", m.FlattenPositionsAtDailyLossLimit())
	w.BoolField("ClosePositionsAtEndOfDayIsSet", m.ClosePositionsAtEndOfDayIsSet())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.BoolField("FlattenPositionsWhenUnderMarginIntradayIsSet", m.FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Uint8Field("FlattenPositionsWhenUnderMarginIntraday", m.FlattenPositionsWhenUnderMarginIntraday())
	w.BoolField("FlattenPositionsWhenUnderMarginAtEndOfDayIsSet", m.FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Uint8Field("FlattenPositionsWhenUnderMarginAtEndOfDay", m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.BoolField("SenderSubIDIsSet", m.SenderSubIDIsSet())
	w.StringField("SenderSubID", m.SenderSubID())
	w.BoolField("SenderLocationIdIsSet", m.SenderLocationIdIsSet())
	w.StringField("SenderLocationId", m.SenderLocationId())
	w.BoolField("SelfMatchPreventionIDIsSet", m.SelfMatchPreventionIDIsSet())
	w.StringField("SelfMatchPreventionID", m.SelfMatchPreventionID())
	w.BoolField("CTICodeIsSet", m.CTICodeIsSet())
	w.Int32Field("CTICode", m.CTICode())
	w.BoolField("TradeAccountIsReadOnlyIsSet", m.TradeAccountIsReadOnlyIsSet())
	w.BoolField("TradeAccountIsReadOnly", m.TradeAccountIsReadOnly())
	w.BoolField("MaximumGlobalPositionQuantityIsSet", m.MaximumGlobalPositionQuantityIsSet())
	w.Int32Field("MaximumGlobalPositionQuantity", m.MaximumGlobalPositionQuantity())
	w.BoolField("TradeFeePerContractIsSet", m.TradeFeePerContractIsSet())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	w.BoolField("TradeFeePerShareIsSet", m.TradeFeePerShareIsSet())
	w.Float64Field("TradeFeePerShare", m.TradeFeePerShare())
	w.BoolField("RequireSufficientMarginForNewPositionsIsSet", m.RequireSufficientMarginForNewPositionsIsSet())
	w.Uint8Field("RequireSufficientMarginForNewPositions", m.RequireSufficientMarginForNewPositions())
	w.BoolField("UsePercentOfMarginIsSet", m.UsePercentOfMarginIsSet())
	w.Int32Field("UsePercentOfMargin", m.UsePercentOfMargin())
	w.BoolField("UsePercentOfMarginForDayTradingIsSet", m.UsePercentOfMarginForDayTradingIsSet())
	w.Int32Field("UsePercentOfMarginForDayTrading", m.UsePercentOfMarginForDayTrading())
	w.BoolField("MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet", m.MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Int32Field("MaximumAllowedAccountBalanceForPositionsAsPercentage", m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.BoolField("FirmIDIsSet", m.FirmIDIsSet())
	w.StringField("FirmID", m.FirmID())
	w.BoolField("TradingIsDisabledIsSet", m.TradingIsDisabledIsSet())
	w.BoolField("TradingIsDisabled", m.TradingIsDisabled())
	w.BoolField("DescriptiveNameIsSet", m.DescriptiveNameIsSet())
	w.StringField("DescriptiveName", m.DescriptiveName())
	w.BoolField("IsMasterFirmControlAccountIsSet", m.IsMasterFirmControlAccountIsSet())
	w.BoolField("IsMasterFirmControlAccount", m.IsMasterFirmControlAccount())
	w.BoolField("MinimumRequiredAccountValueIsSet", m.MinimumRequiredAccountValueIsSet())
	w.Float64Field("MinimumRequiredAccountValue", m.MinimumRequiredAccountValue())
	w.BoolField("BeginTimeForDayMarginIsSet", m.BeginTimeForDayMarginIsSet())
	w.Int64Field("BeginTimeForDayMargin", m.BeginTimeForDayMargin())
	w.BoolField("EndTimeForDayMarginIsSet", m.EndTimeForDayMarginIsSet())
	w.Int64Field("EndTimeForDayMargin", m.EndTimeForDayMargin())
	w.BoolField("DayMarginTimeZoneIsSet", m.DayMarginTimeZoneIsSet())
	w.StringField("DayMarginTimeZone", m.DayMarginTimeZone())
	w.BoolField("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.BoolField("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.BoolField("m_UseMasterFirm_SymbolLimitsArrayIsSet", m.UseMasterFirm_SymbolLimitsArrayIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolLimitsArray", m.UseMasterFirm_SymbolLimitsArray())
	w.BoolField("m_UseMasterFirm_TradeFeesIsSet", m.UseMasterFirm_TradeFeesIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFees", m.UseMasterFirm_TradeFees())
	w.BoolField("m_UseMasterFirm_TradeFeePerShareIsSet", m.UseMasterFirm_TradeFeePerShareIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFeePerShare", m.UseMasterFirm_TradeFeePerShare())
	w.BoolField("m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet", m.UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet())
	w.Uint8Field("m_UseMasterFirm_RequireSufficientMarginForNewPositions", m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginIsSet", m.UseMasterFirm_UsePercentOfMarginIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMargin", m.UseMasterFirm_UsePercentOfMargin())
	w.BoolField("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet())
	w.Uint8Field("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.BoolField("m_UseMasterFirm_MinimumRequiredAccountValueIsSet", m.UseMasterFirm_MinimumRequiredAccountValueIsSet())
	w.Uint8Field("m_UseMasterFirm_MinimumRequiredAccountValue", m.UseMasterFirm_MinimumRequiredAccountValue())
	w.BoolField("m_UseMasterFirm_MarginTimeSettingsIsSet", m.UseMasterFirm_MarginTimeSettingsIsSet())
	w.Uint8Field("m_UseMasterFirm_MarginTimeSettings", m.UseMasterFirm_MarginTimeSettings())
	w.BoolField("m_UseMasterFirm_TradingIsDisabledIsSet", m.UseMasterFirm_TradingIsDisabledIsSet())
	w.BoolField("m_UseMasterFirm_TradingIsDisabled", m.UseMasterFirm_TradingIsDisabled())
	w.BoolField("IsTradeStatisticsPublicallySharedIsSet", m.IsTradeStatisticsPublicallySharedIsSet())
	w.BoolField("IsTradeStatisticsPublicallyShared", m.IsTradeStatisticsPublicallyShared())
	w.BoolField("IsReadOnlyFollowingRequestsAllowedIsSet", m.IsReadOnlyFollowingRequestsAllowedIsSet())
	w.BoolField("IsReadOnlyFollowingRequestsAllowed", m.IsReadOnlyFollowingRequestsAllowed())
	w.BoolField("IsTradeAccountSharingAllowedIsSet", m.IsTradeAccountSharingAllowedIsSet())
	w.BoolField("IsTradeAccountSharingAllowed", m.IsTradeAccountSharingAllowed())
	w.BoolField("ReadOnlyShareToAllSCUsernamesIsSet", m.ReadOnlyShareToAllSCUsernamesIsSet())
	w.Uint8Field("ReadOnlyShareToAllSCUsernames", m.ReadOnlyShareToAllSCUsernames())
	w.BoolField("m_UseMasterFirm_SymbolCommissionsArrayIsSet", m.UseMasterFirm_SymbolCommissionsArrayIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArray", m.UseMasterFirm_SymbolCommissionsArray())
	w.BoolField("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet())
	w.Uint8Field("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet", m.UseMasterFirm_UsePercentOfMarginForDayTradingIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTrading", m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.BoolField("m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet", m.UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArrayFullOverride", m.UseMasterFirm_SymbolCommissionsArrayFullOverride())
	w.BoolField("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet())
	w.Uint8Field("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet", m.UseMasterFirm_UsePercentOfMarginFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginFullOverride", m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.BoolField("m_UseMasterFirm_TradeFeesFullOverrideIsSet", m.UseMasterFirm_TradeFeesFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_TradeFeesFullOverride", m.UseMasterFirm_TradeFeesFullOverride())
	w.BoolField("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.BoolField("m_LiquidationOnlyModeIsSet", m.LiquidationOnlyModeIsSet())
	w.Uint8Field("m_LiquidationOnlyMode", m.LiquidationOnlyMode())
	w.BoolField("m_CustomerOrFirmIsSet", m.CustomerOrFirmIsSet())
	w.Uint8Field("m_CustomerOrFirm", m.CustomerOrFirm())
	w.BoolField("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.BoolField("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.BoolField("m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet", m.MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginIntraday", m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.BoolField("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantityIsSet", m.MasterFirm_MaximumOrderQuantityIsSet())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantity", m.MasterFirm_MaximumOrderQuantity())
	w.BoolField("m_ExchangeTraderIdIsSet", m.ExchangeTraderIdIsSet())
	w.StringField("m_ExchangeTraderId", m.ExchangeTraderId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUpdateBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10117 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsNewAccount(r.Bool())
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
	m.SetCurrencyCodeIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrencyCode(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimitInAccountCurrencyIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimitInAccountCurrency(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetPercentOfCashBalanceForDailyNetLossLimit(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsAtDailyLossLimitIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsAtDailyLossLimit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosePositionsAtEndOfDayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosePositionsAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginIntradayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginIntraday(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderSubIDIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderSubID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderLocationIdIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderLocationId(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSelfMatchPreventionIDIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetSelfMatchPreventionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCTICodeIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetCTICode(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccountIsReadOnlyIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccountIsReadOnly(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumGlobalPositionQuantityIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumGlobalPositionQuantity(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerContractIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerContract(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerShareIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerShare(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequireSufficientMarginForNewPositionsIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequireSufficientMarginForNewPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMarginIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMargin(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMarginForDayTradingIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMarginForDayTrading(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetFirmIDIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetFirmID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsDisabledIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsDisabled(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescriptiveNameIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescriptiveName(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsMasterFirmControlAccountIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsMasterFirmControlAccount(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMinimumRequiredAccountValueIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMinimumRequiredAccountValue(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBeginTimeForDayMarginIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetBeginTimeForDayMargin(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetEndTimeForDayMarginIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetEndTimeForDayMargin(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDayMarginTimeZoneIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDayMarginTimeZone(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolLimitsArrayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolLimitsArray(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeesIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFees(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeePerShareIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeePerShare(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_RequireSufficientMarginForNewPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMargin(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MinimumRequiredAccountValue(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MarginTimeSettingsIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MarginTimeSettings(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradingIsDisabledIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradingIsDisabled(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeStatisticsPublicallySharedIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeStatisticsPublicallyShared(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsReadOnlyFollowingRequestsAllowedIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsReadOnlyFollowingRequestsAllowed(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeAccountSharingAllowedIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeAccountSharingAllowed(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetReadOnlyShareToAllSCUsernamesIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetReadOnlyShareToAllSCUsernames(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArrayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArray(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTrading(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeesFullOverrideIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeesFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetLiquidationOnlyModeIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetLiquidationOnlyMode(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetCustomerOrFirmIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetCustomerOrFirm(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(r.Bool())
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
	m.SetExchangeTraderIdIsSet(r.Bool())
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
	m.SetIsNewAccount(r.Bool())
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
	m.SetCurrencyCodeIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrencyCode(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimitInAccountCurrencyIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimitInAccountCurrency(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetPercentOfCashBalanceForDailyNetLossLimit(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsAtDailyLossLimitIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsAtDailyLossLimit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosePositionsAtEndOfDayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosePositionsAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginIntradayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginIntraday(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderSubIDIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderSubID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderLocationIdIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderLocationId(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSelfMatchPreventionIDIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetSelfMatchPreventionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCTICodeIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetCTICode(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccountIsReadOnlyIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccountIsReadOnly(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumGlobalPositionQuantityIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumGlobalPositionQuantity(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerContractIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerContract(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerShareIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerShare(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequireSufficientMarginForNewPositionsIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequireSufficientMarginForNewPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMarginIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMargin(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMarginForDayTradingIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMarginForDayTrading(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetFirmIDIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetFirmID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsDisabledIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsDisabled(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescriptiveNameIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescriptiveName(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsMasterFirmControlAccountIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsMasterFirmControlAccount(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMinimumRequiredAccountValueIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMinimumRequiredAccountValue(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBeginTimeForDayMarginIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetBeginTimeForDayMargin(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetEndTimeForDayMarginIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetEndTimeForDayMargin(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDayMarginTimeZoneIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDayMarginTimeZone(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolLimitsArrayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolLimitsArray(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeesIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFees(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeePerShareIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeePerShare(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_RequireSufficientMarginForNewPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMargin(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MinimumRequiredAccountValue(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MarginTimeSettingsIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MarginTimeSettings(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradingIsDisabledIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradingIsDisabled(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeStatisticsPublicallySharedIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeStatisticsPublicallyShared(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsReadOnlyFollowingRequestsAllowedIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsReadOnlyFollowingRequestsAllowed(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeAccountSharingAllowedIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeAccountSharingAllowed(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetReadOnlyShareToAllSCUsernamesIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetReadOnlyShareToAllSCUsernames(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArrayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArray(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTrading(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeesFullOverrideIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeesFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetLiquidationOnlyModeIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetLiquidationOnlyMode(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetCustomerOrFirmIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetCustomerOrFirm(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(r.Bool())
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
	m.SetExchangeTraderIdIsSet(r.Bool())
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

func (m *TradeAccountDataUpdateBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
			m.SetIsNewAccount(r.Bool())
		case "NewAccountAuthorizedUsername":
			m.SetNewAccountAuthorizedUsername(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "CurrencyCodeIsSet":
			m.SetCurrencyCodeIsSet(r.Bool())
		case "CurrencyCode":
			m.SetCurrencyCode(r.String())
		case "DailyNetLossLimitInAccountCurrencyIsSet":
			m.SetDailyNetLossLimitInAccountCurrencyIsSet(r.Bool())
		case "DailyNetLossLimitInAccountCurrency":
			m.SetDailyNetLossLimitInAccountCurrency(r.Float32())
		case "PercentOfCashBalanceForDailyNetLossLimitIsSet":
			m.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(r.Bool())
		case "PercentOfCashBalanceForDailyNetLossLimit":
			m.SetPercentOfCashBalanceForDailyNetLossLimit(r.Int32())
		case "UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet":
			m.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(r.Bool())
		case "UseTrailingAccountValueToNotAllowIncreaseInPositions":
			m.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(r.Uint8())
		case "DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet":
			m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(r.Bool())
		case "DoNotAllowIncreaseInPositionsAtDailyLossLimit":
			m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
		case "FlattenPositionsAtDailyLossLimitIsSet":
			m.SetFlattenPositionsAtDailyLossLimitIsSet(r.Bool())
		case "FlattenPositionsAtDailyLossLimit":
			m.SetFlattenPositionsAtDailyLossLimit(r.Uint8())
		case "ClosePositionsAtEndOfDayIsSet":
			m.SetClosePositionsAtEndOfDayIsSet(r.Bool())
		case "ClosePositionsAtEndOfDay":
			m.SetClosePositionsAtEndOfDay(r.Uint8())
		case "FlattenPositionsWhenUnderMarginIntradayIsSet":
			m.SetFlattenPositionsWhenUnderMarginIntradayIsSet(r.Bool())
		case "FlattenPositionsWhenUnderMarginIntraday":
			m.SetFlattenPositionsWhenUnderMarginIntraday(r.Uint8())
		case "FlattenPositionsWhenUnderMarginAtEndOfDayIsSet":
			m.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(r.Bool())
		case "FlattenPositionsWhenUnderMarginAtEndOfDay":
			m.SetFlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
		case "SenderSubIDIsSet":
			m.SetSenderSubIDIsSet(r.Bool())
		case "SenderSubID":
			m.SetSenderSubID(r.String())
		case "SenderLocationIdIsSet":
			m.SetSenderLocationIdIsSet(r.Bool())
		case "SenderLocationId":
			m.SetSenderLocationId(r.String())
		case "SelfMatchPreventionIDIsSet":
			m.SetSelfMatchPreventionIDIsSet(r.Bool())
		case "SelfMatchPreventionID":
			m.SetSelfMatchPreventionID(r.String())
		case "CTICodeIsSet":
			m.SetCTICodeIsSet(r.Bool())
		case "CTICode":
			m.SetCTICode(r.Int32())
		case "TradeAccountIsReadOnlyIsSet":
			m.SetTradeAccountIsReadOnlyIsSet(r.Bool())
		case "TradeAccountIsReadOnly":
			m.SetTradeAccountIsReadOnly(r.Bool())
		case "MaximumGlobalPositionQuantityIsSet":
			m.SetMaximumGlobalPositionQuantityIsSet(r.Bool())
		case "MaximumGlobalPositionQuantity":
			m.SetMaximumGlobalPositionQuantity(r.Int32())
		case "TradeFeePerContractIsSet":
			m.SetTradeFeePerContractIsSet(r.Bool())
		case "TradeFeePerContract":
			m.SetTradeFeePerContract(r.Float64())
		case "TradeFeePerShareIsSet":
			m.SetTradeFeePerShareIsSet(r.Bool())
		case "TradeFeePerShare":
			m.SetTradeFeePerShare(r.Float64())
		case "RequireSufficientMarginForNewPositionsIsSet":
			m.SetRequireSufficientMarginForNewPositionsIsSet(r.Bool())
		case "RequireSufficientMarginForNewPositions":
			m.SetRequireSufficientMarginForNewPositions(r.Uint8())
		case "UsePercentOfMarginIsSet":
			m.SetUsePercentOfMarginIsSet(r.Bool())
		case "UsePercentOfMargin":
			m.SetUsePercentOfMargin(r.Int32())
		case "UsePercentOfMarginForDayTradingIsSet":
			m.SetUsePercentOfMarginForDayTradingIsSet(r.Bool())
		case "UsePercentOfMarginForDayTrading":
			m.SetUsePercentOfMarginForDayTrading(r.Int32())
		case "MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet":
			m.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(r.Bool())
		case "MaximumAllowedAccountBalanceForPositionsAsPercentage":
			m.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(r.Int32())
		case "FirmIDIsSet":
			m.SetFirmIDIsSet(r.Bool())
		case "FirmID":
			m.SetFirmID(r.String())
		case "TradingIsDisabledIsSet":
			m.SetTradingIsDisabledIsSet(r.Bool())
		case "TradingIsDisabled":
			m.SetTradingIsDisabled(r.Bool())
		case "DescriptiveNameIsSet":
			m.SetDescriptiveNameIsSet(r.Bool())
		case "DescriptiveName":
			m.SetDescriptiveName(r.String())
		case "IsMasterFirmControlAccountIsSet":
			m.SetIsMasterFirmControlAccountIsSet(r.Bool())
		case "IsMasterFirmControlAccount":
			m.SetIsMasterFirmControlAccount(r.Bool())
		case "MinimumRequiredAccountValueIsSet":
			m.SetMinimumRequiredAccountValueIsSet(r.Bool())
		case "MinimumRequiredAccountValue":
			m.SetMinimumRequiredAccountValue(r.Float64())
		case "BeginTimeForDayMarginIsSet":
			m.SetBeginTimeForDayMarginIsSet(r.Bool())
		case "BeginTimeForDayMargin":
			m.SetBeginTimeForDayMargin(r.Int64())
		case "EndTimeForDayMarginIsSet":
			m.SetEndTimeForDayMarginIsSet(r.Bool())
		case "EndTimeForDayMargin":
			m.SetEndTimeForDayMargin(r.Int64())
		case "DayMarginTimeZoneIsSet":
			m.SetDayMarginTimeZoneIsSet(r.Bool())
		case "DayMarginTimeZone":
			m.SetDayMarginTimeZone(r.String())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(r.Bool())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(r.Uint8())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(r.Bool())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
		case "m_UseMasterFirm_SymbolLimitsArrayIsSet":
			m.SetUseMasterFirm_SymbolLimitsArrayIsSet(r.Bool())
		case "m_UseMasterFirm_SymbolLimitsArray":
			m.SetUseMasterFirm_SymbolLimitsArray(r.Uint8())
		case "m_UseMasterFirm_TradeFeesIsSet":
			m.SetUseMasterFirm_TradeFeesIsSet(r.Bool())
		case "m_UseMasterFirm_TradeFees":
			m.SetUseMasterFirm_TradeFees(r.Uint8())
		case "m_UseMasterFirm_TradeFeePerShareIsSet":
			m.SetUseMasterFirm_TradeFeePerShareIsSet(r.Bool())
		case "m_UseMasterFirm_TradeFeePerShare":
			m.SetUseMasterFirm_TradeFeePerShare(r.Uint8())
		case "m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet":
			m.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(r.Bool())
		case "m_UseMasterFirm_RequireSufficientMarginForNewPositions":
			m.SetUseMasterFirm_RequireSufficientMarginForNewPositions(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginIsSet":
			m.SetUseMasterFirm_UsePercentOfMarginIsSet(r.Bool())
		case "m_UseMasterFirm_UsePercentOfMargin":
			m.SetUseMasterFirm_UsePercentOfMargin(r.Uint8())
		case "m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet":
			m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(r.Bool())
		case "m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage":
			m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(r.Uint8())
		case "m_UseMasterFirm_MinimumRequiredAccountValueIsSet":
			m.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(r.Bool())
		case "m_UseMasterFirm_MinimumRequiredAccountValue":
			m.SetUseMasterFirm_MinimumRequiredAccountValue(r.Uint8())
		case "m_UseMasterFirm_MarginTimeSettingsIsSet":
			m.SetUseMasterFirm_MarginTimeSettingsIsSet(r.Bool())
		case "m_UseMasterFirm_MarginTimeSettings":
			m.SetUseMasterFirm_MarginTimeSettings(r.Uint8())
		case "m_UseMasterFirm_TradingIsDisabledIsSet":
			m.SetUseMasterFirm_TradingIsDisabledIsSet(r.Bool())
		case "m_UseMasterFirm_TradingIsDisabled":
			m.SetUseMasterFirm_TradingIsDisabled(r.Bool())
		case "IsTradeStatisticsPublicallySharedIsSet":
			m.SetIsTradeStatisticsPublicallySharedIsSet(r.Bool())
		case "IsTradeStatisticsPublicallyShared":
			m.SetIsTradeStatisticsPublicallyShared(r.Bool())
		case "IsReadOnlyFollowingRequestsAllowedIsSet":
			m.SetIsReadOnlyFollowingRequestsAllowedIsSet(r.Bool())
		case "IsReadOnlyFollowingRequestsAllowed":
			m.SetIsReadOnlyFollowingRequestsAllowed(r.Bool())
		case "IsTradeAccountSharingAllowedIsSet":
			m.SetIsTradeAccountSharingAllowedIsSet(r.Bool())
		case "IsTradeAccountSharingAllowed":
			m.SetIsTradeAccountSharingAllowed(r.Bool())
		case "ReadOnlyShareToAllSCUsernamesIsSet":
			m.SetReadOnlyShareToAllSCUsernamesIsSet(r.Bool())
		case "ReadOnlyShareToAllSCUsernames":
			m.SetReadOnlyShareToAllSCUsernames(r.Uint8())
		case "m_UseMasterFirm_SymbolCommissionsArrayIsSet":
			m.SetUseMasterFirm_SymbolCommissionsArrayIsSet(r.Bool())
		case "m_UseMasterFirm_SymbolCommissionsArray":
			m.SetUseMasterFirm_SymbolCommissionsArray(r.Uint8())
		case "m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet":
			m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(r.Bool())
		case "m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit":
			m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(r.Bool())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTrading":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTrading(r.Uint8())
		case "m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet":
			m.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(r.Bool())
		case "m_UseMasterFirm_SymbolCommissionsArrayFullOverride":
			m.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(r.Uint8())
		case "m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet":
			m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(r.Bool())
		case "m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders":
			m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet":
			m.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(r.Bool())
		case "m_UseMasterFirm_UsePercentOfMarginFullOverride":
			m.SetUseMasterFirm_UsePercentOfMarginFullOverride(r.Uint8())
		case "m_UseMasterFirm_TradeFeesFullOverrideIsSet":
			m.SetUseMasterFirm_TradeFeesFullOverrideIsSet(r.Bool())
		case "m_UseMasterFirm_TradeFeesFullOverride":
			m.SetUseMasterFirm_TradeFeesFullOverride(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(r.Bool())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(r.Uint8())
		case "m_LiquidationOnlyModeIsSet":
			m.SetLiquidationOnlyModeIsSet(r.Bool())
		case "m_LiquidationOnlyMode":
			m.SetLiquidationOnlyMode(r.Uint8())
		case "m_CustomerOrFirmIsSet":
			m.SetCustomerOrFirmIsSet(r.Bool())
		case "m_CustomerOrFirm":
			m.SetCustomerOrFirm(r.Uint8())
		case "m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet":
			m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(r.Bool())
		case "m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet":
			m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet":
			m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(r.Bool())
		case "m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue":
			m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(r.Bool())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginIntraday":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(r.Bool())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(r.Uint8())
		case "m_MasterFirm_MaximumOrderQuantityIsSet":
			m.SetMasterFirm_MaximumOrderQuantityIsSet(r.Uint32())
		case "m_MasterFirm_MaximumOrderQuantity":
			m.SetMasterFirm_MaximumOrderQuantity(r.Uint32())
		case "m_ExchangeTraderIdIsSet":
			m.SetExchangeTraderIdIsSet(r.Bool())
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
			m.SetIsNewAccount(r.Bool())
		case "NewAccountAuthorizedUsername":
			m.SetNewAccountAuthorizedUsername(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "CurrencyCodeIsSet":
			m.SetCurrencyCodeIsSet(r.Bool())
		case "CurrencyCode":
			m.SetCurrencyCode(r.String())
		case "DailyNetLossLimitInAccountCurrencyIsSet":
			m.SetDailyNetLossLimitInAccountCurrencyIsSet(r.Bool())
		case "DailyNetLossLimitInAccountCurrency":
			m.SetDailyNetLossLimitInAccountCurrency(r.Float32())
		case "PercentOfCashBalanceForDailyNetLossLimitIsSet":
			m.SetPercentOfCashBalanceForDailyNetLossLimitIsSet(r.Bool())
		case "PercentOfCashBalanceForDailyNetLossLimit":
			m.SetPercentOfCashBalanceForDailyNetLossLimit(r.Int32())
		case "UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet":
			m.SetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet(r.Bool())
		case "UseTrailingAccountValueToNotAllowIncreaseInPositions":
			m.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(r.Uint8())
		case "DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet":
			m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(r.Bool())
		case "DoNotAllowIncreaseInPositionsAtDailyLossLimit":
			m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
		case "FlattenPositionsAtDailyLossLimitIsSet":
			m.SetFlattenPositionsAtDailyLossLimitIsSet(r.Bool())
		case "FlattenPositionsAtDailyLossLimit":
			m.SetFlattenPositionsAtDailyLossLimit(r.Uint8())
		case "ClosePositionsAtEndOfDayIsSet":
			m.SetClosePositionsAtEndOfDayIsSet(r.Bool())
		case "ClosePositionsAtEndOfDay":
			m.SetClosePositionsAtEndOfDay(r.Uint8())
		case "FlattenPositionsWhenUnderMarginIntradayIsSet":
			m.SetFlattenPositionsWhenUnderMarginIntradayIsSet(r.Bool())
		case "FlattenPositionsWhenUnderMarginIntraday":
			m.SetFlattenPositionsWhenUnderMarginIntraday(r.Uint8())
		case "FlattenPositionsWhenUnderMarginAtEndOfDayIsSet":
			m.SetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet(r.Bool())
		case "FlattenPositionsWhenUnderMarginAtEndOfDay":
			m.SetFlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
		case "SenderSubIDIsSet":
			m.SetSenderSubIDIsSet(r.Bool())
		case "SenderSubID":
			m.SetSenderSubID(r.String())
		case "SenderLocationIdIsSet":
			m.SetSenderLocationIdIsSet(r.Bool())
		case "SenderLocationId":
			m.SetSenderLocationId(r.String())
		case "SelfMatchPreventionIDIsSet":
			m.SetSelfMatchPreventionIDIsSet(r.Bool())
		case "SelfMatchPreventionID":
			m.SetSelfMatchPreventionID(r.String())
		case "CTICodeIsSet":
			m.SetCTICodeIsSet(r.Bool())
		case "CTICode":
			m.SetCTICode(r.Int32())
		case "TradeAccountIsReadOnlyIsSet":
			m.SetTradeAccountIsReadOnlyIsSet(r.Bool())
		case "TradeAccountIsReadOnly":
			m.SetTradeAccountIsReadOnly(r.Bool())
		case "MaximumGlobalPositionQuantityIsSet":
			m.SetMaximumGlobalPositionQuantityIsSet(r.Bool())
		case "MaximumGlobalPositionQuantity":
			m.SetMaximumGlobalPositionQuantity(r.Int32())
		case "TradeFeePerContractIsSet":
			m.SetTradeFeePerContractIsSet(r.Bool())
		case "TradeFeePerContract":
			m.SetTradeFeePerContract(r.Float64())
		case "TradeFeePerShareIsSet":
			m.SetTradeFeePerShareIsSet(r.Bool())
		case "TradeFeePerShare":
			m.SetTradeFeePerShare(r.Float64())
		case "RequireSufficientMarginForNewPositionsIsSet":
			m.SetRequireSufficientMarginForNewPositionsIsSet(r.Bool())
		case "RequireSufficientMarginForNewPositions":
			m.SetRequireSufficientMarginForNewPositions(r.Uint8())
		case "UsePercentOfMarginIsSet":
			m.SetUsePercentOfMarginIsSet(r.Bool())
		case "UsePercentOfMargin":
			m.SetUsePercentOfMargin(r.Int32())
		case "UsePercentOfMarginForDayTradingIsSet":
			m.SetUsePercentOfMarginForDayTradingIsSet(r.Bool())
		case "UsePercentOfMarginForDayTrading":
			m.SetUsePercentOfMarginForDayTrading(r.Int32())
		case "MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet":
			m.SetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(r.Bool())
		case "MaximumAllowedAccountBalanceForPositionsAsPercentage":
			m.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(r.Int32())
		case "FirmIDIsSet":
			m.SetFirmIDIsSet(r.Bool())
		case "FirmID":
			m.SetFirmID(r.String())
		case "TradingIsDisabledIsSet":
			m.SetTradingIsDisabledIsSet(r.Bool())
		case "TradingIsDisabled":
			m.SetTradingIsDisabled(r.Bool())
		case "DescriptiveNameIsSet":
			m.SetDescriptiveNameIsSet(r.Bool())
		case "DescriptiveName":
			m.SetDescriptiveName(r.String())
		case "IsMasterFirmControlAccountIsSet":
			m.SetIsMasterFirmControlAccountIsSet(r.Bool())
		case "IsMasterFirmControlAccount":
			m.SetIsMasterFirmControlAccount(r.Bool())
		case "MinimumRequiredAccountValueIsSet":
			m.SetMinimumRequiredAccountValueIsSet(r.Bool())
		case "MinimumRequiredAccountValue":
			m.SetMinimumRequiredAccountValue(r.Float64())
		case "BeginTimeForDayMarginIsSet":
			m.SetBeginTimeForDayMarginIsSet(r.Bool())
		case "BeginTimeForDayMargin":
			m.SetBeginTimeForDayMargin(r.Int64())
		case "EndTimeForDayMarginIsSet":
			m.SetEndTimeForDayMarginIsSet(r.Bool())
		case "EndTimeForDayMargin":
			m.SetEndTimeForDayMargin(r.Int64())
		case "DayMarginTimeZoneIsSet":
			m.SetDayMarginTimeZoneIsSet(r.Bool())
		case "DayMarginTimeZone":
			m.SetDayMarginTimeZone(r.String())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet(r.Bool())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(r.Uint8())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet(r.Bool())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
		case "m_UseMasterFirm_SymbolLimitsArrayIsSet":
			m.SetUseMasterFirm_SymbolLimitsArrayIsSet(r.Bool())
		case "m_UseMasterFirm_SymbolLimitsArray":
			m.SetUseMasterFirm_SymbolLimitsArray(r.Uint8())
		case "m_UseMasterFirm_TradeFeesIsSet":
			m.SetUseMasterFirm_TradeFeesIsSet(r.Bool())
		case "m_UseMasterFirm_TradeFees":
			m.SetUseMasterFirm_TradeFees(r.Uint8())
		case "m_UseMasterFirm_TradeFeePerShareIsSet":
			m.SetUseMasterFirm_TradeFeePerShareIsSet(r.Bool())
		case "m_UseMasterFirm_TradeFeePerShare":
			m.SetUseMasterFirm_TradeFeePerShare(r.Uint8())
		case "m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet":
			m.SetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet(r.Bool())
		case "m_UseMasterFirm_RequireSufficientMarginForNewPositions":
			m.SetUseMasterFirm_RequireSufficientMarginForNewPositions(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginIsSet":
			m.SetUseMasterFirm_UsePercentOfMarginIsSet(r.Bool())
		case "m_UseMasterFirm_UsePercentOfMargin":
			m.SetUseMasterFirm_UsePercentOfMargin(r.Uint8())
		case "m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet":
			m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet(r.Bool())
		case "m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage":
			m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(r.Uint8())
		case "m_UseMasterFirm_MinimumRequiredAccountValueIsSet":
			m.SetUseMasterFirm_MinimumRequiredAccountValueIsSet(r.Bool())
		case "m_UseMasterFirm_MinimumRequiredAccountValue":
			m.SetUseMasterFirm_MinimumRequiredAccountValue(r.Uint8())
		case "m_UseMasterFirm_MarginTimeSettingsIsSet":
			m.SetUseMasterFirm_MarginTimeSettingsIsSet(r.Bool())
		case "m_UseMasterFirm_MarginTimeSettings":
			m.SetUseMasterFirm_MarginTimeSettings(r.Uint8())
		case "m_UseMasterFirm_TradingIsDisabledIsSet":
			m.SetUseMasterFirm_TradingIsDisabledIsSet(r.Bool())
		case "m_UseMasterFirm_TradingIsDisabled":
			m.SetUseMasterFirm_TradingIsDisabled(r.Bool())
		case "IsTradeStatisticsPublicallySharedIsSet":
			m.SetIsTradeStatisticsPublicallySharedIsSet(r.Bool())
		case "IsTradeStatisticsPublicallyShared":
			m.SetIsTradeStatisticsPublicallyShared(r.Bool())
		case "IsReadOnlyFollowingRequestsAllowedIsSet":
			m.SetIsReadOnlyFollowingRequestsAllowedIsSet(r.Bool())
		case "IsReadOnlyFollowingRequestsAllowed":
			m.SetIsReadOnlyFollowingRequestsAllowed(r.Bool())
		case "IsTradeAccountSharingAllowedIsSet":
			m.SetIsTradeAccountSharingAllowedIsSet(r.Bool())
		case "IsTradeAccountSharingAllowed":
			m.SetIsTradeAccountSharingAllowed(r.Bool())
		case "ReadOnlyShareToAllSCUsernamesIsSet":
			m.SetReadOnlyShareToAllSCUsernamesIsSet(r.Bool())
		case "ReadOnlyShareToAllSCUsernames":
			m.SetReadOnlyShareToAllSCUsernames(r.Uint8())
		case "m_UseMasterFirm_SymbolCommissionsArrayIsSet":
			m.SetUseMasterFirm_SymbolCommissionsArrayIsSet(r.Bool())
		case "m_UseMasterFirm_SymbolCommissionsArray":
			m.SetUseMasterFirm_SymbolCommissionsArray(r.Uint8())
		case "m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet":
			m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet(r.Bool())
		case "m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit":
			m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet(r.Bool())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTrading":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTrading(r.Uint8())
		case "m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet":
			m.SetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet(r.Bool())
		case "m_UseMasterFirm_SymbolCommissionsArrayFullOverride":
			m.SetUseMasterFirm_SymbolCommissionsArrayFullOverride(r.Uint8())
		case "m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet":
			m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet(r.Bool())
		case "m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders":
			m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet":
			m.SetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet(r.Bool())
		case "m_UseMasterFirm_UsePercentOfMarginFullOverride":
			m.SetUseMasterFirm_UsePercentOfMarginFullOverride(r.Uint8())
		case "m_UseMasterFirm_TradeFeesFullOverrideIsSet":
			m.SetUseMasterFirm_TradeFeesFullOverrideIsSet(r.Bool())
		case "m_UseMasterFirm_TradeFeesFullOverride":
			m.SetUseMasterFirm_TradeFeesFullOverride(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet(r.Bool())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(r.Uint8())
		case "m_LiquidationOnlyModeIsSet":
			m.SetLiquidationOnlyModeIsSet(r.Bool())
		case "m_LiquidationOnlyMode":
			m.SetLiquidationOnlyMode(r.Uint8())
		case "m_CustomerOrFirmIsSet":
			m.SetCustomerOrFirmIsSet(r.Bool())
		case "m_CustomerOrFirm":
			m.SetCustomerOrFirm(r.Uint8())
		case "m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet":
			m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet(r.Bool())
		case "m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet":
			m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet":
			m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet(r.Bool())
		case "m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue":
			m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet(r.Bool())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginIntraday":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet(r.Bool())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(r.Uint8())
		case "m_MasterFirm_MaximumOrderQuantityIsSet":
			m.SetMasterFirm_MaximumOrderQuantityIsSet(r.Uint32())
		case "m_MasterFirm_MaximumOrderQuantity":
			m.SetMasterFirm_MaximumOrderQuantity(r.Uint32())
		case "m_ExchangeTraderIdIsSet":
			m.SetExchangeTraderIdIsSet(r.Bool())
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
