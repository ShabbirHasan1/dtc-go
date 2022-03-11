package model

import (
	"github.com/moontrade/dtc-go/message"
)

// OpenOrdersRequestPointer This is a message from the Client to the Server requesting the currently
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
type OpenOrdersRequestPointer struct {
	message.VLSPointer
}

// OpenOrdersRequestPointerBuilder This is a message from the Client to the Server requesting the currently
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
type OpenOrdersRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocOpenOrdersRequest() OpenOrdersRequestPointerBuilder {
	a := OpenOrdersRequestPointerBuilder{message.AllocVLSBuilder(24)}
	a.Ptr.SetBytes(0, _OpenOrdersRequestDefault)
	return a
}

func AllocOpenOrdersRequestFrom(b []byte) OpenOrdersRequestPointer {
	return OpenOrdersRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size             = OpenOrdersRequestSize  (24)
//     Type             = OPEN_ORDERS_REQUEST  (300)
//     BaseSize         = OpenOrdersRequestSize  (24)
//     RequestID        = 0
//     RequestAllOrders = 1
//     ServerOrderID    = ""
//     TradeAccount     = ""
// }
func (m OpenOrdersRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _OpenOrdersRequestDefault)
}

// ToBuilder
func (m OpenOrdersRequestPointer) ToBuilder() OpenOrdersRequestPointerBuilder {
	return OpenOrdersRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *OpenOrdersRequestPointerBuilder) Finish() OpenOrdersRequestPointer {
	return OpenOrdersRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestPointer) RequestAllOrders() int32 {
	return message.Int32VLS(m.Ptr, 16, 12)
}

// RequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestPointerBuilder) RequestAllOrders() int32 {
	return message.Int32VLS(m.Ptr, 16, 12)
}

// ServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestPointer) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// ServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestPointerBuilder) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// TradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// TradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// SetRequestID A unique request identifier for this request.
func (m OpenOrdersRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestPointerBuilder) SetRequestAllOrders(value int32) {
	message.SetInt32VLS(m.Ptr, 16, 12, value)
}

// SetServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m *OpenOrdersRequestPointerBuilder) SetServerOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetTradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m *OpenOrdersRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// Copy
func (m OpenOrdersRequestPointer) Copy(to OpenOrdersRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m OpenOrdersRequestPointerBuilder) Copy(to OpenOrdersRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m OpenOrdersRequestPointer) CopyPointer(to OpenOrdersRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m OpenOrdersRequestPointerBuilder) CopyPointer(to OpenOrdersRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m OpenOrdersRequestPointer) CopyTo(to OpenOrdersRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m OpenOrdersRequestPointerBuilder) CopyTo(to OpenOrdersRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m OpenOrdersRequestPointer) CopyToPointer(to OpenOrdersRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m OpenOrdersRequestPointerBuilder) CopyToPointer(to OpenOrdersRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}
