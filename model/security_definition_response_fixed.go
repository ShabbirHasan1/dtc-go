package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                         = SecurityDefinitionResponseFixedSize  (356)
//     Type                         = SECURITY_DEFINITION_RESPONSE  (507)
//     RequestID                    = 0
//     Symbol                       = ""
//     Exchange                     = ""
//     SecurityType                 = 0
//     Description                  = ""
//     MinPriceIncrement            = 0
//     PriceDisplayFormat           = PRICE_DISPLAY_FORMAT_UNSET  (-1)
//     CurrencyValuePerIncrement    = 0
//     IsFinalMessage               = 0
//     FloatToIntPriceMultiplier    = 1.000000
//     IntToFloatPriceDivisor       = 1.000000
//     UnderlyingSymbol             = ""
//     UpdatesBidAskOnly            = 0
//     StrikePrice                  = 0
//     PutOrCall                    = PC_UNSET  (0)
//     ShortInterest                = 0
//     SecurityExpirationDate       = 0
//     BuyRolloverInterest          = 0
//     SellRolloverInterest         = 0
//     EarningsPerShare             = 0
//     SharesOutstanding            = 0
//     IntToFloatQuantityDivisor    = 0
//     HasMarketDepthData           = 1
//     DisplayPriceMultiplier       = 1.000000
//     ExchangeSymbol               = ""
//     InitialMarginRequirement     = 0
//     MaintenanceMarginRequirement = 0
//     Currency                     = ""
//     ContractSize                 = 0
//     OpenInterest                 = 0
//     RolloverDate                 = 0
//     IsDelayed                    = 0
// }
var _SecurityDefinitionResponseFixedDefault = []byte{100, 1, 251, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SecurityDefinitionResponseFixedSize = 356

// SecurityDefinitionResponseFixed This is a response from the Server in response to a SymbolsForExchangeRequest,
// UNDERLYING_SymbolsForExchangeRequest, SymbolsForUnderlyingRequest, SecurityDefinitionForSymbolRequest,
// SymbolSearchRequest message.
//
// If there are no symbols to return in response to a request, the Server
// needs to send through one of these messages with the RequestID set to
// the same RequestID value that the request message set it to, and IsFinalMessage
// = 1. Leave all other member fields in the default state and the Client
// will recognize there are no symbols.
//
// The Client must always send a SecurityDefinitionForSymbolRequest message
// to the Server to obtain the IntegerToFloatPriceDivisor and FloatToIntPriceMultiplier
// values in the Security Definition Response message when the Server uses
// the integer market data and order messages.
type SecurityDefinitionResponseFixed struct {
	message.Fixed
}

// SecurityDefinitionResponseFixedBuilder This is a response from the Server in response to a SymbolsForExchangeRequest,
// UNDERLYING_SymbolsForExchangeRequest, SymbolsForUnderlyingRequest, SecurityDefinitionForSymbolRequest,
// SymbolSearchRequest message.
//
// If there are no symbols to return in response to a request, the Server
// needs to send through one of these messages with the RequestID set to
// the same RequestID value that the request message set it to, and IsFinalMessage
// = 1. Leave all other member fields in the default state and the Client
// will recognize there are no symbols.
//
// The Client must always send a SecurityDefinitionForSymbolRequest message
// to the Server to obtain the IntegerToFloatPriceDivisor and FloatToIntPriceMultiplier
// values in the Security Definition Response message when the Server uses
// the integer market data and order messages.
type SecurityDefinitionResponseFixedBuilder struct {
	message.Fixed
}

func NewSecurityDefinitionResponseFixedFrom(b []byte) SecurityDefinitionResponseFixed {
	return SecurityDefinitionResponseFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapSecurityDefinitionResponseFixed(b []byte) SecurityDefinitionResponseFixed {
	return SecurityDefinitionResponseFixed{Fixed: message.WrapFixed(b)}
}

func NewSecurityDefinitionResponseFixed() SecurityDefinitionResponseFixedBuilder {
	a := SecurityDefinitionResponseFixedBuilder{message.NewFixed(356)}
	a.Unsafe().SetBytes(0, _SecurityDefinitionResponseFixedDefault)
	return a
}

// Clear
// {
//     Size                         = SecurityDefinitionResponseFixedSize  (356)
//     Type                         = SECURITY_DEFINITION_RESPONSE  (507)
//     RequestID                    = 0
//     Symbol                       = ""
//     Exchange                     = ""
//     SecurityType                 = 0
//     Description                  = ""
//     MinPriceIncrement            = 0
//     PriceDisplayFormat           = PRICE_DISPLAY_FORMAT_UNSET  (-1)
//     CurrencyValuePerIncrement    = 0
//     IsFinalMessage               = 0
//     FloatToIntPriceMultiplier    = 1.000000
//     IntToFloatPriceDivisor       = 1.000000
//     UnderlyingSymbol             = ""
//     UpdatesBidAskOnly            = 0
//     StrikePrice                  = 0
//     PutOrCall                    = PC_UNSET  (0)
//     ShortInterest                = 0
//     SecurityExpirationDate       = 0
//     BuyRolloverInterest          = 0
//     SellRolloverInterest         = 0
//     EarningsPerShare             = 0
//     SharesOutstanding            = 0
//     IntToFloatQuantityDivisor    = 0
//     HasMarketDepthData           = 1
//     DisplayPriceMultiplier       = 1.000000
//     ExchangeSymbol               = ""
//     InitialMarginRequirement     = 0
//     MaintenanceMarginRequirement = 0
//     Currency                     = ""
//     ContractSize                 = 0
//     OpenInterest                 = 0
//     RolloverDate                 = 0
//     IsDelayed                    = 0
// }
func (m SecurityDefinitionResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SecurityDefinitionResponseFixedDefault)
}

// ToBuilder
func (m SecurityDefinitionResponseFixed) ToBuilder() SecurityDefinitionResponseFixedBuilder {
	return SecurityDefinitionResponseFixedBuilder{m.Fixed}
}

// Finish
func (m SecurityDefinitionResponseFixedBuilder) Finish() SecurityDefinitionResponseFixed {
	return SecurityDefinitionResponseFixed{m.Fixed}
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponseFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponseFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// Symbol This is the Symbol the Security Definition is for.
//
// When the Server responds with a single SecurityDefinitionResponse message
// and there are no security definitions to return for the original request,
// this will be empty.
//
// This field should be empty when this Security Definition message is in
// response to UnderlyingSymbolsForExchangeRequest.
func (m SecurityDefinitionResponseFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol This is the Symbol the Security Definition is for.
//
// When the Server responds with a single SecurityDefinitionResponse message
// and there are no security definitions to return for the original request,
// this will be empty.
//
// This field should be empty when this Security Definition message is in
// response to UnderlyingSymbolsForExchangeRequest.
func (m SecurityDefinitionResponseFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Exchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponseFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// Exchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponseFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// SecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseFixed) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 92, 88))
}

// SecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseFixedBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 92, 88))
}

// Description The text description for the Symbol.
func (m SecurityDefinitionResponseFixed) Description() string {
	return message.StringFixed(m.Unsafe(), 156, 92)
}

// Description The text description for the Symbol.
func (m SecurityDefinitionResponseFixedBuilder) Description() string {
	return message.StringFixed(m.Unsafe(), 156, 92)
}

// MinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseFixed) MinPriceIncrement() float32 {
	return message.Float32Fixed(m.Unsafe(), 160, 156)
}

// MinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseFixedBuilder) MinPriceIncrement() float32 {
	return message.Float32Fixed(m.Unsafe(), 160, 156)
}

// PriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseFixed) PriceDisplayFormat() PriceDisplayFormatEnum {
	return PriceDisplayFormatEnum(message.Int32Fixed(m.Unsafe(), 164, 160))
}

// PriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseFixedBuilder) PriceDisplayFormat() PriceDisplayFormatEnum {
	return PriceDisplayFormatEnum(message.Int32Fixed(m.Unsafe(), 164, 160))
}

// CurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseFixed) CurrencyValuePerIncrement() float32 {
	return message.Float32Fixed(m.Unsafe(), 168, 164)
}

// CurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseFixedBuilder) CurrencyValuePerIncrement() float32 {
	return message.Float32Fixed(m.Unsafe(), 168, 164)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseFixed) IsFinalMessage() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 169, 168)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseFixedBuilder) IsFinalMessage() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 169, 168)
}

// FloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponseFixed) FloatToIntPriceMultiplier() float32 {
	return message.Float32Fixed(m.Unsafe(), 176, 172)
}

// FloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponseFixedBuilder) FloatToIntPriceMultiplier() float32 {
	return message.Float32Fixed(m.Unsafe(), 176, 172)
}

// IntToFloatPriceDivisor
func (m SecurityDefinitionResponseFixed) IntToFloatPriceDivisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 180, 176)
}

// IntToFloatPriceDivisor
func (m SecurityDefinitionResponseFixedBuilder) IntToFloatPriceDivisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 180, 176)
}

// UnderlyingSymbol This is the underlying symbol for the Symbol field if the Symbol has an
// underlying symbol.
//
// The Server must set this when the SecurityDefinitionResponse message in
// response to UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest.
// response to UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest.
//
// This is typically used with futures. A futures symbol for a specific contract
// year and month will have an underlying symbol equivalent to the Symbol
// without the year and month characters.
func (m SecurityDefinitionResponseFixed) UnderlyingSymbol() string {
	return message.StringFixed(m.Unsafe(), 212, 180)
}

