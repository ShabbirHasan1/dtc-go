package model

import (
	"github.com/moontrade/dtc-go/message"
)

// CancelReplaceOrderIntPointer The CancelReplaceOrder_INT is a message from the Client to Server to cancel
// and replace (modify) an order.
//
// It is identical to CancelReplaceOrder except the prices are as integers.
// It is identical to CancelReplaceOrder except the prices are as integers.
//
// When the cancel and replace operation is completed, an OrderUpdate message
// is sent by the Server with the OrderUpdateReasonfield set to ORDER_CANCEL_REPLACE_COMPLETE.
// If the cancel and replace operation cannot be completed, an OrderUpdate
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
type CancelReplaceOrderIntPointer struct {
	message.VLSPointer
}

// CancelReplaceOrderIntPointerBuilder The CancelReplaceOrder_INT is a message from the Client to Server to cancel
// and replace (modify) an order.
//
// It is identical to CancelReplaceOrder except the prices are as integers.
// It is identical to CancelReplaceOrder except the prices are as integers.
//
// When the cancel and replace operation is completed, an OrderUpdate message
// is sent by the Server with the OrderUpdateReasonfield set to ORDER_CANCEL_REPLACE_COMPLETE.
// If the cancel and replace operation cannot be completed, an OrderUpdate
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
type CancelReplaceOrderIntPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocCancelReplaceOrderInt() CancelReplaceOrderIntPointerBuilder {
	a := CancelReplaceOrderIntPointerBuilder{message.AllocVLSBuilder(80)}
	a.Ptr.SetBytes(0, _CancelReplaceOrderIntDefault)
	return a
}

func AllocCancelReplaceOrderIntFrom(b []byte) CancelReplaceOrderIntPointer {
	return CancelReplaceOrderIntPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                       = CancelReplaceOrderIntSize  (80)
//     Type                       = CANCEL_REPLACE_ORDER_INT  (205)
//     BaseSize                   = CancelReplaceOrderIntSize  (80)
//     ServerOrderID              = ""
//     ClientOrderID              = ""
//     Price1                     = 0
//     Price2                     = 0
//     Divisor                    = 1.000000
//     Quantity                   = 0
//     Price1IsSet                = 1
//     Price2IsSet                = 1
//     Unused                     = 0
//     TimeInForce                = 0
//     GoodTillDateTime           = 0
//     UpdatePrice1OffsetToParent = 0
// }
func (m CancelReplaceOrderIntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CancelReplaceOrderIntDefault)
}

// ToBuilder
func (m CancelReplaceOrderIntPointer) ToBuilder() CancelReplaceOrderIntPointerBuilder {
	return CancelReplaceOrderIntPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *CancelReplaceOrderIntPointerBuilder) Finish() CancelReplaceOrderIntPointer {
	return CancelReplaceOrderIntPointer{m.VLSPointerBuilder.Finish()}
}

// ServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrderIntPointer) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// ServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrderIntPointerBuilder) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// ClientOrderID This is the Client's own order identifier for the order.
//
// This must be the same throughout the life of the order. If the Server
// sees that this order identifier has changed in relation to the ServerOrderID,
// then it should reject this message with a OrderUpdate message with the
// OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED
//
// In the case where the order modification cannot be performed because the
// ServerOrderID does not exist, the Server will send a OrderUpdate message
// with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED and ClientOrderID
// set to the given ClientOrderID in this message. ServerOrderID will be
// unset because an invalid server order identifier was given.
func (m CancelReplaceOrderIntPointer) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// ClientOrderID This is the Client's own order identifier for the order.
//
// This must be the same throughout the life of the order. If the Server
// sees that this order identifier has changed in relation to the ServerOrderID,
// then it should reject this message with a OrderUpdate message with the
// OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED
//
// In the case where the order modification cannot be performed because the
// ServerOrderID does not exist, the Server will send a OrderUpdate message
// with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED and ClientOrderID
// set to the given ClientOrderID in this message. ServerOrderID will be
// unset because an invalid server order identifier was given.
func (m CancelReplaceOrderIntPointerBuilder) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Price1 For orders that require a price, this is the new order price.
//
// This is an integer value. This is calculated by multiplying the actual
// order price by the FloatToIntPriceMultiplier field from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m CancelReplaceOrderIntPointer) Price1() int64 {
	return message.Int64VLS(m.Ptr, 24, 16)
}

// Price1 For orders that require a price, this is the new order price.
//
// This is an integer value. This is calculated by multiplying the actual
// order price by the FloatToIntPriceMultiplier field from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m CancelReplaceOrderIntPointerBuilder) Price1() int64 {
	return message.Int64VLS(m.Ptr, 24, 16)
}

// Price2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This is an integer value. This is calculated by multiplying the actual
// order price by the FloatToIntPriceMultiplier field from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m CancelReplaceOrderIntPointer) Price2() int64 {
	return message.Int64VLS(m.Ptr, 32, 24)
}

