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

func (m TradeAccountDataDelete) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10118,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataDeleteBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10118,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataDeletePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10118,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataDeletePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10118,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataDelete) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10118)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataDeleteBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10118)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataDeletePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10118)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataDeletePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10118)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataDeleteBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 10118 {
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
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataDeletePointerBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 10118 {
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
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataDeleteBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 10118 {
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

func (m *TradeAccountDataDeletePointerBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 10118 {
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

func (m *TradeAccountDataDeleteBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataDeletePointerBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}