// UnderlyingSymbol This is the underlying symbol for the Symbol field if the Symbol has an
// underlying symbol.
//
// The Server must set this when the SecurityDefinitionResponse message in
// response to UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest.
// response to UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest.
//
// This is typically used with futures. A futures symbol for a specific contract
// year and month will have an underlying symbol equivalent to the Symbol
// without the year and month characters.
func (m SecurityDefinitionResponseFixedBuilder) UnderlyingSymbol() string {
	return message.StringFixed(m.Unsafe(), 212, 180)
}

// UpdatesBidAskOnly This is set to 1 when the Symbol does not provide MarketDataUpdateTrade
// messages and only provides MarketDataUpdateBidAsk messages when there
// is market activity for the Symbol.
//
// Otherwise, when this is set to 0, MarketDataUpdateTrade messages will
// be received after subscribing to market data, when there is trading activity.
// be received after subscribing to market data, when there is trading activity.
func (m SecurityDefinitionResponseFixed) UpdatesBidAskOnly() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 213, 212)
}

// UpdatesBidAskOnly This is set to 1 when the Symbol does not provide MarketDataUpdateTrade
// messages and only provides MarketDataUpdateBidAsk messages when there
// is market activity for the Symbol.
//
// Otherwise, when this is set to 0, MarketDataUpdateTrade messages will
// be received after subscribing to market data, when there is trading activity.
// be received after subscribing to market data, when there is trading activity.
func (m SecurityDefinitionResponseFixedBuilder) UpdatesBidAskOnly() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 213, 212)
}

// StrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponseFixed) StrikePrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 220, 216)
}

// StrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponseFixedBuilder) StrikePrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 220, 216)
}

// PutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseFixed) PutOrCall() PutCallEnum {
	return PutCallEnum(message.Uint8Fixed(m.Unsafe(), 221, 220))
}

// PutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseFixedBuilder) PutOrCall() PutCallEnum {
	return PutCallEnum(message.Uint8Fixed(m.Unsafe(), 221, 220))
}

// ShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseFixed) ShortInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 228, 224)
}

// ShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseFixedBuilder) ShortInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 228, 224)
}

// SecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseFixed) SecurityExpirationDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 232, 228))
}

// SecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseFixedBuilder) SecurityExpirationDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 232, 228))
}

// BuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixed) BuyRolloverInterest() float32 {
	return message.Float32Fixed(m.Unsafe(), 236, 232)
}

// BuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedBuilder) BuyRolloverInterest() float32 {
	return message.Float32Fixed(m.Unsafe(), 236, 232)
}

// SellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixed) SellRolloverInterest() float32 {
	return message.Float32Fixed(m.Unsafe(), 240, 236)
}

// SellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedBuilder) SellRolloverInterest() float32 {
	return message.Float32Fixed(m.Unsafe(), 240, 236)
}

// EarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseFixed) EarningsPerShare() float32 {
	return message.Float32Fixed(m.Unsafe(), 244, 240)
}

// EarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseFixedBuilder) EarningsPerShare() float32 {
	return message.Float32Fixed(m.Unsafe(), 244, 240)
}

// SharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseFixed) SharesOutstanding() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 248, 244)
}

// SharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseFixedBuilder) SharesOutstanding() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 248, 244)
}

// IntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseFixed) IntToFloatQuantityDivisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 252, 248)
}

// IntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseFixedBuilder) IntToFloatQuantityDivisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 252, 248)
}

// HasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponseFixed) HasMarketDepthData() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 253, 252)
}

// HasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponseFixedBuilder) HasMarketDepthData() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 253, 252)
}

// DisplayPriceMultiplier This is an optional field for the Server to set.
//
// The default for this is 1.0.
//
// This sets the multiplier to use in the case where the Client should multiply
// the values in market data messages by some number other than 1.0 before
// displaying them to the user.
//
// It is recommended that a Server does not use this and instead transmit
// to the Client values as the actual floating-point values.
//
// This should not be confused with the integer market data messages and
// the IntToFloatPriceDivisor field used with those messages. DisplayPriceMultiplier
// is for when the Server transmits market data values using floating-point
// types and where those values may have a fractional component, but where
// it is necessary to still multiply the original value by this multiplier
// to get the final value to display to the user.
//
// When this is set to a value other than 1.0, then the MinPriceIncrement
// and the PriceDisplayFormat fields are relative to the market data values
// after the DisplayPriceMultiplier is applied.
func (m SecurityDefinitionResponseFixed) DisplayPriceMultiplier() float32 {
	return message.Float32Fixed(m.Unsafe(), 260, 256)
}

