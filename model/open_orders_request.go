package model

import (
	"github.com/moontrade/dtc-go/message"
)

const OpenOrdersRequestSize = 24

const OpenOrdersRequestFixedSize = 76

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

// {
//     Size             = OpenOrdersRequestFixedSize  (76)
//     Type             = OPEN_ORDERS_REQUEST  (300)
//     RequestID        = 0
//     RequestAllOrders = 1
//     ServerOrderID    = ""
//     TradeAccount     = ""
// }
var _OpenOrdersRequestFixedDefault = []byte{76, 0, 44, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocOpenOrdersRequest() OpenOrdersRequestPointerBuilder {
	a := OpenOrdersRequestPointerBuilder{message.AllocVLSBuilder(24)}
	a.Ptr.SetBytes(0, _OpenOrdersRequestDefault)
	return a
}

func AllocOpenOrdersRequestFrom(b []byte) OpenOrdersRequestPointer {
	return OpenOrdersRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m OpenOrdersRequest) ToBuilder() OpenOrdersRequestBuilder {
	return OpenOrdersRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m OpenOrdersRequestFixed) ToBuilder() OpenOrdersRequestFixedBuilder {
	return OpenOrdersRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m OpenOrdersRequestPointer) ToBuilder() OpenOrdersRequestPointerBuilder {
	return OpenOrdersRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m OpenOrdersRequestFixedPointer) ToBuilder() OpenOrdersRequestFixedPointerBuilder {
	return OpenOrdersRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m OpenOrdersRequestBuilder) Finish() OpenOrdersRequest {
	return OpenOrdersRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m OpenOrdersRequestFixedBuilder) Finish() OpenOrdersRequestFixed {
	return OpenOrdersRequestFixed{m.Fixed}
}

// Finish
func (m *OpenOrdersRequestPointerBuilder) Finish() OpenOrdersRequestPointer {
	return OpenOrdersRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *OpenOrdersRequestFixedPointerBuilder) Finish() OpenOrdersRequestFixedPointer {
	return OpenOrdersRequestFixedPointer{m.FixedPointer}
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
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
func (m OpenOrdersRequest) RequestAllOrders() int32 {
	return message.Int32VLS(m.Unsafe(), 16, 12)
}

// RequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestBuilder) RequestAllOrders() int32 {
	return message.Int32VLS(m.Unsafe(), 16, 12)
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
func (m OpenOrdersRequest) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// ServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestBuilder) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
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
func (m OpenOrdersRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// TradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
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

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m OpenOrdersRequestFixed) RequestAllOrders() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// RequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestFixedBuilder) RequestAllOrders() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
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
func (m OpenOrdersRequestFixed) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 44, 12)
}

// ServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestFixedBuilder) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 44, 12)
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
func (m OpenOrdersRequestFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 76, 44)
}

// TradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 76, 44)
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
func (m OpenOrdersRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID A unique request identifier for this request.
func (m OpenOrdersRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestBuilder) SetRequestAllOrders(value int32) {
	message.SetInt32VLS(m.Unsafe(), 16, 12, value)
}

// SetRequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestPointerBuilder) SetRequestAllOrders(value int32) {
	message.SetInt32VLS(m.Ptr, 16, 12, value)
}

// SetServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m *OpenOrdersRequestBuilder) SetServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m *OpenOrdersRequestPointerBuilder) SetServerOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetTradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m *OpenOrdersRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetTradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m *OpenOrdersRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetRequestID A unique request identifier for this request.
func (m OpenOrdersRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID A unique request identifier for this request.
func (m OpenOrdersRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestFixedBuilder) SetRequestAllOrders(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetRequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestFixedPointerBuilder) SetRequestAllOrders(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestFixedBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 44, 12, value)
}

// SetServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestFixedPointerBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 44, 12, value)
}

// SetTradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 76, 44, value)
}

// SetTradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 76, 44, value)
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
