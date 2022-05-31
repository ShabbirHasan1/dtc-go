package message

import (
	"unsafe"
)

type Fixed struct {
	Pointer
}

func (m Fixed) IsVLS() bool {
	return false
}

func (m Fixed) BaseSize() uint16 {
	return m.Size()
}

func WrapFixed(b []byte) Fixed {
	if len(b) < 4 {
		return Fixed{}
	}
	h := *(*_bytes)(unsafe.Pointer(&b))
	m := Fixed{Pointer{p: h.Data}}
	if m.p == nil {
		return m
	}
	return m
}

func WrapFixedPointer(p Pointer) Fixed {
	return Fixed{p}
}

func NewFixed(b []byte) Fixed {
	if len(b) < 4 {
		return Fixed{}
	}
	p := Pointer{p: Alloc(uintptr(len(b)))}
	p.SetBytes(0, b)
	return Fixed{p}
}
