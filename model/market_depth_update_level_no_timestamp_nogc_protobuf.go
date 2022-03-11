//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevelNoTimestampPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat32())
		case 3:
			m.SetQuantity(r.ReadFloat32())
		case 4:
			m.SetNumOrders(r.ReadUint16())
		case 5:
			m.SetSide(r.ReadInt8())
		case 6:
			m.SetUpdateType(r.ReadInt8())
		case 7:
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

func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 141)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.Price())
	w.WriteFixed32Float32(3, m.Quantity())
	w.WriteUvarint16(4, m.NumOrders())
	w.WriteVarint8(5, m.Side())
	w.WriteVarint8(6, m.UpdateType())
	w.WriteUvarint8(7, uint8(m.FinalUpdateInBatch()))
	return w.Finish(), nil
}
