//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthSnapshotLevel_IntPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":132,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.Side()))
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Int32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint16(m.Level())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	return w.Finish(), nil
}

func (m MarketDepthSnapshotLevel_IntPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":132,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.Side()))
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Int32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint16(m.Level())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthSnapshotLevel_IntPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 132)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Int32Field("Price", m.Price())
	w.Int32Field("Quantity", m.Quantity())
	w.Uint16Field("Level", m.Level())
	w.Uint8Field("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint32Field("NumOrders", m.NumOrders())
	return w.Finish(), nil
}

func (m MarketDepthSnapshotLevel_IntPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 132)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Int32Field("Price", m.Price())
	w.Int32Field("Quantity", m.Quantity())
	w.Uint16Field("Level", m.Level())
	w.Uint8Field("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint32Field("NumOrders", m.NumOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthSnapshotLevel_IntPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 132 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSide(AtBidOrAskEnum(r.Uint16()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLevel(r.Uint16())
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
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetNumOrders(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthSnapshotLevel_IntPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 132 {
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
		case "Side":
			m.SetSide(AtBidOrAskEnum(r.Uint16()))
		case "Price":
			m.SetPrice(r.Int32())
		case "Quantity":
			m.SetQuantity(r.Int32())
		case "Level":
			m.SetLevel(r.Uint16())
		case "IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Uint8())
		case "IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Uint8())
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "NumOrders":
			m.SetNumOrders(r.Uint32())
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
