package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeActivityDataSize = 280

// {
//     Size                             = TradeActivityDataSize  (280)
//     Type                             = TRADE_ACTIVITY_DATA  (10114)
//     BaseSize                         = TradeActivityDataSize  (280)
//     ActivityType                     = 0
//     DataDateTimeUTC                  = 0
//     Symbol                           = ""
//     OrderActionSource                = ""
//     InternalOrderID                  = 0
//     ServiceOrderID                   = ""
//     ExchangeOrderID                  = ""
//     FIXClientOrderID                 = ""
//     OrderTypeName                    = ""
//     Quantity                         = 0
//     BuySell                          = 0
//     Price1                           = 0
//     Price2                           = 0
//     NewOrderStatus                   = 0
//     FillPrice                        = 0.000000
//     OrderFilledQuantity              = 0
//     HighPriceDuringPosition          = 0
//     LowPriceDuringPosition           = 0
//     LastPriceDuringPosition          = 0
//     TradeAccount                     = ""
//     ParentInternalOrderID            = 0
//     OpenClose                        = 0
//     IsSimulated                      = false
//     IsAutomatedOrder                 = false
//     IsChartReplaying                 = false
//     FillExecutionServiceID           = ""
//     PositionQuantity                 = 0
//     SourceChartNumber                = 0
//     SourceChartbookFileName          = ""
//     TimeInForce                      = 0
//     SymbolServiceCode                = ""
//     Note                             = ""
//     OriginatingClientUsername        = ""
//     TradeAccountBalance              = 0
//     SupportsPositionQuantityField    = 0
//     IsBillable                       = false
//     QuantityForBilling               = 0
//     OrderRoutingServiceCode          = ""
//     SubAccountIdentifier             = 0
//     AuditTrail_TransactDateTimeUTC   = 0
//     AuditTrail_MessageDirection      = 0
//     AuditTrail_OperatorID            = ""
//     AuditTrail_SelfMatchPreventionID = ""
//     AuditTrail_SessionID             = ""
//     AuditTrail_ExecutingFirmID       = ""
//     AuditTrail_FixMessageType        = ""
//     AuditTrail_CustomerTypeIndicator = 0
//     AuditTrail_CustomerOrFirm        = 0
//     AuditTrail_ExecutionReportID     = ""
//     AuditTrail_SpreadLegLinkID       = ""
//     AuditTrail_SecurityDesc          = ""
//     AuditTrail_MarketSegmentID       = ""
//     AuditTrail_IFMFlag               = 0
//     AuditTrail_DisplayQuantity       = 0
//     AuditTrail_CountryOfOrigin       = ""
//     AuditTrail_FillQuantity          = 0
//     AuditTrail_RemainingQuantity     = 0
//     AuditTrail_AggressorFlag         = 0
//     AuditTrail_SourceOfCancellation  = 0
//     AuditTrail_OrdRejReason          = ""
//     IsSnapshot                       = false
//     IsFirstMessageInBatch            = false
//     IsLastMessageInBatch             = false
// }
var _TradeActivityDataDefault = []byte{24, 1, 130, 39, 24, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeActivityData struct {
	message.VLS
}

type TradeActivityDataBuilder struct {
	message.VLSBuilder
}

type TradeActivityDataPointer struct {
	message.VLSPointer
}

type TradeActivityDataPointerBuilder struct {
	message.VLSPointerBuilder
}

func NewTradeActivityDataFrom(b []byte) TradeActivityData {
	return TradeActivityData{VLS: message.NewVLSFrom(b)}
}

func WrapTradeActivityData(b []byte) TradeActivityData {
	return TradeActivityData{VLS: message.WrapVLS(b)}
}

func NewTradeActivityData() TradeActivityDataBuilder {
	a := TradeActivityDataBuilder{message.NewVLSBuilder(280)}
	a.Unsafe().SetBytes(0, _TradeActivityDataDefault)
	return a
}

func AllocTradeActivityData() TradeActivityDataPointerBuilder {
	a := TradeActivityDataPointerBuilder{message.AllocVLSBuilder(280)}
	a.Ptr.SetBytes(0, _TradeActivityDataDefault)
	return a
}

func AllocTradeActivityDataFrom(b []byte) TradeActivityDataPointer {
	return TradeActivityDataPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                             = TradeActivityDataSize  (280)
//     Type                             = TRADE_ACTIVITY_DATA  (10114)
//     BaseSize                         = TradeActivityDataSize  (280)
//     ActivityType                     = 0
//     DataDateTimeUTC                  = 0
//     Symbol                           = ""
//     OrderActionSource                = ""
//     InternalOrderID                  = 0
//     ServiceOrderID                   = ""
//     ExchangeOrderID                  = ""
//     FIXClientOrderID                 = ""
//     OrderTypeName                    = ""
//     Quantity                         = 0
//     BuySell                          = 0
//     Price1                           = 0
//     Price2                           = 0
//     NewOrderStatus                   = 0
//     FillPrice                        = 0.000000
//     OrderFilledQuantity              = 0
//     HighPriceDuringPosition          = 0
//     LowPriceDuringPosition           = 0
//     LastPriceDuringPosition          = 0
//     TradeAccount                     = ""
//     ParentInternalOrderID            = 0
//     OpenClose                        = 0
//     IsSimulated                      = false
//     IsAutomatedOrder                 = false
//     IsChartReplaying                 = false
//     FillExecutionServiceID           = ""
//     PositionQuantity                 = 0
//     SourceChartNumber                = 0
//     SourceChartbookFileName          = ""
//     TimeInForce                      = 0
//     SymbolServiceCode                = ""
//     Note                             = ""
//     OriginatingClientUsername        = ""
//     TradeAccountBalance              = 0
//     SupportsPositionQuantityField    = 0
//     IsBillable                       = false
//     QuantityForBilling               = 0
//     OrderRoutingServiceCode          = ""
//     SubAccountIdentifier             = 0
//     AuditTrail_TransactDateTimeUTC   = 0
//     AuditTrail_MessageDirection      = 0
//     AuditTrail_OperatorID            = ""
//     AuditTrail_SelfMatchPreventionID = ""
//     AuditTrail_SessionID             = ""
//     AuditTrail_ExecutingFirmID       = ""
//     AuditTrail_FixMessageType        = ""
//     AuditTrail_CustomerTypeIndicator = 0
//     AuditTrail_CustomerOrFirm        = 0
//     AuditTrail_ExecutionReportID     = ""
//     AuditTrail_SpreadLegLinkID       = ""
//     AuditTrail_SecurityDesc          = ""
//     AuditTrail_MarketSegmentID       = ""
//     AuditTrail_IFMFlag               = 0
//     AuditTrail_DisplayQuantity       = 0
//     AuditTrail_CountryOfOrigin       = ""
//     AuditTrail_FillQuantity          = 0
//     AuditTrail_RemainingQuantity     = 0
//     AuditTrail_AggressorFlag         = 0
//     AuditTrail_SourceOfCancellation  = 0
//     AuditTrail_OrdRejReason          = ""
//     IsSnapshot                       = false
//     IsFirstMessageInBatch            = false
//     IsLastMessageInBatch             = false
// }
func (m TradeActivityDataBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeActivityDataDefault)
}

// Clear
// {
//     Size                             = TradeActivityDataSize  (280)
//     Type                             = TRADE_ACTIVITY_DATA  (10114)
//     BaseSize                         = TradeActivityDataSize  (280)
//     ActivityType                     = 0
//     DataDateTimeUTC                  = 0
//     Symbol                           = ""
//     OrderActionSource                = ""
//     InternalOrderID                  = 0
//     ServiceOrderID                   = ""
//     ExchangeOrderID                  = ""
//     FIXClientOrderID                 = ""
//     OrderTypeName                    = ""
//     Quantity                         = 0
//     BuySell                          = 0
//     Price1                           = 0
//     Price2                           = 0
//     NewOrderStatus                   = 0
//     FillPrice                        = 0.000000
//     OrderFilledQuantity              = 0
//     HighPriceDuringPosition          = 0
//     LowPriceDuringPosition           = 0
//     LastPriceDuringPosition          = 0
//     TradeAccount                     = ""
//     ParentInternalOrderID            = 0
//     OpenClose                        = 0
//     IsSimulated                      = false
//     IsAutomatedOrder                 = false
//     IsChartReplaying                 = false
//     FillExecutionServiceID           = ""
//     PositionQuantity                 = 0
//     SourceChartNumber                = 0
//     SourceChartbookFileName          = ""
//     TimeInForce                      = 0
//     SymbolServiceCode                = ""
//     Note                             = ""
//     OriginatingClientUsername        = ""
//     TradeAccountBalance              = 0
//     SupportsPositionQuantityField    = 0
//     IsBillable                       = false
//     QuantityForBilling               = 0
//     OrderRoutingServiceCode          = ""
//     SubAccountIdentifier             = 0
//     AuditTrail_TransactDateTimeUTC   = 0
//     AuditTrail_MessageDirection      = 0
//     AuditTrail_OperatorID            = ""
//     AuditTrail_SelfMatchPreventionID = ""
//     AuditTrail_SessionID             = ""
//     AuditTrail_ExecutingFirmID       = ""
//     AuditTrail_FixMessageType        = ""
//     AuditTrail_CustomerTypeIndicator = 0
//     AuditTrail_CustomerOrFirm        = 0
//     AuditTrail_ExecutionReportID     = ""
//     AuditTrail_SpreadLegLinkID       = ""
//     AuditTrail_SecurityDesc          = ""
//     AuditTrail_MarketSegmentID       = ""
//     AuditTrail_IFMFlag               = 0
//     AuditTrail_DisplayQuantity       = 0
//     AuditTrail_CountryOfOrigin       = ""
//     AuditTrail_FillQuantity          = 0
//     AuditTrail_RemainingQuantity     = 0
//     AuditTrail_AggressorFlag         = 0
//     AuditTrail_SourceOfCancellation  = 0
//     AuditTrail_OrdRejReason          = ""
//     IsSnapshot                       = false
//     IsFirstMessageInBatch            = false
//     IsLastMessageInBatch             = false
// }
func (m TradeActivityDataPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeActivityDataDefault)
}

// ToBuilder
func (m TradeActivityData) ToBuilder() TradeActivityDataBuilder {
	return TradeActivityDataBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradeActivityDataPointer) ToBuilder() TradeActivityDataPointerBuilder {
	return TradeActivityDataPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m TradeActivityDataBuilder) Finish() TradeActivityData {
	return TradeActivityData{m.VLSBuilder.Finish()}
}

// Finish
func (m *TradeActivityDataPointerBuilder) Finish() TradeActivityDataPointer {
	return TradeActivityDataPointer{m.VLSPointerBuilder.Finish()}
}

// ActivityType
func (m TradeActivityData) ActivityType() uint8 {
	return message.Uint8VLS(m.Unsafe(), 7, 6)
}

// ActivityType
func (m TradeActivityDataBuilder) ActivityType() uint8 {
	return message.Uint8VLS(m.Unsafe(), 7, 6)
}

// ActivityType
func (m TradeActivityDataPointer) ActivityType() uint8 {
	return message.Uint8VLS(m.Ptr, 7, 6)
}

// ActivityType
func (m TradeActivityDataPointerBuilder) ActivityType() uint8 {
	return message.Uint8VLS(m.Ptr, 7, 6)
}

// DataDateTimeUTC
func (m TradeActivityData) DataDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 15, 7)
}

// DataDateTimeUTC
func (m TradeActivityDataBuilder) DataDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 15, 7)
}

// DataDateTimeUTC
func (m TradeActivityDataPointer) DataDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 15, 7)
}

// DataDateTimeUTC
func (m TradeActivityDataPointerBuilder) DataDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 15, 7)
}

// Symbol
func (m TradeActivityData) Symbol() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
}

// Symbol
func (m TradeActivityDataBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
}

// Symbol
func (m TradeActivityDataPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 19, 15)
}

// Symbol
func (m TradeActivityDataPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 19, 15)
}

// OrderActionSource
func (m TradeActivityData) OrderActionSource() string {
	return message.StringVLS(m.Unsafe(), 23, 19)
}

// OrderActionSource
func (m TradeActivityDataBuilder) OrderActionSource() string {
	return message.StringVLS(m.Unsafe(), 23, 19)
}

// OrderActionSource
func (m TradeActivityDataPointer) OrderActionSource() string {
	return message.StringVLS(m.Ptr, 23, 19)
}

// OrderActionSource
func (m TradeActivityDataPointerBuilder) OrderActionSource() string {
	return message.StringVLS(m.Ptr, 23, 19)
}

