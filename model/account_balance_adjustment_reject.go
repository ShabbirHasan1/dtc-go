package model

import (
	"github.com/moontrade/dtc-go/message"
)

const AccountBalanceAdjustmentRejectSize = 16

const AccountBalanceAdjustmentRejectFixedSize = 104

// {
//     Size       = AccountBalanceAdjustmentRejectSize  (16)
//     Type       = ACCOUNT_BALANCE_ADJUSTMENT_REJECT  (608)
//     BaseSize   = AccountBalanceAdjustmentRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _AccountBalanceAdjustmentRejectDefault = []byte{16, 0, 96, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size       = AccountBalanceAdjustmentRejectFixedSize  (104)
//     Type       = ACCOUNT_BALANCE_ADJUSTMENT_REJECT  (608)
//     RequestID  = 0
//     RejectText = ""
// }
var _AccountBalanceAdjustmentRejectFixedDefault = []byte{104, 0, 96, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type AccountBalanceAdjustmentReject struct {
	message.VLS
}

type AccountBalanceAdjustmentRejectBuilder struct {
	message.VLSBuilder
}

type AccountBalanceAdjustmentRejectFixed struct {
	message.Fixed
}

type AccountBalanceAdjustmentRejectFixedBuilder struct {
	message.Fixed
}

type AccountBalanceAdjustmentRejectPointer struct {
	message.VLSPointer
}

type AccountBalanceAdjustmentRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

type AccountBalanceAdjustmentRejectFixedPointer struct {
	message.FixedPointer
}

type AccountBalanceAdjustmentRejectFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocAccountBalanceAdjustmentReject() AccountBalanceAdjustmentRejectPointerBuilder {
	a := AccountBalanceAdjustmentRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _AccountBalanceAdjustmentRejectDefault)
	return a
}

func AllocAccountBalanceAdjustmentRejectFrom(b []byte) AccountBalanceAdjustmentRejectPointer {
	return AccountBalanceAdjustmentRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size       = AccountBalanceAdjustmentRejectSize  (16)
//     Type       = ACCOUNT_BALANCE_ADJUSTMENT_REJECT  (608)
//     BaseSize   = AccountBalanceAdjustmentRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m AccountBalanceAdjustmentRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceAdjustmentRejectDefault)
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
func (m AccountBalanceAdjustmentReject) ToBuilder() AccountBalanceAdjustmentRejectBuilder {
	return AccountBalanceAdjustmentRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m AccountBalanceAdjustmentRejectFixed) ToBuilder() AccountBalanceAdjustmentRejectFixedBuilder {
	return AccountBalanceAdjustmentRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m AccountBalanceAdjustmentRejectPointer) ToBuilder() AccountBalanceAdjustmentRejectPointerBuilder {
	return AccountBalanceAdjustmentRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m AccountBalanceAdjustmentRejectFixedPointer) ToBuilder() AccountBalanceAdjustmentRejectFixedPointerBuilder {
	return AccountBalanceAdjustmentRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m AccountBalanceAdjustmentRejectBuilder) Finish() AccountBalanceAdjustmentReject {
	return AccountBalanceAdjustmentReject{m.VLSBuilder.Finish()}
}

// Finish
func (m AccountBalanceAdjustmentRejectFixedBuilder) Finish() AccountBalanceAdjustmentRejectFixed {
	return AccountBalanceAdjustmentRejectFixed{m.Fixed}
}

// Finish
func (m *AccountBalanceAdjustmentRejectPointerBuilder) Finish() AccountBalanceAdjustmentRejectPointer {
	return AccountBalanceAdjustmentRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *AccountBalanceAdjustmentRejectFixedPointerBuilder) Finish() AccountBalanceAdjustmentRejectFixedPointer {
	return AccountBalanceAdjustmentRejectFixedPointer{m.FixedPointer}
}

// RequestID
func (m AccountBalanceAdjustmentReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m AccountBalanceAdjustmentRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
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
func (m AccountBalanceAdjustmentReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText
func (m AccountBalanceAdjustmentRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText
func (m AccountBalanceAdjustmentRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText
func (m AccountBalanceAdjustmentRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RequestID
func (m AccountBalanceAdjustmentRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m AccountBalanceAdjustmentRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m AccountBalanceAdjustmentRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m AccountBalanceAdjustmentRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
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
func (m AccountBalanceAdjustmentRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID
func (m AccountBalanceAdjustmentRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText
func (m *AccountBalanceAdjustmentRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText
func (m *AccountBalanceAdjustmentRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRequestID
func (m AccountBalanceAdjustmentRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m AccountBalanceAdjustmentRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText
func (m AccountBalanceAdjustmentRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText
func (m AccountBalanceAdjustmentRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
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
