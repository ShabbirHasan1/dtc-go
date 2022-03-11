//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewOCOOrderInt) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":207,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID_1())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderType_1()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell_1()))
	w.Data = append(w.Data, ',')
	w.Int64(m.Price1_1())
	w.Data = append(w.Data, ',')
	w.Int64(m.Price2_1())
	w.Data = append(w.Data, ',')
	w.Int64(m.Quantity_1())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID_2())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderType_2()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell_2()))
	w.Data = append(w.Data, ',')
	w.Int64(m.Price1_2())
	w.Data = append(w.Data, ',')
	w.Int64(m.Price2_2())
	w.Data = append(w.Data, ',')
	w.Int64(m.Quantity_2())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.TimeInForce()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.GoodTillDateTime()))
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.String(m.ParentTriggerClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Float32(m.Divisor())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenOrClose()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.PartialFillHandling()))
	return w.Finish(), nil
}

func (m SubmitNewOCOOrderIntBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":207,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID_1())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderType_1()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell_1()))
	w.Data = append(w.Data, ',')
	w.Int64(m.Price1_1())
	w.Data = append(w.Data, ',')
	w.Int64(m.Price2_1())
	w.Data = append(w.Data, ',')
	w.Int64(m.Quantity_1())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID_2())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderType_2()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell_2()))
	w.Data = append(w.Data, ',')
	w.Int64(m.Price1_2())
	w.Data = append(w.Data, ',')
	w.Int64(m.Price2_2())
	w.Data = append(w.Data, ',')
	w.Int64(m.Quantity_2())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.TimeInForce()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.GoodTillDateTime()))
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.String(m.ParentTriggerClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Float32(m.Divisor())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenOrClose()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.PartialFillHandling()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewOCOOrderInt) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 207)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ClientOrderID_1", m.ClientOrderID_1())
	w.Int32Field("OrderType_1", int32(m.OrderType_1()))
	w.Int32Field("BuySell_1", int32(m.BuySell_1()))
	w.Int64Field("Price1_1", m.Price1_1())
	w.Int64Field("Price2_1", m.Price2_1())
	w.Int64Field("Quantity_1", m.Quantity_1())
	w.StringField("ClientOrderID_2", m.ClientOrderID_2())
	w.Int32Field("OrderType_2", int32(m.OrderType_2()))
	w.Int32Field("BuySell_2", int32(m.BuySell_2()))
	w.Int64Field("Price1_2", m.Price1_2())
	w.Int64Field("Price2_2", m.Price2_2())
	w.Int64Field("Quantity_2", m.Quantity_2())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("IsAutomatedOrder", m.IsAutomatedOrder())
	w.StringField("ParentTriggerClientOrderID", m.ParentTriggerClientOrderID())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Float32Field("Divisor", m.Divisor())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Int8Field("PartialFillHandling", int8(m.PartialFillHandling()))
	return w.Finish(), nil
}

func (m SubmitNewOCOOrderIntBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 207)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ClientOrderID_1", m.ClientOrderID_1())
	w.Int32Field("OrderType_1", int32(m.OrderType_1()))
	w.Int32Field("BuySell_1", int32(m.BuySell_1()))
	w.Int64Field("Price1_1", m.Price1_1())
	w.Int64Field("Price2_1", m.Price2_1())
	w.Int64Field("Quantity_1", m.Quantity_1())
	w.StringField("ClientOrderID_2", m.ClientOrderID_2())
	w.Int32Field("OrderType_2", int32(m.OrderType_2()))
	w.Int32Field("BuySell_2", int32(m.BuySell_2()))
	w.Int64Field("Price1_2", m.Price1_2())
	w.Int64Field("Price2_2", m.Price2_2())
	w.Int64Field("Quantity_2", m.Quantity_2())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("IsAutomatedOrder", m.IsAutomatedOrder())
	w.StringField("ParentTriggerClientOrderID", m.ParentTriggerClientOrderID())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Float32Field("Divisor", m.Divisor())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Int8Field("PartialFillHandling", int8(m.PartialFillHandling()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewOCOOrderIntBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 207 {
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
	m.SetClientOrderID_1(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderType_1(OrderTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell_1(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1_1(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2_1(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity_1(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID_2(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderType_2(OrderTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell_2(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1_2(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2_2(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity_2(r.Int64())
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
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsAutomatedOrder(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetParentTriggerClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetFreeFormText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenOrClose(OpenCloseTradeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPartialFillHandling(PartialFillHandlingEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewOCOOrderIntBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 207 {
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
		case "ClientOrderID_1":
			m.SetClientOrderID_1(r.String())
		case "OrderType_1":
			m.SetOrderType_1(OrderTypeEnum(r.Int32()))
		case "BuySell_1":
			m.SetBuySell_1(BuySellEnum(r.Int32()))
		case "Price1_1":
			m.SetPrice1_1(r.Int64())
		case "Price2_1":
			m.SetPrice2_1(r.Int64())
		case "Quantity_1":
			m.SetQuantity_1(r.Int64())
		case "ClientOrderID_2":
			m.SetClientOrderID_2(r.String())
		case "OrderType_2":
			m.SetOrderType_2(OrderTypeEnum(r.Int32()))
		case "BuySell_2":
			m.SetBuySell_2(BuySellEnum(r.Int32()))
		case "Price1_2":
			m.SetPrice1_2(r.Int64())
		case "Price2_2":
			m.SetPrice2_2(r.Int64())
		case "Quantity_2":
			m.SetQuantity_2(r.Int64())
		case "TimeInForce":
			m.SetTimeInForce(TimeInForceEnum(r.Int32()))
		case "GoodTillDateTime":
			m.SetGoodTillDateTime(DateTime(r.Int64()))
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "IsAutomatedOrder":
			m.SetIsAutomatedOrder(r.Uint8())
		case "ParentTriggerClientOrderID":
			m.SetParentTriggerClientOrderID(r.String())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
		case "Divisor":
			m.SetDivisor(r.Float32())
		case "OpenOrClose":
			m.SetOpenOrClose(OpenCloseTradeEnum(r.Int32()))
		case "PartialFillHandling":
			m.SetPartialFillHandling(PartialFillHandlingEnum(r.Int8()))
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