// InternalOrderID
func (m TradeActivityData) InternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 31, 23)
}

// InternalOrderID
func (m TradeActivityDataBuilder) InternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 31, 23)
}

// InternalOrderID
func (m TradeActivityDataPointer) InternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 31, 23)
}

// InternalOrderID
func (m TradeActivityDataPointerBuilder) InternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 31, 23)
}

// ServiceOrderID
func (m TradeActivityData) ServiceOrderID() string {
	return message.StringVLS(m.Unsafe(), 35, 31)
}

// ServiceOrderID
func (m TradeActivityDataBuilder) ServiceOrderID() string {
	return message.StringVLS(m.Unsafe(), 35, 31)
}

// ServiceOrderID
func (m TradeActivityDataPointer) ServiceOrderID() string {
	return message.StringVLS(m.Ptr, 35, 31)
}

// ServiceOrderID
func (m TradeActivityDataPointerBuilder) ServiceOrderID() string {
	return message.StringVLS(m.Ptr, 35, 31)
}

// ExchangeOrderID
func (m TradeActivityData) ExchangeOrderID() string {
	return message.StringVLS(m.Unsafe(), 39, 35)
}

// ExchangeOrderID
func (m TradeActivityDataBuilder) ExchangeOrderID() string {
	return message.StringVLS(m.Unsafe(), 39, 35)
}

// ExchangeOrderID
func (m TradeActivityDataPointer) ExchangeOrderID() string {
	return message.StringVLS(m.Ptr, 39, 35)
}

// ExchangeOrderID
func (m TradeActivityDataPointerBuilder) ExchangeOrderID() string {
	return message.StringVLS(m.Ptr, 39, 35)
}

// FIXClientOrderID
func (m TradeActivityData) FIXClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 43, 39)
}

// FIXClientOrderID
func (m TradeActivityDataBuilder) FIXClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 43, 39)
}

// FIXClientOrderID
func (m TradeActivityDataPointer) FIXClientOrderID() string {
	return message.StringVLS(m.Ptr, 43, 39)
}

// FIXClientOrderID
func (m TradeActivityDataPointerBuilder) FIXClientOrderID() string {
	return message.StringVLS(m.Ptr, 43, 39)
}

// OrderTypeName
func (m TradeActivityData) OrderTypeName() string {
	return message.StringVLS(m.Unsafe(), 47, 43)
}

// OrderTypeName
func (m TradeActivityDataBuilder) OrderTypeName() string {
	return message.StringVLS(m.Unsafe(), 47, 43)
}

// OrderTypeName
func (m TradeActivityDataPointer) OrderTypeName() string {
	return message.StringVLS(m.Ptr, 47, 43)
}

// OrderTypeName
func (m TradeActivityDataPointerBuilder) OrderTypeName() string {
	return message.StringVLS(m.Ptr, 47, 43)
}

// Quantity
func (m TradeActivityData) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 55, 47)
}

// Quantity
func (m TradeActivityDataBuilder) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 55, 47)
}

// Quantity
func (m TradeActivityDataPointer) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 55, 47)
}

// Quantity
func (m TradeActivityDataPointerBuilder) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 55, 47)
}

// BuySell
func (m TradeActivityData) BuySell() uint8 {
	return message.Uint8VLS(m.Unsafe(), 56, 55)
}

// BuySell
func (m TradeActivityDataBuilder) BuySell() uint8 {
	return message.Uint8VLS(m.Unsafe(), 56, 55)
}

// BuySell
func (m TradeActivityDataPointer) BuySell() uint8 {
	return message.Uint8VLS(m.Ptr, 56, 55)
}

// BuySell
func (m TradeActivityDataPointerBuilder) BuySell() uint8 {
	return message.Uint8VLS(m.Ptr, 56, 55)
}

// Price1
func (m TradeActivityData) Price1() float64 {
	return message.Float64VLS(m.Unsafe(), 64, 56)
}

// Price1
func (m TradeActivityDataBuilder) Price1() float64 {
	return message.Float64VLS(m.Unsafe(), 64, 56)
}

// Price1
func (m TradeActivityDataPointer) Price1() float64 {
	return message.Float64VLS(m.Ptr, 64, 56)
}

// Price1
func (m TradeActivityDataPointerBuilder) Price1() float64 {
	return message.Float64VLS(m.Ptr, 64, 56)
}

// Price2
func (m TradeActivityData) Price2() float64 {
	return message.Float64VLS(m.Unsafe(), 72, 64)
}

// Price2
func (m TradeActivityDataBuilder) Price2() float64 {
	return message.Float64VLS(m.Unsafe(), 72, 64)
}

// Price2
func (m TradeActivityDataPointer) Price2() float64 {
	return message.Float64VLS(m.Ptr, 72, 64)
}

// Price2
func (m TradeActivityDataPointerBuilder) Price2() float64 {
	return message.Float64VLS(m.Ptr, 72, 64)
}

// NewOrderStatus
func (m TradeActivityData) NewOrderStatus() uint8 {
	return message.Uint8VLS(m.Unsafe(), 73, 72)
}

// NewOrderStatus
func (m TradeActivityDataBuilder) NewOrderStatus() uint8 {
	return message.Uint8VLS(m.Unsafe(), 73, 72)
}

// NewOrderStatus
func (m TradeActivityDataPointer) NewOrderStatus() uint8 {
	return message.Uint8VLS(m.Ptr, 73, 72)
}

// NewOrderStatus
func (m TradeActivityDataPointerBuilder) NewOrderStatus() uint8 {
	return message.Uint8VLS(m.Ptr, 73, 72)
}

// FillPrice
func (m TradeActivityData) FillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 81, 73)
}

// FillPrice
func (m TradeActivityDataBuilder) FillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 81, 73)
}

// FillPrice
func (m TradeActivityDataPointer) FillPrice() float64 {
	return message.Float64VLS(m.Ptr, 81, 73)
}

// FillPrice
func (m TradeActivityDataPointerBuilder) FillPrice() float64 {
	return message.Float64VLS(m.Ptr, 81, 73)
}

// OrderFilledQuantity
func (m TradeActivityData) OrderFilledQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 89, 81)
}

// OrderFilledQuantity
func (m TradeActivityDataBuilder) OrderFilledQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 89, 81)
}

// OrderFilledQuantity
func (m TradeActivityDataPointer) OrderFilledQuantity() float64 {
	return message.Float64VLS(m.Ptr, 89, 81)
}

// OrderFilledQuantity
func (m TradeActivityDataPointerBuilder) OrderFilledQuantity() float64 {
	return message.Float64VLS(m.Ptr, 89, 81)
}

// HighPriceDuringPosition
func (m TradeActivityData) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 97, 89)
}

// HighPriceDuringPosition
func (m TradeActivityDataBuilder) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 97, 89)
}

// HighPriceDuringPosition
func (m TradeActivityDataPointer) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 97, 89)
}

// HighPriceDuringPosition
func (m TradeActivityDataPointerBuilder) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 97, 89)
}

// LowPriceDuringPosition
func (m TradeActivityData) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 105, 97)
}

// LowPriceDuringPosition
func (m TradeActivityDataBuilder) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 105, 97)
}

// LowPriceDuringPosition
func (m TradeActivityDataPointer) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 105, 97)
}

// LowPriceDuringPosition
func (m TradeActivityDataPointerBuilder) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 105, 97)
}

// LastPriceDuringPosition
func (m TradeActivityData) LastPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 113, 105)
}

// LastPriceDuringPosition
func (m TradeActivityDataBuilder) LastPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 113, 105)
}

// LastPriceDuringPosition
func (m TradeActivityDataPointer) LastPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 113, 105)
}

// LastPriceDuringPosition
func (m TradeActivityDataPointerBuilder) LastPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 113, 105)
}

// TradeAccount
func (m TradeActivityData) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 117, 113)
}

// TradeAccount
func (m TradeActivityDataBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 117, 113)
}

// TradeAccount
func (m TradeActivityDataPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 117, 113)
}

// TradeAccount
func (m TradeActivityDataPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 117, 113)
}

// ParentInternalOrderID
func (m TradeActivityData) ParentInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 125, 117)
}

// ParentInternalOrderID
func (m TradeActivityDataBuilder) ParentInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 125, 117)
}

// ParentInternalOrderID
func (m TradeActivityDataPointer) ParentInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 125, 117)
}

// ParentInternalOrderID
func (m TradeActivityDataPointerBuilder) ParentInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 125, 117)
}

// OpenClose
func (m TradeActivityData) OpenClose() uint8 {
	return message.Uint8VLS(m.Unsafe(), 126, 125)
}

// OpenClose
func (m TradeActivityDataBuilder) OpenClose() uint8 {
	return message.Uint8VLS(m.Unsafe(), 126, 125)
}

// OpenClose
func (m TradeActivityDataPointer) OpenClose() uint8 {
	return message.Uint8VLS(m.Ptr, 126, 125)
}

// OpenClose
func (m TradeActivityDataPointerBuilder) OpenClose() uint8 {
	return message.Uint8VLS(m.Ptr, 126, 125)
}

// IsSimulated
func (m TradeActivityData) IsSimulated() bool {
	return message.BoolVLS(m.Unsafe(), 127, 126)
}

// IsSimulated
func (m TradeActivityDataBuilder) IsSimulated() bool {
	return message.BoolVLS(m.Unsafe(), 127, 126)
}

// IsSimulated
func (m TradeActivityDataPointer) IsSimulated() bool {
	return message.BoolVLS(m.Ptr, 127, 126)
}

// IsSimulated
func (m TradeActivityDataPointerBuilder) IsSimulated() bool {
	return message.BoolVLS(m.Ptr, 127, 126)
}

// IsAutomatedOrder
func (m TradeActivityData) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Unsafe(), 128, 127)
}

// IsAutomatedOrder
func (m TradeActivityDataBuilder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Unsafe(), 128, 127)
}

// IsAutomatedOrder
func (m TradeActivityDataPointer) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Ptr, 128, 127)
}

// IsAutomatedOrder
func (m TradeActivityDataPointerBuilder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Ptr, 128, 127)
}

// IsChartReplaying
func (m TradeActivityData) IsChartReplaying() bool {
	return message.BoolVLS(m.Unsafe(), 129, 128)
}

// IsChartReplaying
func (m TradeActivityDataBuilder) IsChartReplaying() bool {
	return message.BoolVLS(m.Unsafe(), 129, 128)
}

// IsChartReplaying
func (m TradeActivityDataPointer) IsChartReplaying() bool {
	return message.BoolVLS(m.Ptr, 129, 128)
}

// IsChartReplaying
func (m TradeActivityDataPointerBuilder) IsChartReplaying() bool {
	return message.BoolVLS(m.Ptr, 129, 128)
}

// FillExecutionServiceID
func (m TradeActivityData) FillExecutionServiceID() string {
	return message.StringVLS(m.Unsafe(), 133, 129)
}

// FillExecutionServiceID
func (m TradeActivityDataBuilder) FillExecutionServiceID() string {
	return message.StringVLS(m.Unsafe(), 133, 129)
}

// FillExecutionServiceID
func (m TradeActivityDataPointer) FillExecutionServiceID() string {
	return message.StringVLS(m.Ptr, 133, 129)
}

// FillExecutionServiceID
func (m TradeActivityDataPointerBuilder) FillExecutionServiceID() string {
	return message.StringVLS(m.Ptr, 133, 129)
}

// PositionQuantity
func (m TradeActivityData) PositionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 141, 133)
}

// PositionQuantity
func (m TradeActivityDataBuilder) PositionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 141, 133)
}

// PositionQuantity
func (m TradeActivityDataPointer) PositionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 141, 133)
}

// PositionQuantity
func (m TradeActivityDataPointerBuilder) PositionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 141, 133)
}

// SourceChartNumber
func (m TradeActivityData) SourceChartNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 145, 141)
}

// SourceChartNumber
func (m TradeActivityDataBuilder) SourceChartNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 145, 141)
}

// SourceChartNumber
func (m TradeActivityDataPointer) SourceChartNumber() int32 {
	return message.Int32VLS(m.Ptr, 145, 141)
}

// SourceChartNumber
func (m TradeActivityDataPointerBuilder) SourceChartNumber() int32 {
	return message.Int32VLS(m.Ptr, 145, 141)
}

// SourceChartbookFileName
func (m TradeActivityData) SourceChartbookFileName() string {
	return message.StringVLS(m.Unsafe(), 149, 145)
}