// DisplayPriceMultiplier This is an optional field for the Server to set.
//
// The default for this is 1.0.
//
// This sets the multiplier to use in the case where the Client should multiply
// the values in market data messages by some number other than 1.0 before
// displaying them to the user.
//
// It is recommended that a Server does not use this and instead transmit
// to the Client values as the actual floating-point values.
//
// This should not be confused with the integer market data messages and
// the IntToFloatPriceDivisor field used with those messages. DisplayPriceMultiplier
// is for when the Server transmits market data values using floating-point
// types and where those values may have a fractional component, but where
// it is necessary to still multiply the original value by this multiplier
// to get the final value to display to the user.
//
// When this is set to a value other than 1.0, then the MinPriceIncrement
// and the PriceDisplayFormat fields are relative to the market data values
// after the DisplayPriceMultiplier is applied.
func (m SecurityDefinitionResponseFixedBuilder) DisplayPriceMultiplier() float32 {
	return message.Float32Fixed(m.Unsafe(), 260, 256)
}

// ExchangeSymbol This is an optional field. This is the exchange symbol which corresponds
// with the Symbol field.
//
// This field should be empty when this Security Definition message is in
// response to UNDERLYING_SymbolsForExchangeRequest.
func (m SecurityDefinitionResponseFixed) ExchangeSymbol() string {
	return message.StringFixed(m.Unsafe(), 324, 260)
}

// ExchangeSymbol This is an optional field. This is the exchange symbol which corresponds
// with the Symbol field.
//
// This field should be empty when this Security Definition message is in
// response to UNDERLYING_SymbolsForExchangeRequest.
func (m SecurityDefinitionResponseFixedBuilder) ExchangeSymbol() string {
	return message.StringFixed(m.Unsafe(), 324, 260)
}

// InitialMarginRequirement This field applies to the Futures Security Type.
//
// This is the initial margin requirement as specified by the exchange, if
// available.
func (m SecurityDefinitionResponseFixed) InitialMarginRequirement() float32 {
	return message.Float32Fixed(m.Unsafe(), 328, 324)
}

// InitialMarginRequirement This field applies to the Futures Security Type.
//
// This is the initial margin requirement as specified by the exchange, if
// available.
func (m SecurityDefinitionResponseFixedBuilder) InitialMarginRequirement() float32 {
	return message.Float32Fixed(m.Unsafe(), 328, 324)
}

// MaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponseFixed) MaintenanceMarginRequirement() float32 {
	return message.Float32Fixed(m.Unsafe(), 332, 328)
}

// MaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponseFixedBuilder) MaintenanceMarginRequirement() float32 {
	return message.Float32Fixed(m.Unsafe(), 332, 328)
}

// Currency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponseFixed) Currency() string {
	return message.StringFixed(m.Unsafe(), 340, 332)
}

// Currency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponseFixedBuilder) Currency() string {
	return message.StringFixed(m.Unsafe(), 340, 332)
}

// ContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponseFixed) ContractSize() float32 {
	return message.Float32Fixed(m.Unsafe(), 344, 340)
}

// ContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponseFixedBuilder) ContractSize() float32 {
	return message.Float32Fixed(m.Unsafe(), 344, 340)
}

// OpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponseFixed) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 348, 344)
}

// OpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponseFixedBuilder) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 348, 344)
}

// RolloverDate This field applies to the Futures Security Type.
//
// This is the rollover date for the symbol according to the typical time
// where trading transitions from one contract month to the next.
func (m SecurityDefinitionResponseFixed) RolloverDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 352, 348))
}

// RolloverDate This field applies to the Futures Security Type.
//
// This is the rollover date for the symbol according to the typical time
// where trading transitions from one contract month to the next.
func (m SecurityDefinitionResponseFixedBuilder) RolloverDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 352, 348))
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseFixed) IsDelayed() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 353, 352)
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseFixedBuilder) IsDelayed() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 353, 352)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponseFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbol This is the Symbol the Security Definition is for.
//
// When the Server responds with a single SecurityDefinitionResponse message
// and there are no security definitions to return for the original request,
// this will be empty.
//
// This field should be empty when this Security Definition message is in
// response to UnderlyingSymbolsForExchangeRequest.
func (m SecurityDefinitionResponseFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetExchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponseFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 88, 72, value)
}

// SetSecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseFixedBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 92, 88, int32(value))
}

// SetDescription The text description for the Symbol.
func (m SecurityDefinitionResponseFixedBuilder) SetDescription(value string) {
	message.SetStringFixed(m.Unsafe(), 156, 92, value)
}

// SetMinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseFixedBuilder) SetMinPriceIncrement(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 160, 156, value)
}

// SetPriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseFixedBuilder) SetPriceDisplayFormat(value PriceDisplayFormatEnum) {
	message.SetInt32Fixed(m.Unsafe(), 164, 160, int32(value))
}

// SetCurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseFixedBuilder) SetCurrencyValuePerIncrement(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 168, 164, value)
}

// SetIsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseFixedBuilder) SetIsFinalMessage(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 169, 168, value)
}

// SetFloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponseFixedBuilder) SetFloatToIntPriceMultiplier(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 176, 172, value)
}

// SetIntToFloatPriceDivisor
func (m SecurityDefinitionResponseFixedBuilder) SetIntToFloatPriceDivisor(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 180, 176, value)
}

