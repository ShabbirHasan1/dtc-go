package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size      = SecurityDefinitionForSymbolRequestSize  (20)
//     Type      = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     BaseSize  = SecurityDefinitionForSymbolRequestSize  (20)
//     RequestID = 0
//     Symbol    = ""
//     Exchange  = ""
// }
var _SecurityDefinitionForSymbolRequestDefault = []byte{20, 0, 250, 1, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SecurityDefinitionForSymbolRequestSize = 20

// SecurityDefinitionForSymbolRequest This is a message from the Client to the Server for requesting Security
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
type SecurityDefinitionForSymbolRequest struct {
	message.VLS
}

// SecurityDefinitionForSymbolRequestBuilder This is a message from the Client to the Server for requesting Security
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
type SecurityDefinitionForSymbolRequestBuilder struct {
	message.VLSBuilder
}

func NewSecurityDefinitionForSymbolRequestFrom(b []byte) SecurityDefinitionForSymbolRequest {
	return SecurityDefinitionForSymbolRequest{VLS: message.NewVLSFrom(b)}
}

func WrapSecurityDefinitionForSymbolRequest(b []byte) SecurityDefinitionForSymbolRequest {
	return SecurityDefinitionForSymbolRequest{VLS: message.WrapVLS(b)}
}

func NewSecurityDefinitionForSymbolRequest() SecurityDefinitionForSymbolRequestBuilder {
	a := SecurityDefinitionForSymbolRequestBuilder{message.NewVLSBuilder(20)}
	a.Unsafe().SetBytes(0, _SecurityDefinitionForSymbolRequestDefault)
	return a
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
func (m SecurityDefinitionForSymbolRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SecurityDefinitionForSymbolRequestDefault)
}

// ToBuilder
func (m SecurityDefinitionForSymbolRequest) ToBuilder() SecurityDefinitionForSymbolRequestBuilder {
	return SecurityDefinitionForSymbolRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m SecurityDefinitionForSymbolRequestBuilder) Finish() SecurityDefinitionForSymbolRequest {
	return SecurityDefinitionForSymbolRequest{m.VLSBuilder.Finish()}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetSymbol The symbol to return a security definition for.
func (m *SecurityDefinitionForSymbolRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetExchange The optional exchange for the Symbol.
func (m *SecurityDefinitionForSymbolRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// Copy
func (m SecurityDefinitionForSymbolRequest) Copy(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// Copy
func (m SecurityDefinitionForSymbolRequestBuilder) Copy(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequest) CopyPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestBuilder) CopyPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequest) CopyToPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestBuilder) CopyToPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequest) CopyTo(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestBuilder) CopyTo(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}
