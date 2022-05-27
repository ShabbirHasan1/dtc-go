package rust

import (
	"crypto/tls"
	"fmt"
	"testing"

	"github.com/moontrade/dtc-go/codegen/schema"
)

func TestDialDenali(t *testing.T) {
	config := &tls.Config{
		//Rand:                        nil,
		//Time:                        nil,
		//Certificates:                nil,
		//NameToCertificate:           nil,
		//GetCertificate:              nil,
		//GetClientCertificate:        nil,
		//GetConfigForClient:          nil,
		//VerifyPeerCertificate:       nil,
		//VerifyConnection:            nil,
		//RootCAs:                     nil,
		//NextProtos:                  nil,
		ServerName: "*.sierracharts.com",
		//ClientAuth:                  0,
		//ClientCAs:                   nil,
		InsecureSkipVerify: false,
		//CipherSuites:                nil,
		//PreferServerCipherSuites:    false,
		//SessionTicketsDisabled:      false,
		//SessionTicketKey:            [32]byte{},
		//ClientSessionCache:          nil,
		//MinVersion:                  0,
		//MaxVersion:                  0,
		//CurvePreferences:            nil,
		//DynamicRecordSizingDisabled: false,
		//Renegotiation:               0,
		//KeyLogWriter:                nil,
	}
	c, err := tls.Dial("tcp4", "futurestrading1.sierracharts.com:11091", config)
	if err != nil {
		t.Fatal(err)
	}
	c.Handshake()
}

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
