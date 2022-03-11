package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size             = OpenOrdersRequestSize  (24)
//     Type             = OPEN_ORDERS_REQUEST  (300)
//     BaseSize         = OpenOrdersRequestSize  (24)
//     RequestID        = 0
//     RequestAllOrders = 1
//     ServerOrderID    = ""
//     TradeAccount     = ""
// }
var _OpenOrdersRequestDefault = []byte{24, 0, 44, 1, 24, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const OpenOrdersRequestSize = 24

// OpenOrdersRequest This is a message from the Client to the Server requesting the currently
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
type OpenOrdersRequest struct {
	message.VLS
}

// OpenOrdersRequestBuilder This is a message from the Client to the Server requesting the currently
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
type OpenOrdersRequestBuilder struct {
	message.VLSBuilder
}

func NewOpenOrdersRequestFrom(b []byte) OpenOrdersRequest {
	return OpenOrdersRequest{VLS: message.NewVLSFrom(b)}
}

func WrapOpenOrdersRequest(b []byte) OpenOrdersRequest {
	return OpenOrdersRequest{VLS: message.WrapVLS(b)}
}

func NewOpenOrdersRequest() OpenOrdersRequestBuilder {
	a := OpenOrdersRequestBuilder{message.NewVLSBuilder(24)}
	a.Unsafe().SetBytes(0, _OpenOrdersRequestDefault)
	return a
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
func (m OpenOrdersRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _OpenOrdersRequestDefault)
}

// ToBuilder
func (m OpenOrdersRequest) ToBuilder() OpenOrdersRequestBuilder {
	return OpenOrdersRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m OpenOrdersRequestBuilder) Finish() OpenOrdersRequest {
	return OpenOrdersRequest{m.VLSBuilder.Finish()}
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequest) RequestAllOrders() int32 {
	return message.Int32VLS(m.Unsafe(), 16, 12)
}

// RequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestBuilder) RequestAllOrders() int32 {
	return message.Int32VLS(m.Unsafe(), 16, 12)
}

// ServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequest) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// ServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestBuilder) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// TradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// TradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// SetRequestID A unique request identifier for this request.
func (m OpenOrdersRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestBuilder) SetRequestAllOrders(value int32) {
	message.SetInt32VLS(m.Unsafe(), 16, 12, value)
}

// SetServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m *OpenOrdersRequestBuilder) SetServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetTradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m *OpenOrdersRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// Copy
func (m OpenOrdersRequest) Copy(to OpenOrdersRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m OpenOrdersRequestBuilder) Copy(to OpenOrdersRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m OpenOrdersRequest) CopyPointer(to OpenOrdersRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m OpenOrdersRequestBuilder) CopyPointer(to OpenOrdersRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m OpenOrdersRequest) CopyToPointer(to OpenOrdersRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m OpenOrdersRequestBuilder) CopyToPointer(to OpenOrdersRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m OpenOrdersRequest) CopyTo(to OpenOrdersRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m OpenOrdersRequestBuilder) CopyTo(to OpenOrdersRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}
