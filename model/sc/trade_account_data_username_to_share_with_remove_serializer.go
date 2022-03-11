//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUsernameToShareWithRemove) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10129,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsReadOnly())
	return w.Finish(), nil
}

func (m TradeAccountDataUsernameToShareWithRemoveBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10129,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsReadOnly())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUsernameToShareWithRemovePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10129,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsReadOnly())
	return w.Finish(), nil
}

func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10129,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsReadOnly())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUsernameToShareWithRemove) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10129)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Username", m.Username())
	w.BoolField("IsReadOnly", m.IsReadOnly())
	return w.Finish(), nil
}

func (m TradeAccountDataUsernameToShareWithRemoveBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10129)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Username", m.Username())
	w.BoolField("IsReadOnly", m.IsReadOnly())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUsernameToShareWithRemovePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10129)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Username", m.Username())
	w.BoolField("IsReadOnly", m.IsReadOnly())
	return w.Finish(), nil
}

func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10129)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Username", m.Username())
	w.BoolField("IsReadOnly", m.IsReadOnly())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUsernameToShareWithRemoveBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10129 {
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
	m.SetIsReadOnly(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUsernameToShareWithRemovePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10129 {
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
	m.SetIsReadOnly(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUsernameToShareWithRemoveBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10129 {
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
		case "IsReadOnly":
			m.SetIsReadOnly(r.Bool())
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

func (m *TradeAccountDataUsernameToShareWithRemovePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10129 {
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
		case "IsReadOnly":
			m.SetIsReadOnly(r.Bool())
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
