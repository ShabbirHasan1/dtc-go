package model

import (
	"github.com/moontrade/dtc-go/message"
)

// UnderlyingSymbolsForExchangeRequestFixedPointer This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestFixedPointer struct {
	message.FixedPointer
}

// UnderlyingSymbolsForExchangeRequestFixedPointerBuilder This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocUnderlyingSymbolsForExchangeRequestFixed() UnderlyingSymbolsForExchangeRequestFixedPointerBuilder {
	a := UnderlyingSymbolsForExchangeRequestFixedPointerBuilder{message.AllocFixed(28)}
	a.Ptr.SetBytes(0, _UnderlyingSymbolsForExchangeRequestFixedDefault)
	return a
}

func AllocUnderlyingSymbolsForExchangeRequestFixedFrom(b []byte) UnderlyingSymbolsForExchangeRequestFixedPointer {
	return UnderlyingSymbolsForExchangeRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size         = UnderlyingSymbolsForExchangeRequestFixedSize  (28)
//     Type         = UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST  (503)
//     RequestID    = 0
//     Exchange     = ""
//     SecurityType = 0
// }
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _UnderlyingSymbolsForExchangeRequestFixedDefault)
}

// ToBuilder
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) ToBuilder() UnderlyingSymbolsForExchangeRequestFixedPointerBuilder {
	return UnderlyingSymbolsForExchangeRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) Finish() UnderlyingSymbolsForExchangeRequestFixedPointer {
	return UnderlyingSymbolsForExchangeRequestFixedPointer{m.FixedPointer}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 24, 8)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 24, 8)
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Ptr, 28, 24))
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Ptr, 28, 24))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetExchange .
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 24, 8, value)
}

// SetSecurityType .
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 28, 24, int32(value))
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) Copy(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) Copy(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) CopyPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) CopyPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) CopyTo(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) CopyTo(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) CopyToPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) CopyToPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}
