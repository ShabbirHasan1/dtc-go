//go:build !tinygo

// generated by github.com/moontrade/dtc-go/codegen/go at 2022-03-14 08:04:12.929995 +0800 WITA m=+0.027582959
package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataSymbolCommissionUpdate) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10120,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDeleted())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	return w.Finish(), nil
}

func (m TradeAccountDataSymbolCommissionUpdateBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10120,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDeleted())
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

func (m TradeAccountDataSymbolCommissionUpdatePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10120,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDeleted())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeFeePerContract())
	return w.Finish(), nil
}

func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10120,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDeleted())
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

func (m TradeAccountDataSymbolCommissionUpdate) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10120)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsDeleted", m.IsDeleted())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Symbol", m.Symbol())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	return w.Finish(), nil
}

func (m TradeAccountDataSymbolCommissionUpdateBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10120)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsDeleted", m.IsDeleted())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Symbol", m.Symbol())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataSymbolCommissionUpdatePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10120)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsDeleted", m.IsDeleted())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Symbol", m.Symbol())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	return w.Finish(), nil
}

func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10120)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsDeleted", m.IsDeleted())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Symbol", m.Symbol())
	w.Float64Field("TradeFeePerContract", m.TradeFeePerContract())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataSymbolCommissionUpdateBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 10120 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsDeleted(r.Bool())
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

func (m *TradeAccountDataSymbolCommissionUpdatePointerBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 10120 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsDeleted(r.Bool())
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

func (m *TradeAccountDataSymbolCommissionUpdateBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 10120 {
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
		case "IsDeleted":
			m.SetIsDeleted(r.Bool())
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

func (m *TradeAccountDataSymbolCommissionUpdatePointerBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 10120 {
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
		case "IsDeleted":
			m.SetIsDeleted(r.Bool())
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

func (m *TradeAccountDataSymbolCommissionUpdateBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataSymbolCommissionUpdatePointerBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}