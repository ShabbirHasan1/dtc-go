package model

import (
	"github.com/moontrade/dtc-go/message"
)

type JournalEntriesRejectFixedPointer struct {
	message.FixedPointer
}

type JournalEntriesRejectFixedPointerBuilder struct {
	message.FixedPointer
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
//     Size       = JournalEntriesRejectFixedSize  (104)
//     Type       = JOURNAL_ENTRIES_REJECT  (705)
//     RequestID  = 0
//     RejectText = ""
// }
func (m JournalEntriesRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntriesRejectFixedDefault)
}

// ToBuilder
func (m JournalEntriesRejectFixedPointer) ToBuilder() JournalEntriesRejectFixedPointerBuilder {
	return JournalEntriesRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *JournalEntriesRejectFixedPointerBuilder) Finish() JournalEntriesRejectFixedPointer {
	return JournalEntriesRejectFixedPointer{m.FixedPointer}
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
func (m JournalEntriesRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText
func (m JournalEntriesRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID
func (m JournalEntriesRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText
func (m JournalEntriesRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
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
