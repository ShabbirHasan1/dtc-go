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

func (m SubmitNewSingleOrder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":208,\"F\":["...)
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
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxShowQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":208,\"F\":["...)
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
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxShowQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":208,\"F\":["...)
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
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxShowQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":208,\"F\":["...)
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
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxShowQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 208)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Float64Field("MaxShowQuantity", m.MaxShowQuantity())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 208)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Float64Field("MaxShowQuantity", m.MaxShowQuantity())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 208)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Float64Field("MaxShowQuantity", m.MaxShowQuantity())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 208)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Float64Field("MaxShowQuantity", m.MaxShowQuantity())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 208 {
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
	m.SetPrice1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
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
	m.SetMaxShowQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 208 {
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
	m.SetPrice1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
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
	m.SetMaxShowQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 208 {
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
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
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
		case "MaxShowQuantity":
			m.SetMaxShowQuantity(r.Float64())
		case "Price1AsString":
			m.SetPrice1AsString(r.String())
		case "Price2AsString":
			m.SetPrice2AsString(r.String())
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

func (m *SubmitNewSingleOrderPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 208 {
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
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
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
		case "MaxShowQuantity":
			m.SetMaxShowQuantity(r.Float64())
		case "Price1AsString":
			m.SetPrice1AsString(r.String())
		case "Price2AsString":
			m.SetPrice2AsString(r.String())
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

func (m SubmitNewSingleOrderFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":208,\"F\":["...)
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
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxShowQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":208,\"F\":["...)
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
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxShowQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":208,\"F\":["...)
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
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxShowQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":208,\"F\":["...)
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
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxShowQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 208)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Float64Field("MaxShowQuantity", m.MaxShowQuantity())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 208)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Float64Field("MaxShowQuantity", m.MaxShowQuantity())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 208)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Float64Field("MaxShowQuantity", m.MaxShowQuantity())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

func (m SubmitNewSingleOrderFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 208)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Float64Field("MaxShowQuantity", m.MaxShowQuantity())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 208 {
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
	m.SetPrice1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
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
	m.SetMaxShowQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 208 {
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
	m.SetPrice1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
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
	m.SetMaxShowQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 208 {
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
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
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
		case "MaxShowQuantity":
			m.SetMaxShowQuantity(r.Float64())
		case "Price1AsString":
			m.SetPrice1AsString(r.String())
		case "Price2AsString":
			m.SetPrice2AsString(r.String())
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

func (m *SubmitNewSingleOrderFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 208 {
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
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
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
		case "MaxShowQuantity":
			m.SetMaxShowQuantity(r.Float64())
		case "Price1AsString":
			m.SetPrice1AsString(r.String())
		case "Price2AsString":
			m.SetPrice2AsString(r.String())
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

func (m *SubmitNewSingleOrderBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice1(r.ReadFloat64())
		case 8:
			m.SetPrice2(r.ReadFloat64())
		case 9:
			m.SetQuantity(r.ReadFloat64())
		case 10:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 11:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 12:
			m.SetIsAutomatedOrder(r.ReadBool())
		case 13:
			m.SetIsParentOrder(r.ReadBool())
		case 14:
			m.SetFreeFormText(r.ReadString())
		case 15:
			m.SetOpenOrClose(OpenCloseTradeEnum(r.ReadInt32()))
		case 16:
			m.SetMaxShowQuantity(r.ReadFloat64())
		case 17:
			m.SetPrice1AsString(r.ReadString())
		case 18:
			m.SetPrice2AsString(r.ReadString())
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

func (m *SubmitNewSingleOrderPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice1(r.ReadFloat64())
		case 8:
			m.SetPrice2(r.ReadFloat64())
		case 9:
			m.SetQuantity(r.ReadFloat64())
		case 10:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 11:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 12:
			m.SetIsAutomatedOrder(r.ReadBool())
		case 13:
			m.SetIsParentOrder(r.ReadBool())
		case 14:
			m.SetFreeFormText(r.ReadString())
		case 15:
			m.SetOpenOrClose(OpenCloseTradeEnum(r.ReadInt32()))
		case 16:
			m.SetMaxShowQuantity(r.ReadFloat64())
		case 17:
			m.SetPrice1AsString(r.ReadString())
		case 18:
			m.SetPrice2AsString(r.ReadString())
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

func (m SubmitNewSingleOrderBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 208)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.OrderType()))
	w.WriteVarint32(6, int32(m.BuySell()))
	w.WriteFixed64Float64(7, m.Price1())
	w.WriteFixed64Float64(8, m.Price2())
	w.WriteFixed64Float64(9, m.Quantity())
	w.WriteVarint32(10, int32(m.TimeInForce()))
	w.WriteFixed64Int64(11, int64(m.GoodTillDateTime()))
	w.WriteBool(12, m.IsAutomatedOrder())
	w.WriteBool(13, m.IsParentOrder())
	w.WriteString(14, m.FreeFormText())
	w.WriteVarint32(15, int32(m.OpenOrClose()))
	w.WriteFixed64Float64(16, m.MaxShowQuantity())
	w.WriteString(17, m.Price1AsString())
	w.WriteString(18, m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 208)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.OrderType()))
	w.WriteVarint32(6, int32(m.BuySell()))
	w.WriteFixed64Float64(7, m.Price1())
	w.WriteFixed64Float64(8, m.Price2())
	w.WriteFixed64Float64(9, m.Quantity())
	w.WriteVarint32(10, int32(m.TimeInForce()))
	w.WriteFixed64Int64(11, int64(m.GoodTillDateTime()))
	w.WriteBool(12, m.IsAutomatedOrder())
	w.WriteBool(13, m.IsParentOrder())
	w.WriteString(14, m.FreeFormText())
	w.WriteVarint32(15, int32(m.OpenOrClose()))
	w.WriteFixed64Float64(16, m.MaxShowQuantity())
	w.WriteString(17, m.Price1AsString())
	w.WriteString(18, m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice1(r.ReadFloat64())
		case 8:
			m.SetPrice2(r.ReadFloat64())
		case 9:
			m.SetQuantity(r.ReadFloat64())
		case 10:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 11:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 12:
			m.SetIsAutomatedOrder(r.ReadBool())
		case 13:
			m.SetIsParentOrder(r.ReadBool())
		case 14:
			m.SetFreeFormText(r.ReadString())
		case 15:
			m.SetOpenOrClose(OpenCloseTradeEnum(r.ReadInt32()))
		case 16:
			m.SetMaxShowQuantity(r.ReadFloat64())
		case 17:
			m.SetPrice1AsString(r.ReadString())
		case 18:
			m.SetPrice2AsString(r.ReadString())
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

func (m *SubmitNewSingleOrderFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice1(r.ReadFloat64())
		case 8:
			m.SetPrice2(r.ReadFloat64())
		case 9:
			m.SetQuantity(r.ReadFloat64())
		case 10:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 11:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 12:
			m.SetIsAutomatedOrder(r.ReadBool())
		case 13:
			m.SetIsParentOrder(r.ReadBool())
		case 14:
			m.SetFreeFormText(r.ReadString())
		case 15:
			m.SetOpenOrClose(OpenCloseTradeEnum(r.ReadInt32()))
		case 16:
			m.SetMaxShowQuantity(r.ReadFloat64())
		case 17:
			m.SetPrice1AsString(r.ReadString())
		case 18:
			m.SetPrice2AsString(r.ReadString())
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

func (m SubmitNewSingleOrderFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 208)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.OrderType()))
	w.WriteVarint32(6, int32(m.BuySell()))
	w.WriteFixed64Float64(7, m.Price1())
	w.WriteFixed64Float64(8, m.Price2())
	w.WriteFixed64Float64(9, m.Quantity())
	w.WriteVarint32(10, int32(m.TimeInForce()))
	w.WriteFixed64Int64(11, int64(m.GoodTillDateTime()))
	w.WriteBool(12, m.IsAutomatedOrder())
	w.WriteBool(13, m.IsParentOrder())
	w.WriteString(14, m.FreeFormText())
	w.WriteVarint32(15, int32(m.OpenOrClose()))
	w.WriteFixed64Float64(16, m.MaxShowQuantity())
	w.WriteString(17, m.Price1AsString())
	w.WriteString(18, m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewSingleOrderFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 208)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.OrderType()))
	w.WriteVarint32(6, int32(m.BuySell()))
	w.WriteFixed64Float64(7, m.Price1())
	w.WriteFixed64Float64(8, m.Price2())
	w.WriteFixed64Float64(9, m.Quantity())
	w.WriteVarint32(10, int32(m.TimeInForce()))
	w.WriteFixed64Int64(11, int64(m.GoodTillDateTime()))
	w.WriteBool(12, m.IsAutomatedOrder())
	w.WriteBool(13, m.IsParentOrder())
	w.WriteString(14, m.FreeFormText())
	w.WriteVarint32(15, int32(m.OpenOrClose()))
	w.WriteFixed64Float64(16, m.MaxShowQuantity())
	w.WriteString(17, m.Price1AsString())
	w.WriteString(18, m.Price2AsString())
	return w.Finish(), nil
}