// SourceChartbookFileName
func (m TradeActivityDataBuilder) SourceChartbookFileName() string {
	return message.StringVLS(m.Unsafe(), 149, 145)
}

// SourceChartbookFileName
func (m TradeActivityDataPointer) SourceChartbookFileName() string {
	return message.StringVLS(m.Ptr, 149, 145)
}

// SourceChartbookFileName
func (m TradeActivityDataPointerBuilder) SourceChartbookFileName() string {
	return message.StringVLS(m.Ptr, 149, 145)
}

// TimeInForce
func (m TradeActivityData) TimeInForce() int32 {
	return message.Int32VLS(m.Unsafe(), 153, 149)
}

// TimeInForce
func (m TradeActivityDataBuilder) TimeInForce() int32 {
	return message.Int32VLS(m.Unsafe(), 153, 149)
}

// TimeInForce
func (m TradeActivityDataPointer) TimeInForce() int32 {
	return message.Int32VLS(m.Ptr, 153, 149)
}

// TimeInForce
func (m TradeActivityDataPointerBuilder) TimeInForce() int32 {
	return message.Int32VLS(m.Ptr, 153, 149)
}

// SymbolServiceCode
func (m TradeActivityData) SymbolServiceCode() string {
	return message.StringVLS(m.Unsafe(), 157, 153)
}

// SymbolServiceCode
func (m TradeActivityDataBuilder) SymbolServiceCode() string {
	return message.StringVLS(m.Unsafe(), 157, 153)
}

// SymbolServiceCode
func (m TradeActivityDataPointer) SymbolServiceCode() string {
	return message.StringVLS(m.Ptr, 157, 153)
}

// SymbolServiceCode
func (m TradeActivityDataPointerBuilder) SymbolServiceCode() string {
	return message.StringVLS(m.Ptr, 157, 153)
}

// Note
func (m TradeActivityData) Note() string {
	return message.StringVLS(m.Unsafe(), 161, 157)
}

// Note
func (m TradeActivityDataBuilder) Note() string {
	return message.StringVLS(m.Unsafe(), 161, 157)
}

// Note
func (m TradeActivityDataPointer) Note() string {
	return message.StringVLS(m.Ptr, 161, 157)
}

// Note
func (m TradeActivityDataPointerBuilder) Note() string {
	return message.StringVLS(m.Ptr, 161, 157)
}

// OriginatingClientUsername
func (m TradeActivityData) OriginatingClientUsername() string {
	return message.StringVLS(m.Unsafe(), 165, 161)
}

// OriginatingClientUsername
func (m TradeActivityDataBuilder) OriginatingClientUsername() string {
	return message.StringVLS(m.Unsafe(), 165, 161)
}

// OriginatingClientUsername
func (m TradeActivityDataPointer) OriginatingClientUsername() string {
	return message.StringVLS(m.Ptr, 165, 161)
}

// OriginatingClientUsername
func (m TradeActivityDataPointerBuilder) OriginatingClientUsername() string {
	return message.StringVLS(m.Ptr, 165, 161)
}

// TradeAccountBalance
func (m TradeActivityData) TradeAccountBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 173, 165)
}

// TradeAccountBalance
func (m TradeActivityDataBuilder) TradeAccountBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 173, 165)
}

// TradeAccountBalance
func (m TradeActivityDataPointer) TradeAccountBalance() float64 {
	return message.Float64VLS(m.Ptr, 173, 165)
}

// TradeAccountBalance
func (m TradeActivityDataPointerBuilder) TradeAccountBalance() float64 {
	return message.Float64VLS(m.Ptr, 173, 165)
}

// SupportsPositionQuantityField
func (m TradeActivityData) SupportsPositionQuantityField() uint8 {
	return message.Uint8VLS(m.Unsafe(), 174, 173)
}

// SupportsPositionQuantityField
func (m TradeActivityDataBuilder) SupportsPositionQuantityField() uint8 {
	return message.Uint8VLS(m.Unsafe(), 174, 173)
}

// SupportsPositionQuantityField
func (m TradeActivityDataPointer) SupportsPositionQuantityField() uint8 {
	return message.Uint8VLS(m.Ptr, 174, 173)
}

// SupportsPositionQuantityField
func (m TradeActivityDataPointerBuilder) SupportsPositionQuantityField() uint8 {
	return message.Uint8VLS(m.Ptr, 174, 173)
}

// IsBillable
func (m TradeActivityData) IsBillable() bool {
	return message.BoolVLS(m.Unsafe(), 175, 174)
}

// IsBillable
func (m TradeActivityDataBuilder) IsBillable() bool {
	return message.BoolVLS(m.Unsafe(), 175, 174)
}

// IsBillable
func (m TradeActivityDataPointer) IsBillable() bool {
	return message.BoolVLS(m.Ptr, 175, 174)
}

// IsBillable
func (m TradeActivityDataPointerBuilder) IsBillable() bool {
	return message.BoolVLS(m.Ptr, 175, 174)
}

// QuantityForBilling
func (m TradeActivityData) QuantityForBilling() int32 {
	return message.Int32VLS(m.Unsafe(), 179, 175)
}

// QuantityForBilling
func (m TradeActivityDataBuilder) QuantityForBilling() int32 {
	return message.Int32VLS(m.Unsafe(), 179, 175)
}

// QuantityForBilling
func (m TradeActivityDataPointer) QuantityForBilling() int32 {
	return message.Int32VLS(m.Ptr, 179, 175)
}

// QuantityForBilling
func (m TradeActivityDataPointerBuilder) QuantityForBilling() int32 {
	return message.Int32VLS(m.Ptr, 179, 175)
}

// OrderRoutingServiceCode
func (m TradeActivityData) OrderRoutingServiceCode() string {
	return message.StringVLS(m.Unsafe(), 183, 179)
}

// OrderRoutingServiceCode
func (m TradeActivityDataBuilder) OrderRoutingServiceCode() string {
	return message.StringVLS(m.Unsafe(), 183, 179)
}

// OrderRoutingServiceCode
func (m TradeActivityDataPointer) OrderRoutingServiceCode() string {
	return message.StringVLS(m.Ptr, 183, 179)
}

// OrderRoutingServiceCode
func (m TradeActivityDataPointerBuilder) OrderRoutingServiceCode() string {
	return message.StringVLS(m.Ptr, 183, 179)
}

// SubAccountIdentifier
func (m TradeActivityData) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 187, 183)
}

// SubAccountIdentifier
func (m TradeActivityDataBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 187, 183)
}

// SubAccountIdentifier
func (m TradeActivityDataPointer) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 187, 183)
}

// SubAccountIdentifier
func (m TradeActivityDataPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 187, 183)
}

// AuditTrail_TransactDateTimeUTC
func (m TradeActivityData) AuditTrail_TransactDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 195, 187)
}

// AuditTrail_TransactDateTimeUTC
func (m TradeActivityDataBuilder) AuditTrail_TransactDateTimeUTC() int64 {
	return message.Int64VLS(m.Unsafe(), 195, 187)
}

// AuditTrail_TransactDateTimeUTC
func (m TradeActivityDataPointer) AuditTrail_TransactDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 195, 187)
}

// AuditTrail_TransactDateTimeUTC
func (m TradeActivityDataPointerBuilder) AuditTrail_TransactDateTimeUTC() int64 {
	return message.Int64VLS(m.Ptr, 195, 187)
}

// AuditTrail_MessageDirection
func (m TradeActivityData) AuditTrail_MessageDirection() int32 {
	return message.Int32VLS(m.Unsafe(), 199, 195)
}

// AuditTrail_MessageDirection
func (m TradeActivityDataBuilder) AuditTrail_MessageDirection() int32 {
	return message.Int32VLS(m.Unsafe(), 199, 195)
}

// AuditTrail_MessageDirection
func (m TradeActivityDataPointer) AuditTrail_MessageDirection() int32 {
	return message.Int32VLS(m.Ptr, 199, 195)
}

// AuditTrail_MessageDirection
func (m TradeActivityDataPointerBuilder) AuditTrail_MessageDirection() int32 {
	return message.Int32VLS(m.Ptr, 199, 195)
}

// AuditTrail_OperatorID
func (m TradeActivityData) AuditTrail_OperatorID() string {
	return message.StringVLS(m.Unsafe(), 203, 199)
}

// AuditTrail_OperatorID
func (m TradeActivityDataBuilder) AuditTrail_OperatorID() string {
	return message.StringVLS(m.Unsafe(), 203, 199)
}

// AuditTrail_OperatorID
func (m TradeActivityDataPointer) AuditTrail_OperatorID() string {
	return message.StringVLS(m.Ptr, 203, 199)
}

// AuditTrail_OperatorID
func (m TradeActivityDataPointerBuilder) AuditTrail_OperatorID() string {
	return message.StringVLS(m.Ptr, 203, 199)
}

// AuditTrail_SelfMatchPreventionID
func (m TradeActivityData) AuditTrail_SelfMatchPreventionID() string {
	return message.StringVLS(m.Unsafe(), 207, 203)
}

// AuditTrail_SelfMatchPreventionID
func (m TradeActivityDataBuilder) AuditTrail_SelfMatchPreventionID() string {
	return message.StringVLS(m.Unsafe(), 207, 203)
}

// AuditTrail_SelfMatchPreventionID
func (m TradeActivityDataPointer) AuditTrail_SelfMatchPreventionID() string {
	return message.StringVLS(m.Ptr, 207, 203)
}

// AuditTrail_SelfMatchPreventionID
func (m TradeActivityDataPointerBuilder) AuditTrail_SelfMatchPreventionID() string {
	return message.StringVLS(m.Ptr, 207, 203)
}

// AuditTrail_SessionID
func (m TradeActivityData) AuditTrail_SessionID() string {
	return message.StringVLS(m.Unsafe(), 211, 207)
}

// AuditTrail_SessionID
func (m TradeActivityDataBuilder) AuditTrail_SessionID() string {
	return message.StringVLS(m.Unsafe(), 211, 207)
}

// AuditTrail_SessionID
func (m TradeActivityDataPointer) AuditTrail_SessionID() string {
	return message.StringVLS(m.Ptr, 211, 207)
}

// AuditTrail_SessionID
func (m TradeActivityDataPointerBuilder) AuditTrail_SessionID() string {
	return message.StringVLS(m.Ptr, 211, 207)
}

// AuditTrail_ExecutingFirmID
func (m TradeActivityData) AuditTrail_ExecutingFirmID() string {
	return message.StringVLS(m.Unsafe(), 215, 211)
}

// AuditTrail_ExecutingFirmID
func (m TradeActivityDataBuilder) AuditTrail_ExecutingFirmID() string {
	return message.StringVLS(m.Unsafe(), 215, 211)
}

// AuditTrail_ExecutingFirmID
func (m TradeActivityDataPointer) AuditTrail_ExecutingFirmID() string {
	return message.StringVLS(m.Ptr, 215, 211)
}

// AuditTrail_ExecutingFirmID
func (m TradeActivityDataPointerBuilder) AuditTrail_ExecutingFirmID() string {
	return message.StringVLS(m.Ptr, 215, 211)
}

// AuditTrail_FixMessageType
func (m TradeActivityData) AuditTrail_FixMessageType() string {
	return message.StringVLS(m.Unsafe(), 219, 215)
}

// AuditTrail_FixMessageType
func (m TradeActivityDataBuilder) AuditTrail_FixMessageType() string {
	return message.StringVLS(m.Unsafe(), 219, 215)
}

// AuditTrail_FixMessageType
func (m TradeActivityDataPointer) AuditTrail_FixMessageType() string {
	return message.StringVLS(m.Ptr, 219, 215)
}

// AuditTrail_FixMessageType
func (m TradeActivityDataPointerBuilder) AuditTrail_FixMessageType() string {
	return message.StringVLS(m.Ptr, 219, 215)
}

// AuditTrail_CustomerTypeIndicator
func (m TradeActivityData) AuditTrail_CustomerTypeIndicator() int16 {
	return message.Int16VLS(m.Unsafe(), 221, 219)
}

// AuditTrail_CustomerTypeIndicator
func (m TradeActivityDataBuilder) AuditTrail_CustomerTypeIndicator() int16 {
	return message.Int16VLS(m.Unsafe(), 221, 219)
}

// AuditTrail_CustomerTypeIndicator
func (m TradeActivityDataPointer) AuditTrail_CustomerTypeIndicator() int16 {
	return message.Int16VLS(m.Ptr, 221, 219)
}

