package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = JournalEntryAddSize  (24)
//     Type         = JOURNAL_ENTRY_ADD  (703)
//     BaseSize     = JournalEntryAddSize  (24)
//     JournalEntry = ""
//     DateTime     = 0
// }
var _JournalEntryAddDefault = []byte{24, 0, 191, 2, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const JournalEntryAddSize = 24

type JournalEntryAdd struct {
	message.VLS
}

type JournalEntryAddBuilder struct {
	message.VLSBuilder
}

func NewJournalEntryAddFrom(b []byte) JournalEntryAdd {
	return JournalEntryAdd{VLS: message.NewVLSFrom(b)}
}

func WrapJournalEntryAdd(b []byte) JournalEntryAdd {
	return JournalEntryAdd{VLS: message.WrapVLS(b)}
}

func NewJournalEntryAdd() JournalEntryAddBuilder {
	a := JournalEntryAddBuilder{message.NewVLSBuilder(24)}
	a.Unsafe().SetBytes(0, _JournalEntryAddDefault)
	return a
}

// Clear
// {
//     Size         = JournalEntryAddSize  (24)
//     Type         = JOURNAL_ENTRY_ADD  (703)
//     BaseSize     = JournalEntryAddSize  (24)
//     JournalEntry = ""
//     DateTime     = 0
// }
func (m JournalEntryAddBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntryAddDefault)
}

// ToBuilder
func (m JournalEntryAdd) ToBuilder() JournalEntryAddBuilder {
	return JournalEntryAddBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m JournalEntryAddBuilder) Finish() JournalEntryAdd {
	return JournalEntryAdd{m.VLSBuilder.Finish()}
}

// JournalEntry
func (m JournalEntryAdd) JournalEntry() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// JournalEntry
func (m JournalEntryAddBuilder) JournalEntry() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// DateTime
func (m JournalEntryAdd) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 24, 16))
}

// DateTime
func (m JournalEntryAddBuilder) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 24, 16))
}

// SetJournalEntry
func (m *JournalEntryAddBuilder) SetJournalEntry(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetDateTime
func (m JournalEntryAddBuilder) SetDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 24, 16, int64(value))
}

// Copy
func (m JournalEntryAdd) Copy(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m JournalEntryAddBuilder) Copy(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAdd) CopyPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddBuilder) CopyPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAdd) CopyToPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddBuilder) CopyToPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAdd) CopyTo(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddBuilder) CopyTo(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}
