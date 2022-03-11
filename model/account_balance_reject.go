package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = AccountBalanceRejectSize  (16)
//     Type       = ACCOUNT_BALANCE_REJECT  (602)
//     BaseSize   = AccountBalanceRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _AccountBalanceRejectDefault = []byte{16, 0, 90, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AccountBalanceRejectSize = 16

type AccountBalanceReject struct {
	message.VLS
}

type AccountBalanceRejectBuilder struct {
	message.VLSBuilder
}

func NewAccountBalanceRejectFrom(b []byte) AccountBalanceReject {
	return AccountBalanceReject{VLS: message.NewVLSFrom(b)}
}

func WrapAccountBalanceReject(b []byte) AccountBalanceReject {
	return AccountBalanceReject{VLS: message.WrapVLS(b)}
}

func NewAccountBalanceReject() AccountBalanceRejectBuilder {
	a := AccountBalanceRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _AccountBalanceRejectDefault)
	return a
}

// Clear
// {
//     Size       = AccountBalanceRejectSize  (16)
//     Type       = ACCOUNT_BALANCE_REJECT  (602)
//     BaseSize   = AccountBalanceRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m AccountBalanceRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceRejectDefault)
}

// ToBuilder
func (m AccountBalanceReject) ToBuilder() AccountBalanceRejectBuilder {
	return AccountBalanceRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m AccountBalanceRejectBuilder) Finish() AccountBalanceReject {
	return AccountBalanceReject{m.VLSBuilder.Finish()}
}

// RequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetRequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m *AccountBalanceRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// Copy
func (m AccountBalanceReject) Copy(to AccountBalanceRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m AccountBalanceRejectBuilder) Copy(to AccountBalanceRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceReject) CopyPointer(to AccountBalanceRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceRejectBuilder) CopyPointer(to AccountBalanceRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceReject) CopyToPointer(to AccountBalanceRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceRejectBuilder) CopyToPointer(to AccountBalanceRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceReject) CopyTo(to AccountBalanceRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceRejectBuilder) CopyTo(to AccountBalanceRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
