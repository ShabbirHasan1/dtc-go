package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketDepthUpdateLevel_IntPointer struct {
	message.FixedPointer
}

type MarketDepthUpdateLevel_IntPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDepthUpdateLevel_Int() MarketDepthUpdateLevel_IntPointerBuilder {
	a := MarketDepthUpdateLevel_IntPointerBuilder{message.AllocFixed(40)}
	a.Ptr.SetBytes(0, _MarketDepthUpdateLevel_IntDefault)
	return a
}

func AllocMarketDepthUpdateLevel_IntFrom(b []byte) MarketDepthUpdateLevel_IntPointer {
	return MarketDepthUpdateLevel_IntPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = MarketDepthUpdateLevel_IntSize  (40)
//     Type       = MARKET_DEPTH_UPDATE_LEVEL_INT  (133)
//     SymbolID   = 0
//     Side       = 0
//     Price      = 0
//     Quantity   = 0
//     UpdateType = 0
//     DateTime   = 0
//     NumOrders  = 0
// }
func (m MarketDepthUpdateLevel_IntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthUpdateLevel_IntDefault)
}

// ToBuilder
func (m MarketDepthUpdateLevel_IntPointer) ToBuilder() MarketDepthUpdateLevel_IntPointerBuilder {
	return MarketDepthUpdateLevel_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDepthUpdateLevel_IntPointerBuilder) Finish() MarketDepthUpdateLevel_IntPointer {
	return MarketDepthUpdateLevel_IntPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDepthUpdateLevel_IntPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketDepthUpdateLevel_IntPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Side
func (m MarketDepthUpdateLevel_IntPointer) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// Side
func (m MarketDepthUpdateLevel_IntPointerBuilder) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// Price
func (m MarketDepthUpdateLevel_IntPointer) Price() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// Price
func (m MarketDepthUpdateLevel_IntPointerBuilder) Price() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// Quantity
func (m MarketDepthUpdateLevel_IntPointer) Quantity() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// Quantity
func (m MarketDepthUpdateLevel_IntPointerBuilder) Quantity() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// UpdateType
func (m MarketDepthUpdateLevel_IntPointer) UpdateType() MarketDepthUpdateTypeEnum {
	return MarketDepthUpdateTypeEnum(message.Uint8Fixed(m.Ptr, 21, 20))
}

// UpdateType
func (m MarketDepthUpdateLevel_IntPointerBuilder) UpdateType() MarketDepthUpdateTypeEnum {
	return MarketDepthUpdateTypeEnum(message.Uint8Fixed(m.Ptr, 21, 20))
}

// DateTime
func (m MarketDepthUpdateLevel_IntPointer) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 32, 24))
}

// DateTime
func (m MarketDepthUpdateLevel_IntPointerBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 32, 24))
}

// NumOrders
func (m MarketDepthUpdateLevel_IntPointer) NumOrders() uint32 {
	return message.Uint32Fixed(m.Ptr, 36, 32)
}

// NumOrders
func (m MarketDepthUpdateLevel_IntPointerBuilder) NumOrders() uint32 {
	return message.Uint32Fixed(m.Ptr, 36, 32)
}

// SetSymbolID
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetSide
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetSide(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Ptr, 10, 8, uint16(value))
}

// SetPrice
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 16, 12, value)
}

// SetQuantity
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetQuantity(value int32) {
	message.SetInt32Fixed(m.Ptr, 20, 16, value)
}

// SetUpdateType
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetUpdateType(value MarketDepthUpdateTypeEnum) {
	message.SetUint8Fixed(m.Ptr, 21, 20, uint8(value))
}

// SetDateTime
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, float64(value))
}

// SetNumOrders
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetNumOrders(value uint32) {
	message.SetUint32Fixed(m.Ptr, 36, 32, value)
}

// Copy
func (m MarketDepthUpdateLevel_IntPointer) Copy(to MarketDepthUpdateLevel_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

// Copy
func (m MarketDepthUpdateLevel_IntPointerBuilder) Copy(to MarketDepthUpdateLevel_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

// CopyPointer
func (m MarketDepthUpdateLevel_IntPointer) CopyPointer(to MarketDepthUpdateLevel_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

// CopyPointer
func (m MarketDepthUpdateLevel_IntPointerBuilder) CopyPointer(to MarketDepthUpdateLevel_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}
