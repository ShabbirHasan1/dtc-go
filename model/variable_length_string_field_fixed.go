package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Offset = 0
//     Length = 0
// }
var _VariableLengthStringFieldFixedDefault = []byte{0, 0, 0, 0}

const VariableLengthStringFieldFixedSize = 4

type VariableLengthStringFieldFixed struct {
	message.Fixed
}

type VariableLengthStringFieldFixedBuilder struct {
	message.Fixed
}

func NewVariableLengthStringFieldFixedFrom(b []byte) VariableLengthStringFieldFixed {
	return VariableLengthStringFieldFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapVariableLengthStringFieldFixed(b []byte) VariableLengthStringFieldFixed {
	return VariableLengthStringFieldFixed{Fixed: message.WrapFixed(b)}
}

func NewVariableLengthStringFieldFixed() VariableLengthStringFieldFixedBuilder {
	a := VariableLengthStringFieldFixedBuilder{message.NewFixed(4)}
	a.Unsafe().SetBytes(0, _VariableLengthStringFieldFixedDefault)
	return a
}

// Clear
// {
//     Offset = 0
//     Length = 0
// }
func (m VariableLengthStringFieldFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _VariableLengthStringFieldFixedDefault)
}

// ToBuilder
func (m VariableLengthStringFieldFixed) ToBuilder() VariableLengthStringFieldFixedBuilder {
	return VariableLengthStringFieldFixedBuilder{m.Fixed}
}

// Finish
func (m VariableLengthStringFieldFixedBuilder) Finish() VariableLengthStringFieldFixed {
	return VariableLengthStringFieldFixed{m.Fixed}
}

// Offset
func (m VariableLengthStringFieldFixed) Offset() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 2, 0)
}

// Offset
func (m VariableLengthStringFieldFixedBuilder) Offset() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 2, 0)
}

// Length
func (m VariableLengthStringFieldFixed) Length() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 4, 2)
}

// Length
func (m VariableLengthStringFieldFixedBuilder) Length() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 4, 2)
}

// SetOffset
func (m VariableLengthStringFieldFixedBuilder) SetOffset(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 2, 0, value)
}

// SetLength
func (m VariableLengthStringFieldFixedBuilder) SetLength(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 4, 2, value)
}

// Copy
func (m VariableLengthStringFieldFixed) Copy(to VariableLengthStringFieldFixedBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// Copy
func (m VariableLengthStringFieldFixedBuilder) Copy(to VariableLengthStringFieldFixedBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyPointer
func (m VariableLengthStringFieldFixed) CopyPointer(to VariableLengthStringFieldFixedPointerBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyPointer
func (m VariableLengthStringFieldFixedBuilder) CopyPointer(to VariableLengthStringFieldFixedPointerBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyToPointer
func (m VariableLengthStringFieldFixed) CopyToPointer(to VariableLengthStringFieldFixedPointerBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyToPointer
func (m VariableLengthStringFieldFixedBuilder) CopyToPointer(to VariableLengthStringFieldFixedPointerBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyTo
func (m VariableLengthStringFieldFixed) CopyTo(to VariableLengthStringFieldFixedBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}

// CopyTo
func (m VariableLengthStringFieldFixedBuilder) CopyTo(to VariableLengthStringFieldFixedBuilder) {
	to.SetOffset(m.Offset())
	to.SetLength(m.Length())
}
