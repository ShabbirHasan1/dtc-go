//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataResponseTrailerPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10130,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataResponseTrailerPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10130,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataResponseTrailerPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10130)
	w.Uint32Field("RequestID", m.RequestID())
	w.Uint8Field("m_IsSnapshot", m.IsSnapshot())
	w.Uint8Field("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataResponseTrailerPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10130)
	w.Uint32Field("RequestID", m.RequestID())
	w.Uint8Field("m_IsSnapshot", m.IsSnapshot())
	w.Uint8Field("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
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
	m.SetIsSnapshot(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFirstMessageInBatch(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Uint8())
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
			m.SetIsSnapshot(r.Uint8())
		case "m_IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Uint8())
		case "m_IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Uint8())
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
