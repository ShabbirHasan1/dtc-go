package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeOrderPointer struct {
	message.VLSPointer
}

type TradeOrderPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeOrder() TradeOrderPointerBuilder {
	a := TradeOrderPointerBuilder{message.AllocVLSBuilder(535)}
	a.Ptr.SetBytes(0, _TradeOrderDefault)
	return a
}

func AllocTradeOrderFrom(b []byte) TradeOrderPointer {
	return TradeOrderPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m TradeOrderPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeOrderDefault)
}

// ToBuilder
func (m TradeOrderPointer) ToBuilder() TradeOrderPointerBuilder {
	return TradeOrderPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeOrderPointerBuilder) Finish() TradeOrderPointer {
	return TradeOrderPointer{m.VLSPointerBuilder.Finish()}
}

// IsOrderDeleted
func (m TradeOrderPointer) IsOrderDeleted() uint8 {
	return message.Uint8VLS(m.Ptr, 7, 6)
}

// IsOrderDeleted
func (m TradeOrderPointerBuilder) IsOrderDeleted() uint8 {
	return message.Uint8VLS(m.Ptr, 7, 6)
}

// InternalOrderID
func (m TradeOrderPointer) InternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 15, 7)
}

// InternalOrderID
func (m TradeOrderPointerBuilder) InternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 15, 7)
}

// OrderStatusCode
func (m TradeOrderPointer) OrderStatusCode() uint16 {
	return message.Uint16VLS(m.Ptr, 17, 15)
}

// OrderStatusCode
func (m TradeOrderPointerBuilder) OrderStatusCode() uint16 {
	return message.Uint16VLS(m.Ptr, 17, 15)
}

// OrderStatusBeforePendingModify
func (m TradeOrderPointer) OrderStatusBeforePendingModify() uint16 {
	return message.Uint16VLS(m.Ptr, 19, 17)
}

// OrderStatusBeforePendingModify
func (m TradeOrderPointerBuilder) OrderStatusBeforePendingModify() uint16 {
	return message.Uint16VLS(m.Ptr, 19, 17)
}

// OrderStatusBeforePendingCancel
func (m TradeOrderPointer) OrderStatusBeforePendingCancel() uint16 {
	return message.Uint16VLS(m.Ptr, 21, 19)
}

// OrderStatusBeforePendingCancel
func (m TradeOrderPointerBuilder) OrderStatusBeforePendingCancel() uint16 {
	return message.Uint16VLS(m.Ptr, 21, 19)
}

// ServiceOrderID
func (m TradeOrderPointer) ServiceOrderID() string {
	return message.StringVLS(m.Ptr, 25, 21)
}

// ServiceOrderID
func (m TradeOrderPointerBuilder) ServiceOrderID() string {
	return message.StringVLS(m.Ptr, 25, 21)
}

// ActualSymbol
func (m TradeOrderPointer) ActualSymbol() string {
	return message.StringVLS(m.Ptr, 29, 25)
}

// ActualSymbol
func (m TradeOrderPointerBuilder) ActualSymbol() string {
	return message.StringVLS(m.Ptr, 29, 25)
}

// OrderType
func (m TradeOrderPointer) OrderType() int32 {
	return message.Int32VLS(m.Ptr, 33, 29)
}

// OrderType
func (m TradeOrderPointerBuilder) OrderType() int32 {
	return message.Int32VLS(m.Ptr, 33, 29)
}

// BuySell
func (m TradeOrderPointer) BuySell() uint16 {
	return message.Uint16VLS(m.Ptr, 35, 33)
}

// BuySell
func (m TradeOrderPointerBuilder) BuySell() uint16 {
	return message.Uint16VLS(m.Ptr, 35, 33)
}

// Price1
func (m TradeOrderPointer) Price1() float64 {
	return message.Float64VLS(m.Ptr, 43, 35)
}

// Price1
func (m TradeOrderPointerBuilder) Price1() float64 {
	return message.Float64VLS(m.Ptr, 43, 35)
}

// Price2
func (m TradeOrderPointer) Price2() float64 {
	return message.Float64VLS(m.Ptr, 51, 43)
}

// Price2
func (m TradeOrderPointerBuilder) Price2() float64 {
	return message.Float64VLS(m.Ptr, 51, 43)
}

// OrderQuantity
func (m TradeOrderPointer) OrderQuantity() float64 {
	return message.Float64VLS(m.Ptr, 59, 51)
}

// OrderQuantity
func (m TradeOrderPointerBuilder) OrderQuantity() float64 {
	return message.Float64VLS(m.Ptr, 59, 51)
}

// FilledQuantity
func (m TradeOrderPointer) FilledQuantity() float64 {
	return message.Float64VLS(m.Ptr, 67, 59)
}

// FilledQuantity
func (m TradeOrderPointerBuilder) FilledQuantity() float64 {
	return message.Float64VLS(m.Ptr, 67, 59)
}

// AverageFillPrice
func (m TradeOrderPointer) AverageFillPrice() float64 {
	return message.Float64VLS(m.Ptr, 75, 67)
}

// AverageFillPrice
func (m TradeOrderPointerBuilder) AverageFillPrice() float64 {
	return message.Float64VLS(m.Ptr, 75, 67)
}

// RealtimeFillStatus
func (m TradeOrderPointer) RealtimeFillStatus() int32 {
	return message.Int32VLS(m.Ptr, 79, 75)
}

// RealtimeFillStatus
func (m TradeOrderPointerBuilder) RealtimeFillStatus() int32 {
	return message.Int32VLS(m.Ptr, 79, 75)
}

// IsRestingOrderDuringFill
func (m TradeOrderPointer) IsRestingOrderDuringFill() uint8 {
	return message.Uint8VLS(m.Ptr, 80, 79)
}

// IsRestingOrderDuringFill
func (m TradeOrderPointerBuilder) IsRestingOrderDuringFill() uint8 {
	return message.Uint8VLS(m.Ptr, 80, 79)
}

// OrderRejectType
func (m TradeOrderPointer) OrderRejectType() int32 {
	return message.Int32VLS(m.Ptr, 84, 80)
}

// OrderRejectType
func (m TradeOrderPointerBuilder) OrderRejectType() int32 {
	return message.Int32VLS(m.Ptr, 84, 80)
}

// TradeAccount
func (m TradeOrderPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 88, 84)
}

// TradeAccount
func (m TradeOrderPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 88, 84)
}

// SubAccountIdentifier
func (m TradeOrderPointer) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 92, 88)
}

// SubAccountIdentifier
func (m TradeOrderPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 92, 88)
}

// InternalOrderIDModifierForService
func (m TradeOrderPointer) InternalOrderIDModifierForService() int32 {
	return message.Int32VLS(m.Ptr, 96, 92)
}

// InternalOrderIDModifierForService
func (m TradeOrderPointerBuilder) InternalOrderIDModifierForService() int32 {
	return message.Int32VLS(m.Ptr, 96, 92)
}

// FIXClientOrderID
func (m TradeOrderPointer) FIXClientOrderID() string {
	return message.StringVLS(m.Ptr, 100, 96)
}

// FIXClientOrderID
func (m TradeOrderPointerBuilder) FIXClientOrderID() string {
	return message.StringVLS(m.Ptr, 100, 96)
}

