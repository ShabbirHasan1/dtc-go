package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                  = MarketDepthSnapshotLevelSize  (56)
//     Type                  = MARKET_DEPTH_SNAPSHOT_LEVEL  (122)
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
var _MarketDepthSnapshotLevelDefault = []byte{56, 0, 122, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDepthSnapshotLevelSize = 56

// MarketDepthSnapshotLevel This is a message sent by Server to provide the initial market depth data
// entries to the Client after the Client subscribes to market data or separately
// subscribes to market depth data. The Client will need to separately subscribe
// to market depth data if the Server requires it.
//
// Each message provides a single entry of depth data. Therefore, the Server
// will send multiple MarketDepthSnapshotLevel messages in a series in order
// for the Client to build up its initial market depth book.
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
type MarketDepthSnapshotLevel struct {
	message.Fixed
}

// MarketDepthSnapshotLevelBuilder This is a message sent by Server to provide the initial market depth data
// entries to the Client after the Client subscribes to market data or separately
// subscribes to market depth data. The Client will need to separately subscribe
// to market depth data if the Server requires it.
//
// Each message provides a single entry of depth data. Therefore, the Server
// will send multiple MarketDepthSnapshotLevel messages in a series in order
// for the Client to build up its initial market depth book.
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
type MarketDepthSnapshotLevelBuilder struct {
	message.Fixed
}

func NewMarketDepthSnapshotLevelFrom(b []byte) MarketDepthSnapshotLevel {
	return MarketDepthSnapshotLevel{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDepthSnapshotLevel(b []byte) MarketDepthSnapshotLevel {
	return MarketDepthSnapshotLevel{Fixed: message.WrapFixed(b)}
}

func NewMarketDepthSnapshotLevel() MarketDepthSnapshotLevelBuilder {
	a := MarketDepthSnapshotLevelBuilder{message.NewFixed(56)}
	a.Unsafe().SetBytes(0, _MarketDepthSnapshotLevelDefault)
	return a
}

// Clear
// {
//     Size                  = MarketDepthSnapshotLevelSize  (56)
//     Type                  = MARKET_DEPTH_SNAPSHOT_LEVEL  (122)
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
func (m MarketDepthSnapshotLevelBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthSnapshotLevelDefault)
}

// ToBuilder
func (m MarketDepthSnapshotLevel) ToBuilder() MarketDepthSnapshotLevelBuilder {
	return MarketDepthSnapshotLevelBuilder{m.Fixed}
}

// Finish
func (m MarketDepthSnapshotLevelBuilder) Finish() MarketDepthSnapshotLevel {
	return MarketDepthSnapshotLevel{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevel) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Side Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevel) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 10, 8))
}

// Side Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelBuilder) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 10, 8))
}

// Price This is the price of the market depth entry.
func (m MarketDepthSnapshotLevel) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// Price This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// Quantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevel) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// Quantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// Level This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevel) Level() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 34, 32)
}

// Level This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelBuilder) Level() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 34, 32)
}

// IsFirstMessageInBatch Set to 1 if this is the first message in the batch of messages.
func (m MarketDepthSnapshotLevel) IsFirstMessageInBatch() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 35, 34)
}

// IsFirstMessageInBatch Set to 1 if this is the first message in the batch of messages.
func (m MarketDepthSnapshotLevelBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 35, 34)
}

// IsLastMessageInBatch Set to 1 if this is the last message in a batch of messages. If there
// is only a single message to be sent, in case the market depth book is
// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
// will equal 1.
func (m MarketDepthSnapshotLevel) IsLastMessageInBatch() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 36, 35)
}

// IsLastMessageInBatch Set to 1 if this is the last message in a batch of messages. If there
// is only a single message to be sent, in case the market depth book is
// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
// will equal 1.
func (m MarketDepthSnapshotLevelBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 36, 35)
}

// DateTime
func (m MarketDepthSnapshotLevel) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 48, 40))
}

// DateTime
func (m MarketDepthSnapshotLevelBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 48, 40))
}

// NumOrders
func (m MarketDepthSnapshotLevel) NumOrders() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 52, 48)
}

// NumOrders
func (m MarketDepthSnapshotLevelBuilder) NumOrders() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 52, 48)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSide Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelBuilder) SetSide(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Unsafe(), 10, 8, uint16(value))
}

// SetPrice This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetQuantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 32, 24, value)
}

// SetLevel This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelBuilder) SetLevel(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 34, 32, value)
}

// SetIsFirstMessageInBatch Set to 1 if this is the first message in the batch of messages.
func (m MarketDepthSnapshotLevelBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 35, 34, value)
}

// SetIsLastMessageInBatch Set to 1 if this is the last message in a batch of messages. If there
// is only a single message to be sent, in case the market depth book is
// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
// will equal 1.
func (m MarketDepthSnapshotLevelBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 36, 35, value)
}

// SetDateTime
func (m MarketDepthSnapshotLevelBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 48, 40, float64(value))
}

// SetNumOrders
func (m MarketDepthSnapshotLevelBuilder) SetNumOrders(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 52, 48, value)
}

// Copy
func (m MarketDepthSnapshotLevel) Copy(to MarketDepthSnapshotLevelBuilder) {
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
func (m MarketDepthSnapshotLevelBuilder) Copy(to MarketDepthSnapshotLevelBuilder) {
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
func (m MarketDepthSnapshotLevel) CopyPointer(to MarketDepthSnapshotLevelPointerBuilder) {
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
func (m MarketDepthSnapshotLevelBuilder) CopyPointer(to MarketDepthSnapshotLevelPointerBuilder) {
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
