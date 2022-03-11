package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateBidAskSize = 40

// {
//     Size        = MarketDataUpdateBidAskSize  (40)
//     Type        = MARKET_DATA_UPDATE_BID_ASK  (108)
//     SymbolID    = 0
//     BidPrice    = math.MaxFloat64
//     BidQuantity = 0
//     AskPrice    = math.MaxFloat64
//     AskQuantity = 0
//     DateTime    = 0
// }
var _MarketDataUpdateBidAskDefault = []byte{40, 0, 108, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateBidAsk The Server sends this market data feed message to the Client when the
// best bid or ask price or size changes.
type MarketDataUpdateBidAsk struct {
	message.Fixed
}

// MarketDataUpdateBidAskBuilder The Server sends this market data feed message to the Client when the
// best bid or ask price or size changes.
type MarketDataUpdateBidAskBuilder struct {
	message.Fixed
}

// MarketDataUpdateBidAskPointer The Server sends this market data feed message to the Client when the
// best bid or ask price or size changes.
type MarketDataUpdateBidAskPointer struct {
	message.FixedPointer
}

// MarketDataUpdateBidAskPointerBuilder The Server sends this market data feed message to the Client when the
// best bid or ask price or size changes.
type MarketDataUpdateBidAskPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataUpdateBidAskFrom(b []byte) MarketDataUpdateBidAsk {
	return MarketDataUpdateBidAsk{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateBidAsk(b []byte) MarketDataUpdateBidAsk {
	return MarketDataUpdateBidAsk{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateBidAsk() MarketDataUpdateBidAskBuilder {
	a := MarketDataUpdateBidAskBuilder{message.NewFixed(40)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateBidAskDefault)
	return a
}

func AllocMarketDataUpdateBidAsk() MarketDataUpdateBidAskPointerBuilder {
	a := MarketDataUpdateBidAskPointerBuilder{message.AllocFixed(40)}
	a.Ptr.SetBytes(0, _MarketDataUpdateBidAskDefault)
	return a
}

func AllocMarketDataUpdateBidAskFrom(b []byte) MarketDataUpdateBidAskPointer {
	return MarketDataUpdateBidAskPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size        = MarketDataUpdateBidAskSize  (40)
//     Type        = MARKET_DATA_UPDATE_BID_ASK  (108)
//     SymbolID    = 0
//     BidPrice    = math.MaxFloat64
//     BidQuantity = 0
//     AskPrice    = math.MaxFloat64
//     AskQuantity = 0
//     DateTime    = 0
// }
func (m MarketDataUpdateBidAskBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateBidAskDefault)
}

// Clear
// {
//     Size        = MarketDataUpdateBidAskSize  (40)
//     Type        = MARKET_DATA_UPDATE_BID_ASK  (108)
//     SymbolID    = 0
//     BidPrice    = math.MaxFloat64
//     BidQuantity = 0
//     AskPrice    = math.MaxFloat64
//     AskQuantity = 0
//     DateTime    = 0
// }
func (m MarketDataUpdateBidAskPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateBidAskDefault)
}

// ToBuilder
func (m MarketDataUpdateBidAsk) ToBuilder() MarketDataUpdateBidAskBuilder {
	return MarketDataUpdateBidAskBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateBidAskPointer) ToBuilder() MarketDataUpdateBidAskPointerBuilder {
	return MarketDataUpdateBidAskPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateBidAskBuilder) Finish() MarketDataUpdateBidAsk {
	return MarketDataUpdateBidAsk{m.Fixed}
}

// Finish
func (m *MarketDataUpdateBidAskPointerBuilder) Finish() MarketDataUpdateBidAskPointer {
	return MarketDataUpdateBidAskPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAsk) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAsk) BidPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskBuilder) BidPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskPointer) BidPrice() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskPointerBuilder) BidPrice() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// BidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAsk) BidQuantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 20, 16)
}

// BidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskBuilder) BidQuantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 20, 16)
}

// BidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskPointer) BidQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 20, 16)
}

// BidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskPointerBuilder) BidQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 20, 16)
}

// AskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAsk) AskPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// AskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskBuilder) AskPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// AskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskPointer) AskPrice() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
}

// AskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskPointerBuilder) AskPrice() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
}

// AskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAsk) AskQuantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 36, 32)
}

// AskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskBuilder) AskQuantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 36, 32)
}

// AskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskPointer) AskQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 36, 32)
}

// AskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskPointerBuilder) AskQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 36, 32)
}

// DateTime The Date-Time of the Bid and Ask update.
func (m MarketDataUpdateBidAsk) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 40, 36))
}

// DateTime The Date-Time of the Bid and Ask update.
func (m MarketDataUpdateBidAskBuilder) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 40, 36))
}

// DateTime The Date-Time of the Bid and Ask update.
func (m MarketDataUpdateBidAskPointer) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 40, 36))
}

// DateTime The Date-Time of the Bid and Ask update.
func (m MarketDataUpdateBidAskPointerBuilder) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 40, 36))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetBidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskBuilder) SetBidPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, value)
}

// SetBidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskPointerBuilder) SetBidPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, value)
}

// SetBidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskBuilder) SetBidQuantity(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 20, 16, value)
}

// SetBidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskPointerBuilder) SetBidQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 20, 16, value)
}

// SetAskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskBuilder) SetAskPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 32, 24, value)
}

// SetAskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskPointerBuilder) SetAskPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, value)
}

// SetAskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskBuilder) SetAskQuantity(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 36, 32, value)
}

// SetAskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskPointerBuilder) SetAskQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 36, 32, value)
}

// SetDateTime The Date-Time of the Bid and Ask update.
func (m MarketDataUpdateBidAskBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 40, 36, uint32(value))
}

// SetDateTime The Date-Time of the Bid and Ask update.
func (m MarketDataUpdateBidAskPointerBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 40, 36, uint32(value))
}

// Copy
func (m MarketDataUpdateBidAsk) Copy(to MarketDataUpdateBidAskBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateBidAskBuilder) Copy(to MarketDataUpdateBidAskBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateBidAsk) CopyPointer(to MarketDataUpdateBidAskPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateBidAskBuilder) CopyPointer(to MarketDataUpdateBidAskPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateBidAskPointer) Copy(to MarketDataUpdateBidAskBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateBidAskPointerBuilder) Copy(to MarketDataUpdateBidAskBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateBidAskPointer) CopyPointer(to MarketDataUpdateBidAskPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateBidAskPointerBuilder) CopyPointer(to MarketDataUpdateBidAskPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}
