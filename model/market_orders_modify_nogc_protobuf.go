//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersModifyPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPriorPrice(r.ReadFloat64())
		case 6:
			m.SetPriorQuantity(r.ReadUint32())
		case 7:
			m.SetPriorOrderID(r.ReadUint64())
		case 8:
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
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

func (m MarketOrdersModifyPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 153)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed64Float64(2, m.Price())
	w.WriteUvarint32(3, m.Quantity())
	w.WriteUvarint64(4, m.OrderID())
	w.WriteFixed64Float64(5, m.PriorPrice())
	w.WriteUvarint32(6, m.PriorQuantity())
	w.WriteUvarint64(7, m.PriorOrderID())
	w.WriteVarint64(8, int64(m.DateTime()))
	return w.Finish(), nil
}
