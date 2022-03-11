package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size               = MarketDataUpdateSessionOpen_IntSize  (16)
//     Type               = MARKET_DATA_UPDATE_SESSION_OPEN_INT  (128)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
var _MarketDataUpdateSessionOpen_IntDefault = []byte{16, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataUpdateSessionOpen_IntSize = 16

type MarketDataUpdateSessionOpen_Int struct {
	message.Fixed
}

type MarketDataUpdateSessionOpen_IntBuilder struct {
	message.Fixed
}

func NewMarketDataUpdateSessionOpen_IntFrom(b []byte) MarketDataUpdateSessionOpen_Int {
	return MarketDataUpdateSessionOpen_Int{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateSessionOpen_Int(b []byte) MarketDataUpdateSessionOpen_Int {
	return MarketDataUpdateSessionOpen_Int{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateSessionOpen_Int() MarketDataUpdateSessionOpen_IntBuilder {
	a := MarketDataUpdateSessionOpen_IntBuilder{message.NewFixed(16)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateSessionOpen_IntDefault)
	return a
}

// Clear
// {
//     Size               = MarketDataUpdateSessionOpen_IntSize  (16)
//     Type               = MARKET_DATA_UPDATE_SESSION_OPEN_INT  (128)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
func (m MarketDataUpdateSessionOpen_IntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateSessionOpen_IntDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionOpen_Int) ToBuilder() MarketDataUpdateSessionOpen_IntBuilder {
	return MarketDataUpdateSessionOpen_IntBuilder{m.Fixed}
}

// Finish
func (m MarketDataUpdateSessionOpen_IntBuilder) Finish() MarketDataUpdateSessionOpen_Int {
	return MarketDataUpdateSessionOpen_Int{m.Fixed}
}

// SymbolID
func (m MarketDataUpdateSessionOpen_Int) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDataUpdateSessionOpen_IntBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Price
func (m MarketDataUpdateSessionOpen_Int) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// Price
func (m MarketDataUpdateSessionOpen_IntBuilder) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// TradingSessionDate
func (m MarketDataUpdateSessionOpen_Int) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 16, 12))
}

// TradingSessionDate
func (m MarketDataUpdateSessionOpen_IntBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 16, 12))
}

// SetSymbolID
func (m MarketDataUpdateSessionOpen_IntBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetPrice
func (m MarketDataUpdateSessionOpen_IntBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionOpen_IntBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 16, 12, uint32(value))
}

// Copy
func (m MarketDataUpdateSessionOpen_Int) Copy(to MarketDataUpdateSessionOpen_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateSessionOpen_IntBuilder) Copy(to MarketDataUpdateSessionOpen_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionOpen_Int) CopyPointer(to MarketDataUpdateSessionOpen_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionOpen_IntBuilder) CopyPointer(to MarketDataUpdateSessionOpen_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}
