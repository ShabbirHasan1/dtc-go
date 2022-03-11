package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketOrdersModifyPointer struct {
	message.FixedPointer
}

type MarketOrdersModifyPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketOrdersModify() MarketOrdersModifyPointerBuilder {
	a := MarketOrdersModifyPointerBuilder{message.AllocFixed(56)}
	a.Ptr.SetBytes(0, _MarketOrdersModifyDefault)
	return a
}

func AllocMarketOrdersModifyFrom(b []byte) MarketOrdersModifyPointer {
	return MarketOrdersModifyPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = MarketOrdersModifySize  (56)
//     Type          = MARKET_ORDERS_MODIFY  (153)
//     SymbolID      = 0
//     Quantity      = 0
//     Price         = 0
//     OrderID       = 0
//     PriorPrice    = 0
//     PriorQuantity = 0
//     PriorOrderID  = 0
//     DateTime      = 0
// }
func (m MarketOrdersModifyPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersModifyDefault)
}

// ToBuilder
func (m MarketOrdersModifyPointer) ToBuilder() MarketOrdersModifyPointerBuilder {
	return MarketOrdersModifyPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketOrdersModifyPointerBuilder) Finish() MarketOrdersModifyPointer {
	return MarketOrdersModifyPointer{m.FixedPointer}
}

// SymbolID
func (m MarketOrdersModifyPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketOrdersModifyPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Quantity
func (m MarketOrdersModifyPointer) Quantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// Quantity
func (m MarketOrdersModifyPointerBuilder) Quantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// Price
func (m MarketOrdersModifyPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 20, 12)
}

// Price
func (m MarketOrdersModifyPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 20, 12)
}

// OrderID
func (m MarketOrdersModifyPointer) OrderID() uint64 {
	return message.Uint64Fixed(m.Ptr, 28, 20)
}

// OrderID
func (m MarketOrdersModifyPointerBuilder) OrderID() uint64 {
	return message.Uint64Fixed(m.Ptr, 28, 20)
}

// PriorPrice
func (m MarketOrdersModifyPointer) PriorPrice() float64 {
	return message.Float64Fixed(m.Ptr, 36, 28)
}

// PriorPrice
func (m MarketOrdersModifyPointerBuilder) PriorPrice() float64 {
	return message.Float64Fixed(m.Ptr, 36, 28)
}

// PriorQuantity
func (m MarketOrdersModifyPointer) PriorQuantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 40, 36)
}

// PriorQuantity
func (m MarketOrdersModifyPointerBuilder) PriorQuantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 40, 36)
}

// PriorOrderID
func (m MarketOrdersModifyPointer) PriorOrderID() uint64 {
	return message.Uint64Fixed(m.Ptr, 48, 40)
}

// PriorOrderID
func (m MarketOrdersModifyPointerBuilder) PriorOrderID() uint64 {
	return message.Uint64Fixed(m.Ptr, 48, 40)
}

// DateTime
func (m MarketOrdersModifyPointer) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 56, 48))
}

// DateTime
func (m MarketOrdersModifyPointerBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 56, 48))
}

// SetSymbolID
func (m MarketOrdersModifyPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetQuantity
func (m MarketOrdersModifyPointerBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 12, 8, value)
}

// SetPrice
func (m MarketOrdersModifyPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 20, 12, value)
}

// SetOrderID
func (m MarketOrdersModifyPointerBuilder) SetOrderID(value uint64) {
	message.SetUint64Fixed(m.Ptr, 28, 20, value)
}

// SetPriorPrice
func (m MarketOrdersModifyPointerBuilder) SetPriorPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 36, 28, value)
}

// SetPriorQuantity
func (m MarketOrdersModifyPointerBuilder) SetPriorQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 40, 36, value)
}

// SetPriorOrderID
func (m MarketOrdersModifyPointerBuilder) SetPriorOrderID(value uint64) {
	message.SetUint64Fixed(m.Ptr, 48, 40, value)
}

// SetDateTime
func (m MarketOrdersModifyPointerBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 56, 48, int64(value))
}

// Copy
func (m MarketOrdersModifyPointer) Copy(to MarketOrdersModifyBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetPriorPrice(m.PriorPrice())
	to.SetPriorQuantity(m.PriorQuantity())
	to.SetPriorOrderID(m.PriorOrderID())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketOrdersModifyPointerBuilder) Copy(to MarketOrdersModifyBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetPriorPrice(m.PriorPrice())
	to.SetPriorQuantity(m.PriorQuantity())
	to.SetPriorOrderID(m.PriorOrderID())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketOrdersModifyPointer) CopyPointer(to MarketOrdersModifyPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetPriorPrice(m.PriorPrice())
	to.SetPriorQuantity(m.PriorQuantity())
	to.SetPriorOrderID(m.PriorOrderID())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketOrdersModifyPointerBuilder) CopyPointer(to MarketOrdersModifyPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetPriorPrice(m.PriorPrice())
	to.SetPriorQuantity(m.PriorQuantity())
	to.SetPriorOrderID(m.PriorOrderID())
	to.SetDateTime(m.DateTime())
}
