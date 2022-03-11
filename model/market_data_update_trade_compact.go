package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateTradeCompactSize = 24

// {
//     Size       = MarketDataUpdateTradeCompactSize  (24)
//     Type       = MARKET_DATA_UPDATE_TRADE_COMPACT  (112)
//     Price      = 0
//     Volume     = 0
//     DateTime   = 0
//     SymbolID   = 0
//     AtBidOrAsk = BID_ASK_UNSET  (0)
// }
var _MarketDataUpdateTradeCompactDefault = []byte{24, 0, 112, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateTradeCompact Sent by the Server to the Client when a trade occurs. This message is
// a more compact MarketDataUpdateTrade. For the price it uses a 4 byte float.
// a more compact MarketDataUpdateTrade. For the price it uses a 4 byte float.
type MarketDataUpdateTradeCompact struct {
	message.Fixed
}

// MarketDataUpdateTradeCompactBuilder Sent by the Server to the Client when a trade occurs. This message is
// a more compact MarketDataUpdateTrade. For the price it uses a 4 byte float.
// a more compact MarketDataUpdateTrade. For the price it uses a 4 byte float.
type MarketDataUpdateTradeCompactBuilder struct {
	message.Fixed
}

// MarketDataUpdateTradeCompactPointer Sent by the Server to the Client when a trade occurs. This message is
// a more compact MarketDataUpdateTrade. For the price it uses a 4 byte float.
// a more compact MarketDataUpdateTrade. For the price it uses a 4 byte float.
type MarketDataUpdateTradeCompactPointer struct {
	message.FixedPointer
}

// MarketDataUpdateTradeCompactPointerBuilder Sent by the Server to the Client when a trade occurs. This message is
// a more compact MarketDataUpdateTrade. For the price it uses a 4 byte float.
// a more compact MarketDataUpdateTrade. For the price it uses a 4 byte float.
type MarketDataUpdateTradeCompactPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataUpdateTradeCompactFrom(b []byte) MarketDataUpdateTradeCompact {
	return MarketDataUpdateTradeCompact{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateTradeCompact(b []byte) MarketDataUpdateTradeCompact {
	return MarketDataUpdateTradeCompact{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateTradeCompact() MarketDataUpdateTradeCompactBuilder {
	a := MarketDataUpdateTradeCompactBuilder{message.NewFixed(24)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateTradeCompactDefault)
	return a
}

func AllocMarketDataUpdateTradeCompact() MarketDataUpdateTradeCompactPointerBuilder {
	a := MarketDataUpdateTradeCompactPointerBuilder{message.AllocFixed(24)}
	a.Ptr.SetBytes(0, _MarketDataUpdateTradeCompactDefault)
	return a
}

func AllocMarketDataUpdateTradeCompactFrom(b []byte) MarketDataUpdateTradeCompactPointer {
	return MarketDataUpdateTradeCompactPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = MarketDataUpdateTradeCompactSize  (24)
//     Type       = MARKET_DATA_UPDATE_TRADE_COMPACT  (112)
//     Price      = 0
//     Volume     = 0
//     DateTime   = 0
//     SymbolID   = 0
//     AtBidOrAsk = BID_ASK_UNSET  (0)
// }
func (m MarketDataUpdateTradeCompactBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateTradeCompactDefault)
}

// Clear
// {
//     Size       = MarketDataUpdateTradeCompactSize  (24)
//     Type       = MARKET_DATA_UPDATE_TRADE_COMPACT  (112)
//     Price      = 0
//     Volume     = 0
//     DateTime   = 0
//     SymbolID   = 0
//     AtBidOrAsk = BID_ASK_UNSET  (0)
// }
func (m MarketDataUpdateTradeCompactPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateTradeCompactDefault)
}

// ToBuilder
func (m MarketDataUpdateTradeCompact) ToBuilder() MarketDataUpdateTradeCompactBuilder {
	return MarketDataUpdateTradeCompactBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateTradeCompactPointer) ToBuilder() MarketDataUpdateTradeCompactPointerBuilder {
	return MarketDataUpdateTradeCompactPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateTradeCompactBuilder) Finish() MarketDataUpdateTradeCompact {
	return MarketDataUpdateTradeCompact{m.Fixed}
}

// Finish
func (m *MarketDataUpdateTradeCompactPointerBuilder) Finish() MarketDataUpdateTradeCompactPointer {
	return MarketDataUpdateTradeCompactPointer{m.FixedPointer}
}

// Price The price of the trade.
func (m MarketDataUpdateTradeCompact) Price() float32 {
	return message.Float32Fixed(m.Unsafe(), 8, 4)
}

// Price The price of the trade.
func (m MarketDataUpdateTradeCompactBuilder) Price() float32 {
	return message.Float32Fixed(m.Unsafe(), 8, 4)
}

// Price The price of the trade.
func (m MarketDataUpdateTradeCompactPointer) Price() float32 {
	return message.Float32Fixed(m.Ptr, 8, 4)
}

// Price The price of the trade.
func (m MarketDataUpdateTradeCompactPointerBuilder) Price() float32 {
	return message.Float32Fixed(m.Ptr, 8, 4)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeCompact) Volume() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeCompactBuilder) Volume() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeCompactPointer) Volume() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTradeCompactPointerBuilder) Volume() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// DateTime The timestamp of the trade in UNIX time format. This does not contain
// the milliseconds for compactness.
func (m MarketDataUpdateTradeCompact) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 16, 12))
}

