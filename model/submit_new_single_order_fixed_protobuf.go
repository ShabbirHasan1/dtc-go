//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice1(r.ReadFloat64())
		case 8:
			m.SetPrice2(r.ReadFloat64())
		case 9:
			m.SetQuantity(r.ReadFloat64())
		case 10:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 11:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 12:
			m.SetIsAutomatedOrder(r.ReadUint8())
		case 13:
			m.SetIsParentOrder(r.ReadUint8())
		case 14:
			m.SetFreeFormText(r.ReadString())
		case 15:
			m.SetOpenOrClose(OpenCloseTradeEnum(r.ReadInt32()))
		case 16:
			m.SetMaxShowQuantity(r.ReadFloat64())
		case 17:
			m.SetPrice1AsString(r.ReadString())
		case 18:
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

func (m SubmitNewSingleOrderFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 208)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.OrderType()))
	w.WriteVarint32(6, int32(m.BuySell()))
	w.WriteFixed64Float64(7, m.Price1())
	w.WriteFixed64Float64(8, m.Price2())
	w.WriteFixed64Float64(9, m.Quantity())
	w.WriteVarint32(10, int32(m.TimeInForce()))
	w.WriteFixed64Int64(11, int64(m.GoodTillDateTime()))
	w.WriteUvarint8(12, m.IsAutomatedOrder())
	w.WriteUvarint8(13, m.IsParentOrder())
	w.WriteString(14, m.FreeFormText())
	w.WriteVarint32(15, int32(m.OpenOrClose()))
	w.WriteFixed64Float64(16, m.MaxShowQuantity())
	w.WriteString(17, m.Price1AsString())
	w.WriteString(18, m.Price2AsString())
	return w.Finish(), nil
}
