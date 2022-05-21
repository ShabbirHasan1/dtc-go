// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-16 08:00:54.519198 +0800 WITA m=+0.055446543

package v8

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketOrdersModifySize = 60

//     Size           uint16                       = MarketOrdersModifySize  (60)
//     Type           uint16                       = MARKET_ORDERS_MODIFY  (153)
//     SymbolID       uint32                       = 0
//     Side           BuySellEnum                  = BUY_SELL_UNSET  (0)
//     Quantity       uint32                       = 0
//     Price          float64                      = 0
//     OrderID        uint64                       = 0
//     PriorPrice     float64                      = 0
//     PriorQuantity  uint32                       = 0
//     PriorOrderID   uint64                       = 0
//     DateTime       DateTimeWithMicrosecondsInt  = 0
var _MarketOrdersModifyDefault = []byte{60, 0, 153, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketOrdersModify struct {
	message.Fixed
}

type MarketOrdersModifyBuilder struct {
	message.Fixed
}

type MarketOrdersModifyPointer struct {
	message.FixedPointer
}

type MarketOrdersModifyPointerBuilder struct {
	message.FixedPointer
}

func NewMarketOrdersModifyFrom(b []byte) MarketOrdersModify {
	return MarketOrdersModify{Fixed: message.NewFixed(b)}
}

func WrapMarketOrdersModify(b []byte) MarketOrdersModify {
	return MarketOrdersModify{Fixed: message.WrapFixed(b)}
}

func NewMarketOrdersModify() MarketOrdersModifyBuilder {
	return MarketOrdersModifyBuilder{message.NewFixed(_MarketOrdersModifyDefault)}
}

func AllocMarketOrdersModify() MarketOrdersModifyPointerBuilder {
	return MarketOrdersModifyPointerBuilder{message.AllocFixed(_MarketOrdersModifyDefault)}
}

func AllocMarketOrdersModifyFrom(b []byte) MarketOrdersModifyPointer {
	return MarketOrdersModifyPointer{FixedPointer: message.AllocFixed(b)}
}

// Clear
//     Size           uint16                       = MarketOrdersModifySize  (60)
//     Type           uint16                       = MARKET_ORDERS_MODIFY  (153)
//     SymbolID       uint32                       = 0
//     Side           BuySellEnum                  = BUY_SELL_UNSET  (0)
//     Quantity       uint32                       = 0
//     Price          float64                      = 0
//     OrderID        uint64                       = 0
//     PriorPrice     float64                      = 0
//     PriorQuantity  uint32                       = 0
//     PriorOrderID   uint64                       = 0
//     DateTime       DateTimeWithMicrosecondsInt  = 0
func (m MarketOrdersModifyBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersModifyDefault)
}

// Clear
//     Size           uint16                       = MarketOrdersModifySize  (60)
//     Type           uint16                       = MARKET_ORDERS_MODIFY  (153)
//     SymbolID       uint32                       = 0
//     Side           BuySellEnum                  = BUY_SELL_UNSET  (0)
//     Quantity       uint32                       = 0
//     Price          float64                      = 0
//     OrderID        uint64                       = 0
//     PriorPrice     float64                      = 0
//     PriorQuantity  uint32                       = 0
//     PriorOrderID   uint64                       = 0
//     DateTime       DateTimeWithMicrosecondsInt  = 0
func (m MarketOrdersModifyPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersModifyDefault)
}

// ToBuilder
func (m MarketOrdersModify) ToBuilder() MarketOrdersModifyBuilder {
	return MarketOrdersModifyBuilder{m.Fixed}
}

// ToBuilder
func (m MarketOrdersModifyPointer) ToBuilder() MarketOrdersModifyPointerBuilder {
	return MarketOrdersModifyPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketOrdersModifyBuilder) Finish() MarketOrdersModify {
	return MarketOrdersModify{m.Fixed}
}

// Finish
func (m *MarketOrdersModifyPointerBuilder) Finish() MarketOrdersModifyPointer {
	return MarketOrdersModifyPointer{m.FixedPointer}
}

// SymbolID
func (m MarketOrdersModify) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketOrdersModifyBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketOrdersModifyPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketOrdersModifyPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Side
func (m MarketOrdersModify) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Side
func (m MarketOrdersModifyBuilder) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Side
func (m MarketOrdersModifyPointer) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// Side
func (m MarketOrdersModifyPointerBuilder) Side() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// Quantity
func (m MarketOrdersModify) Quantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 16, 12)
}

// Quantity
func (m MarketOrdersModifyBuilder) Quantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 16, 12)
}

// Quantity
func (m MarketOrdersModifyPointer) Quantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 16, 12)
}

