// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-16 08:00:54.519198 +0800 WITA m=+0.055446543

package v8

import (
	"github.com/moontrade/dtc-go/message"
)

const JournalEntryAddSize = 24

const JournalEntryAddFixedSize = 272

//     Size          uint16    = JournalEntryAddSize  (24)
//     Type          uint16    = JOURNAL_ENTRY_ADD  (703)
//     BaseSize      uint16    = JournalEntryAddSize  (24)
//     JournalEntry  string    = ""
//     DateTime      DateTime  = 0
var _JournalEntryAddDefault = []byte{24, 0, 191, 2, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size          uint16       = JournalEntryAddFixedSize  (272)
//     Type          uint16       = JOURNAL_ENTRY_ADD  (703)
//     JournalEntry  string[256]  = ""
//     DateTime      DateTime     = 0
var _JournalEntryAddFixedDefault = []byte{16, 1, 191, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type JournalEntryAdd struct {
	message.VLS
}

type JournalEntryAddBuilder struct {
	message.VLSBuilder
}

type JournalEntryAddFixed struct {
	message.Fixed
}

type JournalEntryAddFixedBuilder struct {
	message.Fixed
}

type JournalEntryAddPointer struct {
	message.VLSPointer
}

type JournalEntryAddPointerBuilder struct {
	message.VLSPointerBuilder
}

type JournalEntryAddFixedPointer struct {
	message.FixedPointer
}

type JournalEntryAddFixedPointerBuilder struct {
	message.FixedPointer
}

func NewJournalEntryAddFrom(b []byte) JournalEntryAdd {
	return JournalEntryAdd{VLS: message.NewVLS(b)}
}

func WrapJournalEntryAdd(b []byte) JournalEntryAdd {
	return JournalEntryAdd{VLS: message.WrapVLS(b)}
}

func NewJournalEntryAdd() JournalEntryAddBuilder {
	return JournalEntryAddBuilder{message.NewVLSBuilder(_JournalEntryAddDefault)}
}

func NewJournalEntryAddFixedFrom(b []byte) JournalEntryAddFixed {
	return JournalEntryAddFixed{Fixed: message.NewFixed(b)}
}

func WrapJournalEntryAddFixed(b []byte) JournalEntryAddFixed {
	return JournalEntryAddFixed{Fixed: message.WrapFixed(b)}
}

func NewJournalEntryAddFixed() JournalEntryAddFixedBuilder {
	return JournalEntryAddFixedBuilder{message.NewFixed(_JournalEntryAddFixedDefault)}
}

func AllocJournalEntryAdd() JournalEntryAddPointerBuilder {
	return JournalEntryAddPointerBuilder{message.AllocVLSBuilder(_JournalEntryAddDefault)}
}

func AllocJournalEntryAddFrom(b []byte) JournalEntryAddPointer {
	return JournalEntryAddPointer{VLSPointer: message.AllocVLS(b)}
}

func AllocJournalEntryAddFixed() JournalEntryAddFixedPointerBuilder {
	return JournalEntryAddFixedPointerBuilder{message.AllocFixed(_JournalEntryAddFixedDefault)}
}

func AllocJournalEntryAddFixedFrom(b []byte) JournalEntryAddFixedPointer {
	return JournalEntryAddFixedPointer{FixedPointer: message.AllocFixed(b)}
}

// Clear
//     Size          uint16    = JournalEntryAddSize  (24)
//     Type          uint16    = JOURNAL_ENTRY_ADD  (703)
//     BaseSize      uint16    = JournalEntryAddSize  (24)
//     JournalEntry  string    = ""
//     DateTime      DateTime  = 0
func (m JournalEntryAddBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntryAddDefault)
}

// Clear
//     Size          uint16       = JournalEntryAddFixedSize  (272)
//     Type          uint16       = JOURNAL_ENTRY_ADD  (703)
//     JournalEntry  string[256]  = ""
//     DateTime      DateTime     = 0
func (m JournalEntryAddFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntryAddFixedDefault)
}

// Clear
//     Size          uint16    = JournalEntryAddSize  (24)
//     Type          uint16    = JOURNAL_ENTRY_ADD  (703)
//     BaseSize      uint16    = JournalEntryAddSize  (24)
//     JournalEntry  string    = ""
//     DateTime      DateTime  = 0
func (m JournalEntryAddPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntryAddDefault)
}

