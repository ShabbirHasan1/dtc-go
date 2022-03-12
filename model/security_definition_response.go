package model

import (
	"github.com/moontrade/dtc-go/message"
)

const SecurityDefinitionResponseSize = 136

const SecurityDefinitionResponseFixedSize = 356

// {
//     Size                         = SecurityDefinitionResponseSize  (136)
//     Type                         = SECURITY_DEFINITION_RESPONSE  (507)
//     BaseSize                     = SecurityDefinitionResponseSize  (136)
//     RequestID                    = 0
//     Symbol                       = ""
//     Exchange                     = ""
//     SecurityType                 = SECURITY_TYPE_UNSET  (0)
//     Description                  = ""
//     MinPriceIncrement            = 0
//     PriceDisplayFormat           = PRICE_DISPLAY_FORMAT_DECIMAL_0  (0)
//     CurrencyValuePerIncrement    = 0
//     IsFinalMessage               = false
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
//     IsDelayed                    = false
// }
var _SecurityDefinitionResponseDefault = []byte{136, 0, 251, 1, 136, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size                         = SecurityDefinitionResponseFixedSize  (356)
//     Type                         = SECURITY_DEFINITION_RESPONSE  (507)
//     RequestID                    = 0
//     Symbol                       = ""
//     Exchange                     = ""
//     SecurityType                 = SECURITY_TYPE_UNSET  (0)
//     Description                  = ""
//     MinPriceIncrement            = 0
//     PriceDisplayFormat           = PRICE_DISPLAY_FORMAT_UNSET  (-1)
//     CurrencyValuePerIncrement    = 0
//     IsFinalMessage               = false
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
//     IsDelayed                    = false
// }
var _SecurityDefinitionResponseFixedDefault = []byte{100, 1, 251, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocSecurityDefinitionResponse() SecurityDefinitionResponsePointerBuilder {
	a := SecurityDefinitionResponsePointerBuilder{message.AllocVLSBuilder(136)}
	a.Ptr.SetBytes(0, _SecurityDefinitionResponseDefault)
	return a
}

func AllocSecurityDefinitionResponseFrom(b []byte) SecurityDefinitionResponsePointer {
	return SecurityDefinitionResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size                         = SecurityDefinitionResponseSize  (136)
//     Type                         = SECURITY_DEFINITION_RESPONSE  (507)
//     BaseSize                     = SecurityDefinitionResponseSize  (136)
//     RequestID                    = 0
//     Symbol                       = ""
//     Exchange                     = ""
//     SecurityType                 = SECURITY_TYPE_UNSET  (0)
//     Description                  = ""
//     MinPriceIncrement            = 0
//     PriceDisplayFormat           = PRICE_DISPLAY_FORMAT_DECIMAL_0  (0)
//     CurrencyValuePerIncrement    = 0
//     IsFinalMessage               = false
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
//     IsDelayed                    = false
// }
func (m SecurityDefinitionResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SecurityDefinitionResponseDefault)
}

// Clear
// {
//     Size                         = SecurityDefinitionResponseFixedSize  (356)
//     Type                         = SECURITY_DEFINITION_RESPONSE  (507)
//     RequestID                    = 0
//     Symbol                       = ""
//     Exchange                     = ""
//     SecurityType                 = SECURITY_TYPE_UNSET  (0)
//     Description                  = ""
//     MinPriceIncrement            = 0
//     PriceDisplayFormat           = PRICE_DISPLAY_FORMAT_UNSET  (-1)
//     CurrencyValuePerIncrement    = 0
//     IsFinalMessage               = false
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
//     IsDelayed                    = false
// }
func (m SecurityDefinitionResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SecurityDefinitionResponseFixedDefault)
}

// Clear
// {
//     Size                         = SecurityDefinitionResponseSize  (136)
//     Type                         = SECURITY_DEFINITION_RESPONSE  (507)
//     BaseSize                     = SecurityDefinitionResponseSize  (136)
//     RequestID                    = 0
//     Symbol                       = ""
//     Exchange                     = ""
//     SecurityType                 = SECURITY_TYPE_UNSET  (0)
//     Description                  = ""
//     MinPriceIncrement            = 0
//     PriceDisplayFormat           = PRICE_DISPLAY_FORMAT_DECIMAL_0  (0)
//     CurrencyValuePerIncrement    = 0
//     IsFinalMessage               = false
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
//     IsDelayed                    = false
// }
func (m SecurityDefinitionResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SecurityDefinitionResponseDefault)
}

