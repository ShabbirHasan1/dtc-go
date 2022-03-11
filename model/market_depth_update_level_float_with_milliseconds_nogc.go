package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDepthUpdateLevelFloatWithMillisecondsPointer Sent by the Server to the Client to Update/Insert or Delete a particular
// market depth price level in the market depth book maintained by the Client.
// market depth price level in the market depth book maintained by the Client.
//
// This message is a more compact version of the MarketDepthUpdateLevel message.
// For the Price and Quantity fields, it uses a 4 byte float for compactness.
// It also supports millisecond precision for the timestamp.
type MarketDepthUpdateLevelFloatWithMillisecondsPointer struct {
	message.FixedPointer
}

// MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder Sent by the Server to the Client to Update/Insert or Delete a particular
// market depth price level in the market depth book maintained by the Client.
// market depth price level in the market depth book maintained by the Client.
//
// This message is a more compact version of the MarketDepthUpdateLevel message.
// For the Price and Quantity fields, it uses a 4 byte float for compactness.
// It also supports millisecond precision for the timestamp.
type MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDepthUpdateLevelFloatWithMilliseconds() MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder {
	a := MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder{message.AllocFixed(29)}
	a.Ptr.SetBytes(0, _MarketDepthUpdateLevelFloatWithMillisecondsDefault)
	return a
}

func AllocMarketDepthUpdateLevelFloatWithMillisecondsFrom(b []byte) MarketDepthUpdateLevelFloatWithMillisecondsPointer {
	return MarketDepthUpdateLevelFloatWithMillisecondsPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size               = MarketDepthUpdateLevelFloatWithMillisecondsSize  (29)
//     Type               = MARKET_DEPTH_UPDATE_LEVEL_FLOAT_WITH_MILLISECONDS  (140)
//     SymbolID           = 0
//     DateTime           = 0
//     Price              = 0
//     Quantity           = 0
//     Side               = 0
//     UpdateType         = 0
//     NumOrders          = 0
//     FinalUpdateInBatch = 0
// }
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthUpdateLevelFloatWithMillisecondsDefault)
}

// ToBuilder
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) ToBuilder() MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder {
	return MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) Finish() MarketDepthUpdateLevelFloatWithMillisecondsPointer {
	return MarketDepthUpdateLevelFloatWithMillisecondsPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// DateTime The Date-Time of the market depth update with millisecond precision.
//
// This is an integer representing the number of milliseconds since the UNIX
// Epoch.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) DateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 16, 8))
}

// DateTime The Date-Time of the market depth update with millisecond precision.
//
// This is an integer representing the number of milliseconds since the UNIX
// Epoch.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) DateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 16, 8))
}

// Price The price level to insert, update or delete.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) Price() float32 {
	return message.Float32Fixed(m.Ptr, 20, 16)
}

// Price The price level to insert, update or delete.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) Price() float32 {
	return message.Float32Fixed(m.Ptr, 20, 16)
}

// Quantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) Quantity() float32 {
	return message.Float32Fixed(m.Ptr, 24, 20)
}

// Quantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) Quantity() float32 {
	return message.Float32Fixed(m.Ptr, 24, 20)
}

// Side Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) Side() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Ptr, 25, 24))
}

// Side Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) Side() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Ptr, 25, 24))
}

// UpdateType Specifies whether this is a MARKET_DEPTH_INSERT_UPDATE_LEVEL operation
// or a MARKET_DEPTH_DELETE_LEVEL operation.
//
// MARKET_DEPTH_INSERT_UPDATE_LEVEL: Insert or update in the market depth
// book on the specified side, the particular Price and Volume specified.
// It is an insert operation of the price level does not exist. It is an
// update operation if the price level already exists. In the case of insert,
// the other levels in the market depth book need to be shifted to make room
// for the new level.
//
// MARKET_DEPTH_DELETE_LEVEL: Remove from the market depth book on the specified
// side, the specified Price level. The other levels need to be shifted to
// fill in the missing level. In this particular case the Quantity is ignored
// and will be 0.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) UpdateType() MarketDepthUpdateTypeEnum {
	return MarketDepthUpdateTypeEnum(message.Uint8Fixed(m.Ptr, 26, 25))
}

