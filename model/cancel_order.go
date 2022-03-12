package model

import (
	"github.com/moontrade/dtc-go/message"
)

const CancelOrderSize = 18

const CancelOrderFixedSize = 100

// {
//     Size          = CancelOrderSize  (18)
//     Type          = CANCEL_ORDER  (203)
//     BaseSize      = CancelOrderSize  (18)
//     ServerOrderID = ""
//     ClientOrderID = ""
//     TradeAccount  = ""
// }
var _CancelOrderDefault = []byte{18, 0, 203, 0, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size          = CancelOrderFixedSize  (100)
//     Type          = CANCEL_ORDER  (203)
//     ServerOrderID = ""
//     ClientOrderID = ""
//     TradeAccount  = ""
// }
var _CancelOrderFixedDefault = []byte{100, 0, 203, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocCancelOrder() CancelOrderPointerBuilder {
	a := CancelOrderPointerBuilder{message.AllocVLSBuilder(18)}
	a.Ptr.SetBytes(0, _CancelOrderDefault)
	return a
}

func AllocCancelOrderFrom(b []byte) CancelOrderPointer {
	return CancelOrderPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m CancelOrder) ToBuilder() CancelOrderBuilder {
	return CancelOrderBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m CancelOrderFixed) ToBuilder() CancelOrderFixedBuilder {
	return CancelOrderFixedBuilder{m.Fixed}
}

// ToBuilder
func (m CancelOrderPointer) ToBuilder() CancelOrderPointerBuilder {
	return CancelOrderPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m CancelOrderFixedPointer) ToBuilder() CancelOrderFixedPointerBuilder {
	return CancelOrderFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m CancelOrderBuilder) Finish() CancelOrder {
	return CancelOrder{m.VLSBuilder.Finish()}
}

// Finish
func (m CancelOrderFixedBuilder) Finish() CancelOrderFixed {
	return CancelOrderFixed{m.Fixed}
}

// Finish
func (m *CancelOrderPointerBuilder) Finish() CancelOrderPointer {
	return CancelOrderPointer{m.VLSPointerBuilder.Finish()}
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
func (m CancelOrder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount
func (m CancelOrderBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount
func (m CancelOrderPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TradeAccount
func (m CancelOrderPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
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
func (m CancelOrderFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 100, 68)
}

// TradeAccount
func (m CancelOrderFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 100, 68)
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
func (m *CancelOrderBuilder) SetServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
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
func (m *CancelOrderBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
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
func (m *CancelOrderBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetTradeAccount
func (m *CancelOrderPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
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
func (m CancelOrderFixedBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 68, 36, value)
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
func (m CancelOrderFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 100, 68, value)
}

// SetTradeAccount
func (m CancelOrderFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 100, 68, value)
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
