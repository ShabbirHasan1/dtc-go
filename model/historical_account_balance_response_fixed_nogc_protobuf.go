//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalanceResponseFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 3:
			m.SetCashBalance(r.ReadFloat64())
		case 4:
			m.SetAccountCurrency(r.ReadString())
		case 5:
			m.SetTradeAccount(r.ReadString())
		case 6:
			m.SetIsFinalResponse(r.ReadUint8())
		case 7:
			m.SetNoAccountBalances(r.ReadUint8())
		case 8:
			m.SetInfoText(r.ReadString())
		case 9:
			m.SetTransactionId(r.ReadString())
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

func (m HistoricalAccountBalanceResponseFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 606)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, float64(m.DateTime()))
	w.WriteFixed64Float64(3, m.CashBalance())
	w.WriteString(4, m.AccountCurrency())
	w.WriteString(5, m.TradeAccount())
	w.WriteUvarint8(6, m.IsFinalResponse())
	w.WriteUvarint8(7, m.NoAccountBalances())
	w.WriteString(8, m.InfoText())
	w.WriteString(9, m.TransactionId())
	return w.Finish(), nil
}
