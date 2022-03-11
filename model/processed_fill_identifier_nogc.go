package model

import (
	"github.com/moontrade/dtc-go/message"
)

type ProcessedFillIdentifierPointer struct {
	message.VLSPointer
}

type ProcessedFillIdentifierPointerBuilder struct {
	message.VLSPointerBuilder
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
//     IsSnapshot            = 0
//     IsFirstMessageInBatch = 0
//     IsLastMessageInBatch  = 0
// }
func (m ProcessedFillIdentifierPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ProcessedFillIdentifierDefault)
}

// ToBuilder
func (m ProcessedFillIdentifierPointer) ToBuilder() ProcessedFillIdentifierPointerBuilder {
	return ProcessedFillIdentifierPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *ProcessedFillIdentifierPointerBuilder) Finish() ProcessedFillIdentifierPointer {
	return ProcessedFillIdentifierPointer{m.VLSPointerBuilder.Finish()}
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
func (m ProcessedFillIdentifierPointer) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// IsSnapshot
func (m ProcessedFillIdentifierPointerBuilder) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// IsFirstMessageInBatch
func (m ProcessedFillIdentifierPointer) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 12, 11)
}

// IsFirstMessageInBatch
func (m ProcessedFillIdentifierPointerBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 12, 11)
}

// IsLastMessageInBatch
func (m ProcessedFillIdentifierPointer) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 13, 12)
}

// IsLastMessageInBatch
func (m ProcessedFillIdentifierPointerBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 13, 12)
}

// SetFillIdentifier
func (m *ProcessedFillIdentifierPointerBuilder) SetFillIdentifier(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetIsSnapshot
func (m ProcessedFillIdentifierPointerBuilder) SetIsSnapshot(value uint8) {
	message.SetUint8VLS(m.Ptr, 11, 10, value)
}

// SetIsFirstMessageInBatch
func (m ProcessedFillIdentifierPointerBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Ptr, 12, 11, value)
}

// SetIsLastMessageInBatch
func (m ProcessedFillIdentifierPointerBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Ptr, 13, 12, value)
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
