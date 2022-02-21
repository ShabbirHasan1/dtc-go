package dtc

import (
	"github.com/moontrade/nogc"
	"math"
	"unsafe"
)

type vlsBuilder struct {
	p   nogc.Pointer
	cap int
}

func (v *vlsBuilder) AsPointer() nogc.Pointer {
	return v.p
}

func (v *vlsBuilder) IsVLS() bool {
	return true
}

func (v *vlsBuilder) Size() uint16 {
	if v.p == 0 {
		return 0
	}
	return v.p.UInt16LE(0)
}
func (v *vlsBuilder) Type() uint16 {
	if v.p == 0 {
		return 0
	}
	return v.p.UInt16LE(2)
}

func (v *vlsBuilder) BaseSize() uint16 {
	if v.p == 0 {
		return 0
	}
	return v.p.UInt16LE(4)
}

type MessageVLSPointerBuilder struct {
	vlsBuilder
	growBy int
}

func MessageVLSPointerBuilderOf(b MessageBuffer) *MessageVLSPointerBuilder {
	if b == nil {
		return &MessageVLSPointerBuilder{growBy: 32}
	}
	builder, ok := b.(*MessageVLSPointerBuilder)
	if ok && builder != nil {
		return builder
	}
	return &MessageVLSPointerBuilder{}
}

func (v *MessageVLSPointerBuilder) Finish() MessageVLSPointer {
	if v == nil {
		return MessageVLSPointer{}
	}
	p := v.p
	if p == 0 {
		return MessageVLSPointer{}
	}
	v.p = 0
	return MessageVLSPointer{MessagePointer{p}}
}

func MessageVLSPointerBuilderReset(b MessageBuffer, from *MessageVLSPointer, baseSize, flex uintptr, growBy int) *MessageVLSPointerBuilder {
	var builder *MessageVLSPointerBuilder
	if b != nil {
		builder, _ = b.(*MessageVLSPointerBuilder)
	}
	if builder == nil {
		builder = &MessageVLSPointerBuilder{}
	}
	builder.reset(from, 16, flex, growBy)
	return builder
}

func MessageVLSBuilderReset(b MessageBuffer, from *MessageVLS, baseSize, flex uintptr, growBy int) *MessageVLSBuilder {
	var builder *MessageVLSBuilder
	if b != nil {
		builder, _ = b.(*MessageVLSBuilder)
	}
	if builder == nil {
		builder = &MessageVLSBuilder{}
	}
	builder.reset(from, 16, flex, growBy)
	return builder
}

func (v *MessageVLSPointerBuilder) reset(from *MessageVLSPointer, baseSize, flex uintptr, growBy int) {
	// Wrap existing allocation?
	if from != nil && from.p != 0 {
		// Free existing if needed
		if v.p != 0 {
			nogc.Free(v.p)
		}
		v.p = from.p
		v.cap = int(v.Size())
		return
	}
	v.p = nogc.Alloc(baseSize + flex)
	if v.p == 0 {
		panic(ErrOutOfMemory)
	}
	v.cap = int(baseSize + flex)
	if v.Size() < uint16(baseSize) {
		v.p.SetUInt16LE(0, uint16(baseSize))
		v.p.SetUInt16LE(4, uint16(baseSize))
	} else {
		v.p.SetUInt16LE(4, uint16(baseSize))
	}
	v.growBy = growBy
	if v.growBy < 0 {
		v.growBy = 0
	}
}

func (v *MessageVLSPointerBuilder) Extend(by int) {
	if v.growBy > by {
		by = v.growBy
	}
	newCap := uintptr(v.cap + by)
	v.p, newCap = nogc.ReallocCap(v.p, newCap)
	v.cap = int(newCap)
}

type MessageVLSBuilder struct {
	vlsBuilder
	ref    unsafe.Pointer
	growBy int
}

func MessageVLSBuilderOf(b MessageBuffer) *MessageVLSBuilder {
	if b == nil {
		return &MessageVLSBuilder{growBy: 32}
	}
	builder, ok := b.(*MessageVLSBuilder)
	if ok && builder != nil {
		return builder
	}
	return &MessageVLSBuilder{}
}

func (v *MessageVLSBuilder) Finish() MessageVLS {
	if v == nil {
		return MessageVLS{}
	}
	ref := v.ref
	if ref == nil {
		return MessageVLS{}
	}
	v.ref = nil
	v.p = 0
	return MessageVLS{MessageGC{ref}}
}