// DateTime The timestamp of the trade in UNIX time format. This does not contain
// the milliseconds for compactness.
func (m MarketDataUpdateTradeCompactBuilder) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 16, 12))
}

// DateTime The timestamp of the trade in UNIX time format. This does not contain
// the milliseconds for compactness.
func (m MarketDataUpdateTradeCompactPointer) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 16, 12))
}

// DateTime The timestamp of the trade in UNIX time format. This does not contain
// the milliseconds for compactness.
func (m MarketDataUpdateTradeCompactPointerBuilder) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 16, 12))
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeCompact) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 20, 16)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeCompactBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 20, 16)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeCompactPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 20, 16)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeCompactPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 20, 16)
}

// AtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeCompact) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 22, 20))
}

// AtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeCompactBuilder) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 22, 20))
}

// AtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeCompactPointer) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 22, 20))
}

// AtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeCompactPointerBuilder) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 22, 20))
}

// SetPrice The price of the trade.
func (m MarketDataUpdateTradeCompactBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 8, 4, value)
}

// SetPrice The price of the trade.
func (m MarketDataUpdateTradeCompactPointerBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 8, 4, value)
}

// SetVolume The volume of the trade.
func (m MarketDataUpdateTradeCompactBuilder) SetVolume(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 12, 8, value)
}

// SetVolume The volume of the trade.
func (m MarketDataUpdateTradeCompactPointerBuilder) SetVolume(value float32) {
	message.SetFloat32Fixed(m.Ptr, 12, 8, value)
}

// SetDateTime The timestamp of the trade in UNIX time format. This does not contain
// the milliseconds for compactness.
func (m MarketDataUpdateTradeCompactBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 16, 12, uint32(value))
}

// SetDateTime The timestamp of the trade in UNIX time format. This does not contain
// the milliseconds for compactness.
func (m MarketDataUpdateTradeCompactPointerBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 16, 12, uint32(value))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeCompactBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 20, 16, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradeCompactPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 20, 16, value)
}

// SetAtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeCompactBuilder) SetAtBidOrAsk(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Unsafe(), 22, 20, uint16(value))
}

// SetAtBidOrAsk Indicator whether the trade occurred at the Bid or Ask price.
func (m MarketDataUpdateTradeCompactPointerBuilder) SetAtBidOrAsk(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Ptr, 22, 20, uint16(value))
}

// Copy
func (m MarketDataUpdateTradeCompact) Copy(to MarketDataUpdateTradeCompactBuilder) {
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
}

// Copy
func (m MarketDataUpdateTradeCompactBuilder) Copy(to MarketDataUpdateTradeCompactBuilder) {
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
}

// CopyPointer
func (m MarketDataUpdateTradeCompact) CopyPointer(to MarketDataUpdateTradeCompactPointerBuilder) {
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
}

// CopyPointer
func (m MarketDataUpdateTradeCompactBuilder) CopyPointer(to MarketDataUpdateTradeCompactPointerBuilder) {
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
}

// Copy
func (m MarketDataUpdateTradeCompactPointer) Copy(to MarketDataUpdateTradeCompactBuilder) {
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
}

// Copy
func (m MarketDataUpdateTradeCompactPointerBuilder) Copy(to MarketDataUpdateTradeCompactBuilder) {
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
}

// CopyPointer
func (m MarketDataUpdateTradeCompactPointer) CopyPointer(to MarketDataUpdateTradeCompactPointerBuilder) {
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
}

// CopyPointer
func (m MarketDataUpdateTradeCompactPointerBuilder) CopyPointer(to MarketDataUpdateTradeCompactPointerBuilder) {
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
}
