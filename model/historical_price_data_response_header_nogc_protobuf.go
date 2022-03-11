//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataResponseHeaderPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetRecordInterval(HistoricalDataIntervalEnum(r.ReadInt32()))
		case 3:
			m.SetUseZLibCompression(r.ReadUint8())
		case 4:
			m.SetNoRecordsToReturn(r.ReadUint8())
		case 5:
			m.SetIntToFloatPriceDivisor(r.ReadFloat32())
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

func (m HistoricalPriceDataResponseHeaderPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 801)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, int32(m.RecordInterval()))
	w.WriteUvarint8(3, m.UseZLibCompression())
	w.WriteUvarint8(4, m.NoRecordsToReturn())
	w.WriteFixed32Float32(5, m.IntToFloatPriceDivisor())
	return w.Finish(), nil
}
