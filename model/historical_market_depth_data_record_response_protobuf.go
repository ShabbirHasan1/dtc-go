//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRecordResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetCommand(r.ReadUint8())
		case 4:
			m.SetFlags(r.ReadUint8())
		case 5:
			m.SetNumOrders(r.ReadUint16())
		case 6:
			m.SetPrice(r.ReadFloat32())
		case 7:
			m.SetQuantity(r.ReadUint32())
		case 8:
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

func (m HistoricalMarketDepthDataRecordResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 903)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint64(2, int64(m.StartDateTime()))
	w.WriteUvarint8(3, m.Command())
	w.WriteUvarint8(4, m.Flags())
	w.WriteUvarint16(5, m.NumOrders())
	w.WriteFixed32Float32(6, m.Price())
	w.WriteUvarint32(7, m.Quantity())
	w.WriteUvarint8(8, m.IsFinalRecord())
	return w.Finish(), nil
}
