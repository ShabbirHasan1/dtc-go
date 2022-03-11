//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogoffFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":5,\"F\":["...)
	w.String(m.Reason())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotReconnect())
	return w.Finish(), nil
}

func (m LogoffFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":5,\"F\":["...)
	w.String(m.Reason())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogoffFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 5)
	w.StringField("Reason", m.Reason())
	w.Uint8Field("DoNotReconnect", m.DoNotReconnect())
	return w.Finish(), nil
}

func (m LogoffFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 5)
	w.StringField("Reason", m.Reason())
	w.Uint8Field("DoNotReconnect", m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogoffFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 5 {
		return message.ErrWrongType
	}
	m.SetReason(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDoNotReconnect(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogoffFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 5 {
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
		case "Reason":
			m.SetReason(r.String())
		case "DoNotReconnect":
			m.SetDoNotReconnect(r.Uint8())
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
