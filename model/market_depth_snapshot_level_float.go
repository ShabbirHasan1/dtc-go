package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _MarketDepthSnapshotLevelFloatDefault = []byte{24, 0, 145, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDepthSnapshotLevelFloatSize = 24

// MarketDepthSnapshotLevelFloat This is a message sent by Server to provide the initial market depth data
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
type MarketDepthSnapshotLevelFloat struct {
	message.Fixed
}

// MarketDepthSnapshotLevelFloatBuilder This is a message sent by Server to provide the initial market depth data
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
type MarketDepthSnapshotLevelFloatBuilder struct {
	message.Fixed
}

func NewMarketDepthSnapshotLevelFloatFrom(b []byte) MarketDepthSnapshotLevelFloat {
	return MarketDepthSnapshotLevelFloat{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDepthSnapshotLevelFloat(b []byte) MarketDepthSnapshotLevelFloat {
	return MarketDepthSnapshotLevelFloat{Fixed: message.WrapFixed(b)}
}

func NewMarketDepthSnapshotLevelFloat() MarketDepthSnapshotLevelFloatBuilder {
	a := MarketDepthSnapshotLevelFloatBuilder{message.NewFixed(24)}
	a.Unsafe().SetBytes(0, _MarketDepthSnapshotLevelFloatDefault)
	return a
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
func (m MarketDepthSnapshotLevelFloatBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthSnapshotLevelFloatDefault)
}

// ToBuilder
func (m MarketDepthSnapshotLevelFloat) ToBuilder() MarketDepthSnapshotLevelFloatBuilder {
	return MarketDepthSnapshotLevelFloatBuilder{m.Fixed}
}

// Finish
func (m MarketDepthSnapshotLevelFloatBuilder) Finish() MarketDepthSnapshotLevelFloat {
	return MarketDepthSnapshotLevelFloat{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelFloat) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelFloatBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Price This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelFloat) Price() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// Price This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelFloatBuilder) Price() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// Quantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelFloat) Quantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 16, 12)
}

// Quantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelFloatBuilder) Quantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 16, 12)
}

// NumOrders The number of orders at the Price.
func (m MarketDepthSnapshotLevelFloat) NumOrders() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 20, 16)
}

// NumOrders The number of orders at the Price.
func (m MarketDepthSnapshotLevelFloatBuilder) NumOrders() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 20, 16)
}

// Level This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelFloat) Level() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 22, 20)
}

// Level This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelFloatBuilder) Level() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 22, 20)
}

// Side Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelFloat) Side() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Unsafe(), 23, 22))
}

// Side Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelFloatBuilder) Side() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Unsafe(), 23, 22))
}

// FinalUpdateInBatch An indicator whether this is the final message or not in a batch of updates.
// An indicator whether this is the final message or not in a batch of updates.
func (m MarketDepthSnapshotLevelFloat) FinalUpdateInBatch() FinalUpdateInBatchEnum {
	return FinalUpdateInBatchEnum(message.Uint8Fixed(m.Unsafe(), 24, 23))
}

// FinalUpdateInBatch An indicator whether this is the final message or not in a batch of updates.
// An indicator whether this is the final message or not in a batch of updates.
func (m MarketDepthSnapshotLevelFloatBuilder) FinalUpdateInBatch() FinalUpdateInBatchEnum {
	return FinalUpdateInBatchEnum(message.Uint8Fixed(m.Unsafe(), 24, 23))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelFloatBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetPrice This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelFloatBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 12, 8, value)
}

// SetQuantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelFloatBuilder) SetQuantity(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 16, 12, value)
}

// SetNumOrders The number of orders at the Price.
func (m MarketDepthSnapshotLevelFloatBuilder) SetNumOrders(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 20, 16, value)
}

// SetLevel This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelFloatBuilder) SetLevel(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 22, 20, value)
}

// SetSide Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelFloatBuilder) SetSide(value AtBidOrAskEnum8) {
	message.SetUint8Fixed(m.Unsafe(), 23, 22, uint8(value))
}

// SetFinalUpdateInBatch An indicator whether this is the final message or not in a batch of updates.
// An indicator whether this is the final message or not in a batch of updates.
func (m MarketDepthSnapshotLevelFloatBuilder) SetFinalUpdateInBatch(value FinalUpdateInBatchEnum) {
	message.SetUint8Fixed(m.Unsafe(), 24, 23, uint8(value))
}

// Copy
func (m MarketDepthSnapshotLevelFloat) Copy(to MarketDepthSnapshotLevelFloatBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetLevel(m.Level())
	to.SetSide(m.Side())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// Copy
func (m MarketDepthSnapshotLevelFloatBuilder) Copy(to MarketDepthSnapshotLevelFloatBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetLevel(m.Level())
	to.SetSide(m.Side())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// CopyPointer
func (m MarketDepthSnapshotLevelFloat) CopyPointer(to MarketDepthSnapshotLevelFloatPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetLevel(m.Level())
	to.SetSide(m.Side())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// CopyPointer
func (m MarketDepthSnapshotLevelFloatBuilder) CopyPointer(to MarketDepthSnapshotLevelFloatPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetLevel(m.Level())
	to.SetSide(m.Side())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}
