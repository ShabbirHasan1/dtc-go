//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceUpdate) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":600,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.Float64(m.BalanceAvailableForNewPositions())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.SecuritiesValue())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPositionsProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyProfitLoss())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TransactionIdentifier())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingAccountValueToLimitPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DailyNetLossLimitReached())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.TransactionDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFull())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFullPositionsOnly())
	w.Data = append(w.Data, ',')
	w.Float64(m.PeakMarginRequirement())
	return w.Finish(), nil
}

func (m AccountBalanceUpdateBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":600,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.Float64(m.BalanceAvailableForNewPositions())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.SecuritiesValue())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPositionsProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyProfitLoss())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TransactionIdentifier())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingAccountValueToLimitPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DailyNetLossLimitReached())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.TransactionDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFull())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFullPositionsOnly())
	w.Data = append(w.Data, ',')
	w.Float64(m.PeakMarginRequirement())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceUpdatePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":600,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.Float64(m.BalanceAvailableForNewPositions())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.SecuritiesValue())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPositionsProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyProfitLoss())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TransactionIdentifier())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingAccountValueToLimitPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DailyNetLossLimitReached())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.TransactionDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFull())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFullPositionsOnly())
	w.Data = append(w.Data, ',')
	w.Float64(m.PeakMarginRequirement())
	return w.Finish(), nil
}

func (m AccountBalanceUpdatePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":600,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.Float64(m.BalanceAvailableForNewPositions())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.SecuritiesValue())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPositionsProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyProfitLoss())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TransactionIdentifier())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingAccountValueToLimitPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DailyNetLossLimitReached())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.TransactionDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFull())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFullPositionsOnly())
	w.Data = append(w.Data, ',')
	w.Float64(m.PeakMarginRequirement())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceUpdate) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 600)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("CashBalance", m.CashBalance())
	w.Float64Field("BalanceAvailableForNewPositions", m.BalanceAvailableForNewPositions())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("SecuritiesValue", m.SecuritiesValue())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("OpenPositionsProfitLoss", m.OpenPositionsProfitLoss())
	w.Float64Field("DailyProfitLoss", m.DailyProfitLoss())
	w.StringField("InfoText", m.InfoText())
	w.Uint64Field("TransactionIdentifier", m.TransactionIdentifier())
	w.Float64Field("DailyNetLossLimit", m.DailyNetLossLimit())
	w.Float64Field("TrailingAccountValueToLimitPositions", m.TrailingAccountValueToLimitPositions())
	w.Uint8Field("DailyNetLossLimitReached", m.DailyNetLossLimitReached())
	w.BoolField("IsUnderRequiredMargin", m.IsUnderRequiredMargin())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.BoolField("TradingIsDisabled", m.TradingIsDisabled())
	w.StringField("Description", m.Description())
	w.BoolField("IsUnderRequiredAccountValue", m.IsUnderRequiredAccountValue())
	w.Int64Field("TransactionDateTime", int64(m.TransactionDateTime()))
	w.Float64Field("MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Float64Field("PeakMarginRequirement", m.PeakMarginRequirement())
	return w.Finish(), nil
}

