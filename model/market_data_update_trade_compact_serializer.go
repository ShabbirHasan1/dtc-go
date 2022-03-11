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

func (m MarketDataUpdateTradeCompact) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":112,\"F\":["...)
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeCompactBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":112,\"F\":["...)
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeCompactPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":112,\"F\":["...)
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeCompactPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":112,\"F\":["...)
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Float32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeCompact) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 112)
	w.Float32Field("Price", m.Price())
	w.Float32Field("Volume", m.Volume())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeCompactBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 112)
	w.Float32Field("Price", m.Price())
	w.Float32Field("Volume", m.Volume())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeCompactPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 112)
	w.Float32Field("Price", m.Price())
	w.Float32Field("Volume", m.Volume())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeCompactPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 112)
	w.Float32Field("Price", m.Price())
	w.Float32Field("Volume", m.Volume())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeCompactBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 112 {
		return message.ErrWrongType
	}
	m.SetPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeCompactPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 112 {
		return message.ErrWrongType
	}
	m.SetPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeCompactBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 112 {
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
		case "Price":
			m.SetPrice(r.Float32())
		case "Volume":
			m.SetVolume(r.Float32())
		case "DateTime":
			m.SetDateTime(DateTime4Byte(r.Uint32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
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

func (m *MarketDataUpdateTradeCompactPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 112 {
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
		case "Price":
			m.SetPrice(r.Float32())
		case "Volume":
			m.SetVolume(r.Float32())
		case "DateTime":
			m.SetDateTime(DateTime4Byte(r.Uint32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
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

func (m *MarketDataUpdateTradeCompactBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat32())
		case 2:
			m.SetVolume(r.ReadFloat32())
		case 3:
			m.SetDateTime(DateTime4Byte(r.ReadUint32()))
		case 4:
			m.SetSymbolID(r.ReadUint32())
		case 5:
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.ReadUint16()))
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

func (m *MarketDataUpdateTradeCompactPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat32())
		case 2:
			m.SetVolume(r.ReadFloat32())
		case 3:
			m.SetDateTime(DateTime4Byte(r.ReadUint32()))
		case 4:
			m.SetSymbolID(r.ReadUint32())
		case 5:
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.ReadUint16()))
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

func (m MarketDataUpdateTradeCompactBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 112)
	w.WriteFixed32Float32(1, m.Price())
	w.WriteFixed32Float32(2, m.Volume())
	w.WriteFixed32Uint32(3, uint32(m.DateTime()))
	w.WriteUvarint32(4, m.SymbolID())
	w.WriteUvarint16(5, uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeCompactPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 112)
	w.WriteFixed32Float32(1, m.Price())
	w.WriteFixed32Float32(2, m.Volume())
	w.WriteFixed32Uint32(3, uint32(m.DateTime()))
	w.WriteUvarint32(4, m.SymbolID())
	w.WriteUvarint16(5, uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}
