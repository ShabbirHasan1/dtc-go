package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size     = MarketDataUpdateTradingSessionDateSize  (12)
//     Type     = MARKET_DATA_UPDATE_TRADING_SESSION_DATE  (136)
//     SymbolID = 0
//     Date     = 0
// }
var _MarketDataUpdateTradingSessionDateDefault = []byte{12, 0, 136, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataUpdateTradingSessionDateSize = 12

// MarketDataUpdateTradingSessionDate Sent by the Server to the Client to update the trading session Date.
type MarketDataUpdateTradingSessionDate struct {
	message.Fixed
}

// MarketDataUpdateTradingSessionDateBuilder Sent by the Server to the Client to update the trading session Date.
type MarketDataUpdateTradingSessionDateBuilder struct {
	message.Fixed
}

func NewMarketDataUpdateTradingSessionDateFrom(b []byte) MarketDataUpdateTradingSessionDate {
	return MarketDataUpdateTradingSessionDate{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateTradingSessionDate(b []byte) MarketDataUpdateTradingSessionDate {
	return MarketDataUpdateTradingSessionDate{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateTradingSessionDate() MarketDataUpdateTradingSessionDateBuilder {
	a := MarketDataUpdateTradingSessionDateBuilder{message.NewFixed(12)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateTradingSessionDateDefault)
	return a
}

// Clear
// {
//     Size     = MarketDataUpdateTradingSessionDateSize  (12)
//     Type     = MARKET_DATA_UPDATE_TRADING_SESSION_DATE  (136)
//     SymbolID = 0
//     Date     = 0
// }
func (m MarketDataUpdateTradingSessionDateBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateTradingSessionDateDefault)
}

// ToBuilder
func (m MarketDataUpdateTradingSessionDate) ToBuilder() MarketDataUpdateTradingSessionDateBuilder {
	return MarketDataUpdateTradingSessionDateBuilder{m.Fixed}
}

// Finish
func (m MarketDataUpdateTradingSessionDateBuilder) Finish() MarketDataUpdateTradingSessionDate {
	return MarketDataUpdateTradingSessionDate{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradingSessionDate) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradingSessionDateBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Date The date of the current trading session. The time component is not normally
// considered relevant in this case.
func (m MarketDataUpdateTradingSessionDate) Date() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 12, 8))
}

// Date The date of the current trading session. The time component is not normally
// considered relevant in this case.
func (m MarketDataUpdateTradingSessionDateBuilder) Date() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 12, 8))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradingSessionDateBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetDate The date of the current trading session. The time component is not normally
// considered relevant in this case.
func (m MarketDataUpdateTradingSessionDateBuilder) SetDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 12, 8, uint32(value))
}

// Copy
func (m MarketDataUpdateTradingSessionDate) Copy(to MarketDataUpdateTradingSessionDateBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetDate(m.Date())
}

// Copy
func (m MarketDataUpdateTradingSessionDateBuilder) Copy(to MarketDataUpdateTradingSessionDateBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetDate(m.Date())
}

// CopyPointer
func (m MarketDataUpdateTradingSessionDate) CopyPointer(to MarketDataUpdateTradingSessionDatePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetDate(m.Date())
}

// CopyPointer
func (m MarketDataUpdateTradingSessionDateBuilder) CopyPointer(to MarketDataUpdateTradingSessionDatePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetDate(m.Date())
}
