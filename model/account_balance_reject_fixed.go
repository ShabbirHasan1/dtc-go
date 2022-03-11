package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = AccountBalanceRejectFixedSize  (104)
//     Type       = ACCOUNT_BALANCE_REJECT  (602)
//     RequestID  = 0
//     RejectText = ""
// }
var _AccountBalanceRejectFixedDefault = []byte{104, 0, 90, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AccountBalanceRejectFixedSize = 104

type AccountBalanceRejectFixed struct {
	message.Fixed
}

type AccountBalanceRejectFixedBuilder struct {
	message.Fixed
}

func NewAccountBalanceRejectFixedFrom(b []byte) AccountBalanceRejectFixed {
	return AccountBalanceRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapAccountBalanceRejectFixed(b []byte) AccountBalanceRejectFixed {
	return AccountBalanceRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewAccountBalanceRejectFixed() AccountBalanceRejectFixedBuilder {
	a := AccountBalanceRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _AccountBalanceRejectFixedDefault)
	return a
}

// Clear
// {
//     Size       = AccountBalanceRejectFixedSize  (104)
//     Type       = ACCOUNT_BALANCE_REJECT  (602)
//     RequestID  = 0
//     RejectText = ""
// }
func (m AccountBalanceRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceRejectFixedDefault)
}

// ToBuilder
func (m AccountBalanceRejectFixed) ToBuilder() AccountBalanceRejectFixedBuilder {
	return AccountBalanceRejectFixedBuilder{m.Fixed}
}

// Finish
func (m AccountBalanceRejectFixedBuilder) Finish() AccountBalanceRejectFixed {
	return AccountBalanceRejectFixed{m.Fixed}
}

// RequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// SetRequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// Copy
func (m AccountBalanceRejectFixed) Copy(to AccountBalanceRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m AccountBalanceRejectFixedBuilder) Copy(to AccountBalanceRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceRejectFixed) CopyPointer(to AccountBalanceRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceRejectFixedBuilder) CopyPointer(to AccountBalanceRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceRejectFixed) CopyToPointer(to AccountBalanceRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceRejectFixedBuilder) CopyToPointer(to AccountBalanceRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceRejectFixed) CopyTo(to AccountBalanceRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceRejectFixedBuilder) CopyTo(to AccountBalanceRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
