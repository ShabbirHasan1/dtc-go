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

func (m MarketDepthSnapshotLevel) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":122,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.Side()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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

func (m MarketDepthSnapshotLevelBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":122,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.Side()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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

func (m MarketDepthSnapshotLevelPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":122,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.Side()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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

func (m MarketDepthSnapshotLevelPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":122,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.Side()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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

func (m MarketDepthSnapshotLevel) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 122)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Quantity", m.Quantity())
	w.Uint16Field("Level", m.Level())
	w.BoolField("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint32Field("NumOrders", m.NumOrders())
	return w.Finish(), nil
}

func (m MarketDepthSnapshotLevelBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 122)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Quantity", m.Quantity())
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

func (m MarketDepthSnapshotLevelPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 122)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Quantity", m.Quantity())
	w.Uint16Field("Level", m.Level())
	w.BoolField("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint32Field("NumOrders", m.NumOrders())
	return w.Finish(), nil
}

func (m MarketDepthSnapshotLevelPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 122)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Quantity", m.Quantity())
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

func (m *MarketDepthSnapshotLevelBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 122 {
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
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
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

func (m *MarketDepthSnapshotLevelPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 122 {
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
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
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

func (m *MarketDepthSnapshotLevelBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 122 {
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
			m.SetPrice(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
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

func (m *MarketDepthSnapshotLevelPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 122 {
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
			m.SetPrice(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
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

func (m *MarketDepthSnapshotLevelBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat64())
		case 4:
			m.SetQuantity(r.ReadFloat64())
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

func (m *MarketDepthSnapshotLevelPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat64())
		case 4:
			m.SetQuantity(r.ReadFloat64())
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

func (m MarketDepthSnapshotLevelBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 122)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint16(2, uint16(m.Side()))
	w.WriteFixed64Float64(3, m.Price())
	w.WriteFixed64Float64(4, m.Quantity())
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

func (m MarketDepthSnapshotLevelPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 122)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint16(2, uint16(m.Side()))
	w.WriteFixed64Float64(3, m.Price())
	w.WriteFixed64Float64(4, m.Quantity())
	w.WriteUvarint16(5, m.Level())
	w.WriteBool(6, m.IsFirstMessageInBatch())
	w.WriteBool(7, m.IsLastMessageInBatch())
	w.WriteFixed64Float64(8, float64(m.DateTime()))
	w.WriteUvarint32(9, m.NumOrders())
	return w.Finish(), nil
}
