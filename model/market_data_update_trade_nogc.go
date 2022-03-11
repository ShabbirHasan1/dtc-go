package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDataUpdateTradePointer The Server sends this market data feed message to the Client when a trade
// occurs.
type MarketDataUpdateTradePointer struct {
	message.FixedPointer
}

// MarketDataUpdateTradePointerBuilder The Server sends this market data feed message to the Client when a trade
// occurs.
type MarketDataUpdateTradePointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataUpdateTrade() MarketDataUpdateTradePointerBuilder {
	a := MarketDataUpdateTradePointerBuilder{message.AllocFixed(40)}
	a.Ptr.SetBytes(0, _MarketDataUpdateTradeDefault)
	return a
}

func AllocMarketDataUpdateTradeFrom(b []byte) MarketDataUpdateTradePointer {
	return MarketDataUpdateTradePointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = MarketDataUpdateTradeSize  (40)
//     Type       = MARKET_DATA_UPDATE_TRADE  (107)
//     SymbolID   = 0
//     AtBidOrAsk = 0
//     Price      = 0
//     Volume     = 0
//     DateTime   = 0
// }
func (m MarketDataUpdateTradePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateTradeDefault)
}

// ToBuilder
func (m MarketDataUpdateTradePointer) ToBuilder() MarketDataUpdateTradePointerBuilder {
	return MarketDataUpdateTradePointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateTradePointerBuilder) Finish() MarketDataUpdateTradePointer {
	return MarketDataUpdateTradePointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradePointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradePointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// AtBidOrAsk Indicator whether the trade occurred at the bid or ask.
func (m MarketDataUpdateTradePointer) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// AtBidOrAsk Indicator whether the trade occurred at the bid or ask.
func (m MarketDataUpdateTradePointerBuilder) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// Price The price of the trade.
func (m MarketDataUpdateTradePointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// Price The price of the trade.
func (m MarketDataUpdateTradePointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradePointer) Volume() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradePointerBuilder) Volume() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
}

// DateTime The Date-Time of the trade.
func (m MarketDataUpdateTradePointer) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 40, 32))
}

// DateTime The Date-Time of the trade.
func (m MarketDataUpdateTradePointerBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 40, 32))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradePointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetAtBidOrAsk Indicator whether the trade occurred at the bid or ask.
func (m MarketDataUpdateTradePointerBuilder) SetAtBidOrAsk(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Ptr, 10, 8, uint16(value))
}

// SetPrice The price of the trade.
func (m MarketDataUpdateTradePointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetVolume The volume of the trade.
func (m MarketDataUpdateTradePointerBuilder) SetVolume(value float64) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, value)
}

// SetDateTime The Date-Time of the trade.
func (m MarketDataUpdateTradePointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 40, 32, float64(value))
}

// Copy
func (m MarketDataUpdateTradePointer) Copy(to MarketDataUpdateTradeBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateTradePointerBuilder) Copy(to MarketDataUpdateTradeBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateTradePointer) CopyPointer(to MarketDataUpdateTradePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateTradePointerBuilder) CopyPointer(to MarketDataUpdateTradePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}
