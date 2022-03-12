package model

import (
	"github.com/moontrade/dtc-go/message"
)

const SymbolsForUnderlyingRequestSize = 24

const SymbolsForUnderlyingRequestFixedSize = 60

// {
//     Size             = SymbolsForUnderlyingRequestSize  (24)
//     Type             = SYMBOLS_FOR_UNDERLYING_REQUEST  (504)
//     BaseSize         = SymbolsForUnderlyingRequestSize  (24)
//     RequestID        = 0
//     UnderlyingSymbol = ""
//     Exchange         = ""
//     SecurityType     = SECURITY_TYPE_UNSET  (0)
// }
var _SymbolsForUnderlyingRequestDefault = []byte{24, 0, 248, 1, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size             = SymbolsForUnderlyingRequestFixedSize  (60)
//     Type             = SYMBOLS_FOR_UNDERLYING_REQUEST  (504)
//     RequestID        = 0
//     UnderlyingSymbol = ""
//     Exchange         = ""
//     SecurityType     = SECURITY_TYPE_UNSET  (0)
// }
var _SymbolsForUnderlyingRequestFixedDefault = []byte{60, 0, 248, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocSymbolsForUnderlyingRequest() SymbolsForUnderlyingRequestPointerBuilder {
	a := SymbolsForUnderlyingRequestPointerBuilder{message.AllocVLSBuilder(24)}
	a.Ptr.SetBytes(0, _SymbolsForUnderlyingRequestDefault)
	return a
}

func AllocSymbolsForUnderlyingRequestFrom(b []byte) SymbolsForUnderlyingRequestPointer {
	return SymbolsForUnderlyingRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size             = SymbolsForUnderlyingRequestSize  (24)
//     Type             = SYMBOLS_FOR_UNDERLYING_REQUEST  (504)
//     BaseSize         = SymbolsForUnderlyingRequestSize  (24)
//     RequestID        = 0
//     UnderlyingSymbol = ""
//     Exchange         = ""
//     SecurityType     = SECURITY_TYPE_UNSET  (0)
// }
func (m SymbolsForUnderlyingRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SymbolsForUnderlyingRequestDefault)
}

// Clear
// {
//     Size             = SymbolsForUnderlyingRequestFixedSize  (60)
//     Type             = SYMBOLS_FOR_UNDERLYING_REQUEST  (504)
//     RequestID        = 0
//     UnderlyingSymbol = ""
//     Exchange         = ""
//     SecurityType     = SECURITY_TYPE_UNSET  (0)
// }
func (m SymbolsForUnderlyingRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SymbolsForUnderlyingRequestFixedDefault)
}

// Clear
// {
//     Size             = SymbolsForUnderlyingRequestSize  (24)
//     Type             = SYMBOLS_FOR_UNDERLYING_REQUEST  (504)
//     BaseSize         = SymbolsForUnderlyingRequestSize  (24)
//     RequestID        = 0
//     UnderlyingSymbol = ""
//     Exchange         = ""
//     SecurityType     = SECURITY_TYPE_UNSET  (0)
// }
func (m SymbolsForUnderlyingRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SymbolsForUnderlyingRequestDefault)
}

// Clear
// {
//     Size             = SymbolsForUnderlyingRequestFixedSize  (60)
//     Type             = SYMBOLS_FOR_UNDERLYING_REQUEST  (504)
//     RequestID        = 0
//     UnderlyingSymbol = ""
//     Exchange         = ""
//     SecurityType     = SECURITY_TYPE_UNSET  (0)
// }
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SymbolsForUnderlyingRequestFixedDefault)
}

// ToBuilder
func (m SymbolsForUnderlyingRequest) ToBuilder() SymbolsForUnderlyingRequestBuilder {
	return SymbolsForUnderlyingRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m SymbolsForUnderlyingRequestFixed) ToBuilder() SymbolsForUnderlyingRequestFixedBuilder {
	return SymbolsForUnderlyingRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m SymbolsForUnderlyingRequestPointer) ToBuilder() SymbolsForUnderlyingRequestPointerBuilder {
	return SymbolsForUnderlyingRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m SymbolsForUnderlyingRequestFixedPointer) ToBuilder() SymbolsForUnderlyingRequestFixedPointerBuilder {
	return SymbolsForUnderlyingRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m SymbolsForUnderlyingRequestBuilder) Finish() SymbolsForUnderlyingRequest {
	return SymbolsForUnderlyingRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m SymbolsForUnderlyingRequestFixedBuilder) Finish() SymbolsForUnderlyingRequestFixed {
	return SymbolsForUnderlyingRequestFixed{m.Fixed}
}

// Finish
func (m *SymbolsForUnderlyingRequestPointerBuilder) Finish() SymbolsForUnderlyingRequestPointer {
	return SymbolsForUnderlyingRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *SymbolsForUnderlyingRequestFixedPointerBuilder) Finish() SymbolsForUnderlyingRequestFixedPointer {
	return SymbolsForUnderlyingRequestFixedPointer{m.FixedPointer}
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
func (m SymbolsForUnderlyingRequest) UnderlyingSymbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// UnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestBuilder) UnderlyingSymbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
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
func (m SymbolsForUnderlyingRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
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
func (m SymbolsForUnderlyingRequest) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestPointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestPointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 24, 20))
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
func (m SymbolsForUnderlyingRequestFixed) UnderlyingSymbol() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// UnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestFixedBuilder) UnderlyingSymbol() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
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
func (m SymbolsForUnderlyingRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 56, 40)
}

// Exchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 56, 40)
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
func (m SymbolsForUnderlyingRequestFixed) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 60, 56))
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 60, 56))
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
func (m SymbolsForUnderlyingRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetUnderlyingSymbol The underlying symbol.
func (m *SymbolsForUnderlyingRequestBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetUnderlyingSymbol The underlying symbol.
func (m *SymbolsForUnderlyingRequestPointerBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetExchange The exchange of the symbols to search for.
func (m *SymbolsForUnderlyingRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetExchange The exchange of the symbols to search for.
func (m *SymbolsForUnderlyingRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetSecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, int32(value))
}

// SetSecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Ptr, 24, 20, int32(value))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetUnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestFixedBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetUnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetExchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 56, 40, value)
}

// SetExchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 56, 40, value)
}

// SetSecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 60, 56, int32(value))
}

// SetSecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixedPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 60, 56, int32(value))
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
