//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserInformationPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10136,\"F\":["...)
	w.String(m.OperatorID())
	w.Data = append(w.Data, ',')
	w.String(m.LocationID())
	return w.Finish(), nil
}

func (m UserInformationPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10136,\"F\":["...)
	w.String(m.OperatorID())
	w.Data = append(w.Data, ',')
	w.String(m.LocationID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserInformationPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10136)
	w.StringField("OperatorID", m.OperatorID())
	w.StringField("LocationID", m.LocationID())
	return w.Finish(), nil
}

func (m UserInformationPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10136)
	w.StringField("OperatorID", m.OperatorID())
	w.StringField("LocationID", m.LocationID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserInformationPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10136 {
		return message.ErrWrongType
	}
	m.SetOperatorID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetLocationID(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserInformationPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10136 {
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
		case "OperatorID":
			m.SetOperatorID(r.String())
		case "LocationID":
			m.SetLocationID(r.String())
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
