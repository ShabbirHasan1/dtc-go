//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillResponseFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetTotalNumberMessages(r.ReadInt32())
		case 3:
			m.SetMessageNumber(r.ReadInt32())
		case 4:
			m.SetSymbol(r.ReadString())
		case 5:
			m.SetExchange(r.ReadString())
		case 6:
			m.SetServerOrderID(r.ReadString())
		case 7:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 8:
			m.SetPrice(r.ReadFloat64())
		case 9:
			m.SetDateTime(DateTime(r.ReadInt64()))
		case 10:
			m.SetQuantity(r.ReadFloat64())
		case 11:
			m.SetUniqueExecutionID(r.ReadString())
		case 12:
			m.SetTradeAccount(r.ReadString())
		case 13:
			m.SetOpenClose(OpenCloseTradeEnum(r.ReadInt32()))
		case 14:
			m.SetNoOrderFills(r.ReadUint8())
		case 15:
			m.SetInfoText(r.ReadString())
		case 16:
			m.SetHighPriceDuringPosition(r.ReadFloat64())
		case 17:
			m.SetLowPriceDuringPosition(r.ReadFloat64())
		case 18:
			m.SetPositionQuantity(r.ReadFloat64())
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

func (m HistoricalOrderFillResponseFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 304)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.TotalNumberMessages())
	w.WriteVarint32(3, m.MessageNumber())
	w.WriteString(4, m.Symbol())
	w.WriteString(5, m.Exchange())
	w.WriteString(6, m.ServerOrderID())
	w.WriteVarint32(7, int32(m.BuySell()))
	w.WriteFixed64Float64(8, m.Price())
	w.WriteFixed64Int64(9, int64(m.DateTime()))
	w.WriteFixed64Float64(10, m.Quantity())
	w.WriteString(11, m.UniqueExecutionID())
	w.WriteString(12, m.TradeAccount())
	w.WriteVarint32(13, int32(m.OpenClose()))
	w.WriteUvarint8(14, m.NoOrderFills())
	w.WriteString(15, m.InfoText())
	w.WriteFixed64Float64(16, m.HighPriceDuringPosition())
	w.WriteFixed64Float64(17, m.LowPriceDuringPosition())
	w.WriteFixed64Float64(18, m.PositionQuantity())
	return w.Finish(), nil
}
