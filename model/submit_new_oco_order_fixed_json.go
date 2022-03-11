//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewOCOOrderFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":201,\"F\":["...)
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
	w.Float64(m.Price1_1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2_1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity_1())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID_2())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderType_2()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell_2()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1_2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2_2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity_2())
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
	w.Int32(int32(m.OpenOrClose()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.PartialFillHandling()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseOffsets())
	w.Data = append(w.Data, ',')
	w.Float64(m.OffsetFromParent1())
	w.Data = append(w.Data, ',')
	w.Float64(m.OffsetFromParent2())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MaintainSamePricesOnParentFill())
	w.Data = append(w.Data, ',')
	w.String(m.Price1_1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2_1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price1_2AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2_2AsString())
	return w.Finish(), nil
}

func (m SubmitNewOCOOrderFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":201,\"F\":["...)
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
	w.Float64(m.Price1_1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2_1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity_1())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID_2())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderType_2()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell_2()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1_2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2_2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity_2())
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
	w.Int32(int32(m.OpenOrClose()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.PartialFillHandling()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseOffsets())
	w.Data = append(w.Data, ',')
	w.Float64(m.OffsetFromParent1())
	w.Data = append(w.Data, ',')
	w.Float64(m.OffsetFromParent2())
	w.Data = append(w.Data, ',')
	w.Uint8(m.MaintainSamePricesOnParentFill())
	w.Data = append(w.Data, ',')
	w.String(m.Price1_1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2_1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price1_2AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2_2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitNewOCOOrderFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 201)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ClientOrderID_1", m.ClientOrderID_1())
	w.Int32Field("OrderType_1", int32(m.OrderType_1()))
	w.Int32Field("BuySell_1", int32(m.BuySell_1()))
	w.Float64Field("Price1_1", m.Price1_1())
	w.Float64Field("Price2_1", m.Price2_1())
	w.Float64Field("Quantity_1", m.Quantity_1())
	w.StringField("ClientOrderID_2", m.ClientOrderID_2())
	w.Int32Field("OrderType_2", int32(m.OrderType_2()))
	w.Int32Field("BuySell_2", int32(m.BuySell_2()))
	w.Float64Field("Price1_2", m.Price1_2())
	w.Float64Field("Price2_2", m.Price2_2())
	w.Float64Field("Quantity_2", m.Quantity_2())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("IsAutomatedOrder", m.IsAutomatedOrder())
	w.StringField("ParentTriggerClientOrderID", m.ParentTriggerClientOrderID())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Int8Field("PartialFillHandling", int8(m.PartialFillHandling()))
	w.Uint8Field("UseOffsets", m.UseOffsets())
	w.Float64Field("OffsetFromParent1", m.OffsetFromParent1())
	w.Float64Field("OffsetFromParent2", m.OffsetFromParent2())
	w.Uint8Field("MaintainSamePricesOnParentFill", m.MaintainSamePricesOnParentFill())
	w.StringField("Price1_1AsString", m.Price1_1AsString())
	w.StringField("Price2_1AsString", m.Price2_1AsString())
	w.StringField("Price1_2AsString", m.Price1_2AsString())
	w.StringField("Price2_2AsString", m.Price2_2AsString())
	return w.Finish(), nil
}

func (m SubmitNewOCOOrderFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 201)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ClientOrderID_1", m.ClientOrderID_1())
	w.Int32Field("OrderType_1", int32(m.OrderType_1()))
	w.Int32Field("BuySell_1", int32(m.BuySell_1()))
	w.Float64Field("Price1_1", m.Price1_1())
	w.Float64Field("Price2_1", m.Price2_1())
	w.Float64Field("Quantity_1", m.Quantity_1())
	w.StringField("ClientOrderID_2", m.ClientOrderID_2())
	w.Int32Field("OrderType_2", int32(m.OrderType_2()))
	w.Int32Field("BuySell_2", int32(m.BuySell_2()))
	w.Float64Field("Price1_2", m.Price1_2())
	w.Float64Field("Price2_2", m.Price2_2())
	w.Float64Field("Quantity_2", m.Quantity_2())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("IsAutomatedOrder", m.IsAutomatedOrder())
	w.StringField("ParentTriggerClientOrderID", m.ParentTriggerClientOrderID())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Int8Field("PartialFillHandling", int8(m.PartialFillHandling()))
	w.Uint8Field("UseOffsets", m.UseOffsets())
	w.Float64Field("OffsetFromParent1", m.OffsetFromParent1())
	w.Float64Field("OffsetFromParent2", m.OffsetFromParent2())
	w.Uint8Field("MaintainSamePricesOnParentFill", m.MaintainSamePricesOnParentFill())
	w.StringField("Price1_1AsString", m.Price1_1AsString())
	w.StringField("Price2_1AsString", m.Price2_1AsString())
	w.StringField("Price1_2AsString", m.Price1_2AsString())
	w.StringField("Price2_2AsString", m.Price2_2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewOCOOrderFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 201 {
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
	m.SetPrice1_1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2_1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity_1(r.Float64())
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
	m.SetPrice1_2(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2_2(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity_2(r.Float64())
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
	m.SetOpenOrClose(OpenCloseTradeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPartialFillHandling(PartialFillHandlingEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUseOffsets(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetOffsetFromParent1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOffsetFromParent2(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaintainSamePricesOnParentFill(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1_1AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2_1AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1_2AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2_2AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewOCOOrderFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 201 {
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
			m.SetPrice1_1(r.Float64())
		case "Price2_1":
			m.SetPrice2_1(r.Float64())
		case "Quantity_1":
			m.SetQuantity_1(r.Float64())
		case "ClientOrderID_2":
			m.SetClientOrderID_2(r.String())
		case "OrderType_2":
			m.SetOrderType_2(OrderTypeEnum(r.Int32()))
		case "BuySell_2":
			m.SetBuySell_2(BuySellEnum(r.Int32()))
		case "Price1_2":
			m.SetPrice1_2(r.Float64())
		case "Price2_2":
			m.SetPrice2_2(r.Float64())
		case "Quantity_2":
			m.SetQuantity_2(r.Float64())
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
		case "OpenOrClose":
			m.SetOpenOrClose(OpenCloseTradeEnum(r.Int32()))
		case "PartialFillHandling":
			m.SetPartialFillHandling(PartialFillHandlingEnum(r.Int8()))
		case "UseOffsets":
			m.SetUseOffsets(r.Uint8())
		case "OffsetFromParent1":
			m.SetOffsetFromParent1(r.Float64())
		case "OffsetFromParent2":
			m.SetOffsetFromParent2(r.Float64())
		case "MaintainSamePricesOnParentFill":
			m.SetMaintainSamePricesOnParentFill(r.Uint8())
		case "Price1_1AsString":
			m.SetPrice1_1AsString(r.String())
		case "Price2_1AsString":
			m.SetPrice2_1AsString(r.String())
		case "Price1_2AsString":
			m.SetPrice1_2AsString(r.String())
		case "Price2_2AsString":
			m.SetPrice2_2AsString(r.String())
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
