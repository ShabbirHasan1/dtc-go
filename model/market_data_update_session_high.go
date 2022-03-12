package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateSessionHighSize = 24

// {
//     Size               = MarketDataUpdateSessionHighSize  (24)
//     Type               = MARKET_DATA_UPDATE_SESSION_HIGH  (114)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
var _MarketDataUpdateSessionHighDefault = []byte{24, 0, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

// MarketDataUpdateSessionHighPointer Sent by the Server to the Client to update the session High as the High
// price changes throughout the session.
type MarketDataUpdateSessionHighPointer struct {
	message.FixedPointer
}

// MarketDataUpdateSessionHighPointerBuilder Sent by the Server to the Client to update the session High as the High
// price changes throughout the session.
type MarketDataUpdateSessionHighPointerBuilder struct {
	message.FixedPointer
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

func AllocMarketDataUpdateSessionHigh() MarketDataUpdateSessionHighPointerBuilder {
	a := MarketDataUpdateSessionHighPointerBuilder{message.AllocFixed(24)}
	a.Ptr.SetBytes(0, _MarketDataUpdateSessionHighDefault)
	return a
}

func AllocMarketDataUpdateSessionHighFrom(b []byte) MarketDataUpdateSessionHighPointer {
	return MarketDataUpdateSessionHighPointer{FixedPointer: message.AllocFixedFrom(b)}
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

// Clear
// {
//     Size               = MarketDataUpdateSessionHighSize  (24)
//     Type               = MARKET_DATA_UPDATE_SESSION_HIGH  (114)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
func (m MarketDataUpdateSessionHighPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateSessionHighDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionHigh) ToBuilder() MarketDataUpdateSessionHighBuilder {
	return MarketDataUpdateSessionHighBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateSessionHighPointer) ToBuilder() MarketDataUpdateSessionHighPointerBuilder {
	return MarketDataUpdateSessionHighPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateSessionHighBuilder) Finish() MarketDataUpdateSessionHigh {
	return MarketDataUpdateSessionHigh{m.Fixed}
}

// Finish
func (m *MarketDataUpdateSessionHighPointerBuilder) Finish() MarketDataUpdateSessionHighPointer {
	return MarketDataUpdateSessionHighPointer{m.FixedPointer}
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

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionHighPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionHighPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Price The session High price.
func (m MarketDataUpdateSessionHigh) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// Price The session High price.
func (m MarketDataUpdateSessionHighBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// Price The session High price.
func (m MarketDataUpdateSessionHighPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// Price The session High price.
func (m MarketDataUpdateSessionHighPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// TradingSessionDate
func (m MarketDataUpdateSessionHigh) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionHighBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionHighPointer) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionHighPointerBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 20, 16))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionHighBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionHighPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetPrice The session High price.
func (m MarketDataUpdateSessionHighBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, value)
}

// SetPrice The session High price.
func (m MarketDataUpdateSessionHighPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionHighBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 20, 16, uint32(value))
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionHighPointerBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 20, 16, uint32(value))
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

// Copy
func (m MarketDataUpdateSessionHighPointer) Copy(to MarketDataUpdateSessionHighBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateSessionHighPointerBuilder) Copy(to MarketDataUpdateSessionHighBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionHighPointer) CopyPointer(to MarketDataUpdateSessionHighPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionHighPointerBuilder) CopyPointer(to MarketDataUpdateSessionHighPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}
