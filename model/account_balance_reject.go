package model

import (
	"github.com/moontrade/dtc-go/message"
)

const AccountBalanceRejectSize = 16

const AccountBalanceRejectFixedSize = 104

// {
//     Size       = AccountBalanceRejectSize  (16)
//     Type       = ACCOUNT_BALANCE_REJECT  (602)
//     BaseSize   = AccountBalanceRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _AccountBalanceRejectDefault = []byte{16, 0, 90, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size       = AccountBalanceRejectFixedSize  (104)
//     Type       = ACCOUNT_BALANCE_REJECT  (602)
//     RequestID  = 0
//     RejectText = ""
// }
var _AccountBalanceRejectFixedDefault = []byte{104, 0, 90, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type AccountBalanceReject struct {
	message.VLS
}

type AccountBalanceRejectBuilder struct {
	message.VLSBuilder
}

type AccountBalanceRejectFixed struct {
	message.Fixed
}

type AccountBalanceRejectFixedBuilder struct {
	message.Fixed
}

type AccountBalanceRejectPointer struct {
	message.VLSPointer
}

type AccountBalanceRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

type AccountBalanceRejectFixedPointer struct {
	message.FixedPointer
}

type AccountBalanceRejectFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocAccountBalanceReject() AccountBalanceRejectPointerBuilder {
	a := AccountBalanceRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _AccountBalanceRejectDefault)
	return a
}

func AllocAccountBalanceRejectFrom(b []byte) AccountBalanceRejectPointer {
	return AccountBalanceRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size       = AccountBalanceRejectSize  (16)
//     Type       = ACCOUNT_BALANCE_REJECT  (602)
//     BaseSize   = AccountBalanceRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m AccountBalanceRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceRejectDefault)
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
func (m AccountBalanceReject) ToBuilder() AccountBalanceRejectBuilder {
	return AccountBalanceRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m AccountBalanceRejectFixed) ToBuilder() AccountBalanceRejectFixedBuilder {
	return AccountBalanceRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m AccountBalanceRejectPointer) ToBuilder() AccountBalanceRejectPointerBuilder {
	return AccountBalanceRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m AccountBalanceRejectFixedPointer) ToBuilder() AccountBalanceRejectFixedPointerBuilder {
	return AccountBalanceRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m AccountBalanceRejectBuilder) Finish() AccountBalanceReject {
	return AccountBalanceReject{m.VLSBuilder.Finish()}
}

// Finish
func (m AccountBalanceRejectFixedBuilder) Finish() AccountBalanceRejectFixed {
	return AccountBalanceRejectFixed{m.Fixed}
}

// Finish
func (m *AccountBalanceRejectPointerBuilder) Finish() AccountBalanceRejectPointer {
	return AccountBalanceRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *AccountBalanceRejectFixedPointerBuilder) Finish() AccountBalanceRejectFixedPointer {
	return AccountBalanceRejectFixedPointer{m.FixedPointer}
}

// RequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
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
func (m AccountBalanceReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m AccountBalanceRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
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
func (m AccountBalanceRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m *AccountBalanceRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m *AccountBalanceRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The unique request identifier sent in the corresponding request.
func (m AccountBalanceRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
func (m AccountBalanceRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
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
