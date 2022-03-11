package model

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeAccountsRequestSize = 8

// {
//     Size      = TradeAccountsRequestSize  (8)
//     Type      = TRADE_ACCOUNTS_REQUEST  (400)
//     RequestID = 0
// }
var _TradeAccountsRequestDefault = []byte{8, 0, 144, 1, 0, 0, 0, 0}

// TradeAccountsRequest This is a message from the Client to the Server to request all of the
// account identifiers for the logged in Username.
//
// If there are no accounts available, then the Server needs to respond with
// at least one TradeAccountResponse message containing an empty Trade Account.
// at least one TradeAccountResponse message containing an empty Trade Account.
type TradeAccountsRequest struct {
	message.Fixed
}

// TradeAccountsRequestBuilder This is a message from the Client to the Server to request all of the
// account identifiers for the logged in Username.
//
// If there are no accounts available, then the Server needs to respond with
// at least one TradeAccountResponse message containing an empty Trade Account.
// at least one TradeAccountResponse message containing an empty Trade Account.
type TradeAccountsRequestBuilder struct {
	message.Fixed
}

// TradeAccountsRequestPointer This is a message from the Client to the Server to request all of the
// account identifiers for the logged in Username.
//
// If there are no accounts available, then the Server needs to respond with
// at least one TradeAccountResponse message containing an empty Trade Account.
// at least one TradeAccountResponse message containing an empty Trade Account.
type TradeAccountsRequestPointer struct {
	message.FixedPointer
}

// TradeAccountsRequestPointerBuilder This is a message from the Client to the Server to request all of the
// account identifiers for the logged in Username.
//
// If there are no accounts available, then the Server needs to respond with
// at least one TradeAccountResponse message containing an empty Trade Account.
// at least one TradeAccountResponse message containing an empty Trade Account.
type TradeAccountsRequestPointerBuilder struct {
	message.FixedPointer
}

func NewTradeAccountsRequestFrom(b []byte) TradeAccountsRequest {
	return TradeAccountsRequest{Fixed: message.NewFixedFrom(b)}
}

func WrapTradeAccountsRequest(b []byte) TradeAccountsRequest {
	return TradeAccountsRequest{Fixed: message.WrapFixed(b)}
}

func NewTradeAccountsRequest() TradeAccountsRequestBuilder {
	a := TradeAccountsRequestBuilder{message.NewFixed(8)}
	a.Unsafe().SetBytes(0, _TradeAccountsRequestDefault)
	return a
}

func AllocTradeAccountsRequest() TradeAccountsRequestPointerBuilder {
	a := TradeAccountsRequestPointerBuilder{message.AllocFixed(8)}
	a.Ptr.SetBytes(0, _TradeAccountsRequestDefault)
	return a
}

func AllocTradeAccountsRequestFrom(b []byte) TradeAccountsRequestPointer {
	return TradeAccountsRequestPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size      = TradeAccountsRequestSize  (8)
//     Type      = TRADE_ACCOUNTS_REQUEST  (400)
//     RequestID = 0
// }
func (m TradeAccountsRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountsRequestDefault)
}

// Clear
// {
//     Size      = TradeAccountsRequestSize  (8)
//     Type      = TRADE_ACCOUNTS_REQUEST  (400)
//     RequestID = 0
// }
func (m TradeAccountsRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountsRequestDefault)
}

// ToBuilder
func (m TradeAccountsRequest) ToBuilder() TradeAccountsRequestBuilder {
	return TradeAccountsRequestBuilder{m.Fixed}
}

// ToBuilder
func (m TradeAccountsRequestPointer) ToBuilder() TradeAccountsRequestPointerBuilder {
	return TradeAccountsRequestPointerBuilder{m.FixedPointer}
}

// Finish
func (m TradeAccountsRequestBuilder) Finish() TradeAccountsRequest {
	return TradeAccountsRequest{m.Fixed}
}

// Finish
func (m *TradeAccountsRequestPointerBuilder) Finish() TradeAccountsRequestPointer {
	return TradeAccountsRequestPointer{m.FixedPointer}
}

// RequestID A unique request identifier for this request.
func (m TradeAccountsRequest) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID A unique request identifier for this request.
func (m TradeAccountsRequestBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID A unique request identifier for this request.
func (m TradeAccountsRequestPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID A unique request identifier for this request.
func (m TradeAccountsRequestPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// SetRequestID A unique request identifier for this request.
func (m TradeAccountsRequestBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID A unique request identifier for this request.
func (m TradeAccountsRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// Copy
func (m TradeAccountsRequest) Copy(to TradeAccountsRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// Copy
func (m TradeAccountsRequestBuilder) Copy(to TradeAccountsRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m TradeAccountsRequest) CopyPointer(to TradeAccountsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m TradeAccountsRequestBuilder) CopyPointer(to TradeAccountsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}

// Copy
func (m TradeAccountsRequestPointer) Copy(to TradeAccountsRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// Copy
func (m TradeAccountsRequestPointerBuilder) Copy(to TradeAccountsRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m TradeAccountsRequestPointer) CopyPointer(to TradeAccountsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m TradeAccountsRequestPointerBuilder) CopyPointer(to TradeAccountsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}