package dtc

import (
	"errors"
	"github.com/moontrade/dtc-go/json"
	"github.com/moontrade/nogc"
	"unsafe"
)

var (
	ErrJSONCompactNotAvailable = errors.New("json compact not available")
	ErrJSONCompactDetected     = errors.New("json compact detected")
	ErrWrongType               = errors.New("wrong type")
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

type MessageBuilder interface {
	UnmarshalJSON(r *json.Reader) error

	UnmarshalProtobuf(b []byte) error
}

type MessageBuffer interface {
	Clear()
}

type MessageGC struct {
	ref unsafe.Pointer
}

func (m MessageGC) Close() error {
	return nil
}

func (m MessageGC) AsPointer() nogc.Pointer {
	return nogc.Pointer(m.ref)
}

func (m MessageGC) Size() uint16 {
	if m.ref == nil {
		return 0
	}
	return m.AsPointer().UInt16LE(0)
}

func (m MessageGC) Type() uint16 {
	if m.ref == nil {
		return 0
	}
	return m.AsPointer().UInt16LE(2)
}

func (m MessageGC) IsGC() bool {
	return true
}

type MessagePointer struct {
	p nogc.Pointer
}

func (m *MessagePointer) Take() MessagePointer {
	p := m.p
	if p != 0 {
		m.p = 0
	}
	return MessagePointer{p}
}

func (m MessagePointer) AsPointer() nogc.Pointer {
	return m.p
}

func (m MessagePointer) Size() uint16 {
	if m.p == 0 {
		return 0
	}
	return m.p.UInt16LE(0)
}

func (m MessagePointer) Type() uint16 {
	if m.p == 0 {
		return 0
	}
	return m.p.UInt16LE(2)
}

func (m MessagePointer) IsGC() bool {
	return false
}

type MessageFixed struct {
	MessageGC
}

func (m MessageFixed) IsVLS() bool {
	return false
}

func (m MessageFixed) BaseSize() uint16 {
	return m.Size()
}

type MessageFixedPointer struct {
	MessagePointer
}

//func (m *MessageFixedPointer) TakePointer() MessageFixedPointer {
//	p := m.p
//	if p != 0 {
//		m.p = 0
//	}
//	return MessageFixedPointer{MessagePointer{p}}
//}

func (m MessageFixedPointer) Close() error {
	if m.MessagePointer.p != 0 {
		m.MessagePointer.p.Free()
		m.MessagePointer.p = 0
	}
	return nil
}

func (m MessageFixedPointer) IsVLS() bool {
	return true
}

func (m MessageFixedPointer) BaseSize() uint16 {
	return m.Size()
}

func (m MessageFixedPointer) IsGC() bool {
	return false
}

type MessageVLS struct {
	MessageGC
}

func (m MessageVLS) IsVLS() bool {
	return true
}

func (m MessageVLS) BaseSize() uint16 {
	if m.ref == nil {
		return 0
	}
	return m.AsPointer().UInt16LE(4)
}

type MessageVLSPointer struct {
	MessagePointer
}

func NewVLSFromBytes(b []byte) MessageVLS {
	if len(b) < 6 {
		return MessageVLS{}
	}
	m := MessageVLS{MessageGC{ref: gcAlloc(uintptr(len(b)))}}
	if m.ref == nil {
		return m
	}
	nogc.Pointer(m.ref).SetBytes(0, b)
	return m
}

func WrapVLS(b []byte) MessageVLS {
	if len(b) < 6 {
		return MessageVLS{}
	}
	header := *(*_bytes)(unsafe.Pointer(&b))
	m := MessageVLS{MessageGC{ref: header.Data}}
	if m.ref == nil {
		return m
	}
	return m
}

func NewVLSFromBytesOfType(b []byte, _type uint16) MessageVLS {
	if len(b) < 6 {
		return MessageVLS{}
	}
	if _type != nogc.Pointer(unsafe.Pointer(&b[2])).UInt16LE(0) {
		return MessageVLS{}
	}
	m := MessageVLS{MessageGC{ref: gcAlloc(uintptr(len(b)))}}
	if m.ref == nil {
		return m
	}
	nogc.Pointer(m.ref).SetBytes(0, b)
	return m
}

type _bytes struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

func WrapFixedFromBytes(b []byte) MessageFixed {
	if len(b) < 4 {
		return MessageFixed{}
	}
	h := *(*_bytes)(unsafe.Pointer(&b))
	m := MessageFixed{MessageGC{ref: h.Data}}
	if m.ref == nil {
		return m
	}
	return m
}

func WrapFixedPointer(p nogc.Pointer) MessageFixedPointer {
	return MessageFixedPointer{MessagePointer{p}}
}

func WrapVLSFromBytes(b []byte) MessageVLS {
	if len(b) < 6 {
		return MessageVLS{}
	}
	h := *(*_bytes)(unsafe.Pointer(&b))
	m := MessageVLS{MessageGC{ref: h.Data}}
	if m.ref == nil {
		return m
	}
	return m
}

func WrapVLSPointer(p nogc.Pointer) MessageVLSPointer {
	return MessageVLSPointer{MessagePointer{p}}
}

func WrapVLSFromBytesOfType(b []byte, _type uint16) MessageVLS {
	if len(b) < 6 {
		return MessageVLS{}
	}
	if _type != nogc.Pointer(unsafe.Pointer(&b[2])).UInt16LE(0) {
		return MessageVLS{}
	}

	h := *(*_bytes)(unsafe.Pointer(&b))
	m := MessageVLS{MessageGC{ref: h.Data}}
	if m.ref == nil {
		return m
	}
	nogc.Pointer(m.ref).SetBytes(0, b)
	return m
}

func NewFixedFromBytes(b []byte) MessageFixed {
	if len(b) < 4 {
		return MessageFixed{}
	}
	m := MessageFixed{MessageGC{gcAlloc(uintptr(len(b)))}}
	if m.ref == nil {
		return m
	}
	nogc.Copy(m.ref, unsafe.Pointer(&b[0]), uintptr(len(b)))
	return m
}

func AllocVLSFromBytes(b []byte) MessageVLSPointer {
	if len(b) < 6 {
		return MessageVLSPointer{}
	}
	m := MessageVLSPointer{MessagePointer{nogc.Alloc(uintptr(len(b)))}}
	if m.p == 0 {
		return m
	}
	m.p.SetBytes(0, b)
	return m
}

func AllocVLSFromBytesOfType(b []byte, _type uint16) MessageVLSPointer {
	if len(b) < 6 {
		return MessageVLSPointer{}
	}
	if _type != nogc.Pointer(unsafe.Pointer(&b[2])).UInt16LE(0) {
		return MessageVLSPointer{}
	}
	m := MessageVLSPointer{MessagePointer{nogc.Alloc(uintptr(len(b)))}}
	if m.p == 0 {
		return m
	}
	m.p.SetBytes(0, b)
	return m
}

func (m MessageVLSPointer) Close() error {
	if m.MessagePointer.p != 0 {
		m.MessagePointer.p.Free()
		m.MessagePointer.p = 0
	}
	return nil
}

func (m MessageVLSPointer) IsVLS() bool {
	return true
}

func (m MessageVLSPointer) BaseSize() uint16 {
	if m.p == 0 {
		return 0
	}
	return m.p.UInt16LE(4)
}

type FixedString struct {
	Value string
}