// SetUnderlyingSymbol This is the underlying symbol for the Symbol field if the Symbol has an
// underlying symbol.
//
// The Server must set this when the SecurityDefinitionResponse message in
// response to UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest.
// response to UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest.
//
// This is typically used with futures. A futures symbol for a specific contract
// year and month will have an underlying symbol equivalent to the Symbol
// without the year and month characters.
func (m SecurityDefinitionResponseFixedBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 212, 180, value)
}

// SetUpdatesBidAskOnly This is set to 1 when the Symbol does not provide MarketDataUpdateTrade
// messages and only provides MarketDataUpdateBidAsk messages when there
// is market activity for the Symbol.
//
// Otherwise, when this is set to 0, MarketDataUpdateTrade messages will
// be received after subscribing to market data, when there is trading activity.
// be received after subscribing to market data, when there is trading activity.
func (m SecurityDefinitionResponseFixedBuilder) SetUpdatesBidAskOnly(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 213, 212, value)
}

// SetStrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponseFixedBuilder) SetStrikePrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 220, 216, value)
}

// SetPutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseFixedBuilder) SetPutOrCall(value PutCallEnum) {
	message.SetUint8Fixed(m.Unsafe(), 221, 220, uint8(value))
}

// SetShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseFixedBuilder) SetShortInterest(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 228, 224, value)
}

// SetSecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseFixedBuilder) SetSecurityExpirationDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 232, 228, uint32(value))
}

// SetBuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedBuilder) SetBuyRolloverInterest(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 236, 232, value)
}

// SetSellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedBuilder) SetSellRolloverInterest(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 240, 236, value)
}

// SetEarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseFixedBuilder) SetEarningsPerShare(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 244, 240, value)
}

// SetSharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseFixedBuilder) SetSharesOutstanding(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 248, 244, value)
}

// SetIntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseFixedBuilder) SetIntToFloatQuantityDivisor(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 252, 248, value)
}

// SetHasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponseFixedBuilder) SetHasMarketDepthData(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 253, 252, value)
}

// SetDisplayPriceMultiplier This is an optional field for the Server to set.
//
// The default for this is 1.0.
//
// This sets the multiplier to use in the case where the Client should multiply
// the values in market data messages by some number other than 1.0 before
// displaying them to the user.
//
// It is recommended that a Server does not use this and instead transmit
// to the Client values as the actual floating-point values.
//
// This should not be confused with the integer market data messages and
// the IntToFloatPriceDivisor field used with those messages. DisplayPriceMultiplier
// is for when the Server transmits market data values using floating-point
// types and where those values may have a fractional component, but where
// it is necessary to still multiply the original value by this multiplier
// to get the final value to display to the user.
//
// When this is set to a value other than 1.0, then the MinPriceIncrement
// and the PriceDisplayFormat fields are relative to the market data values
// after the DisplayPriceMultiplier is applied.
func (m SecurityDefinitionResponseFixedBuilder) SetDisplayPriceMultiplier(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 260, 256, value)
}

// SetExchangeSymbol This is an optional field. This is the exchange symbol which corresponds
// with the Symbol field.
//
// This field should be empty when this Security Definition message is in
// response to UNDERLYING_SymbolsForExchangeRequest.
func (m SecurityDefinitionResponseFixedBuilder) SetExchangeSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 324, 260, value)
}

// SetInitialMarginRequirement This field applies to the Futures Security Type.
//
// This is the initial margin requirement as specified by the exchange, if
// available.
func (m SecurityDefinitionResponseFixedBuilder) SetInitialMarginRequirement(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 328, 324, value)
}

// SetMaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponseFixedBuilder) SetMaintenanceMarginRequirement(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 332, 328, value)
}

// SetCurrency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponseFixedBuilder) SetCurrency(value string) {
	message.SetStringFixed(m.Unsafe(), 340, 332, value)
}

// SetContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponseFixedBuilder) SetContractSize(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 344, 340, value)
}

// SetOpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponseFixedBuilder) SetOpenInterest(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 348, 344, value)
}

// SetRolloverDate This field applies to the Futures Security Type.
//
// This is the rollover date for the symbol according to the typical time
// where trading transitions from one contract month to the next.
func (m SecurityDefinitionResponseFixedBuilder) SetRolloverDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 352, 348, uint32(value))
}

// SetIsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseFixedBuilder) SetIsDelayed(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 353, 352, value)
}

