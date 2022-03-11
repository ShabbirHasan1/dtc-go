package model

import (
	"github.com/moontrade/dtc-go/message"
)

const UnderlyingSymbolsForExchangeRequestSize = 20

const UnderlyingSymbolsForExchangeRequestFixedSize = 28

// {
//     Size         = UnderlyingSymbolsForExchangeRequestSize  (20)
//     Type         = UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST  (503)
//     BaseSize     = UnderlyingSymbolsForExchangeRequestSize  (20)
//     RequestID    = 0
//     Exchange     = ""
//     SecurityType = SECURITY_TYPE_UNSET  (0)
// }
var _UnderlyingSymbolsForExchangeRequestDefault = []byte{20, 0, 247, 1, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size         = UnderlyingSymbolsForExchangeRequestFixedSize  (28)
//     Type         = UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST  (503)
//     RequestID    = 0
//     Exchange     = ""
//     SecurityType = SECURITY_TYPE_UNSET  (0)
// }
var _UnderlyingSymbolsForExchangeRequestFixedDefault = []byte{28, 0, 247, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// UnderlyingSymbolsForExchangeRequest This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequest struct {
	message.VLS
}

// UnderlyingSymbolsForExchangeRequestBuilder This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestBuilder struct {
	message.VLSBuilder
}

// UnderlyingSymbolsForExchangeRequestFixed This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestFixed struct {
	message.Fixed
}

// UnderlyingSymbolsForExchangeRequestFixedBuilder This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestFixedBuilder struct {
	message.Fixed
}

// UnderlyingSymbolsForExchangeRequestPointer This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestPointer struct {
	message.VLSPointer
}

// UnderlyingSymbolsForExchangeRequestPointerBuilder This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

// UnderlyingSymbolsForExchangeRequestFixedPointer This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestFixedPointer struct {
	message.FixedPointer
}

// UnderlyingSymbolsForExchangeRequestFixedPointerBuilder This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type UnderlyingSymbolsForExchangeRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func NewUnderlyingSymbolsForExchangeRequestFrom(b []byte) UnderlyingSymbolsForExchangeRequest {
	return UnderlyingSymbolsForExchangeRequest{VLS: message.NewVLSFrom(b)}
}

func WrapUnderlyingSymbolsForExchangeRequest(b []byte) UnderlyingSymbolsForExchangeRequest {
	return UnderlyingSymbolsForExchangeRequest{VLS: message.WrapVLS(b)}
}

func NewUnderlyingSymbolsForExchangeRequest() UnderlyingSymbolsForExchangeRequestBuilder {
	a := UnderlyingSymbolsForExchangeRequestBuilder{message.NewVLSBuilder(20)}
	a.Unsafe().SetBytes(0, _UnderlyingSymbolsForExchangeRequestDefault)
	return a
}

func NewUnderlyingSymbolsForExchangeRequestFixedFrom(b []byte) UnderlyingSymbolsForExchangeRequestFixed {
	return UnderlyingSymbolsForExchangeRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapUnderlyingSymbolsForExchangeRequestFixed(b []byte) UnderlyingSymbolsForExchangeRequestFixed {
	return UnderlyingSymbolsForExchangeRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewUnderlyingSymbolsForExchangeRequestFixed() UnderlyingSymbolsForExchangeRequestFixedBuilder {
	a := UnderlyingSymbolsForExchangeRequestFixedBuilder{message.NewFixed(28)}
	a.Unsafe().SetBytes(0, _UnderlyingSymbolsForExchangeRequestFixedDefault)
	return a
}

func AllocUnderlyingSymbolsForExchangeRequest() UnderlyingSymbolsForExchangeRequestPointerBuilder {
	a := UnderlyingSymbolsForExchangeRequestPointerBuilder{message.AllocVLSBuilder(20)}
	a.Ptr.SetBytes(0, _UnderlyingSymbolsForExchangeRequestDefault)
	return a
}

func AllocUnderlyingSymbolsForExchangeRequestFrom(b []byte) UnderlyingSymbolsForExchangeRequestPointer {
	return UnderlyingSymbolsForExchangeRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocUnderlyingSymbolsForExchangeRequestFixed() UnderlyingSymbolsForExchangeRequestFixedPointerBuilder {
	a := UnderlyingSymbolsForExchangeRequestFixedPointerBuilder{message.AllocFixed(28)}
	a.Ptr.SetBytes(0, _UnderlyingSymbolsForExchangeRequestFixedDefault)
	return a
}

func AllocUnderlyingSymbolsForExchangeRequestFixedFrom(b []byte) UnderlyingSymbolsForExchangeRequestFixedPointer {
	return UnderlyingSymbolsForExchangeRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size         = UnderlyingSymbolsForExchangeRequestSize  (20)
//     Type         = UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST  (503)
//     BaseSize     = UnderlyingSymbolsForExchangeRequestSize  (20)
//     RequestID    = 0
//     Exchange     = ""
//     SecurityType = SECURITY_TYPE_UNSET  (0)
// }
func (m UnderlyingSymbolsForExchangeRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _UnderlyingSymbolsForExchangeRequestDefault)
}

// Clear
// {
//     Size         = UnderlyingSymbolsForExchangeRequestFixedSize  (28)
//     Type         = UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST  (503)
//     RequestID    = 0
//     Exchange     = ""
//     SecurityType = SECURITY_TYPE_UNSET  (0)
// }
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _UnderlyingSymbolsForExchangeRequestFixedDefault)
}

// Clear
// {
//     Size         = UnderlyingSymbolsForExchangeRequestSize  (20)
//     Type         = UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST  (503)
//     BaseSize     = UnderlyingSymbolsForExchangeRequestSize  (20)
//     RequestID    = 0
//     Exchange     = ""
//     SecurityType = SECURITY_TYPE_UNSET  (0)
// }
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _UnderlyingSymbolsForExchangeRequestDefault)
}

