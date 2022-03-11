package model

import (
	"github.com/moontrade/dtc-go/message"
)

type JournalEntryResponseFixedPointer struct {
	message.FixedPointer
}

type JournalEntryResponseFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocJournalEntryResponseFixed() JournalEntryResponseFixedPointerBuilder {
	a := JournalEntryResponseFixedPointerBuilder{message.AllocFixed(280)}
	a.Ptr.SetBytes(0, _JournalEntryResponseFixedDefault)
	return a
}

func AllocJournalEntryResponseFixedFrom(b []byte) JournalEntryResponseFixedPointer {
	return JournalEntryResponseFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size            = JournalEntryResponseFixedSize  (280)
//     Type            = JOURNAL_ENTRY_RESPONSE  (706)
//     JournalEntry    = ""
//     DateTime        = 0
//     IsFinalResponse = 0
// }
func (m JournalEntryResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntryResponseFixedDefault)
}

// ToBuilder
func (m JournalEntryResponseFixedPointer) ToBuilder() JournalEntryResponseFixedPointerBuilder {
	return JournalEntryResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *JournalEntryResponseFixedPointerBuilder) Finish() JournalEntryResponseFixedPointer {
	return JournalEntryResponseFixedPointer{m.FixedPointer}
}

// JournalEntry
func (m JournalEntryResponseFixedPointer) JournalEntry() string {
	return message.StringFixed(m.Ptr, 260, 4)
}

// JournalEntry
func (m JournalEntryResponseFixedPointerBuilder) JournalEntry() string {
	return message.StringFixed(m.Ptr, 260, 4)
}

// DateTime
func (m JournalEntryResponseFixedPointer) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 272, 264))
}

// DateTime
func (m JournalEntryResponseFixedPointerBuilder) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 272, 264))
}

// IsFinalResponse
func (m JournalEntryResponseFixedPointer) IsFinalResponse() uint8 {
	return message.Uint8Fixed(m.Ptr, 273, 272)
}

// IsFinalResponse
func (m JournalEntryResponseFixedPointerBuilder) IsFinalResponse() uint8 {
	return message.Uint8Fixed(m.Ptr, 273, 272)
}

// SetJournalEntry
func (m JournalEntryResponseFixedPointerBuilder) SetJournalEntry(value string) {
	message.SetStringFixed(m.Ptr, 260, 4, value)
}

// SetDateTime
func (m JournalEntryResponseFixedPointerBuilder) SetDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 272, 264, int64(value))
}

// SetIsFinalResponse
func (m JournalEntryResponseFixedPointerBuilder) SetIsFinalResponse(value uint8) {
	message.SetUint8Fixed(m.Ptr, 273, 272, value)
}

// Copy
func (m JournalEntryResponseFixedPointer) Copy(to JournalEntryResponseFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// Copy
func (m JournalEntryResponseFixedPointerBuilder) Copy(to JournalEntryResponseFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyPointer
func (m JournalEntryResponseFixedPointer) CopyPointer(to JournalEntryResponseFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyPointer
func (m JournalEntryResponseFixedPointerBuilder) CopyPointer(to JournalEntryResponseFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyTo
func (m JournalEntryResponseFixedPointer) CopyTo(to JournalEntryResponseBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyTo
func (m JournalEntryResponseFixedPointerBuilder) CopyTo(to JournalEntryResponseBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyToPointer
func (m JournalEntryResponseFixedPointer) CopyToPointer(to JournalEntryResponsePointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyToPointer
func (m JournalEntryResponseFixedPointerBuilder) CopyToPointer(to JournalEntryResponsePointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}
