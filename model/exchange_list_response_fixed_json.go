//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ExchangeListResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":501,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	return w.Finish(), nil
}

func (m ExchangeListResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":501,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ExchangeListResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 501)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Uint8Field("IsFinalMessage", m.IsFinalMessage())
	w.StringField("Description", m.Description())
	return w.Finish(), nil
}

func (m ExchangeListResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 501)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Uint8Field("IsFinalMessage", m.IsFinalMessage())
	w.StringField("Description", m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ExchangeListResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 501 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalMessage(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ExchangeListResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 501 {
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
			m.SetRequestID(r.Int32())
		case "Exchange":
			m.SetExchange(r.String())
		case "IsFinalMessage":
			m.SetIsFinalMessage(r.Uint8())
		case "Description":
			m.SetDescription(r.String())
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
