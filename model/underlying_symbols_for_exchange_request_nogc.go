package model

import (
	"github.com/moontrade/dtc-go/message"
)

// UnderlyingSymbolsForExchangeRequestPointer This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestPointer struct {
	message.VLSPointer
}

// UnderlyingSymbolsForExchangeRequestPointerBuilder This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocUnderlyingSymbolsForExchangeRequest() UnderlyingSymbolsForExchangeRequestPointerBuilder {
	a := UnderlyingSymbolsForExchangeRequestPointerBuilder{message.AllocVLSBuilder(20)}
	a.Ptr.SetBytes(0, _UnderlyingSymbolsForExchangeRequestDefault)
	return a
}

func AllocUnderlyingSymbolsForExchangeRequestFrom(b []byte) UnderlyingSymbolsForExchangeRequestPointer {
	return UnderlyingSymbolsForExchangeRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _UnderlyingSymbolsForExchangeRequestDefault)
}

// ToBuilder
func (m UnderlyingSymbolsForExchangeRequestPointer) ToBuilder() UnderlyingSymbolsForExchangeRequestPointerBuilder {
	return UnderlyingSymbolsForExchangeRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *UnderlyingSymbolsForExchangeRequestPointerBuilder) Finish() UnderlyingSymbolsForExchangeRequestPointer {
	return UnderlyingSymbolsForExchangeRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestPointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 20, 16))
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 20, 16))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetExchange .
func (m *UnderlyingSymbolsForExchangeRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetSecurityType .
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Ptr, 20, 16, int32(value))
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestPointer) Copy(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) Copy(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestPointer) CopyPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) CopyPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestPointer) CopyTo(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) CopyTo(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestPointer) CopyToPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) CopyToPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}
