package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
//     Size       = SecurityDefinitionRejectFixedSize  (104)
//     Type       = SECURITY_DEFINITION_REJECT  (509)
//     RequestID  = 0
//     RejectText = ""
// }
func (m SecurityDefinitionRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SecurityDefinitionRejectFixedDefault)
}

// ToBuilder
func (m SecurityDefinitionRejectFixedPointer) ToBuilder() SecurityDefinitionRejectFixedPointerBuilder {
	return SecurityDefinitionRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *SecurityDefinitionRejectFixedPointerBuilder) Finish() SecurityDefinitionRejectFixedPointer {
	return SecurityDefinitionRejectFixedPointer{m.FixedPointer}
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
func (m SecurityDefinitionRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText Text reason for rejection.
func (m SecurityDefinitionRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original request message.
func (m SecurityDefinitionRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Text reason for rejection.
func (m SecurityDefinitionRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
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
