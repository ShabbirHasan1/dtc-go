package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const ProcessedFillIdentifierSize = 13

// {
//     Size                  = ProcessedFillIdentifierSize  (13)
//     Type                  = PROCESSED_FILL_IDENTIFIER  (10132)
//     BaseSize              = ProcessedFillIdentifierSize  (13)
//     FillIdentifier        = ""
//     IsSnapshot            = false
//     IsFirstMessageInBatch = false
//     IsLastMessageInBatch  = false
// }
var _ProcessedFillIdentifierDefault = []byte{13, 0, 148, 39, 13, 0, 0, 0, 0, 0, 0, 0, 0}

type ProcessedFillIdentifier struct {
	message.VLS
}

type ProcessedFillIdentifierBuilder struct {
	message.VLSBuilder
}

type ProcessedFillIdentifierPointer struct {
	message.VLSPointer
}

type ProcessedFillIdentifierPointerBuilder struct {
	message.VLSPointerBuilder
}

func NewProcessedFillIdentifierFrom(b []byte) ProcessedFillIdentifier {
	return ProcessedFillIdentifier{VLS: message.NewVLSFrom(b)}
}

func WrapProcessedFillIdentifier(b []byte) ProcessedFillIdentifier {
	return ProcessedFillIdentifier{VLS: message.WrapVLS(b)}
}

func NewProcessedFillIdentifier() ProcessedFillIdentifierBuilder {
	a := ProcessedFillIdentifierBuilder{message.NewVLSBuilder(13)}
	a.Unsafe().SetBytes(0, _ProcessedFillIdentifierDefault)
	return a
}

func AllocProcessedFillIdentifier() ProcessedFillIdentifierPointerBuilder {
	a := ProcessedFillIdentifierPointerBuilder{message.AllocVLSBuilder(13)}
	a.Ptr.SetBytes(0, _ProcessedFillIdentifierDefault)
	return a
}

func AllocProcessedFillIdentifierFrom(b []byte) ProcessedFillIdentifierPointer {
	return ProcessedFillIdentifierPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                  = ProcessedFillIdentifierSize  (13)
//     Type                  = PROCESSED_FILL_IDENTIFIER  (10132)
//     BaseSize              = ProcessedFillIdentifierSize  (13)
//     FillIdentifier        = ""
//     IsSnapshot            = false
//     IsFirstMessageInBatch = false
//     IsLastMessageInBatch  = false
// }
func (m ProcessedFillIdentifierBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ProcessedFillIdentifierDefault)
}

// Clear
// {
//     Size                  = ProcessedFillIdentifierSize  (13)
//     Type                  = PROCESSED_FILL_IDENTIFIER  (10132)
//     BaseSize              = ProcessedFillIdentifierSize  (13)
//     FillIdentifier        = ""
//     IsSnapshot            = false
//     IsFirstMessageInBatch = false
//     IsLastMessageInBatch  = false
// }
func (m ProcessedFillIdentifierPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ProcessedFillIdentifierDefault)
}

// ToBuilder
func (m ProcessedFillIdentifier) ToBuilder() ProcessedFillIdentifierBuilder {
	return ProcessedFillIdentifierBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m ProcessedFillIdentifierPointer) ToBuilder() ProcessedFillIdentifierPointerBuilder {
	return ProcessedFillIdentifierPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m ProcessedFillIdentifierBuilder) Finish() ProcessedFillIdentifier {
	return ProcessedFillIdentifier{m.VLSBuilder.Finish()}
}

// Finish
func (m *ProcessedFillIdentifierPointerBuilder) Finish() ProcessedFillIdentifierPointer {
	return ProcessedFillIdentifierPointer{m.VLSPointerBuilder.Finish()}
}

// FillIdentifier
func (m ProcessedFillIdentifier) FillIdentifier() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// FillIdentifier
func (m ProcessedFillIdentifierBuilder) FillIdentifier() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// FillIdentifier
func (m ProcessedFillIdentifierPointer) FillIdentifier() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// FillIdentifier
func (m ProcessedFillIdentifierPointerBuilder) FillIdentifier() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// IsSnapshot
func (m ProcessedFillIdentifier) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsSnapshot
func (m ProcessedFillIdentifierBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsSnapshot
func (m ProcessedFillIdentifierPointer) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// IsSnapshot
func (m ProcessedFillIdentifierPointerBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// IsFirstMessageInBatch
func (m ProcessedFillIdentifier) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 12, 11)
}

// IsFirstMessageInBatch
func (m ProcessedFillIdentifierBuilder) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 12, 11)
}

// IsFirstMessageInBatch
func (m ProcessedFillIdentifierPointer) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 12, 11)
}

// IsFirstMessageInBatch
func (m ProcessedFillIdentifierPointerBuilder) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 12, 11)
}

// IsLastMessageInBatch
func (m ProcessedFillIdentifier) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 13, 12)
}

// IsLastMessageInBatch
func (m ProcessedFillIdentifierBuilder) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 13, 12)
}

// IsLastMessageInBatch
func (m ProcessedFillIdentifierPointer) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 13, 12)
}

// IsLastMessageInBatch
func (m ProcessedFillIdentifierPointerBuilder) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 13, 12)
}

// SetFillIdentifier
func (m *ProcessedFillIdentifierBuilder) SetFillIdentifier(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetFillIdentifier
func (m *ProcessedFillIdentifierPointerBuilder) SetFillIdentifier(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetIsSnapshot
func (m ProcessedFillIdentifierBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Unsafe(), 11, 10, value)
}

// SetIsSnapshot
func (m ProcessedFillIdentifierPointerBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Ptr, 11, 10, value)
}

// SetIsFirstMessageInBatch
func (m ProcessedFillIdentifierBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolVLS(m.Unsafe(), 12, 11, value)
}

// SetIsFirstMessageInBatch
func (m ProcessedFillIdentifierPointerBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolVLS(m.Ptr, 12, 11, value)
}

// SetIsLastMessageInBatch
func (m ProcessedFillIdentifierBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolVLS(m.Unsafe(), 13, 12, value)
}

// SetIsLastMessageInBatch
func (m ProcessedFillIdentifierPointerBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolVLS(m.Ptr, 13, 12, value)
}

// Copy
func (m ProcessedFillIdentifier) Copy(to ProcessedFillIdentifierBuilder) {
	to.SetFillIdentifier(m.FillIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// Copy
func (m ProcessedFillIdentifierBuilder) Copy(to ProcessedFillIdentifierBuilder) {
	to.SetFillIdentifier(m.FillIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// CopyPointer
func (m ProcessedFillIdentifier) CopyPointer(to ProcessedFillIdentifierPointerBuilder) {
	to.SetFillIdentifier(m.FillIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// CopyPointer
func (m ProcessedFillIdentifierBuilder) CopyPointer(to ProcessedFillIdentifierPointerBuilder) {
	to.SetFillIdentifier(m.FillIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// Copy
func (m ProcessedFillIdentifierPointer) Copy(to ProcessedFillIdentifierBuilder) {
	to.SetFillIdentifier(m.FillIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// Copy
func (m ProcessedFillIdentifierPointerBuilder) Copy(to ProcessedFillIdentifierBuilder) {
	to.SetFillIdentifier(m.FillIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// CopyPointer
func (m ProcessedFillIdentifierPointer) CopyPointer(to ProcessedFillIdentifierPointerBuilder) {
	to.SetFillIdentifier(m.FillIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}

// CopyPointer
func (m ProcessedFillIdentifierPointerBuilder) CopyPointer(to ProcessedFillIdentifierPointerBuilder) {
	to.SetFillIdentifier(m.FillIdentifier())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
}
