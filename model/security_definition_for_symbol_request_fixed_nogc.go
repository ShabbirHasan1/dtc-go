package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SecurityDefinitionForSymbolRequestFixedPointer This is a message from the Client to the Server for requesting Security
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
type SecurityDefinitionForSymbolRequestFixedPointer struct {
	message.FixedPointer
}

// SecurityDefinitionForSymbolRequestFixedPointerBuilder This is a message from the Client to the Server for requesting Security
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
type SecurityDefinitionForSymbolRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocSecurityDefinitionForSymbolRequestFixed() SecurityDefinitionForSymbolRequestFixedPointerBuilder {
	a := SecurityDefinitionForSymbolRequestFixedPointerBuilder{message.AllocFixed(88)}
	a.Ptr.SetBytes(0, _SecurityDefinitionForSymbolRequestFixedDefault)
	return a
}

func AllocSecurityDefinitionForSymbolRequestFixedFrom(b []byte) SecurityDefinitionForSymbolRequestFixedPointer {
	return SecurityDefinitionForSymbolRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size      = SecurityDefinitionForSymbolRequestFixedSize  (88)
//     Type      = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     RequestID = 0
//     Symbol    = ""
//     Exchange  = ""
// }
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SecurityDefinitionForSymbolRequestFixedDefault)
}

// ToBuilder
func (m SecurityDefinitionForSymbolRequestFixedPointer) ToBuilder() SecurityDefinitionForSymbolRequestFixedPointerBuilder {
	return SecurityDefinitionForSymbolRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *SecurityDefinitionForSymbolRequestFixedPointerBuilder) Finish() SecurityDefinitionForSymbolRequestFixedPointer {
	return SecurityDefinitionForSymbolRequestFixedPointer{m.FixedPointer}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 88, 72)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 88, 72)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetSymbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 72, 8, value)
}

// SetExchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 88, 72, value)
}

// Copy
func (m SecurityDefinitionForSymbolRequestFixedPointer) Copy(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// Copy
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) Copy(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestFixedPointer) CopyPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) CopyPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestFixedPointer) CopyTo(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) CopyTo(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestFixedPointer) CopyToPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) CopyToPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}
