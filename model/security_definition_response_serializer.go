//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SecurityDefinitionResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":507,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Float32(m.MinPriceIncrement())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.PriceDisplayFormat()))
	w.Data = append(w.Data, ',')
	w.Float32(m.CurrencyValuePerIncrement())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.Float32(m.FloatToIntPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatPriceDivisor())
	w.Data = append(w.Data, ',')
	w.String(m.UnderlyingSymbol())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UpdatesBidAskOnly())
	w.Data = append(w.Data, ',')
	w.Float32(m.StrikePrice())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.PutOrCall()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.ShortInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SecurityExpirationDate()))
	w.Data = append(w.Data, ',')
	w.Float32(m.BuyRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.SellRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.EarningsPerShare())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SharesOutstanding())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatQuantityDivisor())
	w.Data = append(w.Data, ',')
	w.Uint8(m.HasMarketDepthData())
	w.Data = append(w.Data, ',')
	w.Float32(m.DisplayPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeSymbol())
	w.Data = append(w.Data, ',')
	w.Float32(m.InitialMarginRequirement())
	w.Data = append(w.Data, ',')
	w.Float32(m.MaintenanceMarginRequirement())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.Float32(m.ContractSize())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.RolloverDate()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDelayed())
	return w.Finish(), nil
}

func (m SecurityDefinitionResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":507,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Float32(m.MinPriceIncrement())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.PriceDisplayFormat()))
	w.Data = append(w.Data, ',')
	w.Float32(m.CurrencyValuePerIncrement())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.Float32(m.FloatToIntPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatPriceDivisor())
	w.Data = append(w.Data, ',')
	w.String(m.UnderlyingSymbol())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UpdatesBidAskOnly())
	w.Data = append(w.Data, ',')
	w.Float32(m.StrikePrice())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.PutOrCall()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.ShortInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SecurityExpirationDate()))
	w.Data = append(w.Data, ',')
	w.Float32(m.BuyRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.SellRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.EarningsPerShare())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SharesOutstanding())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatQuantityDivisor())
	w.Data = append(w.Data, ',')
	w.Uint8(m.HasMarketDepthData())
	w.Data = append(w.Data, ',')
	w.Float32(m.DisplayPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeSymbol())
	w.Data = append(w.Data, ',')
	w.Float32(m.InitialMarginRequirement())
	w.Data = append(w.Data, ',')
	w.Float32(m.MaintenanceMarginRequirement())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.Float32(m.ContractSize())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.RolloverDate()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDelayed())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SecurityDefinitionResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":507,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Float32(m.MinPriceIncrement())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.PriceDisplayFormat()))
	w.Data = append(w.Data, ',')
	w.Float32(m.CurrencyValuePerIncrement())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.Float32(m.FloatToIntPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatPriceDivisor())
	w.Data = append(w.Data, ',')
	w.String(m.UnderlyingSymbol())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UpdatesBidAskOnly())
	w.Data = append(w.Data, ',')
	w.Float32(m.StrikePrice())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.PutOrCall()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.ShortInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SecurityExpirationDate()))
	w.Data = append(w.Data, ',')
	w.Float32(m.BuyRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.SellRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.EarningsPerShare())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SharesOutstanding())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatQuantityDivisor())
	w.Data = append(w.Data, ',')
	w.Uint8(m.HasMarketDepthData())
	w.Data = append(w.Data, ',')
	w.Float32(m.DisplayPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeSymbol())
	w.Data = append(w.Data, ',')
	w.Float32(m.InitialMarginRequirement())
	w.Data = append(w.Data, ',')
	w.Float32(m.MaintenanceMarginRequirement())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.Float32(m.ContractSize())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.RolloverDate()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDelayed())
	return w.Finish(), nil
}

func (m SecurityDefinitionResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":507,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Float32(m.MinPriceIncrement())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.PriceDisplayFormat()))
	w.Data = append(w.Data, ',')
	w.Float32(m.CurrencyValuePerIncrement())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.Float32(m.FloatToIntPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatPriceDivisor())
	w.Data = append(w.Data, ',')
	w.String(m.UnderlyingSymbol())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UpdatesBidAskOnly())
	w.Data = append(w.Data, ',')
	w.Float32(m.StrikePrice())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.PutOrCall()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.ShortInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SecurityExpirationDate()))
	w.Data = append(w.Data, ',')
	w.Float32(m.BuyRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.SellRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.EarningsPerShare())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SharesOutstanding())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatQuantityDivisor())
	w.Data = append(w.Data, ',')
	w.Uint8(m.HasMarketDepthData())
	w.Data = append(w.Data, ',')
	w.Float32(m.DisplayPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeSymbol())
	w.Data = append(w.Data, ',')
	w.Float32(m.InitialMarginRequirement())
	w.Data = append(w.Data, ',')
	w.Float32(m.MaintenanceMarginRequirement())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.Float32(m.ContractSize())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.RolloverDate()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDelayed())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SecurityDefinitionResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 507)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.StringField("Description", m.Description())
	w.Float32Field("MinPriceIncrement", m.MinPriceIncrement())
	w.Int32Field("PriceDisplayFormat", int32(m.PriceDisplayFormat()))
	w.Float32Field("CurrencyValuePerIncrement", m.CurrencyValuePerIncrement())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.Float32Field("FloatToIntPriceMultiplier", m.FloatToIntPriceMultiplier())
	w.Float32Field("IntToFloatPriceDivisor", m.IntToFloatPriceDivisor())
	w.StringField("UnderlyingSymbol", m.UnderlyingSymbol())
	w.Uint8Field("UpdatesBidAskOnly", m.UpdatesBidAskOnly())
	w.Float32Field("StrikePrice", m.StrikePrice())
	w.Uint8Field("PutOrCall", uint8(m.PutOrCall()))
	w.Uint32Field("ShortInterest", m.ShortInterest())
	w.Uint32Field("SecurityExpirationDate", uint32(m.SecurityExpirationDate()))
	w.Float32Field("BuyRolloverInterest", m.BuyRolloverInterest())
	w.Float32Field("SellRolloverInterest", m.SellRolloverInterest())
	w.Float32Field("EarningsPerShare", m.EarningsPerShare())
	w.Uint32Field("SharesOutstanding", m.SharesOutstanding())
	w.Float32Field("IntToFloatQuantityDivisor", m.IntToFloatQuantityDivisor())
	w.Uint8Field("HasMarketDepthData", m.HasMarketDepthData())
	w.Float32Field("DisplayPriceMultiplier", m.DisplayPriceMultiplier())
	w.StringField("ExchangeSymbol", m.ExchangeSymbol())
	w.Float32Field("InitialMarginRequirement", m.InitialMarginRequirement())
	w.Float32Field("MaintenanceMarginRequirement", m.MaintenanceMarginRequirement())
	w.StringField("Currency", m.Currency())
	w.Float32Field("ContractSize", m.ContractSize())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Uint32Field("RolloverDate", uint32(m.RolloverDate()))
	w.BoolField("IsDelayed", m.IsDelayed())
	return w.Finish(), nil
}

