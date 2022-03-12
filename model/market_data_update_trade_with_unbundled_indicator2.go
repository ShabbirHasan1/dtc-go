package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateTradeWithUnbundledIndicator2Size = 27

// {
//     Size                    = MarketDataUpdateTradeWithUnbundledIndicator2Size  (27)
//     Type                    = MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_2  (146)
//     SymbolID                = 0
//     Price                   = 0
//     Volume                  = 0
//     DateTime                = 0
//     AtBidOrAsk              = BID_ASK_UNSET_8  (0)
//     UnbundledTradeIndicator = UNBUNDLED_TRADE_NONE  (0)
//     TradeCondition          = TRADE_CONDITION_NONE  (0)
// }
var _MarketDataUpdateTradeWithUnbundledIndicator2Default = []byte{27, 0, 146, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateTradeWithUnbundledIndicator2 Sent by the Server to the Client when a trade occurs. This message has
// additional fields as compared to the MarketDataUpdateTrade message and
// also supports microsecond time stamping.
type MarketDataUpdateTradeWithUnbundledIndicator2 struct {
	message.Fixed
}

// MarketDataUpdateTradeWithUnbundledIndicator2Builder Sent by the Server to the Client when a trade occurs. This message has
// additional fields as compared to the MarketDataUpdateTrade message and
// also supports microsecond time stamping.
type MarketDataUpdateTradeWithUnbundledIndicator2Builder struct {
	message.Fixed
}

// MarketDataUpdateTradeWithUnbundledIndicator2Pointer Sent by the Server to the Client when a trade occurs. This message has
// additional fields as compared to the MarketDataUpdateTrade message and
// also supports microsecond time stamping.
type MarketDataUpdateTradeWithUnbundledIndicator2Pointer struct {
	message.FixedPointer
}

// MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder Sent by the Server to the Client when a trade occurs. This message has
// additional fields as compared to the MarketDataUpdateTrade message and
// also supports microsecond time stamping.
type MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataUpdateTradeWithUnbundledIndicator2From(b []byte) MarketDataUpdateTradeWithUnbundledIndicator2 {
	return MarketDataUpdateTradeWithUnbundledIndicator2{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateTradeWithUnbundledIndicator2(b []byte) MarketDataUpdateTradeWithUnbundledIndicator2 {
	return MarketDataUpdateTradeWithUnbundledIndicator2{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateTradeWithUnbundledIndicator2() MarketDataUpdateTradeWithUnbundledIndicator2Builder {
	a := MarketDataUpdateTradeWithUnbundledIndicator2Builder{message.NewFixed(27)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateTradeWithUnbundledIndicator2Default)
	return a
}

func AllocMarketDataUpdateTradeWithUnbundledIndicator2() MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder {
	a := MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder{message.AllocFixed(27)}
	a.Ptr.SetBytes(0, _MarketDataUpdateTradeWithUnbundledIndicator2Default)
	return a
}

func AllocMarketDataUpdateTradeWithUnbundledIndicator2From(b []byte) MarketDataUpdateTradeWithUnbundledIndicator2Pointer {
	return MarketDataUpdateTradeWithUnbundledIndicator2Pointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                    = MarketDataUpdateTradeWithUnbundledIndicator2Size  (27)
//     Type                    = MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_2  (146)
//     SymbolID                = 0
//     Price                   = 0
//     Volume                  = 0
//     DateTime                = 0
//     AtBidOrAsk              = BID_ASK_UNSET_8  (0)
//     UnbundledTradeIndicator = UNBUNDLED_TRADE_NONE  (0)
//     TradeCondition          = TRADE_CONDITION_NONE  (0)
// }
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateTradeWithUnbundledIndicator2Default)
}

// Clear
// {
//     Size                    = MarketDataUpdateTradeWithUnbundledIndicator2Size  (27)
//     Type                    = MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_2  (146)
//     SymbolID                = 0
//     Price                   = 0
//     Volume                  = 0
//     DateTime                = 0
//     AtBidOrAsk              = BID_ASK_UNSET_8  (0)
//     UnbundledTradeIndicator = UNBUNDLED_TRADE_NONE  (0)
//     TradeCondition          = TRADE_CONDITION_NONE  (0)
// }
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateTradeWithUnbundledIndicator2Default)
}

// ToBuilder
func (m MarketDataUpdateTradeWithUnbundledIndicator2) ToBuilder() MarketDataUpdateTradeWithUnbundledIndicator2Builder {
	return MarketDataUpdateTradeWithUnbundledIndicator2Builder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) ToBuilder() MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder {
	return MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) Finish() MarketDataUpdateTradeWithUnbundledIndicator2 {
	return MarketDataUpdateTradeWithUnbundledIndicator2{m.Fixed}
}

// Finish
func (m *MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) Finish() MarketDataUpdateTradeWithUnbundledIndicator2Pointer {
	return MarketDataUpdateTradeWithUnbundledIndicator2Pointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeWithUnbundledIndicator2) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Price The price of the trade.
func (m MarketDataUpdateTradeWithUnbundledIndicator2) Price() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// Price The price of the trade.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) Price() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// Price The price of the trade.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) Price() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// Price The price of the trade.
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) Price() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeWithUnbundledIndicator2) Volume() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 16, 12)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) Volume() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 16, 12)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) Volume() uint32 {
	return message.Uint32Fixed(m.Ptr, 16, 12)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) Volume() uint32 {
	return message.Uint32Fixed(m.Ptr, 16, 12)
}

// DateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateTradeWithUnbundledIndicator2) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 24, 16))
}

// DateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 24, 16))
}

// DateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 24, 16))
}

// DateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 24, 16))
}

// AtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeWithUnbundledIndicator2) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Unsafe(), 25, 24))
}

// AtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Unsafe(), 25, 24))
}

// AtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Ptr, 25, 24))
}

// AtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(message.Uint8Fixed(m.Ptr, 25, 24))
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
func (m MarketDataUpdateTradeWithUnbundledIndicator2) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(message.Int8Fixed(m.Unsafe(), 26, 25))
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
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(message.Int8Fixed(m.Unsafe(), 26, 25))
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
func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(message.Int8Fixed(m.Ptr, 26, 25))
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
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(message.Int8Fixed(m.Ptr, 26, 25))
}

// TradeCondition This is an optional field. It usually applies to stock symbols.
//
// This field indicates a special condition which applies to the trade. The
// possible values are listed below.
//
// TRADE_CONDITION_NONE = 0
// TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
// TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
func (m MarketDataUpdateTradeWithUnbundledIndicator2) TradeCondition() TradeConditionEnum {
	return TradeConditionEnum(message.Int8Fixed(m.Unsafe(), 27, 26))
}

// TradeCondition This is an optional field. It usually applies to stock symbols.
//
// This field indicates a special condition which applies to the trade. The
// possible values are listed below.
//
// TRADE_CONDITION_NONE = 0
// TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
// TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) TradeCondition() TradeConditionEnum {
	return TradeConditionEnum(message.Int8Fixed(m.Unsafe(), 27, 26))
}

// TradeCondition This is an optional field. It usually applies to stock symbols.
//
// This field indicates a special condition which applies to the trade. The
// possible values are listed below.
//
// TRADE_CONDITION_NONE = 0
// TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
// TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) TradeCondition() TradeConditionEnum {
	return TradeConditionEnum(message.Int8Fixed(m.Ptr, 27, 26))
}

// TradeCondition This is an optional field. It usually applies to stock symbols.
//
// This field indicates a special condition which applies to the trade. The
// possible values are listed below.
//
// TRADE_CONDITION_NONE = 0
// TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
// TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) TradeCondition() TradeConditionEnum {
	return TradeConditionEnum(message.Int8Fixed(m.Ptr, 27, 26))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetPrice The price of the trade.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 12, 8, value)
}

// SetPrice The price of the trade.
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 12, 8, value)
}

// SetVolume The volume of the trade.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) SetVolume(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 16, 12, value)
}

// SetVolume The volume of the trade.
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) SetVolume(value uint32) {
	message.SetUint32Fixed(m.Ptr, 16, 12, value)
}

// SetDateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 24, 16, int64(value))
}

// SetDateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 24, 16, int64(value))
}

// SetAtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) SetAtBidOrAsk(value AtBidOrAskEnum8) {
	message.SetUint8Fixed(m.Unsafe(), 25, 24, uint8(value))
}

// SetAtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) SetAtBidOrAsk(value AtBidOrAskEnum8) {
	message.SetUint8Fixed(m.Ptr, 25, 24, uint8(value))
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
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) SetUnbundledTradeIndicator(value UnbundledTradeIndicatorEnum) {
	message.SetInt8Fixed(m.Unsafe(), 26, 25, int8(value))
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
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) SetUnbundledTradeIndicator(value UnbundledTradeIndicatorEnum) {
	message.SetInt8Fixed(m.Ptr, 26, 25, int8(value))
}

// SetTradeCondition This is an optional field. It usually applies to stock symbols.
//
// This field indicates a special condition which applies to the trade. The
// possible values are listed below.
//
// TRADE_CONDITION_NONE = 0
// TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
// TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) SetTradeCondition(value TradeConditionEnum) {
	message.SetInt8Fixed(m.Unsafe(), 27, 26, int8(value))
}

// SetTradeCondition This is an optional field. It usually applies to stock symbols.
//
// This field indicates a special condition which applies to the trade. The
// possible values are listed below.
//
// TRADE_CONDITION_NONE = 0
// TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
// TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) SetTradeCondition(value TradeConditionEnum) {
	message.SetInt8Fixed(m.Ptr, 27, 26, int8(value))
}

// Copy
func (m MarketDataUpdateTradeWithUnbundledIndicator2) Copy(to MarketDataUpdateTradeWithUnbundledIndicator2Builder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// Copy
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) Copy(to MarketDataUpdateTradeWithUnbundledIndicator2Builder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// CopyPointer
func (m MarketDataUpdateTradeWithUnbundledIndicator2) CopyPointer(to MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// CopyPointer
func (m MarketDataUpdateTradeWithUnbundledIndicator2Builder) CopyPointer(to MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// Copy
func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) Copy(to MarketDataUpdateTradeWithUnbundledIndicator2Builder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// Copy
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) Copy(to MarketDataUpdateTradeWithUnbundledIndicator2Builder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// CopyPointer
func (m MarketDataUpdateTradeWithUnbundledIndicator2Pointer) CopyPointer(to MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}

// CopyPointer
func (m MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) CopyPointer(to MarketDataUpdateTradeWithUnbundledIndicator2PointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
}
