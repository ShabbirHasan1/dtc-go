package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketOrdersRemovePointer struct {
	message.FixedPointer
}

type MarketOrdersRemovePointerBuilder struct {
	message.FixedPointer
}

func AllocMarketOrdersRemove() MarketOrdersRemovePointerBuilder {
	a := MarketOrdersRemovePointerBuilder{message.AllocFixed(40)}
	a.Ptr.SetBytes(0, _MarketOrdersRemoveDefault)
	return a
}

func AllocMarketOrdersRemoveFrom(b []byte) MarketOrdersRemovePointer {
	return MarketOrdersRemovePointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size     = MarketOrdersRemoveSize  (40)
//     Type     = MARKET_ORDERS_REMOVE  (154)
//     SymbolID = 0
//     Quantity = 0
//     Price    = 0
//     OrderID  = 0
//     DateTime = 0
//     Side     = BUY_SELL_UNSET  (0)
// }
func (m MarketOrdersRemovePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersRemoveDefault)
}

// ToBuilder
func (m MarketOrdersRemovePointer) ToBuilder() MarketOrdersRemovePointerBuilder {
	return MarketOrdersRemovePointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketOrdersRemovePointerBuilder) Finish() MarketOrdersRemovePointer {
	return MarketOrdersRemovePointer{m.FixedPointer}
}

// SymbolID
func (m MarketOrdersRemovePointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketOrdersRemovePointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Quantity
func (m MarketOrdersRemovePointer) Quantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// Quantity
func (m MarketOrdersRemovePointerBuilder) Quantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// Price
func (m MarketOrdersRemovePointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 20, 12)
}

// Price
func (m MarketOrdersRemovePointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 20, 12)
}

// OrderID
func (m MarketOrdersRemovePointer) OrderID() uint64 {
	return message.Uint64Fixed(m.Ptr, 28, 20)
}

// OrderID
func (m MarketOrdersRemovePointerBuilder) OrderID() uint64 {
	return message.Uint64Fixed(m.Ptr, 28, 20)
}

// DateTime
func (m MarketOrdersRemovePointer) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 36, 28))
}

// DateTime
func (m MarketOrdersRemovePointerBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 36, 28))
}

// Side
func (m MarketOrdersRemovePointer) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 40, 36))
}

// Side
func (m MarketOrdersRemovePointerBuilder) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 40, 36))
}

// SetSymbolID
func (m MarketOrdersRemovePointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetQuantity
func (m MarketOrdersRemovePointerBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 12, 8, value)
}

// SetPrice
func (m MarketOrdersRemovePointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 20, 12, value)
}

// SetOrderID
func (m MarketOrdersRemovePointerBuilder) SetOrderID(value uint64) {
	message.SetUint64Fixed(m.Ptr, 28, 20, value)
}

// SetDateTime
func (m MarketOrdersRemovePointerBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 36, 28, int64(value))
}

// SetSide
func (m MarketOrdersRemovePointerBuilder) SetSide(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 40, 36, int32(value))
}

// Copy
func (m MarketOrdersRemovePointer) Copy(to MarketOrdersRemoveBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
	to.SetSide(m.Side())
}

// Copy
func (m MarketOrdersRemovePointerBuilder) Copy(to MarketOrdersRemoveBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
	to.SetSide(m.Side())
}

// CopyPointer
func (m MarketOrdersRemovePointer) CopyPointer(to MarketOrdersRemovePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
	to.SetSide(m.Side())
}

// CopyPointer
func (m MarketOrdersRemovePointerBuilder) CopyPointer(to MarketOrdersRemovePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
	to.SetSide(m.Side())
}
