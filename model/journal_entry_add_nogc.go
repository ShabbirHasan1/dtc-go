package model

import (
	"github.com/moontrade/dtc-go/message"
)

type JournalEntryAddPointer struct {
	message.VLSPointer
}

type JournalEntryAddPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocJournalEntryAdd() JournalEntryAddPointerBuilder {
	a := JournalEntryAddPointerBuilder{message.AllocVLSBuilder(24)}
	a.Ptr.SetBytes(0, _JournalEntryAddDefault)
	return a
}

func AllocJournalEntryAddFrom(b []byte) JournalEntryAddPointer {
	return JournalEntryAddPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = JournalEntryAddSize  (24)
//     Type         = JOURNAL_ENTRY_ADD  (703)
//     BaseSize     = JournalEntryAddSize  (24)
//     JournalEntry = ""
//     DateTime     = 0
// }
func (m JournalEntryAddPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntryAddDefault)
}

// ToBuilder
func (m JournalEntryAddPointer) ToBuilder() JournalEntryAddPointerBuilder {
	return JournalEntryAddPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *JournalEntryAddPointerBuilder) Finish() JournalEntryAddPointer {
	return JournalEntryAddPointer{m.VLSPointerBuilder.Finish()}
}

// JournalEntry
func (m JournalEntryAddPointer) JournalEntry() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// JournalEntry
func (m JournalEntryAddPointerBuilder) JournalEntry() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// DateTime
func (m JournalEntryAddPointer) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 24, 16))
}

// DateTime
func (m JournalEntryAddPointerBuilder) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 24, 16))
}

// SetJournalEntry
func (m *JournalEntryAddPointerBuilder) SetJournalEntry(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetDateTime
func (m JournalEntryAddPointerBuilder) SetDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 24, 16, int64(value))
}

// Copy
func (m JournalEntryAddPointer) Copy(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m JournalEntryAddPointerBuilder) Copy(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddPointer) CopyPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddPointerBuilder) CopyPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddPointer) CopyTo(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddPointerBuilder) CopyTo(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddPointer) CopyToPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddPointerBuilder) CopyToPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}
