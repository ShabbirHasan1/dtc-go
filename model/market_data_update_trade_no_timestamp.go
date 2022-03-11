package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _MarketDataUpdateTradeNoTimestampDefault = []byte{19, 0, 142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataUpdateTradeNoTimestampSize = 19

// MarketDataUpdateTradeNoTimestamp This message is optional.
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
type MarketDataUpdateTradeNoTimestamp struct {
	message.Fixed
}

// MarketDataUpdateTradeNoTimestampBuilder This message is optional.
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
type MarketDataUpdateTradeNoTimestampBuilder struct {
	message.Fixed
}

func NewMarketDataUpdateTradeNoTimestampFrom(b []byte) MarketDataUpdateTradeNoTimestamp {
	return MarketDataUpdateTradeNoTimestamp{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateTradeNoTimestamp(b []byte) MarketDataUpdateTradeNoTimestamp {
	return MarketDataUpdateTradeNoTimestamp{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateTradeNoTimestamp() MarketDataUpdateTradeNoTimestampBuilder {
	a := MarketDataUpdateTradeNoTimestampBuilder{message.NewFixed(19)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateTradeNoTimestampDefault)
	return a
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
func (m MarketDataUpdateTradeNoTimestampBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateTradeNoTimestampDefault)
}

// ToBuilder
func (m MarketDataUpdateTradeNoTimestamp) ToBuilder() MarketDataUpdateTradeNoTimestampBuilder {
	return MarketDataUpdateTradeNoTimestampBuilder{m.Fixed}
}

// Finish
func (m MarketDataUpdateTradeNoTimestampBuilder) Finish() MarketDataUpdateTradeNoTimestamp {
	return MarketDataUpdateTradeNoTimestamp{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeNoTimestamp) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeNoTimestampBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Price The price of the trade.
func (m MarketDataUpdateTradeNoTimestamp) Price() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// Price The price of the trade.
func (m MarketDataUpdateTradeNoTimestampBuilder) Price() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeNoTimestamp) Volume() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 16, 12)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeNoTimestampBuilder) Volume() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 16, 12)
}

// AtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeNoTimestamp) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Unsafe(), 17, 16))
}

// AtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeNoTimestampBuilder) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Unsafe(), 17, 16))
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
func (m MarketDataUpdateTradeNoTimestamp) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(message.Int8Fixed(m.Unsafe(), 18, 17))
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
func (m MarketDataUpdateTradeNoTimestampBuilder) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(message.Int8Fixed(m.Unsafe(), 18, 17))
}

// TradeCondition This is an optional field. It usually applies to stock symbols.
//
// This field indicates a special condition which applies to the trade. The
// possible values are listed below.
//
// TRADE_CONDITION_NONE = 0
// TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
// TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
func (m MarketDataUpdateTradeNoTimestamp) TradeCondition() TradeConditionEnum {
	return TradeConditionEnum(message.Int8Fixed(m.Unsafe(), 19, 18))
}

// TradeCondition This is an optional field. It usually applies to stock symbols.
//
// This field indicates a special condition which applies to the trade. The
// possible values are listed below.
//
// TRADE_CONDITION_NONE = 0
// TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
// TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
func (m MarketDataUpdateTradeNoTimestampBuilder) TradeCondition() TradeConditionEnum {
	return TradeConditionEnum(message.Int8Fixed(m.Unsafe(), 19, 18))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeNoTimestampBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetPrice The price of the trade.
func (m MarketDataUpdateTradeNoTimestampBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 12, 8, value)
}

// SetVolume The volume of the trade.
func (m MarketDataUpdateTradeNoTimestampBuilder) SetVolume(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 16, 12, value)
}

// SetAtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeNoTimestampBuilder) SetAtBidOrAsk(value AtBidOrAskEnum8) {
	message.SetUint8Fixed(m.Unsafe(), 17, 16, uint8(value))
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
func (m MarketDataUpdateTradeNoTimestampBuilder) SetUnbundledTradeIndicator(value UnbundledTradeIndicatorEnum) {
	message.SetInt8Fixed(m.Unsafe(), 18, 17, int8(value))
}

// SetTradeCondition This is an optional field. It usually applies to stock symbols.
//
// This field indicates a special condition which applies to the trade. The
// possible values are listed below.
//
// TRADE_CONDITION_NONE = 0
// TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
// TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
func (m MarketDataUpdateTradeNoTimestampBuilder) SetTradeCondition(value TradeConditionEnum) {
	message.SetInt8Fixed(m.Unsafe(), 19, 18, int8(value))
}

// Copy
func (m MarketDataUpdateTradeNoTimestamp) Copy(to MarketDataUpdateTradeNoTimestampBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// Copy
func (m MarketDataUpdateTradeNoTimestampBuilder) Copy(to MarketDataUpdateTradeNoTimestampBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// CopyPointer
func (m MarketDataUpdateTradeNoTimestamp) CopyPointer(to MarketDataUpdateTradeNoTimestampPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// CopyPointer
func (m MarketDataUpdateTradeNoTimestampBuilder) CopyPointer(to MarketDataUpdateTradeNoTimestampPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}
