package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketDataUpdateSessionSettlement_IntPointer struct {
	message.FixedPointer
}

type MarketDataUpdateSessionSettlement_IntPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataUpdateSessionSettlement_Int() MarketDataUpdateSessionSettlement_IntPointerBuilder {
	a := MarketDataUpdateSessionSettlement_IntPointerBuilder{message.AllocFixed(16)}
	a.Ptr.SetBytes(0, _MarketDataUpdateSessionSettlement_IntDefault)
	return a
}

func AllocMarketDataUpdateSessionSettlement_IntFrom(b []byte) MarketDataUpdateSessionSettlement_IntPointer {
	return MarketDataUpdateSessionSettlement_IntPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size     = MarketDataUpdateSessionSettlement_IntSize  (16)
//     Type     = MARKET_DATA_UPDATE_SESSION_SETTLEMENT_INT  (131)
//     SymbolID = 0
//     Price    = 0
//     DateTime = 0
// }
func (m MarketDataUpdateSessionSettlement_IntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateSessionSettlement_IntDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionSettlement_IntPointer) ToBuilder() MarketDataUpdateSessionSettlement_IntPointerBuilder {
	return MarketDataUpdateSessionSettlement_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateSessionSettlement_IntPointerBuilder) Finish() MarketDataUpdateSessionSettlement_IntPointer {
	return MarketDataUpdateSessionSettlement_IntPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDataUpdateSessionSettlement_IntPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketDataUpdateSessionSettlement_IntPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Price
func (m MarketDataUpdateSessionSettlement_IntPointer) Price() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// Price
func (m MarketDataUpdateSessionSettlement_IntPointerBuilder) Price() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// DateTime
func (m MarketDataUpdateSessionSettlement_IntPointer) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 16, 12))
}

// DateTime
func (m MarketDataUpdateSessionSettlement_IntPointerBuilder) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 16, 12))
}

// SetSymbolID
func (m MarketDataUpdateSessionSettlement_IntPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetPrice
func (m MarketDataUpdateSessionSettlement_IntPointerBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetDateTime
func (m MarketDataUpdateSessionSettlement_IntPointerBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 16, 12, uint32(value))
}

// Copy
func (m MarketDataUpdateSessionSettlement_IntPointer) Copy(to MarketDataUpdateSessionSettlement_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateSessionSettlement_IntPointerBuilder) Copy(to MarketDataUpdateSessionSettlement_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateSessionSettlement_IntPointer) CopyPointer(to MarketDataUpdateSessionSettlement_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateSessionSettlement_IntPointerBuilder) CopyPointer(to MarketDataUpdateSessionSettlement_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
}
