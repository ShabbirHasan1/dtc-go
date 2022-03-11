package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDepthSnapshotLevelSize = 56

// {
//     Size                  = MarketDepthSnapshotLevelSize  (56)
//     Type                  = MARKET_DEPTH_SNAPSHOT_LEVEL  (122)
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
var _MarketDepthSnapshotLevelDefault = []byte{56, 0, 122, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

// MarketDepthSnapshotLevelPointer This is a message sent by Server to provide the initial market depth data
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
type MarketDepthSnapshotLevelPointer struct {
	message.FixedPointer
}

// MarketDepthSnapshotLevelPointerBuilder This is a message sent by Server to provide the initial market depth data
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
type MarketDepthSnapshotLevelPointerBuilder struct {
	message.FixedPointer
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

func AllocMarketDepthSnapshotLevel() MarketDepthSnapshotLevelPointerBuilder {
	a := MarketDepthSnapshotLevelPointerBuilder{message.AllocFixed(56)}
	a.Ptr.SetBytes(0, _MarketDepthSnapshotLevelDefault)
	return a
}

func AllocMarketDepthSnapshotLevelFrom(b []byte) MarketDepthSnapshotLevelPointer {
	return MarketDepthSnapshotLevelPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                  = MarketDepthSnapshotLevelSize  (56)
//     Type                  = MARKET_DEPTH_SNAPSHOT_LEVEL  (122)
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
func (m MarketDepthSnapshotLevelBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthSnapshotLevelDefault)
}

// Clear
// {
//     Size                  = MarketDepthSnapshotLevelSize  (56)
//     Type                  = MARKET_DEPTH_SNAPSHOT_LEVEL  (122)
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
func (m MarketDepthSnapshotLevelPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthSnapshotLevelDefault)
}

// ToBuilder
func (m MarketDepthSnapshotLevel) ToBuilder() MarketDepthSnapshotLevelBuilder {
	return MarketDepthSnapshotLevelBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDepthSnapshotLevelPointer) ToBuilder() MarketDepthSnapshotLevelPointerBuilder {
	return MarketDepthSnapshotLevelPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDepthSnapshotLevelBuilder) Finish() MarketDepthSnapshotLevel {
	return MarketDepthSnapshotLevel{m.Fixed}
}

// Finish
func (m *MarketDepthSnapshotLevelPointerBuilder) Finish() MarketDepthSnapshotLevelPointer {
	return MarketDepthSnapshotLevelPointer{m.FixedPointer}
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

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
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

// Side Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelPointer) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// Side Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelPointerBuilder) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// Price This is the price of the market depth entry.
func (m MarketDepthSnapshotLevel) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// Price This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// Price This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// Price This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// Quantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevel) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// Quantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// Quantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelPointer) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
}

// Quantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelPointerBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
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

// Level This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelPointer) Level() uint16 {
	return message.Uint16Fixed(m.Ptr, 34, 32)
}

// Level This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelPointerBuilder) Level() uint16 {
	return message.Uint16Fixed(m.Ptr, 34, 32)
}

// IsFirstMessageInBatch Set to 1 if this is the first message in the batch of messages.
func (m MarketDepthSnapshotLevel) IsFirstMessageInBatch() bool {
	return message.BoolFixed(m.Unsafe(), 35, 34)
}

// IsFirstMessageInBatch Set to 1 if this is the first message in the batch of messages.
func (m MarketDepthSnapshotLevelBuilder) IsFirstMessageInBatch() bool {
	return message.BoolFixed(m.Unsafe(), 35, 34)
}

// IsFirstMessageInBatch Set to 1 if this is the first message in the batch of messages.
func (m MarketDepthSnapshotLevelPointer) IsFirstMessageInBatch() bool {
	return message.BoolFixed(m.Ptr, 35, 34)
}

// IsFirstMessageInBatch Set to 1 if this is the first message in the batch of messages.
func (m MarketDepthSnapshotLevelPointerBuilder) IsFirstMessageInBatch() bool {
	return message.BoolFixed(m.Ptr, 35, 34)
}

// IsLastMessageInBatch Set to 1 if this is the last message in a batch of messages. If there
// is only a single message to be sent, in case the market depth book is
// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
// will equal 1.
func (m MarketDepthSnapshotLevel) IsLastMessageInBatch() bool {
	return message.BoolFixed(m.Unsafe(), 36, 35)
}

// IsLastMessageInBatch Set to 1 if this is the last message in a batch of messages. If there
// is only a single message to be sent, in case the market depth book is
// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
// will equal 1.
func (m MarketDepthSnapshotLevelBuilder) IsLastMessageInBatch() bool {
	return message.BoolFixed(m.Unsafe(), 36, 35)
}

