package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _MarketOrdersModifyDefault = []byte{56, 0, 153, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketOrdersModifySize = 56

type MarketOrdersModify struct {
	message.Fixed
}

type MarketOrdersModifyBuilder struct {
	message.Fixed
}

func NewMarketOrdersModifyFrom(b []byte) MarketOrdersModify {
	return MarketOrdersModify{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketOrdersModify(b []byte) MarketOrdersModify {
	return MarketOrdersModify{Fixed: message.WrapFixed(b)}
}

func NewMarketOrdersModify() MarketOrdersModifyBuilder {
	a := MarketOrdersModifyBuilder{message.NewFixed(56)}
	a.Unsafe().SetBytes(0, _MarketOrdersModifyDefault)
	return a
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
func (m MarketOrdersModifyBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersModifyDefault)
}

// ToBuilder
func (m MarketOrdersModify) ToBuilder() MarketOrdersModifyBuilder {
	return MarketOrdersModifyBuilder{m.Fixed}
}

// Finish
func (m MarketOrdersModifyBuilder) Finish() MarketOrdersModify {
	return MarketOrdersModify{m.Fixed}
}

// SymbolID
func (m MarketOrdersModify) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketOrdersModifyBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Quantity
func (m MarketOrdersModify) Quantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// Quantity
func (m MarketOrdersModifyBuilder) Quantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// Price
func (m MarketOrdersModify) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 20, 12)
}

// Price
func (m MarketOrdersModifyBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 20, 12)
}

// OrderID
func (m MarketOrdersModify) OrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 28, 20)
}

// OrderID
func (m MarketOrdersModifyBuilder) OrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 28, 20)
}

// PriorPrice
func (m MarketOrdersModify) PriorPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 36, 28)
}

// PriorPrice
func (m MarketOrdersModifyBuilder) PriorPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 36, 28)
}

// PriorQuantity
func (m MarketOrdersModify) PriorQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 40, 36)
}

// PriorQuantity
func (m MarketOrdersModifyBuilder) PriorQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 40, 36)
}

// PriorOrderID
func (m MarketOrdersModify) PriorOrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 48, 40)
}

// PriorOrderID
func (m MarketOrdersModifyBuilder) PriorOrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 48, 40)
}

// DateTime
func (m MarketOrdersModify) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 56, 48))
}

// DateTime
func (m MarketOrdersModifyBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 56, 48))
}

// SetSymbolID
func (m MarketOrdersModifyBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetQuantity
func (m MarketOrdersModifyBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 12, 8, value)
}

// SetPrice
func (m MarketOrdersModifyBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 20, 12, value)
}

// SetOrderID
func (m MarketOrdersModifyBuilder) SetOrderID(value uint64) {
	message.SetUint64Fixed(m.Unsafe(), 28, 20, value)
}

// SetPriorPrice
func (m MarketOrdersModifyBuilder) SetPriorPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 36, 28, value)
}

// SetPriorQuantity
func (m MarketOrdersModifyBuilder) SetPriorQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 40, 36, value)
}

// SetPriorOrderID
func (m MarketOrdersModifyBuilder) SetPriorOrderID(value uint64) {
	message.SetUint64Fixed(m.Unsafe(), 48, 40, value)
}

// SetDateTime
func (m MarketOrdersModifyBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 56, 48, int64(value))
}

// Copy
func (m MarketOrdersModify) Copy(to MarketOrdersModifyBuilder) {
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
func (m MarketOrdersModifyBuilder) Copy(to MarketOrdersModifyBuilder) {
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
func (m MarketOrdersModify) CopyPointer(to MarketOrdersModifyPointerBuilder) {
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
func (m MarketOrdersModifyBuilder) CopyPointer(to MarketOrdersModifyPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetPriorPrice(m.PriorPrice())
	to.SetPriorQuantity(m.PriorQuantity())
	to.SetPriorOrderID(m.PriorOrderID())
	to.SetDateTime(m.DateTime())
}
