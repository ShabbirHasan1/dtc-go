package model

import (
	"github.com/moontrade/dtc-go/message"
)

const SymbolsForExchangeRequestSize = 28

const SymbolsForExchangeRequestFixedSize = 96

// {
//     Size          = SymbolsForExchangeRequestSize  (28)
//     Type          = SYMBOLS_FOR_EXCHANGE_REQUEST  (502)
//     BaseSize      = SymbolsForExchangeRequestSize  (28)
//     RequestID     = 0
//     Exchange      = ""
//     SecurityType  = SECURITY_TYPE_UNSET  (0)
//     RequestAction = 0
//     Symbol        = ""
// }
var _SymbolsForExchangeRequestDefault = []byte{28, 0, 246, 1, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size          = SymbolsForExchangeRequestFixedSize  (96)
//     Type          = SYMBOLS_FOR_EXCHANGE_REQUEST  (502)
//     RequestID     = 0
//     Exchange      = ""
//     SecurityType  = SECURITY_TYPE_UNSET  (0)
//     RequestAction = 0
//     Symbol        = ""
// }
var _SymbolsForExchangeRequestFixedDefault = []byte{96, 0, 246, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// SymbolsForExchangeRequest This is a message from the Client to the Server to request all of the
// Symbols for a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForExchangeRequest struct {
	message.VLS
}

// SymbolsForExchangeRequestBuilder This is a message from the Client to the Server to request all of the
// Symbols for a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForExchangeRequestBuilder struct {
	message.VLSBuilder
}

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

func NewSymbolsForExchangeRequestFrom(b []byte) SymbolsForExchangeRequest {
	return SymbolsForExchangeRequest{VLS: message.NewVLSFrom(b)}
}

func WrapSymbolsForExchangeRequest(b []byte) SymbolsForExchangeRequest {
	return SymbolsForExchangeRequest{VLS: message.WrapVLS(b)}
}

func NewSymbolsForExchangeRequest() SymbolsForExchangeRequestBuilder {
	a := SymbolsForExchangeRequestBuilder{message.NewVLSBuilder(28)}
	a.Unsafe().SetBytes(0, _SymbolsForExchangeRequestDefault)
	return a
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

func AllocSymbolsForExchangeRequest() SymbolsForExchangeRequestPointerBuilder {
	a := SymbolsForExchangeRequestPointerBuilder{message.AllocVLSBuilder(28)}
	a.Ptr.SetBytes(0, _SymbolsForExchangeRequestDefault)
	return a
}

func AllocSymbolsForExchangeRequestFrom(b []byte) SymbolsForExchangeRequestPointer {
	return SymbolsForExchangeRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size          = SymbolsForExchangeRequestSize  (28)
//     Type          = SYMBOLS_FOR_EXCHANGE_REQUEST  (502)
//     BaseSize      = SymbolsForExchangeRequestSize  (28)
//     RequestID     = 0
//     Exchange      = ""
//     SecurityType  = SECURITY_TYPE_UNSET  (0)
//     RequestAction = 0
//     Symbol        = ""
// }
func (m SymbolsForExchangeRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SymbolsForExchangeRequestDefault)
}

// Clear
// {
//     Size          = SymbolsForExchangeRequestFixedSize  (96)
//     Type          = SYMBOLS_FOR_EXCHANGE_REQUEST  (502)
//     RequestID     = 0
//     Exchange      = ""
//     SecurityType  = SECURITY_TYPE_UNSET  (0)
//     RequestAction = 0
//     Symbol        = ""
// }
func (m SymbolsForExchangeRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SymbolsForExchangeRequestFixedDefault)
}

// Clear
// {
//     Size          = SymbolsForExchangeRequestSize  (28)
//     Type          = SYMBOLS_FOR_EXCHANGE_REQUEST  (502)
//     BaseSize      = SymbolsForExchangeRequestSize  (28)
//     RequestID     = 0
//     Exchange      = ""
//     SecurityType  = SECURITY_TYPE_UNSET  (0)
//     RequestAction = 0
//     Symbol        = ""
// }
func (m SymbolsForExchangeRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SymbolsForExchangeRequestDefault)
}

// Clear
// {
//     Size          = SymbolsForExchangeRequestFixedSize  (96)
//     Type          = SYMBOLS_FOR_EXCHANGE_REQUEST  (502)
//     RequestID     = 0
//     Exchange      = ""
//     SecurityType  = SECURITY_TYPE_UNSET  (0)
//     RequestAction = 0
//     Symbol        = ""
// }
func (m SymbolsForExchangeRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SymbolsForExchangeRequestFixedDefault)
}