// SequenceNumberBasedClientOrderID
func (m TradeOrderPointer) SequenceNumberBasedClientOrderID() uint32 {
	return message.Uint32VLS(m.Ptr, 104, 100)
}

// SequenceNumberBasedClientOrderID
func (m TradeOrderPointerBuilder) SequenceNumberBasedClientOrderID() uint32 {
	return message.Uint32VLS(m.Ptr, 104, 100)
}

// ClientOrderIDForDTCServer
func (m TradeOrderPointer) ClientOrderIDForDTCServer() string {
	return message.StringVLS(m.Ptr, 108, 104)
}

// ClientOrderIDForDTCServer
func (m TradeOrderPointerBuilder) ClientOrderIDForDTCServer() string {
	return message.StringVLS(m.Ptr, 108, 104)
}

// PreviousClientOrderIDForDTCServer
func (m TradeOrderPointer) PreviousClientOrderIDForDTCServer() string {
	return message.StringVLS(m.Ptr, 112, 108)
}

// PreviousClientOrderIDForDTCServer
func (m TradeOrderPointerBuilder) PreviousClientOrderIDForDTCServer() string {
	return message.StringVLS(m.Ptr, 112, 108)
}

// ExchangeOrderID
func (m TradeOrderPointer) ExchangeOrderID() string {
	return message.StringVLS(m.Ptr, 116, 112)
}

// ExchangeOrderID
func (m TradeOrderPointerBuilder) ExchangeOrderID() string {
	return message.StringVLS(m.Ptr, 116, 112)
}

// OriginatingClientUsername
func (m TradeOrderPointer) OriginatingClientUsername() string {
	return message.StringVLS(m.Ptr, 120, 116)
}

// OriginatingClientUsername
func (m TradeOrderPointerBuilder) OriginatingClientUsername() string {
	return message.StringVLS(m.Ptr, 120, 116)
}

// EntryDateTime
func (m TradeOrderPointer) EntryDateTime() int64 {
	return message.Int64VLS(m.Ptr, 128, 120)
}

// EntryDateTime
func (m TradeOrderPointerBuilder) EntryDateTime() int64 {
	return message.Int64VLS(m.Ptr, 128, 120)
}

// LastActionDateTime
func (m TradeOrderPointer) LastActionDateTime() int64 {
	return message.Int64VLS(m.Ptr, 136, 128)
}

// LastActionDateTime
func (m TradeOrderPointerBuilder) LastActionDateTime() int64 {
	return message.Int64VLS(m.Ptr, 136, 128)
}

// ServiceUpdateDateTimeUTC
func (m TradeOrderPointer) ServiceUpdateDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 144, 136)
}

// ServiceUpdateDateTimeUTC
func (m TradeOrderPointerBuilder) ServiceUpdateDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 144, 136)
}

// OrderEntryTimeForService
func (m TradeOrderPointer) OrderEntryTimeForService() uint32 {
	return message.Uint32VLS(m.Ptr, 148, 144)
}

// OrderEntryTimeForService
func (m TradeOrderPointerBuilder) OrderEntryTimeForService() uint32 {
	return message.Uint32VLS(m.Ptr, 148, 144)
}

// LastModifyTimeForService
func (m TradeOrderPointer) LastModifyTimeForService() uint32 {
	return message.Uint32VLS(m.Ptr, 152, 148)
}

// LastModifyTimeForService
func (m TradeOrderPointerBuilder) LastModifyTimeForService() uint32 {
	return message.Uint32VLS(m.Ptr, 152, 148)
}

// GoodTillDateTime
func (m TradeOrderPointer) GoodTillDateTime() int64 {
	return message.Int64VLS(m.Ptr, 160, 152)
}

// GoodTillDateTime
func (m TradeOrderPointerBuilder) GoodTillDateTime() int64 {
	return message.Int64VLS(m.Ptr, 160, 152)
}

// TimeInForce
func (m TradeOrderPointer) TimeInForce() int32 {
	return message.Int32VLS(m.Ptr, 164, 160)
}

// TimeInForce
func (m TradeOrderPointerBuilder) TimeInForce() int32 {
	return message.Int32VLS(m.Ptr, 164, 160)
}

// OpenClose
func (m TradeOrderPointer) OpenClose() uint16 {
	return message.Uint16VLS(m.Ptr, 166, 164)
}

// OpenClose
func (m TradeOrderPointerBuilder) OpenClose() uint16 {
	return message.Uint16VLS(m.Ptr, 166, 164)
}

// TrailStopOffset1
func (m TradeOrderPointer) TrailStopOffset1() float64 {
	return message.Float64VLS(m.Ptr, 174, 166)
}

// TrailStopOffset1
func (m TradeOrderPointerBuilder) TrailStopOffset1() float64 {
	return message.Float64VLS(m.Ptr, 174, 166)
}

// TrailStopStep
func (m TradeOrderPointer) TrailStopStep() float64 {
	return message.Float64VLS(m.Ptr, 182, 174)
}

// TrailStopStep
func (m TradeOrderPointerBuilder) TrailStopStep() float64 {
	return message.Float64VLS(m.Ptr, 182, 174)
}

// TrailTriggerPrice
func (m TradeOrderPointer) TrailTriggerPrice() float64 {
	return message.Float64VLS(m.Ptr, 190, 182)
}

// TrailTriggerPrice
func (m TradeOrderPointerBuilder) TrailTriggerPrice() float64 {
	return message.Float64VLS(m.Ptr, 190, 182)
}

// TrailingStopTriggerOffset
func (m TradeOrderPointer) TrailingStopTriggerOffset() float64 {
	return message.Float64VLS(m.Ptr, 198, 190)
}

// TrailingStopTriggerOffset
func (m TradeOrderPointerBuilder) TrailingStopTriggerOffset() float64 {
	return message.Float64VLS(m.Ptr, 198, 190)
}

// TrailTriggerHit
func (m TradeOrderPointer) TrailTriggerHit() uint8 {
	return message.Uint8VLS(m.Ptr, 199, 198)
}

// TrailTriggerHit
func (m TradeOrderPointerBuilder) TrailTriggerHit() uint8 {
	return message.Uint8VLS(m.Ptr, 199, 198)
}

// TrailToBreakEvenStopOffset
func (m TradeOrderPointer) TrailToBreakEvenStopOffset() float64 {
	return message.Float64VLS(m.Ptr, 207, 199)
}

// TrailToBreakEvenStopOffset
func (m TradeOrderPointerBuilder) TrailToBreakEvenStopOffset() float64 {
	return message.Float64VLS(m.Ptr, 207, 199)
}

// MaximumChaseAmountAsPrice
func (m TradeOrderPointer) MaximumChaseAmountAsPrice() float64 {
	return message.Float64VLS(m.Ptr, 215, 207)
}

// MaximumChaseAmountAsPrice
func (m TradeOrderPointerBuilder) MaximumChaseAmountAsPrice() float64 {
	return message.Float64VLS(m.Ptr, 215, 207)
}

// InitialChaseOrderPrice1
func (m TradeOrderPointer) InitialChaseOrderPrice1() float64 {
	return message.Float64VLS(m.Ptr, 223, 215)
}

// InitialChaseOrderPrice1
func (m TradeOrderPointerBuilder) InitialChaseOrderPrice1() float64 {
	return message.Float64VLS(m.Ptr, 223, 215)
}

