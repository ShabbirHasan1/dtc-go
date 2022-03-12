package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketOrdersAddSize = 40

// {
//     Size     = MarketOrdersAddSize  (40)
//     Type     = MARKET_ORDERS_ADD  (152)
//     SymbolID = 0
//     Side     = BUY_SELL_UNSET  (0)
//     Quantity = 0
//     Price    = 0
//     OrderID  = 0
//     DateTime = 0
// }
var _MarketOrdersAddDefault = []byte{40, 0, 152, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketOrdersAdd struct {
	message.Fixed
}

type MarketOrdersAddBuilder struct {
	message.Fixed
}

type MarketOrdersAddPointer struct {
	message.FixedPointer
}

type MarketOrdersAddPointerBuilder struct {
	message.FixedPointer
}

func NewMarketOrdersAddFrom(b []byte) MarketOrdersAdd {
	return MarketOrdersAdd{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketOrdersAdd(b []byte) MarketOrdersAdd {
	return MarketOrdersAdd{Fixed: message.WrapFixed(b)}
}

func NewMarketOrdersAdd() MarketOrdersAddBuilder {
	a := MarketOrdersAddBuilder{message.NewFixed(40)}
	a.Unsafe().SetBytes(0, _MarketOrdersAddDefault)
	return a
}

func AllocMarketOrdersAdd() MarketOrdersAddPointerBuilder {
	a := MarketOrdersAddPointerBuilder{message.AllocFixed(40)}
	a.Ptr.SetBytes(0, _MarketOrdersAddDefault)
	return a
}

func AllocMarketOrdersAddFrom(b []byte) MarketOrdersAddPointer {
	return MarketOrdersAddPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size     = MarketOrdersAddSize  (40)
//     Type     = MARKET_ORDERS_ADD  (152)
//     SymbolID = 0
//     Side     = BUY_SELL_UNSET  (0)
//     Quantity = 0
//     Price    = 0
//     OrderID  = 0
//     DateTime = 0
// }
func (m MarketOrdersAddBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersAddDefault)
}

// Clear
// {
//     Size     = MarketOrdersAddSize  (40)
//     Type     = MARKET_ORDERS_ADD  (152)
//     SymbolID = 0
//     Side     = BUY_SELL_UNSET  (0)
//     Quantity = 0
//     Price    = 0
//     OrderID  = 0
//     DateTime = 0
// }
func (m MarketOrdersAddPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersAddDefault)
}

// ToBuilder
func (m MarketOrdersAdd) ToBuilder() MarketOrdersAddBuilder {
	return MarketOrdersAddBuilder{m.Fixed}
}

// ToBuilder
func (m MarketOrdersAddPointer) ToBuilder() MarketOrdersAddPointerBuilder {
	return MarketOrdersAddPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketOrdersAddBuilder) Finish() MarketOrdersAdd {
	return MarketOrdersAdd{m.Fixed}
}

// Finish
func (m *MarketOrdersAddPointerBuilder) Finish() MarketOrdersAddPointer {
	return MarketOrdersAddPointer{m.FixedPointer}
}

// SymbolID
func (m MarketOrdersAdd) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketOrdersAddBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketOrdersAddPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketOrdersAddPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Side
func (m MarketOrdersAdd) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Side
func (m MarketOrdersAddBuilder) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Side
func (m MarketOrdersAddPointer) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// Side
func (m MarketOrdersAddPointerBuilder) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// Quantity
func (m MarketOrdersAdd) Quantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 16, 12)
}

// Quantity
func (m MarketOrdersAddBuilder) Quantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 16, 12)
}

// Quantity
func (m MarketOrdersAddPointer) Quantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 16, 12)
}

// Quantity
func (m MarketOrdersAddPointerBuilder) Quantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 16, 12)
}

// Price
func (m MarketOrdersAdd) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// Price
func (m MarketOrdersAddBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// Price
func (m MarketOrdersAddPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// Price
func (m MarketOrdersAddPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// OrderID
func (m MarketOrdersAdd) OrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 32, 24)
}

// OrderID
func (m MarketOrdersAddBuilder) OrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 32, 24)
}

// OrderID
func (m MarketOrdersAddPointer) OrderID() uint64 {
	return message.Uint64Fixed(m.Ptr, 32, 24)
}

// OrderID
func (m MarketOrdersAddPointerBuilder) OrderID() uint64 {
	return message.Uint64Fixed(m.Ptr, 32, 24)
}

// DateTime
func (m MarketOrdersAdd) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 40, 32))
}

// DateTime
func (m MarketOrdersAddBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 40, 32))
}

// DateTime
func (m MarketOrdersAddPointer) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 40, 32))
}

// DateTime
func (m MarketOrdersAddPointerBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 40, 32))
}

// SetSymbolID
func (m MarketOrdersAddBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID
func (m MarketOrdersAddPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetSide
func (m MarketOrdersAddBuilder) SetSide(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, int32(value))
}

// SetSide
func (m MarketOrdersAddPointerBuilder) SetSide(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
}

// SetQuantity
func (m MarketOrdersAddBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 16, 12, value)
}

// SetQuantity
func (m MarketOrdersAddPointerBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 16, 12, value)
}

// SetPrice
func (m MarketOrdersAddBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetPrice
func (m MarketOrdersAddPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetOrderID
func (m MarketOrdersAddBuilder) SetOrderID(value uint64) {
	message.SetUint64Fixed(m.Unsafe(), 32, 24, value)
}

// SetOrderID
func (m MarketOrdersAddPointerBuilder) SetOrderID(value uint64) {
	message.SetUint64Fixed(m.Ptr, 32, 24, value)
}

// SetDateTime
func (m MarketOrdersAddBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 40, 32, int64(value))
}

// SetDateTime
func (m MarketOrdersAddPointerBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 40, 32, int64(value))
}

// Copy
func (m MarketOrdersAdd) Copy(to MarketOrdersAddBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketOrdersAddBuilder) Copy(to MarketOrdersAddBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketOrdersAdd) CopyPointer(to MarketOrdersAddPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketOrdersAddBuilder) CopyPointer(to MarketOrdersAddPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketOrdersAddPointer) Copy(to MarketOrdersAddBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketOrdersAddPointerBuilder) Copy(to MarketOrdersAddBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketOrdersAddPointer) CopyPointer(to MarketOrdersAddPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketOrdersAddPointerBuilder) CopyPointer(to MarketOrdersAddPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
}
