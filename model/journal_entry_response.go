package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size            = JournalEntryResponseSize  (32)
//     Type            = JOURNAL_ENTRY_RESPONSE  (706)
//     BaseSize        = JournalEntryResponseSize  (32)
//     JournalEntry    = ""
//     DateTime        = 0
//     IsFinalResponse = 0
// }
var _JournalEntryResponseDefault = []byte{32, 0, 194, 2, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const JournalEntryResponseSize = 32

type JournalEntryResponse struct {
	message.VLS
}

type JournalEntryResponseBuilder struct {
	message.VLSBuilder
}

func NewJournalEntryResponseFrom(b []byte) JournalEntryResponse {
	return JournalEntryResponse{VLS: message.NewVLSFrom(b)}
}

func WrapJournalEntryResponse(b []byte) JournalEntryResponse {
	return JournalEntryResponse{VLS: message.WrapVLS(b)}
}

func NewJournalEntryResponse() JournalEntryResponseBuilder {
	a := JournalEntryResponseBuilder{message.NewVLSBuilder(32)}
	a.Unsafe().SetBytes(0, _JournalEntryResponseDefault)
	return a
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
func (m JournalEntryResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntryResponseDefault)
}

// ToBuilder
func (m JournalEntryResponse) ToBuilder() JournalEntryResponseBuilder {
	return JournalEntryResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m JournalEntryResponseBuilder) Finish() JournalEntryResponse {
	return JournalEntryResponse{m.VLSBuilder.Finish()}
}

// JournalEntry
func (m JournalEntryResponse) JournalEntry() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// JournalEntry
func (m JournalEntryResponseBuilder) JournalEntry() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// DateTime
func (m JournalEntryResponse) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 24, 16))
}

// DateTime
func (m JournalEntryResponseBuilder) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 24, 16))
}

// IsFinalResponse
func (m JournalEntryResponse) IsFinalResponse() uint8 {
	return message.Uint8VLS(m.Unsafe(), 25, 24)
}

// IsFinalResponse
func (m JournalEntryResponseBuilder) IsFinalResponse() uint8 {
	return message.Uint8VLS(m.Unsafe(), 25, 24)
}

// SetJournalEntry
func (m *JournalEntryResponseBuilder) SetJournalEntry(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetDateTime
func (m JournalEntryResponseBuilder) SetDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 24, 16, int64(value))
}

// SetIsFinalResponse
func (m JournalEntryResponseBuilder) SetIsFinalResponse(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 25, 24, value)
}

// Copy
func (m JournalEntryResponse) Copy(to JournalEntryResponseBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// Copy
func (m JournalEntryResponseBuilder) Copy(to JournalEntryResponseBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyPointer
func (m JournalEntryResponse) CopyPointer(to JournalEntryResponsePointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyPointer
func (m JournalEntryResponseBuilder) CopyPointer(to JournalEntryResponsePointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyToPointer
func (m JournalEntryResponse) CopyToPointer(to JournalEntryResponseFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyToPointer
func (m JournalEntryResponseBuilder) CopyToPointer(to JournalEntryResponseFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyTo
func (m JournalEntryResponse) CopyTo(to JournalEntryResponseFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyTo
func (m JournalEntryResponseBuilder) CopyTo(to JournalEntryResponseFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}
