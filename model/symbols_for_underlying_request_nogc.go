package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SymbolsForUnderlyingRequestPointer This is a message from the Client to the Server for requesting all of
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
type SymbolsForUnderlyingRequestPointer struct {
	message.VLSPointer
}

// SymbolsForUnderlyingRequestPointerBuilder This is a message from the Client to the Server for requesting all of
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
type SymbolsForUnderlyingRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocSymbolsForUnderlyingRequest() SymbolsForUnderlyingRequestPointerBuilder {
	a := SymbolsForUnderlyingRequestPointerBuilder{message.AllocVLSBuilder(24)}
	a.Ptr.SetBytes(0, _SymbolsForUnderlyingRequestDefault)
	return a
}

func AllocSymbolsForUnderlyingRequestFrom(b []byte) SymbolsForUnderlyingRequestPointer {
	return SymbolsForUnderlyingRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size             = SymbolsForUnderlyingRequestSize  (24)
//     Type             = SYMBOLS_FOR_UNDERLYING_REQUEST  (504)
//     BaseSize         = SymbolsForUnderlyingRequestSize  (24)
//     RequestID        = 0
//     UnderlyingSymbol = ""
//     Exchange         = ""
//     SecurityType     = 0
// }
func (m SymbolsForUnderlyingRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SymbolsForUnderlyingRequestDefault)
}

// ToBuilder
func (m SymbolsForUnderlyingRequestPointer) ToBuilder() SymbolsForUnderlyingRequestPointerBuilder {
	return SymbolsForUnderlyingRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *SymbolsForUnderlyingRequestPointerBuilder) Finish() SymbolsForUnderlyingRequestPointer {
	return SymbolsForUnderlyingRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// UnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestPointer) UnderlyingSymbol() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// UnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestPointerBuilder) UnderlyingSymbol() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Exchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Exchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestPointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestPointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetUnderlyingSymbol The underlying symbol.
func (m *SymbolsForUnderlyingRequestPointerBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetExchange The exchange of the symbols to search for.
func (m *SymbolsForUnderlyingRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetSecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Ptr, 24, 20, int32(value))
}

// Copy
func (m SymbolsForUnderlyingRequestPointer) Copy(to SymbolsForUnderlyingRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m SymbolsForUnderlyingRequestPointerBuilder) Copy(to SymbolsForUnderlyingRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m SymbolsForUnderlyingRequestPointer) CopyPointer(to SymbolsForUnderlyingRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m SymbolsForUnderlyingRequestPointerBuilder) CopyPointer(to SymbolsForUnderlyingRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m SymbolsForUnderlyingRequestPointer) CopyTo(to SymbolsForUnderlyingRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m SymbolsForUnderlyingRequestPointerBuilder) CopyTo(to SymbolsForUnderlyingRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m SymbolsForUnderlyingRequestPointer) CopyToPointer(to SymbolsForUnderlyingRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m SymbolsForUnderlyingRequestPointerBuilder) CopyToPointer(to SymbolsForUnderlyingRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}
