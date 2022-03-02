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
	// AsPointer gets the raw unsafe underlying pointer of the message
	AsPointer() nogc.Pointer

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
	UnmarshalJSON(r *json.Reader) error

	UnmarshalProtobuf(b []byte) error
}

type Buffer interface {
	Clear()
}

type GC struct {
	ref unsafe.Pointer
}

func (m GC) Close() error {
	return nil
}

func (m GC) AsPointer() nogc.Pointer {
	return nogc.Pointer(m.ref)
}

func (m GC) Size() uint16 {
	if m.ref == nil {
		return 0
	}
	return m.AsPointer().UInt16LE(0)
}

func (m GC) Type() uint16 {
	if m.ref == nil {
		return 0
	}
	return m.AsPointer().UInt16LE(2)
}

func (m GC) IsGC() bool {
	return true
}

type Pointer struct {
	p nogc.Pointer
}

func (m *Pointer) Take() Pointer {
	p := m.p
	if p != 0 {
		m.p = 0
	}
	return Pointer{p}
}

func (m Pointer) AsPointer() nogc.Pointer {
	return m.p
}

func (m Pointer) Size() uint16 {
	if m.p == 0 {
		return 0
	}
	return m.p.UInt16LE(0)
}

func (m Pointer) Type() uint16 {
	if m.p == 0 {
		return 0
	}
	return m.p.UInt16LE(2)
}

func (m Pointer) IsGC() bool {
	return false
}
