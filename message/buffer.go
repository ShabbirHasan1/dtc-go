package message

import (
	"github.com/moontrade/nogc"
	"math"
	"strings"
	"unsafe"
)

func CopyFixed(dst, src nogc.Pointer) bool {
	if dst == 0 || src == 0 {
		return false
	}
	var (
		dstSize = dst.Uint16LE(0)
		dstType = dst.Uint16LE(2)
		srcSize = src.Uint16LE(0)
		srcType = src.Uint16LE(2)
	)
	if dstType != srcType {
		return false
	}
	if dstSize < 4 || srcSize < 4 {
		return false
	}
	if dstSize < srcSize {
		nogc.Copy((dst + 4).Unsafe(), (src + 4).Unsafe(), uintptr(dstSize)-4)
	} else if dstSize > srcSize {
		nogc.Copy((dst + 4).Unsafe(), (src + 4).Unsafe(), uintptr(srcSize)-4)
	} else {
		nogc.Copy(dst.Unsafe(), src.Unsafe(), uintptr(dstSize))
	}
	return true
}

//func CopyVLS(dst, src nogc.Pointer) bool {
//	if dst == 0 || src == 0 {
//		return false
//	}
//	var (
//		dstSize = dst.Uint16LE(0)
//		dstType = dst.Uint16LE(2)
//		dstBase = src.Uint16LE(4)
//		srcSize = src.Uint16LE(0)
//		srcType = src.Uint16LE(2)
//		srcBase = src.Uint16LE(4)
//	)
//	if dstType != srcType {
//		return false
//	}
//	if dstSize < 7 || srcSize < 7 {
//		return false
//	}
//	if dstSize < srcSize {
//		nogc.Copy((dst + 4).Unsafe(), (src + 4).Unsafe(), uintptr(dstSize)-4)
//	} else if dstSize > srcSize {
//		nogc.Copy((dst + 4).Unsafe(), (src + 4).Unsafe(), uintptr(srcSize)-4)
//	} else {
//		nogc.Copy(dst.Unsafe(), src.Unsafe(), uintptr(dstSize))
//	}
//	return true
//}

func Size(ptr unsafe.Pointer) uint16 {
	if ptr == nil {
		return 0
	}
	return nogc.Pointer(ptr).Uint16LE(0)
}

func Type(ptr unsafe.Pointer) uint16 {
	if ptr == nil {
		return 0
	}
	return nogc.Pointer(ptr).Uint16LE(2)
}

