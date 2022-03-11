//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":304,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.String(m.UniqueExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenClose()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoOrderFills())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionQuantity())
	return w.Finish(), nil
}

func (m HistoricalOrderFillResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":304,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.String(m.UniqueExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenClose()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoOrderFills())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 304)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price", m.Price())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float64Field("Quantity", m.Quantity())
	w.StringField("UniqueExecutionID", m.UniqueExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("OpenClose", int32(m.OpenClose()))
	w.Uint8Field("NoOrderFills", m.NoOrderFills())
	w.StringField("InfoText", m.InfoText())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	return w.Finish(), nil
}

func (m HistoricalOrderFillResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 304)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price", m.Price())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float64Field("Quantity", m.Quantity())
	w.StringField("UniqueExecutionID", m.UniqueExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("OpenClose", int32(m.OpenClose()))
	w.Uint8Field("NoOrderFills", m.NoOrderFills())
	w.StringField("InfoText", m.InfoText())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 304 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetUniqueExecutionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenClose(OpenCloseTradeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetNoOrderFills(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetHighPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLowPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPositionQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 304 {
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
		case "RequestID":
			m.SetRequestID(r.Int32())
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "Price":
			m.SetPrice(r.Float64())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "UniqueExecutionID":
			m.SetUniqueExecutionID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "OpenClose":
			m.SetOpenClose(OpenCloseTradeEnum(r.Int32()))
		case "NoOrderFills":
			m.SetNoOrderFills(r.Uint8())
		case "InfoText":
			m.SetInfoText(r.String())
		case "HighPriceDuringPosition":
			m.SetHighPriceDuringPosition(r.Float64())
		case "LowPriceDuringPosition":
			m.SetLowPriceDuringPosition(r.Float64())
		case "PositionQuantity":
			m.SetPositionQuantity(r.Float64())
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
