package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateTrade_IntSize = 32

// {
//     Size       = MarketDataUpdateTrade_IntSize  (32)
//     Type       = MARKET_DATA_UPDATE_TRADE_INT  (126)
//     SymbolID   = 0
//     AtBidOrAsk = BID_ASK_UNSET  (0)
//     Price      = 0
//     Volume     = 0
//     DateTime   = 0
// }
var _MarketDataUpdateTrade_IntDefault = []byte{32, 0, 126, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketDataUpdateTrade_Int struct {
	message.Fixed
}

type MarketDataUpdateTrade_IntBuilder struct {
	message.Fixed
}

type MarketDataUpdateTrade_IntPointer struct {
	message.FixedPointer
}

type MarketDataUpdateTrade_IntPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataUpdateTrade_IntFrom(b []byte) MarketDataUpdateTrade_Int {
	return MarketDataUpdateTrade_Int{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateTrade_Int(b []byte) MarketDataUpdateTrade_Int {
	return MarketDataUpdateTrade_Int{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateTrade_Int() MarketDataUpdateTrade_IntBuilder {
	a := MarketDataUpdateTrade_IntBuilder{message.NewFixed(32)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateTrade_IntDefault)
	return a
}

func AllocMarketDataUpdateTrade_Int() MarketDataUpdateTrade_IntPointerBuilder {
	a := MarketDataUpdateTrade_IntPointerBuilder{message.AllocFixed(32)}
	a.Ptr.SetBytes(0, _MarketDataUpdateTrade_IntDefault)
	return a
}

func AllocMarketDataUpdateTrade_IntFrom(b []byte) MarketDataUpdateTrade_IntPointer {
	return MarketDataUpdateTrade_IntPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = MarketDataUpdateTrade_IntSize  (32)
//     Type       = MARKET_DATA_UPDATE_TRADE_INT  (126)
//     SymbolID   = 0
//     AtBidOrAsk = BID_ASK_UNSET  (0)
//     Price      = 0
//     Volume     = 0
//     DateTime   = 0
// }
func (m MarketDataUpdateTrade_IntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateTrade_IntDefault)
}

// Clear
// {
//     Size       = MarketDataUpdateTrade_IntSize  (32)
//     Type       = MARKET_DATA_UPDATE_TRADE_INT  (126)
//     SymbolID   = 0
//     AtBidOrAsk = BID_ASK_UNSET  (0)
//     Price      = 0
//     Volume     = 0
//     DateTime   = 0
// }
func (m MarketDataUpdateTrade_IntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateTrade_IntDefault)
}

// ToBuilder
func (m MarketDataUpdateTrade_Int) ToBuilder() MarketDataUpdateTrade_IntBuilder {
	return MarketDataUpdateTrade_IntBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateTrade_IntPointer) ToBuilder() MarketDataUpdateTrade_IntPointerBuilder {
	return MarketDataUpdateTrade_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateTrade_IntBuilder) Finish() MarketDataUpdateTrade_Int {
	return MarketDataUpdateTrade_Int{m.Fixed}
}

// Finish
func (m *MarketDataUpdateTrade_IntPointerBuilder) Finish() MarketDataUpdateTrade_IntPointer {
	return MarketDataUpdateTrade_IntPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDataUpdateTrade_Int) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDataUpdateTrade_IntBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDataUpdateTrade_IntPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketDataUpdateTrade_IntPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// AtBidOrAsk
func (m MarketDataUpdateTrade_Int) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 10, 8))
}

// AtBidOrAsk
func (m MarketDataUpdateTrade_IntBuilder) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 10, 8))
}

// AtBidOrAsk
func (m MarketDataUpdateTrade_IntPointer) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// AtBidOrAsk
func (m MarketDataUpdateTrade_IntPointerBuilder) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 10, 8))
}

// Price
func (m MarketDataUpdateTrade_Int) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// Price
func (m MarketDataUpdateTrade_IntBuilder) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// Price
func (m MarketDataUpdateTrade_IntPointer) Price() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// Price
func (m MarketDataUpdateTrade_IntPointerBuilder) Price() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// Volume
func (m MarketDataUpdateTrade_Int) Volume() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// Volume
func (m MarketDataUpdateTrade_IntBuilder) Volume() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// Volume
func (m MarketDataUpdateTrade_IntPointer) Volume() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// Volume
func (m MarketDataUpdateTrade_IntPointerBuilder) Volume() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// DateTime
func (m MarketDataUpdateTrade_Int) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 32, 24))
}

// DateTime
func (m MarketDataUpdateTrade_IntBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 32, 24))
}

// DateTime
func (m MarketDataUpdateTrade_IntPointer) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 32, 24))
}

// DateTime
func (m MarketDataUpdateTrade_IntPointerBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 32, 24))
}

// SetSymbolID
func (m MarketDataUpdateTrade_IntBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID
func (m MarketDataUpdateTrade_IntPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetAtBidOrAsk
func (m MarketDataUpdateTrade_IntBuilder) SetAtBidOrAsk(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Unsafe(), 10, 8, uint16(value))
}

// SetAtBidOrAsk
func (m MarketDataUpdateTrade_IntPointerBuilder) SetAtBidOrAsk(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Ptr, 10, 8, uint16(value))
}

// SetPrice
func (m MarketDataUpdateTrade_IntBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 16, 12, value)
}

// SetPrice
func (m MarketDataUpdateTrade_IntPointerBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 16, 12, value)
}

// SetVolume
func (m MarketDataUpdateTrade_IntBuilder) SetVolume(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 20, 16, value)
}

// SetVolume
func (m MarketDataUpdateTrade_IntPointerBuilder) SetVolume(value int32) {
	message.SetInt32Fixed(m.Ptr, 20, 16, value)
}

// SetDateTime
func (m MarketDataUpdateTrade_IntBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 32, 24, float64(value))
}

// SetDateTime
func (m MarketDataUpdateTrade_IntPointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, float64(value))
}

// Copy
func (m MarketDataUpdateTrade_Int) Copy(to MarketDataUpdateTrade_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateTrade_IntBuilder) Copy(to MarketDataUpdateTrade_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateTrade_Int) CopyPointer(to MarketDataUpdateTrade_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateTrade_IntBuilder) CopyPointer(to MarketDataUpdateTrade_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateTrade_IntPointer) Copy(to MarketDataUpdateTrade_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateTrade_IntPointerBuilder) Copy(to MarketDataUpdateTrade_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateTrade_IntPointer) CopyPointer(to MarketDataUpdateTrade_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateTrade_IntPointerBuilder) CopyPointer(to MarketDataUpdateTrade_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}