// AuditTrail_CustomerTypeIndicator
func (m TradeActivityDataPointerBuilder) AuditTrail_CustomerTypeIndicator() int16 {
	return message.Int16VLS(m.Ptr, 221, 219)
}

// AuditTrail_CustomerOrFirm
func (m TradeActivityData) AuditTrail_CustomerOrFirm() int16 {
	return message.Int16VLS(m.Unsafe(), 223, 221)
}

// AuditTrail_CustomerOrFirm
func (m TradeActivityDataBuilder) AuditTrail_CustomerOrFirm() int16 {
	return message.Int16VLS(m.Unsafe(), 223, 221)
}

// AuditTrail_CustomerOrFirm
func (m TradeActivityDataPointer) AuditTrail_CustomerOrFirm() int16 {
	return message.Int16VLS(m.Ptr, 223, 221)
}

// AuditTrail_CustomerOrFirm
func (m TradeActivityDataPointerBuilder) AuditTrail_CustomerOrFirm() int16 {
	return message.Int16VLS(m.Ptr, 223, 221)
}

// AuditTrail_ExecutionReportID
func (m TradeActivityData) AuditTrail_ExecutionReportID() string {
	return message.StringVLS(m.Unsafe(), 227, 223)
}

// AuditTrail_ExecutionReportID
func (m TradeActivityDataBuilder) AuditTrail_ExecutionReportID() string {
	return message.StringVLS(m.Unsafe(), 227, 223)
}

// AuditTrail_ExecutionReportID
func (m TradeActivityDataPointer) AuditTrail_ExecutionReportID() string {
	return message.StringVLS(m.Ptr, 227, 223)
}

// AuditTrail_ExecutionReportID
func (m TradeActivityDataPointerBuilder) AuditTrail_ExecutionReportID() string {
	return message.StringVLS(m.Ptr, 227, 223)
}

// AuditTrail_SpreadLegLinkID
func (m TradeActivityData) AuditTrail_SpreadLegLinkID() string {
	return message.StringVLS(m.Unsafe(), 231, 227)
}

// AuditTrail_SpreadLegLinkID
func (m TradeActivityDataBuilder) AuditTrail_SpreadLegLinkID() string {
	return message.StringVLS(m.Unsafe(), 231, 227)
}

// AuditTrail_SpreadLegLinkID
func (m TradeActivityDataPointer) AuditTrail_SpreadLegLinkID() string {
	return message.StringVLS(m.Ptr, 231, 227)
}

// AuditTrail_SpreadLegLinkID
func (m TradeActivityDataPointerBuilder) AuditTrail_SpreadLegLinkID() string {
	return message.StringVLS(m.Ptr, 231, 227)
}

// AuditTrail_SecurityDesc
func (m TradeActivityData) AuditTrail_SecurityDesc() string {
	return message.StringVLS(m.Unsafe(), 235, 231)
}

// AuditTrail_SecurityDesc
func (m TradeActivityDataBuilder) AuditTrail_SecurityDesc() string {
	return message.StringVLS(m.Unsafe(), 235, 231)
}

// AuditTrail_SecurityDesc
func (m TradeActivityDataPointer) AuditTrail_SecurityDesc() string {
	return message.StringVLS(m.Ptr, 235, 231)
}

// AuditTrail_SecurityDesc
func (m TradeActivityDataPointerBuilder) AuditTrail_SecurityDesc() string {
	return message.StringVLS(m.Ptr, 235, 231)
}

// AuditTrail_MarketSegmentID
func (m TradeActivityData) AuditTrail_MarketSegmentID() string {
	return message.StringVLS(m.Unsafe(), 239, 235)
}

// AuditTrail_MarketSegmentID
func (m TradeActivityDataBuilder) AuditTrail_MarketSegmentID() string {
	return message.StringVLS(m.Unsafe(), 239, 235)
}

// AuditTrail_MarketSegmentID
func (m TradeActivityDataPointer) AuditTrail_MarketSegmentID() string {
	return message.StringVLS(m.Ptr, 239, 235)
}

// AuditTrail_MarketSegmentID
func (m TradeActivityDataPointerBuilder) AuditTrail_MarketSegmentID() string {
	return message.StringVLS(m.Ptr, 239, 235)
}

// AuditTrail_IFMFlag
func (m TradeActivityData) AuditTrail_IFMFlag() uint8 {
	return message.Uint8VLS(m.Unsafe(), 240, 239)
}

// AuditTrail_IFMFlag
func (m TradeActivityDataBuilder) AuditTrail_IFMFlag() uint8 {
	return message.Uint8VLS(m.Unsafe(), 240, 239)
}

// AuditTrail_IFMFlag
func (m TradeActivityDataPointer) AuditTrail_IFMFlag() uint8 {
	return message.Uint8VLS(m.Ptr, 240, 239)
}

// AuditTrail_IFMFlag
func (m TradeActivityDataPointerBuilder) AuditTrail_IFMFlag() uint8 {
	return message.Uint8VLS(m.Ptr, 240, 239)
}

// AuditTrail_DisplayQuantity
func (m TradeActivityData) AuditTrail_DisplayQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 248, 240)
}

// AuditTrail_DisplayQuantity
func (m TradeActivityDataBuilder) AuditTrail_DisplayQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 248, 240)
}

// AuditTrail_DisplayQuantity
func (m TradeActivityDataPointer) AuditTrail_DisplayQuantity() float64 {
	return message.Float64VLS(m.Ptr, 248, 240)
}

// AuditTrail_DisplayQuantity
func (m TradeActivityDataPointerBuilder) AuditTrail_DisplayQuantity() float64 {
	return message.Float64VLS(m.Ptr, 248, 240)
}

// AuditTrail_CountryOfOrigin
func (m TradeActivityData) AuditTrail_CountryOfOrigin() string {
	return message.StringVLS(m.Unsafe(), 252, 248)
}

// AuditTrail_CountryOfOrigin
func (m TradeActivityDataBuilder) AuditTrail_CountryOfOrigin() string {
	return message.StringVLS(m.Unsafe(), 252, 248)
}

// AuditTrail_CountryOfOrigin
func (m TradeActivityDataPointer) AuditTrail_CountryOfOrigin() string {
	return message.StringVLS(m.Ptr, 252, 248)
}

// AuditTrail_CountryOfOrigin
func (m TradeActivityDataPointerBuilder) AuditTrail_CountryOfOrigin() string {
	return message.StringVLS(m.Ptr, 252, 248)
}

// AuditTrail_FillQuantity
func (m TradeActivityData) AuditTrail_FillQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 260, 252)
}

// AuditTrail_FillQuantity
func (m TradeActivityDataBuilder) AuditTrail_FillQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 260, 252)
}

// AuditTrail_FillQuantity
func (m TradeActivityDataPointer) AuditTrail_FillQuantity() float64 {
	return message.Float64VLS(m.Ptr, 260, 252)
}

// AuditTrail_FillQuantity
func (m TradeActivityDataPointerBuilder) AuditTrail_FillQuantity() float64 {
	return message.Float64VLS(m.Ptr, 260, 252)
}

// AuditTrail_RemainingQuantity
func (m TradeActivityData) AuditTrail_RemainingQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 268, 260)
}

// AuditTrail_RemainingQuantity
func (m TradeActivityDataBuilder) AuditTrail_RemainingQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 268, 260)
}

// AuditTrail_RemainingQuantity
func (m TradeActivityDataPointer) AuditTrail_RemainingQuantity() float64 {
	return message.Float64VLS(m.Ptr, 268, 260)
}

// AuditTrail_RemainingQuantity
func (m TradeActivityDataPointerBuilder) AuditTrail_RemainingQuantity() float64 {
	return message.Float64VLS(m.Ptr, 268, 260)
}

// AuditTrail_AggressorFlag
func (m TradeActivityData) AuditTrail_AggressorFlag() uint8 {
	return message.Uint8VLS(m.Unsafe(), 269, 268)
}

// AuditTrail_AggressorFlag
func (m TradeActivityDataBuilder) AuditTrail_AggressorFlag() uint8 {
	return message.Uint8VLS(m.Unsafe(), 269, 268)
}

// AuditTrail_AggressorFlag
func (m TradeActivityDataPointer) AuditTrail_AggressorFlag() uint8 {
	return message.Uint8VLS(m.Ptr, 269, 268)
}

// AuditTrail_AggressorFlag
func (m TradeActivityDataPointerBuilder) AuditTrail_AggressorFlag() uint8 {
	return message.Uint8VLS(m.Ptr, 269, 268)
}

// AuditTrail_SourceOfCancellation
func (m TradeActivityData) AuditTrail_SourceOfCancellation() int32 {
	return message.Int32VLS(m.Unsafe(), 273, 269)
}

// AuditTrail_SourceOfCancellation
func (m TradeActivityDataBuilder) AuditTrail_SourceOfCancellation() int32 {
	return message.Int32VLS(m.Unsafe(), 273, 269)
}

// AuditTrail_SourceOfCancellation
func (m TradeActivityDataPointer) AuditTrail_SourceOfCancellation() int32 {
	return message.Int32VLS(m.Ptr, 273, 269)
}

// AuditTrail_SourceOfCancellation
func (m TradeActivityDataPointerBuilder) AuditTrail_SourceOfCancellation() int32 {
	return message.Int32VLS(m.Ptr, 273, 269)
}

// AuditTrail_OrdRejReason
func (m TradeActivityData) AuditTrail_OrdRejReason() string {
	return message.StringVLS(m.Unsafe(), 277, 273)
}

// AuditTrail_OrdRejReason
func (m TradeActivityDataBuilder) AuditTrail_OrdRejReason() string {
	return message.StringVLS(m.Unsafe(), 277, 273)
}

// AuditTrail_OrdRejReason
func (m TradeActivityDataPointer) AuditTrail_OrdRejReason() string {
	return message.StringVLS(m.Ptr, 277, 273)
}

// AuditTrail_OrdRejReason
func (m TradeActivityDataPointerBuilder) AuditTrail_OrdRejReason() string {
	return message.StringVLS(m.Ptr, 277, 273)
}

// IsSnapshot
func (m TradeActivityData) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 278, 277)
}

// IsSnapshot
func (m TradeActivityDataBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 278, 277)
}

// IsSnapshot
func (m TradeActivityDataPointer) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 278, 277)
}

// IsSnapshot
func (m TradeActivityDataPointerBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 278, 277)
}

// IsFirstMessageInBatch
func (m TradeActivityData) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 279, 278)
}

// IsFirstMessageInBatch
func (m TradeActivityDataBuilder) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 279, 278)
}

// IsFirstMessageInBatch
func (m TradeActivityDataPointer) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 279, 278)
}

// IsFirstMessageInBatch
func (m TradeActivityDataPointerBuilder) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 279, 278)
}

// IsLastMessageInBatch
func (m TradeActivityData) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 280, 279)
}

// IsLastMessageInBatch
func (m TradeActivityDataBuilder) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 280, 279)
}

// IsLastMessageInBatch
func (m TradeActivityDataPointer) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 280, 279)
}

// IsLastMessageInBatch
func (m TradeActivityDataPointerBuilder) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 280, 279)
}

// SetActivityType
func (m TradeActivityDataBuilder) SetActivityType(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 7, 6, value)
}

// SetActivityType
func (m TradeActivityDataPointerBuilder) SetActivityType(value uint8) {
	message.SetUint8VLS(m.Ptr, 7, 6, value)
}

// SetDataDateTimeUTC
func (m TradeActivityDataBuilder) SetDataDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Unsafe(), 15, 7, value)
}

// SetDataDateTimeUTC
func (m TradeActivityDataPointerBuilder) SetDataDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 15, 7, value)
}

// SetSymbol
func (m *TradeActivityDataBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 19, 15, value)
}

// SetSymbol
func (m *TradeActivityDataPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 19, 15, value)
}

// SetOrderActionSource
func (m *TradeActivityDataBuilder) SetOrderActionSource(value string) {
	message.SetStringVLS(&m.VLSBuilder, 23, 19, value)
}

// SetOrderActionSource
func (m *TradeActivityDataPointerBuilder) SetOrderActionSource(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 23, 19, value)
}

// SetInternalOrderID
func (m TradeActivityDataBuilder) SetInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 31, 23, value)
}

// SetInternalOrderID
func (m TradeActivityDataPointerBuilder) SetInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Ptr, 31, 23, value)
}

// SetServiceOrderID
func (m *TradeActivityDataBuilder) SetServiceOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 35, 31, value)
}

// SetServiceOrderID
func (m *TradeActivityDataPointerBuilder) SetServiceOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 35, 31, value)
}

