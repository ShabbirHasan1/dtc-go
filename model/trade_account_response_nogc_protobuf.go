//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetTotalNumberMessages(r.ReadInt32())
		case 2:
			m.SetMessageNumber(r.ReadInt32())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetRequestID(r.ReadInt32())
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

func (m TradeAccountResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 401)
	w.WriteVarint32(1, m.TotalNumberMessages())
	w.WriteVarint32(2, m.MessageNumber())
	w.WriteString(3, m.TradeAccount())
	w.WriteVarint32(4, m.RequestID())
	return w.Finish(), nil
}
