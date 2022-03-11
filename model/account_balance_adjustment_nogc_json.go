//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":607,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.CreditAmount())
	w.Data = append(w.Data, ',')
	w.Float64(m.DebitAmount())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.String(m.Reason())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":607,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.CreditAmount())
	w.Data = append(w.Data, ',')
	w.Float64(m.DebitAmount())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.String(m.Reason())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 607)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("CreditAmount", m.CreditAmount())
	w.Float64Field("DebitAmount", m.DebitAmount())
	w.StringField("Currency", m.Currency())
	w.StringField("Reason", m.Reason())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 607)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("CreditAmount", m.CreditAmount())
	w.Float64Field("DebitAmount", m.DebitAmount())
	w.StringField("Currency", m.Currency())
	w.StringField("Reason", m.Reason())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 607 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCreditAmount(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDebitAmount(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetReason(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 607 {
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "CreditAmount":
			m.SetCreditAmount(r.Float64())
		case "DebitAmount":
			m.SetDebitAmount(r.Float64())
		case "Currency":
			m.SetCurrency(r.String())
		case "Reason":
			m.SetReason(r.String())
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
