package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size      = TradeAccountsRequestSize  (8)
//     Type      = TRADE_ACCOUNTS_REQUEST  (400)
//     RequestID = 0
// }
var _TradeAccountsRequestDefault = []byte{8, 0, 144, 1, 0, 0, 0, 0}

const TradeAccountsRequestSize = 8

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

// Clear
// {
//     Size      = TradeAccountsRequestSize  (8)
//     Type      = TRADE_ACCOUNTS_REQUEST  (400)
//     RequestID = 0
// }
func (m TradeAccountsRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountsRequestDefault)
}

// ToBuilder
func (m TradeAccountsRequest) ToBuilder() TradeAccountsRequestBuilder {
	return TradeAccountsRequestBuilder{m.Fixed}
}

// Finish
func (m TradeAccountsRequestBuilder) Finish() TradeAccountsRequest {
	return TradeAccountsRequest{m.Fixed}
}

// RequestID A unique request identifier for this request.
func (m TradeAccountsRequest) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID A unique request identifier for this request.
func (m TradeAccountsRequestBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// SetRequestID A unique request identifier for this request.
func (m TradeAccountsRequestBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
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
