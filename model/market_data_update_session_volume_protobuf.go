//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateSessionVolumeBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetVolume(r.ReadFloat64())
		case 3:
			m.SetTradingSessionDate(DateTime4Byte(r.ReadUint32()))
		case 4:
			m.SetIsFinalSessionVolume(r.ReadUint8())
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

func (m MarketDataUpdateSessionVolumeBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 113)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed64Float64(2, m.Volume())
	w.WriteUvarint32(3, uint32(m.TradingSessionDate()))
	w.WriteUvarint8(4, m.IsFinalSessionVolume())
	return w.Finish(), nil
}
