//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalAccountBalanceResponseFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":606,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFinalResponse())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.String(m.TransactionId())
	return w.Finish(), nil
}

func (m HistoricalAccountBalanceResponseFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":606,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFinalResponse())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.String(m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalAccountBalanceResponseFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 606)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Float64Field("CashBalance", m.CashBalance())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("IsFinalResponse", m.IsFinalResponse())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.StringField("InfoText", m.InfoText())
	w.StringField("TransactionId", m.TransactionId())
	return w.Finish(), nil
}

func (m HistoricalAccountBalanceResponseFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 606)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Float64Field("CashBalance", m.CashBalance())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("IsFinalResponse", m.IsFinalResponse())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.StringField("InfoText", m.InfoText())
	w.StringField("TransactionId", m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalanceResponseFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 606 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetCashBalance(r.Float64())
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
	m.SetIsFinalResponse(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoAccountBalances(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionId(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalanceResponseFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 606 {
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
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "CashBalance":
			m.SetCashBalance(r.Float64())
		case "AccountCurrency":
			m.SetAccountCurrency(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "IsFinalResponse":
			m.SetIsFinalResponse(r.Uint8())
		case "NoAccountBalances":
			m.SetNoAccountBalances(r.Uint8())
		case "InfoText":
			m.SetInfoText(r.String())
		case "TransactionId":
			m.SetTransactionId(r.String())
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
