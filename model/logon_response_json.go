//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":2,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Result()))
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.String(m.ReconnectAddress())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_1())
	w.Data = append(w.Data, ',')
	w.String(m.ServerName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDepthUpdatesBestBidAndAsk())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradingIsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OCOOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OrderCancelReplaceSupported())
	w.Data = append(w.Data, ',')
	w.String(m.SymbolExchangeDelimiter())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SecurityDefinitionsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.HistoricalPriceDataSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ResubscribeWhenMarketDataFeedAvailable())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDepthIsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OneHistoricalPriceDataRequestPerConnection())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BracketOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseIntegerPriceOrderMessages())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDataSupported())
	return w.Finish(), nil
}

func (m LogonResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":2,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Result()))
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.String(m.ReconnectAddress())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_1())
	w.Data = append(w.Data, ',')
	w.String(m.ServerName())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDepthUpdatesBestBidAndAsk())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradingIsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OCOOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OrderCancelReplaceSupported())
	w.Data = append(w.Data, ',')
	w.String(m.SymbolExchangeDelimiter())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SecurityDefinitionsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.HistoricalPriceDataSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ResubscribeWhenMarketDataFeedAvailable())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDepthIsSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OneHistoricalPriceDataRequestPerConnection())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BracketOrdersSupported())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseIntegerPriceOrderMessages())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 2)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Result", int32(m.Result()))
	w.StringField("ResultText", m.ResultText())
	w.StringField("ReconnectAddress", m.ReconnectAddress())
	w.Int32Field("Integer_1", m.Integer_1())
	w.StringField("ServerName", m.ServerName())
	w.Uint8Field("MarketDepthUpdatesBestBidAndAsk", m.MarketDepthUpdatesBestBidAndAsk())
	w.Uint8Field("TradingIsSupported", m.TradingIsSupported())
	w.Uint8Field("OCOOrdersSupported", m.OCOOrdersSupported())
	w.Uint8Field("OrderCancelReplaceSupported", m.OrderCancelReplaceSupported())
	w.StringField("SymbolExchangeDelimiter", m.SymbolExchangeDelimiter())
	w.Uint8Field("SecurityDefinitionsSupported", m.SecurityDefinitionsSupported())
	w.Uint8Field("HistoricalPriceDataSupported", m.HistoricalPriceDataSupported())
	w.Uint8Field("ResubscribeWhenMarketDataFeedAvailable", m.ResubscribeWhenMarketDataFeedAvailable())
	w.Uint8Field("MarketDepthIsSupported", m.MarketDepthIsSupported())
	w.Uint8Field("OneHistoricalPriceDataRequestPerConnection", m.OneHistoricalPriceDataRequestPerConnection())
	w.Uint8Field("BracketOrdersSupported", m.BracketOrdersSupported())
	w.Uint8Field("UseIntegerPriceOrderMessages", m.UseIntegerPriceOrderMessages())
	w.Uint8Field("UsesMultiplePositionsPerSymbolAndTradeAccount", m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.Uint8Field("MarketDataSupported", m.MarketDataSupported())
	return w.Finish(), nil
}

func (m LogonResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 2)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Result", int32(m.Result()))
	w.StringField("ResultText", m.ResultText())
	w.StringField("ReconnectAddress", m.ReconnectAddress())
	w.Int32Field("Integer_1", m.Integer_1())
	w.StringField("ServerName", m.ServerName())
	w.Uint8Field("MarketDepthUpdatesBestBidAndAsk", m.MarketDepthUpdatesBestBidAndAsk())
	w.Uint8Field("TradingIsSupported", m.TradingIsSupported())
	w.Uint8Field("OCOOrdersSupported", m.OCOOrdersSupported())
	w.Uint8Field("OrderCancelReplaceSupported", m.OrderCancelReplaceSupported())
	w.StringField("SymbolExchangeDelimiter", m.SymbolExchangeDelimiter())
	w.Uint8Field("SecurityDefinitionsSupported", m.SecurityDefinitionsSupported())
	w.Uint8Field("HistoricalPriceDataSupported", m.HistoricalPriceDataSupported())
	w.Uint8Field("ResubscribeWhenMarketDataFeedAvailable", m.ResubscribeWhenMarketDataFeedAvailable())
	w.Uint8Field("MarketDepthIsSupported", m.MarketDepthIsSupported())
	w.Uint8Field("OneHistoricalPriceDataRequestPerConnection", m.OneHistoricalPriceDataRequestPerConnection())
	w.Uint8Field("BracketOrdersSupported", m.BracketOrdersSupported())
	w.Uint8Field("UseIntegerPriceOrderMessages", m.UseIntegerPriceOrderMessages())
	w.Uint8Field("UsesMultiplePositionsPerSymbolAndTradeAccount", m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.Uint8Field("MarketDataSupported", m.MarketDataSupported())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 2 {
		return message.ErrWrongType
	}
	m.SetProtocolVersion(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetResult(LogonStatusEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetResultText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetReconnectAddress(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetInteger_1(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerName(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDepthUpdatesBestBidAndAsk(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingIsSupported(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOOrdersSupported(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderCancelReplaceSupported(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolExchangeDelimiter(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityDefinitionsSupported(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetHistoricalPriceDataSupported(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetResubscribeWhenMarketDataFeedAvailable(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDepthIsSupported(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetOneHistoricalPriceDataRequestPerConnection(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetBracketOrdersSupported(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseIntegerPriceOrderMessages(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDataSupported(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 2 {
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
		case "ProtocolVersion":
			m.SetProtocolVersion(r.Int32())
		case "Result":
			m.SetResult(LogonStatusEnum(r.Int32()))
		case "ResultText":
			m.SetResultText(r.String())
		case "ReconnectAddress":
			m.SetReconnectAddress(r.String())
		case "Integer_1":
			m.SetInteger_1(r.Int32())
		case "ServerName":
			m.SetServerName(r.String())
		case "MarketDepthUpdatesBestBidAndAsk":
			m.SetMarketDepthUpdatesBestBidAndAsk(r.Uint8())
		case "TradingIsSupported":
			m.SetTradingIsSupported(r.Uint8())
		case "OCOOrdersSupported":
			m.SetOCOOrdersSupported(r.Uint8())
		case "OrderCancelReplaceSupported":
			m.SetOrderCancelReplaceSupported(r.Uint8())
		case "SymbolExchangeDelimiter":
			m.SetSymbolExchangeDelimiter(r.String())
		case "SecurityDefinitionsSupported":
			m.SetSecurityDefinitionsSupported(r.Uint8())
		case "HistoricalPriceDataSupported":
			m.SetHistoricalPriceDataSupported(r.Uint8())
		case "ResubscribeWhenMarketDataFeedAvailable":
			m.SetResubscribeWhenMarketDataFeedAvailable(r.Uint8())
		case "MarketDepthIsSupported":
			m.SetMarketDepthIsSupported(r.Uint8())
		case "OneHistoricalPriceDataRequestPerConnection":
			m.SetOneHistoricalPriceDataRequestPerConnection(r.Uint8())
		case "BracketOrdersSupported":
			m.SetBracketOrdersSupported(r.Uint8())
		case "UseIntegerPriceOrderMessages":
			m.SetUseIntegerPriceOrderMessages(r.Uint8())
		case "UsesMultiplePositionsPerSymbolAndTradeAccount":
			m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.Uint8())
		case "MarketDataSupported":
			m.SetMarketDataSupported(r.Uint8())
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
