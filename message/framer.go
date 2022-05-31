package message

import "encoding/binary"

type Framer struct {
	buf []byte
}

func (f *Framer) Push(b []byte, fn func(b []byte)) {
	if len(b) == 0 {
		return
	}
	switch len(f.buf) {
	case 0:
	case 1:
		f.buf = append(f.buf, b[0])
		b = b[1:]
		fallthrough
	default:
		s := int(binary.LittleEndian.Uint16(f.buf))

		if len(f.buf)+len(b) < s {
			return
		}
		f.buf = append(f.buf, b[0:s-len(f.buf)]...)
		b = b[s-len(f.buf):]
		fn(f.buf)
		f.buf = f.buf[0:0]
	}

	for len(b) > 0 {
		if len(b) == 1 {
			f.buf = append(f.buf, b[0])
			return
		}
		s := int(binary.LittleEndian.Uint16(b))
		if len(b) < s {
			f.buf = append(f.buf, b...)
			return
		}
		fn(b[0:s])
		b = b[s:]
	}
}
