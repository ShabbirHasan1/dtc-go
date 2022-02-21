package codegen

import (
	"bytes"
	protoparser "github.com/yoheimuta/go-protoparser/v4"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"os"
)

func ParseProto(path string, option ...protoparser.Option) (*parser.Proto, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return protoparser.Parse(bytes.NewBuffer(b), option...)
}
