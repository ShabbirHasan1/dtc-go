package model

import (
	"github.com/moontrade/dtc-go/message"
)

// OpenOrdersRequestFixedPointer This is a message from the Client to the Server requesting the currently
// open orders.
//
// The Server will send open/working orders in response to this request through
// OrderUpdate messages.
//
// The Server will not return canceled or filled orders.
//
// When the Server responds to this request, it needs to respond with a separate
// OrderUpdate for each order.
//
// When the Server responds to this request, OrderUpdateReason in the OrderUpdate
// message must be set to OpenOrdersRequest_RESPONSE indicating the orders
// are being restated.
//
// If there are no Open orders, the Server will send back 1 OrderUpdate message
// with only the TotalNumberMessages, MessageNumber, RequestID, OrderUpdateReason,
// NoOrders = 1 fields set in the OrderUpdate message.
type OpenOrdersRequestFixedPointer struct {
	message.FixedPointer
}

// OpenOrdersRequestFixedPointerBuilder This is a message from the Client to the Server requesting the currently
// open orders.
//
// The Server will send open/working orders in response to this request through
// OrderUpdate messages.
//
// The Server will not return canceled or filled orders.
//
// When the Server responds to this request, it needs to respond with a separate
// OrderUpdate for each order.
//
// When the Server responds to this request, OrderUpdateReason in the OrderUpdate
// message must be set to OpenOrdersRequest_RESPONSE indicating the orders
// are being restated.
//
// If there are no Open orders, the Server will send back 1 OrderUpdate message
// with only the TotalNumberMessages, MessageNumber, RequestID, OrderUpdateReason,
// NoOrders = 1 fields set in the OrderUpdate message.
type OpenOrdersRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocOpenOrdersRequestFixed() OpenOrdersRequestFixedPointerBuilder {
	a := OpenOrdersRequestFixedPointerBuilder{message.AllocFixed(76)}
	a.Ptr.SetBytes(0, _OpenOrdersRequestFixedDefault)
	return a
}

func AllocOpenOrdersRequestFixedFrom(b []byte) OpenOrdersRequestFixedPointer {
	return OpenOrdersRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size             = OpenOrdersRequestFixedSize  (76)
//     Type             = OPEN_ORDERS_REQUEST  (300)
//     RequestID        = 0
//     RequestAllOrders = 1
//     ServerOrderID    = ""
//     TradeAccount     = ""
// }
func (m OpenOrdersRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _OpenOrdersRequestFixedDefault)
}

// ToBuilder
func (m OpenOrdersRequestFixedPointer) ToBuilder() OpenOrdersRequestFixedPointerBuilder {
	return OpenOrdersRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *OpenOrdersRequestFixedPointerBuilder) Finish() OpenOrdersRequestFixedPointer {
	return OpenOrdersRequestFixedPointer{m.FixedPointer}
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestFixedPointer) RequestAllOrders() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// RequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestFixedPointerBuilder) RequestAllOrders() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// ServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestFixedPointer) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 44, 12)
}

// ServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestFixedPointerBuilder) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 44, 12)
}

// TradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 76, 44)
}

// TradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 76, 44)
}

// SetRequestID A unique request identifier for this request.
func (m OpenOrdersRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestFixedPointerBuilder) SetRequestAllOrders(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestFixedPointerBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 44, 12, value)
}

// SetTradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 76, 44, value)
}

// Copy
func (m OpenOrdersRequestFixedPointer) Copy(to OpenOrdersRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m OpenOrdersRequestFixedPointerBuilder) Copy(to OpenOrdersRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m OpenOrdersRequestFixedPointer) CopyPointer(to OpenOrdersRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m OpenOrdersRequestFixedPointerBuilder) CopyPointer(to OpenOrdersRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m OpenOrdersRequestFixedPointer) CopyTo(to OpenOrdersRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m OpenOrdersRequestFixedPointerBuilder) CopyTo(to OpenOrdersRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m OpenOrdersRequestFixedPointer) CopyToPointer(to OpenOrdersRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m OpenOrdersRequestFixedPointerBuilder) CopyToPointer(to OpenOrdersRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}
