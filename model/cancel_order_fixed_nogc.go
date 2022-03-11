package model

import (
	"github.com/moontrade/dtc-go/message"
)

// CancelOrderFixedPointer This is a message from the Client to the Server requesting a previously
// sent order to be canceled.
type CancelOrderFixedPointer struct {
	message.FixedPointer
}

// CancelOrderFixedPointerBuilder This is a message from the Client to the Server requesting a previously
// sent order to be canceled.
type CancelOrderFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocCancelOrderFixed() CancelOrderFixedPointerBuilder {
	a := CancelOrderFixedPointerBuilder{message.AllocFixed(100)}
	a.Ptr.SetBytes(0, _CancelOrderFixedDefault)
	return a
}

func AllocCancelOrderFixedFrom(b []byte) CancelOrderFixedPointer {
	return CancelOrderFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = CancelOrderFixedSize  (100)
//     Type          = CANCEL_ORDER  (203)
//     ServerOrderID = ""
//     ClientOrderID = ""
//     TradeAccount  = ""
// }
func (m CancelOrderFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CancelOrderFixedDefault)
}

// ToBuilder
func (m CancelOrderFixedPointer) ToBuilder() CancelOrderFixedPointerBuilder {
	return CancelOrderFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *CancelOrderFixedPointerBuilder) Finish() CancelOrderFixedPointer {
	return CancelOrderFixedPointer{m.FixedPointer}
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
func (m CancelOrderFixedPointer) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 36, 4)
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
func (m CancelOrderFixedPointerBuilder) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 36, 4)
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
func (m CancelOrderFixedPointer) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 68, 36)
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
func (m CancelOrderFixedPointerBuilder) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 68, 36)
}

// TradeAccount
func (m CancelOrderFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 100, 68)
}

// TradeAccount
func (m CancelOrderFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 100, 68)
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
func (m CancelOrderFixedPointerBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 36, 4, value)
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
func (m CancelOrderFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 68, 36, value)
}

// SetTradeAccount
func (m CancelOrderFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 100, 68, value)
}

// Copy
func (m CancelOrderFixedPointer) Copy(to CancelOrderFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m CancelOrderFixedPointerBuilder) Copy(to CancelOrderFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CancelOrderFixedPointer) CopyPointer(to CancelOrderFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CancelOrderFixedPointerBuilder) CopyPointer(to CancelOrderFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CancelOrderFixedPointer) CopyTo(to CancelOrderBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CancelOrderFixedPointerBuilder) CopyTo(to CancelOrderBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CancelOrderFixedPointer) CopyToPointer(to CancelOrderPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CancelOrderFixedPointerBuilder) CopyToPointer(to CancelOrderPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}
