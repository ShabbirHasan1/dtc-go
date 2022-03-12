package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketOrdersSnapshotMessageBoundarySize = 9

// {
//     Size            = MarketOrdersSnapshotMessageBoundarySize  (9)
//     Type            = MARKET_ORDERS_SNAPSHOT_MESSAGE_BOUNDARY  (155)
//     SymbolID        = 0
//     MessageBoundary = MESSAGE_SET_BOUNDARY_UNSET  (0)
// }
var _MarketOrdersSnapshotMessageBoundaryDefault = []byte{9, 0, 155, 0, 0, 0, 0, 0, 0}

type MarketOrdersSnapshotMessageBoundary struct {
	message.Fixed
}

type MarketOrdersSnapshotMessageBoundaryBuilder struct {
	message.Fixed
}

type MarketOrdersSnapshotMessageBoundaryPointer struct {
	message.FixedPointer
}

type MarketOrdersSnapshotMessageBoundaryPointerBuilder struct {
	message.FixedPointer
}

func NewMarketOrdersSnapshotMessageBoundaryFrom(b []byte) MarketOrdersSnapshotMessageBoundary {
	return MarketOrdersSnapshotMessageBoundary{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketOrdersSnapshotMessageBoundary(b []byte) MarketOrdersSnapshotMessageBoundary {
	return MarketOrdersSnapshotMessageBoundary{Fixed: message.WrapFixed(b)}
}

func NewMarketOrdersSnapshotMessageBoundary() MarketOrdersSnapshotMessageBoundaryBuilder {
	a := MarketOrdersSnapshotMessageBoundaryBuilder{message.NewFixed(9)}
	a.Unsafe().SetBytes(0, _MarketOrdersSnapshotMessageBoundaryDefault)
	return a
}

func AllocMarketOrdersSnapshotMessageBoundary() MarketOrdersSnapshotMessageBoundaryPointerBuilder {
	a := MarketOrdersSnapshotMessageBoundaryPointerBuilder{message.AllocFixed(9)}
	a.Ptr.SetBytes(0, _MarketOrdersSnapshotMessageBoundaryDefault)
	return a
}

func AllocMarketOrdersSnapshotMessageBoundaryFrom(b []byte) MarketOrdersSnapshotMessageBoundaryPointer {
	return MarketOrdersSnapshotMessageBoundaryPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size            = MarketOrdersSnapshotMessageBoundarySize  (9)
//     Type            = MARKET_ORDERS_SNAPSHOT_MESSAGE_BOUNDARY  (155)
//     SymbolID        = 0
//     MessageBoundary = MESSAGE_SET_BOUNDARY_UNSET  (0)
// }
func (m MarketOrdersSnapshotMessageBoundaryBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketOrdersSnapshotMessageBoundaryDefault)
}

// Clear
// {
//     Size            = MarketOrdersSnapshotMessageBoundarySize  (9)
//     Type            = MARKET_ORDERS_SNAPSHOT_MESSAGE_BOUNDARY  (155)
//     SymbolID        = 0
//     MessageBoundary = MESSAGE_SET_BOUNDARY_UNSET  (0)
// }
func (m MarketOrdersSnapshotMessageBoundaryPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketOrdersSnapshotMessageBoundaryDefault)
}

// ToBuilder
func (m MarketOrdersSnapshotMessageBoundary) ToBuilder() MarketOrdersSnapshotMessageBoundaryBuilder {
	return MarketOrdersSnapshotMessageBoundaryBuilder{m.Fixed}
}

// ToBuilder
func (m MarketOrdersSnapshotMessageBoundaryPointer) ToBuilder() MarketOrdersSnapshotMessageBoundaryPointerBuilder {
	return MarketOrdersSnapshotMessageBoundaryPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketOrdersSnapshotMessageBoundaryBuilder) Finish() MarketOrdersSnapshotMessageBoundary {
	return MarketOrdersSnapshotMessageBoundary{m.Fixed}
}

// Finish
func (m *MarketOrdersSnapshotMessageBoundaryPointerBuilder) Finish() MarketOrdersSnapshotMessageBoundaryPointer {
	return MarketOrdersSnapshotMessageBoundaryPointer{m.FixedPointer}
}

// SymbolID
func (m MarketOrdersSnapshotMessageBoundary) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketOrdersSnapshotMessageBoundaryBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketOrdersSnapshotMessageBoundaryPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketOrdersSnapshotMessageBoundaryPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// MessageBoundary
func (m MarketOrdersSnapshotMessageBoundary) MessageBoundary() MessageSetBoundaryEnum {
	return MessageSetBoundaryEnum(message.Uint8Fixed(m.Unsafe(), 9, 8))
}

// MessageBoundary
func (m MarketOrdersSnapshotMessageBoundaryBuilder) MessageBoundary() MessageSetBoundaryEnum {
	return MessageSetBoundaryEnum(message.Uint8Fixed(m.Unsafe(), 9, 8))
}

// MessageBoundary
func (m MarketOrdersSnapshotMessageBoundaryPointer) MessageBoundary() MessageSetBoundaryEnum {
	return MessageSetBoundaryEnum(message.Uint8Fixed(m.Ptr, 9, 8))
}

// MessageBoundary
func (m MarketOrdersSnapshotMessageBoundaryPointerBuilder) MessageBoundary() MessageSetBoundaryEnum {
	return MessageSetBoundaryEnum(message.Uint8Fixed(m.Ptr, 9, 8))
}

// SetSymbolID
func (m MarketOrdersSnapshotMessageBoundaryBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID
func (m MarketOrdersSnapshotMessageBoundaryPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetMessageBoundary
func (m MarketOrdersSnapshotMessageBoundaryBuilder) SetMessageBoundary(value MessageSetBoundaryEnum) {
	message.SetUint8Fixed(m.Unsafe(), 9, 8, uint8(value))
}

// SetMessageBoundary
func (m MarketOrdersSnapshotMessageBoundaryPointerBuilder) SetMessageBoundary(value MessageSetBoundaryEnum) {
	message.SetUint8Fixed(m.Ptr, 9, 8, uint8(value))
}

// Copy
func (m MarketOrdersSnapshotMessageBoundary) Copy(to MarketOrdersSnapshotMessageBoundaryBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetMessageBoundary(m.MessageBoundary())
}

// Copy
func (m MarketOrdersSnapshotMessageBoundaryBuilder) Copy(to MarketOrdersSnapshotMessageBoundaryBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetMessageBoundary(m.MessageBoundary())
}

// CopyPointer
func (m MarketOrdersSnapshotMessageBoundary) CopyPointer(to MarketOrdersSnapshotMessageBoundaryPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetMessageBoundary(m.MessageBoundary())
}

// CopyPointer
func (m MarketOrdersSnapshotMessageBoundaryBuilder) CopyPointer(to MarketOrdersSnapshotMessageBoundaryPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetMessageBoundary(m.MessageBoundary())
}

// Copy
func (m MarketOrdersSnapshotMessageBoundaryPointer) Copy(to MarketOrdersSnapshotMessageBoundaryBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetMessageBoundary(m.MessageBoundary())
}

// Copy
func (m MarketOrdersSnapshotMessageBoundaryPointerBuilder) Copy(to MarketOrdersSnapshotMessageBoundaryBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetMessageBoundary(m.MessageBoundary())
}

// CopyPointer
func (m MarketOrdersSnapshotMessageBoundaryPointer) CopyPointer(to MarketOrdersSnapshotMessageBoundaryPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetMessageBoundary(m.MessageBoundary())
}

// CopyPointer
func (m MarketOrdersSnapshotMessageBoundaryPointerBuilder) CopyPointer(to MarketOrdersSnapshotMessageBoundaryPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetMessageBoundary(m.MessageBoundary())
}
