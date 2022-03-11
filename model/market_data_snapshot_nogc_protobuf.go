//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataSnapshotPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSessionSettlementPrice(r.ReadFloat64())
		case 3:
			m.SetSessionOpenPrice(r.ReadFloat64())
		case 4:
			m.SetSessionHighPrice(r.ReadFloat64())
		case 5:
			m.SetSessionLowPrice(r.ReadFloat64())
		case 6:
			m.SetSessionVolume(r.ReadFloat64())
		case 7:
			m.SetSessionNumTrades(r.ReadUint32())
		case 8:
			m.SetOpenInterest(r.ReadUint32())
		case 9:
			m.SetBidPrice(r.ReadFloat64())
		case 10:
			m.SetAskPrice(r.ReadFloat64())
		case 11:
			m.SetAskQuantity(r.ReadFloat64())
		case 12:
			m.SetBidQuantity(r.ReadFloat64())
		case 13:
			m.SetLastTradePrice(r.ReadFloat64())
		case 14:
			m.SetLastTradeVolume(r.ReadFloat64())
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
		case 20:
			m.SetMarketDepthUpdateDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
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

func (m MarketDataSnapshotPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 104)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed64Float64(2, m.SessionSettlementPrice())
	w.WriteFixed64Float64(3, m.SessionOpenPrice())
	w.WriteFixed64Float64(4, m.SessionHighPrice())
	w.WriteFixed64Float64(5, m.SessionLowPrice())
	w.WriteFixed64Float64(6, m.SessionVolume())
	w.WriteUvarint32(7, m.SessionNumTrades())
	w.WriteUvarint32(8, m.OpenInterest())
	w.WriteFixed64Float64(9, m.BidPrice())
	w.WriteFixed64Float64(10, m.AskPrice())
	w.WriteFixed64Float64(11, m.AskQuantity())
	w.WriteFixed64Float64(12, m.BidQuantity())
	w.WriteFixed64Float64(13, m.LastTradePrice())
	w.WriteFixed64Float64(14, m.LastTradeVolume())
	w.WriteFixed64Float64(15, float64(m.LastTradeDateTime()))
	w.WriteFixed64Float64(16, float64(m.BidAskDateTime()))
	w.WriteUvarint32(17, uint32(m.SessionSettlementDateTime()))
	w.WriteUvarint32(18, uint32(m.TradingSessionDate()))
	w.WriteVarint8(19, int8(m.TradingStatus()))
	w.WriteFixed64Float64(20, float64(m.MarketDepthUpdateDateTime()))
	return w.Finish(), nil
}
