package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                                                     = TradeOrderSize  (535)
//     Type                                                     = SC_TRADE_ORDER  (10110)
//     BaseSize                                                 = TradeOrderSize  (535)
//     IsOrderDeleted                                           = 0
//     InternalOrderID                                          = 0
//     OrderStatusCode                                          = 0
//     OrderStatusBeforePendingModify                           = 0
//     OrderStatusBeforePendingCancel                           = 0
//     ServiceOrderID                                           = ""
//     ActualSymbol                                             = ""
//     OrderType                                                = 0
//     BuySell                                                  = 0
//     Price1                                                   = 0
//     Price2                                                   = 0
//     OrderQuantity                                            = 0
//     FilledQuantity                                           = 0
//     AverageFillPrice                                         = 0
//     RealtimeFillStatus                                       = 0
//     IsRestingOrderDuringFill                                 = 0
//     OrderRejectType                                          = 0
//     TradeAccount                                             = ""
//     SubAccountIdentifier                                     = 0
//     InternalOrderIDModifierForService                        = 0
//     FIXClientOrderID                                         = ""
//     SequenceNumberBasedClientOrderID                         = 0
//     ClientOrderIDForDTCServer                                = ""
//     PreviousClientOrderIDForDTCServer                        = ""
//     ExchangeOrderID                                          = ""
//     OriginatingClientUsername                                = ""
//     EntryDateTime                                            = 0
//     LastActionDateTime                                       = 0
//     ServiceUpdateDateTimeUTC                                 = 0
//     OrderEntryTimeForService                                 = 0
//     LastModifyTimeForService                                 = 0
//     GoodTillDateTime                                         = 0
//     TimeInForce                                              = 0
//     OpenClose                                                = 0
//     TrailStopOffset1                                         = 0
//     TrailStopStep                                            = 0
//     TrailTriggerPrice                                        = 0
//     TrailingStopTriggerOffset                                = 0
//     TrailTriggerHit                                          = 0
//     TrailToBreakEvenStopOffset                               = 0
//     MaximumChaseAmountAsPrice                                = 0
//     InitialChaseOrderPrice1                                  = 0
//     InitialLastTradePriceForChaseOrders                      = 0
//     TrailingStopTriggerOCOGroupNumber                        = 0
//     LastModifyPrice1                                         = 0
//     LastModifyQuantity                                       = 0
//     CumulativeOrderQuantityFromParentFills                   = 0
//     PriorFilledQuantity                                      = 0
//     TickSize                                                 = 0
//     ValueFormat                                              = 0
//     PriceMultiplier                                          = 0
//     ParentInternalOrderID                                    = 0
//     TargetChildInternalOrderID                               = 0
//     StopChildInternalOrderID                                 = 0
//     AttachedOrderPriceOffset1                                = 0
//     LinkInternalOrderID                                      = 0
//     OCOGroupInternalOrderID                                  = 0
//     OCOSiblingInternalOrderID                                = 0
//     DisableChildAndSiblingRelatedActions                     = 0
//     OCOManagedByService                                      = 0
//     BracketOrderServerManaged                                = 0
//     LastOrderActionSource                                    = ""
//     StopLimitOrderStopPriceTriggered                         = 0
//     OCOFullSiblingCancelOnPartialFill                        = 0
//     ReverseOnCompleteFill                                    = 0
//     SupportScaleIn                                           = 0
//     SupportScaleOut                                          = 0
//     SourceChartNumber                                        = 0
//     SourceChartbookFileName                                  = ""
//     IsAutomatedOrder                                         = 0
//     SimulatedOrder                                           = 0
//     IsChartReplaying                                         = 0
//     AttachedOrderOCOGroupNumber                              = 0
//     LastFillExecutionServiceID                               = ""
//     FillCount                                                = 0
//     LastFillQuantity                                         = 0
//     LastFillPrice                                            = 0
//     LastFillDateTimeUTC                                      = 0
//     RejectedStopOCOSiblingInternalOrderID                    = 0
//     RejectedStopReplacementMarketOrderQuantity               = 0
//     EvaluatingForFill                                        = 0
//     LastProcessedTimeSalesRecordSequenceForPrices            = 0
//     IsMarketDataManagementOfOrderEnabled                     = 0
//     TextTag                                                  = ""
//     TimedOutOrderRequestedStatusDateTime                     = 0
//     RequestedStatusForTimedOutOrder                          = 0
//     SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled = 0
//     QuantityToIncreaseFromParentFill                         = 0
//     MoveToBreakevenStopReferencePrice                        = 0
//     QuantityTriggeredStop_QuantityForTrigger                 = 0
//     AccumulatedTradeVolumeForTriggeredStop                   = 0
//     BidAskQuantityStopInitialTriggerMet                      = 0
//     NeedToOverrideLock                                       = 0
//     CurrentMarketPrice                                       = 0
//     CurrentMarketDateTime                                    = 0
//     SupportOrderFillBilling                                  = 0
//     IsBillable                                               = 0
//     QuantityForBilling                                       = 0
//     NumberOfFailedOrderModifications                         = 0
//     DTCServerIndex                                           = 0
//     ClearingFirmID                                           = ""
//     SenderSubID                                              = ""
//     SenderLocationId                                         = ""
//     SelfMatchPreventionID                                    = ""
//     CTICode                                                  = 0
//     ObtainOrderActionDateTimeFromLastTradeTimeInChart        = 0
//     MaximumShowQuantity                                      = 0
//     OrderSubmitted                                           = 0
//     IsSnapshot                                               = 0
//     IsFirstMessageInBatch                                    = 0
//     IsLastMessageInBatch                                     = 0
//     ExternalLastActionDateTimeUTC                            = 0
// }
var _TradeOrderDefault = []byte{23, 2, 126, 39, 23, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeOrderSize = 535

type TradeOrder struct {
	message.VLS
}

type TradeOrderBuilder struct {
	message.VLSBuilder
}

func NewTradeOrderFrom(b []byte) TradeOrder {
	return TradeOrder{VLS: message.NewVLSFrom(b)}
}

func WrapTradeOrder(b []byte) TradeOrder {
	return TradeOrder{VLS: message.WrapVLS(b)}
}

func NewTradeOrder() TradeOrderBuilder {
	a := TradeOrderBuilder{message.NewVLSBuilder(535)}
	a.Unsafe().SetBytes(0, _TradeOrderDefault)
	return a
}

// Clear
// {
//     Size                                                     = TradeOrderSize  (535)
//     Type                                                     = SC_TRADE_ORDER  (10110)
//     BaseSize                                                 = TradeOrderSize  (535)
//     IsOrderDeleted                                           = 0
//     InternalOrderID                                          = 0
//     OrderStatusCode                                          = 0
//     OrderStatusBeforePendingModify                           = 0
//     OrderStatusBeforePendingCancel                           = 0
//     ServiceOrderID                                           = ""
//     ActualSymbol                                             = ""
//     OrderType                                                = 0
//     BuySell                                                  = 0
//     Price1                                                   = 0
//     Price2                                                   = 0
//     OrderQuantity                                            = 0
//     FilledQuantity                                           = 0
//     AverageFillPrice                                         = 0
//     RealtimeFillStatus                                       = 0
//     IsRestingOrderDuringFill                                 = 0
//     OrderRejectType                                          = 0
//     TradeAccount                                             = ""
//     SubAccountIdentifier                                     = 0
//     InternalOrderIDModifierForService                        = 0
//     FIXClientOrderID                                         = ""
//     SequenceNumberBasedClientOrderID                         = 0
//     ClientOrderIDForDTCServer                                = ""
//     PreviousClientOrderIDForDTCServer                        = ""
//     ExchangeOrderID                                          = ""
//     OriginatingClientUsername                                = ""
//     EntryDateTime                                            = 0
//     LastActionDateTime                                       = 0
//     ServiceUpdateDateTimeUTC                                 = 0
//     OrderEntryTimeForService                                 = 0
//     LastModifyTimeForService                                 = 0
//     GoodTillDateTime                                         = 0
//     TimeInForce                                              = 0
//     OpenClose                                                = 0
//     TrailStopOffset1                                         = 0
//     TrailStopStep                                            = 0
//     TrailTriggerPrice                                        = 0
//     TrailingStopTriggerOffset                                = 0
//     TrailTriggerHit                                          = 0
//     TrailToBreakEvenStopOffset                               = 0
//     MaximumChaseAmountAsPrice                                = 0
//     InitialChaseOrderPrice1                                  = 0
//     InitialLastTradePriceForChaseOrders                      = 0
//     TrailingStopTriggerOCOGroupNumber                        = 0
//     LastModifyPrice1                                         = 0
//     LastModifyQuantity                                       = 0
//     CumulativeOrderQuantityFromParentFills                   = 0
//     PriorFilledQuantity                                      = 0
//     TickSize                                                 = 0
//     ValueFormat                                              = 0
//     PriceMultiplier                                          = 0
//     ParentInternalOrderID                                    = 0
//     TargetChildInternalOrderID                               = 0
//     StopChildInternalOrderID                                 = 0
//     AttachedOrderPriceOffset1                                = 0
//     LinkInternalOrderID                                      = 0
//     OCOGroupInternalOrderID                                  = 0
//     OCOSiblingInternalOrderID                                = 0
//     DisableChildAndSiblingRelatedActions                     = 0
//     OCOManagedByService                                      = 0
//     BracketOrderServerManaged                                = 0
//     LastOrderActionSource                                    = ""
//     StopLimitOrderStopPriceTriggered                         = 0
//     OCOFullSiblingCancelOnPartialFill                        = 0
//     ReverseOnCompleteFill                                    = 0
//     SupportScaleIn                                           = 0
//     SupportScaleOut                                          = 0
//     SourceChartNumber                                        = 0
//     SourceChartbookFileName                                  = ""
//     IsAutomatedOrder                                         = 0
//     SimulatedOrder                                           = 0
//     IsChartReplaying                                         = 0
//     AttachedOrderOCOGroupNumber                              = 0
//     LastFillExecutionServiceID                               = ""
//     FillCount                                                = 0
//     LastFillQuantity                                         = 0
//     LastFillPrice                                            = 0
//     LastFillDateTimeUTC                                      = 0
//     RejectedStopOCOSiblingInternalOrderID                    = 0
//     RejectedStopReplacementMarketOrderQuantity               = 0
//     EvaluatingForFill                                        = 0
//     LastProcessedTimeSalesRecordSequenceForPrices            = 0
//     IsMarketDataManagementOfOrderEnabled                     = 0
//     TextTag                                                  = ""
//     TimedOutOrderRequestedStatusDateTime                     = 0
//     RequestedStatusForTimedOutOrder                          = 0
//     SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled = 0
//     QuantityToIncreaseFromParentFill                         = 0
//     MoveToBreakevenStopReferencePrice                        = 0
//     QuantityTriggeredStop_QuantityForTrigger                 = 0
//     AccumulatedTradeVolumeForTriggeredStop                   = 0
//     BidAskQuantityStopInitialTriggerMet                      = 0
//     NeedToOverrideLock                                       = 0
//     CurrentMarketPrice                                       = 0
//     CurrentMarketDateTime                                    = 0
//     SupportOrderFillBilling                                  = 0
//     IsBillable                                               = 0
//     QuantityForBilling                                       = 0
//     NumberOfFailedOrderModifications                         = 0
//     DTCServerIndex                                           = 0
//     ClearingFirmID                                           = ""
//     SenderSubID                                              = ""
//     SenderLocationId                                         = ""
//     SelfMatchPreventionID                                    = ""
//     CTICode                                                  = 0
//     ObtainOrderActionDateTimeFromLastTradeTimeInChart        = 0
//     MaximumShowQuantity                                      = 0
//     OrderSubmitted                                           = 0
//     IsSnapshot                                               = 0
//     IsFirstMessageInBatch                                    = 0
//     IsLastMessageInBatch                                     = 0
//     ExternalLastActionDateTimeUTC                            = 0
// }
func (m TradeOrderBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeOrderDefault)
}

