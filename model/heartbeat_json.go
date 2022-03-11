//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m Heartbeat) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":3,\"F\":["...)
	w.Uint32(m.NumDroppedMessages())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.CurrentDateTime()))
	return w.Finish(), nil
}

func (m HeartbeatBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":3,\"F\":["...)
	w.Uint32(m.NumDroppedMessages())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.CurrentDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m Heartbeat) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 3)
	w.Uint32Field("NumDroppedMessages", m.NumDroppedMessages())
	w.Int64Field("CurrentDateTime", int64(m.CurrentDateTime()))
	return w.Finish(), nil
}

func (m HeartbeatBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 3)
	w.Uint32Field("NumDroppedMessages", m.NumDroppedMessages())
	w.Int64Field("CurrentDateTime", int64(m.CurrentDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HeartbeatBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 3 {
		return message.ErrWrongType
	}
	m.SetNumDroppedMessages(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HeartbeatBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 3 {
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
		case "NumDroppedMessages":
			m.SetNumDroppedMessages(r.Uint32())
		case "CurrentDateTime":
			m.SetCurrentDateTime(DateTime(r.Int64()))
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
