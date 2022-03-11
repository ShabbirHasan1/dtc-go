package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketDataUpdateBidAsk_IntPointer struct {
	message.FixedPointer
}

type MarketDataUpdateBidAsk_IntPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataUpdateBidAsk_Int() MarketDataUpdateBidAsk_IntPointerBuilder {
	a := MarketDataUpdateBidAsk_IntPointerBuilder{message.AllocFixed(28)}
	a.Ptr.SetBytes(0, _MarketDataUpdateBidAsk_IntDefault)
	return a
}

func AllocMarketDataUpdateBidAsk_IntFrom(b []byte) MarketDataUpdateBidAsk_IntPointer {
	return MarketDataUpdateBidAsk_IntPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size        = MarketDataUpdateBidAsk_IntSize  (28)
//     Type        = MARKET_DATA_UPDATE_BID_ASK_INT  (127)
//     SymbolID    = 0
//     BidPrice    = 32767
//     BidQuantity = 0
//     AskPrice    = 32767
//     AskQuantity = 0
//     DateTime    = 0
// }
func (m MarketDataUpdateBidAsk_IntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateBidAsk_IntDefault)
}

// ToBuilder
func (m MarketDataUpdateBidAsk_IntPointer) ToBuilder() MarketDataUpdateBidAsk_IntPointerBuilder {
	return MarketDataUpdateBidAsk_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateBidAsk_IntPointerBuilder) Finish() MarketDataUpdateBidAsk_IntPointer {
	return MarketDataUpdateBidAsk_IntPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDataUpdateBidAsk_IntPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// BidPrice
func (m MarketDataUpdateBidAsk_IntPointer) BidPrice() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// BidPrice
func (m MarketDataUpdateBidAsk_IntPointerBuilder) BidPrice() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// BidQuantity
func (m MarketDataUpdateBidAsk_IntPointer) BidQuantity() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// BidQuantity
func (m MarketDataUpdateBidAsk_IntPointerBuilder) BidQuantity() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// AskPrice
func (m MarketDataUpdateBidAsk_IntPointer) AskPrice() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// AskPrice
func (m MarketDataUpdateBidAsk_IntPointerBuilder) AskPrice() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// AskQuantity
func (m MarketDataUpdateBidAsk_IntPointer) AskQuantity() int32 {
	return message.Int32Fixed(m.Ptr, 24, 20)
}

// AskQuantity
func (m MarketDataUpdateBidAsk_IntPointerBuilder) AskQuantity() int32 {
	return message.Int32Fixed(m.Ptr, 24, 20)
}

// DateTime
func (m MarketDataUpdateBidAsk_IntPointer) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 28, 24))
}

// DateTime
func (m MarketDataUpdateBidAsk_IntPointerBuilder) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 28, 24))
}

// SetSymbolID
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetBidPrice
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SetBidPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetBidQuantity
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SetBidQuantity(value int32) {
	message.SetInt32Fixed(m.Ptr, 16, 12, value)
}

// SetAskPrice
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SetAskPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 20, 16, value)
}

// SetAskQuantity
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SetAskQuantity(value int32) {
	message.SetInt32Fixed(m.Ptr, 24, 20, value)
}

// SetDateTime
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 28, 24, uint32(value))
}

// Copy
func (m MarketDataUpdateBidAsk_IntPointer) Copy(to MarketDataUpdateBidAsk_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateBidAsk_IntPointerBuilder) Copy(to MarketDataUpdateBidAsk_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateBidAsk_IntPointer) CopyPointer(to MarketDataUpdateBidAsk_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateBidAsk_IntPointerBuilder) CopyPointer(to MarketDataUpdateBidAsk_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}
