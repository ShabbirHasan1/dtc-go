package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDataUpdateOpenInterestPointer The MarketDataUpdateOpenInterest message is sent by the Server to the
// Client to update the OpenInterest field previously sent through the MarketDataSnapshot
// message.
type MarketDataUpdateOpenInterestPointer struct {
	message.FixedPointer
}

// MarketDataUpdateOpenInterestPointerBuilder The MarketDataUpdateOpenInterest message is sent by the Server to the
// Client to update the OpenInterest field previously sent through the MarketDataSnapshot
// message.
type MarketDataUpdateOpenInterestPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataUpdateOpenInterest() MarketDataUpdateOpenInterestPointerBuilder {
	a := MarketDataUpdateOpenInterestPointerBuilder{message.AllocFixed(16)}
	a.Ptr.SetBytes(0, _MarketDataUpdateOpenInterestDefault)
	return a
}

func AllocMarketDataUpdateOpenInterestFrom(b []byte) MarketDataUpdateOpenInterestPointer {
	return MarketDataUpdateOpenInterestPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size               = MarketDataUpdateOpenInterestSize  (16)
//     Type               = MARKET_DATA_UPDATE_OPEN_INTEREST  (124)
//     SymbolID           = 0
//     OpenInterest       = 0
//     TradingSessionDate = 0
// }
func (m MarketDataUpdateOpenInterestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateOpenInterestDefault)
}

// ToBuilder
func (m MarketDataUpdateOpenInterestPointer) ToBuilder() MarketDataUpdateOpenInterestPointerBuilder {
	return MarketDataUpdateOpenInterestPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateOpenInterestPointerBuilder) Finish() MarketDataUpdateOpenInterestPointer {
	return MarketDataUpdateOpenInterestPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateOpenInterestPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateOpenInterestPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// OpenInterest The open interest for the symbol.
func (m MarketDataUpdateOpenInterestPointer) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// OpenInterest The open interest for the symbol.
func (m MarketDataUpdateOpenInterestPointerBuilder) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// TradingSessionDate
func (m MarketDataUpdateOpenInterestPointer) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 16, 12))
}

// TradingSessionDate
func (m MarketDataUpdateOpenInterestPointerBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 16, 12))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateOpenInterestPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetOpenInterest The open interest for the symbol.
func (m MarketDataUpdateOpenInterestPointerBuilder) SetOpenInterest(value uint32) {
	message.SetUint32Fixed(m.Ptr, 12, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateOpenInterestPointerBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 16, 12, uint32(value))
}

// Copy
func (m MarketDataUpdateOpenInterestPointer) Copy(to MarketDataUpdateOpenInterestBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetOpenInterest(m.OpenInterest())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateOpenInterestPointerBuilder) Copy(to MarketDataUpdateOpenInterestBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetOpenInterest(m.OpenInterest())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateOpenInterestPointer) CopyPointer(to MarketDataUpdateOpenInterestPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetOpenInterest(m.OpenInterest())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateOpenInterestPointerBuilder) CopyPointer(to MarketDataUpdateOpenInterestPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetOpenInterest(m.OpenInterest())
	to.SetTradingSessionDate(m.TradingSessionDate())
}
