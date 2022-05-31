package main

import (
	"github.com/moontrade/dtc-go/internal/unsafecgo"
)

func main() {
	//cgo.CGO()
	unsafecgo.NonBlocking((*byte)(nil), 0, 0)
}
