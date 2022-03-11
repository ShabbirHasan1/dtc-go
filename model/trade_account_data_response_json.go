//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10116,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeAccountNotExist())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSimulated())
	w.Data = append(w.Data, ',')
	w.String(m.CurrencyCode())
	w.Data = append(w.Data, ',')
	w.Float64(m.CurrentCashBalance())
	w.Data = append(w.Data, ',')
	w.Float64(m.AvailableFundsForNewPositions())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Float64(m.AccountValue())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPositionsProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyProfitLoss())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TransactionIdentifierForCashBalanceAdjustment())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastTransactionDateTime())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingAccountValueAtWhichToNotAllowNewPositions())
	w.Data = append(w.Data, ',')
	w.Float64(m.CalculatedDailyNetLossLimitInAccountCurrency())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DailyNetLossLimitHasBeenReached())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastResetDailyNetLossManagementVariablesDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsUnderRequiredMargin())
	w.Data = append(w.Data, ',')
	w.Float32(m.DailyNetLossLimitInAccountCurrency())
	w.Data = append(w.Data, ',')
	w.Int32(m.PercentOfCashBalanceForDailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.String(m.SenderSubID())
	w.Data = append(w.Data, ',')
	w.String(m.SenderLocationId())
	w.Data = append(w.Data, ',')
	w.String(m.SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.Int32(m.CTICode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeAccountIsReadOnly())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumGlobalPositionQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.String(m.FirmID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.String(m.DescriptiveName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsMasterFirmControlAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Int64(m.BeginTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Int64(m.EndTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.String(m.DayMarginTimeZone())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsDeleted())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolLimitsArray())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFees())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MarginTimeSettings())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsTradeStatisticsPublicallyShared())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsReadOnlyFollowingRequestsAllowed())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsTradeAccountSharingAllowed())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReadOnlyShareToAllSCUsernames())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArray())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPositionsProfitLossBasedOnSettlement())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFull())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFullPositionsOnly())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeesFullOverride())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.Data = append(w.Data, ',')
	w.Float64(m.PeakMarginRequirement())
	w.Data = append(w.Data, ',')
	w.Uint8(m.LiquidationOnlyMode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginIntradayTriggered())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMinimumAccountValueTriggered())
	w.Data = append(w.Data, ',')
	w.Float64(m.AccountValueAtEndOfDayCaptureTime())
	w.Data = append(w.Data, ',')
	w.Int64(m.EndOfDayCaptureTime())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CustomerOrFirm())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantity())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastTriggerDateTimeUTCForDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OpenPositionsProfitLossIsDelayed())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeTraderId())
	return w.Finish(), nil
}

