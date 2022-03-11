package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
func (m MarketDataUpdateBidAskPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateBidAskDefault)
}

// ToBuilder
func (m MarketDataUpdateBidAskPointer) ToBuilder() MarketDataUpdateBidAskPointerBuilder {
	return MarketDataUpdateBidAskPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateBidAskPointerBuilder) Finish() MarketDataUpdateBidAskPointer {
	return MarketDataUpdateBidAskPointer{m.FixedPointer}
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
func (m MarketDataUpdateBidAskPointer) BidPrice() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskPointerBuilder) BidPrice() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
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
func (m MarketDataUpdateBidAskPointer) AskPrice() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
}

// AskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskPointerBuilder) AskPrice() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
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
func (m MarketDataUpdateBidAskPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetBidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskPointerBuilder) SetBidPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, value)
}

// SetBidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskPointerBuilder) SetBidQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 20, 16, value)
}

// SetAskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskPointerBuilder) SetAskPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, value)
}

// SetAskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskPointerBuilder) SetAskQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 36, 32, value)
}

// SetDateTime The Date-Time of the Bid and Ask update.
func (m MarketDataUpdateBidAskPointerBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 40, 36, uint32(value))
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
