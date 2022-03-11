package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateSessionLow_IntSize = 16

// {
//     Size               = MarketDataUpdateSessionLow_IntSize  (16)
//     Type               = MARKET_DATA_UPDATE_SESSION_LOW_INT  (130)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
var _MarketDataUpdateSessionLow_IntDefault = []byte{16, 0, 130, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketDataUpdateSessionLow_Int struct {
	message.Fixed
}

type MarketDataUpdateSessionLow_IntBuilder struct {
	message.Fixed
}

type MarketDataUpdateSessionLow_IntPointer struct {
	message.FixedPointer
}

type MarketDataUpdateSessionLow_IntPointerBuilder struct {
	message.FixedPointer
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

func AllocMarketDataUpdateSessionLow_Int() MarketDataUpdateSessionLow_IntPointerBuilder {
	a := MarketDataUpdateSessionLow_IntPointerBuilder{message.AllocFixed(16)}
	a.Ptr.SetBytes(0, _MarketDataUpdateSessionLow_IntDefault)
	return a
}

func AllocMarketDataUpdateSessionLow_IntFrom(b []byte) MarketDataUpdateSessionLow_IntPointer {
	return MarketDataUpdateSessionLow_IntPointer{FixedPointer: message.AllocFixedFrom(b)}
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

// Clear
// {
//     Size               = MarketDataUpdateSessionLow_IntSize  (16)
//     Type               = MARKET_DATA_UPDATE_SESSION_LOW_INT  (130)
//     SymbolID           = 0
//     Price              = 0
//     TradingSessionDate = 0
// }
func (m MarketDataUpdateSessionLow_IntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateSessionLow_IntDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionLow_Int) ToBuilder() MarketDataUpdateSessionLow_IntBuilder {
	return MarketDataUpdateSessionLow_IntBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateSessionLow_IntPointer) ToBuilder() MarketDataUpdateSessionLow_IntPointerBuilder {
	return MarketDataUpdateSessionLow_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateSessionLow_IntBuilder) Finish() MarketDataUpdateSessionLow_Int {
	return MarketDataUpdateSessionLow_Int{m.Fixed}
}

// Finish
func (m *MarketDataUpdateSessionLow_IntPointerBuilder) Finish() MarketDataUpdateSessionLow_IntPointer {
	return MarketDataUpdateSessionLow_IntPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDataUpdateSessionLow_Int) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDataUpdateSessionLow_IntBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDataUpdateSessionLow_IntPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketDataUpdateSessionLow_IntPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Price
func (m MarketDataUpdateSessionLow_Int) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// Price
func (m MarketDataUpdateSessionLow_IntBuilder) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// Price
func (m MarketDataUpdateSessionLow_IntPointer) Price() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// Price
func (m MarketDataUpdateSessionLow_IntPointerBuilder) Price() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// TradingSessionDate
func (m MarketDataUpdateSessionLow_Int) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 16, 12))
}

// TradingSessionDate
func (m MarketDataUpdateSessionLow_IntBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 16, 12))
}

// TradingSessionDate
func (m MarketDataUpdateSessionLow_IntPointer) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 16, 12))
}

// TradingSessionDate
func (m MarketDataUpdateSessionLow_IntPointerBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 16, 12))
}

// SetSymbolID
func (m MarketDataUpdateSessionLow_IntBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID
func (m MarketDataUpdateSessionLow_IntPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetPrice
func (m MarketDataUpdateSessionLow_IntBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetPrice
func (m MarketDataUpdateSessionLow_IntPointerBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionLow_IntBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 16, 12, uint32(value))
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionLow_IntPointerBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 16, 12, uint32(value))
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

// Copy
func (m MarketDataUpdateSessionLow_IntPointer) Copy(to MarketDataUpdateSessionLow_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateSessionLow_IntPointerBuilder) Copy(to MarketDataUpdateSessionLow_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionLow_IntPointer) CopyPointer(to MarketDataUpdateSessionLow_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionLow_IntPointerBuilder) CopyPointer(to MarketDataUpdateSessionLow_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetTradingSessionDate(m.TradingSessionDate())
}
