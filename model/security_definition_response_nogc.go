package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SecurityDefinitionResponsePointer This is a response from the Server in response to a SymbolsForExchangeRequest,
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
type SecurityDefinitionResponsePointer struct {
	message.VLSPointer
}

// SecurityDefinitionResponsePointerBuilder This is a response from the Server in response to a SymbolsForExchangeRequest,
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
type SecurityDefinitionResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocSecurityDefinitionResponse() SecurityDefinitionResponsePointerBuilder {
	a := SecurityDefinitionResponsePointerBuilder{message.AllocVLSBuilder(136)}
	a.Ptr.SetBytes(0, _SecurityDefinitionResponseDefault)
	return a
}

func AllocSecurityDefinitionResponseFrom(b []byte) SecurityDefinitionResponsePointer {
	return SecurityDefinitionResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m SecurityDefinitionResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SecurityDefinitionResponseDefault)
}

// ToBuilder
func (m SecurityDefinitionResponsePointer) ToBuilder() SecurityDefinitionResponsePointerBuilder {
	return SecurityDefinitionResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *SecurityDefinitionResponsePointerBuilder) Finish() SecurityDefinitionResponsePointer {
	return SecurityDefinitionResponsePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponsePointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponsePointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// Symbol This is the Symbol the Security Definition is for.
//
// When the Server responds with a single SecurityDefinitionResponse message
// and there are no security definitions to return for the original request,
// this will be empty.
//
// This field should be empty when this Security Definition message is in
// response to UnderlyingSymbolsForExchangeRequest.
func (m SecurityDefinitionResponsePointer) Symbol() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Symbol This is the Symbol the Security Definition is for.
//
// When the Server responds with a single SecurityDefinitionResponse message
// and there are no security definitions to return for the original request,
// this will be empty.
//
// This field should be empty when this Security Definition message is in
// response to UnderlyingSymbolsForExchangeRequest.
func (m SecurityDefinitionResponsePointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Exchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponsePointer) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Exchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponsePointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// SecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponsePointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// SecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponsePointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// Description The text description for the Symbol.
func (m SecurityDefinitionResponsePointer) Description() string {
	return message.StringVLS(m.Ptr, 28, 24)
}

// Description The text description for the Symbol.
func (m SecurityDefinitionResponsePointerBuilder) Description() string {
	return message.StringVLS(m.Ptr, 28, 24)
}

// MinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponsePointer) MinPriceIncrement() float32 {
	return message.Float32VLS(m.Ptr, 32, 28)
}

// MinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponsePointerBuilder) MinPriceIncrement() float32 {
	return message.Float32VLS(m.Ptr, 32, 28)
}

// PriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponsePointer) PriceDisplayFormat() PriceDisplayFormatEnum {
	return PriceDisplayFormatEnum(message.Int32VLS(m.Ptr, 36, 32))
}

// PriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponsePointerBuilder) PriceDisplayFormat() PriceDisplayFormatEnum {
	return PriceDisplayFormatEnum(message.Int32VLS(m.Ptr, 36, 32))
}

// CurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponsePointer) CurrencyValuePerIncrement() float32 {
	return message.Float32VLS(m.Ptr, 40, 36)
}

// CurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponsePointerBuilder) CurrencyValuePerIncrement() float32 {
	return message.Float32VLS(m.Ptr, 40, 36)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponsePointer) IsFinalMessage() uint8 {
	return message.Uint8VLS(m.Ptr, 41, 40)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponsePointerBuilder) IsFinalMessage() uint8 {
	return message.Uint8VLS(m.Ptr, 41, 40)
}

// FloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponsePointer) FloatToIntPriceMultiplier() float32 {
	return message.Float32VLS(m.Ptr, 48, 44)
}

// FloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponsePointerBuilder) FloatToIntPriceMultiplier() float32 {
	return message.Float32VLS(m.Ptr, 48, 44)
}

// IntToFloatPriceDivisor
func (m SecurityDefinitionResponsePointer) IntToFloatPriceDivisor() float32 {
	return message.Float32VLS(m.Ptr, 52, 48)
}

