//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice1(r.ReadFloat64())
		case 4:
			m.SetPrice2(r.ReadFloat64())
		case 5:
			m.SetQuantity(r.ReadFloat64())
		case 6:
			m.SetPrice1IsSet(r.ReadUint8())
		case 7:
			m.SetPrice2IsSet(r.ReadUint8())
		case 9:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 10:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 11:
			m.SetUpdatePrice1OffsetToParent(r.ReadUint8())
		case 12:
			m.SetTradeAccount(r.ReadString())
		case 13:
			m.SetPrice1AsString(r.ReadString())
		case 14:
			m.SetPrice2AsString(r.ReadString())
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

func (m CancelReplaceOrderFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 204)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteFixed64Float64(3, m.Price1())
	w.WriteFixed64Float64(4, m.Price2())
	w.WriteFixed64Float64(5, m.Quantity())
	w.WriteUvarint8(6, m.Price1IsSet())
	w.WriteUvarint8(7, m.Price2IsSet())
	w.WriteVarint32(9, int32(m.TimeInForce()))
	w.WriteVarint64(10, int64(m.GoodTillDateTime()))
	w.WriteUvarint8(11, m.UpdatePrice1OffsetToParent())
	w.WriteString(12, m.TradeAccount())
	w.WriteString(13, m.Price1AsString())
	w.WriteString(14, m.Price2AsString())
	return w.Finish(), nil
}