// SetExchangeOrderID
func (m *TradeActivityDataBuilder) SetExchangeOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 39, 35, value)
}

// SetExchangeOrderID
func (m *TradeActivityDataPointerBuilder) SetExchangeOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 39, 35, value)
}

// SetFIXClientOrderID
func (m *TradeActivityDataBuilder) SetFIXClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 43, 39, value)
}

// SetFIXClientOrderID
func (m *TradeActivityDataPointerBuilder) SetFIXClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 43, 39, value)
}

// SetOrderTypeName
func (m *TradeActivityDataBuilder) SetOrderTypeName(value string) {
	message.SetStringVLS(&m.VLSBuilder, 47, 43, value)
}

// SetOrderTypeName
func (m *TradeActivityDataPointerBuilder) SetOrderTypeName(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 47, 43, value)
}

// SetQuantity
func (m TradeActivityDataBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 55, 47, value)
}

// SetQuantity
func (m TradeActivityDataPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 55, 47, value)
}

// SetBuySell
func (m TradeActivityDataBuilder) SetBuySell(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 56, 55, value)
}

// SetBuySell
func (m TradeActivityDataPointerBuilder) SetBuySell(value uint8) {
	message.SetUint8VLS(m.Ptr, 56, 55, value)
}

// SetPrice1
func (m TradeActivityDataBuilder) SetPrice1(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 64, 56, value)
}

// SetPrice1
func (m TradeActivityDataPointerBuilder) SetPrice1(value float64) {
	message.SetFloat64VLS(m.Ptr, 64, 56, value)
}

// SetPrice2
func (m TradeActivityDataBuilder) SetPrice2(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 72, 64, value)
}

// SetPrice2
func (m TradeActivityDataPointerBuilder) SetPrice2(value float64) {
	message.SetFloat64VLS(m.Ptr, 72, 64, value)
}

// SetNewOrderStatus
func (m TradeActivityDataBuilder) SetNewOrderStatus(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 73, 72, value)
}

// SetNewOrderStatus
func (m TradeActivityDataPointerBuilder) SetNewOrderStatus(value uint8) {
	message.SetUint8VLS(m.Ptr, 73, 72, value)
}

// SetFillPrice
func (m TradeActivityDataBuilder) SetFillPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 81, 73, value)
}

// SetFillPrice
func (m TradeActivityDataPointerBuilder) SetFillPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 81, 73, value)
}

// SetOrderFilledQuantity
func (m TradeActivityDataBuilder) SetOrderFilledQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 89, 81, value)
}

// SetOrderFilledQuantity
func (m TradeActivityDataPointerBuilder) SetOrderFilledQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 89, 81, value)
}

// SetHighPriceDuringPosition
func (m TradeActivityDataBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 97, 89, value)
}

// SetHighPriceDuringPosition
func (m TradeActivityDataPointerBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 97, 89, value)
}

// SetLowPriceDuringPosition
func (m TradeActivityDataBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 105, 97, value)
}

// SetLowPriceDuringPosition
func (m TradeActivityDataPointerBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 105, 97, value)
}

// SetLastPriceDuringPosition
func (m TradeActivityDataBuilder) SetLastPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 113, 105, value)
}

// SetLastPriceDuringPosition
func (m TradeActivityDataPointerBuilder) SetLastPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 113, 105, value)
}

// SetTradeAccount
func (m *TradeActivityDataBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 117, 113, value)
}

// SetTradeAccount
func (m *TradeActivityDataPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 117, 113, value)
}

// SetParentInternalOrderID
func (m TradeActivityDataBuilder) SetParentInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 125, 117, value)
}

// SetParentInternalOrderID
func (m TradeActivityDataPointerBuilder) SetParentInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Ptr, 125, 117, value)
}

// SetOpenClose
func (m TradeActivityDataBuilder) SetOpenClose(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 126, 125, value)
}

// SetOpenClose
func (m TradeActivityDataPointerBuilder) SetOpenClose(value uint8) {
	message.SetUint8VLS(m.Ptr, 126, 125, value)
}

// SetIsSimulated
func (m TradeActivityDataBuilder) SetIsSimulated(value bool) {
	message.SetBoolVLS(m.Unsafe(), 127, 126, value)
}

// SetIsSimulated
func (m TradeActivityDataPointerBuilder) SetIsSimulated(value bool) {
	message.SetBoolVLS(m.Ptr, 127, 126, value)
}

// SetIsAutomatedOrder
func (m TradeActivityDataBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolVLS(m.Unsafe(), 128, 127, value)
}

// SetIsAutomatedOrder
func (m TradeActivityDataPointerBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolVLS(m.Ptr, 128, 127, value)
}

// SetIsChartReplaying
func (m TradeActivityDataBuilder) SetIsChartReplaying(value bool) {
	message.SetBoolVLS(m.Unsafe(), 129, 128, value)
}

// SetIsChartReplaying
func (m TradeActivityDataPointerBuilder) SetIsChartReplaying(value bool) {
	message.SetBoolVLS(m.Ptr, 129, 128, value)
}

// SetFillExecutionServiceID
func (m *TradeActivityDataBuilder) SetFillExecutionServiceID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 133, 129, value)
}

// SetFillExecutionServiceID
func (m *TradeActivityDataPointerBuilder) SetFillExecutionServiceID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 133, 129, value)
}

// SetPositionQuantity
func (m TradeActivityDataBuilder) SetPositionQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 141, 133, value)
}

// SetPositionQuantity
func (m TradeActivityDataPointerBuilder) SetPositionQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 141, 133, value)
}

// SetSourceChartNumber
func (m TradeActivityDataBuilder) SetSourceChartNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 145, 141, value)
}

// SetSourceChartNumber
func (m TradeActivityDataPointerBuilder) SetSourceChartNumber(value int32) {
	message.SetInt32VLS(m.Ptr, 145, 141, value)
}

// SetSourceChartbookFileName
func (m *TradeActivityDataBuilder) SetSourceChartbookFileName(value string) {
	message.SetStringVLS(&m.VLSBuilder, 149, 145, value)
}

// SetSourceChartbookFileName
func (m *TradeActivityDataPointerBuilder) SetSourceChartbookFileName(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 149, 145, value)
}

// SetTimeInForce
func (m TradeActivityDataBuilder) SetTimeInForce(value int32) {
	message.SetInt32VLS(m.Unsafe(), 153, 149, value)
}

// SetTimeInForce
func (m TradeActivityDataPointerBuilder) SetTimeInForce(value int32) {
	message.SetInt32VLS(m.Ptr, 153, 149, value)
}

// SetSymbolServiceCode
func (m *TradeActivityDataBuilder) SetSymbolServiceCode(value string) {
	message.SetStringVLS(&m.VLSBuilder, 157, 153, value)
}

// SetSymbolServiceCode
func (m *TradeActivityDataPointerBuilder) SetSymbolServiceCode(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 157, 153, value)
}

// SetNote
func (m *TradeActivityDataBuilder) SetNote(value string) {
	message.SetStringVLS(&m.VLSBuilder, 161, 157, value)
}

// SetNote
func (m *TradeActivityDataPointerBuilder) SetNote(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 161, 157, value)
}

// SetOriginatingClientUsername
func (m *TradeActivityDataBuilder) SetOriginatingClientUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 165, 161, value)
}

// SetOriginatingClientUsername
func (m *TradeActivityDataPointerBuilder) SetOriginatingClientUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 165, 161, value)
}

// SetTradeAccountBalance
func (m TradeActivityDataBuilder) SetTradeAccountBalance(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 173, 165, value)
}

// SetTradeAccountBalance
func (m TradeActivityDataPointerBuilder) SetTradeAccountBalance(value float64) {
	message.SetFloat64VLS(m.Ptr, 173, 165, value)
}

// SetSupportsPositionQuantityField
func (m TradeActivityDataBuilder) SetSupportsPositionQuantityField(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 174, 173, value)
}

// SetSupportsPositionQuantityField
func (m TradeActivityDataPointerBuilder) SetSupportsPositionQuantityField(value uint8) {
	message.SetUint8VLS(m.Ptr, 174, 173, value)
}

// SetIsBillable
func (m TradeActivityDataBuilder) SetIsBillable(value bool) {
	message.SetBoolVLS(m.Unsafe(), 175, 174, value)
}

// SetIsBillable
func (m TradeActivityDataPointerBuilder) SetIsBillable(value bool) {
	message.SetBoolVLS(m.Ptr, 175, 174, value)
}

// SetQuantityForBilling
func (m TradeActivityDataBuilder) SetQuantityForBilling(value int32) {
	message.SetInt32VLS(m.Unsafe(), 179, 175, value)
}

// SetQuantityForBilling
func (m TradeActivityDataPointerBuilder) SetQuantityForBilling(value int32) {
	message.SetInt32VLS(m.Ptr, 179, 175, value)
}

// SetOrderRoutingServiceCode
func (m *TradeActivityDataBuilder) SetOrderRoutingServiceCode(value string) {
	message.SetStringVLS(&m.VLSBuilder, 183, 179, value)
}

// SetOrderRoutingServiceCode
func (m *TradeActivityDataPointerBuilder) SetOrderRoutingServiceCode(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 183, 179, value)
}

// SetSubAccountIdentifier
func (m TradeActivityDataBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 187, 183, value)
}

// SetSubAccountIdentifier
func (m TradeActivityDataPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Ptr, 187, 183, value)
}

// SetAuditTrail_TransactDateTimeUTC
func (m TradeActivityDataBuilder) SetAuditTrail_TransactDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Unsafe(), 195, 187, value)
}

// SetAuditTrail_TransactDateTimeUTC
func (m TradeActivityDataPointerBuilder) SetAuditTrail_TransactDateTimeUTC(value int64) {
	message.SetInt64VLS(m.Ptr, 195, 187, value)
}

// SetAuditTrail_MessageDirection
func (m TradeActivityDataBuilder) SetAuditTrail_MessageDirection(value int32) {
	message.SetInt32VLS(m.Unsafe(), 199, 195, value)
}

// SetAuditTrail_MessageDirection
func (m TradeActivityDataPointerBuilder) SetAuditTrail_MessageDirection(value int32) {
	message.SetInt32VLS(m.Ptr, 199, 195, value)
}

// SetAuditTrail_OperatorID
func (m *TradeActivityDataBuilder) SetAuditTrail_OperatorID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 203, 199, value)
}

// SetAuditTrail_OperatorID
func (m *TradeActivityDataPointerBuilder) SetAuditTrail_OperatorID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 203, 199, value)
}

// SetAuditTrail_SelfMatchPreventionID
func (m *TradeActivityDataBuilder) SetAuditTrail_SelfMatchPreventionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 207, 203, value)
}

// SetAuditTrail_SelfMatchPreventionID
func (m *TradeActivityDataPointerBuilder) SetAuditTrail_SelfMatchPreventionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 207, 203, value)
}

// SetAuditTrail_SessionID
func (m *TradeActivityDataBuilder) SetAuditTrail_SessionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 211, 207, value)
}

// SetAuditTrail_SessionID
func (m *TradeActivityDataPointerBuilder) SetAuditTrail_SessionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 211, 207, value)
}

// SetAuditTrail_ExecutingFirmID
func (m *TradeActivityDataBuilder) SetAuditTrail_ExecutingFirmID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 215, 211, value)
}

// SetAuditTrail_ExecutingFirmID
func (m *TradeActivityDataPointerBuilder) SetAuditTrail_ExecutingFirmID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 215, 211, value)
}

// SetAuditTrail_FixMessageType
func (m *TradeActivityDataBuilder) SetAuditTrail_FixMessageType(value string) {
	message.SetStringVLS(&m.VLSBuilder, 219, 215, value)
}

// SetAuditTrail_FixMessageType
func (m *TradeActivityDataPointerBuilder) SetAuditTrail_FixMessageType(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 219, 215, value)
}

// SetAuditTrail_CustomerTypeIndicator
func (m TradeActivityDataBuilder) SetAuditTrail_CustomerTypeIndicator(value int16) {
	message.SetInt16VLS(m.Unsafe(), 221, 219, value)
}

// SetAuditTrail_CustomerTypeIndicator
func (m TradeActivityDataPointerBuilder) SetAuditTrail_CustomerTypeIndicator(value int16) {
	message.SetInt16VLS(m.Ptr, 221, 219, value)
}

// SetAuditTrail_CustomerOrFirm
func (m TradeActivityDataBuilder) SetAuditTrail_CustomerOrFirm(value int16) {
	message.SetInt16VLS(m.Unsafe(), 223, 221, value)
}

