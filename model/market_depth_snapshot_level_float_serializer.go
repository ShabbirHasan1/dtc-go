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

func (m MarketDepthSnapshotLevelFloat) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":145,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Uint16(m.Level())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.Side()))
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

func (m MarketDepthSnapshotLevelFloatBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":145,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Uint16(m.Level())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.Side()))
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthSnapshotLevelFloatPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":145,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Uint16(m.Level())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.Side()))
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

func (m MarketDepthSnapshotLevelFloatPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":145,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Uint16(m.Level())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.Side()))
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthSnapshotLevelFloat) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 145)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	w.Uint32Field("NumOrders", m.NumOrders())
	w.Uint16Field("Level", m.Level())
	w.Uint8Field("Side", uint8(m.Side()))
	w.Uint8Field("FinalUpdateInBatch", uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

func (m MarketDepthSnapshotLevelFloatBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 145)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	w.Uint32Field("NumOrders", m.NumOrders())
	w.Uint16Field("Level", m.Level())
	w.Uint8Field("Side", uint8(m.Side()))
	w.Uint8Field("FinalUpdateInBatch", uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthSnapshotLevelFloatPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 145)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	w.Uint32Field("NumOrders", m.NumOrders())
	w.Uint16Field("Level", m.Level())
	w.Uint8Field("Side", uint8(m.Side()))
	w.Uint8Field("FinalUpdateInBatch", uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

func (m MarketDepthSnapshotLevelFloatPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 145)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Float32Field("Quantity", m.Quantity())
	w.Uint32Field("NumOrders", m.NumOrders())
	w.Uint16Field("Level", m.Level())
	w.Uint8Field("Side", uint8(m.Side()))
	w.Uint8Field("FinalUpdateInBatch", uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthSnapshotLevelFloatBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 145 {
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
	m.SetNumOrders(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLevel(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetSide(AtBidOrAskEnum8(r.Uint8()))
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

func (m *MarketDepthSnapshotLevelFloatPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 145 {
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
	m.SetNumOrders(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLevel(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetSide(AtBidOrAskEnum8(r.Uint8()))
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

func (m *MarketDepthSnapshotLevelFloatBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 145 {
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
			m.SetNumOrders(r.Uint32())
		case "Level":
			m.SetLevel(r.Uint16())
		case "Side":
			m.SetSide(AtBidOrAskEnum8(r.Uint8()))
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

func (m *MarketDepthSnapshotLevelFloatPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 145 {
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
			m.SetNumOrders(r.Uint32())
		case "Level":
			m.SetLevel(r.Uint16())
		case "Side":
			m.SetSide(AtBidOrAskEnum8(r.Uint8()))
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

func (m *MarketDepthSnapshotLevelFloatBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetNumOrders(r.ReadUint32())
		case 5:
			m.SetLevel(r.ReadUint16())
		case 6:
			m.SetSide(AtBidOrAskEnum8(r.ReadUint8()))
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

func (m *MarketDepthSnapshotLevelFloatPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetNumOrders(r.ReadUint32())
		case 5:
			m.SetLevel(r.ReadUint16())
		case 6:
			m.SetSide(AtBidOrAskEnum8(r.ReadUint8()))
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

func (m MarketDepthSnapshotLevelFloatBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 145)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.Price())
	w.WriteFixed32Float32(3, m.Quantity())
	w.WriteUvarint32(4, m.NumOrders())
	w.WriteUvarint16(5, m.Level())
	w.WriteUvarint8(6, uint8(m.Side()))
	w.WriteUvarint8(7, uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthSnapshotLevelFloatPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 145)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.Price())
	w.WriteFixed32Float32(3, m.Quantity())
	w.WriteUvarint32(4, m.NumOrders())
	w.WriteUvarint16(5, m.Level())
	w.WriteUvarint8(6, uint8(m.Side()))
	w.WriteUvarint8(7, uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}