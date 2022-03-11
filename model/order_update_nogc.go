package model

import (
	"github.com/moontrade/dtc-go/message"
)

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

func AllocOrderUpdate() OrderUpdatePointerBuilder {
	a := OrderUpdatePointerBuilder{message.AllocVLSBuilder(208)}
	a.Ptr.SetBytes(0, _OrderUpdateDefault)
	return a
}

func AllocOrderUpdateFrom(b []byte) OrderUpdatePointer {
	return OrderUpdatePointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     OrderStatus                 = 0
//     OrderUpdateReason           = 0
//     OrderType                   = 0
//     BuySell                     = 0
//     Price1                      = math.MaxFloat64
//     Price2                      = math.MaxFloat64
//     TimeInForce                 = 0
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
//     OpenOrClose                 = 0
//     PreviousClientOrderID       = ""
//     FreeFormText                = ""
//     OrderReceivedDateTime       = 0
//     LatestTransactionDateTime   = 0
// }
func (m OrderUpdatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _OrderUpdateDefault)
}

// ToBuilder
func (m OrderUpdatePointer) ToBuilder() OrderUpdatePointerBuilder {
	return OrderUpdatePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *OrderUpdatePointerBuilder) Finish() OrderUpdatePointer {
	return OrderUpdatePointer{m.VLSPointerBuilder.Finish()}
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
func (m OrderUpdatePointer) Symbol() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Symbol The symbol for the order.
func (m OrderUpdatePointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 24, 20)
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
func (m OrderUpdatePointer) LatestTransactionDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Ptr, 208, 200))
}

// LatestTransactionDateTime
func (m OrderUpdatePointerBuilder) LatestTransactionDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64VLS(m.Ptr, 208, 200))
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
func (m OrderUpdatePointerBuilder) SetTotalNumMessages(value int32) {
	message.SetInt32VLS(m.Ptr, 16, 12, value)
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
func (m *OrderUpdatePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
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
func (m *OrderUpdatePointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 40, 36, value)
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
func (m OrderUpdatePointerBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32VLS(m.Ptr, 56, 52, int32(value))
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
func (m OrderUpdatePointerBuilder) SetPrice1(value float64) {
	message.SetFloat64VLS(m.Ptr, 72, 64, value)
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
func (m OrderUpdatePointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Ptr, 84, 80, int32(value))
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
func (m OrderUpdatePointerBuilder) SetOrderQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 104, 96, value)
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
func (m OrderUpdatePointerBuilder) SetRemainingQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 120, 112, value)
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
func (m OrderUpdatePointerBuilder) SetLastFillPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 136, 128, value)
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
func (m OrderUpdatePointerBuilder) SetLastFillQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 152, 144, value)
}

// SetLastFillExecutionID The unique identifier for the most recent fill.
//
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
// Only necessary to set if OrderUpdateReason is ORDER_FILLED or ORDER_FILLED_PARTIALLY.
func (m *OrderUpdatePointerBuilder) SetLastFillExecutionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 156, 152, value)
}

// SetTradeAccount The trade account the order belongs to.
func (m *OrderUpdatePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 160, 156, value)
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
func (m OrderUpdatePointerBuilder) SetNoOrders(value uint8) {
	message.SetUint8VLS(m.Ptr, 165, 164, value)
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
func (m *OrderUpdatePointerBuilder) SetOCOLinkedOrderServerOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 174, 170, value)
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
func (m *OrderUpdatePointerBuilder) SetPreviousClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 184, 180, value)
}

// SetFreeFormText This is the optional free-form text that was originally set with a new
// order using the new order messages.
func (m *OrderUpdatePointerBuilder) SetFreeFormText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 188, 184, value)
}

// SetOrderReceivedDateTime This is the Date-Time when the original order was received by the Server.
// Order modifications normally will not cause this Date-Time to be updated.
// Order modifications normally will not cause this Date-Time to be updated.
func (m OrderUpdatePointerBuilder) SetOrderReceivedDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Ptr, 200, 192, int64(value))
}

// SetLatestTransactionDateTime
func (m OrderUpdatePointerBuilder) SetLatestTransactionDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64VLS(m.Ptr, 208, 200, float64(value))
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
