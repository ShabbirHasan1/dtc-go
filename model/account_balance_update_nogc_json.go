//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

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
	w.Uint8(m.IsUnderRequiredMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsUnderRequiredAccountValue())
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
	w.Uint8(m.IsUnderRequiredMargin())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ClosePositionsAtEndOfDay())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradingIsDisabled())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsUnderRequiredAccountValue())
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
	w.Uint8Field("IsUnderRequiredMargin", m.IsUnderRequiredMargin())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.Uint8Field("TradingIsDisabled", m.TradingIsDisabled())
	w.StringField("Description", m.Description())
	w.Uint8Field("IsUnderRequiredAccountValue", m.IsUnderRequiredAccountValue())
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
	w.Uint8Field("IsUnderRequiredMargin", m.IsUnderRequiredMargin())
	w.Uint8Field("ClosePositionsAtEndOfDay", m.ClosePositionsAtEndOfDay())
	w.Uint8Field("TradingIsDisabled", m.TradingIsDisabled())
	w.StringField("Description", m.Description())
	w.Uint8Field("IsUnderRequiredAccountValue", m.IsUnderRequiredAccountValue())
	w.Int64Field("TransactionDateTime", int64(m.TransactionDateTime()))
	w.Float64Field("MarginRequirementFull", m.MarginRequirementFull())
	w.Float64Field("MarginRequirementFullPositionsOnly", m.MarginRequirementFullPositionsOnly())
	w.Float64Field("PeakMarginRequirement", m.PeakMarginRequirement())
	return w.Finish(), nil
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
	m.SetIsUnderRequiredMargin(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosePositionsAtEndOfDay(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsDisabled(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsUnderRequiredAccountValue(r.Uint8())
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
			m.SetIsUnderRequiredMargin(r.Uint8())
		case "ClosePositionsAtEndOfDay":
			m.SetClosePositionsAtEndOfDay(r.Uint8())
		case "TradingIsDisabled":
			m.SetTradingIsDisabled(r.Uint8())
		case "Description":
			m.SetDescription(r.String())
		case "IsUnderRequiredAccountValue":
			m.SetIsUnderRequiredAccountValue(r.Uint8())
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
