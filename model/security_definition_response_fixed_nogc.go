package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SecurityDefinitionResponseFixedPointer This is a response from the Server in response to a SymbolsForExchangeRequest,
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
type SecurityDefinitionResponseFixedPointer struct {
	message.FixedPointer
}

// SecurityDefinitionResponseFixedPointerBuilder This is a response from the Server in response to a SymbolsForExchangeRequest,
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
type SecurityDefinitionResponseFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocSecurityDefinitionResponseFixed() SecurityDefinitionResponseFixedPointerBuilder {
	a := SecurityDefinitionResponseFixedPointerBuilder{message.AllocFixed(356)}
	a.Ptr.SetBytes(0, _SecurityDefinitionResponseFixedDefault)
	return a
}

func AllocSecurityDefinitionResponseFixedFrom(b []byte) SecurityDefinitionResponseFixedPointer {
	return SecurityDefinitionResponseFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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
func (m SecurityDefinitionResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SecurityDefinitionResponseFixedDefault)
}

// ToBuilder
func (m SecurityDefinitionResponseFixedPointer) ToBuilder() SecurityDefinitionResponseFixedPointerBuilder {
	return SecurityDefinitionResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *SecurityDefinitionResponseFixedPointerBuilder) Finish() SecurityDefinitionResponseFixedPointer {
	return SecurityDefinitionResponseFixedPointer{m.FixedPointer}
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponseFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponseFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Symbol This is the Symbol the Security Definition is for.
//
// When the Server responds with a single SecurityDefinitionResponse message
// and there are no security definitions to return for the original request,
// this will be empty.
//
// This field should be empty when this Security Definition message is in
// response to UnderlyingSymbolsForExchangeRequest.
func (m SecurityDefinitionResponseFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Symbol This is the Symbol the Security Definition is for.
//
// When the Server responds with a single SecurityDefinitionResponse message
// and there are no security definitions to return for the original request,
// this will be empty.
//
// This field should be empty when this Security Definition message is in
// response to UnderlyingSymbolsForExchangeRequest.
func (m SecurityDefinitionResponseFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Exchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponseFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 88, 72)
}

// Exchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponseFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 88, 72)
}

// SecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseFixedPointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Ptr, 92, 88))
}

// SecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Ptr, 92, 88))
}

// Description The text description for the Symbol.
func (m SecurityDefinitionResponseFixedPointer) Description() string {
	return message.StringFixed(m.Ptr, 156, 92)
}

// Description The text description for the Symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) Description() string {
	return message.StringFixed(m.Ptr, 156, 92)
}

// MinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseFixedPointer) MinPriceIncrement() float32 {
	return message.Float32Fixed(m.Ptr, 160, 156)
}

// MinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseFixedPointerBuilder) MinPriceIncrement() float32 {
	return message.Float32Fixed(m.Ptr, 160, 156)
}

// PriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseFixedPointer) PriceDisplayFormat() PriceDisplayFormatEnum {
	return PriceDisplayFormatEnum(message.Int32Fixed(m.Ptr, 164, 160))
}

// PriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseFixedPointerBuilder) PriceDisplayFormat() PriceDisplayFormatEnum {
	return PriceDisplayFormatEnum(message.Int32Fixed(m.Ptr, 164, 160))
}

// CurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseFixedPointer) CurrencyValuePerIncrement() float32 {
	return message.Float32Fixed(m.Ptr, 168, 164)
}

// CurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseFixedPointerBuilder) CurrencyValuePerIncrement() float32 {
	return message.Float32Fixed(m.Ptr, 168, 164)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseFixedPointer) IsFinalMessage() uint8 {
	return message.Uint8Fixed(m.Ptr, 169, 168)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseFixedPointerBuilder) IsFinalMessage() uint8 {
	return message.Uint8Fixed(m.Ptr, 169, 168)
}

// FloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponseFixedPointer) FloatToIntPriceMultiplier() float32 {
	return message.Float32Fixed(m.Ptr, 176, 172)
}

// FloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponseFixedPointerBuilder) FloatToIntPriceMultiplier() float32 {
	return message.Float32Fixed(m.Ptr, 176, 172)
}

// IntToFloatPriceDivisor
func (m SecurityDefinitionResponseFixedPointer) IntToFloatPriceDivisor() float32 {
	return message.Float32Fixed(m.Ptr, 180, 176)
}

