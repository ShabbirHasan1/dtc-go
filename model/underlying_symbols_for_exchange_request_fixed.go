package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = UnderlyingSymbolsForExchangeRequestFixedSize  (28)
//     Type         = UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST  (503)
//     RequestID    = 0
//     Exchange     = ""
//     SecurityType = 0
// }
var _UnderlyingSymbolsForExchangeRequestFixedDefault = []byte{28, 0, 247, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const UnderlyingSymbolsForExchangeRequestFixedSize = 28

// UnderlyingSymbolsForExchangeRequestFixed This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestFixed struct {
	message.Fixed
}

// UnderlyingSymbolsForExchangeRequestFixedBuilder This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestFixedBuilder struct {
	message.Fixed
}

func NewUnderlyingSymbolsForExchangeRequestFixedFrom(b []byte) UnderlyingSymbolsForExchangeRequestFixed {
	return UnderlyingSymbolsForExchangeRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapUnderlyingSymbolsForExchangeRequestFixed(b []byte) UnderlyingSymbolsForExchangeRequestFixed {
	return UnderlyingSymbolsForExchangeRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewUnderlyingSymbolsForExchangeRequestFixed() UnderlyingSymbolsForExchangeRequestFixedBuilder {
	a := UnderlyingSymbolsForExchangeRequestFixedBuilder{message.NewFixed(28)}
	a.Unsafe().SetBytes(0, _UnderlyingSymbolsForExchangeRequestFixedDefault)
	return a
}

// Clear
// {
//     Size         = UnderlyingSymbolsForExchangeRequestFixedSize  (28)
//     Type         = UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST  (503)
//     RequestID    = 0
//     Exchange     = ""
//     SecurityType = 0
// }
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _UnderlyingSymbolsForExchangeRequestFixedDefault)
}

// ToBuilder
func (m UnderlyingSymbolsForExchangeRequestFixed) ToBuilder() UnderlyingSymbolsForExchangeRequestFixedBuilder {
	return UnderlyingSymbolsForExchangeRequestFixedBuilder{m.Fixed}
}

// Finish
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) Finish() UnderlyingSymbolsForExchangeRequestFixed {
	return UnderlyingSymbolsForExchangeRequestFixed{m.Fixed}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 24, 8)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 24, 8)
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestFixed) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 28, 24))
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 28, 24))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetExchange .
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 24, 8, value)
}

// SetSecurityType .
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 28, 24, int32(value))
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestFixed) Copy(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) Copy(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestFixed) CopyPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) CopyPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestFixed) CopyToPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) CopyToPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestFixed) CopyTo(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) CopyTo(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}
