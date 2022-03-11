package model

import (
	"github.com/moontrade/dtc-go/message"
)

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

func AllocSecurityDefinitionReject() SecurityDefinitionRejectPointerBuilder {
	a := SecurityDefinitionRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _SecurityDefinitionRejectDefault)
	return a
}

func AllocSecurityDefinitionRejectFrom(b []byte) SecurityDefinitionRejectPointer {
	return SecurityDefinitionRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
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

// ToBuilder
func (m SecurityDefinitionRejectPointer) ToBuilder() SecurityDefinitionRejectPointerBuilder {
	return SecurityDefinitionRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *SecurityDefinitionRejectPointerBuilder) Finish() SecurityDefinitionRejectPointer {
	return SecurityDefinitionRejectPointer{m.VLSPointerBuilder.Finish()}
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
func (m SecurityDefinitionRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Text reason for rejection.
func (m SecurityDefinitionRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Text reason for rejection.
func (m *SecurityDefinitionRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
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
