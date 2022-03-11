//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeCompactPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat32())
		case 2:
			m.SetVolume(r.ReadFloat32())
		case 3:
			m.SetDateTime(DateTime4Byte(r.ReadUint32()))
		case 4:
			m.SetSymbolID(r.ReadUint32())
		case 5:
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.ReadUint16()))
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

func (m MarketDataUpdateTradeCompactPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 112)
	w.WriteFixed32Float32(1, m.Price())
	w.WriteFixed32Float32(2, m.Volume())
	w.WriteFixed32Uint32(3, uint32(m.DateTime()))
	w.WriteUvarint32(4, m.SymbolID())
	w.WriteUvarint16(5, uint16(m.AtBidOrAsk()))
	return w.Finish(), nil
}
