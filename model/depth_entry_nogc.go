package model

import (
	"github.com/moontrade/dtc-go/message"
)

type DepthEntryPointer struct {
	message.FixedPointer
}

type DepthEntryPointerBuilder struct {
	message.FixedPointer
}

func AllocDepthEntry() DepthEntryPointerBuilder {
	a := DepthEntryPointerBuilder{message.AllocFixed(16)}
	a.Ptr.SetBytes(0, _DepthEntryDefault)
	return a
}

func AllocDepthEntryFrom(b []byte) DepthEntryPointer {
	return DepthEntryPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Price    = 0
//     Quantity = 0
// }
func (m DepthEntryPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _DepthEntryDefault)
}

// ToBuilder
func (m DepthEntryPointer) ToBuilder() DepthEntryPointerBuilder {
	return DepthEntryPointerBuilder{m.FixedPointer}
}

// Finish
func (m *DepthEntryPointerBuilder) Finish() DepthEntryPointer {
	return DepthEntryPointer{m.FixedPointer}
}

// Price
func (m DepthEntryPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 8, 0)
}

// Price
func (m DepthEntryPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 8, 0)
}

// Quantity
func (m DepthEntryPointer) Quantity() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// Quantity
func (m DepthEntryPointerBuilder) Quantity() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// SetPrice
func (m DepthEntryPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 8, 0, value)
}

// SetQuantity
func (m DepthEntryPointerBuilder) SetQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 12, 8, value)
}

// Copy
func (m DepthEntryPointer) Copy(to DepthEntryBuilder) {
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
}

// Copy
func (m DepthEntryPointerBuilder) Copy(to DepthEntryBuilder) {
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
}

// CopyPointer
func (m DepthEntryPointer) CopyPointer(to DepthEntryPointerBuilder) {
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
}

// CopyPointer
func (m DepthEntryPointerBuilder) CopyPointer(to DepthEntryPointerBuilder) {
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
}
