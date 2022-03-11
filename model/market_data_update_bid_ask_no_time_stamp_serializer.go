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

func (m MarketDataUpdateBidAskNoTimeStamp) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":143,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Uint32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Uint32(m.AskQuantity())
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskNoTimeStampBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":143,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Uint32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Uint32(m.AskQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskNoTimeStampPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":143,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Uint32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Uint32(m.AskQuantity())
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskNoTimeStampPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":143,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Uint32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Uint32(m.AskQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskNoTimeStamp) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 143)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("BidPrice", m.BidPrice())
	w.Uint32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Uint32Field("AskQuantity", m.AskQuantity())
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskNoTimeStampBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 143)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("BidPrice", m.BidPrice())
	w.Uint32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Uint32Field("AskQuantity", m.AskQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskNoTimeStampPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 143)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("BidPrice", m.BidPrice())
	w.Uint32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Uint32Field("AskQuantity", m.AskQuantity())
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAskNoTimeStampPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 143)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("BidPrice", m.BidPrice())
	w.Uint32Field("BidQuantity", m.BidQuantity())
	w.Float32Field("AskPrice", m.AskPrice())
	w.Uint32Field("AskQuantity", m.AskQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskNoTimeStampBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 143 {
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
	m.SetBidQuantity(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskQuantity(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskNoTimeStampPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 143 {
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
	m.SetBidQuantity(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskQuantity(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskNoTimeStampBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 143 {
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
			m.SetBidQuantity(r.Uint32())
		case "AskPrice":
			m.SetAskPrice(r.Float32())
		case "AskQuantity":
			m.SetAskQuantity(r.Uint32())
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

func (m *MarketDataUpdateBidAskNoTimeStampPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 143 {
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
			m.SetBidQuantity(r.Uint32())
		case "AskPrice":
			m.SetAskPrice(r.Float32())
		case "AskQuantity":
			m.SetAskQuantity(r.Uint32())
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

func (m *MarketDataUpdateBidAskNoTimeStampBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidQuantity(r.ReadUint32())
		case 4:
			m.SetAskPrice(r.ReadFloat32())
		case 5:
			m.SetAskQuantity(r.ReadUint32())
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

func (m *MarketDataUpdateBidAskNoTimeStampPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidQuantity(r.ReadUint32())
		case 4:
			m.SetAskPrice(r.ReadFloat32())
		case 5:
			m.SetAskQuantity(r.ReadUint32())
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

func (m MarketDataUpdateBidAskNoTimeStampBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 143)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.BidPrice())
	w.WriteUvarint32(3, m.BidQuantity())
	w.WriteFixed32Float32(4, m.AskPrice())
	w.WriteUvarint32(5, m.AskQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAskNoTimeStampPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 143)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.BidPrice())
	w.WriteUvarint32(3, m.BidQuantity())
	w.WriteFixed32Float32(4, m.AskPrice())
	w.WriteUvarint32(5, m.AskQuantity())
	return w.Finish(), nil
}