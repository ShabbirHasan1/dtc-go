//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SymbolsForExchangeRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetExchange(r.ReadString())
		case 3:
			m.SetSecurityType(SecurityTypeEnum(r.ReadInt32()))
		case 4:
			m.SetRequestAction(RequestActionEnum(r.ReadInt32()))
		case 5:
			m.SetSymbol(r.ReadString())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolsForExchangeRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 502)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteVarint32(3, int32(m.SecurityType()))
	w.WriteVarint32(4, int32(m.RequestAction()))
	w.WriteString(5, m.Symbol())
	return w.Finish(), nil
}
