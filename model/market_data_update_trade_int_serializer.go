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

func (m MarketDataUpdateTrade_Int) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":126,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Int32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTrade_IntBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":126,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Int32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTrade_IntPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":126,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Int32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTrade_IntPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":126,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Int32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTrade_Int) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 126)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Int32Field("Price", m.Price())
	w.Int32Field("Volume", m.Volume())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTrade_IntBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 126)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Int32Field("Price", m.Price())
	w.Int32Field("Volume", m.Volume())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTrade_IntPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 126)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Int32Field("Price", m.Price())
	w.Int32Field("Volume", m.Volume())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTrade_IntPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 126)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Int32Field("Price", m.Price())
	w.Int32Field("Volume", m.Volume())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTrade_IntBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 126 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTrade_IntPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 126 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTrade_IntBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 126 {
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
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
		case "Price":
			m.SetPrice(r.Int32())
		case "Volume":
			m.SetVolume(r.Int32())
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
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

func (m *MarketDataUpdateTrade_IntPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 126 {
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
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
		case "Price":
			m.SetPrice(r.Int32())
		case "Volume":
			m.SetVolume(r.Int32())
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
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

func (m *MarketDataUpdateTrade_IntBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.ReadUint16()))
		case 3:
			m.SetPrice(r.ReadInt32())
		case 4:
			m.SetVolume(r.ReadInt32())
		case 5:
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
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

func (m *MarketDataUpdateTrade_IntPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.ReadUint16()))
		case 3:
			m.SetPrice(r.ReadInt32())
		case 4:
			m.SetVolume(r.ReadInt32())
		case 5:
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
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

func (m MarketDataUpdateTrade_IntBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 126)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint16(2, uint16(m.AtBidOrAsk()))
	w.WriteVarint32(3, m.Price())
	w.WriteVarint32(4, m.Volume())
	w.WriteFixed64Float64(5, float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTrade_IntPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 126)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint16(2, uint16(m.AtBidOrAsk()))
	w.WriteVarint32(3, m.Price())
	w.WriteVarint32(4, m.Volume())
	w.WriteFixed64Float64(5, float64(m.DateTime()))
	return w.Finish(), nil
}
