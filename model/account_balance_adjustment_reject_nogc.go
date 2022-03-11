package model

import (
	"github.com/moontrade/dtc-go/message"
)

type AccountBalanceAdjustmentRejectPointer struct {
	message.VLSPointer
}

type AccountBalanceAdjustmentRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocAccountBalanceAdjustmentReject() AccountBalanceAdjustmentRejectPointerBuilder {
	a := AccountBalanceAdjustmentRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _AccountBalanceAdjustmentRejectDefault)
	return a
}

func AllocAccountBalanceAdjustmentRejectFrom(b []byte) AccountBalanceAdjustmentRejectPointer {
	return AccountBalanceAdjustmentRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size       = AccountBalanceAdjustmentRejectSize  (16)
//     Type       = ACCOUNT_BALANCE_ADJUSTMENT_REJECT  (608)
//     BaseSize   = AccountBalanceAdjustmentRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m AccountBalanceAdjustmentRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceAdjustmentRejectDefault)
}

// ToBuilder
func (m AccountBalanceAdjustmentRejectPointer) ToBuilder() AccountBalanceAdjustmentRejectPointerBuilder {
	return AccountBalanceAdjustmentRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *AccountBalanceAdjustmentRejectPointerBuilder) Finish() AccountBalanceAdjustmentRejectPointer {
	return AccountBalanceAdjustmentRejectPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m AccountBalanceAdjustmentRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID
func (m AccountBalanceAdjustmentRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText
func (m AccountBalanceAdjustmentRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText
func (m AccountBalanceAdjustmentRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetRequestID
func (m AccountBalanceAdjustmentRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText
func (m *AccountBalanceAdjustmentRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// Copy
func (m AccountBalanceAdjustmentRejectPointer) Copy(to AccountBalanceAdjustmentRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m AccountBalanceAdjustmentRejectPointerBuilder) Copy(to AccountBalanceAdjustmentRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceAdjustmentRejectPointer) CopyPointer(to AccountBalanceAdjustmentRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceAdjustmentRejectPointerBuilder) CopyPointer(to AccountBalanceAdjustmentRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceAdjustmentRejectPointer) CopyTo(to AccountBalanceAdjustmentRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceAdjustmentRejectPointerBuilder) CopyTo(to AccountBalanceAdjustmentRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceAdjustmentRejectPointer) CopyToPointer(to AccountBalanceAdjustmentRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceAdjustmentRejectPointerBuilder) CopyToPointer(to AccountBalanceAdjustmentRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