// Copy
func (m SecurityDefinitionResponseFixed) Copy(to SecurityDefinitionResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetDescription(m.Description())
	to.SetMinPriceIncrement(m.MinPriceIncrement())
	to.SetPriceDisplayFormat(m.PriceDisplayFormat())
	to.SetCurrencyValuePerIncrement(m.CurrencyValuePerIncrement())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetFloatToIntPriceMultiplier(m.FloatToIntPriceMultiplier())
	to.SetIntToFloatPriceDivisor(m.IntToFloatPriceDivisor())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetUpdatesBidAskOnly(m.UpdatesBidAskOnly())
	to.SetStrikePrice(m.StrikePrice())
	to.SetPutOrCall(m.PutOrCall())
	to.SetShortInterest(m.ShortInterest())
	to.SetSecurityExpirationDate(m.SecurityExpirationDate())
	to.SetBuyRolloverInterest(m.BuyRolloverInterest())
	to.SetSellRolloverInterest(m.SellRolloverInterest())
	to.SetEarningsPerShare(m.EarningsPerShare())
	to.SetSharesOutstanding(m.SharesOutstanding())
	to.SetIntToFloatQuantityDivisor(m.IntToFloatQuantityDivisor())
	to.SetHasMarketDepthData(m.HasMarketDepthData())
	to.SetDisplayPriceMultiplier(m.DisplayPriceMultiplier())
	to.SetExchangeSymbol(m.ExchangeSymbol())
	to.SetInitialMarginRequirement(m.InitialMarginRequirement())
	to.SetMaintenanceMarginRequirement(m.MaintenanceMarginRequirement())
	to.SetCurrency(m.Currency())
	to.SetContractSize(m.ContractSize())
	to.SetOpenInterest(m.OpenInterest())
	to.SetRolloverDate(m.RolloverDate())
	to.SetIsDelayed(m.IsDelayed())
}

// Copy
func (m SecurityDefinitionResponseFixedBuilder) Copy(to SecurityDefinitionResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetDescription(m.Description())
	to.SetMinPriceIncrement(m.MinPriceIncrement())
	to.SetPriceDisplayFormat(m.PriceDisplayFormat())
	to.SetCurrencyValuePerIncrement(m.CurrencyValuePerIncrement())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetFloatToIntPriceMultiplier(m.FloatToIntPriceMultiplier())
	to.SetIntToFloatPriceDivisor(m.IntToFloatPriceDivisor())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetUpdatesBidAskOnly(m.UpdatesBidAskOnly())
	to.SetStrikePrice(m.StrikePrice())
	to.SetPutOrCall(m.PutOrCall())
	to.SetShortInterest(m.ShortInterest())
	to.SetSecurityExpirationDate(m.SecurityExpirationDate())
	to.SetBuyRolloverInterest(m.BuyRolloverInterest())
	to.SetSellRolloverInterest(m.SellRolloverInterest())
	to.SetEarningsPerShare(m.EarningsPerShare())
	to.SetSharesOutstanding(m.SharesOutstanding())
	to.SetIntToFloatQuantityDivisor(m.IntToFloatQuantityDivisor())
	to.SetHasMarketDepthData(m.HasMarketDepthData())
	to.SetDisplayPriceMultiplier(m.DisplayPriceMultiplier())
	to.SetExchangeSymbol(m.ExchangeSymbol())
	to.SetInitialMarginRequirement(m.InitialMarginRequirement())
	to.SetMaintenanceMarginRequirement(m.MaintenanceMarginRequirement())
	to.SetCurrency(m.Currency())
	to.SetContractSize(m.ContractSize())
	to.SetOpenInterest(m.OpenInterest())
	to.SetRolloverDate(m.RolloverDate())
	to.SetIsDelayed(m.IsDelayed())
}

// CopyPointer
func (m SecurityDefinitionResponseFixed) CopyPointer(to SecurityDefinitionResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetDescription(m.Description())
	to.SetMinPriceIncrement(m.MinPriceIncrement())
	to.SetPriceDisplayFormat(m.PriceDisplayFormat())
	to.SetCurrencyValuePerIncrement(m.CurrencyValuePerIncrement())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetFloatToIntPriceMultiplier(m.FloatToIntPriceMultiplier())
	to.SetIntToFloatPriceDivisor(m.IntToFloatPriceDivisor())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetUpdatesBidAskOnly(m.UpdatesBidAskOnly())
	to.SetStrikePrice(m.StrikePrice())
	to.SetPutOrCall(m.PutOrCall())
	to.SetShortInterest(m.ShortInterest())
	to.SetSecurityExpirationDate(m.SecurityExpirationDate())
	to.SetBuyRolloverInterest(m.BuyRolloverInterest())
	to.SetSellRolloverInterest(m.SellRolloverInterest())
	to.SetEarningsPerShare(m.EarningsPerShare())
	to.SetSharesOutstanding(m.SharesOutstanding())
	to.SetIntToFloatQuantityDivisor(m.IntToFloatQuantityDivisor())
	to.SetHasMarketDepthData(m.HasMarketDepthData())
	to.SetDisplayPriceMultiplier(m.DisplayPriceMultiplier())
	to.SetExchangeSymbol(m.ExchangeSymbol())
	to.SetInitialMarginRequirement(m.InitialMarginRequirement())
	to.SetMaintenanceMarginRequirement(m.MaintenanceMarginRequirement())
	to.SetCurrency(m.Currency())
	to.SetContractSize(m.ContractSize())
	to.SetOpenInterest(m.OpenInterest())
	to.SetRolloverDate(m.RolloverDate())
	to.SetIsDelayed(m.IsDelayed())
}

