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

func (m SubmitNewSingleOrderInt) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsParentOrder())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenOrClose()))
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderIntBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsParentOrder())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenOrClose()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderIntPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsParentOrder())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenOrClose()))
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderIntPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsParentOrder())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenOrClose()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderInt) MarshalJSONTo(b []byte) ([]byte, error) {
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
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderIntBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
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
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderIntPointer) MarshalJSONTo(b []byte) ([]byte, error) {
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
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderIntPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
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
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderIntBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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
	m.SetIsAutomatedOrder(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsParentOrder(r.Bool())
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
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderIntPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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
	m.SetIsAutomatedOrder(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsParentOrder(r.Bool())
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

func (m *SubmitNewSingleOrderIntBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
			m.SetIsAutomatedOrder(r.Bool())
		case "IsParentOrder":
			m.SetIsParentOrder(r.Bool())
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderIntPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
			m.SetIsAutomatedOrder(r.Bool())
		case "IsParentOrder":
			m.SetIsParentOrder(r.Bool())
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
	w.Bool(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsParentOrder())
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
	w.Bool(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsParentOrder())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenOrClose()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderIntFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsParentOrder())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenOrClose()))
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderIntFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsParentOrder())
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
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
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
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderIntFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
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
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderIntFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
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
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
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
	m.SetIsAutomatedOrder(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsParentOrder(r.Bool())
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
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderIntFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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
	m.SetIsAutomatedOrder(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsParentOrder(r.Bool())
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
			m.SetIsAutomatedOrder(r.Bool())
		case "IsParentOrder":
			m.SetIsParentOrder(r.Bool())
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderIntFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
			m.SetIsAutomatedOrder(r.Bool())
		case "IsParentOrder":
			m.SetIsParentOrder(r.Bool())
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

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderIntBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 2:
			m.SetExchange(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetClientOrderID(r.ReadString())
		case 5:
			m.SetOrderType(OrderTypeEnum(r.ReadInt32()))
		case 6:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 7:
			m.SetPrice1(r.ReadInt64())
		case 8:
			m.SetPrice2(r.ReadInt64())
		case 9:
			m.SetDivisor(r.ReadFloat32())
		case 10:
			m.SetQuantity(r.ReadInt64())
		case 11:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 12:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 13:
			m.SetIsAutomatedOrder(r.ReadBool())
		case 14:
			m.SetIsParentOrder(r.ReadBool())
		case 15:
			m.SetFreeFormText(r.ReadString())
		case 16:
			m.SetOpenOrClose(OpenCloseTradeEnum(r.ReadInt32()))
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

func (m *SubmitNewSingleOrderIntPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 2:
			m.SetExchange(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetClientOrderID(r.ReadString())
		case 5:
			m.SetOrderType(OrderTypeEnum(r.ReadInt32()))
		case 6:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 7:
			m.SetPrice1(r.ReadInt64())
		case 8:
			m.SetPrice2(r.ReadInt64())
		case 9:
			m.SetDivisor(r.ReadFloat32())
		case 10:
			m.SetQuantity(r.ReadInt64())
		case 11:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 12:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 13:
			m.SetIsAutomatedOrder(r.ReadBool())
		case 14:
			m.SetIsParentOrder(r.ReadBool())
		case 15:
			m.SetFreeFormText(r.ReadString())
		case 16:
			m.SetOpenOrClose(OpenCloseTradeEnum(r.ReadInt32()))
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

func (m SubmitNewSingleOrderIntBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 206)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.OrderType()))
	w.WriteVarint32(6, int32(m.BuySell()))
	w.WriteVarint64(7, m.Price1())
	w.WriteVarint64(8, m.Price2())
	w.WriteFixed32Float32(9, m.Divisor())
	w.WriteVarint64(10, m.Quantity())
	w.WriteVarint32(11, int32(m.TimeInForce()))
	w.WriteFixed64Int64(12, int64(m.GoodTillDateTime()))
	w.WriteBool(13, m.IsAutomatedOrder())
	w.WriteBool(14, m.IsParentOrder())
	w.WriteString(15, m.FreeFormText())
	w.WriteVarint32(16, int32(m.OpenOrClose()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderIntPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 206)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.OrderType()))
	w.WriteVarint32(6, int32(m.BuySell()))
	w.WriteVarint64(7, m.Price1())
	w.WriteVarint64(8, m.Price2())
	w.WriteFixed32Float32(9, m.Divisor())
	w.WriteVarint64(10, m.Quantity())
	w.WriteVarint32(11, int32(m.TimeInForce()))
	w.WriteFixed64Int64(12, int64(m.GoodTillDateTime()))
	w.WriteBool(13, m.IsAutomatedOrder())
	w.WriteBool(14, m.IsParentOrder())
	w.WriteString(15, m.FreeFormText())
	w.WriteVarint32(16, int32(m.OpenOrClose()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderIntFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 2:
			m.SetExchange(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetClientOrderID(r.ReadString())
		case 5:
			m.SetOrderType(OrderTypeEnum(r.ReadInt32()))
		case 6:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 7:
			m.SetPrice1(r.ReadInt64())
		case 8:
			m.SetPrice2(r.ReadInt64())
		case 9:
			m.SetDivisor(r.ReadFloat32())
		case 10:
			m.SetQuantity(r.ReadInt64())
		case 11:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 12:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 13:
			m.SetIsAutomatedOrder(r.ReadBool())
		case 14:
			m.SetIsParentOrder(r.ReadBool())
		case 15:
			m.SetFreeFormText(r.ReadString())
		case 16:
			m.SetOpenOrClose(OpenCloseTradeEnum(r.ReadInt32()))
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

func (m *SubmitNewSingleOrderIntFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 2:
			m.SetExchange(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetClientOrderID(r.ReadString())
		case 5:
			m.SetOrderType(OrderTypeEnum(r.ReadInt32()))
		case 6:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 7:
			m.SetPrice1(r.ReadInt64())
		case 8:
			m.SetPrice2(r.ReadInt64())
		case 9:
			m.SetDivisor(r.ReadFloat32())
		case 10:
			m.SetQuantity(r.ReadInt64())
		case 11:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 12:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 13:
			m.SetIsAutomatedOrder(r.ReadBool())
		case 14:
			m.SetIsParentOrder(r.ReadBool())
		case 15:
			m.SetFreeFormText(r.ReadString())
		case 16:
			m.SetOpenOrClose(OpenCloseTradeEnum(r.ReadInt32()))
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

func (m SubmitNewSingleOrderIntFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 206)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.OrderType()))
	w.WriteVarint32(6, int32(m.BuySell()))
	w.WriteVarint64(7, m.Price1())
	w.WriteVarint64(8, m.Price2())
	w.WriteFixed32Float32(9, m.Divisor())
	w.WriteVarint64(10, m.Quantity())
	w.WriteVarint32(11, int32(m.TimeInForce()))
	w.WriteFixed64Int64(12, int64(m.GoodTillDateTime()))
	w.WriteBool(13, m.IsAutomatedOrder())
	w.WriteBool(14, m.IsParentOrder())
	w.WriteString(15, m.FreeFormText())
	w.WriteVarint32(16, int32(m.OpenOrClose()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderIntFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 206)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.OrderType()))
	w.WriteVarint32(6, int32(m.BuySell()))
	w.WriteVarint64(7, m.Price1())
	w.WriteVarint64(8, m.Price2())
	w.WriteFixed32Float32(9, m.Divisor())
	w.WriteVarint64(10, m.Quantity())
	w.WriteVarint32(11, int32(m.TimeInForce()))
	w.WriteFixed64Int64(12, int64(m.GoodTillDateTime()))
	w.WriteBool(13, m.IsAutomatedOrder())
	w.WriteBool(14, m.IsParentOrder())
	w.WriteString(15, m.FreeFormText())
	w.WriteVarint32(16, int32(m.OpenOrClose()))
	return w.Finish(), nil
}
