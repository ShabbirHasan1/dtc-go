package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDepthSnapshotLevelFloatPointer This is a message sent by Server to provide the initial market depth data
// entries to the Client after the Client subscribes to market data or separately
// subscribes to market depth data. The Client will need to separately subscribe
// to market depth data if the Server requires it.
//
// Each message provides a single entry of depth data. Therefore, the Server
// will send multiple MarketDepthSnapshotLevelFloat messages in a series
// in order for the Client to build up its initial market depth book.
//
// The first message will be identified by the IsFirstMessageInBatch field
// being set to 1. The last message will be identified by the IsLastMessageInBatch
// field being set to 1.
//
// In the case where the market depth book is empty, the Server still needs
// to send through one single message with the SymbolID set, IsFirstMessageInBatch
// equal to 1 and IsLastMessageInBatch equal to 1. All other members will
// be at the default values. The Client will understand this as an empty
// book.
type MarketDepthSnapshotLevelFloatPointer struct {
	message.FixedPointer
}

// MarketDepthSnapshotLevelFloatPointerBuilder This is a message sent by Server to provide the initial market depth data
// entries to the Client after the Client subscribes to market data or separately
// subscribes to market depth data. The Client will need to separately subscribe
// to market depth data if the Server requires it.
//
// Each message provides a single entry of depth data. Therefore, the Server
// will send multiple MarketDepthSnapshotLevelFloat messages in a series
// in order for the Client to build up its initial market depth book.
//
// The first message will be identified by the IsFirstMessageInBatch field
// being set to 1. The last message will be identified by the IsLastMessageInBatch
// field being set to 1.
//
// In the case where the market depth book is empty, the Server still needs
// to send through one single message with the SymbolID set, IsFirstMessageInBatch
// equal to 1 and IsLastMessageInBatch equal to 1. All other members will
// be at the default values. The Client will understand this as an empty
// book.
type MarketDepthSnapshotLevelFloatPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDepthSnapshotLevelFloat() MarketDepthSnapshotLevelFloatPointerBuilder {
	a := MarketDepthSnapshotLevelFloatPointerBuilder{message.AllocFixed(24)}
	a.Ptr.SetBytes(0, _MarketDepthSnapshotLevelFloatDefault)
	return a
}

func AllocMarketDepthSnapshotLevelFloatFrom(b []byte) MarketDepthSnapshotLevelFloatPointer {
	return MarketDepthSnapshotLevelFloatPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size               = MarketDepthSnapshotLevelFloatSize  (24)
//     Type               = MARKET_DEPTH_SNAPSHOT_LEVEL_FLOAT  (145)
//     SymbolID           = 0
//     Price              = 0
//     Quantity           = 0
//     NumOrders          = 0
//     Level              = 0
//     Side               = 0
//     FinalUpdateInBatch = 0
// }
func (m MarketDepthSnapshotLevelFloatPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthSnapshotLevelFloatDefault)
}

// ToBuilder
func (m MarketDepthSnapshotLevelFloatPointer) ToBuilder() MarketDepthSnapshotLevelFloatPointerBuilder {
	return MarketDepthSnapshotLevelFloatPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDepthSnapshotLevelFloatPointerBuilder) Finish() MarketDepthSnapshotLevelFloatPointer {
	return MarketDepthSnapshotLevelFloatPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelFloatPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Price This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelFloatPointer) Price() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// Price This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) Price() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// Quantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelFloatPointer) Quantity() float32 {
	return message.Float32Fixed(m.Ptr, 16, 12)
}

// Quantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) Quantity() float32 {
	return message.Float32Fixed(m.Ptr, 16, 12)
}

// NumOrders The number of orders at the Price.
func (m MarketDepthSnapshotLevelFloatPointer) NumOrders() uint32 {
	return message.Uint32Fixed(m.Ptr, 20, 16)
}

// NumOrders The number of orders at the Price.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) NumOrders() uint32 {
	return message.Uint32Fixed(m.Ptr, 20, 16)
}

// Level This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelFloatPointer) Level() uint16 {
	return message.Uint16Fixed(m.Ptr, 22, 20)
}

// Level This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) Level() uint16 {
	return message.Uint16Fixed(m.Ptr, 22, 20)
}

// Side Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelFloatPointer) Side() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Ptr, 23, 22))
}

// Side Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) Side() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Ptr, 23, 22))
}

// FinalUpdateInBatch An indicator whether this is the final message or not in a batch of updates.
// An indicator whether this is the final message or not in a batch of updates.
func (m MarketDepthSnapshotLevelFloatPointer) FinalUpdateInBatch() FinalUpdateInBatchEnum {
	return FinalUpdateInBatchEnum(message.Uint8Fixed(m.Ptr, 24, 23))
}

// FinalUpdateInBatch An indicator whether this is the final message or not in a batch of updates.
// An indicator whether this is the final message or not in a batch of updates.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) FinalUpdateInBatch() FinalUpdateInBatchEnum {
	return FinalUpdateInBatchEnum(message.Uint8Fixed(m.Ptr, 24, 23))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetPrice This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 12, 8, value)
}

// SetQuantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) SetQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 16, 12, value)
}

// SetNumOrders The number of orders at the Price.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) SetNumOrders(value uint32) {
	message.SetUint32Fixed(m.Ptr, 20, 16, value)
}

// SetLevel This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) SetLevel(value uint16) {
	message.SetUint16Fixed(m.Ptr, 22, 20, value)
}

// SetSide Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) SetSide(value AtBidOrAskEnum8) {
	message.SetUint8Fixed(m.Ptr, 23, 22, uint8(value))
}

// SetFinalUpdateInBatch An indicator whether this is the final message or not in a batch of updates.
// An indicator whether this is the final message or not in a batch of updates.
func (m MarketDepthSnapshotLevelFloatPointerBuilder) SetFinalUpdateInBatch(value FinalUpdateInBatchEnum) {
	message.SetUint8Fixed(m.Ptr, 24, 23, uint8(value))
}

// Copy
func (m MarketDepthSnapshotLevelFloatPointer) Copy(to MarketDepthSnapshotLevelFloatBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetLevel(m.Level())
	to.SetSide(m.Side())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// Copy
func (m MarketDepthSnapshotLevelFloatPointerBuilder) Copy(to MarketDepthSnapshotLevelFloatBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetLevel(m.Level())
	to.SetSide(m.Side())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// CopyPointer
func (m MarketDepthSnapshotLevelFloatPointer) CopyPointer(to MarketDepthSnapshotLevelFloatPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetLevel(m.Level())
	to.SetSide(m.Side())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// CopyPointer
func (m MarketDepthSnapshotLevelFloatPointerBuilder) CopyPointer(to MarketDepthSnapshotLevelFloatPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetLevel(m.Level())
	to.SetSide(m.Side())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}
