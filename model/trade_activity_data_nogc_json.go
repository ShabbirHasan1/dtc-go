//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeActivityDataPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10114,\"F\":["...)
	w.Uint8(m.ActivityType())
	w.Data = append(w.Data, ',')
	w.Int64(m.DataDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.OrderActionSource())
	w.Data = append(w.Data, ',')
	w.Uint64(m.InternalOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ServiceOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.FIXClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.OrderTypeName())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BuySell())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NewOrderStatus())
	w.Data = append(w.Data, ',')
	w.Float64(m.FillPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.OrderFilledQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint64(m.ParentInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OpenClose())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSimulated())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsChartReplaying())
	w.Data = append(w.Data, ',')
	w.String(m.FillExecutionServiceID())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.SourceChartNumber())
	w.Data = append(w.Data, ',')
	w.String(m.SourceChartbookFileName())
	w.Data = append(w.Data, ',')
	w.Int32(m.TimeInForce())
	w.Data = append(w.Data, ',')
	w.String(m.SymbolServiceCode())
	w.Data = append(w.Data, ',')
	w.String(m.Note())
	w.Data = append(w.Data, ',')
	w.String(m.OriginatingClientUsername())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeAccountBalance())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportsPositionQuantityField())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsBillable())
	w.Data = append(w.Data, ',')
	w.Int32(m.QuantityForBilling())
	w.Data = append(w.Data, ',')
	w.String(m.OrderRoutingServiceCode())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Int64(m.AuditTrail_TransactDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Int32(m.AuditTrail_MessageDirection())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_OperatorID())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_SessionID())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_ExecutingFirmID())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_FixMessageType())
	w.Data = append(w.Data, ',')
	w.Int16(m.AuditTrail_CustomerTypeIndicator())
	w.Data = append(w.Data, ',')
	w.Int16(m.AuditTrail_CustomerOrFirm())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_ExecutionReportID())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_SpreadLegLinkID())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_SecurityDesc())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_MarketSegmentID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.AuditTrail_IFMFlag())
	w.Data = append(w.Data, ',')
	w.Float64(m.AuditTrail_DisplayQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_CountryOfOrigin())
	w.Data = append(w.Data, ',')
	w.Float64(m.AuditTrail_FillQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AuditTrail_RemainingQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.AuditTrail_AggressorFlag())
	w.Data = append(w.Data, ',')
	w.Int32(m.AuditTrail_SourceOfCancellation())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_OrdRejReason())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	return w.Finish(), nil
}