// InitialLastTradePriceForChaseOrders
func (m TradeOrderPointer) InitialLastTradePriceForChaseOrders() float64 {
	return message.Float64VLS(m.Ptr, 231, 223)
}

// InitialLastTradePriceForChaseOrders
func (m TradeOrderPointerBuilder) InitialLastTradePriceForChaseOrders() float64 {
	return message.Float64VLS(m.Ptr, 231, 223)
}

// TrailingStopTriggerOCOGroupNumber
func (m TradeOrderPointer) TrailingStopTriggerOCOGroupNumber() int32 {
	return message.Int32VLS(m.Ptr, 235, 231)
}

// TrailingStopTriggerOCOGroupNumber
func (m TradeOrderPointerBuilder) TrailingStopTriggerOCOGroupNumber() int32 {
	return message.Int32VLS(m.Ptr, 235, 231)
}

// LastModifyPrice1
func (m TradeOrderPointer) LastModifyPrice1() float64 {
	return message.Float64VLS(m.Ptr, 243, 235)
}

// LastModifyPrice1
func (m TradeOrderPointerBuilder) LastModifyPrice1() float64 {
	return message.Float64VLS(m.Ptr, 243, 235)
}

// LastModifyQuantity
func (m TradeOrderPointer) LastModifyQuantity() float64 {
	return message.Float64VLS(m.Ptr, 251, 243)
}

// LastModifyQuantity
func (m TradeOrderPointerBuilder) LastModifyQuantity() float64 {
	return message.Float64VLS(m.Ptr, 251, 243)
}

// CumulativeOrderQuantityFromParentFills
func (m TradeOrderPointer) CumulativeOrderQuantityFromParentFills() float64 {
	return message.Float64VLS(m.Ptr, 259, 251)
}

// CumulativeOrderQuantityFromParentFills
func (m TradeOrderPointerBuilder) CumulativeOrderQuantityFromParentFills() float64 {
	return message.Float64VLS(m.Ptr, 259, 251)
}

// PriorFilledQuantity
func (m TradeOrderPointer) PriorFilledQuantity() float64 {
	return message.Float64VLS(m.Ptr, 267, 259)
}

// PriorFilledQuantity
func (m TradeOrderPointerBuilder) PriorFilledQuantity() float64 {
	return message.Float64VLS(m.Ptr, 267, 259)
}

// TickSize
func (m TradeOrderPointer) TickSize() float32 {
	return message.Float32VLS(m.Ptr, 271, 267)
}

// TickSize
func (m TradeOrderPointerBuilder) TickSize() float32 {
	return message.Float32VLS(m.Ptr, 271, 267)
}

// ValueFormat
func (m TradeOrderPointer) ValueFormat() int32 {
	return message.Int32VLS(m.Ptr, 275, 271)
}

// ValueFormat
func (m TradeOrderPointerBuilder) ValueFormat() int32 {
	return message.Int32VLS(m.Ptr, 275, 271)
}

// PriceMultiplier
func (m TradeOrderPointer) PriceMultiplier() float32 {
	return message.Float32VLS(m.Ptr, 279, 275)
}

// PriceMultiplier
func (m TradeOrderPointerBuilder) PriceMultiplier() float32 {
	return message.Float32VLS(m.Ptr, 279, 275)
}

// ParentInternalOrderID
func (m TradeOrderPointer) ParentInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 287, 279)
}

// ParentInternalOrderID
func (m TradeOrderPointerBuilder) ParentInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 287, 279)
}

// TargetChildInternalOrderID
func (m TradeOrderPointer) TargetChildInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 295, 287)
}

// TargetChildInternalOrderID
func (m TradeOrderPointerBuilder) TargetChildInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 295, 287)
}

// StopChildInternalOrderID
func (m TradeOrderPointer) StopChildInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 303, 295)
}

// StopChildInternalOrderID
func (m TradeOrderPointerBuilder) StopChildInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 303, 295)
}

// AttachedOrderPriceOffset1
func (m TradeOrderPointer) AttachedOrderPriceOffset1() float64 {
	return message.Float64VLS(m.Ptr, 311, 303)
}

// AttachedOrderPriceOffset1
func (m TradeOrderPointerBuilder) AttachedOrderPriceOffset1() float64 {
	return message.Float64VLS(m.Ptr, 311, 303)
}

// LinkInternalOrderID
func (m TradeOrderPointer) LinkInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 319, 311)
}

// LinkInternalOrderID
func (m TradeOrderPointerBuilder) LinkInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 319, 311)
}

// OCOGroupInternalOrderID
func (m TradeOrderPointer) OCOGroupInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 327, 319)
}

// OCOGroupInternalOrderID
func (m TradeOrderPointerBuilder) OCOGroupInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 327, 319)
}

// OCOSiblingInternalOrderID
func (m TradeOrderPointer) OCOSiblingInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 335, 327)
}

// OCOSiblingInternalOrderID
func (m TradeOrderPointerBuilder) OCOSiblingInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 335, 327)
}

// DisableChildAndSiblingRelatedActions
func (m TradeOrderPointer) DisableChildAndSiblingRelatedActions() uint8 {
	return message.Uint8VLS(m.Ptr, 336, 335)
}

// DisableChildAndSiblingRelatedActions
func (m TradeOrderPointerBuilder) DisableChildAndSiblingRelatedActions() uint8 {
	return message.Uint8VLS(m.Ptr, 336, 335)
}

// OCOManagedByService
func (m TradeOrderPointer) OCOManagedByService() uint8 {
	return message.Uint8VLS(m.Ptr, 337, 336)
}

// OCOManagedByService
func (m TradeOrderPointerBuilder) OCOManagedByService() uint8 {
	return message.Uint8VLS(m.Ptr, 337, 336)
}

// BracketOrderServerManaged
func (m TradeOrderPointer) BracketOrderServerManaged() uint8 {
	return message.Uint8VLS(m.Ptr, 338, 337)
}

// BracketOrderServerManaged
func (m TradeOrderPointerBuilder) BracketOrderServerManaged() uint8 {
	return message.Uint8VLS(m.Ptr, 338, 337)
}

// LastOrderActionSource
func (m TradeOrderPointer) LastOrderActionSource() string {
	return message.StringVLS(m.Ptr, 342, 338)
}

// LastOrderActionSource
func (m TradeOrderPointerBuilder) LastOrderActionSource() string {
	return message.StringVLS(m.Ptr, 342, 338)
}

// StopLimitOrderStopPriceTriggered
func (m TradeOrderPointer) StopLimitOrderStopPriceTriggered() uint8 {
	return message.Uint8VLS(m.Ptr, 343, 342)
}

// StopLimitOrderStopPriceTriggered
func (m TradeOrderPointerBuilder) StopLimitOrderStopPriceTriggered() uint8 {
	return message.Uint8VLS(m.Ptr, 343, 342)
}

// OCOFullSiblingCancelOnPartialFill
func (m TradeOrderPointer) OCOFullSiblingCancelOnPartialFill() uint8 {
	return message.Uint8VLS(m.Ptr, 344, 343)
}

// OCOFullSiblingCancelOnPartialFill
func (m TradeOrderPointerBuilder) OCOFullSiblingCancelOnPartialFill() uint8 {
	return message.Uint8VLS(m.Ptr, 344, 343)
}

