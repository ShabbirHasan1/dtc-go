package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size               = MarketDataUpdateOpenInterestSize  (16)
//     Type               = MARKET_DATA_UPDATE_OPEN_INTEREST  (124)
//     SymbolID           = 0
//     OpenInterest       = 0
//     TradingSessionDate = 0
// }
var _MarketDataUpdateOpenInterestDefault = []byte{16, 0, 124, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataUpdateOpenInterestSize = 16

// MarketDataUpdateOpenInterest The MarketDataUpdateOpenInterest message is sent by the Server to the
// Client to update the OpenInterest field previously sent through the MarketDataSnapshot
// message.
type MarketDataUpdateOpenInterest struct {
	message.Fixed
}

// MarketDataUpdateOpenInterestBuilder The MarketDataUpdateOpenInterest message is sent by the Server to the
// Client to update the OpenInterest field previously sent through the MarketDataSnapshot
// message.
type MarketDataUpdateOpenInterestBuilder struct {
	message.Fixed
}

func NewMarketDataUpdateOpenInterestFrom(b []byte) MarketDataUpdateOpenInterest {
	return MarketDataUpdateOpenInterest{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateOpenInterest(b []byte) MarketDataUpdateOpenInterest {
	return MarketDataUpdateOpenInterest{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateOpenInterest() MarketDataUpdateOpenInterestBuilder {
	a := MarketDataUpdateOpenInterestBuilder{message.NewFixed(16)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateOpenInterestDefault)
	return a
}

// Clear
// {
//     Size               = MarketDataUpdateOpenInterestSize  (16)
//     Type               = MARKET_DATA_UPDATE_OPEN_INTEREST  (124)
//     SymbolID           = 0
//     OpenInterest       = 0
//     TradingSessionDate = 0
// }
func (m MarketDataUpdateOpenInterestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateOpenInterestDefault)
}

// ToBuilder
func (m MarketDataUpdateOpenInterest) ToBuilder() MarketDataUpdateOpenInterestBuilder {
	return MarketDataUpdateOpenInterestBuilder{m.Fixed}
}

// Finish
func (m MarketDataUpdateOpenInterestBuilder) Finish() MarketDataUpdateOpenInterest {
	return MarketDataUpdateOpenInterest{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateOpenInterest) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateOpenInterestBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// OpenInterest The open interest for the symbol.
func (m MarketDataUpdateOpenInterest) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// OpenInterest The open interest for the symbol.
func (m MarketDataUpdateOpenInterestBuilder) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// TradingSessionDate
func (m MarketDataUpdateOpenInterest) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 16, 12))
}

// TradingSessionDate
func (m MarketDataUpdateOpenInterestBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 16, 12))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateOpenInterestBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetOpenInterest The open interest for the symbol.
func (m MarketDataUpdateOpenInterestBuilder) SetOpenInterest(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 12, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateOpenInterestBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 16, 12, uint32(value))
}

// Copy
func (m MarketDataUpdateOpenInterest) Copy(to MarketDataUpdateOpenInterestBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetOpenInterest(m.OpenInterest())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateOpenInterestBuilder) Copy(to MarketDataUpdateOpenInterestBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetOpenInterest(m.OpenInterest())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateOpenInterest) CopyPointer(to MarketDataUpdateOpenInterestPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetOpenInterest(m.OpenInterest())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateOpenInterestBuilder) CopyPointer(to MarketDataUpdateOpenInterestPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetOpenInterest(m.OpenInterest())
	to.SetTradingSessionDate(m.TradingSessionDate())
}
