package dtc

import (
	"errors"
	"github.com/moontrade/nogc"
	"strings"
	"unsafe"
)

var (
	ErrJSONCompactNotAvailable = errors.New("json compact not available")
)

type Message interface {
	AsPointer() nogc.Pointer
	Size() uint16
	Type() uint16
	BaseSize() uint16
	IsVLS() bool
}

type MessageBuffer interface {
	//Clear()
}

type MessageGC struct {
	ref unsafe.Pointer
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

func (m *MessageFixedPointer) TakePointer() MessageFixedPointer {
	p := m.p
	if p != 0 {
		m.p = 0
	}
	return MessageFixedPointer{MessagePointer{p}}
}

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

func WrapVLSFromBytes(b []byte) MessageVLS {
	if len(b) < 6 {
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

func StringFixed(p nogc.Pointer, baseSize, bounds uint16, offset, maxLength int) string {
	if p == 0 {
		return ""
	}
	if baseSize < bounds {
		return ""
	}
	s := p.String(offset, maxLength)
	index := strings.IndexByte(s, 0)
	if index > -1 {
		return s[0:index]
	}
	return s
}

func SetStringFixed(p nogc.Pointer, baseSize, bounds uint16, offset, maxLength int, value string) {
	if p == 0 {
		return
	}
	if baseSize < bounds {
		return
	}
	if len(value) >= maxLength {
		value = value[0 : maxLength-1]
	}
	//nogc.Zero(p.Pointer(offset).Unsafe(), uintptr(maxLength))
	p.SetString(offset, value)
	p.SetUInt8(offset+len(value), 0)
}

func Int8(p nogc.Pointer, baseSize, bounds uint16, offset int) int8 {
	if p == 0 {
		return 0
	}
	if baseSize < bounds {
		return 0
	}
	return p.Int8(offset)
}

func SetInt8(p nogc.Pointer, baseSize, bounds uint16, offset int, value int8) {
	if p == 0 {
		return
	}
	if baseSize < bounds {
		return
	}
	p.SetInt8(offset, value)
}

func UInt8(p nogc.Pointer, baseSize, bounds uint16, offset int) uint8 {
	if p == 0 {
		return 0
	}
	if baseSize < bounds {
		return 0
	}
	return p.UInt8(offset)
}

func SetUInt8(p nogc.Pointer, baseSize, bounds uint16, offset int, value uint8) {
	if p == 0 {
		return
	}
	if baseSize < bounds {
		return
	}
	p.SetUInt8(offset, value)
}

func Int16(p nogc.Pointer, baseSize, bounds uint16, offset int) int16 {
	if p == 0 {
		return 0
	}
	if baseSize < bounds {
		return 0
	}
	return p.Int16LE(offset)
}

func SetInt16(p nogc.Pointer, baseSize, bounds uint16, offset int, value int16) {
	if p == 0 {
		return
	}
	if baseSize < bounds {
		return
	}
	p.SetInt16LE(offset, value)
}

func UInt16(p nogc.Pointer, baseSize, bounds uint16, offset int) uint16 {
	if p == 0 {
		return 0
	}
	if baseSize < bounds {
		return 0
	}
	return p.UInt16LE(offset)
}

func SetUInt16(p nogc.Pointer, baseSize, bounds uint16, offset int, value uint16) {
	if p == 0 {
		return
	}
	if baseSize < bounds {
		return
	}
	p.SetUInt16LE(offset, value)
}

func Int32(p nogc.Pointer, baseSize, bounds uint16, offset int) int32 {
	if p == 0 {
		return 0
	}
	if baseSize < bounds {
		return 0
	}
	return p.Int32LE(offset)
}
func SetInt32(p nogc.Pointer, baseSize, bounds uint16, offset int, value int32) {
	if p == 0 {
		return
	}
	if baseSize < bounds {
		return
	}
	p.SetInt32LE(offset, value)
}

func UInt32(p nogc.Pointer, baseSize, bounds uint16, offset int) uint32 {
	if p == 0 {
		return 0
	}
	if baseSize < bounds {
		return 0
	}
	return p.UInt32LE(offset)
}
func SetUInt32(p nogc.Pointer, baseSize, bounds uint16, offset int, value uint32) {
	if p == 0 {
		return
	}
	if baseSize < bounds {
		return
	}
	p.SetUInt32LE(offset, value)
}

func Int64(p nogc.Pointer, baseSize, bounds uint16, offset int) int64 {
	if p == 0 {
		return 0
	}
	if baseSize < bounds {
		return 0
	}
	return p.Int64LE(offset)
}
func SetInt64(p nogc.Pointer, baseSize, bounds uint16, offset int, value int64) {
	if p == 0 {
		return
	}
	if baseSize < bounds {
		return
	}
	p.SetInt64LE(offset, value)
}

func UInt64(p nogc.Pointer, baseSize, bounds uint16, offset int) uint64 {
	if p == 0 {
		return 0
	}
	if baseSize < bounds {
		return 0
	}
	return p.UInt64LE(offset)
}
func SetUInt64(p nogc.Pointer, baseSize, bounds uint16, offset int, value uint64) {
	if p == 0 {
		return
	}
	if baseSize < bounds {
		return
	}
	p.SetUInt64LE(offset, value)
}

func Float32(p nogc.Pointer, baseSize, bounds uint16, offset int) float32 {
	if p == 0 {
		return 0
	}
	if baseSize < bounds {
		return 0
	}
	return p.Float32LE(offset)
}

func SetFloat32(p nogc.Pointer, baseSize, bounds uint16, offset int, value float32) {
	if p == 0 {
		return
	}
	if baseSize < bounds {
		return
	}
	p.SetFloat32LE(offset, value)
}

func Float64(p nogc.Pointer, baseSize, bounds uint16, offset int) float64 {
	if p == 0 {
		return 0
	}
	if baseSize < bounds {
		return 0
	}
	return p.Float64LE(offset)
}

func SetFloat64(p nogc.Pointer, baseSize, bounds uint16, offset int, value float64) {
	if p == 0 {
		return
	}
	if baseSize < bounds {
		return
	}
	p.SetFloat64LE(offset, value)
}