// ReverseOnCompleteFill
func (m TradeOrderPointer) ReverseOnCompleteFill() uint8 {
	return message.Uint8VLS(m.Ptr, 345, 344)
}

// ReverseOnCompleteFill
func (m TradeOrderPointerBuilder) ReverseOnCompleteFill() uint8 {
	return message.Uint8VLS(m.Ptr, 345, 344)
}

// SupportScaleIn
func (m TradeOrderPointer) SupportScaleIn() uint8 {
	return message.Uint8VLS(m.Ptr, 346, 345)
}

// SupportScaleIn
func (m TradeOrderPointerBuilder) SupportScaleIn() uint8 {
	return message.Uint8VLS(m.Ptr, 346, 345)
}

// SupportScaleOut
func (m TradeOrderPointer) SupportScaleOut() uint8 {
	return message.Uint8VLS(m.Ptr, 347, 346)
}

// SupportScaleOut
func (m TradeOrderPointerBuilder) SupportScaleOut() uint8 {
	return message.Uint8VLS(m.Ptr, 347, 346)
}

// SourceChartNumber
func (m TradeOrderPointer) SourceChartNumber() int32 {
	return message.Int32VLS(m.Ptr, 351, 347)
}

// SourceChartNumber
func (m TradeOrderPointerBuilder) SourceChartNumber() int32 {
	return message.Int32VLS(m.Ptr, 351, 347)
}

// SourceChartbookFileName
func (m TradeOrderPointer) SourceChartbookFileName() string {
	return message.StringVLS(m.Ptr, 355, 351)
}

// SourceChartbookFileName
func (m TradeOrderPointerBuilder) SourceChartbookFileName() string {
	return message.StringVLS(m.Ptr, 355, 351)
}

// IsAutomatedOrder
func (m TradeOrderPointer) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 356, 355)
}

// IsAutomatedOrder
func (m TradeOrderPointerBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 356, 355)
}

// SimulatedOrder
func (m TradeOrderPointer) SimulatedOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 357, 356)
}

// SimulatedOrder
func (m TradeOrderPointerBuilder) SimulatedOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 357, 356)
}

// IsChartReplaying
func (m TradeOrderPointer) IsChartReplaying() uint8 {
	return message.Uint8VLS(m.Ptr, 358, 357)
}

// IsChartReplaying
func (m TradeOrderPointerBuilder) IsChartReplaying() uint8 {
	return message.Uint8VLS(m.Ptr, 358, 357)
}

// AttachedOrderOCOGroupNumber
func (m TradeOrderPointer) AttachedOrderOCOGroupNumber() int32 {
	return message.Int32VLS(m.Ptr, 362, 358)
}

// AttachedOrderOCOGroupNumber
func (m TradeOrderPointerBuilder) AttachedOrderOCOGroupNumber() int32 {
	return message.Int32VLS(m.Ptr, 362, 358)
}

// LastFillExecutionServiceID
func (m TradeOrderPointer) LastFillExecutionServiceID() string {
	return message.StringVLS(m.Ptr, 366, 362)
}

// LastFillExecutionServiceID
func (m TradeOrderPointerBuilder) LastFillExecutionServiceID() string {
	return message.StringVLS(m.Ptr, 366, 362)
}

// FillCount
func (m TradeOrderPointer) FillCount() int32 {
	return message.Int32VLS(m.Ptr, 370, 366)
}

// FillCount
func (m TradeOrderPointerBuilder) FillCount() int32 {
	return message.Int32VLS(m.Ptr, 370, 366)
}

// LastFillQuantity
func (m TradeOrderPointer) LastFillQuantity() float64 {
	return message.Float64VLS(m.Ptr, 378, 370)
}

// LastFillQuantity
func (m TradeOrderPointerBuilder) LastFillQuantity() float64 {
	return message.Float64VLS(m.Ptr, 378, 370)
}

// LastFillPrice
func (m TradeOrderPointer) LastFillPrice() float64 {
	return message.Float64VLS(m.Ptr, 386, 378)
}

// LastFillPrice
func (m TradeOrderPointerBuilder) LastFillPrice() float64 {
	return message.Float64VLS(m.Ptr, 386, 378)
}

// LastFillDateTimeUTC
func (m TradeOrderPointer) LastFillDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 394, 386)
}

// LastFillDateTimeUTC
func (m TradeOrderPointerBuilder) LastFillDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 394, 386)
}

// RejectedStopOCOSiblingInternalOrderID
func (m TradeOrderPointer) RejectedStopOCOSiblingInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 402, 394)
}

// RejectedStopOCOSiblingInternalOrderID
func (m TradeOrderPointerBuilder) RejectedStopOCOSiblingInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 402, 394)
}

// RejectedStopReplacementMarketOrderQuantity
func (m TradeOrderPointer) RejectedStopReplacementMarketOrderQuantity() float64 {
	return message.Float64VLS(m.Ptr, 410, 402)
}

// RejectedStopReplacementMarketOrderQuantity
func (m TradeOrderPointerBuilder) RejectedStopReplacementMarketOrderQuantity() float64 {
	return message.Float64VLS(m.Ptr, 410, 402)
}

// EvaluatingForFill
func (m TradeOrderPointer) EvaluatingForFill() uint8 {
	return message.Uint8VLS(m.Ptr, 411, 410)
}

// EvaluatingForFill
func (m TradeOrderPointerBuilder) EvaluatingForFill() uint8 {
	return message.Uint8VLS(m.Ptr, 411, 410)
}

// LastProcessedTimeSalesRecordSequenceForPrices
func (m TradeOrderPointer) LastProcessedTimeSalesRecordSequenceForPrices() uint32 {
	return message.Uint32VLS(m.Ptr, 415, 411)
}

// LastProcessedTimeSalesRecordSequenceForPrices
func (m TradeOrderPointerBuilder) LastProcessedTimeSalesRecordSequenceForPrices() uint32 {
	return message.Uint32VLS(m.Ptr, 415, 411)
}

// IsMarketDataManagementOfOrderEnabled
func (m TradeOrderPointer) IsMarketDataManagementOfOrderEnabled() uint8 {
	return message.Uint8VLS(m.Ptr, 416, 415)
}

// IsMarketDataManagementOfOrderEnabled
func (m TradeOrderPointerBuilder) IsMarketDataManagementOfOrderEnabled() uint8 {
	return message.Uint8VLS(m.Ptr, 416, 415)
}

// TextTag
func (m TradeOrderPointer) TextTag() string {
	return message.StringVLS(m.Ptr, 420, 416)
}

// TextTag
func (m TradeOrderPointerBuilder) TextTag() string {
	return message.StringVLS(m.Ptr, 420, 416)
}

// TimedOutOrderRequestedStatusDateTime
func (m TradeOrderPointer) TimedOutOrderRequestedStatusDateTime() int64 {
	return message.Int64VLS(m.Ptr, 428, 420)
}

// TimedOutOrderRequestedStatusDateTime
func (m TradeOrderPointerBuilder) TimedOutOrderRequestedStatusDateTime() int64 {
	return message.Int64VLS(m.Ptr, 428, 420)
}

// RequestedStatusForTimedOutOrder
func (m TradeOrderPointer) RequestedStatusForTimedOutOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 429, 428)
}

// RequestedStatusForTimedOutOrder
func (m TradeOrderPointerBuilder) RequestedStatusForTimedOutOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 429, 428)
}

// SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled
func (m TradeOrderPointer) SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled() uint8 {
	return message.Uint8VLS(m.Ptr, 430, 429)
}

// SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled
func (m TradeOrderPointerBuilder) SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled() uint8 {
	return message.Uint8VLS(m.Ptr, 430, 429)
}

// QuantityToIncreaseFromParentFill
func (m TradeOrderPointer) QuantityToIncreaseFromParentFill() float64 {
	return message.Float64VLS(m.Ptr, 438, 430)
}

// QuantityToIncreaseFromParentFill
func (m TradeOrderPointerBuilder) QuantityToIncreaseFromParentFill() float64 {
	return message.Float64VLS(m.Ptr, 438, 430)
}

// MoveToBreakevenStopReferencePrice
func (m TradeOrderPointer) MoveToBreakevenStopReferencePrice() float64 {
	return message.Float64VLS(m.Ptr, 446, 438)
}

// MoveToBreakevenStopReferencePrice
func (m TradeOrderPointerBuilder) MoveToBreakevenStopReferencePrice() float64 {
	return message.Float64VLS(m.Ptr, 446, 438)
}

// QuantityTriggeredStop_QuantityForTrigger
func (m TradeOrderPointer) QuantityTriggeredStop_QuantityForTrigger() float64 {
	return message.Float64VLS(m.Ptr, 454, 446)
}

// QuantityTriggeredStop_QuantityForTrigger
func (m TradeOrderPointerBuilder) QuantityTriggeredStop_QuantityForTrigger() float64 {
	return message.Float64VLS(m.Ptr, 454, 446)
}

// AccumulatedTradeVolumeForTriggeredStop
func (m TradeOrderPointer) AccumulatedTradeVolumeForTriggeredStop() float64 {
	return message.Float64VLS(m.Ptr, 462, 454)
}

// AccumulatedTradeVolumeForTriggeredStop
func (m TradeOrderPointerBuilder) AccumulatedTradeVolumeForTriggeredStop() float64 {
	return message.Float64VLS(m.Ptr, 462, 454)
}

// BidAskQuantityStopInitialTriggerMet
func (m TradeOrderPointer) BidAskQuantityStopInitialTriggerMet() uint8 {
	return message.Uint8VLS(m.Ptr, 463, 462)
}

// BidAskQuantityStopInitialTriggerMet
func (m TradeOrderPointerBuilder) BidAskQuantityStopInitialTriggerMet() uint8 {
	return message.Uint8VLS(m.Ptr, 463, 462)
}

// NeedToOverrideLock
func (m TradeOrderPointer) NeedToOverrideLock() uint8 {
	return message.Uint8VLS(m.Ptr, 464, 463)
}

// NeedToOverrideLock
func (m TradeOrderPointerBuilder) NeedToOverrideLock() uint8 {
	return message.Uint8VLS(m.Ptr, 464, 463)
}

// CurrentMarketPrice
func (m TradeOrderPointer) CurrentMarketPrice() float64 {
	return message.Float64VLS(m.Ptr, 472, 464)
}

// CurrentMarketPrice
func (m TradeOrderPointerBuilder) CurrentMarketPrice() float64 {
	return message.Float64VLS(m.Ptr, 472, 464)
}

// CurrentMarketDateTime
func (m TradeOrderPointer) CurrentMarketDateTime() int64 {
	return message.Int64VLS(m.Ptr, 480, 472)
}

// CurrentMarketDateTime
func (m TradeOrderPointerBuilder) CurrentMarketDateTime() int64 {
	return message.Int64VLS(m.Ptr, 480, 472)
}

// SupportOrderFillBilling
func (m TradeOrderPointer) SupportOrderFillBilling() uint8 {
	return message.Uint8VLS(m.Ptr, 481, 480)
}

// SupportOrderFillBilling
func (m TradeOrderPointerBuilder) SupportOrderFillBilling() uint8 {
	return message.Uint8VLS(m.Ptr, 481, 480)
}

// IsBillable
func (m TradeOrderPointer) IsBillable() uint8 {
	return message.Uint8VLS(m.Ptr, 482, 481)
}

// IsBillable
func (m TradeOrderPointerBuilder) IsBillable() uint8 {
	return message.Uint8VLS(m.Ptr, 482, 481)
}

// QuantityForBilling
func (m TradeOrderPointer) QuantityForBilling() int32 {
	return message.Int32VLS(m.Ptr, 486, 482)
}

// QuantityForBilling
func (m TradeOrderPointerBuilder) QuantityForBilling() int32 {
	return message.Int32VLS(m.Ptr, 486, 482)
}

// NumberOfFailedOrderModifications
func (m TradeOrderPointer) NumberOfFailedOrderModifications() uint32 {
	return message.Uint32VLS(m.Ptr, 490, 486)
}

// NumberOfFailedOrderModifications
func (m TradeOrderPointerBuilder) NumberOfFailedOrderModifications() uint32 {
	return message.Uint32VLS(m.Ptr, 490, 486)
}

// DTCServerIndex
func (m TradeOrderPointer) DTCServerIndex() int32 {
	return message.Int32VLS(m.Ptr, 494, 490)
}

// DTCServerIndex
func (m TradeOrderPointerBuilder) DTCServerIndex() int32 {
	return message.Int32VLS(m.Ptr, 494, 490)
}

// ClearingFirmID
func (m TradeOrderPointer) ClearingFirmID() string {
	return message.StringVLS(m.Ptr, 498, 494)
}

// ClearingFirmID
func (m TradeOrderPointerBuilder) ClearingFirmID() string {
	return message.StringVLS(m.Ptr, 498, 494)
}

// SenderSubID
func (m TradeOrderPointer) SenderSubID() string {
	return message.StringVLS(m.Ptr, 502, 498)
}

// SenderSubID
func (m TradeOrderPointerBuilder) SenderSubID() string {
	return message.StringVLS(m.Ptr, 502, 498)
}

// SenderLocationId
func (m TradeOrderPointer) SenderLocationId() string {
	return message.StringVLS(m.Ptr, 506, 502)
}

// SenderLocationId
func (m TradeOrderPointerBuilder) SenderLocationId() string {
	return message.StringVLS(m.Ptr, 506, 502)
}

// SelfMatchPreventionID
func (m TradeOrderPointer) SelfMatchPreventionID() string {
	return message.StringVLS(m.Ptr, 510, 506)
}

// SelfMatchPreventionID
func (m TradeOrderPointerBuilder) SelfMatchPreventionID() string {
	return message.StringVLS(m.Ptr, 510, 506)
}

// CTICode
func (m TradeOrderPointer) CTICode() int32 {
	return message.Int32VLS(m.Ptr, 514, 510)
}

// CTICode
func (m TradeOrderPointerBuilder) CTICode() int32 {
	return message.Int32VLS(m.Ptr, 514, 510)
}

// ObtainOrderActionDateTimeFromLastTradeTimeInChart
func (m TradeOrderPointer) ObtainOrderActionDateTimeFromLastTradeTimeInChart() uint8 {
	return message.Uint8VLS(m.Ptr, 515, 514)
}

// ObtainOrderActionDateTimeFromLastTradeTimeInChart
func (m TradeOrderPointerBuilder) ObtainOrderActionDateTimeFromLastTradeTimeInChart() uint8 {
	return message.Uint8VLS(m.Ptr, 515, 514)
}

