package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = SecurityDefinitionRejectSize  (16)
//     Type       = SECURITY_DEFINITION_REJECT  (509)
//     BaseSize   = SecurityDefinitionRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _SecurityDefinitionRejectDefault = []byte{16, 0, 253, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SecurityDefinitionRejectSize = 16

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

// ToBuilder
func (m SecurityDefinitionReject) ToBuilder() SecurityDefinitionRejectBuilder {
	return SecurityDefinitionRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m SecurityDefinitionRejectBuilder) Finish() SecurityDefinitionReject {
	return SecurityDefinitionReject{m.VLSBuilder.Finish()}
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

// RejectText Text reason for rejection.
func (m SecurityDefinitionReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Text reason for rejection.
func (m SecurityDefinitionRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText Text reason for rejection.
func (m *SecurityDefinitionRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
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