// CopyPointer
func (m SecurityDefinitionResponseFixedBuilder) CopyPointer(to SecurityDefinitionResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetDescription(m.Description())
	to.SetMinPriceIncrement(m.MinPriceIncrement())
	to.SetPriceDisplayFormat(m.PriceDisplayFormat())
	to.SetCurrencyValuePerIncrement(m.CurrencyValuePerIncrement())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetFloatToIntPriceMultiplier(m.FloatToIntPriceMultiplier())
	to.SetIntToFloatPriceDivisor(m.IntToFloatPriceDivisor())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetUpdatesBidAskOnly(m.UpdatesBidAskOnly())
	to.SetStrikePrice(m.StrikePrice())
	to.SetPutOrCall(m.PutOrCall())
	to.SetShortInterest(m.ShortInterest())
	to.SetSecurityExpirationDate(m.SecurityExpirationDate())
	to.SetBuyRolloverInterest(m.BuyRolloverInterest())
	to.SetSellRolloverInterest(m.SellRolloverInterest())
	to.SetEarningsPerShare(m.EarningsPerShare())
	to.SetSharesOutstanding(m.SharesOutstanding())
	to.SetIntToFloatQuantityDivisor(m.IntToFloatQuantityDivisor())
	to.SetHasMarketDepthData(m.HasMarketDepthData())
	to.SetDisplayPriceMultiplier(m.DisplayPriceMultiplier())
	to.SetExchangeSymbol(m.ExchangeSymbol())
	to.SetInitialMarginRequirement(m.InitialMarginRequirement())
	to.SetMaintenanceMarginRequirement(m.MaintenanceMarginRequirement())
	to.SetCurrency(m.Currency())
	to.SetContractSize(m.ContractSize())
	to.SetOpenInterest(m.OpenInterest())
	to.SetRolloverDate(m.RolloverDate())
	to.SetIsDelayed(m.IsDelayed())
}

// CopyToPointer
func (m SecurityDefinitionResponseFixed) CopyToPointer(to SecurityDefinitionResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetDescription(m.Description())
	to.SetMinPriceIncrement(m.MinPriceIncrement())
	to.SetPriceDisplayFormat(m.PriceDisplayFormat())
	to.SetCurrencyValuePerIncrement(m.CurrencyValuePerIncrement())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetFloatToIntPriceMultiplier(m.FloatToIntPriceMultiplier())
	to.SetIntToFloatPriceDivisor(m.IntToFloatPriceDivisor())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetUpdatesBidAskOnly(m.UpdatesBidAskOnly())
	to.SetStrikePrice(m.StrikePrice())
	to.SetPutOrCall(m.PutOrCall())
	to.SetShortInterest(m.ShortInterest())
	to.SetSecurityExpirationDate(m.SecurityExpirationDate())
	to.SetBuyRolloverInterest(m.BuyRolloverInterest())
	to.SetSellRolloverInterest(m.SellRolloverInterest())
	to.SetEarningsPerShare(m.EarningsPerShare())
	to.SetSharesOutstanding(m.SharesOutstanding())
	to.SetIntToFloatQuantityDivisor(m.IntToFloatQuantityDivisor())
	to.SetHasMarketDepthData(m.HasMarketDepthData())
	to.SetDisplayPriceMultiplier(m.DisplayPriceMultiplier())
	to.SetExchangeSymbol(m.ExchangeSymbol())
	to.SetInitialMarginRequirement(m.InitialMarginRequirement())
	to.SetMaintenanceMarginRequirement(m.MaintenanceMarginRequirement())
	to.SetCurrency(m.Currency())
	to.SetContractSize(m.ContractSize())
	to.SetOpenInterest(m.OpenInterest())
	to.SetRolloverDate(m.RolloverDate())
	to.SetIsDelayed(m.IsDelayed())
}

