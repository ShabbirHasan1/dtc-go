package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketOrdersRemoveSize = 40

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

type MarketOrdersRemove struct {
	message.Fixed
}

type MarketOrdersRemoveBuilder struct {
	message.Fixed
}

type MarketOrdersRemovePointer struct {
	message.FixedPointer
}

type MarketOrdersRemovePointerBuilder struct {
	message.FixedPointer
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
func (m MarketOrdersRemoveBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersRemoveDefault)
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
func (m MarketOrdersRemove) ToBuilder() MarketOrdersRemoveBuilder {
	return MarketOrdersRemoveBuilder{m.Fixed}
}

// ToBuilder
func (m MarketOrdersRemovePointer) ToBuilder() MarketOrdersRemovePointerBuilder {
	return MarketOrdersRemovePointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketOrdersRemoveBuilder) Finish() MarketOrdersRemove {
	return MarketOrdersRemove{m.Fixed}
}

// Finish
func (m *MarketOrdersRemovePointerBuilder) Finish() MarketOrdersRemovePointer {
	return MarketOrdersRemovePointer{m.FixedPointer}
}

// SymbolID
func (m MarketOrdersRemove) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketOrdersRemoveBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
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
func (m MarketOrdersRemove) Quantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// Quantity
func (m MarketOrdersRemoveBuilder) Quantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
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
func (m MarketOrdersRemove) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 20, 12)
}

// Price
func (m MarketOrdersRemoveBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 20, 12)
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
func (m MarketOrdersRemove) OrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 28, 20)
}

// OrderID
func (m MarketOrdersRemoveBuilder) OrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 28, 20)
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
func (m MarketOrdersRemove) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 36, 28))
}

// DateTime
func (m MarketOrdersRemoveBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 36, 28))
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
func (m MarketOrdersRemove) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 40, 36))
}

// Side
func (m MarketOrdersRemoveBuilder) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 40, 36))
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
func (m MarketOrdersRemoveBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID
func (m MarketOrdersRemovePointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetQuantity
func (m MarketOrdersRemoveBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 12, 8, value)
}

// SetQuantity
func (m MarketOrdersRemovePointerBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 12, 8, value)
}

// SetPrice
func (m MarketOrdersRemoveBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 20, 12, value)
}

// SetPrice
func (m MarketOrdersRemovePointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 20, 12, value)
}

// SetOrderID
func (m MarketOrdersRemoveBuilder) SetOrderID(value uint64) {
	message.SetUint64Fixed(m.Unsafe(), 28, 20, value)
}

// SetOrderID
func (m MarketOrdersRemovePointerBuilder) SetOrderID(value uint64) {
	message.SetUint64Fixed(m.Ptr, 28, 20, value)
}

// SetDateTime
func (m MarketOrdersRemoveBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 36, 28, int64(value))
}

// SetDateTime
func (m MarketOrdersRemovePointerBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 36, 28, int64(value))
}

// SetSide
func (m MarketOrdersRemoveBuilder) SetSide(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 40, 36, int32(value))
}

// SetSide
func (m MarketOrdersRemovePointerBuilder) SetSide(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 40, 36, int32(value))
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
