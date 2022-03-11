package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size               = MarketDataUpdateSessionLow_IntSize  (16)
//     Type               = MARKET_DATA_UPDATE_SESSION_LOW_INT  (130)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
var _MarketDataUpdateSessionLow_IntDefault = []byte{16, 0, 130, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataUpdateSessionLow_IntSize = 16

type MarketDataUpdateSessionLow_Int struct {
	message.Fixed
}

type MarketDataUpdateSessionLow_IntBuilder struct {
	message.Fixed
}

func NewMarketDataUpdateSessionLow_IntFrom(b []byte) MarketDataUpdateSessionLow_Int {
	return MarketDataUpdateSessionLow_Int{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateSessionLow_Int(b []byte) MarketDataUpdateSessionLow_Int {
	return MarketDataUpdateSessionLow_Int{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateSessionLow_Int() MarketDataUpdateSessionLow_IntBuilder {
	a := MarketDataUpdateSessionLow_IntBuilder{message.NewFixed(16)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateSessionLow_IntDefault)
	return a
}

// Clear
// {
//     Size               = MarketDataUpdateSessionLow_IntSize  (16)
//     Type               = MARKET_DATA_UPDATE_SESSION_LOW_INT  (130)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
func (m MarketDataUpdateSessionLow_IntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateSessionLow_IntDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionLow_Int) ToBuilder() MarketDataUpdateSessionLow_IntBuilder {
	return MarketDataUpdateSessionLow_IntBuilder{m.Fixed}
}

// Finish
func (m MarketDataUpdateSessionLow_IntBuilder) Finish() MarketDataUpdateSessionLow_Int {
	return MarketDataUpdateSessionLow_Int{m.Fixed}
}

// SymbolID
func (m MarketDataUpdateSessionLow_Int) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDataUpdateSessionLow_IntBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Price
func (m MarketDataUpdateSessionLow_Int) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// Price
func (m MarketDataUpdateSessionLow_IntBuilder) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// TradingSessionDate
func (m MarketDataUpdateSessionLow_Int) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 16, 12))
}

// TradingSessionDate
func (m MarketDataUpdateSessionLow_IntBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 16, 12))
}

// SetSymbolID
func (m MarketDataUpdateSessionLow_IntBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetPrice
func (m MarketDataUpdateSessionLow_IntBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionLow_IntBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 16, 12, uint32(value))
}

// Copy
func (m MarketDataUpdateSessionLow_Int) Copy(to MarketDataUpdateSessionLow_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateSessionLow_IntBuilder) Copy(to MarketDataUpdateSessionLow_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionLow_Int) CopyPointer(to MarketDataUpdateSessionLow_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionLow_IntBuilder) CopyPointer(to MarketDataUpdateSessionLow_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}
