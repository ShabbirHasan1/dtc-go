package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDataUpdateTradeNoTimestampPointer This message is optional.
//
// Sent by the Server to the Client when a trade occurs. This message is
// identical to the MarketDataUpdateTrade_WITH_UNBUNDLED_INDICATOR_2 message
// except it does not have a timestamp. It needs to be sent when there is
// no change with the timestamp for the trade as compared to the prior trade.
// no change with the timestamp for the trade as compared to the prior trade.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received trade timestamp to know what the timestamp is for
// this message.
type MarketDataUpdateTradeNoTimestampPointer struct {
	message.FixedPointer
}

// MarketDataUpdateTradeNoTimestampPointerBuilder This message is optional.
//
// Sent by the Server to the Client when a trade occurs. This message is
// identical to the MarketDataUpdateTrade_WITH_UNBUNDLED_INDICATOR_2 message
// except it does not have a timestamp. It needs to be sent when there is
// no change with the timestamp for the trade as compared to the prior trade.
// no change with the timestamp for the trade as compared to the prior trade.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received trade timestamp to know what the timestamp is for
// this message.
type MarketDataUpdateTradeNoTimestampPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataUpdateTradeNoTimestamp() MarketDataUpdateTradeNoTimestampPointerBuilder {
	a := MarketDataUpdateTradeNoTimestampPointerBuilder{message.AllocFixed(19)}
	a.Ptr.SetBytes(0, _MarketDataUpdateTradeNoTimestampDefault)
	return a
}

func AllocMarketDataUpdateTradeNoTimestampFrom(b []byte) MarketDataUpdateTradeNoTimestampPointer {
	return MarketDataUpdateTradeNoTimestampPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                    = MarketDataUpdateTradeNoTimestampSize  (19)
//     Type                    = MARKET_DATA_UPDATE_TRADE_NO_TIMESTAMP  (142)
//     SymbolID                = 0
//     Price                   = 0
//     Volume                  = 0
//     AtBidOrAsk              = 0
//     UnbundledTradeIndicator = UNBUNDLED_TRADE_NONE  (0)
//     TradeCondition          = 0
// }
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateTradeNoTimestampDefault)
}

// ToBuilder
func (m MarketDataUpdateTradeNoTimestampPointer) ToBuilder() MarketDataUpdateTradeNoTimestampPointerBuilder {
	return MarketDataUpdateTradeNoTimestampPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateTradeNoTimestampPointerBuilder) Finish() MarketDataUpdateTradeNoTimestampPointer {
	return MarketDataUpdateTradeNoTimestampPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeNoTimestampPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Price The price of the trade.
func (m MarketDataUpdateTradeNoTimestampPointer) Price() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// Price The price of the trade.
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) Price() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeNoTimestampPointer) Volume() uint32 {
	return message.Uint32Fixed(m.Ptr, 16, 12)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) Volume() uint32 {
	return message.Uint32Fixed(m.Ptr, 16, 12)
}

// AtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeNoTimestampPointer) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Ptr, 17, 16))
}

// AtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Ptr, 17, 16))
}

// UnbundledTradeIndicator This is an optional field.
//
// For symbols on exchanges which support reporting individual trades which
// are part of a larger summary trade, this message field is used. It indicates
// whether the trade is part of a larger summary trrade and reported as an
// unbundled individual trade.
//
// The possible values are listed below.
//
// UNBUNDLED_TRADE_NONE = 0
// FIRST_SUB_TRADE_OF_UNBUNDLED_TRADE = 1
// LAST_SUB_TRADE_OF_UNBUNDLED_TRADE = 2
func (m MarketDataUpdateTradeNoTimestampPointer) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(message.Int8Fixed(m.Ptr, 18, 17))
}

// UnbundledTradeIndicator This is an optional field.
//
// For symbols on exchanges which support reporting individual trades which
// are part of a larger summary trade, this message field is used. It indicates
// whether the trade is part of a larger summary trrade and reported as an
// unbundled individual trade.
//
// The possible values are listed below.
//
// UNBUNDLED_TRADE_NONE = 0
// FIRST_SUB_TRADE_OF_UNBUNDLED_TRADE = 1
// LAST_SUB_TRADE_OF_UNBUNDLED_TRADE = 2
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(message.Int8Fixed(m.Ptr, 18, 17))
}

// TradeCondition This is an optional field. It usually applies to stock symbols.
//
// This field indicates a special condition which applies to the trade. The
// possible values are listed below.
//
// TRADE_CONDITION_NONE = 0
// TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
// TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
func (m MarketDataUpdateTradeNoTimestampPointer) TradeCondition() TradeConditionEnum {
	return TradeConditionEnum(message.Int8Fixed(m.Ptr, 19, 18))
}

// TradeCondition This is an optional field. It usually applies to stock symbols.
//
// This field indicates a special condition which applies to the trade. The
// possible values are listed below.
//
// TRADE_CONDITION_NONE = 0
// TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
// TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) TradeCondition() TradeConditionEnum {
	return TradeConditionEnum(message.Int8Fixed(m.Ptr, 19, 18))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetPrice The price of the trade.
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 12, 8, value)
}

// SetVolume The volume of the trade.
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) SetVolume(value uint32) {
	message.SetUint32Fixed(m.Ptr, 16, 12, value)
}

// SetAtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) SetAtBidOrAsk(value AtBidOrAskEnum8) {
	message.SetUint8Fixed(m.Ptr, 17, 16, uint8(value))
}

// SetUnbundledTradeIndicator This is an optional field.
//
// For symbols on exchanges which support reporting individual trades which
// are part of a larger summary trade, this message field is used. It indicates
// whether the trade is part of a larger summary trrade and reported as an
// unbundled individual trade.
//
// The possible values are listed below.
//
// UNBUNDLED_TRADE_NONE = 0
// FIRST_SUB_TRADE_OF_UNBUNDLED_TRADE = 1
// LAST_SUB_TRADE_OF_UNBUNDLED_TRADE = 2
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) SetUnbundledTradeIndicator(value UnbundledTradeIndicatorEnum) {
	message.SetInt8Fixed(m.Ptr, 18, 17, int8(value))
}

// SetTradeCondition This is an optional field. It usually applies to stock symbols.
//
// This field indicates a special condition which applies to the trade. The
// possible values are listed below.
//
// TRADE_CONDITION_NONE = 0
// TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
// TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) SetTradeCondition(value TradeConditionEnum) {
	message.SetInt8Fixed(m.Ptr, 19, 18, int8(value))
}

// Copy
func (m MarketDataUpdateTradeNoTimestampPointer) Copy(to MarketDataUpdateTradeNoTimestampBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// Copy
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) Copy(to MarketDataUpdateTradeNoTimestampBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// CopyPointer
func (m MarketDataUpdateTradeNoTimestampPointer) CopyPointer(to MarketDataUpdateTradeNoTimestampPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// CopyPointer
func (m MarketDataUpdateTradeNoTimestampPointerBuilder) CopyPointer(to MarketDataUpdateTradeNoTimestampPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}
