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

func (m MarketDataUpdateSessionHigh_Int) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":129,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionHigh_IntBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":129,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionHigh_IntPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":129,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionHigh_IntPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":129,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionHigh_Int) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 129)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("Price", m.Price())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionHigh_IntBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 129)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("Price", m.Price())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionHigh_IntPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 129)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("Price", m.Price())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionHigh_IntPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 129)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("Price", m.Price())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateSessionHigh_IntBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 129 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Int32())
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

func (m *MarketDataUpdateSessionHigh_IntPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 129 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Int32())
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

func (m *MarketDataUpdateSessionHigh_IntBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 129 {
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
			m.SetPrice(r.Int32())
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

func (m *MarketDataUpdateSessionHigh_IntPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 129 {
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
			m.SetPrice(r.Int32())
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

func (m *MarketDataUpdateSessionHigh_IntBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadInt32())
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

func (m *MarketDataUpdateSessionHigh_IntPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadInt32())
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

func (m MarketDataUpdateSessionHigh_IntBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 129)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint32(2, m.Price())
	w.WriteUvarint32(3, uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionHigh_IntPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 129)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint32(2, m.Price())
	w.WriteUvarint32(3, uint32(m.TradingSessionDate()))
	return w.Finish(), nil
}