func (m TradeActivityDataPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10114,\"F\":["...)
	w.Uint8(m.ActivityType())
	w.Data = append(w.Data, ',')
	w.Int64(m.DataDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.OrderActionSource())
	w.Data = append(w.Data, ',')
	w.Uint64(m.InternalOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ServiceOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.FIXClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.OrderTypeName())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BuySell())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NewOrderStatus())
	w.Data = append(w.Data, ',')
	w.Float64(m.FillPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.OrderFilledQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint64(m.ParentInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OpenClose())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSimulated())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsChartReplaying())
	w.Data = append(w.Data, ',')
	w.String(m.FillExecutionServiceID())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionQuantity())
	w.Data = append(w.Data, ',')
	w.Int32(m.SourceChartNumber())
	w.Data = append(w.Data, ',')
	w.String(m.SourceChartbookFileName())
	w.Data = append(w.Data, ',')
	w.Int32(m.TimeInForce())
	w.Data = append(w.Data, ',')
	w.String(m.SymbolServiceCode())
	w.Data = append(w.Data, ',')
	w.String(m.Note())
	w.Data = append(w.Data, ',')
	w.String(m.OriginatingClientUsername())
	w.Data = append(w.Data, ',')
	w.Float64(m.TradeAccountBalance())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportsPositionQuantityField())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsBillable())
	w.Data = append(w.Data, ',')
	w.Int32(m.QuantityForBilling())
	w.Data = append(w.Data, ',')
	w.String(m.OrderRoutingServiceCode())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Int64(m.AuditTrail_TransactDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Int32(m.AuditTrail_MessageDirection())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_OperatorID())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_SessionID())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_ExecutingFirmID())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_FixMessageType())
	w.Data = append(w.Data, ',')
	w.Int16(m.AuditTrail_CustomerTypeIndicator())
	w.Data = append(w.Data, ',')
	w.Int16(m.AuditTrail_CustomerOrFirm())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_ExecutionReportID())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_SpreadLegLinkID())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_SecurityDesc())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_MarketSegmentID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.AuditTrail_IFMFlag())
	w.Data = append(w.Data, ',')
	w.Float64(m.AuditTrail_DisplayQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_CountryOfOrigin())
	w.Data = append(w.Data, ',')
	w.Float64(m.AuditTrail_FillQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AuditTrail_RemainingQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.AuditTrail_AggressorFlag())
	w.Data = append(w.Data, ',')
	w.Int32(m.AuditTrail_SourceOfCancellation())
	w.Data = append(w.Data, ',')
	w.String(m.AuditTrail_OrdRejReason())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeActivityDataPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10114)
	w.Uint8Field("ActivityType", m.ActivityType())
	w.Int64Field("DataDateTimeUTC", m.DataDateTimeUTC())
	w.StringField("Symbol", m.Symbol())
	w.StringField("OrderActionSource", m.OrderActionSource())
	w.Uint64Field("InternalOrderID", m.InternalOrderID())
	w.StringField("ServiceOrderID", m.ServiceOrderID())
	w.StringField("ExchangeOrderID", m.ExchangeOrderID())
	w.StringField("FIXClientOrderID", m.FIXClientOrderID())
	w.StringField("OrderTypeName", m.OrderTypeName())
	w.Float64Field("Quantity", m.Quantity())
	w.Uint8Field("BuySell", m.BuySell())
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Uint8Field("NewOrderStatus", m.NewOrderStatus())
	w.Float64Field("FillPrice", m.FillPrice())
	w.Float64Field("OrderFilledQuantity", m.OrderFilledQuantity())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("LastPriceDuringPosition", m.LastPriceDuringPosition())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint64Field("ParentInternalOrderID", m.ParentInternalOrderID())
	w.Uint8Field("OpenClose", m.OpenClose())
	w.Uint8Field("IsSimulated", m.IsSimulated())
	w.Uint8Field("IsAutomatedOrder", m.IsAutomatedOrder())
	w.Uint8Field("IsChartReplaying", m.IsChartReplaying())
	w.StringField("FillExecutionServiceID", m.FillExecutionServiceID())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	w.Int32Field("SourceChartNumber", m.SourceChartNumber())
	w.StringField("SourceChartbookFileName", m.SourceChartbookFileName())
	w.Int32Field("TimeInForce", m.TimeInForce())
	w.StringField("SymbolServiceCode", m.SymbolServiceCode())
	w.StringField("Note", m.Note())
	w.StringField("OriginatingClientUsername", m.OriginatingClientUsername())
	w.Float64Field("TradeAccountBalance", m.TradeAccountBalance())
	w.Uint8Field("SupportsPositionQuantityField", m.SupportsPositionQuantityField())
	w.Uint8Field("IsBillable", m.IsBillable())
	w.Int32Field("QuantityForBilling", m.QuantityForBilling())
	w.StringField("OrderRoutingServiceCode", m.OrderRoutingServiceCode())
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	w.Int64Field("AuditTrail_TransactDateTimeUTC", m.AuditTrail_TransactDateTimeUTC())
	w.Int32Field("AuditTrail_MessageDirection", m.AuditTrail_MessageDirection())
	w.StringField("AuditTrail_OperatorID", m.AuditTrail_OperatorID())
	w.StringField("AuditTrail_SelfMatchPreventionID", m.AuditTrail_SelfMatchPreventionID())
	w.StringField("AuditTrail_SessionID", m.AuditTrail_SessionID())
	w.StringField("AuditTrail_ExecutingFirmID", m.AuditTrail_ExecutingFirmID())
	w.StringField("AuditTrail_FixMessageType", m.AuditTrail_FixMessageType())
	w.Int16Field("AuditTrail_CustomerTypeIndicator", m.AuditTrail_CustomerTypeIndicator())
	w.Int16Field("AuditTrail_CustomerOrFirm", m.AuditTrail_CustomerOrFirm())
	w.StringField("AuditTrail_ExecutionReportID", m.AuditTrail_ExecutionReportID())
	w.StringField("AuditTrail_SpreadLegLinkID", m.AuditTrail_SpreadLegLinkID())
	w.StringField("AuditTrail_SecurityDesc", m.AuditTrail_SecurityDesc())
	w.StringField("AuditTrail_MarketSegmentID", m.AuditTrail_MarketSegmentID())
	w.Uint8Field("AuditTrail_IFMFlag", m.AuditTrail_IFMFlag())
	w.Float64Field("AuditTrail_DisplayQuantity", m.AuditTrail_DisplayQuantity())
	w.StringField("AuditTrail_CountryOfOrigin", m.AuditTrail_CountryOfOrigin())
	w.Float64Field("AuditTrail_FillQuantity", m.AuditTrail_FillQuantity())
	w.Float64Field("AuditTrail_RemainingQuantity", m.AuditTrail_RemainingQuantity())
	w.Uint8Field("AuditTrail_AggressorFlag", m.AuditTrail_AggressorFlag())
	w.Int32Field("AuditTrail_SourceOfCancellation", m.AuditTrail_SourceOfCancellation())
	w.StringField("AuditTrail_OrdRejReason", m.AuditTrail_OrdRejReason())
	w.Uint8Field("IsSnapshot", m.IsSnapshot())
	w.Uint8Field("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("IsLastMessageInBatch", m.IsLastMessageInBatch())
	return w.Finish(), nil
}

