package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SymbolsForExchangeRequestFixedPointer This is a message from the Client to the Server to request all of the
// Symbols for a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForExchangeRequestFixedPointer struct {
	message.FixedPointer
}

// SymbolsForExchangeRequestFixedPointerBuilder This is a message from the Client to the Server to request all of the
// Symbols for a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForExchangeRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocSymbolsForExchangeRequestFixed() SymbolsForExchangeRequestFixedPointerBuilder {
	a := SymbolsForExchangeRequestFixedPointerBuilder{message.AllocFixed(96)}
	a.Ptr.SetBytes(0, _SymbolsForExchangeRequestFixedDefault)
	return a
}

func AllocSymbolsForExchangeRequestFixedFrom(b []byte) SymbolsForExchangeRequestFixedPointer {
	return SymbolsForExchangeRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = SymbolsForExchangeRequestFixedSize  (96)
//     Type          = SYMBOLS_FOR_EXCHANGE_REQUEST  (502)
//     RequestID     = 0
//     Exchange      = ""
//     SecurityType  = 0
//     RequestAction = 0
//     Symbol        = ""
// }
func (m SymbolsForExchangeRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SymbolsForExchangeRequestFixedDefault)
}

// ToBuilder
func (m SymbolsForExchangeRequestFixedPointer) ToBuilder() SymbolsForExchangeRequestFixedPointerBuilder {
	return SymbolsForExchangeRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *SymbolsForExchangeRequestFixedPointerBuilder) Finish() SymbolsForExchangeRequestFixedPointer {
	return SymbolsForExchangeRequestFixedPointer{m.FixedPointer}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Exchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 24, 8)
}

// Exchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 24, 8)
}

// SecurityType The optional Security Type.
func (m SymbolsForExchangeRequestFixedPointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Ptr, 28, 24))
}

// SecurityType The optional Security Type.
func (m SymbolsForExchangeRequestFixedPointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Ptr, 28, 24))
}

// RequestAction
func (m SymbolsForExchangeRequestFixedPointer) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Ptr, 32, 28))
}

// RequestAction
func (m SymbolsForExchangeRequestFixedPointerBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Ptr, 32, 28))
}

// Symbol
func (m SymbolsForExchangeRequestFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 96, 32)
}

// Symbol
func (m SymbolsForExchangeRequestFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 96, 32)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetExchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 24, 8, value)
}

// SetSecurityType The optional Security Type.
func (m SymbolsForExchangeRequestFixedPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 28, 24, int32(value))
}

// SetRequestAction
func (m SymbolsForExchangeRequestFixedPointerBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32Fixed(m.Ptr, 32, 28, int32(value))
}

// SetSymbol
func (m SymbolsForExchangeRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 96, 32, value)
}

// Copy
func (m SymbolsForExchangeRequestFixedPointer) Copy(to SymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// Copy
func (m SymbolsForExchangeRequestFixedPointerBuilder) Copy(to SymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m SymbolsForExchangeRequestFixedPointer) CopyPointer(to SymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m SymbolsForExchangeRequestFixedPointerBuilder) CopyPointer(to SymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyTo
func (m SymbolsForExchangeRequestFixedPointer) CopyTo(to SymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyTo
func (m SymbolsForExchangeRequestFixedPointerBuilder) CopyTo(to SymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyToPointer
func (m SymbolsForExchangeRequestFixedPointer) CopyToPointer(to SymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyToPointer
func (m SymbolsForExchangeRequestFixedPointerBuilder) CopyToPointer(to SymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}
