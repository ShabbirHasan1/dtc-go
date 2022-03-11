package model

import (
	"github.com/moontrade/dtc-go/message"
)

type JournalEntryAddFixedPointer struct {
	message.FixedPointer
}

type JournalEntryAddFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocJournalEntryAddFixed() JournalEntryAddFixedPointerBuilder {
	a := JournalEntryAddFixedPointerBuilder{message.AllocFixed(272)}
	a.Ptr.SetBytes(0, _JournalEntryAddFixedDefault)
	return a
}

func AllocJournalEntryAddFixedFrom(b []byte) JournalEntryAddFixedPointer {
	return JournalEntryAddFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size         = JournalEntryAddFixedSize  (272)
//     Type         = JOURNAL_ENTRY_ADD  (703)
//     JournalEntry = ""
//     DateTime     = 0
// }
func (m JournalEntryAddFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntryAddFixedDefault)
}

// ToBuilder
func (m JournalEntryAddFixedPointer) ToBuilder() JournalEntryAddFixedPointerBuilder {
	return JournalEntryAddFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *JournalEntryAddFixedPointerBuilder) Finish() JournalEntryAddFixedPointer {
	return JournalEntryAddFixedPointer{m.FixedPointer}
}

// JournalEntry
func (m JournalEntryAddFixedPointer) JournalEntry() string {
	return message.StringFixed(m.Ptr, 260, 4)
}

// JournalEntry
func (m JournalEntryAddFixedPointerBuilder) JournalEntry() string {
	return message.StringFixed(m.Ptr, 260, 4)
}

// DateTime
func (m JournalEntryAddFixedPointer) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 272, 264))
}

// DateTime
func (m JournalEntryAddFixedPointerBuilder) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 272, 264))
}

// SetJournalEntry
func (m JournalEntryAddFixedPointerBuilder) SetJournalEntry(value string) {
	message.SetStringFixed(m.Ptr, 260, 4, value)
}

// SetDateTime
func (m JournalEntryAddFixedPointerBuilder) SetDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 272, 264, int64(value))
}

// Copy
func (m JournalEntryAddFixedPointer) Copy(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m JournalEntryAddFixedPointerBuilder) Copy(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddFixedPointer) CopyPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddFixedPointerBuilder) CopyPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddFixedPointer) CopyTo(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddFixedPointerBuilder) CopyTo(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddFixedPointer) CopyToPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddFixedPointerBuilder) CopyToPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}
