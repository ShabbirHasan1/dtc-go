package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _MarketDataUpdateBidAskFloatWithMicrosecondsDefault = []byte{32, 0, 144, 0, 0, 0, 0, 0, 255, 255, 127, 127, 0, 0, 0, 0, 255, 255, 127, 127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataUpdateBidAskFloatWithMicrosecondsSize = 32

// MarketDataUpdateBidAskFloatWithMicroseconds This message is optional.
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
type MarketDataUpdateBidAskFloatWithMicroseconds struct {
	message.Fixed
}

// MarketDataUpdateBidAskFloatWithMicrosecondsBuilder This message is optional.
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
type MarketDataUpdateBidAskFloatWithMicrosecondsBuilder struct {
	message.Fixed
}

func NewMarketDataUpdateBidAskFloatWithMicrosecondsFrom(b []byte) MarketDataUpdateBidAskFloatWithMicroseconds {
	return MarketDataUpdateBidAskFloatWithMicroseconds{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateBidAskFloatWithMicroseconds(b []byte) MarketDataUpdateBidAskFloatWithMicroseconds {
	return MarketDataUpdateBidAskFloatWithMicroseconds{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateBidAskFloatWithMicroseconds() MarketDataUpdateBidAskFloatWithMicrosecondsBuilder {
	a := MarketDataUpdateBidAskFloatWithMicrosecondsBuilder{message.NewFixed(32)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateBidAskFloatWithMicrosecondsDefault)
	return a
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
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateBidAskFloatWithMicrosecondsDefault)
}

// ToBuilder
func (m MarketDataUpdateBidAskFloatWithMicroseconds) ToBuilder() MarketDataUpdateBidAskFloatWithMicrosecondsBuilder {
	return MarketDataUpdateBidAskFloatWithMicrosecondsBuilder{m.Fixed}
}

// Finish
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) Finish() MarketDataUpdateBidAskFloatWithMicroseconds {
	return MarketDataUpdateBidAskFloatWithMicroseconds{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) BidPrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) BidPrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// BidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) BidQuantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 16, 12)
}

// BidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) BidQuantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 16, 12)
}

// AskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) AskPrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 20, 16)
}

// AskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) AskPrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 20, 16)
}

// AskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) AskQuantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 24, 20)
}

// AskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) AskQuantity() float32 {
	return message.Float32Fixed(m.Unsafe(), 24, 20)
}

// DateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateBidAskFloatWithMicroseconds) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 32, 24))
}

// DateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 32, 24))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetBidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) SetBidPrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 12, 8, value)
}

// SetBidQuantity The current number of contracts/shares at the bid price.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) SetBidQuantity(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 16, 12, value)
}

// SetAskPrice The current ask or offer price. Leave unset if there is no price available.
// The current ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) SetAskPrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 20, 16, value)
}

// SetAskQuantity The current number of contracts/shares at the ask price.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) SetAskQuantity(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 24, 20, value)
}

// SetDateTime The timestamp of the trade in UNIX microseconds time format.
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 32, 24, int64(value))
}

// Copy
func (m MarketDataUpdateBidAskFloatWithMicroseconds) Copy(to MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) Copy(to MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateBidAskFloatWithMicroseconds) CopyPointer(to MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateBidAskFloatWithMicrosecondsBuilder) CopyPointer(to MarketDataUpdateBidAskFloatWithMicrosecondsPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetDateTime(m.DateTime())
}
