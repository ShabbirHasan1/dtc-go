//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m DepthEntry) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":0,\"F\":["...)
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Quantity())
	return w.Finish(), nil
}

func (m DepthEntryBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":0,\"F\":["...)
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Quantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m DepthEntry) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 0)
	w.Float64Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	return w.Finish(), nil
}

func (m DepthEntryBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 0)
	w.Float64Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *DepthEntryBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 0 {
		return message.ErrWrongType
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *DepthEntryBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 0 {
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
		case "Price":
			m.SetPrice(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float32())
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