// ToBuilder
func (m SymbolsForExchangeRequest) ToBuilder() SymbolsForExchangeRequestBuilder {
	return SymbolsForExchangeRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m SymbolsForExchangeRequestFixed) ToBuilder() SymbolsForExchangeRequestFixedBuilder {
	return SymbolsForExchangeRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m SymbolsForExchangeRequestPointer) ToBuilder() SymbolsForExchangeRequestPointerBuilder {
	return SymbolsForExchangeRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m SymbolsForExchangeRequestFixedPointer) ToBuilder() SymbolsForExchangeRequestFixedPointerBuilder {
	return SymbolsForExchangeRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m SymbolsForExchangeRequestBuilder) Finish() SymbolsForExchangeRequest {
	return SymbolsForExchangeRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m SymbolsForExchangeRequestFixedBuilder) Finish() SymbolsForExchangeRequestFixed {
	return SymbolsForExchangeRequestFixed{m.Fixed}
}

// Finish
func (m *SymbolsForExchangeRequestPointerBuilder) Finish() SymbolsForExchangeRequestPointer {
	return SymbolsForExchangeRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *SymbolsForExchangeRequestFixedPointerBuilder) Finish() SymbolsForExchangeRequestFixedPointer {
	return SymbolsForExchangeRequestFixedPointer{m.FixedPointer}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
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
func (m SymbolsForExchangeRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Exchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
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
func (m SymbolsForExchangeRequest) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 20, 16))
}

// SecurityType The optional Security Type.
func (m SymbolsForExchangeRequestBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 20, 16))
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
func (m SymbolsForExchangeRequest) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// RequestAction
func (m SymbolsForExchangeRequestBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Unsafe(), 24, 20))
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
func (m SymbolsForExchangeRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Symbol
func (m SymbolsForExchangeRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Symbol
func (m SymbolsForExchangeRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 28, 24)
}

// Symbol
func (m SymbolsForExchangeRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 28, 24)
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
func (m SymbolsForExchangeRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 24, 8)
}

// Exchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 24, 8)
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
func (m SymbolsForExchangeRequestFixed) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 28, 24))
}

// SecurityType The optional Security Type.
func (m SymbolsForExchangeRequestFixedBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 28, 24))
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
func (m SymbolsForExchangeRequestFixed) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Unsafe(), 32, 28))
}

// RequestAction
func (m SymbolsForExchangeRequestFixedBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Unsafe(), 32, 28))
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
func (m SymbolsForExchangeRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 96, 32)
}

// Symbol
func (m SymbolsForExchangeRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 96, 32)
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
func (m SymbolsForExchangeRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetExchange The Exchange to return the available symbols listed on that Exchange.
func (m *SymbolsForExchangeRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetExchange The Exchange to return the available symbols listed on that Exchange.
func (m *SymbolsForExchangeRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetSecurityType The optional Security Type.
func (m SymbolsForExchangeRequestBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 20, 16, int32(value))
}

// SetSecurityType The optional Security Type.
func (m SymbolsForExchangeRequestPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Ptr, 20, 16, int32(value))
}

// SetRequestAction
func (m SymbolsForExchangeRequestBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, int32(value))
}

// SetRequestAction
func (m SymbolsForExchangeRequestPointerBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Ptr, 24, 20, int32(value))
}

// SetSymbol
func (m *SymbolsForExchangeRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 28, 24, value)
}

// SetSymbol
func (m *SymbolsForExchangeRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 28, 24, value)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForExchangeRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetExchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 24, 8, value)
}

// SetExchange The Exchange to return the available symbols listed on that Exchange.
func (m SymbolsForExchangeRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 24, 8, value)
}

// SetSecurityType The optional Security Type.
func (m SymbolsForExchangeRequestFixedBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 28, 24, int32(value))
}

// SetSecurityType The optional Security Type.
func (m SymbolsForExchangeRequestFixedPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 28, 24, int32(value))
}

// SetRequestAction
func (m SymbolsForExchangeRequestFixedBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32Fixed(m.Unsafe(), 32, 28, int32(value))
}

// SetRequestAction
func (m SymbolsForExchangeRequestFixedPointerBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32Fixed(m.Ptr, 32, 28, int32(value))
}

// SetSymbol
func (m SymbolsForExchangeRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 96, 32, value)
}

// SetSymbol
func (m SymbolsForExchangeRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 96, 32, value)
}

// Copy
func (m SymbolsForExchangeRequest) Copy(to SymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// Copy
func (m SymbolsForExchangeRequestBuilder) Copy(to SymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m SymbolsForExchangeRequest) CopyPointer(to SymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyPointer
func (m SymbolsForExchangeRequestBuilder) CopyPointer(to SymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyToPointer
func (m SymbolsForExchangeRequest) CopyToPointer(to SymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyToPointer
func (m SymbolsForExchangeRequestBuilder) CopyToPointer(to SymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyTo
func (m SymbolsForExchangeRequest) CopyTo(to SymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
}

// CopyTo
func (m SymbolsForExchangeRequestBuilder) CopyTo(to SymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetRequestAction(m.RequestAction())
	to.SetSymbol(m.Symbol())
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