func (m SecurityDefinitionResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 507)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.StringField("Description", m.Description())
	w.Float32Field("MinPriceIncrement", m.MinPriceIncrement())
	w.Int32Field("PriceDisplayFormat", int32(m.PriceDisplayFormat()))
	w.Float32Field("CurrencyValuePerIncrement", m.CurrencyValuePerIncrement())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.Float32Field("FloatToIntPriceMultiplier", m.FloatToIntPriceMultiplier())
	w.Float32Field("IntToFloatPriceDivisor", m.IntToFloatPriceDivisor())
	w.StringField("UnderlyingSymbol", m.UnderlyingSymbol())
	w.Uint8Field("UpdatesBidAskOnly", m.UpdatesBidAskOnly())
	w.Float32Field("StrikePrice", m.StrikePrice())
	w.Uint8Field("PutOrCall", uint8(m.PutOrCall()))
	w.Uint32Field("ShortInterest", m.ShortInterest())
	w.Uint32Field("SecurityExpirationDate", uint32(m.SecurityExpirationDate()))
	w.Float32Field("BuyRolloverInterest", m.BuyRolloverInterest())
	w.Float32Field("SellRolloverInterest", m.SellRolloverInterest())
	w.Float32Field("EarningsPerShare", m.EarningsPerShare())
	w.Uint32Field("SharesOutstanding", m.SharesOutstanding())
	w.Float32Field("IntToFloatQuantityDivisor", m.IntToFloatQuantityDivisor())
	w.Uint8Field("HasMarketDepthData", m.HasMarketDepthData())
	w.Float32Field("DisplayPriceMultiplier", m.DisplayPriceMultiplier())
	w.StringField("ExchangeSymbol", m.ExchangeSymbol())
	w.Float32Field("InitialMarginRequirement", m.InitialMarginRequirement())
	w.Float32Field("MaintenanceMarginRequirement", m.MaintenanceMarginRequirement())
	w.StringField("Currency", m.Currency())
	w.Float32Field("ContractSize", m.ContractSize())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Uint32Field("RolloverDate", uint32(m.RolloverDate()))
	w.BoolField("IsDelayed", m.IsDelayed())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SecurityDefinitionResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 507)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.StringField("Description", m.Description())
	w.Float32Field("MinPriceIncrement", m.MinPriceIncrement())
	w.Int32Field("PriceDisplayFormat", int32(m.PriceDisplayFormat()))
	w.Float32Field("CurrencyValuePerIncrement", m.CurrencyValuePerIncrement())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.Float32Field("FloatToIntPriceMultiplier", m.FloatToIntPriceMultiplier())
	w.Float32Field("IntToFloatPriceDivisor", m.IntToFloatPriceDivisor())
	w.StringField("UnderlyingSymbol", m.UnderlyingSymbol())
	w.Uint8Field("UpdatesBidAskOnly", m.UpdatesBidAskOnly())
	w.Float32Field("StrikePrice", m.StrikePrice())
	w.Uint8Field("PutOrCall", uint8(m.PutOrCall()))
	w.Uint32Field("ShortInterest", m.ShortInterest())
	w.Uint32Field("SecurityExpirationDate", uint32(m.SecurityExpirationDate()))
	w.Float32Field("BuyRolloverInterest", m.BuyRolloverInterest())
	w.Float32Field("SellRolloverInterest", m.SellRolloverInterest())
	w.Float32Field("EarningsPerShare", m.EarningsPerShare())
	w.Uint32Field("SharesOutstanding", m.SharesOutstanding())
	w.Float32Field("IntToFloatQuantityDivisor", m.IntToFloatQuantityDivisor())
	w.Uint8Field("HasMarketDepthData", m.HasMarketDepthData())
	w.Float32Field("DisplayPriceMultiplier", m.DisplayPriceMultiplier())
	w.StringField("ExchangeSymbol", m.ExchangeSymbol())
	w.Float32Field("InitialMarginRequirement", m.InitialMarginRequirement())
	w.Float32Field("MaintenanceMarginRequirement", m.MaintenanceMarginRequirement())
	w.StringField("Currency", m.Currency())
	w.Float32Field("ContractSize", m.ContractSize())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Uint32Field("RolloverDate", uint32(m.RolloverDate()))
	w.BoolField("IsDelayed", m.IsDelayed())
	return w.Finish(), nil
}