// IsLastMessageInBatch Set to 1 if this is the last message in a batch of messages. If there
// is only a single message to be sent, in case the market depth book is
// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
// will equal 1.
func (m MarketDepthSnapshotLevelPointer) IsLastMessageInBatch() bool {
	return message.BoolFixed(m.Ptr, 36, 35)
}

// IsLastMessageInBatch Set to 1 if this is the last message in a batch of messages. If there
// is only a single message to be sent, in case the market depth book is
// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
// will equal 1.
func (m MarketDepthSnapshotLevelPointerBuilder) IsLastMessageInBatch() bool {
	return message.BoolFixed(m.Ptr, 36, 35)
}

// DateTime
func (m MarketDepthSnapshotLevel) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 48, 40))
}

// DateTime
func (m MarketDepthSnapshotLevelBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 48, 40))
}

// DateTime
func (m MarketDepthSnapshotLevelPointer) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 48, 40))
}

// DateTime
func (m MarketDepthSnapshotLevelPointerBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 48, 40))
}

// NumOrders
func (m MarketDepthSnapshotLevel) NumOrders() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 52, 48)
}

// NumOrders
func (m MarketDepthSnapshotLevelBuilder) NumOrders() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 52, 48)
}

// NumOrders
func (m MarketDepthSnapshotLevelPointer) NumOrders() uint32 {
	return message.Uint32Fixed(m.Ptr, 52, 48)
}

// NumOrders
func (m MarketDepthSnapshotLevelPointerBuilder) NumOrders() uint32 {
	return message.Uint32Fixed(m.Ptr, 52, 48)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest/MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthSnapshotLevelPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetSide Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelBuilder) SetSide(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Unsafe(), 10, 8, uint16(value))
}

// SetSide Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
// = 2, if this is an ask side market depth entry.
func (m MarketDepthSnapshotLevelPointerBuilder) SetSide(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Ptr, 10, 8, uint16(value))
}

// SetPrice This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetPrice This is the price of the market depth entry.
func (m MarketDepthSnapshotLevelPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetQuantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 32, 24, value)
}

// SetQuantity This is the quantity of orders at the Price.
func (m MarketDepthSnapshotLevelPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, value)
}

// SetLevel This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelBuilder) SetLevel(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 34, 32, value)
}

// SetLevel This indicates the level of the price within the market depth book. The
// minimum value is 1. There is no maximum value. A value of 1 is considered
// the best bid or ask data.
func (m MarketDepthSnapshotLevelPointerBuilder) SetLevel(value uint16) {
	message.SetUint16Fixed(m.Ptr, 34, 32, value)
}

// SetIsFirstMessageInBatch Set to 1 if this is the first message in the batch of messages.
func (m MarketDepthSnapshotLevelBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolFixed(m.Unsafe(), 35, 34, value)
}

// SetIsFirstMessageInBatch Set to 1 if this is the first message in the batch of messages.
func (m MarketDepthSnapshotLevelPointerBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolFixed(m.Ptr, 35, 34, value)
}

// SetIsLastMessageInBatch Set to 1 if this is the last message in a batch of messages. If there
// is only a single message to be sent, in case the market depth book is
// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
// will equal 1.
func (m MarketDepthSnapshotLevelBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolFixed(m.Unsafe(), 36, 35, value)
}

// SetIsLastMessageInBatch Set to 1 if this is the last message in a batch of messages. If there
// is only a single message to be sent, in case the market depth book is
// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
// will equal 1.
func (m MarketDepthSnapshotLevelPointerBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolFixed(m.Ptr, 36, 35, value)
}

// SetDateTime
func (m MarketDepthSnapshotLevelBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 48, 40, float64(value))
}

// SetDateTime
func (m MarketDepthSnapshotLevelPointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 48, 40, float64(value))
}

// SetNumOrders
func (m MarketDepthSnapshotLevelBuilder) SetNumOrders(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 52, 48, value)
}

// SetNumOrders
func (m MarketDepthSnapshotLevelPointerBuilder) SetNumOrders(value uint32) {
	message.SetUint32Fixed(m.Ptr, 52, 48, value)
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

// Copy
func (m MarketDepthSnapshotLevelPointer) Copy(to MarketDepthSnapshotLevelBuilder) {
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
func (m MarketDepthSnapshotLevelPointerBuilder) Copy(to MarketDepthSnapshotLevelBuilder) {
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
func (m MarketDepthSnapshotLevelPointer) CopyPointer(to MarketDepthSnapshotLevelPointerBuilder) {
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
func (m MarketDepthSnapshotLevelPointerBuilder) CopyPointer(to MarketDepthSnapshotLevelPointerBuilder) {
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
