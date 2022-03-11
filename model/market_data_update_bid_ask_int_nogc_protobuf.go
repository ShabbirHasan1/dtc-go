//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAsk_IntPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidPrice(r.ReadInt32())
		case 3:
			m.SetBidQuantity(r.ReadInt32())
		case 4:
			m.SetAskPrice(r.ReadInt32())
		case 5:
			m.SetAskQuantity(r.ReadInt32())
		case 6:
			m.SetDateTime(DateTime4Byte(r.ReadUint32()))
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

func (m MarketDataUpdateBidAsk_IntPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 127)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteVarint32(2, m.BidPrice())
	w.WriteVarint32(3, m.BidQuantity())
	w.WriteVarint32(4, m.AskPrice())
	w.WriteVarint32(5, m.AskQuantity())
	w.WriteFixed32Uint32(6, uint32(m.DateTime()))
	return w.Finish(), nil
}
