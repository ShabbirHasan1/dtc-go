package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size               = MarketDataUpdateSessionLowSize  (24)
//     Type               = MARKET_DATA_UPDATE_SESSION_LOW  (115)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
var _MarketDataUpdateSessionLowDefault = []byte{24, 0, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataUpdateSessionLowSize = 24

// MarketDataUpdateSessionLow Sent by the Server to the Client to update the session Low as the Low
// price changes throughout the session.
type MarketDataUpdateSessionLow struct {
	message.Fixed
}

// MarketDataUpdateSessionLowBuilder Sent by the Server to the Client to update the session Low as the Low
// price changes throughout the session.
type MarketDataUpdateSessionLowBuilder struct {
	message.Fixed
}

func NewMarketDataUpdateSessionLowFrom(b []byte) MarketDataUpdateSessionLow {
	return MarketDataUpdateSessionLow{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateSessionLow(b []byte) MarketDataUpdateSessionLow {
	return MarketDataUpdateSessionLow{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateSessionLow() MarketDataUpdateSessionLowBuilder {
	a := MarketDataUpdateSessionLowBuilder{message.NewFixed(24)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateSessionLowDefault)
	return a
}

// Clear
// {
//     Size               = MarketDataUpdateSessionLowSize  (24)
//     Type               = MARKET_DATA_UPDATE_SESSION_LOW  (115)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
func (m MarketDataUpdateSessionLowBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateSessionLowDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionLow) ToBuilder() MarketDataUpdateSessionLowBuilder {
	return MarketDataUpdateSessionLowBuilder{m.Fixed}
}

// Finish
func (m MarketDataUpdateSessionLowBuilder) Finish() MarketDataUpdateSessionLow {
	return MarketDataUpdateSessionLow{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionLow) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionLowBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Price The session Low price.
func (m MarketDataUpdateSessionLow) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// Price The session Low price.
func (m MarketDataUpdateSessionLowBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// TradingSessionDate
func (m MarketDataUpdateSessionLow) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionLowBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionLowBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetPrice The session Low price.
func (m MarketDataUpdateSessionLowBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionLowBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 20, 16, uint32(value))
}

// Copy
func (m MarketDataUpdateSessionLow) Copy(to MarketDataUpdateSessionLowBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateSessionLowBuilder) Copy(to MarketDataUpdateSessionLowBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionLow) CopyPointer(to MarketDataUpdateSessionLowPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionLowBuilder) CopyPointer(to MarketDataUpdateSessionLowPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}