// ToBuilder
func (m TradeOrder) ToBuilder() TradeOrderBuilder {
	return TradeOrderBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeOrderBuilder) Finish() TradeOrder {
	return TradeOrder{m.VLSBuilder.Finish()}
}

// IsOrderDeleted
func (m TradeOrder) IsOrderDeleted() uint8 {
	return message.Uint8VLS(m.Unsafe(), 7, 6)
}

// IsOrderDeleted
func (m TradeOrderBuilder) IsOrderDeleted() uint8 {
	return message.Uint8VLS(m.Unsafe(), 7, 6)
}

// InternalOrderID
func (m TradeOrder) InternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 15, 7)
}

// InternalOrderID
func (m TradeOrderBuilder) InternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 15, 7)
}

// OrderStatusCode
func (m TradeOrder) OrderStatusCode() uint16 {
	return message.Uint16VLS(m.Unsafe(), 17, 15)
}

// OrderStatusCode
func (m TradeOrderBuilder) OrderStatusCode() uint16 {
	return message.Uint16VLS(m.Unsafe(), 17, 15)
}

// OrderStatusBeforePendingModify
func (m TradeOrder) OrderStatusBeforePendingModify() uint16 {
	return message.Uint16VLS(m.Unsafe(), 19, 17)
}

// OrderStatusBeforePendingModify
func (m TradeOrderBuilder) OrderStatusBeforePendingModify() uint16 {
	return message.Uint16VLS(m.Unsafe(), 19, 17)
}

// OrderStatusBeforePendingCancel
func (m TradeOrder) OrderStatusBeforePendingCancel() uint16 {
	return message.Uint16VLS(m.Unsafe(), 21, 19)
}

// OrderStatusBeforePendingCancel
func (m TradeOrderBuilder) OrderStatusBeforePendingCancel() uint16 {
	return message.Uint16VLS(m.Unsafe(), 21, 19)
}

// ServiceOrderID
func (m TradeOrder) ServiceOrderID() string {
	return message.StringVLS(m.Unsafe(), 25, 21)
}

// ServiceOrderID
func (m TradeOrderBuilder) ServiceOrderID() string {
	return message.StringVLS(m.Unsafe(), 25, 21)
}

// ActualSymbol
func (m TradeOrder) ActualSymbol() string {
	return message.StringVLS(m.Unsafe(), 29, 25)
}

// ActualSymbol
func (m TradeOrderBuilder) ActualSymbol() string {
	return message.StringVLS(m.Unsafe(), 29, 25)
}

// OrderType
func (m TradeOrder) OrderType() int32 {
	return message.Int32VLS(m.Unsafe(), 33, 29)
}

// OrderType
func (m TradeOrderBuilder) OrderType() int32 {
	return message.Int32VLS(m.Unsafe(), 33, 29)
}

// BuySell
func (m TradeOrder) BuySell() uint16 {
	return message.Uint16VLS(m.Unsafe(), 35, 33)
}

// BuySell
func (m TradeOrderBuilder) BuySell() uint16 {
	return message.Uint16VLS(m.Unsafe(), 35, 33)
}

// Price1
func (m TradeOrder) Price1() float64 {
	return message.Float64VLS(m.Unsafe(), 43, 35)
}

// Price1
func (m TradeOrderBuilder) Price1() float64 {
	return message.Float64VLS(m.Unsafe(), 43, 35)
}

// Price2
func (m TradeOrder) Price2() float64 {
	return message.Float64VLS(m.Unsafe(), 51, 43)
}

// Price2
func (m TradeOrderBuilder) Price2() float64 {
	return message.Float64VLS(m.Unsafe(), 51, 43)
}

// OrderQuantity
func (m TradeOrder) OrderQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 59, 51)
}

// OrderQuantity
func (m TradeOrderBuilder) OrderQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 59, 51)
}

// FilledQuantity
func (m TradeOrder) FilledQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 67, 59)
}

// FilledQuantity
func (m TradeOrderBuilder) FilledQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 67, 59)
}

// AverageFillPrice
func (m TradeOrder) AverageFillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 75, 67)
}

// AverageFillPrice
func (m TradeOrderBuilder) AverageFillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 75, 67)
}

// RealtimeFillStatus
func (m TradeOrder) RealtimeFillStatus() int32 {
	return message.Int32VLS(m.Unsafe(), 79, 75)
}

// RealtimeFillStatus
func (m TradeOrderBuilder) RealtimeFillStatus() int32 {
	return message.Int32VLS(m.Unsafe(), 79, 75)
}

// IsRestingOrderDuringFill
func (m TradeOrder) IsRestingOrderDuringFill() uint8 {
	return message.Uint8VLS(m.Unsafe(), 80, 79)
}

// IsRestingOrderDuringFill
func (m TradeOrderBuilder) IsRestingOrderDuringFill() uint8 {
	return message.Uint8VLS(m.Unsafe(), 80, 79)
}

// OrderRejectType
func (m TradeOrder) OrderRejectType() int32 {
	return message.Int32VLS(m.Unsafe(), 84, 80)
}

// OrderRejectType
func (m TradeOrderBuilder) OrderRejectType() int32 {
	return message.Int32VLS(m.Unsafe(), 84, 80)
}

// TradeAccount
func (m TradeOrder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 88, 84)
}

// TradeAccount
func (m TradeOrderBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 88, 84)
}

// SubAccountIdentifier
func (m TradeOrder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 92, 88)
}

// SubAccountIdentifier
func (m TradeOrderBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 92, 88)
}

// InternalOrderIDModifierForService
func (m TradeOrder) InternalOrderIDModifierForService() int32 {
	return message.Int32VLS(m.Unsafe(), 96, 92)
}

// InternalOrderIDModifierForService
func (m TradeOrderBuilder) InternalOrderIDModifierForService() int32 {
	return message.Int32VLS(m.Unsafe(), 96, 92)
}

// FIXClientOrderID
func (m TradeOrder) FIXClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 100, 96)
}

// FIXClientOrderID
func (m TradeOrderBuilder) FIXClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 100, 96)
}

// SequenceNumberBasedClientOrderID
func (m TradeOrder) SequenceNumberBasedClientOrderID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 104, 100)
}

// SequenceNumberBasedClientOrderID
func (m TradeOrderBuilder) SequenceNumberBasedClientOrderID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 104, 100)
}

// ClientOrderIDForDTCServer
func (m TradeOrder) ClientOrderIDForDTCServer() string {
	return message.StringVLS(m.Unsafe(), 108, 104)
}

// ClientOrderIDForDTCServer
func (m TradeOrderBuilder) ClientOrderIDForDTCServer() string {
	return message.StringVLS(m.Unsafe(), 108, 104)
}

// PreviousClientOrderIDForDTCServer
func (m TradeOrder) PreviousClientOrderIDForDTCServer() string {
	return message.StringVLS(m.Unsafe(), 112, 108)
}

// PreviousClientOrderIDForDTCServer
func (m TradeOrderBuilder) PreviousClientOrderIDForDTCServer() string {
	return message.StringVLS(m.Unsafe(), 112, 108)
}

// ExchangeOrderID
func (m TradeOrder) ExchangeOrderID() string {
	return message.StringVLS(m.Unsafe(), 116, 112)
}

// ExchangeOrderID
func (m TradeOrderBuilder) ExchangeOrderID() string {
	return message.StringVLS(m.Unsafe(), 116, 112)
}

// OriginatingClientUsername
func (m TradeOrder) OriginatingClientUsername() string {
	return message.StringVLS(m.Unsafe(), 120, 116)
}

// OriginatingClientUsername
func (m TradeOrderBuilder) OriginatingClientUsername() string {
	return message.StringVLS(m.Unsafe(), 120, 116)
}

// EntryDateTime
func (m TradeOrder) EntryDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 128, 120)
}

// EntryDateTime
func (m TradeOrderBuilder) EntryDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 128, 120)
}

// LastActionDateTime
func (m TradeOrder) LastActionDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 136, 128)
}

// LastActionDateTime
func (m TradeOrderBuilder) LastActionDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 136, 128)
}

// ServiceUpdateDateTimeUTC
func (m TradeOrder) ServiceUpdateDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 144, 136)
}

// ServiceUpdateDateTimeUTC
func (m TradeOrderBuilder) ServiceUpdateDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 144, 136)
}

// OrderEntryTimeForService
func (m TradeOrder) OrderEntryTimeForService() uint32 {
	return message.Uint32VLS(m.Unsafe(), 148, 144)
}

// OrderEntryTimeForService
func (m TradeOrderBuilder) OrderEntryTimeForService() uint32 {
	return message.Uint32VLS(m.Unsafe(), 148, 144)
}

// LastModifyTimeForService
func (m TradeOrder) LastModifyTimeForService() uint32 {
	return message.Uint32VLS(m.Unsafe(), 152, 148)
}

// LastModifyTimeForService
func (m TradeOrderBuilder) LastModifyTimeForService() uint32 {
	return message.Uint32VLS(m.Unsafe(), 152, 148)
}

// GoodTillDateTime
func (m TradeOrder) GoodTillDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 160, 152)
}

// GoodTillDateTime
func (m TradeOrderBuilder) GoodTillDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 160, 152)
}

// TimeInForce
func (m TradeOrder) TimeInForce() int32 {
	return message.Int32VLS(m.Unsafe(), 164, 160)
}

// TimeInForce
func (m TradeOrderBuilder) TimeInForce() int32 {
	return message.Int32VLS(m.Unsafe(), 164, 160)
}

// OpenClose
func (m TradeOrder) OpenClose() uint16 {
	return message.Uint16VLS(m.Unsafe(), 166, 164)
}

// OpenClose
func (m TradeOrderBuilder) OpenClose() uint16 {
	return message.Uint16VLS(m.Unsafe(), 166, 164)
}

// TrailStopOffset1
func (m TradeOrder) TrailStopOffset1() float64 {
	return message.Float64VLS(m.Unsafe(), 174, 166)
}

// TrailStopOffset1
func (m TradeOrderBuilder) TrailStopOffset1() float64 {
	return message.Float64VLS(m.Unsafe(), 174, 166)
}

// TrailStopStep
func (m TradeOrder) TrailStopStep() float64 {
	return message.Float64VLS(m.Unsafe(), 182, 174)
}

// TrailStopStep
func (m TradeOrderBuilder) TrailStopStep() float64 {
	return message.Float64VLS(m.Unsafe(), 182, 174)
}

// TrailTriggerPrice
func (m TradeOrder) TrailTriggerPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 190, 182)
}

// TrailTriggerPrice
func (m TradeOrderBuilder) TrailTriggerPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 190, 182)
}

// TrailingStopTriggerOffset
func (m TradeOrder) TrailingStopTriggerOffset() float64 {
	return message.Float64VLS(m.Unsafe(), 198, 190)
}

