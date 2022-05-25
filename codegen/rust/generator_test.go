package rust

import (
	"fmt"
	"testing"

	"github.com/moontrade/dtc-go/codegen/schema"
)

func TestGenerator(t *testing.T) {
	schema, err := schema.LoadSchemaFromCHeaders("../testdata/docs.json", "", "../testdata/DTCProtocol.h", "../testdata/DTCProtocolVLS.h", "../testdata/DTCProtocol_NonStandard.h")
	// schema, err := schema.LoadSchemaFromCHeaders("../testdata/docs.json", "", "../testdata/DTCProtocol.h", "../testdata/DTCProtocolVLS.h")
	if err != nil {
		t.Fatal(err)
	}
	c := DefaultConfig("dtc-rs/src/v8")
	c.NonStandard = true
	g, err := NewGenerator(c, schema)
	if err != nil {
		t.Fatal(err)
	}
	err = g.Run()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGeneratorNonStandard(t *testing.T) {
	schema, err := schema.LoadSchemaFromCHeaders("../testdata/docs.json", "", "../testdata/DTCProtocol.h", "../testdata/DTCProtocolVLS.h", "../testdata/DTCProtocol_NonStandard.h")
	if err != nil {
		t.Fatal(err)
	}
	c := DefaultConfig("out")
	c.NonStandard = true
	g, err := NewGenerator(c, schema)
	if err != nil {
		t.Fatal(err)
	}
	err = g.Run()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(g)
}
