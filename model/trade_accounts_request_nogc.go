package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
func (m TradeAccountsRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountsRequestDefault)
}

// ToBuilder
func (m TradeAccountsRequestPointer) ToBuilder() TradeAccountsRequestPointerBuilder {
	return TradeAccountsRequestPointerBuilder{m.FixedPointer}
}

// Finish
func (m *TradeAccountsRequestPointerBuilder) Finish() TradeAccountsRequestPointer {
	return TradeAccountsRequestPointer{m.FixedPointer}
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
func (m TradeAccountsRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
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