// SetAuditTrail_CustomerOrFirm
func (m TradeActivityDataPointerBuilder) SetAuditTrail_CustomerOrFirm(value int16) {
	message.SetInt16VLS(m.Ptr, 223, 221, value)
}

// SetAuditTrail_ExecutionReportID
func (m *TradeActivityDataBuilder) SetAuditTrail_ExecutionReportID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 227, 223, value)
}

// SetAuditTrail_ExecutionReportID
func (m *TradeActivityDataPointerBuilder) SetAuditTrail_ExecutionReportID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 227, 223, value)
}

// SetAuditTrail_SpreadLegLinkID
func (m *TradeActivityDataBuilder) SetAuditTrail_SpreadLegLinkID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 231, 227, value)
}

// SetAuditTrail_SpreadLegLinkID
func (m *TradeActivityDataPointerBuilder) SetAuditTrail_SpreadLegLinkID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 231, 227, value)
}

// SetAuditTrail_SecurityDesc
func (m *TradeActivityDataBuilder) SetAuditTrail_SecurityDesc(value string) {
	message.SetStringVLS(&m.VLSBuilder, 235, 231, value)
}

// SetAuditTrail_SecurityDesc
func (m *TradeActivityDataPointerBuilder) SetAuditTrail_SecurityDesc(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 235, 231, value)
}

// SetAuditTrail_MarketSegmentID
func (m *TradeActivityDataBuilder) SetAuditTrail_MarketSegmentID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 239, 235, value)
}

// SetAuditTrail_MarketSegmentID
func (m *TradeActivityDataPointerBuilder) SetAuditTrail_MarketSegmentID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 239, 235, value)
}

// SetAuditTrail_IFMFlag
func (m TradeActivityDataBuilder) SetAuditTrail_IFMFlag(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 240, 239, value)
}

// SetAuditTrail_IFMFlag
func (m TradeActivityDataPointerBuilder) SetAuditTrail_IFMFlag(value uint8) {
	message.SetUint8VLS(m.Ptr, 240, 239, value)
}

// SetAuditTrail_DisplayQuantity
func (m TradeActivityDataBuilder) SetAuditTrail_DisplayQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 248, 240, value)
}

// SetAuditTrail_DisplayQuantity
func (m TradeActivityDataPointerBuilder) SetAuditTrail_DisplayQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 248, 240, value)
}

// SetAuditTrail_CountryOfOrigin
func (m *TradeActivityDataBuilder) SetAuditTrail_CountryOfOrigin(value string) {
	message.SetStringVLS(&m.VLSBuilder, 252, 248, value)
}

// SetAuditTrail_CountryOfOrigin
func (m *TradeActivityDataPointerBuilder) SetAuditTrail_CountryOfOrigin(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 252, 248, value)
}

// SetAuditTrail_FillQuantity
func (m TradeActivityDataBuilder) SetAuditTrail_FillQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 260, 252, value)
}

// SetAuditTrail_FillQuantity
func (m TradeActivityDataPointerBuilder) SetAuditTrail_FillQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 260, 252, value)
}

// SetAuditTrail_RemainingQuantity
func (m TradeActivityDataBuilder) SetAuditTrail_RemainingQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 268, 260, value)
}

// SetAuditTrail_RemainingQuantity
func (m TradeActivityDataPointerBuilder) SetAuditTrail_RemainingQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 268, 260, value)
}

// SetAuditTrail_AggressorFlag
func (m TradeActivityDataBuilder) SetAuditTrail_AggressorFlag(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 269, 268, value)
}

// SetAuditTrail_AggressorFlag
func (m TradeActivityDataPointerBuilder) SetAuditTrail_AggressorFlag(value uint8) {
	message.SetUint8VLS(m.Ptr, 269, 268, value)
}

// SetAuditTrail_SourceOfCancellation
func (m TradeActivityDataBuilder) SetAuditTrail_SourceOfCancellation(value int32) {
	message.SetInt32VLS(m.Unsafe(), 273, 269, value)
}

// SetAuditTrail_SourceOfCancellation
func (m TradeActivityDataPointerBuilder) SetAuditTrail_SourceOfCancellation(value int32) {
	message.SetInt32VLS(m.Ptr, 273, 269, value)
}

// SetAuditTrail_OrdRejReason
func (m *TradeActivityDataBuilder) SetAuditTrail_OrdRejReason(value string) {
	message.SetStringVLS(&m.VLSBuilder, 277, 273, value)
}

// SetAuditTrail_OrdRejReason
func (m *TradeActivityDataPointerBuilder) SetAuditTrail_OrdRejReason(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 277, 273, value)
}

// SetIsSnapshot
func (m TradeActivityDataBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Unsafe(), 278, 277, value)
}

// SetIsSnapshot
func (m TradeActivityDataPointerBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Ptr, 278, 277, value)
}

// SetIsFirstMessageInBatch
func (m TradeActivityDataBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolVLS(m.Unsafe(), 279, 278, value)
}

// SetIsFirstMessageInBatch
func (m TradeActivityDataPointerBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolVLS(m.Ptr, 279, 278, value)
}

// SetIsLastMessageInBatch
func (m TradeActivityDataBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolVLS(m.Unsafe(), 280, 279, value)
}

// SetIsLastMessageInBatch
func (m TradeActivityDataPointerBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolVLS(m.Ptr, 280, 279, value)
}

