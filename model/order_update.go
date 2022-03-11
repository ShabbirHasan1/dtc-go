package model

import (
	"github.com/moontrade/dtc-go/message"
)

const OrderUpdateSize = 208

const OrderUpdateFixedSize = 688

// {
//     Size                        = OrderUpdateSize  (208)
//     Type                        = ORDER_UPDATE  (301)
//     BaseSize                    = OrderUpdateSize  (208)
//     RequestID                   = 0
//     TotalNumMessages            = 0
//     MessageNumber               = 0
//     Symbol                      = ""
//     Exchange                    = ""
//     PreviousServerOrderID       = ""
//     ServerOrderID               = ""
//     ClientOrderID               = ""
//     ExchangeOrderID             = ""
//     OrderStatus                 = ORDER_STATUS_UNSPECIFIED  (0)
//     OrderUpdateReason           = ORDER_UPDATE_REASON_UNSET  (0)
//     OrderType                   = ORDER_TYPE_UNSET  (0)
//     BuySell                     = BUY_SELL_UNSET  (0)
//     Price1                      = math.MaxFloat64
//     Price2                      = math.MaxFloat64
//     TimeInForce                 = TIF_UNSET  (0)
//     GoodTillDateTime            = 0
//     OrderQuantity               = math.MaxFloat64
//     FilledQuantity              = math.MaxFloat64
//     RemainingQuantity           = math.MaxFloat64
//     AverageFillPrice            = math.MaxFloat64
//     LastFillPrice               = math.MaxFloat64
//     LastFillDateTime            = 0
//     LastFillQuantity            = math.MaxFloat64
//     LastFillExecutionID         = ""
//     TradeAccount                = ""
//     InfoText                    = ""
//     NoOrders                    = 0
//     ParentServerOrderID         = ""
//     OCOLinkedOrderServerOrderID = ""
//     OpenOrClose                 = TRADE_UNSET  (0)
//     PreviousClientOrderID       = ""
//     FreeFormText                = ""
//     OrderReceivedDateTime       = 0
//     LatestTransactionDateTime   = 0
// }
var _OrderUpdateDefault = []byte{208, 0, 45, 1, 208, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size                        = OrderUpdateFixedSize  (688)
//     Type                        = ORDER_UPDATE  (301)
//     RequestID                   = 0
//     TotalNumMessages            = 0
//     MessageNumber               = 0
//     Symbol                      = ""
//     Exchange                    = ""
//     PreviousServerOrderID       = ""
//     ServerOrderID               = ""
//     ClientOrderID               = ""
//     ExchangeOrderID             = ""
//     OrderStatus                 = ORDER_STATUS_UNSPECIFIED  (0)
//     OrderUpdateReason           = ORDER_UPDATE_REASON_UNSET  (0)
//     OrderType                   = ORDER_TYPE_UNSET  (0)
//     BuySell                     = BUY_SELL_UNSET  (0)
//     Price1                      = math.MaxFloat64
//     Price2                      = math.MaxFloat64
//     TimeInForce                 = TIF_UNSET  (0)
//     GoodTillDateTime            = 0
//     OrderQuantity               = math.MaxFloat64
//     FilledQuantity              = math.MaxFloat64
//     RemainingQuantity           = math.MaxFloat64
//     AverageFillPrice            = math.MaxFloat64
//     LastFillPrice               = math.MaxFloat64
//     LastFillDateTime            = 0
//     LastFillQuantity            = math.MaxFloat64
//     LastFillExecutionID         = ""
//     TradeAccount                = ""
//     InfoText                    = ""
//     NoOrders                    = 0
//     ParentServerOrderID         = ""
//     OCOLinkedOrderServerOrderID = ""
//     OpenOrClose                 = TRADE_UNSET  (0)
//     PreviousClientOrderID       = ""
//     FreeFormText                = ""
//     OrderReceivedDateTime       = 0
//     LatestTransactionDateTime   = 0
// }
var _OrderUpdateFixedDefault = []byte{176, 2, 45, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// OrderUpdate The OrderUpdate is a unified message from the Server to the Client which
// communicates the complete details of an order, the Order Status, and the
// reason for sending the message (OrderUpdateReason).
//
// DTC uses this single unified message to provide an update for an order.
// The OrderUpdateReason field provides a clear indication for each reason
// this message is being sent.
type OrderUpdate struct {
	message.VLS
}

// OrderUpdateBuilder The OrderUpdate is a unified message from the Server to the Client which
// communicates the complete details of an order, the Order Status, and the
// reason for sending the message (OrderUpdateReason).
//
// DTC uses this single unified message to provide an update for an order.
// The OrderUpdateReason field provides a clear indication for each reason
// this message is being sent.
type OrderUpdateBuilder struct {
	message.VLSBuilder
}

// OrderUpdateFixed The OrderUpdate is a unified message from the Server to the Client which
// communicates the complete details of an order, the Order Status, and the
// reason for sending the message (OrderUpdateReason).
//
// DTC uses this single unified message to provide an update for an order.
// The OrderUpdateReason field provides a clear indication for each reason
// this message is being sent.
type OrderUpdateFixed struct {
	message.Fixed
}

// OrderUpdateFixedBuilder The OrderUpdate is a unified message from the Server to the Client which
// communicates the complete details of an order, the Order Status, and the
// reason for sending the message (OrderUpdateReason).
//
// DTC uses this single unified message to provide an update for an order.
// The OrderUpdateReason field provides a clear indication for each reason
// this message is being sent.
type OrderUpdateFixedBuilder struct {
	message.Fixed
}

// OrderUpdatePointer The OrderUpdate is a unified message from the Server to the Client which
// communicates the complete details of an order, the Order Status, and the
// reason for sending the message (OrderUpdateReason).
//
// DTC uses this single unified message to provide an update for an order.
// The OrderUpdateReason field provides a clear indication for each reason
// this message is being sent.
type OrderUpdatePointer struct {
	message.VLSPointer
}

// OrderUpdatePointerBuilder The OrderUpdate is a unified message from the Server to the Client which
// communicates the complete details of an order, the Order Status, and the
// reason for sending the message (OrderUpdateReason).
//
// DTC uses this single unified message to provide an update for an order.
// The OrderUpdateReason field provides a clear indication for each reason
// this message is being sent.
type OrderUpdatePointerBuilder struct {
	message.VLSPointerBuilder
}

// OrderUpdateFixedPointer The OrderUpdate is a unified message from the Server to the Client which
// communicates the complete details of an order, the Order Status, and the
// reason for sending the message (OrderUpdateReason).
//
// DTC uses this single unified message to provide an update for an order.
// The OrderUpdateReason field provides a clear indication for each reason
// this message is being sent.
type OrderUpdateFixedPointer struct {
	message.FixedPointer
}

// OrderUpdateFixedPointerBuilder The OrderUpdate is a unified message from the Server to the Client which
// communicates the complete details of an order, the Order Status, and the
// reason for sending the message (OrderUpdateReason).
//
// DTC uses this single unified message to provide an update for an order.
// The OrderUpdateReason field provides a clear indication for each reason
// this message is being sent.
type OrderUpdateFixedPointerBuilder struct {
	message.FixedPointer
}

func NewOrderUpdateFrom(b []byte) OrderUpdate {
	return OrderUpdate{VLS: message.NewVLSFrom(b)}
}

func WrapOrderUpdate(b []byte) OrderUpdate {
	return OrderUpdate{VLS: message.WrapVLS(b)}
}

func NewOrderUpdate() OrderUpdateBuilder {
	a := OrderUpdateBuilder{message.NewVLSBuilder(208)}
	a.Unsafe().SetBytes(0, _OrderUpdateDefault)
	return a
}

func NewOrderUpdateFixedFrom(b []byte) OrderUpdateFixed {
	return OrderUpdateFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapOrderUpdateFixed(b []byte) OrderUpdateFixed {
	return OrderUpdateFixed{Fixed: message.WrapFixed(b)}
}

func NewOrderUpdateFixed() OrderUpdateFixedBuilder {
	a := OrderUpdateFixedBuilder{message.NewFixed(688)}
	a.Unsafe().SetBytes(0, _OrderUpdateFixedDefault)
	return a
}

func AllocOrderUpdate() OrderUpdatePointerBuilder {
	a := OrderUpdatePointerBuilder{message.AllocVLSBuilder(208)}
	a.Ptr.SetBytes(0, _OrderUpdateDefault)
	return a
}

func AllocOrderUpdateFrom(b []byte) OrderUpdatePointer {
	return OrderUpdatePointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocOrderUpdateFixed() OrderUpdateFixedPointerBuilder {
	a := OrderUpdateFixedPointerBuilder{message.AllocFixed(688)}
	a.Ptr.SetBytes(0, _OrderUpdateFixedDefault)
	return a
}

func AllocOrderUpdateFixedFrom(b []byte) OrderUpdateFixedPointer {
	return OrderUpdateFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                        = OrderUpdateSize  (208)
//     Type                        = ORDER_UPDATE  (301)
//     BaseSize                    = OrderUpdateSize  (208)
//     RequestID                   = 0
//     TotalNumMessages            = 0
//     MessageNumber               = 0
//     Symbol                      = ""
//     Exchange                    = ""
//     PreviousServerOrderID       = ""
//     ServerOrderID               = ""
//     ClientOrderID               = ""
//     ExchangeOrderID             = ""
//     OrderStatus                 = ORDER_STATUS_UNSPECIFIED  (0)
//     OrderUpdateReason           = ORDER_UPDATE_REASON_UNSET  (0)
//     OrderType                   = ORDER_TYPE_UNSET  (0)
//     BuySell                     = BUY_SELL_UNSET  (0)
//     Price1                      = math.MaxFloat64
//     Price2                      = math.MaxFloat64
//     TimeInForce                 = TIF_UNSET  (0)
//     GoodTillDateTime            = 0
//     OrderQuantity               = math.MaxFloat64
//     FilledQuantity              = math.MaxFloat64
//     RemainingQuantity           = math.MaxFloat64
//     AverageFillPrice            = math.MaxFloat64
//     LastFillPrice               = math.MaxFloat64
//     LastFillDateTime            = 0
//     LastFillQuantity            = math.MaxFloat64
//     LastFillExecutionID         = ""
//     TradeAccount                = ""
//     InfoText                    = ""
//     NoOrders                    = 0
//     ParentServerOrderID         = ""
//     OCOLinkedOrderServerOrderID = ""
//     OpenOrClose                 = TRADE_UNSET  (0)
//     PreviousClientOrderID       = ""
//     FreeFormText                = ""
//     OrderReceivedDateTime       = 0
//     LatestTransactionDateTime   = 0
// }
func (m OrderUpdateBuilder) Clear() {
	m.Unsafe().SetBytes(0, _OrderUpdateDefault)
}

// Clear
// {
//     Size                        = OrderUpdateFixedSize  (688)
//     Type                        = ORDER_UPDATE  (301)
//     RequestID                   = 0
//     TotalNumMessages            = 0
//     MessageNumber               = 0
//     Symbol                      = ""
//     Exchange                    = ""
//     PreviousServerOrderID       = ""
//     ServerOrderID               = ""
//     ClientOrderID               = ""
//     ExchangeOrderID             = ""
//     OrderStatus                 = ORDER_STATUS_UNSPECIFIED  (0)
//     OrderUpdateReason           = ORDER_UPDATE_REASON_UNSET  (0)
//     OrderType                   = ORDER_TYPE_UNSET  (0)
//     BuySell                     = BUY_SELL_UNSET  (0)
//     Price1                      = math.MaxFloat64
//     Price2                      = math.MaxFloat64
//     TimeInForce                 = TIF_UNSET  (0)
//     GoodTillDateTime            = 0
//     OrderQuantity               = math.MaxFloat64
//     FilledQuantity              = math.MaxFloat64
//     RemainingQuantity           = math.MaxFloat64
//     AverageFillPrice            = math.MaxFloat64
//     LastFillPrice               = math.MaxFloat64
//     LastFillDateTime            = 0
//     LastFillQuantity            = math.MaxFloat64
//     LastFillExecutionID         = ""
//     TradeAccount                = ""
//     InfoText                    = ""
//     NoOrders                    = 0
//     ParentServerOrderID         = ""
//     OCOLinkedOrderServerOrderID = ""
//     OpenOrClose                 = TRADE_UNSET  (0)
//     PreviousClientOrderID       = ""
//     FreeFormText                = ""
//     OrderReceivedDateTime       = 0
//     LatestTransactionDateTime   = 0
// }
func (m OrderUpdateFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _OrderUpdateFixedDefault)
}

// Clear
// {
//     Size                        = OrderUpdateSize  (208)
//     Type                        = ORDER_UPDATE  (301)
//     BaseSize                    = OrderUpdateSize  (208)
//     RequestID                   = 0
//     TotalNumMessages            = 0
//     MessageNumber               = 0
//     Symbol                      = ""
//     Exchange                    = ""
//     PreviousServerOrderID       = ""
//     ServerOrderID               = ""
//     ClientOrderID               = ""
//     ExchangeOrderID             = ""
//     OrderStatus                 = ORDER_STATUS_UNSPECIFIED  (0)
//     OrderUpdateReason           = ORDER_UPDATE_REASON_UNSET  (0)
//     OrderType                   = ORDER_TYPE_UNSET  (0)
//     BuySell                     = BUY_SELL_UNSET  (0)
//     Price1                      = math.MaxFloat64
//     Price2                      = math.MaxFloat64
//     TimeInForce                 = TIF_UNSET  (0)
//     GoodTillDateTime            = 0
//     OrderQuantity               = math.MaxFloat64
//     FilledQuantity              = math.MaxFloat64
//     RemainingQuantity           = math.MaxFloat64
//     AverageFillPrice            = math.MaxFloat64
//     LastFillPrice               = math.MaxFloat64
//     LastFillDateTime            = 0
//     LastFillQuantity            = math.MaxFloat64
//     LastFillExecutionID         = ""
//     TradeAccount                = ""
//     InfoText                    = ""
//     NoOrders                    = 0
//     ParentServerOrderID         = ""
//     OCOLinkedOrderServerOrderID = ""
//     OpenOrClose                 = TRADE_UNSET  (0)
//     PreviousClientOrderID       = ""
//     FreeFormText                = ""
//     OrderReceivedDateTime       = 0
//     LatestTransactionDateTime   = 0
// }
func (m OrderUpdatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _OrderUpdateDefault)
}

// Clear
// {
//     Size                        = OrderUpdateFixedSize  (688)
//     Type                        = ORDER_UPDATE  (301)
//     RequestID                   = 0
//     TotalNumMessages            = 0
//     MessageNumber               = 0
//     Symbol                      = ""
//     Exchange                    = ""
//     PreviousServerOrderID       = ""
//     ServerOrderID               = ""
//     ClientOrderID               = ""
//     ExchangeOrderID             = ""
//     OrderStatus                 = ORDER_STATUS_UNSPECIFIED  (0)
//     OrderUpdateReason           = ORDER_UPDATE_REASON_UNSET  (0)
//     OrderType                   = ORDER_TYPE_UNSET  (0)
//     BuySell                     = BUY_SELL_UNSET  (0)
//     Price1                      = math.MaxFloat64
//     Price2                      = math.MaxFloat64
//     TimeInForce                 = TIF_UNSET  (0)
//     GoodTillDateTime            = 0
//     OrderQuantity               = math.MaxFloat64
//     FilledQuantity              = math.MaxFloat64
//     RemainingQuantity           = math.MaxFloat64
//     AverageFillPrice            = math.MaxFloat64
//     LastFillPrice               = math.MaxFloat64
//     LastFillDateTime            = 0
//     LastFillQuantity            = math.MaxFloat64
//     LastFillExecutionID         = ""
//     TradeAccount                = ""
//     InfoText                    = ""
//     NoOrders                    = 0
//     ParentServerOrderID         = ""
//     OCOLinkedOrderServerOrderID = ""
//     OpenOrClose                 = TRADE_UNSET  (0)
//     PreviousClientOrderID       = ""
//     FreeFormText                = ""
//     OrderReceivedDateTime       = 0
//     LatestTransactionDateTime   = 0
// }
func (m OrderUpdateFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _OrderUpdateFixedDefault)
}

// ToBuilder
func (m OrderUpdate) ToBuilder() OrderUpdateBuilder {
	return OrderUpdateBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m OrderUpdateFixed) ToBuilder() OrderUpdateFixedBuilder {
	return OrderUpdateFixedBuilder{m.Fixed}
}

// ToBuilder
func (m OrderUpdatePointer) ToBuilder() OrderUpdatePointerBuilder {
	return OrderUpdatePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m OrderUpdateFixedPointer) ToBuilder() OrderUpdateFixedPointerBuilder {
	return OrderUpdateFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m OrderUpdateBuilder) Finish() OrderUpdate {
	return OrderUpdate{m.VLSBuilder.Finish()}
}

// Finish
func (m OrderUpdateFixedBuilder) Finish() OrderUpdateFixed {
	return OrderUpdateFixed{m.Fixed}
}

// Finish
func (m *OrderUpdatePointerBuilder) Finish() OrderUpdatePointer {
	return OrderUpdatePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *OrderUpdateFixedPointerBuilder) Finish() OrderUpdateFixedPointer {
	return OrderUpdateFixedPointer{m.FixedPointer}
}

// RequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdate) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdateBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdatePointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdatePointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// TotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdate) TotalNumMessages() int32 {
	return message.Int32VLS(m.Unsafe(), 16, 12)
}

// TotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdateBuilder) TotalNumMessages() int32 {
	return message.Int32VLS(m.Unsafe(), 16, 12)
}

// TotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdatePointer) TotalNumMessages() int32 {
	return message.Int32VLS(m.Ptr, 16, 12)
}

// TotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdatePointerBuilder) TotalNumMessages() int32 {
	return message.Int32VLS(m.Ptr, 16, 12)
}

// MessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdate) MessageNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 20, 16)
}

// MessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdateBuilder) MessageNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 20, 16)
}

// MessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdatePointer) MessageNumber() int32 {
	return message.Int32VLS(m.Ptr, 20, 16)
}

// MessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdatePointerBuilder) MessageNumber() int32 {
	return message.Int32VLS(m.Ptr, 20, 16)
}

// Symbol The symbol for the order.
func (m OrderUpdate) Symbol() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Symbol The symbol for the order.
func (m OrderUpdateBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Symbol The symbol for the order.
func (m OrderUpdatePointer) Symbol() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Symbol The symbol for the order.
func (m OrderUpdatePointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Exchange The optional exchange for the symbol.
func (m OrderUpdate) Exchange() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Exchange The optional exchange for the symbol.
func (m OrderUpdateBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Exchange The optional exchange for the symbol.
func (m OrderUpdatePointer) Exchange() string {
	return message.StringVLS(m.Ptr, 28, 24)
}

// Exchange The optional exchange for the symbol.
func (m OrderUpdatePointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 28, 24)
}

// PreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdate) PreviousServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 32, 28)
}

// PreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdateBuilder) PreviousServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 32, 28)
}

// PreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdatePointer) PreviousServerOrderID() string {
	return message.StringVLS(m.Ptr, 32, 28)
}

// PreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdatePointerBuilder) PreviousServerOrderID() string {
	return message.StringVLS(m.Ptr, 32, 28)
}

// ServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdate) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
}

// ServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdateBuilder) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
}

// ServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdatePointer) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 36, 32)
}

// ServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdatePointerBuilder) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 36, 32)
}

// ClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdate) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// ClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdateBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// ClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdatePointer) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 40, 36)
}

// ClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdatePointerBuilder) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 40, 36)
}

// ExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdate) ExchangeOrderID() string {
	return message.StringVLS(m.Unsafe(), 44, 40)
}

// ExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdateBuilder) ExchangeOrderID() string {
	return message.StringVLS(m.Unsafe(), 44, 40)
}

// ExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdatePointer) ExchangeOrderID() string {
	return message.StringVLS(m.Ptr, 44, 40)
}

// ExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdatePointerBuilder) ExchangeOrderID() string {
	return message.StringVLS(m.Ptr, 44, 40)
}

// OrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdate) OrderStatus() OrderStatusEnum {
	return OrderStatusEnum(message.Int32VLS(m.Unsafe(), 48, 44))
}

// OrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdateBuilder) OrderStatus() OrderStatusEnum {
	return OrderStatusEnum(message.Int32VLS(m.Unsafe(), 48, 44))
}

// OrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdatePointer) OrderStatus() OrderStatusEnum {
	return OrderStatusEnum(message.Int32VLS(m.Ptr, 48, 44))
}

// OrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdatePointerBuilder) OrderStatus() OrderStatusEnum {
	return OrderStatusEnum(message.Int32VLS(m.Ptr, 48, 44))
}

// OrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdate) OrderUpdateReason() OrderUpdateReasonEnum {
	return OrderUpdateReasonEnum(message.Int32VLS(m.Unsafe(), 52, 48))
}

// OrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdateBuilder) OrderUpdateReason() OrderUpdateReasonEnum {
	return OrderUpdateReasonEnum(message.Int32VLS(m.Unsafe(), 52, 48))
}

// OrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdatePointer) OrderUpdateReason() OrderUpdateReasonEnum {
	return OrderUpdateReasonEnum(message.Int32VLS(m.Ptr, 52, 48))
}

// OrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdatePointerBuilder) OrderUpdateReason() OrderUpdateReasonEnum {
	return OrderUpdateReasonEnum(message.Int32VLS(m.Ptr, 52, 48))
}

// OrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdate) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 56, 52))
}

// OrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdateBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 56, 52))
}

// OrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdatePointer) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 56, 52))
}

// OrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdatePointerBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 56, 52))
}

// BuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdate) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 60, 56))
}

// BuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdateBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 60, 56))
}

// BuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdatePointer) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 60, 56))
}

// BuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdatePointerBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 60, 56))
}

// Price1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdate) Price1() float64 {
	return message.Float64VLS(m.Unsafe(), 72, 64)
}

// Price1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) Price1() float64 {
	return message.Float64VLS(m.Unsafe(), 72, 64)
}

// Price1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointer) Price1() float64 {
	return message.Float64VLS(m.Ptr, 72, 64)
}

// Price1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) Price1() float64 {
	return message.Float64VLS(m.Ptr, 72, 64)
}

// Price2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdate) Price2() float64 {
	return message.Float64VLS(m.Unsafe(), 80, 72)
}

// Price2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) Price2() float64 {
	return message.Float64VLS(m.Unsafe(), 80, 72)
}

// Price2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointer) Price2() float64 {
	return message.Float64VLS(m.Ptr, 80, 72)
}

// Price2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) Price2() float64 {
	return message.Float64VLS(m.Ptr, 80, 72)
}

// TimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdate) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Unsafe(), 84, 80))
}

// TimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdateBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Unsafe(), 84, 80))
}

// TimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdatePointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Ptr, 84, 80))
}

// TimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdatePointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Ptr, 84, 80))
}

// GoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdate) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 96, 88))
}

// GoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdateBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 96, 88))
}

// GoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdatePointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 96, 88))
}

// GoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdatePointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 96, 88))
}

// OrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdate) OrderQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 104, 96)
}

// OrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) OrderQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 104, 96)
}

// OrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointer) OrderQuantity() float64 {
	return message.Float64VLS(m.Ptr, 104, 96)
}

// OrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) OrderQuantity() float64 {
	return message.Float64VLS(m.Ptr, 104, 96)
}

// FilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdate) FilledQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 112, 104)
}

// FilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) FilledQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 112, 104)
}

// FilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointer) FilledQuantity() float64 {
	return message.Float64VLS(m.Ptr, 112, 104)
}

// FilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) FilledQuantity() float64 {
	return message.Float64VLS(m.Ptr, 112, 104)
}

// RemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdate) RemainingQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 120, 112)
}

// RemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) RemainingQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 120, 112)
}

// RemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointer) RemainingQuantity() float64 {
	return message.Float64VLS(m.Ptr, 120, 112)
}

// RemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) RemainingQuantity() float64 {
	return message.Float64VLS(m.Ptr, 120, 112)
}

// AverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdate) AverageFillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 128, 120)
}

// AverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) AverageFillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 128, 120)
}

// AverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointer) AverageFillPrice() float64 {
	return message.Float64VLS(m.Ptr, 128, 120)
}

// AverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) AverageFillPrice() float64 {
	return message.Float64VLS(m.Ptr, 128, 120)
}

// LastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdate) LastFillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 136, 128)
}

// LastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) LastFillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 136, 128)
}

// LastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointer) LastFillPrice() float64 {
	return message.Float64VLS(m.Ptr, 136, 128)
}

// LastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) LastFillPrice() float64 {
	return message.Float64VLS(m.Ptr, 136, 128)
}

// LastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdate) LastFillDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 144, 136))
}

// LastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateBuilder) LastFillDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 144, 136))
}

// LastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdatePointer) LastFillDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Ptr, 144, 136))
}

// LastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdatePointerBuilder) LastFillDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Ptr, 144, 136))
}

// LastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdate) LastFillQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 152, 144)
}

// LastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) LastFillQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 152, 144)
}

// LastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointer) LastFillQuantity() float64 {
	return message.Float64VLS(m.Ptr, 152, 144)
}

// LastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) LastFillQuantity() float64 {
	return message.Float64VLS(m.Ptr, 152, 144)
}

// LastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdate) LastFillExecutionID() string {
	return message.StringVLS(m.Unsafe(), 156, 152)
}

// LastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateBuilder) LastFillExecutionID() string {
	return message.StringVLS(m.Unsafe(), 156, 152)
}

// LastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdatePointer) LastFillExecutionID() string {
	return message.StringVLS(m.Ptr, 156, 152)
}

// LastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdatePointerBuilder) LastFillExecutionID() string {
	return message.StringVLS(m.Ptr, 156, 152)
}

// TradeAccount The trade account the order belongs to.
func (m OrderUpdate) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 160, 156)
}

// TradeAccount The trade account the order belongs to.
func (m OrderUpdateBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 160, 156)
}

