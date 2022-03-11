package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDataUpdateSessionOpenPointer Sent by the Server to the Client to update the session Open.
type MarketDataUpdateSessionOpenPointer struct {
	message.FixedPointer
}

// MarketDataUpdateSessionOpenPointerBuilder Sent by the Server to the Client to update the session Open.
type MarketDataUpdateSessionOpenPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataUpdateSessionOpen() MarketDataUpdateSessionOpenPointerBuilder {
	a := MarketDataUpdateSessionOpenPointerBuilder{message.AllocFixed(24)}
	a.Ptr.SetBytes(0, _MarketDataUpdateSessionOpenDefault)
	return a
}

func AllocMarketDataUpdateSessionOpenFrom(b []byte) MarketDataUpdateSessionOpenPointer {
	return MarketDataUpdateSessionOpenPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size               = MarketDataUpdateSessionOpenSize  (24)
//     Type               = MARKET_DATA_UPDATE_SESSION_OPEN  (120)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
func (m MarketDataUpdateSessionOpenPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateSessionOpenDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionOpenPointer) ToBuilder() MarketDataUpdateSessionOpenPointerBuilder {
	return MarketDataUpdateSessionOpenPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateSessionOpenPointerBuilder) Finish() MarketDataUpdateSessionOpenPointer {
	return MarketDataUpdateSessionOpenPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionOpenPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionOpenPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Price The session Open price.
func (m MarketDataUpdateSessionOpenPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// Price The session Open price.
func (m MarketDataUpdateSessionOpenPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// TradingSessionDate
func (m MarketDataUpdateSessionOpenPointer) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionOpenPointerBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 20, 16))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionOpenPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetPrice The session Open price.
func (m MarketDataUpdateSessionOpenPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionOpenPointerBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 20, 16, uint32(value))
}

// Copy
func (m MarketDataUpdateSessionOpenPointer) Copy(to MarketDataUpdateSessionOpenBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateSessionOpenPointerBuilder) Copy(to MarketDataUpdateSessionOpenBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionOpenPointer) CopyPointer(to MarketDataUpdateSessionOpenPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionOpenPointerBuilder) CopyPointer(to MarketDataUpdateSessionOpenPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}
