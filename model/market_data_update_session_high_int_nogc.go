package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketDataUpdateSessionHigh_IntPointer struct {
	message.FixedPointer
}

type MarketDataUpdateSessionHigh_IntPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataUpdateSessionHigh_Int() MarketDataUpdateSessionHigh_IntPointerBuilder {
	a := MarketDataUpdateSessionHigh_IntPointerBuilder{message.AllocFixed(16)}
	a.Ptr.SetBytes(0, _MarketDataUpdateSessionHigh_IntDefault)
	return a
}

func AllocMarketDataUpdateSessionHigh_IntFrom(b []byte) MarketDataUpdateSessionHigh_IntPointer {
	return MarketDataUpdateSessionHigh_IntPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size               = MarketDataUpdateSessionHigh_IntSize  (16)
//     Type               = MARKET_DATA_UPDATE_SESSION_HIGH_INT  (129)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
func (m MarketDataUpdateSessionHigh_IntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateSessionHigh_IntDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionHigh_IntPointer) ToBuilder() MarketDataUpdateSessionHigh_IntPointerBuilder {
	return MarketDataUpdateSessionHigh_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateSessionHigh_IntPointerBuilder) Finish() MarketDataUpdateSessionHigh_IntPointer {
	return MarketDataUpdateSessionHigh_IntPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDataUpdateSessionHigh_IntPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketDataUpdateSessionHigh_IntPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Price
func (m MarketDataUpdateSessionHigh_IntPointer) Price() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// Price
func (m MarketDataUpdateSessionHigh_IntPointerBuilder) Price() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// TradingSessionDate
func (m MarketDataUpdateSessionHigh_IntPointer) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 16, 12))
}

// TradingSessionDate
func (m MarketDataUpdateSessionHigh_IntPointerBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 16, 12))
}

// SetSymbolID
func (m MarketDataUpdateSessionHigh_IntPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetPrice
func (m MarketDataUpdateSessionHigh_IntPointerBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionHigh_IntPointerBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 16, 12, uint32(value))
}

// Copy
func (m MarketDataUpdateSessionHigh_IntPointer) Copy(to MarketDataUpdateSessionHigh_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateSessionHigh_IntPointerBuilder) Copy(to MarketDataUpdateSessionHigh_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionHigh_IntPointer) CopyPointer(to MarketDataUpdateSessionHigh_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionHigh_IntPointerBuilder) CopyPointer(to MarketDataUpdateSessionHigh_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}
