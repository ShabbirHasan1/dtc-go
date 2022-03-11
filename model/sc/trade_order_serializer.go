//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeOrder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10110,\"F\":["...)
	w.Bool(m.IsOrderDeleted())
	w.Data = append(w.Data, ',')
	w.Uint64(m.InternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OrderStatusCode())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OrderStatusBeforePendingModify())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OrderStatusBeforePendingCancel())
	w.Data = append(w.Data, ',')
	w.String(m.ServiceOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ActualSymbol())
	w.Data = append(w.Data, ',')
	w.Int32(m.OrderType())
	w.Data = append(w.Data, ',')
	w.Uint16(m.BuySell())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.OrderQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.FilledQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AverageFillPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.RealtimeFillStatus())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsRestingOrderDuringFill())
	w.Data = append(w.Data, ',')
	w.Int32(m.OrderRejectType())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Int32(m.InternalOrderIDModifierForService())
	w.Data = append(w.Data, ',')
	w.String(m.FIXClientOrderID())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SequenceNumberBasedClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderIDForDTCServer())
	w.Data = append(w.Data, ',')
	w.String(m.PreviousClientOrderIDForDTCServer())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.OriginatingClientUsername())
	w.Data = append(w.Data, ',')
	w.Int64(m.EntryDateTime())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastActionDateTime())
	w.Data = append(w.Data, ',')
	w.Int64(m.ServiceUpdateDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OrderEntryTimeForService())
	w.Data = append(w.Data, ',')
	w.Uint32(m.LastModifyTimeForService())
	w.Data = append(w.Data, ',')
	w.Int64(m.GoodTillDateTime())
	w.Data = append(w.Data, ',')
	w.Int32(m.TimeInForce())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OpenClose())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailStopOffset1())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailStopStep())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailTriggerPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingStopTriggerOffset())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TrailTriggerHit())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailToBreakEvenStopOffset())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaximumChaseAmountAsPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.InitialChaseOrderPrice1())
	w.Data = append(w.Data, ',')
	w.Float64(m.InitialLastTradePriceForChaseOrders())
	w.Data = append(w.Data, ',')
	w.Int32(m.TrailingStopTriggerOCOGroupNumber())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastModifyPrice1())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastModifyQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.CumulativeOrderQuantityFromParentFills())
	w.Data = append(w.Data, ',')
	w.Float64(m.PriorFilledQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.TickSize())
	w.Data = append(w.Data, ',')
	w.Int32(m.ValueFormat())
	w.Data = append(w.Data, ',')
	w.Float32(m.PriceMultiplier())
	w.Data = append(w.Data, ',')
	w.Uint64(m.ParentInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TargetChildInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.StopChildInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.AttachedOrderPriceOffset1())
	w.Data = append(w.Data, ',')
	w.Uint64(m.LinkInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OCOGroupInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OCOSiblingInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DisableChildAndSiblingRelatedActions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OCOManagedByService())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BracketOrderServerManaged())
	w.Data = append(w.Data, ',')
	w.String(m.LastOrderActionSource())
	w.Data = append(w.Data, ',')
	w.Uint8(m.StopLimitOrderStopPriceTriggered())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OCOFullSiblingCancelOnPartialFill())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReverseOnCompleteFill())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportScaleIn())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportScaleOut())
	w.Data = append(w.Data, ',')
	w.Int32(m.SourceChartNumber())
	w.Data = append(w.Data, ',')
	w.String(m.SourceChartbookFileName())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SimulatedOrder())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsChartReplaying())
	w.Data = append(w.Data, ',')
	w.Int32(m.AttachedOrderOCOGroupNumber())
	w.Data = append(w.Data, ',')
	w.String(m.LastFillExecutionServiceID())
	w.Data = append(w.Data, ',')
	w.Int32(m.FillCount())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastFillQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastFillPrice())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastFillDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint64(m.RejectedStopOCOSiblingInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.RejectedStopReplacementMarketOrderQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.EvaluatingForFill())
	w.Data = append(w.Data, ',')
	w.Uint32(m.LastProcessedTimeSalesRecordSequenceForPrices())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsMarketDataManagementOfOrderEnabled())
	w.Data = append(w.Data, ',')
	w.String(m.TextTag())
	w.Data = append(w.Data, ',')
	w.Int64(m.TimedOutOrderRequestedStatusDateTime())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequestedStatusForTimedOutOrder())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityToIncreaseFromParentFill())
	w.Data = append(w.Data, ',')
	w.Float64(m.MoveToBreakevenStopReferencePrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityTriggeredStop_QuantityForTrigger())
	w.Data = append(w.Data, ',')
	w.Float64(m.AccumulatedTradeVolumeForTriggeredStop())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BidAskQuantityStopInitialTriggerMet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NeedToOverrideLock())
	w.Data = append(w.Data, ',')
	w.Float64(m.CurrentMarketPrice())
	w.Data = append(w.Data, ',')
	w.Int64(m.CurrentMarketDateTime())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportOrderFillBilling())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsBillable())
	w.Data = append(w.Data, ',')
	w.Int32(m.QuantityForBilling())
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumberOfFailedOrderModifications())
	w.Data = append(w.Data, ',')
	w.Int32(m.DTCServerIndex())
	w.Data = append(w.Data, ',')
	w.String(m.ClearingFirmID())
	w.Data = append(w.Data, ',')
	w.String(m.SenderSubID())
	w.Data = append(w.Data, ',')
	w.String(m.SenderLocationId())
	w.Data = append(w.Data, ',')
	w.String(m.SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.Int32(m.CTICode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ObtainOrderActionDateTimeFromLastTradeTimeInChart())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaximumShowQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OrderSubmitted())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Int64(m.ExternalLastActionDateTimeUTC())
	return w.Finish(), nil
}

func (m TradeOrderBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10110,\"F\":["...)
	w.Bool(m.IsOrderDeleted())
	w.Data = append(w.Data, ',')
	w.Uint64(m.InternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OrderStatusCode())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OrderStatusBeforePendingModify())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OrderStatusBeforePendingCancel())
	w.Data = append(w.Data, ',')
	w.String(m.ServiceOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ActualSymbol())
	w.Data = append(w.Data, ',')
	w.Int32(m.OrderType())
	w.Data = append(w.Data, ',')
	w.Uint16(m.BuySell())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.OrderQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.FilledQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AverageFillPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.RealtimeFillStatus())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsRestingOrderDuringFill())
	w.Data = append(w.Data, ',')
	w.Int32(m.OrderRejectType())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Int32(m.InternalOrderIDModifierForService())
	w.Data = append(w.Data, ',')
	w.String(m.FIXClientOrderID())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SequenceNumberBasedClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderIDForDTCServer())
	w.Data = append(w.Data, ',')
	w.String(m.PreviousClientOrderIDForDTCServer())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.OriginatingClientUsername())
	w.Data = append(w.Data, ',')
	w.Int64(m.EntryDateTime())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastActionDateTime())
	w.Data = append(w.Data, ',')
	w.Int64(m.ServiceUpdateDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OrderEntryTimeForService())
	w.Data = append(w.Data, ',')
	w.Uint32(m.LastModifyTimeForService())
	w.Data = append(w.Data, ',')
	w.Int64(m.GoodTillDateTime())
	w.Data = append(w.Data, ',')
	w.Int32(m.TimeInForce())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OpenClose())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailStopOffset1())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailStopStep())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailTriggerPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingStopTriggerOffset())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TrailTriggerHit())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailToBreakEvenStopOffset())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaximumChaseAmountAsPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.InitialChaseOrderPrice1())
	w.Data = append(w.Data, ',')
	w.Float64(m.InitialLastTradePriceForChaseOrders())
	w.Data = append(w.Data, ',')
	w.Int32(m.TrailingStopTriggerOCOGroupNumber())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastModifyPrice1())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastModifyQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.CumulativeOrderQuantityFromParentFills())
	w.Data = append(w.Data, ',')
	w.Float64(m.PriorFilledQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.TickSize())
	w.Data = append(w.Data, ',')
	w.Int32(m.ValueFormat())
	w.Data = append(w.Data, ',')
	w.Float32(m.PriceMultiplier())
	w.Data = append(w.Data, ',')
	w.Uint64(m.ParentInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TargetChildInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.StopChildInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.AttachedOrderPriceOffset1())
	w.Data = append(w.Data, ',')
	w.Uint64(m.LinkInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OCOGroupInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OCOSiblingInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DisableChildAndSiblingRelatedActions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OCOManagedByService())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BracketOrderServerManaged())
	w.Data = append(w.Data, ',')
	w.String(m.LastOrderActionSource())
	w.Data = append(w.Data, ',')
	w.Uint8(m.StopLimitOrderStopPriceTriggered())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OCOFullSiblingCancelOnPartialFill())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReverseOnCompleteFill())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportScaleIn())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportScaleOut())
	w.Data = append(w.Data, ',')
	w.Int32(m.SourceChartNumber())
	w.Data = append(w.Data, ',')
	w.String(m.SourceChartbookFileName())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SimulatedOrder())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsChartReplaying())
	w.Data = append(w.Data, ',')
	w.Int32(m.AttachedOrderOCOGroupNumber())
	w.Data = append(w.Data, ',')
	w.String(m.LastFillExecutionServiceID())
	w.Data = append(w.Data, ',')
	w.Int32(m.FillCount())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastFillQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastFillPrice())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastFillDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint64(m.RejectedStopOCOSiblingInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.RejectedStopReplacementMarketOrderQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.EvaluatingForFill())
	w.Data = append(w.Data, ',')
	w.Uint32(m.LastProcessedTimeSalesRecordSequenceForPrices())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsMarketDataManagementOfOrderEnabled())
	w.Data = append(w.Data, ',')
	w.String(m.TextTag())
	w.Data = append(w.Data, ',')
	w.Int64(m.TimedOutOrderRequestedStatusDateTime())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequestedStatusForTimedOutOrder())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityToIncreaseFromParentFill())
	w.Data = append(w.Data, ',')
	w.Float64(m.MoveToBreakevenStopReferencePrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityTriggeredStop_QuantityForTrigger())
	w.Data = append(w.Data, ',')
	w.Float64(m.AccumulatedTradeVolumeForTriggeredStop())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BidAskQuantityStopInitialTriggerMet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NeedToOverrideLock())
	w.Data = append(w.Data, ',')
	w.Float64(m.CurrentMarketPrice())
	w.Data = append(w.Data, ',')
	w.Int64(m.CurrentMarketDateTime())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportOrderFillBilling())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsBillable())
	w.Data = append(w.Data, ',')
	w.Int32(m.QuantityForBilling())
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumberOfFailedOrderModifications())
	w.Data = append(w.Data, ',')
	w.Int32(m.DTCServerIndex())
	w.Data = append(w.Data, ',')
	w.String(m.ClearingFirmID())
	w.Data = append(w.Data, ',')
	w.String(m.SenderSubID())
	w.Data = append(w.Data, ',')
	w.String(m.SenderLocationId())
	w.Data = append(w.Data, ',')
	w.String(m.SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.Int32(m.CTICode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ObtainOrderActionDateTimeFromLastTradeTimeInChart())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaximumShowQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OrderSubmitted())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Int64(m.ExternalLastActionDateTimeUTC())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeOrderPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10110,\"F\":["...)
	w.Bool(m.IsOrderDeleted())
	w.Data = append(w.Data, ',')
	w.Uint64(m.InternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OrderStatusCode())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OrderStatusBeforePendingModify())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OrderStatusBeforePendingCancel())
	w.Data = append(w.Data, ',')
	w.String(m.ServiceOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ActualSymbol())
	w.Data = append(w.Data, ',')
	w.Int32(m.OrderType())
	w.Data = append(w.Data, ',')
	w.Uint16(m.BuySell())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.OrderQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.FilledQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AverageFillPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.RealtimeFillStatus())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsRestingOrderDuringFill())
	w.Data = append(w.Data, ',')
	w.Int32(m.OrderRejectType())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Int32(m.InternalOrderIDModifierForService())
	w.Data = append(w.Data, ',')
	w.String(m.FIXClientOrderID())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SequenceNumberBasedClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderIDForDTCServer())
	w.Data = append(w.Data, ',')
	w.String(m.PreviousClientOrderIDForDTCServer())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.OriginatingClientUsername())
	w.Data = append(w.Data, ',')
	w.Int64(m.EntryDateTime())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastActionDateTime())
	w.Data = append(w.Data, ',')
	w.Int64(m.ServiceUpdateDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OrderEntryTimeForService())
	w.Data = append(w.Data, ',')
	w.Uint32(m.LastModifyTimeForService())
	w.Data = append(w.Data, ',')
	w.Int64(m.GoodTillDateTime())
	w.Data = append(w.Data, ',')
	w.Int32(m.TimeInForce())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OpenClose())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailStopOffset1())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailStopStep())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailTriggerPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingStopTriggerOffset())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TrailTriggerHit())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailToBreakEvenStopOffset())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaximumChaseAmountAsPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.InitialChaseOrderPrice1())
	w.Data = append(w.Data, ',')
	w.Float64(m.InitialLastTradePriceForChaseOrders())
	w.Data = append(w.Data, ',')
	w.Int32(m.TrailingStopTriggerOCOGroupNumber())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastModifyPrice1())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastModifyQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.CumulativeOrderQuantityFromParentFills())
	w.Data = append(w.Data, ',')
	w.Float64(m.PriorFilledQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.TickSize())
	w.Data = append(w.Data, ',')
	w.Int32(m.ValueFormat())
	w.Data = append(w.Data, ',')
	w.Float32(m.PriceMultiplier())
	w.Data = append(w.Data, ',')
	w.Uint64(m.ParentInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TargetChildInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.StopChildInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.AttachedOrderPriceOffset1())
	w.Data = append(w.Data, ',')
	w.Uint64(m.LinkInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OCOGroupInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OCOSiblingInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DisableChildAndSiblingRelatedActions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OCOManagedByService())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BracketOrderServerManaged())
	w.Data = append(w.Data, ',')
	w.String(m.LastOrderActionSource())
	w.Data = append(w.Data, ',')
	w.Uint8(m.StopLimitOrderStopPriceTriggered())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OCOFullSiblingCancelOnPartialFill())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReverseOnCompleteFill())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportScaleIn())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportScaleOut())
	w.Data = append(w.Data, ',')
	w.Int32(m.SourceChartNumber())
	w.Data = append(w.Data, ',')
	w.String(m.SourceChartbookFileName())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SimulatedOrder())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsChartReplaying())
	w.Data = append(w.Data, ',')
	w.Int32(m.AttachedOrderOCOGroupNumber())
	w.Data = append(w.Data, ',')
	w.String(m.LastFillExecutionServiceID())
	w.Data = append(w.Data, ',')
	w.Int32(m.FillCount())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastFillQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastFillPrice())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastFillDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint64(m.RejectedStopOCOSiblingInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.RejectedStopReplacementMarketOrderQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.EvaluatingForFill())
	w.Data = append(w.Data, ',')
	w.Uint32(m.LastProcessedTimeSalesRecordSequenceForPrices())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsMarketDataManagementOfOrderEnabled())
	w.Data = append(w.Data, ',')
	w.String(m.TextTag())
	w.Data = append(w.Data, ',')
	w.Int64(m.TimedOutOrderRequestedStatusDateTime())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequestedStatusForTimedOutOrder())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityToIncreaseFromParentFill())
	w.Data = append(w.Data, ',')
	w.Float64(m.MoveToBreakevenStopReferencePrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityTriggeredStop_QuantityForTrigger())
	w.Data = append(w.Data, ',')
	w.Float64(m.AccumulatedTradeVolumeForTriggeredStop())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BidAskQuantityStopInitialTriggerMet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NeedToOverrideLock())
	w.Data = append(w.Data, ',')
	w.Float64(m.CurrentMarketPrice())
	w.Data = append(w.Data, ',')
	w.Int64(m.CurrentMarketDateTime())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportOrderFillBilling())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsBillable())
	w.Data = append(w.Data, ',')
	w.Int32(m.QuantityForBilling())
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumberOfFailedOrderModifications())
	w.Data = append(w.Data, ',')
	w.Int32(m.DTCServerIndex())
	w.Data = append(w.Data, ',')
	w.String(m.ClearingFirmID())
	w.Data = append(w.Data, ',')
	w.String(m.SenderSubID())
	w.Data = append(w.Data, ',')
	w.String(m.SenderLocationId())
	w.Data = append(w.Data, ',')
	w.String(m.SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.Int32(m.CTICode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ObtainOrderActionDateTimeFromLastTradeTimeInChart())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaximumShowQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OrderSubmitted())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Int64(m.ExternalLastActionDateTimeUTC())
	return w.Finish(), nil
}