// IntToFloatPriceDivisor
func (m SecurityDefinitionResponseFixedPointerBuilder) IntToFloatPriceDivisor() float32 {
	return message.Float32Fixed(m.Ptr, 180, 176)
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
func (m SecurityDefinitionResponseFixedPointer) UnderlyingSymbol() string {
	return message.StringFixed(m.Ptr, 212, 180)
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
func (m SecurityDefinitionResponseFixedPointerBuilder) UnderlyingSymbol() string {
	return message.StringFixed(m.Ptr, 212, 180)
}

// UpdatesBidAskOnly This is set to 1 when the Symbol does not provide MarketDataUpdateTrade
// messages and only provides MarketDataUpdateBidAsk messages when there
// is market activity for the Symbol.
//
// Otherwise, when this is set to 0, MarketDataUpdateTrade messages will
// be received after subscribing to market data, when there is trading activity.
// be received after subscribing to market data, when there is trading activity.
func (m SecurityDefinitionResponseFixedPointer) UpdatesBidAskOnly() uint8 {
	return message.Uint8Fixed(m.Ptr, 213, 212)
}

// UpdatesBidAskOnly This is set to 1 when the Symbol does not provide MarketDataUpdateTrade
// messages and only provides MarketDataUpdateBidAsk messages when there
// is market activity for the Symbol.
//
// Otherwise, when this is set to 0, MarketDataUpdateTrade messages will
// be received after subscribing to market data, when there is trading activity.
// be received after subscribing to market data, when there is trading activity.
func (m SecurityDefinitionResponseFixedPointerBuilder) UpdatesBidAskOnly() uint8 {
	return message.Uint8Fixed(m.Ptr, 213, 212)
}

// StrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponseFixedPointer) StrikePrice() float32 {
	return message.Float32Fixed(m.Ptr, 220, 216)
}

// StrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponseFixedPointerBuilder) StrikePrice() float32 {
	return message.Float32Fixed(m.Ptr, 220, 216)
}

// PutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseFixedPointer) PutOrCall() PutCallEnum {
	return PutCallEnum(message.Uint8Fixed(m.Ptr, 221, 220))
}

// PutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseFixedPointerBuilder) PutOrCall() PutCallEnum {
	return PutCallEnum(message.Uint8Fixed(m.Ptr, 221, 220))
}

// ShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseFixedPointer) ShortInterest() uint32 {
	return message.Uint32Fixed(m.Ptr, 228, 224)
}

// ShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseFixedPointerBuilder) ShortInterest() uint32 {
	return message.Uint32Fixed(m.Ptr, 228, 224)
}

// SecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseFixedPointer) SecurityExpirationDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 232, 228))
}

// SecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseFixedPointerBuilder) SecurityExpirationDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 232, 228))
}

// BuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedPointer) BuyRolloverInterest() float32 {
	return message.Float32Fixed(m.Ptr, 236, 232)
}

// BuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) BuyRolloverInterest() float32 {
	return message.Float32Fixed(m.Ptr, 236, 232)
}

// SellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedPointer) SellRolloverInterest() float32 {
	return message.Float32Fixed(m.Ptr, 240, 236)
}

// SellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) SellRolloverInterest() float32 {
	return message.Float32Fixed(m.Ptr, 240, 236)
}

// EarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseFixedPointer) EarningsPerShare() float32 {
	return message.Float32Fixed(m.Ptr, 244, 240)
}

// EarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseFixedPointerBuilder) EarningsPerShare() float32 {
	return message.Float32Fixed(m.Ptr, 244, 240)
}

// SharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseFixedPointer) SharesOutstanding() uint32 {
	return message.Uint32Fixed(m.Ptr, 248, 244)
}

// SharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseFixedPointerBuilder) SharesOutstanding() uint32 {
	return message.Uint32Fixed(m.Ptr, 248, 244)
}

// IntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseFixedPointer) IntToFloatQuantityDivisor() float32 {
	return message.Float32Fixed(m.Ptr, 252, 248)
}

// IntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseFixedPointerBuilder) IntToFloatQuantityDivisor() float32 {
	return message.Float32Fixed(m.Ptr, 252, 248)
}

// HasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponseFixedPointer) HasMarketDepthData() uint8 {
	return message.Uint8Fixed(m.Ptr, 253, 252)
}

// HasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) HasMarketDepthData() uint8 {
	return message.Uint8Fixed(m.Ptr, 253, 252)
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
func (m SecurityDefinitionResponseFixedPointer) DisplayPriceMultiplier() float32 {
	return message.Float32Fixed(m.Ptr, 260, 256)
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
func (m SecurityDefinitionResponseFixedPointerBuilder) DisplayPriceMultiplier() float32 {
	return message.Float32Fixed(m.Ptr, 260, 256)
}

// ExchangeSymbol This is an optional field. This is the exchange symbol which corresponds
// with the Symbol field.
//
// This field should be empty when this Security Definition message is in
// response to UNDERLYING_SymbolsForExchangeRequest.
func (m SecurityDefinitionResponseFixedPointer) ExchangeSymbol() string {
	return message.StringFixed(m.Ptr, 324, 260)
}

// ExchangeSymbol This is an optional field. This is the exchange symbol which corresponds
// with the Symbol field.
//
// This field should be empty when this Security Definition message is in
// response to UNDERLYING_SymbolsForExchangeRequest.
func (m SecurityDefinitionResponseFixedPointerBuilder) ExchangeSymbol() string {
	return message.StringFixed(m.Ptr, 324, 260)
}

// InitialMarginRequirement This field applies to the Futures Security Type.
//
// This is the initial margin requirement as specified by the exchange, if
// available.
func (m SecurityDefinitionResponseFixedPointer) InitialMarginRequirement() float32 {
	return message.Float32Fixed(m.Ptr, 328, 324)
}

// InitialMarginRequirement This field applies to the Futures Security Type.
//
// This is the initial margin requirement as specified by the exchange, if
// available.
func (m SecurityDefinitionResponseFixedPointerBuilder) InitialMarginRequirement() float32 {
	return message.Float32Fixed(m.Ptr, 328, 324)
}

// MaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponseFixedPointer) MaintenanceMarginRequirement() float32 {
	return message.Float32Fixed(m.Ptr, 332, 328)
}

// MaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponseFixedPointerBuilder) MaintenanceMarginRequirement() float32 {
	return message.Float32Fixed(m.Ptr, 332, 328)
}

// Currency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponseFixedPointer) Currency() string {
	return message.StringFixed(m.Ptr, 340, 332)
}

// Currency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponseFixedPointerBuilder) Currency() string {
	return message.StringFixed(m.Ptr, 340, 332)
}

// ContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponseFixedPointer) ContractSize() float32 {
	return message.Float32Fixed(m.Ptr, 344, 340)
}

// ContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponseFixedPointerBuilder) ContractSize() float32 {
	return message.Float32Fixed(m.Ptr, 344, 340)
}

// OpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponseFixedPointer) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Ptr, 348, 344)
}

// OpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponseFixedPointerBuilder) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Ptr, 348, 344)
}

// RolloverDate This field applies to the Futures Security Type.
//
// This is the rollover date for the symbol according to the typical time
// where trading transitions from one contract month to the next.
func (m SecurityDefinitionResponseFixedPointer) RolloverDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 352, 348))
}

// RolloverDate This field applies to the Futures Security Type.
//
// This is the rollover date for the symbol according to the typical time
// where trading transitions from one contract month to the next.
func (m SecurityDefinitionResponseFixedPointerBuilder) RolloverDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 352, 348))
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseFixedPointer) IsDelayed() uint8 {
	return message.Uint8Fixed(m.Ptr, 353, 352)
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseFixedPointerBuilder) IsDelayed() uint8 {
	return message.Uint8Fixed(m.Ptr, 353, 352)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetSymbol This is the Symbol the Security Definition is for.
//
// When the Server responds with a single SecurityDefinitionResponse message
// and there are no security definitions to return for the original request,
// this will be empty.
//
// This field should be empty when this Security Definition message is in
// response to UnderlyingSymbolsForExchangeRequest.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 72, 8, value)
}

// SetExchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 88, 72, value)
}

// SetSecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 92, 88, int32(value))
}

// SetDescription The text description for the Symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetDescription(value string) {
	message.SetStringFixed(m.Ptr, 156, 92, value)
}

// SetMinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetMinPriceIncrement(value float32) {
	message.SetFloat32Fixed(m.Ptr, 160, 156, value)
}

// SetPriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetPriceDisplayFormat(value PriceDisplayFormatEnum) {
	message.SetInt32Fixed(m.Ptr, 164, 160, int32(value))
}

// SetCurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetCurrencyValuePerIncrement(value float32) {
	message.SetFloat32Fixed(m.Ptr, 168, 164, value)
}

// SetIsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetIsFinalMessage(value uint8) {
	message.SetUint8Fixed(m.Ptr, 169, 168, value)
}

// SetFloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetFloatToIntPriceMultiplier(value float32) {
	message.SetFloat32Fixed(m.Ptr, 176, 172, value)
}

