package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = SecurityDefinitionRejectFixedSize  (104)
//     Type       = SECURITY_DEFINITION_REJECT  (509)
//     RequestID  = 0
//     RejectText = ""
// }
var _SecurityDefinitionRejectFixedDefault = []byte{104, 0, 253, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SecurityDefinitionRejectFixedSize = 104

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

// ToBuilder
func (m SecurityDefinitionRejectFixed) ToBuilder() SecurityDefinitionRejectFixedBuilder {
	return SecurityDefinitionRejectFixedBuilder{m.Fixed}
}

// Finish
func (m SecurityDefinitionRejectFixedBuilder) Finish() SecurityDefinitionRejectFixed {
	return SecurityDefinitionRejectFixed{m.Fixed}
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

// RejectText Text reason for rejection.
func (m SecurityDefinitionRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Text reason for rejection.
func (m SecurityDefinitionRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText Text reason for rejection.
func (m SecurityDefinitionRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
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
