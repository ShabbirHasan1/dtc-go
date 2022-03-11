//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ClientDeviceUpdate) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10139,\"F\":["...)
	w.Int64(m.ClientDeviceIdentifier())
	return w.Finish(), nil
}

func (m ClientDeviceUpdateBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10139,\"F\":["...)
	w.Int64(m.ClientDeviceIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ClientDeviceUpdate) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10139)
	w.Int64Field("ClientDeviceIdentifier", m.ClientDeviceIdentifier())
	return w.Finish(), nil
}

func (m ClientDeviceUpdateBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10139)
	w.Int64Field("ClientDeviceIdentifier", m.ClientDeviceIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ClientDeviceUpdateBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10139 {
		return message.ErrWrongType
	}
	m.SetClientDeviceIdentifier(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ClientDeviceUpdateBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10139 {
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
		case "ClientDeviceIdentifier":
			m.SetClientDeviceIdentifier(r.Int64())
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
