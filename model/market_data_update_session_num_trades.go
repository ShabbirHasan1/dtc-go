package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateSessionNumTradesSize = 16

// {
//     Size               = MarketDataUpdateSessionNumTradesSize  (16)
//     Type               = MARKET_DATA_UPDATE_SESSION_NUM_TRADES  (135)
//     SymbolID           = 0
//     NumTrades          = 0
//     TradingSessionDate = 0
// }
var _MarketDataUpdateSessionNumTradesDefault = []byte{16, 0, 135, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateSessionNumTrades Sent by the Server to the Client to update the trading session number
// of trades.
type MarketDataUpdateSessionNumTrades struct {
	message.Fixed
}

// MarketDataUpdateSessionNumTradesBuilder Sent by the Server to the Client to update the trading session number
// of trades.
type MarketDataUpdateSessionNumTradesBuilder struct {
	message.Fixed
}

// MarketDataUpdateSessionNumTradesPointer Sent by the Server to the Client to update the trading session number
// of trades.
type MarketDataUpdateSessionNumTradesPointer struct {
	message.FixedPointer
}

// MarketDataUpdateSessionNumTradesPointerBuilder Sent by the Server to the Client to update the trading session number
// of trades.
type MarketDataUpdateSessionNumTradesPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataUpdateSessionNumTradesFrom(b []byte) MarketDataUpdateSessionNumTrades {
	return MarketDataUpdateSessionNumTrades{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateSessionNumTrades(b []byte) MarketDataUpdateSessionNumTrades {
	return MarketDataUpdateSessionNumTrades{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateSessionNumTrades() MarketDataUpdateSessionNumTradesBuilder {
	a := MarketDataUpdateSessionNumTradesBuilder{message.NewFixed(16)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateSessionNumTradesDefault)
	return a
}

func AllocMarketDataUpdateSessionNumTrades() MarketDataUpdateSessionNumTradesPointerBuilder {
	a := MarketDataUpdateSessionNumTradesPointerBuilder{message.AllocFixed(16)}
	a.Ptr.SetBytes(0, _MarketDataUpdateSessionNumTradesDefault)
	return a
}

func AllocMarketDataUpdateSessionNumTradesFrom(b []byte) MarketDataUpdateSessionNumTradesPointer {
	return MarketDataUpdateSessionNumTradesPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size               = MarketDataUpdateSessionNumTradesSize  (16)
//     Type               = MARKET_DATA_UPDATE_SESSION_NUM_TRADES  (135)
//     SymbolID           = 0
//     NumTrades          = 0
//     TradingSessionDate = 0
// }
func (m MarketDataUpdateSessionNumTradesBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateSessionNumTradesDefault)
}

// Clear
// {
//     Size               = MarketDataUpdateSessionNumTradesSize  (16)
//     Type               = MARKET_DATA_UPDATE_SESSION_NUM_TRADES  (135)
//     SymbolID           = 0
//     NumTrades          = 0
//     TradingSessionDate = 0
// }
func (m MarketDataUpdateSessionNumTradesPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateSessionNumTradesDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionNumTrades) ToBuilder() MarketDataUpdateSessionNumTradesBuilder {
	return MarketDataUpdateSessionNumTradesBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateSessionNumTradesPointer) ToBuilder() MarketDataUpdateSessionNumTradesPointerBuilder {
	return MarketDataUpdateSessionNumTradesPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateSessionNumTradesBuilder) Finish() MarketDataUpdateSessionNumTrades {
	return MarketDataUpdateSessionNumTrades{m.Fixed}
}

// Finish
func (m *MarketDataUpdateSessionNumTradesPointerBuilder) Finish() MarketDataUpdateSessionNumTradesPointer {
	return MarketDataUpdateSessionNumTradesPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionNumTrades) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionNumTradesBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionNumTradesPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionNumTradesPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// NumTrades The number of trades which have occurred during the current trading session.
// The number of trades which have occurred during the current trading session.
func (m MarketDataUpdateSessionNumTrades) NumTrades() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// NumTrades The number of trades which have occurred during the current trading session.
// The number of trades which have occurred during the current trading session.
func (m MarketDataUpdateSessionNumTradesBuilder) NumTrades() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// NumTrades The number of trades which have occurred during the current trading session.
// The number of trades which have occurred during the current trading session.
func (m MarketDataUpdateSessionNumTradesPointer) NumTrades() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// NumTrades The number of trades which have occurred during the current trading session.
// The number of trades which have occurred during the current trading session.
func (m MarketDataUpdateSessionNumTradesPointerBuilder) NumTrades() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// TradingSessionDate
func (m MarketDataUpdateSessionNumTrades) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 16, 12))
}

// TradingSessionDate
func (m MarketDataUpdateSessionNumTradesBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 16, 12))
}

// TradingSessionDate
func (m MarketDataUpdateSessionNumTradesPointer) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 16, 12))
}

// TradingSessionDate
func (m MarketDataUpdateSessionNumTradesPointerBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 16, 12))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionNumTradesBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionNumTradesPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetNumTrades The number of trades which have occurred during the current trading session.
// The number of trades which have occurred during the current trading session.
func (m MarketDataUpdateSessionNumTradesBuilder) SetNumTrades(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetNumTrades The number of trades which have occurred during the current trading session.
// The number of trades which have occurred during the current trading session.
func (m MarketDataUpdateSessionNumTradesPointerBuilder) SetNumTrades(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionNumTradesBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 16, 12, uint32(value))
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionNumTradesPointerBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 16, 12, uint32(value))
}

// Copy
func (m MarketDataUpdateSessionNumTrades) Copy(to MarketDataUpdateSessionNumTradesBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetNumTrades(m.NumTrades())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateSessionNumTradesBuilder) Copy(to MarketDataUpdateSessionNumTradesBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetNumTrades(m.NumTrades())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionNumTrades) CopyPointer(to MarketDataUpdateSessionNumTradesPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetNumTrades(m.NumTrades())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionNumTradesBuilder) CopyPointer(to MarketDataUpdateSessionNumTradesPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetNumTrades(m.NumTrades())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateSessionNumTradesPointer) Copy(to MarketDataUpdateSessionNumTradesBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetNumTrades(m.NumTrades())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// Copy
func (m MarketDataUpdateSessionNumTradesPointerBuilder) Copy(to MarketDataUpdateSessionNumTradesBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetNumTrades(m.NumTrades())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionNumTradesPointer) CopyPointer(to MarketDataUpdateSessionNumTradesPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetNumTrades(m.NumTrades())
	to.SetTradingSessionDate(m.TradingSessionDate())
}

// CopyPointer
func (m MarketDataUpdateSessionNumTradesPointerBuilder) CopyPointer(to MarketDataUpdateSessionNumTradesPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetNumTrades(m.NumTrades())
	to.SetTradingSessionDate(m.TradingSessionDate())
}
