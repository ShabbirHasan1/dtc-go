package message

import (
	"github.com/moontrade/nogc"
	"unsafe"
)

type Fixed struct {
	GCPointer
}

func (m Fixed) IsVLS() bool {
	return false
}

func (m Fixed) BaseSize() uint16 {
	return m.Size()
}

func NewFixed(size uintptr) Fixed {
	return Fixed{GCPointer{gcAlloc(size)}}
}

func WrapFixed(b []byte) Fixed {
	if len(b) < 4 {
		return Fixed{}
	}
	h := *(*_bytes)(unsafe.Pointer(&b))
	m := Fixed{GCPointer{Ptr: h.Data}}
	if m.Ptr == nil {
		return m
	}
	return m
}

func NewFixedFrom(b []byte) Fixed {
	if len(b) < 4 {
		return Fixed{}
	}
	m := Fixed{GCPointer{gcAlloc(uintptr(len(b)))}}
	if m.Ptr == nil {
		return m
	}
	nogc.Copy(m.Ptr, unsafe.Pointer(&b[0]), uintptr(len(b)))
	return m
}

type FixedPointer struct {
	Pointer
}

func (m FixedPointer) Close() error {
	if m.Pointer.Ptr != 0 {
		m.Pointer.Ptr.Free()
		m.Pointer.Ptr = 0
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
	if m.Ptr == 0 {
		panic(ErrOutOfMemory)
	}
	m.Ptr.SetBytes(0, b)
	return m
}
