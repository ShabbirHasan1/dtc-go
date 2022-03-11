package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                  = MarketDepthSnapshotLevel_IntSize  (40)
//     Type                  = MARKET_DEPTH_SNAPSHOT_LEVEL_INT  (132)
//     SymbolID              = 0
//     Side                  = 0
//     Price                 = 0
//     Quantity              = 0
//     Level                 = 0
//     IsFirstMessageInBatch = 0
//     IsLastMessageInBatch  = 0
//     DateTime              = 0
//     NumOrders             = 0
// }
var _MarketDepthSnapshotLevel_IntDefault = []byte{40, 0, 132, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDepthSnapshotLevel_IntSize = 40

type MarketDepthSnapshotLevel_Int struct {
	message.Fixed
}

type MarketDepthSnapshotLevel_IntBuilder struct {
	message.Fixed
}

func NewMarketDepthSnapshotLevel_IntFrom(b []byte) MarketDepthSnapshotLevel_Int {
	return MarketDepthSnapshotLevel_Int{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDepthSnapshotLevel_Int(b []byte) MarketDepthSnapshotLevel_Int {
	return MarketDepthSnapshotLevel_Int{Fixed: message.WrapFixed(b)}
}

func NewMarketDepthSnapshotLevel_Int() MarketDepthSnapshotLevel_IntBuilder {
	a := MarketDepthSnapshotLevel_IntBuilder{message.NewFixed(40)}
	a.Unsafe().SetBytes(0, _MarketDepthSnapshotLevel_IntDefault)
	return a
}

// Clear
// {
//     Size                  = MarketDepthSnapshotLevel_IntSize  (40)
//     Type                  = MARKET_DEPTH_SNAPSHOT_LEVEL_INT  (132)
//     SymbolID              = 0
//     Side                  = 0
//     Price                 = 0
//     Quantity              = 0
//     Level                 = 0
//     IsFirstMessageInBatch = 0
//     IsLastMessageInBatch  = 0
//     DateTime              = 0
//     NumOrders             = 0
// }
func (m MarketDepthSnapshotLevel_IntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthSnapshotLevel_IntDefault)
}

// ToBuilder
func (m MarketDepthSnapshotLevel_Int) ToBuilder() MarketDepthSnapshotLevel_IntBuilder {
	return MarketDepthSnapshotLevel_IntBuilder{m.Fixed}
}

// Finish
func (m MarketDepthSnapshotLevel_IntBuilder) Finish() MarketDepthSnapshotLevel_Int {
	return MarketDepthSnapshotLevel_Int{m.Fixed}
}

// SymbolID
func (m MarketDepthSnapshotLevel_Int) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDepthSnapshotLevel_IntBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Side
func (m MarketDepthSnapshotLevel_Int) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 10, 8))
}

// Side
func (m MarketDepthSnapshotLevel_IntBuilder) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 10, 8))
}

// Price
func (m MarketDepthSnapshotLevel_Int) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// Price
func (m MarketDepthSnapshotLevel_IntBuilder) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// Quantity
func (m MarketDepthSnapshotLevel_Int) Quantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// Quantity
func (m MarketDepthSnapshotLevel_IntBuilder) Quantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// Level
func (m MarketDepthSnapshotLevel_Int) Level() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 22, 20)
}

// Level
func (m MarketDepthSnapshotLevel_IntBuilder) Level() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 22, 20)
}

// IsFirstMessageInBatch
func (m MarketDepthSnapshotLevel_Int) IsFirstMessageInBatch() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 23, 22)
}

// IsFirstMessageInBatch
func (m MarketDepthSnapshotLevel_IntBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 23, 22)
}

// IsLastMessageInBatch
func (m MarketDepthSnapshotLevel_Int) IsLastMessageInBatch() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 24, 23)
}

// IsLastMessageInBatch
func (m MarketDepthSnapshotLevel_IntBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 24, 23)
}

// DateTime
func (m MarketDepthSnapshotLevel_Int) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 32, 24))
}

// DateTime
func (m MarketDepthSnapshotLevel_IntBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 32, 24))
}

// NumOrders
func (m MarketDepthSnapshotLevel_Int) NumOrders() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 36, 32)
}

// NumOrders
func (m MarketDepthSnapshotLevel_IntBuilder) NumOrders() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 36, 32)
}

// SetSymbolID
func (m MarketDepthSnapshotLevel_IntBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSide
func (m MarketDepthSnapshotLevel_IntBuilder) SetSide(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Unsafe(), 10, 8, uint16(value))
}

// SetPrice
func (m MarketDepthSnapshotLevel_IntBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 16, 12, value)
}

// SetQuantity
func (m MarketDepthSnapshotLevel_IntBuilder) SetQuantity(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 20, 16, value)
}

// SetLevel
func (m MarketDepthSnapshotLevel_IntBuilder) SetLevel(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 22, 20, value)
}

// SetIsFirstMessageInBatch
func (m MarketDepthSnapshotLevel_IntBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 23, 22, value)
}

// SetIsLastMessageInBatch
func (m MarketDepthSnapshotLevel_IntBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 24, 23, value)
}

// SetDateTime
func (m MarketDepthSnapshotLevel_IntBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 32, 24, float64(value))
}

// SetNumOrders
func (m MarketDepthSnapshotLevel_IntBuilder) SetNumOrders(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 36, 32, value)
}

// Copy
func (m MarketDepthSnapshotLevel_Int) Copy(to MarketDepthSnapshotLevel_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetLevel(m.Level())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

// Copy
func (m MarketDepthSnapshotLevel_IntBuilder) Copy(to MarketDepthSnapshotLevel_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetLevel(m.Level())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

// CopyPointer
func (m MarketDepthSnapshotLevel_Int) CopyPointer(to MarketDepthSnapshotLevel_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetLevel(m.Level())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

// CopyPointer
func (m MarketDepthSnapshotLevel_IntBuilder) CopyPointer(to MarketDepthSnapshotLevel_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetLevel(m.Level())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}
