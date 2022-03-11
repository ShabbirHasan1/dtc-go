package model

import (
	"github.com/moontrade/dtc-go/message"
)

const SecurityDefinitionForSymbolRequestSize = 20

const SecurityDefinitionForSymbolRequestFixedSize = 88

// {
//     Size      = SecurityDefinitionForSymbolRequestSize  (20)
//     Type      = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     BaseSize  = SecurityDefinitionForSymbolRequestSize  (20)
//     RequestID = 0
//     Symbol    = ""
//     Exchange  = ""
// }
var _SecurityDefinitionForSymbolRequestDefault = []byte{20, 0, 250, 1, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size      = SecurityDefinitionForSymbolRequestFixedSize  (88)
//     Type      = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     RequestID = 0
//     Symbol    = ""
//     Exchange  = ""
// }
var _SecurityDefinitionForSymbolRequestFixedDefault = []byte{88, 0, 250, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

// SecurityDefinitionForSymbolRequestFixed This is a message from the Client to the Server for requesting Security
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
type SecurityDefinitionForSymbolRequestFixed struct {
	message.Fixed
}

// SecurityDefinitionForSymbolRequestFixedBuilder This is a message from the Client to the Server for requesting Security
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
type SecurityDefinitionForSymbolRequestFixedBuilder struct {
	message.Fixed
}

// SecurityDefinitionForSymbolRequestPointer This is a message from the Client to the Server for requesting Security
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
type SecurityDefinitionForSymbolRequestPointer struct {
	message.VLSPointer
}

// SecurityDefinitionForSymbolRequestPointerBuilder This is a message from the Client to the Server for requesting Security
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
type SecurityDefinitionForSymbolRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

// SecurityDefinitionForSymbolRequestFixedPointer This is a message from the Client to the Server for requesting Security
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
type SecurityDefinitionForSymbolRequestFixedPointer struct {
	message.FixedPointer
}

// SecurityDefinitionForSymbolRequestFixedPointerBuilder This is a message from the Client to the Server for requesting Security
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
type SecurityDefinitionForSymbolRequestFixedPointerBuilder struct {
	message.FixedPointer
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

func NewSecurityDefinitionForSymbolRequestFixedFrom(b []byte) SecurityDefinitionForSymbolRequestFixed {
	return SecurityDefinitionForSymbolRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapSecurityDefinitionForSymbolRequestFixed(b []byte) SecurityDefinitionForSymbolRequestFixed {
	return SecurityDefinitionForSymbolRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewSecurityDefinitionForSymbolRequestFixed() SecurityDefinitionForSymbolRequestFixedBuilder {
	a := SecurityDefinitionForSymbolRequestFixedBuilder{message.NewFixed(88)}
	a.Unsafe().SetBytes(0, _SecurityDefinitionForSymbolRequestFixedDefault)
	return a
}

func AllocSecurityDefinitionForSymbolRequest() SecurityDefinitionForSymbolRequestPointerBuilder {
	a := SecurityDefinitionForSymbolRequestPointerBuilder{message.AllocVLSBuilder(20)}
	a.Ptr.SetBytes(0, _SecurityDefinitionForSymbolRequestDefault)
	return a
}

func AllocSecurityDefinitionForSymbolRequestFrom(b []byte) SecurityDefinitionForSymbolRequestPointer {
	return SecurityDefinitionForSymbolRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocSecurityDefinitionForSymbolRequestFixed() SecurityDefinitionForSymbolRequestFixedPointerBuilder {
	a := SecurityDefinitionForSymbolRequestFixedPointerBuilder{message.AllocFixed(88)}
	a.Ptr.SetBytes(0, _SecurityDefinitionForSymbolRequestFixedDefault)
	return a
}

func AllocSecurityDefinitionForSymbolRequestFixedFrom(b []byte) SecurityDefinitionForSymbolRequestFixedPointer {
	return SecurityDefinitionForSymbolRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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

// Clear
// {
//     Size      = SecurityDefinitionForSymbolRequestFixedSize  (88)
//     Type      = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     RequestID = 0
//     Symbol    = ""
//     Exchange  = ""
// }
func (m SecurityDefinitionForSymbolRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SecurityDefinitionForSymbolRequestFixedDefault)
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
func (m SecurityDefinitionForSymbolRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SecurityDefinitionForSymbolRequestDefault)
}

// Clear
// {
//     Size      = SecurityDefinitionForSymbolRequestFixedSize  (88)
//     Type      = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     RequestID = 0
//     Symbol    = ""
//     Exchange  = ""
// }
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SecurityDefinitionForSymbolRequestFixedDefault)
}

// ToBuilder
func (m SecurityDefinitionForSymbolRequest) ToBuilder() SecurityDefinitionForSymbolRequestBuilder {
	return SecurityDefinitionForSymbolRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m SecurityDefinitionForSymbolRequestFixed) ToBuilder() SecurityDefinitionForSymbolRequestFixedBuilder {
	return SecurityDefinitionForSymbolRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m SecurityDefinitionForSymbolRequestPointer) ToBuilder() SecurityDefinitionForSymbolRequestPointerBuilder {
	return SecurityDefinitionForSymbolRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m SecurityDefinitionForSymbolRequestFixedPointer) ToBuilder() SecurityDefinitionForSymbolRequestFixedPointerBuilder {
	return SecurityDefinitionForSymbolRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m SecurityDefinitionForSymbolRequestBuilder) Finish() SecurityDefinitionForSymbolRequest {
	return SecurityDefinitionForSymbolRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m SecurityDefinitionForSymbolRequestFixedBuilder) Finish() SecurityDefinitionForSymbolRequestFixed {
	return SecurityDefinitionForSymbolRequestFixed{m.Fixed}
}

// Finish
func (m *SecurityDefinitionForSymbolRequestPointerBuilder) Finish() SecurityDefinitionForSymbolRequestPointer {
	return SecurityDefinitionForSymbolRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *SecurityDefinitionForSymbolRequestFixedPointerBuilder) Finish() SecurityDefinitionForSymbolRequestFixedPointer {
	return SecurityDefinitionForSymbolRequestFixedPointer{m.FixedPointer}
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

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Symbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 88, 72)
}

// Exchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 88, 72)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetSymbol The symbol to return a security definition for.
func (m *SecurityDefinitionForSymbolRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetSymbol The symbol to return a security definition for.
func (m *SecurityDefinitionForSymbolRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetExchange The optional exchange for the Symbol.
func (m *SecurityDefinitionForSymbolRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetExchange The optional exchange for the Symbol.
func (m *SecurityDefinitionForSymbolRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetSymbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetSymbol The symbol to return a security definition for.
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 72, 8, value)
}

// SetExchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 88, 72, value)
}

// SetExchange The optional exchange for the Symbol.
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 88, 72, value)
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

// Copy
func (m SecurityDefinitionForSymbolRequestFixed) Copy(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// Copy
func (m SecurityDefinitionForSymbolRequestFixedBuilder) Copy(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestFixed) CopyPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestFixedBuilder) CopyPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestFixed) CopyToPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestFixedBuilder) CopyToPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestFixed) CopyTo(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestFixedBuilder) CopyTo(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// Copy
func (m SecurityDefinitionForSymbolRequestPointer) Copy(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// Copy
func (m SecurityDefinitionForSymbolRequestPointerBuilder) Copy(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestPointer) CopyPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestPointerBuilder) CopyPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestPointer) CopyTo(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestPointerBuilder) CopyTo(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestPointer) CopyToPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestPointerBuilder) CopyToPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// Copy
func (m SecurityDefinitionForSymbolRequestFixedPointer) Copy(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// Copy
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) Copy(to SecurityDefinitionForSymbolRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestFixedPointer) CopyPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyPointer
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) CopyPointer(to SecurityDefinitionForSymbolRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestFixedPointer) CopyTo(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyTo
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) CopyTo(to SecurityDefinitionForSymbolRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestFixedPointer) CopyToPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}

// CopyToPointer
func (m SecurityDefinitionForSymbolRequestFixedPointerBuilder) CopyToPointer(to SecurityDefinitionForSymbolRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
}
