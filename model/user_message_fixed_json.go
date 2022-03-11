//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessageFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":700,\"F\":["...)
	w.String(m.UserMessage())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsPopupMessage())
	return w.Finish(), nil
}

func (m UserMessageFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":700,\"F\":["...)
	w.String(m.UserMessage())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessageFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 700)
	w.StringField("UserMessage", m.UserMessage())
	w.Uint8Field("IsPopupMessage", m.IsPopupMessage())
	return w.Finish(), nil
}

func (m UserMessageFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 700)
	w.StringField("UserMessage", m.UserMessage())
	w.Uint8Field("IsPopupMessage", m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessageFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 700 {
		return message.ErrWrongType
	}
	m.SetUserMessage(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsPopupMessage(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessageFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 700 {
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
		case "UserMessage":
			m.SetUserMessage(r.String())
		case "IsPopupMessage":
			m.SetIsPopupMessage(r.Uint8())
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
