package model

import (
	"github.com/moontrade/dtc-go/message"
)

type AccountBalanceRejectPointer struct {
	message.VLSPointer
}

type AccountBalanceRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocAccountBalanceReject() AccountBalanceRejectPointerBuilder {
	a := AccountBalanceRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _AccountBalanceRejectDefault)
	return a
}

func AllocAccountBalanceRejectFrom(b []byte) AccountBalanceRejectPointer {
	return AccountBalanceRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size       = AccountBalanceRejectSize  (16)
//     Type       = ACCOUNT_BALANCE_REJECT  (602)
//     BaseSize   = AccountBalanceRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m AccountBalanceRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceRejectDefault)
}

// ToBuilder
func (m AccountBalanceRejectPointer) ToBuilder() AccountBalanceRejectPointerBuilder {
	return AccountBalanceRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *AccountBalanceRejectPointerBuilder) Finish() AccountBalanceRejectPointer {
	return AccountBalanceRejectPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetRequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m *AccountBalanceRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// Copy
func (m AccountBalanceRejectPointer) Copy(to AccountBalanceRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m AccountBalanceRejectPointerBuilder) Copy(to AccountBalanceRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceRejectPointer) CopyPointer(to AccountBalanceRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceRejectPointerBuilder) CopyPointer(to AccountBalanceRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceRejectPointer) CopyTo(to AccountBalanceRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceRejectPointerBuilder) CopyTo(to AccountBalanceRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceRejectPointer) CopyToPointer(to AccountBalanceRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceRejectPointerBuilder) CopyToPointer(to AccountBalanceRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