// TradeAccount The trade account the order belongs to.
func (m OrderUpdatePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 160, 156)
}

// TradeAccount The trade account the order belongs to.
func (m OrderUpdatePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 160, 156)
}

// InfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdate) InfoText() string {
	return message.StringVLS(m.Unsafe(), 164, 160)
}

// InfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdateBuilder) InfoText() string {
	return message.StringVLS(m.Unsafe(), 164, 160)
}

// InfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdatePointer) InfoText() string {
	return message.StringVLS(m.Ptr, 164, 160)
}

// InfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdatePointerBuilder) InfoText() string {
	return message.StringVLS(m.Ptr, 164, 160)
}

// NoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdate) NoOrders() uint8 {
	return message.Uint8VLS(m.Unsafe(), 165, 164)
}

// NoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdateBuilder) NoOrders() uint8 {
	return message.Uint8VLS(m.Unsafe(), 165, 164)
}

// NoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdatePointer) NoOrders() uint8 {
	return message.Uint8VLS(m.Ptr, 165, 164)
}

// NoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdatePointerBuilder) NoOrders() uint8 {
	return message.Uint8VLS(m.Ptr, 165, 164)
}

// ParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdate) ParentServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 170, 166)
}

// ParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdateBuilder) ParentServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 170, 166)
}

// ParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdatePointer) ParentServerOrderID() string {
	return message.StringVLS(m.Ptr, 170, 166)
}

// ParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdatePointerBuilder) ParentServerOrderID() string {
	return message.StringVLS(m.Ptr, 170, 166)
}

// OCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdate) OCOLinkedOrderServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 174, 170)
}

// OCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdateBuilder) OCOLinkedOrderServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 174, 170)
}

// OCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdatePointer) OCOLinkedOrderServerOrderID() string {
	return message.StringVLS(m.Ptr, 174, 170)
}

// OCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdatePointerBuilder) OCOLinkedOrderServerOrderID() string {
	return message.StringVLS(m.Ptr, 174, 170)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdate) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 180, 176))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdateBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 180, 176))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdatePointer) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Ptr, 180, 176))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdatePointerBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Ptr, 180, 176))
}

// PreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdate) PreviousClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 184, 180)
}

// PreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdateBuilder) PreviousClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 184, 180)
}

// PreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdatePointer) PreviousClientOrderID() string {
	return message.StringVLS(m.Ptr, 184, 180)
}

// PreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdatePointerBuilder) PreviousClientOrderID() string {
	return message.StringVLS(m.Ptr, 184, 180)
}

// FreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdate) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 188, 184)
}

// FreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdateBuilder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 188, 184)
}

// FreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdatePointer) FreeFormText() string {
	return message.StringVLS(m.Ptr, 188, 184)
}

// FreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdatePointerBuilder) FreeFormText() string {
	return message.StringVLS(m.Ptr, 188, 184)
}

// OrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdate) OrderReceivedDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 200, 192))
}

// OrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdateBuilder) OrderReceivedDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 200, 192))
}

// OrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdatePointer) OrderReceivedDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Ptr, 200, 192))
}

// OrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdatePointerBuilder) OrderReceivedDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Ptr, 200, 192))
}

// LatestTransactionDateTime
func (m OrderUpdate) LatestTransactionDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 208, 200))
}

// LatestTransactionDateTime
func (m OrderUpdateBuilder) LatestTransactionDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Unsafe(), 208, 200))
}

// LatestTransactionDateTime
func (m OrderUpdatePointer) LatestTransactionDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Ptr, 208, 200))
}

// LatestTransactionDateTime
func (m OrderUpdatePointerBuilder) LatestTransactionDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Ptr, 208, 200))
}

// RequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdateFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdateFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdateFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdateFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// TotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdateFixed) TotalNumMessages() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// TotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdateFixedBuilder) TotalNumMessages() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// TotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdateFixedPointer) TotalNumMessages() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// TotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdateFixedPointerBuilder) TotalNumMessages() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// MessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdateFixed) MessageNumber() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// MessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdateFixedBuilder) MessageNumber() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// MessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdateFixedPointer) MessageNumber() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// MessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdateFixedPointerBuilder) MessageNumber() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// Symbol The symbol for the order.
func (m OrderUpdateFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 80, 16)
}

// Symbol The symbol for the order.
func (m OrderUpdateFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 80, 16)
}

// Symbol The symbol for the order.
func (m OrderUpdateFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 80, 16)
}

// Symbol The symbol for the order.
func (m OrderUpdateFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 80, 16)
}

// Exchange The optional exchange for the symbol.
func (m OrderUpdateFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 96, 80)
}

// Exchange The optional exchange for the symbol.
func (m OrderUpdateFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 96, 80)
}

// Exchange The optional exchange for the symbol.
func (m OrderUpdateFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 96, 80)
}

// Exchange The optional exchange for the symbol.
func (m OrderUpdateFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 96, 80)
}

// PreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdateFixed) PreviousServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 128, 96)
}

// PreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdateFixedBuilder) PreviousServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 128, 96)
}

// PreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdateFixedPointer) PreviousServerOrderID() string {
	return message.StringFixed(m.Ptr, 128, 96)
}

// PreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdateFixedPointerBuilder) PreviousServerOrderID() string {
	return message.StringFixed(m.Ptr, 128, 96)
}

// ServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdateFixed) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 160, 128)
}

// ServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdateFixedBuilder) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 160, 128)
}

// ServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdateFixedPointer) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 160, 128)
}

// ServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdateFixedPointerBuilder) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 160, 128)
}

// ClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdateFixed) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 192, 160)
}

// ClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdateFixedBuilder) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 192, 160)
}

// ClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdateFixedPointer) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 192, 160)
}

// ClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdateFixedPointerBuilder) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 192, 160)
}

// ExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdateFixed) ExchangeOrderID() string {
	return message.StringFixed(m.Unsafe(), 224, 192)
}

// ExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdateFixedBuilder) ExchangeOrderID() string {
	return message.StringFixed(m.Unsafe(), 224, 192)
}

// ExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdateFixedPointer) ExchangeOrderID() string {
	return message.StringFixed(m.Ptr, 224, 192)
}

// ExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdateFixedPointerBuilder) ExchangeOrderID() string {
	return message.StringFixed(m.Ptr, 224, 192)
}

// OrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdateFixed) OrderStatus() OrderStatusEnum {
	return OrderStatusEnum(message.Int32Fixed(m.Unsafe(), 228, 224))
}

// OrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdateFixedBuilder) OrderStatus() OrderStatusEnum {
	return OrderStatusEnum(message.Int32Fixed(m.Unsafe(), 228, 224))
}

// OrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdateFixedPointer) OrderStatus() OrderStatusEnum {
	return OrderStatusEnum(message.Int32Fixed(m.Ptr, 228, 224))
}

// OrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdateFixedPointerBuilder) OrderStatus() OrderStatusEnum {
	return OrderStatusEnum(message.Int32Fixed(m.Ptr, 228, 224))
}

// OrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdateFixed) OrderUpdateReason() OrderUpdateReasonEnum {
	return OrderUpdateReasonEnum(message.Int32Fixed(m.Unsafe(), 232, 228))
}

// OrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdateFixedBuilder) OrderUpdateReason() OrderUpdateReasonEnum {
	return OrderUpdateReasonEnum(message.Int32Fixed(m.Unsafe(), 232, 228))
}

// OrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdateFixedPointer) OrderUpdateReason() OrderUpdateReasonEnum {
	return OrderUpdateReasonEnum(message.Int32Fixed(m.Ptr, 232, 228))
}

// OrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdateFixedPointerBuilder) OrderUpdateReason() OrderUpdateReasonEnum {
	return OrderUpdateReasonEnum(message.Int32Fixed(m.Ptr, 232, 228))
}

// OrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdateFixed) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 236, 232))
}

// OrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdateFixedBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 236, 232))
}

// OrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdateFixedPointer) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 236, 232))
}

// OrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdateFixedPointerBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 236, 232))
}

// BuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdateFixed) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 240, 236))
}

// BuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdateFixedBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 240, 236))
}

// BuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdateFixedPointer) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 240, 236))
}

// BuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdateFixedPointerBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 240, 236))
}

// Price1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixed) Price1() float64 {
	return message.Float64Fixed(m.Unsafe(), 248, 240)
}

// Price1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) Price1() float64 {
	return message.Float64Fixed(m.Unsafe(), 248, 240)
}

// Price1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) Price1() float64 {
	return message.Float64Fixed(m.Ptr, 248, 240)
}

// Price1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) Price1() float64 {
	return message.Float64Fixed(m.Ptr, 248, 240)
}

// Price2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixed) Price2() float64 {
	return message.Float64Fixed(m.Unsafe(), 256, 248)
}