// Copy
func (m TradeActivityData) Copy(to TradeActivityDataBuilder) {
	to.SetActivityType(m.ActivityType())
	to.SetDataDateTimeUTC(m.DataDateTimeUTC())
	to.SetSymbol(m.Symbol())
	to.SetOrderActionSource(m.OrderActionSource())
	to.SetInternalOrderID(m.InternalOrderID())
	to.SetServiceOrderID(m.ServiceOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetFIXClientOrderID(m.FIXClientOrderID())
	to.SetOrderTypeName(m.OrderTypeName())
	to.SetQuantity(m.Quantity())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetNewOrderStatus(m.NewOrderStatus())
	to.SetFillPrice(m.FillPrice())
	to.SetOrderFilledQuantity(m.OrderFilledQuantity())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetLastPriceDuringPosition(m.LastPriceDuringPosition())
	to.SetTradeAccount(m.TradeAccount())
	to.SetParentInternalOrderID(m.ParentInternalOrderID())
	to.SetOpenClose(m.OpenClose())
	to.SetIsSimulated(m.IsSimulated())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsChartReplaying(m.IsChartReplaying())
	to.SetFillExecutionServiceID(m.FillExecutionServiceID())
	to.SetPositionQuantity(m.PositionQuantity())
	to.SetSourceChartNumber(m.SourceChartNumber())
	to.SetSourceChartbookFileName(m.SourceChartbookFileName())
	to.SetTimeInForce(m.TimeInForce())
	to.SetSymbolServiceCode(m.SymbolServiceCode())
	to.SetNote(m.Note())
	to.SetOriginatingClientUsername(m.OriginatingClientUsername())
	to.SetTradeAccountBalance(m.TradeAccountBalance())
	to.SetSupportsPositionQuantityField(m.SupportsPositionQuantityField())
	to.SetIsBillable(m.IsBillable())
	to.SetQuantityForBilling(m.QuantityForBilling())
	to.SetOrderRoutingServiceCode(m.OrderRoutingServiceCode())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetAuditTrail_TransactDateTimeUTC(m.AuditTrail_TransactDateTimeUTC())
	to.SetAuditTrail_MessageDirection(m.AuditTrail_MessageDirection())
	to.SetAuditTrail_OperatorID(m.AuditTrail_OperatorID())
	to.SetAuditTrail_SelfMatchPreventionID(m.AuditTrail_SelfMatchPreventionID())
	to.SetAuditTrail_SessionID(m.AuditTrail_SessionID())
	to.SetAuditTrail_ExecutingFirmID(m.AuditTrail_ExecutingFirmID())
	to.SetAuditTrail_FixMessageType(m.AuditTrail_FixMessageType())
	to.SetAuditTrail_CustomerTypeIndicator(m.AuditTrail_CustomerTypeIndicator())
	to.SetAuditTrail_CustomerOrFirm(m.AuditTrail_CustomerOrFirm())
	to.SetAuditTrail_ExecutionReportID(m.AuditTrail_ExecutionReportID())
	to.SetAuditTrail_SpreadLegLinkID(m.AuditTrail_SpreadLegLinkID())
	to.SetAuditTrail_SecurityDesc(m.AuditTrail_SecurityDesc())
	to.SetAuditTrail_MarketSegmentID(m.AuditTrail_MarketSegmentID())
	to.SetAuditTrail_IFMFlag(m.AuditTrail_IFMFlag())
	to.SetAuditTrail_DisplayQuantity(m.AuditTrail_DisplayQuantity())
	to.SetAuditTrail_CountryOfOrigin(m.AuditTrail_CountryOfOrigin())
	to.SetAuditTrail_FillQuantity(m.AuditTrail_FillQuantity())
	to.SetAuditTrail_RemainingQuantity(m.AuditTrail_RemainingQuantity())
	to.SetAuditTrail_AggressorFlag(m.AuditTrail_AggressorFlag())
	to.SetAuditTrail_SourceOfCancellation(m.AuditTrail_SourceOfCancellation())
	to.SetAuditTrail_OrdRejReason(m.AuditTrail_OrdRejReason())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// Copy
func (m TradeActivityDataBuilder) Copy(to TradeActivityDataBuilder) {
	to.SetActivityType(m.ActivityType())
	to.SetDataDateTimeUTC(m.DataDateTimeUTC())
	to.SetSymbol(m.Symbol())
	to.SetOrderActionSource(m.OrderActionSource())
	to.SetInternalOrderID(m.InternalOrderID())
	to.SetServiceOrderID(m.ServiceOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetFIXClientOrderID(m.FIXClientOrderID())
	to.SetOrderTypeName(m.OrderTypeName())
	to.SetQuantity(m.Quantity())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetNewOrderStatus(m.NewOrderStatus())
	to.SetFillPrice(m.FillPrice())
	to.SetOrderFilledQuantity(m.OrderFilledQuantity())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetLastPriceDuringPosition(m.LastPriceDuringPosition())
	to.SetTradeAccount(m.TradeAccount())
	to.SetParentInternalOrderID(m.ParentInternalOrderID())
	to.SetOpenClose(m.OpenClose())
	to.SetIsSimulated(m.IsSimulated())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsChartReplaying(m.IsChartReplaying())
	to.SetFillExecutionServiceID(m.FillExecutionServiceID())
	to.SetPositionQuantity(m.PositionQuantity())
	to.SetSourceChartNumber(m.SourceChartNumber())
	to.SetSourceChartbookFileName(m.SourceChartbookFileName())
	to.SetTimeInForce(m.TimeInForce())
	to.SetSymbolServiceCode(m.SymbolServiceCode())
	to.SetNote(m.Note())
	to.SetOriginatingClientUsername(m.OriginatingClientUsername())
	to.SetTradeAccountBalance(m.TradeAccountBalance())
	to.SetSupportsPositionQuantityField(m.SupportsPositionQuantityField())
	to.SetIsBillable(m.IsBillable())
	to.SetQuantityForBilling(m.QuantityForBilling())
	to.SetOrderRoutingServiceCode(m.OrderRoutingServiceCode())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetAuditTrail_TransactDateTimeUTC(m.AuditTrail_TransactDateTimeUTC())
	to.SetAuditTrail_MessageDirection(m.AuditTrail_MessageDirection())
	to.SetAuditTrail_OperatorID(m.AuditTrail_OperatorID())
	to.SetAuditTrail_SelfMatchPreventionID(m.AuditTrail_SelfMatchPreventionID())
	to.SetAuditTrail_SessionID(m.AuditTrail_SessionID())
	to.SetAuditTrail_ExecutingFirmID(m.AuditTrail_ExecutingFirmID())
	to.SetAuditTrail_FixMessageType(m.AuditTrail_FixMessageType())
	to.SetAuditTrail_CustomerTypeIndicator(m.AuditTrail_CustomerTypeIndicator())
	to.SetAuditTrail_CustomerOrFirm(m.AuditTrail_CustomerOrFirm())
	to.SetAuditTrail_ExecutionReportID(m.AuditTrail_ExecutionReportID())
	to.SetAuditTrail_SpreadLegLinkID(m.AuditTrail_SpreadLegLinkID())
	to.SetAuditTrail_SecurityDesc(m.AuditTrail_SecurityDesc())
	to.SetAuditTrail_MarketSegmentID(m.AuditTrail_MarketSegmentID())
	to.SetAuditTrail_IFMFlag(m.AuditTrail_IFMFlag())
	to.SetAuditTrail_DisplayQuantity(m.AuditTrail_DisplayQuantity())
	to.SetAuditTrail_CountryOfOrigin(m.AuditTrail_CountryOfOrigin())
	to.SetAuditTrail_FillQuantity(m.AuditTrail_FillQuantity())
	to.SetAuditTrail_RemainingQuantity(m.AuditTrail_RemainingQuantity())
	to.SetAuditTrail_AggressorFlag(m.AuditTrail_AggressorFlag())
	to.SetAuditTrail_SourceOfCancellation(m.AuditTrail_SourceOfCancellation())
	to.SetAuditTrail_OrdRejReason(m.AuditTrail_OrdRejReason())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// CopyPointer
func (m TradeActivityData) CopyPointer(to TradeActivityDataPointerBuilder) {
	to.SetActivityType(m.ActivityType())
	to.SetDataDateTimeUTC(m.DataDateTimeUTC())
	to.SetSymbol(m.Symbol())
	to.SetOrderActionSource(m.OrderActionSource())
	to.SetInternalOrderID(m.InternalOrderID())
	to.SetServiceOrderID(m.ServiceOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetFIXClientOrderID(m.FIXClientOrderID())
	to.SetOrderTypeName(m.OrderTypeName())
	to.SetQuantity(m.Quantity())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetNewOrderStatus(m.NewOrderStatus())
	to.SetFillPrice(m.FillPrice())
	to.SetOrderFilledQuantity(m.OrderFilledQuantity())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetLastPriceDuringPosition(m.LastPriceDuringPosition())
	to.SetTradeAccount(m.TradeAccount())
	to.SetParentInternalOrderID(m.ParentInternalOrderID())
	to.SetOpenClose(m.OpenClose())
	to.SetIsSimulated(m.IsSimulated())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsChartReplaying(m.IsChartReplaying())
	to.SetFillExecutionServiceID(m.FillExecutionServiceID())
	to.SetPositionQuantity(m.PositionQuantity())
	to.SetSourceChartNumber(m.SourceChartNumber())
	to.SetSourceChartbookFileName(m.SourceChartbookFileName())
	to.SetTimeInForce(m.TimeInForce())
	to.SetSymbolServiceCode(m.SymbolServiceCode())
	to.SetNote(m.Note())
	to.SetOriginatingClientUsername(m.OriginatingClientUsername())
	to.SetTradeAccountBalance(m.TradeAccountBalance())
	to.SetSupportsPositionQuantityField(m.SupportsPositionQuantityField())
	to.SetIsBillable(m.IsBillable())
	to.SetQuantityForBilling(m.QuantityForBilling())
	to.SetOrderRoutingServiceCode(m.OrderRoutingServiceCode())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetAuditTrail_TransactDateTimeUTC(m.AuditTrail_TransactDateTimeUTC())
	to.SetAuditTrail_MessageDirection(m.AuditTrail_MessageDirection())
	to.SetAuditTrail_OperatorID(m.AuditTrail_OperatorID())
	to.SetAuditTrail_SelfMatchPreventionID(m.AuditTrail_SelfMatchPreventionID())
	to.SetAuditTrail_SessionID(m.AuditTrail_SessionID())
	to.SetAuditTrail_ExecutingFirmID(m.AuditTrail_ExecutingFirmID())
	to.SetAuditTrail_FixMessageType(m.AuditTrail_FixMessageType())
	to.SetAuditTrail_CustomerTypeIndicator(m.AuditTrail_CustomerTypeIndicator())
	to.SetAuditTrail_CustomerOrFirm(m.AuditTrail_CustomerOrFirm())
	to.SetAuditTrail_ExecutionReportID(m.AuditTrail_ExecutionReportID())
	to.SetAuditTrail_SpreadLegLinkID(m.AuditTrail_SpreadLegLinkID())
	to.SetAuditTrail_SecurityDesc(m.AuditTrail_SecurityDesc())
	to.SetAuditTrail_MarketSegmentID(m.AuditTrail_MarketSegmentID())
	to.SetAuditTrail_IFMFlag(m.AuditTrail_IFMFlag())
	to.SetAuditTrail_DisplayQuantity(m.AuditTrail_DisplayQuantity())
	to.SetAuditTrail_CountryOfOrigin(m.AuditTrail_CountryOfOrigin())
	to.SetAuditTrail_FillQuantity(m.AuditTrail_FillQuantity())
	to.SetAuditTrail_RemainingQuantity(m.AuditTrail_RemainingQuantity())
	to.SetAuditTrail_AggressorFlag(m.AuditTrail_AggressorFlag())
	to.SetAuditTrail_SourceOfCancellation(m.AuditTrail_SourceOfCancellation())
	to.SetAuditTrail_OrdRejReason(m.AuditTrail_OrdRejReason())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// CopyPointer
func (m TradeActivityDataBuilder) CopyPointer(to TradeActivityDataPointerBuilder) {
	to.SetActivityType(m.ActivityType())
	to.SetDataDateTimeUTC(m.DataDateTimeUTC())
	to.SetSymbol(m.Symbol())
	to.SetOrderActionSource(m.OrderActionSource())
	to.SetInternalOrderID(m.InternalOrderID())
	to.SetServiceOrderID(m.ServiceOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetFIXClientOrderID(m.FIXClientOrderID())
	to.SetOrderTypeName(m.OrderTypeName())
	to.SetQuantity(m.Quantity())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetNewOrderStatus(m.NewOrderStatus())
	to.SetFillPrice(m.FillPrice())
	to.SetOrderFilledQuantity(m.OrderFilledQuantity())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetLastPriceDuringPosition(m.LastPriceDuringPosition())
	to.SetTradeAccount(m.TradeAccount())
	to.SetParentInternalOrderID(m.ParentInternalOrderID())
	to.SetOpenClose(m.OpenClose())
	to.SetIsSimulated(m.IsSimulated())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsChartReplaying(m.IsChartReplaying())
	to.SetFillExecutionServiceID(m.FillExecutionServiceID())
	to.SetPositionQuantity(m.PositionQuantity())
	to.SetSourceChartNumber(m.SourceChartNumber())
	to.SetSourceChartbookFileName(m.SourceChartbookFileName())
	to.SetTimeInForce(m.TimeInForce())
	to.SetSymbolServiceCode(m.SymbolServiceCode())
	to.SetNote(m.Note())
	to.SetOriginatingClientUsername(m.OriginatingClientUsername())
	to.SetTradeAccountBalance(m.TradeAccountBalance())
	to.SetSupportsPositionQuantityField(m.SupportsPositionQuantityField())
	to.SetIsBillable(m.IsBillable())
	to.SetQuantityForBilling(m.QuantityForBilling())
	to.SetOrderRoutingServiceCode(m.OrderRoutingServiceCode())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetAuditTrail_TransactDateTimeUTC(m.AuditTrail_TransactDateTimeUTC())
	to.SetAuditTrail_MessageDirection(m.AuditTrail_MessageDirection())
	to.SetAuditTrail_OperatorID(m.AuditTrail_OperatorID())
	to.SetAuditTrail_SelfMatchPreventionID(m.AuditTrail_SelfMatchPreventionID())
	to.SetAuditTrail_SessionID(m.AuditTrail_SessionID())
	to.SetAuditTrail_ExecutingFirmID(m.AuditTrail_ExecutingFirmID())
	to.SetAuditTrail_FixMessageType(m.AuditTrail_FixMessageType())
	to.SetAuditTrail_CustomerTypeIndicator(m.AuditTrail_CustomerTypeIndicator())
	to.SetAuditTrail_CustomerOrFirm(m.AuditTrail_CustomerOrFirm())
	to.SetAuditTrail_ExecutionReportID(m.AuditTrail_ExecutionReportID())
	to.SetAuditTrail_SpreadLegLinkID(m.AuditTrail_SpreadLegLinkID())
	to.SetAuditTrail_SecurityDesc(m.AuditTrail_SecurityDesc())
	to.SetAuditTrail_MarketSegmentID(m.AuditTrail_MarketSegmentID())
	to.SetAuditTrail_IFMFlag(m.AuditTrail_IFMFlag())
	to.SetAuditTrail_DisplayQuantity(m.AuditTrail_DisplayQuantity())
	to.SetAuditTrail_CountryOfOrigin(m.AuditTrail_CountryOfOrigin())
	to.SetAuditTrail_FillQuantity(m.AuditTrail_FillQuantity())
	to.SetAuditTrail_RemainingQuantity(m.AuditTrail_RemainingQuantity())
	to.SetAuditTrail_AggressorFlag(m.AuditTrail_AggressorFlag())
	to.SetAuditTrail_SourceOfCancellation(m.AuditTrail_SourceOfCancellation())
	to.SetAuditTrail_OrdRejReason(m.AuditTrail_OrdRejReason())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// Copy
func (m TradeActivityDataPointer) Copy(to TradeActivityDataBuilder) {
	to.SetActivityType(m.ActivityType())
	to.SetDataDateTimeUTC(m.DataDateTimeUTC())
	to.SetSymbol(m.Symbol())
	to.SetOrderActionSource(m.OrderActionSource())
	to.SetInternalOrderID(m.InternalOrderID())
	to.SetServiceOrderID(m.ServiceOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetFIXClientOrderID(m.FIXClientOrderID())
	to.SetOrderTypeName(m.OrderTypeName())
	to.SetQuantity(m.Quantity())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetNewOrderStatus(m.NewOrderStatus())
	to.SetFillPrice(m.FillPrice())
	to.SetOrderFilledQuantity(m.OrderFilledQuantity())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetLastPriceDuringPosition(m.LastPriceDuringPosition())
	to.SetTradeAccount(m.TradeAccount())
	to.SetParentInternalOrderID(m.ParentInternalOrderID())
	to.SetOpenClose(m.OpenClose())
	to.SetIsSimulated(m.IsSimulated())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsChartReplaying(m.IsChartReplaying())
	to.SetFillExecutionServiceID(m.FillExecutionServiceID())
	to.SetPositionQuantity(m.PositionQuantity())
	to.SetSourceChartNumber(m.SourceChartNumber())
	to.SetSourceChartbookFileName(m.SourceChartbookFileName())
	to.SetTimeInForce(m.TimeInForce())
	to.SetSymbolServiceCode(m.SymbolServiceCode())
	to.SetNote(m.Note())
	to.SetOriginatingClientUsername(m.OriginatingClientUsername())
	to.SetTradeAccountBalance(m.TradeAccountBalance())
	to.SetSupportsPositionQuantityField(m.SupportsPositionQuantityField())
	to.SetIsBillable(m.IsBillable())
	to.SetQuantityForBilling(m.QuantityForBilling())
	to.SetOrderRoutingServiceCode(m.OrderRoutingServiceCode())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetAuditTrail_TransactDateTimeUTC(m.AuditTrail_TransactDateTimeUTC())
	to.SetAuditTrail_MessageDirection(m.AuditTrail_MessageDirection())
	to.SetAuditTrail_OperatorID(m.AuditTrail_OperatorID())
	to.SetAuditTrail_SelfMatchPreventionID(m.AuditTrail_SelfMatchPreventionID())
	to.SetAuditTrail_SessionID(m.AuditTrail_SessionID())
	to.SetAuditTrail_ExecutingFirmID(m.AuditTrail_ExecutingFirmID())
	to.SetAuditTrail_FixMessageType(m.AuditTrail_FixMessageType())
	to.SetAuditTrail_CustomerTypeIndicator(m.AuditTrail_CustomerTypeIndicator())
	to.SetAuditTrail_CustomerOrFirm(m.AuditTrail_CustomerOrFirm())
	to.SetAuditTrail_ExecutionReportID(m.AuditTrail_ExecutionReportID())
	to.SetAuditTrail_SpreadLegLinkID(m.AuditTrail_SpreadLegLinkID())
	to.SetAuditTrail_SecurityDesc(m.AuditTrail_SecurityDesc())
	to.SetAuditTrail_MarketSegmentID(m.AuditTrail_MarketSegmentID())
	to.SetAuditTrail_IFMFlag(m.AuditTrail_IFMFlag())
	to.SetAuditTrail_DisplayQuantity(m.AuditTrail_DisplayQuantity())
	to.SetAuditTrail_CountryOfOrigin(m.AuditTrail_CountryOfOrigin())
	to.SetAuditTrail_FillQuantity(m.AuditTrail_FillQuantity())
	to.SetAuditTrail_RemainingQuantity(m.AuditTrail_RemainingQuantity())
	to.SetAuditTrail_AggressorFlag(m.AuditTrail_AggressorFlag())
	to.SetAuditTrail_SourceOfCancellation(m.AuditTrail_SourceOfCancellation())
	to.SetAuditTrail_OrdRejReason(m.AuditTrail_OrdRejReason())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// Copy
func (m TradeActivityDataPointerBuilder) Copy(to TradeActivityDataBuilder) {
	to.SetActivityType(m.ActivityType())
	to.SetDataDateTimeUTC(m.DataDateTimeUTC())
	to.SetSymbol(m.Symbol())
	to.SetOrderActionSource(m.OrderActionSource())
	to.SetInternalOrderID(m.InternalOrderID())
	to.SetServiceOrderID(m.ServiceOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetFIXClientOrderID(m.FIXClientOrderID())
	to.SetOrderTypeName(m.OrderTypeName())
	to.SetQuantity(m.Quantity())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetNewOrderStatus(m.NewOrderStatus())
	to.SetFillPrice(m.FillPrice())
	to.SetOrderFilledQuantity(m.OrderFilledQuantity())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetLastPriceDuringPosition(m.LastPriceDuringPosition())
	to.SetTradeAccount(m.TradeAccount())
	to.SetParentInternalOrderID(m.ParentInternalOrderID())
	to.SetOpenClose(m.OpenClose())
	to.SetIsSimulated(m.IsSimulated())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsChartReplaying(m.IsChartReplaying())
	to.SetFillExecutionServiceID(m.FillExecutionServiceID())
	to.SetPositionQuantity(m.PositionQuantity())
	to.SetSourceChartNumber(m.SourceChartNumber())
	to.SetSourceChartbookFileName(m.SourceChartbookFileName())
	to.SetTimeInForce(m.TimeInForce())
	to.SetSymbolServiceCode(m.SymbolServiceCode())
	to.SetNote(m.Note())
	to.SetOriginatingClientUsername(m.OriginatingClientUsername())
	to.SetTradeAccountBalance(m.TradeAccountBalance())
	to.SetSupportsPositionQuantityField(m.SupportsPositionQuantityField())
	to.SetIsBillable(m.IsBillable())
	to.SetQuantityForBilling(m.QuantityForBilling())
	to.SetOrderRoutingServiceCode(m.OrderRoutingServiceCode())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetAuditTrail_TransactDateTimeUTC(m.AuditTrail_TransactDateTimeUTC())
	to.SetAuditTrail_MessageDirection(m.AuditTrail_MessageDirection())
	to.SetAuditTrail_OperatorID(m.AuditTrail_OperatorID())
	to.SetAuditTrail_SelfMatchPreventionID(m.AuditTrail_SelfMatchPreventionID())
	to.SetAuditTrail_SessionID(m.AuditTrail_SessionID())
	to.SetAuditTrail_ExecutingFirmID(m.AuditTrail_ExecutingFirmID())
	to.SetAuditTrail_FixMessageType(m.AuditTrail_FixMessageType())
	to.SetAuditTrail_CustomerTypeIndicator(m.AuditTrail_CustomerTypeIndicator())
	to.SetAuditTrail_CustomerOrFirm(m.AuditTrail_CustomerOrFirm())
	to.SetAuditTrail_ExecutionReportID(m.AuditTrail_ExecutionReportID())
	to.SetAuditTrail_SpreadLegLinkID(m.AuditTrail_SpreadLegLinkID())
	to.SetAuditTrail_SecurityDesc(m.AuditTrail_SecurityDesc())
	to.SetAuditTrail_MarketSegmentID(m.AuditTrail_MarketSegmentID())
	to.SetAuditTrail_IFMFlag(m.AuditTrail_IFMFlag())
	to.SetAuditTrail_DisplayQuantity(m.AuditTrail_DisplayQuantity())
	to.SetAuditTrail_CountryOfOrigin(m.AuditTrail_CountryOfOrigin())
	to.SetAuditTrail_FillQuantity(m.AuditTrail_FillQuantity())
	to.SetAuditTrail_RemainingQuantity(m.AuditTrail_RemainingQuantity())
	to.SetAuditTrail_AggressorFlag(m.AuditTrail_AggressorFlag())
	to.SetAuditTrail_SourceOfCancellation(m.AuditTrail_SourceOfCancellation())
	to.SetAuditTrail_OrdRejReason(m.AuditTrail_OrdRejReason())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// CopyPointer
func (m TradeActivityDataPointer) CopyPointer(to TradeActivityDataPointerBuilder) {
	to.SetActivityType(m.ActivityType())
	to.SetDataDateTimeUTC(m.DataDateTimeUTC())
	to.SetSymbol(m.Symbol())
	to.SetOrderActionSource(m.OrderActionSource())
	to.SetInternalOrderID(m.InternalOrderID())
	to.SetServiceOrderID(m.ServiceOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetFIXClientOrderID(m.FIXClientOrderID())
	to.SetOrderTypeName(m.OrderTypeName())
	to.SetQuantity(m.Quantity())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetNewOrderStatus(m.NewOrderStatus())
	to.SetFillPrice(m.FillPrice())
	to.SetOrderFilledQuantity(m.OrderFilledQuantity())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetLastPriceDuringPosition(m.LastPriceDuringPosition())
	to.SetTradeAccount(m.TradeAccount())
	to.SetParentInternalOrderID(m.ParentInternalOrderID())
	to.SetOpenClose(m.OpenClose())
	to.SetIsSimulated(m.IsSimulated())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsChartReplaying(m.IsChartReplaying())
	to.SetFillExecutionServiceID(m.FillExecutionServiceID())
	to.SetPositionQuantity(m.PositionQuantity())
	to.SetSourceChartNumber(m.SourceChartNumber())
	to.SetSourceChartbookFileName(m.SourceChartbookFileName())
	to.SetTimeInForce(m.TimeInForce())
	to.SetSymbolServiceCode(m.SymbolServiceCode())
	to.SetNote(m.Note())
	to.SetOriginatingClientUsername(m.OriginatingClientUsername())
	to.SetTradeAccountBalance(m.TradeAccountBalance())
	to.SetSupportsPositionQuantityField(m.SupportsPositionQuantityField())
	to.SetIsBillable(m.IsBillable())
	to.SetQuantityForBilling(m.QuantityForBilling())
	to.SetOrderRoutingServiceCode(m.OrderRoutingServiceCode())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetAuditTrail_TransactDateTimeUTC(m.AuditTrail_TransactDateTimeUTC())
	to.SetAuditTrail_MessageDirection(m.AuditTrail_MessageDirection())
	to.SetAuditTrail_OperatorID(m.AuditTrail_OperatorID())
	to.SetAuditTrail_SelfMatchPreventionID(m.AuditTrail_SelfMatchPreventionID())
	to.SetAuditTrail_SessionID(m.AuditTrail_SessionID())
	to.SetAuditTrail_ExecutingFirmID(m.AuditTrail_ExecutingFirmID())
	to.SetAuditTrail_FixMessageType(m.AuditTrail_FixMessageType())
	to.SetAuditTrail_CustomerTypeIndicator(m.AuditTrail_CustomerTypeIndicator())
	to.SetAuditTrail_CustomerOrFirm(m.AuditTrail_CustomerOrFirm())
	to.SetAuditTrail_ExecutionReportID(m.AuditTrail_ExecutionReportID())
	to.SetAuditTrail_SpreadLegLinkID(m.AuditTrail_SpreadLegLinkID())
	to.SetAuditTrail_SecurityDesc(m.AuditTrail_SecurityDesc())
	to.SetAuditTrail_MarketSegmentID(m.AuditTrail_MarketSegmentID())
	to.SetAuditTrail_IFMFlag(m.AuditTrail_IFMFlag())
	to.SetAuditTrail_DisplayQuantity(m.AuditTrail_DisplayQuantity())
	to.SetAuditTrail_CountryOfOrigin(m.AuditTrail_CountryOfOrigin())
	to.SetAuditTrail_FillQuantity(m.AuditTrail_FillQuantity())
	to.SetAuditTrail_RemainingQuantity(m.AuditTrail_RemainingQuantity())
	to.SetAuditTrail_AggressorFlag(m.AuditTrail_AggressorFlag())
	to.SetAuditTrail_SourceOfCancellation(m.AuditTrail_SourceOfCancellation())
	to.SetAuditTrail_OrdRejReason(m.AuditTrail_OrdRejReason())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// CopyPointer
func (m TradeActivityDataPointerBuilder) CopyPointer(to TradeActivityDataPointerBuilder) {
	to.SetActivityType(m.ActivityType())
	to.SetDataDateTimeUTC(m.DataDateTimeUTC())
	to.SetSymbol(m.Symbol())
	to.SetOrderActionSource(m.OrderActionSource())
	to.SetInternalOrderID(m.InternalOrderID())
	to.SetServiceOrderID(m.ServiceOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetFIXClientOrderID(m.FIXClientOrderID())
	to.SetOrderTypeName(m.OrderTypeName())
	to.SetQuantity(m.Quantity())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetNewOrderStatus(m.NewOrderStatus())
	to.SetFillPrice(m.FillPrice())
	to.SetOrderFilledQuantity(m.OrderFilledQuantity())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetLastPriceDuringPosition(m.LastPriceDuringPosition())
	to.SetTradeAccount(m.TradeAccount())
	to.SetParentInternalOrderID(m.ParentInternalOrderID())
	to.SetOpenClose(m.OpenClose())
	to.SetIsSimulated(m.IsSimulated())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsChartReplaying(m.IsChartReplaying())
	to.SetFillExecutionServiceID(m.FillExecutionServiceID())
	to.SetPositionQuantity(m.PositionQuantity())
	to.SetSourceChartNumber(m.SourceChartNumber())
	to.SetSourceChartbookFileName(m.SourceChartbookFileName())
	to.SetTimeInForce(m.TimeInForce())
	to.SetSymbolServiceCode(m.SymbolServiceCode())
	to.SetNote(m.Note())
	to.SetOriginatingClientUsername(m.OriginatingClientUsername())
	to.SetTradeAccountBalance(m.TradeAccountBalance())
	to.SetSupportsPositionQuantityField(m.SupportsPositionQuantityField())
	to.SetIsBillable(m.IsBillable())
	to.SetQuantityForBilling(m.QuantityForBilling())
	to.SetOrderRoutingServiceCode(m.OrderRoutingServiceCode())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetAuditTrail_TransactDateTimeUTC(m.AuditTrail_TransactDateTimeUTC())
	to.SetAuditTrail_MessageDirection(m.AuditTrail_MessageDirection())
	to.SetAuditTrail_OperatorID(m.AuditTrail_OperatorID())
	to.SetAuditTrail_SelfMatchPreventionID(m.AuditTrail_SelfMatchPreventionID())
	to.SetAuditTrail_SessionID(m.AuditTrail_SessionID())
	to.SetAuditTrail_ExecutingFirmID(m.AuditTrail_ExecutingFirmID())
	to.SetAuditTrail_FixMessageType(m.AuditTrail_FixMessageType())
	to.SetAuditTrail_CustomerTypeIndicator(m.AuditTrail_CustomerTypeIndicator())
	to.SetAuditTrail_CustomerOrFirm(m.AuditTrail_CustomerOrFirm())
	to.SetAuditTrail_ExecutionReportID(m.AuditTrail_ExecutionReportID())
	to.SetAuditTrail_SpreadLegLinkID(m.AuditTrail_SpreadLegLinkID())
	to.SetAuditTrail_SecurityDesc(m.AuditTrail_SecurityDesc())
	to.SetAuditTrail_MarketSegmentID(m.AuditTrail_MarketSegmentID())
	to.SetAuditTrail_IFMFlag(m.AuditTrail_IFMFlag())
	to.SetAuditTrail_DisplayQuantity(m.AuditTrail_DisplayQuantity())
	to.SetAuditTrail_CountryOfOrigin(m.AuditTrail_CountryOfOrigin())
	to.SetAuditTrail_FillQuantity(m.AuditTrail_FillQuantity())
	to.SetAuditTrail_RemainingQuantity(m.AuditTrail_RemainingQuantity())
	to.SetAuditTrail_AggressorFlag(m.AuditTrail_AggressorFlag())
	to.SetAuditTrail_SourceOfCancellation(m.AuditTrail_SourceOfCancellation())
	to.SetAuditTrail_OrdRejReason(m.AuditTrail_OrdRejReason())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}