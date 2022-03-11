//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataRequestFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestAction(RequestActionEnum(r.ReadInt32()))
		case 2:
			m.SetSymbolID(r.ReadUint32())
		case 3:
			m.SetSymbol(r.ReadString())
		case 4:
			m.SetExchange(r.ReadString())
		case 5:
			m.SetIntervalForSnapshotUpdatesInMilliseconds(r.ReadUint32())
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

func (m MarketDataRequestFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 101)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteUvarint32(5, m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}
