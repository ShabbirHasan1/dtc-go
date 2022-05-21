// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-16 08:00:54.519198 +0800 WITA m=+0.055446543

package v8

import (
	"github.com/moontrade/dtc-go/message"
)

const SecurityDefinitionForSymbolRequestSize = 20

const SecurityDefinitionForSymbolRequestFixedSize = 88

//     Size       uint16  = SecurityDefinitionForSymbolRequestSize  (20)
//     Type       uint16  = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     BaseSize   uint16  = SecurityDefinitionForSymbolRequestSize  (20)
//     RequestID  int32   = 0
//     Symbol     string  = ""
//     Exchange   string  = ""
var _SecurityDefinitionForSymbolRequestDefault = []byte{20, 0, 250, 1, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size       uint16      = SecurityDefinitionForSymbolRequestFixedSize  (88)
//     Type       uint16      = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     RequestID  int32       = 0
//     Symbol     string[64]  = ""
//     Exchange   string[16]  = ""
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
	return SecurityDefinitionForSymbolRequest{VLS: message.NewVLS(b)}
}

func WrapSecurityDefinitionForSymbolRequest(b []byte) SecurityDefinitionForSymbolRequest {
	return SecurityDefinitionForSymbolRequest{VLS: message.WrapVLS(b)}
}

func NewSecurityDefinitionForSymbolRequest() SecurityDefinitionForSymbolRequestBuilder {
	return SecurityDefinitionForSymbolRequestBuilder{message.NewVLSBuilder(_SecurityDefinitionForSymbolRequestDefault)}
}

func NewSecurityDefinitionForSymbolRequestFixedFrom(b []byte) SecurityDefinitionForSymbolRequestFixed {
	return SecurityDefinitionForSymbolRequestFixed{Fixed: message.NewFixed(b)}
}

func WrapSecurityDefinitionForSymbolRequestFixed(b []byte) SecurityDefinitionForSymbolRequestFixed {
	return SecurityDefinitionForSymbolRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewSecurityDefinitionForSymbolRequestFixed() SecurityDefinitionForSymbolRequestFixedBuilder {
	return SecurityDefinitionForSymbolRequestFixedBuilder{message.NewFixed(_SecurityDefinitionForSymbolRequestFixedDefault)}
}

func AllocSecurityDefinitionForSymbolRequest() SecurityDefinitionForSymbolRequestPointerBuilder {
	return SecurityDefinitionForSymbolRequestPointerBuilder{message.AllocVLSBuilder(_SecurityDefinitionForSymbolRequestDefault)}
}

func AllocSecurityDefinitionForSymbolRequestFrom(b []byte) SecurityDefinitionForSymbolRequestPointer {
	return SecurityDefinitionForSymbolRequestPointer{VLSPointer: message.AllocVLS(b)}
}

func AllocSecurityDefinitionForSymbolRequestFixed() SecurityDefinitionForSymbolRequestFixedPointerBuilder {
	return SecurityDefinitionForSymbolRequestFixedPointerBuilder{message.AllocFixed(_SecurityDefinitionForSymbolRequestFixedDefault)}
}

func AllocSecurityDefinitionForSymbolRequestFixedFrom(b []byte) SecurityDefinitionForSymbolRequestFixedPointer {
	return SecurityDefinitionForSymbolRequestFixedPointer{FixedPointer: message.AllocFixed(b)}
}

// Clear
//     Size       uint16  = SecurityDefinitionForSymbolRequestSize  (20)
//     Type       uint16  = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     BaseSize   uint16  = SecurityDefinitionForSymbolRequestSize  (20)
//     RequestID  int32   = 0
//     Symbol     string  = ""
//     Exchange   string  = ""
func (m SecurityDefinitionForSymbolRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SecurityDefinitionForSymbolRequestDefault)
}

// Clear
//     Size       uint16      = SecurityDefinitionForSymbolRequestFixedSize  (88)
//     Type       uint16      = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     RequestID  int32       = 0
//     Symbol     string[64]  = ""
//     Exchange   string[16]  = ""
func (m SecurityDefinitionForSymbolRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SecurityDefinitionForSymbolRequestFixedDefault)
}

// Clear
//     Size       uint16  = SecurityDefinitionForSymbolRequestSize  (20)
//     Type       uint16  = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     BaseSize   uint16  = SecurityDefinitionForSymbolRequestSize  (20)
//     RequestID  int32   = 0
//     Symbol     string  = ""
//     Exchange   string  = ""
func (m SecurityDefinitionForSymbolRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SecurityDefinitionForSymbolRequestDefault)
}

// Clear
//     Size       uint16      = SecurityDefinitionForSymbolRequestFixedSize  (88)
//     Type       uint16      = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  (506)
//     RequestID  int32       = 0
//     Symbol     string[64]  = ""
//     Exchange   string[16]  = ""
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
