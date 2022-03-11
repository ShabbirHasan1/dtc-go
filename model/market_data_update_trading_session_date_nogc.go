package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDataUpdateTradingSessionDatePointer Sent by the Server to the Client to update the trading session Date.
type MarketDataUpdateTradingSessionDatePointer struct {
	message.FixedPointer
}

// MarketDataUpdateTradingSessionDatePointerBuilder Sent by the Server to the Client to update the trading session Date.
type MarketDataUpdateTradingSessionDatePointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataUpdateTradingSessionDate() MarketDataUpdateTradingSessionDatePointerBuilder {
	a := MarketDataUpdateTradingSessionDatePointerBuilder{message.AllocFixed(12)}
	a.Ptr.SetBytes(0, _MarketDataUpdateTradingSessionDateDefault)
	return a
}

func AllocMarketDataUpdateTradingSessionDateFrom(b []byte) MarketDataUpdateTradingSessionDatePointer {
	return MarketDataUpdateTradingSessionDatePointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size     = MarketDataUpdateTradingSessionDateSize  (12)
//     Type     = MARKET_DATA_UPDATE_TRADING_SESSION_DATE  (136)
//     SymbolID = 0
//     Date     = 0
// }
func (m MarketDataUpdateTradingSessionDatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateTradingSessionDateDefault)
}

// ToBuilder
func (m MarketDataUpdateTradingSessionDatePointer) ToBuilder() MarketDataUpdateTradingSessionDatePointerBuilder {
	return MarketDataUpdateTradingSessionDatePointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateTradingSessionDatePointerBuilder) Finish() MarketDataUpdateTradingSessionDatePointer {
	return MarketDataUpdateTradingSessionDatePointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradingSessionDatePointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradingSessionDatePointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Date The date of the current trading session. The time component is not normally
// considered relevant in this case.
func (m MarketDataUpdateTradingSessionDatePointer) Date() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 12, 8))
}

// Date The date of the current trading session. The time component is not normally
// considered relevant in this case.
func (m MarketDataUpdateTradingSessionDatePointerBuilder) Date() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 12, 8))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTradingSessionDatePointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetDate The date of the current trading session. The time component is not normally
// considered relevant in this case.
func (m MarketDataUpdateTradingSessionDatePointerBuilder) SetDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 12, 8, uint32(value))
}

// Copy
func (m MarketDataUpdateTradingSessionDatePointer) Copy(to MarketDataUpdateTradingSessionDateBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetDate(m.Date())
}

// Copy
func (m MarketDataUpdateTradingSessionDatePointerBuilder) Copy(to MarketDataUpdateTradingSessionDateBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetDate(m.Date())
}

// CopyPointer
func (m MarketDataUpdateTradingSessionDatePointer) CopyPointer(to MarketDataUpdateTradingSessionDatePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetDate(m.Date())
}

// CopyPointer
func (m MarketDataUpdateTradingSessionDatePointerBuilder) CopyPointer(to MarketDataUpdateTradingSessionDatePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetDate(m.Date())
}
