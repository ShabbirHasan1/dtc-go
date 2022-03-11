//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceUpdateBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetCashBalance(r.ReadFloat64())
		case 3:
			m.SetBalanceAvailableForNewPositions(r.ReadFloat64())
		case 4:
			m.SetAccountCurrency(r.ReadString())
		case 5:
			m.SetTradeAccount(r.ReadString())
		case 6:
			m.SetSecuritiesValue(r.ReadFloat64())
		case 7:
			m.SetMarginRequirement(r.ReadFloat64())
		case 8:
			m.SetTotalNumberMessages(r.ReadInt32())
		case 9:
			m.SetMessageNumber(r.ReadInt32())
		case 10:
			m.SetNoAccountBalances(r.ReadUint8())
		case 11:
			m.SetUnsolicited(r.ReadUint8())
		case 12:
			m.SetOpenPositionsProfitLoss(r.ReadFloat64())
		case 13:
			m.SetDailyProfitLoss(r.ReadFloat64())
		case 14:
			m.SetInfoText(r.ReadString())
		case 15:
			m.SetTransactionIdentifier(r.ReadUint64())
		case 16:
			m.SetDailyNetLossLimit(r.ReadFloat64())
		case 17:
			m.SetTrailingAccountValueToLimitPositions(r.ReadFloat64())
		case 18:
			m.SetDailyNetLossLimitReached(r.ReadUint8())
		case 19:
			m.SetIsUnderRequiredMargin(r.ReadUint8())
		case 20:
			m.SetClosePositionsAtEndOfDay(r.ReadUint8())
		case 21:
			m.SetTradingIsDisabled(r.ReadUint8())
		case 22:
			m.SetDescription(r.ReadString())
		case 23:
			m.SetIsUnderRequiredAccountValue(r.ReadUint8())
		case 24:
			m.SetTransactionDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 25:
			m.SetMarginRequirementFull(r.ReadFloat64())
		case 26:
			m.SetMarginRequirementFullPositionsOnly(r.ReadFloat64())
		case 27:
			m.SetPeakMarginRequirement(r.ReadFloat64())
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

func (m AccountBalanceUpdateBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 600)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, m.CashBalance())
	w.WriteFixed64Float64(3, m.BalanceAvailableForNewPositions())
	w.WriteString(4, m.AccountCurrency())
	w.WriteString(5, m.TradeAccount())
	w.WriteFixed64Float64(6, m.SecuritiesValue())
	w.WriteFixed64Float64(7, m.MarginRequirement())
	w.WriteVarint32(8, m.TotalNumberMessages())
	w.WriteVarint32(9, m.MessageNumber())
	w.WriteUvarint8(10, m.NoAccountBalances())
	w.WriteUvarint8(11, m.Unsolicited())
	w.WriteFixed64Float64(12, m.OpenPositionsProfitLoss())
	w.WriteFixed64Float64(13, m.DailyProfitLoss())
	w.WriteString(14, m.InfoText())
	w.WriteUvarint64(15, m.TransactionIdentifier())
	w.WriteFixed64Float64(16, m.DailyNetLossLimit())
	w.WriteFixed64Float64(17, m.TrailingAccountValueToLimitPositions())
	w.WriteUvarint8(18, m.DailyNetLossLimitReached())
	w.WriteUvarint8(19, m.IsUnderRequiredMargin())
	w.WriteUvarint8(20, m.ClosePositionsAtEndOfDay())
	w.WriteUvarint8(21, m.TradingIsDisabled())
	w.WriteString(22, m.Description())
	w.WriteUvarint8(23, m.IsUnderRequiredAccountValue())
	w.WriteFixed64Int64(24, int64(m.TransactionDateTime()))
	w.WriteFixed64Float64(25, m.MarginRequirementFull())
	w.WriteFixed64Float64(26, m.MarginRequirementFullPositionsOnly())
	w.WriteFixed64Float64(27, m.PeakMarginRequirement())
	return w.Finish(), nil
}
