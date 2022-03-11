package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size             = SymbolsForUnderlyingRequestFixedSize  (60)
//     Type             = SYMBOLS_FOR_UNDERLYING_REQUEST  (504)
//     RequestID        = 0
//     UnderlyingSymbol = ""
//     Exchange         = ""
//     SecurityType     = 0
// }
var _SymbolsForUnderlyingRequestFixedDefault = []byte{60, 0, 248, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SymbolsForUnderlyingRequestFixedSize = 60

// SymbolsForUnderlyingRequestFixed This is a message from the Client to the Server for requesting all of
// the symbols for a particular underlying symbol.
//
// For example, all of the futures contracts for a particular underlying
// futures symbol or all of the option symbols for a specific futures or
// stock symbol.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForUnderlyingRequestFixed struct {
	message.Fixed
}

// SymbolsForUnderlyingRequestFixedBuilder This is a message from the Client to the Server for requesting all of
// the symbols for a particular underlying symbol.
//
// For example, all of the futures contracts for a particular underlying
// futures symbol or all of the option symbols for a specific futures or
// stock symbol.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForUnderlyingRequestFixedBuilder struct {
	message.Fixed
}

func NewSymbolsForUnderlyingRequestFixedFrom(b []byte) SymbolsForUnderlyingRequestFixed {
	return SymbolsForUnderlyingRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapSymbolsForUnderlyingRequestFixed(b []byte) SymbolsForUnderlyingRequestFixed {
	return SymbolsForUnderlyingRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewSymbolsForUnderlyingRequestFixed() SymbolsForUnderlyingRequestFixedBuilder {
	a := SymbolsForUnderlyingRequestFixedBuilder{message.NewFixed(60)}
	a.Unsafe().SetBytes(0, _SymbolsForUnderlyingRequestFixedDefault)
	return a
}

// Clear
// {
//     Size             = SymbolsForUnderlyingRequestFixedSize  (60)
//     Type             = SYMBOLS_FOR_UNDERLYING_REQUEST  (504)
//     RequestID        = 0
//     UnderlyingSymbol = ""
//     Exchange         = ""
//     SecurityType     = 0
// }
func (m SymbolsForUnderlyingRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SymbolsForUnderlyingRequestFixedDefault)
}

// ToBuilder
func (m SymbolsForUnderlyingRequestFixed) ToBuilder() SymbolsForUnderlyingRequestFixedBuilder {
	return SymbolsForUnderlyingRequestFixedBuilder{m.Fixed}
}

// Finish
func (m SymbolsForUnderlyingRequestFixedBuilder) Finish() SymbolsForUnderlyingRequestFixed {
	return SymbolsForUnderlyingRequestFixed{m.Fixed}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// UnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestFixed) UnderlyingSymbol() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// UnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestFixedBuilder) UnderlyingSymbol() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// Exchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 56, 40)
}

// Exchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 56, 40)
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixed) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 60, 56))
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 60, 56))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetUnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestFixedBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetExchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 56, 40, value)
}

// SetSecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 60, 56, int32(value))
}

// Copy
func (m SymbolsForUnderlyingRequestFixed) Copy(to SymbolsForUnderlyingRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m SymbolsForUnderlyingRequestFixedBuilder) Copy(to SymbolsForUnderlyingRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m SymbolsForUnderlyingRequestFixed) CopyPointer(to SymbolsForUnderlyingRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m SymbolsForUnderlyingRequestFixedBuilder) CopyPointer(to SymbolsForUnderlyingRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m SymbolsForUnderlyingRequestFixed) CopyToPointer(to SymbolsForUnderlyingRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m SymbolsForUnderlyingRequestFixedBuilder) CopyToPointer(to SymbolsForUnderlyingRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m SymbolsForUnderlyingRequestFixed) CopyTo(to SymbolsForUnderlyingRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m SymbolsForUnderlyingRequestFixedBuilder) CopyTo(to SymbolsForUnderlyingRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}
