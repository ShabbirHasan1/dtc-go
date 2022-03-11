package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = CancelOrderFixedSize  (100)
//     Type          = CANCEL_ORDER  (203)
//     ServerOrderID = ""
//     ClientOrderID = ""
//     TradeAccount  = ""
// }
var _CancelOrderFixedDefault = []byte{100, 0, 203, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const CancelOrderFixedSize = 100

// CancelOrderFixed This is a message from the Client to the Server requesting a previously
// sent order to be canceled.
type CancelOrderFixed struct {
	message.Fixed
}

// CancelOrderFixedBuilder This is a message from the Client to the Server requesting a previously
// sent order to be canceled.
type CancelOrderFixedBuilder struct {
	message.Fixed
}

func NewCancelOrderFixedFrom(b []byte) CancelOrderFixed {
	return CancelOrderFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapCancelOrderFixed(b []byte) CancelOrderFixed {
	return CancelOrderFixed{Fixed: message.WrapFixed(b)}
}

func NewCancelOrderFixed() CancelOrderFixedBuilder {
	a := CancelOrderFixedBuilder{message.NewFixed(100)}
	a.Unsafe().SetBytes(0, _CancelOrderFixedDefault)
	return a
}

// Clear
// {
//     Size          = CancelOrderFixedSize  (100)
//     Type          = CANCEL_ORDER  (203)
//     ServerOrderID = ""
//     ClientOrderID = ""
//     TradeAccount  = ""
// }
func (m CancelOrderFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CancelOrderFixedDefault)
}

// ToBuilder
func (m CancelOrderFixed) ToBuilder() CancelOrderFixedBuilder {
	return CancelOrderFixedBuilder{m.Fixed}
}

// Finish
func (m CancelOrderFixedBuilder) Finish() CancelOrderFixed {
	return CancelOrderFixed{m.Fixed}
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
func (m CancelOrderFixed) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 36, 4)
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
func (m CancelOrderFixedBuilder) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 36, 4)
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
func (m CancelOrderFixed) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 68, 36)
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
func (m CancelOrderFixedBuilder) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 68, 36)
}

// TradeAccount
func (m CancelOrderFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 100, 68)
}

// TradeAccount
func (m CancelOrderFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 100, 68)
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
func (m CancelOrderFixedBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 36, 4, value)
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
func (m CancelOrderFixedBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 68, 36, value)
}

// SetTradeAccount
func (m CancelOrderFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 100, 68, value)
}

// Copy
func (m CancelOrderFixed) Copy(to CancelOrderFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m CancelOrderFixedBuilder) Copy(to CancelOrderFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CancelOrderFixed) CopyPointer(to CancelOrderFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CancelOrderFixedBuilder) CopyPointer(to CancelOrderFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CancelOrderFixed) CopyToPointer(to CancelOrderPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CancelOrderFixedBuilder) CopyToPointer(to CancelOrderPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CancelOrderFixed) CopyTo(to CancelOrderBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CancelOrderFixedBuilder) CopyTo(to CancelOrderBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}
