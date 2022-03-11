package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size      = SecurityDefinitionForSymbolRequestFixedSize  (88)
//     Type      = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     RequestID = 0
//     Symbol    = ""
//     Exchange  = ""
// }
var _SecurityDefinitionForSymbolRequestFixedDefault = []byte{88, 0, 250, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SecurityDefinitionForSymbolRequestFixedSize = 88

// SecurityDefinitionForSymbolRequestFixed This is a message from the Client to the Server for requesting Security
// Definition data for a specific symbol.
//
// The Server will return a single SecurityDefinitionResponse message in
// response to this request.
//
// The Client must always send a SecurityDefinitionForSymbolRequest message
// to the Server in order to obtain the IntegerToFloatPriceDivisor value
// in case the Server uses the integer market data messages.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SecurityDefinitionForSymbolRequestFixed struct {
	message.Fixed
}

// SecurityDefinitionForSymbolRequestFixedBuilder This is a message from the Client to the Server for requesting Security
// Definition data for a specific symbol.
//
// The Server will return a single SecurityDefinitionResponse message in
// response to this request.
//
// The Client must always send a SecurityDefinitionForSymbolRequest message
// to the Server in order to obtain the IntegerToFloatPriceDivisor value
// in case the Server uses the integer market data messages.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SecurityDefinitionForSymbolRequestFixedBuilder struct {
	message.Fixed
}

func NewSecurityDefinitionForSymbolRequestFixedFrom(b []byte) SecurityDefinitionForSymbolRequestFixed {
	return SecurityDefinitionForSymbolRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapSecurityDefinitionForSymbolRequestFixed(b []byte) SecurityDefinitionForSymbolRequestFixed {
	return SecurityDefinitionForSymbolRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewSecurityDefinitionForSymbolRequestFixed() SecurityDefinitionForSymbolRequestFixedBuilder {
	a := SecurityDefinitionForSymbolRequestFixedBuilder{message.NewFixed(88)}
	a.Unsafe().SetBytes(0, _SecurityDefinitionForSymbolRequestFixedDefault)
	return a
}

// Clear
// {
//     Size      = SecurityDefinitionForSymbolRequestFixedSize  (88)
//     Type      = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     RequestID = 0
//     Symbol    = ""
//     Exchange  = ""
// }
func (m SecurityDefinitionForSymbolRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SecurityDefinitionForSymbolRequestFixedDefault)
}

// ToBuilder
func (m SecurityDefinitionForSymbolRequestFixed) ToBuilder() SecurityDefinitionForSymbolRequestFixedBuilder {
	return SecurityDefinitionForSymbolRequestFixedBuilder{m.Fixed}
}

// Finish
func (m SecurityDefinitionForSymbolRequestFixedBuilder) Finish() SecurityDefinitionForSymbolRequestFixed {
	return SecurityDefinitionForSymbolRequestFixed{m.Fixed}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetExchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 88, 72, value)
}

// Copy
func (m SecurityDefinitionForSymbolRequestFixed) Copy(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// Copy
func (m SecurityDefinitionForSymbolRequestFixedBuilder) Copy(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestFixed) CopyPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestFixedBuilder) CopyPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestFixed) CopyToPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestFixedBuilder) CopyToPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestFixed) CopyTo(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestFixedBuilder) CopyTo(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}
