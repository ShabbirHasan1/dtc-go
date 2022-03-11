//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetProtocolVersion(r.ReadInt32())
		case 2:
			m.SetResult(LogonStatusEnum(r.ReadInt32()))
		case 3:
			m.SetResultText(r.ReadString())
		case 4:
			m.SetReconnectAddress(r.ReadString())
		case 5:
			m.SetInteger_1(r.ReadInt32())
		case 6:
			m.SetServerName(r.ReadString())
		case 7:
			m.SetMarketDepthUpdatesBestBidAndAsk(r.ReadUint8())
		case 8:
			m.SetTradingIsSupported(r.ReadUint8())
		case 9:
			m.SetOCOOrdersSupported(r.ReadUint8())
		case 10:
			m.SetOrderCancelReplaceSupported(r.ReadUint8())
		case 11:
			m.SetSymbolExchangeDelimiter(r.ReadString())
		case 12:
			m.SetSecurityDefinitionsSupported(r.ReadUint8())
		case 13:
			m.SetHistoricalPriceDataSupported(r.ReadUint8())
		case 14:
			m.SetResubscribeWhenMarketDataFeedAvailable(r.ReadUint8())
		case 15:
			m.SetMarketDepthIsSupported(r.ReadUint8())
		case 16:
			m.SetOneHistoricalPriceDataRequestPerConnection(r.ReadUint8())
		case 17:
			m.SetBracketOrdersSupported(r.ReadUint8())
		case 18:
			m.SetUseIntegerPriceOrderMessages(r.ReadUint8())
		case 19:
			m.SetUsesMultiplePositionsPerSymbolAndTradeAccount(r.ReadUint8())
		case 20:
			m.SetMarketDataSupported(r.ReadUint8())
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

func (m LogonResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 2)
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteVarint32(2, int32(m.Result()))
	w.WriteString(3, m.ResultText())
	w.WriteString(4, m.ReconnectAddress())
	w.WriteVarint32(5, m.Integer_1())
	w.WriteString(6, m.ServerName())
	w.WriteUvarint8(7, m.MarketDepthUpdatesBestBidAndAsk())
	w.WriteUvarint8(8, m.TradingIsSupported())
	w.WriteUvarint8(9, m.OCOOrdersSupported())
	w.WriteUvarint8(10, m.OrderCancelReplaceSupported())
	w.WriteString(11, m.SymbolExchangeDelimiter())
	w.WriteUvarint8(12, m.SecurityDefinitionsSupported())
	w.WriteUvarint8(13, m.HistoricalPriceDataSupported())
	w.WriteUvarint8(14, m.ResubscribeWhenMarketDataFeedAvailable())
	w.WriteUvarint8(15, m.MarketDepthIsSupported())
	w.WriteUvarint8(16, m.OneHistoricalPriceDataRequestPerConnection())
	w.WriteUvarint8(17, m.BracketOrdersSupported())
	w.WriteUvarint8(18, m.UseIntegerPriceOrderMessages())
	w.WriteUvarint8(19, m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	w.WriteUvarint8(20, m.MarketDataSupported())
	return w.Finish(), nil
}
