package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDepthUpdateLevelPointer Sent by the Server to the Client to Update/Insert or Delete a particular
// market depth price level in the market depth book maintained by the Client.
// market depth price level in the market depth book maintained by the Client.
//
// Each MarketDepthUpdateLevel message updates one level of market depth
// on one side. An insert/update/delete model is used for market depth.
//
// The Client will need to determine the based upon the price, what particular
// market depth level is being updated, inserted or deleted.
//
// It is for this reason, that an insert/update is considered as one update
// type since it is possible to determine whether it is an insert or update
// based upon the existence of the price level in the existing market depth
// book on the Client side.
//
// What this means is that when the UpdateType field is MARKET_DEPTH_INSERT_UPDATE_LEVEL,
// it is considered an insert if the price level is not found on the particular
// side of the market depth being updated. It is considered an update, if
// the price level is found on the particular side of market depth being
// updated.
//
// This message uses a double datatype for the Price field. There is no level
// index. It is the responsibility of the Client to determine where in its
// market depth array it is maintaining where the insert/update/delete operation
// needs to occur.
//
// Since floating-point comparisons are not always precise, there should
// be a comparison made only to the number of decimal places the symbol specifies
// in its security definition. This can be determined through the SecurityDefinitionResponse::PriceDisplayFormat
// field.
type MarketDepthUpdateLevelPointer struct {
	message.FixedPointer
}

// MarketDepthUpdateLevelPointerBuilder Sent by the Server to the Client to Update/Insert or Delete a particular
// market depth price level in the market depth book maintained by the Client.
// market depth price level in the market depth book maintained by the Client.
//
// Each MarketDepthUpdateLevel message updates one level of market depth
// on one side. An insert/update/delete model is used for market depth.
//
// The Client will need to determine the based upon the price, what particular
// market depth level is being updated, inserted or deleted.
//
// It is for this reason, that an insert/update is considered as one update
// type since it is possible to determine whether it is an insert or update
// based upon the existence of the price level in the existing market depth
// book on the Client side.
//
// What this means is that when the UpdateType field is MARKET_DEPTH_INSERT_UPDATE_LEVEL,
// it is considered an insert if the price level is not found on the particular
// side of the market depth being updated. It is considered an update, if
// the price level is found on the particular side of market depth being
// updated.
//
// This message uses a double datatype for the Price field. There is no level
// index. It is the responsibility of the Client to determine where in its
// market depth array it is maintaining where the insert/update/delete operation
// needs to occur.
//
// Since floating-point comparisons are not always precise, there should
// be a comparison made only to the number of decimal places the symbol specifies
// in its security definition. This can be determined through the SecurityDefinitionResponse::PriceDisplayFormat
// field.
type MarketDepthUpdateLevelPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDepthUpdateLevel() MarketDepthUpdateLevelPointerBuilder {
	a := MarketDepthUpdateLevelPointerBuilder{message.AllocFixed(56)}
	a.Ptr.SetBytes(0, _MarketDepthUpdateLevelDefault)
	return a
}

func AllocMarketDepthUpdateLevelFrom(b []byte) MarketDepthUpdateLevelPointer {
	return MarketDepthUpdateLevelPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = MarketDepthUpdateLevelSize  (56)
//     Type       = MARKET_DEPTH_UPDATE_LEVEL  (106)
//     SymbolID   = 0
//     Side       = 0
//     Price      = 0
//     Quantity   = 0
//     UpdateType = 0
//     DateTime   = 0
//     NumOrders  = 0
// }
func (m MarketDepthUpdateLevelPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthUpdateLevelDefault)
}

// ToBuilder
func (m MarketDepthUpdateLevelPointer) ToBuilder() MarketDepthUpdateLevelPointerBuilder {
	return MarketDepthUpdateLevelPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDepthUpdateLevelPointerBuilder) Finish() MarketDepthUpdateLevelPointer {
	return MarketDepthUpdateLevelPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevelPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevelPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Side Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevelPointer) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// Side Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevelPointerBuilder) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// Price The price level to insert, update or delete.
func (m MarketDepthUpdateLevelPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// Price The price level to insert, update or delete.
func (m MarketDepthUpdateLevelPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// Quantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevelPointer) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
}

// Quantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevelPointerBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
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
func (m MarketDepthUpdateLevelPointer) UpdateType() MarketDepthUpdateTypeEnum {
	return MarketDepthUpdateTypeEnum(message.Uint8Fixed(m.Ptr, 33, 32))
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
func (m MarketDepthUpdateLevelPointerBuilder) UpdateType() MarketDepthUpdateTypeEnum {
	return MarketDepthUpdateTypeEnum(message.Uint8Fixed(m.Ptr, 33, 32))
}

// DateTime The Date-Time of the market depth update.
func (m MarketDepthUpdateLevelPointer) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 48, 40))
}

// DateTime The Date-Time of the market depth update.
func (m MarketDepthUpdateLevelPointerBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 48, 40))
}

// NumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevelPointer) NumOrders() uint32 {
	return message.Uint32Fixed(m.Ptr, 52, 48)
}

// NumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevelPointerBuilder) NumOrders() uint32 {
	return message.Uint32Fixed(m.Ptr, 52, 48)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevelPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetSide Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevelPointerBuilder) SetSide(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Ptr, 10, 8, uint16(value))
}

// SetPrice The price level to insert, update or delete.
func (m MarketDepthUpdateLevelPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetQuantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevelPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, value)
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
func (m MarketDepthUpdateLevelPointerBuilder) SetUpdateType(value MarketDepthUpdateTypeEnum) {
	message.SetUint8Fixed(m.Ptr, 33, 32, uint8(value))
}

// SetDateTime The Date-Time of the market depth update.
func (m MarketDepthUpdateLevelPointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 48, 40, float64(value))
}

// SetNumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevelPointerBuilder) SetNumOrders(value uint32) {
	message.SetUint32Fixed(m.Ptr, 52, 48, value)
}

// Copy
func (m MarketDepthUpdateLevelPointer) Copy(to MarketDepthUpdateLevelBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

// Copy
func (m MarketDepthUpdateLevelPointerBuilder) Copy(to MarketDepthUpdateLevelBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

// CopyPointer
func (m MarketDepthUpdateLevelPointer) CopyPointer(to MarketDepthUpdateLevelPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

// CopyPointer
func (m MarketDepthUpdateLevelPointerBuilder) CopyPointer(to MarketDepthUpdateLevelPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}
