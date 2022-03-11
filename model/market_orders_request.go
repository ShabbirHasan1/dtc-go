package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketOrdersRequestSize = 28

const MarketOrdersRequestFixedSize = 96

// {
//     Size                           = MarketOrdersRequestSize  (28)
//     Type                           = MARKET_ORDERS_REQUEST  (150)
//     BaseSize                       = MarketOrdersRequestSize  (28)
//     RequestAction                  = 0
//     SymbolID                       = 0
//     Symbol                         = ""
//     Exchange                       = ""
//     SendQuantitiesGreaterOrEqualTo = 0
// }
var _MarketOrdersRequestDefault = []byte{28, 0, 150, 0, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size                           = MarketOrdersRequestFixedSize  (96)
//     Type                           = MARKET_ORDERS_REQUEST  (150)
//     RequestAction                  = SUBSCRIBE  (1)
//     SymbolID                       = 0
//     Symbol                         = ""
//     Exchange                       = ""
//     SendQuantitiesGreaterOrEqualTo = 0
// }
var _MarketOrdersRequestFixedDefault = []byte{96, 0, 150, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketOrdersRequest struct {
	message.VLS
}

type MarketOrdersRequestBuilder struct {
	message.VLSBuilder
}

type MarketOrdersRequestFixed struct {
	message.Fixed
}

type MarketOrdersRequestFixedBuilder struct {
	message.Fixed
}

type MarketOrdersRequestPointer struct {
	message.VLSPointer
}

type MarketOrdersRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

type MarketOrdersRequestFixedPointer struct {
	message.FixedPointer
}

type MarketOrdersRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func NewMarketOrdersRequestFrom(b []byte) MarketOrdersRequest {
	return MarketOrdersRequest{VLS: message.NewVLSFrom(b)}
}

func WrapMarketOrdersRequest(b []byte) MarketOrdersRequest {
	return MarketOrdersRequest{VLS: message.WrapVLS(b)}
}

func NewMarketOrdersRequest() MarketOrdersRequestBuilder {
	a := MarketOrdersRequestBuilder{message.NewVLSBuilder(28)}
	a.Unsafe().SetBytes(0, _MarketOrdersRequestDefault)
	return a
}

func NewMarketOrdersRequestFixedFrom(b []byte) MarketOrdersRequestFixed {
	return MarketOrdersRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketOrdersRequestFixed(b []byte) MarketOrdersRequestFixed {
	return MarketOrdersRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewMarketOrdersRequestFixed() MarketOrdersRequestFixedBuilder {
	a := MarketOrdersRequestFixedBuilder{message.NewFixed(96)}
	a.Unsafe().SetBytes(0, _MarketOrdersRequestFixedDefault)
	return a
}

func AllocMarketOrdersRequest() MarketOrdersRequestPointerBuilder {
	a := MarketOrdersRequestPointerBuilder{message.AllocVLSBuilder(28)}
	a.Ptr.SetBytes(0, _MarketOrdersRequestDefault)
	return a
}

func AllocMarketOrdersRequestFrom(b []byte) MarketOrdersRequestPointer {
	return MarketOrdersRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocMarketOrdersRequestFixed() MarketOrdersRequestFixedPointerBuilder {
	a := MarketOrdersRequestFixedPointerBuilder{message.AllocFixed(96)}
	a.Ptr.SetBytes(0, _MarketOrdersRequestFixedDefault)
	return a
}

func AllocMarketOrdersRequestFixedFrom(b []byte) MarketOrdersRequestFixedPointer {
	return MarketOrdersRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                           = MarketOrdersRequestSize  (28)
//     Type                           = MARKET_ORDERS_REQUEST  (150)
//     BaseSize                       = MarketOrdersRequestSize  (28)
//     RequestAction                  = 0
//     SymbolID                       = 0
//     Symbol                         = ""
//     Exchange                       = ""
//     SendQuantitiesGreaterOrEqualTo = 0
// }
func (m MarketOrdersRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersRequestDefault)
}

// Clear
// {
//     Size                           = MarketOrdersRequestFixedSize  (96)
//     Type                           = MARKET_ORDERS_REQUEST  (150)
//     RequestAction                  = SUBSCRIBE  (1)
//     SymbolID                       = 0
//     Symbol                         = ""
//     Exchange                       = ""
//     SendQuantitiesGreaterOrEqualTo = 0
// }
func (m MarketOrdersRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersRequestFixedDefault)
}

// Clear
// {
//     Size                           = MarketOrdersRequestSize  (28)
//     Type                           = MARKET_ORDERS_REQUEST  (150)
//     BaseSize                       = MarketOrdersRequestSize  (28)
//     RequestAction                  = 0
//     SymbolID                       = 0
//     Symbol                         = ""
//     Exchange                       = ""
//     SendQuantitiesGreaterOrEqualTo = 0
// }
func (m MarketOrdersRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersRequestDefault)
}

// Clear
// {
//     Size                           = MarketOrdersRequestFixedSize  (96)
//     Type                           = MARKET_ORDERS_REQUEST  (150)
//     RequestAction                  = SUBSCRIBE  (1)
//     SymbolID                       = 0
//     Symbol                         = ""
//     Exchange                       = ""
//     SendQuantitiesGreaterOrEqualTo = 0
// }
func (m MarketOrdersRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersRequestFixedDefault)
}

// ToBuilder
func (m MarketOrdersRequest) ToBuilder() MarketOrdersRequestBuilder {
	return MarketOrdersRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m MarketOrdersRequestFixed) ToBuilder() MarketOrdersRequestFixedBuilder {
	return MarketOrdersRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m MarketOrdersRequestPointer) ToBuilder() MarketOrdersRequestPointerBuilder {
	return MarketOrdersRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m MarketOrdersRequestFixedPointer) ToBuilder() MarketOrdersRequestFixedPointerBuilder {
	return MarketOrdersRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketOrdersRequestBuilder) Finish() MarketOrdersRequest {
	return MarketOrdersRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m MarketOrdersRequestFixedBuilder) Finish() MarketOrdersRequestFixed {
	return MarketOrdersRequestFixed{m.Fixed}
}

// Finish
func (m *MarketOrdersRequestPointerBuilder) Finish() MarketOrdersRequestPointer {
	return MarketOrdersRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *MarketOrdersRequestFixedPointerBuilder) Finish() MarketOrdersRequestFixedPointer {
	return MarketOrdersRequestFixedPointer{m.FixedPointer}
}

// RequestAction
func (m MarketOrdersRequest) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Unsafe(), 12, 8))
}

