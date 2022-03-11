package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Price    = 0
//     Quantity = 0
// }
var _DepthEntryDefault = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const DepthEntrySize = 16

type DepthEntry struct {
	message.Fixed
}

type DepthEntryBuilder struct {
	message.Fixed
}

func NewDepthEntryFrom(b []byte) DepthEntry {
	return DepthEntry{Fixed: message.NewFixedFrom(b)}
}

func WrapDepthEntry(b []byte) DepthEntry {
	return DepthEntry{Fixed: message.WrapFixed(b)}
}

func NewDepthEntry() DepthEntryBuilder {
	a := DepthEntryBuilder{message.NewFixed(16)}
	a.Unsafe().SetBytes(0, _DepthEntryDefault)
	return a
}

// Clear
// {
//     Price    = 0
//     Quantity = 0
// }
func (m DepthEntryBuilder) Clear() {
	m.Unsafe().SetBytes(0, _DepthEntryDefault)
}

// ToBuilder
func (m DepthEntry) ToBuilder() DepthEntryBuilder {
	return DepthEntryBuilder{m.Fixed}
}

// Finish
func (m DepthEntryBuilder) Finish() DepthEntry {
	return DepthEntry{m.Fixed}
}

// Price
func (m DepthEntry) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 8, 0)
}

// Price
func (m DepthEntryBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 8, 0)
}

// Quantity
func (m DepthEntry) Quantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// Quantity
func (m DepthEntryBuilder) Quantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// SetPrice
func (m DepthEntryBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 8, 0, value)
}

// SetQuantity
func (m DepthEntryBuilder) SetQuantity(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 12, 8, value)
}

// Copy
func (m DepthEntry) Copy(to DepthEntryBuilder) {
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
}

// Copy
func (m DepthEntryBuilder) Copy(to DepthEntryBuilder) {
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
}

// CopyPointer
func (m DepthEntry) CopyPointer(to DepthEntryPointerBuilder) {
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
}

// CopyPointer
func (m DepthEntryBuilder) CopyPointer(to DepthEntryPointerBuilder) {
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
}
