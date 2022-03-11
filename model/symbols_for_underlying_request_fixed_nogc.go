package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SymbolsForUnderlyingRequestFixedPointer This is a message from the Client to the Server for requesting all of
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
type SymbolsForUnderlyingRequestFixedPointer struct {
	message.FixedPointer
}

// SymbolsForUnderlyingRequestFixedPointerBuilder This is a message from the Client to the Server for requesting all of
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
type SymbolsForUnderlyingRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocSymbolsForUnderlyingRequestFixed() SymbolsForUnderlyingRequestFixedPointerBuilder {
	a := SymbolsForUnderlyingRequestFixedPointerBuilder{message.AllocFixed(60)}
	a.Ptr.SetBytes(0, _SymbolsForUnderlyingRequestFixedDefault)
	return a
}

func AllocSymbolsForUnderlyingRequestFixedFrom(b []byte) SymbolsForUnderlyingRequestFixedPointer {
	return SymbolsForUnderlyingRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SymbolsForUnderlyingRequestFixedDefault)
}

// ToBuilder
func (m SymbolsForUnderlyingRequestFixedPointer) ToBuilder() SymbolsForUnderlyingRequestFixedPointerBuilder {
	return SymbolsForUnderlyingRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *SymbolsForUnderlyingRequestFixedPointerBuilder) Finish() SymbolsForUnderlyingRequestFixedPointer {
	return SymbolsForUnderlyingRequestFixedPointer{m.FixedPointer}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// UnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestFixedPointer) UnderlyingSymbol() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// UnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) UnderlyingSymbol() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// Exchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 56, 40)
}

// Exchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 56, 40)
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedPointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Ptr, 60, 56))
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Ptr, 60, 56))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetUnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetExchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 56, 40, value)
}

// SetSecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 60, 56, int32(value))
}

// Copy
func (m SymbolsForUnderlyingRequestFixedPointer) Copy(to SymbolsForUnderlyingRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) Copy(to SymbolsForUnderlyingRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m SymbolsForUnderlyingRequestFixedPointer) CopyPointer(to SymbolsForUnderlyingRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) CopyPointer(to SymbolsForUnderlyingRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m SymbolsForUnderlyingRequestFixedPointer) CopyTo(to SymbolsForUnderlyingRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) CopyTo(to SymbolsForUnderlyingRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m SymbolsForUnderlyingRequestFixedPointer) CopyToPointer(to SymbolsForUnderlyingRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) CopyToPointer(to SymbolsForUnderlyingRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}
