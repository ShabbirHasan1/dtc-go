//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataTickRecordResponse_IntPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 3:
			m.SetPrice(r.ReadInt32())
		case 4:
			m.SetVolume(r.ReadInt32())
		case 5:
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.ReadUint16()))
		case 6:
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

func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 806)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, float64(m.DateTime()))
	w.WriteVarint32(3, m.Price())
	w.WriteVarint32(4, m.Volume())
	w.WriteUvarint16(5, uint16(m.AtBidOrAsk()))
	w.WriteUvarint8(6, m.IsFinalRecord())
	return w.Finish(), nil
}