// RequestAction
func (m MarketOrdersRequestBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Unsafe(), 12, 8))
}

// RequestAction
func (m MarketOrdersRequestPointer) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Ptr, 12, 8))
}

// RequestAction
func (m MarketOrdersRequestPointerBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Ptr, 12, 8))
}

// SymbolID
func (m MarketOrdersRequest) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 16, 12)
}

// SymbolID
func (m MarketOrdersRequestBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 16, 12)
}

// SymbolID
func (m MarketOrdersRequestPointer) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 16, 12)
}

// SymbolID
func (m MarketOrdersRequestPointerBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 16, 12)
}

// Symbol
func (m MarketOrdersRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Symbol
func (m MarketOrdersRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Symbol
func (m MarketOrdersRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Symbol
func (m MarketOrdersRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Exchange
func (m MarketOrdersRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Exchange
func (m MarketOrdersRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Exchange
func (m MarketOrdersRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Exchange
func (m MarketOrdersRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequest) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32VLS(m.Unsafe(), 28, 24)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestBuilder) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32VLS(m.Unsafe(), 28, 24)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestPointer) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32VLS(m.Ptr, 28, 24)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestPointerBuilder) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32VLS(m.Ptr, 28, 24)
}

// RequestAction
func (m MarketOrdersRequestFixed) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Unsafe(), 8, 4))
}

// RequestAction
func (m MarketOrdersRequestFixedBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Unsafe(), 8, 4))
}

// RequestAction
func (m MarketOrdersRequestFixedPointer) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Ptr, 8, 4))
}

// RequestAction
func (m MarketOrdersRequestFixedPointerBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Ptr, 8, 4))
}

// SymbolID
func (m MarketOrdersRequestFixed) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// SymbolID
func (m MarketOrdersRequestFixedBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// SymbolID
func (m MarketOrdersRequestFixedPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// SymbolID
func (m MarketOrdersRequestFixedPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// Symbol
func (m MarketOrdersRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 76, 12)
}

// Symbol
func (m MarketOrdersRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 76, 12)
}

// Symbol
func (m MarketOrdersRequestFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 76, 12)
}

