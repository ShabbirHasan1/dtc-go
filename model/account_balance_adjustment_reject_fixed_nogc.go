package model

import (
	"github.com/moontrade/dtc-go/message"
)

type AccountBalanceAdjustmentRejectFixedPointer struct {
	message.FixedPointer
}

type AccountBalanceAdjustmentRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocAccountBalanceAdjustmentRejectFixed() AccountBalanceAdjustmentRejectFixedPointerBuilder {
	a := AccountBalanceAdjustmentRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _AccountBalanceAdjustmentRejectFixedDefault)
	return a
}

func AllocAccountBalanceAdjustmentRejectFixedFrom(b []byte) AccountBalanceAdjustmentRejectFixedPointer {
	return AccountBalanceAdjustmentRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = AccountBalanceAdjustmentRejectFixedSize  (104)
//     Type       = ACCOUNT_BALANCE_ADJUSTMENT_REJECT  (608)
//     RequestID  = 0
//     RejectText = ""
// }
func (m AccountBalanceAdjustmentRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceAdjustmentRejectFixedDefault)
}

// ToBuilder
func (m AccountBalanceAdjustmentRejectFixedPointer) ToBuilder() AccountBalanceAdjustmentRejectFixedPointerBuilder {
	return AccountBalanceAdjustmentRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *AccountBalanceAdjustmentRejectFixedPointerBuilder) Finish() AccountBalanceAdjustmentRejectFixedPointer {
	return AccountBalanceAdjustmentRejectFixedPointer{m.FixedPointer}
}

// RequestID
func (m AccountBalanceAdjustmentRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m AccountBalanceAdjustmentRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText
func (m AccountBalanceAdjustmentRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText
func (m AccountBalanceAdjustmentRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID
func (m AccountBalanceAdjustmentRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText
func (m AccountBalanceAdjustmentRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m AccountBalanceAdjustmentRejectFixedPointer) Copy(to AccountBalanceAdjustmentRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m AccountBalanceAdjustmentRejectFixedPointerBuilder) Copy(to AccountBalanceAdjustmentRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceAdjustmentRejectFixedPointer) CopyPointer(to AccountBalanceAdjustmentRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceAdjustmentRejectFixedPointerBuilder) CopyPointer(to AccountBalanceAdjustmentRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceAdjustmentRejectFixedPointer) CopyTo(to AccountBalanceAdjustmentRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceAdjustmentRejectFixedPointerBuilder) CopyTo(to AccountBalanceAdjustmentRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceAdjustmentRejectFixedPointer) CopyToPointer(to AccountBalanceAdjustmentRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceAdjustmentRejectFixedPointerBuilder) CopyToPointer(to AccountBalanceAdjustmentRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
