//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRecordResponse_IntPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetStartDateTime(DateTime(r.ReadInt64()))
		case 3:
			m.SetOpenPrice(r.ReadInt32())
		case 4:
			m.SetHighPrice(r.ReadInt32())
		case 5:
			m.SetLowPrice(r.ReadInt32())
		case 6:
			m.SetLastPrice(r.ReadInt32())
		case 7:
			m.SetVolume(r.ReadInt32())
		case 9:
			m.SetBidVolume(r.ReadInt32())
		case 10:
			m.SetAskVolume(r.ReadInt32())
		case 11:
			m.SetIsFinalRecord(r.ReadUint8())
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

func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 805)
	w.WriteUvarint32(8, m.NumTrades())
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Int64(2, int64(m.StartDateTime()))
	w.WriteVarint32(3, m.OpenPrice())
	w.WriteVarint32(4, m.HighPrice())
	w.WriteVarint32(5, m.LowPrice())
	w.WriteVarint32(6, m.LastPrice())
	w.WriteVarint32(7, m.Volume())
	w.WriteVarint32(9, m.BidVolume())
	w.WriteVarint32(10, m.AskVolume())
	w.WriteUvarint8(11, m.IsFinalRecord())
	return w.Finish(), nil
}