func (v *MessageVLSBuilder) reset(from *MessageVLS, baseSize, flex uintptr, growBy int) {
	if from != nil {
		v.ref = from.ref
		v.p = nogc.Pointer(v.ref)
		v.cap = int(v.Size())
		return
	}
	v.ref = gcAlloc(baseSize + flex)
	if v.ref == nil {
		panic(ErrOutOfMemory)
	}
	v.p = nogc.Pointer(v.ref)
	v.cap = int(baseSize + flex)
	if v.Size() < uint16(baseSize) {
		v.p.SetUInt16LE(0, uint16(baseSize))
		v.p.SetUInt16LE(4, uint16(baseSize))
	} else {
		v.p.SetUInt16LE(4, uint16(baseSize))
	}
	v.growBy = growBy
	if v.growBy < 0 {
		v.growBy = 0
	}
}

func (v *MessageVLSBuilder) Extend(by int) {
	if v.growBy > by {
		by = v.growBy
	}
	newCap := uintptr(v.cap + by)
	old := v.ref
	v.ref = gcAlloc(newCap)
	if v.ref == nil {
		v.p = 0
		return
	}
	nogc.Copy(v.ref, old, uintptr(v.Size()))
	v.cap = int(newCap)
}

//type VLS struct {
//	Offset uint16
//	Length uint16
//}

func StringVLS(p nogc.Pointer, bounds uint16, offset int) string {
	if p == 0 {
		return ""
	}
	var (
		size     = p.UInt16LE(0)
		baseSize = p.UInt16LE(4)
	)
	if baseSize < bounds {
		return ""
	}
	vlsOffset := p.UInt16LE(offset)
	vlsLength := p.UInt16LE(offset + 2)
	if vlsLength > 4096 {
		vlsLength = 4096
	}
	if vlsOffset == 0 || vlsLength == 0 {
		return ""
	}
	if size < vlsOffset+vlsLength {
		return ""
	}

	if p.Byte(int(vlsOffset+vlsLength)-1) == 0 {
		return p.String(int(vlsOffset), int(vlsLength)-1)
	} else {
		return p.String(int(vlsOffset), int(vlsLength))
	}
}

func SetStringVLSPointer(b *MessageVLSPointerBuilder, baseSize, bounds uint16, offset int, value string) {
	if baseSize < bounds {
		return
	}
	vlsOffset := int(b.p.UInt16LE(offset))
	vlsLength := int(b.p.UInt16LE(offset + 2))
	if vlsLength > 4096 {
		vlsLength = 4096
	}
	newLength := len(value) + 1
	if newLength > 4096 {
		value = value[0:4095]
		newLength = 4096
	}
	if vlsLength >= newLength {
		// Set new length
		b.p.SetUInt16LE(offset+2, uint16(len(value)+1))
		b.p.SetString(vlsOffset, value)
		b.p.SetByte(vlsOffset+len(value), 0)
		return
	}

	newSize := int(b.Size()) + newLength
	if newSize > math.MaxUint16 {
		return
	}
	vlsOffset = int(b.Size())
	if b.cap < newSize {
		b.Extend(newLength)
		if b.p == 0 {
			return
		}
	}

	// Set new size
	b.p.SetUInt16LE(0, uint16(newSize))
	// Set VLS offset
	b.p.SetUInt16LE(offset, b.Size())
	// Set VLS length
	b.p.SetUInt16LE(offset+2, uint16(newLength))
	// Set string
	b.p.SetString(int(b.Size()), value)
	// Set null terminator
	b.p.SetByte(int(b.Size())+len(value), 0)
	// Update to new length
	b.p.SetUInt16LE(0, uint16(newSize))
}

func SetStringVLS(b *MessageVLSBuilder, baseSize, bounds uint16, offset int, value string) {
	if baseSize < bounds {
		return
	}
	vlsOffset := int(b.p.UInt16LE(offset))
	vlsLength := int(b.p.UInt16LE(offset + 2))
	if vlsLength > 4096 {
		vlsLength = 4096
	}
	newLength := len(value) + 1
	if newLength > 4096 {
		value = value[0:4095]
		newLength = 4096
	}
	if vlsLength >= newLength {
		// Set new length
		b.p.SetUInt16LE(offset+2, uint16(len(value)+1))
		b.p.SetString(vlsOffset, value)
		b.p.SetByte(vlsOffset+len(value), 0)
		return
	}

	newSize := int(b.Size()) + newLength
	if newSize > math.MaxUint16 {
		return
	}
	vlsOffset = int(b.Size())
	if b.cap < newSize {
		b.Extend(newLength)
		if b.p == 0 {
			return
		}
	}

	// Set new size
	b.p.SetUInt16LE(0, uint16(newSize))
	// Set VLS offset
	b.p.SetUInt16LE(offset, b.Size())
	// Set VLS length
	b.p.SetUInt16LE(offset+2, uint16(newLength))
	// Set string
	b.p.SetString(int(b.Size()), value)
	// Set null terminator
	b.p.SetByte(int(b.Size())+len(value), 0)
	// Update to new length
	b.p.SetUInt16LE(0, uint16(newSize))
}