// MaximumShowQuantity
func (m TradeOrderPointer) MaximumShowQuantity() float64 {
	return message.Float64VLS(m.Ptr, 523, 515)
}

// MaximumShowQuantity
func (m TradeOrderPointerBuilder) MaximumShowQuantity() float64 {
	return message.Float64VLS(m.Ptr, 523, 515)
}

// OrderSubmitted
func (m TradeOrderPointer) OrderSubmitted() uint8 {
	return message.Uint8VLS(m.Ptr, 524, 523)
}

// OrderSubmitted
func (m TradeOrderPointerBuilder) OrderSubmitted() uint8 {
	return message.Uint8VLS(m.Ptr, 524, 523)
}

// IsSnapshot
func (m TradeOrderPointer) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Ptr, 525, 524)
}

// IsSnapshot
func (m TradeOrderPointerBuilder) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Ptr, 525, 524)
}

// IsFirstMessageInBatch
func (m TradeOrderPointer) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 526, 525)
}

// IsFirstMessageInBatch
func (m TradeOrderPointerBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 526, 525)
}

// IsLastMessageInBatch
func (m TradeOrderPointer) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 527, 526)
}

// IsLastMessageInBatch
func (m TradeOrderPointerBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 527, 526)
}

// ExternalLastActionDateTimeUTC
func (m TradeOrderPointer) ExternalLastActionDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 535, 527)
}

// ExternalLastActionDateTimeUTC
func (m TradeOrderPointerBuilder) ExternalLastActionDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 535, 527)
}

// SetIsOrderDeleted
func (m TradeOrderPointerBuilder) SetIsOrderDeleted(value uint8) {
	message.SetUint8VLS(m.Ptr, 7, 6, value)
}

// SetInternalOrderID
func (m TradeOrderPointerBuilder) SetInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Ptr, 15, 7, value)
}

// SetOrderStatusCode
func (m TradeOrderPointerBuilder) SetOrderStatusCode(value uint16) {
	message.SetUint16VLS(m.Ptr, 17, 15, value)
}

// SetOrderStatusBeforePendingModify
func (m TradeOrderPointerBuilder) SetOrderStatusBeforePendingModify(value uint16) {
	message.SetUint16VLS(m.Ptr, 19, 17, value)
}

// SetOrderStatusBeforePendingCancel
func (m TradeOrderPointerBuilder) SetOrderStatusBeforePendingCancel(value uint16) {
	message.SetUint16VLS(m.Ptr, 21, 19, value)
}

// SetServiceOrderID
func (m *TradeOrderPointerBuilder) SetServiceOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 25, 21, value)
}

// SetActualSymbol
func (m *TradeOrderPointerBuilder) SetActualSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 29, 25, value)
}

// SetOrderType
func (m TradeOrderPointerBuilder) SetOrderType(value int32) {
	message.SetInt32VLS(m.Ptr, 33, 29, value)
}

// SetBuySell
func (m TradeOrderPointerBuilder) SetBuySell(value uint16) {
	message.SetUint16VLS(m.Ptr, 35, 33, value)
}

// SetPrice1
func (m TradeOrderPointerBuilder) SetPrice1(value float64) {
	message.SetFloat64VLS(m.Ptr, 43, 35, value)
}

// SetPrice2
func (m TradeOrderPointerBuilder) SetPrice2(value float64) {
	message.SetFloat64VLS(m.Ptr, 51, 43, value)
}

// SetOrderQuantity
func (m TradeOrderPointerBuilder) SetOrderQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 59, 51, value)
}

// SetFilledQuantity
func (m TradeOrderPointerBuilder) SetFilledQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 67, 59, value)
}

// SetAverageFillPrice
func (m TradeOrderPointerBuilder) SetAverageFillPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 75, 67, value)
}

// SetRealtimeFillStatus
func (m TradeOrderPointerBuilder) SetRealtimeFillStatus(value int32) {
	message.SetInt32VLS(m.Ptr, 79, 75, value)
}

// SetIsRestingOrderDuringFill
func (m TradeOrderPointerBuilder) SetIsRestingOrderDuringFill(value uint8) {
	message.SetUint8VLS(m.Ptr, 80, 79, value)
}

// SetOrderRejectType
func (m TradeOrderPointerBuilder) SetOrderRejectType(value int32) {
	message.SetInt32VLS(m.Ptr, 84, 80, value)
}

// SetTradeAccount
func (m *TradeOrderPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 88, 84, value)
}

// SetSubAccountIdentifier
func (m TradeOrderPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Ptr, 92, 88, value)
}

// SetInternalOrderIDModifierForService
func (m TradeOrderPointerBuilder) SetInternalOrderIDModifierForService(value int32) {
	message.SetInt32VLS(m.Ptr, 96, 92, value)
}

// SetFIXClientOrderID
func (m *TradeOrderPointerBuilder) SetFIXClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 100, 96, value)
}

// SetSequenceNumberBasedClientOrderID
func (m TradeOrderPointerBuilder) SetSequenceNumberBasedClientOrderID(value uint32) {
	message.SetUint32VLS(m.Ptr, 104, 100, value)
}

// SetClientOrderIDForDTCServer
func (m *TradeOrderPointerBuilder) SetClientOrderIDForDTCServer(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 108, 104, value)
}

// SetPreviousClientOrderIDForDTCServer
func (m *TradeOrderPointerBuilder) SetPreviousClientOrderIDForDTCServer(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 112, 108, value)
}

// SetExchangeOrderID
func (m *TradeOrderPointerBuilder) SetExchangeOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 116, 112, value)
}

// SetOriginatingClientUsername
func (m *TradeOrderPointerBuilder) SetOriginatingClientUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 120, 116, value)
}

// SetEntryDateTime
func (m TradeOrderPointerBuilder) SetEntryDateTime(value int64) {
	message.SetInt64VLS(m.Ptr, 128, 120, value)
}

// SetLastActionDateTime
func (m TradeOrderPointerBuilder) SetLastActionDateTime(value int64) {
	message.SetInt64VLS(m.Ptr, 136, 128, value)
}

// SetServiceUpdateDateTimeUTC
func (m TradeOrderPointerBuilder) SetServiceUpdateDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 144, 136, value)
}

// SetOrderEntryTimeForService
func (m TradeOrderPointerBuilder) SetOrderEntryTimeForService(value uint32) {
	message.SetUint32VLS(m.Ptr, 148, 144, value)
}

// SetLastModifyTimeForService
func (m TradeOrderPointerBuilder) SetLastModifyTimeForService(value uint32) {
	message.SetUint32VLS(m.Ptr, 152, 148, value)
}

// SetGoodTillDateTime
func (m TradeOrderPointerBuilder) SetGoodTillDateTime(value int64) {
	message.SetInt64VLS(m.Ptr, 160, 152, value)
}

// SetTimeInForce
func (m TradeOrderPointerBuilder) SetTimeInForce(value int32) {
	message.SetInt32VLS(m.Ptr, 164, 160, value)
}

// SetOpenClose
func (m TradeOrderPointerBuilder) SetOpenClose(value uint16) {
	message.SetUint16VLS(m.Ptr, 166, 164, value)
}

