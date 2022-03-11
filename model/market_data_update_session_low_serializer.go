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

func (m MarketDataUpdateSessionLow) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":115,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionLowBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":115,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionLowPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":115,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionLowPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":115,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionLow) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 115)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("Price", m.Price())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionLowBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 115)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("Price", m.Price())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionLowPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 115)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("Price", m.Price())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionLowPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 115)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("Price", m.Price())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateSessionLowBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 115 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateSessionLowPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 115 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateSessionLowBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 115 {
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
			m.SetPrice(r.Float64())
		case "TradingSessionDate":
			m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
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

func (m *MarketDataUpdateSessionLowPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 115 {
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
			m.SetPrice(r.Float64())
		case "TradingSessionDate":
			m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
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

func (m *MarketDataUpdateSessionLowBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat64())
		case 3:
			m.SetTradingSessionDate(DateTime4Byte(r.ReadUint32()))
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

func (m *MarketDataUpdateSessionLowPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat64())
		case 3:
			m.SetTradingSessionDate(DateTime4Byte(r.ReadUint32()))
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

func (m MarketDataUpdateSessionLowBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 115)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed64Float64(2, m.Price())
	w.WriteUvarint32(3, uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionLowPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 115)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed64Float64(2, m.Price())
	w.WriteUvarint32(3, uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}
