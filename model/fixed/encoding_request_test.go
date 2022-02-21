package fixed

import (
	"fmt"
	"github.com/moontrade/dtc-go/model/types"
	"testing"
)

func TestEncodingRequestFixed_Encoding(t *testing.T) {
	b := AllocEncodingRequest()
	b.SetEncoding(4)
	b.SetProtocolType("DTC")

	//b.SetProtocolVersion(8)
	r := b.Finish()

	var bb types.EncodingRequestBuilder = b
	bb.Encoding()

	fmt.Println("EncodingRequest {")
	fmt.Println("\t", "Size", "				=", r.Size())
	fmt.Println("\t", "Type", "				=", r.Type())
	fmt.Println("\t", "ProtocolVersion", "	=", r.ProtocolVersion())
	fmt.Println("\t", "Encoding", "			=", r.Encoding())
	fmt.Println("\t", "ProtocolType", "		=", r.ProtocolType())
	fmt.Println("}")
}
