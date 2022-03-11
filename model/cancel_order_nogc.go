package model

import (
	"github.com/moontrade/dtc-go/message"
)

// CancelOrderPointer This is a message from the Client to the Server requesting a previously
// sent order to be canceled.
type CancelOrderPointer struct {
	message.VLSPointer
}

// CancelOrderPointerBuilder This is a message from the Client to the Server requesting a previously
// sent order to be canceled.
type CancelOrderPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocCancelOrder() CancelOrderPointerBuilder {
	a := CancelOrderPointerBuilder{message.AllocVLSBuilder(18)}
	a.Ptr.SetBytes(0, _CancelOrderDefault)
	return a
}

func AllocCancelOrderFrom(b []byte) CancelOrderPointer {
	return CancelOrderPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m CancelOrderPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CancelOrderDefault)
}

// ToBuilder
func (m CancelOrderPointer) ToBuilder() CancelOrderPointerBuilder {
	return CancelOrderPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *CancelOrderPointerBuilder) Finish() CancelOrderPointer {
	return CancelOrderPointer{m.VLSPointerBuilder.Finish()}
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
func (m CancelOrderPointer) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 10, 6)
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
func (m CancelOrderPointerBuilder) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 10, 6)
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
func (m CancelOrderPointer) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 14, 10)
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
func (m CancelOrderPointerBuilder) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m CancelOrderPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TradeAccount
func (m CancelOrderPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
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
func (m *CancelOrderPointerBuilder) SetServerOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
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
func (m *CancelOrderPointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *CancelOrderPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// Copy
func (m CancelOrderPointer) Copy(to CancelOrderBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m CancelOrderPointerBuilder) Copy(to CancelOrderBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CancelOrderPointer) CopyPointer(to CancelOrderPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CancelOrderPointerBuilder) CopyPointer(to CancelOrderPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CancelOrderPointer) CopyTo(to CancelOrderFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CancelOrderPointerBuilder) CopyTo(to CancelOrderFixedBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CancelOrderPointer) CopyToPointer(to CancelOrderFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CancelOrderPointerBuilder) CopyToPointer(to CancelOrderFixedPointerBuilder) {
	to.SetServerOrderID(m.ServerOrderID())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetTradeAccount(m.TradeAccount())
}