// TrailingStopTriggerOffset
func (m TradeOrderBuilder) TrailingStopTriggerOffset() float64 {
	return message.Float64VLS(m.Unsafe(), 198, 190)
}

// TrailTriggerHit
func (m TradeOrder) TrailTriggerHit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 199, 198)
}

// TrailTriggerHit
func (m TradeOrderBuilder) TrailTriggerHit() uint8 {
	return message.Uint8VLS(m.Unsafe(), 199, 198)
}

// TrailToBreakEvenStopOffset
func (m TradeOrder) TrailToBreakEvenStopOffset() float64 {
	return message.Float64VLS(m.Unsafe(), 207, 199)
}

// TrailToBreakEvenStopOffset
func (m TradeOrderBuilder) TrailToBreakEvenStopOffset() float64 {
	return message.Float64VLS(m.Unsafe(), 207, 199)
}

// MaximumChaseAmountAsPrice
func (m TradeOrder) MaximumChaseAmountAsPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 215, 207)
}

// MaximumChaseAmountAsPrice
func (m TradeOrderBuilder) MaximumChaseAmountAsPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 215, 207)
}

// InitialChaseOrderPrice1
func (m TradeOrder) InitialChaseOrderPrice1() float64 {
	return message.Float64VLS(m.Unsafe(), 223, 215)
}

// InitialChaseOrderPrice1
func (m TradeOrderBuilder) InitialChaseOrderPrice1() float64 {
	return message.Float64VLS(m.Unsafe(), 223, 215)
}

// InitialLastTradePriceForChaseOrders
func (m TradeOrder) InitialLastTradePriceForChaseOrders() float64 {
	return message.Float64VLS(m.Unsafe(), 231, 223)
}

// InitialLastTradePriceForChaseOrders
func (m TradeOrderBuilder) InitialLastTradePriceForChaseOrders() float64 {
	return message.Float64VLS(m.Unsafe(), 231, 223)
}

// TrailingStopTriggerOCOGroupNumber
func (m TradeOrder) TrailingStopTriggerOCOGroupNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 235, 231)
}

// TrailingStopTriggerOCOGroupNumber
func (m TradeOrderBuilder) TrailingStopTriggerOCOGroupNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 235, 231)
}

// LastModifyPrice1
func (m TradeOrder) LastModifyPrice1() float64 {
	return message.Float64VLS(m.Unsafe(), 243, 235)
}

// LastModifyPrice1
func (m TradeOrderBuilder) LastModifyPrice1() float64 {
	return message.Float64VLS(m.Unsafe(), 243, 235)
}

// LastModifyQuantity
func (m TradeOrder) LastModifyQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 251, 243)
}

// LastModifyQuantity
func (m TradeOrderBuilder) LastModifyQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 251, 243)
}

// CumulativeOrderQuantityFromParentFills
func (m TradeOrder) CumulativeOrderQuantityFromParentFills() float64 {
	return message.Float64VLS(m.Unsafe(), 259, 251)
}

// CumulativeOrderQuantityFromParentFills
func (m TradeOrderBuilder) CumulativeOrderQuantityFromParentFills() float64 {
	return message.Float64VLS(m.Unsafe(), 259, 251)
}

// PriorFilledQuantity
func (m TradeOrder) PriorFilledQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 267, 259)
}

// PriorFilledQuantity
func (m TradeOrderBuilder) PriorFilledQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 267, 259)
}

// TickSize
func (m TradeOrder) TickSize() float32 {
	return message.Float32VLS(m.Unsafe(), 271, 267)
}

// TickSize
func (m TradeOrderBuilder) TickSize() float32 {
	return message.Float32VLS(m.Unsafe(), 271, 267)
}

// ValueFormat
func (m TradeOrder) ValueFormat() int32 {
	return message.Int32VLS(m.Unsafe(), 275, 271)
}

// ValueFormat
func (m TradeOrderBuilder) ValueFormat() int32 {
	return message.Int32VLS(m.Unsafe(), 275, 271)
}

// PriceMultiplier
func (m TradeOrder) PriceMultiplier() float32 {
	return message.Float32VLS(m.Unsafe(), 279, 275)
}

// PriceMultiplier
func (m TradeOrderBuilder) PriceMultiplier() float32 {
	return message.Float32VLS(m.Unsafe(), 279, 275)
}

// ParentInternalOrderID
func (m TradeOrder) ParentInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 287, 279)
}

// ParentInternalOrderID
func (m TradeOrderBuilder) ParentInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 287, 279)
}

// TargetChildInternalOrderID
func (m TradeOrder) TargetChildInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 295, 287)
}

// TargetChildInternalOrderID
func (m TradeOrderBuilder) TargetChildInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 295, 287)
}

// StopChildInternalOrderID
func (m TradeOrder) StopChildInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 303, 295)
}

// StopChildInternalOrderID
func (m TradeOrderBuilder) StopChildInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 303, 295)
}

// AttachedOrderPriceOffset1
func (m TradeOrder) AttachedOrderPriceOffset1() float64 {
	return message.Float64VLS(m.Unsafe(), 311, 303)
}

// AttachedOrderPriceOffset1
func (m TradeOrderBuilder) AttachedOrderPriceOffset1() float64 {
	return message.Float64VLS(m.Unsafe(), 311, 303)
}

// LinkInternalOrderID
func (m TradeOrder) LinkInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 319, 311)
}

// LinkInternalOrderID
func (m TradeOrderBuilder) LinkInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 319, 311)
}

// OCOGroupInternalOrderID
func (m TradeOrder) OCOGroupInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 327, 319)
}

// OCOGroupInternalOrderID
func (m TradeOrderBuilder) OCOGroupInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 327, 319)
}

// OCOSiblingInternalOrderID
func (m TradeOrder) OCOSiblingInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 335, 327)
}

// OCOSiblingInternalOrderID
func (m TradeOrderBuilder) OCOSiblingInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 335, 327)
}

// DisableChildAndSiblingRelatedActions
func (m TradeOrder) DisableChildAndSiblingRelatedActions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 336, 335)
}

// DisableChildAndSiblingRelatedActions
func (m TradeOrderBuilder) DisableChildAndSiblingRelatedActions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 336, 335)
}

// OCOManagedByService
func (m TradeOrder) OCOManagedByService() uint8 {
	return message.Uint8VLS(m.Unsafe(), 337, 336)
}

// OCOManagedByService
func (m TradeOrderBuilder) OCOManagedByService() uint8 {
	return message.Uint8VLS(m.Unsafe(), 337, 336)
}

// BracketOrderServerManaged
func (m TradeOrder) BracketOrderServerManaged() uint8 {
	return message.Uint8VLS(m.Unsafe(), 338, 337)
}

// BracketOrderServerManaged
func (m TradeOrderBuilder) BracketOrderServerManaged() uint8 {
	return message.Uint8VLS(m.Unsafe(), 338, 337)
}

// LastOrderActionSource
func (m TradeOrder) LastOrderActionSource() string {
	return message.StringVLS(m.Unsafe(), 342, 338)
}

// LastOrderActionSource
func (m TradeOrderBuilder) LastOrderActionSource() string {
	return message.StringVLS(m.Unsafe(), 342, 338)
}

// StopLimitOrderStopPriceTriggered
func (m TradeOrder) StopLimitOrderStopPriceTriggered() uint8 {
	return message.Uint8VLS(m.Unsafe(), 343, 342)
}

// StopLimitOrderStopPriceTriggered
func (m TradeOrderBuilder) StopLimitOrderStopPriceTriggered() uint8 {
	return message.Uint8VLS(m.Unsafe(), 343, 342)
}

// OCOFullSiblingCancelOnPartialFill
func (m TradeOrder) OCOFullSiblingCancelOnPartialFill() uint8 {
	return message.Uint8VLS(m.Unsafe(), 344, 343)
}

// OCOFullSiblingCancelOnPartialFill
func (m TradeOrderBuilder) OCOFullSiblingCancelOnPartialFill() uint8 {
	return message.Uint8VLS(m.Unsafe(), 344, 343)
}

// ReverseOnCompleteFill
func (m TradeOrder) ReverseOnCompleteFill() uint8 {
	return message.Uint8VLS(m.Unsafe(), 345, 344)
}

// ReverseOnCompleteFill
func (m TradeOrderBuilder) ReverseOnCompleteFill() uint8 {
	return message.Uint8VLS(m.Unsafe(), 345, 344)
}

// SupportScaleIn
func (m TradeOrder) SupportScaleIn() uint8 {
	return message.Uint8VLS(m.Unsafe(), 346, 345)
}

// SupportScaleIn
func (m TradeOrderBuilder) SupportScaleIn() uint8 {
	return message.Uint8VLS(m.Unsafe(), 346, 345)
}

// SupportScaleOut
func (m TradeOrder) SupportScaleOut() uint8 {
	return message.Uint8VLS(m.Unsafe(), 347, 346)
}

// SupportScaleOut
func (m TradeOrderBuilder) SupportScaleOut() uint8 {
	return message.Uint8VLS(m.Unsafe(), 347, 346)
}

// SourceChartNumber
func (m TradeOrder) SourceChartNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 351, 347)
}

// SourceChartNumber
func (m TradeOrderBuilder) SourceChartNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 351, 347)
}

// SourceChartbookFileName
func (m TradeOrder) SourceChartbookFileName() string {
	return message.StringVLS(m.Unsafe(), 355, 351)
}

// SourceChartbookFileName
func (m TradeOrderBuilder) SourceChartbookFileName() string {
	return message.StringVLS(m.Unsafe(), 355, 351)
}

// IsAutomatedOrder
func (m TradeOrder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 356, 355)
}

// IsAutomatedOrder
func (m TradeOrderBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 356, 355)
}

// SimulatedOrder
func (m TradeOrder) SimulatedOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 357, 356)
}

// SimulatedOrder
func (m TradeOrderBuilder) SimulatedOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 357, 356)
}

// IsChartReplaying
func (m TradeOrder) IsChartReplaying() uint8 {
	return message.Uint8VLS(m.Unsafe(), 358, 357)
}

// IsChartReplaying
func (m TradeOrderBuilder) IsChartReplaying() uint8 {
	return message.Uint8VLS(m.Unsafe(), 358, 357)
}

// AttachedOrderOCOGroupNumber
func (m TradeOrder) AttachedOrderOCOGroupNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 362, 358)
}

// AttachedOrderOCOGroupNumber
func (m TradeOrderBuilder) AttachedOrderOCOGroupNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 362, 358)
}

// LastFillExecutionServiceID
func (m TradeOrder) LastFillExecutionServiceID() string {
	return message.StringVLS(m.Unsafe(), 366, 362)
}

// LastFillExecutionServiceID
func (m TradeOrderBuilder) LastFillExecutionServiceID() string {
	return message.StringVLS(m.Unsafe(), 366, 362)
}

// FillCount
func (m TradeOrder) FillCount() int32 {
	return message.Int32VLS(m.Unsafe(), 370, 366)
}

// FillCount
func (m TradeOrderBuilder) FillCount() int32 {
	return message.Int32VLS(m.Unsafe(), 370, 366)
}

// LastFillQuantity
func (m TradeOrder) LastFillQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 378, 370)
}

// LastFillQuantity
func (m TradeOrderBuilder) LastFillQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 378, 370)
}

// LastFillPrice
func (m TradeOrder) LastFillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 386, 378)
}

// LastFillPrice
func (m TradeOrderBuilder) LastFillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 386, 378)
}