// UpdateType Specifies whether this is a MARKET_DEPTH_INSERT_UPDATE_LEVEL operation
// or a MARKET_DEPTH_DELETE_LEVEL operation.
//
// MARKET_DEPTH_INSERT_UPDATE_LEVEL: Insert or update in the market depth
// book on the specified side, the particular Price and Volume specified.
// It is an insert operation of the price level does not exist. It is an
// update operation if the price level already exists. In the case of insert,
// the other levels in the market depth book need to be shifted to make room
// for the new level.
//
// MARKET_DEPTH_DELETE_LEVEL: Remove from the market depth book on the specified
// side, the specified Price level. The other levels need to be shifted to
// fill in the missing level. In this particular case the Quantity is ignored
// and will be 0.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) UpdateType() MarketDepthUpdateTypeEnum {
	return MarketDepthUpdateTypeEnum(message.Uint8Fixed(m.Ptr, 26, 25))
}

// NumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) NumOrders() uint16 {
	return message.Uint16Fixed(m.Ptr, 28, 26)
}

// NumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) NumOrders() uint16 {
	return message.Uint16Fixed(m.Ptr, 28, 26)
}

// FinalUpdateInBatch An indicator whether this is the final message or not in a batch of updates.
// An indicator whether this is the final message or not in a batch of updates.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) FinalUpdateInBatch() FinalUpdateInBatchEnum {
	return FinalUpdateInBatchEnum(message.Uint8Fixed(m.Ptr, 29, 28))
}

// FinalUpdateInBatch An indicator whether this is the final message or not in a batch of updates.
// An indicator whether this is the final message or not in a batch of updates.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) FinalUpdateInBatch() FinalUpdateInBatchEnum {
	return FinalUpdateInBatchEnum(message.Uint8Fixed(m.Ptr, 29, 28))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetDateTime The Date-Time of the market depth update with millisecond precision.
//
// This is an integer representing the number of milliseconds since the UNIX
// Epoch.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) SetDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
}

// SetPrice The price level to insert, update or delete.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 20, 16, value)
}

// SetQuantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) SetQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 24, 20, value)
}

// SetSide Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) SetSide(value AtBidOrAskEnum8) {
	message.SetUint8Fixed(m.Ptr, 25, 24, uint8(value))
}

// SetUpdateType Specifies whether this is a MARKET_DEPTH_INSERT_UPDATE_LEVEL operation
// or a MARKET_DEPTH_DELETE_LEVEL operation.
//
// MARKET_DEPTH_INSERT_UPDATE_LEVEL: Insert or update in the market depth
// book on the specified side, the particular Price and Volume specified.
// It is an insert operation of the price level does not exist. It is an
// update operation if the price level already exists. In the case of insert,
// the other levels in the market depth book need to be shifted to make room
// for the new level.
//
// MARKET_DEPTH_DELETE_LEVEL: Remove from the market depth book on the specified
// side, the specified Price level. The other levels need to be shifted to
// fill in the missing level. In this particular case the Quantity is ignored
// and will be 0.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) SetUpdateType(value MarketDepthUpdateTypeEnum) {
	message.SetUint8Fixed(m.Ptr, 26, 25, uint8(value))
}

// SetNumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) SetNumOrders(value uint16) {
	message.SetUint16Fixed(m.Ptr, 28, 26, value)
}

// SetFinalUpdateInBatch An indicator whether this is the final message or not in a batch of updates.
// An indicator whether this is the final message or not in a batch of updates.
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) SetFinalUpdateInBatch(value FinalUpdateInBatchEnum) {
	message.SetUint8Fixed(m.Ptr, 29, 28, uint8(value))
}

// Copy
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) Copy(to MarketDepthUpdateLevelFloatWithMillisecondsBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetSide(m.Side())
	to.SetUpdateType(m.UpdateType())
	to.SetNumOrders(m.NumOrders())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// Copy
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) Copy(to MarketDepthUpdateLevelFloatWithMillisecondsBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetSide(m.Side())
	to.SetUpdateType(m.UpdateType())
	to.SetNumOrders(m.NumOrders())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// CopyPointer
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointer) CopyPointer(to MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetSide(m.Side())
	to.SetUpdateType(m.UpdateType())
	to.SetNumOrders(m.NumOrders())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// CopyPointer
func (m MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) CopyPointer(to MarketDepthUpdateLevelFloatWithMillisecondsPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetSide(m.Side())
	to.SetUpdateType(m.UpdateType())
	to.SetNumOrders(m.NumOrders())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}
