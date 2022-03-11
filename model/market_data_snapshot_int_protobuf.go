//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataSnapshot_IntBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSessionSettlementPrice(r.ReadInt32())
		case 3:
			m.SetSessionOpenPrice(r.ReadInt32())
		case 4:
			m.SetSessionHighPrice(r.ReadInt32())
		case 5:
			m.SetSessionLowPrice(r.ReadInt32())
		case 6:
			m.SetSessionVolume(r.ReadInt32())
		case 7:
			m.SetSessionNumTrades(r.ReadUint32())
		case 8:
			m.SetOpenInterest(r.ReadUint32())
		case 9:
			m.SetBidPrice(r.ReadInt32())
		case 10:
			m.SetAskPrice(r.ReadInt32())
		case 11:
			m.SetAskQuantity(r.ReadInt32())
		case 12:
			m.SetBidQuantity(r.ReadInt32())
		case 13:
			m.SetLastTradePrice(r.ReadInt32())
		case 14:
			m.SetLastTradeVolume(r.ReadInt32())
		case 15:
			m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 16:
			m.SetBidAskDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 17:
			m.SetSessionSettlementDateTime(DateTime4Byte(r.ReadUint32()))
		case 18:
			m.SetTradingSessionDate(DateTime4Byte(r.ReadUint32()))
		case 19:
			m.SetTradingStatus(TradingStatusEnum(r.ReadInt8()))
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

func (m MarketDataSnapshot_IntBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 125)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint32(2, m.SessionSettlementPrice())
	w.WriteVarint32(3, m.SessionOpenPrice())
	w.WriteVarint32(4, m.SessionHighPrice())
	w.WriteVarint32(5, m.SessionLowPrice())
	w.WriteVarint32(6, m.SessionVolume())
	w.WriteUvarint32(7, m.SessionNumTrades())
	w.WriteUvarint32(8, m.OpenInterest())
	w.WriteVarint32(9, m.BidPrice())
	w.WriteVarint32(10, m.AskPrice())
	w.WriteVarint32(11, m.AskQuantity())
	w.WriteVarint32(12, m.BidQuantity())
	w.WriteVarint32(13, m.LastTradePrice())
	w.WriteVarint32(14, m.LastTradeVolume())
	w.WriteFixed64Float64(15, float64(m.LastTradeDateTime()))
	w.WriteFixed64Float64(16, float64(m.BidAskDateTime()))
	w.WriteUvarint32(17, uint32(m.SessionSettlementDateTime()))
	w.WriteUvarint32(18, uint32(m.TradingSessionDate()))
	w.WriteVarint8(19, int8(m.TradingStatus()))
	return w.Finish(), nil
}
