package schema

import (
	"os"
	"testing"
)

func TestCPPLayout(t *testing.T) {
	schema := NewSchema()
	err := schema.AddCHeaders("../testdata/DTCProtocol.h", "../testdata/DTCProtocolVLS.h", "../testdata/DTCProtocol_NonStandard.h")
	if err != nil {
		t.Fatal(err)
	}
	if err = schema.Validate(); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile("testdata/DTCLayout.txt")
	if err != nil {
		t.Fatal(err)
	}
	if err = printCPPLayoutErrors(schema, data); err != nil {
		t.Fatal(err)
	}
}

func TestPrintCPPLayout(t *testing.T) {
	schema := NewSchema()
	err := schema.AddCHeaders("../testdata/DTCProtocol.h", "../testdata/DTCProtocolVLS.h", "../testdata/DTCProtocol_NonStandard.h")
	if err != nil {
		t.Fatal(err)
	}
	printCPPLayoutCode(schema)
}
