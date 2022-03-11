//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderIntFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 2:
			m.SetClientOrderID(r.ReadString())
		case 3:
			m.SetPrice1(r.ReadInt64())
		case 4:
			m.SetPrice2(r.ReadInt64())
		case 5:
			m.SetDivisor(r.ReadFloat32())
		case 6:
			m.SetQuantity(r.ReadInt64())
		case 7:
			m.SetPrice1IsSet(r.ReadUint8())
		case 8:
			m.SetPrice2IsSet(r.ReadUint8())
		case 10:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 11:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 12:
			m.SetUpdatePrice1OffsetToParent(r.ReadUint8())
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

func (m CancelReplaceOrderIntFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 205)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteVarint64(3, m.Price1())
	w.WriteVarint64(4, m.Price2())
	w.WriteFixed32Float32(5, m.Divisor())
	w.WriteVarint64(6, m.Quantity())
	w.WriteUvarint8(7, m.Price1IsSet())
	w.WriteUvarint8(8, m.Price2IsSet())
	w.WriteVarint32(10, int32(m.TimeInForce()))
	w.WriteVarint64(11, int64(m.GoodTillDateTime()))
	w.WriteUvarint8(12, m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}