// Quantity
func (m MarketOrdersModifyPointerBuilder) Quantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 16, 12)
}

// Price
func (m MarketOrdersModify) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// Price
func (m MarketOrdersModifyBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// Price
func (m MarketOrdersModifyPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// Price
func (m MarketOrdersModifyPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// OrderID
func (m MarketOrdersModify) OrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 32, 24)
}

// OrderID
func (m MarketOrdersModifyBuilder) OrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 32, 24)
}

// OrderID
func (m MarketOrdersModifyPointer) OrderID() uint64 {
	return message.Uint64Fixed(m.Ptr, 32, 24)
}

// OrderID
func (m MarketOrdersModifyPointerBuilder) OrderID() uint64 {
	return message.Uint64Fixed(m.Ptr, 32, 24)
}

// PriorPrice
func (m MarketOrdersModify) PriorPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 40, 32)
}

// PriorPrice
func (m MarketOrdersModifyBuilder) PriorPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 40, 32)
}

// PriorPrice
func (m MarketOrdersModifyPointer) PriorPrice() float64 {
	return message.Float64Fixed(m.Ptr, 40, 32)
}

// PriorPrice
func (m MarketOrdersModifyPointerBuilder) PriorPrice() float64 {
	return message.Float64Fixed(m.Ptr, 40, 32)
}

// PriorQuantity
func (m MarketOrdersModify) PriorQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 44, 40)
}

// PriorQuantity
func (m MarketOrdersModifyBuilder) PriorQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 44, 40)
}

// PriorQuantity
func (m MarketOrdersModifyPointer) PriorQuantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 44, 40)
}

// PriorQuantity
func (m MarketOrdersModifyPointerBuilder) PriorQuantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 44, 40)
}

// PriorOrderID
func (m MarketOrdersModify) PriorOrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 52, 44)
}

// PriorOrderID
func (m MarketOrdersModifyBuilder) PriorOrderID() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 52, 44)
}

// PriorOrderID
func (m MarketOrdersModifyPointer) PriorOrderID() uint64 {
	return message.Uint64Fixed(m.Ptr, 52, 44)
}

// PriorOrderID
func (m MarketOrdersModifyPointerBuilder) PriorOrderID() uint64 {
	return message.Uint64Fixed(m.Ptr, 52, 44)
}

// DateTime
func (m MarketOrdersModify) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 60, 52))
}

// DateTime
func (m MarketOrdersModifyBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 60, 52))
}

// DateTime
func (m MarketOrdersModifyPointer) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 60, 52))
}

// DateTime
func (m MarketOrdersModifyPointerBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 60, 52))
}

// SetSymbolID
func (m MarketOrdersModifyBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID
func (m MarketOrdersModifyPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetSide
func (m MarketOrdersModifyBuilder) SetSide(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, int32(value))
}

// SetSide
func (m MarketOrdersModifyPointerBuilder) SetSide(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
}

// SetQuantity
func (m MarketOrdersModifyBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 16, 12, value)
}

// SetQuantity
func (m MarketOrdersModifyPointerBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 16, 12, value)
}

// SetPrice
func (m MarketOrdersModifyBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetPrice
func (m MarketOrdersModifyPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetOrderID
func (m MarketOrdersModifyBuilder) SetOrderID(value uint64) {
	message.SetUint64Fixed(m.Unsafe(), 32, 24, value)
}

// SetOrderID
func (m MarketOrdersModifyPointerBuilder) SetOrderID(value uint64) {
	message.SetUint64Fixed(m.Ptr, 32, 24, value)
}

// SetPriorPrice
func (m MarketOrdersModifyBuilder) SetPriorPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 40, 32, value)
}

// SetPriorPrice
func (m MarketOrdersModifyPointerBuilder) SetPriorPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 40, 32, value)
}

// SetPriorQuantity
func (m MarketOrdersModifyBuilder) SetPriorQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 44, 40, value)
}

// SetPriorQuantity
func (m MarketOrdersModifyPointerBuilder) SetPriorQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 44, 40, value)
}

// SetPriorOrderID
func (m MarketOrdersModifyBuilder) SetPriorOrderID(value uint64) {
	message.SetUint64Fixed(m.Unsafe(), 52, 44, value)
}

// SetPriorOrderID
func (m MarketOrdersModifyPointerBuilder) SetPriorOrderID(value uint64) {
	message.SetUint64Fixed(m.Ptr, 52, 44, value)
}

// SetDateTime
func (m MarketOrdersModifyBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 60, 52, int64(value))
}

// SetDateTime
func (m MarketOrdersModifyPointerBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 60, 52, int64(value))
}
