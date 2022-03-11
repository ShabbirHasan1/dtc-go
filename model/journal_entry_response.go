package model

import (
	"github.com/moontrade/dtc-go/message"
)

const JournalEntryResponseSize = 32

const JournalEntryResponseFixedSize = 280

// {
//     Size            = JournalEntryResponseSize  (32)
//     Type            = JOURNAL_ENTRY_RESPONSE  (706)
//     BaseSize        = JournalEntryResponseSize  (32)
//     JournalEntry    = ""
//     DateTime        = 0
//     IsFinalResponse = false
// }
var _JournalEntryResponseDefault = []byte{32, 0, 194, 2, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size            = JournalEntryResponseFixedSize  (280)
//     Type            = JOURNAL_ENTRY_RESPONSE  (706)
//     JournalEntry    = ""
//     DateTime        = 0
//     IsFinalResponse = false
// }
var _JournalEntryResponseFixedDefault = []byte{24, 1, 194, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type JournalEntryResponse struct {
	message.VLS
}

type JournalEntryResponseBuilder struct {
	message.VLSBuilder
}

type JournalEntryResponseFixed struct {
	message.Fixed
}

type JournalEntryResponseFixedBuilder struct {
	message.Fixed
}

type JournalEntryResponsePointer struct {
	message.VLSPointer
}

type JournalEntryResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

type JournalEntryResponseFixedPointer struct {
	message.FixedPointer
}

type JournalEntryResponseFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocJournalEntryResponse() JournalEntryResponsePointerBuilder {
	a := JournalEntryResponsePointerBuilder{message.AllocVLSBuilder(32)}
	a.Ptr.SetBytes(0, _JournalEntryResponseDefault)
	return a
}

func AllocJournalEntryResponseFrom(b []byte) JournalEntryResponsePointer {
	return JournalEntryResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size            = JournalEntryResponseSize  (32)
//     Type            = JOURNAL_ENTRY_RESPONSE  (706)
//     BaseSize        = JournalEntryResponseSize  (32)
//     JournalEntry    = ""
//     DateTime        = 0
//     IsFinalResponse = false
// }
func (m JournalEntryResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntryResponseDefault)
}

// Clear
// {
//     Size            = JournalEntryResponseFixedSize  (280)
//     Type            = JOURNAL_ENTRY_RESPONSE  (706)
//     JournalEntry    = ""
//     DateTime        = 0
//     IsFinalResponse = false
// }
func (m JournalEntryResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _JournalEntryResponseFixedDefault)
}

// Clear
// {
//     Size            = JournalEntryResponseSize  (32)
//     Type            = JOURNAL_ENTRY_RESPONSE  (706)
//     BaseSize        = JournalEntryResponseSize  (32)
//     JournalEntry    = ""
//     DateTime        = 0
//     IsFinalResponse = false
// }
func (m JournalEntryResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntryResponseDefault)
}

// Clear
// {
//     Size            = JournalEntryResponseFixedSize  (280)
//     Type            = JOURNAL_ENTRY_RESPONSE  (706)
//     JournalEntry    = ""
//     DateTime        = 0
//     IsFinalResponse = false
// }
func (m JournalEntryResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _JournalEntryResponseFixedDefault)
}

// ToBuilder
func (m JournalEntryResponse) ToBuilder() JournalEntryResponseBuilder {
	return JournalEntryResponseBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m JournalEntryResponseFixed) ToBuilder() JournalEntryResponseFixedBuilder {
	return JournalEntryResponseFixedBuilder{m.Fixed}
}

// ToBuilder
func (m JournalEntryResponsePointer) ToBuilder() JournalEntryResponsePointerBuilder {
	return JournalEntryResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m JournalEntryResponseFixedPointer) ToBuilder() JournalEntryResponseFixedPointerBuilder {
	return JournalEntryResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m JournalEntryResponseBuilder) Finish() JournalEntryResponse {
	return JournalEntryResponse{m.VLSBuilder.Finish()}
}

// Finish
func (m JournalEntryResponseFixedBuilder) Finish() JournalEntryResponseFixed {
	return JournalEntryResponseFixed{m.Fixed}
}

