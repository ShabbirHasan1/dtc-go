package model

import (
	"github.com/moontrade/dtc-go/message"
)

const SecurityDefinitionRejectSize = 16

const SecurityDefinitionRejectFixedSize = 104

// {
//     Size       = SecurityDefinitionRejectSize  (16)
//     Type       = SECURITY_DEFINITION_REJECT  (509)
//     BaseSize   = SecurityDefinitionRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _SecurityDefinitionRejectDefault = []byte{16, 0, 253, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size       = SecurityDefinitionRejectFixedSize  (104)
//     Type       = SECURITY_DEFINITION_REJECT  (509)
//     RequestID  = 0
//     RejectText = ""
// }
var _SecurityDefinitionRejectFixedDefault = []byte{104, 0, 253, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// SecurityDefinitionReject This is a message from the Server to the Client indicating the Server
// is rejecting one of the following messages: SymbolsForExchangeRequest,
// UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest, SecurityDefinitionForSymbolRequest,
// SymbolSearchRequest.
//
// If there are no symbols to send in response to one of these messages above,
// then the Server should not send a SecurityDefinitionReject message and
// instead send a SecurityDefinitionResponse with only the RequestID and
// IsFinalMessage fields set. This will be a clear indication to the Client
// that the request returned no Symbols.
type SecurityDefinitionReject struct {
	message.VLS
}

// SecurityDefinitionRejectBuilder This is a message from the Server to the Client indicating the Server
// is rejecting one of the following messages: SymbolsForExchangeRequest,
// UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest, SecurityDefinitionForSymbolRequest,
// SymbolSearchRequest.
//
// If there are no symbols to send in response to one of these messages above,
// then the Server should not send a SecurityDefinitionReject message and
// instead send a SecurityDefinitionResponse with only the RequestID and
// IsFinalMessage fields set. This will be a clear indication to the Client
// that the request returned no Symbols.
type SecurityDefinitionRejectBuilder struct {
	message.VLSBuilder
}

// SecurityDefinitionRejectFixed This is a message from the Server to the Client indicating the Server
// is rejecting one of the following messages: SymbolsForExchangeRequest,
// UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest, SecurityDefinitionForSymbolRequest,
// SymbolSearchRequest.
//
// If there are no symbols to send in response to one of these messages above,
// then the Server should not send a SecurityDefinitionReject message and
// instead send a SecurityDefinitionResponse with only the RequestID and
// IsFinalMessage fields set. This will be a clear indication to the Client
// that the request returned no Symbols.
type SecurityDefinitionRejectFixed struct {
	message.Fixed
}

// SecurityDefinitionRejectFixedBuilder This is a message from the Server to the Client indicating the Server
// is rejecting one of the following messages: SymbolsForExchangeRequest,
// UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest, SecurityDefinitionForSymbolRequest,
// SymbolSearchRequest.
//
// If there are no symbols to send in response to one of these messages above,
// then the Server should not send a SecurityDefinitionReject message and
// instead send a SecurityDefinitionResponse with only the RequestID and
// IsFinalMessage fields set. This will be a clear indication to the Client
// that the request returned no Symbols.
type SecurityDefinitionRejectFixedBuilder struct {
	message.Fixed
}

// SecurityDefinitionRejectPointer This is a message from the Server to the Client indicating the Server
// is rejecting one of the following messages: SymbolsForExchangeRequest,
// UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest, SecurityDefinitionForSymbolRequest,
// SymbolSearchRequest.
//
// If there are no symbols to send in response to one of these messages above,
// then the Server should not send a SecurityDefinitionReject message and
// instead send a SecurityDefinitionResponse with only the RequestID and
// IsFinalMessage fields set. This will be a clear indication to the Client
// that the request returned no Symbols.
type SecurityDefinitionRejectPointer struct {
	message.VLSPointer
}

// SecurityDefinitionRejectPointerBuilder This is a message from the Server to the Client indicating the Server
// is rejecting one of the following messages: SymbolsForExchangeRequest,
// UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest, SecurityDefinitionForSymbolRequest,
// SymbolSearchRequest.
//
// If there are no symbols to send in response to one of these messages above,
// then the Server should not send a SecurityDefinitionReject message and
// instead send a SecurityDefinitionResponse with only the RequestID and
// IsFinalMessage fields set. This will be a clear indication to the Client
// that the request returned no Symbols.
type SecurityDefinitionRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

// SecurityDefinitionRejectFixedPointer This is a message from the Server to the Client indicating the Server
// is rejecting one of the following messages: SymbolsForExchangeRequest,
// UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest, SecurityDefinitionForSymbolRequest,
// SymbolSearchRequest.
//
// If there are no symbols to send in response to one of these messages above,
// then the Server should not send a SecurityDefinitionReject message and
// instead send a SecurityDefinitionResponse with only the RequestID and
// IsFinalMessage fields set. This will be a clear indication to the Client
// that the request returned no Symbols.
type SecurityDefinitionRejectFixedPointer struct {
	message.FixedPointer
}

// SecurityDefinitionRejectFixedPointerBuilder This is a message from the Server to the Client indicating the Server
// is rejecting one of the following messages: SymbolsForExchangeRequest,
// UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest, SecurityDefinitionForSymbolRequest,
// SymbolSearchRequest.
//
// If there are no symbols to send in response to one of these messages above,
// then the Server should not send a SecurityDefinitionReject message and
// instead send a SecurityDefinitionResponse with only the RequestID and
// IsFinalMessage fields set. This will be a clear indication to the Client
// that the request returned no Symbols.
type SecurityDefinitionRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func NewSecurityDefinitionRejectFrom(b []byte) SecurityDefinitionReject {
	return SecurityDefinitionReject{VLS: message.NewVLSFrom(b)}
}

func WrapSecurityDefinitionReject(b []byte) SecurityDefinitionReject {
	return SecurityDefinitionReject{VLS: message.WrapVLS(b)}
}

func NewSecurityDefinitionReject() SecurityDefinitionRejectBuilder {
	a := SecurityDefinitionRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _SecurityDefinitionRejectDefault)
	return a
}