// SetTrailStopOffset1
func (m TradeOrderPointerBuilder) SetTrailStopOffset1(value float64) {
	message.SetFloat64VLS(m.Ptr, 174, 166, value)
}

// SetTrailStopStep
func (m TradeOrderPointerBuilder) SetTrailStopStep(value float64) {
	message.SetFloat64VLS(m.Ptr, 182, 174, value)
}

// SetTrailTriggerPrice
func (m TradeOrderPointerBuilder) SetTrailTriggerPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 190, 182, value)
}

// SetTrailingStopTriggerOffset
func (m TradeOrderPointerBuilder) SetTrailingStopTriggerOffset(value float64) {
	message.SetFloat64VLS(m.Ptr, 198, 190, value)
}

// SetTrailTriggerHit
func (m TradeOrderPointerBuilder) SetTrailTriggerHit(value uint8) {
	message.SetUint8VLS(m.Ptr, 199, 198, value)
}

// SetTrailToBreakEvenStopOffset
func (m TradeOrderPointerBuilder) SetTrailToBreakEvenStopOffset(value float64) {
	message.SetFloat64VLS(m.Ptr, 207, 199, value)
}

// SetMaximumChaseAmountAsPrice
func (m TradeOrderPointerBuilder) SetMaximumChaseAmountAsPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 215, 207, value)
}

// SetInitialChaseOrderPrice1
func (m TradeOrderPointerBuilder) SetInitialChaseOrderPrice1(value float64) {
	message.SetFloat64VLS(m.Ptr, 223, 215, value)
}

// SetInitialLastTradePriceForChaseOrders
func (m TradeOrderPointerBuilder) SetInitialLastTradePriceForChaseOrders(value float64) {
	message.SetFloat64VLS(m.Ptr, 231, 223, value)
}

// SetTrailingStopTriggerOCOGroupNumber
func (m TradeOrderPointerBuilder) SetTrailingStopTriggerOCOGroupNumber(value int32) {
	message.SetInt32VLS(m.Ptr, 235, 231, value)
}

// SetLastModifyPrice1
func (m TradeOrderPointerBuilder) SetLastModifyPrice1(value float64) {
	message.SetFloat64VLS(m.Ptr, 243, 235, value)
}

// SetLastModifyQuantity
func (m TradeOrderPointerBuilder) SetLastModifyQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 251, 243, value)
}

// SetCumulativeOrderQuantityFromParentFills
func (m TradeOrderPointerBuilder) SetCumulativeOrderQuantityFromParentFills(value float64) {
	message.SetFloat64VLS(m.Ptr, 259, 251, value)
}

// SetPriorFilledQuantity
func (m TradeOrderPointerBuilder) SetPriorFilledQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 267, 259, value)
}

// SetTickSize
func (m TradeOrderPointerBuilder) SetTickSize(value float32) {
	message.SetFloat32VLS(m.Ptr, 271, 267, value)
}

// SetValueFormat
func (m TradeOrderPointerBuilder) SetValueFormat(value int32) {
	message.SetInt32VLS(m.Ptr, 275, 271, value)
}

// SetPriceMultiplier
func (m TradeOrderPointerBuilder) SetPriceMultiplier(value float32) {
	message.SetFloat32VLS(m.Ptr, 279, 275, value)
}

// SetParentInternalOrderID
func (m TradeOrderPointerBuilder) SetParentInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Ptr, 287, 279, value)
}

// SetTargetChildInternalOrderID
func (m TradeOrderPointerBuilder) SetTargetChildInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Ptr, 295, 287, value)
}

// SetStopChildInternalOrderID
func (m TradeOrderPointerBuilder) SetStopChildInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Ptr, 303, 295, value)
}

// SetAttachedOrderPriceOffset1
func (m TradeOrderPointerBuilder) SetAttachedOrderPriceOffset1(value float64) {
	message.SetFloat64VLS(m.Ptr, 311, 303, value)
}

// SetLinkInternalOrderID
func (m TradeOrderPointerBuilder) SetLinkInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Ptr, 319, 311, value)
}

// SetOCOGroupInternalOrderID
func (m TradeOrderPointerBuilder) SetOCOGroupInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Ptr, 327, 319, value)
}

// SetOCOSiblingInternalOrderID
func (m TradeOrderPointerBuilder) SetOCOSiblingInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Ptr, 335, 327, value)
}

// SetDisableChildAndSiblingRelatedActions
func (m TradeOrderPointerBuilder) SetDisableChildAndSiblingRelatedActions(value uint8) {
	message.SetUint8VLS(m.Ptr, 336, 335, value)
}

// SetOCOManagedByService
func (m TradeOrderPointerBuilder) SetOCOManagedByService(value uint8) {
	message.SetUint8VLS(m.Ptr, 337, 336, value)
}

// SetBracketOrderServerManaged
func (m TradeOrderPointerBuilder) SetBracketOrderServerManaged(value uint8) {
	message.SetUint8VLS(m.Ptr, 338, 337, value)
}

// SetLastOrderActionSource
func (m *TradeOrderPointerBuilder) SetLastOrderActionSource(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 342, 338, value)
}

// SetStopLimitOrderStopPriceTriggered
func (m TradeOrderPointerBuilder) SetStopLimitOrderStopPriceTriggered(value uint8) {
	message.SetUint8VLS(m.Ptr, 343, 342, value)
}

// SetOCOFullSiblingCancelOnPartialFill
func (m TradeOrderPointerBuilder) SetOCOFullSiblingCancelOnPartialFill(value uint8) {
	message.SetUint8VLS(m.Ptr, 344, 343, value)
}

// SetReverseOnCompleteFill
func (m TradeOrderPointerBuilder) SetReverseOnCompleteFill(value uint8) {
	message.SetUint8VLS(m.Ptr, 345, 344, value)
}

// SetSupportScaleIn
func (m TradeOrderPointerBuilder) SetSupportScaleIn(value uint8) {
	message.SetUint8VLS(m.Ptr, 346, 345, value)
}

// SetSupportScaleOut
func (m TradeOrderPointerBuilder) SetSupportScaleOut(value uint8) {
	message.SetUint8VLS(m.Ptr, 347, 346, value)
}

// SetSourceChartNumber
func (m TradeOrderPointerBuilder) SetSourceChartNumber(value int32) {
	message.SetInt32VLS(m.Ptr, 351, 347, value)
}

// SetSourceChartbookFileName
func (m *TradeOrderPointerBuilder) SetSourceChartbookFileName(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 355, 351, value)
}

// SetIsAutomatedOrder
func (m TradeOrderPointerBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8VLS(m.Ptr, 356, 355, value)
}

// SetSimulatedOrder
func (m TradeOrderPointerBuilder) SetSimulatedOrder(value uint8) {
	message.SetUint8VLS(m.Ptr, 357, 356, value)
}

// SetIsChartReplaying
func (m TradeOrderPointerBuilder) SetIsChartReplaying(value uint8) {
	message.SetUint8VLS(m.Ptr, 358, 357, value)
}

// SetAttachedOrderOCOGroupNumber
func (m TradeOrderPointerBuilder) SetAttachedOrderOCOGroupNumber(value int32) {
	message.SetInt32VLS(m.Ptr, 362, 358, value)
}

// SetLastFillExecutionServiceID
func (m *TradeOrderPointerBuilder) SetLastFillExecutionServiceID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 366, 362, value)
}