// Finish
func (m *JournalEntryResponsePointerBuilder) Finish() JournalEntryResponsePointer {
	return JournalEntryResponsePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *JournalEntryResponseFixedPointerBuilder) Finish() JournalEntryResponseFixedPointer {
	return JournalEntryResponseFixedPointer{m.FixedPointer}
}

// JournalEntry
func (m JournalEntryResponse) JournalEntry() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// JournalEntry
func (m JournalEntryResponseBuilder) JournalEntry() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
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
func (m JournalEntryResponse) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 24, 16))
}

// DateTime
func (m JournalEntryResponseBuilder) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 24, 16))
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
func (m JournalEntryResponse) IsFinalResponse() bool {
	return message.BoolVLS(m.Unsafe(), 25, 24)
}

// IsFinalResponse
func (m JournalEntryResponseBuilder) IsFinalResponse() bool {
	return message.BoolVLS(m.Unsafe(), 25, 24)
}

// IsFinalResponse
func (m JournalEntryResponsePointer) IsFinalResponse() bool {
	return message.BoolVLS(m.Ptr, 25, 24)
}

// IsFinalResponse
func (m JournalEntryResponsePointerBuilder) IsFinalResponse() bool {
	return message.BoolVLS(m.Ptr, 25, 24)
}

// JournalEntry
func (m JournalEntryResponseFixed) JournalEntry() string {
	return message.StringFixed(m.Unsafe(), 260, 4)
}

// JournalEntry
func (m JournalEntryResponseFixedBuilder) JournalEntry() string {
	return message.StringFixed(m.Unsafe(), 260, 4)
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
func (m JournalEntryResponseFixed) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 272, 264))
}

// DateTime
func (m JournalEntryResponseFixedBuilder) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 272, 264))
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
func (m JournalEntryResponseFixed) IsFinalResponse() bool {
	return message.BoolFixed(m.Unsafe(), 273, 272)
}

// IsFinalResponse
func (m JournalEntryResponseFixedBuilder) IsFinalResponse() bool {
	return message.BoolFixed(m.Unsafe(), 273, 272)
}

// IsFinalResponse
func (m JournalEntryResponseFixedPointer) IsFinalResponse() bool {
	return message.BoolFixed(m.Ptr, 273, 272)
}

// IsFinalResponse
func (m JournalEntryResponseFixedPointerBuilder) IsFinalResponse() bool {
	return message.BoolFixed(m.Ptr, 273, 272)
}

// SetJournalEntry
func (m *JournalEntryResponseBuilder) SetJournalEntry(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetJournalEntry
func (m *JournalEntryResponsePointerBuilder) SetJournalEntry(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetDateTime
func (m JournalEntryResponseBuilder) SetDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 24, 16, int64(value))
}

// SetDateTime
func (m JournalEntryResponsePointerBuilder) SetDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 24, 16, int64(value))
}

// SetIsFinalResponse
func (m JournalEntryResponseBuilder) SetIsFinalResponse(value bool) {
	message.SetBoolVLS(m.Unsafe(), 25, 24, value)
}

// SetIsFinalResponse
func (m JournalEntryResponsePointerBuilder) SetIsFinalResponse(value bool) {
	message.SetBoolVLS(m.Ptr, 25, 24, value)
}

// SetJournalEntry
func (m JournalEntryResponseFixedBuilder) SetJournalEntry(value string) {
	message.SetStringFixed(m.Unsafe(), 260, 4, value)
}

// SetJournalEntry
func (m JournalEntryResponseFixedPointerBuilder) SetJournalEntry(value string) {
	message.SetStringFixed(m.Ptr, 260, 4, value)
}

// SetDateTime
func (m JournalEntryResponseFixedBuilder) SetDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 272, 264, int64(value))
}

// SetDateTime
func (m JournalEntryResponseFixedPointerBuilder) SetDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 272, 264, int64(value))
}

// SetIsFinalResponse
func (m JournalEntryResponseFixedBuilder) SetIsFinalResponse(value bool) {
	message.SetBoolFixed(m.Unsafe(), 273, 272, value)
}

// SetIsFinalResponse
func (m JournalEntryResponseFixedPointerBuilder) SetIsFinalResponse(value bool) {
	message.SetBoolFixed(m.Ptr, 273, 272, value)
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
