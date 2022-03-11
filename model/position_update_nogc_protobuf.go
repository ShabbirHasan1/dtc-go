//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *PositionUpdatePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetQuantity(r.ReadFloat64())
		case 7:
			m.SetAveragePrice(r.ReadFloat64())
		case 8:
			m.SetPositionIdentifier(r.ReadString())
		case 9:
			m.SetTradeAccount(r.ReadString())
		case 10:
			m.SetNoPositions(r.ReadUint8())
		case 11:
			m.SetUnsolicited(r.ReadUint8())
		case 12:
			m.SetMarginRequirement(r.ReadFloat64())
		case 13:
			m.SetEntryDateTime(DateTime4Byte(r.ReadUint32()))
		case 14:
			m.SetOpenProfitLoss(r.ReadFloat64())
		case 15:
			m.SetHighPriceDuringPosition(r.ReadFloat64())
		case 16:
			m.SetLowPriceDuringPosition(r.ReadFloat64())
		case 17:
			m.SetQuantityLimit(r.ReadFloat64())
		case 18:
			m.SetMaxPotentialPostionQuantity(r.ReadFloat64())
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

func (m PositionUpdatePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 306)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.TotalNumberMessages())
	w.WriteVarint32(3, m.MessageNumber())
	w.WriteString(4, m.Symbol())
	w.WriteString(5, m.Exchange())
	w.WriteFixed64Float64(6, m.Quantity())
	w.WriteFixed64Float64(7, m.AveragePrice())
	w.WriteString(8, m.PositionIdentifier())
	w.WriteString(9, m.TradeAccount())
	w.WriteUvarint8(10, m.NoPositions())
	w.WriteUvarint8(11, m.Unsolicited())
	w.WriteFixed64Float64(12, m.MarginRequirement())
	w.WriteUvarint32(13, uint32(m.EntryDateTime()))
	w.WriteFixed64Float64(14, m.OpenProfitLoss())
	w.WriteFixed64Float64(15, m.HighPriceDuringPosition())
	w.WriteFixed64Float64(16, m.LowPriceDuringPosition())
	w.WriteFixed64Float64(17, m.QuantityLimit())
	w.WriteFixed64Float64(18, m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}
