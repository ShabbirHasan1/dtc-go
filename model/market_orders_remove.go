package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _MarketOrdersRemoveDefault = []byte{40, 0, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketOrdersRemoveSize = 40

type MarketOrdersRemove struct {
	message.Fixed
}

type MarketOrdersRemoveBuilder struct {
	message.Fixed
}

func NewMarketOrdersRemoveFrom(b []byte) MarketOrdersRemove {
	return MarketOrdersRemove{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketOrdersRemove(b []byte) MarketOrdersRemove {
	return MarketOrdersRemove{Fixed: message.WrapFixed(b)}
}

func NewMarketOrdersRemove() MarketOrdersRemoveBuilder {
	a := MarketOrdersRemoveBuilder{message.NewFixed(40)}
	a.Unsafe().SetBytes(0, _MarketOrdersRemoveDefault)
	return a
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
func (m MarketOrdersRemoveBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersRemoveDefault)
}

// ToBuilder
func (m MarketOrdersRemove) ToBuilder() MarketOrdersRemoveBuilder {
	return MarketOrdersRemoveBuilder{m.Fixed}
}

// Finish
func (m MarketOrdersRemoveBuilder) Finish() MarketOrdersRemove {
	return MarketOrdersRemove{m.Fixed}
}

// SymbolID
func (m MarketOrdersRemove) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketOrdersRemoveBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Quantity
func (m MarketOrdersRemove) Quantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// Quantity
func (m MarketOrdersRemoveBuilder) Quantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// Price
func (m MarketOrdersRemove) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 20, 12)
}

// Price
func (m MarketOrdersRemoveBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 20, 12)
}

// OrderID
func (m MarketOrdersRemove) OrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 28, 20)
}

// OrderID
func (m MarketOrdersRemoveBuilder) OrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 28, 20)
}

// DateTime
func (m MarketOrdersRemove) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 36, 28))
}

// DateTime
func (m MarketOrdersRemoveBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 36, 28))
}

// Side
func (m MarketOrdersRemove) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 40, 36))
}

// Side
func (m MarketOrdersRemoveBuilder) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 40, 36))
}

// SetSymbolID
func (m MarketOrdersRemoveBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetQuantity
func (m MarketOrdersRemoveBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 12, 8, value)
}

// SetPrice
func (m MarketOrdersRemoveBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 20, 12, value)
}

// SetOrderID
func (m MarketOrdersRemoveBuilder) SetOrderID(value uint64) {
	message.SetUint64Fixed(m.Unsafe(), 28, 20, value)
}

// SetDateTime
func (m MarketOrdersRemoveBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 36, 28, int64(value))
}

// SetSide
func (m MarketOrdersRemoveBuilder) SetSide(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 40, 36, int32(value))
}

// Copy
func (m MarketOrdersRemove) Copy(to MarketOrdersRemoveBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
	to.SetSide(m.Side())
}

// Copy
func (m MarketOrdersRemoveBuilder) Copy(to MarketOrdersRemoveBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
	to.SetSide(m.Side())
}

// CopyPointer
func (m MarketOrdersRemove) CopyPointer(to MarketOrdersRemovePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
	to.SetSide(m.Side())
}

// CopyPointer
func (m MarketOrdersRemoveBuilder) CopyPointer(to MarketOrdersRemovePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
	to.SetSide(m.Side())
}
