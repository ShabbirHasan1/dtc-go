//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDateTime(DateTimeWithMillisecondsInt(r.ReadInt64()))
		case 3:
			m.SetPrice(r.ReadFloat32())
		case 4:
			m.SetQuantity(r.ReadFloat32())
		case 5:
			m.SetSide(AtBidOrAskEnum8(r.ReadUint8()))
		case 6:
			m.SetUpdateType(MarketDepthUpdateTypeEnum(r.ReadUint8()))
		case 7:
			m.SetNumOrders(r.ReadUint16())
		case 8:
			m.SetFinalUpdateInBatch(FinalUpdateInBatchEnum(r.ReadUint8()))
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

func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 140)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint64(2, int64(m.DateTime()))
	w.WriteFixed32Float32(3, m.Price())
	w.WriteFixed32Float32(4, m.Quantity())
	w.WriteUvarint8(5, uint8(m.Side()))
	w.WriteUvarint8(6, uint8(m.UpdateType()))
	w.WriteUvarint16(7, m.NumOrders())
	w.WriteUvarint8(8, uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}
