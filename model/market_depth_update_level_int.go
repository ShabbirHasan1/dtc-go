package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDepthUpdateLevel_IntSize = 40

// {
//     Size       = MarketDepthUpdateLevel_IntSize  (40)
//     Type       = MARKET_DEPTH_UPDATE_LEVEL_INT  (133)
//     SymbolID   = 0
//     Side       = BID_ASK_UNSET  (0)
//     Price      = 0
//     Quantity   = 0
//     UpdateType = MARKET_DEPTH_UNSET  (0)
//     DateTime   = 0
//     NumOrders  = 0
// }
var _MarketDepthUpdateLevel_IntDefault = []byte{40, 0, 133, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketDepthUpdateLevel_Int struct {
	message.Fixed
}

type MarketDepthUpdateLevel_IntBuilder struct {
	message.Fixed
}

type MarketDepthUpdateLevel_IntPointer struct {
	message.FixedPointer
}

type MarketDepthUpdateLevel_IntPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDepthUpdateLevel_IntFrom(b []byte) MarketDepthUpdateLevel_Int {
	return MarketDepthUpdateLevel_Int{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDepthUpdateLevel_Int(b []byte) MarketDepthUpdateLevel_Int {
	return MarketDepthUpdateLevel_Int{Fixed: message.WrapFixed(b)}
}

func NewMarketDepthUpdateLevel_Int() MarketDepthUpdateLevel_IntBuilder {
	a := MarketDepthUpdateLevel_IntBuilder{message.NewFixed(40)}
	a.Unsafe().SetBytes(0, _MarketDepthUpdateLevel_IntDefault)
	return a
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
//     Side       = BID_ASK_UNSET  (0)
//     Price      = 0
//     Quantity   = 0
//     UpdateType = MARKET_DEPTH_UNSET  (0)
//     DateTime   = 0
//     NumOrders  = 0
// }
func (m MarketDepthUpdateLevel_IntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthUpdateLevel_IntDefault)
}

// Clear
// {
//     Size       = MarketDepthUpdateLevel_IntSize  (40)
//     Type       = MARKET_DEPTH_UPDATE_LEVEL_INT  (133)
//     SymbolID   = 0
//     Side       = BID_ASK_UNSET  (0)
//     Price      = 0
//     Quantity   = 0
//     UpdateType = MARKET_DEPTH_UNSET  (0)
//     DateTime   = 0
//     NumOrders  = 0
// }
func (m MarketDepthUpdateLevel_IntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthUpdateLevel_IntDefault)
}

// ToBuilder
func (m MarketDepthUpdateLevel_Int) ToBuilder() MarketDepthUpdateLevel_IntBuilder {
	return MarketDepthUpdateLevel_IntBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDepthUpdateLevel_IntPointer) ToBuilder() MarketDepthUpdateLevel_IntPointerBuilder {
	return MarketDepthUpdateLevel_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDepthUpdateLevel_IntBuilder) Finish() MarketDepthUpdateLevel_Int {
	return MarketDepthUpdateLevel_Int{m.Fixed}
}

// Finish
func (m *MarketDepthUpdateLevel_IntPointerBuilder) Finish() MarketDepthUpdateLevel_IntPointer {
	return MarketDepthUpdateLevel_IntPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDepthUpdateLevel_Int) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDepthUpdateLevel_IntBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
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
func (m MarketDepthUpdateLevel_Int) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 10, 8))
}

// Side
func (m MarketDepthUpdateLevel_IntBuilder) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 10, 8))
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
func (m MarketDepthUpdateLevel_Int) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// Price
func (m MarketDepthUpdateLevel_IntBuilder) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
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
func (m MarketDepthUpdateLevel_Int) Quantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// Quantity
func (m MarketDepthUpdateLevel_IntBuilder) Quantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
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
func (m MarketDepthUpdateLevel_Int) UpdateType() MarketDepthUpdateTypeEnum {
	return MarketDepthUpdateTypeEnum(message.Uint8Fixed(m.Unsafe(), 21, 20))
}

// UpdateType
func (m MarketDepthUpdateLevel_IntBuilder) UpdateType() MarketDepthUpdateTypeEnum {
	return MarketDepthUpdateTypeEnum(message.Uint8Fixed(m.Unsafe(), 21, 20))
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
func (m MarketDepthUpdateLevel_Int) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 32, 24))
}

// DateTime
func (m MarketDepthUpdateLevel_IntBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 32, 24))
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
func (m MarketDepthUpdateLevel_Int) NumOrders() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 36, 32)
}

// NumOrders
func (m MarketDepthUpdateLevel_IntBuilder) NumOrders() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 36, 32)
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
func (m MarketDepthUpdateLevel_IntBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetSide
func (m MarketDepthUpdateLevel_IntBuilder) SetSide(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Unsafe(), 10, 8, uint16(value))
}

// SetSide
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetSide(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Ptr, 10, 8, uint16(value))
}

// SetPrice
func (m MarketDepthUpdateLevel_IntBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 16, 12, value)
}

// SetPrice
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 16, 12, value)
}

// SetQuantity
func (m MarketDepthUpdateLevel_IntBuilder) SetQuantity(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 20, 16, value)
}

// SetQuantity
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetQuantity(value int32) {
	message.SetInt32Fixed(m.Ptr, 20, 16, value)
}

// SetUpdateType
func (m MarketDepthUpdateLevel_IntBuilder) SetUpdateType(value MarketDepthUpdateTypeEnum) {
	message.SetUint8Fixed(m.Unsafe(), 21, 20, uint8(value))
}

// SetUpdateType
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetUpdateType(value MarketDepthUpdateTypeEnum) {
	message.SetUint8Fixed(m.Ptr, 21, 20, uint8(value))
}

// SetDateTime
func (m MarketDepthUpdateLevel_IntBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 32, 24, float64(value))
}

// SetDateTime
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, float64(value))
}

// SetNumOrders
func (m MarketDepthUpdateLevel_IntBuilder) SetNumOrders(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 36, 32, value)
}

// SetNumOrders
func (m MarketDepthUpdateLevel_IntPointerBuilder) SetNumOrders(value uint32) {
	message.SetUint32Fixed(m.Ptr, 36, 32, value)
}

// Copy
func (m MarketDepthUpdateLevel_Int) Copy(to MarketDepthUpdateLevel_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

// Copy
func (m MarketDepthUpdateLevel_IntBuilder) Copy(to MarketDepthUpdateLevel_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

// CopyPointer
func (m MarketDepthUpdateLevel_Int) CopyPointer(to MarketDepthUpdateLevel_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

// CopyPointer
func (m MarketDepthUpdateLevel_IntBuilder) CopyPointer(to MarketDepthUpdateLevel_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
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
