package model

import (
	"github.com/moontrade/dtc-go/message"
)

// CancelReplaceOrderIntFixedPointer The CancelReplaceOrder_INT is a message from the Client to Server to cancel
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
type CancelReplaceOrderIntFixedPointer struct {
	message.FixedPointer
}

// CancelReplaceOrderIntFixedPointerBuilder The CancelReplaceOrder_INT is a message from the Client to Server to cancel
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
type CancelReplaceOrderIntFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocCancelReplaceOrderIntFixed() CancelReplaceOrderIntFixedPointerBuilder {
	a := CancelReplaceOrderIntFixedPointerBuilder{message.AllocFixed(136)}
	a.Ptr.SetBytes(0, _CancelReplaceOrderIntFixedDefault)
	return a
}

func AllocCancelReplaceOrderIntFixedFrom(b []byte) CancelReplaceOrderIntFixedPointer {
	return CancelReplaceOrderIntFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                       = CancelReplaceOrderIntFixedSize  (136)
//     Type                       = CANCEL_REPLACE_ORDER_INT  (205)
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
func (m CancelReplaceOrderIntFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CancelReplaceOrderIntFixedDefault)
}

// ToBuilder
func (m CancelReplaceOrderIntFixedPointer) ToBuilder() CancelReplaceOrderIntFixedPointerBuilder {
	return CancelReplaceOrderIntFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *CancelReplaceOrderIntFixedPointerBuilder) Finish() CancelReplaceOrderIntFixedPointer {
	return CancelReplaceOrderIntFixedPointer{m.FixedPointer}
}

// ServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrderIntFixedPointer) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 36, 4)
}

// ServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrderIntFixedPointerBuilder) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 36, 4)
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
func (m CancelReplaceOrderIntFixedPointer) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 68, 36)
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
func (m CancelReplaceOrderIntFixedPointerBuilder) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 68, 36)
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
func (m CancelReplaceOrderIntFixedPointer) Price1() int64 {
	return message.Int64Fixed(m.Ptr, 80, 72)
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
func (m CancelReplaceOrderIntFixedPointerBuilder) Price1() int64 {
	return message.Int64Fixed(m.Ptr, 80, 72)
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
func (m CancelReplaceOrderIntFixedPointer) Price2() int64 {
	return message.Int64Fixed(m.Ptr, 88, 80)
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
func (m CancelReplaceOrderIntFixedPointerBuilder) Price2() int64 {
	return message.Int64Fixed(m.Ptr, 88, 80)
}

// Divisor This is the FloatToIntPriceMultiplier field received from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// The Server needs to divide the Price1 or Price2 fields by this Divisor
// to arrive at the original floating-point value.
func (m CancelReplaceOrderIntFixedPointer) Divisor() float32 {
	return message.Float32Fixed(m.Ptr, 92, 88)
}

// Divisor This is the FloatToIntPriceMultiplier field received from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// The Server needs to divide the Price1 or Price2 fields by this Divisor
// to arrive at the original floating-point value.
func (m CancelReplaceOrderIntFixedPointerBuilder) Divisor() float32 {
	return message.Float32Fixed(m.Ptr, 92, 88)
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
func (m CancelReplaceOrderIntFixedPointer) Quantity() int64 {
	return message.Int64Fixed(m.Ptr, 104, 96)
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
func (m CancelReplaceOrderIntFixedPointerBuilder) Quantity() int64 {
	return message.Int64Fixed(m.Ptr, 104, 96)
}

// Price1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntFixedPointer) Price1IsSet() uint8 {
	return message.Uint8Fixed(m.Ptr, 105, 104)
}

// Price1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntFixedPointerBuilder) Price1IsSet() uint8 {
	return message.Uint8Fixed(m.Ptr, 105, 104)
}

// Price2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntFixedPointer) Price2IsSet() uint8 {
	return message.Uint8Fixed(m.Ptr, 106, 105)
}

// Price2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntFixedPointerBuilder) Price2IsSet() uint8 {
	return message.Uint8Fixed(m.Ptr, 106, 105)
}

// Unused
func (m CancelReplaceOrderIntFixedPointer) Unused() int32 {
	return message.Int32Fixed(m.Ptr, 112, 108)
}

// Unused
func (m CancelReplaceOrderIntFixedPointerBuilder) Unused() int32 {
	return message.Int32Fixed(m.Ptr, 112, 108)
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
func (m CancelReplaceOrderIntFixedPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 116, 112))
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
func (m CancelReplaceOrderIntFixedPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 116, 112))
}

// GoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderIntFixedPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 128, 120))
}

// GoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderIntFixedPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 128, 120))
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
func (m CancelReplaceOrderIntFixedPointer) UpdatePrice1OffsetToParent() uint8 {
	return message.Uint8Fixed(m.Ptr, 129, 128)
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
func (m CancelReplaceOrderIntFixedPointerBuilder) UpdatePrice1OffsetToParent() uint8 {
	return message.Uint8Fixed(m.Ptr, 129, 128)
}

// SetServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrderIntFixedPointerBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 36, 4, value)
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
func (m CancelReplaceOrderIntFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 68, 36, value)
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
func (m CancelReplaceOrderIntFixedPointerBuilder) SetPrice1(value int64) {
	message.SetInt64Fixed(m.Ptr, 80, 72, value)
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
func (m CancelReplaceOrderIntFixedPointerBuilder) SetPrice2(value int64) {
	message.SetInt64Fixed(m.Ptr, 88, 80, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier field received from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// The Server needs to divide the Price1 or Price2 fields by this Divisor
// to arrive at the original floating-point value.
func (m CancelReplaceOrderIntFixedPointerBuilder) SetDivisor(value float32) {
	message.SetFloat32Fixed(m.Ptr, 92, 88, value)
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
func (m CancelReplaceOrderIntFixedPointerBuilder) SetQuantity(value int64) {
	message.SetInt64Fixed(m.Ptr, 104, 96, value)
}

// SetPrice1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntFixedPointerBuilder) SetPrice1IsSet(value uint8) {
	message.SetUint8Fixed(m.Ptr, 105, 104, value)
}

// SetPrice2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntFixedPointerBuilder) SetPrice2IsSet(value uint8) {
	message.SetUint8Fixed(m.Ptr, 106, 105, value)
}

// SetUnused
func (m CancelReplaceOrderIntFixedPointerBuilder) SetUnused(value int32) {
	message.SetInt32Fixed(m.Ptr, 112, 108, value)
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
func (m CancelReplaceOrderIntFixedPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Ptr, 116, 112, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderIntFixedPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 128, 120, int64(value))
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
func (m CancelReplaceOrderIntFixedPointerBuilder) SetUpdatePrice1OffsetToParent(value uint8) {
	message.SetUint8Fixed(m.Ptr, 129, 128, value)
}

// Copy
func (m CancelReplaceOrderIntFixedPointer) Copy(to CancelReplaceOrderIntFixedBuilder) {
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
func (m CancelReplaceOrderIntFixedPointerBuilder) Copy(to CancelReplaceOrderIntFixedBuilder) {
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
func (m CancelReplaceOrderIntFixedPointer) CopyPointer(to CancelReplaceOrderIntFixedPointerBuilder) {
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
func (m CancelReplaceOrderIntFixedPointerBuilder) CopyPointer(to CancelReplaceOrderIntFixedPointerBuilder) {
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
func (m CancelReplaceOrderIntFixedPointer) CopyTo(to CancelReplaceOrderIntBuilder) {
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
func (m CancelReplaceOrderIntFixedPointerBuilder) CopyTo(to CancelReplaceOrderIntBuilder) {
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
func (m CancelReplaceOrderIntFixedPointer) CopyToPointer(to CancelReplaceOrderIntPointerBuilder) {
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
func (m CancelReplaceOrderIntFixedPointerBuilder) CopyToPointer(to CancelReplaceOrderIntPointerBuilder) {
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