func (m SecurityDefinitionResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 507)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.StringField("Description", m.Description())
	w.Float32Field("MinPriceIncrement", m.MinPriceIncrement())
	w.Int32Field("PriceDisplayFormat", int32(m.PriceDisplayFormat()))
	w.Float32Field("CurrencyValuePerIncrement", m.CurrencyValuePerIncrement())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.Float32Field("FloatToIntPriceMultiplier", m.FloatToIntPriceMultiplier())
	w.Float32Field("IntToFloatPriceDivisor", m.IntToFloatPriceDivisor())
	w.StringField("UnderlyingSymbol", m.UnderlyingSymbol())
	w.Uint8Field("UpdatesBidAskOnly", m.UpdatesBidAskOnly())
	w.Float32Field("StrikePrice", m.StrikePrice())
	w.Uint8Field("PutOrCall", uint8(m.PutOrCall()))
	w.Uint32Field("ShortInterest", m.ShortInterest())
	w.Uint32Field("SecurityExpirationDate", uint32(m.SecurityExpirationDate()))
	w.Float32Field("BuyRolloverInterest", m.BuyRolloverInterest())
	w.Float32Field("SellRolloverInterest", m.SellRolloverInterest())
	w.Float32Field("EarningsPerShare", m.EarningsPerShare())
	w.Uint32Field("SharesOutstanding", m.SharesOutstanding())
	w.Float32Field("IntToFloatQuantityDivisor", m.IntToFloatQuantityDivisor())
	w.Uint8Field("HasMarketDepthData", m.HasMarketDepthData())
	w.Float32Field("DisplayPriceMultiplier", m.DisplayPriceMultiplier())
	w.StringField("ExchangeSymbol", m.ExchangeSymbol())
	w.Float32Field("InitialMarginRequirement", m.InitialMarginRequirement())
	w.Float32Field("MaintenanceMarginRequirement", m.MaintenanceMarginRequirement())
	w.StringField("Currency", m.Currency())
	w.Float32Field("ContractSize", m.ContractSize())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Uint32Field("RolloverDate", uint32(m.RolloverDate()))
	w.BoolField("IsDelayed", m.IsDelayed())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SecurityDefinitionResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 507 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityType(SecurityTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetMinPriceIncrement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriceDisplayFormat(PriceDisplayFormatEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrencyValuePerIncrement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalMessage(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetFloatToIntPriceMultiplier(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIntToFloatPriceDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnderlyingSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdatesBidAskOnly(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetStrikePrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPutOrCall(PutCallEnum(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetShortInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityExpirationDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetBuyRolloverInterest(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSellRolloverInterest(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetEarningsPerShare(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSharesOutstanding(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIntToFloatQuantityDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetHasMarketDepthData(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDisplayPriceMultiplier(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchangeSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetInitialMarginRequirement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaintenanceMarginRequirement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetContractSize(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRolloverDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetIsDelayed(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SecurityDefinitionResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 507 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityType(SecurityTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetMinPriceIncrement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriceDisplayFormat(PriceDisplayFormatEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrencyValuePerIncrement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalMessage(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetFloatToIntPriceMultiplier(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIntToFloatPriceDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnderlyingSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdatesBidAskOnly(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetStrikePrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPutOrCall(PutCallEnum(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetShortInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityExpirationDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetBuyRolloverInterest(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSellRolloverInterest(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetEarningsPerShare(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSharesOutstanding(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIntToFloatQuantityDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetHasMarketDepthData(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDisplayPriceMultiplier(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchangeSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetInitialMarginRequirement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaintenanceMarginRequirement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetContractSize(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRolloverDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetIsDelayed(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SecurityDefinitionResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 507 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "RequestID":
			m.SetRequestID(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
		case "Description":
			m.SetDescription(r.String())
		case "MinPriceIncrement":
			m.SetMinPriceIncrement(r.Float32())
		case "PriceDisplayFormat":
			m.SetPriceDisplayFormat(PriceDisplayFormatEnum(r.Int32()))
		case "CurrencyValuePerIncrement":
			m.SetCurrencyValuePerIncrement(r.Float32())
		case "IsFinalMessage":
			m.SetIsFinalMessage(r.Bool())
		case "FloatToIntPriceMultiplier":
			m.SetFloatToIntPriceMultiplier(r.Float32())
		case "IntToFloatPriceDivisor":
			m.SetIntToFloatPriceDivisor(r.Float32())
		case "UnderlyingSymbol":
			m.SetUnderlyingSymbol(r.String())
		case "UpdatesBidAskOnly":
			m.SetUpdatesBidAskOnly(r.Uint8())
		case "StrikePrice":
			m.SetStrikePrice(r.Float32())
		case "PutOrCall":
			m.SetPutOrCall(PutCallEnum(r.Uint8()))
		case "ShortInterest":
			m.SetShortInterest(r.Uint32())
		case "SecurityExpirationDate":
			m.SetSecurityExpirationDate(DateTime4Byte(r.Uint32()))
		case "BuyRolloverInterest":
			m.SetBuyRolloverInterest(r.Float32())
		case "SellRolloverInterest":
			m.SetSellRolloverInterest(r.Float32())
		case "EarningsPerShare":
			m.SetEarningsPerShare(r.Float32())
		case "SharesOutstanding":
			m.SetSharesOutstanding(r.Uint32())
		case "IntToFloatQuantityDivisor":
			m.SetIntToFloatQuantityDivisor(r.Float32())
		case "HasMarketDepthData":
			m.SetHasMarketDepthData(r.Uint8())
		case "DisplayPriceMultiplier":
			m.SetDisplayPriceMultiplier(r.Float32())
		case "ExchangeSymbol":
			m.SetExchangeSymbol(r.String())
		case "InitialMarginRequirement":
			m.SetInitialMarginRequirement(r.Float32())
		case "MaintenanceMarginRequirement":
			m.SetMaintenanceMarginRequirement(r.Float32())
		case "Currency":
			m.SetCurrency(r.String())
		case "ContractSize":
			m.SetContractSize(r.Float32())
		case "OpenInterest":
			m.SetOpenInterest(r.Uint32())
		case "RolloverDate":
			m.SetRolloverDate(DateTime4Byte(r.Uint32()))
		case "IsDelayed":
			m.SetIsDelayed(r.Bool())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SecurityDefinitionResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 507 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "RequestID":
			m.SetRequestID(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
		case "Description":
			m.SetDescription(r.String())
		case "MinPriceIncrement":
			m.SetMinPriceIncrement(r.Float32())
		case "PriceDisplayFormat":
			m.SetPriceDisplayFormat(PriceDisplayFormatEnum(r.Int32()))
		case "CurrencyValuePerIncrement":
			m.SetCurrencyValuePerIncrement(r.Float32())
		case "IsFinalMessage":
			m.SetIsFinalMessage(r.Bool())
		case "FloatToIntPriceMultiplier":
			m.SetFloatToIntPriceMultiplier(r.Float32())
		case "IntToFloatPriceDivisor":
			m.SetIntToFloatPriceDivisor(r.Float32())
		case "UnderlyingSymbol":
			m.SetUnderlyingSymbol(r.String())
		case "UpdatesBidAskOnly":
			m.SetUpdatesBidAskOnly(r.Uint8())
		case "StrikePrice":
			m.SetStrikePrice(r.Float32())
		case "PutOrCall":
			m.SetPutOrCall(PutCallEnum(r.Uint8()))
		case "ShortInterest":
			m.SetShortInterest(r.Uint32())
		case "SecurityExpirationDate":
			m.SetSecurityExpirationDate(DateTime4Byte(r.Uint32()))
		case "BuyRolloverInterest":
			m.SetBuyRolloverInterest(r.Float32())
		case "SellRolloverInterest":
			m.SetSellRolloverInterest(r.Float32())
		case "EarningsPerShare":
			m.SetEarningsPerShare(r.Float32())
		case "SharesOutstanding":
			m.SetSharesOutstanding(r.Uint32())
		case "IntToFloatQuantityDivisor":
			m.SetIntToFloatQuantityDivisor(r.Float32())
		case "HasMarketDepthData":
			m.SetHasMarketDepthData(r.Uint8())
		case "DisplayPriceMultiplier":
			m.SetDisplayPriceMultiplier(r.Float32())
		case "ExchangeSymbol":
			m.SetExchangeSymbol(r.String())
		case "InitialMarginRequirement":
			m.SetInitialMarginRequirement(r.Float32())
		case "MaintenanceMarginRequirement":
			m.SetMaintenanceMarginRequirement(r.Float32())
		case "Currency":
			m.SetCurrency(r.String())
		case "ContractSize":
			m.SetContractSize(r.Float32())
		case "OpenInterest":
			m.SetOpenInterest(r.Uint32())
		case "RolloverDate":
			m.SetRolloverDate(DateTime4Byte(r.Uint32()))
		case "IsDelayed":
			m.SetIsDelayed(r.Bool())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SecurityDefinitionResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":507,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Float32(m.MinPriceIncrement())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.PriceDisplayFormat()))
	w.Data = append(w.Data, ',')
	w.Float32(m.CurrencyValuePerIncrement())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.Float32(m.FloatToIntPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatPriceDivisor())
	w.Data = append(w.Data, ',')
	w.String(m.UnderlyingSymbol())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UpdatesBidAskOnly())
	w.Data = append(w.Data, ',')
	w.Float32(m.StrikePrice())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.PutOrCall()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.ShortInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SecurityExpirationDate()))
	w.Data = append(w.Data, ',')
	w.Float32(m.BuyRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.SellRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.EarningsPerShare())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SharesOutstanding())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatQuantityDivisor())
	w.Data = append(w.Data, ',')
	w.Uint8(m.HasMarketDepthData())
	w.Data = append(w.Data, ',')
	w.Float32(m.DisplayPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeSymbol())
	w.Data = append(w.Data, ',')
	w.Float32(m.InitialMarginRequirement())
	w.Data = append(w.Data, ',')
	w.Float32(m.MaintenanceMarginRequirement())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.Float32(m.ContractSize())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.RolloverDate()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDelayed())
	return w.Finish(), nil
}

func (m SecurityDefinitionResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":507,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Float32(m.MinPriceIncrement())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.PriceDisplayFormat()))
	w.Data = append(w.Data, ',')
	w.Float32(m.CurrencyValuePerIncrement())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.Float32(m.FloatToIntPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatPriceDivisor())
	w.Data = append(w.Data, ',')
	w.String(m.UnderlyingSymbol())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UpdatesBidAskOnly())
	w.Data = append(w.Data, ',')
	w.Float32(m.StrikePrice())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.PutOrCall()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.ShortInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SecurityExpirationDate()))
	w.Data = append(w.Data, ',')
	w.Float32(m.BuyRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.SellRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.EarningsPerShare())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SharesOutstanding())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatQuantityDivisor())
	w.Data = append(w.Data, ',')
	w.Uint8(m.HasMarketDepthData())
	w.Data = append(w.Data, ',')
	w.Float32(m.DisplayPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeSymbol())
	w.Data = append(w.Data, ',')
	w.Float32(m.InitialMarginRequirement())
	w.Data = append(w.Data, ',')
	w.Float32(m.MaintenanceMarginRequirement())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.Float32(m.ContractSize())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.RolloverDate()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDelayed())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SecurityDefinitionResponseFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":507,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Float32(m.MinPriceIncrement())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.PriceDisplayFormat()))
	w.Data = append(w.Data, ',')
	w.Float32(m.CurrencyValuePerIncrement())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.Float32(m.FloatToIntPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatPriceDivisor())
	w.Data = append(w.Data, ',')
	w.String(m.UnderlyingSymbol())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UpdatesBidAskOnly())
	w.Data = append(w.Data, ',')
	w.Float32(m.StrikePrice())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.PutOrCall()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.ShortInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SecurityExpirationDate()))
	w.Data = append(w.Data, ',')
	w.Float32(m.BuyRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.SellRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.EarningsPerShare())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SharesOutstanding())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatQuantityDivisor())
	w.Data = append(w.Data, ',')
	w.Uint8(m.HasMarketDepthData())
	w.Data = append(w.Data, ',')
	w.Float32(m.DisplayPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeSymbol())
	w.Data = append(w.Data, ',')
	w.Float32(m.InitialMarginRequirement())
	w.Data = append(w.Data, ',')
	w.Float32(m.MaintenanceMarginRequirement())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.Float32(m.ContractSize())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.RolloverDate()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDelayed())
	return w.Finish(), nil
}

func (m SecurityDefinitionResponseFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":507,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	w.Data = append(w.Data, ',')
	w.Float32(m.MinPriceIncrement())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.PriceDisplayFormat()))
	w.Data = append(w.Data, ',')
	w.Float32(m.CurrencyValuePerIncrement())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.Float32(m.FloatToIntPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatPriceDivisor())
	w.Data = append(w.Data, ',')
	w.String(m.UnderlyingSymbol())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UpdatesBidAskOnly())
	w.Data = append(w.Data, ',')
	w.Float32(m.StrikePrice())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.PutOrCall()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.ShortInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.SecurityExpirationDate()))
	w.Data = append(w.Data, ',')
	w.Float32(m.BuyRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.SellRolloverInterest())
	w.Data = append(w.Data, ',')
	w.Float32(m.EarningsPerShare())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SharesOutstanding())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatQuantityDivisor())
	w.Data = append(w.Data, ',')
	w.Uint8(m.HasMarketDepthData())
	w.Data = append(w.Data, ',')
	w.Float32(m.DisplayPriceMultiplier())
	w.Data = append(w.Data, ',')
	w.String(m.ExchangeSymbol())
	w.Data = append(w.Data, ',')
	w.Float32(m.InitialMarginRequirement())
	w.Data = append(w.Data, ',')
	w.Float32(m.MaintenanceMarginRequirement())
	w.Data = append(w.Data, ',')
	w.String(m.Currency())
	w.Data = append(w.Data, ',')
	w.Float32(m.ContractSize())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.RolloverDate()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDelayed())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SecurityDefinitionResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 507)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.StringField("Description", m.Description())
	w.Float32Field("MinPriceIncrement", m.MinPriceIncrement())
	w.Int32Field("PriceDisplayFormat", int32(m.PriceDisplayFormat()))
	w.Float32Field("CurrencyValuePerIncrement", m.CurrencyValuePerIncrement())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.Float32Field("FloatToIntPriceMultiplier", m.FloatToIntPriceMultiplier())
	w.Float32Field("IntToFloatPriceDivisor", m.IntToFloatPriceDivisor())
	w.StringField("UnderlyingSymbol", m.UnderlyingSymbol())
	w.Uint8Field("UpdatesBidAskOnly", m.UpdatesBidAskOnly())
	w.Float32Field("StrikePrice", m.StrikePrice())
	w.Uint8Field("PutOrCall", uint8(m.PutOrCall()))
	w.Uint32Field("ShortInterest", m.ShortInterest())
	w.Uint32Field("SecurityExpirationDate", uint32(m.SecurityExpirationDate()))
	w.Float32Field("BuyRolloverInterest", m.BuyRolloverInterest())
	w.Float32Field("SellRolloverInterest", m.SellRolloverInterest())
	w.Float32Field("EarningsPerShare", m.EarningsPerShare())
	w.Uint32Field("SharesOutstanding", m.SharesOutstanding())
	w.Float32Field("IntToFloatQuantityDivisor", m.IntToFloatQuantityDivisor())
	w.Uint8Field("HasMarketDepthData", m.HasMarketDepthData())
	w.Float32Field("DisplayPriceMultiplier", m.DisplayPriceMultiplier())
	w.StringField("ExchangeSymbol", m.ExchangeSymbol())
	w.Float32Field("InitialMarginRequirement", m.InitialMarginRequirement())
	w.Float32Field("MaintenanceMarginRequirement", m.MaintenanceMarginRequirement())
	w.StringField("Currency", m.Currency())
	w.Float32Field("ContractSize", m.ContractSize())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Uint32Field("RolloverDate", uint32(m.RolloverDate()))
	w.BoolField("IsDelayed", m.IsDelayed())
	return w.Finish(), nil
}

func (m SecurityDefinitionResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 507)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.StringField("Description", m.Description())
	w.Float32Field("MinPriceIncrement", m.MinPriceIncrement())
	w.Int32Field("PriceDisplayFormat", int32(m.PriceDisplayFormat()))
	w.Float32Field("CurrencyValuePerIncrement", m.CurrencyValuePerIncrement())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.Float32Field("FloatToIntPriceMultiplier", m.FloatToIntPriceMultiplier())
	w.Float32Field("IntToFloatPriceDivisor", m.IntToFloatPriceDivisor())
	w.StringField("UnderlyingSymbol", m.UnderlyingSymbol())
	w.Uint8Field("UpdatesBidAskOnly", m.UpdatesBidAskOnly())
	w.Float32Field("StrikePrice", m.StrikePrice())
	w.Uint8Field("PutOrCall", uint8(m.PutOrCall()))
	w.Uint32Field("ShortInterest", m.ShortInterest())
	w.Uint32Field("SecurityExpirationDate", uint32(m.SecurityExpirationDate()))
	w.Float32Field("BuyRolloverInterest", m.BuyRolloverInterest())
	w.Float32Field("SellRolloverInterest", m.SellRolloverInterest())
	w.Float32Field("EarningsPerShare", m.EarningsPerShare())
	w.Uint32Field("SharesOutstanding", m.SharesOutstanding())
	w.Float32Field("IntToFloatQuantityDivisor", m.IntToFloatQuantityDivisor())
	w.Uint8Field("HasMarketDepthData", m.HasMarketDepthData())
	w.Float32Field("DisplayPriceMultiplier", m.DisplayPriceMultiplier())
	w.StringField("ExchangeSymbol", m.ExchangeSymbol())
	w.Float32Field("InitialMarginRequirement", m.InitialMarginRequirement())
	w.Float32Field("MaintenanceMarginRequirement", m.MaintenanceMarginRequirement())
	w.StringField("Currency", m.Currency())
	w.Float32Field("ContractSize", m.ContractSize())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Uint32Field("RolloverDate", uint32(m.RolloverDate()))
	w.BoolField("IsDelayed", m.IsDelayed())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SecurityDefinitionResponseFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 507)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.StringField("Description", m.Description())
	w.Float32Field("MinPriceIncrement", m.MinPriceIncrement())
	w.Int32Field("PriceDisplayFormat", int32(m.PriceDisplayFormat()))
	w.Float32Field("CurrencyValuePerIncrement", m.CurrencyValuePerIncrement())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.Float32Field("FloatToIntPriceMultiplier", m.FloatToIntPriceMultiplier())
	w.Float32Field("IntToFloatPriceDivisor", m.IntToFloatPriceDivisor())
	w.StringField("UnderlyingSymbol", m.UnderlyingSymbol())
	w.Uint8Field("UpdatesBidAskOnly", m.UpdatesBidAskOnly())
	w.Float32Field("StrikePrice", m.StrikePrice())
	w.Uint8Field("PutOrCall", uint8(m.PutOrCall()))
	w.Uint32Field("ShortInterest", m.ShortInterest())
	w.Uint32Field("SecurityExpirationDate", uint32(m.SecurityExpirationDate()))
	w.Float32Field("BuyRolloverInterest", m.BuyRolloverInterest())
	w.Float32Field("SellRolloverInterest", m.SellRolloverInterest())
	w.Float32Field("EarningsPerShare", m.EarningsPerShare())
	w.Uint32Field("SharesOutstanding", m.SharesOutstanding())
	w.Float32Field("IntToFloatQuantityDivisor", m.IntToFloatQuantityDivisor())
	w.Uint8Field("HasMarketDepthData", m.HasMarketDepthData())
	w.Float32Field("DisplayPriceMultiplier", m.DisplayPriceMultiplier())
	w.StringField("ExchangeSymbol", m.ExchangeSymbol())
	w.Float32Field("InitialMarginRequirement", m.InitialMarginRequirement())
	w.Float32Field("MaintenanceMarginRequirement", m.MaintenanceMarginRequirement())
	w.StringField("Currency", m.Currency())
	w.Float32Field("ContractSize", m.ContractSize())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Uint32Field("RolloverDate", uint32(m.RolloverDate()))
	w.BoolField("IsDelayed", m.IsDelayed())
	return w.Finish(), nil
}

func (m SecurityDefinitionResponseFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 507)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.StringField("Description", m.Description())
	w.Float32Field("MinPriceIncrement", m.MinPriceIncrement())
	w.Int32Field("PriceDisplayFormat", int32(m.PriceDisplayFormat()))
	w.Float32Field("CurrencyValuePerIncrement", m.CurrencyValuePerIncrement())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.Float32Field("FloatToIntPriceMultiplier", m.FloatToIntPriceMultiplier())
	w.Float32Field("IntToFloatPriceDivisor", m.IntToFloatPriceDivisor())
	w.StringField("UnderlyingSymbol", m.UnderlyingSymbol())
	w.Uint8Field("UpdatesBidAskOnly", m.UpdatesBidAskOnly())
	w.Float32Field("StrikePrice", m.StrikePrice())
	w.Uint8Field("PutOrCall", uint8(m.PutOrCall()))
	w.Uint32Field("ShortInterest", m.ShortInterest())
	w.Uint32Field("SecurityExpirationDate", uint32(m.SecurityExpirationDate()))
	w.Float32Field("BuyRolloverInterest", m.BuyRolloverInterest())
	w.Float32Field("SellRolloverInterest", m.SellRolloverInterest())
	w.Float32Field("EarningsPerShare", m.EarningsPerShare())
	w.Uint32Field("SharesOutstanding", m.SharesOutstanding())
	w.Float32Field("IntToFloatQuantityDivisor", m.IntToFloatQuantityDivisor())
	w.Uint8Field("HasMarketDepthData", m.HasMarketDepthData())
	w.Float32Field("DisplayPriceMultiplier", m.DisplayPriceMultiplier())
	w.StringField("ExchangeSymbol", m.ExchangeSymbol())
	w.Float32Field("InitialMarginRequirement", m.InitialMarginRequirement())
	w.Float32Field("MaintenanceMarginRequirement", m.MaintenanceMarginRequirement())
	w.StringField("Currency", m.Currency())
	w.Float32Field("ContractSize", m.ContractSize())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Uint32Field("RolloverDate", uint32(m.RolloverDate()))
	w.BoolField("IsDelayed", m.IsDelayed())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SecurityDefinitionResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 507 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityType(SecurityTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetMinPriceIncrement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriceDisplayFormat(PriceDisplayFormatEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrencyValuePerIncrement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalMessage(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetFloatToIntPriceMultiplier(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIntToFloatPriceDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnderlyingSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdatesBidAskOnly(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetStrikePrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPutOrCall(PutCallEnum(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetShortInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityExpirationDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetBuyRolloverInterest(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSellRolloverInterest(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetEarningsPerShare(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSharesOutstanding(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIntToFloatQuantityDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetHasMarketDepthData(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDisplayPriceMultiplier(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchangeSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetInitialMarginRequirement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaintenanceMarginRequirement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetContractSize(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRolloverDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetIsDelayed(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SecurityDefinitionResponseFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 507 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityType(SecurityTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetMinPriceIncrement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPriceDisplayFormat(PriceDisplayFormatEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrencyValuePerIncrement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalMessage(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetFloatToIntPriceMultiplier(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIntToFloatPriceDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnderlyingSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdatesBidAskOnly(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetStrikePrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPutOrCall(PutCallEnum(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetShortInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSecurityExpirationDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetBuyRolloverInterest(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSellRolloverInterest(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetEarningsPerShare(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSharesOutstanding(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIntToFloatQuantityDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetHasMarketDepthData(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetDisplayPriceMultiplier(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchangeSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetInitialMarginRequirement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaintenanceMarginRequirement(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetContractSize(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRolloverDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetIsDelayed(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SecurityDefinitionResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 507 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "RequestID":
			m.SetRequestID(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
		case "Description":
			m.SetDescription(r.String())
		case "MinPriceIncrement":
			m.SetMinPriceIncrement(r.Float32())
		case "PriceDisplayFormat":
			m.SetPriceDisplayFormat(PriceDisplayFormatEnum(r.Int32()))
		case "CurrencyValuePerIncrement":
			m.SetCurrencyValuePerIncrement(r.Float32())
		case "IsFinalMessage":
			m.SetIsFinalMessage(r.Bool())
		case "FloatToIntPriceMultiplier":
			m.SetFloatToIntPriceMultiplier(r.Float32())
		case "IntToFloatPriceDivisor":
			m.SetIntToFloatPriceDivisor(r.Float32())
		case "UnderlyingSymbol":
			m.SetUnderlyingSymbol(r.String())
		case "UpdatesBidAskOnly":
			m.SetUpdatesBidAskOnly(r.Uint8())
		case "StrikePrice":
			m.SetStrikePrice(r.Float32())
		case "PutOrCall":
			m.SetPutOrCall(PutCallEnum(r.Uint8()))
		case "ShortInterest":
			m.SetShortInterest(r.Uint32())
		case "SecurityExpirationDate":
			m.SetSecurityExpirationDate(DateTime4Byte(r.Uint32()))
		case "BuyRolloverInterest":
			m.SetBuyRolloverInterest(r.Float32())
		case "SellRolloverInterest":
			m.SetSellRolloverInterest(r.Float32())
		case "EarningsPerShare":
			m.SetEarningsPerShare(r.Float32())
		case "SharesOutstanding":
			m.SetSharesOutstanding(r.Uint32())
		case "IntToFloatQuantityDivisor":
			m.SetIntToFloatQuantityDivisor(r.Float32())
		case "HasMarketDepthData":
			m.SetHasMarketDepthData(r.Uint8())
		case "DisplayPriceMultiplier":
			m.SetDisplayPriceMultiplier(r.Float32())
		case "ExchangeSymbol":
			m.SetExchangeSymbol(r.String())
		case "InitialMarginRequirement":
			m.SetInitialMarginRequirement(r.Float32())
		case "MaintenanceMarginRequirement":
			m.SetMaintenanceMarginRequirement(r.Float32())
		case "Currency":
			m.SetCurrency(r.String())
		case "ContractSize":
			m.SetContractSize(r.Float32())
		case "OpenInterest":
			m.SetOpenInterest(r.Uint32())
		case "RolloverDate":
			m.SetRolloverDate(DateTime4Byte(r.Uint32()))
		case "IsDelayed":
			m.SetIsDelayed(r.Bool())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SecurityDefinitionResponseFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 507 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "RequestID":
			m.SetRequestID(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
		case "Description":
			m.SetDescription(r.String())
		case "MinPriceIncrement":
			m.SetMinPriceIncrement(r.Float32())
		case "PriceDisplayFormat":
			m.SetPriceDisplayFormat(PriceDisplayFormatEnum(r.Int32()))
		case "CurrencyValuePerIncrement":
			m.SetCurrencyValuePerIncrement(r.Float32())
		case "IsFinalMessage":
			m.SetIsFinalMessage(r.Bool())
		case "FloatToIntPriceMultiplier":
			m.SetFloatToIntPriceMultiplier(r.Float32())
		case "IntToFloatPriceDivisor":
			m.SetIntToFloatPriceDivisor(r.Float32())
		case "UnderlyingSymbol":
			m.SetUnderlyingSymbol(r.String())
		case "UpdatesBidAskOnly":
			m.SetUpdatesBidAskOnly(r.Uint8())
		case "StrikePrice":
			m.SetStrikePrice(r.Float32())
		case "PutOrCall":
			m.SetPutOrCall(PutCallEnum(r.Uint8()))
		case "ShortInterest":
			m.SetShortInterest(r.Uint32())
		case "SecurityExpirationDate":
			m.SetSecurityExpirationDate(DateTime4Byte(r.Uint32()))
		case "BuyRolloverInterest":
			m.SetBuyRolloverInterest(r.Float32())
		case "SellRolloverInterest":
			m.SetSellRolloverInterest(r.Float32())
		case "EarningsPerShare":
			m.SetEarningsPerShare(r.Float32())
		case "SharesOutstanding":
			m.SetSharesOutstanding(r.Uint32())
		case "IntToFloatQuantityDivisor":
			m.SetIntToFloatQuantityDivisor(r.Float32())
		case "HasMarketDepthData":
			m.SetHasMarketDepthData(r.Uint8())
		case "DisplayPriceMultiplier":
			m.SetDisplayPriceMultiplier(r.Float32())
		case "ExchangeSymbol":
			m.SetExchangeSymbol(r.String())
		case "InitialMarginRequirement":
			m.SetInitialMarginRequirement(r.Float32())
		case "MaintenanceMarginRequirement":
			m.SetMaintenanceMarginRequirement(r.Float32())
		case "Currency":
			m.SetCurrency(r.String())
		case "ContractSize":
			m.SetContractSize(r.Float32())
		case "OpenInterest":
			m.SetOpenInterest(r.Uint32())
		case "RolloverDate":
			m.SetRolloverDate(DateTime4Byte(r.Uint32()))
		case "IsDelayed":
			m.SetIsDelayed(r.Bool())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SecurityDefinitionResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetIsFinalMessage(r.ReadBool())
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
			m.SetIsDelayed(r.ReadBool())
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
			m.SetIsFinalMessage(r.ReadBool())
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
			m.SetIsDelayed(r.ReadBool())
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

func (m SecurityDefinitionResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 507)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint32(4, int32(m.SecurityType()))
	w.WriteString(5, m.Description())
	w.WriteFixed32Float32(6, m.MinPriceIncrement())
	w.WriteVarint32(7, int32(m.PriceDisplayFormat()))
	w.WriteFixed32Float32(8, m.CurrencyValuePerIncrement())
	w.WriteBool(9, m.IsFinalMessage())
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
	w.WriteBool(32, m.IsDelayed())
	return w.Finish(), nil
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
	w.WriteBool(9, m.IsFinalMessage())
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
	w.WriteBool(32, m.IsDelayed())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SecurityDefinitionResponseFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetIsFinalMessage(r.ReadBool())
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
			m.SetIsDelayed(r.ReadBool())
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
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SecurityDefinitionResponseFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetIsFinalMessage(r.ReadBool())
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
			m.SetIsDelayed(r.ReadBool())
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

func (m SecurityDefinitionResponseFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 507)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint32(4, int32(m.SecurityType()))
	w.WriteString(5, m.Description())
	w.WriteFixed32Float32(6, m.MinPriceIncrement())
	w.WriteVarint32(7, int32(m.PriceDisplayFormat()))
	w.WriteFixed32Float32(8, m.CurrencyValuePerIncrement())
	w.WriteBool(9, m.IsFinalMessage())
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
	w.WriteBool(32, m.IsDelayed())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SecurityDefinitionResponseFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 507)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint32(4, int32(m.SecurityType()))
	w.WriteString(5, m.Description())
	w.WriteFixed32Float32(6, m.MinPriceIncrement())
	w.WriteVarint32(7, int32(m.PriceDisplayFormat()))
	w.WriteFixed32Float32(8, m.CurrencyValuePerIncrement())
	w.WriteBool(9, m.IsFinalMessage())
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
	w.WriteBool(32, m.IsDelayed())
	return w.Finish(), nil
}
