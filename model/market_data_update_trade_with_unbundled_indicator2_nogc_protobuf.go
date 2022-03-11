//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat32())
		case 3:
			m.SetVolume(r.ReadUint32())
		case 4:
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 5:
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.ReadUint8()))
		case 6:
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.ReadInt8()))
		case 7:
			m.SetTradeCondition(TradeConditionEnum(r.ReadInt8()))
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

func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 146)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.Price())
	w.WriteUvarint32(3, m.Volume())
	w.WriteFixed64Int64(4, int64(m.DateTime()))
	w.WriteUvarint8(5, uint8(m.AtBidOrAsk()))
	w.WriteVarint8(6, int8(m.UnbundledTradeIndicator()))
	w.WriteVarint8(7, int8(m.TradeCondition()))
	return w.Finish(), nil
}
