package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDataUpdateBidAskFloatWithMicrosecondsPointer This message is optional.
//
// Sent by the Server to the Client when there is an update to the Bid Ask
// prices and/or quantities. This message is identical to the MarketDataUpdateBidAsk
// message except it does not have a timestamp. It needs to be sent when
// there is no change with the timestamp for the Bid Ask update as compared
// to the prior update.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received Bid Ask update timestamp to know what the timestamp
// is for this message.
type MarketDataUpdateBidAskFloatWithMicrosecondsPointer struct {
	message.FixedPointer
}

// MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder This message is optional.
//
// Sent by the Server to the Client when there is an update to the Bid Ask
// prices and/or quantities. This message is identical to the MarketDataUpdateBidAsk
// message except it does not have a timestamp. It needs to be sent when
// there is no change with the timestamp for the Bid Ask update as compared
// to the prior update.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received Bid Ask update timestamp to know what the timestamp
// is for this message.
type MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataUpdateBidAskFloatWithMicroseconds() MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder {
	a := MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder{message.AllocFixed(32)}
	a.Ptr.SetBytes(0, _MarketDataUpdateBidAskFloatWithMicrosecondsDefault)
	return a
}

func AllocMarketDataUpdateBidAskFloatWithMicrosecondsFrom(b []byte) MarketDataUpdateBidAskFloatWithMicrosecondsPointer {
	return MarketDataUpdateBidAskFloatWithMicrosecondsPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size        = MarketDataUpdateBidAskFloatWithMicrosecondsSize  (32)
//     Type        = MARKET_DATA_UPDATE_BID_ASK_FLOAT_WITH_MICROSECONDS  (144)
//     SymbolID    = 0
//     BidPrice    = math.MaxFloat32
//     BidQuantity = 0
//     AskPrice    = math.MaxFloat32
//     AskQuantity = 0
//     DateTime    = 0
// }
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateBidAskFloatWithMicrosecondsDefault)
}

// ToBuilder
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointer) ToBuilder() MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder {
	return MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) Finish() MarketDataUpdateBidAskFloatWithMicrosecondsPointer {
	return MarketDataUpdateBidAskFloatWithMicrosecondsPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointer) BidPrice() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) BidPrice() float32 {
	return message.Float32Fixed(m.Ptr, 12, 8)
}

// BidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointer) BidQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 16, 12)
}

// BidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) BidQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 16, 12)
}

// AskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointer) AskPrice() float32 {
	return message.Float32Fixed(m.Ptr, 20, 16)
}

// AskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) AskPrice() float32 {
	return message.Float32Fixed(m.Ptr, 20, 16)
}

// AskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointer) AskQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 24, 20)
}

// AskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) AskQuantity() float32 {
	return message.Float32Fixed(m.Ptr, 24, 20)
}

// DateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointer) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 32, 24))
}

// DateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 32, 24))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetBidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) SetBidPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 12, 8, value)
}

// SetBidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) SetBidQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 16, 12, value)
}

// SetAskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) SetAskPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 20, 16, value)
}

// SetAskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) SetAskQuantity(value float32) {
	message.SetFloat32Fixed(m.Ptr, 24, 20, value)
}

// SetDateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 32, 24, int64(value))
}

// Copy
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointer) Copy(to MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) Copy(to MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointer) CopyPointer(to MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) CopyPointer(to MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}
