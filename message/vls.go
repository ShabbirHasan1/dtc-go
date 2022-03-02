package message

import (
	"github.com/moontrade/nogc"
	"unsafe"
)

type VLS struct {
	GC
}

func (m VLS) IsVLS() bool {
	return true
}

func (m VLS) BaseSize() uint16 {
	if m.ref == nil {
		return 0
	}
	return m.AsPointer().UInt16LE(4)
}

func NewVLS(size uintptr) VLS {
	return VLS{GC{gcAlloc(size)}}
}

func WrapVLSFromBytes(b []byte) VLS {
	if len(b) < 6 {
		return VLS{}
	}
	h := *(*_bytes)(unsafe.Pointer(&b))
	m := VLS{GC{ref: h.Data}}
	if m.ref == nil {
		return m
	}
	return m
}

func WrapVLSFromBytesOfType(b []byte, _type uint16) VLS {
	if len(b) < 6 {
		return VLS{}
	}
	if _type != nogc.Pointer(unsafe.Pointer(&b[2])).UInt16LE(0) {
		return VLS{}
	}

	h := *(*_bytes)(unsafe.Pointer(&b))
	m := VLS{GC{ref: h.Data}}
	if m.ref == nil {
		return m
	}
	nogc.Pointer(m.ref).SetBytes(0, b)
	return m
}

type VLSPointer struct {
	Pointer
}

func (m VLSPointer) Close() error {
	if m.Pointer.p != 0 {
		m.Pointer.p.Free()
		m.Pointer.p = 0
	}
	return nil
}

func (m VLSPointer) IsVLS() bool {
	return true
}

func (m VLSPointer) BaseSize() uint16 {
	if m.p == 0 {
		return 0
	}
	return m.p.UInt16LE(4)
}

func AllocVLS(size uintptr) VLSPointer {
	return VLSPointer{Pointer{nogc.Alloc(size)}}
}

func AllocVLSFrom(b []byte) VLSPointer {
	if len(b) == 0 {
		return VLSPointer{}
	}
	m := VLSPointer{Pointer{nogc.Alloc(uintptr(len(b)))}}
	if m.p == 0 {
		panic(ErrOutOfMemory)
	}
	m.p.SetBytes(0, b)
	return m
}

func AllocVLSFromBytes(b []byte) VLSPointer {
	if len(b) < 6 {
		return VLSPointer{}
	}
	m := VLSPointer{Pointer{nogc.Alloc(uintptr(len(b)))}}
	if m.p == 0 {
		return m
	}
	m.p.SetBytes(0, b)
	return m
}

func WrapVLSPointer(p nogc.Pointer) VLSPointer {
	return VLSPointer{Pointer{p}}
}

func AllocVLSFromBytesOfType(b []byte, _type uint16) VLSPointer {
	if len(b) < 6 {
		return VLSPointer{}
	}
	if _type != nogc.Pointer(unsafe.Pointer(&b[2])).UInt16LE(0) {
		return VLSPointer{}
	}
	m := VLSPointer{Pointer{nogc.Alloc(uintptr(len(b)))}}
	if m.p == 0 {
		return m
	}
	m.p.SetBytes(0, b)
	return m
}

func NewVLSFromBytes(b []byte) VLS {
	if len(b) < 6 {
		return VLS{}
	}
	m := VLS{GC{ref: gcAlloc(uintptr(len(b)))}}
	if m.ref == nil {
		return m
	}
	nogc.Pointer(m.ref).SetBytes(0, b)
	return m
}

func WrapVLS(b []byte) VLS {
	if len(b) < 6 {
		return VLS{}
	}
	header := *(*_bytes)(unsafe.Pointer(&b))
	m := VLS{GC{ref: header.Data}}
	if m.ref == nil {
		return m
	}
	return m
}

func NewVLSFromBytesOfType(b []byte, _type uint16) VLS {
	if len(b) < 6 {
		return VLS{}
	}
	if _type != nogc.Pointer(unsafe.Pointer(&b[2])).UInt16LE(0) {
		return VLS{}
	}
	m := VLS{GC{ref: gcAlloc(uintptr(len(b)))}}
	if m.ref == nil {
		return m
	}
	nogc.Pointer(m.ref).SetBytes(0, b)
	return m
}

type vlsBuilder struct {
	p   nogc.Pointer
	cap int
}

func (v *vlsBuilder) Clear() {
	if v.p == 0 {
		return
	}
	size := v.Size()
	if size == 0 {
		return
	}
	// Zero out
	nogc.Zero(v.p.Unsafe(), uintptr(size))
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

type VLSPointerBuilder struct {
	vlsBuilder
	growBy int
}

func (m *VLSPointerBuilder) IsGC() bool {
	return false
}

func (m *VLSPointerBuilder) Close() error {
	if m.p == 0 {
		return nil
	}
	nogc.Free(m.p)
	m.p = 0
	return nil
}

func VLSPointerBuilderOf(b Buffer) *VLSPointerBuilder {
	if b == nil {
		return &VLSPointerBuilder{growBy: 32}
	}
	builder, ok := b.(*VLSPointerBuilder)
	if ok && builder != nil {
		return builder
	}
	return &VLSPointerBuilder{}
}

func (v *VLSPointerBuilder) Finish() VLSPointer {
	if v == nil {
		return VLSPointer{}
	}
	p := v.p
	if p == 0 {
		return VLSPointer{}
	}
	v.p = 0
	return VLSPointer{Pointer{p}}
}

func VLSPointerBuilderReset(b Buffer, from *VLSPointer, baseSize, flex uintptr, growBy int) *VLSPointerBuilder {
	var builder *VLSPointerBuilder
	if b != nil {
		builder, _ = b.(*VLSPointerBuilder)
	}
	if builder == nil {
		builder = &VLSPointerBuilder{}
	}
	builder.reset(from, 16, flex, growBy)
	return builder
}

func VLSBuilderReset(b Buffer, from *VLS, baseSize, flex uintptr, growBy int) *VLSBuilder {
	var builder *VLSBuilder
	if b != nil {
		builder, _ = b.(*VLSBuilder)
	}
	if builder == nil {
		builder = &VLSBuilder{}
	}
	builder.reset(from, 16, flex, growBy)
	return builder
}

func (v *VLSPointerBuilder) reset(from *VLSPointer, baseSize, flex uintptr, growBy int) {
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

func (v *VLSPointerBuilder) Extend(by int) {
	if v.growBy > by {
		by = v.growBy
	}
	newCap := uintptr(v.cap + by)
	v.p, newCap = nogc.ReallocCap(v.p, newCap)
	v.cap = int(newCap)
}

type VLSBuilder struct {
	vlsBuilder
	ref    unsafe.Pointer
	growBy int
}

func (m *VLSBuilder) Close() error {
	m.ref = nil
	m.p = 0
	return nil
}

func (m *VLSBuilder) IsGC() bool {
	return true
}

func VLSBuilderOf(b Buffer) *VLSBuilder {
	if b == nil {
		return &VLSBuilder{growBy: 32}
	}
	builder, ok := b.(*VLSBuilder)
	if ok && builder != nil {
		return builder
	}
	return &VLSBuilder{}
}

func (v *VLSBuilder) Finish() VLS {
	if v == nil {
		return VLS{}
	}
	ref := v.ref
	if ref == nil {
		return VLS{}
	}
	v.ref = nil
	v.p = 0
	return VLS{GC{ref}}
}

func (v *VLSBuilder) reset(from *VLS, baseSize, flex uintptr, growBy int) {
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

func (v *VLSBuilder) Extend(by int) {
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
