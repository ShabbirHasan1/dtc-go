package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketDataUpdateTrade_IntPointer struct {
	message.FixedPointer
}

type MarketDataUpdateTrade_IntPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataUpdateTrade_Int() MarketDataUpdateTrade_IntPointerBuilder {
	a := MarketDataUpdateTrade_IntPointerBuilder{message.AllocFixed(32)}
	a.Ptr.SetBytes(0, _MarketDataUpdateTrade_IntDefault)
	return a
}

func AllocMarketDataUpdateTrade_IntFrom(b []byte) MarketDataUpdateTrade_IntPointer {
	return MarketDataUpdateTrade_IntPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = MarketDataUpdateTrade_IntSize  (32)
//     Type       = MARKET_DATA_UPDATE_TRADE_INT  (126)
//     SymbolID   = 0
//     AtBidOrAsk = 0
//     Price      = 0
//     Volume     = 0
//     DateTime   = 0
// }
func (m MarketDataUpdateTrade_IntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateTrade_IntDefault)
}

// ToBuilder
func (m MarketDataUpdateTrade_IntPointer) ToBuilder() MarketDataUpdateTrade_IntPointerBuilder {
	return MarketDataUpdateTrade_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateTrade_IntPointerBuilder) Finish() MarketDataUpdateTrade_IntPointer {
	return MarketDataUpdateTrade_IntPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDataUpdateTrade_IntPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketDataUpdateTrade_IntPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// AtBidOrAsk
func (m MarketDataUpdateTrade_IntPointer) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// AtBidOrAsk
func (m MarketDataUpdateTrade_IntPointerBuilder) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// Price
func (m MarketDataUpdateTrade_IntPointer) Price() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// Price
func (m MarketDataUpdateTrade_IntPointerBuilder) Price() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// Volume
func (m MarketDataUpdateTrade_IntPointer) Volume() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// Volume
func (m MarketDataUpdateTrade_IntPointerBuilder) Volume() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// DateTime
func (m MarketDataUpdateTrade_IntPointer) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 32, 24))
}

// DateTime
func (m MarketDataUpdateTrade_IntPointerBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 32, 24))
}

// SetSymbolID
func (m MarketDataUpdateTrade_IntPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetAtBidOrAsk
func (m MarketDataUpdateTrade_IntPointerBuilder) SetAtBidOrAsk(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Ptr, 10, 8, uint16(value))
}

// SetPrice
func (m MarketDataUpdateTrade_IntPointerBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 16, 12, value)
}

// SetVolume
func (m MarketDataUpdateTrade_IntPointerBuilder) SetVolume(value int32) {
	message.SetInt32Fixed(m.Ptr, 20, 16, value)
}

// SetDateTime
func (m MarketDataUpdateTrade_IntPointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, float64(value))
}

// Copy
func (m MarketDataUpdateTrade_IntPointer) Copy(to MarketDataUpdateTrade_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateTrade_IntPointerBuilder) Copy(to MarketDataUpdateTrade_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateTrade_IntPointer) CopyPointer(to MarketDataUpdateTrade_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateTrade_IntPointerBuilder) CopyPointer(to MarketDataUpdateTrade_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}
