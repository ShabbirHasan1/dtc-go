//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
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
