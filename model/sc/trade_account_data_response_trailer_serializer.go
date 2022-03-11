//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataResponseTrailer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10130,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataResponseTrailerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10130,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataResponseTrailerPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10130,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataResponseTrailerPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10130,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataResponseTrailer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10130)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("m_IsSnapshot", m.IsSnapshot())
	w.BoolField("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataResponseTrailerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10130)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("m_IsSnapshot", m.IsSnapshot())
	w.BoolField("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataResponseTrailerPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10130)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("m_IsSnapshot", m.IsSnapshot())
	w.BoolField("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataResponseTrailerPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10130)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("m_IsSnapshot", m.IsSnapshot())
	w.BoolField("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataResponseTrailerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10130 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSnapshot(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFirstMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataResponseTrailerPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10130 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSnapshot(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFirstMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataResponseTrailerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10130 {
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
		case "m_IsSnapshot":
			m.SetIsSnapshot(r.Bool())
		case "m_IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Bool())
		case "m_IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Bool())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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

func (m *TradeAccountDataResponseTrailerPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10130 {
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
		case "m_IsSnapshot":
			m.SetIsSnapshot(r.Bool())
		case "m_IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Bool())
		case "m_IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Bool())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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