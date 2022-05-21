//go:build !tinygo

// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-16 08:00:54.519198 +0800 WITA m=+0.055446543

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTrade) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":107,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":107,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":107,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":107,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTrade) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 107)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Volume", m.Volume())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 107)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Volume", m.Volume())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 107)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Volume", m.Volume())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 107)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Volume", m.Volume())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 107 {
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
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Float64())
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

func (m *MarketDataUpdateTradePointerBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 107 {
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
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Float64())
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

func (m *MarketDataUpdateTradeBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 107 {
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
			m.SetPrice(r.Float64())
		case "Volume":
			m.SetVolume(r.Float64())
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

func (m *MarketDataUpdateTradePointerBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 107 {
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
			m.SetPrice(r.Float64())
		case "Volume":
			m.SetVolume(r.Float64())
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

func (m *MarketDataUpdateTradeBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradePointerBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat64())
		case 4:
			m.SetVolume(r.ReadFloat64())
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

func (m *MarketDataUpdateTradePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat64())
		case 4:
			m.SetVolume(r.ReadFloat64())
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

func (m MarketDataUpdateTradeBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 107)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint16(2, uint16(m.AtBidOrAsk()))
	w.WriteFixed64Float64(3, m.Price())
	w.WriteFixed64Float64(4, m.Volume())
	w.WriteFixed64Float64(5, float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 107)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint16(2, uint16(m.AtBidOrAsk()))
	w.WriteFixed64Float64(3, m.Price())
	w.WriteFixed64Float64(4, m.Volume())
	w.WriteFixed64Float64(5, float64(m.DateTime()))
	return w.Finish(), nil
}