func (m TradeAccountDataResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10116,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeAccountNotExist())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSimulated())
	w.Data = append(w.Data, ',')
	w.String(m.CurrencyCode())
	w.Data = append(w.Data, ',')
	w.Float64(m.CurrentCashBalance())
	w.Data = append(w.Data, ',')
	w.Float64(m.AvailableFundsForNewPositions())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Float64(m.AccountValue())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPositionsProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyProfitLoss())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TransactionIdentifierForCashBalanceAdjustment())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastTransactionDateTime())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingAccountValueAtWhichToNotAllowNewPositions())
	w.Data = append(w.Data, ',')
	w.Float64(m.CalculatedDailyNetLossLimitInAccountCurrency())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DailyNetLossLimitHasBeenReached())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastResetDailyNetLossManagementVariablesDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsUnderRequiredMargin())
	w.Data = append(w.Data, ',')
	w.Float32(m.DailyNetLossLimitInAccountCurrency())
	w.Data = append(w.Data, ',')
	w.Int32(m.PercentOfCashBalanceForDailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.String(m.SenderSubID())
	w.Data = append(w.Data, ',')
	w.String(m.SenderLocationId())
	w.Data = append(w.Data, ',')
	w.String(m.SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.Int32(m.CTICode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeAccountIsReadOnly())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumGlobalPositionQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Int32(m.UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.String(m.FirmID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.String(m.DescriptiveName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsMasterFirmControlAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Int64(m.BeginTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.Int64(m.EndTimeForDayMargin())
	w.Data = append(w.Data, ',')
	w.String(m.DayMarginTimeZone())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsDeleted())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolLimitsArray())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFees())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeePerShare())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MinimumRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_MarginTimeSettings())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsTradeStatisticsPublicallyShared())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsReadOnlyFollowingRequestsAllowed())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsTradeAccountSharingAllowed())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReadOnlyShareToAllSCUsernames())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_SymbolCommissionsArray())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPositionsProfitLossBasedOnSettlement())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFull())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFullPositionsOnly())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_TradeFeesFullOverride())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.Data = append(w.Data, ',')
	w.Float64(m.PeakMarginRequirement())
	w.Data = append(w.Data, ',')
	w.Uint8(m.LiquidationOnlyMode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMarginIntradayTriggered())
	w.Data = append(w.Data, ',')
	w.Uint8(m.FlattenPositionsWhenUnderMinimumAccountValueTriggered())
	w.Data = append(w.Data, ',')
	w.Float64(m.AccountValueAtEndOfDayCaptureTime())
	w.Data = append(w.Data, ',')
	w.Int64(m.EndOfDayCaptureTime())
	w.Data = append(w.Data, ',')
	w.Uint8(m.CustomerOrFirm())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MasterFirm_MaximumOrderQuantity())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastTriggerDateTimeUTCForDailyLossLimit())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OpenPositionsProfitLossIsDelayed())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeTraderId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10116)
	w.Uint32Field("m_RequestID", m.RequestID())
	w.Uint8Field("m_TradeAccountNotExist", m.TradeAccountNotExist())
	w.StringField("m_TradeAccount", m.TradeAccount())
	w.Uint8Field("m_IsSimulated", m.IsSimulated())
	w.StringField("m_CurrencyCode", m.CurrencyCode())
	w.Float64Field("m_CurrentCashBalance", m.CurrentCashBalance())
	w.Float64Field("m_AvailableFundsForNewPositions", m.AvailableFundsForNewPositions())
	w.Float64Field("m_MarginRequirement", m.MarginRequirement())
	w.Float64Field("m_AccountValue", m.AccountValue())
	w.Float64Field("m_OpenPositionsProfitLoss", m.OpenPositionsProfitLoss())
	w.Float64Field("m_DailyProfitLoss", m.DailyProfitLoss())
	w.Uint64Field("m_TransactionIdentifierForCashBalanceAdjustment", m.TransactionIdentifierForCashBalanceAdjustment())
	w.Int64Field("m_LastTransactionDateTime", m.LastTransactionDateTime())
	w.Float64Field("m_TrailingAccountValueAtWhichToNotAllowNewPositions", m.TrailingAccountValueAtWhichToNotAllowNewPositions())
	w.Float64Field("m_CalculatedDailyNetLossLimitInAccountCurrency", m.CalculatedDailyNetLossLimitInAccountCurrency())
	w.Uint8Field("m_DailyNetLossLimitHasBeenReached", m.DailyNetLossLimitHasBeenReached())
	w.Int64Field("m_LastResetDailyNetLossManagementVariablesDateTimeUTC", m.LastResetDailyNetLossManagementVariablesDateTimeUTC())
	w.Uint8Field("m_IsUnderRequiredMargin", m.IsUnderRequiredMargin())
	w.Float32Field("m_DailyNetLossLimitInAccountCurrency", m.DailyNetLossLimitInAccountCurrency())
	w.Int32Field("m_PercentOfCashBalanceForDailyNetLossLimit", m.PercentOfCashBalanceForDailyNetLossLimit())
	w.Uint8Field("m_UseTrailingAccountValueToNotAllowIncreaseInPositions", m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.Uint8Field("m_DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Uint8Field("m_FlattenPositionsAtDailyLossLimit", m.FlattenPositionsAtDailyLossLimit())
	w.Uint8Field("m_ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.Uint8Field("m_FlattenPositionsWhenUnderMarginIntraday", m.FlattenPositionsWhenUnderMarginIntraday())
	w.Uint8Field("m_FlattenPositionsWhenUnderMarginAtEndOfDay", m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.StringField("m_SenderSubID", m.SenderSubID())
	w.StringField("m_SenderLocationId", m.SenderLocationId())
	w.StringField("m_SelfMatchPreventionID", m.SelfMatchPreventionID())
	w.Int32Field("m_CTICode", m.CTICode())
	w.Uint8Field("m_TradeAccountIsReadOnly", m.TradeAccountIsReadOnly())
	w.Int32Field("m_MaximumGlobalPositionQuantity", m.MaximumGlobalPositionQuantity())
	w.Float64Field("m_TradeFeePerContract", m.TradeFeePerContract())
	w.Float64Field("m_TradeFeePerShare", m.TradeFeePerShare())
	w.Uint8Field("m_RequireSufficientMarginForNewPositions", m.RequireSufficientMarginForNewPositions())
	w.Int32Field("m_UsePercentOfMargin", m.UsePercentOfMargin())
	w.Int32Field("m_UsePercentOfMarginForDayTrading", m.UsePercentOfMarginForDayTrading())
	w.Int32Field("m_MaximumAllowedAccountBalanceForPositionsAsPercentage", m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.StringField("m_FirmID", m.FirmID())
	w.Uint8Field("m_TradingIsDisabled", m.TradingIsDisabled())
	w.StringField("m_DescriptiveName", m.DescriptiveName())
	w.Uint8Field("m_IsMasterFirmControlAccount", m.IsMasterFirmControlAccount())
	w.Float64Field("m_MinimumRequiredAccountValue", m.MinimumRequiredAccountValue())
	w.Int64Field("m_BeginTimeForDayMargin", m.BeginTimeForDayMargin())
	w.Int64Field("m_EndTimeForDayMargin", m.EndTimeForDayMargin())
	w.StringField("m_DayMarginTimeZone", m.DayMarginTimeZone())
	w.Uint8Field("m_IsSnapshot", m.IsSnapshot())
	w.Uint8Field("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Uint8Field("m_IsDeleted", m.IsDeleted())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Uint8Field("m_UseMasterFirm_SymbolLimitsArray", m.UseMasterFirm_SymbolLimitsArray())
	w.Uint8Field("m_UseMasterFirm_TradeFees", m.UseMasterFirm_TradeFees())
	w.Uint8Field("m_UseMasterFirm_TradeFeePerShare", m.UseMasterFirm_TradeFeePerShare())
	w.Uint8Field("m_UseMasterFirm_RequireSufficientMarginForNewPositions", m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMargin", m.UseMasterFirm_UsePercentOfMargin())
	w.Uint8Field("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Uint8Field("m_UseMasterFirm_MinimumRequiredAccountValue", m.UseMasterFirm_MinimumRequiredAccountValue())
	w.Uint8Field("m_UseMasterFirm_MarginTimeSettings", m.UseMasterFirm_MarginTimeSettings())
	w.Uint8Field("m_UseMasterFirm_TradingIsDisabled", m.UseMasterFirm_TradingIsDisabled())
	w.Uint8Field("m_IsTradeStatisticsPublicallyShared", m.IsTradeStatisticsPublicallyShared())
	w.Uint8Field("m_IsReadOnlyFollowingRequestsAllowed", m.IsReadOnlyFollowingRequestsAllowed())
	w.Uint8Field("m_IsTradeAccountSharingAllowed", m.IsTradeAccountSharingAllowed())
	w.Uint8Field("m_ReadOnlyShareToAllSCUsernames", m.ReadOnlyShareToAllSCUsernames())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArray", m.UseMasterFirm_SymbolCommissionsArray())
	w.Uint8Field("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTrading", m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.Float64Field("m_OpenPositionsProfitLossBasedOnSettlement", m.OpenPositionsProfitLossBasedOnSettlement())
	w.Float64Field("m_MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("m_MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Uint8Field("m_UseMasterFirm_TradeFeesFullOverride", m.UseMasterFirm_TradeFeesFullOverride())
	w.Uint8Field("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginFullOverride", m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.Float64Field("m_PeakMarginRequirement", m.PeakMarginRequirement())
	w.Uint8Field("m_LiquidationOnlyMode", m.LiquidationOnlyMode())
	w.Uint8Field("m_FlattenPositionsWhenUnderMarginIntradayTriggered", m.FlattenPositionsWhenUnderMarginIntradayTriggered())
	w.Uint8Field("m_FlattenPositionsWhenUnderMinimumAccountValueTriggered", m.FlattenPositionsWhenUnderMinimumAccountValueTriggered())
	w.Float64Field("m_AccountValueAtEndOfDayCaptureTime", m.AccountValueAtEndOfDayCaptureTime())
	w.Int64Field("m_EndOfDayCaptureTime", m.EndOfDayCaptureTime())
	w.Uint8Field("m_CustomerOrFirm", m.CustomerOrFirm())
	w.Uint8Field("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginIntraday", m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantity", m.MasterFirm_MaximumOrderQuantity())
	w.Int64Field("m_LastTriggerDateTimeUTCForDailyLossLimit", m.LastTriggerDateTimeUTCForDailyLossLimit())
	w.Uint8Field("m_OpenPositionsProfitLossIsDelayed", m.OpenPositionsProfitLossIsDelayed())
	w.StringField("m_ExchangeTraderId", m.ExchangeTraderId())
	return w.Finish(), nil
}

func (m TradeAccountDataResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10116)
	w.Uint32Field("m_RequestID", m.RequestID())
	w.Uint8Field("m_TradeAccountNotExist", m.TradeAccountNotExist())
	w.StringField("m_TradeAccount", m.TradeAccount())
	w.Uint8Field("m_IsSimulated", m.IsSimulated())
	w.StringField("m_CurrencyCode", m.CurrencyCode())
	w.Float64Field("m_CurrentCashBalance", m.CurrentCashBalance())
	w.Float64Field("m_AvailableFundsForNewPositions", m.AvailableFundsForNewPositions())
	w.Float64Field("m_MarginRequirement", m.MarginRequirement())
	w.Float64Field("m_AccountValue", m.AccountValue())
	w.Float64Field("m_OpenPositionsProfitLoss", m.OpenPositionsProfitLoss())
	w.Float64Field("m_DailyProfitLoss", m.DailyProfitLoss())
	w.Uint64Field("m_TransactionIdentifierForCashBalanceAdjustment", m.TransactionIdentifierForCashBalanceAdjustment())
	w.Int64Field("m_LastTransactionDateTime", m.LastTransactionDateTime())
	w.Float64Field("m_TrailingAccountValueAtWhichToNotAllowNewPositions", m.TrailingAccountValueAtWhichToNotAllowNewPositions())
	w.Float64Field("m_CalculatedDailyNetLossLimitInAccountCurrency", m.CalculatedDailyNetLossLimitInAccountCurrency())
	w.Uint8Field("m_DailyNetLossLimitHasBeenReached", m.DailyNetLossLimitHasBeenReached())
	w.Int64Field("m_LastResetDailyNetLossManagementVariablesDateTimeUTC", m.LastResetDailyNetLossManagementVariablesDateTimeUTC())
	w.Uint8Field("m_IsUnderRequiredMargin", m.IsUnderRequiredMargin())
	w.Float32Field("m_DailyNetLossLimitInAccountCurrency", m.DailyNetLossLimitInAccountCurrency())
	w.Int32Field("m_PercentOfCashBalanceForDailyNetLossLimit", m.PercentOfCashBalanceForDailyNetLossLimit())
	w.Uint8Field("m_UseTrailingAccountValueToNotAllowIncreaseInPositions", m.UseTrailingAccountValueToNotAllowIncreaseInPositions())
	w.Uint8Field("m_DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Uint8Field("m_FlattenPositionsAtDailyLossLimit", m.FlattenPositionsAtDailyLossLimit())
	w.Uint8Field("m_ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.Uint8Field("m_FlattenPositionsWhenUnderMarginIntraday", m.FlattenPositionsWhenUnderMarginIntraday())
	w.Uint8Field("m_FlattenPositionsWhenUnderMarginAtEndOfDay", m.FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.StringField("m_SenderSubID", m.SenderSubID())
	w.StringField("m_SenderLocationId", m.SenderLocationId())
	w.StringField("m_SelfMatchPreventionID", m.SelfMatchPreventionID())
	w.Int32Field("m_CTICode", m.CTICode())
	w.Uint8Field("m_TradeAccountIsReadOnly", m.TradeAccountIsReadOnly())
	w.Int32Field("m_MaximumGlobalPositionQuantity", m.MaximumGlobalPositionQuantity())
	w.Float64Field("m_TradeFeePerContract", m.TradeFeePerContract())
	w.Float64Field("m_TradeFeePerShare", m.TradeFeePerShare())
	w.Uint8Field("m_RequireSufficientMarginForNewPositions", m.RequireSufficientMarginForNewPositions())
	w.Int32Field("m_UsePercentOfMargin", m.UsePercentOfMargin())
	w.Int32Field("m_UsePercentOfMarginForDayTrading", m.UsePercentOfMarginForDayTrading())
	w.Int32Field("m_MaximumAllowedAccountBalanceForPositionsAsPercentage", m.MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.StringField("m_FirmID", m.FirmID())
	w.Uint8Field("m_TradingIsDisabled", m.TradingIsDisabled())
	w.StringField("m_DescriptiveName", m.DescriptiveName())
	w.Uint8Field("m_IsMasterFirmControlAccount", m.IsMasterFirmControlAccount())
	w.Float64Field("m_MinimumRequiredAccountValue", m.MinimumRequiredAccountValue())
	w.Int64Field("m_BeginTimeForDayMargin", m.BeginTimeForDayMargin())
	w.Int64Field("m_EndTimeForDayMargin", m.EndTimeForDayMargin())
	w.StringField("m_DayMarginTimeZone", m.DayMarginTimeZone())
	w.Uint8Field("m_IsSnapshot", m.IsSnapshot())
	w.Uint8Field("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Uint8Field("m_IsDeleted", m.IsDeleted())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday", m.UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday())
	w.Uint8Field("m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay", m.UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay())
	w.Uint8Field("m_UseMasterFirm_SymbolLimitsArray", m.UseMasterFirm_SymbolLimitsArray())
	w.Uint8Field("m_UseMasterFirm_TradeFees", m.UseMasterFirm_TradeFees())
	w.Uint8Field("m_UseMasterFirm_TradeFeePerShare", m.UseMasterFirm_TradeFeePerShare())
	w.Uint8Field("m_UseMasterFirm_RequireSufficientMarginForNewPositions", m.UseMasterFirm_RequireSufficientMarginForNewPositions())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMargin", m.UseMasterFirm_UsePercentOfMargin())
	w.Uint8Field("m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage", m.UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage())
	w.Uint8Field("m_UseMasterFirm_MinimumRequiredAccountValue", m.UseMasterFirm_MinimumRequiredAccountValue())
	w.Uint8Field("m_UseMasterFirm_MarginTimeSettings", m.UseMasterFirm_MarginTimeSettings())
	w.Uint8Field("m_UseMasterFirm_TradingIsDisabled", m.UseMasterFirm_TradingIsDisabled())
	w.Uint8Field("m_IsTradeStatisticsPublicallyShared", m.IsTradeStatisticsPublicallyShared())
	w.Uint8Field("m_IsReadOnlyFollowingRequestsAllowed", m.IsReadOnlyFollowingRequestsAllowed())
	w.Uint8Field("m_IsTradeAccountSharingAllowed", m.IsTradeAccountSharingAllowed())
	w.Uint8Field("m_ReadOnlyShareToAllSCUsernames", m.ReadOnlyShareToAllSCUsernames())
	w.Uint8Field("m_UseMasterFirm_SymbolCommissionsArray", m.UseMasterFirm_SymbolCommissionsArray())
	w.Uint8Field("m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit", m.UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTrading", m.UseMasterFirm_UsePercentOfMarginForDayTrading())
	w.Float64Field("m_OpenPositionsProfitLossBasedOnSettlement", m.OpenPositionsProfitLossBasedOnSettlement())
	w.Float64Field("m_MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("m_MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Uint8Field("m_UseMasterFirm_TradeFeesFullOverride", m.UseMasterFirm_TradeFeesFullOverride())
	w.Uint8Field("m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders", m.UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginFullOverride", m.UseMasterFirm_UsePercentOfMarginFullOverride())
	w.Uint8Field("m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride", m.UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride())
	w.Float64Field("m_PeakMarginRequirement", m.PeakMarginRequirement())
	w.Uint8Field("m_LiquidationOnlyMode", m.LiquidationOnlyMode())
	w.Uint8Field("m_FlattenPositionsWhenUnderMarginIntradayTriggered", m.FlattenPositionsWhenUnderMarginIntradayTriggered())
	w.Uint8Field("m_FlattenPositionsWhenUnderMinimumAccountValueTriggered", m.FlattenPositionsWhenUnderMinimumAccountValueTriggered())
	w.Float64Field("m_AccountValueAtEndOfDayCaptureTime", m.AccountValueAtEndOfDayCaptureTime())
	w.Int64Field("m_EndOfDayCaptureTime", m.EndOfDayCaptureTime())
	w.Uint8Field("m_CustomerOrFirm", m.CustomerOrFirm())
	w.Uint8Field("m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet", m.MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue", m.MasterFirm_FlattenCancelWhenUnderMinimumAccountValue())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginIntraday", m.MasterFirm_FlattenCancelWhenUnderMarginIntraday())
	w.Uint8Field("m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay", m.MasterFirm_FlattenCancelWhenUnderMarginEndOfDay())
	w.Uint32Field("m_MasterFirm_MaximumOrderQuantity", m.MasterFirm_MaximumOrderQuantity())
	w.Int64Field("m_LastTriggerDateTimeUTCForDailyLossLimit", m.LastTriggerDateTimeUTCForDailyLossLimit())
	w.Uint8Field("m_OpenPositionsProfitLossIsDelayed", m.OpenPositionsProfitLossIsDelayed())
	w.StringField("m_ExchangeTraderId", m.ExchangeTraderId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10116 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccountNotExist(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSimulated(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrencyCode(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentCashBalance(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAvailableFundsForNewPositions(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccountValue(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenPositionsProfitLoss(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyProfitLoss(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionIdentifierForCashBalanceAdjustment(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTransactionDateTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailingAccountValueAtWhichToNotAllowNewPositions(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetCalculatedDailyNetLossLimitInAccountCurrency(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimitHasBeenReached(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastResetDailyNetLossManagementVariablesDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsUnderRequiredMargin(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimitInAccountCurrency(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPercentOfCashBalanceForDailyNetLossLimit(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsAtDailyLossLimit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosePositionsAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginIntraday(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderSubID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderLocationId(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSelfMatchPreventionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCTICode(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccountIsReadOnly(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumGlobalPositionQuantity(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerContract(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerShare(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequireSufficientMarginForNewPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMargin(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsePercentOfMarginForDayTrading(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetFirmID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsDisabled(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescriptiveName(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsMasterFirmControlAccount(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMinimumRequiredAccountValue(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBeginTimeForDayMargin(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetEndTimeForDayMargin(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDayMarginTimeZone(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSnapshot(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFirstMessageInBatch(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsDeleted(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolLimitsArray(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFees(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeePerShare(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_RequireSufficientMarginForNewPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMargin(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MinimumRequiredAccountValue(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_MarginTimeSettings(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradingIsDisabled(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeStatisticsPublicallyShared(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsReadOnlyFollowingRequestsAllowed(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsTradeAccountSharingAllowed(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetReadOnlyShareToAllSCUsernames(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_SymbolCommissionsArray(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTrading(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenPositionsProfitLossBasedOnSettlement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirementFull(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirementFullPositionsOnly(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_TradeFeesFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetPeakMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLiquidationOnlyMode(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMarginIntradayTriggered(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccountValueAtEndOfDayCaptureTime(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetEndOfDayCaptureTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetCustomerOrFirm(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMasterFirm_MaximumOrderQuantity(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTriggerDateTimeUTCForDailyLossLimit(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenPositionsProfitLossIsDelayed(r.Uint8())
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

func (m *TradeAccountDataResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10116 {
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
		case "m_RequestID":
			m.SetRequestID(r.Uint32())
		case "m_TradeAccountNotExist":
			m.SetTradeAccountNotExist(r.Uint8())
		case "m_TradeAccount":
			m.SetTradeAccount(r.String())
		case "m_IsSimulated":
			m.SetIsSimulated(r.Uint8())
		case "m_CurrencyCode":
			m.SetCurrencyCode(r.String())
		case "m_CurrentCashBalance":
			m.SetCurrentCashBalance(r.Float64())
		case "m_AvailableFundsForNewPositions":
			m.SetAvailableFundsForNewPositions(r.Float64())
		case "m_MarginRequirement":
			m.SetMarginRequirement(r.Float64())
		case "m_AccountValue":
			m.SetAccountValue(r.Float64())
		case "m_OpenPositionsProfitLoss":
			m.SetOpenPositionsProfitLoss(r.Float64())
		case "m_DailyProfitLoss":
			m.SetDailyProfitLoss(r.Float64())
		case "m_TransactionIdentifierForCashBalanceAdjustment":
			m.SetTransactionIdentifierForCashBalanceAdjustment(r.Uint64())
		case "m_LastTransactionDateTime":
			m.SetLastTransactionDateTime(r.Int64())
		case "m_TrailingAccountValueAtWhichToNotAllowNewPositions":
			m.SetTrailingAccountValueAtWhichToNotAllowNewPositions(r.Float64())
		case "m_CalculatedDailyNetLossLimitInAccountCurrency":
			m.SetCalculatedDailyNetLossLimitInAccountCurrency(r.Float64())
		case "m_DailyNetLossLimitHasBeenReached":
			m.SetDailyNetLossLimitHasBeenReached(r.Uint8())
		case "m_LastResetDailyNetLossManagementVariablesDateTimeUTC":
			m.SetLastResetDailyNetLossManagementVariablesDateTimeUTC(r.Int64())
		case "m_IsUnderRequiredMargin":
			m.SetIsUnderRequiredMargin(r.Uint8())
		case "m_DailyNetLossLimitInAccountCurrency":
			m.SetDailyNetLossLimitInAccountCurrency(r.Float32())
		case "m_PercentOfCashBalanceForDailyNetLossLimit":
			m.SetPercentOfCashBalanceForDailyNetLossLimit(r.Int32())
		case "m_UseTrailingAccountValueToNotAllowIncreaseInPositions":
			m.SetUseTrailingAccountValueToNotAllowIncreaseInPositions(r.Uint8())
		case "m_DoNotAllowIncreaseInPositionsAtDailyLossLimit":
			m.SetDoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
		case "m_FlattenPositionsAtDailyLossLimit":
			m.SetFlattenPositionsAtDailyLossLimit(r.Uint8())
		case "m_ClosePositionsAtEndOfDay":
			m.SetClosePositionsAtEndOfDay(r.Uint8())
		case "m_FlattenPositionsWhenUnderMarginIntraday":
			m.SetFlattenPositionsWhenUnderMarginIntraday(r.Uint8())
		case "m_FlattenPositionsWhenUnderMarginAtEndOfDay":
			m.SetFlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
		case "m_SenderSubID":
			m.SetSenderSubID(r.String())
		case "m_SenderLocationId":
			m.SetSenderLocationId(r.String())
		case "m_SelfMatchPreventionID":
			m.SetSelfMatchPreventionID(r.String())
		case "m_CTICode":
			m.SetCTICode(r.Int32())
		case "m_TradeAccountIsReadOnly":
			m.SetTradeAccountIsReadOnly(r.Uint8())
		case "m_MaximumGlobalPositionQuantity":
			m.SetMaximumGlobalPositionQuantity(r.Int32())
		case "m_TradeFeePerContract":
			m.SetTradeFeePerContract(r.Float64())
		case "m_TradeFeePerShare":
			m.SetTradeFeePerShare(r.Float64())
		case "m_RequireSufficientMarginForNewPositions":
			m.SetRequireSufficientMarginForNewPositions(r.Uint8())
		case "m_UsePercentOfMargin":
			m.SetUsePercentOfMargin(r.Int32())
		case "m_UsePercentOfMarginForDayTrading":
			m.SetUsePercentOfMarginForDayTrading(r.Int32())
		case "m_MaximumAllowedAccountBalanceForPositionsAsPercentage":
			m.SetMaximumAllowedAccountBalanceForPositionsAsPercentage(r.Int32())
		case "m_FirmID":
			m.SetFirmID(r.String())
		case "m_TradingIsDisabled":
			m.SetTradingIsDisabled(r.Uint8())
		case "m_DescriptiveName":
			m.SetDescriptiveName(r.String())
		case "m_IsMasterFirmControlAccount":
			m.SetIsMasterFirmControlAccount(r.Uint8())
		case "m_MinimumRequiredAccountValue":
			m.SetMinimumRequiredAccountValue(r.Float64())
		case "m_BeginTimeForDayMargin":
			m.SetBeginTimeForDayMargin(r.Int64())
		case "m_EndTimeForDayMargin":
			m.SetEndTimeForDayMargin(r.Int64())
		case "m_DayMarginTimeZone":
			m.SetDayMarginTimeZone(r.String())
		case "m_IsSnapshot":
			m.SetIsSnapshot(r.Uint8())
		case "m_IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Uint8())
		case "m_IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Uint8())
		case "m_IsDeleted":
			m.SetIsDeleted(r.Uint8())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday(r.Uint8())
		case "m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay":
			m.SetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay(r.Uint8())
		case "m_UseMasterFirm_SymbolLimitsArray":
			m.SetUseMasterFirm_SymbolLimitsArray(r.Uint8())
		case "m_UseMasterFirm_TradeFees":
			m.SetUseMasterFirm_TradeFees(r.Uint8())
		case "m_UseMasterFirm_TradeFeePerShare":
			m.SetUseMasterFirm_TradeFeePerShare(r.Uint8())
		case "m_UseMasterFirm_RequireSufficientMarginForNewPositions":
			m.SetUseMasterFirm_RequireSufficientMarginForNewPositions(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMargin":
			m.SetUseMasterFirm_UsePercentOfMargin(r.Uint8())
		case "m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage":
			m.SetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage(r.Uint8())
		case "m_UseMasterFirm_MinimumRequiredAccountValue":
			m.SetUseMasterFirm_MinimumRequiredAccountValue(r.Uint8())
		case "m_UseMasterFirm_MarginTimeSettings":
			m.SetUseMasterFirm_MarginTimeSettings(r.Uint8())
		case "m_UseMasterFirm_TradingIsDisabled":
			m.SetUseMasterFirm_TradingIsDisabled(r.Uint8())
		case "m_IsTradeStatisticsPublicallyShared":
			m.SetIsTradeStatisticsPublicallyShared(r.Uint8())
		case "m_IsReadOnlyFollowingRequestsAllowed":
			m.SetIsReadOnlyFollowingRequestsAllowed(r.Uint8())
		case "m_IsTradeAccountSharingAllowed":
			m.SetIsTradeAccountSharingAllowed(r.Uint8())
		case "m_ReadOnlyShareToAllSCUsernames":
			m.SetReadOnlyShareToAllSCUsernames(r.Uint8())
		case "m_UseMasterFirm_SymbolCommissionsArray":
			m.SetUseMasterFirm_SymbolCommissionsArray(r.Uint8())
		case "m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit":
			m.SetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTrading":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTrading(r.Uint8())
		case "m_OpenPositionsProfitLossBasedOnSettlement":
			m.SetOpenPositionsProfitLossBasedOnSettlement(r.Float64())
		case "m_MarginRequirementFull":
			m.SetMarginRequirementFull(r.Float64())
		case "m_MarginRequirementFullPositionsOnly":
			m.SetMarginRequirementFullPositionsOnly(r.Float64())
		case "m_UseMasterFirm_TradeFeesFullOverride":
			m.SetUseMasterFirm_TradeFeesFullOverride(r.Uint8())
		case "m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders":
			m.SetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginFullOverride":
			m.SetUseMasterFirm_UsePercentOfMarginFullOverride(r.Uint8())
		case "m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride":
			m.SetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride(r.Uint8())
		case "m_PeakMarginRequirement":
			m.SetPeakMarginRequirement(r.Float64())
		case "m_LiquidationOnlyMode":
			m.SetLiquidationOnlyMode(r.Uint8())
		case "m_FlattenPositionsWhenUnderMarginIntradayTriggered":
			m.SetFlattenPositionsWhenUnderMarginIntradayTriggered(r.Uint8())
		case "m_FlattenPositionsWhenUnderMinimumAccountValueTriggered":
			m.SetFlattenPositionsWhenUnderMinimumAccountValueTriggered(r.Uint8())
		case "m_AccountValueAtEndOfDayCaptureTime":
			m.SetAccountValueAtEndOfDayCaptureTime(r.Float64())
		case "m_EndOfDayCaptureTime":
			m.SetEndOfDayCaptureTime(r.Int64())
		case "m_CustomerOrFirm":
			m.SetCustomerOrFirm(r.Uint8())
		case "m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet":
			m.SetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue":
			m.SetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginIntraday":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginIntraday(r.Uint8())
		case "m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay":
			m.SetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay(r.Uint8())
		case "m_MasterFirm_MaximumOrderQuantity":
			m.SetMasterFirm_MaximumOrderQuantity(r.Uint32())
		case "m_LastTriggerDateTimeUTCForDailyLossLimit":
			m.SetLastTriggerDateTimeUTCForDailyLossLimit(r.Int64())
		case "m_OpenPositionsProfitLossIsDelayed":
			m.SetOpenPositionsProfitLossIsDelayed(r.Uint8())
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
