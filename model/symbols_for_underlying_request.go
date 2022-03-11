package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size             = SymbolsForUnderlyingRequestSize  (24)
//     Type             = SYMBOLS_FOR_UNDERLYING_REQUEST  (504)
//     BaseSize         = SymbolsForUnderlyingRequestSize  (24)
//     RequestID        = 0
//     UnderlyingSymbol = ""
//     Exchange         = ""
//     SecurityType     = 0
// }
var _SymbolsForUnderlyingRequestDefault = []byte{24, 0, 248, 1, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SymbolsForUnderlyingRequestSize = 24

// SymbolsForUnderlyingRequest This is a message from the Client to the Server for requesting all of
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
type SymbolsForUnderlyingRequest struct {
	message.VLS
}

// SymbolsForUnderlyingRequestBuilder This is a message from the Client to the Server for requesting all of
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
type SymbolsForUnderlyingRequestBuilder struct {
	message.VLSBuilder
}

func NewSymbolsForUnderlyingRequestFrom(b []byte) SymbolsForUnderlyingRequest {
	return SymbolsForUnderlyingRequest{VLS: message.NewVLSFrom(b)}
}

func WrapSymbolsForUnderlyingRequest(b []byte) SymbolsForUnderlyingRequest {
	return SymbolsForUnderlyingRequest{VLS: message.WrapVLS(b)}
}

func NewSymbolsForUnderlyingRequest() SymbolsForUnderlyingRequestBuilder {
	a := SymbolsForUnderlyingRequestBuilder{message.NewVLSBuilder(24)}
	a.Unsafe().SetBytes(0, _SymbolsForUnderlyingRequestDefault)
	return a
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
func (m SymbolsForUnderlyingRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SymbolsForUnderlyingRequestDefault)
}

// ToBuilder
func (m SymbolsForUnderlyingRequest) ToBuilder() SymbolsForUnderlyingRequestBuilder {
	return SymbolsForUnderlyingRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m SymbolsForUnderlyingRequestBuilder) Finish() SymbolsForUnderlyingRequest {
	return SymbolsForUnderlyingRequest{m.VLSBuilder.Finish()}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// UnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequest) UnderlyingSymbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// UnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestBuilder) UnderlyingSymbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Exchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequest) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetUnderlyingSymbol The underlying symbol.
func (m *SymbolsForUnderlyingRequestBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetExchange The exchange of the symbols to search for.
func (m *SymbolsForUnderlyingRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetSecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, int32(value))
}

// Copy
func (m SymbolsForUnderlyingRequest) Copy(to SymbolsForUnderlyingRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m SymbolsForUnderlyingRequestBuilder) Copy(to SymbolsForUnderlyingRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m SymbolsForUnderlyingRequest) CopyPointer(to SymbolsForUnderlyingRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m SymbolsForUnderlyingRequestBuilder) CopyPointer(to SymbolsForUnderlyingRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m SymbolsForUnderlyingRequest) CopyToPointer(to SymbolsForUnderlyingRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m SymbolsForUnderlyingRequestBuilder) CopyToPointer(to SymbolsForUnderlyingRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m SymbolsForUnderlyingRequest) CopyTo(to SymbolsForUnderlyingRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m SymbolsForUnderlyingRequestBuilder) CopyTo(to SymbolsForUnderlyingRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}