// Price2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This is an integer value. This is calculated by multiplying the actual
// order price by the FloatToIntPriceMultiplier field from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m CancelReplaceOrderIntPointerBuilder) Price2() int64 {
	return message.Int64VLS(m.Ptr, 32, 24)
}

// Divisor This is the FloatToIntPriceMultiplier field received from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// The Server needs to divide the Price1 or Price2 fields by this Divisor
// to arrive at the original floating-point value.
func (m CancelReplaceOrderIntPointer) Divisor() float32 {
	return message.Float32VLS(m.Ptr, 36, 32)
}

// Divisor This is the FloatToIntPriceMultiplier field received from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// The Server needs to divide the Price1 or Price2 fields by this Divisor
// to arrive at the original floating-point value.
func (m CancelReplaceOrderIntPointerBuilder) Divisor() float32 {
	return message.Float32VLS(m.Ptr, 36, 32)
}

// Quantity This is the new order quantity. It this is 0, then this means the order
// quantity must not be changed by the Server.
//
// If the order has partially filled, then this is going to be the order
// quantity which also includes the amount which has partially filled.
//
// For example, if the original quantity was 10 and there has been a partial
// fill of 3, the Client wants a fill of 2 more making a total of 5, then
// the Client will set this to 5.
func (m CancelReplaceOrderIntPointer) Quantity() int64 {
	return message.Int64VLS(m.Ptr, 48, 40)
}

// Quantity This is the new order quantity. It this is 0, then this means the order
// quantity must not be changed by the Server.
//
// If the order has partially filled, then this is going to be the order
// quantity which also includes the amount which has partially filled.
//
// For example, if the original quantity was 10 and there has been a partial
// fill of 3, the Client wants a fill of 2 more making a total of 5, then
// the Client will set this to 5.
func (m CancelReplaceOrderIntPointerBuilder) Quantity() int64 {
	return message.Int64VLS(m.Ptr, 48, 40)
}

// Price1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntPointer) Price1IsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 49, 48)
}

// Price1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntPointerBuilder) Price1IsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 49, 48)
}

// Price2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntPointer) Price2IsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 50, 49)
}

// Price2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntPointerBuilder) Price2IsSet() uint8 {
	return message.Uint8VLS(m.Ptr, 50, 49)
}

// Unused
func (m CancelReplaceOrderIntPointer) Unused() int32 {
	return message.Int32VLS(m.Ptr, 56, 52)
}

// Unused
func (m CancelReplaceOrderIntPointerBuilder) Unused() int32 {
	return message.Int32VLS(m.Ptr, 56, 52)
}

// TimeInForce The Time in Force for the order. For a list of Time in Force values, refer
// to TimeInForceEnum.
//
// The default value is TIF_UNSET.
//
// When this field is set to a value other than TIF_UNSET, it indicates that
// the TimeInForce is being changed.
//
// If the server does not support changing the Time in Force of the order,
// it needs to reject this CancelReplaceOrder_INT message and send an OrderUpdate
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
//
// The server is under no obligation to support changing the Time in Force.
// The server is under no obligation to support changing the Time in Force.
func (m CancelReplaceOrderIntPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Ptr, 60, 56))
}

// TimeInForce The Time in Force for the order. For a list of Time in Force values, refer
// to TimeInForceEnum.
//
// The default value is TIF_UNSET.
//
// When this field is set to a value other than TIF_UNSET, it indicates that
// the TimeInForce is being changed.
//
// If the server does not support changing the Time in Force of the order,
// it needs to reject this CancelReplaceOrder_INT message and send an OrderUpdate
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
//
// The server is under no obligation to support changing the Time in Force.
// The server is under no obligation to support changing the Time in Force.
func (m CancelReplaceOrderIntPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Ptr, 60, 56))
}

// GoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderIntPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 72, 64))
}

// GoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderIntPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 72, 64))
}

// UpdatePrice1OffsetToParent This is an optional field. If modifying a child order which is part of
// a Server managed bracket order, then when this variable is set to 1 it
// provides an indication to the Server to update the internal server managed
// price offset to the parent order that this child order has to the parent.
// price offset to the parent order that this child order has to the parent.
//
// This will ensure the Server will maintain the proper offset of the child
// order to the fill price of the parent order when the parent order fills.
// order to the fill price of the parent order when the parent order fills.
func (m CancelReplaceOrderIntPointer) UpdatePrice1OffsetToParent() uint8 {
	return message.Uint8VLS(m.Ptr, 73, 72)
}

// UpdatePrice1OffsetToParent This is an optional field. If modifying a child order which is part of
// a Server managed bracket order, then when this variable is set to 1 it
// provides an indication to the Server to update the internal server managed
// price offset to the parent order that this child order has to the parent.
// price offset to the parent order that this child order has to the parent.
//
// This will ensure the Server will maintain the proper offset of the child
// order to the fill price of the parent order when the parent order fills.
// order to the fill price of the parent order when the parent order fills.
func (m CancelReplaceOrderIntPointerBuilder) UpdatePrice1OffsetToParent() uint8 {
	return message.Uint8VLS(m.Ptr, 73, 72)
}

// SetServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m *CancelReplaceOrderIntPointerBuilder) SetServerOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetClientOrderID This is the Client's own order identifier for the order.
//
// This must be the same throughout the life of the order. If the Server
// sees that this order identifier has changed in relation to the ServerOrderID,
// then it should reject this message with a OrderUpdate message with the
// OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED
//
// In the case where the order modification cannot be performed because the
// ServerOrderID does not exist, the Server will send a OrderUpdate message
// with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED and ClientOrderID
// set to the given ClientOrderID in this message. ServerOrderID will be
// unset because an invalid server order identifier was given.
func (m *CancelReplaceOrderIntPointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetPrice1 For orders that require a price, this is the new order price.
//
// This is an integer value. This is calculated by multiplying the actual
// order price by the FloatToIntPriceMultiplier field from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m CancelReplaceOrderIntPointerBuilder) SetPrice1(value int64) {
	message.SetInt64VLS(m.Ptr, 24, 16, value)
}

// SetPrice2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This is an integer value. This is calculated by multiplying the actual
// order price by the FloatToIntPriceMultiplier field from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m CancelReplaceOrderIntPointerBuilder) SetPrice2(value int64) {
	message.SetInt64VLS(m.Ptr, 32, 24, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier field received from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// The Server needs to divide the Price1 or Price2 fields by this Divisor
// to arrive at the original floating-point value.
func (m CancelReplaceOrderIntPointerBuilder) SetDivisor(value float32) {
	message.SetFloat32VLS(m.Ptr, 36, 32, value)
}

// SetQuantity This is the new order quantity. It this is 0, then this means the order
// quantity must not be changed by the Server.
//
// If the order has partially filled, then this is going to be the order
// quantity which also includes the amount which has partially filled.
//
// For example, if the original quantity was 10 and there has been a partial
// fill of 3, the Client wants a fill of 2 more making a total of 5, then
// the Client will set this to 5.
func (m CancelReplaceOrderIntPointerBuilder) SetQuantity(value int64) {
	message.SetInt64VLS(m.Ptr, 48, 40, value)
}

// SetPrice1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntPointerBuilder) SetPrice1IsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 49, 48, value)
}

// SetPrice2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntPointerBuilder) SetPrice2IsSet(value uint8) {
	message.SetUint8VLS(m.Ptr, 50, 49, value)
}

// SetUnused
func (m CancelReplaceOrderIntPointerBuilder) SetUnused(value int32) {
	message.SetInt32VLS(m.Ptr, 56, 52, value)
}

// SetTimeInForce The Time in Force for the order. For a list of Time in Force values, refer
// to TimeInForceEnum.
//
// The default value is TIF_UNSET.
//
// When this field is set to a value other than TIF_UNSET, it indicates that
// the TimeInForce is being changed.
//
// If the server does not support changing the Time in Force of the order,
// it needs to reject this CancelReplaceOrder_INT message and send an OrderUpdate
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
//
// The server is under no obligation to support changing the Time in Force.
// The server is under no obligation to support changing the Time in Force.
func (m CancelReplaceOrderIntPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Ptr, 60, 56, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderIntPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 72, 64, int64(value))
}

// SetUpdatePrice1OffsetToParent This is an optional field. If modifying a child order which is part of
// a Server managed bracket order, then when this variable is set to 1 it
// provides an indication to the Server to update the internal server managed
// price offset to the parent order that this child order has to the parent.
// price offset to the parent order that this child order has to the parent.
//
// This will ensure the Server will maintain the proper offset of the child
// order to the fill price of the parent order when the parent order fills.
// order to the fill price of the parent order when the parent order fills.
func (m CancelReplaceOrderIntPointerBuilder) SetUpdatePrice1OffsetToParent(value uint8) {
	message.SetUint8VLS(m.Ptr, 73, 72, value)
}

// Copy
func (m CancelReplaceOrderIntPointer) Copy(to CancelReplaceOrderIntBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
}

// Copy
func (m CancelReplaceOrderIntPointerBuilder) Copy(to CancelReplaceOrderIntBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
}

// CopyPointer
func (m CancelReplaceOrderIntPointer) CopyPointer(to CancelReplaceOrderIntPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
}

// CopyPointer
func (m CancelReplaceOrderIntPointerBuilder) CopyPointer(to CancelReplaceOrderIntPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
}

// CopyTo
func (m CancelReplaceOrderIntPointer) CopyTo(to CancelReplaceOrderIntFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
}

// CopyTo
func (m CancelReplaceOrderIntPointerBuilder) CopyTo(to CancelReplaceOrderIntFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
}

// CopyToPointer
func (m CancelReplaceOrderIntPointer) CopyToPointer(to CancelReplaceOrderIntFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
}

// CopyToPointer
func (m CancelReplaceOrderIntPointerBuilder) CopyToPointer(to CancelReplaceOrderIntFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
}