// Clear
//     Size          uint16       = JournalEntryAddFixedSize  (272)
//     Type          uint16       = JOURNAL_ENTRY_ADD  (703)
//     JournalEntry  string[256]  = ""
//     DateTime      DateTime     = 0
func (m JournalEntryAddFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntryAddFixedDefault)
}

// ToBuilder
func (m JournalEntryAdd) ToBuilder() JournalEntryAddBuilder {
	return JournalEntryAddBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m JournalEntryAddFixed) ToBuilder() JournalEntryAddFixedBuilder {
	return JournalEntryAddFixedBuilder{m.Fixed}
}

// ToBuilder
func (m JournalEntryAddPointer) ToBuilder() JournalEntryAddPointerBuilder {
	return JournalEntryAddPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m JournalEntryAddFixedPointer) ToBuilder() JournalEntryAddFixedPointerBuilder {
	return JournalEntryAddFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m JournalEntryAddBuilder) Finish() JournalEntryAdd {
	return JournalEntryAdd{m.VLSBuilder.Finish()}
}

// Finish
func (m JournalEntryAddFixedBuilder) Finish() JournalEntryAddFixed {
	return JournalEntryAddFixed{m.Fixed}
}

// Finish
func (m *JournalEntryAddPointerBuilder) Finish() JournalEntryAddPointer {
	return JournalEntryAddPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *JournalEntryAddFixedPointerBuilder) Finish() JournalEntryAddFixedPointer {
	return JournalEntryAddFixedPointer{m.FixedPointer}
}

// JournalEntry
func (m JournalEntryAdd) JournalEntry() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// JournalEntry
func (m JournalEntryAddBuilder) JournalEntry() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// JournalEntry
func (m JournalEntryAddPointer) JournalEntry() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// JournalEntry
func (m JournalEntryAddPointerBuilder) JournalEntry() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// DateTime
func (m JournalEntryAdd) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 24, 16))
}

// DateTime
func (m JournalEntryAddBuilder) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 24, 16))
}

// DateTime
func (m JournalEntryAddPointer) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 24, 16))
}

// DateTime
func (m JournalEntryAddPointerBuilder) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 24, 16))
}

// JournalEntry
func (m JournalEntryAddFixed) JournalEntry() string {
	return message.StringFixed(m.Unsafe(), 260, 4)
}

// JournalEntry
func (m JournalEntryAddFixedBuilder) JournalEntry() string {
	return message.StringFixed(m.Unsafe(), 260, 4)
}

// JournalEntry
func (m JournalEntryAddFixedPointer) JournalEntry() string {
	return message.StringFixed(m.Ptr, 260, 4)
}

// JournalEntry
func (m JournalEntryAddFixedPointerBuilder) JournalEntry() string {
	return message.StringFixed(m.Ptr, 260, 4)
}

// DateTime
func (m JournalEntryAddFixed) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 272, 264))
}

// DateTime
func (m JournalEntryAddFixedBuilder) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 272, 264))
}

// DateTime
func (m JournalEntryAddFixedPointer) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 272, 264))
}

// DateTime
func (m JournalEntryAddFixedPointerBuilder) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 272, 264))
}

// SetJournalEntry
func (m *JournalEntryAddBuilder) SetJournalEntry(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetJournalEntry
func (m *JournalEntryAddPointerBuilder) SetJournalEntry(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetDateTime
func (m JournalEntryAddBuilder) SetDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 24, 16, int64(value))
}

// SetDateTime
func (m JournalEntryAddPointerBuilder) SetDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 24, 16, int64(value))
}

// SetJournalEntry
func (m JournalEntryAddFixedBuilder) SetJournalEntry(value string) {
	message.SetStringFixed(m.Unsafe(), 260, 4, value)
}

// SetJournalEntry
func (m JournalEntryAddFixedPointerBuilder) SetJournalEntry(value string) {
	message.SetStringFixed(m.Ptr, 260, 4, value)
}

// SetDateTime
func (m JournalEntryAddFixedBuilder) SetDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 272, 264, int64(value))
}

// SetDateTime
func (m JournalEntryAddFixedPointerBuilder) SetDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 272, 264, int64(value))
}