func BaseSize(ptr unsafe.Pointer) uint16 {
	if ptr == nil {
		return 0
	}
	return nogc.Pointer(ptr).Uint16LE(4)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Bool
////////////////////////////////////////////////////////////////////////////////////////////

func BoolFixed(p nogc.Pointer, bounds uint16, offset int) bool {
	if p == 0 || p.Uint16LE(0) < bounds {
		return false
	}
	return p.Uint8(offset) != 0
}
func BoolVLS(p nogc.Pointer, bounds uint16, offset int) bool {
	if p == 0 || p.Uint16LE(4) < bounds {
		return false
	}
	return p.Uint8(offset) != 0
}
func SetBoolFixed(p nogc.Pointer, bounds uint16, offset int, value bool) {
	if p == 0 || p.Uint16LE(0) < bounds {
		return
	}
	if value {
		p.SetUint8(offset, 1)
	} else {
		p.SetUint8(offset, 0)
	}
}
func SetBoolVLS(p nogc.Pointer, bounds uint16, offset int, value bool) {
	if p == 0 || p.Uint16LE(4) < bounds {
		return
	}
	if value {
		p.SetUint8(offset, 1)
	} else {
		p.SetUint8(offset, 0)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////
// Int8
////////////////////////////////////////////////////////////////////////////////////////////

func Int8Fixed(p nogc.Pointer, bounds uint16, offset int) int8 {
	if p == 0 || p.Uint16LE(0) < bounds {
		return 0
	}
	return p.Int8(offset)
}
func Int8VLS(p nogc.Pointer, bounds uint16, offset int) int8 {
	if p == 0 || p.Uint16LE(4) < bounds {
		return 0
	}
	return p.Int8(offset)
}
func SetInt8Fixed(p nogc.Pointer, bounds uint16, offset int, value int8) {
	if p == 0 || p.Uint16LE(0) < bounds {
		return
	}
	p.SetInt8(offset, value)
}
func SetInt8VLS(p nogc.Pointer, bounds uint16, offset int, value int8) {
	if p == 0 || p.Uint16LE(4) < bounds {
		return
	}
	p.SetInt8(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Uint8
////////////////////////////////////////////////////////////////////////////////////////////

func Uint8Fixed(p nogc.Pointer, bounds uint16, offset int) uint8 {
	if p == 0 || p.Uint16LE(0) < bounds {
		return 0
	}
	return p.Uint8(offset)
}
func Uint8VLS(p nogc.Pointer, bounds uint16, offset int) uint8 {
	if p == 0 || p.Uint16LE(4) < bounds {
		return 0
	}
	return p.Uint8(offset)
}
func SetUint8Fixed(p nogc.Pointer, bounds uint16, offset int, value uint8) {
	if p == 0 || p.Uint16LE(0) < bounds {
		return
	}
	p.SetUint8(offset, value)
}
func SetUint8VLS(p nogc.Pointer, bounds uint16, offset int, value uint8) {
	if p == 0 || p.Uint16LE(4) < bounds {
		return
	}
	p.SetUint8(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Int16
////////////////////////////////////////////////////////////////////////////////////////////

func Int16Fixed(p nogc.Pointer, bounds uint16, offset int) int16 {
	if p == 0 || p.Uint16LE(0) < bounds {
		return 0
	}
	return p.Int16LE(offset)
}
func Int16VLS(p nogc.Pointer, bounds uint16, offset int) int16 {
	if p == 0 || p.Uint16LE(4) < bounds {
		return 0
	}
	return p.Int16LE(offset)
}
func SetInt16Fixed(p nogc.Pointer, bounds uint16, offset int, value int16) {
	if p == 0 || p.Uint16LE(0) < bounds {
		return
	}
	p.SetInt16LE(offset, value)
}
func SetInt16VLS(p nogc.Pointer, bounds uint16, offset int, value int16) {
	if p != 0 || p.Uint16LE(4) < bounds {
		return
	}
	p.SetInt16LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Uint16
////////////////////////////////////////////////////////////////////////////////////////////

func Uint16Fixed(p nogc.Pointer, bounds uint16, offset int) uint16 {
	if p == 0 || p.Uint16LE(0) < bounds {
		return 0
	}
	return p.Uint16LE(offset)
}
func Uint16VLS(p nogc.Pointer, bounds uint16, offset int) uint16 {
	if p == 0 || p.Uint16LE(4) < bounds {
		return 0
	}
	return p.Uint16LE(offset)
}
func SetUint16Fixed(p nogc.Pointer, bounds uint16, offset int, value uint16) {
	if p == 0 || p.Uint16LE(0) < bounds {
		return
	}
	p.SetUint16LE(offset, value)
}
func SetUint16VLS(p nogc.Pointer, bounds uint16, offset int, value uint16) {
	if p == 0 || p.Uint16LE(4) < bounds {
		return
	}
	p.SetUint16LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Int32
////////////////////////////////////////////////////////////////////////////////////////////

func Int32Fixed(p nogc.Pointer, bounds uint16, offset int) int32 {
	if p == 0 || p.Uint16LE(0) < bounds {
		return 0
	}
	return p.Int32LE(offset)
}
func Int32VLS(p nogc.Pointer, bounds uint16, offset int) int32 {
	if p == 0 || p.Uint16LE(4) < bounds {
		return 0
	}
	return p.Int32LE(offset)
}
func SetInt32Fixed(p nogc.Pointer, bounds uint16, offset int, value int32) {
	if p == 0 || p.Uint16LE(0) < bounds {
		return
	}
	p.SetInt32LE(offset, value)
}
func SetInt32VLS(p nogc.Pointer, bounds uint16, offset int, value int32) {
	if p == 0 || p.Uint16LE(4) < bounds {
		return
	}
	p.SetInt32LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Uint32
////////////////////////////////////////////////////////////////////////////////////////////

func Uint32Fixed(p nogc.Pointer, bounds uint16, offset int) uint32 {
	if p == 0 || p.Uint16LE(0) < bounds {
		return 0
	}
	return p.Uint32LE(offset)
}
func Uint32VLS(p nogc.Pointer, bounds uint16, offset int) uint32 {
	if p == 0 || p.Uint16LE(4) < bounds {
		return 0
	}
	return p.Uint32LE(offset)
}
func SetUint32Fixed(p nogc.Pointer, bounds uint16, offset int, value uint32) {
	if p == 0 || p.Uint16LE(0) < bounds {
		return
	}
	p.SetUint32LE(offset, value)
}
func SetUint32VLS(p nogc.Pointer, bounds uint16, offset int, value uint32) {
	if p == 0 || p.Uint16LE(4) < bounds {
		return
	}
	p.SetUint32LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Int64
////////////////////////////////////////////////////////////////////////////////////////////

func Int64Fixed(p nogc.Pointer, bounds uint16, offset int) int64 {
	if p == 0 || p.Uint16LE(0) < bounds {
		return 0
	}
	return p.Int64LE(offset)
}
func Int64VLS(p nogc.Pointer, bounds uint16, offset int) int64 {
	if p == 0 || p.Uint16LE(4) < bounds {
		return 0
	}
	return p.Int64LE(offset)
}
func SetInt64Fixed(p nogc.Pointer, bounds uint16, offset int, value int64) {
	if p == 0 || p.Uint16LE(0) < bounds {
		return
	}
	p.SetInt64LE(offset, value)
}
func SetInt64VLS(p nogc.Pointer, bounds uint16, offset int, value int64) {
	if p == 0 || p.Uint16LE(4) < bounds {
		return
	}
	p.SetInt64LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Uint64
////////////////////////////////////////////////////////////////////////////////////////////

func Uint64Fixed(p nogc.Pointer, bounds uint16, offset int) uint64 {
	if p == 0 || p.Uint16LE(0) < bounds {
		return 0
	}
	return p.Uint64LE(offset)
}
func Uint64VLS(p nogc.Pointer, bounds uint16, offset int) uint64 {
	if p == 0 || p.Uint16LE(4) < bounds {
		return 0
	}
	return p.Uint64LE(offset)
}
func SetUint64Fixed(p nogc.Pointer, bounds uint16, offset int, value uint64) {
	if p == 0 || p.Uint16LE(0) < bounds {
		return
	}
	p.SetUint64LE(offset, value)
}
func SetUint64VLS(p nogc.Pointer, bounds uint16, offset int, value uint64) {
	if p == 0 || p.Uint16LE(4) < bounds {
		return
	}
	p.SetUint64LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Float32
////////////////////////////////////////////////////////////////////////////////////////////

func Float32Fixed(p nogc.Pointer, bounds uint16, offset int) float32 {
	if p == 0 || p.Uint16LE(0) < bounds {
		return 0
	}
	return p.Float32LE(offset)
}
func Float32VLS(p nogc.Pointer, bounds uint16, offset int) float32 {
	if p == 0 || p.Uint16LE(4) < bounds {
		return 0
	}
	return p.Float32LE(offset)
}
func SetFloat32Fixed(p nogc.Pointer, bounds uint16, offset int, value float32) {
	if p == 0 || p.Uint16LE(0) < bounds {
		return
	}
	p.SetFloat32LE(offset, value)
}
func SetFloat32VLS(p nogc.Pointer, bounds uint16, offset int, value float32) {
	if p == 0 || p.Uint16LE(4) < bounds {
		return
	}
	p.SetFloat32LE(offset, value)
}

////////////////////////////////////////////////////////////////////////////////////////////
// Float64
////////////////////////////////////////////////////////////////////////////////////////////

func Float64Fixed(p nogc.Pointer, bounds uint16, offset int) float64 {
	if p == 0 || p.Uint16LE(0) < bounds {
		return 0
	}
	return p.Float64LE(offset)
}
func Float64VLS(p nogc.Pointer, bounds uint16, offset int) float64 {
	if p == 0 || p.Uint16LE(4) < bounds {
		return 0
	}
	return p.Float64LE(offset)
}
func SetFloat64Fixed(p nogc.Pointer, bounds uint16, offset int, value float64) {
	if p == 0 || p.Uint16LE(0) < bounds {
		return
	}
	p.SetFloat64LE(offset, value)
}
func SetFloat64VLS(p nogc.Pointer, bounds uint16, offset int, value float64) {
	if p == 0 || p.Uint16LE(4) < bounds {
		return
	}
	p.SetFloat64LE(offset, value)
}

func StringFixed(p nogc.Pointer, bounds uint16, offset int) string {
	if p == 0 {
		return ""
	}
	if p == 0 || p.Uint16LE(0) < bounds {
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
	if p == 0 || p.Uint16LE(0) < bounds {
		return
	}
	maxLength := int(bounds) - offset
	if len(value) >= maxLength {
		value = value[0 : maxLength-1]
	}
	//nogc.Zero(Ptr.Pointer(offset).Unsafe(), uintptr(maxLength))
	p.SetString(offset, value)
	p.SetUint8(offset+len(value), 0)
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
		size     = p.Uint16LE(0)
		baseSize = p.Uint16LE(4)
	)
	if baseSize < bounds {
		return ""
	}
	vlsOffset := p.Uint16LE(offset)
	vlsLength := p.Uint16LE(offset + 2)
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
	vlsOffset := int(p.Uint16LE(offset))
	vlsLength := int(p.Uint16LE(offset + 2))
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
		p.SetUint16LE(offset+2, uint16(len(value)+1))
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
	p.SetUint16LE(0, uint16(newSize))
	// Set VLS offset
	p.SetUint16LE(offset, uint16(vlsOffset))
	// Set VLS length
	p.SetUint16LE(offset+2, uint16(newLength))
	// Set string
	p.SetString(vlsOffset, value)
	// Set null terminator
	p.SetByte(vlsOffset+len(value), 0)
}

// SetStringVLS replaces existing VLS if new one fits, otherwise appends to end possibly growing
// the existing allocation in order to do so.
func SetStringVLS(b *VLSBuilder, bounds uint16, offset int, value string) {
	if b == nil || b.BaseSize() < bounds {
		return
	}

	p := b.Ptr
	vlsOffset := int(p.Uint16LE(offset))
	vlsLength := int(p.Uint16LE(offset + 2))
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
		p.SetUint16LE(offset+2, uint16(len(value)+1))
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
	p.SetUint16LE(0, uint16(newSize))
	// Set VLS offset
	p.SetUint16LE(offset, uint16(vlsOffset))
	// Set VLS length
	p.SetUint16LE(offset+2, uint16(newLength))
	// Set string
	p.SetString(vlsOffset, value)
	// Set null terminator
	p.SetByte(vlsOffset+len(value), 0)
}

// SetStringVLS replaces existing VLS if new one fits, otherwise appends to end possibly growing
// the existing allocation in order to do so.
func SetStringVLS0(builder nogc.Pointer, bounds uint16, offset int, value string) {
	if builder == 0 {
		return
	}
	p := nogc.Pointer(builder.Uintptr(0))
	if p == 0 {
		return
	}
	var (
		isAlloc  = builder.Uintptr(int(unsafe.Offsetof(VLSBuilder{}.ref))) == 0
		c        = builder.Int(int(unsafe.Offsetof(VLSBuilder{}.cap)))
		size     = p.Uint16LE(0)
		baseSize = p.Uint16LE(4)
	)
	if baseSize < bounds {
		return
	}

	var (
		vlsOffset = p.Uint16LE(offset)
		vlsLength = p.Uint16LE(offset + 2)
	)
	newLength := len(value) + 1
	if newLength > 4096 {
		value = value[0:4095]
		newLength = 4096
	}
	if vlsLength >= uint16(newLength) {
		// Set new length
		p.SetUint16LE(offset+2, uint16(len(value)+1))
		p.SetString(int(vlsOffset), value)
		p.SetByte(int(vlsOffset)+len(value), 0)
		return
	}
	newSize := int(size) + newLength
	if newSize > math.MaxUint16 {
		return
	}
	vlsOffset = size
	if c < newSize {
		if isAlloc {
			b := (*VLSPointerBuilder)(builder.Unsafe())
			growBy := b.growBy
			if growBy < int(vlsLength) {
				growBy = int(vlsLength)
			}
			newPtr, newCap := nogc.ReallocCap(p, uintptr(c+growBy))
			if newPtr == 0 {
				return
			}
			p = newPtr
			b.Ptr = newPtr
			b.cap = int(newCap)
		} else {
			b := (*VLSBuilder)(builder.Unsafe())
			growBy := b.growBy
			if growBy < int(vlsLength) {
				growBy = int(vlsLength)
			}
			newCap := uintptr(c + growBy)

			old := b.Ptr
			newRef := gcAlloc(newCap)
			if newRef == nil {
				return
			}

			b.Ptr = nogc.Pointer(newRef)
			p = b.Ptr
			b.cap = int(newCap)
			b.ref = newRef
			nogc.Copy(newRef, old.Unsafe(), uintptr(size))
		}
	}

	// Set new size
	p.SetUint16LE(0, uint16(newSize))
	// Set VLS offset
	p.SetUint16LE(offset, vlsOffset)
	// Set VLS length
	p.SetUint16LE(offset+2, uint16(newLength))
	// Set string
	p.SetString(int(vlsOffset), value)
	// Set null terminator
	p.SetByte(int(vlsOffset)+len(value), 0)
}
