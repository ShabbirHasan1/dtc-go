//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetCreditAmount(r.ReadFloat64())
		case 3:
			m.SetDebitAmount(r.ReadFloat64())
		case 4:
			m.SetCurrency(r.ReadString())
		case 5:
			m.SetReason(r.ReadString())
		case 6:
			m.SetTradeAccount(r.ReadString())
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

func (m AccountBalanceAdjustmentFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 607)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, m.CreditAmount())
	w.WriteFixed64Float64(3, m.DebitAmount())
	w.WriteString(4, m.Currency())
	w.WriteString(5, m.Reason())
	w.WriteString(6, m.TradeAccount())
	return w.Finish(), nil
}
