package message

import (
	"github.com/moontrade/nogc"
	"math"
	"strings"
)

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

//type VLS_t struct {
//	Offset uint16
//	Length uint16
//}

// StringVLS returns an unsafe Go string
func StringVLS(p nogc.Pointer, bounds uint16, offset int) string {
	if p == 0 {
		return ""
	}
	var (
		size     = p.UInt16LE(0)
		baseSize = p.UInt16LE(4)
	)
	if baseSize < bounds {
		return ""
	}
	vlsOffset := p.UInt16LE(offset)
	vlsLength := p.UInt16LE(offset + 2)
	if vlsLength > 4096 {
		vlsLength = 4096
	}
	if vlsOffset == 0 || vlsLength == 0 {
		return ""
	}
	if size < vlsOffset+vlsLength {
		return ""
	}

	if p.Byte(int(vlsOffset+vlsLength)-1) == 0 {
		return p.String(int(vlsOffset), int(vlsLength)-1)
	} else {
		return p.String(int(vlsOffset), int(vlsLength))
	}
}

// SetStringVLSPointer replaces existing VLS if new one fits, otherwise appends to end possibly growing
// the existing allocation in order to do so.
func SetStringVLSPointer(b *VLSPointerBuilder, baseSize, bounds uint16, offset int, value string) {
	if baseSize < bounds {
		return
	}
	vlsOffset := int(b.p.UInt16LE(offset))
	vlsLength := int(b.p.UInt16LE(offset + 2))
	if vlsLength > 4096 {
		vlsLength = 4096
	}
	newLength := len(value) + 1
	if newLength > 4096 {
		value = value[0:4095]
		newLength = 4096
	}
	if vlsLength >= newLength {
		// Set new length
		b.p.SetUInt16LE(offset+2, uint16(len(value)+1))
		b.p.SetString(vlsOffset, value)
		b.p.SetByte(vlsOffset+len(value), 0)
		return
	}

	newSize := int(b.Size()) + newLength
	if newSize > math.MaxUint16 {
		return
	}
	vlsOffset = int(b.Size())
	if b.cap < newSize {
		b.Extend(newLength)
		if b.p == 0 {
			return
		}
	}

	// Set new size
	b.p.SetUInt16LE(0, uint16(newSize))
	// Set VLS offset
	b.p.SetUInt16LE(offset, b.Size())
	// Set VLS length
	b.p.SetUInt16LE(offset+2, uint16(newLength))
	// Set string
	b.p.SetString(int(b.Size()), value)
	// Set null terminator
	b.p.SetByte(int(b.Size())+len(value), 0)
	// Update to new length
	b.p.SetUInt16LE(0, uint16(newSize))
}

// SetStringVLS replaces existing VLS if new one fits, otherwise appends to end possibly growing
// the existing allocation in order to do so.
func SetStringVLS(b *VLSBuilder, baseSize, bounds uint16, offset int, value string) {
	if baseSize < bounds {
		return
	}
	vlsOffset := int(b.p.UInt16LE(offset))
	vlsLength := int(b.p.UInt16LE(offset + 2))
	if vlsLength > 4096 {
		vlsLength = 4096
	}
	newLength := len(value) + 1
	if newLength > 4096 {
		value = value[0:4095]
		newLength = 4096
	}
	if vlsLength >= newLength {
		// Set new length
		b.p.SetUInt16LE(offset+2, uint16(len(value)+1))
		b.p.SetString(vlsOffset, value)
		b.p.SetByte(vlsOffset+len(value), 0)
		return
	}

	newSize := int(b.Size()) + newLength
	if newSize > math.MaxUint16 {
		return
	}
	vlsOffset = int(b.Size())
	if b.cap < newSize {
		b.Extend(newLength)
		if b.p == 0 {
			return
		}
	}

	// Set new size
	b.p.SetUInt16LE(0, uint16(newSize))
	// Set VLS offset
	b.p.SetUInt16LE(offset, b.Size())
	// Set VLS length
	b.p.SetUInt16LE(offset+2, uint16(newLength))
	// Set string
	b.p.SetString(int(b.Size()), value)
	// Set null terminator
	b.p.SetByte(int(b.Size())+len(value), 0)
	// Update to new length
	b.p.SetUInt16LE(0, uint16(newSize))
}