// SetIntToFloatPriceDivisor
func (m SecurityDefinitionResponseFixedPointerBuilder) SetIntToFloatPriceDivisor(value float32) {
	message.SetFloat32Fixed(m.Ptr, 180, 176, value)
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
func (m SecurityDefinitionResponseFixedPointerBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringFixed(m.Ptr, 212, 180, value)
}

// SetUpdatesBidAskOnly This is set to 1 when the Symbol does not provide MarketDataUpdateTrade
// messages and only provides MarketDataUpdateBidAsk messages when there
// is market activity for the Symbol.
//
// Otherwise, when this is set to 0, MarketDataUpdateTrade messages will
// be received after subscribing to market data, when there is trading activity.
// be received after subscribing to market data, when there is trading activity.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetUpdatesBidAskOnly(value uint8) {
	message.SetUint8Fixed(m.Ptr, 213, 212, value)
}

// SetStrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetStrikePrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 220, 216, value)
}

// SetPutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetPutOrCall(value PutCallEnum) {
	message.SetUint8Fixed(m.Ptr, 221, 220, uint8(value))
}

// SetShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetShortInterest(value uint32) {
	message.SetUint32Fixed(m.Ptr, 228, 224, value)
}

// SetSecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetSecurityExpirationDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 232, 228, uint32(value))
}

// SetBuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetBuyRolloverInterest(value float32) {
	message.SetFloat32Fixed(m.Ptr, 236, 232, value)
}

// SetSellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetSellRolloverInterest(value float32) {
	message.SetFloat32Fixed(m.Ptr, 240, 236, value)
}

// SetEarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetEarningsPerShare(value float32) {
	message.SetFloat32Fixed(m.Ptr, 244, 240, value)
}

// SetSharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetSharesOutstanding(value uint32) {
	message.SetUint32Fixed(m.Ptr, 248, 244, value)
}

// SetIntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetIntToFloatQuantityDivisor(value float32) {
	message.SetFloat32Fixed(m.Ptr, 252, 248, value)
}

// SetHasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetHasMarketDepthData(value uint8) {
	message.SetUint8Fixed(m.Ptr, 253, 252, value)
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
func (m SecurityDefinitionResponseFixedPointerBuilder) SetDisplayPriceMultiplier(value float32) {
	message.SetFloat32Fixed(m.Ptr, 260, 256, value)
}

// SetExchangeSymbol This is an optional field. This is the exchange symbol which corresponds
// with the Symbol field.
//
// This field should be empty when this Security Definition message is in
// response to UNDERLYING_SymbolsForExchangeRequest.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetExchangeSymbol(value string) {
	message.SetStringFixed(m.Ptr, 324, 260, value)
}

// SetInitialMarginRequirement This field applies to the Futures Security Type.
//
// This is the initial margin requirement as specified by the exchange, if
// available.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetInitialMarginRequirement(value float32) {
	message.SetFloat32Fixed(m.Ptr, 328, 324, value)
}

// SetMaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetMaintenanceMarginRequirement(value float32) {
	message.SetFloat32Fixed(m.Ptr, 332, 328, value)
}

// SetCurrency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetCurrency(value string) {
	message.SetStringFixed(m.Ptr, 340, 332, value)
}

// SetContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetContractSize(value float32) {
	message.SetFloat32Fixed(m.Ptr, 344, 340, value)
}

// SetOpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetOpenInterest(value uint32) {
	message.SetUint32Fixed(m.Ptr, 348, 344, value)
}

// SetRolloverDate This field applies to the Futures Security Type.
//
// This is the rollover date for the symbol according to the typical time
// where trading transitions from one contract month to the next.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetRolloverDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 352, 348, uint32(value))
}

// SetIsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetIsDelayed(value uint8) {
	message.SetUint8Fixed(m.Ptr, 353, 352, value)
}

// Copy
func (m SecurityDefinitionResponseFixedPointer) Copy(to SecurityDefinitionResponseFixedBuilder) {
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
func (m SecurityDefinitionResponseFixedPointerBuilder) Copy(to SecurityDefinitionResponseFixedBuilder) {
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
func (m SecurityDefinitionResponseFixedPointer) CopyPointer(to SecurityDefinitionResponseFixedPointerBuilder) {
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
func (m SecurityDefinitionResponseFixedPointerBuilder) CopyPointer(to SecurityDefinitionResponseFixedPointerBuilder) {
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
func (m SecurityDefinitionResponseFixedPointer) CopyTo(to SecurityDefinitionResponseBuilder) {
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
func (m SecurityDefinitionResponseFixedPointerBuilder) CopyTo(to SecurityDefinitionResponseBuilder) {
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
func (m SecurityDefinitionResponseFixedPointer) CopyToPointer(to SecurityDefinitionResponsePointerBuilder) {
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
func (m SecurityDefinitionResponseFixedPointerBuilder) CopyToPointer(to SecurityDefinitionResponsePointerBuilder) {
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
