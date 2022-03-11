//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SecurityDefinitionResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSecurityType(SecurityTypeEnum(r.ReadInt32()))
		case 5:
			m.SetDescription(r.ReadString())
		case 6:
			m.SetMinPriceIncrement(r.ReadFloat32())
		case 7:
			m.SetPriceDisplayFormat(PriceDisplayFormatEnum(r.ReadInt32()))
		case 8:
			m.SetCurrencyValuePerIncrement(r.ReadFloat32())
		case 9:
			m.SetIsFinalMessage(r.ReadUint8())
		case 10:
			m.SetFloatToIntPriceMultiplier(r.ReadFloat32())
		case 11:
			m.SetIntToFloatPriceDivisor(r.ReadFloat32())
		case 12:
			m.SetUnderlyingSymbol(r.ReadString())
		case 13:
			m.SetUpdatesBidAskOnly(r.ReadUint8())
		case 14:
			m.SetStrikePrice(r.ReadFloat32())
		case 15:
			m.SetPutOrCall(PutCallEnum(r.ReadUint8()))
		case 16:
			m.SetShortInterest(r.ReadUint32())
		case 17:
			m.SetSecurityExpirationDate(DateTime4Byte(r.ReadUint32()))
		case 18:
			m.SetBuyRolloverInterest(r.ReadFloat32())
		case 19:
			m.SetSellRolloverInterest(r.ReadFloat32())
		case 20:
			m.SetEarningsPerShare(r.ReadFloat32())
		case 21:
			m.SetSharesOutstanding(r.ReadUint32())
		case 22:
			m.SetIntToFloatQuantityDivisor(r.ReadFloat32())
		case 23:
			m.SetHasMarketDepthData(r.ReadUint8())
		case 24:
			m.SetDisplayPriceMultiplier(r.ReadFloat32())
		case 25:
			m.SetExchangeSymbol(r.ReadString())
		case 26:
			m.SetInitialMarginRequirement(r.ReadFloat32())
		case 27:
			m.SetMaintenanceMarginRequirement(r.ReadFloat32())
		case 28:
			m.SetCurrency(r.ReadString())
		case 29:
			m.SetContractSize(r.ReadFloat32())
		case 30:
			m.SetOpenInterest(r.ReadUint32())
		case 31:
			m.SetRolloverDate(DateTime4Byte(r.ReadUint32()))
		case 32:
			m.SetIsDelayed(r.ReadUint8())
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

func (m SecurityDefinitionResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 507)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint32(4, int32(m.SecurityType()))
	w.WriteString(5, m.Description())
	w.WriteFixed32Float32(6, m.MinPriceIncrement())
	w.WriteVarint32(7, int32(m.PriceDisplayFormat()))
	w.WriteFixed32Float32(8, m.CurrencyValuePerIncrement())
	w.WriteUvarint8(9, m.IsFinalMessage())
	w.WriteFixed32Float32(10, m.FloatToIntPriceMultiplier())
	w.WriteFixed32Float32(11, m.IntToFloatPriceDivisor())
	w.WriteString(12, m.UnderlyingSymbol())
	w.WriteUvarint8(13, m.UpdatesBidAskOnly())
	w.WriteFixed32Float32(14, m.StrikePrice())
	w.WriteUvarint8(15, uint8(m.PutOrCall()))
	w.WriteUvarint32(16, m.ShortInterest())
	w.WriteFixed32Uint32(17, uint32(m.SecurityExpirationDate()))
	w.WriteFixed32Float32(18, m.BuyRolloverInterest())
	w.WriteFixed32Float32(19, m.SellRolloverInterest())
	w.WriteFixed32Float32(20, m.EarningsPerShare())
	w.WriteUvarint32(21, m.SharesOutstanding())
	w.WriteFixed32Float32(22, m.IntToFloatQuantityDivisor())
	w.WriteUvarint8(23, m.HasMarketDepthData())
	w.WriteFixed32Float32(24, m.DisplayPriceMultiplier())
	w.WriteString(25, m.ExchangeSymbol())
	w.WriteFixed32Float32(26, m.InitialMarginRequirement())
	w.WriteFixed32Float32(27, m.MaintenanceMarginRequirement())
	w.WriteString(28, m.Currency())
	w.WriteFixed32Float32(29, m.ContractSize())
	w.WriteUvarint32(30, m.OpenInterest())
	w.WriteFixed32Uint32(31, uint32(m.RolloverDate()))
	w.WriteUvarint8(32, m.IsDelayed())
	return w.Finish(), nil
}