func (m TradeActivityDataPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10114)
	w.Uint8Field("ActivityType", m.ActivityType())
	w.Int64Field("DataDateTimeUTC", m.DataDateTimeUTC())
	w.StringField("Symbol", m.Symbol())
	w.StringField("OrderActionSource", m.OrderActionSource())
	w.Uint64Field("InternalOrderID", m.InternalOrderID())
	w.StringField("ServiceOrderID", m.ServiceOrderID())
	w.StringField("ExchangeOrderID", m.ExchangeOrderID())
	w.StringField("FIXClientOrderID", m.FIXClientOrderID())
	w.StringField("OrderTypeName", m.OrderTypeName())
	w.Float64Field("Quantity", m.Quantity())
	w.Uint8Field("BuySell", m.BuySell())
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Uint8Field("NewOrderStatus", m.NewOrderStatus())
	w.Float64Field("FillPrice", m.FillPrice())
	w.Float64Field("OrderFilledQuantity", m.OrderFilledQuantity())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("LastPriceDuringPosition", m.LastPriceDuringPosition())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint64Field("ParentInternalOrderID", m.ParentInternalOrderID())
	w.Uint8Field("OpenClose", m.OpenClose())
	w.Uint8Field("IsSimulated", m.IsSimulated())
	w.Uint8Field("IsAutomatedOrder", m.IsAutomatedOrder())
	w.Uint8Field("IsChartReplaying", m.IsChartReplaying())
	w.StringField("FillExecutionServiceID", m.FillExecutionServiceID())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	w.Int32Field("SourceChartNumber", m.SourceChartNumber())
	w.StringField("SourceChartbookFileName", m.SourceChartbookFileName())
	w.Int32Field("TimeInForce", m.TimeInForce())
	w.StringField("SymbolServiceCode", m.SymbolServiceCode())
	w.StringField("Note", m.Note())
	w.StringField("OriginatingClientUsername", m.OriginatingClientUsername())
	w.Float64Field("TradeAccountBalance", m.TradeAccountBalance())
	w.Uint8Field("SupportsPositionQuantityField", m.SupportsPositionQuantityField())
	w.Uint8Field("IsBillable", m.IsBillable())
	w.Int32Field("QuantityForBilling", m.QuantityForBilling())
	w.StringField("OrderRoutingServiceCode", m.OrderRoutingServiceCode())
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	w.Int64Field("AuditTrail_TransactDateTimeUTC", m.AuditTrail_TransactDateTimeUTC())
	w.Int32Field("AuditTrail_MessageDirection", m.AuditTrail_MessageDirection())
	w.StringField("AuditTrail_OperatorID", m.AuditTrail_OperatorID())
	w.StringField("AuditTrail_SelfMatchPreventionID", m.AuditTrail_SelfMatchPreventionID())
	w.StringField("AuditTrail_SessionID", m.AuditTrail_SessionID())
	w.StringField("AuditTrail_ExecutingFirmID", m.AuditTrail_ExecutingFirmID())
	w.StringField("AuditTrail_FixMessageType", m.AuditTrail_FixMessageType())
	w.Int16Field("AuditTrail_CustomerTypeIndicator", m.AuditTrail_CustomerTypeIndicator())
	w.Int16Field("AuditTrail_CustomerOrFirm", m.AuditTrail_CustomerOrFirm())
	w.StringField("AuditTrail_ExecutionReportID", m.AuditTrail_ExecutionReportID())
	w.StringField("AuditTrail_SpreadLegLinkID", m.AuditTrail_SpreadLegLinkID())
	w.StringField("AuditTrail_SecurityDesc", m.AuditTrail_SecurityDesc())
	w.StringField("AuditTrail_MarketSegmentID", m.AuditTrail_MarketSegmentID())
	w.Uint8Field("AuditTrail_IFMFlag", m.AuditTrail_IFMFlag())
	w.Float64Field("AuditTrail_DisplayQuantity", m.AuditTrail_DisplayQuantity())
	w.StringField("AuditTrail_CountryOfOrigin", m.AuditTrail_CountryOfOrigin())
	w.Float64Field("AuditTrail_FillQuantity", m.AuditTrail_FillQuantity())
	w.Float64Field("AuditTrail_RemainingQuantity", m.AuditTrail_RemainingQuantity())
	w.Uint8Field("AuditTrail_AggressorFlag", m.AuditTrail_AggressorFlag())
	w.Int32Field("AuditTrail_SourceOfCancellation", m.AuditTrail_SourceOfCancellation())
	w.StringField("AuditTrail_OrdRejReason", m.AuditTrail_OrdRejReason())
	w.Uint8Field("IsSnapshot", m.IsSnapshot())
	w.Uint8Field("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("IsLastMessageInBatch", m.IsLastMessageInBatch())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeActivityDataPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10114 {
		return message.ErrWrongType
	}
	m.SetActivityType(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDataDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderActionSource(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetServiceOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchangeOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetFIXClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderTypeName(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(r.Uint8())
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
	m.SetNewOrderStatus(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFillPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderFilledQuantity(r.Float64())
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
	m.SetLastPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetParentInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenClose(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSimulated(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsAutomatedOrder(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsChartReplaying(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFillExecutionServiceID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPositionQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetSourceChartNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSourceChartbookFileName(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTimeInForce(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolServiceCode(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNote(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOriginatingClientUsername(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccountBalance(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetSupportsPositionQuantityField(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsBillable(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantityForBilling(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderRoutingServiceCode(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_TransactDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_MessageDirection(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_OperatorID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_SelfMatchPreventionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_SessionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_ExecutingFirmID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_FixMessageType(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_CustomerTypeIndicator(r.Int16())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_CustomerOrFirm(r.Int16())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_ExecutionReportID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_SpreadLegLinkID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_SecurityDesc(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_MarketSegmentID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_IFMFlag(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_DisplayQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_CountryOfOrigin(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_FillQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_RemainingQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_AggressorFlag(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_SourceOfCancellation(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAuditTrail_OrdRejReason(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSnapshot(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFirstMessageInBatch(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeActivityDataPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10114 {
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
		case "ActivityType":
			m.SetActivityType(r.Uint8())
		case "DataDateTimeUTC":
			m.SetDataDateTimeUTC(r.Int64())
		case "Symbol":
			m.SetSymbol(r.String())
		case "OrderActionSource":
			m.SetOrderActionSource(r.String())
		case "InternalOrderID":
			m.SetInternalOrderID(r.Uint64())
		case "ServiceOrderID":
			m.SetServiceOrderID(r.String())
		case "ExchangeOrderID":
			m.SetExchangeOrderID(r.String())
		case "FIXClientOrderID":
			m.SetFIXClientOrderID(r.String())
		case "OrderTypeName":
			m.SetOrderTypeName(r.String())
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "BuySell":
			m.SetBuySell(r.Uint8())
		case "Price1":
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "NewOrderStatus":
			m.SetNewOrderStatus(r.Uint8())
		case "FillPrice":
			m.SetFillPrice(r.Float64())
		case "OrderFilledQuantity":
			m.SetOrderFilledQuantity(r.Float64())
		case "HighPriceDuringPosition":
			m.SetHighPriceDuringPosition(r.Float64())
		case "LowPriceDuringPosition":
			m.SetLowPriceDuringPosition(r.Float64())
		case "LastPriceDuringPosition":
			m.SetLastPriceDuringPosition(r.Float64())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "ParentInternalOrderID":
			m.SetParentInternalOrderID(r.Uint64())
		case "OpenClose":
			m.SetOpenClose(r.Uint8())
		case "IsSimulated":
			m.SetIsSimulated(r.Uint8())
		case "IsAutomatedOrder":
			m.SetIsAutomatedOrder(r.Uint8())
		case "IsChartReplaying":
			m.SetIsChartReplaying(r.Uint8())
		case "FillExecutionServiceID":
			m.SetFillExecutionServiceID(r.String())
		case "PositionQuantity":
			m.SetPositionQuantity(r.Float64())
		case "SourceChartNumber":
			m.SetSourceChartNumber(r.Int32())
		case "SourceChartbookFileName":
			m.SetSourceChartbookFileName(r.String())
		case "TimeInForce":
			m.SetTimeInForce(r.Int32())
		case "SymbolServiceCode":
			m.SetSymbolServiceCode(r.String())
		case "Note":
			m.SetNote(r.String())
		case "OriginatingClientUsername":
			m.SetOriginatingClientUsername(r.String())
		case "TradeAccountBalance":
			m.SetTradeAccountBalance(r.Float64())
		case "SupportsPositionQuantityField":
			m.SetSupportsPositionQuantityField(r.Uint8())
		case "IsBillable":
			m.SetIsBillable(r.Uint8())
		case "QuantityForBilling":
			m.SetQuantityForBilling(r.Int32())
		case "OrderRoutingServiceCode":
			m.SetOrderRoutingServiceCode(r.String())
		case "SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
		case "AuditTrail_TransactDateTimeUTC":
			m.SetAuditTrail_TransactDateTimeUTC(r.Int64())
		case "AuditTrail_MessageDirection":
			m.SetAuditTrail_MessageDirection(r.Int32())
		case "AuditTrail_OperatorID":
			m.SetAuditTrail_OperatorID(r.String())
		case "AuditTrail_SelfMatchPreventionID":
			m.SetAuditTrail_SelfMatchPreventionID(r.String())
		case "AuditTrail_SessionID":
			m.SetAuditTrail_SessionID(r.String())
		case "AuditTrail_ExecutingFirmID":
			m.SetAuditTrail_ExecutingFirmID(r.String())
		case "AuditTrail_FixMessageType":
			m.SetAuditTrail_FixMessageType(r.String())
		case "AuditTrail_CustomerTypeIndicator":
			m.SetAuditTrail_CustomerTypeIndicator(r.Int16())
		case "AuditTrail_CustomerOrFirm":
			m.SetAuditTrail_CustomerOrFirm(r.Int16())
		case "AuditTrail_ExecutionReportID":
			m.SetAuditTrail_ExecutionReportID(r.String())
		case "AuditTrail_SpreadLegLinkID":
			m.SetAuditTrail_SpreadLegLinkID(r.String())
		case "AuditTrail_SecurityDesc":
			m.SetAuditTrail_SecurityDesc(r.String())
		case "AuditTrail_MarketSegmentID":
			m.SetAuditTrail_MarketSegmentID(r.String())
		case "AuditTrail_IFMFlag":
			m.SetAuditTrail_IFMFlag(r.Uint8())
		case "AuditTrail_DisplayQuantity":
			m.SetAuditTrail_DisplayQuantity(r.Float64())
		case "AuditTrail_CountryOfOrigin":
			m.SetAuditTrail_CountryOfOrigin(r.String())
		case "AuditTrail_FillQuantity":
			m.SetAuditTrail_FillQuantity(r.Float64())
		case "AuditTrail_RemainingQuantity":
			m.SetAuditTrail_RemainingQuantity(r.Float64())
		case "AuditTrail_AggressorFlag":
			m.SetAuditTrail_AggressorFlag(r.Uint8())
		case "AuditTrail_SourceOfCancellation":
			m.SetAuditTrail_SourceOfCancellation(r.Int32())
		case "AuditTrail_OrdRejReason":
			m.SetAuditTrail_OrdRejReason(r.String())
		case "IsSnapshot":
			m.SetIsSnapshot(r.Uint8())
		case "IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Uint8())
		case "IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Uint8())
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