// IntToFloatPriceDivisor
func (m SecurityDefinitionResponsePointerBuilder) IntToFloatPriceDivisor() float32 {
	return message.Float32VLS(m.Ptr, 52, 48)
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
func (m SecurityDefinitionResponsePointer) UnderlyingSymbol() string {
	return message.StringVLS(m.Ptr, 56, 52)
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
func (m SecurityDefinitionResponsePointerBuilder) UnderlyingSymbol() string {
	return message.StringVLS(m.Ptr, 56, 52)
}

// UpdatesBidAskOnly This is set to 1 when the Symbol does not provide MarketDataUpdateTrade
// messages and only provides MarketDataUpdateBidAsk messages when there
// is market activity for the Symbol.
//
// Otherwise, when this is set to 0, MarketDataUpdateTrade messages will
// be received after subscribing to market data, when there is trading activity.
// be received after subscribing to market data, when there is trading activity.
func (m SecurityDefinitionResponsePointer) UpdatesBidAskOnly() uint8 {
	return message.Uint8VLS(m.Ptr, 57, 56)
}

// UpdatesBidAskOnly This is set to 1 when the Symbol does not provide MarketDataUpdateTrade
// messages and only provides MarketDataUpdateBidAsk messages when there
// is market activity for the Symbol.
//
// Otherwise, when this is set to 0, MarketDataUpdateTrade messages will
// be received after subscribing to market data, when there is trading activity.
// be received after subscribing to market data, when there is trading activity.
func (m SecurityDefinitionResponsePointerBuilder) UpdatesBidAskOnly() uint8 {
	return message.Uint8VLS(m.Ptr, 57, 56)
}

// StrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponsePointer) StrikePrice() float32 {
	return message.Float32VLS(m.Ptr, 64, 60)
}

// StrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponsePointerBuilder) StrikePrice() float32 {
	return message.Float32VLS(m.Ptr, 64, 60)
}

// PutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponsePointer) PutOrCall() PutCallEnum {
	return PutCallEnum(message.Uint8VLS(m.Ptr, 65, 64))
}

// PutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponsePointerBuilder) PutOrCall() PutCallEnum {
	return PutCallEnum(message.Uint8VLS(m.Ptr, 65, 64))
}

// ShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponsePointer) ShortInterest() uint32 {
	return message.Uint32VLS(m.Ptr, 72, 68)
}

// ShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponsePointerBuilder) ShortInterest() uint32 {
	return message.Uint32VLS(m.Ptr, 72, 68)
}

// SecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponsePointer) SecurityExpirationDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Ptr, 76, 72))
}

// SecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponsePointerBuilder) SecurityExpirationDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Ptr, 76, 72))
}

// BuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponsePointer) BuyRolloverInterest() float32 {
	return message.Float32VLS(m.Ptr, 80, 76)
}

// BuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponsePointerBuilder) BuyRolloverInterest() float32 {
	return message.Float32VLS(m.Ptr, 80, 76)
}

// SellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponsePointer) SellRolloverInterest() float32 {
	return message.Float32VLS(m.Ptr, 84, 80)
}

// SellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponsePointerBuilder) SellRolloverInterest() float32 {
	return message.Float32VLS(m.Ptr, 84, 80)
}

// EarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponsePointer) EarningsPerShare() float32 {
	return message.Float32VLS(m.Ptr, 88, 84)
}

// EarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponsePointerBuilder) EarningsPerShare() float32 {
	return message.Float32VLS(m.Ptr, 88, 84)
}

// SharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponsePointer) SharesOutstanding() uint32 {
	return message.Uint32VLS(m.Ptr, 92, 88)
}

// SharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponsePointerBuilder) SharesOutstanding() uint32 {
	return message.Uint32VLS(m.Ptr, 92, 88)
}

// IntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponsePointer) IntToFloatQuantityDivisor() float32 {
	return message.Float32VLS(m.Ptr, 96, 92)
}

// IntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponsePointerBuilder) IntToFloatQuantityDivisor() float32 {
	return message.Float32VLS(m.Ptr, 96, 92)
}

// HasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponsePointer) HasMarketDepthData() uint8 {
	return message.Uint8VLS(m.Ptr, 97, 96)
}

// HasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponsePointerBuilder) HasMarketDepthData() uint8 {
	return message.Uint8VLS(m.Ptr, 97, 96)
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
func (m SecurityDefinitionResponsePointer) DisplayPriceMultiplier() float32 {
	return message.Float32VLS(m.Ptr, 104, 100)
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
func (m SecurityDefinitionResponsePointerBuilder) DisplayPriceMultiplier() float32 {
	return message.Float32VLS(m.Ptr, 104, 100)
}

// ExchangeSymbol This is an optional field. This is the exchange symbol which corresponds
// with the Symbol field.
//
// This field should be empty when this Security Definition message is in
// response to UNDERLYING_SymbolsForExchangeRequest.
func (m SecurityDefinitionResponsePointer) ExchangeSymbol() string {
	return message.StringVLS(m.Ptr, 108, 104)
}

// ExchangeSymbol This is an optional field. This is the exchange symbol which corresponds
// with the Symbol field.
//
// This field should be empty when this Security Definition message is in
// response to UNDERLYING_SymbolsForExchangeRequest.
func (m SecurityDefinitionResponsePointerBuilder) ExchangeSymbol() string {
	return message.StringVLS(m.Ptr, 108, 104)
}

// InitialMarginRequirement This field applies to the Futures Security Type.
//
// This is the initial margin requirement as specified by the exchange, if
// available.
func (m SecurityDefinitionResponsePointer) InitialMarginRequirement() float32 {
	return message.Float32VLS(m.Ptr, 112, 108)
}

// InitialMarginRequirement This field applies to the Futures Security Type.
//
// This is the initial margin requirement as specified by the exchange, if
// available.
func (m SecurityDefinitionResponsePointerBuilder) InitialMarginRequirement() float32 {
	return message.Float32VLS(m.Ptr, 112, 108)
}

// MaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponsePointer) MaintenanceMarginRequirement() float32 {
	return message.Float32VLS(m.Ptr, 116, 112)
}

// MaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponsePointerBuilder) MaintenanceMarginRequirement() float32 {
	return message.Float32VLS(m.Ptr, 116, 112)
}

// Currency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponsePointer) Currency() string {
	return message.StringVLS(m.Ptr, 120, 116)
}

// Currency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponsePointerBuilder) Currency() string {
	return message.StringVLS(m.Ptr, 120, 116)
}

// ContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponsePointer) ContractSize() float32 {
	return message.Float32VLS(m.Ptr, 124, 120)
}

// ContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponsePointerBuilder) ContractSize() float32 {
	return message.Float32VLS(m.Ptr, 124, 120)
}

// OpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponsePointer) OpenInterest() uint32 {
	return message.Uint32VLS(m.Ptr, 128, 124)
}

// OpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponsePointerBuilder) OpenInterest() uint32 {
	return message.Uint32VLS(m.Ptr, 128, 124)
}

// RolloverDate This field applies to the Futures Security Type.
//
// This is the rollover date for the symbol according to the typical time
// where trading transitions from one contract month to the next.
func (m SecurityDefinitionResponsePointer) RolloverDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Ptr, 132, 128))
}

// RolloverDate This field applies to the Futures Security Type.
//
// This is the rollover date for the symbol according to the typical time
// where trading transitions from one contract month to the next.
func (m SecurityDefinitionResponsePointerBuilder) RolloverDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Ptr, 132, 128))
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponsePointer) IsDelayed() uint8 {
	return message.Uint8VLS(m.Ptr, 133, 132)
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponsePointerBuilder) IsDelayed() uint8 {
	return message.Uint8VLS(m.Ptr, 133, 132)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetSymbol This is the Symbol the Security Definition is for.
//
// When the Server responds with a single SecurityDefinitionResponse message
// and there are no security definitions to return for the original request,
// this will be empty.
//
// This field should be empty when this Security Definition message is in
// response to UnderlyingSymbolsForExchangeRequest.
func (m *SecurityDefinitionResponsePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetExchange This is the Exchange for the Symbol. This field is optional.
func (m *SecurityDefinitionResponsePointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetSecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponsePointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Ptr, 24, 20, int32(value))
}

// SetDescription The text description for the Symbol.
func (m *SecurityDefinitionResponsePointerBuilder) SetDescription(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 28, 24, value)
}

// SetMinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponsePointerBuilder) SetMinPriceIncrement(value float32) {
	message.SetFloat32VLS(m.Ptr, 32, 28, value)
}

// SetPriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponsePointerBuilder) SetPriceDisplayFormat(value PriceDisplayFormatEnum) {
	message.SetInt32VLS(m.Ptr, 36, 32, int32(value))
}

// SetCurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponsePointerBuilder) SetCurrencyValuePerIncrement(value float32) {
	message.SetFloat32VLS(m.Ptr, 40, 36, value)
}

// SetIsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponsePointerBuilder) SetIsFinalMessage(value uint8) {
	message.SetUint8VLS(m.Ptr, 41, 40, value)
}

// SetFloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponsePointerBuilder) SetFloatToIntPriceMultiplier(value float32) {
	message.SetFloat32VLS(m.Ptr, 48, 44, value)
}

// SetIntToFloatPriceDivisor
func (m SecurityDefinitionResponsePointerBuilder) SetIntToFloatPriceDivisor(value float32) {
	message.SetFloat32VLS(m.Ptr, 52, 48, value)
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
func (m *SecurityDefinitionResponsePointerBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 56, 52, value)
}

// SetUpdatesBidAskOnly This is set to 1 when the Symbol does not provide MarketDataUpdateTrade
// messages and only provides MarketDataUpdateBidAsk messages when there
// is market activity for the Symbol.
//
// Otherwise, when this is set to 0, MarketDataUpdateTrade messages will
// be received after subscribing to market data, when there is trading activity.
// be received after subscribing to market data, when there is trading activity.
func (m SecurityDefinitionResponsePointerBuilder) SetUpdatesBidAskOnly(value uint8) {
	message.SetUint8VLS(m.Ptr, 57, 56, value)
}

// SetStrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponsePointerBuilder) SetStrikePrice(value float32) {
	message.SetFloat32VLS(m.Ptr, 64, 60, value)
}

// SetPutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponsePointerBuilder) SetPutOrCall(value PutCallEnum) {
	message.SetUint8VLS(m.Ptr, 65, 64, uint8(value))
}

// SetShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponsePointerBuilder) SetShortInterest(value uint32) {
	message.SetUint32VLS(m.Ptr, 72, 68, value)
}

// SetSecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponsePointerBuilder) SetSecurityExpirationDate(value DateTime4Byte) {
	message.SetUint32VLS(m.Ptr, 76, 72, uint32(value))
}

// SetBuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponsePointerBuilder) SetBuyRolloverInterest(value float32) {
	message.SetFloat32VLS(m.Ptr, 80, 76, value)
}

// SetSellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponsePointerBuilder) SetSellRolloverInterest(value float32) {
	message.SetFloat32VLS(m.Ptr, 84, 80, value)
}

// SetEarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponsePointerBuilder) SetEarningsPerShare(value float32) {
	message.SetFloat32VLS(m.Ptr, 88, 84, value)
}

// SetSharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponsePointerBuilder) SetSharesOutstanding(value uint32) {
	message.SetUint32VLS(m.Ptr, 92, 88, value)
}

// SetIntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponsePointerBuilder) SetIntToFloatQuantityDivisor(value float32) {
	message.SetFloat32VLS(m.Ptr, 96, 92, value)
}

// SetHasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponsePointerBuilder) SetHasMarketDepthData(value uint8) {
	message.SetUint8VLS(m.Ptr, 97, 96, value)
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
func (m SecurityDefinitionResponsePointerBuilder) SetDisplayPriceMultiplier(value float32) {
	message.SetFloat32VLS(m.Ptr, 104, 100, value)
}

// SetExchangeSymbol This is an optional field. This is the exchange symbol which corresponds
// with the Symbol field.
//
// This field should be empty when this Security Definition message is in
// response to UNDERLYING_SymbolsForExchangeRequest.
func (m *SecurityDefinitionResponsePointerBuilder) SetExchangeSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 108, 104, value)
}

// SetInitialMarginRequirement This field applies to the Futures Security Type.
//
// This is the initial margin requirement as specified by the exchange, if
// available.
func (m SecurityDefinitionResponsePointerBuilder) SetInitialMarginRequirement(value float32) {
	message.SetFloat32VLS(m.Ptr, 112, 108, value)
}

// SetMaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponsePointerBuilder) SetMaintenanceMarginRequirement(value float32) {
	message.SetFloat32VLS(m.Ptr, 116, 112, value)
}

// SetCurrency This is the currency that the Symbol trades in or is priced in.
func (m *SecurityDefinitionResponsePointerBuilder) SetCurrency(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 120, 116, value)
}

// SetContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponsePointerBuilder) SetContractSize(value float32) {
	message.SetFloat32VLS(m.Ptr, 124, 120, value)
}

// SetOpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponsePointerBuilder) SetOpenInterest(value uint32) {
	message.SetUint32VLS(m.Ptr, 128, 124, value)
}

// SetRolloverDate This field applies to the Futures Security Type.
//
// This is the rollover date for the symbol according to the typical time
// where trading transitions from one contract month to the next.
func (m SecurityDefinitionResponsePointerBuilder) SetRolloverDate(value DateTime4Byte) {
	message.SetUint32VLS(m.Ptr, 132, 128, uint32(value))
}

// SetIsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponsePointerBuilder) SetIsDelayed(value uint8) {
	message.SetUint8VLS(m.Ptr, 133, 132, value)
}

// Copy
func (m SecurityDefinitionResponsePointer) Copy(to SecurityDefinitionResponseBuilder) {
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
func (m SecurityDefinitionResponsePointerBuilder) Copy(to SecurityDefinitionResponseBuilder) {
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
func (m SecurityDefinitionResponsePointer) CopyPointer(to SecurityDefinitionResponsePointerBuilder) {
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
func (m SecurityDefinitionResponsePointerBuilder) CopyPointer(to SecurityDefinitionResponsePointerBuilder) {
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
func (m SecurityDefinitionResponsePointer) CopyTo(to SecurityDefinitionResponseFixedBuilder) {
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
func (m SecurityDefinitionResponsePointerBuilder) CopyTo(to SecurityDefinitionResponseFixedBuilder) {
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
func (m SecurityDefinitionResponsePointer) CopyToPointer(to SecurityDefinitionResponseFixedPointerBuilder) {
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
func (m SecurityDefinitionResponsePointerBuilder) CopyToPointer(to SecurityDefinitionResponseFixedPointerBuilder) {
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
