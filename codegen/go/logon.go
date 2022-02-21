package _go

import (
	"github.com/moontrade/nogc"
	"unsafe"
)

type Message interface {
	AsPointer() nogc.Pointer
	Size() uint16
	Type() uint16
	BaseSize() uint16
	IsVLS() bool
	//Clear()
}

type MessageBase struct {
	gcRef unsafe.Pointer
}

func (mbp *MessageBase) AsPointer() nogc.Pointer {
	return nogc.Pointer(mbp.gcRef)
}

func (self *MessageBase) Size() uint16 {
	return self.AsPointer().UInt16LE(0)
}

func (self *MessageBase) Type() uint16 {
	return self.AsPointer().UInt16LE(2)
}

type MessageBasePointer struct {
	p nogc.Pointer
}

func (mbp *MessageBasePointer) AsPointer() nogc.Pointer {
	return mbp.p
}

func (self *MessageBasePointer) Size() uint16 {
	return self.AsPointer().UInt16LE(0)
}

func (self *MessageBasePointer) Type() uint16 {
	return self.AsPointer().UInt16LE(2)
}

type MessageFixed struct {
	MessageBase
}

func (self *MessageFixed) IsVLS() bool {
	return false
}

func (self *MessageFixed) BaseSize() uint16 {
	return nogc.Pointer(self.gcRef).UInt16LE(0)
}

type MessageFixedPointer struct {
	MessageBasePointer
}

func (self *MessageFixedPointer) Close() error {
	if self.MessageBasePointer.p != 0 {
		self.MessageBasePointer.p.Free()
		self.MessageBasePointer.p = 0
	}
	return nil
}

func (self *MessageFixedPointer) IsVLS() bool {
	return true
}

func (self *MessageFixedPointer) BaseSize() uint16 {
	return self.AsPointer().UInt16LE(4)
}

type MessageVLS struct {
	MessageBase
}

func (self *MessageVLS) IsVLS() bool {
	return true
}

func (self *MessageVLS) BaseSize() uint16 {
	return self.AsPointer().UInt16LE(4)
}

type MessageVLSPointer struct {
	MessageBasePointer
}

func (self *MessageVLSPointer) Close() error {
	if self.MessageBasePointer.p != 0 {
		self.MessageBasePointer.p.Free()
		self.MessageBasePointer.p = 0
	}
	return nil
}

func (self *MessageVLSPointer) IsVLS() bool {
	return true
}

func (self *MessageVLSPointer) BaseSize() uint16 {
	if self.AsPointer() == 0 {
		return 0
	}
	return self.AsPointer().UInt16LE(4)
}

func UInt16(m Message, offset int, bounds uint16) uint16 {
	p := m.AsPointer()
	if p == 0 {
		return 0
	}
	if m.BaseSize() < bounds {
		return 0
	}
	return m.AsPointer().UInt16LE(offset)
}

type EncodingRequest interface {
	Message

	ProtocolVersion() int32
	Encoding() int32
	ProtocolType() int32
}

type EncodingRequestVLS struct {
	MessageVLS
}

type EncodingRequestFixed struct {
	MessageFixed
}

type EncodingRequestFixedPointer struct {
	MessageFixedPointer
}

type EncodingRequestVLSPointer struct {
	MessageVLSPointer
}

func (self *EncodingRequestVLSPointer) ProtocolVersion() int32 {
	if self.BaseSize() < 12 {
		return 0
	}
	return self.AsPointer().Int32LE(8)
}
func (self *EncodingRequestVLSPointer) Encoding() int32 {
	if self.BaseSize() < 12 {
		return 0
	}
	return self.AsPointer().Int32LE(8)
}
func (self *EncodingRequestVLSPointer) ProtocolType() int32 {
	if self.BaseSize() < 12 {
		return 0
	}
	return self.AsPointer().Int32LE(8)
}
