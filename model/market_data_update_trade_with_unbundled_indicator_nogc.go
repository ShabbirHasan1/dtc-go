package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketDataUpdateTradeWithUnbundledIndicatorPointer struct {
	message.FixedPointer
}

type MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataUpdateTradeWithUnbundledIndicator() MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder {
	a := MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder{message.AllocFixed(40)}
	a.Ptr.SetBytes(0, _MarketDataUpdateTradeWithUnbundledIndicatorDefault)
	return a
}

func AllocMarketDataUpdateTradeWithUnbundledIndicatorFrom(b []byte) MarketDataUpdateTradeWithUnbundledIndicatorPointer {
	return MarketDataUpdateTradeWithUnbundledIndicatorPointer{FixedPointer: message.AllocFixedFrom(b)}
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
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateTradeWithUnbundledIndicatorDefault)
}

// ToBuilder
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) ToBuilder() MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder {
	return MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) Finish() MarketDataUpdateTradeWithUnbundledIndicatorPointer {
	return MarketDataUpdateTradeWithUnbundledIndicatorPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// AtBidOrAsk
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Ptr, 9, 8))
}

// AtBidOrAsk
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Ptr, 9, 8))
}

// UnbundledTradeIndicator
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(message.Int8Fixed(m.Ptr, 10, 9))
}

// UnbundledTradeIndicator
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(message.Int8Fixed(m.Ptr, 10, 9))
}

// TradeCondition
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) TradeCondition() uint8 {
	return message.Uint8Fixed(m.Ptr, 11, 10)
}

// TradeCondition
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) TradeCondition() uint8 {
	return message.Uint8Fixed(m.Ptr, 11, 10)
}

// Reserve_1
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) Reserve_1() uint8 {
	return message.Uint8Fixed(m.Ptr, 12, 11)
}

// Reserve_1
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) Reserve_1() uint8 {
	return message.Uint8Fixed(m.Ptr, 12, 11)
}

// Reserve_2
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) Reserve_2() uint32 {
	return message.Uint32Fixed(m.Ptr, 16, 12)
}

// Reserve_2
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) Reserve_2() uint32 {
	return message.Uint32Fixed(m.Ptr, 16, 12)
}

// Price
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// Price
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// Volume
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) Volume() uint32 {
	return message.Uint32Fixed(m.Ptr, 28, 24)
}

// Volume
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) Volume() uint32 {
	return message.Uint32Fixed(m.Ptr, 28, 24)
}

// Reserve_3
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) Reserve_3() uint32 {
	return message.Uint32Fixed(m.Ptr, 32, 28)
}

// Reserve_3
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) Reserve_3() uint32 {
	return message.Uint32Fixed(m.Ptr, 32, 28)
}

// DateTime
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 40, 32))
}

// DateTime
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 40, 32))
}

// SetSymbolID
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetAtBidOrAsk
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) SetAtBidOrAsk(value AtBidOrAskEnum8) {
	message.SetUint8Fixed(m.Ptr, 9, 8, uint8(value))
}

// SetUnbundledTradeIndicator
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) SetUnbundledTradeIndicator(value UnbundledTradeIndicatorEnum) {
	message.SetInt8Fixed(m.Ptr, 10, 9, int8(value))
}

// SetTradeCondition
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) SetTradeCondition(value uint8) {
	message.SetUint8Fixed(m.Ptr, 11, 10, value)
}

// SetReserve_1
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) SetReserve_1(value uint8) {
	message.SetUint8Fixed(m.Ptr, 12, 11, value)
}

// SetReserve_2
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) SetReserve_2(value uint32) {
	message.SetUint32Fixed(m.Ptr, 16, 12, value)
}

// SetPrice
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetVolume
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) SetVolume(value uint32) {
	message.SetUint32Fixed(m.Ptr, 28, 24, value)
}

// SetReserve_3
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) SetReserve_3(value uint32) {
	message.SetUint32Fixed(m.Ptr, 32, 28, value)
}

// SetDateTime
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 40, 32, float64(value))
}

// Copy
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) Copy(to MarketDataUpdateTradeWithUnbundledIndicatorBuilder) {
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
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) Copy(to MarketDataUpdateTradeWithUnbundledIndicatorBuilder) {
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
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) CopyPointer(to MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) {
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
func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) CopyPointer(to MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) {
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
