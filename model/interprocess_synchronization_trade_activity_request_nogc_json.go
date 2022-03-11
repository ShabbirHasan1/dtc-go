//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m InterprocessSynchronizationTradeActivityRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10137,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(m.StartDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SendOrderActivityOnly())
	return w.Finish(), nil
}

func (m InterprocessSynchronizationTradeActivityRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10137,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(m.StartDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SendOrderActivityOnly())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m InterprocessSynchronizationTradeActivityRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10137)
	w.Uint32Field("RequestID", m.RequestID())
	w.Int64Field("StartDateTimeUTC", m.StartDateTimeUTC())
	w.Uint8Field("SendOrderActivityOnly", m.SendOrderActivityOnly())
	return w.Finish(), nil
}

func (m InterprocessSynchronizationTradeActivityRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10137)
	w.Uint32Field("RequestID", m.RequestID())
	w.Int64Field("StartDateTimeUTC", m.StartDateTimeUTC())
	w.Uint8Field("SendOrderActivityOnly", m.SendOrderActivityOnly())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *InterprocessSynchronizationTradeActivityRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10137 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetSendOrderActivityOnly(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *InterprocessSynchronizationTradeActivityRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10137 {
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
		case "StartDateTimeUTC":
			m.SetStartDateTimeUTC(r.Int64())
		case "SendOrderActivityOnly":
			m.SetSendOrderActivityOnly(r.Uint8())
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
