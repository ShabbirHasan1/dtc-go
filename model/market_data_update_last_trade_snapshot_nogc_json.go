//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateLastTradeSnapshotPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":134,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastTradePrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastTradeVolume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.LastTradeDateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":134,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastTradePrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastTradeVolume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.LastTradeDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateLastTradeSnapshotPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 134)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("LastTradePrice", m.LastTradePrice())
	w.Float64Field("LastTradeVolume", m.LastTradeVolume())
	w.Float64Field("LastTradeDateTime", float64(m.LastTradeDateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 134)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("LastTradePrice", m.LastTradePrice())
	w.Float64Field("LastTradeVolume", m.LastTradeVolume())
	w.Float64Field("LastTradeDateTime", float64(m.LastTradeDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateLastTradeSnapshotPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 134 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTradePrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTradeVolume(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateLastTradeSnapshotPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 134 {
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
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "LastTradePrice":
			m.SetLastTradePrice(r.Float64())
		case "LastTradeVolume":
			m.SetLastTradeVolume(r.Float64())
		case "LastTradeDateTime":
			m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.Float64()))
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