// CopyToPointer
func (m SecurityDefinitionResponseFixedBuilder) CopyToPointer(to SecurityDefinitionResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetDescription(m.Description())
	to.SetMinPriceIncrement(m.MinPriceIncrement())
	to.SetPriceDisplayFormat(m.PriceDisplayFormat())
	to.SetCurrencyValuePerIncrement(m.CurrencyValuePerIncrement())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetFloatToIntPriceMultiplier(m.FloatToIntPriceMultiplier())
	to.SetIntToFloatPriceDivisor(m.IntToFloatPriceDivisor())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetUpdatesBidAskOnly(m.UpdatesBidAskOnly())
	to.SetStrikePrice(m.StrikePrice())
	to.SetPutOrCall(m.PutOrCall())
	to.SetShortInterest(m.ShortInterest())
	to.SetSecurityExpirationDate(m.SecurityExpirationDate())
	to.SetBuyRolloverInterest(m.BuyRolloverInterest())
	to.SetSellRolloverInterest(m.SellRolloverInterest())
	to.SetEarningsPerShare(m.EarningsPerShare())
	to.SetSharesOutstanding(m.SharesOutstanding())
	to.SetIntToFloatQuantityDivisor(m.IntToFloatQuantityDivisor())
	to.SetHasMarketDepthData(m.HasMarketDepthData())
	to.SetDisplayPriceMultiplier(m.DisplayPriceMultiplier())
	to.SetExchangeSymbol(m.ExchangeSymbol())
	to.SetInitialMarginRequirement(m.InitialMarginRequirement())
	to.SetMaintenanceMarginRequirement(m.MaintenanceMarginRequirement())
	to.SetCurrency(m.Currency())
	to.SetContractSize(m.ContractSize())
	to.SetOpenInterest(m.OpenInterest())
	to.SetRolloverDate(m.RolloverDate())
	to.SetIsDelayed(m.IsDelayed())
}

// CopyTo
func (m SecurityDefinitionResponseFixed) CopyTo(to SecurityDefinitionResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetDescription(m.Description())
	to.SetMinPriceIncrement(m.MinPriceIncrement())
	to.SetPriceDisplayFormat(m.PriceDisplayFormat())
	to.SetCurrencyValuePerIncrement(m.CurrencyValuePerIncrement())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetFloatToIntPriceMultiplier(m.FloatToIntPriceMultiplier())
	to.SetIntToFloatPriceDivisor(m.IntToFloatPriceDivisor())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetUpdatesBidAskOnly(m.UpdatesBidAskOnly())
	to.SetStrikePrice(m.StrikePrice())
	to.SetPutOrCall(m.PutOrCall())
	to.SetShortInterest(m.ShortInterest())
	to.SetSecurityExpirationDate(m.SecurityExpirationDate())
	to.SetBuyRolloverInterest(m.BuyRolloverInterest())
	to.SetSellRolloverInterest(m.SellRolloverInterest())
	to.SetEarningsPerShare(m.EarningsPerShare())
	to.SetSharesOutstanding(m.SharesOutstanding())
	to.SetIntToFloatQuantityDivisor(m.IntToFloatQuantityDivisor())
	to.SetHasMarketDepthData(m.HasMarketDepthData())
	to.SetDisplayPriceMultiplier(m.DisplayPriceMultiplier())
	to.SetExchangeSymbol(m.ExchangeSymbol())
	to.SetInitialMarginRequirement(m.InitialMarginRequirement())
	to.SetMaintenanceMarginRequirement(m.MaintenanceMarginRequirement())
	to.SetCurrency(m.Currency())
	to.SetContractSize(m.ContractSize())
	to.SetOpenInterest(m.OpenInterest())
	to.SetRolloverDate(m.RolloverDate())
	to.SetIsDelayed(m.IsDelayed())
}

// CopyTo
func (m SecurityDefinitionResponseFixedBuilder) CopyTo(to SecurityDefinitionResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetDescription(m.Description())
	to.SetMinPriceIncrement(m.MinPriceIncrement())
	to.SetPriceDisplayFormat(m.PriceDisplayFormat())
	to.SetCurrencyValuePerIncrement(m.CurrencyValuePerIncrement())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetFloatToIntPriceMultiplier(m.FloatToIntPriceMultiplier())
	to.SetIntToFloatPriceDivisor(m.IntToFloatPriceDivisor())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetUpdatesBidAskOnly(m.UpdatesBidAskOnly())
	to.SetStrikePrice(m.StrikePrice())
	to.SetPutOrCall(m.PutOrCall())
	to.SetShortInterest(m.ShortInterest())
	to.SetSecurityExpirationDate(m.SecurityExpirationDate())
	to.SetBuyRolloverInterest(m.BuyRolloverInterest())
	to.SetSellRolloverInterest(m.SellRolloverInterest())
	to.SetEarningsPerShare(m.EarningsPerShare())
	to.SetSharesOutstanding(m.SharesOutstanding())
	to.SetIntToFloatQuantityDivisor(m.IntToFloatQuantityDivisor())
	to.SetHasMarketDepthData(m.HasMarketDepthData())
	to.SetDisplayPriceMultiplier(m.DisplayPriceMultiplier())
	to.SetExchangeSymbol(m.ExchangeSymbol())
	to.SetInitialMarginRequirement(m.InitialMarginRequirement())
	to.SetMaintenanceMarginRequirement(m.MaintenanceMarginRequirement())
	to.SetCurrency(m.Currency())
	to.SetContractSize(m.ContractSize())
	to.SetOpenInterest(m.OpenInterest())
	to.SetRolloverDate(m.RolloverDate())
	to.SetIsDelayed(m.IsDelayed())
}
