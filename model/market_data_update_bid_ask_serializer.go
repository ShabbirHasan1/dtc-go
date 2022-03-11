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

func (m MarketDataUpdateBidAsk) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":108,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":108,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":108,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":108,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAsk) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 108)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float64Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 108)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float64Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 108)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float64Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 108)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float64Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 108 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidQuantity(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskQuantity(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 108 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidQuantity(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskQuantity(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 108 {
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
			m.SetBidPrice(r.Float64())
		case "BidQuantity":
			m.SetBidQuantity(r.Float32())
		case "AskPrice":
			m.SetAskPrice(r.Float64())
		case "AskQuantity":
			m.SetAskQuantity(r.Float32())
		case "DateTime":
			m.SetDateTime(DateTime4Byte(r.Uint32()))
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

func (m *MarketDataUpdateBidAskPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 108 {
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
			m.SetBidPrice(r.Float64())
		case "BidQuantity":
			m.SetBidQuantity(r.Float32())
		case "AskPrice":
			m.SetAskPrice(r.Float64())
		case "AskQuantity":
			m.SetAskQuantity(r.Float32())
		case "DateTime":
			m.SetDateTime(DateTime4Byte(r.Uint32()))
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

func (m *MarketDataUpdateBidAskBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidPrice(r.ReadFloat64())
		case 3:
			m.SetBidQuantity(r.ReadFloat32())
		case 4:
			m.SetAskPrice(r.ReadFloat64())
		case 5:
			m.SetAskQuantity(r.ReadFloat32())
		case 6:
			m.SetDateTime(DateTime4Byte(r.ReadUint32()))
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

func (m *MarketDataUpdateBidAskPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidPrice(r.ReadFloat64())
		case 3:
			m.SetBidQuantity(r.ReadFloat32())
		case 4:
			m.SetAskPrice(r.ReadFloat64())
		case 5:
			m.SetAskQuantity(r.ReadFloat32())
		case 6:
			m.SetDateTime(DateTime4Byte(r.ReadUint32()))
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

func (m MarketDataUpdateBidAskBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 108)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed64Float64(2, m.BidPrice())
	w.WriteFixed32Float32(3, m.BidQuantity())
	w.WriteFixed64Float64(4, m.AskPrice())
	w.WriteFixed32Float32(5, m.AskQuantity())
	w.WriteFixed32Uint32(6, uint32(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 108)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed64Float64(2, m.BidPrice())
	w.WriteFixed32Float32(3, m.BidQuantity())
	w.WriteFixed64Float64(4, m.AskPrice())
	w.WriteFixed32Float32(5, m.AskQuantity())
	w.WriteFixed32Uint32(6, uint32(m.DateTime()))
	return w.Finish(), nil
}
