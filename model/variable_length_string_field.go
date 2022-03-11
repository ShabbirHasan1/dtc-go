package model

import (
	"github.com/moontrade/dtc-go/message"
)

const VariableLengthStringFieldSize = 4

// {
//     Offset = 0
//     Length = 0
// }
var _VariableLengthStringFieldDefault = []byte{0, 0, 0, 0}

type VariableLengthStringField struct {
	message.VLS
}

type VariableLengthStringFieldBuilder struct {
	message.VLSBuilder
}

type VariableLengthStringFieldPointer struct {
	message.VLSPointer
}

type VariableLengthStringFieldPointerBuilder struct {
	message.VLSPointerBuilder
}

func NewVariableLengthStringFieldFrom(b []byte) VariableLengthStringField {
	return VariableLengthStringField{VLS: message.NewVLSFrom(b)}
}

func WrapVariableLengthStringField(b []byte) VariableLengthStringField {
	return VariableLengthStringField{VLS: message.WrapVLS(b)}
}

func NewVariableLengthStringField() VariableLengthStringFieldBuilder {
	a := VariableLengthStringFieldBuilder{message.NewVLSBuilder(4)}
	a.Unsafe().SetBytes(0, _VariableLengthStringFieldDefault)
	return a
}

func AllocVariableLengthStringField() VariableLengthStringFieldPointerBuilder {
	a := VariableLengthStringFieldPointerBuilder{message.AllocVLSBuilder(4)}
	a.Ptr.SetBytes(0, _VariableLengthStringFieldDefault)
	return a
}

func AllocVariableLengthStringFieldFrom(b []byte) VariableLengthStringFieldPointer {
	return VariableLengthStringFieldPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Offset = 0
//     Length = 0
// }
func (m VariableLengthStringFieldBuilder) Clear() {
	m.Unsafe().SetBytes(0, _VariableLengthStringFieldDefault)
}

// Clear
// {
//     Offset = 0
//     Length = 0
// }
func (m VariableLengthStringFieldPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _VariableLengthStringFieldDefault)
}

// ToBuilder
func (m VariableLengthStringField) ToBuilder() VariableLengthStringFieldBuilder {
	return VariableLengthStringFieldBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m VariableLengthStringFieldPointer) ToBuilder() VariableLengthStringFieldPointerBuilder {
	return VariableLengthStringFieldPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m VariableLengthStringFieldBuilder) Finish() VariableLengthStringField {
	return VariableLengthStringField{m.VLSBuilder.Finish()}
}

// Finish
func (m *VariableLengthStringFieldPointerBuilder) Finish() VariableLengthStringFieldPointer {
	return VariableLengthStringFieldPointer{m.VLSPointerBuilder.Finish()}
}

// Offset
func (m VariableLengthStringField) Offset() uint16 {
	return message.Uint16VLS(m.Unsafe(), 2, 0)
}

// Offset
func (m VariableLengthStringFieldBuilder) Offset() uint16 {
	return message.Uint16VLS(m.Unsafe(), 2, 0)
}

// Offset
func (m VariableLengthStringFieldPointer) Offset() uint16 {
	return message.Uint16VLS(m.Ptr, 2, 0)
}

// Offset
func (m VariableLengthStringFieldPointerBuilder) Offset() uint16 {
	return message.Uint16VLS(m.Ptr, 2, 0)
}

// Length
func (m VariableLengthStringField) Length() uint16 {
	return message.Uint16VLS(m.Unsafe(), 4, 2)
}

// Length
func (m VariableLengthStringFieldBuilder) Length() uint16 {
	return message.Uint16VLS(m.Unsafe(), 4, 2)
}

// Length
func (m VariableLengthStringFieldPointer) Length() uint16 {
	return message.Uint16VLS(m.Ptr, 4, 2)
}

// Length
func (m VariableLengthStringFieldPointerBuilder) Length() uint16 {
	return message.Uint16VLS(m.Ptr, 4, 2)
}

// SetOffset
func (m VariableLengthStringFieldBuilder) SetOffset(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 2, 0, value)
}

// SetOffset
func (m VariableLengthStringFieldPointerBuilder) SetOffset(value uint16) {
	message.SetUint16VLS(m.Ptr, 2, 0, value)
}

// SetLength
func (m VariableLengthStringFieldBuilder) SetLength(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 4, 2, value)
}

// SetLength
func (m VariableLengthStringFieldPointerBuilder) SetLength(value uint16) {
	message.SetUint16VLS(m.Ptr, 4, 2, value)
}

// Copy
func (m VariableLengthStringField) Copy(to VariableLengthStringFieldBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// Copy
func (m VariableLengthStringFieldBuilder) Copy(to VariableLengthStringFieldBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyPointer
func (m VariableLengthStringField) CopyPointer(to VariableLengthStringFieldPointerBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyPointer
func (m VariableLengthStringFieldBuilder) CopyPointer(to VariableLengthStringFieldPointerBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// Copy
func (m VariableLengthStringFieldPointer) Copy(to VariableLengthStringFieldBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// Copy
func (m VariableLengthStringFieldPointerBuilder) Copy(to VariableLengthStringFieldBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyPointer
func (m VariableLengthStringFieldPointer) CopyPointer(to VariableLengthStringFieldPointerBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyPointer
func (m VariableLengthStringFieldPointerBuilder) CopyPointer(to VariableLengthStringFieldPointerBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}
