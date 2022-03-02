package message

import "unsafe"

type _bytes struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}
