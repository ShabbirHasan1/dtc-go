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

func (m MarketDataUpdateBidAskCompact) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":117,\"F\":["...)
	w.Float32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskCompactBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":117,\"F\":["...)
	w.Float32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskCompactPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":117,\"F\":["...)
	w.Float32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskCompactPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":117,\"F\":["...)
	w.Float32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskCompact) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 117)
	w.Float32Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	w.Uint32Field("SymbolID", m.SymbolID())
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskCompactBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 117)
	w.Float32Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	w.Uint32Field("SymbolID", m.SymbolID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskCompactPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 117)
	w.Float32Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	w.Uint32Field("SymbolID", m.SymbolID())
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskCompactPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 117)
	w.Float32Field("BidPrice", m.BidPrice())
	w.Float32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Float32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	w.Uint32Field("SymbolID", m.SymbolID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskCompactBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 117 {
		return message.ErrWrongType
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
	m.SetDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskCompactPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 117 {
		return message.ErrWrongType
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
	m.SetDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskCompactBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 117 {
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
		case "BidPrice":
			m.SetBidPrice(r.Float32())
		case "BidQuantity":
			m.SetBidQuantity(r.Float32())
		case "AskPrice":
			m.SetAskPrice(r.Float32())
		case "AskQuantity":
			m.SetAskQuantity(r.Float32())
		case "DateTime":
			m.SetDateTime(DateTime4Byte(r.Uint32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
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

func (m *MarketDataUpdateBidAskCompactPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 117 {
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
		case "BidPrice":
			m.SetBidPrice(r.Float32())
		case "BidQuantity":
			m.SetBidQuantity(r.Float32())
		case "AskPrice":
			m.SetAskPrice(r.Float32())
		case "AskQuantity":
			m.SetAskQuantity(r.Float32())
		case "DateTime":
			m.SetDateTime(DateTime4Byte(r.Uint32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
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

func (m *MarketDataUpdateBidAskCompactBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidPrice(r.ReadFloat32())
		case 2:
			m.SetBidQuantity(r.ReadFloat32())
		case 3:
			m.SetAskPrice(r.ReadFloat32())
		case 4:
			m.SetAskQuantity(r.ReadFloat32())
		case 5:
			m.SetDateTime(DateTime4Byte(r.ReadUint32()))
		case 6:
			m.SetSymbolID(r.ReadUint32())
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

func (m *MarketDataUpdateBidAskCompactPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidPrice(r.ReadFloat32())
		case 2:
			m.SetBidQuantity(r.ReadFloat32())
		case 3:
			m.SetAskPrice(r.ReadFloat32())
		case 4:
			m.SetAskQuantity(r.ReadFloat32())
		case 5:
			m.SetDateTime(DateTime4Byte(r.ReadUint32()))
		case 6:
			m.SetSymbolID(r.ReadUint32())
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

func (m MarketDataUpdateBidAskCompactBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 117)
	w.WriteFixed32Float32(1, m.BidPrice())
	w.WriteFixed32Float32(2, m.BidQuantity())
	w.WriteFixed32Float32(3, m.AskPrice())
	w.WriteFixed32Float32(4, m.AskQuantity())
	w.WriteFixed32Uint32(5, uint32(m.DateTime()))
	w.WriteUvarint32(6, m.SymbolID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskCompactPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 117)
	w.WriteFixed32Float32(1, m.BidPrice())
	w.WriteFixed32Float32(2, m.BidQuantity())
	w.WriteFixed32Float32(3, m.AskPrice())
	w.WriteFixed32Float32(4, m.AskQuantity())
	w.WriteFixed32Uint32(5, uint32(m.DateTime()))
	w.WriteUvarint32(6, m.SymbolID())
	return w.Finish(), nil
}
