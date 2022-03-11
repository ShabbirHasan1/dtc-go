//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataSnapshot) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":104,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.SessionSettlementPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.SessionOpenPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.SessionHighPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.SessionLowPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.SessionVolume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SessionNumTrades())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Float64(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastTradePrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastTradeVolume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.LastTradeDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.BidAskDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SessionSettlementDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradingStatus()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.MarketDepthUpdateDateTime()))
	return w.Finish(), nil
}

func (m MarketDataSnapshotBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":104,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.SessionSettlementPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.SessionOpenPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.SessionHighPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.SessionLowPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.SessionVolume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SessionNumTrades())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Float64(m.BidPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.AskPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.AskQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.BidQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastTradePrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastTradeVolume())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.LastTradeDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.BidAskDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SessionSettlementDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradingStatus()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.MarketDepthUpdateDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataSnapshot) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 104)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("SessionSettlementPrice", m.SessionSettlementPrice())
	w.Float64Field("SessionOpenPrice", m.SessionOpenPrice())
	w.Float64Field("SessionHighPrice", m.SessionHighPrice())
	w.Float64Field("SessionLowPrice", m.SessionLowPrice())
	w.Float64Field("SessionVolume", m.SessionVolume())
	w.Uint32Field("SessionNumTrades", m.SessionNumTrades())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Float64Field("BidPrice", m.BidPrice())
	w.Float64Field("AskPrice", m.AskPrice())
	w.Float64Field("AskQuantity", m.AskQuantity())
	w.Float64Field("BidQuantity", m.BidQuantity())
	w.Float64Field("LastTradePrice", m.LastTradePrice())
	w.Float64Field("LastTradeVolume", m.LastTradeVolume())
	w.Float64Field("LastTradeDateTime", float64(m.LastTradeDateTime()))
	w.Float64Field("BidAskDateTime", float64(m.BidAskDateTime()))
	w.Uint32Field("SessionSettlementDateTime", uint32(m.SessionSettlementDateTime()))
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	w.Int8Field("TradingStatus", int8(m.TradingStatus()))
	w.Float64Field("MarketDepthUpdateDateTime", float64(m.MarketDepthUpdateDateTime()))
	return w.Finish(), nil
}

func (m MarketDataSnapshotBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 104)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("SessionSettlementPrice", m.SessionSettlementPrice())
	w.Float64Field("SessionOpenPrice", m.SessionOpenPrice())
	w.Float64Field("SessionHighPrice", m.SessionHighPrice())
	w.Float64Field("SessionLowPrice", m.SessionLowPrice())
	w.Float64Field("SessionVolume", m.SessionVolume())
	w.Uint32Field("SessionNumTrades", m.SessionNumTrades())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Float64Field("BidPrice", m.BidPrice())
	w.Float64Field("AskPrice", m.AskPrice())
	w.Float64Field("AskQuantity", m.AskQuantity())
	w.Float64Field("BidQuantity", m.BidQuantity())
	w.Float64Field("LastTradePrice", m.LastTradePrice())
	w.Float64Field("LastTradeVolume", m.LastTradeVolume())
	w.Float64Field("LastTradeDateTime", float64(m.LastTradeDateTime()))
	w.Float64Field("BidAskDateTime", float64(m.BidAskDateTime()))
	w.Uint32Field("SessionSettlementDateTime", uint32(m.SessionSettlementDateTime()))
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	w.Int8Field("TradingStatus", int8(m.TradingStatus()))
	w.Float64Field("MarketDepthUpdateDateTime", float64(m.MarketDepthUpdateDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataSnapshotBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 104 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionSettlementPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionOpenPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionHighPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionLowPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionVolume(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionNumTrades(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTradePrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTradeVolume(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetBidAskDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionSettlementDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingStatus(TradingStatusEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDepthUpdateDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataSnapshotBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 104 {
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
		case "SessionSettlementPrice":
			m.SetSessionSettlementPrice(r.Float64())
		case "SessionOpenPrice":
			m.SetSessionOpenPrice(r.Float64())
		case "SessionHighPrice":
			m.SetSessionHighPrice(r.Float64())
		case "SessionLowPrice":
			m.SetSessionLowPrice(r.Float64())
		case "SessionVolume":
			m.SetSessionVolume(r.Float64())
		case "SessionNumTrades":
			m.SetSessionNumTrades(r.Uint32())
		case "OpenInterest":
			m.SetOpenInterest(r.Uint32())
		case "BidPrice":
			m.SetBidPrice(r.Float64())
		case "AskPrice":
			m.SetAskPrice(r.Float64())
		case "AskQuantity":
			m.SetAskQuantity(r.Float64())
		case "BidQuantity":
			m.SetBidQuantity(r.Float64())
		case "LastTradePrice":
			m.SetLastTradePrice(r.Float64())
		case "LastTradeVolume":
			m.SetLastTradeVolume(r.Float64())
		case "LastTradeDateTime":
			m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "BidAskDateTime":
			m.SetBidAskDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "SessionSettlementDateTime":
			m.SetSessionSettlementDateTime(DateTime4Byte(r.Uint32()))
		case "TradingSessionDate":
			m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
		case "TradingStatus":
			m.SetTradingStatus(TradingStatusEnum(r.Int8()))
		case "MarketDepthUpdateDateTime":
			m.SetMarketDepthUpdateDateTime(DateTimeWithMilliseconds(r.Float64()))
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
