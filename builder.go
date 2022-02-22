package dtc

import (
	"errors"
	"github.com/moontrade/nogc"
	"unsafe"
)

var (
	ErrOutOfMemory = errors.New("out of memory")
)

func GCAlloc(size uintptr) unsafe.Pointer {
	return gcAlloc(size)
}

func NewFixed(size uintptr) MessageFixed {
	return MessageFixed{MessageGC{gcAlloc(size)}}
}

func NewZeroed(size uintptr) MessageFixed {
	return MessageFixed{MessageGC{gcAllocZeroed(size)}}
}

func AllocFixed(size uintptr) MessageFixedPointer {
	return MessageFixedPointer{MessagePointer{nogc.Alloc(size)}}
}

func AllocCopyFrom(b []byte) MessageFixedPointer {
	if len(b) == 0 {
		return MessageFixedPointer{}
	}
	m := MessageFixedPointer{MessagePointer{nogc.Alloc(uintptr(len(b)))}}
	if m.p == 0 {
		panic(ErrOutOfMemory)
	}
	m.p.SetBytes(0, b)
	return m
}

func NewVLS(size uintptr) MessageVLS {
	return MessageVLS{MessageGC{gcAlloc(size)}}
}

func AllocVLS(size uintptr) MessageVLSPointer {
	return MessageVLSPointer{MessagePointer{nogc.Alloc(size)}}
}
