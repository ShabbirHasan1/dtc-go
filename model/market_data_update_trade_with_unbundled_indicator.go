package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                    = MarketDataUpdateTradeWithUnbundledIndicatorSize  (40)
//     Type                    = MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR  (137)
//     SymbolID                = 0
//     AtBidOrAsk              = 0
//     UnbundledTradeIndicator = UNBUNDLED_TRADE_NONE  (0)
//     TradeCondition          = 0
//     Reserve_1               = 0
//     Reserve_2               = 0
//     Price                   = 0
//     Volume                  = 0
//     Reserve_3               = 0
//     DateTime                = 0
// }
var _MarketDataUpdateTradeWithUnbundledIndicatorDefault = []byte{40, 0, 137, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataUpdateTradeWithUnbundledIndicatorSize = 40

type MarketDataUpdateTradeWithUnbundledIndicator struct {
	message.Fixed
}

type MarketDataUpdateTradeWithUnbundledIndicatorBuilder struct {
	message.Fixed
}

func NewMarketDataUpdateTradeWithUnbundledIndicatorFrom(b []byte) MarketDataUpdateTradeWithUnbundledIndicator {
	return MarketDataUpdateTradeWithUnbundledIndicator{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateTradeWithUnbundledIndicator(b []byte) MarketDataUpdateTradeWithUnbundledIndicator {
	return MarketDataUpdateTradeWithUnbundledIndicator{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateTradeWithUnbundledIndicator() MarketDataUpdateTradeWithUnbundledIndicatorBuilder {
	a := MarketDataUpdateTradeWithUnbundledIndicatorBuilder{message.NewFixed(40)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateTradeWithUnbundledIndicatorDefault)
	return a
}

// Clear
// {
//     Size                    = MarketDataUpdateTradeWithUnbundledIndicatorSize  (40)
//     Type                    = MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR  (137)
//     SymbolID                = 0
//     AtBidOrAsk              = 0
//     UnbundledTradeIndicator = UNBUNDLED_TRADE_NONE  (0)
//     TradeCondition          = 0
//     Reserve_1               = 0
//     Reserve_2               = 0
//     Price                   = 0
//     Volume                  = 0
//     Reserve_3               = 0
//     DateTime                = 0
// }
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateTradeWithUnbundledIndicatorDefault)
}

// ToBuilder
func (m MarketDataUpdateTradeWithUnbundledIndicator) ToBuilder() MarketDataUpdateTradeWithUnbundledIndicatorBuilder {
	return MarketDataUpdateTradeWithUnbundledIndicatorBuilder{m.Fixed}
}

// Finish
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) Finish() MarketDataUpdateTradeWithUnbundledIndicator {
	return MarketDataUpdateTradeWithUnbundledIndicator{m.Fixed}
}

// SymbolID
func (m MarketDataUpdateTradeWithUnbundledIndicator) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// AtBidOrAsk
func (m MarketDataUpdateTradeWithUnbundledIndicator) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Unsafe(), 9, 8))
}

// AtBidOrAsk
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Unsafe(), 9, 8))
}

// UnbundledTradeIndicator
func (m MarketDataUpdateTradeWithUnbundledIndicator) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(message.Int8Fixed(m.Unsafe(), 10, 9))
}

// UnbundledTradeIndicator
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(message.Int8Fixed(m.Unsafe(), 10, 9))
}

// TradeCondition
func (m MarketDataUpdateTradeWithUnbundledIndicator) TradeCondition() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 11, 10)
}

// TradeCondition
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) TradeCondition() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 11, 10)
}

// Reserve_1
func (m MarketDataUpdateTradeWithUnbundledIndicator) Reserve_1() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 12, 11)
}

// Reserve_1
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) Reserve_1() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 12, 11)
}

// Reserve_2
func (m MarketDataUpdateTradeWithUnbundledIndicator) Reserve_2() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 16, 12)
}

// Reserve_2
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) Reserve_2() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 16, 12)
}

// Price
func (m MarketDataUpdateTradeWithUnbundledIndicator) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// Price
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// Volume
func (m MarketDataUpdateTradeWithUnbundledIndicator) Volume() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 28, 24)
}

// Volume
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) Volume() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 28, 24)
}

// Reserve_3
func (m MarketDataUpdateTradeWithUnbundledIndicator) Reserve_3() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 32, 28)
}

// Reserve_3
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) Reserve_3() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 32, 28)
}

// DateTime
func (m MarketDataUpdateTradeWithUnbundledIndicator) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 40, 32))
}

// DateTime
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 40, 32))
}

// SetSymbolID
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetAtBidOrAsk
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) SetAtBidOrAsk(value AtBidOrAskEnum8) {
	message.SetUint8Fixed(m.Unsafe(), 9, 8, uint8(value))
}

// SetUnbundledTradeIndicator
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) SetUnbundledTradeIndicator(value UnbundledTradeIndicatorEnum) {
	message.SetInt8Fixed(m.Unsafe(), 10, 9, int8(value))
}

// SetTradeCondition
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) SetTradeCondition(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 11, 10, value)
}

// SetReserve_1
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) SetReserve_1(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 12, 11, value)
}

// SetReserve_2
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) SetReserve_2(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 16, 12, value)
}

// SetPrice
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetVolume
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) SetVolume(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 28, 24, value)
}

// SetReserve_3
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) SetReserve_3(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 32, 28, value)
}

// SetDateTime
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 40, 32, float64(value))
}

// Copy
func (m MarketDataUpdateTradeWithUnbundledIndicator) Copy(to MarketDataUpdateTradeWithUnbundledIndicatorBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
	to.SetReserve_1(m.Reserve_1())
	to.SetReserve_2(m.Reserve_2())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetReserve_3(m.Reserve_3())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) Copy(to MarketDataUpdateTradeWithUnbundledIndicatorBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
	to.SetReserve_1(m.Reserve_1())
	to.SetReserve_2(m.Reserve_2())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetReserve_3(m.Reserve_3())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateTradeWithUnbundledIndicator) CopyPointer(to MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
	to.SetReserve_1(m.Reserve_1())
	to.SetReserve_2(m.Reserve_2())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetReserve_3(m.Reserve_3())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) CopyPointer(to MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
	to.SetReserve_1(m.Reserve_1())
	to.SetReserve_2(m.Reserve_2())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetReserve_3(m.Reserve_3())
	to.SetDateTime(m.DateTime())
}