// LastFillDateTimeUTC
func (m TradeOrder) LastFillDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 394, 386)
}

// LastFillDateTimeUTC
func (m TradeOrderBuilder) LastFillDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 394, 386)
}

// RejectedStopOCOSiblingInternalOrderID
func (m TradeOrder) RejectedStopOCOSiblingInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 402, 394)
}

// RejectedStopOCOSiblingInternalOrderID
func (m TradeOrderBuilder) RejectedStopOCOSiblingInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 402, 394)
}

// RejectedStopReplacementMarketOrderQuantity
func (m TradeOrder) RejectedStopReplacementMarketOrderQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 410, 402)
}

// RejectedStopReplacementMarketOrderQuantity
func (m TradeOrderBuilder) RejectedStopReplacementMarketOrderQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 410, 402)
}

// EvaluatingForFill
func (m TradeOrder) EvaluatingForFill() uint8 {
	return message.Uint8VLS(m.Unsafe(), 411, 410)
}

// EvaluatingForFill
func (m TradeOrderBuilder) EvaluatingForFill() uint8 {
	return message.Uint8VLS(m.Unsafe(), 411, 410)
}

// LastProcessedTimeSalesRecordSequenceForPrices
func (m TradeOrder) LastProcessedTimeSalesRecordSequenceForPrices() uint32 {
	return message.Uint32VLS(m.Unsafe(), 415, 411)
}

// LastProcessedTimeSalesRecordSequenceForPrices
func (m TradeOrderBuilder) LastProcessedTimeSalesRecordSequenceForPrices() uint32 {
	return message.Uint32VLS(m.Unsafe(), 415, 411)
}

// IsMarketDataManagementOfOrderEnabled
func (m TradeOrder) IsMarketDataManagementOfOrderEnabled() uint8 {
	return message.Uint8VLS(m.Unsafe(), 416, 415)
}

// IsMarketDataManagementOfOrderEnabled
func (m TradeOrderBuilder) IsMarketDataManagementOfOrderEnabled() uint8 {
	return message.Uint8VLS(m.Unsafe(), 416, 415)
}

// TextTag
func (m TradeOrder) TextTag() string {
	return message.StringVLS(m.Unsafe(), 420, 416)
}

// TextTag
func (m TradeOrderBuilder) TextTag() string {
	return message.StringVLS(m.Unsafe(), 420, 416)
}

// TimedOutOrderRequestedStatusDateTime
func (m TradeOrder) TimedOutOrderRequestedStatusDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 428, 420)
}

// TimedOutOrderRequestedStatusDateTime
func (m TradeOrderBuilder) TimedOutOrderRequestedStatusDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 428, 420)
}

// RequestedStatusForTimedOutOrder
func (m TradeOrder) RequestedStatusForTimedOutOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 429, 428)
}

// RequestedStatusForTimedOutOrder
func (m TradeOrderBuilder) RequestedStatusForTimedOutOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 429, 428)
}

// SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled
func (m TradeOrder) SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled() uint8 {
	return message.Uint8VLS(m.Unsafe(), 430, 429)
}

// SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled
func (m TradeOrderBuilder) SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled() uint8 {
	return message.Uint8VLS(m.Unsafe(), 430, 429)
}

// QuantityToIncreaseFromParentFill
func (m TradeOrder) QuantityToIncreaseFromParentFill() float64 {
	return message.Float64VLS(m.Unsafe(), 438, 430)
}

// QuantityToIncreaseFromParentFill
func (m TradeOrderBuilder) QuantityToIncreaseFromParentFill() float64 {
	return message.Float64VLS(m.Unsafe(), 438, 430)
}

// MoveToBreakevenStopReferencePrice
func (m TradeOrder) MoveToBreakevenStopReferencePrice() float64 {
	return message.Float64VLS(m.Unsafe(), 446, 438)
}

// MoveToBreakevenStopReferencePrice
func (m TradeOrderBuilder) MoveToBreakevenStopReferencePrice() float64 {
	return message.Float64VLS(m.Unsafe(), 446, 438)
}

// QuantityTriggeredStop_QuantityForTrigger
func (m TradeOrder) QuantityTriggeredStop_QuantityForTrigger() float64 {
	return message.Float64VLS(m.Unsafe(), 454, 446)
}

// QuantityTriggeredStop_QuantityForTrigger
func (m TradeOrderBuilder) QuantityTriggeredStop_QuantityForTrigger() float64 {
	return message.Float64VLS(m.Unsafe(), 454, 446)
}

// AccumulatedTradeVolumeForTriggeredStop
func (m TradeOrder) AccumulatedTradeVolumeForTriggeredStop() float64 {
	return message.Float64VLS(m.Unsafe(), 462, 454)
}

// AccumulatedTradeVolumeForTriggeredStop
func (m TradeOrderBuilder) AccumulatedTradeVolumeForTriggeredStop() float64 {
	return message.Float64VLS(m.Unsafe(), 462, 454)
}

// BidAskQuantityStopInitialTriggerMet
func (m TradeOrder) BidAskQuantityStopInitialTriggerMet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 463, 462)
}

// BidAskQuantityStopInitialTriggerMet
func (m TradeOrderBuilder) BidAskQuantityStopInitialTriggerMet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 463, 462)
}

// NeedToOverrideLock
func (m TradeOrder) NeedToOverrideLock() uint8 {
	return message.Uint8VLS(m.Unsafe(), 464, 463)
}

// NeedToOverrideLock
func (m TradeOrderBuilder) NeedToOverrideLock() uint8 {
	return message.Uint8VLS(m.Unsafe(), 464, 463)
}

// CurrentMarketPrice
func (m TradeOrder) CurrentMarketPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 472, 464)
}

// CurrentMarketPrice
func (m TradeOrderBuilder) CurrentMarketPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 472, 464)
}

// CurrentMarketDateTime
func (m TradeOrder) CurrentMarketDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 480, 472)
}

// CurrentMarketDateTime
func (m TradeOrderBuilder) CurrentMarketDateTime() int64 {
	return message.Int64VLS(m.Unsafe(), 480, 472)
}

// SupportOrderFillBilling
func (m TradeOrder) SupportOrderFillBilling() uint8 {
	return message.Uint8VLS(m.Unsafe(), 481, 480)
}

// SupportOrderFillBilling
func (m TradeOrderBuilder) SupportOrderFillBilling() uint8 {
	return message.Uint8VLS(m.Unsafe(), 481, 480)
}

// IsBillable
func (m TradeOrder) IsBillable() uint8 {
	return message.Uint8VLS(m.Unsafe(), 482, 481)
}

// IsBillable
func (m TradeOrderBuilder) IsBillable() uint8 {
	return message.Uint8VLS(m.Unsafe(), 482, 481)
}

// QuantityForBilling
func (m TradeOrder) QuantityForBilling() int32 {
	return message.Int32VLS(m.Unsafe(), 486, 482)
}

// QuantityForBilling
func (m TradeOrderBuilder) QuantityForBilling() int32 {
	return message.Int32VLS(m.Unsafe(), 486, 482)
}

// NumberOfFailedOrderModifications
func (m TradeOrder) NumberOfFailedOrderModifications() uint32 {
	return message.Uint32VLS(m.Unsafe(), 490, 486)
}

// NumberOfFailedOrderModifications
func (m TradeOrderBuilder) NumberOfFailedOrderModifications() uint32 {
	return message.Uint32VLS(m.Unsafe(), 490, 486)
}

// DTCServerIndex
func (m TradeOrder) DTCServerIndex() int32 {
	return message.Int32VLS(m.Unsafe(), 494, 490)
}

// DTCServerIndex
func (m TradeOrderBuilder) DTCServerIndex() int32 {
	return message.Int32VLS(m.Unsafe(), 494, 490)
}

// ClearingFirmID
func (m TradeOrder) ClearingFirmID() string {
	return message.StringVLS(m.Unsafe(), 498, 494)
}

// ClearingFirmID
func (m TradeOrderBuilder) ClearingFirmID() string {
	return message.StringVLS(m.Unsafe(), 498, 494)
}

// SenderSubID
func (m TradeOrder) SenderSubID() string {
	return message.StringVLS(m.Unsafe(), 502, 498)
}

// SenderSubID
func (m TradeOrderBuilder) SenderSubID() string {
	return message.StringVLS(m.Unsafe(), 502, 498)
}

// SenderLocationId
func (m TradeOrder) SenderLocationId() string {
	return message.StringVLS(m.Unsafe(), 506, 502)
}

// SenderLocationId
func (m TradeOrderBuilder) SenderLocationId() string {
	return message.StringVLS(m.Unsafe(), 506, 502)
}

// SelfMatchPreventionID
func (m TradeOrder) SelfMatchPreventionID() string {
	return message.StringVLS(m.Unsafe(), 510, 506)
}

// SelfMatchPreventionID
func (m TradeOrderBuilder) SelfMatchPreventionID() string {
	return message.StringVLS(m.Unsafe(), 510, 506)
}

// CTICode
func (m TradeOrder) CTICode() int32 {
	return message.Int32VLS(m.Unsafe(), 514, 510)
}

// CTICode
func (m TradeOrderBuilder) CTICode() int32 {
	return message.Int32VLS(m.Unsafe(), 514, 510)
}

// ObtainOrderActionDateTimeFromLastTradeTimeInChart
func (m TradeOrder) ObtainOrderActionDateTimeFromLastTradeTimeInChart() uint8 {
	return message.Uint8VLS(m.Unsafe(), 515, 514)
}

// ObtainOrderActionDateTimeFromLastTradeTimeInChart
func (m TradeOrderBuilder) ObtainOrderActionDateTimeFromLastTradeTimeInChart() uint8 {
	return message.Uint8VLS(m.Unsafe(), 515, 514)
}

// MaximumShowQuantity
func (m TradeOrder) MaximumShowQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 523, 515)
}

// MaximumShowQuantity
func (m TradeOrderBuilder) MaximumShowQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 523, 515)
}

// OrderSubmitted
func (m TradeOrder) OrderSubmitted() uint8 {
	return message.Uint8VLS(m.Unsafe(), 524, 523)
}

// OrderSubmitted
func (m TradeOrderBuilder) OrderSubmitted() uint8 {
	return message.Uint8VLS(m.Unsafe(), 524, 523)
}

// IsSnapshot
func (m TradeOrder) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Unsafe(), 525, 524)
}

// IsSnapshot
func (m TradeOrderBuilder) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Unsafe(), 525, 524)
}

// IsFirstMessageInBatch
func (m TradeOrder) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 526, 525)
}

// IsFirstMessageInBatch
func (m TradeOrderBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 526, 525)
}

// IsLastMessageInBatch
func (m TradeOrder) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 527, 526)
}

// IsLastMessageInBatch
func (m TradeOrderBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 527, 526)
}

// ExternalLastActionDateTimeUTC
func (m TradeOrder) ExternalLastActionDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 535, 527)
}

// ExternalLastActionDateTimeUTC
func (m TradeOrderBuilder) ExternalLastActionDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 535, 527)
}

// SetIsOrderDeleted
func (m TradeOrderBuilder) SetIsOrderDeleted(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 7, 6, value)
}

// SetInternalOrderID
func (m TradeOrderBuilder) SetInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 15, 7, value)
}

// SetOrderStatusCode
func (m TradeOrderBuilder) SetOrderStatusCode(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 17, 15, value)
}

// SetOrderStatusBeforePendingModify
func (m TradeOrderBuilder) SetOrderStatusBeforePendingModify(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 19, 17, value)
}