// SetFillCount
func (m TradeOrderPointerBuilder) SetFillCount(value int32) {
	message.SetInt32VLS(m.Ptr, 370, 366, value)
}

// SetLastFillQuantity
func (m TradeOrderPointerBuilder) SetLastFillQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 378, 370, value)
}

// SetLastFillPrice
func (m TradeOrderPointerBuilder) SetLastFillPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 386, 378, value)
}

// SetLastFillDateTimeUTC
func (m TradeOrderPointerBuilder) SetLastFillDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 394, 386, value)
}

// SetRejectedStopOCOSiblingInternalOrderID
func (m TradeOrderPointerBuilder) SetRejectedStopOCOSiblingInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Ptr, 402, 394, value)
}

// SetRejectedStopReplacementMarketOrderQuantity
func (m TradeOrderPointerBuilder) SetRejectedStopReplacementMarketOrderQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 410, 402, value)
}

// SetEvaluatingForFill
func (m TradeOrderPointerBuilder) SetEvaluatingForFill(value uint8) {
	message.SetUint8VLS(m.Ptr, 411, 410, value)
}

// SetLastProcessedTimeSalesRecordSequenceForPrices
func (m TradeOrderPointerBuilder) SetLastProcessedTimeSalesRecordSequenceForPrices(value uint32) {
	message.SetUint32VLS(m.Ptr, 415, 411, value)
}

// SetIsMarketDataManagementOfOrderEnabled
func (m TradeOrderPointerBuilder) SetIsMarketDataManagementOfOrderEnabled(value uint8) {
	message.SetUint8VLS(m.Ptr, 416, 415, value)
}

// SetTextTag
func (m *TradeOrderPointerBuilder) SetTextTag(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 420, 416, value)
}

// SetTimedOutOrderRequestedStatusDateTime
func (m TradeOrderPointerBuilder) SetTimedOutOrderRequestedStatusDateTime(value int64) {
	message.SetInt64VLS(m.Ptr, 428, 420, value)
}

// SetRequestedStatusForTimedOutOrder
func (m TradeOrderPointerBuilder) SetRequestedStatusForTimedOutOrder(value uint8) {
	message.SetUint8VLS(m.Ptr, 429, 428, value)
}

// SetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled
func (m TradeOrderPointerBuilder) SetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled(value uint8) {
	message.SetUint8VLS(m.Ptr, 430, 429, value)
}

// SetQuantityToIncreaseFromParentFill
func (m TradeOrderPointerBuilder) SetQuantityToIncreaseFromParentFill(value float64) {
	message.SetFloat64VLS(m.Ptr, 438, 430, value)
}

// SetMoveToBreakevenStopReferencePrice
func (m TradeOrderPointerBuilder) SetMoveToBreakevenStopReferencePrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 446, 438, value)
}

// SetQuantityTriggeredStop_QuantityForTrigger
func (m TradeOrderPointerBuilder) SetQuantityTriggeredStop_QuantityForTrigger(value float64) {
	message.SetFloat64VLS(m.Ptr, 454, 446, value)
}

// SetAccumulatedTradeVolumeForTriggeredStop
func (m TradeOrderPointerBuilder) SetAccumulatedTradeVolumeForTriggeredStop(value float64) {
	message.SetFloat64VLS(m.Ptr, 462, 454, value)
}

// SetBidAskQuantityStopInitialTriggerMet
func (m TradeOrderPointerBuilder) SetBidAskQuantityStopInitialTriggerMet(value uint8) {
	message.SetUint8VLS(m.Ptr, 463, 462, value)
}

// SetNeedToOverrideLock
func (m TradeOrderPointerBuilder) SetNeedToOverrideLock(value uint8) {
	message.SetUint8VLS(m.Ptr, 464, 463, value)
}

// SetCurrentMarketPrice
func (m TradeOrderPointerBuilder) SetCurrentMarketPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 472, 464, value)
}

// SetCurrentMarketDateTime
func (m TradeOrderPointerBuilder) SetCurrentMarketDateTime(value int64) {
	message.SetInt64VLS(m.Ptr, 480, 472, value)
}

// SetSupportOrderFillBilling
func (m TradeOrderPointerBuilder) SetSupportOrderFillBilling(value uint8) {
	message.SetUint8VLS(m.Ptr, 481, 480, value)
}

// SetIsBillable
func (m TradeOrderPointerBuilder) SetIsBillable(value uint8) {
	message.SetUint8VLS(m.Ptr, 482, 481, value)
}

// SetQuantityForBilling
func (m TradeOrderPointerBuilder) SetQuantityForBilling(value int32) {
	message.SetInt32VLS(m.Ptr, 486, 482, value)
}

// SetNumberOfFailedOrderModifications
func (m TradeOrderPointerBuilder) SetNumberOfFailedOrderModifications(value uint32) {
	message.SetUint32VLS(m.Ptr, 490, 486, value)
}

// SetDTCServerIndex
func (m TradeOrderPointerBuilder) SetDTCServerIndex(value int32) {
	message.SetInt32VLS(m.Ptr, 494, 490, value)
}

// SetClearingFirmID
func (m *TradeOrderPointerBuilder) SetClearingFirmID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 498, 494, value)
}

// SetSenderSubID
func (m *TradeOrderPointerBuilder) SetSenderSubID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 502, 498, value)
}

// SetSenderLocationId
func (m *TradeOrderPointerBuilder) SetSenderLocationId(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 506, 502, value)
}

// SetSelfMatchPreventionID
func (m *TradeOrderPointerBuilder) SetSelfMatchPreventionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 510, 506, value)
}

// SetCTICode
func (m TradeOrderPointerBuilder) SetCTICode(value int32) {
	message.SetInt32VLS(m.Ptr, 514, 510, value)
}

// SetObtainOrderActionDateTimeFromLastTradeTimeInChart
func (m TradeOrderPointerBuilder) SetObtainOrderActionDateTimeFromLastTradeTimeInChart(value uint8) {
	message.SetUint8VLS(m.Ptr, 515, 514, value)
}

// SetMaximumShowQuantity
func (m TradeOrderPointerBuilder) SetMaximumShowQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 523, 515, value)
}

// SetOrderSubmitted
func (m TradeOrderPointerBuilder) SetOrderSubmitted(value uint8) {
	message.SetUint8VLS(m.Ptr, 524, 523, value)
}

// SetIsSnapshot
func (m TradeOrderPointerBuilder) SetIsSnapshot(value uint8) {
	message.SetUint8VLS(m.Ptr, 525, 524, value)
}

// SetIsFirstMessageInBatch
func (m TradeOrderPointerBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Ptr, 526, 525, value)
}

// SetIsLastMessageInBatch
func (m TradeOrderPointerBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Ptr, 527, 526, value)
}

// SetExternalLastActionDateTimeUTC
func (m TradeOrderPointerBuilder) SetExternalLastActionDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 535, 527, value)
}

// Copy
func (m TradeOrderPointer) Copy(to TradeOrderBuilder) {
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
func (m TradeOrderPointerBuilder) Copy(to TradeOrderBuilder) {
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
func (m TradeOrderPointer) CopyPointer(to TradeOrderPointerBuilder) {
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
func (m TradeOrderPointerBuilder) CopyPointer(to TradeOrderPointerBuilder) {
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
