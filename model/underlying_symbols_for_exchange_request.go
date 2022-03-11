package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = UnderlyingSymbolsForExchangeRequestSize  (20)
//     Type         = UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST  (503)
//     BaseSize     = UnderlyingSymbolsForExchangeRequestSize  (20)
//     RequestID    = 0
//     Exchange     = ""
//     SecurityType = 0
// }
var _UnderlyingSymbolsForExchangeRequestDefault = []byte{20, 0, 247, 1, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const UnderlyingSymbolsForExchangeRequestSize = 20

// UnderlyingSymbolsForExchangeRequest This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequest struct {
	message.VLS
}

// UnderlyingSymbolsForExchangeRequestBuilder This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestBuilder struct {
	message.VLSBuilder
}

func NewUnderlyingSymbolsForExchangeRequestFrom(b []byte) UnderlyingSymbolsForExchangeRequest {
	return UnderlyingSymbolsForExchangeRequest{VLS: message.NewVLSFrom(b)}
}

func WrapUnderlyingSymbolsForExchangeRequest(b []byte) UnderlyingSymbolsForExchangeRequest {
	return UnderlyingSymbolsForExchangeRequest{VLS: message.WrapVLS(b)}
}

func NewUnderlyingSymbolsForExchangeRequest() UnderlyingSymbolsForExchangeRequestBuilder {
	a := UnderlyingSymbolsForExchangeRequestBuilder{message.NewVLSBuilder(20)}
	a.Unsafe().SetBytes(0, _UnderlyingSymbolsForExchangeRequestDefault)
	return a
}

// Clear
// {
//     Size         = UnderlyingSymbolsForExchangeRequestSize  (20)
//     Type         = UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST  (503)
//     BaseSize     = UnderlyingSymbolsForExchangeRequestSize  (20)
//     RequestID    = 0
//     Exchange     = ""
//     SecurityType = 0
// }
func (m UnderlyingSymbolsForExchangeRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _UnderlyingSymbolsForExchangeRequestDefault)
}

// ToBuilder
func (m UnderlyingSymbolsForExchangeRequest) ToBuilder() UnderlyingSymbolsForExchangeRequestBuilder {
	return UnderlyingSymbolsForExchangeRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m UnderlyingSymbolsForExchangeRequestBuilder) Finish() UnderlyingSymbolsForExchangeRequest {
	return UnderlyingSymbolsForExchangeRequest{m.VLSBuilder.Finish()}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequest) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 20, 16))
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 20, 16))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetExchange .
func (m *UnderlyingSymbolsForExchangeRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetSecurityType .
func (m UnderlyingSymbolsForExchangeRequestBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 20, 16, int32(value))
}

// Copy
func (m UnderlyingSymbolsForExchangeRequest) Copy(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestBuilder) Copy(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequest) CopyPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestBuilder) CopyPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequest) CopyToPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestBuilder) CopyToPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequest) CopyTo(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestBuilder) CopyTo(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}
