//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRecordResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 3:
			m.SetOpenPrice(r.ReadFloat64())
		case 4:
			m.SetHighPrice(r.ReadFloat64())
		case 5:
			m.SetLowPrice(r.ReadFloat64())
		case 6:
			m.SetLastPrice(r.ReadFloat64())
		case 7:
			m.SetVolume(r.ReadFloat64())
		case 9:
			m.SetBidVolume(r.ReadFloat64())
		case 10:
			m.SetAskVolume(r.ReadFloat64())
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

func (m HistoricalPriceDataRecordResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 803)
	w.WriteUvarint32(8, m.NumTrades())
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Int64(2, int64(m.StartDateTime()))
	w.WriteFixed64Float64(3, m.OpenPrice())
	w.WriteFixed64Float64(4, m.HighPrice())
	w.WriteFixed64Float64(5, m.LowPrice())
	w.WriteFixed64Float64(6, m.LastPrice())
	w.WriteFixed64Float64(7, m.Volume())
	w.WriteFixed64Float64(9, m.BidVolume())
	w.WriteFixed64Float64(10, m.AskVolume())
	w.WriteUvarint8(11, m.IsFinalRecord())
	return w.Finish(), nil
}