// SetOrderStatusBeforePendingCancel
func (m TradeOrderBuilder) SetOrderStatusBeforePendingCancel(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 21, 19, value)
}

// SetServiceOrderID
func (m *TradeOrderBuilder) SetServiceOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 25, 21, value)
}

// SetActualSymbol
func (m *TradeOrderBuilder) SetActualSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 29, 25, value)
}

// SetOrderType
func (m TradeOrderBuilder) SetOrderType(value int32) {
	message.SetInt32VLS(m.Unsafe(), 33, 29, value)
}

// SetBuySell
func (m TradeOrderBuilder) SetBuySell(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 35, 33, value)
}

// SetPrice1
func (m TradeOrderBuilder) SetPrice1(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 43, 35, value)
}

// SetPrice2
func (m TradeOrderBuilder) SetPrice2(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 51, 43, value)
}

// SetOrderQuantity
func (m TradeOrderBuilder) SetOrderQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 59, 51, value)
}

// SetFilledQuantity
func (m TradeOrderBuilder) SetFilledQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 67, 59, value)
}

// SetAverageFillPrice
func (m TradeOrderBuilder) SetAverageFillPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 75, 67, value)
}

// SetRealtimeFillStatus
func (m TradeOrderBuilder) SetRealtimeFillStatus(value int32) {
	message.SetInt32VLS(m.Unsafe(), 79, 75, value)
}

// SetIsRestingOrderDuringFill
func (m TradeOrderBuilder) SetIsRestingOrderDuringFill(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 80, 79, value)
}

// SetOrderRejectType
func (m TradeOrderBuilder) SetOrderRejectType(value int32) {
	message.SetInt32VLS(m.Unsafe(), 84, 80, value)
}

// SetTradeAccount
func (m *TradeOrderBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 88, 84, value)
}

// SetSubAccountIdentifier
func (m TradeOrderBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 92, 88, value)
}

// SetInternalOrderIDModifierForService
func (m TradeOrderBuilder) SetInternalOrderIDModifierForService(value int32) {
	message.SetInt32VLS(m.Unsafe(), 96, 92, value)
}

// SetFIXClientOrderID
func (m *TradeOrderBuilder) SetFIXClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 100, 96, value)
}

// SetSequenceNumberBasedClientOrderID
func (m TradeOrderBuilder) SetSequenceNumberBasedClientOrderID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 104, 100, value)
}

// SetClientOrderIDForDTCServer
func (m *TradeOrderBuilder) SetClientOrderIDForDTCServer(value string) {
	message.SetStringVLS(&m.VLSBuilder, 108, 104, value)
}

// SetPreviousClientOrderIDForDTCServer
func (m *TradeOrderBuilder) SetPreviousClientOrderIDForDTCServer(value string) {
	message.SetStringVLS(&m.VLSBuilder, 112, 108, value)
}

// SetExchangeOrderID
func (m *TradeOrderBuilder) SetExchangeOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 116, 112, value)
}

// SetOriginatingClientUsername
func (m *TradeOrderBuilder) SetOriginatingClientUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 120, 116, value)
}

// SetEntryDateTime
func (m TradeOrderBuilder) SetEntryDateTime(value int64) {
	message.SetInt64VLS(m.Unsafe(), 128, 120, value)
}

// SetLastActionDateTime
func (m TradeOrderBuilder) SetLastActionDateTime(value int64) {
	message.SetInt64VLS(m.Unsafe(), 136, 128, value)
}

// SetServiceUpdateDateTimeUTC
func (m TradeOrderBuilder) SetServiceUpdateDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Unsafe(), 144, 136, value)
}

// SetOrderEntryTimeForService
func (m TradeOrderBuilder) SetOrderEntryTimeForService(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 148, 144, value)
}

// SetLastModifyTimeForService
func (m TradeOrderBuilder) SetLastModifyTimeForService(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 152, 148, value)
}

// SetGoodTillDateTime
func (m TradeOrderBuilder) SetGoodTillDateTime(value int64) {
	message.SetInt64VLS(m.Unsafe(), 160, 152, value)
}

// SetTimeInForce
func (m TradeOrderBuilder) SetTimeInForce(value int32) {
	message.SetInt32VLS(m.Unsafe(), 164, 160, value)
}

// SetOpenClose
func (m TradeOrderBuilder) SetOpenClose(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 166, 164, value)
}

// SetTrailStopOffset1
func (m TradeOrderBuilder) SetTrailStopOffset1(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 174, 166, value)
}

// SetTrailStopStep
func (m TradeOrderBuilder) SetTrailStopStep(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 182, 174, value)
}

// SetTrailTriggerPrice
func (m TradeOrderBuilder) SetTrailTriggerPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 190, 182, value)
}

// SetTrailingStopTriggerOffset
func (m TradeOrderBuilder) SetTrailingStopTriggerOffset(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 198, 190, value)
}

// SetTrailTriggerHit
func (m TradeOrderBuilder) SetTrailTriggerHit(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 199, 198, value)
}

// SetTrailToBreakEvenStopOffset
func (m TradeOrderBuilder) SetTrailToBreakEvenStopOffset(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 207, 199, value)
}

// SetMaximumChaseAmountAsPrice
func (m TradeOrderBuilder) SetMaximumChaseAmountAsPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 215, 207, value)
}

// SetInitialChaseOrderPrice1
func (m TradeOrderBuilder) SetInitialChaseOrderPrice1(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 223, 215, value)
}

// SetInitialLastTradePriceForChaseOrders
func (m TradeOrderBuilder) SetInitialLastTradePriceForChaseOrders(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 231, 223, value)
}

// SetTrailingStopTriggerOCOGroupNumber
func (m TradeOrderBuilder) SetTrailingStopTriggerOCOGroupNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 235, 231, value)
}

// SetLastModifyPrice1
func (m TradeOrderBuilder) SetLastModifyPrice1(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 243, 235, value)
}

// SetLastModifyQuantity
func (m TradeOrderBuilder) SetLastModifyQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 251, 243, value)
}

// SetCumulativeOrderQuantityFromParentFills
func (m TradeOrderBuilder) SetCumulativeOrderQuantityFromParentFills(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 259, 251, value)
}

// SetPriorFilledQuantity
func (m TradeOrderBuilder) SetPriorFilledQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 267, 259, value)
}

// SetTickSize
func (m TradeOrderBuilder) SetTickSize(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 271, 267, value)
}

// SetValueFormat
func (m TradeOrderBuilder) SetValueFormat(value int32) {
	message.SetInt32VLS(m.Unsafe(), 275, 271, value)
}

// SetPriceMultiplier
func (m TradeOrderBuilder) SetPriceMultiplier(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 279, 275, value)
}

// SetParentInternalOrderID
func (m TradeOrderBuilder) SetParentInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 287, 279, value)
}

// SetTargetChildInternalOrderID
func (m TradeOrderBuilder) SetTargetChildInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 295, 287, value)
}

// SetStopChildInternalOrderID
func (m TradeOrderBuilder) SetStopChildInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 303, 295, value)
}

// SetAttachedOrderPriceOffset1
func (m TradeOrderBuilder) SetAttachedOrderPriceOffset1(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 311, 303, value)
}

// SetLinkInternalOrderID
func (m TradeOrderBuilder) SetLinkInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 319, 311, value)
}

// SetOCOGroupInternalOrderID
func (m TradeOrderBuilder) SetOCOGroupInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 327, 319, value)
}

// SetOCOSiblingInternalOrderID
func (m TradeOrderBuilder) SetOCOSiblingInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 335, 327, value)
}

// SetDisableChildAndSiblingRelatedActions
func (m TradeOrderBuilder) SetDisableChildAndSiblingRelatedActions(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 336, 335, value)
}

// SetOCOManagedByService
func (m TradeOrderBuilder) SetOCOManagedByService(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 337, 336, value)
}

// SetBracketOrderServerManaged
func (m TradeOrderBuilder) SetBracketOrderServerManaged(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 338, 337, value)
}

// SetLastOrderActionSource
func (m *TradeOrderBuilder) SetLastOrderActionSource(value string) {
	message.SetStringVLS(&m.VLSBuilder, 342, 338, value)
}

// SetStopLimitOrderStopPriceTriggered
func (m TradeOrderBuilder) SetStopLimitOrderStopPriceTriggered(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 343, 342, value)
}

// SetOCOFullSiblingCancelOnPartialFill
func (m TradeOrderBuilder) SetOCOFullSiblingCancelOnPartialFill(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 344, 343, value)
}

// SetReverseOnCompleteFill
func (m TradeOrderBuilder) SetReverseOnCompleteFill(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 345, 344, value)
}

// SetSupportScaleIn
func (m TradeOrderBuilder) SetSupportScaleIn(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 346, 345, value)
}

// SetSupportScaleOut
func (m TradeOrderBuilder) SetSupportScaleOut(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 347, 346, value)
}

// SetSourceChartNumber
func (m TradeOrderBuilder) SetSourceChartNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 351, 347, value)
}

// SetSourceChartbookFileName
func (m *TradeOrderBuilder) SetSourceChartbookFileName(value string) {
	message.SetStringVLS(&m.VLSBuilder, 355, 351, value)
}

// SetIsAutomatedOrder
func (m TradeOrderBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 356, 355, value)
}

// SetSimulatedOrder
func (m TradeOrderBuilder) SetSimulatedOrder(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 357, 356, value)
}

// SetIsChartReplaying
func (m TradeOrderBuilder) SetIsChartReplaying(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 358, 357, value)
}

// SetAttachedOrderOCOGroupNumber
func (m TradeOrderBuilder) SetAttachedOrderOCOGroupNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 362, 358, value)
}

// SetLastFillExecutionServiceID
func (m *TradeOrderBuilder) SetLastFillExecutionServiceID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 366, 362, value)
}

// SetFillCount
func (m TradeOrderBuilder) SetFillCount(value int32) {
	message.SetInt32VLS(m.Unsafe(), 370, 366, value)
}

// SetLastFillQuantity
func (m TradeOrderBuilder) SetLastFillQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 378, 370, value)
}

// SetLastFillPrice
func (m TradeOrderBuilder) SetLastFillPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 386, 378, value)
}

// SetLastFillDateTimeUTC
func (m TradeOrderBuilder) SetLastFillDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Unsafe(), 394, 386, value)
}

// SetRejectedStopOCOSiblingInternalOrderID
func (m TradeOrderBuilder) SetRejectedStopOCOSiblingInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 402, 394, value)
}

// SetRejectedStopReplacementMarketOrderQuantity
func (m TradeOrderBuilder) SetRejectedStopReplacementMarketOrderQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 410, 402, value)
}

// SetEvaluatingForFill
func (m TradeOrderBuilder) SetEvaluatingForFill(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 411, 410, value)
}

// SetLastProcessedTimeSalesRecordSequenceForPrices
func (m TradeOrderBuilder) SetLastProcessedTimeSalesRecordSequenceForPrices(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 415, 411, value)
}

// SetIsMarketDataManagementOfOrderEnabled
func (m TradeOrderBuilder) SetIsMarketDataManagementOfOrderEnabled(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 416, 415, value)
}

// SetTextTag
func (m *TradeOrderBuilder) SetTextTag(value string) {
	message.SetStringVLS(&m.VLSBuilder, 420, 416, value)
}

