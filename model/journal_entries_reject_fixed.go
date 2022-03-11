package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = JournalEntriesRejectFixedSize  (104)
//     Type       = JOURNAL_ENTRIES_REJECT  (705)
//     RequestID  = 0
//     RejectText = ""
// }
var _JournalEntriesRejectFixedDefault = []byte{104, 0, 193, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const JournalEntriesRejectFixedSize = 104

type JournalEntriesRejectFixed struct {
	message.Fixed
}

type JournalEntriesRejectFixedBuilder struct {
	message.Fixed
}

func NewJournalEntriesRejectFixedFrom(b []byte) JournalEntriesRejectFixed {
	return JournalEntriesRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapJournalEntriesRejectFixed(b []byte) JournalEntriesRejectFixed {
	return JournalEntriesRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewJournalEntriesRejectFixed() JournalEntriesRejectFixedBuilder {
	a := JournalEntriesRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _JournalEntriesRejectFixedDefault)
	return a
}

// Clear
// {
//     Size       = JournalEntriesRejectFixedSize  (104)
//     Type       = JOURNAL_ENTRIES_REJECT  (705)
//     RequestID  = 0
//     RejectText = ""
// }
func (m JournalEntriesRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntriesRejectFixedDefault)
}

// ToBuilder
func (m JournalEntriesRejectFixed) ToBuilder() JournalEntriesRejectFixedBuilder {
	return JournalEntriesRejectFixedBuilder{m.Fixed}
}

// Finish
func (m JournalEntriesRejectFixedBuilder) Finish() JournalEntriesRejectFixed {
	return JournalEntriesRejectFixed{m.Fixed}
}

// RequestID
func (m JournalEntriesRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m JournalEntriesRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RejectText
func (m JournalEntriesRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m JournalEntriesRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// SetRequestID
func (m JournalEntriesRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText
func (m JournalEntriesRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// Copy
func (m JournalEntriesRejectFixed) Copy(to JournalEntriesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m JournalEntriesRejectFixedBuilder) Copy(to JournalEntriesRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m JournalEntriesRejectFixed) CopyPointer(to JournalEntriesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m JournalEntriesRejectFixedBuilder) CopyPointer(to JournalEntriesRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m JournalEntriesRejectFixed) CopyToPointer(to JournalEntriesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m JournalEntriesRejectFixedBuilder) CopyToPointer(to JournalEntriesRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m JournalEntriesRejectFixed) CopyTo(to JournalEntriesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m JournalEntriesRejectFixedBuilder) CopyTo(to JournalEntriesRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
