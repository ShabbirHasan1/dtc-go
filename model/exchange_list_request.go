package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size      = ExchangeListRequestSize  (8)
//     Type      = EXCHANGE_LIST_REQUEST  (500)
//     RequestID = 0
// }
var _ExchangeListRequestDefault = []byte{8, 0, 244, 1, 0, 0, 0, 0}

const ExchangeListRequestSize = 8

// ExchangeListRequest This is a message from the Client to the Server to request a list of all
// available exchanges from the Server.
//
// The server will respond with a separate ExchangeListResponse message for
// each exchange.
//
// In the case where the Server does not specify an exchange with its symbols,
// then the Server should provide a single response with an empty Exchange.
// then the Server should provide a single response with an empty Exchange.
type ExchangeListRequest struct {
	message.Fixed
}

// ExchangeListRequestBuilder This is a message from the Client to the Server to request a list of all
// available exchanges from the Server.
//
// The server will respond with a separate ExchangeListResponse message for
// each exchange.
//
// In the case where the Server does not specify an exchange with its symbols,
// then the Server should provide a single response with an empty Exchange.
// then the Server should provide a single response with an empty Exchange.
type ExchangeListRequestBuilder struct {
	message.Fixed
}

func NewExchangeListRequestFrom(b []byte) ExchangeListRequest {
	return ExchangeListRequest{Fixed: message.NewFixedFrom(b)}
}

func WrapExchangeListRequest(b []byte) ExchangeListRequest {
	return ExchangeListRequest{Fixed: message.WrapFixed(b)}
}

func NewExchangeListRequest() ExchangeListRequestBuilder {
	a := ExchangeListRequestBuilder{message.NewFixed(8)}
	a.Unsafe().SetBytes(0, _ExchangeListRequestDefault)
	return a
}

// Clear
// {
//     Size      = ExchangeListRequestSize  (8)
//     Type      = EXCHANGE_LIST_REQUEST  (500)
//     RequestID = 0
// }
func (m ExchangeListRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ExchangeListRequestDefault)
}

// ToBuilder
func (m ExchangeListRequest) ToBuilder() ExchangeListRequestBuilder {
	return ExchangeListRequestBuilder{m.Fixed}
}

// Finish
func (m ExchangeListRequestBuilder) Finish() ExchangeListRequest {
	return ExchangeListRequest{m.Fixed}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the ExchangeListResponse message.
func (m ExchangeListRequest) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the ExchangeListResponse message.
func (m ExchangeListRequestBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the ExchangeListResponse message.
func (m ExchangeListRequestBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// Copy
func (m ExchangeListRequest) Copy(to ExchangeListRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// Copy
func (m ExchangeListRequestBuilder) Copy(to ExchangeListRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m ExchangeListRequest) CopyPointer(to ExchangeListRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m ExchangeListRequestBuilder) CopyPointer(to ExchangeListRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}
