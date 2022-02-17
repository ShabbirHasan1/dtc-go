package codegen

import (
	"fmt"
	"os"
	"testing"
)

func TestCPPLayout(t *testing.T) {
	namespaces := NewNamespaces()
	err := namespaces.AddHeaders("testdata/DTCProtocol.h", "testdata/DTCProtocolVLS.h", "testdata/DTCProtocol_NonStandard.h")
	if err != nil {
		t.Fatal(err)
	}

	s := namespaces.Namespaces["DTC_VLS"].StructsByName["s_Logoff"]
	s = namespaces.Namespaces["DTC_VLS"].StructsByName["s_TradeOrder"]
	fmt.Println(s)
	data, err := os.ReadFile("testdata/DTCLayout.txt")
	if err != nil {
		t.Fatal(err)
	}
	namespaces.PrintCPPLayoutErrors(data)
}

func TestPrintCPPLayout(t *testing.T) {
	namespaces := NewNamespaces()
	err := namespaces.AddHeaders("testdata/DTCProtocol.h", "testdata/DTCProtocolVLS.h", "testdata/DTCProtocol_NonStandard.h")
	if err != nil {
		t.Fatal(err)
	}
	namespaces.PrintCPPLayoutCode()
}
