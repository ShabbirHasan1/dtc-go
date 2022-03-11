//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateBidAskCompactBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetBidPrice(r.ReadFloat32())
		case 2:
			m.SetBidQuantity(r.ReadFloat32())
		case 3:
			m.SetAskPrice(r.ReadFloat32())
		case 4:
			m.SetAskQuantity(r.ReadFloat32())
		case 5:
			m.SetDateTime(DateTime4Byte(r.ReadUint32()))
		case 6:
			m.SetSymbolID(r.ReadUint32())
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

func (m MarketDataUpdateBidAskCompactBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 117)
	w.WriteFixed32Float32(1, m.BidPrice())
	w.WriteFixed32Float32(2, m.BidQuantity())
	w.WriteFixed32Float32(3, m.AskPrice())
	w.WriteFixed32Float32(4, m.AskQuantity())
	w.WriteFixed32Uint32(5, uint32(m.DateTime()))
	w.WriteUvarint32(6, m.SymbolID())
	return w.Finish(), nil
}
