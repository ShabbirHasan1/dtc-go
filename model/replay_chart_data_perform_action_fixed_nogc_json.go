//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ReplayChartDataPerformActionFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10105,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Action()))
	w.Data = append(w.Data, ',')
	w.Float32(m.ReplaySpeed())
	return w.Finish(), nil
}

func (m ReplayChartDataPerformActionFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10105,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Action()))
	w.Data = append(w.Data, ',')
	w.Float32(m.ReplaySpeed())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ReplayChartDataPerformActionFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10105)
	w.Uint32Field("RequestID", m.RequestID())
	w.Int32Field("Action", int32(m.Action()))
	w.Float32Field("ReplaySpeed", m.ReplaySpeed())
	return w.Finish(), nil
}

func (m ReplayChartDataPerformActionFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10105)
	w.Uint32Field("RequestID", m.RequestID())
	w.Int32Field("Action", int32(m.Action()))
	w.Float32Field("ReplaySpeed", m.ReplaySpeed())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ReplayChartDataPerformActionFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10105 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAction(ReplayChartDataActionEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetReplaySpeed(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ReplayChartDataPerformActionFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10105 {
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
		case "Action":
			m.SetAction(ReplayChartDataActionEnum(r.Int32()))
		case "ReplaySpeed":
			m.SetReplaySpeed(r.Float32())
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
