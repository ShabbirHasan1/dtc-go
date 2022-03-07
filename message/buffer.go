package message

import (
	"github.com/moontrade/nogc"
	"math"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////
// Int8
////////////////////////////////////////////////////////////////////////////////////////////

func Int8Fixed(p nogc.Pointer, bounds uint16, offset int) int8 {
	if p == 0 || p.UInt16LE(0) < bounds {
		return 0
	}
	return p.Int8(offset)
}
func Int8VLS(p nogc.Pointer, bounds uint16, offset int) int8 {
	if p == 0 || p.UInt16LE(4) < bounds {
		return 0
	}
	return p.Int8(offset)
}
func SetInt8Fixed(p nogc.Pointer, bounds uint16, offset int, value int8) {
	if p == 0 || p.UInt16LE(0) < bounds {
		return
	}
	p.SetInt8(offset, value)
}
func SetInt8VLS(p nogc.Pointer, bounds uint16, offset int, value int8) {
	if p == 0 || p.UInt16LE(4) < bounds {
		return
	}
	p.SetInt8(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Uint8
////////////////////////////////////////////////////////////////////////////////////////////

func Uint8Fixed(p nogc.Pointer, bounds uint16, offset int) uint8 {
	if p == 0 || p.UInt16LE(0) < bounds {
		return 0
	}
	return p.UInt8(offset)
}
func Uint8VLS(p nogc.Pointer, bounds uint16, offset int) uint8 {
	if p == 0 || p.UInt16LE(4) < bounds {
		return 0
	}
	return p.UInt8(offset)
}
func SetUint8Fixed(p nogc.Pointer, bounds uint16, offset int, value uint8) {
	if p == 0 || p.UInt16LE(0) < bounds {
		return
	}
	p.SetUInt8(offset, value)
}
func SetUint8VLS(p nogc.Pointer, bounds uint16, offset int, value uint8) {
	if p == 0 || p.UInt16LE(4) < bounds {
		return
	}
	p.SetUInt8(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Int16
////////////////////////////////////////////////////////////////////////////////////////////

func Int16Fixed(p nogc.Pointer, bounds uint16, offset int) int16 {
	if p == 0 || p.UInt16LE(0) < bounds {
		return 0
	}
	return p.Int16LE(offset)
}
func Int16VLS(p nogc.Pointer, bounds uint16, offset int) int16 {
	if p == 0 || p.UInt16LE(4) < bounds {
		return 0
	}
	return p.Int16LE(offset)
}
func SetInt16Fixed(p nogc.Pointer, bounds uint16, offset int, value int16) {
	if p == 0 || p.UInt16LE(0) < bounds {
		return
	}
	p.SetInt16LE(offset, value)
}
func SetInt16VLS(p nogc.Pointer, bounds uint16, offset int, value int16) {
	if p != 0 || p.UInt16LE(4) < bounds {
		return
	}
	p.SetInt16LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Uint16
////////////////////////////////////////////////////////////////////////////////////////////

func Uint16Fixed(p nogc.Pointer, bounds uint16, offset int) uint16 {
	if p == 0 || p.UInt16LE(0) < bounds {
		return 0
	}
	return p.UInt16LE(offset)
}
func Uint16VLS(p nogc.Pointer, bounds uint16, offset int) uint16 {
	if p == 0 || p.UInt16LE(4) < bounds {
		return 0
	}
	return p.UInt16LE(offset)
}
func SetUint16Fixed(p nogc.Pointer, bounds uint16, offset int, value uint16) {
	if p == 0 || p.UInt16LE(0) < bounds {
		return
	}
	p.SetUInt16LE(offset, value)
}
func SetUint16VLS(p nogc.Pointer, bounds uint16, offset int, value uint16) {
	if p == 0 || p.UInt16LE(4) < bounds {
		return
	}
	p.SetUInt16LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Int32
////////////////////////////////////////////////////////////////////////////////////////////

func Int32Fixed(p nogc.Pointer, bounds uint16, offset int) int32 {
	if p == 0 || p.UInt16LE(0) < bounds {
		return 0
	}
	return p.Int32LE(offset)
}
func Int32VLS(p nogc.Pointer, bounds uint16, offset int) int32 {
	if p == 0 || p.UInt16LE(4) < bounds {
		return 0
	}
	return p.Int32LE(offset)
}
func SetInt32Fixed(p nogc.Pointer, bounds uint16, offset int, value int32) {
	if p == 0 || p.UInt16LE(0) < bounds {
		return
	}
	p.SetInt32LE(offset, value)
}
func SetInt32VLS(p nogc.Pointer, bounds uint16, offset int, value int32) {
	if p == 0 || p.UInt16LE(4) < bounds {
		return
	}
	p.SetInt32LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Uint32
////////////////////////////////////////////////////////////////////////////////////////////

func Uint32Fixed(p nogc.Pointer, bounds uint16, offset int) uint32 {
	if p == 0 || p.UInt16LE(0) < bounds {
		return 0
	}
	return p.UInt32LE(offset)
}
func Uint32VLS(p nogc.Pointer, bounds uint16, offset int) uint32 {
	if p == 0 || p.UInt16LE(4) < bounds {
		return 0
	}
	return p.UInt32LE(offset)
}
func SetUint32Fixed(p nogc.Pointer, bounds uint16, offset int, value uint32) {
	if p == 0 || p.UInt16LE(0) < bounds {
		return
	}
	p.SetUInt32LE(offset, value)
}
func SetUint32VLS(p nogc.Pointer, bounds uint16, offset int, value uint32) {
	if p == 0 || p.UInt16LE(4) < bounds {
		return
	}
	p.SetUInt32LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Int64
////////////////////////////////////////////////////////////////////////////////////////////

func Int64Fixed(p nogc.Pointer, bounds uint16, offset int) int64 {
	if p == 0 || p.UInt16LE(0) < bounds {
		return 0
	}
	return p.Int64LE(offset)
}
func Int64VLS(p nogc.Pointer, bounds uint16, offset int) int64 {
	if p == 0 || p.UInt16LE(4) < bounds {
		return 0
	}
	return p.Int64LE(offset)
}
func SetInt64Fixed(p nogc.Pointer, bounds uint16, offset int, value int64) {
	if p == 0 || p.UInt16LE(0) < bounds {
		return
	}
	p.SetInt64LE(offset, value)
}
func SetInt64VLS(p nogc.Pointer, bounds uint16, offset int, value int64) {
	if p == 0 || p.UInt16LE(4) < bounds {
		return
	}
	p.SetInt64LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Uint64
////////////////////////////////////////////////////////////////////////////////////////////

func Uint64Fixed(p nogc.Pointer, bounds uint16, offset int) uint64 {
	if p == 0 || p.UInt16LE(0) < bounds {
		return 0
	}
	return p.UInt64LE(offset)
}
func Uint64VLS(p nogc.Pointer, bounds uint16, offset int) uint64 {
	if p == 0 || p.UInt16LE(4) < bounds {
		return 0
	}
	return p.UInt64LE(offset)
}
func SetUint64Fixed(p nogc.Pointer, bounds uint16, offset int, value uint64) {
	if p == 0 || p.UInt16LE(0) < bounds {
		return
	}
	p.SetUInt64LE(offset, value)
}
func SetUint64VLS(p nogc.Pointer, bounds uint16, offset int, value uint64) {
	if p == 0 || p.UInt16LE(4) < bounds {
		return
	}
	p.SetUInt64LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Float32
////////////////////////////////////////////////////////////////////////////////////////////

func Float32Fixed(p nogc.Pointer, bounds uint16, offset int) float32 {
	if p == 0 || p.UInt16LE(0) < bounds {
		return 0
	}
	return p.Float32LE(offset)
}
func Float32VLS(p nogc.Pointer, bounds uint16, offset int) float32 {
	if p == 0 || p.UInt16LE(4) < bounds {
		return 0
	}
	return p.Float32LE(offset)
}
func SetFloat32Fixed(p nogc.Pointer, bounds uint16, offset int, value float32) {
	if p == 0 || p.UInt16LE(0) < bounds {
		return
	}
	p.SetFloat32LE(offset, value)
}
func SetFloat32VLS(p nogc.Pointer, bounds uint16, offset int, value float32) {
	if p == 0 || p.UInt16LE(4) < bounds {
		return
	}
	p.SetFloat32LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Float64
////////////////////////////////////////////////////////////////////////////////////////////

func Float64Fixed(p nogc.Pointer, bounds uint16, offset int) float64 {
	if p == 0 || p.UInt16LE(0) < bounds {
		return 0
	}
	return p.Float64LE(offset)
}
func Float64VLS(p nogc.Pointer, bounds uint16, offset int) float64 {
	if p == 0 || p.UInt16LE(4) < bounds {
		return 0
	}
	return p.Float64LE(offset)
}
func SetFloat64Fixed(p nogc.Pointer, bounds uint16, offset int, value float64) {
	if p == 0 || p.UInt16LE(0) < bounds {
		return
	}
	p.SetFloat64LE(offset, value)
}
func SetFloat64VLS(p nogc.Pointer, bounds uint16, offset int, value float64) {
	if p == 0 || p.UInt16LE(4) < bounds {
		return
	}
	p.SetFloat64LE(offset, value)
}

func StringFixed(p nogc.Pointer, bounds uint16, offset int) string {
	if p == 0 {
		return ""
	}
	if p == 0 || p.UInt16LE(0) < bounds {
		return ""
	}
	s := p.String(offset, int(bounds)-offset)
	index := strings.IndexByte(s, 0)
	if index > -1 {
		return s[0:index]
	}
	return s
}

func SetStringFixed(p nogc.Pointer, bounds uint16, offset int, value string) {
	if p == 0 || p.UInt16LE(0) < bounds {
		return
	}
	maxLength := int(bounds) - offset
	if len(value) >= maxLength {
		value = value[0 : maxLength-1]
	}
	//nogc.Zero(Ptr.Pointer(offset).Unsafe(), uintptr(maxLength))
	p.SetString(offset, value)
	p.SetUInt8(offset+len(value), 0)
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
func SetStringVLSPointer(b *VLSPointerBuilder, bounds uint16, offset int, value string) {
	if b == nil || b.BaseSize() < bounds {
		return
	}
	p := b.Ptr
	vlsOffset := int(p.UInt16LE(offset))
	vlsLength := int(p.UInt16LE(offset + 2))
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
		p.SetUInt16LE(offset+2, uint16(len(value)+1))
		p.SetString(vlsOffset, value)
		p.SetByte(vlsOffset+len(value), 0)
		return
	}

	newSize := int(b.Size()) + newLength
	if newSize > math.MaxUint16 {
		return
	}
	vlsOffset = int(b.Size())
	if b.cap < newSize {
		b.Extend(newLength)
		p = b.Ptr
		if p == 0 {
			return
		}
	}

	// Set new size
	p.SetUInt16LE(0, uint16(newSize))
	// Set VLS offset
	p.SetUInt16LE(offset, b.Size())
	// Set VLS length
	p.SetUInt16LE(offset+2, uint16(newLength))
	// Set string
	p.SetString(int(b.Size()), value)
	// Set null terminator
	p.SetByte(int(b.Size())+len(value), 0)
	// Update to new length
	p.SetUInt16LE(0, uint16(newSize))
}

// SetStringVLS replaces existing VLS if new one fits, otherwise appends to end possibly growing
// the existing allocation in order to do so.
func SetStringVLS(b *VLSBuilder, bounds uint16, offset int, value string) {
	if b == nil || b.BaseSize() < bounds {
		return
	}

	p := b.Ptr
	vlsOffset := int(p.UInt16LE(offset))
	vlsLength := int(p.UInt16LE(offset + 2))
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
		p.SetUInt16LE(offset+2, uint16(len(value)+1))
		p.SetString(vlsOffset, value)
		p.SetByte(vlsOffset+len(value), 0)
		return
	}

	newSize := int(b.Size()) + newLength
	if newSize > math.MaxUint16 {
		return
	}
	vlsOffset = int(b.Size())
	if b.cap < newSize {
		b.Extend(newLength)
		p = b.Ptr
		if p == 0 {
			return
		}
	}

	// Set new size
	p.SetUInt16LE(0, uint16(newSize))
	// Set VLS offset
	p.SetUInt16LE(offset, b.Size())
	// Set VLS length
	p.SetUInt16LE(offset+2, uint16(newLength))
	// Set string
	p.SetString(int(b.Size()), value)
	// Set null terminator
	p.SetByte(int(b.Size())+len(value), 0)
	// Update to new length
	p.SetUInt16LE(0, uint16(newSize))
}
