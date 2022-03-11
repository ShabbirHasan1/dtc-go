package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = AccountBalanceAdjustmentRejectSize  (16)
//     Type       = ACCOUNT_BALANCE_ADJUSTMENT_REJECT  (608)
//     BaseSize   = AccountBalanceAdjustmentRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _AccountBalanceAdjustmentRejectDefault = []byte{16, 0, 96, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AccountBalanceAdjustmentRejectSize = 16

type AccountBalanceAdjustmentReject struct {
	message.VLS
}

type AccountBalanceAdjustmentRejectBuilder struct {
	message.VLSBuilder
}

func NewAccountBalanceAdjustmentRejectFrom(b []byte) AccountBalanceAdjustmentReject {
	return AccountBalanceAdjustmentReject{VLS: message.NewVLSFrom(b)}
}

func WrapAccountBalanceAdjustmentReject(b []byte) AccountBalanceAdjustmentReject {
	return AccountBalanceAdjustmentReject{VLS: message.WrapVLS(b)}
}

func NewAccountBalanceAdjustmentReject() AccountBalanceAdjustmentRejectBuilder {
	a := AccountBalanceAdjustmentRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _AccountBalanceAdjustmentRejectDefault)
	return a
}

// Clear
// {
//     Size       = AccountBalanceAdjustmentRejectSize  (16)
//     Type       = ACCOUNT_BALANCE_ADJUSTMENT_REJECT  (608)
//     BaseSize   = AccountBalanceAdjustmentRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m AccountBalanceAdjustmentRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceAdjustmentRejectDefault)
}

// ToBuilder
func (m AccountBalanceAdjustmentReject) ToBuilder() AccountBalanceAdjustmentRejectBuilder {
	return AccountBalanceAdjustmentRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m AccountBalanceAdjustmentRejectBuilder) Finish() AccountBalanceAdjustmentReject {
	return AccountBalanceAdjustmentReject{m.VLSBuilder.Finish()}
}

// RequestID
func (m AccountBalanceAdjustmentReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m AccountBalanceAdjustmentRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RejectText
func (m AccountBalanceAdjustmentReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText
func (m AccountBalanceAdjustmentRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetRequestID
func (m AccountBalanceAdjustmentRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText
func (m *AccountBalanceAdjustmentRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// Copy
func (m AccountBalanceAdjustmentReject) Copy(to AccountBalanceAdjustmentRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m AccountBalanceAdjustmentRejectBuilder) Copy(to AccountBalanceAdjustmentRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceAdjustmentReject) CopyPointer(to AccountBalanceAdjustmentRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceAdjustmentRejectBuilder) CopyPointer(to AccountBalanceAdjustmentRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceAdjustmentReject) CopyToPointer(to AccountBalanceAdjustmentRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceAdjustmentRejectBuilder) CopyToPointer(to AccountBalanceAdjustmentRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceAdjustmentReject) CopyTo(to AccountBalanceAdjustmentRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceAdjustmentRejectBuilder) CopyTo(to AccountBalanceAdjustmentRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
