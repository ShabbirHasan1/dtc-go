package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateBidAsk_IntSize = 28

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
var _MarketDataUpdateBidAsk_IntDefault = []byte{28, 0, 127, 0, 0, 0, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketDataUpdateBidAsk_Int struct {
	message.Fixed
}

type MarketDataUpdateBidAsk_IntBuilder struct {
	message.Fixed
}

type MarketDataUpdateBidAsk_IntPointer struct {
	message.FixedPointer
}

type MarketDataUpdateBidAsk_IntPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataUpdateBidAsk_IntFrom(b []byte) MarketDataUpdateBidAsk_Int {
	return MarketDataUpdateBidAsk_Int{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateBidAsk_Int(b []byte) MarketDataUpdateBidAsk_Int {
	return MarketDataUpdateBidAsk_Int{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateBidAsk_Int() MarketDataUpdateBidAsk_IntBuilder {
	a := MarketDataUpdateBidAsk_IntBuilder{message.NewFixed(28)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateBidAsk_IntDefault)
	return a
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
func (m MarketDataUpdateBidAsk_IntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateBidAsk_IntDefault)
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
func (m MarketDataUpdateBidAsk_Int) ToBuilder() MarketDataUpdateBidAsk_IntBuilder {
	return MarketDataUpdateBidAsk_IntBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateBidAsk_IntPointer) ToBuilder() MarketDataUpdateBidAsk_IntPointerBuilder {
	return MarketDataUpdateBidAsk_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateBidAsk_IntBuilder) Finish() MarketDataUpdateBidAsk_Int {
	return MarketDataUpdateBidAsk_Int{m.Fixed}
}

// Finish
func (m *MarketDataUpdateBidAsk_IntPointerBuilder) Finish() MarketDataUpdateBidAsk_IntPointer {
	return MarketDataUpdateBidAsk_IntPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDataUpdateBidAsk_Int) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDataUpdateBidAsk_IntBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
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
func (m MarketDataUpdateBidAsk_Int) BidPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// BidPrice
func (m MarketDataUpdateBidAsk_IntBuilder) BidPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
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
func (m MarketDataUpdateBidAsk_Int) BidQuantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// BidQuantity
func (m MarketDataUpdateBidAsk_IntBuilder) BidQuantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
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
func (m MarketDataUpdateBidAsk_Int) AskPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// AskPrice
func (m MarketDataUpdateBidAsk_IntBuilder) AskPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
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
func (m MarketDataUpdateBidAsk_Int) AskQuantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 24, 20)
}

// AskQuantity
func (m MarketDataUpdateBidAsk_IntBuilder) AskQuantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 24, 20)
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
func (m MarketDataUpdateBidAsk_Int) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 28, 24))
}

// DateTime
func (m MarketDataUpdateBidAsk_IntBuilder) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 28, 24))
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
func (m MarketDataUpdateBidAsk_IntBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetBidPrice
func (m MarketDataUpdateBidAsk_IntBuilder) SetBidPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetBidPrice
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SetBidPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetBidQuantity
func (m MarketDataUpdateBidAsk_IntBuilder) SetBidQuantity(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 16, 12, value)
}

// SetBidQuantity
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SetBidQuantity(value int32) {
	message.SetInt32Fixed(m.Ptr, 16, 12, value)
}

// SetAskPrice
func (m MarketDataUpdateBidAsk_IntBuilder) SetAskPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 20, 16, value)
}

// SetAskPrice
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SetAskPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 20, 16, value)
}

// SetAskQuantity
func (m MarketDataUpdateBidAsk_IntBuilder) SetAskQuantity(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 24, 20, value)
}

// SetAskQuantity
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SetAskQuantity(value int32) {
	message.SetInt32Fixed(m.Ptr, 24, 20, value)
}

// SetDateTime
func (m MarketDataUpdateBidAsk_IntBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 28, 24, uint32(value))
}

// SetDateTime
func (m MarketDataUpdateBidAsk_IntPointerBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 28, 24, uint32(value))
}

// Copy
func (m MarketDataUpdateBidAsk_Int) Copy(to MarketDataUpdateBidAsk_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateBidAsk_IntBuilder) Copy(to MarketDataUpdateBidAsk_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateBidAsk_Int) CopyPointer(to MarketDataUpdateBidAsk_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateBidAsk_IntBuilder) CopyPointer(to MarketDataUpdateBidAsk_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
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