func NewSecurityDefinitionRejectFixedFrom(b []byte) SecurityDefinitionRejectFixed {
	return SecurityDefinitionRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapSecurityDefinitionRejectFixed(b []byte) SecurityDefinitionRejectFixed {
	return SecurityDefinitionRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewSecurityDefinitionRejectFixed() SecurityDefinitionRejectFixedBuilder {
	a := SecurityDefinitionRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _SecurityDefinitionRejectFixedDefault)
	return a
}

func AllocSecurityDefinitionReject() SecurityDefinitionRejectPointerBuilder {
	a := SecurityDefinitionRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _SecurityDefinitionRejectDefault)
	return a
}

func AllocSecurityDefinitionRejectFrom(b []byte) SecurityDefinitionRejectPointer {
	return SecurityDefinitionRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocSecurityDefinitionRejectFixed() SecurityDefinitionRejectFixedPointerBuilder {
	a := SecurityDefinitionRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _SecurityDefinitionRejectFixedDefault)
	return a
}

func AllocSecurityDefinitionRejectFixedFrom(b []byte) SecurityDefinitionRejectFixedPointer {
	return SecurityDefinitionRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = SecurityDefinitionRejectSize  (16)
//     Type       = SECURITY_DEFINITION_REJECT  (509)
//     BaseSize   = SecurityDefinitionRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m SecurityDefinitionRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SecurityDefinitionRejectDefault)
}

// Clear
// {
//     Size       = SecurityDefinitionRejectFixedSize  (104)
//     Type       = SECURITY_DEFINITION_REJECT  (509)
//     RequestID  = 0
//     RejectText = ""
// }
func (m SecurityDefinitionRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SecurityDefinitionRejectFixedDefault)
}

// Clear
// {
//     Size       = SecurityDefinitionRejectSize  (16)
//     Type       = SECURITY_DEFINITION_REJECT  (509)
//     BaseSize   = SecurityDefinitionRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m SecurityDefinitionRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SecurityDefinitionRejectDefault)
}

// Clear
// {
//     Size       = SecurityDefinitionRejectFixedSize  (104)
//     Type       = SECURITY_DEFINITION_REJECT  (509)
//     RequestID  = 0
//     RejectText = ""
// }
func (m SecurityDefinitionRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SecurityDefinitionRejectFixedDefault)
}

// ToBuilder
func (m SecurityDefinitionReject) ToBuilder() SecurityDefinitionRejectBuilder {
	return SecurityDefinitionRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m SecurityDefinitionRejectFixed) ToBuilder() SecurityDefinitionRejectFixedBuilder {
	return SecurityDefinitionRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m SecurityDefinitionRejectPointer) ToBuilder() SecurityDefinitionRejectPointerBuilder {
	return SecurityDefinitionRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m SecurityDefinitionRejectFixedPointer) ToBuilder() SecurityDefinitionRejectFixedPointerBuilder {
	return SecurityDefinitionRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m SecurityDefinitionRejectBuilder) Finish() SecurityDefinitionReject {
	return SecurityDefinitionReject{m.VLSBuilder.Finish()}
}