// Clear
// {
//     Size                         = SecurityDefinitionResponseFixedSize  (356)
//     Type                         = SECURITY_DEFINITION_RESPONSE  (507)
//     RequestID                    = 0
//     Symbol                       = ""
//     Exchange                     = ""
//     SecurityType                 = SECURITY_TYPE_UNSET  (0)
//     Description                  = ""
//     MinPriceIncrement            = 0
//     PriceDisplayFormat           = PRICE_DISPLAY_FORMAT_UNSET  (-1)
//     CurrencyValuePerIncrement    = 0
//     IsFinalMessage               = false
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
//     IsDelayed                    = false
// }
func (m SecurityDefinitionResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SecurityDefinitionResponseFixedDefault)
}

// ToBuilder
func (m SecurityDefinitionResponse) ToBuilder() SecurityDefinitionResponseBuilder {
	return SecurityDefinitionResponseBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m SecurityDefinitionResponseFixed) ToBuilder() SecurityDefinitionResponseFixedBuilder {
	return SecurityDefinitionResponseFixedBuilder{m.Fixed}
}

// ToBuilder
func (m SecurityDefinitionResponsePointer) ToBuilder() SecurityDefinitionResponsePointerBuilder {
	return SecurityDefinitionResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m SecurityDefinitionResponseFixedPointer) ToBuilder() SecurityDefinitionResponseFixedPointerBuilder {
	return SecurityDefinitionResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m SecurityDefinitionResponseBuilder) Finish() SecurityDefinitionResponse {
	return SecurityDefinitionResponse{m.VLSBuilder.Finish()}
}

// Finish
func (m SecurityDefinitionResponseFixedBuilder) Finish() SecurityDefinitionResponseFixed {
	return SecurityDefinitionResponseFixed{m.Fixed}
}

// Finish
func (m *SecurityDefinitionResponsePointerBuilder) Finish() SecurityDefinitionResponsePointer {
	return SecurityDefinitionResponsePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *SecurityDefinitionResponseFixedPointerBuilder) Finish() SecurityDefinitionResponseFixedPointer {
	return SecurityDefinitionResponseFixedPointer{m.FixedPointer}
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
func (m SecurityDefinitionResponse) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponseBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
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
func (m SecurityDefinitionResponse) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// SecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
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
func (m SecurityDefinitionResponse) Description() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Description The text description for the Symbol.
func (m SecurityDefinitionResponseBuilder) Description() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
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
func (m SecurityDefinitionResponse) MinPriceIncrement() float32 {
	return message.Float32VLS(m.Unsafe(), 32, 28)
}

// MinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseBuilder) MinPriceIncrement() float32 {
	return message.Float32VLS(m.Unsafe(), 32, 28)
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
func (m SecurityDefinitionResponse) PriceDisplayFormat() PriceDisplayFormatEnum {
	return PriceDisplayFormatEnum(message.Int32VLS(m.Unsafe(), 36, 32))
}

// PriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseBuilder) PriceDisplayFormat() PriceDisplayFormatEnum {
	return PriceDisplayFormatEnum(message.Int32VLS(m.Unsafe(), 36, 32))
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
func (m SecurityDefinitionResponse) CurrencyValuePerIncrement() float32 {
	return message.Float32VLS(m.Unsafe(), 40, 36)
}

// CurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseBuilder) CurrencyValuePerIncrement() float32 {
	return message.Float32VLS(m.Unsafe(), 40, 36)
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
func (m SecurityDefinitionResponse) IsFinalMessage() bool {
	return message.BoolVLS(m.Unsafe(), 41, 40)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseBuilder) IsFinalMessage() bool {
	return message.BoolVLS(m.Unsafe(), 41, 40)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponsePointer) IsFinalMessage() bool {
	return message.BoolVLS(m.Ptr, 41, 40)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponsePointerBuilder) IsFinalMessage() bool {
	return message.BoolVLS(m.Ptr, 41, 40)
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
func (m SecurityDefinitionResponse) IntToFloatPriceDivisor() float32 {
	return message.Float32VLS(m.Unsafe(), 52, 48)
}

// IntToFloatPriceDivisor
func (m SecurityDefinitionResponseBuilder) IntToFloatPriceDivisor() float32 {
	return message.Float32VLS(m.Unsafe(), 52, 48)
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
func (m SecurityDefinitionResponse) StrikePrice() float32 {
	return message.Float32VLS(m.Unsafe(), 64, 60)
}

// StrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponseBuilder) StrikePrice() float32 {
	return message.Float32VLS(m.Unsafe(), 64, 60)
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
func (m SecurityDefinitionResponse) PutOrCall() PutCallEnum {
	return PutCallEnum(message.Uint8VLS(m.Unsafe(), 65, 64))
}

// PutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseBuilder) PutOrCall() PutCallEnum {
	return PutCallEnum(message.Uint8VLS(m.Unsafe(), 65, 64))
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
func (m SecurityDefinitionResponse) ShortInterest() uint32 {
	return message.Uint32VLS(m.Unsafe(), 72, 68)
}

// ShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseBuilder) ShortInterest() uint32 {
	return message.Uint32VLS(m.Unsafe(), 72, 68)
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
func (m SecurityDefinitionResponse) SecurityExpirationDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Unsafe(), 76, 72))
}

// SecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseBuilder) SecurityExpirationDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Unsafe(), 76, 72))
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
func (m SecurityDefinitionResponse) BuyRolloverInterest() float32 {
	return message.Float32VLS(m.Unsafe(), 80, 76)
}

// BuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseBuilder) BuyRolloverInterest() float32 {
	return message.Float32VLS(m.Unsafe(), 80, 76)
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
func (m SecurityDefinitionResponse) SellRolloverInterest() float32 {
	return message.Float32VLS(m.Unsafe(), 84, 80)
}

// SellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseBuilder) SellRolloverInterest() float32 {
	return message.Float32VLS(m.Unsafe(), 84, 80)
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
func (m SecurityDefinitionResponse) EarningsPerShare() float32 {
	return message.Float32VLS(m.Unsafe(), 88, 84)
}

// EarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseBuilder) EarningsPerShare() float32 {
	return message.Float32VLS(m.Unsafe(), 88, 84)
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
func (m SecurityDefinitionResponse) SharesOutstanding() uint32 {
	return message.Uint32VLS(m.Unsafe(), 92, 88)
}

// SharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseBuilder) SharesOutstanding() uint32 {
	return message.Uint32VLS(m.Unsafe(), 92, 88)
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
func (m SecurityDefinitionResponse) IntToFloatQuantityDivisor() float32 {
	return message.Float32VLS(m.Unsafe(), 96, 92)
}

// IntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseBuilder) IntToFloatQuantityDivisor() float32 {
	return message.Float32VLS(m.Unsafe(), 96, 92)
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
func (m SecurityDefinitionResponse) HasMarketDepthData() uint8 {
	return message.Uint8VLS(m.Unsafe(), 97, 96)
}

// HasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponseBuilder) HasMarketDepthData() uint8 {
	return message.Uint8VLS(m.Unsafe(), 97, 96)
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
func (m SecurityDefinitionResponse) Currency() string {
	return message.StringVLS(m.Unsafe(), 120, 116)
}

// Currency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponseBuilder) Currency() string {
	return message.StringVLS(m.Unsafe(), 120, 116)
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
func (m SecurityDefinitionResponse) OpenInterest() uint32 {
	return message.Uint32VLS(m.Unsafe(), 128, 124)
}

// OpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponseBuilder) OpenInterest() uint32 {
	return message.Uint32VLS(m.Unsafe(), 128, 124)
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
func (m SecurityDefinitionResponse) IsDelayed() bool {
	return message.BoolVLS(m.Unsafe(), 133, 132)
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseBuilder) IsDelayed() bool {
	return message.BoolVLS(m.Unsafe(), 133, 132)
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponsePointer) IsDelayed() bool {
	return message.BoolVLS(m.Ptr, 133, 132)
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponsePointerBuilder) IsDelayed() bool {
	return message.BoolVLS(m.Ptr, 133, 132)
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
func (m SecurityDefinitionResponseFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// Exchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponseFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
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
func (m SecurityDefinitionResponseFixed) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 92, 88))
}

// SecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseFixedBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 92, 88))
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
func (m SecurityDefinitionResponseFixed) Description() string {
	return message.StringFixed(m.Unsafe(), 156, 92)
}

// Description The text description for the Symbol.
func (m SecurityDefinitionResponseFixedBuilder) Description() string {
	return message.StringFixed(m.Unsafe(), 156, 92)
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
func (m SecurityDefinitionResponseFixed) MinPriceIncrement() float32 {
	return message.Float32Fixed(m.Unsafe(), 160, 156)
}

// MinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseFixedBuilder) MinPriceIncrement() float32 {
	return message.Float32Fixed(m.Unsafe(), 160, 156)
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
func (m SecurityDefinitionResponseFixed) PriceDisplayFormat() PriceDisplayFormatEnum {
	return PriceDisplayFormatEnum(message.Int32Fixed(m.Unsafe(), 164, 160))
}

// PriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseFixedBuilder) PriceDisplayFormat() PriceDisplayFormatEnum {
	return PriceDisplayFormatEnum(message.Int32Fixed(m.Unsafe(), 164, 160))
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
func (m SecurityDefinitionResponseFixed) CurrencyValuePerIncrement() float32 {
	return message.Float32Fixed(m.Unsafe(), 168, 164)
}

// CurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseFixedBuilder) CurrencyValuePerIncrement() float32 {
	return message.Float32Fixed(m.Unsafe(), 168, 164)
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
func (m SecurityDefinitionResponseFixed) IsFinalMessage() bool {
	return message.BoolFixed(m.Unsafe(), 169, 168)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseFixedBuilder) IsFinalMessage() bool {
	return message.BoolFixed(m.Unsafe(), 169, 168)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseFixedPointer) IsFinalMessage() bool {
	return message.BoolFixed(m.Ptr, 169, 168)
}

// IsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseFixedPointerBuilder) IsFinalMessage() bool {
	return message.BoolFixed(m.Ptr, 169, 168)
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
func (m SecurityDefinitionResponseFixed) IntToFloatPriceDivisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 180, 176)
}

// IntToFloatPriceDivisor
func (m SecurityDefinitionResponseFixedBuilder) IntToFloatPriceDivisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 180, 176)
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
func (m SecurityDefinitionResponseFixed) StrikePrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 220, 216)
}

// StrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponseFixedBuilder) StrikePrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 220, 216)
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
func (m SecurityDefinitionResponseFixed) PutOrCall() PutCallEnum {
	return PutCallEnum(message.Uint8Fixed(m.Unsafe(), 221, 220))
}

// PutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseFixedBuilder) PutOrCall() PutCallEnum {
	return PutCallEnum(message.Uint8Fixed(m.Unsafe(), 221, 220))
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
func (m SecurityDefinitionResponseFixed) ShortInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 228, 224)
}

// ShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseFixedBuilder) ShortInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 228, 224)
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
func (m SecurityDefinitionResponseFixed) SecurityExpirationDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 232, 228))
}

// SecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseFixedBuilder) SecurityExpirationDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 232, 228))
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
func (m SecurityDefinitionResponseFixed) BuyRolloverInterest() float32 {
	return message.Float32Fixed(m.Unsafe(), 236, 232)
}

// BuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedBuilder) BuyRolloverInterest() float32 {
	return message.Float32Fixed(m.Unsafe(), 236, 232)
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
func (m SecurityDefinitionResponseFixed) SellRolloverInterest() float32 {
	return message.Float32Fixed(m.Unsafe(), 240, 236)
}

// SellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedBuilder) SellRolloverInterest() float32 {
	return message.Float32Fixed(m.Unsafe(), 240, 236)
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
func (m SecurityDefinitionResponseFixed) EarningsPerShare() float32 {
	return message.Float32Fixed(m.Unsafe(), 244, 240)
}

// EarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseFixedBuilder) EarningsPerShare() float32 {
	return message.Float32Fixed(m.Unsafe(), 244, 240)
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
func (m SecurityDefinitionResponseFixed) SharesOutstanding() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 248, 244)
}

// SharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseFixedBuilder) SharesOutstanding() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 248, 244)
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
func (m SecurityDefinitionResponseFixed) IntToFloatQuantityDivisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 252, 248)
}

// IntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseFixedBuilder) IntToFloatQuantityDivisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 252, 248)
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
func (m SecurityDefinitionResponseFixed) HasMarketDepthData() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 253, 252)
}

// HasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponseFixedBuilder) HasMarketDepthData() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 253, 252)
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
func (m SecurityDefinitionResponseFixed) Currency() string {
	return message.StringFixed(m.Unsafe(), 340, 332)
}

// Currency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponseFixedBuilder) Currency() string {
	return message.StringFixed(m.Unsafe(), 340, 332)
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
func (m SecurityDefinitionResponseFixed) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 348, 344)
}

// OpenInterest In the case of when a Symbol is a contract type, this field is the number
// of outstanding contracts.
func (m SecurityDefinitionResponseFixedBuilder) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 348, 344)
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
func (m SecurityDefinitionResponseFixed) IsDelayed() bool {
	return message.BoolFixed(m.Unsafe(), 353, 352)
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseFixedBuilder) IsDelayed() bool {
	return message.BoolFixed(m.Unsafe(), 353, 352)
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseFixedPointer) IsDelayed() bool {
	return message.BoolFixed(m.Ptr, 353, 352)
}

// IsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseFixedPointerBuilder) IsDelayed() bool {
	return message.BoolFixed(m.Ptr, 353, 352)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
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
func (m *SecurityDefinitionResponseBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
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
func (m *SecurityDefinitionResponseBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetExchange This is the Exchange for the Symbol. This field is optional.
func (m *SecurityDefinitionResponsePointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetSecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, int32(value))
}

// SetSecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponsePointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Ptr, 24, 20, int32(value))
}

// SetDescription The text description for the Symbol.
func (m *SecurityDefinitionResponseBuilder) SetDescription(value string) {
	message.SetStringVLS(&m.VLSBuilder, 28, 24, value)
}

// SetDescription The text description for the Symbol.
func (m *SecurityDefinitionResponsePointerBuilder) SetDescription(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 28, 24, value)
}

// SetMinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseBuilder) SetMinPriceIncrement(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 32, 28, value)
}

// SetMinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponsePointerBuilder) SetMinPriceIncrement(value float32) {
	message.SetFloat32VLS(m.Ptr, 32, 28, value)
}

// SetPriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseBuilder) SetPriceDisplayFormat(value PriceDisplayFormatEnum) {
	message.SetInt32VLS(m.Unsafe(), 36, 32, int32(value))
}

// SetPriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponsePointerBuilder) SetPriceDisplayFormat(value PriceDisplayFormatEnum) {
	message.SetInt32VLS(m.Ptr, 36, 32, int32(value))
}

// SetCurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseBuilder) SetCurrencyValuePerIncrement(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 40, 36, value)
}

// SetCurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponsePointerBuilder) SetCurrencyValuePerIncrement(value float32) {
	message.SetFloat32VLS(m.Ptr, 40, 36, value)
}

// SetIsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseBuilder) SetIsFinalMessage(value bool) {
	message.SetBoolVLS(m.Unsafe(), 41, 40, value)
}

// SetIsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponsePointerBuilder) SetIsFinalMessage(value bool) {
	message.SetBoolVLS(m.Ptr, 41, 40, value)
}

// SetFloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponseBuilder) SetFloatToIntPriceMultiplier(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 48, 44, value)
}

// SetFloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponsePointerBuilder) SetFloatToIntPriceMultiplier(value float32) {
	message.SetFloat32VLS(m.Ptr, 48, 44, value)
}

// SetIntToFloatPriceDivisor
func (m SecurityDefinitionResponseBuilder) SetIntToFloatPriceDivisor(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 52, 48, value)
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
func (m *SecurityDefinitionResponseBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 56, 52, value)
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
func (m SecurityDefinitionResponseBuilder) SetUpdatesBidAskOnly(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 57, 56, value)
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
func (m SecurityDefinitionResponseBuilder) SetStrikePrice(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 64, 60, value)
}

// SetStrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponsePointerBuilder) SetStrikePrice(value float32) {
	message.SetFloat32VLS(m.Ptr, 64, 60, value)
}

// SetPutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseBuilder) SetPutOrCall(value PutCallEnum) {
	message.SetUint8VLS(m.Unsafe(), 65, 64, uint8(value))
}

// SetPutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponsePointerBuilder) SetPutOrCall(value PutCallEnum) {
	message.SetUint8VLS(m.Ptr, 65, 64, uint8(value))
}

// SetShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseBuilder) SetShortInterest(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 72, 68, value)
}

// SetShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponsePointerBuilder) SetShortInterest(value uint32) {
	message.SetUint32VLS(m.Ptr, 72, 68, value)
}

// SetSecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseBuilder) SetSecurityExpirationDate(value DateTime4Byte) {
	message.SetUint32VLS(m.Unsafe(), 76, 72, uint32(value))
}

// SetSecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponsePointerBuilder) SetSecurityExpirationDate(value DateTime4Byte) {
	message.SetUint32VLS(m.Ptr, 76, 72, uint32(value))
}

// SetBuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseBuilder) SetBuyRolloverInterest(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 80, 76, value)
}

// SetBuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponsePointerBuilder) SetBuyRolloverInterest(value float32) {
	message.SetFloat32VLS(m.Ptr, 80, 76, value)
}

// SetSellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseBuilder) SetSellRolloverInterest(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 84, 80, value)
}

// SetSellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponsePointerBuilder) SetSellRolloverInterest(value float32) {
	message.SetFloat32VLS(m.Ptr, 84, 80, value)
}

// SetEarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseBuilder) SetEarningsPerShare(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 88, 84, value)
}

// SetEarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponsePointerBuilder) SetEarningsPerShare(value float32) {
	message.SetFloat32VLS(m.Ptr, 88, 84, value)
}

// SetSharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseBuilder) SetSharesOutstanding(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 92, 88, value)
}

// SetSharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponsePointerBuilder) SetSharesOutstanding(value uint32) {
	message.SetUint32VLS(m.Ptr, 92, 88, value)
}

// SetIntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseBuilder) SetIntToFloatQuantityDivisor(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 96, 92, value)
}

// SetIntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponsePointerBuilder) SetIntToFloatQuantityDivisor(value float32) {
	message.SetFloat32VLS(m.Ptr, 96, 92, value)
}

// SetHasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponseBuilder) SetHasMarketDepthData(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 97, 96, value)
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
func (m SecurityDefinitionResponseBuilder) SetDisplayPriceMultiplier(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 104, 100, value)
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
func (m *SecurityDefinitionResponseBuilder) SetExchangeSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 108, 104, value)
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
func (m SecurityDefinitionResponseBuilder) SetInitialMarginRequirement(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 112, 108, value)
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
func (m SecurityDefinitionResponseBuilder) SetMaintenanceMarginRequirement(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 116, 112, value)
}

// SetMaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponsePointerBuilder) SetMaintenanceMarginRequirement(value float32) {
	message.SetFloat32VLS(m.Ptr, 116, 112, value)
}

// SetCurrency This is the currency that the Symbol trades in or is priced in.
func (m *SecurityDefinitionResponseBuilder) SetCurrency(value string) {
	message.SetStringVLS(&m.VLSBuilder, 120, 116, value)
}

// SetCurrency This is the currency that the Symbol trades in or is priced in.
func (m *SecurityDefinitionResponsePointerBuilder) SetCurrency(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 120, 116, value)
}

// SetContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponseBuilder) SetContractSize(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 124, 120, value)
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
func (m SecurityDefinitionResponseBuilder) SetOpenInterest(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 128, 124, value)
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
func (m SecurityDefinitionResponseBuilder) SetRolloverDate(value DateTime4Byte) {
	message.SetUint32VLS(m.Unsafe(), 132, 128, uint32(value))
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
func (m SecurityDefinitionResponseBuilder) SetIsDelayed(value bool) {
	message.SetBoolVLS(m.Unsafe(), 133, 132, value)
}

// SetIsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponsePointerBuilder) SetIsDelayed(value bool) {
	message.SetBoolVLS(m.Ptr, 133, 132, value)
}

