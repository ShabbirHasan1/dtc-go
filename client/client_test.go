package client

import (
	v8 "github.com/moontrade/dtc-go/v8/sc"
	"testing"
)

func TestClient(t *testing.T) {
	config := NewConfig().
		SetLogon(v8.NewLogonRequest().
			SetProtocolVersion(8).
			SetGeneralTextData("{D9423186-40FD-423E-A2DA-6D5E5C9C20F9}").
			SetHardwareIdentifier("1441858590").
			SetClientName("Sierra Chart").
			SetTradeAccount("47029"))

	c, err := Dial("127.0.0.1:11099", config)
	//c, err := DialTLS("ds22.sierracharts.com:11091", config)
	if err != nil {
		t.Fatal(err)
	}

	c.Wait()
}