// SetTimedOutOrderRequestedStatusDateTime
func (m TradeOrderBuilder) SetTimedOutOrderRequestedStatusDateTime(value int64) {
	message.SetInt64VLS(m.Unsafe(), 428, 420, value)
}

// SetRequestedStatusForTimedOutOrder
func (m TradeOrderBuilder) SetRequestedStatusForTimedOutOrder(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 429, 428, value)
}

// SetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled
func (m TradeOrderBuilder) SetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 430, 429, value)
}

// SetQuantityToIncreaseFromParentFill
func (m TradeOrderBuilder) SetQuantityToIncreaseFromParentFill(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 438, 430, value)
}

// SetMoveToBreakevenStopReferencePrice
func (m TradeOrderBuilder) SetMoveToBreakevenStopReferencePrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 446, 438, value)
}

// SetQuantityTriggeredStop_QuantityForTrigger
func (m TradeOrderBuilder) SetQuantityTriggeredStop_QuantityForTrigger(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 454, 446, value)
}

// SetAccumulatedTradeVolumeForTriggeredStop
func (m TradeOrderBuilder) SetAccumulatedTradeVolumeForTriggeredStop(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 462, 454, value)
}

// SetBidAskQuantityStopInitialTriggerMet
func (m TradeOrderBuilder) SetBidAskQuantityStopInitialTriggerMet(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 463, 462, value)
}

// SetNeedToOverrideLock
func (m TradeOrderBuilder) SetNeedToOverrideLock(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 464, 463, value)
}

// SetCurrentMarketPrice
func (m TradeOrderBuilder) SetCurrentMarketPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 472, 464, value)
}

// SetCurrentMarketDateTime
func (m TradeOrderBuilder) SetCurrentMarketDateTime(value int64) {
	message.SetInt64VLS(m.Unsafe(), 480, 472, value)
}

// SetSupportOrderFillBilling
func (m TradeOrderBuilder) SetSupportOrderFillBilling(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 481, 480, value)
}

// SetIsBillable
func (m TradeOrderBuilder) SetIsBillable(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 482, 481, value)
}

// SetQuantityForBilling
func (m TradeOrderBuilder) SetQuantityForBilling(value int32) {
	message.SetInt32VLS(m.Unsafe(), 486, 482, value)
}

// SetNumberOfFailedOrderModifications
func (m TradeOrderBuilder) SetNumberOfFailedOrderModifications(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 490, 486, value)
}

// SetDTCServerIndex
func (m TradeOrderBuilder) SetDTCServerIndex(value int32) {
	message.SetInt32VLS(m.Unsafe(), 494, 490, value)
}

// SetClearingFirmID
func (m *TradeOrderBuilder) SetClearingFirmID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 498, 494, value)
}

// SetSenderSubID
func (m *TradeOrderBuilder) SetSenderSubID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 502, 498, value)
}

// SetSenderLocationId
func (m *TradeOrderBuilder) SetSenderLocationId(value string) {
	message.SetStringVLS(&m.VLSBuilder, 506, 502, value)
}

// SetSelfMatchPreventionID
func (m *TradeOrderBuilder) SetSelfMatchPreventionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 510, 506, value)
}

// SetCTICode
func (m TradeOrderBuilder) SetCTICode(value int32) {
	message.SetInt32VLS(m.Unsafe(), 514, 510, value)
}

// SetObtainOrderActionDateTimeFromLastTradeTimeInChart
func (m TradeOrderBuilder) SetObtainOrderActionDateTimeFromLastTradeTimeInChart(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 515, 514, value)
}

// SetMaximumShowQuantity
func (m TradeOrderBuilder) SetMaximumShowQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 523, 515, value)
}

// SetOrderSubmitted
func (m TradeOrderBuilder) SetOrderSubmitted(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 524, 523, value)
}

// SetIsSnapshot
func (m TradeOrderBuilder) SetIsSnapshot(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 525, 524, value)
}

// SetIsFirstMessageInBatch
func (m TradeOrderBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 526, 525, value)
}

// SetIsLastMessageInBatch
func (m TradeOrderBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 527, 526, value)
}

// SetExternalLastActionDateTimeUTC
func (m TradeOrderBuilder) SetExternalLastActionDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Unsafe(), 535, 527, value)
}

// Copy
func (m TradeOrder) Copy(to TradeOrderBuilder) {
	to.SetIsOrderDeleted(m.IsOrderDeleted())
	to.SetInternalOrderID(m.InternalOrderID())
	to.SetOrderStatusCode(m.OrderStatusCode())
	to.SetOrderStatusBeforePendingModify(m.OrderStatusBeforePendingModify())
	to.SetOrderStatusBeforePendingCancel(m.OrderStatusBeforePendingCancel())
	to.SetServiceOrderID(m.ServiceOrderID())
	to.SetActualSymbol(m.ActualSymbol())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetRealtimeFillStatus(m.RealtimeFillStatus())
	to.SetIsRestingOrderDuringFill(m.IsRestingOrderDuringFill())
	to.SetOrderRejectType(m.OrderRejectType())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetInternalOrderIDModifierForService(m.InternalOrderIDModifierForService())
	to.SetFIXClientOrderID(m.FIXClientOrderID())
	to.SetSequenceNumberBasedClientOrderID(m.SequenceNumberBasedClientOrderID())
	to.SetClientOrderIDForDTCServer(m.ClientOrderIDForDTCServer())
	to.SetPreviousClientOrderIDForDTCServer(m.PreviousClientOrderIDForDTCServer())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOriginatingClientUsername(m.OriginatingClientUsername())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetLastActionDateTime(m.LastActionDateTime())
	to.SetServiceUpdateDateTimeUTC(m.ServiceUpdateDateTimeUTC())
	to.SetOrderEntryTimeForService(m.OrderEntryTimeForService())
	to.SetLastModifyTimeForService(m.LastModifyTimeForService())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetTimeInForce(m.TimeInForce())
	to.SetOpenClose(m.OpenClose())
	to.SetTrailStopOffset1(m.TrailStopOffset1())
	to.SetTrailStopStep(m.TrailStopStep())
	to.SetTrailTriggerPrice(m.TrailTriggerPrice())
	to.SetTrailingStopTriggerOffset(m.TrailingStopTriggerOffset())
	to.SetTrailTriggerHit(m.TrailTriggerHit())
	to.SetTrailToBreakEvenStopOffset(m.TrailToBreakEvenStopOffset())
	to.SetMaximumChaseAmountAsPrice(m.MaximumChaseAmountAsPrice())
	to.SetInitialChaseOrderPrice1(m.InitialChaseOrderPrice1())
	to.SetInitialLastTradePriceForChaseOrders(m.InitialLastTradePriceForChaseOrders())
	to.SetTrailingStopTriggerOCOGroupNumber(m.TrailingStopTriggerOCOGroupNumber())
	to.SetLastModifyPrice1(m.LastModifyPrice1())
	to.SetLastModifyQuantity(m.LastModifyQuantity())
	to.SetCumulativeOrderQuantityFromParentFills(m.CumulativeOrderQuantityFromParentFills())
	to.SetPriorFilledQuantity(m.PriorFilledQuantity())
	to.SetTickSize(m.TickSize())
	to.SetValueFormat(m.ValueFormat())
	to.SetPriceMultiplier(m.PriceMultiplier())
	to.SetParentInternalOrderID(m.ParentInternalOrderID())
	to.SetTargetChildInternalOrderID(m.TargetChildInternalOrderID())
	to.SetStopChildInternalOrderID(m.StopChildInternalOrderID())
	to.SetAttachedOrderPriceOffset1(m.AttachedOrderPriceOffset1())
	to.SetLinkInternalOrderID(m.LinkInternalOrderID())
	to.SetOCOGroupInternalOrderID(m.OCOGroupInternalOrderID())
	to.SetOCOSiblingInternalOrderID(m.OCOSiblingInternalOrderID())
	to.SetDisableChildAndSiblingRelatedActions(m.DisableChildAndSiblingRelatedActions())
	to.SetOCOManagedByService(m.OCOManagedByService())
	to.SetBracketOrderServerManaged(m.BracketOrderServerManaged())
	to.SetLastOrderActionSource(m.LastOrderActionSource())
	to.SetStopLimitOrderStopPriceTriggered(m.StopLimitOrderStopPriceTriggered())
	to.SetOCOFullSiblingCancelOnPartialFill(m.OCOFullSiblingCancelOnPartialFill())
	to.SetReverseOnCompleteFill(m.ReverseOnCompleteFill())
	to.SetSupportScaleIn(m.SupportScaleIn())
	to.SetSupportScaleOut(m.SupportScaleOut())
	to.SetSourceChartNumber(m.SourceChartNumber())
	to.SetSourceChartbookFileName(m.SourceChartbookFileName())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetSimulatedOrder(m.SimulatedOrder())
	to.SetIsChartReplaying(m.IsChartReplaying())
	to.SetAttachedOrderOCOGroupNumber(m.AttachedOrderOCOGroupNumber())
	to.SetLastFillExecutionServiceID(m.LastFillExecutionServiceID())
	to.SetFillCount(m.FillCount())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTimeUTC(m.LastFillDateTimeUTC())
	to.SetRejectedStopOCOSiblingInternalOrderID(m.RejectedStopOCOSiblingInternalOrderID())
	to.SetRejectedStopReplacementMarketOrderQuantity(m.RejectedStopReplacementMarketOrderQuantity())
	to.SetEvaluatingForFill(m.EvaluatingForFill())
	to.SetLastProcessedTimeSalesRecordSequenceForPrices(m.LastProcessedTimeSalesRecordSequenceForPrices())
	to.SetIsMarketDataManagementOfOrderEnabled(m.IsMarketDataManagementOfOrderEnabled())
	to.SetTextTag(m.TextTag())
	to.SetTimedOutOrderRequestedStatusDateTime(m.TimedOutOrderRequestedStatusDateTime())
	to.SetRequestedStatusForTimedOutOrder(m.RequestedStatusForTimedOutOrder())
	to.SetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled(m.SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled())
	to.SetQuantityToIncreaseFromParentFill(m.QuantityToIncreaseFromParentFill())
	to.SetMoveToBreakevenStopReferencePrice(m.MoveToBreakevenStopReferencePrice())
	to.SetQuantityTriggeredStop_QuantityForTrigger(m.QuantityTriggeredStop_QuantityForTrigger())
	to.SetAccumulatedTradeVolumeForTriggeredStop(m.AccumulatedTradeVolumeForTriggeredStop())
	to.SetBidAskQuantityStopInitialTriggerMet(m.BidAskQuantityStopInitialTriggerMet())
	to.SetNeedToOverrideLock(m.NeedToOverrideLock())
	to.SetCurrentMarketPrice(m.CurrentMarketPrice())
	to.SetCurrentMarketDateTime(m.CurrentMarketDateTime())
	to.SetSupportOrderFillBilling(m.SupportOrderFillBilling())
	to.SetIsBillable(m.IsBillable())
	to.SetQuantityForBilling(m.QuantityForBilling())
	to.SetNumberOfFailedOrderModifications(m.NumberOfFailedOrderModifications())
	to.SetDTCServerIndex(m.DTCServerIndex())
	to.SetClearingFirmID(m.ClearingFirmID())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICode(m.CTICode())
	to.SetObtainOrderActionDateTimeFromLastTradeTimeInChart(m.ObtainOrderActionDateTimeFromLastTradeTimeInChart())
	to.SetMaximumShowQuantity(m.MaximumShowQuantity())
	to.SetOrderSubmitted(m.OrderSubmitted())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetExternalLastActionDateTimeUTC(m.ExternalLastActionDateTimeUTC())
}

