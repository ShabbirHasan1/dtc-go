package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = CancelOrderSize  (18)
//     Type          = CANCEL_ORDER  (203)
//     BaseSize      = CancelOrderSize  (18)
//     ServerOrderID = ""
//     ClientOrderID = ""
//     TradeAccount  = ""
// }
var _CancelOrderDefault = []byte{18, 0, 203, 0, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const CancelOrderSize = 18

// CancelOrder This is a message from the Client to the Server requesting a previously
// sent order to be canceled.
type CancelOrder struct {
	message.VLS
}

// CancelOrderBuilder This is a message from the Client to the Server requesting a previously
// sent order to be canceled.
type CancelOrderBuilder struct {
	message.VLSBuilder
}

func NewCancelOrderFrom(b []byte) CancelOrder {
	return CancelOrder{VLS: message.NewVLSFrom(b)}
}

func WrapCancelOrder(b []byte) CancelOrder {
	return CancelOrder{VLS: message.WrapVLS(b)}
}

func NewCancelOrder() CancelOrderBuilder {
	a := CancelOrderBuilder{message.NewVLSBuilder(18)}
	a.Unsafe().SetBytes(0, _CancelOrderDefault)
	return a
}

// Clear
// {
//     Size          = CancelOrderSize  (18)
//     Type          = CANCEL_ORDER  (203)
//     BaseSize      = CancelOrderSize  (18)
//     ServerOrderID = ""
//     ClientOrderID = ""
//     TradeAccount  = ""
// }
func (m CancelOrderBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CancelOrderDefault)
}

// ToBuilder
func (m CancelOrder) ToBuilder() CancelOrderBuilder {
	return CancelOrderBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m CancelOrderBuilder) Finish() CancelOrder {
	return CancelOrder{m.VLSBuilder.Finish()}
}

// ServerOrderID This is the order identifier for the order to cancel. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order. The only case in which a ServerOrderID
// would change is in the case of a successful order Cancel and Replace operation.
// would change is in the case of a successful order Cancel and Replace operation.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled. Although the given ClientOrderID
// from the Client must not change.
func (m CancelOrder) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// ServerOrderID This is the order identifier for the order to cancel. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order. The only case in which a ServerOrderID
// would change is in the case of a successful order Cancel and Replace operation.
// would change is in the case of a successful order Cancel and Replace operation.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled. Although the given ClientOrderID
// from the Client must not change.
func (m CancelOrderBuilder) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// ClientOrderID This is the Client's own order identifier for the order.
//
// This must be the same throughout the life of the order. If the Server
// sees that this order identifier has changed in relation to the ServerOrderID
// , then it should reject this message with a OrderUpdate message with the
// OrderUpdateReason set to ORDER_CANCEL_REJECTED.
//
// In the case where the order cancellation cannot be performed because the
// ServerOrderID does not exist, the Server will send a OrderUpdate message
// with the OrderUpdateReason set to ORDER_CANCEL_REJECTED and ClientOrderID
// set to the given ClientOrderID in this message. ServerOrderID will be
// unset because an invalid server order identifier was given.
func (m CancelOrder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// ClientOrderID This is the Client's own order identifier for the order.
//
// This must be the same throughout the life of the order. If the Server
// sees that this order identifier has changed in relation to the ServerOrderID
// , then it should reject this message with a OrderUpdate message with the
// OrderUpdateReason set to ORDER_CANCEL_REJECTED.
//
// In the case where the order cancellation cannot be performed because the
// ServerOrderID does not exist, the Server will send a OrderUpdate message
// with the OrderUpdateReason set to ORDER_CANCEL_REJECTED and ClientOrderID
// set to the given ClientOrderID in this message. ServerOrderID will be
// unset because an invalid server order identifier was given.
func (m CancelOrderBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m CancelOrder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount
func (m CancelOrderBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// SetServerOrderID This is the order identifier for the order to cancel. The Client needs
// to set this to the ServerOrderID field received back in the most recent
// OrderUpdate message for the order. The only case in which a ServerOrderID
// would change is in the case of a successful order Cancel and Replace operation.
// would change is in the case of a successful order Cancel and Replace operation.
//
// The Server will rely upon this ServerOrderID and only this order identifier
// to identify the order to be canceled. Although the given ClientOrderID
// from the Client must not change.
func (m *CancelOrderBuilder) SetServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetClientOrderID This is the Client's own order identifier for the order.
//
// This must be the same throughout the life of the order. If the Server
// sees that this order identifier has changed in relation to the ServerOrderID
// , then it should reject this message with a OrderUpdate message with the
// OrderUpdateReason set to ORDER_CANCEL_REJECTED.
//
// In the case where the order cancellation cannot be performed because the
// ServerOrderID does not exist, the Server will send a OrderUpdate message
// with the OrderUpdateReason set to ORDER_CANCEL_REJECTED and ClientOrderID
// set to the given ClientOrderID in this message. ServerOrderID will be
// unset because an invalid server order identifier was given.
func (m *CancelOrderBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *CancelOrderBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// Copy
func (m CancelOrder) Copy(to CancelOrderBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m CancelOrderBuilder) Copy(to CancelOrderBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CancelOrder) CopyPointer(to CancelOrderPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CancelOrderBuilder) CopyPointer(to CancelOrderPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CancelOrder) CopyToPointer(to CancelOrderFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CancelOrderBuilder) CopyToPointer(to CancelOrderFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CancelOrder) CopyTo(to CancelOrderFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CancelOrderBuilder) CopyTo(to CancelOrderFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}
