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

func (m MarketDepthSnapshotLevel_Int) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	return w.Finish(), nil
}

func (m MarketDepthSnapshotLevel_IntBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	return w.Finish(), nil
}

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
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
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
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthSnapshotLevel_Int) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 132)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Int32Field("Price", m.Price())
	w.Int32Field("Quantity", m.Quantity())
	w.Uint16Field("Level", m.Level())
	w.BoolField("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint32Field("NumOrders", m.NumOrders())
	return w.Finish(), nil
}

func (m MarketDepthSnapshotLevel_IntBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 132)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Int32Field("Price", m.Price())
	w.Int32Field("Quantity", m.Quantity())
	w.Uint16Field("Level", m.Level())
	w.BoolField("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint32Field("NumOrders", m.NumOrders())
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
	w.BoolField("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("IsLastMessageInBatch", m.IsLastMessageInBatch())
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
	w.BoolField("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint32Field("NumOrders", m.NumOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthSnapshotLevel_IntBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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
	m.SetIsFirstMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Bool())
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
	m.SetIsFirstMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Bool())
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

func (m *MarketDepthSnapshotLevel_IntBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
			m.SetIsFirstMessageInBatch(r.Bool())
		case "IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Bool())
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
			m.SetIsFirstMessageInBatch(r.Bool())
		case "IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Bool())
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

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthSnapshotLevel_IntBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSide(AtBidOrAskEnum(r.ReadUint16()))
		case 3:
			m.SetPrice(r.ReadInt32())
		case 4:
			m.SetQuantity(r.ReadInt32())
		case 5:
			m.SetLevel(r.ReadUint16())
		case 6:
			m.SetIsFirstMessageInBatch(r.ReadBool())
		case 7:
			m.SetIsLastMessageInBatch(r.ReadBool())
		case 8:
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 9:
			m.SetNumOrders(r.ReadUint32())
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

func (m *MarketDepthSnapshotLevel_IntPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSide(AtBidOrAskEnum(r.ReadUint16()))
		case 3:
			m.SetPrice(r.ReadInt32())
		case 4:
			m.SetQuantity(r.ReadInt32())
		case 5:
			m.SetLevel(r.ReadUint16())
		case 6:
			m.SetIsFirstMessageInBatch(r.ReadBool())
		case 7:
			m.SetIsLastMessageInBatch(r.ReadBool())
		case 8:
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 9:
			m.SetNumOrders(r.ReadUint32())
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

func (m MarketDepthSnapshotLevel_IntBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 132)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint16(2, uint16(m.Side()))
	w.WriteVarint32(3, m.Price())
	w.WriteVarint32(4, m.Quantity())
	w.WriteUvarint16(5, m.Level())
	w.WriteBool(6, m.IsFirstMessageInBatch())
	w.WriteBool(7, m.IsLastMessageInBatch())
	w.WriteFixed64Float64(8, float64(m.DateTime()))
	w.WriteUvarint32(9, m.NumOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthSnapshotLevel_IntPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 132)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint16(2, uint16(m.Side()))
	w.WriteVarint32(3, m.Price())
	w.WriteVarint32(4, m.Quantity())
	w.WriteUvarint16(5, m.Level())
	w.WriteBool(6, m.IsFirstMessageInBatch())
	w.WriteBool(7, m.IsLastMessageInBatch())
	w.WriteFixed64Float64(8, float64(m.DateTime()))
	w.WriteUvarint32(9, m.NumOrders())
	return w.Finish(), nil
}
