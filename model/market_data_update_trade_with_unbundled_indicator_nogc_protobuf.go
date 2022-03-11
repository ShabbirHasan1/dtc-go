//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.ReadUint8()))
		case 3:
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.ReadInt8()))
		case 4:
			m.SetPrice(r.ReadFloat64())
		case 5:
			m.SetVolume(r.ReadUint32())
		case 6:
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 7:
			m.SetTradeCondition(r.ReadUint8())
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

func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 137)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint8(2, uint8(m.AtBidOrAsk()))
	w.WriteVarint8(3, int8(m.UnbundledTradeIndicator()))
	w.WriteFixed64Float64(4, m.Price())
	w.WriteUvarint32(5, m.Volume())
	w.WriteFixed64Float64(6, float64(m.DateTime()))
	w.WriteUvarint8(7, m.TradeCondition())
	return w.Finish(), nil
}
