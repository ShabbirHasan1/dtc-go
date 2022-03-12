package golang

import (
	"fmt"
	"strconv"
)

type Writer struct {
	b []byte
}

func (w *Writer) WriteByteSlice(b []byte) {
	w.b = append(w.b, "[]byte{"...)
	for i, a := range b {
		w.b = append(w.b, strconv.Itoa(int(a))...)
		if i < len(b)-1 {
			w.b = append(w.b, ',', ' ')
		}
	}
	w.b = append(w.b, '}')
}

func (w *Writer) Indent(times int, format string, args ...interface{}) {
	for i := 0; i < times; i++ {
		w.b = append(w.b, "    "...)
	}
	w.b = append(w.b, fmt.Sprintf(format, args...)...)
}

func (w *Writer) Write(format string, args ...interface{}) {
	w.b = append(w.b, fmt.Sprintf(format, args...)...)
}

func (w *Writer) IndentLine(times int, format string, args ...interface{}) {
	for i := 0; i < times; i++ {
		w.b = append(w.b, "    "...)
	}
	w.b = append(w.b, fmt.Sprintf(format, args...)...)
	w.b = append(w.b, "\n"...)
}

func (w *Writer) Line(format string, args ...interface{}) {
	w.b = append(w.b, fmt.Sprintf(format, args...)...)
	w.b = append(w.b, "\n"...)
}
