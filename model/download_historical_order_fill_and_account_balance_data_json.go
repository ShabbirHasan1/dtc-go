//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m DownloadHistoricalOrderFillAndAccountBalanceData) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10138,\"F\":["...)
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10138,\"F\":["...)
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m DownloadHistoricalOrderFillAndAccountBalanceData) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10138)
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10138)
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10138 {
		return message.ErrWrongType
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10138 {
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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
