package model

import (
	"github.com/moontrade/dtc-go/message"
)

type JournalEntriesRejectPointer struct {
	message.VLSPointer
}

type JournalEntriesRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocJournalEntriesReject() JournalEntriesRejectPointerBuilder {
	a := JournalEntriesRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _JournalEntriesRejectDefault)
	return a
}

func AllocJournalEntriesRejectFrom(b []byte) JournalEntriesRejectPointer {
	return JournalEntriesRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
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

// ToBuilder
func (m JournalEntriesRejectPointer) ToBuilder() JournalEntriesRejectPointerBuilder {
	return JournalEntriesRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *JournalEntriesRejectPointerBuilder) Finish() JournalEntriesRejectPointer {
	return JournalEntriesRejectPointer{m.VLSPointerBuilder.Finish()}
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
func (m JournalEntriesRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText
func (m JournalEntriesRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetRequestID
func (m JournalEntriesRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText
func (m *JournalEntriesRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
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
