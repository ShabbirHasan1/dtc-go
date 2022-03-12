package golang

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen"
	"testing"
)

func TestGenerator(t *testing.T) {
	schema, err := codegen.LoadSchemaFromCHeaders("../testdata/docs.json", "../testdata/DTCProtocol.proto", "../testdata/DTCProtocol.h", "../testdata/DTCProtocolVLS.h")
	if err != nil {
		t.Fatal(err)
	}
	g, err := NewGenerator(DefaultConfig("../../model"), schema)
	if err != nil {
		t.Fatal(err)
	}
	err = g.Run()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(g)
}

func TestGeneratorNonStandard(t *testing.T) {
	schema, err := codegen.LoadSchemaFromCHeaders("../testdata/docs.json", "../testdata/DTCProtocol.proto", "../testdata/DTCProtocol.h", "../testdata/DTCProtocolVLS.h", "../testdata/DTCProtocol_NonStandard.h")
	if err != nil {
		t.Fatal(err)
	}
	c := DefaultConfig("../../model/sc")
	c.NonStandard = true
	c.RootPackage = "github.com/moontrade/dtc-go/model/sc"
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
