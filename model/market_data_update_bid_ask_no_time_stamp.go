package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size        = MarketDataUpdateBidAskNoTimeStampSize  (24)
//     Type        = MARKET_DATA_UPDATE_BID_ASK_NO_TIMESTAMP  (143)
//     SymbolID    = 0
//     BidPrice    = math.MaxFloat32
//     BidQuantity = 0
//     AskPrice    = math.MaxFloat32
//     AskQuantity = 0
// }
var _MarketDataUpdateBidAskNoTimeStampDefault = []byte{24, 0, 143, 0, 0, 0, 0, 0, 255, 255, 127, 127, 0, 0, 0, 0, 255, 255, 127, 127, 0, 0, 0, 0}

const MarketDataUpdateBidAskNoTimeStampSize = 24

// MarketDataUpdateBidAskNoTimeStamp This message is optional.
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
type MarketDataUpdateBidAskNoTimeStamp struct {
	message.Fixed
}

// MarketDataUpdateBidAskNoTimeStampBuilder This message is optional.
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
type MarketDataUpdateBidAskNoTimeStampBuilder struct {
	message.Fixed
}

func NewMarketDataUpdateBidAskNoTimeStampFrom(b []byte) MarketDataUpdateBidAskNoTimeStamp {
	return MarketDataUpdateBidAskNoTimeStamp{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateBidAskNoTimeStamp(b []byte) MarketDataUpdateBidAskNoTimeStamp {
	return MarketDataUpdateBidAskNoTimeStamp{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateBidAskNoTimeStamp() MarketDataUpdateBidAskNoTimeStampBuilder {
	a := MarketDataUpdateBidAskNoTimeStampBuilder{message.NewFixed(24)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateBidAskNoTimeStampDefault)
	return a
}

// Clear
// {
//     Size        = MarketDataUpdateBidAskNoTimeStampSize  (24)
//     Type        = MARKET_DATA_UPDATE_BID_ASK_NO_TIMESTAMP  (143)
//     SymbolID    = 0
//     BidPrice    = math.MaxFloat32
//     BidQuantity = 0
//     AskPrice    = math.MaxFloat32
//     AskQuantity = 0
// }
func (m MarketDataUpdateBidAskNoTimeStampBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateBidAskNoTimeStampDefault)
}

// ToBuilder
func (m MarketDataUpdateBidAskNoTimeStamp) ToBuilder() MarketDataUpdateBidAskNoTimeStampBuilder {
	return MarketDataUpdateBidAskNoTimeStampBuilder{m.Fixed}
}

// Finish
func (m MarketDataUpdateBidAskNoTimeStampBuilder) Finish() MarketDataUpdateBidAskNoTimeStamp {
	return MarketDataUpdateBidAskNoTimeStamp{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskNoTimeStamp) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskNoTimeStampBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskNoTimeStamp) BidPrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// BidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskNoTimeStampBuilder) BidPrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 12, 8)
}

// BidQuantity The current number of contracts/shares at the Bid price.
func (m MarketDataUpdateBidAskNoTimeStamp) BidQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 16, 12)
}

// BidQuantity The current number of contracts/shares at the Bid price.
func (m MarketDataUpdateBidAskNoTimeStampBuilder) BidQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 16, 12)
}

// AskPrice The current Ask or offer price. Leave unset if there is no price available.
// The current Ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskNoTimeStamp) AskPrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 20, 16)
}

// AskPrice The current Ask or offer price. Leave unset if there is no price available.
// The current Ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskNoTimeStampBuilder) AskPrice() float32 {
	return message.Float32Fixed(m.Unsafe(), 20, 16)
}

// AskQuantity The current number of contracts/shares at the Ask price.
func (m MarketDataUpdateBidAskNoTimeStamp) AskQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 24, 20)
}

// AskQuantity The current number of contracts/shares at the Ask price.
func (m MarketDataUpdateBidAskNoTimeStampBuilder) AskQuantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 24, 20)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateBidAskNoTimeStampBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetBidPrice The current Bid price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskNoTimeStampBuilder) SetBidPrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 12, 8, value)
}

// SetBidQuantity The current number of contracts/shares at the Bid price.
func (m MarketDataUpdateBidAskNoTimeStampBuilder) SetBidQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 16, 12, value)
}

// SetAskPrice The current Ask or offer price. Leave unset if there is no price available.
// The current Ask or offer price. Leave unset if there is no price available.
func (m MarketDataUpdateBidAskNoTimeStampBuilder) SetAskPrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 20, 16, value)
}

// SetAskQuantity The current number of contracts/shares at the Ask price.
func (m MarketDataUpdateBidAskNoTimeStampBuilder) SetAskQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 24, 20, value)
}

// Copy
func (m MarketDataUpdateBidAskNoTimeStamp) Copy(to MarketDataUpdateBidAskNoTimeStampBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
}

// Copy
func (m MarketDataUpdateBidAskNoTimeStampBuilder) Copy(to MarketDataUpdateBidAskNoTimeStampBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
}

// CopyPointer
func (m MarketDataUpdateBidAskNoTimeStamp) CopyPointer(to MarketDataUpdateBidAskNoTimeStampPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
}

// CopyPointer
func (m MarketDataUpdateBidAskNoTimeStampBuilder) CopyPointer(to MarketDataUpdateBidAskNoTimeStampPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetBidPrice(m.BidPrice())
	to.SetBidQuantity(m.BidQuantity())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
}