// Copy
func (m TradeOrderBuilder) Copy(to TradeOrderBuilder) {
	to.SetIsOrderDeleted(m.IsOrderDeleted())
	to.SetInternalOrderID(m.InternalOrderID())
	to.SetOrderStatusCode(m.OrderStatusCode())
	to.SetOrderStatusBeforePendingModify(m.OrderStatusBeforePendingModify())
	to.SetOrderStatusBeforePendingCancel(m.OrderStatusBeforePendingCancel())
	to.SetServiceOrderID(m.ServiceOrderID())
	to.SetActualSymbol(m.ActualSymbol())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetRealtimeFillStatus(m.RealtimeFillStatus())
	to.SetIsRestingOrderDuringFill(m.IsRestingOrderDuringFill())
	to.SetOrderRejectType(m.OrderRejectType())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetInternalOrderIDModifierForService(m.InternalOrderIDModifierForService())
	to.SetFIXClientOrderID(m.FIXClientOrderID())
	to.SetSequenceNumberBasedClientOrderID(m.SequenceNumberBasedClientOrderID())
	to.SetClientOrderIDForDTCServer(m.ClientOrderIDForDTCServer())
	to.SetPreviousClientOrderIDForDTCServer(m.PreviousClientOrderIDForDTCServer())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOriginatingClientUsername(m.OriginatingClientUsername())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetLastActionDateTime(m.LastActionDateTime())
	to.SetServiceUpdateDateTimeUTC(m.ServiceUpdateDateTimeUTC())
	to.SetOrderEntryTimeForService(m.OrderEntryTimeForService())
	to.SetLastModifyTimeForService(m.LastModifyTimeForService())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetTimeInForce(m.TimeInForce())
	to.SetOpenClose(m.OpenClose())
	to.SetTrailStopOffset1(m.TrailStopOffset1())
	to.SetTrailStopStep(m.TrailStopStep())
	to.SetTrailTriggerPrice(m.TrailTriggerPrice())
	to.SetTrailingStopTriggerOffset(m.TrailingStopTriggerOffset())
	to.SetTrailTriggerHit(m.TrailTriggerHit())
	to.SetTrailToBreakEvenStopOffset(m.TrailToBreakEvenStopOffset())
	to.SetMaximumChaseAmountAsPrice(m.MaximumChaseAmountAsPrice())
	to.SetInitialChaseOrderPrice1(m.InitialChaseOrderPrice1())
	to.SetInitialLastTradePriceForChaseOrders(m.InitialLastTradePriceForChaseOrders())
	to.SetTrailingStopTriggerOCOGroupNumber(m.TrailingStopTriggerOCOGroupNumber())
	to.SetLastModifyPrice1(m.LastModifyPrice1())
	to.SetLastModifyQuantity(m.LastModifyQuantity())
	to.SetCumulativeOrderQuantityFromParentFills(m.CumulativeOrderQuantityFromParentFills())
	to.SetPriorFilledQuantity(m.PriorFilledQuantity())
	to.SetTickSize(m.TickSize())
	to.SetValueFormat(m.ValueFormat())
	to.SetPriceMultiplier(m.PriceMultiplier())
	to.SetParentInternalOrderID(m.ParentInternalOrderID())
	to.SetTargetChildInternalOrderID(m.TargetChildInternalOrderID())
	to.SetStopChildInternalOrderID(m.StopChildInternalOrderID())
	to.SetAttachedOrderPriceOffset1(m.AttachedOrderPriceOffset1())
	to.SetLinkInternalOrderID(m.LinkInternalOrderID())
	to.SetOCOGroupInternalOrderID(m.OCOGroupInternalOrderID())
	to.SetOCOSiblingInternalOrderID(m.OCOSiblingInternalOrderID())
	to.SetDisableChildAndSiblingRelatedActions(m.DisableChildAndSiblingRelatedActions())
	to.SetOCOManagedByService(m.OCOManagedByService())
	to.SetBracketOrderServerManaged(m.BracketOrderServerManaged())
	to.SetLastOrderActionSource(m.LastOrderActionSource())
	to.SetStopLimitOrderStopPriceTriggered(m.StopLimitOrderStopPriceTriggered())
	to.SetOCOFullSiblingCancelOnPartialFill(m.OCOFullSiblingCancelOnPartialFill())
	to.SetReverseOnCompleteFill(m.ReverseOnCompleteFill())
	to.SetSupportScaleIn(m.SupportScaleIn())
	to.SetSupportScaleOut(m.SupportScaleOut())
	to.SetSourceChartNumber(m.SourceChartNumber())
	to.SetSourceChartbookFileName(m.SourceChartbookFileName())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetSimulatedOrder(m.SimulatedOrder())
	to.SetIsChartReplaying(m.IsChartReplaying())
	to.SetAttachedOrderOCOGroupNumber(m.AttachedOrderOCOGroupNumber())
	to.SetLastFillExecutionServiceID(m.LastFillExecutionServiceID())
	to.SetFillCount(m.FillCount())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTimeUTC(m.LastFillDateTimeUTC())
	to.SetRejectedStopOCOSiblingInternalOrderID(m.RejectedStopOCOSiblingInternalOrderID())
	to.SetRejectedStopReplacementMarketOrderQuantity(m.RejectedStopReplacementMarketOrderQuantity())
	to.SetEvaluatingForFill(m.EvaluatingForFill())
	to.SetLastProcessedTimeSalesRecordSequenceForPrices(m.LastProcessedTimeSalesRecordSequenceForPrices())
	to.SetIsMarketDataManagementOfOrderEnabled(m.IsMarketDataManagementOfOrderEnabled())
	to.SetTextTag(m.TextTag())
	to.SetTimedOutOrderRequestedStatusDateTime(m.TimedOutOrderRequestedStatusDateTime())
	to.SetRequestedStatusForTimedOutOrder(m.RequestedStatusForTimedOutOrder())
	to.SetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled(m.SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled())
	to.SetQuantityToIncreaseFromParentFill(m.QuantityToIncreaseFromParentFill())
	to.SetMoveToBreakevenStopReferencePrice(m.MoveToBreakevenStopReferencePrice())
	to.SetQuantityTriggeredStop_QuantityForTrigger(m.QuantityTriggeredStop_QuantityForTrigger())
	to.SetAccumulatedTradeVolumeForTriggeredStop(m.AccumulatedTradeVolumeForTriggeredStop())
	to.SetBidAskQuantityStopInitialTriggerMet(m.BidAskQuantityStopInitialTriggerMet())
	to.SetNeedToOverrideLock(m.NeedToOverrideLock())
	to.SetCurrentMarketPrice(m.CurrentMarketPrice())
	to.SetCurrentMarketDateTime(m.CurrentMarketDateTime())
	to.SetSupportOrderFillBilling(m.SupportOrderFillBilling())
	to.SetIsBillable(m.IsBillable())
	to.SetQuantityForBilling(m.QuantityForBilling())
	to.SetNumberOfFailedOrderModifications(m.NumberOfFailedOrderModifications())
	to.SetDTCServerIndex(m.DTCServerIndex())
	to.SetClearingFirmID(m.ClearingFirmID())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICode(m.CTICode())
	to.SetObtainOrderActionDateTimeFromLastTradeTimeInChart(m.ObtainOrderActionDateTimeFromLastTradeTimeInChart())
	to.SetMaximumShowQuantity(m.MaximumShowQuantity())
	to.SetOrderSubmitted(m.OrderSubmitted())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetExternalLastActionDateTimeUTC(m.ExternalLastActionDateTimeUTC())
}

// CopyPointer
func (m TradeOrder) CopyPointer(to TradeOrderPointerBuilder) {
	to.SetIsOrderDeleted(m.IsOrderDeleted())
	to.SetInternalOrderID(m.InternalOrderID())
	to.SetOrderStatusCode(m.OrderStatusCode())
	to.SetOrderStatusBeforePendingModify(m.OrderStatusBeforePendingModify())
	to.SetOrderStatusBeforePendingCancel(m.OrderStatusBeforePendingCancel())
	to.SetServiceOrderID(m.ServiceOrderID())
	to.SetActualSymbol(m.ActualSymbol())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetRealtimeFillStatus(m.RealtimeFillStatus())
	to.SetIsRestingOrderDuringFill(m.IsRestingOrderDuringFill())
	to.SetOrderRejectType(m.OrderRejectType())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetInternalOrderIDModifierForService(m.InternalOrderIDModifierForService())
	to.SetFIXClientOrderID(m.FIXClientOrderID())
	to.SetSequenceNumberBasedClientOrderID(m.SequenceNumberBasedClientOrderID())
	to.SetClientOrderIDForDTCServer(m.ClientOrderIDForDTCServer())
	to.SetPreviousClientOrderIDForDTCServer(m.PreviousClientOrderIDForDTCServer())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOriginatingClientUsername(m.OriginatingClientUsername())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetLastActionDateTime(m.LastActionDateTime())
	to.SetServiceUpdateDateTimeUTC(m.ServiceUpdateDateTimeUTC())
	to.SetOrderEntryTimeForService(m.OrderEntryTimeForService())
	to.SetLastModifyTimeForService(m.LastModifyTimeForService())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetTimeInForce(m.TimeInForce())
	to.SetOpenClose(m.OpenClose())
	to.SetTrailStopOffset1(m.TrailStopOffset1())
	to.SetTrailStopStep(m.TrailStopStep())
	to.SetTrailTriggerPrice(m.TrailTriggerPrice())
	to.SetTrailingStopTriggerOffset(m.TrailingStopTriggerOffset())
	to.SetTrailTriggerHit(m.TrailTriggerHit())
	to.SetTrailToBreakEvenStopOffset(m.TrailToBreakEvenStopOffset())
	to.SetMaximumChaseAmountAsPrice(m.MaximumChaseAmountAsPrice())
	to.SetInitialChaseOrderPrice1(m.InitialChaseOrderPrice1())
	to.SetInitialLastTradePriceForChaseOrders(m.InitialLastTradePriceForChaseOrders())
	to.SetTrailingStopTriggerOCOGroupNumber(m.TrailingStopTriggerOCOGroupNumber())
	to.SetLastModifyPrice1(m.LastModifyPrice1())
	to.SetLastModifyQuantity(m.LastModifyQuantity())
	to.SetCumulativeOrderQuantityFromParentFills(m.CumulativeOrderQuantityFromParentFills())
	to.SetPriorFilledQuantity(m.PriorFilledQuantity())
	to.SetTickSize(m.TickSize())
	to.SetValueFormat(m.ValueFormat())
	to.SetPriceMultiplier(m.PriceMultiplier())
	to.SetParentInternalOrderID(m.ParentInternalOrderID())
	to.SetTargetChildInternalOrderID(m.TargetChildInternalOrderID())
	to.SetStopChildInternalOrderID(m.StopChildInternalOrderID())
	to.SetAttachedOrderPriceOffset1(m.AttachedOrderPriceOffset1())
	to.SetLinkInternalOrderID(m.LinkInternalOrderID())
	to.SetOCOGroupInternalOrderID(m.OCOGroupInternalOrderID())
	to.SetOCOSiblingInternalOrderID(m.OCOSiblingInternalOrderID())
	to.SetDisableChildAndSiblingRelatedActions(m.DisableChildAndSiblingRelatedActions())
	to.SetOCOManagedByService(m.OCOManagedByService())
	to.SetBracketOrderServerManaged(m.BracketOrderServerManaged())
	to.SetLastOrderActionSource(m.LastOrderActionSource())
	to.SetStopLimitOrderStopPriceTriggered(m.StopLimitOrderStopPriceTriggered())
	to.SetOCOFullSiblingCancelOnPartialFill(m.OCOFullSiblingCancelOnPartialFill())
	to.SetReverseOnCompleteFill(m.ReverseOnCompleteFill())
	to.SetSupportScaleIn(m.SupportScaleIn())
	to.SetSupportScaleOut(m.SupportScaleOut())
	to.SetSourceChartNumber(m.SourceChartNumber())
	to.SetSourceChartbookFileName(m.SourceChartbookFileName())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetSimulatedOrder(m.SimulatedOrder())
	to.SetIsChartReplaying(m.IsChartReplaying())
	to.SetAttachedOrderOCOGroupNumber(m.AttachedOrderOCOGroupNumber())
	to.SetLastFillExecutionServiceID(m.LastFillExecutionServiceID())
	to.SetFillCount(m.FillCount())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTimeUTC(m.LastFillDateTimeUTC())
	to.SetRejectedStopOCOSiblingInternalOrderID(m.RejectedStopOCOSiblingInternalOrderID())
	to.SetRejectedStopReplacementMarketOrderQuantity(m.RejectedStopReplacementMarketOrderQuantity())
	to.SetEvaluatingForFill(m.EvaluatingForFill())
	to.SetLastProcessedTimeSalesRecordSequenceForPrices(m.LastProcessedTimeSalesRecordSequenceForPrices())
	to.SetIsMarketDataManagementOfOrderEnabled(m.IsMarketDataManagementOfOrderEnabled())
	to.SetTextTag(m.TextTag())
	to.SetTimedOutOrderRequestedStatusDateTime(m.TimedOutOrderRequestedStatusDateTime())
	to.SetRequestedStatusForTimedOutOrder(m.RequestedStatusForTimedOutOrder())
	to.SetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled(m.SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled())
	to.SetQuantityToIncreaseFromParentFill(m.QuantityToIncreaseFromParentFill())
	to.SetMoveToBreakevenStopReferencePrice(m.MoveToBreakevenStopReferencePrice())
	to.SetQuantityTriggeredStop_QuantityForTrigger(m.QuantityTriggeredStop_QuantityForTrigger())
	to.SetAccumulatedTradeVolumeForTriggeredStop(m.AccumulatedTradeVolumeForTriggeredStop())
	to.SetBidAskQuantityStopInitialTriggerMet(m.BidAskQuantityStopInitialTriggerMet())
	to.SetNeedToOverrideLock(m.NeedToOverrideLock())
	to.SetCurrentMarketPrice(m.CurrentMarketPrice())
	to.SetCurrentMarketDateTime(m.CurrentMarketDateTime())
	to.SetSupportOrderFillBilling(m.SupportOrderFillBilling())
	to.SetIsBillable(m.IsBillable())
	to.SetQuantityForBilling(m.QuantityForBilling())
	to.SetNumberOfFailedOrderModifications(m.NumberOfFailedOrderModifications())
	to.SetDTCServerIndex(m.DTCServerIndex())
	to.SetClearingFirmID(m.ClearingFirmID())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICode(m.CTICode())
	to.SetObtainOrderActionDateTimeFromLastTradeTimeInChart(m.ObtainOrderActionDateTimeFromLastTradeTimeInChart())
	to.SetMaximumShowQuantity(m.MaximumShowQuantity())
	to.SetOrderSubmitted(m.OrderSubmitted())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetExternalLastActionDateTimeUTC(m.ExternalLastActionDateTimeUTC())
}