// SetRequestID This is the same RequestID that this message is in response to and was
// given in the original Security Definition request message.
func (m SecurityDefinitionResponseFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
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
func (m SecurityDefinitionResponseFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
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
func (m SecurityDefinitionResponseFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 88, 72, value)
}

// SetExchange This is the Exchange for the Symbol. This field is optional.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 88, 72, value)
}

// SetSecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseFixedBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 92, 88, int32(value))
}

// SetSecurityType The Security Type for the symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 92, 88, int32(value))
}

// SetDescription The text description for the Symbol.
func (m SecurityDefinitionResponseFixedBuilder) SetDescription(value string) {
	message.SetStringFixed(m.Unsafe(), 156, 92, value)
}

// SetDescription The text description for the Symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetDescription(value string) {
	message.SetStringFixed(m.Ptr, 156, 92, value)
}

// SetMinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseFixedBuilder) SetMinPriceIncrement(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 160, 156, value)
}

// SetMinPriceIncrement The minimum amount that prices can change by for the Symbol and minimum
// amount that prices are quoted in.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetMinPriceIncrement(value float32) {
	message.SetFloat32Fixed(m.Ptr, 160, 156, value)
}

// SetPriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseFixedBuilder) SetPriceDisplayFormat(value PriceDisplayFormatEnum) {
	message.SetInt32Fixed(m.Unsafe(), 164, 160, int32(value))
}

// SetPriceDisplayFormat This field specifies the price formatting for display purposes.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetPriceDisplayFormat(value PriceDisplayFormatEnum) {
	message.SetInt32Fixed(m.Ptr, 164, 160, int32(value))
}

// SetCurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseFixedBuilder) SetCurrencyValuePerIncrement(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 168, 164, value)
}

// SetCurrencyValuePerIncrement This field is the currency value per MinPriceIncrement in the Symbols
// currency.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetCurrencyValuePerIncrement(value float32) {
	message.SetFloat32Fixed(m.Ptr, 168, 164, value)
}

// SetIsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseFixedBuilder) SetIsFinalMessage(value bool) {
	message.SetBoolFixed(m.Unsafe(), 169, 168, value)
}

// SetIsFinalMessage Set to a integer value of 1 to indicate this is the final message in batch
// of Security Definition Response messages.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetIsFinalMessage(value bool) {
	message.SetBoolFixed(m.Ptr, 169, 168, value)
}

// SetFloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponseFixedBuilder) SetFloatToIntPriceMultiplier(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 176, 172, value)
}

// SetFloatToIntPriceMultiplier With the integer order entry messages discontinued as of August 2020,
// this field is no longer relevant.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetFloatToIntPriceMultiplier(value float32) {
	message.SetFloat32Fixed(m.Ptr, 176, 172, value)
}

// SetIntToFloatPriceDivisor
func (m SecurityDefinitionResponseFixedBuilder) SetIntToFloatPriceDivisor(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 180, 176, value)
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
func (m SecurityDefinitionResponseFixedBuilder) SetUnderlyingSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 212, 180, value)
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
func (m SecurityDefinitionResponseFixedBuilder) SetUpdatesBidAskOnly(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 213, 212, value)
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
func (m SecurityDefinitionResponseFixedBuilder) SetStrikePrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 220, 216, value)
}

// SetStrikePrice The strike price when the Security Type is an option type.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetStrikePrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 220, 216, value)
}

// SetPutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseFixedBuilder) SetPutOrCall(value PutCallEnum) {
	message.SetUint8Fixed(m.Unsafe(), 221, 220, uint8(value))
}

// SetPutOrCall When the Security Type is an option, this indicates if it is a put or
// call.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetPutOrCall(value PutCallEnum) {
	message.SetUint8Fixed(m.Ptr, 221, 220, uint8(value))
}

// SetShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseFixedBuilder) SetShortInterest(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 228, 224, value)
}

// SetShortInterest The short interest when the Security Type is a stock.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetShortInterest(value uint32) {
	message.SetUint32Fixed(m.Ptr, 228, 224, value)
}

// SetSecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseFixedBuilder) SetSecurityExpirationDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 232, 228, uint32(value))
}