func (m TradeOrderPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10110,\"F\":["...)
	w.Bool(m.IsOrderDeleted())
	w.Data = append(w.Data, ',')
	w.Uint64(m.InternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OrderStatusCode())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OrderStatusBeforePendingModify())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OrderStatusBeforePendingCancel())
	w.Data = append(w.Data, ',')
	w.String(m.ServiceOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ActualSymbol())
	w.Data = append(w.Data, ',')
	w.Int32(m.OrderType())
	w.Data = append(w.Data, ',')
	w.Uint16(m.BuySell())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.OrderQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.FilledQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AverageFillPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.RealtimeFillStatus())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsRestingOrderDuringFill())
	w.Data = append(w.Data, ',')
	w.Int32(m.OrderRejectType())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Int32(m.InternalOrderIDModifierForService())
	w.Data = append(w.Data, ',')
	w.String(m.FIXClientOrderID())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SequenceNumberBasedClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderIDForDTCServer())
	w.Data = append(w.Data, ',')
	w.String(m.PreviousClientOrderIDForDTCServer())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.OriginatingClientUsername())
	w.Data = append(w.Data, ',')
	w.Int64(m.EntryDateTime())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastActionDateTime())
	w.Data = append(w.Data, ',')
	w.Int64(m.ServiceUpdateDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OrderEntryTimeForService())
	w.Data = append(w.Data, ',')
	w.Uint32(m.LastModifyTimeForService())
	w.Data = append(w.Data, ',')
	w.Int64(m.GoodTillDateTime())
	w.Data = append(w.Data, ',')
	w.Int32(m.TimeInForce())
	w.Data = append(w.Data, ',')
	w.Uint16(m.OpenClose())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailStopOffset1())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailStopStep())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailTriggerPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailingStopTriggerOffset())
	w.Data = append(w.Data, ',')
	w.Uint8(m.TrailTriggerHit())
	w.Data = append(w.Data, ',')
	w.Float64(m.TrailToBreakEvenStopOffset())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaximumChaseAmountAsPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.InitialChaseOrderPrice1())
	w.Data = append(w.Data, ',')
	w.Float64(m.InitialLastTradePriceForChaseOrders())
	w.Data = append(w.Data, ',')
	w.Int32(m.TrailingStopTriggerOCOGroupNumber())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastModifyPrice1())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastModifyQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.CumulativeOrderQuantityFromParentFills())
	w.Data = append(w.Data, ',')
	w.Float64(m.PriorFilledQuantity())
	w.Data = append(w.Data, ',')
	w.Float32(m.TickSize())
	w.Data = append(w.Data, ',')
	w.Int32(m.ValueFormat())
	w.Data = append(w.Data, ',')
	w.Float32(m.PriceMultiplier())
	w.Data = append(w.Data, ',')
	w.Uint64(m.ParentInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TargetChildInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.StopChildInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.AttachedOrderPriceOffset1())
	w.Data = append(w.Data, ',')
	w.Uint64(m.LinkInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OCOGroupInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint64(m.OCOSiblingInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DisableChildAndSiblingRelatedActions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OCOManagedByService())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BracketOrderServerManaged())
	w.Data = append(w.Data, ',')
	w.String(m.LastOrderActionSource())
	w.Data = append(w.Data, ',')
	w.Uint8(m.StopLimitOrderStopPriceTriggered())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OCOFullSiblingCancelOnPartialFill())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ReverseOnCompleteFill())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportScaleIn())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportScaleOut())
	w.Data = append(w.Data, ',')
	w.Int32(m.SourceChartNumber())
	w.Data = append(w.Data, ',')
	w.String(m.SourceChartbookFileName())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsAutomatedOrder())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SimulatedOrder())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsChartReplaying())
	w.Data = append(w.Data, ',')
	w.Int32(m.AttachedOrderOCOGroupNumber())
	w.Data = append(w.Data, ',')
	w.String(m.LastFillExecutionServiceID())
	w.Data = append(w.Data, ',')
	w.Int32(m.FillCount())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastFillQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastFillPrice())
	w.Data = append(w.Data, ',')
	w.Int64(m.LastFillDateTimeUTC())
	w.Data = append(w.Data, ',')
	w.Uint64(m.RejectedStopOCOSiblingInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.RejectedStopReplacementMarketOrderQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.EvaluatingForFill())
	w.Data = append(w.Data, ',')
	w.Uint32(m.LastProcessedTimeSalesRecordSequenceForPrices())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsMarketDataManagementOfOrderEnabled())
	w.Data = append(w.Data, ',')
	w.String(m.TextTag())
	w.Data = append(w.Data, ',')
	w.Int64(m.TimedOutOrderRequestedStatusDateTime())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequestedStatusForTimedOutOrder())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityToIncreaseFromParentFill())
	w.Data = append(w.Data, ',')
	w.Float64(m.MoveToBreakevenStopReferencePrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityTriggeredStop_QuantityForTrigger())
	w.Data = append(w.Data, ',')
	w.Float64(m.AccumulatedTradeVolumeForTriggeredStop())
	w.Data = append(w.Data, ',')
	w.Uint8(m.BidAskQuantityStopInitialTriggerMet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NeedToOverrideLock())
	w.Data = append(w.Data, ',')
	w.Float64(m.CurrentMarketPrice())
	w.Data = append(w.Data, ',')
	w.Int64(m.CurrentMarketDateTime())
	w.Data = append(w.Data, ',')
	w.Uint8(m.SupportOrderFillBilling())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsBillable())
	w.Data = append(w.Data, ',')
	w.Int32(m.QuantityForBilling())
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumberOfFailedOrderModifications())
	w.Data = append(w.Data, ',')
	w.Int32(m.DTCServerIndex())
	w.Data = append(w.Data, ',')
	w.String(m.ClearingFirmID())
	w.Data = append(w.Data, ',')
	w.String(m.SenderSubID())
	w.Data = append(w.Data, ',')
	w.String(m.SenderLocationId())
	w.Data = append(w.Data, ',')
	w.String(m.SelfMatchPreventionID())
	w.Data = append(w.Data, ',')
	w.Int32(m.CTICode())
	w.Data = append(w.Data, ',')
	w.Uint8(m.ObtainOrderActionDateTimeFromLastTradeTimeInChart())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaximumShowQuantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.OrderSubmitted())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Int64(m.ExternalLastActionDateTimeUTC())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeOrder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10110)
	w.BoolField("m_IsOrderDeleted", m.IsOrderDeleted())
	w.Uint64Field("m_InternalOrderID", m.InternalOrderID())
	w.Uint16Field("m_OrderStatusCode", m.OrderStatusCode())
	w.Uint16Field("m_OrderStatusBeforePendingModify", m.OrderStatusBeforePendingModify())
	w.Uint16Field("m_OrderStatusBeforePendingCancel", m.OrderStatusBeforePendingCancel())
	w.StringField("m_ServiceOrderID", m.ServiceOrderID())
	w.StringField("m_ActualSymbol", m.ActualSymbol())
	w.Int32Field("m_OrderType", m.OrderType())
	w.Uint16Field("m_BuySell", m.BuySell())
	w.Float64Field("m_Price1", m.Price1())
	w.Float64Field("m_Price2", m.Price2())
	w.Float64Field("m_OrderQuantity", m.OrderQuantity())
	w.Float64Field("m_FilledQuantity", m.FilledQuantity())
	w.Float64Field("m_AverageFillPrice", m.AverageFillPrice())
	w.Int32Field("m_RealtimeFillStatus", m.RealtimeFillStatus())
	w.BoolField("m_IsRestingOrderDuringFill", m.IsRestingOrderDuringFill())
	w.Int32Field("m_OrderRejectType", m.OrderRejectType())
	w.StringField("m_TradeAccount", m.TradeAccount())
	w.Uint32Field("m_SubAccountIdentifier", m.SubAccountIdentifier())
	w.Int32Field("m_InternalOrderIDModifierForService", m.InternalOrderIDModifierForService())
	w.StringField("m_FIXClientOrderID", m.FIXClientOrderID())
	w.Uint32Field("m_SequenceNumberBasedClientOrderID", m.SequenceNumberBasedClientOrderID())
	w.StringField("m_ClientOrderIDForDTCServer", m.ClientOrderIDForDTCServer())
	w.StringField("m_PreviousClientOrderIDForDTCServer", m.PreviousClientOrderIDForDTCServer())
	w.StringField("m_ExchangeOrderID", m.ExchangeOrderID())
	w.StringField("m_OriginatingClientUsername", m.OriginatingClientUsername())
	w.Int64Field("m_EntryDateTime", m.EntryDateTime())
	w.Int64Field("m_LastActionDateTime", m.LastActionDateTime())
	w.Int64Field("m_ServiceUpdateDateTimeUTC", m.ServiceUpdateDateTimeUTC())
	w.Uint32Field("m_OrderEntryTimeForService", m.OrderEntryTimeForService())
	w.Uint32Field("m_LastModifyTimeForService", m.LastModifyTimeForService())
	w.Int64Field("m_GoodTillDateTime", m.GoodTillDateTime())
	w.Int32Field("m_TimeInForce", m.TimeInForce())
	w.Uint16Field("m_OpenClose", m.OpenClose())
	w.Float64Field("m_TrailStopOffset1", m.TrailStopOffset1())
	w.Float64Field("m_TrailStopStep", m.TrailStopStep())
	w.Float64Field("m_TrailTriggerPrice", m.TrailTriggerPrice())
	w.Float64Field("m_TrailingStopTriggerOffset", m.TrailingStopTriggerOffset())
	w.Uint8Field("m_TrailTriggerHit", m.TrailTriggerHit())
	w.Float64Field("m_TrailToBreakEvenStopOffset", m.TrailToBreakEvenStopOffset())
	w.Float64Field("m_MaximumChaseAmountAsPrice", m.MaximumChaseAmountAsPrice())
	w.Float64Field("m_InitialChaseOrderPrice1", m.InitialChaseOrderPrice1())
	w.Float64Field("m_InitialLastTradePriceForChaseOrders", m.InitialLastTradePriceForChaseOrders())
	w.Int32Field("m_TrailingStopTriggerOCOGroupNumber", m.TrailingStopTriggerOCOGroupNumber())
	w.Float64Field("m_LastModifyPrice1", m.LastModifyPrice1())
	w.Float64Field("m_LastModifyQuantity", m.LastModifyQuantity())
	w.Float64Field("m_CumulativeOrderQuantityFromParentFills", m.CumulativeOrderQuantityFromParentFills())
	w.Float64Field("m_PriorFilledQuantity", m.PriorFilledQuantity())
	w.Float32Field("m_TickSize", m.TickSize())
	w.Int32Field("m_ValueFormat", m.ValueFormat())
	w.Float32Field("m_PriceMultiplier", m.PriceMultiplier())
	w.Uint64Field("m_ParentInternalOrderID", m.ParentInternalOrderID())
	w.Uint64Field("m_TargetChildInternalOrderID", m.TargetChildInternalOrderID())
	w.Uint64Field("m_StopChildInternalOrderID", m.StopChildInternalOrderID())
	w.Float64Field("m_AttachedOrderPriceOffset1", m.AttachedOrderPriceOffset1())
	w.Uint64Field("m_LinkInternalOrderID", m.LinkInternalOrderID())
	w.Uint64Field("m_OCOGroupInternalOrderID", m.OCOGroupInternalOrderID())
	w.Uint64Field("m_OCOSiblingInternalOrderID", m.OCOSiblingInternalOrderID())
	w.Uint8Field("m_DisableChildAndSiblingRelatedActions", m.DisableChildAndSiblingRelatedActions())
	w.Uint8Field("m_OCOManagedByService", m.OCOManagedByService())
	w.Uint8Field("m_BracketOrderServerManaged", m.BracketOrderServerManaged())
	w.StringField("m_LastOrderActionSource", m.LastOrderActionSource())
	w.Uint8Field("m_StopLimitOrderStopPriceTriggered", m.StopLimitOrderStopPriceTriggered())
	w.Uint8Field("m_OCOFullSiblingCancelOnPartialFill", m.OCOFullSiblingCancelOnPartialFill())
	w.Uint8Field("m_ReverseOnCompleteFill", m.ReverseOnCompleteFill())
	w.Uint8Field("m_SupportScaleIn", m.SupportScaleIn())
	w.Uint8Field("m_SupportScaleOut", m.SupportScaleOut())
	w.Int32Field("m_SourceChartNumber", m.SourceChartNumber())
	w.StringField("m_SourceChartbookFileName", m.SourceChartbookFileName())
	w.BoolField("m_IsAutomatedOrder", m.IsAutomatedOrder())
	w.Uint8Field("m_SimulatedOrder", m.SimulatedOrder())
	w.BoolField("m_IsChartReplaying", m.IsChartReplaying())
	w.Int32Field("m_AttachedOrderOCOGroupNumber", m.AttachedOrderOCOGroupNumber())
	w.StringField("m_LastFillExecutionServiceID", m.LastFillExecutionServiceID())
	w.Int32Field("m_FillCount", m.FillCount())
	w.Float64Field("m_LastFillQuantity", m.LastFillQuantity())
	w.Float64Field("m_LastFillPrice", m.LastFillPrice())
	w.Int64Field("m_LastFillDateTimeUTC", m.LastFillDateTimeUTC())
	w.Uint64Field("m_RejectedStopOCOSiblingInternalOrderID", m.RejectedStopOCOSiblingInternalOrderID())
	w.Float64Field("m_RejectedStopReplacementMarketOrderQuantity", m.RejectedStopReplacementMarketOrderQuantity())
	w.Uint8Field("m_EvaluatingForFill", m.EvaluatingForFill())
	w.Uint32Field("m_LastProcessedTimeSalesRecordSequenceForPrices", m.LastProcessedTimeSalesRecordSequenceForPrices())
	w.BoolField("m_IsMarketDataManagementOfOrderEnabled", m.IsMarketDataManagementOfOrderEnabled())
	w.StringField("m_TextTag", m.TextTag())
	w.Int64Field("m_TimedOutOrderRequestedStatusDateTime", m.TimedOutOrderRequestedStatusDateTime())
	w.Uint8Field("m_RequestedStatusForTimedOutOrder", m.RequestedStatusForTimedOutOrder())
	w.Uint8Field("m_SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled", m.SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled())
	w.Float64Field("m_QuantityToIncreaseFromParentFill", m.QuantityToIncreaseFromParentFill())
	w.Float64Field("m_MoveToBreakevenStopReferencePrice", m.MoveToBreakevenStopReferencePrice())
	w.Float64Field("m_QuantityTriggeredStop_QuantityForTrigger", m.QuantityTriggeredStop_QuantityForTrigger())
	w.Float64Field("m_AccumulatedTradeVolumeForTriggeredStop", m.AccumulatedTradeVolumeForTriggeredStop())
	w.Uint8Field("m_BidAskQuantityStopInitialTriggerMet", m.BidAskQuantityStopInitialTriggerMet())
	w.Uint8Field("m_NeedToOverrideLock", m.NeedToOverrideLock())
	w.Float64Field("m_CurrentMarketPrice", m.CurrentMarketPrice())
	w.Int64Field("m_CurrentMarketDateTime", m.CurrentMarketDateTime())
	w.Uint8Field("m_SupportOrderFillBilling", m.SupportOrderFillBilling())
	w.BoolField("m_IsBillable", m.IsBillable())
	w.Int32Field("m_QuantityForBilling", m.QuantityForBilling())
	w.Uint32Field("m_NumberOfFailedOrderModifications", m.NumberOfFailedOrderModifications())
	w.Int32Field("m_DTCServerIndex", m.DTCServerIndex())
	w.StringField("m_ClearingFirmID", m.ClearingFirmID())
	w.StringField("m_SenderSubID", m.SenderSubID())
	w.StringField("m_SenderLocationId", m.SenderLocationId())
	w.StringField("m_SelfMatchPreventionID", m.SelfMatchPreventionID())
	w.Int32Field("m_CTICode", m.CTICode())
	w.Uint8Field("m_ObtainOrderActionDateTimeFromLastTradeTimeInChart", m.ObtainOrderActionDateTimeFromLastTradeTimeInChart())
	w.Float64Field("m_MaximumShowQuantity", m.MaximumShowQuantity())
	w.Uint8Field("m_OrderSubmitted", m.OrderSubmitted())
	w.BoolField("m_IsSnapshot", m.IsSnapshot())
	w.BoolField("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Int64Field("m_ExternalLastActionDateTimeUTC", m.ExternalLastActionDateTimeUTC())
	return w.Finish(), nil
}