// CopyPointer
func (m TradeOrderBuilder) CopyPointer(to TradeOrderPointerBuilder) {
	to.SetIsOrderDeleted(m.IsOrderDeleted())
	to.SetInternalOrderID(m.InternalOrderID())
	to.SetOrderStatusCode(m.OrderStatusCode())
	to.SetOrderStatusBeforePendingModify(m.OrderStatusBeforePendingModify())
	to.SetOrderStatusBeforePendingCancel(m.OrderStatusBeforePendingCancel())
	to.SetServiceOrderID(m.ServiceOrderID())
	to.SetActualSymbol(m.ActualSymbol())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetRealtimeFillStatus(m.RealtimeFillStatus())
	to.SetIsRestingOrderDuringFill(m.IsRestingOrderDuringFill())
	to.SetOrderRejectType(m.OrderRejectType())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetInternalOrderIDModifierForService(m.InternalOrderIDModifierForService())
	to.SetFIXClientOrderID(m.FIXClientOrderID())
	to.SetSequenceNumberBasedClientOrderID(m.SequenceNumberBasedClientOrderID())
	to.SetClientOrderIDForDTCServer(m.ClientOrderIDForDTCServer())
	to.SetPreviousClientOrderIDForDTCServer(m.PreviousClientOrderIDForDTCServer())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOriginatingClientUsername(m.OriginatingClientUsername())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetLastActionDateTime(m.LastActionDateTime())
	to.SetServiceUpdateDateTimeUTC(m.ServiceUpdateDateTimeUTC())
	to.SetOrderEntryTimeForService(m.OrderEntryTimeForService())
	to.SetLastModifyTimeForService(m.LastModifyTimeForService())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetTimeInForce(m.TimeInForce())
	to.SetOpenClose(m.OpenClose())
	to.SetTrailStopOffset1(m.TrailStopOffset1())
	to.SetTrailStopStep(m.TrailStopStep())
	to.SetTrailTriggerPrice(m.TrailTriggerPrice())
	to.SetTrailingStopTriggerOffset(m.TrailingStopTriggerOffset())
	to.SetTrailTriggerHit(m.TrailTriggerHit())
	to.SetTrailToBreakEvenStopOffset(m.TrailToBreakEvenStopOffset())
	to.SetMaximumChaseAmountAsPrice(m.MaximumChaseAmountAsPrice())
	to.SetInitialChaseOrderPrice1(m.InitialChaseOrderPrice1())
	to.SetInitialLastTradePriceForChaseOrders(m.InitialLastTradePriceForChaseOrders())
	to.SetTrailingStopTriggerOCOGroupNumber(m.TrailingStopTriggerOCOGroupNumber())
	to.SetLastModifyPrice1(m.LastModifyPrice1())
	to.SetLastModifyQuantity(m.LastModifyQuantity())
	to.SetCumulativeOrderQuantityFromParentFills(m.CumulativeOrderQuantityFromParentFills())
	to.SetPriorFilledQuantity(m.PriorFilledQuantity())
	to.SetTickSize(m.TickSize())
	to.SetValueFormat(m.ValueFormat())
	to.SetPriceMultiplier(m.PriceMultiplier())
	to.SetParentInternalOrderID(m.ParentInternalOrderID())
	to.SetTargetChildInternalOrderID(m.TargetChildInternalOrderID())
	to.SetStopChildInternalOrderID(m.StopChildInternalOrderID())
	to.SetAttachedOrderPriceOffset1(m.AttachedOrderPriceOffset1())
	to.SetLinkInternalOrderID(m.LinkInternalOrderID())
	to.SetOCOGroupInternalOrderID(m.OCOGroupInternalOrderID())
	to.SetOCOSiblingInternalOrderID(m.OCOSiblingInternalOrderID())
	to.SetDisableChildAndSiblingRelatedActions(m.DisableChildAndSiblingRelatedActions())
	to.SetOCOManagedByService(m.OCOManagedByService())
	to.SetBracketOrderServerManaged(m.BracketOrderServerManaged())
	to.SetLastOrderActionSource(m.LastOrderActionSource())
	to.SetStopLimitOrderStopPriceTriggered(m.StopLimitOrderStopPriceTriggered())
	to.SetOCOFullSiblingCancelOnPartialFill(m.OCOFullSiblingCancelOnPartialFill())
	to.SetReverseOnCompleteFill(m.ReverseOnCompleteFill())
	to.SetSupportScaleIn(m.SupportScaleIn())
	to.SetSupportScaleOut(m.SupportScaleOut())
	to.SetSourceChartNumber(m.SourceChartNumber())
	to.SetSourceChartbookFileName(m.SourceChartbookFileName())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetSimulatedOrder(m.SimulatedOrder())
	to.SetIsChartReplaying(m.IsChartReplaying())
	to.SetAttachedOrderOCOGroupNumber(m.AttachedOrderOCOGroupNumber())
	to.SetLastFillExecutionServiceID(m.LastFillExecutionServiceID())
	to.SetFillCount(m.FillCount())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTimeUTC(m.LastFillDateTimeUTC())
	to.SetRejectedStopOCOSiblingInternalOrderID(m.RejectedStopOCOSiblingInternalOrderID())
	to.SetRejectedStopReplacementMarketOrderQuantity(m.RejectedStopReplacementMarketOrderQuantity())
	to.SetEvaluatingForFill(m.EvaluatingForFill())
	to.SetLastProcessedTimeSalesRecordSequenceForPrices(m.LastProcessedTimeSalesRecordSequenceForPrices())
	to.SetIsMarketDataManagementOfOrderEnabled(m.IsMarketDataManagementOfOrderEnabled())
	to.SetTextTag(m.TextTag())
	to.SetTimedOutOrderRequestedStatusDateTime(m.TimedOutOrderRequestedStatusDateTime())
	to.SetRequestedStatusForTimedOutOrder(m.RequestedStatusForTimedOutOrder())
	to.SetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled(m.SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled())
	to.SetQuantityToIncreaseFromParentFill(m.QuantityToIncreaseFromParentFill())
	to.SetMoveToBreakevenStopReferencePrice(m.MoveToBreakevenStopReferencePrice())
	to.SetQuantityTriggeredStop_QuantityForTrigger(m.QuantityTriggeredStop_QuantityForTrigger())
	to.SetAccumulatedTradeVolumeForTriggeredStop(m.AccumulatedTradeVolumeForTriggeredStop())
	to.SetBidAskQuantityStopInitialTriggerMet(m.BidAskQuantityStopInitialTriggerMet())
	to.SetNeedToOverrideLock(m.NeedToOverrideLock())
	to.SetCurrentMarketPrice(m.CurrentMarketPrice())
	to.SetCurrentMarketDateTime(m.CurrentMarketDateTime())
	to.SetSupportOrderFillBilling(m.SupportOrderFillBilling())
	to.SetIsBillable(m.IsBillable())
	to.SetQuantityForBilling(m.QuantityForBilling())
	to.SetNumberOfFailedOrderModifications(m.NumberOfFailedOrderModifications())
	to.SetDTCServerIndex(m.DTCServerIndex())
	to.SetClearingFirmID(m.ClearingFirmID())
	to.SetSenderSubID(m.SenderSubID())
	to.SetSenderLocationId(m.SenderLocationId())
	to.SetSelfMatchPreventionID(m.SelfMatchPreventionID())
	to.SetCTICode(m.CTICode())
	to.SetObtainOrderActionDateTimeFromLastTradeTimeInChart(m.ObtainOrderActionDateTimeFromLastTradeTimeInChart())
	to.SetMaximumShowQuantity(m.MaximumShowQuantity())
	to.SetOrderSubmitted(m.OrderSubmitted())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetExternalLastActionDateTimeUTC(m.ExternalLastActionDateTimeUTC())
}
