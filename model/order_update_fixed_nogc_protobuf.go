//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *OrderUpdateFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetTotalNumMessages(r.ReadInt32())
		case 3:
			m.SetMessageNumber(r.ReadInt32())
		case 4:
			m.SetSymbol(r.ReadString())
		case 5:
			m.SetExchange(r.ReadString())
		case 6:
			m.SetPreviousServerOrderID(r.ReadString())
		case 7:
			m.SetServerOrderID(r.ReadString())
		case 8:
			m.SetClientOrderID(r.ReadString())
		case 9:
			m.SetExchangeOrderID(r.ReadString())
		case 10:
			m.SetOrderStatus(OrderStatusEnum(r.ReadInt32()))
		case 11:
			m.SetOrderUpdateReason(OrderUpdateReasonEnum(r.ReadInt32()))
		case 12:
			m.SetOrderType(OrderTypeEnum(r.ReadInt32()))
		case 13:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 14:
			m.SetPrice1(r.ReadFloat64())
		case 15:
			m.SetPrice2(r.ReadFloat64())
		case 16:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 17:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 18:
			m.SetOrderQuantity(r.ReadFloat64())
		case 19:
			m.SetFilledQuantity(r.ReadFloat64())
		case 20:
			m.SetRemainingQuantity(r.ReadFloat64())
		case 21:
			m.SetAverageFillPrice(r.ReadFloat64())
		case 22:
			m.SetLastFillPrice(r.ReadFloat64())
		case 23:
			m.SetLastFillDateTime(DateTimeWithMillisecondsInt(r.ReadInt64()))
		case 24:
			m.SetLastFillQuantity(r.ReadFloat64())
		case 25:
			m.SetLastFillExecutionID(r.ReadString())
		case 26:
			m.SetTradeAccount(r.ReadString())
		case 27:
			m.SetInfoText(r.ReadString())
		case 28:
			m.SetNoOrders(r.ReadUint8())
		case 29:
			m.SetParentServerOrderID(r.ReadString())
		case 30:
			m.SetOCOLinkedOrderServerOrderID(r.ReadString())
		case 31:
			m.SetOpenOrClose(OpenCloseTradeEnum(r.ReadInt32()))
		case 32:
			m.SetPreviousClientOrderID(r.ReadString())
		case 33:
			m.SetFreeFormText(r.ReadString())
		case 34:
			m.SetOrderReceivedDateTime(DateTimeWithMillisecondsInt(r.ReadInt64()))
		case 35:
			m.SetLatestTransactionDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
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

func (m OrderUpdateFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 301)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.TotalNumMessages())
	w.WriteVarint32(3, m.MessageNumber())
	w.WriteString(4, m.Symbol())
	w.WriteString(5, m.Exchange())
	w.WriteString(6, m.PreviousServerOrderID())
	w.WriteString(7, m.ServerOrderID())
	w.WriteString(8, m.ClientOrderID())
	w.WriteString(9, m.ExchangeOrderID())
	w.WriteVarint32(10, int32(m.OrderStatus()))
	w.WriteVarint32(11, int32(m.OrderUpdateReason()))
	w.WriteVarint32(12, int32(m.OrderType()))
	w.WriteVarint32(13, int32(m.BuySell()))
	w.WriteFixed64Float64(14, m.Price1())
	w.WriteFixed64Float64(15, m.Price2())
	w.WriteVarint32(16, int32(m.TimeInForce()))
	w.WriteFixed64Int64(17, int64(m.GoodTillDateTime()))
	w.WriteFixed64Float64(18, m.OrderQuantity())
	w.WriteFixed64Float64(19, m.FilledQuantity())
	w.WriteFixed64Float64(20, m.RemainingQuantity())
	w.WriteFixed64Float64(21, m.AverageFillPrice())
	w.WriteFixed64Float64(22, m.LastFillPrice())
	w.WriteFixed64Int64(23, int64(m.LastFillDateTime()))
	w.WriteFixed64Float64(24, m.LastFillQuantity())
	w.WriteString(25, m.LastFillExecutionID())
	w.WriteString(26, m.TradeAccount())
	w.WriteString(27, m.InfoText())
	w.WriteUvarint8(28, m.NoOrders())
	w.WriteString(29, m.ParentServerOrderID())
	w.WriteString(30, m.OCOLinkedOrderServerOrderID())
	w.WriteVarint32(31, int32(m.OpenOrClose()))
	w.WriteString(32, m.PreviousClientOrderID())
	w.WriteString(33, m.FreeFormText())
	w.WriteFixed64Int64(34, int64(m.OrderReceivedDateTime()))
	w.WriteFixed64Float64(35, float64(m.LatestTransactionDateTime()))
	return w.Finish(), nil
}
