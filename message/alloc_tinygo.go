//go:build tinygo

package message

func gcAlloc(size uintptr) unsafe.Pointer {
	b := make([]byte, size)
	return unsafe.Pointer(&b[0])
}
