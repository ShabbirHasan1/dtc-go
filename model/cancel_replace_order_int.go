package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _CancelReplaceOrderIntDefault = []byte{80, 0, 205, 0, 80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const CancelReplaceOrderIntSize = 80

// CancelReplaceOrderInt The CancelReplaceOrder_INT is a message from the Client to Server to cancel
// and replace (modify) an order.
//
// It is identical to CancelReplaceOrder except the prices are as integers.
// It is identical to CancelReplaceOrder except the prices are as integers.
//
// When the cancel and replace operation is completed, an OrderUpdate message
// is sent by the Server with the OrderUpdateReason field set to ORDER_CANCEL_REPLACE_COMPLETE.
// If the cancel and replace operation cannot be completed, an OrderUpdate
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
type CancelReplaceOrderInt struct {
	message.VLS
}

// CancelReplaceOrderIntBuilder The CancelReplaceOrder_INT is a message from the Client to Server to cancel
// and replace (modify) an order.
//
// It is identical to CancelReplaceOrder except the prices are as integers.
// It is identical to CancelReplaceOrder except the prices are as integers.
//
// When the cancel and replace operation is completed, an OrderUpdate message
// is sent by the Server with the OrderUpdateReason field set to ORDER_CANCEL_REPLACE_COMPLETE.
// If the cancel and replace operation cannot be completed, an OrderUpdate
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
type CancelReplaceOrderIntBuilder struct {
	message.VLSBuilder
}

func NewCancelReplaceOrderIntFrom(b []byte) CancelReplaceOrderInt {
	return CancelReplaceOrderInt{VLS: message.NewVLSFrom(b)}
}

func WrapCancelReplaceOrderInt(b []byte) CancelReplaceOrderInt {
	return CancelReplaceOrderInt{VLS: message.WrapVLS(b)}
}

func NewCancelReplaceOrderInt() CancelReplaceOrderIntBuilder {
	a := CancelReplaceOrderIntBuilder{message.NewVLSBuilder(80)}
	a.Unsafe().SetBytes(0, _CancelReplaceOrderIntDefault)
	return a
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
func (m CancelReplaceOrderIntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CancelReplaceOrderIntDefault)
}

// ToBuilder
func (m CancelReplaceOrderInt) ToBuilder() CancelReplaceOrderIntBuilder {
	return CancelReplaceOrderIntBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m CancelReplaceOrderIntBuilder) Finish() CancelReplaceOrderInt {
	return CancelReplaceOrderInt{m.VLSBuilder.Finish()}
}

// ServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrderInt) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// ServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrderIntBuilder) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
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
func (m CancelReplaceOrderInt) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
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
func (m CancelReplaceOrderIntBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
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
func (m CancelReplaceOrderInt) Price1() int64 {
	return message.Int64VLS(m.Unsafe(), 24, 16)
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
func (m CancelReplaceOrderIntBuilder) Price1() int64 {
	return message.Int64VLS(m.Unsafe(), 24, 16)
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
func (m CancelReplaceOrderInt) Price2() int64 {
	return message.Int64VLS(m.Unsafe(), 32, 24)
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
func (m CancelReplaceOrderIntBuilder) Price2() int64 {
	return message.Int64VLS(m.Unsafe(), 32, 24)
}

// Divisor This is the FloatToIntPriceMultiplier field received from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// The Server needs to divide the Price1 or Price2 fields by this Divisor
// to arrive at the original floating-point value.
func (m CancelReplaceOrderInt) Divisor() float32 {
	return message.Float32VLS(m.Unsafe(), 36, 32)
}

// Divisor This is the FloatToIntPriceMultiplier field received from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// The Server needs to divide the Price1 or Price2 fields by this Divisor
// to arrive at the original floating-point value.
func (m CancelReplaceOrderIntBuilder) Divisor() float32 {
	return message.Float32VLS(m.Unsafe(), 36, 32)
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
func (m CancelReplaceOrderInt) Quantity() int64 {
	return message.Int64VLS(m.Unsafe(), 48, 40)
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
func (m CancelReplaceOrderIntBuilder) Quantity() int64 {
	return message.Int64VLS(m.Unsafe(), 48, 40)
}

// Price1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderInt) Price1IsSet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 49, 48)
}

// Price1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntBuilder) Price1IsSet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 49, 48)
}

// Price2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderInt) Price2IsSet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 50, 49)
}

// Price2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntBuilder) Price2IsSet() uint8 {
	return message.Uint8VLS(m.Unsafe(), 50, 49)
}

// Unused
func (m CancelReplaceOrderInt) Unused() int32 {
	return message.Int32VLS(m.Unsafe(), 56, 52)
}

// Unused
func (m CancelReplaceOrderIntBuilder) Unused() int32 {
	return message.Int32VLS(m.Unsafe(), 56, 52)
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
func (m CancelReplaceOrderInt) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Unsafe(), 60, 56))
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
func (m CancelReplaceOrderIntBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Unsafe(), 60, 56))
}

// GoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderInt) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 72, 64))
}

// GoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderIntBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 72, 64))
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
func (m CancelReplaceOrderInt) UpdatePrice1OffsetToParent() uint8 {
	return message.Uint8VLS(m.Unsafe(), 73, 72)
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
func (m CancelReplaceOrderIntBuilder) UpdatePrice1OffsetToParent() uint8 {
	return message.Uint8VLS(m.Unsafe(), 73, 72)
}

// SetServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m *CancelReplaceOrderIntBuilder) SetServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
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
func (m *CancelReplaceOrderIntBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
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
func (m CancelReplaceOrderIntBuilder) SetPrice1(value int64) {
	message.SetInt64VLS(m.Unsafe(), 24, 16, value)
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
func (m CancelReplaceOrderIntBuilder) SetPrice2(value int64) {
	message.SetInt64VLS(m.Unsafe(), 32, 24, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier field received from the SecurityDefinitionResponse
// message for the Symbol of the order.
//
// The Server needs to divide the Price1 or Price2 fields by this Divisor
// to arrive at the original floating-point value.
func (m CancelReplaceOrderIntBuilder) SetDivisor(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 36, 32, value)
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
func (m CancelReplaceOrderIntBuilder) SetQuantity(value int64) {
	message.SetInt64VLS(m.Unsafe(), 48, 40, value)
}

// SetPrice1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntBuilder) SetPrice1IsSet(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 49, 48, value)
}

// SetPrice2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderIntBuilder) SetPrice2IsSet(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 50, 49, value)
}

// SetUnused
func (m CancelReplaceOrderIntBuilder) SetUnused(value int32) {
	message.SetInt32VLS(m.Unsafe(), 56, 52, value)
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
func (m CancelReplaceOrderIntBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Unsafe(), 60, 56, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderIntBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 72, 64, int64(value))
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
func (m CancelReplaceOrderIntBuilder) SetUpdatePrice1OffsetToParent(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 73, 72, value)
}

// Copy
func (m CancelReplaceOrderInt) Copy(to CancelReplaceOrderIntBuilder) {
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
func (m CancelReplaceOrderIntBuilder) Copy(to CancelReplaceOrderIntBuilder) {
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
func (m CancelReplaceOrderInt) CopyPointer(to CancelReplaceOrderIntPointerBuilder) {
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
func (m CancelReplaceOrderIntBuilder) CopyPointer(to CancelReplaceOrderIntPointerBuilder) {
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
func (m CancelReplaceOrderInt) CopyToPointer(to CancelReplaceOrderIntFixedPointerBuilder) {
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
func (m CancelReplaceOrderIntBuilder) CopyToPointer(to CancelReplaceOrderIntFixedPointerBuilder) {
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
func (m CancelReplaceOrderInt) CopyTo(to CancelReplaceOrderIntFixedBuilder) {
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
func (m CancelReplaceOrderIntBuilder) CopyTo(to CancelReplaceOrderIntFixedBuilder) {
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
