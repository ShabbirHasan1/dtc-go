//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevelFloatWithMilliseconds) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":140,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.Side()))
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.UpdateType()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

func (m MarketDepthUpdateLevelFloatWithMillisecondsBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":140,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.Side()))
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.UpdateType()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevelFloatWithMilliseconds) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 140)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float32Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	w.Uint8Field("Side", uint8(m.Side()))
	w.Uint8Field("UpdateType", uint8(m.UpdateType()))
	w.Uint16Field("NumOrders", m.NumOrders())
	w.Uint8Field("FinalUpdateInBatch", uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

func (m MarketDepthUpdateLevelFloatWithMillisecondsBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 140)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float32Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	w.Uint8Field("Side", uint8(m.Side()))
	w.Uint8Field("UpdateType", uint8(m.UpdateType()))
	w.Uint16Field("NumOrders", m.NumOrders())
	w.Uint8Field("FinalUpdateInBatch", uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevelFloatWithMillisecondsBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 140 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMillisecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSide(AtBidOrAskEnum8(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdateType(MarketDepthUpdateTypeEnum(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetNumOrders(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetFinalUpdateInBatch(FinalUpdateInBatchEnum(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevelFloatWithMillisecondsBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 140 {
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
		case "DateTime":
			m.SetDateTime(DateTimeWithMillisecondsInt(r.Int64()))
		case "Price":
			m.SetPrice(r.Float32())
		case "Quantity":
			m.SetQuantity(r.Float32())
		case "Side":
			m.SetSide(AtBidOrAskEnum8(r.Uint8()))
		case "UpdateType":
			m.SetUpdateType(MarketDepthUpdateTypeEnum(r.Uint8()))
		case "NumOrders":
			m.SetNumOrders(r.Uint16())
		case "FinalUpdateInBatch":
			m.SetFinalUpdateInBatch(FinalUpdateInBatchEnum(r.Uint8()))
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
