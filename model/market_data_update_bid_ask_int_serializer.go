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

func (m MarketDataUpdateBidAsk_Int) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":127,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAsk_IntBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":127,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAsk_IntPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":127,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAsk_IntPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":127,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAsk_Int) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 127)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("BidPrice", m.BidPrice())
	w.Int32Field("BidQuantity", m.BidQuantity())
	w.Int32Field("AskPrice", m.AskPrice())
	w.Int32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAsk_IntBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 127)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("BidPrice", m.BidPrice())
	w.Int32Field("BidQuantity", m.BidQuantity())
	w.Int32Field("AskPrice", m.AskPrice())
	w.Int32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAsk_IntPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 127)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("BidPrice", m.BidPrice())
	w.Int32Field("BidQuantity", m.BidQuantity())
	w.Int32Field("AskPrice", m.AskPrice())
	w.Int32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateBidAsk_IntPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 127)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("BidPrice", m.BidPrice())
	w.Int32Field("BidQuantity", m.BidQuantity())
	w.Int32Field("AskPrice", m.AskPrice())
	w.Int32Field("AskQuantity", m.AskQuantity())
	w.Uint32Field("DateTime", uint32(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAsk_IntBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 127 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidQuantity(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskQuantity(r.Int32())
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

func (m *MarketDataUpdateBidAsk_IntPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 127 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidQuantity(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskQuantity(r.Int32())
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

func (m *MarketDataUpdateBidAsk_IntBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 127 {
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
			m.SetBidPrice(r.Int32())
		case "BidQuantity":
			m.SetBidQuantity(r.Int32())
		case "AskPrice":
			m.SetAskPrice(r.Int32())
		case "AskQuantity":
			m.SetAskQuantity(r.Int32())
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

func (m *MarketDataUpdateBidAsk_IntPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 127 {
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
			m.SetBidPrice(r.Int32())
		case "BidQuantity":
			m.SetBidQuantity(r.Int32())
		case "AskPrice":
			m.SetAskPrice(r.Int32())
		case "AskQuantity":
			m.SetAskQuantity(r.Int32())
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

func (m *MarketDataUpdateBidAsk_IntBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidPrice(r.ReadInt32())
		case 3:
			m.SetBidQuantity(r.ReadInt32())
		case 4:
			m.SetAskPrice(r.ReadInt32())
		case 5:
			m.SetAskQuantity(r.ReadInt32())
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

func (m *MarketDataUpdateBidAsk_IntPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidPrice(r.ReadInt32())
		case 3:
			m.SetBidQuantity(r.ReadInt32())
		case 4:
			m.SetAskPrice(r.ReadInt32())
		case 5:
			m.SetAskQuantity(r.ReadInt32())
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

func (m MarketDataUpdateBidAsk_IntBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 127)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint32(2, m.BidPrice())
	w.WriteVarint32(3, m.BidQuantity())
	w.WriteVarint32(4, m.AskPrice())
	w.WriteVarint32(5, m.AskQuantity())
	w.WriteFixed32Uint32(6, uint32(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateBidAsk_IntPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 127)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint32(2, m.BidPrice())
	w.WriteVarint32(3, m.BidQuantity())
	w.WriteVarint32(4, m.AskPrice())
	w.WriteVarint32(5, m.AskQuantity())
	w.WriteFixed32Uint32(6, uint32(m.DateTime()))
	return w.Finish(), nil
}
