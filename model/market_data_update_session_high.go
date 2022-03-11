package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size               = MarketDataUpdateSessionHighSize  (24)
//     Type               = MARKET_DATA_UPDATE_SESSION_HIGH  (114)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
var _MarketDataUpdateSessionHighDefault = []byte{24, 0, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataUpdateSessionHighSize = 24

// MarketDataUpdateSessionHigh Sent by the Server to the Client to update the session High as the High
// price changes throughout the session.
type MarketDataUpdateSessionHigh struct {
	message.Fixed
}

// MarketDataUpdateSessionHighBuilder Sent by the Server to the Client to update the session High as the High
// price changes throughout the session.
type MarketDataUpdateSessionHighBuilder struct {
	message.Fixed
}

func NewMarketDataUpdateSessionHighFrom(b []byte) MarketDataUpdateSessionHigh {
	return MarketDataUpdateSessionHigh{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateSessionHigh(b []byte) MarketDataUpdateSessionHigh {
	return MarketDataUpdateSessionHigh{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateSessionHigh() MarketDataUpdateSessionHighBuilder {
	a := MarketDataUpdateSessionHighBuilder{message.NewFixed(24)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateSessionHighDefault)
	return a
}

// Clear
// {
//     Size               = MarketDataUpdateSessionHighSize  (24)
//     Type               = MARKET_DATA_UPDATE_SESSION_HIGH  (114)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
func (m MarketDataUpdateSessionHighBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateSessionHighDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionHigh) ToBuilder() MarketDataUpdateSessionHighBuilder {
	return MarketDataUpdateSessionHighBuilder{m.Fixed}
}

// Finish
func (m MarketDataUpdateSessionHighBuilder) Finish() MarketDataUpdateSessionHigh {
	return MarketDataUpdateSessionHigh{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionHigh) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionHighBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Price The session High price.
func (m MarketDataUpdateSessionHigh) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// Price The session High price.
func (m MarketDataUpdateSessionHighBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// TradingSessionDate
func (m MarketDataUpdateSessionHigh) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionHighBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionHighBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetPrice The session High price.
func (m MarketDataUpdateSessionHighBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionHighBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 20, 16, uint32(value))
}

// Copy
func (m MarketDataUpdateSessionHigh) Copy(to MarketDataUpdateSessionHighBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateSessionHighBuilder) Copy(to MarketDataUpdateSessionHighBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionHigh) CopyPointer(to MarketDataUpdateSessionHighPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionHighBuilder) CopyPointer(to MarketDataUpdateSessionHighPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}
