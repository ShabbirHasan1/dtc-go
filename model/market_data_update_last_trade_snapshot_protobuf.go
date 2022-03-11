//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateLastTradeSnapshotBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetLastTradePrice(r.ReadFloat64())
		case 3:
			m.SetLastTradeVolume(r.ReadFloat64())
		case 4:
			m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
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

func (m MarketDataUpdateLastTradeSnapshotBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 134)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed64Float64(2, m.LastTradePrice())
	w.WriteFixed64Float64(3, m.LastTradeVolume())
	w.WriteFixed64Float64(4, float64(m.LastTradeDateTime()))
	return w.Finish(), nil
}