// Price2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) Price2() float64 {
	return message.Float64Fixed(m.Unsafe(), 256, 248)
}

// Price2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) Price2() float64 {
	return message.Float64Fixed(m.Ptr, 256, 248)
}

// Price2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) Price2() float64 {
	return message.Float64Fixed(m.Ptr, 256, 248)
}

// TimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdateFixed) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Unsafe(), 260, 256))
}

// TimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdateFixedBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Unsafe(), 260, 256))
}

// TimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdateFixedPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 260, 256))
}

// TimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdateFixedPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 260, 256))
}

// GoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdateFixed) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 272, 264))
}

// GoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdateFixedBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 272, 264))
}

// GoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdateFixedPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 272, 264))
}

// GoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdateFixedPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 272, 264))
}

// OrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixed) OrderQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 280, 272)
}

// OrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) OrderQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 280, 272)
}

// OrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) OrderQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 280, 272)
}

// OrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) OrderQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 280, 272)
}

// FilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixed) FilledQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 288, 280)
}

// FilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) FilledQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 288, 280)
}

// FilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) FilledQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 288, 280)
}

// FilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) FilledQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 288, 280)
}

// RemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixed) RemainingQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 296, 288)
}

// RemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) RemainingQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 296, 288)
}

// RemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) RemainingQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 296, 288)
}

// RemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) RemainingQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 296, 288)
}

// AverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixed) AverageFillPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 304, 296)
}

// AverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) AverageFillPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 304, 296)
}

// AverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) AverageFillPrice() float64 {
	return message.Float64Fixed(m.Ptr, 304, 296)
}

// AverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) AverageFillPrice() float64 {
	return message.Float64Fixed(m.Ptr, 304, 296)
}

// LastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixed) LastFillPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 312, 304)
}

// LastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) LastFillPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 312, 304)
}

// LastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) LastFillPrice() float64 {
	return message.Float64Fixed(m.Ptr, 312, 304)
}

// LastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) LastFillPrice() float64 {
	return message.Float64Fixed(m.Ptr, 312, 304)
}

// LastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixed) LastFillDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 320, 312))
}

// LastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedBuilder) LastFillDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 320, 312))
}

// LastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedPointer) LastFillDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 320, 312))
}

// LastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedPointerBuilder) LastFillDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 320, 312))
}

// LastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixed) LastFillQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 328, 320)
}

// LastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) LastFillQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 328, 320)
}

// LastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointer) LastFillQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 328, 320)
}

// LastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) LastFillQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 328, 320)
}

// LastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixed) LastFillExecutionID() string {
	return message.StringFixed(m.Unsafe(), 392, 328)
}

// LastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedBuilder) LastFillExecutionID() string {
	return message.StringFixed(m.Unsafe(), 392, 328)
}

// LastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedPointer) LastFillExecutionID() string {
	return message.StringFixed(m.Ptr, 392, 328)
}

// LastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedPointerBuilder) LastFillExecutionID() string {
	return message.StringFixed(m.Ptr, 392, 328)
}

// TradeAccount The trade account the order belongs to.
func (m OrderUpdateFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 424, 392)
}

// TradeAccount The trade account the order belongs to.
func (m OrderUpdateFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 424, 392)
}

// TradeAccount The trade account the order belongs to.
func (m OrderUpdateFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 424, 392)
}

// TradeAccount The trade account the order belongs to.
func (m OrderUpdateFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 424, 392)
}

// InfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdateFixed) InfoText() string {
	return message.StringFixed(m.Unsafe(), 520, 424)
}

// InfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdateFixedBuilder) InfoText() string {
	return message.StringFixed(m.Unsafe(), 520, 424)
}

// InfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdateFixedPointer) InfoText() string {
	return message.StringFixed(m.Ptr, 520, 424)
}

// InfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdateFixedPointerBuilder) InfoText() string {
	return message.StringFixed(m.Ptr, 520, 424)
}

// NoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdateFixed) NoOrders() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 521, 520)
}

// NoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdateFixedBuilder) NoOrders() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 521, 520)
}

// NoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdateFixedPointer) NoOrders() uint8 {
	return message.Uint8Fixed(m.Ptr, 521, 520)
}

// NoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdateFixedPointerBuilder) NoOrders() uint8 {
	return message.Uint8Fixed(m.Ptr, 521, 520)
}

// ParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixed) ParentServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 553, 521)
}

// ParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedBuilder) ParentServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 553, 521)
}

// ParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedPointer) ParentServerOrderID() string {
	return message.StringFixed(m.Ptr, 553, 521)
}

// ParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedPointerBuilder) ParentServerOrderID() string {
	return message.StringFixed(m.Ptr, 553, 521)
}

// OCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixed) OCOLinkedOrderServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 585, 553)
}

// OCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedBuilder) OCOLinkedOrderServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 585, 553)
}

// OCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedPointer) OCOLinkedOrderServerOrderID() string {
	return message.StringFixed(m.Ptr, 585, 553)
}

// OCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedPointerBuilder) OCOLinkedOrderServerOrderID() string {
	return message.StringFixed(m.Ptr, 585, 553)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdateFixed) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Unsafe(), 592, 588))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdateFixedBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Unsafe(), 592, 588))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdateFixedPointer) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 592, 588))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdateFixedPointerBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 592, 588))
}

// PreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdateFixed) PreviousClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 624, 592)
}

// PreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdateFixedBuilder) PreviousClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 624, 592)
}

// PreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdateFixedPointer) PreviousClientOrderID() string {
	return message.StringFixed(m.Ptr, 624, 592)
}

// PreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdateFixedPointerBuilder) PreviousClientOrderID() string {
	return message.StringFixed(m.Ptr, 624, 592)
}

// FreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdateFixed) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 672, 624)
}

// FreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdateFixedBuilder) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 672, 624)
}

// FreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdateFixedPointer) FreeFormText() string {
	return message.StringFixed(m.Ptr, 672, 624)
}

// FreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdateFixedPointerBuilder) FreeFormText() string {
	return message.StringFixed(m.Ptr, 672, 624)
}

// OrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdateFixed) OrderReceivedDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 680, 672))
}

// OrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdateFixedBuilder) OrderReceivedDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 680, 672))
}

// OrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdateFixedPointer) OrderReceivedDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 680, 672))
}

// OrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdateFixedPointerBuilder) OrderReceivedDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 680, 672))
}

// LatestTransactionDateTime
func (m OrderUpdateFixed) LatestTransactionDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 688, 680))
}

// LatestTransactionDateTime
func (m OrderUpdateFixedBuilder) LatestTransactionDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 688, 680))
}

// LatestTransactionDateTime
func (m OrderUpdateFixedPointer) LatestTransactionDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 688, 680))
}

// LatestTransactionDateTime
func (m OrderUpdateFixedPointerBuilder) LatestTransactionDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 688, 680))
}

// SetRequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdateBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdatePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetTotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdateBuilder) SetTotalNumMessages(value int32) {
	message.SetInt32VLS(m.Unsafe(), 16, 12, value)
}

// SetTotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdatePointerBuilder) SetTotalNumMessages(value int32) {
	message.SetInt32VLS(m.Ptr, 16, 12, value)
}

// SetMessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdateBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 20, 16, value)
}

// SetMessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdatePointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Ptr, 20, 16, value)
}

// SetSymbol The symbol for the order.
func (m *OrderUpdateBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetSymbol The symbol for the order.
func (m *OrderUpdatePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetExchange The optional exchange for the symbol.
func (m *OrderUpdateBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 28, 24, value)
}

// SetExchange The optional exchange for the symbol.
func (m *OrderUpdatePointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 28, 24, value)
}

// SetPreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m *OrderUpdateBuilder) SetPreviousServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 32, 28, value)
}

// SetPreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m *OrderUpdatePointerBuilder) SetPreviousServerOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 32, 28, value)
}

// SetServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m *OrderUpdateBuilder) SetServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 36, 32, value)
}

// SetServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m *OrderUpdatePointerBuilder) SetServerOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 36, 32, value)
}

// SetClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m *OrderUpdateBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 40, 36, value)
}

// SetClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m *OrderUpdatePointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 40, 36, value)
}

// SetExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m *OrderUpdateBuilder) SetExchangeOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 44, 40, value)
}

// SetExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m *OrderUpdatePointerBuilder) SetExchangeOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 44, 40, value)
}

// SetOrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdateBuilder) SetOrderStatus(value OrderStatusEnum) {
	message.SetInt32VLS(m.Unsafe(), 48, 44, int32(value))
}

// SetOrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdatePointerBuilder) SetOrderStatus(value OrderStatusEnum) {
	message.SetInt32VLS(m.Ptr, 48, 44, int32(value))
}

// SetOrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdateBuilder) SetOrderUpdateReason(value OrderUpdateReasonEnum) {
	message.SetInt32VLS(m.Unsafe(), 52, 48, int32(value))
}

// SetOrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdatePointerBuilder) SetOrderUpdateReason(value OrderUpdateReasonEnum) {
	message.SetInt32VLS(m.Ptr, 52, 48, int32(value))
}

// SetOrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdateBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 56, 52, int32(value))
}

// SetOrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdatePointerBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32VLS(m.Ptr, 56, 52, int32(value))
}

// SetBuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdateBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 60, 56, int32(value))
}

// SetBuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdatePointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 60, 56, int32(value))
}

// SetPrice1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) SetPrice1(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 72, 64, value)
}

// SetPrice1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) SetPrice1(value float64) {
	message.SetFloat64VLS(m.Ptr, 72, 64, value)
}

// SetPrice2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) SetPrice2(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 80, 72, value)
}

// SetPrice2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) SetPrice2(value float64) {
	message.SetFloat64VLS(m.Ptr, 80, 72, value)
}

// SetTimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdateBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Unsafe(), 84, 80, int32(value))
}

// SetTimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdatePointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Ptr, 84, 80, int32(value))
}

// SetGoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdateBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 96, 88, int64(value))
}

// SetGoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdatePointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 96, 88, int64(value))
}

// SetOrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) SetOrderQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 104, 96, value)
}

// SetOrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) SetOrderQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 104, 96, value)
}

// SetFilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) SetFilledQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 112, 104, value)
}

// SetFilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) SetFilledQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 112, 104, value)
}

// SetRemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) SetRemainingQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 120, 112, value)
}

// SetRemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) SetRemainingQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 120, 112, value)
}

// SetAverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) SetAverageFillPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 128, 120, value)
}

// SetAverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) SetAverageFillPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 128, 120, value)
}

// SetLastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) SetLastFillPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 136, 128, value)
}

// SetLastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) SetLastFillPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 136, 128, value)
}

// SetLastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateBuilder) SetLastFillDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 144, 136, int64(value))
}

// SetLastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdatePointerBuilder) SetLastFillDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Ptr, 144, 136, int64(value))
}

// SetLastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateBuilder) SetLastFillQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 152, 144, value)
}

// SetLastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdatePointerBuilder) SetLastFillQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 152, 144, value)
}

// SetLastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m *OrderUpdateBuilder) SetLastFillExecutionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 156, 152, value)
}

// SetLastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m *OrderUpdatePointerBuilder) SetLastFillExecutionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 156, 152, value)
}

// SetTradeAccount The trade account the order belongs to.
func (m *OrderUpdateBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 160, 156, value)
}

// SetTradeAccount The trade account the order belongs to.
func (m *OrderUpdatePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 160, 156, value)
}

// SetInfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m *OrderUpdateBuilder) SetInfoText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 164, 160, value)
}

// SetInfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m *OrderUpdatePointerBuilder) SetInfoText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 164, 160, value)
}

// SetNoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdateBuilder) SetNoOrders(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 165, 164, value)
}

// SetNoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdatePointerBuilder) SetNoOrders(value uint8) {
	message.SetUint8VLS(m.Ptr, 165, 164, value)
}

// SetParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m *OrderUpdateBuilder) SetParentServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 170, 166, value)
}

// SetParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m *OrderUpdatePointerBuilder) SetParentServerOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 170, 166, value)
}

// SetOCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m *OrderUpdateBuilder) SetOCOLinkedOrderServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 174, 170, value)
}

// SetOCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m *OrderUpdatePointerBuilder) SetOCOLinkedOrderServerOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 174, 170, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdateBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Unsafe(), 180, 176, int32(value))
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdatePointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Ptr, 180, 176, int32(value))
}

// SetPreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m *OrderUpdateBuilder) SetPreviousClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 184, 180, value)
}

// SetPreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m *OrderUpdatePointerBuilder) SetPreviousClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 184, 180, value)
}

// SetFreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m *OrderUpdateBuilder) SetFreeFormText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 188, 184, value)
}

// SetFreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m *OrderUpdatePointerBuilder) SetFreeFormText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 188, 184, value)
}

// SetOrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdateBuilder) SetOrderReceivedDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 200, 192, int64(value))
}

// SetOrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdatePointerBuilder) SetOrderReceivedDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Ptr, 200, 192, int64(value))
}

// SetLatestTransactionDateTime
func (m OrderUpdateBuilder) SetLatestTransactionDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Unsafe(), 208, 200, float64(value))
}

// SetLatestTransactionDateTime
func (m OrderUpdatePointerBuilder) SetLatestTransactionDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Ptr, 208, 200, float64(value))
}

// SetRequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdateFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID Set to 0 unless this is in response to an OpenOrdersRequest, in which
// case this must be set to the RequestID given in the OpenOrdersRequest.
//
// If this OrderUpdate is unsolicited, for example a real-time fill or other
// unsolicited order event, the Server must leave this at 0.
func (m OrderUpdateFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdateFixedBuilder) SetTotalNumMessages(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetTotalNumMessages This indicates the total number of OrderUpdate messages when a batch of
// reports is being sent in response to an OpenOrdersRequest. If there is
// only one order being sent, this will be 1. The Server must use a value
// of 1 for an unsolicited report. A Client should not rely on this field
// for an unsolicited report.
func (m OrderUpdateFixedPointerBuilder) SetTotalNumMessages(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetMessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdateFixedBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 16, 12, value)
}

// SetMessageNumber This indicates the 1-based index of the OrderUpdate message when a batch
// of reports is being sent in response to an OpenOrdersRequest. If there
// is only one order being sent, this will be 1. Use a value of 1 for an
// unsolicited report. A Client should not rely on this field for an unsolicited
// report.
func (m OrderUpdateFixedPointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Ptr, 16, 12, value)
}

// SetSymbol The symbol for the order.
func (m OrderUpdateFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 80, 16, value)
}

// SetSymbol The symbol for the order.
func (m OrderUpdateFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 80, 16, value)
}

// SetExchange The optional exchange for the symbol.
func (m OrderUpdateFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 96, 80, value)
}

// SetExchange The optional exchange for the symbol.
func (m OrderUpdateFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 96, 80, value)
}

// SetPreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdateFixedBuilder) SetPreviousServerOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 128, 96, value)
}

// SetPreviousServerOrderID Used upon a Cancel and Replace operation (ORDER_CANCEL_REPLACE_COMPLETE)
// where a new Server Order identifier is given.
//
// In this case this field needs to be set to the previous Server Order identifier.
// In this case this field needs to be set to the previous Server Order identifier.
//
// This should be left at the default setting of empty in the case where
// the Server does not change the Server Order identifier upon a Cancel and
// Replace operation.
func (m OrderUpdateFixedPointerBuilder) SetPreviousServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 128, 96, value)
}

// SetServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdateFixedBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 160, 128, value)
}

// SetServerOrderID The ServerOrderID is set by the server and uniquely identifies the order.
// When a new order is submitted by the Client and the Server responds with
// an OrderUpdate, this field needs to be set to the order identifier which
// is good for the life of the order.
//
// This ServerOrderID can optionally change on a Cancel and Replace operation.
// In this case, the PreviousServerOrderID field will contain the previous
// ServerOrderID upon a ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReason.
//
// ServerOrderID must always be set except it is not required in the cases
// when the OrderUpdateReason is one of the following: NEW_ORDER_REJECTED,
// ORDER_CANCEL_REJECTED, ORDER_CANCEL_REPLACE_REJECTED. If it is not set,
// then the ClientOrderID must be set.
func (m OrderUpdateFixedPointerBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 160, 128, value)
}

// SetClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdateFixedBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 192, 160, value)
}

// SetClientOrderID The ClientOrderID is the order identifier provided by the Client. When
// the Client submits a new order, cancels and replaces an existing order,
// or cancels an order, then the Client needs to specify this identifier.
//
// The Client must maintain the same order identifier throughout the life
// of the order.
//
// The Server should persist the ClientOrderID across sessions. A session
// is defined as the period of time from the start of the network connection
// between the Client and Server to the end of that connection.
//
// The Client should only rely upon the ClientOrderID being set when the
// OrderUpdateReason is one of the following: NEW_ORDER_ACCEPTED, NEW_ORDER_REJECTED,
// ORDER_CANCELED, ORDER_CANCEL_REPLACE_COMPLETE, ORDER_CANCEL_REJECTED,
// ORDER_CANCEL_REPLACE_REJECTED.
//
// After a new order has been accepted, the Client will rely upon the given
// ServerOrderID from the server to identify the order and should no longer
// rely upon the given ClientOrderID. However, the Client needs to maintain
// a copy of the ClientOrderID for any subsequent order modifications and
// cancellations because this is a required field for those.
func (m OrderUpdateFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 192, 160, value)
}

// SetExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdateFixedBuilder) SetExchangeOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 224, 192, value)
}

// SetExchangeOrderID The order identifier from the exchange that handles the order. This is
// optional.
func (m OrderUpdateFixedPointerBuilder) SetExchangeOrderID(value string) {
	message.SetStringFixed(m.Ptr, 224, 192, value)
}

// SetOrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdateFixedBuilder) SetOrderStatus(value OrderStatusEnum) {
	message.SetInt32Fixed(m.Unsafe(), 228, 224, int32(value))
}

// SetOrderStatus This is required. Needs to be set to one of the following values by the
// Server to indicate the current status of the order, unless NoneOrder =
// 1:
//
// ORDER_STATUS_UNSPECIFIED (0): The status of the order is unset. Use this
// when NoneOrder = 1, or when the Server does not know the true status of
// an order when the Order Update Report message is sent.
// ORDER_STATUS_ORDER_SENT (1): When a Client sends an order to the Server,
// then the Client internally will set the status to ORDER_STATUS_ORDER_SENT.
// The Server will not set this Status.
// ORDER_STATUS_PENDING_OPEN (2): This means the Server has accepted the
// order but it is not yet considered in a fully working state for any reason.
// order but it is not yet considered in a fully working state for any reason.
// ORDER_STATUS_PENDING_CHILD (3): This status applies to a Limit or Stop
// order attached to a parent order. It will have this status if the parent
// order has not yet filled.
// ORDER_STATUS_OPEN (4): Order is open and working.
// ORDER_STATUS_PENDING_CANCEL_REPLACE (5): Order is pending a Cancel and
// Replace operation. The Server should send a OrderUpdate message with the
// OrderUpdateReason set to GENERAL_OrderUpdate and the OrderStatus set to
// this status to indicate the pending Cancel and Replace operation.
// ORDER_STATUS_PENDING_CANCEL (6): Order is pending a Cancel operation.
// The Server should send a OrderUpdate message with the OrderUpdateReason
// set to GENERAL_OrderUpdate and the OrderStatus set to this status to indicate
// the pending Cancel operation.
// ORDER_STATUS_FILLED (7): Order is filled and no longer working.
// ORDER_STATUS_CANCELED (8): Order is canceled. If the user tries to cancel
// an order that has already been canceled, then continue to return this
// Order Status for it.
// ORDER_STATUS_REJECTED (9): Order has been rejected after the initial order
// submission. It is not working.
// ORDER_STATUS_PARTIALLY_FILLED (10): Order is partially filled and still
// working.
func (m OrderUpdateFixedPointerBuilder) SetOrderStatus(value OrderStatusEnum) {
	message.SetInt32Fixed(m.Ptr, 228, 224, int32(value))
}

// SetOrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdateFixedBuilder) SetOrderUpdateReason(value OrderUpdateReasonEnum) {
	message.SetInt32Fixed(m.Unsafe(), 232, 228, int32(value))
}

// SetOrderUpdateReason This is required. This field needs to be set to one of the following values
// by the Server to indicate the reason the OrderUpdate is being sent.
//
// OpenOrdersRequest_RESPONSE (1): When the OrderUpdate is specifically sent
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// in response to an OpenOrdersRequest request, this is the OrderUpdateReason.
// NEW_ORDER_ACCEPTED (2): This OrderUpdate indicates a new order has been
// accepted.
// GENERAL_OrderUpdate (3): A general order update. For example, when an
// order is in the process of being canceled, the Server may send an OrderUpdate
// message with the OrderStatus set to ORDER_STATUS_PENDING_CANCEL. In this
// particular case the OrderUpdateReason needs to be set to GENERAL_OrderUpdate
// (3).
// ORDER_FILLED (4): Upon a complete fill of the order, this is the OrderUpdateReason.
// This OrderUpdateReason must only be used when an OrderUpdate is sent at
// the moment in time of a fill. A previously filled order that is being
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// restated in response to an OpenOrdersRequest must not use this OrderUpdateReason.
// ORDER_FILLED_PARTIALLY (5): Upon a partial fill of the order, this is
// the OrderUpdateReason. This OrderUpdateReason must only be used when an
// OrderUpdate is sent at the moment in time of a fill. A previously filled
// order that is being restated in response to an OpenOrdersRequest must
// not use this OrderUpdateReason.
// ORDER_CANCELED (6): This OrderUpdateReason indicates the order is now
// successfully canceled.
// ORDER_CANCEL_REPLACE_COMPLETE (7): This OrderUpdateReason indicates the
// order is now successfully canceled and replaced (modified).
// NEW_ORDER_REJECTED (8): After an order has been submitted by the Client,
// it has been rejected for any reason and was never working, the Server
// will send through an OrderUpdate with this OrderUpdateReason. In this
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
// case the Server needs to set the OrderStatus in the OrderUpdate to ORDER_STATUS_REJECTED.
//
// The following fields need to be set for a NEW_ORDER_REJECTED OrderUpdateReason:
// OrderUpdateReason, OrderStatus, MessageNumber, TotalNumMessages, ClientOrderID,
// Symbol, Exchange, TradeAccount.
// ORDER_CANCEL_REJECTED (9): A request by the Client to cancel the order
// with the CancelOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelOrder message from
// the Client is not known, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelOrder message. The ClientOrderID must always be set in this
// case in an OrderUpdate.
//
// In the case where the order has already been canceled for the given ServerOrderID
// in a CancelOrder message, then respond with an OrderUpdate message and
// set the OrderUpdateReason to ORDER_CANCEL_REJECTED and set the OrderStatus
// to ORDER_STATUS_CANCELED.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// that was attempted to be canceled, then set the OrderStatus to ORDER_STATUS_UNSPECIFIED.
// ORDER_CANCEL_REPLACE_REJECTED (10): A request by the Client to cancel
// and replace the order with the CancelReplaceOrder message has been rejected.
// and replace the order with the CancelReplaceOrder message has been rejected.
//
// The current status of the order must be set in the OrderStatus member
// of the OrderUpdate message.
//
// In the case where the given ServerOrderID in a CancelReplaceOrder message
// from the Client is not known, then respond with an OrderUpdate message
// and set the OrderUpdateReason to ORDER_CANCEL_REPLACE_REJECTED and set
// the OrderStatus to ORDER_STATUS_REJECTED.
//
// In the case where the ServerOrderID is not known, then in the OrderUpdate
// message, it should not be set. However, in this case it is necessary to
// set the ClientOrderID in the OrderUpdate message to the given ClientOrderID
// in the CancelReplaceOrder message. The ClientOrderID must always be set
// in this case in an OrderUpdate.
//
// If for some reason the Server is uncertain as to the status of the order
// that was attempted to be canceled and replaced, then set the OrderStatus
// to ORDER_STATUS_UNSPECIFIED.
func (m OrderUpdateFixedPointerBuilder) SetOrderUpdateReason(value OrderUpdateReasonEnum) {
	message.SetInt32Fixed(m.Ptr, 232, 228, int32(value))
}

// SetOrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdateFixedBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 236, 232, int32(value))
}

// SetOrderType The order type. Can be set to one of the following.
//
// ORDER_TYPE_MARKET
// ORDER_TYPE_LIMIT
// ORDER_TYPE_STOP
// ORDER_TYPE_STOP_LIMIT
// ORDER_TYPE_MARKET_IF_TOUCHED
func (m OrderUpdateFixedPointerBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 236, 232, int32(value))
}

// SetBuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdateFixedBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 240, 236, int32(value))
}

// SetBuySell Indicates whether the order is a Buy or Sell order. Can be set to one
// of the following constants: DTC::BUY (1) or DTC::SELL (2).
func (m OrderUpdateFixedPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 240, 236, int32(value))
}

// SetPrice1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) SetPrice1(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 248, 240, value)
}

// SetPrice1 For orders that require a price, this is the order price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetPrice1(value float64) {
	message.SetFloat64Fixed(m.Ptr, 248, 240, value)
}

// SetPrice2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) SetPrice2(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 256, 248, value)
}

// SetPrice2 For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
// For Stop-Limit orders this is the Limit price. Otherwise, this is unset.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetPrice2(value float64) {
	message.SetFloat64Fixed(m.Ptr, 256, 248, value)
}

// SetTimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdateFixedBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Unsafe(), 260, 256, int32(value))
}

// SetTimeInForce The Time in Force of the order. Can be any of the following:
//
// TIF_DAY
// TIF_GOOD_TILL_CANCELED
// TIF_GOOD_TILL_DATE_TIME
// TIF_IMMEDIATE_OR_CANCEL
// TIF_ALL_OR_NONE
// TIF_FILL_OR_KILL
func (m OrderUpdateFixedPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Ptr, 260, 256, int32(value))
}

// SetGoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdateFixedBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 272, 264, int64(value))
}

