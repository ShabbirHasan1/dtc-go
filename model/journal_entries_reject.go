package model

import (
	"github.com/moontrade/dtc-go/message"
)

const JournalEntriesRejectSize = 16

const JournalEntriesRejectFixedSize = 104

// {
//     Size       = JournalEntriesRejectSize  (16)
//     Type       = JOURNAL_ENTRIES_REJECT  (705)
//     BaseSize   = JournalEntriesRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _JournalEntriesRejectDefault = []byte{16, 0, 193, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size       = JournalEntriesRejectFixedSize  (104)
//     Type       = JOURNAL_ENTRIES_REJECT  (705)
//     RequestID  = 0
//     RejectText = ""
// }
var _JournalEntriesRejectFixedDefault = []byte{104, 0, 193, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type JournalEntriesReject struct {
	message.VLS
}

type JournalEntriesRejectBuilder struct {
	message.VLSBuilder
}

type JournalEntriesRejectFixed struct {
	message.Fixed
}

type JournalEntriesRejectFixedBuilder struct {
	message.Fixed
}

type JournalEntriesRejectPointer struct {
	message.VLSPointer
}

type JournalEntriesRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

type JournalEntriesRejectFixedPointer struct {
	message.FixedPointer
}

type JournalEntriesRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func NewJournalEntriesRejectFrom(b []byte) JournalEntriesReject {
	return JournalEntriesReject{VLS: message.NewVLSFrom(b)}
}

func WrapJournalEntriesReject(b []byte) JournalEntriesReject {
	return JournalEntriesReject{VLS: message.WrapVLS(b)}
}

func NewJournalEntriesReject() JournalEntriesRejectBuilder {
	a := JournalEntriesRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _JournalEntriesRejectDefault)
	return a
}

func NewJournalEntriesRejectFixedFrom(b []byte) JournalEntriesRejectFixed {
	return JournalEntriesRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapJournalEntriesRejectFixed(b []byte) JournalEntriesRejectFixed {
	return JournalEntriesRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewJournalEntriesRejectFixed() JournalEntriesRejectFixedBuilder {
	a := JournalEntriesRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _JournalEntriesRejectFixedDefault)
	return a
}

func AllocJournalEntriesReject() JournalEntriesRejectPointerBuilder {
	a := JournalEntriesRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _JournalEntriesRejectDefault)
	return a
}

func AllocJournalEntriesRejectFrom(b []byte) JournalEntriesRejectPointer {
	return JournalEntriesRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocJournalEntriesRejectFixed() JournalEntriesRejectFixedPointerBuilder {
	a := JournalEntriesRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _JournalEntriesRejectFixedDefault)
	return a
}

func AllocJournalEntriesRejectFixedFrom(b []byte) JournalEntriesRejectFixedPointer {
	return JournalEntriesRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = JournalEntriesRejectSize  (16)
//     Type       = JOURNAL_ENTRIES_REJECT  (705)
//     BaseSize   = JournalEntriesRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m JournalEntriesRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntriesRejectDefault)
}

// Clear
// {
//     Size       = JournalEntriesRejectFixedSize  (104)
//     Type       = JOURNAL_ENTRIES_REJECT  (705)
//     RequestID  = 0
//     RejectText = ""
// }
func (m JournalEntriesRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntriesRejectFixedDefault)
}

// Clear
// {
//     Size       = JournalEntriesRejectSize  (16)
//     Type       = JOURNAL_ENTRIES_REJECT  (705)
//     BaseSize   = JournalEntriesRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m JournalEntriesRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntriesRejectDefault)
}

// Clear
// {
//     Size       = JournalEntriesRejectFixedSize  (104)
//     Type       = JOURNAL_ENTRIES_REJECT  (705)
//     RequestID  = 0
//     RejectText = ""
// }
func (m JournalEntriesRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntriesRejectFixedDefault)
}

