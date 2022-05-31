package message

import (
	"unsafe"
)

type VLS struct {
	Pointer
	capacity int
}

func (vls VLS) BaseSize() uint16 {
	return vls.Uint16LE(4)
}

func NewVLS(init []byte) VLS {
	c := len(init) * 2
	if c == 0 {
		return VLS{}
	}
	p := Pointer{p: Alloc(uintptr(c))}
	p.SetBytes(0, init)
	return VLS{Pointer: p, capacity: cap(init)}
}

func WrapVLS(b []byte) VLS {
	if len(b) < 6 {
		return VLS{}
	}
	h := *(*_bytes)(unsafe.Pointer(&b))
	m := VLS{Pointer: Pointer{p: h.Data}, capacity: cap(b)}
	if m.p == nil {
		return m
	}
	return m
}

func WrapVLSUnsafe(p unsafe.Pointer, capacity int) VLS {
	return VLS{Pointer: Pointer{p}, capacity: capacity}
}

func WrapVLSPointer(p Pointer, capacity int) VLS {
	return VLS{Pointer: p, capacity: capacity}
}

func (vls *VLS) Extend(by int) {
	newCap := uintptr(vls.capacity + by)
	old := vls.p
	next := Alloc(newCap)
	if next == nil {
		panic("out of memory")
	}
	Copy(next, old, uintptr(vls.capacity))
	vls.capacity = int(newCap)
}
