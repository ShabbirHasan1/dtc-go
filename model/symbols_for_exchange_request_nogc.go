package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SymbolsForExchangeRequestPointer This is a message from the Client to the Server to request all of the
// Symbols for a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForExchangeRequestPointer struct {
	message.VLSPointer
}

// SymbolsForExchangeRequestPointerBuilder This is a message from the Client to the Server to request all of the
// Symbols for a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForExchangeRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocSymbolsForExchangeRequest() SymbolsForExchangeRequestPointerBuilder {
	a := SymbolsForExchangeRequestPointerBuilder{message.AllocVLSBuilder(28)}
	a.Ptr.SetBytes(0, _SymbolsForExchangeRequestDefault)
	return a
}

func AllocSymbolsForExchangeRequestFrom(b []byte) SymbolsForExchangeRequestPointer {
	return SymbolsForExchangeRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m SymbolsForExchangeRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SymbolsForExchangeRequestDefault)
}

// ToBuilder
func (m SymbolsForExchangeRequestPointer) ToBuilder() SymbolsForExchangeRequestPointerBuilder {
	return SymbolsForExchangeRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *SymbolsForExchangeRequestPointerBuilder) Finish() SymbolsForExchangeRequestPointer {
	return SymbolsForExchangeRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// Exchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Exchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SecurityType The optional Security Type.
func (m SymbolsForExchangeRequestPointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 20, 16))
}

// SecurityType The optional Security Type.
func (m SymbolsForExchangeRequestPointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 20, 16))
}

// RequestAction
func (m SymbolsForExchangeRequestPointer) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// RequestAction
func (m SymbolsForExchangeRequestPointerBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// Symbol
func (m SymbolsForExchangeRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 28, 24)
}

// Symbol
func (m SymbolsForExchangeRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 28, 24)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetExchange The Exchange to return the available symbols listed on that Exchange.
func (m *SymbolsForExchangeRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetSecurityType The optional Security Type.
func (m SymbolsForExchangeRequestPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Ptr, 20, 16, int32(value))
}

// SetRequestAction
func (m SymbolsForExchangeRequestPointerBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Ptr, 24, 20, int32(value))
}

// SetSymbol
func (m *SymbolsForExchangeRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 28, 24, value)
}

// Copy
func (m SymbolsForExchangeRequestPointer) Copy(to SymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// Copy
func (m SymbolsForExchangeRequestPointerBuilder) Copy(to SymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m SymbolsForExchangeRequestPointer) CopyPointer(to SymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m SymbolsForExchangeRequestPointerBuilder) CopyPointer(to SymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyTo
func (m SymbolsForExchangeRequestPointer) CopyTo(to SymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyTo
func (m SymbolsForExchangeRequestPointerBuilder) CopyTo(to SymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyToPointer
func (m SymbolsForExchangeRequestPointer) CopyToPointer(to SymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyToPointer
func (m SymbolsForExchangeRequestPointerBuilder) CopyToPointer(to SymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}
