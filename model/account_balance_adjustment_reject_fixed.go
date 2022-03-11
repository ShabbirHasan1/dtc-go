package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = AccountBalanceAdjustmentRejectFixedSize  (104)
//     Type       = ACCOUNT_BALANCE_ADJUSTMENT_REJECT  (608)
//     RequestID  = 0
//     RejectText = ""
// }
var _AccountBalanceAdjustmentRejectFixedDefault = []byte{104, 0, 96, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AccountBalanceAdjustmentRejectFixedSize = 104

type AccountBalanceAdjustmentRejectFixed struct {
	message.Fixed
}

type AccountBalanceAdjustmentRejectFixedBuilder struct {
	message.Fixed
}

func NewAccountBalanceAdjustmentRejectFixedFrom(b []byte) AccountBalanceAdjustmentRejectFixed {
	return AccountBalanceAdjustmentRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapAccountBalanceAdjustmentRejectFixed(b []byte) AccountBalanceAdjustmentRejectFixed {
	return AccountBalanceAdjustmentRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewAccountBalanceAdjustmentRejectFixed() AccountBalanceAdjustmentRejectFixedBuilder {
	a := AccountBalanceAdjustmentRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _AccountBalanceAdjustmentRejectFixedDefault)
	return a
}

// Clear
// {
//     Size       = AccountBalanceAdjustmentRejectFixedSize  (104)
//     Type       = ACCOUNT_BALANCE_ADJUSTMENT_REJECT  (608)
//     RequestID  = 0
//     RejectText = ""
// }
func (m AccountBalanceAdjustmentRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceAdjustmentRejectFixedDefault)
}

// ToBuilder
func (m AccountBalanceAdjustmentRejectFixed) ToBuilder() AccountBalanceAdjustmentRejectFixedBuilder {
	return AccountBalanceAdjustmentRejectFixedBuilder{m.Fixed}
}

// Finish
func (m AccountBalanceAdjustmentRejectFixedBuilder) Finish() AccountBalanceAdjustmentRejectFixed {
	return AccountBalanceAdjustmentRejectFixed{m.Fixed}
}

// RequestID
func (m AccountBalanceAdjustmentRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m AccountBalanceAdjustmentRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RejectText
func (m AccountBalanceAdjustmentRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m AccountBalanceAdjustmentRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// SetRequestID
func (m AccountBalanceAdjustmentRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText
func (m AccountBalanceAdjustmentRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// Copy
func (m AccountBalanceAdjustmentRejectFixed) Copy(to AccountBalanceAdjustmentRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m AccountBalanceAdjustmentRejectFixedBuilder) Copy(to AccountBalanceAdjustmentRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceAdjustmentRejectFixed) CopyPointer(to AccountBalanceAdjustmentRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m AccountBalanceAdjustmentRejectFixedBuilder) CopyPointer(to AccountBalanceAdjustmentRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceAdjustmentRejectFixed) CopyToPointer(to AccountBalanceAdjustmentRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m AccountBalanceAdjustmentRejectFixedBuilder) CopyToPointer(to AccountBalanceAdjustmentRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceAdjustmentRejectFixed) CopyTo(to AccountBalanceAdjustmentRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m AccountBalanceAdjustmentRejectFixedBuilder) CopyTo(to AccountBalanceAdjustmentRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
