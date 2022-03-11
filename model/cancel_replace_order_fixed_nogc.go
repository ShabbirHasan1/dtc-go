package model

import (
	"github.com/moontrade/dtc-go/message"
)

// CancelReplaceOrderFixedPointer This message is sent by the Client to the Server to cancel and replace
// an existing order. This is also known as an order modification.
//
// When the cancel and replace operation is completed, an OrderUpdate message
// is sent by the Server with the OrderUpdateReasonfield set to ORDER_CANCEL_REPLACE_COMPLETE.
// If the cancel and replace operation cannot be completed, an OrderUpdate
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
type CancelReplaceOrderFixedPointer struct {
	message.FixedPointer
}

// CancelReplaceOrderFixedPointerBuilder This message is sent by the Client to the Server to cancel and replace
// an existing order. This is also known as an order modification.
//
// When the cancel and replace operation is completed, an OrderUpdate message
// is sent by the Server with the OrderUpdateReasonfield set to ORDER_CANCEL_REPLACE_COMPLETE.
// If the cancel and replace operation cannot be completed, an OrderUpdate
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
type CancelReplaceOrderFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocCancelReplaceOrderFixed() CancelReplaceOrderFixedPointerBuilder {
	a := CancelReplaceOrderFixedPointerBuilder{message.AllocFixed(192)}
	a.Ptr.SetBytes(0, _CancelReplaceOrderFixedDefault)
	return a
}

func AllocCancelReplaceOrderFixedFrom(b []byte) CancelReplaceOrderFixedPointer {
	return CancelReplaceOrderFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                       = CancelReplaceOrderFixedSize  (192)
//     Type                       = CANCEL_REPLACE_ORDER  (204)
//     ServerOrderID              = ""
//     ClientOrderID              = ""
//     Price1                     = 0
//     Price2                     = 0
//     Quantity                   = 0
//     Price1IsSet                = 1
//     Price2IsSet                = 1
//     Unused                     = 0
//     TimeInForce                = 0
//     GoodTillDateTime           = 0
//     UpdatePrice1OffsetToParent = 0
//     TradeAccount               = ""
//     Price1AsString             = ""
//     Price2AsString             = ""
// }
func (m CancelReplaceOrderFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CancelReplaceOrderFixedDefault)
}

// ToBuilder
func (m CancelReplaceOrderFixedPointer) ToBuilder() CancelReplaceOrderFixedPointerBuilder {
	return CancelReplaceOrderFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *CancelReplaceOrderFixedPointerBuilder) Finish() CancelReplaceOrderFixedPointer {
	return CancelReplaceOrderFixedPointer{m.FixedPointer}
}

// ServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrderFixedPointer) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 36, 4)
}

// ServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrderFixedPointerBuilder) ServerOrderID() string {
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
func (m CancelReplaceOrderFixedPointer) ClientOrderID() string {
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
func (m CancelReplaceOrderFixedPointerBuilder) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 68, 36)
}

// Price1 For orders that require a price, this is the new order price.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m CancelReplaceOrderFixedPointer) Price1() float64 {
	return message.Float64Fixed(m.Ptr, 80, 72)
}

// Price1 For orders that require a price, this is the new order price.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m CancelReplaceOrderFixedPointerBuilder) Price1() float64 {
	return message.Float64Fixed(m.Ptr, 80, 72)
}

// Price2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m CancelReplaceOrderFixedPointer) Price2() float64 {
	return message.Float64Fixed(m.Ptr, 88, 80)
}

// Price2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m CancelReplaceOrderFixedPointerBuilder) Price2() float64 {
	return message.Float64Fixed(m.Ptr, 88, 80)
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
func (m CancelReplaceOrderFixedPointer) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 96, 88)
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
func (m CancelReplaceOrderFixedPointerBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 96, 88)
}

// Price1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderFixedPointer) Price1IsSet() uint8 {
	return message.Uint8Fixed(m.Ptr, 97, 96)
}

// Price1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderFixedPointerBuilder) Price1IsSet() uint8 {
	return message.Uint8Fixed(m.Ptr, 97, 96)
}

// Price2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderFixedPointer) Price2IsSet() uint8 {
	return message.Uint8Fixed(m.Ptr, 98, 97)
}

// Price2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderFixedPointerBuilder) Price2IsSet() uint8 {
	return message.Uint8Fixed(m.Ptr, 98, 97)
}

// Unused
func (m CancelReplaceOrderFixedPointer) Unused() int32 {
	return message.Int32Fixed(m.Ptr, 104, 100)
}

// Unused
func (m CancelReplaceOrderFixedPointerBuilder) Unused() int32 {
	return message.Int32Fixed(m.Ptr, 104, 100)
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
// it needs to reject this CancelReplaceOrder message and send an OrderUpdate
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
//
// The server is under no obligation to support changing the Time in Force.
// The server is under no obligation to support changing the Time in Force.
func (m CancelReplaceOrderFixedPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 108, 104))
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
// it needs to reject this CancelReplaceOrder message and send an OrderUpdate
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
//
// The server is under no obligation to support changing the Time in Force.
// The server is under no obligation to support changing the Time in Force.
func (m CancelReplaceOrderFixedPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 108, 104))
}

// GoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderFixedPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 120, 112))
}

// GoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderFixedPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 120, 112))
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
func (m CancelReplaceOrderFixedPointer) UpdatePrice1OffsetToParent() uint8 {
	return message.Uint8Fixed(m.Ptr, 121, 120)
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
func (m CancelReplaceOrderFixedPointerBuilder) UpdatePrice1OffsetToParent() uint8 {
	return message.Uint8Fixed(m.Ptr, 121, 120)
}

// TradeAccount
func (m CancelReplaceOrderFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 153, 121)
}

// TradeAccount
func (m CancelReplaceOrderFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 153, 121)
}

// Price1AsString
func (m CancelReplaceOrderFixedPointer) Price1AsString() string {
	return message.StringFixed(m.Ptr, 169, 153)
}

// Price1AsString
func (m CancelReplaceOrderFixedPointerBuilder) Price1AsString() string {
	return message.StringFixed(m.Ptr, 169, 153)
}

// Price2AsString
func (m CancelReplaceOrderFixedPointer) Price2AsString() string {
	return message.StringFixed(m.Ptr, 185, 169)
}

// Price2AsString
func (m CancelReplaceOrderFixedPointerBuilder) Price2AsString() string {
	return message.StringFixed(m.Ptr, 185, 169)
}

// SetServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrderFixedPointerBuilder) SetServerOrderID(value string) {
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
func (m CancelReplaceOrderFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 68, 36, value)
}

// SetPrice1 For orders that require a price, this is the new order price.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m CancelReplaceOrderFixedPointerBuilder) SetPrice1(value float64) {
	message.SetFloat64Fixed(m.Ptr, 80, 72, value)
}

// SetPrice2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m CancelReplaceOrderFixedPointerBuilder) SetPrice2(value float64) {
	message.SetFloat64Fixed(m.Ptr, 88, 80, value)
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
func (m CancelReplaceOrderFixedPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 96, 88, value)
}

// SetPrice1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderFixedPointerBuilder) SetPrice1IsSet(value uint8) {
	message.SetUint8Fixed(m.Ptr, 97, 96, value)
}

// SetPrice2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderFixedPointerBuilder) SetPrice2IsSet(value uint8) {
	message.SetUint8Fixed(m.Ptr, 98, 97, value)
}

// SetUnused
func (m CancelReplaceOrderFixedPointerBuilder) SetUnused(value int32) {
	message.SetInt32Fixed(m.Ptr, 104, 100, value)
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
// it needs to reject this CancelReplaceOrder message and send an OrderUpdate
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
//
// The server is under no obligation to support changing the Time in Force.
// The server is under no obligation to support changing the Time in Force.
func (m CancelReplaceOrderFixedPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Ptr, 108, 104, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderFixedPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 120, 112, int64(value))
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
func (m CancelReplaceOrderFixedPointerBuilder) SetUpdatePrice1OffsetToParent(value uint8) {
	message.SetUint8Fixed(m.Ptr, 121, 120, value)
}

// SetTradeAccount
func (m CancelReplaceOrderFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 153, 121, value)
}

// SetPrice1AsString
func (m CancelReplaceOrderFixedPointerBuilder) SetPrice1AsString(value string) {
	message.SetStringFixed(m.Ptr, 169, 153, value)
}

// SetPrice2AsString
func (m CancelReplaceOrderFixedPointerBuilder) SetPrice2AsString(value string) {
	message.SetStringFixed(m.Ptr, 185, 169, value)
}

// Copy
func (m CancelReplaceOrderFixedPointer) Copy(to CancelReplaceOrderFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// Copy
func (m CancelReplaceOrderFixedPointerBuilder) Copy(to CancelReplaceOrderFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyPointer
func (m CancelReplaceOrderFixedPointer) CopyPointer(to CancelReplaceOrderFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyPointer
func (m CancelReplaceOrderFixedPointerBuilder) CopyPointer(to CancelReplaceOrderFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyTo
func (m CancelReplaceOrderFixedPointer) CopyTo(to CancelReplaceOrderBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyTo
func (m CancelReplaceOrderFixedPointerBuilder) CopyTo(to CancelReplaceOrderBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyToPointer
func (m CancelReplaceOrderFixedPointer) CopyToPointer(to CancelReplaceOrderPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyToPointer
func (m CancelReplaceOrderFixedPointerBuilder) CopyToPointer(to CancelReplaceOrderPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetPrice1IsSet(m.Price1IsSet())
	to.SetPrice2IsSet(m.Price2IsSet())
	to.SetUnused(m.Unused())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetUpdatePrice1OffsetToParent(m.UpdatePrice1OffsetToParent())
	to.SetTradeAccount(m.TradeAccount())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}
