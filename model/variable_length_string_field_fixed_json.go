//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m VariableLengthStringFieldFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":0,\"F\":["...)
	w.Uint16(m.Offset())
	w.Data = append(w.Data, ',')
	w.Uint16(m.Length())
	return w.Finish(), nil
}

func (m VariableLengthStringFieldFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":0,\"F\":["...)
	w.Uint16(m.Offset())
	w.Data = append(w.Data, ',')
	w.Uint16(m.Length())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m VariableLengthStringFieldFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 0)
	w.Uint16Field("Offset", m.Offset())
	w.Uint16Field("Length", m.Length())
	return w.Finish(), nil
}

func (m VariableLengthStringFieldFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 0)
	w.Uint16Field("Offset", m.Offset())
	w.Uint16Field("Length", m.Length())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *VariableLengthStringFieldFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 0 {
		return message.ErrWrongType
	}
	m.SetOffset(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetLength(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *VariableLengthStringFieldFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
		case "Offset":
			m.SetOffset(r.Uint16())
		case "Length":
			m.SetLength(r.Uint16())
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
