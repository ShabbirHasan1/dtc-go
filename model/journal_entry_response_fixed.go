package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size            = JournalEntryResponseFixedSize  (280)
//     Type            = JOURNAL_ENTRY_RESPONSE  (706)
//     JournalEntry    = ""
//     DateTime        = 0
//     IsFinalResponse = 0
// }
var _JournalEntryResponseFixedDefault = []byte{24, 1, 194, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const JournalEntryResponseFixedSize = 280

type JournalEntryResponseFixed struct {
	message.Fixed
}

type JournalEntryResponseFixedBuilder struct {
	message.Fixed
}

func NewJournalEntryResponseFixedFrom(b []byte) JournalEntryResponseFixed {
	return JournalEntryResponseFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapJournalEntryResponseFixed(b []byte) JournalEntryResponseFixed {
	return JournalEntryResponseFixed{Fixed: message.WrapFixed(b)}
}

func NewJournalEntryResponseFixed() JournalEntryResponseFixedBuilder {
	a := JournalEntryResponseFixedBuilder{message.NewFixed(280)}
	a.Unsafe().SetBytes(0, _JournalEntryResponseFixedDefault)
	return a
}

// Clear
// {
//     Size            = JournalEntryResponseFixedSize  (280)
//     Type            = JOURNAL_ENTRY_RESPONSE  (706)
//     JournalEntry    = ""
//     DateTime        = 0
//     IsFinalResponse = 0
// }
func (m JournalEntryResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntryResponseFixedDefault)
}

// ToBuilder
func (m JournalEntryResponseFixed) ToBuilder() JournalEntryResponseFixedBuilder {
	return JournalEntryResponseFixedBuilder{m.Fixed}
}

// Finish
func (m JournalEntryResponseFixedBuilder) Finish() JournalEntryResponseFixed {
	return JournalEntryResponseFixed{m.Fixed}
}

// JournalEntry
func (m JournalEntryResponseFixed) JournalEntry() string {
	return message.StringFixed(m.Unsafe(), 260, 4)
}

// JournalEntry
func (m JournalEntryResponseFixedBuilder) JournalEntry() string {
	return message.StringFixed(m.Unsafe(), 260, 4)
}

// DateTime
func (m JournalEntryResponseFixed) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 272, 264))
}

// DateTime
func (m JournalEntryResponseFixedBuilder) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 272, 264))
}

// IsFinalResponse
func (m JournalEntryResponseFixed) IsFinalResponse() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 273, 272)
}

// IsFinalResponse
func (m JournalEntryResponseFixedBuilder) IsFinalResponse() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 273, 272)
}

// SetJournalEntry
func (m JournalEntryResponseFixedBuilder) SetJournalEntry(value string) {
	message.SetStringFixed(m.Unsafe(), 260, 4, value)
}

// SetDateTime
func (m JournalEntryResponseFixedBuilder) SetDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 272, 264, int64(value))
}

// SetIsFinalResponse
func (m JournalEntryResponseFixedBuilder) SetIsFinalResponse(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 273, 272, value)
}

// Copy
func (m JournalEntryResponseFixed) Copy(to JournalEntryResponseFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// Copy
func (m JournalEntryResponseFixedBuilder) Copy(to JournalEntryResponseFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyPointer
func (m JournalEntryResponseFixed) CopyPointer(to JournalEntryResponseFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyPointer
func (m JournalEntryResponseFixedBuilder) CopyPointer(to JournalEntryResponseFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyToPointer
func (m JournalEntryResponseFixed) CopyToPointer(to JournalEntryResponsePointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyToPointer
func (m JournalEntryResponseFixedBuilder) CopyToPointer(to JournalEntryResponsePointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyTo
func (m JournalEntryResponseFixed) CopyTo(to JournalEntryResponseBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyTo
func (m JournalEntryResponseFixedBuilder) CopyTo(to JournalEntryResponseBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}
