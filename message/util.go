package message

import "unsafe"

type _bytes struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

func CloneString(s string) string {
	to := make([]byte, len(s))
	to = append(to, s...)
	return string(to)
}
