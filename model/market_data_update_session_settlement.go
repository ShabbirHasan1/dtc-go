package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateSessionSettlementSize = 24

// {
//     Size     = MarketDataUpdateSessionSettlementSize  (24)
//     Type     = MARKET_DATA_UPDATE_SESSION_SETTLEMENT  (119)
//     SymbolID = 0
//     Price    = 0
//     DateTime = 0
// }
var _MarketDataUpdateSessionSettlementDefault = []byte{24, 0, 119, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateSessionSettlement Sent by the Server to the Client to update the session settlement price
// when the session settlement price changes.
type MarketDataUpdateSessionSettlement struct {
	message.Fixed
}

// MarketDataUpdateSessionSettlementBuilder Sent by the Server to the Client to update the session settlement price
// when the session settlement price changes.
type MarketDataUpdateSessionSettlementBuilder struct {
	message.Fixed
}

// MarketDataUpdateSessionSettlementPointer Sent by the Server to the Client to update the session settlement price
// when the session settlement price changes.
type MarketDataUpdateSessionSettlementPointer struct {
	message.FixedPointer
}

// MarketDataUpdateSessionSettlementPointerBuilder Sent by the Server to the Client to update the session settlement price
// when the session settlement price changes.
type MarketDataUpdateSessionSettlementPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataUpdateSessionSettlementFrom(b []byte) MarketDataUpdateSessionSettlement {
	return MarketDataUpdateSessionSettlement{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateSessionSettlement(b []byte) MarketDataUpdateSessionSettlement {
	return MarketDataUpdateSessionSettlement{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateSessionSettlement() MarketDataUpdateSessionSettlementBuilder {
	a := MarketDataUpdateSessionSettlementBuilder{message.NewFixed(24)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateSessionSettlementDefault)
	return a
}

func AllocMarketDataUpdateSessionSettlement() MarketDataUpdateSessionSettlementPointerBuilder {
	a := MarketDataUpdateSessionSettlementPointerBuilder{message.AllocFixed(24)}
	a.Ptr.SetBytes(0, _MarketDataUpdateSessionSettlementDefault)
	return a
}

func AllocMarketDataUpdateSessionSettlementFrom(b []byte) MarketDataUpdateSessionSettlementPointer {
	return MarketDataUpdateSessionSettlementPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size     = MarketDataUpdateSessionSettlementSize  (24)
//     Type     = MARKET_DATA_UPDATE_SESSION_SETTLEMENT  (119)
//     SymbolID = 0
//     Price    = 0
//     DateTime = 0
// }
func (m MarketDataUpdateSessionSettlementBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateSessionSettlementDefault)
}

// Clear
// {
//     Size     = MarketDataUpdateSessionSettlementSize  (24)
//     Type     = MARKET_DATA_UPDATE_SESSION_SETTLEMENT  (119)
//     SymbolID = 0
//     Price    = 0
//     DateTime = 0
// }
func (m MarketDataUpdateSessionSettlementPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateSessionSettlementDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionSettlement) ToBuilder() MarketDataUpdateSessionSettlementBuilder {
	return MarketDataUpdateSessionSettlementBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateSessionSettlementPointer) ToBuilder() MarketDataUpdateSessionSettlementPointerBuilder {
	return MarketDataUpdateSessionSettlementPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateSessionSettlementBuilder) Finish() MarketDataUpdateSessionSettlement {
	return MarketDataUpdateSessionSettlement{m.Fixed}
}

// Finish
func (m *MarketDataUpdateSessionSettlementPointerBuilder) Finish() MarketDataUpdateSessionSettlementPointer {
	return MarketDataUpdateSessionSettlementPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionSettlement) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionSettlementBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionSettlementPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionSettlementPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Price The settlement price.
func (m MarketDataUpdateSessionSettlement) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// Price The settlement price.
func (m MarketDataUpdateSessionSettlementBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// Price The settlement price.
func (m MarketDataUpdateSessionSettlementPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// Price The settlement price.
func (m MarketDataUpdateSessionSettlementPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// DateTime That trading date the settlement price is for. The time component is not
// normally considered relevant in this case.
func (m MarketDataUpdateSessionSettlement) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// DateTime That trading date the settlement price is for. The time component is not
// normally considered relevant in this case.
func (m MarketDataUpdateSessionSettlementBuilder) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// DateTime That trading date the settlement price is for. The time component is not
// normally considered relevant in this case.
func (m MarketDataUpdateSessionSettlementPointer) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 20, 16))
}

// DateTime That trading date the settlement price is for. The time component is not
// normally considered relevant in this case.
func (m MarketDataUpdateSessionSettlementPointerBuilder) DateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 20, 16))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionSettlementBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionSettlementPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetPrice The settlement price.
func (m MarketDataUpdateSessionSettlementBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, value)
}

// SetPrice The settlement price.
func (m MarketDataUpdateSessionSettlementPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, value)
}

// SetDateTime That trading date the settlement price is for. The time component is not
// normally considered relevant in this case.
func (m MarketDataUpdateSessionSettlementBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 20, 16, uint32(value))
}

// SetDateTime That trading date the settlement price is for. The time component is not
// normally considered relevant in this case.
func (m MarketDataUpdateSessionSettlementPointerBuilder) SetDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 20, 16, uint32(value))
}

// Copy
func (m MarketDataUpdateSessionSettlement) Copy(to MarketDataUpdateSessionSettlementBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateSessionSettlementBuilder) Copy(to MarketDataUpdateSessionSettlementBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateSessionSettlement) CopyPointer(to MarketDataUpdateSessionSettlementPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateSessionSettlementBuilder) CopyPointer(to MarketDataUpdateSessionSettlementPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateSessionSettlementPointer) Copy(to MarketDataUpdateSessionSettlementBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
}

// Copy
func (m MarketDataUpdateSessionSettlementPointerBuilder) Copy(to MarketDataUpdateSessionSettlementBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateSessionSettlementPointer) CopyPointer(to MarketDataUpdateSessionSettlementPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
}

// CopyPointer
func (m MarketDataUpdateSessionSettlementPointerBuilder) CopyPointer(to MarketDataUpdateSessionSettlementPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
}