// ToBuilder
func (m JournalEntriesReject) ToBuilder() JournalEntriesRejectBuilder {
	return JournalEntriesRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m JournalEntriesRejectFixed) ToBuilder() JournalEntriesRejectFixedBuilder {
	return JournalEntriesRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m JournalEntriesRejectPointer) ToBuilder() JournalEntriesRejectPointerBuilder {
	return JournalEntriesRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m JournalEntriesRejectFixedPointer) ToBuilder() JournalEntriesRejectFixedPointerBuilder {
	return JournalEntriesRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m JournalEntriesRejectBuilder) Finish() JournalEntriesReject {
	return JournalEntriesReject{m.VLSBuilder.Finish()}
}

// Finish
func (m JournalEntriesRejectFixedBuilder) Finish() JournalEntriesRejectFixed {
	return JournalEntriesRejectFixed{m.Fixed}
}

// Finish
func (m *JournalEntriesRejectPointerBuilder) Finish() JournalEntriesRejectPointer {
	return JournalEntriesRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *JournalEntriesRejectFixedPointerBuilder) Finish() JournalEntriesRejectFixedPointer {
	return JournalEntriesRejectFixedPointer{m.FixedPointer}
}

// RequestID
func (m JournalEntriesReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m JournalEntriesRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m JournalEntriesRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID
func (m JournalEntriesRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText
func (m JournalEntriesReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText
func (m JournalEntriesRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText
func (m JournalEntriesRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText
func (m JournalEntriesRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RequestID
func (m JournalEntriesRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m JournalEntriesRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m JournalEntriesRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m JournalEntriesRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText
func (m JournalEntriesRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m JournalEntriesRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m JournalEntriesRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText
func (m JournalEntriesRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID
func (m JournalEntriesRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID
func (m JournalEntriesRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText
func (m *JournalEntriesRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText
func (m *JournalEntriesRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRequestID
func (m JournalEntriesRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m JournalEntriesRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText
func (m JournalEntriesRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText
func (m JournalEntriesRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m JournalEntriesReject) Copy(to JournalEntriesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m JournalEntriesRejectBuilder) Copy(to JournalEntriesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m JournalEntriesReject) CopyPointer(to JournalEntriesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m JournalEntriesRejectBuilder) CopyPointer(to JournalEntriesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m JournalEntriesReject) CopyToPointer(to JournalEntriesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m JournalEntriesRejectBuilder) CopyToPointer(to JournalEntriesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m JournalEntriesReject) CopyTo(to JournalEntriesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m JournalEntriesRejectBuilder) CopyTo(to JournalEntriesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m JournalEntriesRejectFixed) Copy(to JournalEntriesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m JournalEntriesRejectFixedBuilder) Copy(to JournalEntriesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m JournalEntriesRejectFixed) CopyPointer(to JournalEntriesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m JournalEntriesRejectFixedBuilder) CopyPointer(to JournalEntriesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m JournalEntriesRejectFixed) CopyToPointer(to JournalEntriesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m JournalEntriesRejectFixedBuilder) CopyToPointer(to JournalEntriesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m JournalEntriesRejectFixed) CopyTo(to JournalEntriesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m JournalEntriesRejectFixedBuilder) CopyTo(to JournalEntriesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m JournalEntriesRejectPointer) Copy(to JournalEntriesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m JournalEntriesRejectPointerBuilder) Copy(to JournalEntriesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m JournalEntriesRejectPointer) CopyPointer(to JournalEntriesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m JournalEntriesRejectPointerBuilder) CopyPointer(to JournalEntriesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m JournalEntriesRejectPointer) CopyTo(to JournalEntriesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m JournalEntriesRejectPointerBuilder) CopyTo(to JournalEntriesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m JournalEntriesRejectPointer) CopyToPointer(to JournalEntriesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m JournalEntriesRejectPointerBuilder) CopyToPointer(to JournalEntriesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m JournalEntriesRejectFixedPointer) Copy(to JournalEntriesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m JournalEntriesRejectFixedPointerBuilder) Copy(to JournalEntriesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m JournalEntriesRejectFixedPointer) CopyPointer(to JournalEntriesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m JournalEntriesRejectFixedPointerBuilder) CopyPointer(to JournalEntriesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m JournalEntriesRejectFixedPointer) CopyTo(to JournalEntriesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m JournalEntriesRejectFixedPointerBuilder) CopyTo(to JournalEntriesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m JournalEntriesRejectFixedPointer) CopyToPointer(to JournalEntriesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m JournalEntriesRejectFixedPointerBuilder) CopyToPointer(to JournalEntriesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
