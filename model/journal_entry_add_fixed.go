package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = JournalEntryAddFixedSize  (272)
//     Type         = JOURNAL_ENTRY_ADD  (703)
//     JournalEntry = ""
//     DateTime     = 0
// }
var _JournalEntryAddFixedDefault = []byte{16, 1, 191, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const JournalEntryAddFixedSize = 272

type JournalEntryAddFixed struct {
	message.Fixed
}

type JournalEntryAddFixedBuilder struct {
	message.Fixed
}

func NewJournalEntryAddFixedFrom(b []byte) JournalEntryAddFixed {
	return JournalEntryAddFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapJournalEntryAddFixed(b []byte) JournalEntryAddFixed {
	return JournalEntryAddFixed{Fixed: message.WrapFixed(b)}
}

func NewJournalEntryAddFixed() JournalEntryAddFixedBuilder {
	a := JournalEntryAddFixedBuilder{message.NewFixed(272)}
	a.Unsafe().SetBytes(0, _JournalEntryAddFixedDefault)
	return a
}

// Clear
// {
//     Size         = JournalEntryAddFixedSize  (272)
//     Type         = JOURNAL_ENTRY_ADD  (703)
//     JournalEntry = ""
//     DateTime     = 0
// }
func (m JournalEntryAddFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntryAddFixedDefault)
}

// ToBuilder
func (m JournalEntryAddFixed) ToBuilder() JournalEntryAddFixedBuilder {
	return JournalEntryAddFixedBuilder{m.Fixed}
}

// Finish
func (m JournalEntryAddFixedBuilder) Finish() JournalEntryAddFixed {
	return JournalEntryAddFixed{m.Fixed}
}

// JournalEntry
func (m JournalEntryAddFixed) JournalEntry() string {
	return message.StringFixed(m.Unsafe(), 260, 4)
}

// JournalEntry
func (m JournalEntryAddFixedBuilder) JournalEntry() string {
	return message.StringFixed(m.Unsafe(), 260, 4)
}

// DateTime
func (m JournalEntryAddFixed) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 272, 264))
}

// DateTime
func (m JournalEntryAddFixedBuilder) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 272, 264))
}

// SetJournalEntry
func (m JournalEntryAddFixedBuilder) SetJournalEntry(value string) {
	message.SetStringFixed(m.Unsafe(), 260, 4, value)
}

// SetDateTime
func (m JournalEntryAddFixedBuilder) SetDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 272, 264, int64(value))
}

// Copy
func (m JournalEntryAddFixed) Copy(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m JournalEntryAddFixedBuilder) Copy(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddFixed) CopyPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddFixedBuilder) CopyPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddFixed) CopyToPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddFixedBuilder) CopyToPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddFixed) CopyTo(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddFixedBuilder) CopyTo(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}
