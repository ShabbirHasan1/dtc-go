//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderIntFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":206,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderType()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Int64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Int64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float32(m.Divisor())
	w.Data = append(w.Data, ',')
	w.Int64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.TimeInForce()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.GoodTillDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsParentOrder())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenOrClose()))
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderIntFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":206,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderType()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Int64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Int64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float32(m.Divisor())
	w.Data = append(w.Data, ',')
	w.Int64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.TimeInForce()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.GoodTillDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsParentOrder())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenOrClose()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderIntFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 206)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Int64Field("Price1", m.Price1())
	w.Int64Field("Price2", m.Price2())
	w.Float32Field("Divisor", m.Divisor())
	w.Int64Field("Quantity", m.Quantity())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("IsAutomatedOrder", m.IsAutomatedOrder())
	w.Uint8Field("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderIntFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 206)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Int64Field("Price1", m.Price1())
	w.Int64Field("Price2", m.Price2())
	w.Float32Field("Divisor", m.Divisor())
	w.Int64Field("Quantity", m.Quantity())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("IsAutomatedOrder", m.IsAutomatedOrder())
	w.Uint8Field("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderIntFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 206 {
		return message.ErrWrongType
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderType(OrderTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTimeInForce(TimeInForceEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetGoodTillDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetIsAutomatedOrder(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsParentOrder(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFreeFormText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenOrClose(OpenCloseTradeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderIntFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 206 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "OrderType":
			m.SetOrderType(OrderTypeEnum(r.Int32()))
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "Price1":
			m.SetPrice1(r.Int64())
		case "Price2":
			m.SetPrice2(r.Int64())
		case "Divisor":
			m.SetDivisor(r.Float32())
		case "Quantity":
			m.SetQuantity(r.Int64())
		case "TimeInForce":
			m.SetTimeInForce(TimeInForceEnum(r.Int32()))
		case "GoodTillDateTime":
			m.SetGoodTillDateTime(DateTime(r.Int64()))
		case "IsAutomatedOrder":
			m.SetIsAutomatedOrder(r.Uint8())
		case "IsParentOrder":
			m.SetIsParentOrder(r.Uint8())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
		case "OpenOrClose":
			m.SetOpenOrClose(OpenCloseTradeEnum(r.Int32()))
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
