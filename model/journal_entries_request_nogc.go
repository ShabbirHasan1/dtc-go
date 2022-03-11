package model

import (
	"github.com/moontrade/dtc-go/message"
)

type JournalEntriesRequestPointer struct {
	message.FixedPointer
}

type JournalEntriesRequestPointerBuilder struct {
	message.FixedPointer
}

func AllocJournalEntriesRequest() JournalEntriesRequestPointerBuilder {
	a := JournalEntriesRequestPointerBuilder{message.AllocFixed(16)}
	a.Ptr.SetBytes(0, _JournalEntriesRequestDefault)
	return a
}

func AllocJournalEntriesRequestFrom(b []byte) JournalEntriesRequestPointer {
	return JournalEntriesRequestPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = JournalEntriesRequestSize  (16)
//     Type          = JOURNAL_ENTRIES_REQUEST  (704)
//     RequestID     = 0
//     StartDateTime = 0
// }
func (m JournalEntriesRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntriesRequestDefault)
}

// ToBuilder
func (m JournalEntriesRequestPointer) ToBuilder() JournalEntriesRequestPointerBuilder {
	return JournalEntriesRequestPointerBuilder{m.FixedPointer}
}

// Finish
func (m *JournalEntriesRequestPointerBuilder) Finish() JournalEntriesRequestPointer {
	return JournalEntriesRequestPointer{m.FixedPointer}
}

// RequestID
func (m JournalEntriesRequestPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m JournalEntriesRequestPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// StartDateTime
func (m JournalEntriesRequestPointer) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 16, 8))
}

// StartDateTime
func (m JournalEntriesRequestPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 16, 8))
}

// SetRequestID
func (m JournalEntriesRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetStartDateTime
func (m JournalEntriesRequestPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
}

// Copy
func (m JournalEntriesRequestPointer) Copy(to JournalEntriesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m JournalEntriesRequestPointerBuilder) Copy(to JournalEntriesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m JournalEntriesRequestPointer) CopyPointer(to JournalEntriesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m JournalEntriesRequestPointerBuilder) CopyPointer(to JournalEntriesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
}
