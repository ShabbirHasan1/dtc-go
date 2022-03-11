//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":146,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradeCondition()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":146,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradeCondition()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 146)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Int8Field("TradeCondition", int8(m.TradeCondition()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 146)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Int8Field("TradeCondition", int8(m.TradeCondition()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 146 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum8(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeCondition(TradeConditionEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 146 {
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
		case "Price":
			m.SetPrice(r.Float32())
		case "Volume":
			m.SetVolume(r.Uint32())
		case "DateTime":
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.Uint8()))
		case "UnbundledTradeIndicator":
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.Int8()))
		case "TradeCondition":
			m.SetTradeCondition(TradeConditionEnum(r.Int8()))
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
