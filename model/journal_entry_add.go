package model

import (
	"github.com/moontrade/dtc-go/message"
)

const JournalEntryAddSize = 24

const JournalEntryAddFixedSize = 272

// {
//     Size         = JournalEntryAddSize  (24)
//     Type         = JOURNAL_ENTRY_ADD  (703)
//     BaseSize     = JournalEntryAddSize  (24)
//     JournalEntry = ""
//     DateTime     = 0
// }
var _JournalEntryAddDefault = []byte{24, 0, 191, 2, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size         = JournalEntryAddFixedSize  (272)
//     Type         = JOURNAL_ENTRY_ADD  (703)
//     JournalEntry = ""
//     DateTime     = 0
// }
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

func NewJournalEntryAddFixedFrom(b []byte) JournalEntryAddFixed {
	return JournalEntryAddFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapJournalEntryAddFixed(b []byte) JournalEntryAddFixed {
	return JournalEntryAddFixed{Fixed: message.WrapFixed(b)}
}

func NewJournalEntryAddFixed() JournalEntryAddFixedBuilder {
	a := JournalEntryAddFixedBuilder{message.NewFixed(272)}
	a.Unsafe().SetBytes(0, _JournalEntryAddFixedDefault)
	return a
}

func AllocJournalEntryAdd() JournalEntryAddPointerBuilder {
	a := JournalEntryAddPointerBuilder{message.AllocVLSBuilder(24)}
	a.Ptr.SetBytes(0, _JournalEntryAddDefault)
	return a
}

func AllocJournalEntryAddFrom(b []byte) JournalEntryAddPointer {
	return JournalEntryAddPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocJournalEntryAddFixed() JournalEntryAddFixedPointerBuilder {
	a := JournalEntryAddFixedPointerBuilder{message.AllocFixed(272)}
	a.Ptr.SetBytes(0, _JournalEntryAddFixedDefault)
	return a
}

func AllocJournalEntryAddFixedFrom(b []byte) JournalEntryAddFixedPointer {
	return JournalEntryAddFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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

// Clear
// {
//     Size         = JournalEntryAddFixedSize  (272)
//     Type         = JOURNAL_ENTRY_ADD  (703)
//     JournalEntry = ""
//     DateTime     = 0
// }
func (m JournalEntryAddFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntryAddFixedDefault)
}

// Clear
// {
//     Size         = JournalEntryAddSize  (24)
//     Type         = JOURNAL_ENTRY_ADD  (703)
//     BaseSize     = JournalEntryAddSize  (24)
//     JournalEntry = ""
//     DateTime     = 0
// }
func (m JournalEntryAddPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntryAddDefault)
}

// Clear
// {
//     Size         = JournalEntryAddFixedSize  (272)
//     Type         = JOURNAL_ENTRY_ADD  (703)
//     JournalEntry = ""
//     DateTime     = 0
// }
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

// Copy
func (m JournalEntryAddFixed) Copy(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m JournalEntryAddFixedBuilder) Copy(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddFixed) CopyPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddFixedBuilder) CopyPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddFixed) CopyToPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddFixedBuilder) CopyToPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddFixed) CopyTo(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddFixedBuilder) CopyTo(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m JournalEntryAddPointer) Copy(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m JournalEntryAddPointerBuilder) Copy(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddPointer) CopyPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddPointerBuilder) CopyPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddPointer) CopyTo(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddPointerBuilder) CopyTo(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddPointer) CopyToPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddPointerBuilder) CopyToPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m JournalEntryAddFixedPointer) Copy(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m JournalEntryAddFixedPointerBuilder) Copy(to JournalEntryAddFixedBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddFixedPointer) CopyPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m JournalEntryAddFixedPointerBuilder) CopyPointer(to JournalEntryAddFixedPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddFixedPointer) CopyTo(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyTo
func (m JournalEntryAddFixedPointerBuilder) CopyTo(to JournalEntryAddBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddFixedPointer) CopyToPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}

// CopyToPointer
func (m JournalEntryAddFixedPointerBuilder) CopyToPointer(to JournalEntryAddPointerBuilder) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
}
