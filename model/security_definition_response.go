package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                         = SecurityDefinitionResponseSize  (136)
//     Type                         = SECURITY_DEFINITION_RESPONSE  (507)
//     BaseSize                     = SecurityDefinitionResponseSize  (136)
//     RequestID                    = 0
//     Symbol                       = ""
//     Exchange                     = ""
//     SecurityType                 = 0
//     Description                  = ""
//     MinPriceIncrement            = 0
//     PriceDisplayFormat           = 0
//     CurrencyValuePerIncrement    = 0
//     IsFinalMessage               = 0
//     FloatToIntPriceMultiplier    = 1.000000
//     IntToFloatPriceDivisor       = 1.000000
//     UnderlyingSymbol             = ""
//     UpdatesBidAskOnly            = 0
//     StrikePrice                  = 0
//     PutOrCall                    = 0
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
var _SecurityDefinitionResponseDefault = []byte{136, 0, 251, 1, 136, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SecurityDefinitionResponseSize = 136

// SecurityDefinitionResponse This is a response from the Server in response to a SymbolsForExchangeRequest,
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
type SecurityDefinitionResponse struct {
	message.VLS
}

// SecurityDefinitionResponseBuilder This is a response from the Server in response to a SymbolsForExchangeRequest,
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
type SecurityDefinitionResponseBuilder struct {
	message.VLSBuilder
}

func NewSecurityDefinitionResponseFrom(b []byte) SecurityDefinitionResponse {
	return SecurityDefinitionResponse{VLS: message.NewVLSFrom(b)}
}

func WrapSecurityDefinitionResponse(b []byte) SecurityDefinitionResponse {
	return SecurityDefinitionResponse{VLS: message.WrapVLS(b)}
}

func NewSecurityDefinitionResponse() SecurityDefinitionResponseBuilder {
	a := SecurityDefinitionResponseBuilder{message.NewVLSBuilder(136)}
	a.Unsafe().SetBytes(0, _SecurityDefinitionResponseDefault)
	return a
}

// Clear
// {
//     Size                         = SecurityDefinitionResponseSize  (136)
//     Type                         = SECURITY_DEFINITION_RESPONSE  (507)
//     BaseSize                     = SecurityDefinitionResponseSize  (136)
//     RequestID                    = 0
//     Symbol                       = ""
//     Exchange                     = ""
//     SecurityType                 = 0
//     Description                  = ""
//     MinPriceIncrement            = 0
//     PriceDisplayFormat           = 0
//     CurrencyValuePerIncrement    = 0
//     IsFinalMessage               = 0
//     FloatToIntPriceMultiplier    = 1.000000
//     IntToFloatPriceDivisor       = 1.000000
//     UnderlyingSymbol             = ""
//     UpdatesBidAskOnly            = 0
//     StrikePrice                  = 0
//     PutOrCall                    = 0
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
func (m SecurityDefinitionResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SecurityDefinitionResponseDefault)
}

// ToBuilder
func (m SecurityDefinitionResponse) ToBuilder() SecurityDefinitionResponseBuilder {
	return SecurityDefinitionResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m SecurityDefinitionResponseBuilder) Finish() SecurityDefinitionResponse {
	return SecurityDefinitionResponse{m.VLSBuilder.Finish()}
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponse) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponseBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// Symbol This is the Symbol the Security Definition is for.
//
// When the Server responds with a single SecurityDefinitionResponse message
// and there are no security definitions to return for the original request,
// this will be empty.
//
// This field should be empty when this Security Definition message is in
// response to UnderlyingSymbolsForExchangeRequest.
func (m SecurityDefinitionResponse) Symbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Symbol This is the Symbol the Security Definition is for.
//
// When the Server responds with a single SecurityDefinitionResponse message
// and there are no security definitions to return for the original request,
// this will be empty.
//
// This field should be empty when this Security Definition message is in
// response to UnderlyingSymbolsForExchangeRequest.
func (m SecurityDefinitionResponseBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Exchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponse) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponseBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// SecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponse) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// SecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// Description The text description for the Symbol.
func (m SecurityDefinitionResponse) Description() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Description The text description for the Symbol.
func (m SecurityDefinitionResponseBuilder) Description() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// MinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponse) MinPriceIncrement() float32 {
	return message.Float32VLS(m.Unsafe(), 32, 28)
}

// MinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseBuilder) MinPriceIncrement() float32 {
	return message.Float32VLS(m.Unsafe(), 32, 28)
}

// PriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponse) PriceDisplayFormat() PriceDisplayFormatEnum {
	return PriceDisplayFormatEnum(message.Int32VLS(m.Unsafe(), 36, 32))
}

// PriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseBuilder) PriceDisplayFormat() PriceDisplayFormatEnum {
	return PriceDisplayFormatEnum(message.Int32VLS(m.Unsafe(), 36, 32))
}

// CurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponse) CurrencyValuePerIncrement() float32 {
	return message.Float32VLS(m.Unsafe(), 40, 36)
}

// CurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseBuilder) CurrencyValuePerIncrement() float32 {
	return message.Float32VLS(m.Unsafe(), 40, 36)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponse) IsFinalMessage() uint8 {
	return message.Uint8VLS(m.Unsafe(), 41, 40)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseBuilder) IsFinalMessage() uint8 {
	return message.Uint8VLS(m.Unsafe(), 41, 40)
}

// FloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponse) FloatToIntPriceMultiplier() float32 {
	return message.Float32VLS(m.Unsafe(), 48, 44)
}

// FloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponseBuilder) FloatToIntPriceMultiplier() float32 {
	return message.Float32VLS(m.Unsafe(), 48, 44)
}

// IntToFloatPriceDivisor
func (m SecurityDefinitionResponse) IntToFloatPriceDivisor() float32 {
	return message.Float32VLS(m.Unsafe(), 52, 48)
}

// IntToFloatPriceDivisor
func (m SecurityDefinitionResponseBuilder) IntToFloatPriceDivisor() float32 {
	return message.Float32VLS(m.Unsafe(), 52, 48)
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
func (m SecurityDefinitionResponse) UnderlyingSymbol() string {
	return message.StringVLS(m.Unsafe(), 56, 52)
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
func (m SecurityDefinitionResponseBuilder) UnderlyingSymbol() string {
	return message.StringVLS(m.Unsafe(), 56, 52)
}

// UpdatesBidAskOnly This is set to 1 when the Symbol does not provide MarketDataUpdateTrade
// messages and only provides MarketDataUpdateBidAsk messages when there
// is market activity for the Symbol.
//
// Otherwise, when this is set to 0, MarketDataUpdateTrade messages will
// be received after subscribing to market data, when there is trading activity.
// be received after subscribing to market data, when there is trading activity.
func (m SecurityDefinitionResponse) UpdatesBidAskOnly() uint8 {
	return message.Uint8VLS(m.Unsafe(), 57, 56)
}

// UpdatesBidAskOnly This is set to 1 when the Symbol does not provide MarketDataUpdateTrade
// messages and only provides MarketDataUpdateBidAsk messages when there
// is market activity for the Symbol.
//
// Otherwise, when this is set to 0, MarketDataUpdateTrade messages will
// be received after subscribing to market data, when there is trading activity.
// be received after subscribing to market data, when there is trading activity.
func (m SecurityDefinitionResponseBuilder) UpdatesBidAskOnly() uint8 {
	return message.Uint8VLS(m.Unsafe(), 57, 56)
}

// StrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponse) StrikePrice() float32 {
	return message.Float32VLS(m.Unsafe(), 64, 60)
}

// StrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponseBuilder) StrikePrice() float32 {
	return message.Float32VLS(m.Unsafe(), 64, 60)
}

// PutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponse) PutOrCall() PutCallEnum {
	return PutCallEnum(message.Uint8VLS(m.Unsafe(), 65, 64))
}

// PutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseBuilder) PutOrCall() PutCallEnum {
	return PutCallEnum(message.Uint8VLS(m.Unsafe(), 65, 64))
}

// ShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponse) ShortInterest() uint32 {
	return message.Uint32VLS(m.Unsafe(), 72, 68)
}

// ShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseBuilder) ShortInterest() uint32 {
	return message.Uint32VLS(m.Unsafe(), 72, 68)
}

// SecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponse) SecurityExpirationDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Unsafe(), 76, 72))
}

// SecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseBuilder) SecurityExpirationDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Unsafe(), 76, 72))
}

// BuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponse) BuyRolloverInterest() float32 {
	return message.Float32VLS(m.Unsafe(), 80, 76)
}

// BuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseBuilder) BuyRolloverInterest() float32 {
	return message.Float32VLS(m.Unsafe(), 80, 76)
}

// SellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponse) SellRolloverInterest() float32 {
	return message.Float32VLS(m.Unsafe(), 84, 80)
}

// SellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseBuilder) SellRolloverInterest() float32 {
	return message.Float32VLS(m.Unsafe(), 84, 80)
}

// EarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponse) EarningsPerShare() float32 {
	return message.Float32VLS(m.Unsafe(), 88, 84)
}

// EarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseBuilder) EarningsPerShare() float32 {
	return message.Float32VLS(m.Unsafe(), 88, 84)
}

// SharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponse) SharesOutstanding() uint32 {
	return message.Uint32VLS(m.Unsafe(), 92, 88)
}

// SharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseBuilder) SharesOutstanding() uint32 {
	return message.Uint32VLS(m.Unsafe(), 92, 88)
}

// IntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponse) IntToFloatQuantityDivisor() float32 {
	return message.Float32VLS(m.Unsafe(), 96, 92)
}

// IntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseBuilder) IntToFloatQuantityDivisor() float32 {
	return message.Float32VLS(m.Unsafe(), 96, 92)
}

// HasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponse) HasMarketDepthData() uint8 {
	return message.Uint8VLS(m.Unsafe(), 97, 96)
}

// HasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponseBuilder) HasMarketDepthData() uint8 {
	return message.Uint8VLS(m.Unsafe(), 97, 96)
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
func (m SecurityDefinitionResponse) DisplayPriceMultiplier() float32 {
	return message.Float32VLS(m.Unsafe(), 104, 100)
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
func (m SecurityDefinitionResponseBuilder) DisplayPriceMultiplier() float32 {
	return message.Float32VLS(m.Unsafe(), 104, 100)
}

// ExchangeSymbol This is an optional field. This is the exchange symbol which corresponds
// with the Symbol field.
//
// This field should be empty when this Security Definition message is in
// response to UNDERLYING_SymbolsForExchangeRequest.
func (m SecurityDefinitionResponse) ExchangeSymbol() string {
	return message.StringVLS(m.Unsafe(), 108, 104)
}

// ExchangeSymbol This is an optional field. This is the exchange symbol which corresponds
// with the Symbol field.
//
// This field should be empty when this Security Definition message is in
// response to UNDERLYING_SymbolsForExchangeRequest.
func (m SecurityDefinitionResponseBuilder) ExchangeSymbol() string {
	return message.StringVLS(m.Unsafe(), 108, 104)
}

// InitialMarginRequirement This field applies to the Futures Security Type.
//
// This is the initial margin requirement as specified by the exchange, if
// available.
func (m SecurityDefinitionResponse) InitialMarginRequirement() float32 {
	return message.Float32VLS(m.Unsafe(), 112, 108)
}

// InitialMarginRequirement This field applies to the Futures Security Type.
//
// This is the initial margin requirement as specified by the exchange, if
// available.
func (m SecurityDefinitionResponseBuilder) InitialMarginRequirement() float32 {
	return message.Float32VLS(m.Unsafe(), 112, 108)
}

// MaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponse) MaintenanceMarginRequirement() float32 {
	return message.Float32VLS(m.Unsafe(), 116, 112)
}

// MaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponseBuilder) MaintenanceMarginRequirement() float32 {
	return message.Float32VLS(m.Unsafe(), 116, 112)
}

// Currency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponse) Currency() string {
	return message.StringVLS(m.Unsafe(), 120, 116)
}

// Currency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponseBuilder) Currency() string {
	return message.StringVLS(m.Unsafe(), 120, 116)
}

// ContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponse) ContractSize() float32 {
	return message.Float32VLS(m.Unsafe(), 124, 120)
}

// ContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponseBuilder) ContractSize() float32 {
	return message.Float32VLS(m.Unsafe(), 124, 120)
}

// OpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponse) OpenInterest() uint32 {
	return message.Uint32VLS(m.Unsafe(), 128, 124)
}

// OpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponseBuilder) OpenInterest() uint32 {
	return message.Uint32VLS(m.Unsafe(), 128, 124)
}

// RolloverDate This field applies to the Futures Security Type.
//
// This is the rollover date for the symbol according to the typical time
// where trading transitions from one contract month to the next.
func (m SecurityDefinitionResponse) RolloverDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Unsafe(), 132, 128))
}

// RolloverDate This field applies to the Futures Security Type.
//
// This is the rollover date for the symbol according to the typical time
// where trading transitions from one contract month to the next.
func (m SecurityDefinitionResponseBuilder) RolloverDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Unsafe(), 132, 128))
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponse) IsDelayed() uint8 {
	return message.Uint8VLS(m.Unsafe(), 133, 132)
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseBuilder) IsDelayed() uint8 {
	return message.Uint8VLS(m.Unsafe(), 133, 132)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetSymbol This is the Symbol the Security Definition is for.
//
// When the Server responds with a single SecurityDefinitionResponse message
// and there are no security definitions to return for the original request,
// this will be empty.
//
// This field should be empty when this Security Definition message is in
// response to UnderlyingSymbolsForExchangeRequest.
func (m *SecurityDefinitionResponseBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetExchange This is the Exchange for the Symbol. This field is optional.
func (m *SecurityDefinitionResponseBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetSecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, int32(value))
}

// SetDescription The text description for the Symbol.
func (m *SecurityDefinitionResponseBuilder) SetDescription(value string) {
	message.SetStringVLS(&m.VLSBuilder, 28, 24, value)
}

// SetMinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseBuilder) SetMinPriceIncrement(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 32, 28, value)
}

// SetPriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseBuilder) SetPriceDisplayFormat(value PriceDisplayFormatEnum) {
	message.SetInt32VLS(m.Unsafe(), 36, 32, int32(value))
}

// SetCurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseBuilder) SetCurrencyValuePerIncrement(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 40, 36, value)
}

// SetIsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseBuilder) SetIsFinalMessage(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 41, 40, value)
}

// SetFloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponseBuilder) SetFloatToIntPriceMultiplier(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 48, 44, value)
}

// SetIntToFloatPriceDivisor
func (m SecurityDefinitionResponseBuilder) SetIntToFloatPriceDivisor(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 52, 48, value)
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
func (m *SecurityDefinitionResponseBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 56, 52, value)
}

// SetUpdatesBidAskOnly This is set to 1 when the Symbol does not provide MarketDataUpdateTrade
// messages and only provides MarketDataUpdateBidAsk messages when there
// is market activity for the Symbol.
//
// Otherwise, when this is set to 0, MarketDataUpdateTrade messages will
// be received after subscribing to market data, when there is trading activity.
// be received after subscribing to market data, when there is trading activity.
func (m SecurityDefinitionResponseBuilder) SetUpdatesBidAskOnly(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 57, 56, value)
}

// SetStrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponseBuilder) SetStrikePrice(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 64, 60, value)
}

// SetPutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseBuilder) SetPutOrCall(value PutCallEnum) {
	message.SetUint8VLS(m.Unsafe(), 65, 64, uint8(value))
}

// SetShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseBuilder) SetShortInterest(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 72, 68, value)
}

// SetSecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseBuilder) SetSecurityExpirationDate(value DateTime4Byte) {
	message.SetUint32VLS(m.Unsafe(), 76, 72, uint32(value))
}

// SetBuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseBuilder) SetBuyRolloverInterest(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 80, 76, value)
}

// SetSellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseBuilder) SetSellRolloverInterest(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 84, 80, value)
}

// SetEarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseBuilder) SetEarningsPerShare(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 88, 84, value)
}

// SetSharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseBuilder) SetSharesOutstanding(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 92, 88, value)
}

// SetIntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseBuilder) SetIntToFloatQuantityDivisor(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 96, 92, value)
}

// SetHasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponseBuilder) SetHasMarketDepthData(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 97, 96, value)
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
func (m SecurityDefinitionResponseBuilder) SetDisplayPriceMultiplier(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 104, 100, value)
}

// SetExchangeSymbol This is an optional field. This is the exchange symbol which corresponds
// with the Symbol field.
//
// This field should be empty when this Security Definition message is in
// response to UNDERLYING_SymbolsForExchangeRequest.
func (m *SecurityDefinitionResponseBuilder) SetExchangeSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 108, 104, value)
}

// SetInitialMarginRequirement This field applies to the Futures Security Type.
//
// This is the initial margin requirement as specified by the exchange, if
// available.
func (m SecurityDefinitionResponseBuilder) SetInitialMarginRequirement(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 112, 108, value)
}

// SetMaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponseBuilder) SetMaintenanceMarginRequirement(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 116, 112, value)
}

// SetCurrency This is the currency that the Symbol trades in or is priced in.
func (m *SecurityDefinitionResponseBuilder) SetCurrency(value string) {
	message.SetStringVLS(&m.VLSBuilder, 120, 116, value)
}

// SetContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponseBuilder) SetContractSize(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 124, 120, value)
}

// SetOpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponseBuilder) SetOpenInterest(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 128, 124, value)
}

// SetRolloverDate This field applies to the Futures Security Type.
//
// This is the rollover date for the symbol according to the typical time
// where trading transitions from one contract month to the next.
func (m SecurityDefinitionResponseBuilder) SetRolloverDate(value DateTime4Byte) {
	message.SetUint32VLS(m.Unsafe(), 132, 128, uint32(value))
}

// SetIsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseBuilder) SetIsDelayed(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 133, 132, value)
}

// Copy
func (m SecurityDefinitionResponse) Copy(to SecurityDefinitionResponseBuilder) {
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
func (m SecurityDefinitionResponseBuilder) Copy(to SecurityDefinitionResponseBuilder) {
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
func (m SecurityDefinitionResponse) CopyPointer(to SecurityDefinitionResponsePointerBuilder) {
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
func (m SecurityDefinitionResponseBuilder) CopyPointer(to SecurityDefinitionResponsePointerBuilder) {
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
func (m SecurityDefinitionResponse) CopyToPointer(to SecurityDefinitionResponseFixedPointerBuilder) {
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
func (m SecurityDefinitionResponseBuilder) CopyToPointer(to SecurityDefinitionResponseFixedPointerBuilder) {
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
func (m SecurityDefinitionResponse) CopyTo(to SecurityDefinitionResponseFixedBuilder) {
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
func (m SecurityDefinitionResponseBuilder) CopyTo(to SecurityDefinitionResponseFixedBuilder) {
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
