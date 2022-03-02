package message

import (
	"github.com/moontrade/nogc"
	"unsafe"
)

type Fixed struct {
	GC
}

func (m Fixed) IsVLS() bool {
	return false
}

func (m Fixed) BaseSize() uint16 {
	return m.Size()
}

func NewFixed(size uintptr) Fixed {
	return Fixed{GC{gcAlloc(size)}}
}

func WrapFixedFromBytes(b []byte) Fixed {
	if len(b) < 4 {
		return Fixed{}
	}
	h := *(*_bytes)(unsafe.Pointer(&b))
	m := Fixed{GC{ref: h.Data}}
	if m.ref == nil {
		return m
	}
	return m
}

func NewFixedFromBytes(b []byte) Fixed {
	if len(b) < 4 {
		return Fixed{}
	}
	m := Fixed{GC{gcAlloc(uintptr(len(b)))}}
	if m.ref == nil {
		return m
	}
	nogc.Copy(m.ref, unsafe.Pointer(&b[0]), uintptr(len(b)))
	return m
}

type FixedPointer struct {
	Pointer
}

func (m FixedPointer) Close() error {
	if m.Pointer.p != 0 {
		m.Pointer.p.Free()
		m.Pointer.p = 0
	}
	return nil
}

func (m FixedPointer) IsVLS() bool {
	return false
}

func (m FixedPointer) BaseSize() uint16 {
	return m.Size()
}

func (m FixedPointer) IsGC() bool {
	return false
}

func AllocFixed(size uintptr) FixedPointer {
	return FixedPointer{Pointer{nogc.Alloc(size)}}
}

func WrapFixedPointer(p nogc.Pointer) FixedPointer {
	return FixedPointer{Pointer{p}}
}

func AllocFixedFrom(b []byte) FixedPointer {
	if len(b) == 0 {
		return FixedPointer{}
	}
	m := FixedPointer{Pointer{nogc.Alloc(uintptr(len(b)))}}
	if m.p == 0 {
		panic(ErrOutOfMemory)
	}
	m.p.SetBytes(0, b)
	return m
}
