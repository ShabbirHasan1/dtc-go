//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m OrderUpdateFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":301,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.PreviousServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderStatus()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderUpdateReason()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderType()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.TimeInForce()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.GoodTillDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.OrderQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.FilledQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.RemainingQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AverageFillPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastFillPrice())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.LastFillDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.LastFillQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.LastFillExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoOrders())
	w.Data = append(w.Data, ',')
	w.String(m.ParentServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.OCOLinkedOrderServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenOrClose()))
	w.Data = append(w.Data, ',')
	w.String(m.PreviousClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.OrderReceivedDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.LatestTransactionDateTime()))
	return w.Finish(), nil
}

func (m OrderUpdateFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":301,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.PreviousServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderStatus()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderUpdateReason()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OrderType()))
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.TimeInForce()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.GoodTillDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.OrderQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.FilledQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.RemainingQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AverageFillPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastFillPrice())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.LastFillDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.LastFillQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.LastFillExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoOrders())
	w.Data = append(w.Data, ',')
	w.String(m.ParentServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.OCOLinkedOrderServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.OpenOrClose()))
	w.Data = append(w.Data, ',')
	w.String(m.PreviousClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.OrderReceivedDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.LatestTransactionDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m OrderUpdateFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 301)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumMessages", m.TotalNumMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("PreviousServerOrderID", m.PreviousServerOrderID())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("ExchangeOrderID", m.ExchangeOrderID())
	w.Int32Field("OrderStatus", int32(m.OrderStatus()))
	w.Int32Field("OrderUpdateReason", int32(m.OrderUpdateReason()))
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Float64Field("OrderQuantity", m.OrderQuantity())
	w.Float64Field("FilledQuantity", m.FilledQuantity())
	w.Float64Field("RemainingQuantity", m.RemainingQuantity())
	w.Float64Field("AverageFillPrice", m.AverageFillPrice())
	w.Float64Field("LastFillPrice", m.LastFillPrice())
	w.Int64Field("LastFillDateTime", int64(m.LastFillDateTime()))
	w.Float64Field("LastFillQuantity", m.LastFillQuantity())
	w.StringField("LastFillExecutionID", m.LastFillExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("InfoText", m.InfoText())
	w.Uint8Field("NoOrders", m.NoOrders())
	w.StringField("ParentServerOrderID", m.ParentServerOrderID())
	w.StringField("OCOLinkedOrderServerOrderID", m.OCOLinkedOrderServerOrderID())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.StringField("PreviousClientOrderID", m.PreviousClientOrderID())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int64Field("OrderReceivedDateTime", int64(m.OrderReceivedDateTime()))
	w.Float64Field("LatestTransactionDateTime", float64(m.LatestTransactionDateTime()))
	return w.Finish(), nil
}

func (m OrderUpdateFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 301)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumMessages", m.TotalNumMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("PreviousServerOrderID", m.PreviousServerOrderID())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("ExchangeOrderID", m.ExchangeOrderID())
	w.Int32Field("OrderStatus", int32(m.OrderStatus()))
	w.Int32Field("OrderUpdateReason", int32(m.OrderUpdateReason()))
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Float64Field("OrderQuantity", m.OrderQuantity())
	w.Float64Field("FilledQuantity", m.FilledQuantity())
	w.Float64Field("RemainingQuantity", m.RemainingQuantity())
	w.Float64Field("AverageFillPrice", m.AverageFillPrice())
	w.Float64Field("LastFillPrice", m.LastFillPrice())
	w.Int64Field("LastFillDateTime", int64(m.LastFillDateTime()))
	w.Float64Field("LastFillQuantity", m.LastFillQuantity())
	w.StringField("LastFillExecutionID", m.LastFillExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("InfoText", m.InfoText())
	w.Uint8Field("NoOrders", m.NoOrders())
	w.StringField("ParentServerOrderID", m.ParentServerOrderID())
	w.StringField("OCOLinkedOrderServerOrderID", m.OCOLinkedOrderServerOrderID())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.StringField("PreviousClientOrderID", m.PreviousClientOrderID())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int64Field("OrderReceivedDateTime", int64(m.OrderReceivedDateTime()))
	w.Float64Field("LatestTransactionDateTime", float64(m.LatestTransactionDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *OrderUpdateFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 301 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumMessages(r.Int32())
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
	m.SetPreviousServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchangeOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderStatus(OrderStatusEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderUpdateReason(OrderUpdateReasonEnum(r.Int32()))
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
	m.SetTimeInForce(TimeInForceEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetGoodTillDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetFilledQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRemainingQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAverageFillPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastFillPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastFillDateTime(DateTimeWithMillisecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetLastFillQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastFillExecutionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoOrders(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetParentServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOLinkedOrderServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenOrClose(OpenCloseTradeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPreviousClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetFreeFormText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderReceivedDateTime(DateTimeWithMillisecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetLatestTransactionDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *OrderUpdateFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 301 {
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
		case "TotalNumMessages":
			m.SetTotalNumMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "PreviousServerOrderID":
			m.SetPreviousServerOrderID(r.String())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "ExchangeOrderID":
			m.SetExchangeOrderID(r.String())
		case "OrderStatus":
			m.SetOrderStatus(OrderStatusEnum(r.Int32()))
		case "OrderUpdateReason":
			m.SetOrderUpdateReason(OrderUpdateReasonEnum(r.Int32()))
		case "OrderType":
			m.SetOrderType(OrderTypeEnum(r.Int32()))
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "Price1":
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "TimeInForce":
			m.SetTimeInForce(TimeInForceEnum(r.Int32()))
		case "GoodTillDateTime":
			m.SetGoodTillDateTime(DateTime(r.Int64()))
		case "OrderQuantity":
			m.SetOrderQuantity(r.Float64())
		case "FilledQuantity":
			m.SetFilledQuantity(r.Float64())
		case "RemainingQuantity":
			m.SetRemainingQuantity(r.Float64())
		case "AverageFillPrice":
			m.SetAverageFillPrice(r.Float64())
		case "LastFillPrice":
			m.SetLastFillPrice(r.Float64())
		case "LastFillDateTime":
			m.SetLastFillDateTime(DateTimeWithMillisecondsInt(r.Int64()))
		case "LastFillQuantity":
			m.SetLastFillQuantity(r.Float64())
		case "LastFillExecutionID":
			m.SetLastFillExecutionID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "InfoText":
			m.SetInfoText(r.String())
		case "NoOrders":
			m.SetNoOrders(r.Uint8())
		case "ParentServerOrderID":
			m.SetParentServerOrderID(r.String())
		case "OCOLinkedOrderServerOrderID":
			m.SetOCOLinkedOrderServerOrderID(r.String())
		case "OpenOrClose":
			m.SetOpenOrClose(OpenCloseTradeEnum(r.Int32()))
		case "PreviousClientOrderID":
			m.SetPreviousClientOrderID(r.String())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
		case "OrderReceivedDateTime":
			m.SetOrderReceivedDateTime(DateTimeWithMillisecondsInt(r.Int64()))
		case "LatestTransactionDateTime":
			m.SetLatestTransactionDateTime(DateTimeWithMilliseconds(r.Float64()))
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
