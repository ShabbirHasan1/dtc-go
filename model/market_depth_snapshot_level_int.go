package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDepthSnapshotLevel_IntSize = 40

// {
//     Size                  = MarketDepthSnapshotLevel_IntSize  (40)
//     Type                  = MARKET_DEPTH_SNAPSHOT_LEVEL_INT  (132)
//     SymbolID              = 0
//     Side                  = BID_ASK_UNSET  (0)
//     Price                 = 0
//     Quantity              = 0
//     Level                 = 0
//     IsFirstMessageInBatch = false
//     IsLastMessageInBatch  = false
//     DateTime              = 0
//     NumOrders             = 0
// }
var _MarketDepthSnapshotLevel_IntDefault = []byte{40, 0, 132, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketDepthSnapshotLevel_Int struct {
	message.Fixed
}

type MarketDepthSnapshotLevel_IntBuilder struct {
	message.Fixed
}

type MarketDepthSnapshotLevel_IntPointer struct {
	message.FixedPointer
}

type MarketDepthSnapshotLevel_IntPointerBuilder struct {
	message.FixedPointer
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

func AllocMarketDepthSnapshotLevel_Int() MarketDepthSnapshotLevel_IntPointerBuilder {
	a := MarketDepthSnapshotLevel_IntPointerBuilder{message.AllocFixed(40)}
	a.Ptr.SetBytes(0, _MarketDepthSnapshotLevel_IntDefault)
	return a
}

func AllocMarketDepthSnapshotLevel_IntFrom(b []byte) MarketDepthSnapshotLevel_IntPointer {
	return MarketDepthSnapshotLevel_IntPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                  = MarketDepthSnapshotLevel_IntSize  (40)
//     Type                  = MARKET_DEPTH_SNAPSHOT_LEVEL_INT  (132)
//     SymbolID              = 0
//     Side                  = BID_ASK_UNSET  (0)
//     Price                 = 0
//     Quantity              = 0
//     Level                 = 0
//     IsFirstMessageInBatch = false
//     IsLastMessageInBatch  = false
//     DateTime              = 0
//     NumOrders             = 0
// }
func (m MarketDepthSnapshotLevel_IntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthSnapshotLevel_IntDefault)
}

// Clear
// {
//     Size                  = MarketDepthSnapshotLevel_IntSize  (40)
//     Type                  = MARKET_DEPTH_SNAPSHOT_LEVEL_INT  (132)
//     SymbolID              = 0
//     Side                  = BID_ASK_UNSET  (0)
//     Price                 = 0
//     Quantity              = 0
//     Level                 = 0
//     IsFirstMessageInBatch = false
//     IsLastMessageInBatch  = false
//     DateTime              = 0
//     NumOrders             = 0
// }
func (m MarketDepthSnapshotLevel_IntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthSnapshotLevel_IntDefault)
}

// ToBuilder
func (m MarketDepthSnapshotLevel_Int) ToBuilder() MarketDepthSnapshotLevel_IntBuilder {
	return MarketDepthSnapshotLevel_IntBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDepthSnapshotLevel_IntPointer) ToBuilder() MarketDepthSnapshotLevel_IntPointerBuilder {
	return MarketDepthSnapshotLevel_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDepthSnapshotLevel_IntBuilder) Finish() MarketDepthSnapshotLevel_Int {
	return MarketDepthSnapshotLevel_Int{m.Fixed}
}

// Finish
func (m *MarketDepthSnapshotLevel_IntPointerBuilder) Finish() MarketDepthSnapshotLevel_IntPointer {
	return MarketDepthSnapshotLevel_IntPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDepthSnapshotLevel_Int) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDepthSnapshotLevel_IntBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDepthSnapshotLevel_IntPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketDepthSnapshotLevel_IntPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Side
func (m MarketDepthSnapshotLevel_Int) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 10, 8))
}

// Side
func (m MarketDepthSnapshotLevel_IntBuilder) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 10, 8))
}

// Side
func (m MarketDepthSnapshotLevel_IntPointer) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// Side
func (m MarketDepthSnapshotLevel_IntPointerBuilder) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// Price
func (m MarketDepthSnapshotLevel_Int) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// Price
func (m MarketDepthSnapshotLevel_IntBuilder) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// Price
func (m MarketDepthSnapshotLevel_IntPointer) Price() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// Price
func (m MarketDepthSnapshotLevel_IntPointerBuilder) Price() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// Quantity
func (m MarketDepthSnapshotLevel_Int) Quantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// Quantity
func (m MarketDepthSnapshotLevel_IntBuilder) Quantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// Quantity
func (m MarketDepthSnapshotLevel_IntPointer) Quantity() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// Quantity
func (m MarketDepthSnapshotLevel_IntPointerBuilder) Quantity() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// Level
func (m MarketDepthSnapshotLevel_Int) Level() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 22, 20)
}

// Level
func (m MarketDepthSnapshotLevel_IntBuilder) Level() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 22, 20)
}

// Level
func (m MarketDepthSnapshotLevel_IntPointer) Level() uint16 {
	return message.Uint16Fixed(m.Ptr, 22, 20)
}

// Level
func (m MarketDepthSnapshotLevel_IntPointerBuilder) Level() uint16 {
	return message.Uint16Fixed(m.Ptr, 22, 20)
}

// IsFirstMessageInBatch
func (m MarketDepthSnapshotLevel_Int) IsFirstMessageInBatch() bool {
	return message.BoolFixed(m.Unsafe(), 23, 22)
}

