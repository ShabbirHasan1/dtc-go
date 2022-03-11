package model

import (
	"github.com/moontrade/dtc-go/message"
)

type AccountBalanceRejectFixedPointer struct {
	message.FixedPointer
}

type AccountBalanceRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocAccountBalanceRejectFixed() AccountBalanceRejectFixedPointerBuilder {
	a := AccountBalanceRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _AccountBalanceRejectFixedDefault)
	return a
}

func AllocAccountBalanceRejectFixedFrom(b []byte) AccountBalanceRejectFixedPointer {
	return AccountBalanceRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = AccountBalanceRejectFixedSize  (104)
//     Type       = ACCOUNT_BALANCE_REJECT  (602)
//     RequestID  = 0
//     RejectText = ""
// }
func (m AccountBalanceRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceRejectFixedDefault)
}

// ToBuilder
func (m AccountBalanceRejectFixedPointer) ToBuilder() AccountBalanceRejectFixedPointerBuilder {
	return AccountBalanceRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *AccountBalanceRejectFixedPointerBuilder) Finish() AccountBalanceRejectFixedPointer {
	return AccountBalanceRejectFixedPointer{m.FixedPointer}
}

// RequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m AccountBalanceRejectFixedPointer) Copy(to AccountBalanceRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m AccountBalanceRejectFixedPointerBuilder) Copy(to AccountBalanceRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceRejectFixedPointer) CopyPointer(to AccountBalanceRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceRejectFixedPointerBuilder) CopyPointer(to AccountBalanceRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceRejectFixedPointer) CopyTo(to AccountBalanceRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceRejectFixedPointerBuilder) CopyTo(to AccountBalanceRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceRejectFixedPointer) CopyToPointer(to AccountBalanceRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceRejectFixedPointerBuilder) CopyToPointer(to AccountBalanceRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
