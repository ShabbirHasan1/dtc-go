//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthSnapshotLevel_IntPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbolID(r.ReadUint32())
		case 2:
			m.SetSide(AtBidOrAskEnum(r.ReadUint16()))
		case 3:
			m.SetPrice(r.ReadInt32())
		case 4:
			m.SetQuantity(r.ReadInt32())
		case 5:
			m.SetLevel(r.ReadUint16())
		case 6:
			m.SetIsFirstMessageInBatch(r.ReadUint8())
		case 7:
			m.SetIsLastMessageInBatch(r.ReadUint8())
		case 8:
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 9:
			m.SetNumOrders(r.ReadUint32())
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

func (m MarketDepthSnapshotLevel_IntPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 132)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint16(2, uint16(m.Side()))
	w.WriteVarint32(3, m.Price())
	w.WriteVarint32(4, m.Quantity())
	w.WriteUvarint16(5, m.Level())
	w.WriteUvarint8(6, m.IsFirstMessageInBatch())
	w.WriteUvarint8(7, m.IsLastMessageInBatch())
	w.WriteFixed64Float64(8, float64(m.DateTime()))
	w.WriteUvarint32(9, m.NumOrders())
	return w.Finish(), nil
}