// IsFirstMessageInBatch
func (m MarketDepthSnapshotLevel_IntBuilder) IsFirstMessageInBatch() bool {
	return message.BoolFixed(m.Unsafe(), 23, 22)
}

// IsFirstMessageInBatch
func (m MarketDepthSnapshotLevel_IntPointer) IsFirstMessageInBatch() bool {
	return message.BoolFixed(m.Ptr, 23, 22)
}

// IsFirstMessageInBatch
func (m MarketDepthSnapshotLevel_IntPointerBuilder) IsFirstMessageInBatch() bool {
	return message.BoolFixed(m.Ptr, 23, 22)
}

// IsLastMessageInBatch
func (m MarketDepthSnapshotLevel_Int) IsLastMessageInBatch() bool {
	return message.BoolFixed(m.Unsafe(), 24, 23)
}

// IsLastMessageInBatch
func (m MarketDepthSnapshotLevel_IntBuilder) IsLastMessageInBatch() bool {
	return message.BoolFixed(m.Unsafe(), 24, 23)
}

// IsLastMessageInBatch
func (m MarketDepthSnapshotLevel_IntPointer) IsLastMessageInBatch() bool {
	return message.BoolFixed(m.Ptr, 24, 23)
}

// IsLastMessageInBatch
func (m MarketDepthSnapshotLevel_IntPointerBuilder) IsLastMessageInBatch() bool {
	return message.BoolFixed(m.Ptr, 24, 23)
}

// DateTime
func (m MarketDepthSnapshotLevel_Int) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 32, 24))
}

// DateTime
func (m MarketDepthSnapshotLevel_IntBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 32, 24))
}

// DateTime
func (m MarketDepthSnapshotLevel_IntPointer) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 32, 24))
}

// DateTime
func (m MarketDepthSnapshotLevel_IntPointerBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 32, 24))
}

// NumOrders
func (m MarketDepthSnapshotLevel_Int) NumOrders() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 36, 32)
}

// NumOrders
func (m MarketDepthSnapshotLevel_IntBuilder) NumOrders() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 36, 32)
}

// NumOrders
func (m MarketDepthSnapshotLevel_IntPointer) NumOrders() uint32 {
	return message.Uint32Fixed(m.Ptr, 36, 32)
}

// NumOrders
func (m MarketDepthSnapshotLevel_IntPointerBuilder) NumOrders() uint32 {
	return message.Uint32Fixed(m.Ptr, 36, 32)
}

// SetSymbolID
func (m MarketDepthSnapshotLevel_IntBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID
func (m MarketDepthSnapshotLevel_IntPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetSide
func (m MarketDepthSnapshotLevel_IntBuilder) SetSide(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Unsafe(), 10, 8, uint16(value))
}

// SetSide
func (m MarketDepthSnapshotLevel_IntPointerBuilder) SetSide(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Ptr, 10, 8, uint16(value))
}

// SetPrice
func (m MarketDepthSnapshotLevel_IntBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 16, 12, value)
}

// SetPrice
func (m MarketDepthSnapshotLevel_IntPointerBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 16, 12, value)
}

// SetQuantity
func (m MarketDepthSnapshotLevel_IntBuilder) SetQuantity(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 20, 16, value)
}

// SetQuantity
func (m MarketDepthSnapshotLevel_IntPointerBuilder) SetQuantity(value int32) {
	message.SetInt32Fixed(m.Ptr, 20, 16, value)
}

// SetLevel
func (m MarketDepthSnapshotLevel_IntBuilder) SetLevel(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 22, 20, value)
}

// SetLevel
func (m MarketDepthSnapshotLevel_IntPointerBuilder) SetLevel(value uint16) {
	message.SetUint16Fixed(m.Ptr, 22, 20, value)
}

// SetIsFirstMessageInBatch
func (m MarketDepthSnapshotLevel_IntBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolFixed(m.Unsafe(), 23, 22, value)
}

// SetIsFirstMessageInBatch
func (m MarketDepthSnapshotLevel_IntPointerBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolFixed(m.Ptr, 23, 22, value)
}

// SetIsLastMessageInBatch
func (m MarketDepthSnapshotLevel_IntBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolFixed(m.Unsafe(), 24, 23, value)
}

// SetIsLastMessageInBatch
func (m MarketDepthSnapshotLevel_IntPointerBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolFixed(m.Ptr, 24, 23, value)
}

// SetDateTime
func (m MarketDepthSnapshotLevel_IntBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 32, 24, float64(value))
}

// SetDateTime
func (m MarketDepthSnapshotLevel_IntPointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, float64(value))
}

// SetNumOrders
func (m MarketDepthSnapshotLevel_IntBuilder) SetNumOrders(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 36, 32, value)
}

// SetNumOrders
func (m MarketDepthSnapshotLevel_IntPointerBuilder) SetNumOrders(value uint32) {
	message.SetUint32Fixed(m.Ptr, 36, 32, value)
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

// Copy
func (m MarketDepthSnapshotLevel_IntPointer) Copy(to MarketDepthSnapshotLevel_IntBuilder) {
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
func (m MarketDepthSnapshotLevel_IntPointerBuilder) Copy(to MarketDepthSnapshotLevel_IntBuilder) {
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
func (m MarketDepthSnapshotLevel_IntPointer) CopyPointer(to MarketDepthSnapshotLevel_IntPointerBuilder) {
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
func (m MarketDepthSnapshotLevel_IntPointerBuilder) CopyPointer(to MarketDepthSnapshotLevel_IntPointerBuilder) {
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
