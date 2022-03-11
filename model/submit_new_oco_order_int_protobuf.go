//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewOCOOrderIntBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetClientOrderID_1(r.ReadString())
		case 4:
			m.SetOrderType_1(OrderTypeEnum(r.ReadInt32()))
		case 5:
			m.SetBuySell_1(BuySellEnum(r.ReadInt32()))
		case 6:
			m.SetPrice1_1(r.ReadInt64())
		case 7:
			m.SetPrice2_1(r.ReadInt64())
		case 8:
			m.SetQuantity_1(r.ReadInt64())
		case 9:
			m.SetClientOrderID_2(r.ReadString())
		case 10:
			m.SetOrderType_2(OrderTypeEnum(r.ReadInt32()))
		case 11:
			m.SetBuySell_2(BuySellEnum(r.ReadInt32()))
		case 12:
			m.SetPrice1_2(r.ReadInt64())
		case 13:
			m.SetPrice2_2(r.ReadInt64())
		case 14:
			m.SetQuantity_2(r.ReadInt64())
		case 15:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 16:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 17:
			m.SetTradeAccount(r.ReadString())
		case 18:
			m.SetIsAutomatedOrder(r.ReadUint8())
		case 19:
			m.SetParentTriggerClientOrderID(r.ReadString())
		case 20:
			m.SetFreeFormText(r.ReadString())
		case 21:
			m.SetDivisor(r.ReadFloat32())
		case 22:
			m.SetOpenOrClose(OpenCloseTradeEnum(r.ReadInt32()))
		case 23:
			m.SetPartialFillHandling(PartialFillHandlingEnum(r.ReadInt8()))
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

func (m SubmitNewOCOOrderIntBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 207)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.ClientOrderID_1())
	w.WriteVarint32(4, int32(m.OrderType_1()))
	w.WriteVarint32(5, int32(m.BuySell_1()))
	w.WriteVarint64(6, m.Price1_1())
	w.WriteVarint64(7, m.Price2_1())
	w.WriteVarint64(8, m.Quantity_1())
	w.WriteString(9, m.ClientOrderID_2())
	w.WriteVarint32(10, int32(m.OrderType_2()))
	w.WriteVarint32(11, int32(m.BuySell_2()))
	w.WriteVarint64(12, m.Price1_2())
	w.WriteVarint64(13, m.Price2_2())
	w.WriteVarint64(14, m.Quantity_2())
	w.WriteVarint32(15, int32(m.TimeInForce()))
	w.WriteFixed64Int64(16, int64(m.GoodTillDateTime()))
	w.WriteString(17, m.TradeAccount())
	w.WriteUvarint8(18, m.IsAutomatedOrder())
	w.WriteString(19, m.ParentTriggerClientOrderID())
	w.WriteString(20, m.FreeFormText())
	w.WriteFixed32Float32(21, m.Divisor())
	w.WriteVarint32(22, int32(m.OpenOrClose()))
	w.WriteVarint8(23, int8(m.PartialFillHandling()))
	return w.Finish(), nil
}
