// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const CancelReplaceOrderSize = 80

const CancelReplaceOrderFixedSize = 192

//     Size                        uint16           = CancelReplaceOrderSize  (80)
//     Type                        uint16           = CANCEL_REPLACE_ORDER  (204)
//     BaseSize                    uint16           = CancelReplaceOrderSize  (80)
//     ServerOrderID               string           = ""
//     ClientOrderID               string           = ""
//     Price1                      float64          = 0.000000
//     Price2                      float64          = 0.000000
//     Quantity                    float64          = 0.000000
//     Price1IsSet                 bool             = true
//     Price2IsSet                 bool             = true
//     Unused                      int32            = 0
//     TimeInForce                 TimeInForceEnum  = TIF_UNSET  (0)
//     GoodTillDateTime            DateTime         = 0
//     UpdatePrice1OffsetToParent  uint8            = 0
//     TradeAccount                string           = ""
//     Price1AsString              string           = ""
//     Price2AsString              string           = ""
var _CancelReplaceOrderDefault = []byte{80, 0, 204, 0, 80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size                        uint16           = CancelReplaceOrderFixedSize  (192)
//     Type                        uint16           = CANCEL_REPLACE_ORDER  (204)
//     ServerOrderID               string[32]       = ""
//     ClientOrderID               string[32]       = ""
//     Price1                      float64          = 0.000000
//     Price2                      float64          = 0.000000
//     Quantity                    float64          = 0.000000
//     Price1IsSet                 bool             = true
//     Price2IsSet                 bool             = true
//     Unused                      int32            = 0
//     TimeInForce                 TimeInForceEnum  = TIF_UNSET  (0)
//     GoodTillDateTime            DateTime         = 0
//     UpdatePrice1OffsetToParent  uint8            = 0
//     TradeAccount                string[32]       = ""
//     Price1AsString              string[16]       = ""
//     Price2AsString              string[16]       = ""
var _CancelReplaceOrderFixedDefault = []byte{192, 0, 204, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// CancelReplaceOrder This message is sent by the Client to the Server to cancel and replace
// an existing order. This is also known as an order modification.
//
// When the cancel and replace operation is completed, an OrderUpdate message
// is sent by the Server with the OrderUpdateReasonfield set to ORDER_CANCEL_REPLACE_COMPLETE.
// If the cancel and replace operation cannot be completed, an OrderUpdate
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
type CancelReplaceOrder struct {
	p message.VLS
}

// CancelReplaceOrderFixed This message is sent by the Client to the Server to cancel and replace
// an existing order. This is also known as an order modification.
//
// When the cancel and replace operation is completed, an OrderUpdate message
// is sent by the Server with the OrderUpdateReasonfield set to ORDER_CANCEL_REPLACE_COMPLETE.
// If the cancel and replace operation cannot be completed, an OrderUpdate
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
type CancelReplaceOrderFixed struct {
	p message.Fixed
}

func NewCancelReplaceOrderFrom(b []byte) CancelReplaceOrder {
	return CancelReplaceOrder{p: message.NewVLS(b)}
}

func WrapCancelReplaceOrder(b []byte) CancelReplaceOrder {
	return CancelReplaceOrder{p: message.WrapVLS(b)}
}

func NewCancelReplaceOrder() *CancelReplaceOrder {
	return &CancelReplaceOrder{p: message.NewVLS(_CancelReplaceOrderDefault)}
}

func ParseCancelReplaceOrder(b []byte) (CancelReplaceOrder, error) {
	if len(b) < 6 {
		return CancelReplaceOrder{}, message.ErrShortBuffer
	}
	m := WrapCancelReplaceOrder(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return CancelReplaceOrder{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return CancelReplaceOrder{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 80 {
		newSize := len(b) + (80 - baseSize)
		if newSize > message.MaxSize {
			return CancelReplaceOrder{}, message.ErrOverflow
		}
		clone := CancelReplaceOrder{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _CancelReplaceOrderDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(80 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(6)
			if offset > 0 {
				clone.p.SetUint16LE(6, offset+shift)
			}
			offset = clone.p.Uint16LE(10)
			if offset > 0 {
				clone.p.SetUint16LE(10, offset+shift)
			}
			offset = clone.p.Uint16LE(66)
			if offset > 0 {
				clone.p.SetUint16LE(66, offset+shift)
			}
			offset = clone.p.Uint16LE(70)
			if offset > 0 {
				clone.p.SetUint16LE(70, offset+shift)
			}
			offset = clone.p.Uint16LE(74)
			if offset > 0 {
				clone.p.SetUint16LE(74, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewCancelReplaceOrderFixedFrom(b []byte) CancelReplaceOrderFixed {
	return CancelReplaceOrderFixed{p: message.NewFixed(b)}
}

func WrapCancelReplaceOrderFixed(b []byte) CancelReplaceOrderFixed {
	return CancelReplaceOrderFixed{p: message.WrapFixed(b)}
}

func NewCancelReplaceOrderFixed() *CancelReplaceOrderFixed {
	return &CancelReplaceOrderFixed{p: message.NewFixed(_CancelReplaceOrderFixedDefault)}
}

func ParseCancelReplaceOrderFixed(b []byte) (CancelReplaceOrderFixed, error) {
	if len(b) < 4 {
		return CancelReplaceOrderFixed{}, message.ErrShortBuffer
	}
	m := WrapCancelReplaceOrderFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return CancelReplaceOrderFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return CancelReplaceOrderFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 192 {
		clone := *NewCancelReplaceOrderFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _CancelReplaceOrderFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m CancelReplaceOrder) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m CancelReplaceOrder) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m CancelReplaceOrder) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// ServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrder) ServerOrderID() string {
	return m.p.StringVLS(6)
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
func (m CancelReplaceOrder) ClientOrderID() string {
	return m.p.StringVLS(10)
}

// Price1 For orders that require a price, this is the new order price.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m CancelReplaceOrder) Price1() float64 {
	return m.p.Float64LE(16)
}

// Price2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m CancelReplaceOrder) Price2() float64 {
	return m.p.Float64LE(24)
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
func (m CancelReplaceOrder) Quantity() float64 {
	return m.p.Float64LE(32)
}

// Price1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrder) Price1IsSet() bool {
	return m.p.Bool(40)
}

// Price2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrder) Price2IsSet() bool {
	return m.p.Bool(41)
}

// Unused
func (m CancelReplaceOrder) Unused() int32 {
	return m.p.Int32LE(44)
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
func (m CancelReplaceOrder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(m.p.Int32LE(48))
}

// GoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrder) GoodTillDateTime() DateTime {
	return DateTime(m.p.Int64LE(56))
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
func (m CancelReplaceOrder) UpdatePrice1OffsetToParent() uint8 {
	return m.p.Uint8(64)
}

// TradeAccount
func (m CancelReplaceOrder) TradeAccount() string {
	return m.p.StringVLS(66)
}

// Price1AsString
func (m CancelReplaceOrder) Price1AsString() string {
	return m.p.StringVLS(70)
}

// Price2AsString
func (m CancelReplaceOrder) Price2AsString() string {
	return m.p.StringVLS(74)
}

// Size The standard message size field. Automatically set by constructor.
func (m CancelReplaceOrderFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m CancelReplaceOrderFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// ServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m CancelReplaceOrderFixed) ServerOrderID() string {
	return m.p.StringFixed(4, 32)
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
func (m CancelReplaceOrderFixed) ClientOrderID() string {
	return m.p.StringFixed(36, 32)
}

// Price1 For orders that require a price, this is the new order price.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m CancelReplaceOrderFixed) Price1() float64 {
	return m.p.Float64LE(72)
}

// Price2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m CancelReplaceOrderFixed) Price2() float64 {
	return m.p.Float64LE(80)
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
func (m CancelReplaceOrderFixed) Quantity() float64 {
	return m.p.Float64LE(88)
}

// Price1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderFixed) Price1IsSet() bool {
	return m.p.Bool(96)
}

// Price2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m CancelReplaceOrderFixed) Price2IsSet() bool {
	return m.p.Bool(97)
}

// Unused
func (m CancelReplaceOrderFixed) Unused() int32 {
	return m.p.Int32LE(100)
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
func (m CancelReplaceOrderFixed) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(m.p.Int32LE(104))
}

// GoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m CancelReplaceOrderFixed) GoodTillDateTime() DateTime {
	return DateTime(m.p.Int64LE(112))
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
func (m CancelReplaceOrderFixed) UpdatePrice1OffsetToParent() uint8 {
	return m.p.Uint8(120)
}

// TradeAccount
func (m CancelReplaceOrderFixed) TradeAccount() string {
	return m.p.StringFixed(121, 32)
}

// Price1AsString
func (m CancelReplaceOrderFixed) Price1AsString() string {
	return m.p.StringFixed(153, 16)
}

// Price2AsString
func (m CancelReplaceOrderFixed) Price2AsString() string {
	return m.p.StringFixed(169, 16)
}

// SetServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m *CancelReplaceOrder) SetServerOrderID(value string) *CancelReplaceOrder {
	m.p.SetStringVLS(6, value)
	return m
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
func (m *CancelReplaceOrder) SetClientOrderID(value string) *CancelReplaceOrder {
	m.p.SetStringVLS(10, value)
	return m
}

// SetPrice1 For orders that require a price, this is the new order price.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m *CancelReplaceOrder) SetPrice1(value float64) *CancelReplaceOrder {
	m.p.SetFloat64LE(16, value)
	return m
}

// SetPrice2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m *CancelReplaceOrder) SetPrice2(value float64) *CancelReplaceOrder {
	m.p.SetFloat64LE(24, value)
	return m
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
func (m *CancelReplaceOrder) SetQuantity(value float64) *CancelReplaceOrder {
	m.p.SetFloat64LE(32, value)
	return m
}

// SetPrice1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m *CancelReplaceOrder) SetPrice1IsSet(value bool) *CancelReplaceOrder {
	m.p.SetBool(40, value)
	return m
}

// SetPrice2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m *CancelReplaceOrder) SetPrice2IsSet(value bool) *CancelReplaceOrder {
	m.p.SetBool(41, value)
	return m
}

// SetUnused
func (m *CancelReplaceOrder) SetUnused(value int32) *CancelReplaceOrder {
	m.p.SetInt32LE(44, value)
	return m
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
func (m *CancelReplaceOrder) SetTimeInForce(value TimeInForceEnum) *CancelReplaceOrder {
	m.p.SetInt32LE(48, int32(value))
	return m
}

// SetGoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m *CancelReplaceOrder) SetGoodTillDateTime(value DateTime) *CancelReplaceOrder {
	m.p.SetInt64LE(56, int64(value))
	return m
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
func (m *CancelReplaceOrder) SetUpdatePrice1OffsetToParent(value uint8) *CancelReplaceOrder {
	m.p.SetUint8(64, value)
	return m
}

// SetTradeAccount
func (m *CancelReplaceOrder) SetTradeAccount(value string) *CancelReplaceOrder {
	m.p.SetStringVLS(66, value)
	return m
}

// SetPrice1AsString
func (m *CancelReplaceOrder) SetPrice1AsString(value string) *CancelReplaceOrder {
	m.p.SetStringVLS(70, value)
	return m
}

// SetPrice2AsString
func (m *CancelReplaceOrder) SetPrice2AsString(value string) *CancelReplaceOrder {
	m.p.SetStringVLS(74, value)
	return m
}

// SetServerOrderID This is the order identifier for the order to modify. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled and replaced. Although the given
// ClientOrderID by the Client must not change.
func (m *CancelReplaceOrderFixed) SetServerOrderID(value string) *CancelReplaceOrderFixed {
	m.p.SetStringFixed(4, 32, value)
	return m
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
func (m *CancelReplaceOrderFixed) SetClientOrderID(value string) *CancelReplaceOrderFixed {
	m.p.SetStringFixed(36, 32, value)
	return m
}

// SetPrice1 For orders that require a price, this is the new order price.
//
// This value can be left unset indicating to the Server that Price1 must
// not change and only the Quantity. In this case it is necessary to set
// Price1IsSet to a 0 value.
func (m *CancelReplaceOrderFixed) SetPrice1(value float64) *CancelReplaceOrderFixed {
	m.p.SetFloat64LE(72, value)
	return m
}

// SetPrice2 For Stop-Limit orders this is the new Limit price. For other order types
// it is not used.
//
// This value can be left unset indicating to the Server that Price2 must
// not change and only the Quantity. In this case it is necessary to set
// Price2IsSet to a 0 value.
func (m *CancelReplaceOrderFixed) SetPrice2(value float64) *CancelReplaceOrderFixed {
	m.p.SetFloat64LE(80, value)
	return m
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
func (m *CancelReplaceOrderFixed) SetQuantity(value float64) *CancelReplaceOrderFixed {
	m.p.SetFloat64LE(88, value)
	return m
}

// SetPrice1IsSet When this field is set to a nonzero value, it indicates that Price1 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m *CancelReplaceOrderFixed) SetPrice1IsSet(value bool) *CancelReplaceOrderFixed {
	m.p.SetBool(96, value)
	return m
}

// SetPrice2IsSet When this field is set to a nonzero value, it indicates that Price2 is
// set and the server should use the value, if it applies to the order type.
// set and the server should use the value, if it applies to the order type.
//
// The default value is 1.
func (m *CancelReplaceOrderFixed) SetPrice2IsSet(value bool) *CancelReplaceOrderFixed {
	m.p.SetBool(97, value)
	return m
}

// SetUnused
func (m *CancelReplaceOrderFixed) SetUnused(value int32) *CancelReplaceOrderFixed {
	m.p.SetInt32LE(100, value)
	return m
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
func (m *CancelReplaceOrderFixed) SetTimeInForce(value TimeInForceEnum) *CancelReplaceOrderFixed {
	m.p.SetInt32LE(104, int32(value))
	return m
}

// SetGoodTillDateTime In the case of when the TimeInForce field is TIF_GOOD_TILL_DATE_TIME,
// this specifies the expiration Date-Time of the order.
func (m *CancelReplaceOrderFixed) SetGoodTillDateTime(value DateTime) *CancelReplaceOrderFixed {
	m.p.SetInt64LE(112, int64(value))
	return m
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
func (m *CancelReplaceOrderFixed) SetUpdatePrice1OffsetToParent(value uint8) *CancelReplaceOrderFixed {
	m.p.SetUint8(120, value)
	return m
}

// SetTradeAccount
func (m *CancelReplaceOrderFixed) SetTradeAccount(value string) *CancelReplaceOrderFixed {
	m.p.SetStringFixed(121, 32, value)
	return m
}

// SetPrice1AsString
func (m *CancelReplaceOrderFixed) SetPrice1AsString(value string) *CancelReplaceOrderFixed {
	m.p.SetStringFixed(153, 16, value)
	return m
}

// SetPrice2AsString
func (m *CancelReplaceOrderFixed) SetPrice2AsString(value string) *CancelReplaceOrderFixed {
	m.p.SetStringFixed(169, 16, value)
	return m
}

func (m *CancelReplaceOrder) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *CancelReplaceOrder) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *CancelReplaceOrder) Clone() *CancelReplaceOrder {
	return &CancelReplaceOrder{message.WrapVLSPointer(m.p.Clone(uintptr(m.Size())), int(m.Size()))}
}

func (m *CancelReplaceOrderFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *CancelReplaceOrderFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *CancelReplaceOrderFixed) Clone() *CancelReplaceOrderFixed {
	return &CancelReplaceOrderFixed{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m CancelReplaceOrder) Copy(to CancelReplaceOrder) {
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
func (m CancelReplaceOrder) CopyTo(to CancelReplaceOrderFixed) {
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
func (m CancelReplaceOrderFixed) Copy(to CancelReplaceOrderFixed) {
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
func (m CancelReplaceOrderFixed) CopyTo(to CancelReplaceOrder) {
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrder) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *CancelReplaceOrder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 204)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrder) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *CancelReplaceOrder) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 204 {
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
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "Price1":
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "Price1IsSet":
			m.SetPrice1IsSet(r.Bool())
		case "Price2IsSet":
			m.SetPrice2IsSet(r.Bool())
		case "Unused":
			m.SetUnused(r.Int32())
		case "TimeInForce":
			m.SetTimeInForce(TimeInForceEnum(r.Int32()))
		case "GoodTillDateTime":
			m.SetGoodTillDateTime(DateTime(r.Int64()))
		case "UpdatePrice1OffsetToParent":
			m.SetUpdatePrice1OffsetToParent(r.Uint8())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "Price1AsString":
			m.SetPrice1AsString(r.String())
		case "Price2AsString":
			m.SetPrice2AsString(r.String())
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
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderFixed) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *CancelReplaceOrderFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 204)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderFixed) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *CancelReplaceOrderFixed) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 204 {
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
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "Price1":
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "Price1IsSet":
			m.SetPrice1IsSet(r.Bool())
		case "Price2IsSet":
			m.SetPrice2IsSet(r.Bool())
		case "Unused":
			m.SetUnused(r.Int32())
		case "TimeInForce":
			m.SetTimeInForce(TimeInForceEnum(r.Int32()))
		case "GoodTillDateTime":
			m.SetGoodTillDateTime(DateTime(r.Int64()))
		case "UpdatePrice1OffsetToParent":
			m.SetUpdatePrice1OffsetToParent(r.Uint8())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "Price1AsString":
			m.SetPrice1AsString(r.String())
		case "Price2AsString":
			m.SetPrice2AsString(r.String())
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