// SetSecurityExpirationDate The expiration date for the Symbol for symbols which have an expiration
// date.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetSecurityExpirationDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 232, 228, uint32(value))
}

// SetBuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedBuilder) SetBuyRolloverInterest(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 236, 232, value)
}

// SetBuyRolloverInterest The daily interest amount which is deducted for a Buy position. This only
// applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetBuyRolloverInterest(value float32) {
	message.SetFloat32Fixed(m.Ptr, 236, 232, value)
}

// SetSellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedBuilder) SetSellRolloverInterest(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 240, 236, value)
}

// SetSellRolloverInterest The daily interest amount which is deducted for a Sell position. This
// only applies for Forex trading. It is in the quote currency of the symbol.
// only applies for Forex trading. It is in the quote currency of the symbol.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetSellRolloverInterest(value float32) {
	message.SetFloat32Fixed(m.Ptr, 240, 236, value)
}

// SetEarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseFixedBuilder) SetEarningsPerShare(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 244, 240, value)
}

// SetEarningsPerShare The earnings per share as a currency value when the Security Type is a
// stock.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetEarningsPerShare(value float32) {
	message.SetFloat32Fixed(m.Ptr, 244, 240, value)
}

// SetSharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseFixedBuilder) SetSharesOutstanding(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 248, 244, value)
}

// SetSharesOutstanding This is the number of shares outstanding for stocks.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetSharesOutstanding(value uint32) {
	message.SetUint32Fixed(m.Ptr, 248, 244, value)
}

// SetIntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseFixedBuilder) SetIntToFloatQuantityDivisor(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 252, 248, value)
}

// SetIntToFloatQuantityDivisor With the integer order entry and market data messages discontinued as
// of August 2020, this field is no longer relevant.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetIntToFloatQuantityDivisor(value float32) {
	message.SetFloat32Fixed(m.Ptr, 252, 248, value)
}

// SetHasMarketDepthData When HasMarketDepthData is set to 1, it indicates the Symbol has market
// depth data available for it. When this is set to 0, market depth data
// is not supported for the Symbol.
func (m SecurityDefinitionResponseFixedBuilder) SetHasMarketDepthData(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 253, 252, value)
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
func (m SecurityDefinitionResponseFixedBuilder) SetDisplayPriceMultiplier(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 260, 256, value)
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
func (m SecurityDefinitionResponseFixedBuilder) SetExchangeSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 324, 260, value)
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
func (m SecurityDefinitionResponseFixedBuilder) SetInitialMarginRequirement(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 328, 324, value)
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
func (m SecurityDefinitionResponseFixedBuilder) SetMaintenanceMarginRequirement(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 332, 328, value)
}

// SetMaintenanceMarginRequirement This field applies to the Futures Security Type.
//
// This is the maintenance margin requirement as specified by the exchange,
// if available.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetMaintenanceMarginRequirement(value float32) {
	message.SetFloat32Fixed(m.Ptr, 332, 328, value)
}

// SetCurrency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponseFixedBuilder) SetCurrency(value string) {
	message.SetStringFixed(m.Unsafe(), 340, 332, value)
}

// SetCurrency This is the currency that the Symbol trades in or is priced in.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetCurrency(value string) {
	message.SetStringFixed(m.Ptr, 340, 332, value)
}

// SetContractSize In the case of when a Symbol is a contract type, this variable indicates
// the size of the contract.
//
// This is going to be an exchange specific specification.
func (m SecurityDefinitionResponseFixedBuilder) SetContractSize(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 344, 340, value)
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
func (m SecurityDefinitionResponseFixedBuilder) SetOpenInterest(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 348, 344, value)
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
func (m SecurityDefinitionResponseFixedBuilder) SetRolloverDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 352, 348, uint32(value))
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
func (m SecurityDefinitionResponseFixedBuilder) SetIsDelayed(value bool) {
	message.SetBoolFixed(m.Unsafe(), 353, 352, value)
}

// SetIsDelayed This field will be 1 if the market data for the Symbol is intentionally
// delayed by a certain amount of time. Otherwise, this will be 0.
func (m SecurityDefinitionResponseFixedPointerBuilder) SetIsDelayed(value bool) {
	message.SetBoolFixed(m.Ptr, 353, 352, value)
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
