package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size             = OpenOrdersRequestFixedSize  (76)
//     Type             = OPEN_ORDERS_REQUEST  (300)
//     RequestID        = 0
//     RequestAllOrders = 1
//     ServerOrderID    = ""
//     TradeAccount     = ""
// }
var _OpenOrdersRequestFixedDefault = []byte{76, 0, 44, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const OpenOrdersRequestFixedSize = 76

// OpenOrdersRequestFixed This is a message from the Client to the Server requesting the currently
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
type OpenOrdersRequestFixed struct {
	message.Fixed
}

// OpenOrdersRequestFixedBuilder This is a message from the Client to the Server requesting the currently
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
type OpenOrdersRequestFixedBuilder struct {
	message.Fixed
}

func NewOpenOrdersRequestFixedFrom(b []byte) OpenOrdersRequestFixed {
	return OpenOrdersRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapOpenOrdersRequestFixed(b []byte) OpenOrdersRequestFixed {
	return OpenOrdersRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewOpenOrdersRequestFixed() OpenOrdersRequestFixedBuilder {
	a := OpenOrdersRequestFixedBuilder{message.NewFixed(76)}
	a.Unsafe().SetBytes(0, _OpenOrdersRequestFixedDefault)
	return a
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
func (m OpenOrdersRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _OpenOrdersRequestFixedDefault)
}

// ToBuilder
func (m OpenOrdersRequestFixed) ToBuilder() OpenOrdersRequestFixedBuilder {
	return OpenOrdersRequestFixedBuilder{m.Fixed}
}

// Finish
func (m OpenOrdersRequestFixedBuilder) Finish() OpenOrdersRequestFixed {
	return OpenOrdersRequestFixed{m.Fixed}
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestFixed) RequestAllOrders() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// RequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestFixedBuilder) RequestAllOrders() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// ServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestFixed) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 44, 12)
}

// ServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestFixedBuilder) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 44, 12)
}

// TradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 76, 44)
}

// TradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 76, 44)
}

// SetRequestID A unique request identifier for this request.
func (m OpenOrdersRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestFixedBuilder) SetRequestAllOrders(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestFixedBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 44, 12, value)
}

// SetTradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 76, 44, value)
}

// Copy
func (m OpenOrdersRequestFixed) Copy(to OpenOrdersRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m OpenOrdersRequestFixedBuilder) Copy(to OpenOrdersRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m OpenOrdersRequestFixed) CopyPointer(to OpenOrdersRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m OpenOrdersRequestFixedBuilder) CopyPointer(to OpenOrdersRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m OpenOrdersRequestFixed) CopyToPointer(to OpenOrdersRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m OpenOrdersRequestFixedBuilder) CopyToPointer(to OpenOrdersRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m OpenOrdersRequestFixed) CopyTo(to OpenOrdersRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m OpenOrdersRequestFixedBuilder) CopyTo(to OpenOrdersRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}
