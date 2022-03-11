//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderIntFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 2:
			m.SetExchange(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetClientOrderID(r.ReadString())
		case 5:
			m.SetOrderType(OrderTypeEnum(r.ReadInt32()))
		case 6:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 7:
			m.SetPrice1(r.ReadInt64())
		case 8:
			m.SetPrice2(r.ReadInt64())
		case 9:
			m.SetDivisor(r.ReadFloat32())
		case 10:
			m.SetQuantity(r.ReadInt64())
		case 11:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 12:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 13:
			m.SetIsAutomatedOrder(r.ReadUint8())
		case 14:
			m.SetIsParentOrder(r.ReadUint8())
		case 15:
			m.SetFreeFormText(r.ReadString())
		case 16:
			m.SetOpenOrClose(OpenCloseTradeEnum(r.ReadInt32()))
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

func (m SubmitNewSingleOrderIntFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 206)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.OrderType()))
	w.WriteVarint32(6, int32(m.BuySell()))
	w.WriteVarint64(7, m.Price1())
	w.WriteVarint64(8, m.Price2())
	w.WriteFixed32Float32(9, m.Divisor())
	w.WriteVarint64(10, m.Quantity())
	w.WriteVarint32(11, int32(m.TimeInForce()))
	w.WriteFixed64Int64(12, int64(m.GoodTillDateTime()))
	w.WriteUvarint8(13, m.IsAutomatedOrder())
	w.WriteUvarint8(14, m.IsParentOrder())
	w.WriteString(15, m.FreeFormText())
	w.WriteVarint32(16, int32(m.OpenOrClose()))
	return w.Finish(), nil
}
