package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = SymbolsForExchangeRequestFixedSize  (96)
//     Type          = SYMBOLS_FOR_EXCHANGE_REQUEST  (502)
//     RequestID     = 0
//     Exchange      = ""
//     SecurityType  = 0
//     RequestAction = 0
//     Symbol        = ""
// }
var _SymbolsForExchangeRequestFixedDefault = []byte{96, 0, 246, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SymbolsForExchangeRequestFixedSize = 96

// SymbolsForExchangeRequestFixed This is a message from the Client to the Server to request all of the
// Symbols for a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForExchangeRequestFixed struct {
	message.Fixed
}

// SymbolsForExchangeRequestFixedBuilder This is a message from the Client to the Server to request all of the
// Symbols for a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForExchangeRequestFixedBuilder struct {
	message.Fixed
}

func NewSymbolsForExchangeRequestFixedFrom(b []byte) SymbolsForExchangeRequestFixed {
	return SymbolsForExchangeRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapSymbolsForExchangeRequestFixed(b []byte) SymbolsForExchangeRequestFixed {
	return SymbolsForExchangeRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewSymbolsForExchangeRequestFixed() SymbolsForExchangeRequestFixedBuilder {
	a := SymbolsForExchangeRequestFixedBuilder{message.NewFixed(96)}
	a.Unsafe().SetBytes(0, _SymbolsForExchangeRequestFixedDefault)
	return a
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
func (m SymbolsForExchangeRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SymbolsForExchangeRequestFixedDefault)
}

// ToBuilder
func (m SymbolsForExchangeRequestFixed) ToBuilder() SymbolsForExchangeRequestFixedBuilder {
	return SymbolsForExchangeRequestFixedBuilder{m.Fixed}
}

// Finish
func (m SymbolsForExchangeRequestFixedBuilder) Finish() SymbolsForExchangeRequestFixed {
	return SymbolsForExchangeRequestFixed{m.Fixed}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// Exchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 24, 8)
}

// Exchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 24, 8)
}

// SecurityType The optional Security Type.
func (m SymbolsForExchangeRequestFixed) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 28, 24))
}

// SecurityType The optional Security Type.
func (m SymbolsForExchangeRequestFixedBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 28, 24))
}

// RequestAction
func (m SymbolsForExchangeRequestFixed) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Unsafe(), 32, 28))
}

// RequestAction
func (m SymbolsForExchangeRequestFixedBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Unsafe(), 32, 28))
}

// Symbol
func (m SymbolsForExchangeRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 96, 32)
}

// Symbol
func (m SymbolsForExchangeRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 96, 32)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetExchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 24, 8, value)
}

// SetSecurityType The optional Security Type.
func (m SymbolsForExchangeRequestFixedBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 28, 24, int32(value))
}

// SetRequestAction
func (m SymbolsForExchangeRequestFixedBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32Fixed(m.Unsafe(), 32, 28, int32(value))
}

// SetSymbol
func (m SymbolsForExchangeRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 96, 32, value)
}

// Copy
func (m SymbolsForExchangeRequestFixed) Copy(to SymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// Copy
func (m SymbolsForExchangeRequestFixedBuilder) Copy(to SymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m SymbolsForExchangeRequestFixed) CopyPointer(to SymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m SymbolsForExchangeRequestFixedBuilder) CopyPointer(to SymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyToPointer
func (m SymbolsForExchangeRequestFixed) CopyToPointer(to SymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyToPointer
func (m SymbolsForExchangeRequestFixedBuilder) CopyToPointer(to SymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyTo
func (m SymbolsForExchangeRequestFixed) CopyTo(to SymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyTo
func (m SymbolsForExchangeRequestFixedBuilder) CopyTo(to SymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}
