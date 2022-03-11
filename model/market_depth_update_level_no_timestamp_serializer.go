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

func (m MarketDepthUpdateLevelNoTimestamp) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":141,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Int8(m.Side())
	w.Data = append(w.Data, ',')
	w.Int8(m.UpdateType())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

func (m MarketDepthUpdateLevelNoTimestampBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":141,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Int8(m.Side())
	w.Data = append(w.Data, ',')
	w.Int8(m.UpdateType())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevelNoTimestampPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":141,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Int8(m.Side())
	w.Data = append(w.Data, ',')
	w.Int8(m.UpdateType())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":141,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Int8(m.Side())
	w.Data = append(w.Data, ',')
	w.Int8(m.UpdateType())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevelNoTimestamp) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 141)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	w.Uint16Field("NumOrders", m.NumOrders())
	w.Int8Field("Side", m.Side())
	w.Int8Field("UpdateType", m.UpdateType())
	w.Uint8Field("FinalUpdateInBatch", uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

func (m MarketDepthUpdateLevelNoTimestampBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 141)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	w.Uint16Field("NumOrders", m.NumOrders())
	w.Int8Field("Side", m.Side())
	w.Int8Field("UpdateType", m.UpdateType())
	w.Uint8Field("FinalUpdateInBatch", uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevelNoTimestampPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 141)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	w.Uint16Field("NumOrders", m.NumOrders())
	w.Int8Field("Side", m.Side())
	w.Int8Field("UpdateType", m.UpdateType())
	w.Uint8Field("FinalUpdateInBatch", uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 141)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	w.Uint16Field("NumOrders", m.NumOrders())
	w.Int8Field("Side", m.Side())
	w.Int8Field("UpdateType", m.UpdateType())
	w.Uint8Field("FinalUpdateInBatch", uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevelNoTimestampBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 141 {
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
	m.SetQuantity(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumOrders(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetSide(r.Int8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdateType(r.Int8())
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

func (m *MarketDepthUpdateLevelNoTimestampPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 141 {
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
	m.SetQuantity(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumOrders(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetSide(r.Int8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdateType(r.Int8())
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

func (m *MarketDepthUpdateLevelNoTimestampBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 141 {
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
		case "Quantity":
			m.SetQuantity(r.Float32())
		case "NumOrders":
			m.SetNumOrders(r.Uint16())
		case "Side":
			m.SetSide(r.Int8())
		case "UpdateType":
			m.SetUpdateType(r.Int8())
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

func (m *MarketDepthUpdateLevelNoTimestampPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 141 {
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
		case "Quantity":
			m.SetQuantity(r.Float32())
		case "NumOrders":
			m.SetNumOrders(r.Uint16())
		case "Side":
			m.SetSide(r.Int8())
		case "UpdateType":
			m.SetUpdateType(r.Int8())
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

func (m *MarketDepthUpdateLevelNoTimestampBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat32())
		case 3:
			m.SetQuantity(r.ReadFloat32())
		case 4:
			m.SetNumOrders(r.ReadUint16())
		case 5:
			m.SetSide(r.ReadInt8())
		case 6:
			m.SetUpdateType(r.ReadInt8())
		case 7:
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

func (m *MarketDepthUpdateLevelNoTimestampPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat32())
		case 3:
			m.SetQuantity(r.ReadFloat32())
		case 4:
			m.SetNumOrders(r.ReadUint16())
		case 5:
			m.SetSide(r.ReadInt8())
		case 6:
			m.SetUpdateType(r.ReadInt8())
		case 7:
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

func (m MarketDepthUpdateLevelNoTimestampBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 141)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.Price())
	w.WriteFixed32Float32(3, m.Quantity())
	w.WriteUvarint16(4, m.NumOrders())
	w.WriteVarint8(5, m.Side())
	w.WriteVarint8(6, m.UpdateType())
	w.WriteUvarint8(7, uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 141)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.Price())
	w.WriteFixed32Float32(3, m.Quantity())
	w.WriteUvarint16(4, m.NumOrders())
	w.WriteVarint8(5, m.Side())
	w.WriteVarint8(6, m.UpdateType())
	w.WriteUvarint8(7, uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}
