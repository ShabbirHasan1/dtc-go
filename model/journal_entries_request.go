package model

import (
	"github.com/moontrade/dtc-go/message"
)

const JournalEntriesRequestSize = 16

// {
//     Size          = JournalEntriesRequestSize  (16)
//     Type          = JOURNAL_ENTRIES_REQUEST  (704)
//     RequestID     = 0
//     StartDateTime = 0
// }
var _JournalEntriesRequestDefault = []byte{16, 0, 192, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type JournalEntriesRequest struct {
	message.Fixed
}

type JournalEntriesRequestBuilder struct {
	message.Fixed
}

type JournalEntriesRequestPointer struct {
	message.FixedPointer
}

type JournalEntriesRequestPointerBuilder struct {
	message.FixedPointer
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
func (m JournalEntriesRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntriesRequestDefault)
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
func (m JournalEntriesRequest) ToBuilder() JournalEntriesRequestBuilder {
	return JournalEntriesRequestBuilder{m.Fixed}
}

// ToBuilder
func (m JournalEntriesRequestPointer) ToBuilder() JournalEntriesRequestPointerBuilder {
	return JournalEntriesRequestPointerBuilder{m.FixedPointer}
}

// Finish
func (m JournalEntriesRequestBuilder) Finish() JournalEntriesRequest {
	return JournalEntriesRequest{m.Fixed}
}

// Finish
func (m *JournalEntriesRequestPointerBuilder) Finish() JournalEntriesRequestPointer {
	return JournalEntriesRequestPointer{m.FixedPointer}
}

// RequestID
func (m JournalEntriesRequest) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m JournalEntriesRequestBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m JournalEntriesRequest) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// StartDateTime
func (m JournalEntriesRequestBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 16, 8))
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
func (m JournalEntriesRequestBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m JournalEntriesRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetStartDateTime
func (m JournalEntriesRequestBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, int64(value))
}

// SetStartDateTime
func (m JournalEntriesRequestPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
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