// SetGoodTillDateTime The expiration Date and Time of the order in the case when TimeInForce
// is TIF_GOOD_TILL_DATE_TIME.
func (m OrderUpdateFixedPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 272, 264, int64(value))
}

// SetOrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) SetOrderQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 280, 272, value)
}

// SetOrderQuantity The quantity of the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetOrderQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 280, 272, value)
}

// SetFilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) SetFilledQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 288, 280, value)
}

// SetFilledQuantity The number of shares or contracts that have filled in the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetFilledQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 288, 280, value)
}

// SetRemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) SetRemainingQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 296, 288, value)
}

// SetRemainingQuantity The number of shares or contracts that still remain to be filled.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetRemainingQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 296, 288, value)
}

// SetAverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) SetAverageFillPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 304, 296, value)
}

// SetAverageFillPrice The average price of all of the fills for the order.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetAverageFillPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 304, 296, value)
}

// SetLastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) SetLastFillPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 312, 304, value)
}

// SetLastFillPrice The price of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetLastFillPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 312, 304, value)
}

// SetLastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedBuilder) SetLastFillDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 320, 312, int64(value))
}

// SetLastFillDateTime The date and Time of the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedPointerBuilder) SetLastFillDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Ptr, 320, 312, int64(value))
}

// SetLastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedBuilder) SetLastFillQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 328, 320, value)
}

// SetLastFillQuantity The number of contracts/shares that has filled for the specific order
// fill that is currently reported through the OrderUpdate message.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m OrderUpdateFixedPointerBuilder) SetLastFillQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 328, 320, value)
}

// SetLastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedBuilder) SetLastFillExecutionID(value string) {
	message.SetStringFixed(m.Unsafe(), 392, 328, value)
}

// SetLastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m OrderUpdateFixedPointerBuilder) SetLastFillExecutionID(value string) {
	message.SetStringFixed(m.Ptr, 392, 328, value)
}

// SetTradeAccount The trade account the order belongs to.
func (m OrderUpdateFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 424, 392, value)
}

// SetTradeAccount The trade account the order belongs to.
func (m OrderUpdateFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 424, 392, value)
}

// SetInfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdateFixedBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Unsafe(), 520, 424, value)
}

// SetInfoText Free-form text with information to communicate about the order. When an
// order is rejected, this should be set by the Server to indicate the reason
// for the rejection.
func (m OrderUpdateFixedPointerBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Ptr, 520, 424, value)
}

// SetNoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdateFixedBuilder) SetNoOrders(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 521, 520, value)
}

// SetNoOrders Set by the Server to 1 to indicate there are no orders when OpenOrdersRequest
// message has been received and is being responded to. Otherwise, leave
// at the default of 0.
func (m OrderUpdateFixedPointerBuilder) SetNoOrders(value uint8) {
	message.SetUint8Fixed(m.Ptr, 521, 520, value)
}

// SetParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedBuilder) SetParentServerOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 553, 521, value)
}

// SetParentServerOrderID This is the ServerOrderID of the parent order when the order that this
// Order Update Report is for, is a child order in a bracket order. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedPointerBuilder) SetParentServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 553, 521, value)
}

// SetOCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedBuilder) SetOCOLinkedOrderServerOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 585, 553, value)
}

// SetOCOLinkedOrderServerOrderID In the case of an OCO order set submitted with SubmitNewOCOOrder, whether
// it has a Parent order or not, this is the ServerOrderID of the other order
// in the OCO pair. These two orders are considered "linked" together. Otherwise,
// this is an empty text string.
func (m OrderUpdateFixedPointerBuilder) SetOCOLinkedOrderServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 585, 553, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdateFixedBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 592, 588, int32(value))
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m OrderUpdateFixedPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Ptr, 592, 588, int32(value))
}

// SetPreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdateFixedBuilder) SetPreviousClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 624, 592, value)
}

// SetPreviousClientOrderID The PreviousClientOrderID is the previous ClientOrderID provided by the
// Client for the order, if the Client changed it during order cancel and
// replace request or an order cancel request.
//
// A Server only is obligated to provide this field immediately after the
// ClientOrderID has been changed. Subsequent OrderUpdate messages do not
// need to set this field.
func (m OrderUpdateFixedPointerBuilder) SetPreviousClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 624, 592, value)
}

// SetFreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdateFixedBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Unsafe(), 672, 624, value)
}

// SetFreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m OrderUpdateFixedPointerBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Ptr, 672, 624, value)
}

// SetOrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdateFixedBuilder) SetOrderReceivedDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 680, 672, int64(value))
}

// SetOrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdateFixedPointerBuilder) SetOrderReceivedDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Ptr, 680, 672, int64(value))
}

// SetLatestTransactionDateTime
func (m OrderUpdateFixedBuilder) SetLatestTransactionDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 688, 680, float64(value))
}

// SetLatestTransactionDateTime
func (m OrderUpdateFixedPointerBuilder) SetLatestTransactionDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 688, 680, float64(value))
}

// Copy
func (m OrderUpdate) Copy(to OrderUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// Copy
func (m OrderUpdateBuilder) Copy(to OrderUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyPointer
func (m OrderUpdate) CopyPointer(to OrderUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyPointer
func (m OrderUpdateBuilder) CopyPointer(to OrderUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyToPointer
func (m OrderUpdate) CopyToPointer(to OrderUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyToPointer
func (m OrderUpdateBuilder) CopyToPointer(to OrderUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyTo
func (m OrderUpdate) CopyTo(to OrderUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyTo
func (m OrderUpdateBuilder) CopyTo(to OrderUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// Copy
func (m OrderUpdateFixed) Copy(to OrderUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// Copy
func (m OrderUpdateFixedBuilder) Copy(to OrderUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyPointer
func (m OrderUpdateFixed) CopyPointer(to OrderUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyPointer
func (m OrderUpdateFixedBuilder) CopyPointer(to OrderUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyToPointer
func (m OrderUpdateFixed) CopyToPointer(to OrderUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyToPointer
func (m OrderUpdateFixedBuilder) CopyToPointer(to OrderUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyTo
func (m OrderUpdateFixed) CopyTo(to OrderUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyTo
func (m OrderUpdateFixedBuilder) CopyTo(to OrderUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// Copy
func (m OrderUpdatePointer) Copy(to OrderUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// Copy
func (m OrderUpdatePointerBuilder) Copy(to OrderUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyPointer
func (m OrderUpdatePointer) CopyPointer(to OrderUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyPointer
func (m OrderUpdatePointerBuilder) CopyPointer(to OrderUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyTo
func (m OrderUpdatePointer) CopyTo(to OrderUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyTo
func (m OrderUpdatePointerBuilder) CopyTo(to OrderUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyToPointer
func (m OrderUpdatePointer) CopyToPointer(to OrderUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyToPointer
func (m OrderUpdatePointerBuilder) CopyToPointer(to OrderUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// Copy
func (m OrderUpdateFixedPointer) Copy(to OrderUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// Copy
func (m OrderUpdateFixedPointerBuilder) Copy(to OrderUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyPointer
func (m OrderUpdateFixedPointer) CopyPointer(to OrderUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyPointer
func (m OrderUpdateFixedPointerBuilder) CopyPointer(to OrderUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyTo
func (m OrderUpdateFixedPointer) CopyTo(to OrderUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyTo
func (m OrderUpdateFixedPointerBuilder) CopyTo(to OrderUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyToPointer
func (m OrderUpdateFixedPointer) CopyToPointer(to OrderUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}

// CopyToPointer
func (m OrderUpdateFixedPointerBuilder) CopyToPointer(to OrderUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumMessages(m.TotalNumMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetPreviousServerOrderID(m.PreviousServerOrderID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetExchangeOrderID(m.ExchangeOrderID())
	to.SetOrderStatus(m.OrderStatus())
	to.SetOrderUpdateReason(m.OrderUpdateReason())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetOrderQuantity(m.OrderQuantity())
	to.SetFilledQuantity(m.FilledQuantity())
	to.SetRemainingQuantity(m.RemainingQuantity())
	to.SetAverageFillPrice(m.AverageFillPrice())
	to.SetLastFillPrice(m.LastFillPrice())
	to.SetLastFillDateTime(m.LastFillDateTime())
	to.SetLastFillQuantity(m.LastFillQuantity())
	to.SetLastFillExecutionID(m.LastFillExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetInfoText(m.InfoText())
	to.SetNoOrders(m.NoOrders())
	to.SetParentServerOrderID(m.ParentServerOrderID())
	to.SetOCOLinkedOrderServerOrderID(m.OCOLinkedOrderServerOrderID())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPreviousClientOrderID(m.PreviousClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOrderReceivedDateTime(m.OrderReceivedDateTime())
	to.SetLatestTransactionDateTime(m.LatestTransactionDateTime())
}
