//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionVolumePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":113,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFinalSessionVolume())
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionVolumePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":113,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFinalSessionVolume())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionVolumePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 113)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("Volume", m.Volume())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	w.Uint8Field("IsFinalSessionVolume", m.IsFinalSessionVolume())
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionVolumePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 113)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("Volume", m.Volume())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	w.Uint8Field("IsFinalSessionVolume", m.IsFinalSessionVolume())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateSessionVolumePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 113 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalSessionVolume(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateSessionVolumePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 113 {
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
		case "Volume":
			m.SetVolume(r.Float64())
		case "TradingSessionDate":
			m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
		case "IsFinalSessionVolume":
			m.SetIsFinalSessionVolume(r.Uint8())
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
