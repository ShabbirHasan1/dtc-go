package dtc

import "unsafe"

type VLS struct {
	Offset uint16
	Length uint16
}

func (vls *VLS) Value(size *uint16, baseSize uint16) string {
	if vls == nil || size == nil {
		return ""
	}
	if vls.Offset == 0 || vls.Length < 2 {
		return ""
	}
	offset := int64(uintptr(unsafe.Pointer(vls))) - int64(uintptr(unsafe.Pointer(size)))
	if offset < 2 {
		return ""
	}
	if offset > int64(baseSize)-int64(unsafe.Sizeof(VLS{})) {
		return ""
	}
	if vls.Offset+vls.Length > *size {
		return ""
	}
	length := int(vls.Length)
	if length > 4096 {
		length = 4096
	}

	ptr := unsafe.Add(unsafe.Pointer(size), vls.Offset)
	if *(*byte)(unsafe.Add(ptr, length-1)) == 0 {
		return *(*string)(unsafe.Pointer(&_string{
			Data: uintptr(ptr),
			Len:  length - 1,
		}))
	} else {
		return *(*string)(unsafe.Pointer(&_string{
			Data: uintptr(ptr),
			Len:  length,
		}))
	}
}

func Str(vls *VLS, size *uint16, baseSize uint16) string {
	return vls.Value(size, baseSize)
}

func getVLSField(messageSizeField *uint16, baseStructureSizeField uint16, vlsStringField *VLS, vlsFieldOffset uintptr) string {
	if baseStructureSizeField < uint16(vlsFieldOffset)+uint16(unsafe.Sizeof(VLS{})) {
		return ""
	} else if vlsStringField.Offset == 0 || vlsStringField.Length <= 1 {
		return ""
	} else if vlsStringField.Offset+vlsStringField.Length > *messageSizeField {
		return ""
	}
	length := int(vlsStringField.Length)
	if length > 4096 {
		length = 4096
	}

	ptr := unsafe.Add(unsafe.Pointer(messageSizeField), vlsStringField.Offset)

	if *(*byte)(unsafe.Add(ptr, length-1)) == 0 {
		return *(*string)(unsafe.Pointer(&_string{
			Data: uintptr(ptr),
			Len:  length - 1,
		}))
	} else {
		return *(*string)(unsafe.Pointer(&_string{
			Data: uintptr(ptr),
			Len:  length,
		}))
	}
}