// Symbol
func (m MarketOrdersRequestFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 76, 12)
}

// Exchange
func (m MarketOrdersRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 92, 76)
}

// Exchange
func (m MarketOrdersRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 92, 76)
}

// Exchange
func (m MarketOrdersRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 92, 76)
}

// Exchange
func (m MarketOrdersRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 92, 76)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixed) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 96, 92)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixedBuilder) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 96, 92)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixedPointer) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32Fixed(m.Ptr, 96, 92)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixedPointerBuilder) SendQuantitiesGreaterOrEqualTo() uint32 {
	return message.Uint32Fixed(m.Ptr, 96, 92)
}

// SetRequestAction
func (m MarketOrdersRequestBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, int32(value))
}

// SetRequestAction
func (m MarketOrdersRequestPointerBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Ptr, 12, 8, int32(value))
}

// SetSymbolID
func (m MarketOrdersRequestBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 16, 12, value)
}

// SetSymbolID
func (m MarketOrdersRequestPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Ptr, 16, 12, value)
}

// SetSymbol
func (m *MarketOrdersRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetSymbol
func (m *MarketOrdersRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetExchange
func (m *MarketOrdersRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetExchange
func (m *MarketOrdersRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetSendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestBuilder) SetSendQuantitiesGreaterOrEqualTo(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 28, 24, value)
}

// SetSendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestPointerBuilder) SetSendQuantitiesGreaterOrEqualTo(value uint32) {
	message.SetUint32VLS(m.Ptr, 28, 24, value)
}

// SetRequestAction
func (m MarketOrdersRequestFixedBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, int32(value))
}

// SetRequestAction
func (m MarketOrdersRequestFixedPointerBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32Fixed(m.Ptr, 8, 4, int32(value))
}

// SetSymbolID
func (m MarketOrdersRequestFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 12, 8, value)
}

// SetSymbolID
func (m MarketOrdersRequestFixedPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 12, 8, value)
}

// SetSymbol
func (m MarketOrdersRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 76, 12, value)
}

// SetSymbol
func (m MarketOrdersRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 76, 12, value)
}

// SetExchange
func (m MarketOrdersRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 92, 76, value)
}

// SetExchange
func (m MarketOrdersRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 92, 76, value)
}

// SetSendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixedBuilder) SetSendQuantitiesGreaterOrEqualTo(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 96, 92, value)
}

// SetSendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixedPointerBuilder) SetSendQuantitiesGreaterOrEqualTo(value uint32) {
	message.SetUint32Fixed(m.Ptr, 96, 92, value)
}

// Copy
func (m MarketOrdersRequest) Copy(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// Copy
func (m MarketOrdersRequestBuilder) Copy(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequest) CopyPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestBuilder) CopyPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequest) CopyToPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestBuilder) CopyToPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequest) CopyTo(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestBuilder) CopyTo(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// Copy
func (m MarketOrdersRequestFixed) Copy(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// Copy
func (m MarketOrdersRequestFixedBuilder) Copy(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestFixed) CopyPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestFixedBuilder) CopyPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestFixed) CopyToPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestFixedBuilder) CopyToPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestFixed) CopyTo(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestFixedBuilder) CopyTo(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// Copy
func (m MarketOrdersRequestPointer) Copy(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// Copy
func (m MarketOrdersRequestPointerBuilder) Copy(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestPointer) CopyPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestPointerBuilder) CopyPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestPointer) CopyTo(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestPointerBuilder) CopyTo(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestPointer) CopyToPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestPointerBuilder) CopyToPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// Copy
func (m MarketOrdersRequestFixedPointer) Copy(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// Copy
func (m MarketOrdersRequestFixedPointerBuilder) Copy(to MarketOrdersRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestFixedPointer) CopyPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyPointer
func (m MarketOrdersRequestFixedPointerBuilder) CopyPointer(to MarketOrdersRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestFixedPointer) CopyTo(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestFixedPointerBuilder) CopyTo(to MarketOrdersRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestFixedPointer) CopyToPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyToPointer
func (m MarketOrdersRequestFixedPointerBuilder) CopyToPointer(to MarketOrdersRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}
