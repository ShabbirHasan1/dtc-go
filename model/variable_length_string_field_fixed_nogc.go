package model

import (
	"github.com/moontrade/dtc-go/message"
)

type VariableLengthStringFieldFixedPointer struct {
	message.FixedPointer
}

type VariableLengthStringFieldFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocVariableLengthStringFieldFixed() VariableLengthStringFieldFixedPointerBuilder {
	a := VariableLengthStringFieldFixedPointerBuilder{message.AllocFixed(4)}
	a.Ptr.SetBytes(0, _VariableLengthStringFieldFixedDefault)
	return a
}

func AllocVariableLengthStringFieldFixedFrom(b []byte) VariableLengthStringFieldFixedPointer {
	return VariableLengthStringFieldFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Offset = 0
//     Length = 0
// }
func (m VariableLengthStringFieldFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _VariableLengthStringFieldFixedDefault)
}

// ToBuilder
func (m VariableLengthStringFieldFixedPointer) ToBuilder() VariableLengthStringFieldFixedPointerBuilder {
	return VariableLengthStringFieldFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *VariableLengthStringFieldFixedPointerBuilder) Finish() VariableLengthStringFieldFixedPointer {
	return VariableLengthStringFieldFixedPointer{m.FixedPointer}
}

// Offset
func (m VariableLengthStringFieldFixedPointer) Offset() uint16 {
	return message.Uint16Fixed(m.Ptr, 2, 0)
}

// Offset
func (m VariableLengthStringFieldFixedPointerBuilder) Offset() uint16 {
	return message.Uint16Fixed(m.Ptr, 2, 0)
}

// Length
func (m VariableLengthStringFieldFixedPointer) Length() uint16 {
	return message.Uint16Fixed(m.Ptr, 4, 2)
}

// Length
func (m VariableLengthStringFieldFixedPointerBuilder) Length() uint16 {
	return message.Uint16Fixed(m.Ptr, 4, 2)
}

// SetOffset
func (m VariableLengthStringFieldFixedPointerBuilder) SetOffset(value uint16) {
	message.SetUint16Fixed(m.Ptr, 2, 0, value)
}

// SetLength
func (m VariableLengthStringFieldFixedPointerBuilder) SetLength(value uint16) {
	message.SetUint16Fixed(m.Ptr, 4, 2, value)
}

// Copy
func (m VariableLengthStringFieldFixedPointer) Copy(to VariableLengthStringFieldFixedBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// Copy
func (m VariableLengthStringFieldFixedPointerBuilder) Copy(to VariableLengthStringFieldFixedBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyPointer
func (m VariableLengthStringFieldFixedPointer) CopyPointer(to VariableLengthStringFieldFixedPointerBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyPointer
func (m VariableLengthStringFieldFixedPointerBuilder) CopyPointer(to VariableLengthStringFieldFixedPointerBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyTo
func (m VariableLengthStringFieldFixedPointer) CopyTo(to VariableLengthStringFieldFixedBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyTo
func (m VariableLengthStringFieldFixedPointerBuilder) CopyTo(to VariableLengthStringFieldFixedBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyToPointer
func (m VariableLengthStringFieldFixedPointer) CopyToPointer(to VariableLengthStringFieldFixedPointerBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyToPointer
func (m VariableLengthStringFieldFixedPointerBuilder) CopyToPointer(to VariableLengthStringFieldFixedPointerBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}
