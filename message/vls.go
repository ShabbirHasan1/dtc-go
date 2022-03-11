package message

import (
	"github.com/moontrade/nogc"
	"unsafe"
)

type VLS struct {
	GCPointer
}

func (v VLS) ToBuilder() VLSBuilder {
	return VLSBuilder{
		ref: v.GCPointer.Ptr,
		vlsBuilder: vlsBuilder{
			Ptr: nogc.Pointer(v.GCPointer.Ptr),
			cap: int(v.Size()),
		},
	}
}

func (m VLS) IsVLS() bool {
	return true
}

func (m VLS) BaseSize() uint16 {
	if m.Ptr == nil {
		return 0
	}
	return m.Unsafe().Uint16LE(4)
}

func NewVLS(size uintptr) VLS {
	return VLS{GCPointer{gcAlloc(size)}}
}

func WrapVLSFromBytes(b []byte) VLS {
	if len(b) < 6 {
		return VLS{}
	}
	h := *(*_bytes)(unsafe.Pointer(&b))
	m := VLS{GCPointer{Ptr: h.Data}}
	if m.Ptr == nil {
		return m
	}
	return m
}

func WrapVLSFromBytesOfType(b []byte, _type uint16) VLS {
	if len(b) < 6 {
		return VLS{}
	}
	if _type != nogc.Pointer(unsafe.Pointer(&b[2])).Uint16LE(0) {
		return VLS{}
	}

	h := *(*_bytes)(unsafe.Pointer(&b))
	m := VLS{GCPointer{Ptr: h.Data}}
	if m.Ptr == nil {
		return m
	}
	nogc.Pointer(m.Ptr).SetBytes(0, b)
	return m
}

type VLSPointer struct {
	Pointer
}

func (v VLSPointer) IsGC() bool {
	return false
}

func (v VLSPointer) ToBuilder() VLSPointerBuilder {
	return VLSPointerBuilder{
		vlsBuilder: vlsBuilder{
			Ptr:    v.Ptr,
			cap:    int(v.Size()),
			growBy: int(v.Size()),
		},
		ref: 0,
	}
}

func (m VLSPointer) Close() error {
	if m.Pointer.Ptr != 0 {
		m.Pointer.Ptr.Free()
		m.Pointer.Ptr = 0
	}
	return nil
}

func (m VLSPointer) IsVLS() bool {
	return true
}

func (m VLSPointer) BaseSize() uint16 {
	if m.Ptr == 0 {
		return 0
	}
	return m.Ptr.Uint16LE(4)
}

func AllocVLS(size uintptr) VLSPointer {
	return VLSPointer{Pointer{nogc.Alloc(size)}}
}

func NewVLSBuilder(size int) VLSBuilder {
	c := size * 2
	p := gcAlloc(uintptr(c))
	return VLSBuilder{vlsBuilder: vlsBuilder{
		Ptr:    nogc.Pointer(p),
		cap:    c,
		growBy: size,
	}, ref: p}
}

func AllocVLSBuilder(size uintptr) VLSPointerBuilder {
	p, c := nogc.AllocCap(size)
	return VLSPointerBuilder{vlsBuilder: vlsBuilder{
		Ptr:    p,
		cap:    int(c),
		growBy: int(size),
	}, ref: 0}
}

func AllocVLSFrom(b []byte) VLSPointer {
	if len(b) == 0 {
		return VLSPointer{}
	}
	m := VLSPointer{Pointer{nogc.Alloc(uintptr(len(b)))}}
	if m.Ptr == 0 {
		panic(ErrOutOfMemory)
	}
	m.Ptr.SetBytes(0, b)
	return m
}

func WrapVLSPointer(p nogc.Pointer) VLSPointer {
	return VLSPointer{Pointer{p}}
}

func AllocVLSFromBytesOfType(b []byte, _type uint16) VLSPointer {
	if len(b) < 6 {
		return VLSPointer{}
	}
	if _type != nogc.Pointer(unsafe.Pointer(&b[2])).Uint16LE(0) {
		return VLSPointer{}
	}
	m := VLSPointer{Pointer{nogc.Alloc(uintptr(len(b)))}}
	if m.Ptr == 0 {
		return m
	}
	m.Ptr.SetBytes(0, b)
	return m
}

