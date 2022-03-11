//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataAuthorizedUsernameAdd) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10125,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	return w.Finish(), nil
}

func (m TradeAccountDataAuthorizedUsernameAddBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10125,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataAuthorizedUsernameAddPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10125,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	return w.Finish(), nil
}

func (m TradeAccountDataAuthorizedUsernameAddPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10125,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataAuthorizedUsernameAdd) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10125)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Username", m.Username())
	return w.Finish(), nil
}

func (m TradeAccountDataAuthorizedUsernameAddBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10125)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Username", m.Username())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataAuthorizedUsernameAddPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10125)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Username", m.Username())
	return w.Finish(), nil
}

func (m TradeAccountDataAuthorizedUsernameAddPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10125)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Username", m.Username())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataAuthorizedUsernameAddBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10125 {
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
	m.SetUsername(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataAuthorizedUsernameAddPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10125 {
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
	m.SetUsername(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataAuthorizedUsernameAddBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10125 {
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
		case "Username":
			m.SetUsername(r.String())
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

func (m *TradeAccountDataAuthorizedUsernameAddPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10125 {
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
		case "Username":
			m.SetUsername(r.String())
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
