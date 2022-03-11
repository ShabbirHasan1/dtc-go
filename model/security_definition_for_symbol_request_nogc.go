package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SecurityDefinitionForSymbolRequestPointer This is a message from the Client to the Server for requesting Security
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
type SecurityDefinitionForSymbolRequestPointer struct {
	message.VLSPointer
}

// SecurityDefinitionForSymbolRequestPointerBuilder This is a message from the Client to the Server for requesting Security
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
type SecurityDefinitionForSymbolRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocSecurityDefinitionForSymbolRequest() SecurityDefinitionForSymbolRequestPointerBuilder {
	a := SecurityDefinitionForSymbolRequestPointerBuilder{message.AllocVLSBuilder(20)}
	a.Ptr.SetBytes(0, _SecurityDefinitionForSymbolRequestDefault)
	return a
}

func AllocSecurityDefinitionForSymbolRequestFrom(b []byte) SecurityDefinitionForSymbolRequestPointer {
	return SecurityDefinitionForSymbolRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size      = SecurityDefinitionForSymbolRequestSize  (20)
//     Type      = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     BaseSize  = SecurityDefinitionForSymbolRequestSize  (20)
//     RequestID = 0
//     Symbol    = ""
//     Exchange  = ""
// }
func (m SecurityDefinitionForSymbolRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SecurityDefinitionForSymbolRequestDefault)
}

// ToBuilder
func (m SecurityDefinitionForSymbolRequestPointer) ToBuilder() SecurityDefinitionForSymbolRequestPointerBuilder {
	return SecurityDefinitionForSymbolRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *SecurityDefinitionForSymbolRequestPointerBuilder) Finish() SecurityDefinitionForSymbolRequestPointer {
	return SecurityDefinitionForSymbolRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetSymbol The symbol to return a security definition for.
func (m *SecurityDefinitionForSymbolRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetExchange The optional exchange for the Symbol.
func (m *SecurityDefinitionForSymbolRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// Copy
func (m SecurityDefinitionForSymbolRequestPointer) Copy(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// Copy
func (m SecurityDefinitionForSymbolRequestPointerBuilder) Copy(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestPointer) CopyPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestPointerBuilder) CopyPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestPointer) CopyTo(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestPointerBuilder) CopyTo(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestPointer) CopyToPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestPointerBuilder) CopyToPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}
