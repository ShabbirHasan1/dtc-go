package codegen

import (
	"fmt"
	"testing"
)

func TestParseProto(t *testing.T) {
	proto, err := ParseProto("testdata/DTCProtocol.proto")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(proto)
}
