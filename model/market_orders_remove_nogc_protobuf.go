//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersRemovePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat64())
		case 3:
			m.SetQuantity(r.ReadUint32())
		case 4:
			m.SetOrderID(r.ReadUint64())
		case 5:
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 6:
			m.SetSide(BuySellEnum(r.ReadInt32()))
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

func (m MarketOrdersRemovePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 154)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed64Float64(2, m.Price())
	w.WriteUvarint32(3, m.Quantity())
	w.WriteUvarint64(4, m.OrderID())
	w.WriteVarint64(5, int64(m.DateTime()))
	w.WriteVarint32(6, int32(m.Side()))
	return w.Finish(), nil
}
