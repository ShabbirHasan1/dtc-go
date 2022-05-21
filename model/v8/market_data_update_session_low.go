// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-16 08:00:54.519198 +0800 WITA m=+0.055446543

package v8

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateSessionLowSize = 24

//     Size                uint16         = MarketDataUpdateSessionLowSize  (24)
//     Type                uint16         = MARKET_DATA_UPDATE_SESSION_LOW  (115)
//     SymbolID            uint32         = 0
//     Price               float64        = 0
//     TradingSessionDate  DateTime4Byte  = 0
var _MarketDataUpdateSessionLowDefault = []byte{24, 0, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateSessionLow Sent by the Server to the Client to update the session Low as the Low
// price changes throughout the session.
type MarketDataUpdateSessionLow struct {
	message.Fixed
}

// MarketDataUpdateSessionLowBuilder Sent by the Server to the Client to update the session Low as the Low
// price changes throughout the session.
type MarketDataUpdateSessionLowBuilder struct {
	message.Fixed
}

// MarketDataUpdateSessionLowPointer Sent by the Server to the Client to update the session Low as the Low
// price changes throughout the session.
type MarketDataUpdateSessionLowPointer struct {
	message.FixedPointer
}

// MarketDataUpdateSessionLowPointerBuilder Sent by the Server to the Client to update the session Low as the Low
// price changes throughout the session.
type MarketDataUpdateSessionLowPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataUpdateSessionLowFrom(b []byte) MarketDataUpdateSessionLow {
	return MarketDataUpdateSessionLow{Fixed: message.NewFixed(b)}
}

func WrapMarketDataUpdateSessionLow(b []byte) MarketDataUpdateSessionLow {
	return MarketDataUpdateSessionLow{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateSessionLow() MarketDataUpdateSessionLowBuilder {
	return MarketDataUpdateSessionLowBuilder{message.NewFixed(_MarketDataUpdateSessionLowDefault)}
}

func AllocMarketDataUpdateSessionLow() MarketDataUpdateSessionLowPointerBuilder {
	return MarketDataUpdateSessionLowPointerBuilder{message.AllocFixed(_MarketDataUpdateSessionLowDefault)}
}

func AllocMarketDataUpdateSessionLowFrom(b []byte) MarketDataUpdateSessionLowPointer {
	return MarketDataUpdateSessionLowPointer{FixedPointer: message.AllocFixed(b)}
}

// Clear
//     Size                uint16         = MarketDataUpdateSessionLowSize  (24)
//     Type                uint16         = MARKET_DATA_UPDATE_SESSION_LOW  (115)
//     SymbolID            uint32         = 0
//     Price               float64        = 0
//     TradingSessionDate  DateTime4Byte  = 0
func (m MarketDataUpdateSessionLowBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateSessionLowDefault)
}

// Clear
//     Size                uint16         = MarketDataUpdateSessionLowSize  (24)
//     Type                uint16         = MARKET_DATA_UPDATE_SESSION_LOW  (115)
//     SymbolID            uint32         = 0
//     Price               float64        = 0
//     TradingSessionDate  DateTime4Byte  = 0
func (m MarketDataUpdateSessionLowPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateSessionLowDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionLow) ToBuilder() MarketDataUpdateSessionLowBuilder {
	return MarketDataUpdateSessionLowBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateSessionLowPointer) ToBuilder() MarketDataUpdateSessionLowPointerBuilder {
	return MarketDataUpdateSessionLowPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateSessionLowBuilder) Finish() MarketDataUpdateSessionLow {
	return MarketDataUpdateSessionLow{m.Fixed}
}

// Finish
func (m *MarketDataUpdateSessionLowPointerBuilder) Finish() MarketDataUpdateSessionLowPointer {
	return MarketDataUpdateSessionLowPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionLow) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionLowBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionLowPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionLowPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Price The session Low price.
func (m MarketDataUpdateSessionLow) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// Price The session Low price.
func (m MarketDataUpdateSessionLowBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// Price The session Low price.
func (m MarketDataUpdateSessionLowPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// Price The session Low price.
func (m MarketDataUpdateSessionLowPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// TradingSessionDate
func (m MarketDataUpdateSessionLow) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionLowBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionLowPointer) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionLowPointerBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 20, 16))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionLowBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionLowPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetPrice The session Low price.
func (m MarketDataUpdateSessionLowBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, value)
}

// SetPrice The session Low price.
func (m MarketDataUpdateSessionLowPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionLowBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 20, 16, uint32(value))
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionLowPointerBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 20, 16, uint32(value))
}