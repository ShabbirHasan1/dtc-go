//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataSymbolCommissionResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10119,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	return w.Finish(), nil
}

func (m TradeAccountDataSymbolCommissionResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10119,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataSymbolCommissionResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10119,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	return w.Finish(), nil
}

func (m TradeAccountDataSymbolCommissionResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10119,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataSymbolCommissionResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10119)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Symbol", m.Symbol())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	return w.Finish(), nil
}

func (m TradeAccountDataSymbolCommissionResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10119)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Symbol", m.Symbol())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataSymbolCommissionResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10119)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Symbol", m.Symbol())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	return w.Finish(), nil
}

func (m TradeAccountDataSymbolCommissionResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10119)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Symbol", m.Symbol())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataSymbolCommissionResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10119 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerContract(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataSymbolCommissionResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10119 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeFeePerContract(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataSymbolCommissionResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10119 {
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "Symbol":
			m.SetSymbol(r.String())
		case "TradeFeePerContract":
			m.SetTradeFeePerContract(r.Float64())
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

func (m *TradeAccountDataSymbolCommissionResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10119 {
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "Symbol":
			m.SetSymbol(r.String())
		case "TradeFeePerContract":
			m.SetTradeFeePerContract(r.Float64())
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
