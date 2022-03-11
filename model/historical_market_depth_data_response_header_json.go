//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataResponseHeader) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":901,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoRecordsToReturn())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataResponseHeaderBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":901,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoRecordsToReturn())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataResponseHeader) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 901)
	w.Int32Field("RequestID", m.RequestID())
	w.Uint8Field("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("NoRecordsToReturn", m.NoRecordsToReturn())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataResponseHeaderBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 901)
	w.Int32Field("RequestID", m.RequestID())
	w.Uint8Field("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("NoRecordsToReturn", m.NoRecordsToReturn())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataResponseHeaderBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 901 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseZLibCompression(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoRecordsToReturn(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataResponseHeaderBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 901 {
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
		case "UseZLibCompression":
			m.SetUseZLibCompression(r.Uint8())
		case "NoRecordsToReturn":
			m.SetNoRecordsToReturn(r.Uint8())
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
