package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketDataUpdateBidAskCompactPointer struct {
	message.FixedPointer
}

type MarketDataUpdateBidAskCompactPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataUpdateBidAskCompact() MarketDataUpdateBidAskCompactPointerBuilder {
	a := MarketDataUpdateBidAskCompactPointerBuilder{message.AllocFixed(28)}
	a.Ptr.SetBytes(0, _MarketDataUpdateBidAskCompactDefault)
	return a
}

func AllocMarketDataUpdateBidAskCompactFrom(b []byte) MarketDataUpdateBidAskCompactPointer {
	return MarketDataUpdateBidAskCompactPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size        = MarketDataUpdateBidAskCompactSize  (28)
//     Type        = MARKET_DATA_UPDATE_BID_ASK_COMPACT  (117)
//     BidPrice    = math.MaxFloat32
//     BidQuantity = 0
//     AskPrice    = math.MaxFloat32
//     AskQuantity = 0
//     DateTime    = 0
//     SymbolID    = 0
// }
func (m MarketDataUpdateBidAskCompactPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateBidAskCompactDefault)
}

// ToBuilder
func (m MarketDataUpdateBidAskCompactPointer) ToBuilder() MarketDataUpdateBidAskCompactPointerBuilder {
	return MarketDataUpdateBidAskCompactPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateBidAskCompactPointerBuilder) Finish() MarketDataUpdateBidAskCompactPointer {
	return MarketDataUpdateBidAskCompactPointer{m.FixedPointer}
}

// BidPrice
func (m MarketDataUpdateBidAskCompactPointer) BidPrice() float32 {
	return message.Float32Fixed(m.Ptr, 8, 4)
}

// BidPrice
func (m MarketDataUpdateBidAskCompactPointerBuilder) BidPrice() float32 {
	return message.Float32Fixed(m.Ptr, 8, 4)
}

// BidQuantity
func (m MarketDataUpdateBidAskCompactPointer) BidQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// BidQuantity
func (m MarketDataUpdateBidAskCompactPointerBuilder) BidQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// AskPrice
func (m MarketDataUpdateBidAskCompactPointer) AskPrice() float32 {
	return message.Float32Fixed(m.Ptr, 16, 12)
}

// AskPrice
func (m MarketDataUpdateBidAskCompactPointerBuilder) AskPrice() float32 {
	return message.Float32Fixed(m.Ptr, 16, 12)
}

// AskQuantity
func (m MarketDataUpdateBidAskCompactPointer) AskQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 20, 16)
}

// AskQuantity
func (m MarketDataUpdateBidAskCompactPointerBuilder) AskQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 20, 16)
}

// DateTime
func (m MarketDataUpdateBidAskCompactPointer) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 24, 20))
}

// DateTime
func (m MarketDataUpdateBidAskCompactPointerBuilder) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 24, 20))
}

// SymbolID
func (m MarketDataUpdateBidAskCompactPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 28, 24)
}

// SymbolID
func (m MarketDataUpdateBidAskCompactPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 28, 24)
}

// SetBidPrice
func (m MarketDataUpdateBidAskCompactPointerBuilder) SetBidPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 8, 4, value)
}

// SetBidQuantity
func (m MarketDataUpdateBidAskCompactPointerBuilder) SetBidQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 12, 8, value)
}

// SetAskPrice
func (m MarketDataUpdateBidAskCompactPointerBuilder) SetAskPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 16, 12, value)
}

// SetAskQuantity
func (m MarketDataUpdateBidAskCompactPointerBuilder) SetAskQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 20, 16, value)
}

// SetDateTime
func (m MarketDataUpdateBidAskCompactPointerBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 24, 20, uint32(value))
}

// SetSymbolID
func (m MarketDataUpdateBidAskCompactPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 28, 24, value)
}

// Copy
func (m MarketDataUpdateBidAskCompactPointer) Copy(to MarketDataUpdateBidAskCompactBuilder) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}

// Copy
func (m MarketDataUpdateBidAskCompactPointerBuilder) Copy(to MarketDataUpdateBidAskCompactBuilder) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}

// CopyPointer
func (m MarketDataUpdateBidAskCompactPointer) CopyPointer(to MarketDataUpdateBidAskCompactPointerBuilder) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}

// CopyPointer
func (m MarketDataUpdateBidAskCompactPointerBuilder) CopyPointer(to MarketDataUpdateBidAskCompactPointerBuilder) {
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
	to.SetSymbolID(m.SymbolID())
}
