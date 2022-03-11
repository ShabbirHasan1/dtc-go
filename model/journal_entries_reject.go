package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = JournalEntriesRejectSize  (16)
//     Type       = JOURNAL_ENTRIES_REJECT  (705)
//     BaseSize   = JournalEntriesRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _JournalEntriesRejectDefault = []byte{16, 0, 193, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const JournalEntriesRejectSize = 16

type JournalEntriesReject struct {
	message.VLS
}

type JournalEntriesRejectBuilder struct {
	message.VLSBuilder
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

// ToBuilder
func (m JournalEntriesReject) ToBuilder() JournalEntriesRejectBuilder {
	return JournalEntriesRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m JournalEntriesRejectBuilder) Finish() JournalEntriesReject {
	return JournalEntriesReject{m.VLSBuilder.Finish()}
}

// RequestID
func (m JournalEntriesReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m JournalEntriesRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RejectText
func (m JournalEntriesReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText
func (m JournalEntriesRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetRequestID
func (m JournalEntriesRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText
func (m *JournalEntriesRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
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