func NewVLSFrom(b []byte) VLS {
	if len(b) < 6 {
		return VLS{}
	}
	m := VLS{GCPointer{Ptr: gcAlloc(uintptr(len(b)))}}
	if m.Ptr == nil {
		return m
	}
	nogc.Pointer(m.Ptr).SetBytes(0, b)
	return m
}

func WrapVLS(b []byte) VLS {
	if len(b) < 6 {
		return VLS{}
	}
	header := *(*_bytes)(unsafe.Pointer(&b))
	m := VLS{GCPointer{Ptr: header.Data}}
	if m.Ptr == nil {
		return m
	}
	return m
}

type vlsBuilder struct {
	Ptr    nogc.Pointer
	cap    int
	growBy int
}

func (v vlsBuilder) IsGC() bool {
	return false
}

func (v vlsBuilder) Clear() {
	if v.Ptr == 0 {
		return
	}
	size := v.Size()
	if size == 0 {
		return
	}
	// Zero out
	nogc.Zero(v.Ptr.Unsafe(), uintptr(size))
}

func (v vlsBuilder) Unsafe() nogc.Pointer {
	return v.Ptr
}

func (v vlsBuilder) IsVLS() bool {
	return true
}

func (v vlsBuilder) Size() uint16 {
	if v.Ptr == 0 {
		return 0
	}
	return v.Ptr.Uint16LE(0)
}
func (v vlsBuilder) Type() uint16 {
	if v.Ptr == 0 {
		return 0
	}
	return v.Ptr.Uint16LE(2)
}

func (v vlsBuilder) BaseSize() uint16 {
	if v.Ptr == 0 {
		return 0
	}
	return v.Ptr.Uint16LE(4)
}

type VLSPointerBuilder struct {
	vlsBuilder
	ref uintptr
}

func (m VLSPointerBuilder) IsGC() bool {
	return false
}

func (m VLSPointerBuilder) Close() error {
	if m.Ptr == 0 {
		return nil
	}
	nogc.Free(m.Ptr)
	m.Ptr = 0
	return nil
}

func (v *VLSPointerBuilder) Finish() VLSPointer {
	if v == nil {
		return VLSPointer{}
	}
	p := v.Ptr
	if p == 0 {
		return VLSPointer{}
	}
	v.Ptr = 0
	return VLSPointer{Pointer{p}}
}

func (v *VLSPointerBuilder) reset(from *VLSPointer, baseSize, flex uintptr, growBy int) {
	// Wrap existing allocation?
	if from != nil && from.Ptr != 0 {
		// Free existing if needed
		if v.Ptr != 0 {
			nogc.Free(v.Ptr)
		}
		v.Ptr = from.Ptr
		v.cap = int(v.Size())
		return
	}
	v.Ptr = nogc.Alloc(baseSize + flex)
	if v.Ptr == 0 {
		panic(ErrOutOfMemory)
	}
	v.cap = int(baseSize + flex)
	if v.Size() < uint16(baseSize) {
		v.Ptr.SetUint16LE(0, uint16(baseSize))
		v.Ptr.SetUint16LE(4, uint16(baseSize))
	} else {
		v.Ptr.SetUint16LE(4, uint16(baseSize))
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
	v.Ptr, newCap = nogc.ReallocCap(v.Ptr, newCap)
	v.cap = int(newCap)
}

type VLSBuilder struct {
	vlsBuilder
	ref unsafe.Pointer
}

func (m VLSBuilder) Close() error {
	m.ref = nil
	m.Ptr = 0
	return nil
}

func (m VLSBuilder) IsGC() bool {
	return true
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
	v.Ptr = 0
	return VLS{GCPointer{ref}}
}

func (v *VLSBuilder) Extend(by int) {
	if v.growBy > by {
		by = v.growBy
	}
	newCap := uintptr(v.cap + by)
	old := v.ref
	v.ref = gcAlloc(newCap)
	if v.ref == nil {
		v.Ptr = 0
		return
	}
	nogc.Copy(v.ref, old, uintptr(v.Size()))
	v.cap = int(newCap)
}
