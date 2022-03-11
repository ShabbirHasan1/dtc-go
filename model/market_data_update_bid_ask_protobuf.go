//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidPrice(r.ReadFloat64())
		case 3:
			m.SetBidQuantity(r.ReadFloat32())
		case 4:
			m.SetAskPrice(r.ReadFloat64())
		case 5:
			m.SetAskQuantity(r.ReadFloat32())
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

func (m MarketDataUpdateBidAskBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 108)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed64Float64(2, m.BidPrice())
	w.WriteFixed32Float32(3, m.BidQuantity())
	w.WriteFixed64Float64(4, m.AskPrice())
	w.WriteFixed32Float32(5, m.AskQuantity())
	w.WriteFixed32Uint32(6, uint32(m.DateTime()))
	return w.Finish(), nil
}
