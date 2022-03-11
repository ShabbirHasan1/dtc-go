//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersModifyPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":153,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.PriorPrice())
	w.Data = append(w.Data, ',')
	w.Uint32(m.PriorQuantity())
	w.Data = append(w.Data, ',')
	w.Uint64(m.PriorOrderID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketOrdersModifyPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":153,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.PriorPrice())
	w.Data = append(w.Data, ',')
	w.Uint32(m.PriorQuantity())
	w.Data = append(w.Data, ',')
	w.Uint64(m.PriorOrderID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersModifyPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 153)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint32Field("Quantity", m.Quantity())
	w.Float64Field("Price", m.Price())
	w.Uint64Field("OrderID", m.OrderID())
	w.Float64Field("PriorPrice", m.PriorPrice())
	w.Uint32Field("PriorQuantity", m.PriorQuantity())
	w.Uint64Field("PriorOrderID", m.PriorOrderID())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketOrdersModifyPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 153)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint32Field("Quantity", m.Quantity())
	w.Float64Field("Price", m.Price())
	w.Uint64Field("OrderID", m.OrderID())
	w.Float64Field("PriorPrice", m.PriorPrice())
	w.Uint32Field("PriorQuantity", m.PriorQuantity())
	w.Uint64Field("PriorOrderID", m.PriorOrderID())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersModifyPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 153 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriorPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriorQuantity(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriorOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersModifyPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 153 {
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
		case "Quantity":
			m.SetQuantity(r.Uint32())
		case "Price":
			m.SetPrice(r.Float64())
		case "OrderID":
			m.SetOrderID(r.Uint64())
		case "PriorPrice":
			m.SetPriorPrice(r.Float64())
		case "PriorQuantity":
			m.SetPriorQuantity(r.Uint32())
		case "PriorOrderID":
			m.SetPriorOrderID(r.Uint64())
		case "DateTime":
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
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
