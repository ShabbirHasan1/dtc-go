package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateSessionOpenSize = 24

// {
//     Size               = MarketDataUpdateSessionOpenSize  (24)
//     Type               = MARKET_DATA_UPDATE_SESSION_OPEN  (120)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
var _MarketDataUpdateSessionOpenDefault = []byte{24, 0, 120, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateSessionOpen Sent by the Server to the Client to update the session Open.
type MarketDataUpdateSessionOpen struct {
	message.Fixed
}

// MarketDataUpdateSessionOpenBuilder Sent by the Server to the Client to update the session Open.
type MarketDataUpdateSessionOpenBuilder struct {
	message.Fixed
}

// MarketDataUpdateSessionOpenPointer Sent by the Server to the Client to update the session Open.
type MarketDataUpdateSessionOpenPointer struct {
	message.FixedPointer
}

// MarketDataUpdateSessionOpenPointerBuilder Sent by the Server to the Client to update the session Open.
type MarketDataUpdateSessionOpenPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataUpdateSessionOpenFrom(b []byte) MarketDataUpdateSessionOpen {
	return MarketDataUpdateSessionOpen{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateSessionOpen(b []byte) MarketDataUpdateSessionOpen {
	return MarketDataUpdateSessionOpen{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateSessionOpen() MarketDataUpdateSessionOpenBuilder {
	a := MarketDataUpdateSessionOpenBuilder{message.NewFixed(24)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateSessionOpenDefault)
	return a
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
func (m MarketDataUpdateSessionOpenBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateSessionOpenDefault)
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
func (m MarketDataUpdateSessionOpen) ToBuilder() MarketDataUpdateSessionOpenBuilder {
	return MarketDataUpdateSessionOpenBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateSessionOpenPointer) ToBuilder() MarketDataUpdateSessionOpenPointerBuilder {
	return MarketDataUpdateSessionOpenPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateSessionOpenBuilder) Finish() MarketDataUpdateSessionOpen {
	return MarketDataUpdateSessionOpen{m.Fixed}
}

// Finish
func (m *MarketDataUpdateSessionOpenPointerBuilder) Finish() MarketDataUpdateSessionOpenPointer {
	return MarketDataUpdateSessionOpenPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionOpen) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionOpenBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
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
func (m MarketDataUpdateSessionOpen) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// Price The session Open price.
func (m MarketDataUpdateSessionOpenBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
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
func (m MarketDataUpdateSessionOpen) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionOpenBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
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
func (m MarketDataUpdateSessionOpenBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionOpenPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetPrice The session Open price.
func (m MarketDataUpdateSessionOpenBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, value)
}

// SetPrice The session Open price.
func (m MarketDataUpdateSessionOpenPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionOpenBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 20, 16, uint32(value))
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionOpenPointerBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 20, 16, uint32(value))
}

// Copy
func (m MarketDataUpdateSessionOpen) Copy(to MarketDataUpdateSessionOpenBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateSessionOpenBuilder) Copy(to MarketDataUpdateSessionOpenBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionOpen) CopyPointer(to MarketDataUpdateSessionOpenPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionOpenBuilder) CopyPointer(to MarketDataUpdateSessionOpenPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
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
