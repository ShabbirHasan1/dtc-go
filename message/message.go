package message

import (
	"errors"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/nogc"
	"unsafe"
)

var (
	ErrJSONCompactNotAvailable = errors.New("json compact not available")
	ErrJSONCompactDetected     = errors.New("json compact detected")
	ErrWrongType               = errors.New("wrong type")
	ErrOutOfMemory             = errors.New("out of memory")
)

type Factory interface {
	Type() uint16

	Encoding() uint16
}

type Message interface {
	// Unsafe gets the raw unsafe underlying pointer of the message
	Unsafe() nogc.Pointer

	// Size of message
	Size() uint16

	// Type of message
	Type() uint16

	// BaseSize is the size of non VLS part of message if applicable
	BaseSize() uint16

	// IsVLS is true if message contains variable length strings
	IsVLS() bool

	// IsGC is true if this message is managed via Go's Garbage Collector
	// false if it was allocated outside the GC via nogc.Alloc
	IsGC() bool

	// Close frees up memory for messages allocated using nogc.Alloc
	Close() error

	MarshalJSON(b []byte) ([]byte, error)

	MarshalJSONCompact(b []byte) ([]byte, error)

	MarshalProtobuf(b []byte) ([]byte, error)
}

type Builder interface {
	Clear()

	UnmarshalJSON(r *json.Reader) error

	UnmarshalProtobuf(b []byte) error
}

type Buffer interface {
	Clear()
}

// GCPointer is a message allocated using the Go heap and managed by Go's GC.
type GCPointer struct {
	Ptr unsafe.Pointer
}

func (m GCPointer) Close() error {
	return nil
}

func (m GCPointer) Unsafe() nogc.Pointer {
	return nogc.Pointer(m.Ptr)
}

func (m GCPointer) Size() uint16 {
	if m.Ptr == nil {
		return 0
	}
	return nogc.Pointer(m.Ptr).UInt16LE(0)
}

func (m GCPointer) Type() uint16 {
	if m.Ptr == nil {
		return 0
	}
	return nogc.Pointer(m.Ptr).UInt16LE(2)
}

func (m GCPointer) IsGC() bool {
	return true
}

// Pointer is a message allocated outside the Go heap and NOT managed by Go's GC.
// Requires manually freeing the memory by calling Close.
type Pointer struct {
	Ptr nogc.Pointer
}

func (m *Pointer) Take() Pointer {
	p := m.Ptr
	if p != 0 {
		m.Ptr = 0
	}
	return Pointer{p}
}

func (m Pointer) Unsafe() nogc.Pointer {
	return m.Ptr
}

func (m Pointer) Size() uint16 {
	if m.Ptr == 0 {
		return 0
	}
	return m.Ptr.UInt16LE(0)
}

func (m Pointer) Type() uint16 {
	if m.Ptr == 0 {
		return 0
	}
	return m.Ptr.UInt16LE(2)
}

func (m Pointer) IsGC() bool {
	return false
}
