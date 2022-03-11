//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeCompact) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":112,\"F\":["...)
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeCompactBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":112,\"F\":["...)
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeCompact) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 112)
	w.Float32Field("Price", m.Price())
	w.Float32Field("Volume", m.Volume())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeCompactBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 112)
	w.Float32Field("Price", m.Price())
	w.Float32Field("Volume", m.Volume())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeCompactBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 112 {
		return message.ErrWrongType
	}
	m.SetPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeCompactBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 112 {
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
		case "Price":
			m.SetPrice(r.Float32())
		case "Volume":
			m.SetVolume(r.Float32())
		case "DateTime":
			m.SetDateTime(DateTime4Byte(r.Uint32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
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
