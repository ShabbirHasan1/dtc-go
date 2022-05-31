package rust

import (
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"strings"
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

func TestHex(t *testing.T) {
	/*
		EÉÁú@Äº+ZTèÁ-À_~Pÿ ¡88?
		L'a?³hs!
		kamaiur6HRgT6q$hn8{D9423186-40FD-423E-A2DA-6D5E5C9C20F9}7fcd786fa28967227b674f7e87217d8cSierra Chart
	*/
	h := `02 00 00 00 45 02 00 c9 c1 fa 40 00 80 06 00 00
        7f 00 00 01 7f 00 00 01 c4 ba 2b 5a 54 e8 c1 2d
        89 c0 5f 7e 50 18 9f ff 20 80 00 00 a1 00 01 00
        38 00 00 00 08 00 00 00 38 00 07 00 3f 00 0d 00
        4c 00 27 00 00 00 00 00 61 3f b3 68 14 00 00 00
        00 00 00 00 00 00 00 00 73 00 21 00 94 00 0d 00
        00 00 00 00 6b 61 6d 61 69 75 00 72 36 48 52 67
        54 36 71 24 68 6e 38 00 7b 44 39 34 32 33 31 38
        36 2d 34 30 46 44 2d 34 32 33 45 2d 41 32 44 41
        2d 36 44 35 45 35 43 39 43 32 30 46 39 7d 00 37
        66 63 64 37 38 36 66 61 32 38 39 36 37 32 32 37
        62 36 37 34 66 37 65 38 37 32 31 37 64 38 63 00
        53 69 65 72 72 61 20 43 68 61 72 74 00`

	h = `4d 00 20 03`

	h = strings.ReplaceAll(h, "\n", "")
	h = strings.ReplaceAll(h, " ", "")

	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatal(err)
	}

	print("[")
	for i, by := range b {
		print(by)
		if i < len(b)-1 {
			print(",")
		}
	}
	print("]\n")

	println(len(b))
}

// fromHexChar converts a hex character into its value and a success flag.
func fromHexChar(c byte) (byte, bool) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, true
	}

	return 0, false
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
