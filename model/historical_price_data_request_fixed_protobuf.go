//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRequestFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 3:
			m.SetExchange(r.ReadString())
		case 4:
			m.SetRecordInterval(HistoricalDataIntervalEnum(r.ReadInt32()))
		case 5:
			m.SetStartDateTime(DateTime(r.ReadInt64()))
		case 6:
			m.SetEndDateTime(DateTime(r.ReadInt64()))
		case 7:
			m.SetMaxDaysToReturn(r.ReadUint32())
		case 8:
			m.SetUseZLibCompression(r.ReadUint8())
		case 9:
			m.SetRequestDividendAdjustedStockData(r.ReadUint8())
		case 10:
			m.SetInteger_1(r.ReadUint16())
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

func (m HistoricalPriceDataRequestFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 800)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint32(4, int32(m.RecordInterval()))
	w.WriteFixed64Int64(5, int64(m.StartDateTime()))
	w.WriteFixed64Int64(6, int64(m.EndDateTime()))
	w.WriteUvarint32(7, m.MaxDaysToReturn())
	w.WriteUvarint8(8, m.UseZLibCompression())
	w.WriteUvarint8(9, m.RequestDividendAdjustedStockData())
	w.WriteUvarint16(10, m.Integer_1())
	return w.Finish(), nil
}
