package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = JournalEntriesRequestSize  (16)
//     Type          = JOURNAL_ENTRIES_REQUEST  (704)
//     RequestID     = 0
//     StartDateTime = 0
// }
var _JournalEntriesRequestDefault = []byte{16, 0, 192, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const JournalEntriesRequestSize = 16

type JournalEntriesRequest struct {
	message.Fixed
}

type JournalEntriesRequestBuilder struct {
	message.Fixed
}

func NewJournalEntriesRequestFrom(b []byte) JournalEntriesRequest {
	return JournalEntriesRequest{Fixed: message.NewFixedFrom(b)}
}

func WrapJournalEntriesRequest(b []byte) JournalEntriesRequest {
	return JournalEntriesRequest{Fixed: message.WrapFixed(b)}
}

func NewJournalEntriesRequest() JournalEntriesRequestBuilder {
	a := JournalEntriesRequestBuilder{message.NewFixed(16)}
	a.Unsafe().SetBytes(0, _JournalEntriesRequestDefault)
	return a
}

// Clear
// {
//     Size          = JournalEntriesRequestSize  (16)
//     Type          = JOURNAL_ENTRIES_REQUEST  (704)
//     RequestID     = 0
//     StartDateTime = 0
// }
func (m JournalEntriesRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntriesRequestDefault)
}

// ToBuilder
func (m JournalEntriesRequest) ToBuilder() JournalEntriesRequestBuilder {
	return JournalEntriesRequestBuilder{m.Fixed}
}

// Finish
func (m JournalEntriesRequestBuilder) Finish() JournalEntriesRequest {
	return JournalEntriesRequest{m.Fixed}
}

// RequestID
func (m JournalEntriesRequest) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m JournalEntriesRequestBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// StartDateTime
func (m JournalEntriesRequest) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// StartDateTime
func (m JournalEntriesRequestBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// SetRequestID
func (m JournalEntriesRequestBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetStartDateTime
func (m JournalEntriesRequestBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, int64(value))
}

// Copy
func (m JournalEntriesRequest) Copy(to JournalEntriesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m JournalEntriesRequestBuilder) Copy(to JournalEntriesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m JournalEntriesRequest) CopyPointer(to JournalEntriesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m JournalEntriesRequestBuilder) CopyPointer(to JournalEntriesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
}
