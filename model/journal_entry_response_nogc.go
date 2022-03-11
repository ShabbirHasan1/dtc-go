package model

import (
	"github.com/moontrade/dtc-go/message"
)

type JournalEntryResponsePointer struct {
	message.VLSPointer
}

type JournalEntryResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocJournalEntryResponse() JournalEntryResponsePointerBuilder {
	a := JournalEntryResponsePointerBuilder{message.AllocVLSBuilder(32)}
	a.Ptr.SetBytes(0, _JournalEntryResponseDefault)
	return a
}

func AllocJournalEntryResponseFrom(b []byte) JournalEntryResponsePointer {
	return JournalEntryResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size            = JournalEntryResponseSize  (32)
//     Type            = JOURNAL_ENTRY_RESPONSE  (706)
//     BaseSize        = JournalEntryResponseSize  (32)
//     JournalEntry    = ""
//     DateTime        = 0
//     IsFinalResponse = 0
// }
func (m JournalEntryResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntryResponseDefault)
}

// ToBuilder
func (m JournalEntryResponsePointer) ToBuilder() JournalEntryResponsePointerBuilder {
	return JournalEntryResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *JournalEntryResponsePointerBuilder) Finish() JournalEntryResponsePointer {
	return JournalEntryResponsePointer{m.VLSPointerBuilder.Finish()}
}

// JournalEntry
func (m JournalEntryResponsePointer) JournalEntry() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// JournalEntry
func (m JournalEntryResponsePointerBuilder) JournalEntry() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// DateTime
func (m JournalEntryResponsePointer) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 24, 16))
}

// DateTime
func (m JournalEntryResponsePointerBuilder) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 24, 16))
}

// IsFinalResponse
func (m JournalEntryResponsePointer) IsFinalResponse() uint8 {
	return message.Uint8VLS(m.Ptr, 25, 24)
}

// IsFinalResponse
func (m JournalEntryResponsePointerBuilder) IsFinalResponse() uint8 {
	return message.Uint8VLS(m.Ptr, 25, 24)
}

// SetJournalEntry
func (m *JournalEntryResponsePointerBuilder) SetJournalEntry(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetDateTime
func (m JournalEntryResponsePointerBuilder) SetDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 24, 16, int64(value))
}

// SetIsFinalResponse
func (m JournalEntryResponsePointerBuilder) SetIsFinalResponse(value uint8) {
	message.SetUint8VLS(m.Ptr, 25, 24, value)
}

// Copy
func (m JournalEntryResponsePointer) Copy(to JournalEntryResponseBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// Copy
func (m JournalEntryResponsePointerBuilder) Copy(to JournalEntryResponseBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyPointer
func (m JournalEntryResponsePointer) CopyPointer(to JournalEntryResponsePointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyPointer
func (m JournalEntryResponsePointerBuilder) CopyPointer(to JournalEntryResponsePointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyTo
func (m JournalEntryResponsePointer) CopyTo(to JournalEntryResponseFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyTo
func (m JournalEntryResponsePointerBuilder) CopyTo(to JournalEntryResponseFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyToPointer
func (m JournalEntryResponsePointer) CopyToPointer(to JournalEntryResponseFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyToPointer
func (m JournalEntryResponsePointerBuilder) CopyToPointer(to JournalEntryResponseFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}