func (m TradeOrderBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10110)
	w.BoolField("m_IsOrderDeleted", m.IsOrderDeleted())
	w.Uint64Field("m_InternalOrderID", m.InternalOrderID())
	w.Uint16Field("m_OrderStatusCode", m.OrderStatusCode())
	w.Uint16Field("m_OrderStatusBeforePendingModify", m.OrderStatusBeforePendingModify())
	w.Uint16Field("m_OrderStatusBeforePendingCancel", m.OrderStatusBeforePendingCancel())
	w.StringField("m_ServiceOrderID", m.ServiceOrderID())
	w.StringField("m_ActualSymbol", m.ActualSymbol())
	w.Int32Field("m_OrderType", m.OrderType())
	w.Uint16Field("m_BuySell", m.BuySell())
	w.Float64Field("m_Price1", m.Price1())
	w.Float64Field("m_Price2", m.Price2())
	w.Float64Field("m_OrderQuantity", m.OrderQuantity())
	w.Float64Field("m_FilledQuantity", m.FilledQuantity())
	w.Float64Field("m_AverageFillPrice", m.AverageFillPrice())
	w.Int32Field("m_RealtimeFillStatus", m.RealtimeFillStatus())
	w.BoolField("m_IsRestingOrderDuringFill", m.IsRestingOrderDuringFill())
	w.Int32Field("m_OrderRejectType", m.OrderRejectType())
	w.StringField("m_TradeAccount", m.TradeAccount())
	w.Uint32Field("m_SubAccountIdentifier", m.SubAccountIdentifier())
	w.Int32Field("m_InternalOrderIDModifierForService", m.InternalOrderIDModifierForService())
	w.StringField("m_FIXClientOrderID", m.FIXClientOrderID())
	w.Uint32Field("m_SequenceNumberBasedClientOrderID", m.SequenceNumberBasedClientOrderID())
	w.StringField("m_ClientOrderIDForDTCServer", m.ClientOrderIDForDTCServer())
	w.StringField("m_PreviousClientOrderIDForDTCServer", m.PreviousClientOrderIDForDTCServer())
	w.StringField("m_ExchangeOrderID", m.ExchangeOrderID())
	w.StringField("m_OriginatingClientUsername", m.OriginatingClientUsername())
	w.Int64Field("m_EntryDateTime", m.EntryDateTime())
	w.Int64Field("m_LastActionDateTime", m.LastActionDateTime())
	w.Int64Field("m_ServiceUpdateDateTimeUTC", m.ServiceUpdateDateTimeUTC())
	w.Uint32Field("m_OrderEntryTimeForService", m.OrderEntryTimeForService())
	w.Uint32Field("m_LastModifyTimeForService", m.LastModifyTimeForService())
	w.Int64Field("m_GoodTillDateTime", m.GoodTillDateTime())
	w.Int32Field("m_TimeInForce", m.TimeInForce())
	w.Uint16Field("m_OpenClose", m.OpenClose())
	w.Float64Field("m_TrailStopOffset1", m.TrailStopOffset1())
	w.Float64Field("m_TrailStopStep", m.TrailStopStep())
	w.Float64Field("m_TrailTriggerPrice", m.TrailTriggerPrice())
	w.Float64Field("m_TrailingStopTriggerOffset", m.TrailingStopTriggerOffset())
	w.Uint8Field("m_TrailTriggerHit", m.TrailTriggerHit())
	w.Float64Field("m_TrailToBreakEvenStopOffset", m.TrailToBreakEvenStopOffset())
	w.Float64Field("m_MaximumChaseAmountAsPrice", m.MaximumChaseAmountAsPrice())
	w.Float64Field("m_InitialChaseOrderPrice1", m.InitialChaseOrderPrice1())
	w.Float64Field("m_InitialLastTradePriceForChaseOrders", m.InitialLastTradePriceForChaseOrders())
	w.Int32Field("m_TrailingStopTriggerOCOGroupNumber", m.TrailingStopTriggerOCOGroupNumber())
	w.Float64Field("m_LastModifyPrice1", m.LastModifyPrice1())
	w.Float64Field("m_LastModifyQuantity", m.LastModifyQuantity())
	w.Float64Field("m_CumulativeOrderQuantityFromParentFills", m.CumulativeOrderQuantityFromParentFills())
	w.Float64Field("m_PriorFilledQuantity", m.PriorFilledQuantity())
	w.Float32Field("m_TickSize", m.TickSize())
	w.Int32Field("m_ValueFormat", m.ValueFormat())
	w.Float32Field("m_PriceMultiplier", m.PriceMultiplier())
	w.Uint64Field("m_ParentInternalOrderID", m.ParentInternalOrderID())
	w.Uint64Field("m_TargetChildInternalOrderID", m.TargetChildInternalOrderID())
	w.Uint64Field("m_StopChildInternalOrderID", m.StopChildInternalOrderID())
	w.Float64Field("m_AttachedOrderPriceOffset1", m.AttachedOrderPriceOffset1())
	w.Uint64Field("m_LinkInternalOrderID", m.LinkInternalOrderID())
	w.Uint64Field("m_OCOGroupInternalOrderID", m.OCOGroupInternalOrderID())
	w.Uint64Field("m_OCOSiblingInternalOrderID", m.OCOSiblingInternalOrderID())
	w.Uint8Field("m_DisableChildAndSiblingRelatedActions", m.DisableChildAndSiblingRelatedActions())
	w.Uint8Field("m_OCOManagedByService", m.OCOManagedByService())
	w.Uint8Field("m_BracketOrderServerManaged", m.BracketOrderServerManaged())
	w.StringField("m_LastOrderActionSource", m.LastOrderActionSource())
	w.Uint8Field("m_StopLimitOrderStopPriceTriggered", m.StopLimitOrderStopPriceTriggered())
	w.Uint8Field("m_OCOFullSiblingCancelOnPartialFill", m.OCOFullSiblingCancelOnPartialFill())
	w.Uint8Field("m_ReverseOnCompleteFill", m.ReverseOnCompleteFill())
	w.Uint8Field("m_SupportScaleIn", m.SupportScaleIn())
	w.Uint8Field("m_SupportScaleOut", m.SupportScaleOut())
	w.Int32Field("m_SourceChartNumber", m.SourceChartNumber())
	w.StringField("m_SourceChartbookFileName", m.SourceChartbookFileName())
	w.BoolField("m_IsAutomatedOrder", m.IsAutomatedOrder())
	w.Uint8Field("m_SimulatedOrder", m.SimulatedOrder())
	w.BoolField("m_IsChartReplaying", m.IsChartReplaying())
	w.Int32Field("m_AttachedOrderOCOGroupNumber", m.AttachedOrderOCOGroupNumber())
	w.StringField("m_LastFillExecutionServiceID", m.LastFillExecutionServiceID())
	w.Int32Field("m_FillCount", m.FillCount())
	w.Float64Field("m_LastFillQuantity", m.LastFillQuantity())
	w.Float64Field("m_LastFillPrice", m.LastFillPrice())
	w.Int64Field("m_LastFillDateTimeUTC", m.LastFillDateTimeUTC())
	w.Uint64Field("m_RejectedStopOCOSiblingInternalOrderID", m.RejectedStopOCOSiblingInternalOrderID())
	w.Float64Field("m_RejectedStopReplacementMarketOrderQuantity", m.RejectedStopReplacementMarketOrderQuantity())
	w.Uint8Field("m_EvaluatingForFill", m.EvaluatingForFill())
	w.Uint32Field("m_LastProcessedTimeSalesRecordSequenceForPrices", m.LastProcessedTimeSalesRecordSequenceForPrices())
	w.BoolField("m_IsMarketDataManagementOfOrderEnabled", m.IsMarketDataManagementOfOrderEnabled())
	w.StringField("m_TextTag", m.TextTag())
	w.Int64Field("m_TimedOutOrderRequestedStatusDateTime", m.TimedOutOrderRequestedStatusDateTime())
	w.Uint8Field("m_RequestedStatusForTimedOutOrder", m.RequestedStatusForTimedOutOrder())
	w.Uint8Field("m_SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled", m.SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled())
	w.Float64Field("m_QuantityToIncreaseFromParentFill", m.QuantityToIncreaseFromParentFill())
	w.Float64Field("m_MoveToBreakevenStopReferencePrice", m.MoveToBreakevenStopReferencePrice())
	w.Float64Field("m_QuantityTriggeredStop_QuantityForTrigger", m.QuantityTriggeredStop_QuantityForTrigger())
	w.Float64Field("m_AccumulatedTradeVolumeForTriggeredStop", m.AccumulatedTradeVolumeForTriggeredStop())
	w.Uint8Field("m_BidAskQuantityStopInitialTriggerMet", m.BidAskQuantityStopInitialTriggerMet())
	w.Uint8Field("m_NeedToOverrideLock", m.NeedToOverrideLock())
	w.Float64Field("m_CurrentMarketPrice", m.CurrentMarketPrice())
	w.Int64Field("m_CurrentMarketDateTime", m.CurrentMarketDateTime())
	w.Uint8Field("m_SupportOrderFillBilling", m.SupportOrderFillBilling())
	w.BoolField("m_IsBillable", m.IsBillable())
	w.Int32Field("m_QuantityForBilling", m.QuantityForBilling())
	w.Uint32Field("m_NumberOfFailedOrderModifications", m.NumberOfFailedOrderModifications())
	w.Int32Field("m_DTCServerIndex", m.DTCServerIndex())
	w.StringField("m_ClearingFirmID", m.ClearingFirmID())
	w.StringField("m_SenderSubID", m.SenderSubID())
	w.StringField("m_SenderLocationId", m.SenderLocationId())
	w.StringField("m_SelfMatchPreventionID", m.SelfMatchPreventionID())
	w.Int32Field("m_CTICode", m.CTICode())
	w.Uint8Field("m_ObtainOrderActionDateTimeFromLastTradeTimeInChart", m.ObtainOrderActionDateTimeFromLastTradeTimeInChart())
	w.Float64Field("m_MaximumShowQuantity", m.MaximumShowQuantity())
	w.Uint8Field("m_OrderSubmitted", m.OrderSubmitted())
	w.BoolField("m_IsSnapshot", m.IsSnapshot())
	w.BoolField("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Int64Field("m_ExternalLastActionDateTimeUTC", m.ExternalLastActionDateTimeUTC())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeOrderPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10110)
	w.BoolField("m_IsOrderDeleted", m.IsOrderDeleted())
	w.Uint64Field("m_InternalOrderID", m.InternalOrderID())
	w.Uint16Field("m_OrderStatusCode", m.OrderStatusCode())
	w.Uint16Field("m_OrderStatusBeforePendingModify", m.OrderStatusBeforePendingModify())
	w.Uint16Field("m_OrderStatusBeforePendingCancel", m.OrderStatusBeforePendingCancel())
	w.StringField("m_ServiceOrderID", m.ServiceOrderID())
	w.StringField("m_ActualSymbol", m.ActualSymbol())
	w.Int32Field("m_OrderType", m.OrderType())
	w.Uint16Field("m_BuySell", m.BuySell())
	w.Float64Field("m_Price1", m.Price1())
	w.Float64Field("m_Price2", m.Price2())
	w.Float64Field("m_OrderQuantity", m.OrderQuantity())
	w.Float64Field("m_FilledQuantity", m.FilledQuantity())
	w.Float64Field("m_AverageFillPrice", m.AverageFillPrice())
	w.Int32Field("m_RealtimeFillStatus", m.RealtimeFillStatus())
	w.BoolField("m_IsRestingOrderDuringFill", m.IsRestingOrderDuringFill())
	w.Int32Field("m_OrderRejectType", m.OrderRejectType())
	w.StringField("m_TradeAccount", m.TradeAccount())
	w.Uint32Field("m_SubAccountIdentifier", m.SubAccountIdentifier())
	w.Int32Field("m_InternalOrderIDModifierForService", m.InternalOrderIDModifierForService())
	w.StringField("m_FIXClientOrderID", m.FIXClientOrderID())
	w.Uint32Field("m_SequenceNumberBasedClientOrderID", m.SequenceNumberBasedClientOrderID())
	w.StringField("m_ClientOrderIDForDTCServer", m.ClientOrderIDForDTCServer())
	w.StringField("m_PreviousClientOrderIDForDTCServer", m.PreviousClientOrderIDForDTCServer())
	w.StringField("m_ExchangeOrderID", m.ExchangeOrderID())
	w.StringField("m_OriginatingClientUsername", m.OriginatingClientUsername())
	w.Int64Field("m_EntryDateTime", m.EntryDateTime())
	w.Int64Field("m_LastActionDateTime", m.LastActionDateTime())
	w.Int64Field("m_ServiceUpdateDateTimeUTC", m.ServiceUpdateDateTimeUTC())
	w.Uint32Field("m_OrderEntryTimeForService", m.OrderEntryTimeForService())
	w.Uint32Field("m_LastModifyTimeForService", m.LastModifyTimeForService())
	w.Int64Field("m_GoodTillDateTime", m.GoodTillDateTime())
	w.Int32Field("m_TimeInForce", m.TimeInForce())
	w.Uint16Field("m_OpenClose", m.OpenClose())
	w.Float64Field("m_TrailStopOffset1", m.TrailStopOffset1())
	w.Float64Field("m_TrailStopStep", m.TrailStopStep())
	w.Float64Field("m_TrailTriggerPrice", m.TrailTriggerPrice())
	w.Float64Field("m_TrailingStopTriggerOffset", m.TrailingStopTriggerOffset())
	w.Uint8Field("m_TrailTriggerHit", m.TrailTriggerHit())
	w.Float64Field("m_TrailToBreakEvenStopOffset", m.TrailToBreakEvenStopOffset())
	w.Float64Field("m_MaximumChaseAmountAsPrice", m.MaximumChaseAmountAsPrice())
	w.Float64Field("m_InitialChaseOrderPrice1", m.InitialChaseOrderPrice1())
	w.Float64Field("m_InitialLastTradePriceForChaseOrders", m.InitialLastTradePriceForChaseOrders())
	w.Int32Field("m_TrailingStopTriggerOCOGroupNumber", m.TrailingStopTriggerOCOGroupNumber())
	w.Float64Field("m_LastModifyPrice1", m.LastModifyPrice1())
	w.Float64Field("m_LastModifyQuantity", m.LastModifyQuantity())
	w.Float64Field("m_CumulativeOrderQuantityFromParentFills", m.CumulativeOrderQuantityFromParentFills())
	w.Float64Field("m_PriorFilledQuantity", m.PriorFilledQuantity())
	w.Float32Field("m_TickSize", m.TickSize())
	w.Int32Field("m_ValueFormat", m.ValueFormat())
	w.Float32Field("m_PriceMultiplier", m.PriceMultiplier())
	w.Uint64Field("m_ParentInternalOrderID", m.ParentInternalOrderID())
	w.Uint64Field("m_TargetChildInternalOrderID", m.TargetChildInternalOrderID())
	w.Uint64Field("m_StopChildInternalOrderID", m.StopChildInternalOrderID())
	w.Float64Field("m_AttachedOrderPriceOffset1", m.AttachedOrderPriceOffset1())
	w.Uint64Field("m_LinkInternalOrderID", m.LinkInternalOrderID())
	w.Uint64Field("m_OCOGroupInternalOrderID", m.OCOGroupInternalOrderID())
	w.Uint64Field("m_OCOSiblingInternalOrderID", m.OCOSiblingInternalOrderID())
	w.Uint8Field("m_DisableChildAndSiblingRelatedActions", m.DisableChildAndSiblingRelatedActions())
	w.Uint8Field("m_OCOManagedByService", m.OCOManagedByService())
	w.Uint8Field("m_BracketOrderServerManaged", m.BracketOrderServerManaged())
	w.StringField("m_LastOrderActionSource", m.LastOrderActionSource())
	w.Uint8Field("m_StopLimitOrderStopPriceTriggered", m.StopLimitOrderStopPriceTriggered())
	w.Uint8Field("m_OCOFullSiblingCancelOnPartialFill", m.OCOFullSiblingCancelOnPartialFill())
	w.Uint8Field("m_ReverseOnCompleteFill", m.ReverseOnCompleteFill())
	w.Uint8Field("m_SupportScaleIn", m.SupportScaleIn())
	w.Uint8Field("m_SupportScaleOut", m.SupportScaleOut())
	w.Int32Field("m_SourceChartNumber", m.SourceChartNumber())
	w.StringField("m_SourceChartbookFileName", m.SourceChartbookFileName())
	w.BoolField("m_IsAutomatedOrder", m.IsAutomatedOrder())
	w.Uint8Field("m_SimulatedOrder", m.SimulatedOrder())
	w.BoolField("m_IsChartReplaying", m.IsChartReplaying())
	w.Int32Field("m_AttachedOrderOCOGroupNumber", m.AttachedOrderOCOGroupNumber())
	w.StringField("m_LastFillExecutionServiceID", m.LastFillExecutionServiceID())
	w.Int32Field("m_FillCount", m.FillCount())
	w.Float64Field("m_LastFillQuantity", m.LastFillQuantity())
	w.Float64Field("m_LastFillPrice", m.LastFillPrice())
	w.Int64Field("m_LastFillDateTimeUTC", m.LastFillDateTimeUTC())
	w.Uint64Field("m_RejectedStopOCOSiblingInternalOrderID", m.RejectedStopOCOSiblingInternalOrderID())
	w.Float64Field("m_RejectedStopReplacementMarketOrderQuantity", m.RejectedStopReplacementMarketOrderQuantity())
	w.Uint8Field("m_EvaluatingForFill", m.EvaluatingForFill())
	w.Uint32Field("m_LastProcessedTimeSalesRecordSequenceForPrices", m.LastProcessedTimeSalesRecordSequenceForPrices())
	w.BoolField("m_IsMarketDataManagementOfOrderEnabled", m.IsMarketDataManagementOfOrderEnabled())
	w.StringField("m_TextTag", m.TextTag())
	w.Int64Field("m_TimedOutOrderRequestedStatusDateTime", m.TimedOutOrderRequestedStatusDateTime())
	w.Uint8Field("m_RequestedStatusForTimedOutOrder", m.RequestedStatusForTimedOutOrder())
	w.Uint8Field("m_SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled", m.SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled())
	w.Float64Field("m_QuantityToIncreaseFromParentFill", m.QuantityToIncreaseFromParentFill())
	w.Float64Field("m_MoveToBreakevenStopReferencePrice", m.MoveToBreakevenStopReferencePrice())
	w.Float64Field("m_QuantityTriggeredStop_QuantityForTrigger", m.QuantityTriggeredStop_QuantityForTrigger())
	w.Float64Field("m_AccumulatedTradeVolumeForTriggeredStop", m.AccumulatedTradeVolumeForTriggeredStop())
	w.Uint8Field("m_BidAskQuantityStopInitialTriggerMet", m.BidAskQuantityStopInitialTriggerMet())
	w.Uint8Field("m_NeedToOverrideLock", m.NeedToOverrideLock())
	w.Float64Field("m_CurrentMarketPrice", m.CurrentMarketPrice())
	w.Int64Field("m_CurrentMarketDateTime", m.CurrentMarketDateTime())
	w.Uint8Field("m_SupportOrderFillBilling", m.SupportOrderFillBilling())
	w.BoolField("m_IsBillable", m.IsBillable())
	w.Int32Field("m_QuantityForBilling", m.QuantityForBilling())
	w.Uint32Field("m_NumberOfFailedOrderModifications", m.NumberOfFailedOrderModifications())
	w.Int32Field("m_DTCServerIndex", m.DTCServerIndex())
	w.StringField("m_ClearingFirmID", m.ClearingFirmID())
	w.StringField("m_SenderSubID", m.SenderSubID())
	w.StringField("m_SenderLocationId", m.SenderLocationId())
	w.StringField("m_SelfMatchPreventionID", m.SelfMatchPreventionID())
	w.Int32Field("m_CTICode", m.CTICode())
	w.Uint8Field("m_ObtainOrderActionDateTimeFromLastTradeTimeInChart", m.ObtainOrderActionDateTimeFromLastTradeTimeInChart())
	w.Float64Field("m_MaximumShowQuantity", m.MaximumShowQuantity())
	w.Uint8Field("m_OrderSubmitted", m.OrderSubmitted())
	w.BoolField("m_IsSnapshot", m.IsSnapshot())
	w.BoolField("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Int64Field("m_ExternalLastActionDateTimeUTC", m.ExternalLastActionDateTimeUTC())
	return w.Finish(), nil
}

func (m TradeOrderPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10110)
	w.BoolField("m_IsOrderDeleted", m.IsOrderDeleted())
	w.Uint64Field("m_InternalOrderID", m.InternalOrderID())
	w.Uint16Field("m_OrderStatusCode", m.OrderStatusCode())
	w.Uint16Field("m_OrderStatusBeforePendingModify", m.OrderStatusBeforePendingModify())
	w.Uint16Field("m_OrderStatusBeforePendingCancel", m.OrderStatusBeforePendingCancel())
	w.StringField("m_ServiceOrderID", m.ServiceOrderID())
	w.StringField("m_ActualSymbol", m.ActualSymbol())
	w.Int32Field("m_OrderType", m.OrderType())
	w.Uint16Field("m_BuySell", m.BuySell())
	w.Float64Field("m_Price1", m.Price1())
	w.Float64Field("m_Price2", m.Price2())
	w.Float64Field("m_OrderQuantity", m.OrderQuantity())
	w.Float64Field("m_FilledQuantity", m.FilledQuantity())
	w.Float64Field("m_AverageFillPrice", m.AverageFillPrice())
	w.Int32Field("m_RealtimeFillStatus", m.RealtimeFillStatus())
	w.BoolField("m_IsRestingOrderDuringFill", m.IsRestingOrderDuringFill())
	w.Int32Field("m_OrderRejectType", m.OrderRejectType())
	w.StringField("m_TradeAccount", m.TradeAccount())
	w.Uint32Field("m_SubAccountIdentifier", m.SubAccountIdentifier())
	w.Int32Field("m_InternalOrderIDModifierForService", m.InternalOrderIDModifierForService())
	w.StringField("m_FIXClientOrderID", m.FIXClientOrderID())
	w.Uint32Field("m_SequenceNumberBasedClientOrderID", m.SequenceNumberBasedClientOrderID())
	w.StringField("m_ClientOrderIDForDTCServer", m.ClientOrderIDForDTCServer())
	w.StringField("m_PreviousClientOrderIDForDTCServer", m.PreviousClientOrderIDForDTCServer())
	w.StringField("m_ExchangeOrderID", m.ExchangeOrderID())
	w.StringField("m_OriginatingClientUsername", m.OriginatingClientUsername())
	w.Int64Field("m_EntryDateTime", m.EntryDateTime())
	w.Int64Field("m_LastActionDateTime", m.LastActionDateTime())
	w.Int64Field("m_ServiceUpdateDateTimeUTC", m.ServiceUpdateDateTimeUTC())
	w.Uint32Field("m_OrderEntryTimeForService", m.OrderEntryTimeForService())
	w.Uint32Field("m_LastModifyTimeForService", m.LastModifyTimeForService())
	w.Int64Field("m_GoodTillDateTime", m.GoodTillDateTime())
	w.Int32Field("m_TimeInForce", m.TimeInForce())
	w.Uint16Field("m_OpenClose", m.OpenClose())
	w.Float64Field("m_TrailStopOffset1", m.TrailStopOffset1())
	w.Float64Field("m_TrailStopStep", m.TrailStopStep())
	w.Float64Field("m_TrailTriggerPrice", m.TrailTriggerPrice())
	w.Float64Field("m_TrailingStopTriggerOffset", m.TrailingStopTriggerOffset())
	w.Uint8Field("m_TrailTriggerHit", m.TrailTriggerHit())
	w.Float64Field("m_TrailToBreakEvenStopOffset", m.TrailToBreakEvenStopOffset())
	w.Float64Field("m_MaximumChaseAmountAsPrice", m.MaximumChaseAmountAsPrice())
	w.Float64Field("m_InitialChaseOrderPrice1", m.InitialChaseOrderPrice1())
	w.Float64Field("m_InitialLastTradePriceForChaseOrders", m.InitialLastTradePriceForChaseOrders())
	w.Int32Field("m_TrailingStopTriggerOCOGroupNumber", m.TrailingStopTriggerOCOGroupNumber())
	w.Float64Field("m_LastModifyPrice1", m.LastModifyPrice1())
	w.Float64Field("m_LastModifyQuantity", m.LastModifyQuantity())
	w.Float64Field("m_CumulativeOrderQuantityFromParentFills", m.CumulativeOrderQuantityFromParentFills())
	w.Float64Field("m_PriorFilledQuantity", m.PriorFilledQuantity())
	w.Float32Field("m_TickSize", m.TickSize())
	w.Int32Field("m_ValueFormat", m.ValueFormat())
	w.Float32Field("m_PriceMultiplier", m.PriceMultiplier())
	w.Uint64Field("m_ParentInternalOrderID", m.ParentInternalOrderID())
	w.Uint64Field("m_TargetChildInternalOrderID", m.TargetChildInternalOrderID())
	w.Uint64Field("m_StopChildInternalOrderID", m.StopChildInternalOrderID())
	w.Float64Field("m_AttachedOrderPriceOffset1", m.AttachedOrderPriceOffset1())
	w.Uint64Field("m_LinkInternalOrderID", m.LinkInternalOrderID())
	w.Uint64Field("m_OCOGroupInternalOrderID", m.OCOGroupInternalOrderID())
	w.Uint64Field("m_OCOSiblingInternalOrderID", m.OCOSiblingInternalOrderID())
	w.Uint8Field("m_DisableChildAndSiblingRelatedActions", m.DisableChildAndSiblingRelatedActions())
	w.Uint8Field("m_OCOManagedByService", m.OCOManagedByService())
	w.Uint8Field("m_BracketOrderServerManaged", m.BracketOrderServerManaged())
	w.StringField("m_LastOrderActionSource", m.LastOrderActionSource())
	w.Uint8Field("m_StopLimitOrderStopPriceTriggered", m.StopLimitOrderStopPriceTriggered())
	w.Uint8Field("m_OCOFullSiblingCancelOnPartialFill", m.OCOFullSiblingCancelOnPartialFill())
	w.Uint8Field("m_ReverseOnCompleteFill", m.ReverseOnCompleteFill())
	w.Uint8Field("m_SupportScaleIn", m.SupportScaleIn())
	w.Uint8Field("m_SupportScaleOut", m.SupportScaleOut())
	w.Int32Field("m_SourceChartNumber", m.SourceChartNumber())
	w.StringField("m_SourceChartbookFileName", m.SourceChartbookFileName())
	w.BoolField("m_IsAutomatedOrder", m.IsAutomatedOrder())
	w.Uint8Field("m_SimulatedOrder", m.SimulatedOrder())
	w.BoolField("m_IsChartReplaying", m.IsChartReplaying())
	w.Int32Field("m_AttachedOrderOCOGroupNumber", m.AttachedOrderOCOGroupNumber())
	w.StringField("m_LastFillExecutionServiceID", m.LastFillExecutionServiceID())
	w.Int32Field("m_FillCount", m.FillCount())
	w.Float64Field("m_LastFillQuantity", m.LastFillQuantity())
	w.Float64Field("m_LastFillPrice", m.LastFillPrice())
	w.Int64Field("m_LastFillDateTimeUTC", m.LastFillDateTimeUTC())
	w.Uint64Field("m_RejectedStopOCOSiblingInternalOrderID", m.RejectedStopOCOSiblingInternalOrderID())
	w.Float64Field("m_RejectedStopReplacementMarketOrderQuantity", m.RejectedStopReplacementMarketOrderQuantity())
	w.Uint8Field("m_EvaluatingForFill", m.EvaluatingForFill())
	w.Uint32Field("m_LastProcessedTimeSalesRecordSequenceForPrices", m.LastProcessedTimeSalesRecordSequenceForPrices())
	w.BoolField("m_IsMarketDataManagementOfOrderEnabled", m.IsMarketDataManagementOfOrderEnabled())
	w.StringField("m_TextTag", m.TextTag())
	w.Int64Field("m_TimedOutOrderRequestedStatusDateTime", m.TimedOutOrderRequestedStatusDateTime())
	w.Uint8Field("m_RequestedStatusForTimedOutOrder", m.RequestedStatusForTimedOutOrder())
	w.Uint8Field("m_SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled", m.SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled())
	w.Float64Field("m_QuantityToIncreaseFromParentFill", m.QuantityToIncreaseFromParentFill())
	w.Float64Field("m_MoveToBreakevenStopReferencePrice", m.MoveToBreakevenStopReferencePrice())
	w.Float64Field("m_QuantityTriggeredStop_QuantityForTrigger", m.QuantityTriggeredStop_QuantityForTrigger())
	w.Float64Field("m_AccumulatedTradeVolumeForTriggeredStop", m.AccumulatedTradeVolumeForTriggeredStop())
	w.Uint8Field("m_BidAskQuantityStopInitialTriggerMet", m.BidAskQuantityStopInitialTriggerMet())
	w.Uint8Field("m_NeedToOverrideLock", m.NeedToOverrideLock())
	w.Float64Field("m_CurrentMarketPrice", m.CurrentMarketPrice())
	w.Int64Field("m_CurrentMarketDateTime", m.CurrentMarketDateTime())
	w.Uint8Field("m_SupportOrderFillBilling", m.SupportOrderFillBilling())
	w.BoolField("m_IsBillable", m.IsBillable())
	w.Int32Field("m_QuantityForBilling", m.QuantityForBilling())
	w.Uint32Field("m_NumberOfFailedOrderModifications", m.NumberOfFailedOrderModifications())
	w.Int32Field("m_DTCServerIndex", m.DTCServerIndex())
	w.StringField("m_ClearingFirmID", m.ClearingFirmID())
	w.StringField("m_SenderSubID", m.SenderSubID())
	w.StringField("m_SenderLocationId", m.SenderLocationId())
	w.StringField("m_SelfMatchPreventionID", m.SelfMatchPreventionID())
	w.Int32Field("m_CTICode", m.CTICode())
	w.Uint8Field("m_ObtainOrderActionDateTimeFromLastTradeTimeInChart", m.ObtainOrderActionDateTimeFromLastTradeTimeInChart())
	w.Float64Field("m_MaximumShowQuantity", m.MaximumShowQuantity())
	w.Uint8Field("m_OrderSubmitted", m.OrderSubmitted())
	w.BoolField("m_IsSnapshot", m.IsSnapshot())
	w.BoolField("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.Int64Field("m_ExternalLastActionDateTimeUTC", m.ExternalLastActionDateTimeUTC())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeOrderBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10110 {
		return message.ErrWrongType
	}
	m.SetIsOrderDeleted(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderStatusCode(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderStatusBeforePendingModify(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderStatusBeforePendingCancel(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetServiceOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetActualSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderType(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(r.Uint16())
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
	m.SetOrderQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetFilledQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAverageFillPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRealtimeFillStatus(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsRestingOrderDuringFill(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderRejectType(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetInternalOrderIDModifierForService(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetFIXClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSequenceNumberBasedClientOrderID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderIDForDTCServer(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPreviousClientOrderIDForDTCServer(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchangeOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOriginatingClientUsername(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetEntryDateTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastActionDateTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetServiceUpdateDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderEntryTimeForService(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastModifyTimeForService(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetGoodTillDateTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTimeInForce(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenClose(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailStopOffset1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailStopStep(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailTriggerPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailingStopTriggerOffset(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailTriggerHit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailToBreakEvenStopOffset(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumChaseAmountAsPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetInitialChaseOrderPrice1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetInitialLastTradePriceForChaseOrders(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailingStopTriggerOCOGroupNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastModifyPrice1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastModifyQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetCumulativeOrderQuantityFromParentFills(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriorFilledQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTickSize(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetValueFormat(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriceMultiplier(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetParentInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTargetChildInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetStopChildInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAttachedOrderPriceOffset1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLinkInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOGroupInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOSiblingInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDisableChildAndSiblingRelatedActions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOManagedByService(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetBracketOrderServerManaged(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastOrderActionSource(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStopLimitOrderStopPriceTriggered(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOFullSiblingCancelOnPartialFill(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetReverseOnCompleteFill(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSupportScaleIn(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSupportScaleOut(r.Uint8())
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
	m.SetIsAutomatedOrder(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetSimulatedOrder(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsChartReplaying(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetAttachedOrderOCOGroupNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastFillExecutionServiceID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetFillCount(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastFillQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastFillPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastFillDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectedStopOCOSiblingInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectedStopReplacementMarketOrderQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetEvaluatingForFill(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastProcessedTimeSalesRecordSequenceForPrices(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsMarketDataManagementOfOrderEnabled(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTextTag(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTimedOutOrderRequestedStatusDateTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestedStatusForTimedOutOrder(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantityToIncreaseFromParentFill(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMoveToBreakevenStopReferencePrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantityTriggeredStop_QuantityForTrigger(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccumulatedTradeVolumeForTriggeredStop(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidAskQuantityStopInitialTriggerMet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetNeedToOverrideLock(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentMarketPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentMarketDateTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetSupportOrderFillBilling(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsBillable(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantityForBilling(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfFailedOrderModifications(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDTCServerIndex(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetClearingFirmID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderSubID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderLocationId(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSelfMatchPreventionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCTICode(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetObtainOrderActionDateTimeFromLastTradeTimeInChart(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumShowQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderSubmitted(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSnapshot(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFirstMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetExternalLastActionDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeOrderPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10110 {
		return message.ErrWrongType
	}
	m.SetIsOrderDeleted(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderStatusCode(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderStatusBeforePendingModify(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderStatusBeforePendingCancel(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetServiceOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetActualSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderType(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(r.Uint16())
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
	m.SetOrderQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetFilledQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAverageFillPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRealtimeFillStatus(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsRestingOrderDuringFill(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderRejectType(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetInternalOrderIDModifierForService(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetFIXClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSequenceNumberBasedClientOrderID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderIDForDTCServer(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPreviousClientOrderIDForDTCServer(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchangeOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetOriginatingClientUsername(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetEntryDateTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastActionDateTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetServiceUpdateDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderEntryTimeForService(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastModifyTimeForService(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetGoodTillDateTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTimeInForce(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenClose(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailStopOffset1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailStopStep(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailTriggerPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailingStopTriggerOffset(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailTriggerHit(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailToBreakEvenStopOffset(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumChaseAmountAsPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetInitialChaseOrderPrice1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetInitialLastTradePriceForChaseOrders(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTrailingStopTriggerOCOGroupNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastModifyPrice1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastModifyQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetCumulativeOrderQuantityFromParentFills(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriorFilledQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTickSize(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetValueFormat(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriceMultiplier(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetParentInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTargetChildInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetStopChildInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAttachedOrderPriceOffset1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLinkInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOGroupInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOSiblingInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDisableChildAndSiblingRelatedActions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOManagedByService(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetBracketOrderServerManaged(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastOrderActionSource(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStopLimitOrderStopPriceTriggered(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetOCOFullSiblingCancelOnPartialFill(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetReverseOnCompleteFill(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSupportScaleIn(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSupportScaleOut(r.Uint8())
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
	m.SetIsAutomatedOrder(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetSimulatedOrder(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsChartReplaying(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetAttachedOrderOCOGroupNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastFillExecutionServiceID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetFillCount(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastFillQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastFillPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastFillDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectedStopOCOSiblingInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectedStopReplacementMarketOrderQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetEvaluatingForFill(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastProcessedTimeSalesRecordSequenceForPrices(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsMarketDataManagementOfOrderEnabled(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTextTag(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTimedOutOrderRequestedStatusDateTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestedStatusForTimedOutOrder(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantityToIncreaseFromParentFill(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMoveToBreakevenStopReferencePrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantityTriggeredStop_QuantityForTrigger(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccumulatedTradeVolumeForTriggeredStop(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidAskQuantityStopInitialTriggerMet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetNeedToOverrideLock(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentMarketPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentMarketDateTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetSupportOrderFillBilling(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsBillable(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantityForBilling(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfFailedOrderModifications(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDTCServerIndex(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetClearingFirmID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderSubID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSenderLocationId(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSelfMatchPreventionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCTICode(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetObtainOrderActionDateTimeFromLastTradeTimeInChart(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumShowQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOrderSubmitted(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSnapshot(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFirstMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetExternalLastActionDateTimeUTC(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeOrderBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10110 {
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
		case "m_IsOrderDeleted":
			m.SetIsOrderDeleted(r.Bool())
		case "m_InternalOrderID":
			m.SetInternalOrderID(r.Uint64())
		case "m_OrderStatusCode":
			m.SetOrderStatusCode(r.Uint16())
		case "m_OrderStatusBeforePendingModify":
			m.SetOrderStatusBeforePendingModify(r.Uint16())
		case "m_OrderStatusBeforePendingCancel":
			m.SetOrderStatusBeforePendingCancel(r.Uint16())
		case "m_ServiceOrderID":
			m.SetServiceOrderID(r.String())
		case "m_ActualSymbol":
			m.SetActualSymbol(r.String())
		case "m_OrderType":
			m.SetOrderType(r.Int32())
		case "m_BuySell":
			m.SetBuySell(r.Uint16())
		case "m_Price1":
			m.SetPrice1(r.Float64())
		case "m_Price2":
			m.SetPrice2(r.Float64())
		case "m_OrderQuantity":
			m.SetOrderQuantity(r.Float64())
		case "m_FilledQuantity":
			m.SetFilledQuantity(r.Float64())
		case "m_AverageFillPrice":
			m.SetAverageFillPrice(r.Float64())
		case "m_RealtimeFillStatus":
			m.SetRealtimeFillStatus(r.Int32())
		case "m_IsRestingOrderDuringFill":
			m.SetIsRestingOrderDuringFill(r.Bool())
		case "m_OrderRejectType":
			m.SetOrderRejectType(r.Int32())
		case "m_TradeAccount":
			m.SetTradeAccount(r.String())
		case "m_SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
		case "m_InternalOrderIDModifierForService":
			m.SetInternalOrderIDModifierForService(r.Int32())
		case "m_FIXClientOrderID":
			m.SetFIXClientOrderID(r.String())
		case "m_SequenceNumberBasedClientOrderID":
			m.SetSequenceNumberBasedClientOrderID(r.Uint32())
		case "m_ClientOrderIDForDTCServer":
			m.SetClientOrderIDForDTCServer(r.String())
		case "m_PreviousClientOrderIDForDTCServer":
			m.SetPreviousClientOrderIDForDTCServer(r.String())
		case "m_ExchangeOrderID":
			m.SetExchangeOrderID(r.String())
		case "m_OriginatingClientUsername":
			m.SetOriginatingClientUsername(r.String())
		case "m_EntryDateTime":
			m.SetEntryDateTime(r.Int64())
		case "m_LastActionDateTime":
			m.SetLastActionDateTime(r.Int64())
		case "m_ServiceUpdateDateTimeUTC":
			m.SetServiceUpdateDateTimeUTC(r.Int64())
		case "m_OrderEntryTimeForService":
			m.SetOrderEntryTimeForService(r.Uint32())
		case "m_LastModifyTimeForService":
			m.SetLastModifyTimeForService(r.Uint32())
		case "m_GoodTillDateTime":
			m.SetGoodTillDateTime(r.Int64())
		case "m_TimeInForce":
			m.SetTimeInForce(r.Int32())
		case "m_OpenClose":
			m.SetOpenClose(r.Uint16())
		case "m_TrailStopOffset1":
			m.SetTrailStopOffset1(r.Float64())
		case "m_TrailStopStep":
			m.SetTrailStopStep(r.Float64())
		case "m_TrailTriggerPrice":
			m.SetTrailTriggerPrice(r.Float64())
		case "m_TrailingStopTriggerOffset":
			m.SetTrailingStopTriggerOffset(r.Float64())
		case "m_TrailTriggerHit":
			m.SetTrailTriggerHit(r.Uint8())
		case "m_TrailToBreakEvenStopOffset":
			m.SetTrailToBreakEvenStopOffset(r.Float64())
		case "m_MaximumChaseAmountAsPrice":
			m.SetMaximumChaseAmountAsPrice(r.Float64())
		case "m_InitialChaseOrderPrice1":
			m.SetInitialChaseOrderPrice1(r.Float64())
		case "m_InitialLastTradePriceForChaseOrders":
			m.SetInitialLastTradePriceForChaseOrders(r.Float64())
		case "m_TrailingStopTriggerOCOGroupNumber":
			m.SetTrailingStopTriggerOCOGroupNumber(r.Int32())
		case "m_LastModifyPrice1":
			m.SetLastModifyPrice1(r.Float64())
		case "m_LastModifyQuantity":
			m.SetLastModifyQuantity(r.Float64())
		case "m_CumulativeOrderQuantityFromParentFills":
			m.SetCumulativeOrderQuantityFromParentFills(r.Float64())
		case "m_PriorFilledQuantity":
			m.SetPriorFilledQuantity(r.Float64())
		case "m_TickSize":
			m.SetTickSize(r.Float32())
		case "m_ValueFormat":
			m.SetValueFormat(r.Int32())
		case "m_PriceMultiplier":
			m.SetPriceMultiplier(r.Float32())
		case "m_ParentInternalOrderID":
			m.SetParentInternalOrderID(r.Uint64())
		case "m_TargetChildInternalOrderID":
			m.SetTargetChildInternalOrderID(r.Uint64())
		case "m_StopChildInternalOrderID":
			m.SetStopChildInternalOrderID(r.Uint64())
		case "m_AttachedOrderPriceOffset1":
			m.SetAttachedOrderPriceOffset1(r.Float64())
		case "m_LinkInternalOrderID":
			m.SetLinkInternalOrderID(r.Uint64())
		case "m_OCOGroupInternalOrderID":
			m.SetOCOGroupInternalOrderID(r.Uint64())
		case "m_OCOSiblingInternalOrderID":
			m.SetOCOSiblingInternalOrderID(r.Uint64())
		case "m_DisableChildAndSiblingRelatedActions":
			m.SetDisableChildAndSiblingRelatedActions(r.Uint8())
		case "m_OCOManagedByService":
			m.SetOCOManagedByService(r.Uint8())
		case "m_BracketOrderServerManaged":
			m.SetBracketOrderServerManaged(r.Uint8())
		case "m_LastOrderActionSource":
			m.SetLastOrderActionSource(r.String())
		case "m_StopLimitOrderStopPriceTriggered":
			m.SetStopLimitOrderStopPriceTriggered(r.Uint8())
		case "m_OCOFullSiblingCancelOnPartialFill":
			m.SetOCOFullSiblingCancelOnPartialFill(r.Uint8())
		case "m_ReverseOnCompleteFill":
			m.SetReverseOnCompleteFill(r.Uint8())
		case "m_SupportScaleIn":
			m.SetSupportScaleIn(r.Uint8())
		case "m_SupportScaleOut":
			m.SetSupportScaleOut(r.Uint8())
		case "m_SourceChartNumber":
			m.SetSourceChartNumber(r.Int32())
		case "m_SourceChartbookFileName":
			m.SetSourceChartbookFileName(r.String())
		case "m_IsAutomatedOrder":
			m.SetIsAutomatedOrder(r.Bool())
		case "m_SimulatedOrder":
			m.SetSimulatedOrder(r.Uint8())
		case "m_IsChartReplaying":
			m.SetIsChartReplaying(r.Bool())
		case "m_AttachedOrderOCOGroupNumber":
			m.SetAttachedOrderOCOGroupNumber(r.Int32())
		case "m_LastFillExecutionServiceID":
			m.SetLastFillExecutionServiceID(r.String())
		case "m_FillCount":
			m.SetFillCount(r.Int32())
		case "m_LastFillQuantity":
			m.SetLastFillQuantity(r.Float64())
		case "m_LastFillPrice":
			m.SetLastFillPrice(r.Float64())
		case "m_LastFillDateTimeUTC":
			m.SetLastFillDateTimeUTC(r.Int64())
		case "m_RejectedStopOCOSiblingInternalOrderID":
			m.SetRejectedStopOCOSiblingInternalOrderID(r.Uint64())
		case "m_RejectedStopReplacementMarketOrderQuantity":
			m.SetRejectedStopReplacementMarketOrderQuantity(r.Float64())
		case "m_EvaluatingForFill":
			m.SetEvaluatingForFill(r.Uint8())
		case "m_LastProcessedTimeSalesRecordSequenceForPrices":
			m.SetLastProcessedTimeSalesRecordSequenceForPrices(r.Uint32())
		case "m_IsMarketDataManagementOfOrderEnabled":
			m.SetIsMarketDataManagementOfOrderEnabled(r.Bool())
		case "m_TextTag":
			m.SetTextTag(r.String())
		case "m_TimedOutOrderRequestedStatusDateTime":
			m.SetTimedOutOrderRequestedStatusDateTime(r.Int64())
		case "m_RequestedStatusForTimedOutOrder":
			m.SetRequestedStatusForTimedOutOrder(r.Uint8())
		case "m_SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled":
			m.SetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled(r.Uint8())
		case "m_QuantityToIncreaseFromParentFill":
			m.SetQuantityToIncreaseFromParentFill(r.Float64())
		case "m_MoveToBreakevenStopReferencePrice":
			m.SetMoveToBreakevenStopReferencePrice(r.Float64())
		case "m_QuantityTriggeredStop_QuantityForTrigger":
			m.SetQuantityTriggeredStop_QuantityForTrigger(r.Float64())
		case "m_AccumulatedTradeVolumeForTriggeredStop":
			m.SetAccumulatedTradeVolumeForTriggeredStop(r.Float64())
		case "m_BidAskQuantityStopInitialTriggerMet":
			m.SetBidAskQuantityStopInitialTriggerMet(r.Uint8())
		case "m_NeedToOverrideLock":
			m.SetNeedToOverrideLock(r.Uint8())
		case "m_CurrentMarketPrice":
			m.SetCurrentMarketPrice(r.Float64())
		case "m_CurrentMarketDateTime":
			m.SetCurrentMarketDateTime(r.Int64())
		case "m_SupportOrderFillBilling":
			m.SetSupportOrderFillBilling(r.Uint8())
		case "m_IsBillable":
			m.SetIsBillable(r.Bool())
		case "m_QuantityForBilling":
			m.SetQuantityForBilling(r.Int32())
		case "m_NumberOfFailedOrderModifications":
			m.SetNumberOfFailedOrderModifications(r.Uint32())
		case "m_DTCServerIndex":
			m.SetDTCServerIndex(r.Int32())
		case "m_ClearingFirmID":
			m.SetClearingFirmID(r.String())
		case "m_SenderSubID":
			m.SetSenderSubID(r.String())
		case "m_SenderLocationId":
			m.SetSenderLocationId(r.String())
		case "m_SelfMatchPreventionID":
			m.SetSelfMatchPreventionID(r.String())
		case "m_CTICode":
			m.SetCTICode(r.Int32())
		case "m_ObtainOrderActionDateTimeFromLastTradeTimeInChart":
			m.SetObtainOrderActionDateTimeFromLastTradeTimeInChart(r.Uint8())
		case "m_MaximumShowQuantity":
			m.SetMaximumShowQuantity(r.Float64())
		case "m_OrderSubmitted":
			m.SetOrderSubmitted(r.Uint8())
		case "m_IsSnapshot":
			m.SetIsSnapshot(r.Bool())
		case "m_IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Bool())
		case "m_IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Bool())
		case "m_ExternalLastActionDateTimeUTC":
			m.SetExternalLastActionDateTimeUTC(r.Int64())
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

func (m *TradeOrderPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10110 {
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
		case "m_IsOrderDeleted":
			m.SetIsOrderDeleted(r.Bool())
		case "m_InternalOrderID":
			m.SetInternalOrderID(r.Uint64())
		case "m_OrderStatusCode":
			m.SetOrderStatusCode(r.Uint16())
		case "m_OrderStatusBeforePendingModify":
			m.SetOrderStatusBeforePendingModify(r.Uint16())
		case "m_OrderStatusBeforePendingCancel":
			m.SetOrderStatusBeforePendingCancel(r.Uint16())
		case "m_ServiceOrderID":
			m.SetServiceOrderID(r.String())
		case "m_ActualSymbol":
			m.SetActualSymbol(r.String())
		case "m_OrderType":
			m.SetOrderType(r.Int32())
		case "m_BuySell":
			m.SetBuySell(r.Uint16())
		case "m_Price1":
			m.SetPrice1(r.Float64())
		case "m_Price2":
			m.SetPrice2(r.Float64())
		case "m_OrderQuantity":
			m.SetOrderQuantity(r.Float64())
		case "m_FilledQuantity":
			m.SetFilledQuantity(r.Float64())
		case "m_AverageFillPrice":
			m.SetAverageFillPrice(r.Float64())
		case "m_RealtimeFillStatus":
			m.SetRealtimeFillStatus(r.Int32())
		case "m_IsRestingOrderDuringFill":
			m.SetIsRestingOrderDuringFill(r.Bool())
		case "m_OrderRejectType":
			m.SetOrderRejectType(r.Int32())
		case "m_TradeAccount":
			m.SetTradeAccount(r.String())
		case "m_SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
		case "m_InternalOrderIDModifierForService":
			m.SetInternalOrderIDModifierForService(r.Int32())
		case "m_FIXClientOrderID":
			m.SetFIXClientOrderID(r.String())
		case "m_SequenceNumberBasedClientOrderID":
			m.SetSequenceNumberBasedClientOrderID(r.Uint32())
		case "m_ClientOrderIDForDTCServer":
			m.SetClientOrderIDForDTCServer(r.String())
		case "m_PreviousClientOrderIDForDTCServer":
			m.SetPreviousClientOrderIDForDTCServer(r.String())
		case "m_ExchangeOrderID":
			m.SetExchangeOrderID(r.String())
		case "m_OriginatingClientUsername":
			m.SetOriginatingClientUsername(r.String())
		case "m_EntryDateTime":
			m.SetEntryDateTime(r.Int64())
		case "m_LastActionDateTime":
			m.SetLastActionDateTime(r.Int64())
		case "m_ServiceUpdateDateTimeUTC":
			m.SetServiceUpdateDateTimeUTC(r.Int64())
		case "m_OrderEntryTimeForService":
			m.SetOrderEntryTimeForService(r.Uint32())
		case "m_LastModifyTimeForService":
			m.SetLastModifyTimeForService(r.Uint32())
		case "m_GoodTillDateTime":
			m.SetGoodTillDateTime(r.Int64())
		case "m_TimeInForce":
			m.SetTimeInForce(r.Int32())
		case "m_OpenClose":
			m.SetOpenClose(r.Uint16())
		case "m_TrailStopOffset1":
			m.SetTrailStopOffset1(r.Float64())
		case "m_TrailStopStep":
			m.SetTrailStopStep(r.Float64())
		case "m_TrailTriggerPrice":
			m.SetTrailTriggerPrice(r.Float64())
		case "m_TrailingStopTriggerOffset":
			m.SetTrailingStopTriggerOffset(r.Float64())
		case "m_TrailTriggerHit":
			m.SetTrailTriggerHit(r.Uint8())
		case "m_TrailToBreakEvenStopOffset":
			m.SetTrailToBreakEvenStopOffset(r.Float64())
		case "m_MaximumChaseAmountAsPrice":
			m.SetMaximumChaseAmountAsPrice(r.Float64())
		case "m_InitialChaseOrderPrice1":
			m.SetInitialChaseOrderPrice1(r.Float64())
		case "m_InitialLastTradePriceForChaseOrders":
			m.SetInitialLastTradePriceForChaseOrders(r.Float64())
		case "m_TrailingStopTriggerOCOGroupNumber":
			m.SetTrailingStopTriggerOCOGroupNumber(r.Int32())
		case "m_LastModifyPrice1":
			m.SetLastModifyPrice1(r.Float64())
		case "m_LastModifyQuantity":
			m.SetLastModifyQuantity(r.Float64())
		case "m_CumulativeOrderQuantityFromParentFills":
			m.SetCumulativeOrderQuantityFromParentFills(r.Float64())
		case "m_PriorFilledQuantity":
			m.SetPriorFilledQuantity(r.Float64())
		case "m_TickSize":
			m.SetTickSize(r.Float32())
		case "m_ValueFormat":
			m.SetValueFormat(r.Int32())
		case "m_PriceMultiplier":
			m.SetPriceMultiplier(r.Float32())
		case "m_ParentInternalOrderID":
			m.SetParentInternalOrderID(r.Uint64())
		case "m_TargetChildInternalOrderID":
			m.SetTargetChildInternalOrderID(r.Uint64())
		case "m_StopChildInternalOrderID":
			m.SetStopChildInternalOrderID(r.Uint64())
		case "m_AttachedOrderPriceOffset1":
			m.SetAttachedOrderPriceOffset1(r.Float64())
		case "m_LinkInternalOrderID":
			m.SetLinkInternalOrderID(r.Uint64())
		case "m_OCOGroupInternalOrderID":
			m.SetOCOGroupInternalOrderID(r.Uint64())
		case "m_OCOSiblingInternalOrderID":
			m.SetOCOSiblingInternalOrderID(r.Uint64())
		case "m_DisableChildAndSiblingRelatedActions":
			m.SetDisableChildAndSiblingRelatedActions(r.Uint8())
		case "m_OCOManagedByService":
			m.SetOCOManagedByService(r.Uint8())
		case "m_BracketOrderServerManaged":
			m.SetBracketOrderServerManaged(r.Uint8())
		case "m_LastOrderActionSource":
			m.SetLastOrderActionSource(r.String())
		case "m_StopLimitOrderStopPriceTriggered":
			m.SetStopLimitOrderStopPriceTriggered(r.Uint8())
		case "m_OCOFullSiblingCancelOnPartialFill":
			m.SetOCOFullSiblingCancelOnPartialFill(r.Uint8())
		case "m_ReverseOnCompleteFill":
			m.SetReverseOnCompleteFill(r.Uint8())
		case "m_SupportScaleIn":
			m.SetSupportScaleIn(r.Uint8())
		case "m_SupportScaleOut":
			m.SetSupportScaleOut(r.Uint8())
		case "m_SourceChartNumber":
			m.SetSourceChartNumber(r.Int32())
		case "m_SourceChartbookFileName":
			m.SetSourceChartbookFileName(r.String())
		case "m_IsAutomatedOrder":
			m.SetIsAutomatedOrder(r.Bool())
		case "m_SimulatedOrder":
			m.SetSimulatedOrder(r.Uint8())
		case "m_IsChartReplaying":
			m.SetIsChartReplaying(r.Bool())
		case "m_AttachedOrderOCOGroupNumber":
			m.SetAttachedOrderOCOGroupNumber(r.Int32())
		case "m_LastFillExecutionServiceID":
			m.SetLastFillExecutionServiceID(r.String())
		case "m_FillCount":
			m.SetFillCount(r.Int32())
		case "m_LastFillQuantity":
			m.SetLastFillQuantity(r.Float64())
		case "m_LastFillPrice":
			m.SetLastFillPrice(r.Float64())
		case "m_LastFillDateTimeUTC":
			m.SetLastFillDateTimeUTC(r.Int64())
		case "m_RejectedStopOCOSiblingInternalOrderID":
			m.SetRejectedStopOCOSiblingInternalOrderID(r.Uint64())
		case "m_RejectedStopReplacementMarketOrderQuantity":
			m.SetRejectedStopReplacementMarketOrderQuantity(r.Float64())
		case "m_EvaluatingForFill":
			m.SetEvaluatingForFill(r.Uint8())
		case "m_LastProcessedTimeSalesRecordSequenceForPrices":
			m.SetLastProcessedTimeSalesRecordSequenceForPrices(r.Uint32())
		case "m_IsMarketDataManagementOfOrderEnabled":
			m.SetIsMarketDataManagementOfOrderEnabled(r.Bool())
		case "m_TextTag":
			m.SetTextTag(r.String())
		case "m_TimedOutOrderRequestedStatusDateTime":
			m.SetTimedOutOrderRequestedStatusDateTime(r.Int64())
		case "m_RequestedStatusForTimedOutOrder":
			m.SetRequestedStatusForTimedOutOrder(r.Uint8())
		case "m_SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled":
			m.SetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled(r.Uint8())
		case "m_QuantityToIncreaseFromParentFill":
			m.SetQuantityToIncreaseFromParentFill(r.Float64())
		case "m_MoveToBreakevenStopReferencePrice":
			m.SetMoveToBreakevenStopReferencePrice(r.Float64())
		case "m_QuantityTriggeredStop_QuantityForTrigger":
			m.SetQuantityTriggeredStop_QuantityForTrigger(r.Float64())
		case "m_AccumulatedTradeVolumeForTriggeredStop":
			m.SetAccumulatedTradeVolumeForTriggeredStop(r.Float64())
		case "m_BidAskQuantityStopInitialTriggerMet":
			m.SetBidAskQuantityStopInitialTriggerMet(r.Uint8())
		case "m_NeedToOverrideLock":
			m.SetNeedToOverrideLock(r.Uint8())
		case "m_CurrentMarketPrice":
			m.SetCurrentMarketPrice(r.Float64())
		case "m_CurrentMarketDateTime":
			m.SetCurrentMarketDateTime(r.Int64())
		case "m_SupportOrderFillBilling":
			m.SetSupportOrderFillBilling(r.Uint8())
		case "m_IsBillable":
			m.SetIsBillable(r.Bool())
		case "m_QuantityForBilling":
			m.SetQuantityForBilling(r.Int32())
		case "m_NumberOfFailedOrderModifications":
			m.SetNumberOfFailedOrderModifications(r.Uint32())
		case "m_DTCServerIndex":
			m.SetDTCServerIndex(r.Int32())
		case "m_ClearingFirmID":
			m.SetClearingFirmID(r.String())
		case "m_SenderSubID":
			m.SetSenderSubID(r.String())
		case "m_SenderLocationId":
			m.SetSenderLocationId(r.String())
		case "m_SelfMatchPreventionID":
			m.SetSelfMatchPreventionID(r.String())
		case "m_CTICode":
			m.SetCTICode(r.Int32())
		case "m_ObtainOrderActionDateTimeFromLastTradeTimeInChart":
			m.SetObtainOrderActionDateTimeFromLastTradeTimeInChart(r.Uint8())
		case "m_MaximumShowQuantity":
			m.SetMaximumShowQuantity(r.Float64())
		case "m_OrderSubmitted":
			m.SetOrderSubmitted(r.Uint8())
		case "m_IsSnapshot":
			m.SetIsSnapshot(r.Bool())
		case "m_IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Bool())
		case "m_IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Bool())
		case "m_ExternalLastActionDateTimeUTC":
			m.SetExternalLastActionDateTimeUTC(r.Int64())
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