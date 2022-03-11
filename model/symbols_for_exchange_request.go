package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = SymbolsForExchangeRequestSize  (28)
//     Type          = SYMBOLS_FOR_EXCHANGE_REQUEST  (502)
//     BaseSize      = SymbolsForExchangeRequestSize  (28)
//     RequestID     = 0
//     Exchange      = ""
//     SecurityType  = 0
//     RequestAction = 0
//     Symbol        = ""
// }
var _SymbolsForExchangeRequestDefault = []byte{28, 0, 246, 1, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SymbolsForExchangeRequestSize = 28

// SymbolsForExchangeRequest This is a message from the Client to the Server to request all of the
// Symbols for a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForExchangeRequest struct {
	message.VLS
}

// SymbolsForExchangeRequestBuilder This is a message from the Client to the Server to request all of the
// Symbols for a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForExchangeRequestBuilder struct {
	message.VLSBuilder
}

func NewSymbolsForExchangeRequestFrom(b []byte) SymbolsForExchangeRequest {
	return SymbolsForExchangeRequest{VLS: message.NewVLSFrom(b)}
}

func WrapSymbolsForExchangeRequest(b []byte) SymbolsForExchangeRequest {
	return SymbolsForExchangeRequest{VLS: message.WrapVLS(b)}
}

func NewSymbolsForExchangeRequest() SymbolsForExchangeRequestBuilder {
	a := SymbolsForExchangeRequestBuilder{message.NewVLSBuilder(28)}
	a.Unsafe().SetBytes(0, _SymbolsForExchangeRequestDefault)
	return a
}

// Clear
// {
//     Size          = SymbolsForExchangeRequestSize  (28)
//     Type          = SYMBOLS_FOR_EXCHANGE_REQUEST  (502)
//     BaseSize      = SymbolsForExchangeRequestSize  (28)
//     RequestID     = 0
//     Exchange      = ""
//     SecurityType  = 0
//     RequestAction = 0
//     Symbol        = ""
// }
func (m SymbolsForExchangeRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SymbolsForExchangeRequestDefault)
}

// ToBuilder
func (m SymbolsForExchangeRequest) ToBuilder() SymbolsForExchangeRequestBuilder {
	return SymbolsForExchangeRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m SymbolsForExchangeRequestBuilder) Finish() SymbolsForExchangeRequest {
	return SymbolsForExchangeRequest{m.VLSBuilder.Finish()}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// Exchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Exchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SecurityType The optional Security Type.
func (m SymbolsForExchangeRequest) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 20, 16))
}

// SecurityType The optional Security Type.
func (m SymbolsForExchangeRequestBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 20, 16))
}

// RequestAction
func (m SymbolsForExchangeRequest) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// RequestAction
func (m SymbolsForExchangeRequestBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// Symbol
func (m SymbolsForExchangeRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Symbol
func (m SymbolsForExchangeRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetExchange The Exchange to return the available symbols listed on that Exchange.
func (m *SymbolsForExchangeRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetSecurityType The optional Security Type.
func (m SymbolsForExchangeRequestBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 20, 16, int32(value))
}

// SetRequestAction
func (m SymbolsForExchangeRequestBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, int32(value))
}

// SetSymbol
func (m *SymbolsForExchangeRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 28, 24, value)
}

// Copy
func (m SymbolsForExchangeRequest) Copy(to SymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// Copy
func (m SymbolsForExchangeRequestBuilder) Copy(to SymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m SymbolsForExchangeRequest) CopyPointer(to SymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m SymbolsForExchangeRequestBuilder) CopyPointer(to SymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyToPointer
func (m SymbolsForExchangeRequest) CopyToPointer(to SymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyToPointer
func (m SymbolsForExchangeRequestBuilder) CopyToPointer(to SymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyTo
func (m SymbolsForExchangeRequest) CopyTo(to SymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyTo
func (m SymbolsForExchangeRequestBuilder) CopyTo(to SymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}
