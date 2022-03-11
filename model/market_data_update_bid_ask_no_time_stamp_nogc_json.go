//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

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
