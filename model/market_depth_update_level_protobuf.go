//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevelBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat64())
		case 4:
			m.SetQuantity(r.ReadFloat64())
		case 5:
			m.SetUpdateType(MarketDepthUpdateTypeEnum(r.ReadUint8()))
		case 6:
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 7:
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

func (m MarketDepthUpdateLevelBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 106)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint16(2, uint16(m.Side()))
	w.WriteFixed64Float64(3, m.Price())
	w.WriteFixed64Float64(4, m.Quantity())
	w.WriteUvarint8(5, uint8(m.UpdateType()))
	w.WriteFixed64Float64(6, float64(m.DateTime()))
	w.WriteUvarint32(7, m.NumOrders())
	return w.Finish(), nil
}