// Finish
func (m SecurityDefinitionRejectFixedBuilder) Finish() SecurityDefinitionRejectFixed {
	return SecurityDefinitionRejectFixed{m.Fixed}
}

// Finish
func (m *SecurityDefinitionRejectPointerBuilder) Finish() SecurityDefinitionRejectPointer {
	return SecurityDefinitionRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *SecurityDefinitionRejectFixedPointerBuilder) Finish() SecurityDefinitionRejectFixedPointer {
	return SecurityDefinitionRejectFixedPointer{m.FixedPointer}
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText Text reason for rejection.
func (m SecurityDefinitionReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Text reason for rejection.
func (m SecurityDefinitionRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Text reason for rejection.
func (m SecurityDefinitionRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Text reason for rejection.
func (m SecurityDefinitionRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText Text reason for rejection.
func (m SecurityDefinitionRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Text reason for rejection.
func (m SecurityDefinitionRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Text reason for rejection.
func (m SecurityDefinitionRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText Text reason for rejection.
func (m SecurityDefinitionRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Text reason for rejection.
func (m *SecurityDefinitionRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText Text reason for rejection.
func (m *SecurityDefinitionRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Text reason for rejection.
func (m SecurityDefinitionRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText Text reason for rejection.
func (m SecurityDefinitionRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m SecurityDefinitionReject) Copy(to SecurityDefinitionRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m SecurityDefinitionRejectBuilder) Copy(to SecurityDefinitionRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m SecurityDefinitionReject) CopyPointer(to SecurityDefinitionRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m SecurityDefinitionRejectBuilder) CopyPointer(to SecurityDefinitionRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m SecurityDefinitionReject) CopyToPointer(to SecurityDefinitionRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m SecurityDefinitionRejectBuilder) CopyToPointer(to SecurityDefinitionRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m SecurityDefinitionReject) CopyTo(to SecurityDefinitionRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m SecurityDefinitionRejectBuilder) CopyTo(to SecurityDefinitionRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m SecurityDefinitionRejectFixed) Copy(to SecurityDefinitionRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m SecurityDefinitionRejectFixedBuilder) Copy(to SecurityDefinitionRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m SecurityDefinitionRejectFixed) CopyPointer(to SecurityDefinitionRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m SecurityDefinitionRejectFixedBuilder) CopyPointer(to SecurityDefinitionRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m SecurityDefinitionRejectFixed) CopyToPointer(to SecurityDefinitionRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m SecurityDefinitionRejectFixedBuilder) CopyToPointer(to SecurityDefinitionRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m SecurityDefinitionRejectFixed) CopyTo(to SecurityDefinitionRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m SecurityDefinitionRejectFixedBuilder) CopyTo(to SecurityDefinitionRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m SecurityDefinitionRejectPointer) Copy(to SecurityDefinitionRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m SecurityDefinitionRejectPointerBuilder) Copy(to SecurityDefinitionRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m SecurityDefinitionRejectPointer) CopyPointer(to SecurityDefinitionRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m SecurityDefinitionRejectPointerBuilder) CopyPointer(to SecurityDefinitionRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m SecurityDefinitionRejectPointer) CopyTo(to SecurityDefinitionRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m SecurityDefinitionRejectPointerBuilder) CopyTo(to SecurityDefinitionRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m SecurityDefinitionRejectPointer) CopyToPointer(to SecurityDefinitionRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m SecurityDefinitionRejectPointerBuilder) CopyToPointer(to SecurityDefinitionRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m SecurityDefinitionRejectFixedPointer) Copy(to SecurityDefinitionRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m SecurityDefinitionRejectFixedPointerBuilder) Copy(to SecurityDefinitionRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m SecurityDefinitionRejectFixedPointer) CopyPointer(to SecurityDefinitionRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m SecurityDefinitionRejectFixedPointerBuilder) CopyPointer(to SecurityDefinitionRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m SecurityDefinitionRejectFixedPointer) CopyTo(to SecurityDefinitionRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m SecurityDefinitionRejectFixedPointerBuilder) CopyTo(to SecurityDefinitionRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m SecurityDefinitionRejectFixedPointer) CopyToPointer(to SecurityDefinitionRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m SecurityDefinitionRejectFixedPointerBuilder) CopyToPointer(to SecurityDefinitionRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}