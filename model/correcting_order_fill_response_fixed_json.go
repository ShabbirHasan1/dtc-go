//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CorrectingOrderFillResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":310,\"F\":["...)
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsError())
	return w.Finish(), nil
}

func (m CorrectingOrderFillResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":310,\"F\":["...)
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CorrectingOrderFillResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 310)
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("ResultText", m.ResultText())
	w.Uint8Field("IsError", m.IsError())
	return w.Finish(), nil
}

func (m CorrectingOrderFillResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 310)
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("ResultText", m.ResultText())
	w.Uint8Field("IsError", m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 310 {
		return message.ErrWrongType
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetResultText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsError(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 310 {
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
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "ResultText":
			m.SetResultText(r.String())
		case "IsError":
			m.SetIsError(r.Uint8())
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
