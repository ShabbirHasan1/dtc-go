//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

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