func (m AccountBalanceUpdateBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 600)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("CashBalance", m.CashBalance())
	w.Float64Field("BalanceAvailableForNewPositions", m.BalanceAvailableForNewPositions())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("SecuritiesValue", m.SecuritiesValue())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("OpenPositionsProfitLoss", m.OpenPositionsProfitLoss())
	w.Float64Field("DailyProfitLoss", m.DailyProfitLoss())
	w.StringField("InfoText", m.InfoText())
	w.Uint64Field("TransactionIdentifier", m.TransactionIdentifier())
	w.Float64Field("DailyNetLossLimit", m.DailyNetLossLimit())
	w.Float64Field("TrailingAccountValueToLimitPositions", m.TrailingAccountValueToLimitPositions())
	w.Uint8Field("DailyNetLossLimitReached", m.DailyNetLossLimitReached())
	w.BoolField("IsUnderRequiredMargin", m.IsUnderRequiredMargin())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.BoolField("TradingIsDisabled", m.TradingIsDisabled())
	w.StringField("Description", m.Description())
	w.BoolField("IsUnderRequiredAccountValue", m.IsUnderRequiredAccountValue())
	w.Int64Field("TransactionDateTime", int64(m.TransactionDateTime()))
	w.Float64Field("MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Float64Field("PeakMarginRequirement", m.PeakMarginRequirement())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceUpdatePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 600)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("CashBalance", m.CashBalance())
	w.Float64Field("BalanceAvailableForNewPositions", m.BalanceAvailableForNewPositions())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("SecuritiesValue", m.SecuritiesValue())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("OpenPositionsProfitLoss", m.OpenPositionsProfitLoss())
	w.Float64Field("DailyProfitLoss", m.DailyProfitLoss())
	w.StringField("InfoText", m.InfoText())
	w.Uint64Field("TransactionIdentifier", m.TransactionIdentifier())
	w.Float64Field("DailyNetLossLimit", m.DailyNetLossLimit())
	w.Float64Field("TrailingAccountValueToLimitPositions", m.TrailingAccountValueToLimitPositions())
	w.Uint8Field("DailyNetLossLimitReached", m.DailyNetLossLimitReached())
	w.BoolField("IsUnderRequiredMargin", m.IsUnderRequiredMargin())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.BoolField("TradingIsDisabled", m.TradingIsDisabled())
	w.StringField("Description", m.Description())
	w.BoolField("IsUnderRequiredAccountValue", m.IsUnderRequiredAccountValue())
	w.Int64Field("TransactionDateTime", int64(m.TransactionDateTime()))
	w.Float64Field("MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Float64Field("PeakMarginRequirement", m.PeakMarginRequirement())
	return w.Finish(), nil
}

func (m AccountBalanceUpdatePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 600)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("CashBalance", m.CashBalance())
	w.Float64Field("BalanceAvailableForNewPositions", m.BalanceAvailableForNewPositions())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("SecuritiesValue", m.SecuritiesValue())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("OpenPositionsProfitLoss", m.OpenPositionsProfitLoss())
	w.Float64Field("DailyProfitLoss", m.DailyProfitLoss())
	w.StringField("InfoText", m.InfoText())
	w.Uint64Field("TransactionIdentifier", m.TransactionIdentifier())
	w.Float64Field("DailyNetLossLimit", m.DailyNetLossLimit())
	w.Float64Field("TrailingAccountValueToLimitPositions", m.TrailingAccountValueToLimitPositions())
	w.Uint8Field("DailyNetLossLimitReached", m.DailyNetLossLimitReached())
	w.BoolField("IsUnderRequiredMargin", m.IsUnderRequiredMargin())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.BoolField("TradingIsDisabled", m.TradingIsDisabled())
	w.StringField("Description", m.Description())
	w.BoolField("IsUnderRequiredAccountValue", m.IsUnderRequiredAccountValue())
	w.Int64Field("TransactionDateTime", int64(m.TransactionDateTime()))
	w.Float64Field("MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Float64Field("PeakMarginRequirement", m.PeakMarginRequirement())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceUpdateBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 600 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCashBalance(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBalanceAvailableForNewPositions(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccountCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecuritiesValue(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoAccountBalances(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnsolicited(r.Uint8())
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
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionIdentifier(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimit(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailingAccountValueToLimitPositions(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimitReached(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsUnderRequiredMargin(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosePositionsAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsDisabled(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsUnderRequiredAccountValue(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
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
	m.SetPeakMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceUpdatePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 600 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCashBalance(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBalanceAvailableForNewPositions(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccountCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecuritiesValue(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoAccountBalances(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnsolicited(r.Uint8())
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
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionIdentifier(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimit(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailingAccountValueToLimitPositions(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimitReached(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsUnderRequiredMargin(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosePositionsAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsDisabled(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsUnderRequiredAccountValue(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
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
	m.SetPeakMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceUpdateBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 600 {
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
			m.SetRequestID(r.Int32())
		case "CashBalance":
			m.SetCashBalance(r.Float64())
		case "BalanceAvailableForNewPositions":
			m.SetBalanceAvailableForNewPositions(r.Float64())
		case "AccountCurrency":
			m.SetAccountCurrency(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "SecuritiesValue":
			m.SetSecuritiesValue(r.Float64())
		case "MarginRequirement":
			m.SetMarginRequirement(r.Float64())
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "NoAccountBalances":
			m.SetNoAccountBalances(r.Uint8())
		case "Unsolicited":
			m.SetUnsolicited(r.Uint8())
		case "OpenPositionsProfitLoss":
			m.SetOpenPositionsProfitLoss(r.Float64())
		case "DailyProfitLoss":
			m.SetDailyProfitLoss(r.Float64())
		case "InfoText":
			m.SetInfoText(r.String())
		case "TransactionIdentifier":
			m.SetTransactionIdentifier(r.Uint64())
		case "DailyNetLossLimit":
			m.SetDailyNetLossLimit(r.Float64())
		case "TrailingAccountValueToLimitPositions":
			m.SetTrailingAccountValueToLimitPositions(r.Float64())
		case "DailyNetLossLimitReached":
			m.SetDailyNetLossLimitReached(r.Uint8())
		case "IsUnderRequiredMargin":
			m.SetIsUnderRequiredMargin(r.Bool())
		case "ClosePositionsAtEndOfDay":
			m.SetClosePositionsAtEndOfDay(r.Uint8())
		case "TradingIsDisabled":
			m.SetTradingIsDisabled(r.Bool())
		case "Description":
			m.SetDescription(r.String())
		case "IsUnderRequiredAccountValue":
			m.SetIsUnderRequiredAccountValue(r.Bool())
		case "TransactionDateTime":
			m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "MarginRequirementFull":
			m.SetMarginRequirementFull(r.Float64())
		case "MarginRequirementFullPositionsOnly":
			m.SetMarginRequirementFullPositionsOnly(r.Float64())
		case "PeakMarginRequirement":
			m.SetPeakMarginRequirement(r.Float64())
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

func (m *AccountBalanceUpdatePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 600 {
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
			m.SetRequestID(r.Int32())
		case "CashBalance":
			m.SetCashBalance(r.Float64())
		case "BalanceAvailableForNewPositions":
			m.SetBalanceAvailableForNewPositions(r.Float64())
		case "AccountCurrency":
			m.SetAccountCurrency(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "SecuritiesValue":
			m.SetSecuritiesValue(r.Float64())
		case "MarginRequirement":
			m.SetMarginRequirement(r.Float64())
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "NoAccountBalances":
			m.SetNoAccountBalances(r.Uint8())
		case "Unsolicited":
			m.SetUnsolicited(r.Uint8())
		case "OpenPositionsProfitLoss":
			m.SetOpenPositionsProfitLoss(r.Float64())
		case "DailyProfitLoss":
			m.SetDailyProfitLoss(r.Float64())
		case "InfoText":
			m.SetInfoText(r.String())
		case "TransactionIdentifier":
			m.SetTransactionIdentifier(r.Uint64())
		case "DailyNetLossLimit":
			m.SetDailyNetLossLimit(r.Float64())
		case "TrailingAccountValueToLimitPositions":
			m.SetTrailingAccountValueToLimitPositions(r.Float64())
		case "DailyNetLossLimitReached":
			m.SetDailyNetLossLimitReached(r.Uint8())
		case "IsUnderRequiredMargin":
			m.SetIsUnderRequiredMargin(r.Bool())
		case "ClosePositionsAtEndOfDay":
			m.SetClosePositionsAtEndOfDay(r.Uint8())
		case "TradingIsDisabled":
			m.SetTradingIsDisabled(r.Bool())
		case "Description":
			m.SetDescription(r.String())
		case "IsUnderRequiredAccountValue":
			m.SetIsUnderRequiredAccountValue(r.Bool())
		case "TransactionDateTime":
			m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "MarginRequirementFull":
			m.SetMarginRequirementFull(r.Float64())
		case "MarginRequirementFullPositionsOnly":
			m.SetMarginRequirementFullPositionsOnly(r.Float64())
		case "PeakMarginRequirement":
			m.SetPeakMarginRequirement(r.Float64())
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
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceUpdateFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":600,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.Float64(m.BalanceAvailableForNewPositions())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.SecuritiesValue())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPositionsProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyProfitLoss())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TransactionIdentifier())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingAccountValueToLimitPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DailyNetLossLimitReached())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.TransactionDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFull())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFullPositionsOnly())
	w.Data = append(w.Data, ',')
	w.Float64(m.PeakMarginRequirement())
	return w.Finish(), nil
}

func (m AccountBalanceUpdateFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":600,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.Float64(m.BalanceAvailableForNewPositions())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.SecuritiesValue())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPositionsProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyProfitLoss())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TransactionIdentifier())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingAccountValueToLimitPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DailyNetLossLimitReached())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.TransactionDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFull())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFullPositionsOnly())
	w.Data = append(w.Data, ',')
	w.Float64(m.PeakMarginRequirement())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceUpdateFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":600,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.Float64(m.BalanceAvailableForNewPositions())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.SecuritiesValue())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPositionsProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyProfitLoss())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TransactionIdentifier())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingAccountValueToLimitPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DailyNetLossLimitReached())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.TransactionDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFull())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFullPositionsOnly())
	w.Data = append(w.Data, ',')
	w.Float64(m.PeakMarginRequirement())
	return w.Finish(), nil
}

func (m AccountBalanceUpdateFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":600,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.Float64(m.BalanceAvailableForNewPositions())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.SecuritiesValue())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPositionsProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyProfitLoss())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TransactionIdentifier())
	w.Data = append(w.Data, ',')
	w.Float64(m.DailyNetLossLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingAccountValueToLimitPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DailyNetLossLimitReached())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Bool(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsUnderRequiredAccountValue())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.TransactionDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFull())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirementFullPositionsOnly())
	w.Data = append(w.Data, ',')
	w.Float64(m.PeakMarginRequirement())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceUpdateFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 600)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("CashBalance", m.CashBalance())
	w.Float64Field("BalanceAvailableForNewPositions", m.BalanceAvailableForNewPositions())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("SecuritiesValue", m.SecuritiesValue())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("OpenPositionsProfitLoss", m.OpenPositionsProfitLoss())
	w.Float64Field("DailyProfitLoss", m.DailyProfitLoss())
	w.StringField("InfoText", m.InfoText())
	w.Uint64Field("TransactionIdentifier", m.TransactionIdentifier())
	w.Float64Field("DailyNetLossLimit", m.DailyNetLossLimit())
	w.Float64Field("TrailingAccountValueToLimitPositions", m.TrailingAccountValueToLimitPositions())
	w.Uint8Field("DailyNetLossLimitReached", m.DailyNetLossLimitReached())
	w.BoolField("IsUnderRequiredMargin", m.IsUnderRequiredMargin())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.BoolField("TradingIsDisabled", m.TradingIsDisabled())
	w.StringField("Description", m.Description())
	w.BoolField("IsUnderRequiredAccountValue", m.IsUnderRequiredAccountValue())
	w.Int64Field("TransactionDateTime", int64(m.TransactionDateTime()))
	w.Float64Field("MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Float64Field("PeakMarginRequirement", m.PeakMarginRequirement())
	return w.Finish(), nil
}

func (m AccountBalanceUpdateFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 600)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("CashBalance", m.CashBalance())
	w.Float64Field("BalanceAvailableForNewPositions", m.BalanceAvailableForNewPositions())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("SecuritiesValue", m.SecuritiesValue())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("OpenPositionsProfitLoss", m.OpenPositionsProfitLoss())
	w.Float64Field("DailyProfitLoss", m.DailyProfitLoss())
	w.StringField("InfoText", m.InfoText())
	w.Uint64Field("TransactionIdentifier", m.TransactionIdentifier())
	w.Float64Field("DailyNetLossLimit", m.DailyNetLossLimit())
	w.Float64Field("TrailingAccountValueToLimitPositions", m.TrailingAccountValueToLimitPositions())
	w.Uint8Field("DailyNetLossLimitReached", m.DailyNetLossLimitReached())
	w.BoolField("IsUnderRequiredMargin", m.IsUnderRequiredMargin())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.BoolField("TradingIsDisabled", m.TradingIsDisabled())
	w.StringField("Description", m.Description())
	w.BoolField("IsUnderRequiredAccountValue", m.IsUnderRequiredAccountValue())
	w.Int64Field("TransactionDateTime", int64(m.TransactionDateTime()))
	w.Float64Field("MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Float64Field("PeakMarginRequirement", m.PeakMarginRequirement())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceUpdateFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 600)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("CashBalance", m.CashBalance())
	w.Float64Field("BalanceAvailableForNewPositions", m.BalanceAvailableForNewPositions())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("SecuritiesValue", m.SecuritiesValue())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("OpenPositionsProfitLoss", m.OpenPositionsProfitLoss())
	w.Float64Field("DailyProfitLoss", m.DailyProfitLoss())
	w.StringField("InfoText", m.InfoText())
	w.Uint64Field("TransactionIdentifier", m.TransactionIdentifier())
	w.Float64Field("DailyNetLossLimit", m.DailyNetLossLimit())
	w.Float64Field("TrailingAccountValueToLimitPositions", m.TrailingAccountValueToLimitPositions())
	w.Uint8Field("DailyNetLossLimitReached", m.DailyNetLossLimitReached())
	w.BoolField("IsUnderRequiredMargin", m.IsUnderRequiredMargin())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.BoolField("TradingIsDisabled", m.TradingIsDisabled())
	w.StringField("Description", m.Description())
	w.BoolField("IsUnderRequiredAccountValue", m.IsUnderRequiredAccountValue())
	w.Int64Field("TransactionDateTime", int64(m.TransactionDateTime()))
	w.Float64Field("MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Float64Field("PeakMarginRequirement", m.PeakMarginRequirement())
	return w.Finish(), nil
}

func (m AccountBalanceUpdateFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 600)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("CashBalance", m.CashBalance())
	w.Float64Field("BalanceAvailableForNewPositions", m.BalanceAvailableForNewPositions())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("SecuritiesValue", m.SecuritiesValue())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("OpenPositionsProfitLoss", m.OpenPositionsProfitLoss())
	w.Float64Field("DailyProfitLoss", m.DailyProfitLoss())
	w.StringField("InfoText", m.InfoText())
	w.Uint64Field("TransactionIdentifier", m.TransactionIdentifier())
	w.Float64Field("DailyNetLossLimit", m.DailyNetLossLimit())
	w.Float64Field("TrailingAccountValueToLimitPositions", m.TrailingAccountValueToLimitPositions())
	w.Uint8Field("DailyNetLossLimitReached", m.DailyNetLossLimitReached())
	w.BoolField("IsUnderRequiredMargin", m.IsUnderRequiredMargin())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.BoolField("TradingIsDisabled", m.TradingIsDisabled())
	w.StringField("Description", m.Description())
	w.BoolField("IsUnderRequiredAccountValue", m.IsUnderRequiredAccountValue())
	w.Int64Field("TransactionDateTime", int64(m.TransactionDateTime()))
	w.Float64Field("MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Float64Field("PeakMarginRequirement", m.PeakMarginRequirement())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceUpdateFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 600 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCashBalance(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBalanceAvailableForNewPositions(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccountCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecuritiesValue(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoAccountBalances(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnsolicited(r.Uint8())
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
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionIdentifier(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimit(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailingAccountValueToLimitPositions(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimitReached(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsUnderRequiredMargin(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosePositionsAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsDisabled(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsUnderRequiredAccountValue(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
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
	m.SetPeakMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceUpdateFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 600 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCashBalance(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBalanceAvailableForNewPositions(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccountCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecuritiesValue(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoAccountBalances(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnsolicited(r.Uint8())
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
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionIdentifier(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimit(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailingAccountValueToLimitPositions(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDailyNetLossLimitReached(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsUnderRequiredMargin(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosePositionsAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsDisabled(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsUnderRequiredAccountValue(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
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
	m.SetPeakMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceUpdateFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 600 {
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
			m.SetRequestID(r.Int32())
		case "CashBalance":
			m.SetCashBalance(r.Float64())
		case "BalanceAvailableForNewPositions":
			m.SetBalanceAvailableForNewPositions(r.Float64())
		case "AccountCurrency":
			m.SetAccountCurrency(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "SecuritiesValue":
			m.SetSecuritiesValue(r.Float64())
		case "MarginRequirement":
			m.SetMarginRequirement(r.Float64())
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "NoAccountBalances":
			m.SetNoAccountBalances(r.Uint8())
		case "Unsolicited":
			m.SetUnsolicited(r.Uint8())
		case "OpenPositionsProfitLoss":
			m.SetOpenPositionsProfitLoss(r.Float64())
		case "DailyProfitLoss":
			m.SetDailyProfitLoss(r.Float64())
		case "InfoText":
			m.SetInfoText(r.String())
		case "TransactionIdentifier":
			m.SetTransactionIdentifier(r.Uint64())
		case "DailyNetLossLimit":
			m.SetDailyNetLossLimit(r.Float64())
		case "TrailingAccountValueToLimitPositions":
			m.SetTrailingAccountValueToLimitPositions(r.Float64())
		case "DailyNetLossLimitReached":
			m.SetDailyNetLossLimitReached(r.Uint8())
		case "IsUnderRequiredMargin":
			m.SetIsUnderRequiredMargin(r.Bool())
		case "ClosePositionsAtEndOfDay":
			m.SetClosePositionsAtEndOfDay(r.Uint8())
		case "TradingIsDisabled":
			m.SetTradingIsDisabled(r.Bool())
		case "Description":
			m.SetDescription(r.String())
		case "IsUnderRequiredAccountValue":
			m.SetIsUnderRequiredAccountValue(r.Bool())
		case "TransactionDateTime":
			m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "MarginRequirementFull":
			m.SetMarginRequirementFull(r.Float64())
		case "MarginRequirementFullPositionsOnly":
			m.SetMarginRequirementFullPositionsOnly(r.Float64())
		case "PeakMarginRequirement":
			m.SetPeakMarginRequirement(r.Float64())
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

func (m *AccountBalanceUpdateFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 600 {
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
			m.SetRequestID(r.Int32())
		case "CashBalance":
			m.SetCashBalance(r.Float64())
		case "BalanceAvailableForNewPositions":
			m.SetBalanceAvailableForNewPositions(r.Float64())
		case "AccountCurrency":
			m.SetAccountCurrency(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "SecuritiesValue":
			m.SetSecuritiesValue(r.Float64())
		case "MarginRequirement":
			m.SetMarginRequirement(r.Float64())
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "NoAccountBalances":
			m.SetNoAccountBalances(r.Uint8())
		case "Unsolicited":
			m.SetUnsolicited(r.Uint8())
		case "OpenPositionsProfitLoss":
			m.SetOpenPositionsProfitLoss(r.Float64())
		case "DailyProfitLoss":
			m.SetDailyProfitLoss(r.Float64())
		case "InfoText":
			m.SetInfoText(r.String())
		case "TransactionIdentifier":
			m.SetTransactionIdentifier(r.Uint64())
		case "DailyNetLossLimit":
			m.SetDailyNetLossLimit(r.Float64())
		case "TrailingAccountValueToLimitPositions":
			m.SetTrailingAccountValueToLimitPositions(r.Float64())
		case "DailyNetLossLimitReached":
			m.SetDailyNetLossLimitReached(r.Uint8())
		case "IsUnderRequiredMargin":
			m.SetIsUnderRequiredMargin(r.Bool())
		case "ClosePositionsAtEndOfDay":
			m.SetClosePositionsAtEndOfDay(r.Uint8())
		case "TradingIsDisabled":
			m.SetTradingIsDisabled(r.Bool())
		case "Description":
			m.SetDescription(r.String())
		case "IsUnderRequiredAccountValue":
			m.SetIsUnderRequiredAccountValue(r.Bool())
		case "TransactionDateTime":
			m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "MarginRequirementFull":
			m.SetMarginRequirementFull(r.Float64())
		case "MarginRequirementFullPositionsOnly":
			m.SetMarginRequirementFullPositionsOnly(r.Float64())
		case "PeakMarginRequirement":
			m.SetPeakMarginRequirement(r.Float64())
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
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceUpdateBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetCashBalance(r.ReadFloat64())
		case 3:
			m.SetBalanceAvailableForNewPositions(r.ReadFloat64())
		case 4:
			m.SetAccountCurrency(r.ReadString())
		case 5:
			m.SetTradeAccount(r.ReadString())
		case 6:
			m.SetSecuritiesValue(r.ReadFloat64())
		case 7:
			m.SetMarginRequirement(r.ReadFloat64())
		case 8:
			m.SetTotalNumberMessages(r.ReadInt32())
		case 9:
			m.SetMessageNumber(r.ReadInt32())
		case 10:
			m.SetNoAccountBalances(r.ReadUint8())
		case 11:
			m.SetUnsolicited(r.ReadUint8())
		case 12:
			m.SetOpenPositionsProfitLoss(r.ReadFloat64())
		case 13:
			m.SetDailyProfitLoss(r.ReadFloat64())
		case 14:
			m.SetInfoText(r.ReadString())
		case 15:
			m.SetTransactionIdentifier(r.ReadUint64())
		case 16:
			m.SetDailyNetLossLimit(r.ReadFloat64())
		case 17:
			m.SetTrailingAccountValueToLimitPositions(r.ReadFloat64())
		case 18:
			m.SetDailyNetLossLimitReached(r.ReadUint8())
		case 19:
			m.SetIsUnderRequiredMargin(r.ReadBool())
		case 20:
			m.SetClosePositionsAtEndOfDay(r.ReadUint8())
		case 21:
			m.SetTradingIsDisabled(r.ReadBool())
		case 22:
			m.SetDescription(r.ReadString())
		case 23:
			m.SetIsUnderRequiredAccountValue(r.ReadBool())
		case 24:
			m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 25:
			m.SetMarginRequirementFull(r.ReadFloat64())
		case 26:
			m.SetMarginRequirementFullPositionsOnly(r.ReadFloat64())
		case 27:
			m.SetPeakMarginRequirement(r.ReadFloat64())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceUpdatePointerBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetCashBalance(r.ReadFloat64())
		case 3:
			m.SetBalanceAvailableForNewPositions(r.ReadFloat64())
		case 4:
			m.SetAccountCurrency(r.ReadString())
		case 5:
			m.SetTradeAccount(r.ReadString())
		case 6:
			m.SetSecuritiesValue(r.ReadFloat64())
		case 7:
			m.SetMarginRequirement(r.ReadFloat64())
		case 8:
			m.SetTotalNumberMessages(r.ReadInt32())
		case 9:
			m.SetMessageNumber(r.ReadInt32())
		case 10:
			m.SetNoAccountBalances(r.ReadUint8())
		case 11:
			m.SetUnsolicited(r.ReadUint8())
		case 12:
			m.SetOpenPositionsProfitLoss(r.ReadFloat64())
		case 13:
			m.SetDailyProfitLoss(r.ReadFloat64())
		case 14:
			m.SetInfoText(r.ReadString())
		case 15:
			m.SetTransactionIdentifier(r.ReadUint64())
		case 16:
			m.SetDailyNetLossLimit(r.ReadFloat64())
		case 17:
			m.SetTrailingAccountValueToLimitPositions(r.ReadFloat64())
		case 18:
			m.SetDailyNetLossLimitReached(r.ReadUint8())
		case 19:
			m.SetIsUnderRequiredMargin(r.ReadBool())
		case 20:
			m.SetClosePositionsAtEndOfDay(r.ReadUint8())
		case 21:
			m.SetTradingIsDisabled(r.ReadBool())
		case 22:
			m.SetDescription(r.ReadString())
		case 23:
			m.SetIsUnderRequiredAccountValue(r.ReadBool())
		case 24:
			m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 25:
			m.SetMarginRequirementFull(r.ReadFloat64())
		case 26:
			m.SetMarginRequirementFullPositionsOnly(r.ReadFloat64())
		case 27:
			m.SetPeakMarginRequirement(r.ReadFloat64())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceUpdateBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 600)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, m.CashBalance())
	w.WriteFixed64Float64(3, m.BalanceAvailableForNewPositions())
	w.WriteString(4, m.AccountCurrency())
	w.WriteString(5, m.TradeAccount())
	w.WriteFixed64Float64(6, m.SecuritiesValue())
	w.WriteFixed64Float64(7, m.MarginRequirement())
	w.WriteVarint32(8, m.TotalNumberMessages())
	w.WriteVarint32(9, m.MessageNumber())
	w.WriteUvarint8(10, m.NoAccountBalances())
	w.WriteUvarint8(11, m.Unsolicited())
	w.WriteFixed64Float64(12, m.OpenPositionsProfitLoss())
	w.WriteFixed64Float64(13, m.DailyProfitLoss())
	w.WriteString(14, m.InfoText())
	w.WriteUvarint64(15, m.TransactionIdentifier())
	w.WriteFixed64Float64(16, m.DailyNetLossLimit())
	w.WriteFixed64Float64(17, m.TrailingAccountValueToLimitPositions())
	w.WriteUvarint8(18, m.DailyNetLossLimitReached())
	w.WriteBool(19, m.IsUnderRequiredMargin())
	w.WriteUvarint8(20, m.ClosePositionsAtEndOfDay())
	w.WriteBool(21, m.TradingIsDisabled())
	w.WriteString(22, m.Description())
	w.WriteBool(23, m.IsUnderRequiredAccountValue())
	w.WriteFixed64Int64(24, int64(m.TransactionDateTime()))
	w.WriteFixed64Float64(25, m.MarginRequirementFull())
	w.WriteFixed64Float64(26, m.MarginRequirementFullPositionsOnly())
	w.WriteFixed64Float64(27, m.PeakMarginRequirement())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceUpdatePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 600)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, m.CashBalance())
	w.WriteFixed64Float64(3, m.BalanceAvailableForNewPositions())
	w.WriteString(4, m.AccountCurrency())
	w.WriteString(5, m.TradeAccount())
	w.WriteFixed64Float64(6, m.SecuritiesValue())
	w.WriteFixed64Float64(7, m.MarginRequirement())
	w.WriteVarint32(8, m.TotalNumberMessages())
	w.WriteVarint32(9, m.MessageNumber())
	w.WriteUvarint8(10, m.NoAccountBalances())
	w.WriteUvarint8(11, m.Unsolicited())
	w.WriteFixed64Float64(12, m.OpenPositionsProfitLoss())
	w.WriteFixed64Float64(13, m.DailyProfitLoss())
	w.WriteString(14, m.InfoText())
	w.WriteUvarint64(15, m.TransactionIdentifier())
	w.WriteFixed64Float64(16, m.DailyNetLossLimit())
	w.WriteFixed64Float64(17, m.TrailingAccountValueToLimitPositions())
	w.WriteUvarint8(18, m.DailyNetLossLimitReached())
	w.WriteBool(19, m.IsUnderRequiredMargin())
	w.WriteUvarint8(20, m.ClosePositionsAtEndOfDay())
	w.WriteBool(21, m.TradingIsDisabled())
	w.WriteString(22, m.Description())
	w.WriteBool(23, m.IsUnderRequiredAccountValue())
	w.WriteFixed64Int64(24, int64(m.TransactionDateTime()))
	w.WriteFixed64Float64(25, m.MarginRequirementFull())
	w.WriteFixed64Float64(26, m.MarginRequirementFullPositionsOnly())
	w.WriteFixed64Float64(27, m.PeakMarginRequirement())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceUpdateFixedBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetCashBalance(r.ReadFloat64())
		case 3:
			m.SetBalanceAvailableForNewPositions(r.ReadFloat64())
		case 4:
			m.SetAccountCurrency(r.ReadString())
		case 5:
			m.SetTradeAccount(r.ReadString())
		case 6:
			m.SetSecuritiesValue(r.ReadFloat64())
		case 7:
			m.SetMarginRequirement(r.ReadFloat64())
		case 8:
			m.SetTotalNumberMessages(r.ReadInt32())
		case 9:
			m.SetMessageNumber(r.ReadInt32())
		case 10:
			m.SetNoAccountBalances(r.ReadUint8())
		case 11:
			m.SetUnsolicited(r.ReadUint8())
		case 12:
			m.SetOpenPositionsProfitLoss(r.ReadFloat64())
		case 13:
			m.SetDailyProfitLoss(r.ReadFloat64())
		case 14:
			m.SetInfoText(r.ReadString())
		case 15:
			m.SetTransactionIdentifier(r.ReadUint64())
		case 16:
			m.SetDailyNetLossLimit(r.ReadFloat64())
		case 17:
			m.SetTrailingAccountValueToLimitPositions(r.ReadFloat64())
		case 18:
			m.SetDailyNetLossLimitReached(r.ReadUint8())
		case 19:
			m.SetIsUnderRequiredMargin(r.ReadBool())
		case 20:
			m.SetClosePositionsAtEndOfDay(r.ReadUint8())
		case 21:
			m.SetTradingIsDisabled(r.ReadBool())
		case 22:
			m.SetDescription(r.ReadString())
		case 23:
			m.SetIsUnderRequiredAccountValue(r.ReadBool())
		case 24:
			m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 25:
			m.SetMarginRequirementFull(r.ReadFloat64())
		case 26:
			m.SetMarginRequirementFullPositionsOnly(r.ReadFloat64())
		case 27:
			m.SetPeakMarginRequirement(r.ReadFloat64())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceUpdateFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetCashBalance(r.ReadFloat64())
		case 3:
			m.SetBalanceAvailableForNewPositions(r.ReadFloat64())
		case 4:
			m.SetAccountCurrency(r.ReadString())
		case 5:
			m.SetTradeAccount(r.ReadString())
		case 6:
			m.SetSecuritiesValue(r.ReadFloat64())
		case 7:
			m.SetMarginRequirement(r.ReadFloat64())
		case 8:
			m.SetTotalNumberMessages(r.ReadInt32())
		case 9:
			m.SetMessageNumber(r.ReadInt32())
		case 10:
			m.SetNoAccountBalances(r.ReadUint8())
		case 11:
			m.SetUnsolicited(r.ReadUint8())
		case 12:
			m.SetOpenPositionsProfitLoss(r.ReadFloat64())
		case 13:
			m.SetDailyProfitLoss(r.ReadFloat64())
		case 14:
			m.SetInfoText(r.ReadString())
		case 15:
			m.SetTransactionIdentifier(r.ReadUint64())
		case 16:
			m.SetDailyNetLossLimit(r.ReadFloat64())
		case 17:
			m.SetTrailingAccountValueToLimitPositions(r.ReadFloat64())
		case 18:
			m.SetDailyNetLossLimitReached(r.ReadUint8())
		case 19:
			m.SetIsUnderRequiredMargin(r.ReadBool())
		case 20:
			m.SetClosePositionsAtEndOfDay(r.ReadUint8())
		case 21:
			m.SetTradingIsDisabled(r.ReadBool())
		case 22:
			m.SetDescription(r.ReadString())
		case 23:
			m.SetIsUnderRequiredAccountValue(r.ReadBool())
		case 24:
			m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 25:
			m.SetMarginRequirementFull(r.ReadFloat64())
		case 26:
			m.SetMarginRequirementFullPositionsOnly(r.ReadFloat64())
		case 27:
			m.SetPeakMarginRequirement(r.ReadFloat64())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceUpdateFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 600)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, m.CashBalance())
	w.WriteFixed64Float64(3, m.BalanceAvailableForNewPositions())
	w.WriteString(4, m.AccountCurrency())
	w.WriteString(5, m.TradeAccount())
	w.WriteFixed64Float64(6, m.SecuritiesValue())
	w.WriteFixed64Float64(7, m.MarginRequirement())
	w.WriteVarint32(8, m.TotalNumberMessages())
	w.WriteVarint32(9, m.MessageNumber())
	w.WriteUvarint8(10, m.NoAccountBalances())
	w.WriteUvarint8(11, m.Unsolicited())
	w.WriteFixed64Float64(12, m.OpenPositionsProfitLoss())
	w.WriteFixed64Float64(13, m.DailyProfitLoss())
	w.WriteString(14, m.InfoText())
	w.WriteUvarint64(15, m.TransactionIdentifier())
	w.WriteFixed64Float64(16, m.DailyNetLossLimit())
	w.WriteFixed64Float64(17, m.TrailingAccountValueToLimitPositions())
	w.WriteUvarint8(18, m.DailyNetLossLimitReached())
	w.WriteBool(19, m.IsUnderRequiredMargin())
	w.WriteUvarint8(20, m.ClosePositionsAtEndOfDay())
	w.WriteBool(21, m.TradingIsDisabled())
	w.WriteString(22, m.Description())
	w.WriteBool(23, m.IsUnderRequiredAccountValue())
	w.WriteFixed64Int64(24, int64(m.TransactionDateTime()))
	w.WriteFixed64Float64(25, m.MarginRequirementFull())
	w.WriteFixed64Float64(26, m.MarginRequirementFullPositionsOnly())
	w.WriteFixed64Float64(27, m.PeakMarginRequirement())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceUpdateFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 600)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, m.CashBalance())
	w.WriteFixed64Float64(3, m.BalanceAvailableForNewPositions())
	w.WriteString(4, m.AccountCurrency())
	w.WriteString(5, m.TradeAccount())
	w.WriteFixed64Float64(6, m.SecuritiesValue())
	w.WriteFixed64Float64(7, m.MarginRequirement())
	w.WriteVarint32(8, m.TotalNumberMessages())
	w.WriteVarint32(9, m.MessageNumber())
	w.WriteUvarint8(10, m.NoAccountBalances())
	w.WriteUvarint8(11, m.Unsolicited())
	w.WriteFixed64Float64(12, m.OpenPositionsProfitLoss())
	w.WriteFixed64Float64(13, m.DailyProfitLoss())
	w.WriteString(14, m.InfoText())
	w.WriteUvarint64(15, m.TransactionIdentifier())
	w.WriteFixed64Float64(16, m.DailyNetLossLimit())
	w.WriteFixed64Float64(17, m.TrailingAccountValueToLimitPositions())
	w.WriteUvarint8(18, m.DailyNetLossLimitReached())
	w.WriteBool(19, m.IsUnderRequiredMargin())
	w.WriteUvarint8(20, m.ClosePositionsAtEndOfDay())
	w.WriteBool(21, m.TradingIsDisabled())
	w.WriteString(22, m.Description())
	w.WriteBool(23, m.IsUnderRequiredAccountValue())
	w.WriteFixed64Int64(24, int64(m.TransactionDateTime()))
	w.WriteFixed64Float64(25, m.MarginRequirementFull())
	w.WriteFixed64Float64(26, m.MarginRequirementFullPositionsOnly())
	w.WriteFixed64Float64(27, m.PeakMarginRequirement())
	return w.Finish(), nil
}
