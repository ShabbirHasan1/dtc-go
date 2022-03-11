//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersAdd) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":152,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Side()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OrderID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketOrdersAddBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":152,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Side()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OrderID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersAddPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":152,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Side()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OrderID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketOrdersAddPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":152,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Side()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OrderID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersAdd) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 152)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("Side", int32(m.Side()))
	w.Uint32Field("Quantity", m.Quantity())
	w.Float64Field("Price", m.Price())
	w.Uint64Field("OrderID", m.OrderID())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketOrdersAddBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 152)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("Side", int32(m.Side()))
	w.Uint32Field("Quantity", m.Quantity())
	w.Float64Field("Price", m.Price())
	w.Uint64Field("OrderID", m.OrderID())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersAddPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 152)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("Side", int32(m.Side()))
	w.Uint32Field("Quantity", m.Quantity())
	w.Float64Field("Price", m.Price())
	w.Uint64Field("OrderID", m.OrderID())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketOrdersAddPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 152)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("Side", int32(m.Side()))
	w.Uint32Field("Quantity", m.Quantity())
	w.Float64Field("Price", m.Price())
	w.Uint64Field("OrderID", m.OrderID())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersAddBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 152 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSide(BuySellEnum(r.Int32()))
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
	m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersAddPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 152 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSide(BuySellEnum(r.Int32()))
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
	m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersAddBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 152 {
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
			m.SetSide(BuySellEnum(r.Int32()))
		case "Quantity":
			m.SetQuantity(r.Uint32())
		case "Price":
			m.SetPrice(r.Float64())
		case "OrderID":
			m.SetOrderID(r.Uint64())
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersAddPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 152 {
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
			m.SetSide(BuySellEnum(r.Int32()))
		case "Quantity":
			m.SetQuantity(r.Uint32())
		case "Price":
			m.SetPrice(r.Float64())
		case "OrderID":
			m.SetOrderID(r.Uint64())
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

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersAddBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetSymbolID(r.ReadUint32())
		case 2:
			m.SetSide(BuySellEnum(r.ReadInt32()))
		case 3:
			m.SetPrice(r.ReadFloat64())
		case 4:
			m.SetQuantity(r.ReadUint32())
		case 5:
			m.SetOrderID(r.ReadUint64())
		case 6:
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersAddPointerBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetSymbolID(r.ReadUint32())
		case 2:
			m.SetSide(BuySellEnum(r.ReadInt32()))
		case 3:
			m.SetPrice(r.ReadFloat64())
		case 4:
			m.SetQuantity(r.ReadUint32())
		case 5:
			m.SetOrderID(r.ReadUint64())
		case 6:
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersAddBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 152)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint32(2, int32(m.Side()))
	w.WriteFixed64Float64(3, m.Price())
	w.WriteUvarint32(4, m.Quantity())
	w.WriteUvarint64(5, m.OrderID())
	w.WriteVarint64(6, int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersAddPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 152)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint32(2, int32(m.Side()))
	w.WriteFixed64Float64(3, m.Price())
	w.WriteUvarint32(4, m.Quantity())
	w.WriteUvarint64(5, m.OrderID())
	w.WriteVarint64(6, int64(m.DateTime()))
	return w.Finish(), nil
}
