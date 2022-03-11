//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillsRequestFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 3:
			m.SetNumberOfDays(r.ReadInt32())
		case 4:
			m.SetTradeAccount(r.ReadString())
		case 5:
			m.SetStartDateTime(DateTime(r.ReadInt64()))
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

func (m HistoricalOrderFillsRequestFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 303)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.ServerOrderID())
	w.WriteVarint32(3, m.NumberOfDays())
	w.WriteString(4, m.TradeAccount())
	w.WriteFixed64Int64(5, int64(m.StartDateTime()))
	return w.Finish(), nil
}
