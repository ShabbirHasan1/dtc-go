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

func (m MarketDataUpdateTradingSessionDate) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":136,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.Date()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradingSessionDateBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":136,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.Date()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradingSessionDatePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":136,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.Date()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradingSessionDatePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":136,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.Date()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradingSessionDate) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 136)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint32Field("Date", uint32(m.Date()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradingSessionDateBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 136)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint32Field("Date", uint32(m.Date()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradingSessionDatePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 136)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint32Field("Date", uint32(m.Date()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradingSessionDatePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 136)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint32Field("Date", uint32(m.Date()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradingSessionDateBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 136 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradingSessionDatePointerBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 136 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradingSessionDateBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 136 {
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
		case "Date":
			m.SetDate(DateTime4Byte(r.Uint32()))
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

func (m *MarketDataUpdateTradingSessionDatePointerBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 136 {
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
		case "Date":
			m.SetDate(DateTime4Byte(r.Uint32()))
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

func (m *MarketDataUpdateTradingSessionDateBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradingSessionDatePointerBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradingSessionDateBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDate(DateTime4Byte(r.ReadUint32()))
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

func (m *MarketDataUpdateTradingSessionDatePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDate(DateTime4Byte(r.ReadUint32()))
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

func (m MarketDataUpdateTradingSessionDateBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 136)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint32(2, uint32(m.Date()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradingSessionDatePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 136)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint32(2, uint32(m.Date()))
	return w.Finish(), nil
}
