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
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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

func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) MarshalJSONTo(b []byte) ([]byte, error) {
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

func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
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
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevelFloatWithMillisecondsBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDateTime(DateTimeWithMillisecondsInt(r.ReadInt64()))
		case 3:
			m.SetPrice(r.ReadFloat32())
		case 4:
			m.SetQuantity(r.ReadFloat32())
		case 5:
			m.SetSide(AtBidOrAskEnum8(r.ReadUint8()))
		case 6:
			m.SetUpdateType(MarketDepthUpdateTypeEnum(r.ReadUint8()))
		case 7:
			m.SetNumOrders(r.ReadUint16())
		case 8:
			m.SetFinalUpdateInBatch(FinalUpdateInBatchEnum(r.ReadUint8()))
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

func (m *MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDateTime(DateTimeWithMillisecondsInt(r.ReadInt64()))
		case 3:
			m.SetPrice(r.ReadFloat32())
		case 4:
			m.SetQuantity(r.ReadFloat32())
		case 5:
			m.SetSide(AtBidOrAskEnum8(r.ReadUint8()))
		case 6:
			m.SetUpdateType(MarketDepthUpdateTypeEnum(r.ReadUint8()))
		case 7:
			m.SetNumOrders(r.ReadUint16())
		case 8:
			m.SetFinalUpdateInBatch(FinalUpdateInBatchEnum(r.ReadUint8()))
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

func (m MarketDepthUpdateLevelFloatWithMillisecondsBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 140)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint64(2, int64(m.DateTime()))
	w.WriteFixed32Float32(3, m.Price())
	w.WriteFixed32Float32(4, m.Quantity())
	w.WriteUvarint8(5, uint8(m.Side()))
	w.WriteUvarint8(6, uint8(m.UpdateType()))
	w.WriteUvarint16(7, m.NumOrders())
	w.WriteUvarint8(8, uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 140)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint64(2, int64(m.DateTime()))
	w.WriteFixed32Float32(3, m.Price())
	w.WriteFixed32Float32(4, m.Quantity())
	w.WriteUvarint8(5, uint8(m.Side()))
	w.WriteUvarint8(6, uint8(m.UpdateType()))
	w.WriteUvarint16(7, m.NumOrders())
	w.WriteUvarint8(8, uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}