// Clear
// {
//     Size         = UnderlyingSymbolsForExchangeRequestFixedSize  (28)
//     Type         = UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST  (503)
//     RequestID    = 0
//     Exchange     = ""
//     SecurityType = SECURITY_TYPE_UNSET  (0)
// }
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _UnderlyingSymbolsForExchangeRequestFixedDefault)
}

// ToBuilder
func (m UnderlyingSymbolsForExchangeRequest) ToBuilder() UnderlyingSymbolsForExchangeRequestBuilder {
	return UnderlyingSymbolsForExchangeRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m UnderlyingSymbolsForExchangeRequestFixed) ToBuilder() UnderlyingSymbolsForExchangeRequestFixedBuilder {
	return UnderlyingSymbolsForExchangeRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m UnderlyingSymbolsForExchangeRequestPointer) ToBuilder() UnderlyingSymbolsForExchangeRequestPointerBuilder {
	return UnderlyingSymbolsForExchangeRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) ToBuilder() UnderlyingSymbolsForExchangeRequestFixedPointerBuilder {
	return UnderlyingSymbolsForExchangeRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m UnderlyingSymbolsForExchangeRequestBuilder) Finish() UnderlyingSymbolsForExchangeRequest {
	return UnderlyingSymbolsForExchangeRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) Finish() UnderlyingSymbolsForExchangeRequestFixed {
	return UnderlyingSymbolsForExchangeRequestFixed{m.Fixed}
}

// Finish
func (m *UnderlyingSymbolsForExchangeRequestPointerBuilder) Finish() UnderlyingSymbolsForExchangeRequestPointer {
	return UnderlyingSymbolsForExchangeRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) Finish() UnderlyingSymbolsForExchangeRequestFixedPointer {
	return UnderlyingSymbolsForExchangeRequestFixedPointer{m.FixedPointer}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequest) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 20, 16))
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 20, 16))
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestPointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 20, 16))
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 20, 16))
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 24, 8)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 24, 8)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 24, 8)
}

// Exchange .
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 24, 8)
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestFixed) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 28, 24))
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 28, 24))
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Ptr, 28, 24))
}

// SecurityType .
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Ptr, 28, 24))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetExchange .
func (m *UnderlyingSymbolsForExchangeRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetExchange .
func (m *UnderlyingSymbolsForExchangeRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetSecurityType .
func (m UnderlyingSymbolsForExchangeRequestBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 20, 16, int32(value))
}

// SetSecurityType .
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Ptr, 20, 16, int32(value))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetExchange .
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 24, 8, value)
}

// SetExchange .
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 24, 8, value)
}

// SetSecurityType .
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 28, 24, int32(value))
}

// SetSecurityType .
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 28, 24, int32(value))
}

// Copy
func (m UnderlyingSymbolsForExchangeRequest) Copy(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestBuilder) Copy(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequest) CopyPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestBuilder) CopyPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequest) CopyToPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestBuilder) CopyToPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequest) CopyTo(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestBuilder) CopyTo(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestFixed) Copy(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) Copy(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestFixed) CopyPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) CopyPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestFixed) CopyToPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) CopyToPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestFixed) CopyTo(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) CopyTo(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestPointer) Copy(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) Copy(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestPointer) CopyPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) CopyPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestPointer) CopyTo(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) CopyTo(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestPointer) CopyToPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) CopyToPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) Copy(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) Copy(to UnderlyingSymbolsForExchangeRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) CopyPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyPointer
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) CopyPointer(to UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) CopyTo(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) CopyTo(to UnderlyingSymbolsForExchangeRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestFixedPointer) CopyToPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyToPointer
func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) CopyToPointer(to UnderlyingSymbolsForExchangeRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}
