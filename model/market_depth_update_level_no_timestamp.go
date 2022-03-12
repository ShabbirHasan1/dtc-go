package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDepthUpdateLevelNoTimestampSize = 21

// {
//     Size               = MarketDepthUpdateLevelNoTimestampSize  (21)
//     Type               = MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP  (141)
//     SymbolID           = 0
//     Price              = 0
//     Quantity           = 0
//     NumOrders          = 0
//     Side               = 0
//     UpdateType         = 0
//     FinalUpdateInBatch = FINAL_UPDATE_UNSET  (0)
// }
var _MarketDepthUpdateLevelNoTimestampDefault = []byte{21, 0, 141, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDepthUpdateLevelNoTimestamp Sent by the Server to the Client to Update/Insert or Delete a particular
// market depth price level in the market depth book maintained by the Client.
// market depth price level in the market depth book maintained by the Client.
//
// This message is identical to the MarketDepthUpdateLevel message except
// it has no timestamp field. It needs to be sent when there is no change
// with the timestamp for the market depth update as compared to the prior
// update.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received market depth update timestamp to know what the
// timestamp is for this message.
type MarketDepthUpdateLevelNoTimestamp struct {
	message.Fixed
}

// MarketDepthUpdateLevelNoTimestampBuilder Sent by the Server to the Client to Update/Insert or Delete a particular
// market depth price level in the market depth book maintained by the Client.
// market depth price level in the market depth book maintained by the Client.
//
// This message is identical to the MarketDepthUpdateLevel message except
// it has no timestamp field. It needs to be sent when there is no change
// with the timestamp for the market depth update as compared to the prior
// update.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received market depth update timestamp to know what the
// timestamp is for this message.
type MarketDepthUpdateLevelNoTimestampBuilder struct {
	message.Fixed
}

// MarketDepthUpdateLevelNoTimestampPointer Sent by the Server to the Client to Update/Insert or Delete a particular
// market depth price level in the market depth book maintained by the Client.
// market depth price level in the market depth book maintained by the Client.
//
// This message is identical to the MarketDepthUpdateLevel message except
// it has no timestamp field. It needs to be sent when there is no change
// with the timestamp for the market depth update as compared to the prior
// update.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received market depth update timestamp to know what the
// timestamp is for this message.
type MarketDepthUpdateLevelNoTimestampPointer struct {
	message.FixedPointer
}

// MarketDepthUpdateLevelNoTimestampPointerBuilder Sent by the Server to the Client to Update/Insert or Delete a particular
// market depth price level in the market depth book maintained by the Client.
// market depth price level in the market depth book maintained by the Client.
//
// This message is identical to the MarketDepthUpdateLevel message except
// it has no timestamp field. It needs to be sent when there is no change
// with the timestamp for the market depth update as compared to the prior
// update.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received market depth update timestamp to know what the
// timestamp is for this message.
type MarketDepthUpdateLevelNoTimestampPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDepthUpdateLevelNoTimestampFrom(b []byte) MarketDepthUpdateLevelNoTimestamp {
	return MarketDepthUpdateLevelNoTimestamp{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDepthUpdateLevelNoTimestamp(b []byte) MarketDepthUpdateLevelNoTimestamp {
	return MarketDepthUpdateLevelNoTimestamp{Fixed: message.WrapFixed(b)}
}

func NewMarketDepthUpdateLevelNoTimestamp() MarketDepthUpdateLevelNoTimestampBuilder {
	a := MarketDepthUpdateLevelNoTimestampBuilder{message.NewFixed(21)}
	a.Unsafe().SetBytes(0, _MarketDepthUpdateLevelNoTimestampDefault)
	return a
}

func AllocMarketDepthUpdateLevelNoTimestamp() MarketDepthUpdateLevelNoTimestampPointerBuilder {
	a := MarketDepthUpdateLevelNoTimestampPointerBuilder{message.AllocFixed(21)}
	a.Ptr.SetBytes(0, _MarketDepthUpdateLevelNoTimestampDefault)
	return a
}

func AllocMarketDepthUpdateLevelNoTimestampFrom(b []byte) MarketDepthUpdateLevelNoTimestampPointer {
	return MarketDepthUpdateLevelNoTimestampPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size               = MarketDepthUpdateLevelNoTimestampSize  (21)
//     Type               = MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP  (141)
//     SymbolID           = 0
//     Price              = 0
//     Quantity           = 0
//     NumOrders          = 0
//     Side               = 0
//     UpdateType         = 0
//     FinalUpdateInBatch = FINAL_UPDATE_UNSET  (0)
// }
func (m MarketDepthUpdateLevelNoTimestampBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthUpdateLevelNoTimestampDefault)
}

// Clear
// {
//     Size               = MarketDepthUpdateLevelNoTimestampSize  (21)
//     Type               = MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP  (141)
//     SymbolID           = 0
//     Price              = 0
//     Quantity           = 0
//     NumOrders          = 0
//     Side               = 0
//     UpdateType         = 0
//     FinalUpdateInBatch = FINAL_UPDATE_UNSET  (0)
// }
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthUpdateLevelNoTimestampDefault)
}

// ToBuilder
func (m MarketDepthUpdateLevelNoTimestamp) ToBuilder() MarketDepthUpdateLevelNoTimestampBuilder {
	return MarketDepthUpdateLevelNoTimestampBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDepthUpdateLevelNoTimestampPointer) ToBuilder() MarketDepthUpdateLevelNoTimestampPointerBuilder {
	return MarketDepthUpdateLevelNoTimestampPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDepthUpdateLevelNoTimestampBuilder) Finish() MarketDepthUpdateLevelNoTimestamp {
	return MarketDepthUpdateLevelNoTimestamp{m.Fixed}
}

// Finish
func (m *MarketDepthUpdateLevelNoTimestampPointerBuilder) Finish() MarketDepthUpdateLevelNoTimestampPointer {
	return MarketDepthUpdateLevelNoTimestampPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevelNoTimestamp) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevelNoTimestampBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevelNoTimestampPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Price The price level to insert, update or delete.
func (m MarketDepthUpdateLevelNoTimestamp) Price() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// Price The price level to insert, update or delete.
func (m MarketDepthUpdateLevelNoTimestampBuilder) Price() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// Price The price level to insert, update or delete.
func (m MarketDepthUpdateLevelNoTimestampPointer) Price() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// Price The price level to insert, update or delete.
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) Price() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// Quantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevelNoTimestamp) Quantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 16, 12)
}

// Quantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevelNoTimestampBuilder) Quantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 16, 12)
}

// Quantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevelNoTimestampPointer) Quantity() float32 {
	return message.Float32Fixed(m.Ptr, 16, 12)
}

// Quantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) Quantity() float32 {
	return message.Float32Fixed(m.Ptr, 16, 12)
}

// NumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevelNoTimestamp) NumOrders() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 18, 16)
}

// NumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevelNoTimestampBuilder) NumOrders() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 18, 16)
}

// NumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevelNoTimestampPointer) NumOrders() uint16 {
	return message.Uint16Fixed(m.Ptr, 18, 16)
}

// NumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) NumOrders() uint16 {
	return message.Uint16Fixed(m.Ptr, 18, 16)
}

// Side Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevelNoTimestamp) Side() int8 {
	return message.Int8Fixed(m.Unsafe(), 19, 18)
}

// Side Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevelNoTimestampBuilder) Side() int8 {
	return message.Int8Fixed(m.Unsafe(), 19, 18)
}

// Side Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevelNoTimestampPointer) Side() int8 {
	return message.Int8Fixed(m.Ptr, 19, 18)
}

// Side Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) Side() int8 {
	return message.Int8Fixed(m.Ptr, 19, 18)
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
func (m MarketDepthUpdateLevelNoTimestamp) UpdateType() int8 {
	return message.Int8Fixed(m.Unsafe(), 20, 19)
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
func (m MarketDepthUpdateLevelNoTimestampBuilder) UpdateType() int8 {
	return message.Int8Fixed(m.Unsafe(), 20, 19)
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
func (m MarketDepthUpdateLevelNoTimestampPointer) UpdateType() int8 {
	return message.Int8Fixed(m.Ptr, 20, 19)
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
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) UpdateType() int8 {
	return message.Int8Fixed(m.Ptr, 20, 19)
}

// FinalUpdateInBatch An indicator whether this is the beginning or final update in a batch
// of updates.
func (m MarketDepthUpdateLevelNoTimestamp) FinalUpdateInBatch() FinalUpdateInBatchEnum {
	return FinalUpdateInBatchEnum(message.Uint8Fixed(m.Unsafe(), 21, 20))
}

// FinalUpdateInBatch An indicator whether this is the beginning or final update in a batch
// of updates.
func (m MarketDepthUpdateLevelNoTimestampBuilder) FinalUpdateInBatch() FinalUpdateInBatchEnum {
	return FinalUpdateInBatchEnum(message.Uint8Fixed(m.Unsafe(), 21, 20))
}

// FinalUpdateInBatch An indicator whether this is the beginning or final update in a batch
// of updates.
func (m MarketDepthUpdateLevelNoTimestampPointer) FinalUpdateInBatch() FinalUpdateInBatchEnum {
	return FinalUpdateInBatchEnum(message.Uint8Fixed(m.Ptr, 21, 20))
}

// FinalUpdateInBatch An indicator whether this is the beginning or final update in a batch
// of updates.
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) FinalUpdateInBatch() FinalUpdateInBatchEnum {
	return FinalUpdateInBatchEnum(message.Uint8Fixed(m.Ptr, 21, 20))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevelNoTimestampBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetPrice The price level to insert, update or delete.
func (m MarketDepthUpdateLevelNoTimestampBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 12, 8, value)
}

// SetPrice The price level to insert, update or delete.
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 12, 8, value)
}

// SetQuantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevelNoTimestampBuilder) SetQuantity(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 16, 12, value)
}

// SetQuantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) SetQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 16, 12, value)
}

// SetNumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevelNoTimestampBuilder) SetNumOrders(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 18, 16, value)
}

// SetNumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) SetNumOrders(value uint16) {
	message.SetUint16Fixed(m.Ptr, 18, 16, value)
}

// SetSide Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevelNoTimestampBuilder) SetSide(value int8) {
	message.SetInt8Fixed(m.Unsafe(), 19, 18, value)
}

// SetSide Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) SetSide(value int8) {
	message.SetInt8Fixed(m.Ptr, 19, 18, value)
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
func (m MarketDepthUpdateLevelNoTimestampBuilder) SetUpdateType(value int8) {
	message.SetInt8Fixed(m.Unsafe(), 20, 19, value)
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
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) SetUpdateType(value int8) {
	message.SetInt8Fixed(m.Ptr, 20, 19, value)
}

// SetFinalUpdateInBatch An indicator whether this is the beginning or final update in a batch
// of updates.
func (m MarketDepthUpdateLevelNoTimestampBuilder) SetFinalUpdateInBatch(value FinalUpdateInBatchEnum) {
	message.SetUint8Fixed(m.Unsafe(), 21, 20, uint8(value))
}

// SetFinalUpdateInBatch An indicator whether this is the beginning or final update in a batch
// of updates.
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) SetFinalUpdateInBatch(value FinalUpdateInBatchEnum) {
	message.SetUint8Fixed(m.Ptr, 21, 20, uint8(value))
}

// Copy
func (m MarketDepthUpdateLevelNoTimestamp) Copy(to MarketDepthUpdateLevelNoTimestampBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetSide(m.Side())
	to.SetUpdateType(m.UpdateType())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// Copy
func (m MarketDepthUpdateLevelNoTimestampBuilder) Copy(to MarketDepthUpdateLevelNoTimestampBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetSide(m.Side())
	to.SetUpdateType(m.UpdateType())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// CopyPointer
func (m MarketDepthUpdateLevelNoTimestamp) CopyPointer(to MarketDepthUpdateLevelNoTimestampPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetSide(m.Side())
	to.SetUpdateType(m.UpdateType())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// CopyPointer
func (m MarketDepthUpdateLevelNoTimestampBuilder) CopyPointer(to MarketDepthUpdateLevelNoTimestampPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetSide(m.Side())
	to.SetUpdateType(m.UpdateType())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// Copy
func (m MarketDepthUpdateLevelNoTimestampPointer) Copy(to MarketDepthUpdateLevelNoTimestampBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetSide(m.Side())
	to.SetUpdateType(m.UpdateType())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// Copy
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) Copy(to MarketDepthUpdateLevelNoTimestampBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetSide(m.Side())
	to.SetUpdateType(m.UpdateType())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// CopyPointer
func (m MarketDepthUpdateLevelNoTimestampPointer) CopyPointer(to MarketDepthUpdateLevelNoTimestampPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetSide(m.Side())
	to.SetUpdateType(m.UpdateType())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}

// CopyPointer
func (m MarketDepthUpdateLevelNoTimestampPointerBuilder) CopyPointer(to MarketDepthUpdateLevelNoTimestampPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetNumOrders(m.NumOrders())
	to.SetSide(m.Side())
	to.SetUpdateType(m.UpdateType())
	to.SetFinalUpdateInBatch(m.FinalUpdateInBatch())
}
