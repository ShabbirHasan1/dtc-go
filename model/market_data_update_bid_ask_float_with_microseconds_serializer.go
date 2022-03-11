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

func (m MarketDataUpdateBidAskFloatWithMicroseconds) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":144,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":144,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":144,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":144,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskFloatWithMicroseconds) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 144)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 144)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 144)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 144)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 144 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidQuantity(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskQuantity(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 144 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidQuantity(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskQuantity(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 144 {
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
		case "BidPrice":
			m.SetBidPrice(r.Float32())
		case "BidQuantity":
			m.SetBidQuantity(r.Float32())
		case "AskPrice":
			m.SetAskPrice(r.Float32())
		case "AskQuantity":
			m.SetAskQuantity(r.Float32())
		case "DateTime":
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
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

func (m *MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 144 {
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
		case "BidPrice":
			m.SetBidPrice(r.Float32())
		case "BidQuantity":
			m.SetBidQuantity(r.Float32())
		case "AskPrice":
			m.SetAskPrice(r.Float32())
		case "AskQuantity":
			m.SetAskQuantity(r.Float32())
		case "DateTime":
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
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

func (m *MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidPrice(r.ReadFloat32())
		case 3:
			m.SetBidQuantity(r.ReadFloat32())
		case 4:
			m.SetAskPrice(r.ReadFloat32())
		case 5:
			m.SetAskQuantity(r.ReadFloat32())
		case 6:
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
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

func (m *MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidPrice(r.ReadFloat32())
		case 3:
			m.SetBidQuantity(r.ReadFloat32())
		case 4:
			m.SetAskPrice(r.ReadFloat32())
		case 5:
			m.SetAskQuantity(r.ReadFloat32())
		case 6:
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
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

func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 144)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.BidPrice())
	w.WriteFixed32Float32(3, m.BidQuantity())
	w.WriteFixed32Float32(4, m.AskPrice())
	w.WriteFixed32Float32(5, m.AskQuantity())
	w.WriteVarint64(6, int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 144)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.BidPrice())
	w.WriteFixed32Float32(3, m.BidQuantity())
	w.WriteFixed32Float32(4, m.AskPrice())
	w.WriteFixed32Float32(5, m.AskQuantity())
	w.WriteVarint64(6, int64(m.DateTime()))
	return w.Finish(), nil
}
