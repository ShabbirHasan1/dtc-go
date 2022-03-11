package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                  = ProcessedFillIdentifierSize  (13)
//     Type                  = PROCESSED_FILL_IDENTIFIER  (10132)
//     BaseSize              = ProcessedFillIdentifierSize  (13)
//     FillIdentifier        = ""
//     IsSnapshot            = 0
//     IsFirstMessageInBatch = 0
//     IsLastMessageInBatch  = 0
// }
var _ProcessedFillIdentifierDefault = []byte{13, 0, 148, 39, 13, 0, 0, 0, 0, 0, 0, 0, 0}

const ProcessedFillIdentifierSize = 13

type ProcessedFillIdentifier struct {
	message.VLS
}

type ProcessedFillIdentifierBuilder struct {
	message.VLSBuilder
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
func (m ProcessedFillIdentifierBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ProcessedFillIdentifierDefault)
}

// ToBuilder
func (m ProcessedFillIdentifier) ToBuilder() ProcessedFillIdentifierBuilder {
	return ProcessedFillIdentifierBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m ProcessedFillIdentifierBuilder) Finish() ProcessedFillIdentifier {
	return ProcessedFillIdentifier{m.VLSBuilder.Finish()}
}

// FillIdentifier
func (m ProcessedFillIdentifier) FillIdentifier() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// FillIdentifier
func (m ProcessedFillIdentifierBuilder) FillIdentifier() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// IsSnapshot
func (m ProcessedFillIdentifier) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// IsSnapshot
func (m ProcessedFillIdentifierBuilder) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// IsFirstMessageInBatch
func (m ProcessedFillIdentifier) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 12, 11)
}

// IsFirstMessageInBatch
func (m ProcessedFillIdentifierBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 12, 11)
}

// IsLastMessageInBatch
func (m ProcessedFillIdentifier) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 13, 12)
}

// IsLastMessageInBatch
func (m ProcessedFillIdentifierBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 13, 12)
}

// SetFillIdentifier
func (m *ProcessedFillIdentifierBuilder) SetFillIdentifier(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetIsSnapshot
func (m ProcessedFillIdentifierBuilder) SetIsSnapshot(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 11, 10, value)
}

// SetIsFirstMessageInBatch
func (m ProcessedFillIdentifierBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 12, 11, value)
}

// SetIsLastMessageInBatch
func (m ProcessedFillIdentifierBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 13, 12, value)
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
