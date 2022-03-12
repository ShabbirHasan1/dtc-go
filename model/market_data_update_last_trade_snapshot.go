package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateLastTradeSnapshotSize = 32

// {
//     Size              = MarketDataUpdateLastTradeSnapshotSize  (32)
//     Type              = MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT  (134)
//     SymbolID          = 0
//     LastTradePrice    = 0
//     LastTradeVolume   = 0
//     LastTradeDateTime = 0
// }
var _MarketDataUpdateLastTradeSnapshotDefault = []byte{32, 0, 134, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateLastTradeSnapshot Sent by the Server to the Client to update the last trade price, volume
// and date-time fields under conditions when there is not a trade.
//
// This message type does not signify a trade has occurred. It should never
// be interpreted by the Client in that way.
type MarketDataUpdateLastTradeSnapshot struct {
	message.Fixed
}

// MarketDataUpdateLastTradeSnapshotBuilder Sent by the Server to the Client to update the last trade price, volume
// and date-time fields under conditions when there is not a trade.
//
// This message type does not signify a trade has occurred. It should never
// be interpreted by the Client in that way.
type MarketDataUpdateLastTradeSnapshotBuilder struct {
	message.Fixed
}

// MarketDataUpdateLastTradeSnapshotPointer Sent by the Server to the Client to update the last trade price, volume
// and date-time fields under conditions when there is not a trade.
//
// This message type does not signify a trade has occurred. It should never
// be interpreted by the Client in that way.
type MarketDataUpdateLastTradeSnapshotPointer struct {
	message.FixedPointer
}

// MarketDataUpdateLastTradeSnapshotPointerBuilder Sent by the Server to the Client to update the last trade price, volume
// and date-time fields under conditions when there is not a trade.
//
// This message type does not signify a trade has occurred. It should never
// be interpreted by the Client in that way.
type MarketDataUpdateLastTradeSnapshotPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataUpdateLastTradeSnapshotFrom(b []byte) MarketDataUpdateLastTradeSnapshot {
	return MarketDataUpdateLastTradeSnapshot{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateLastTradeSnapshot(b []byte) MarketDataUpdateLastTradeSnapshot {
	return MarketDataUpdateLastTradeSnapshot{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateLastTradeSnapshot() MarketDataUpdateLastTradeSnapshotBuilder {
	a := MarketDataUpdateLastTradeSnapshotBuilder{message.NewFixed(32)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateLastTradeSnapshotDefault)
	return a
}

func AllocMarketDataUpdateLastTradeSnapshot() MarketDataUpdateLastTradeSnapshotPointerBuilder {
	a := MarketDataUpdateLastTradeSnapshotPointerBuilder{message.AllocFixed(32)}
	a.Ptr.SetBytes(0, _MarketDataUpdateLastTradeSnapshotDefault)
	return a
}

func AllocMarketDataUpdateLastTradeSnapshotFrom(b []byte) MarketDataUpdateLastTradeSnapshotPointer {
	return MarketDataUpdateLastTradeSnapshotPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size              = MarketDataUpdateLastTradeSnapshotSize  (32)
//     Type              = MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT  (134)
//     SymbolID          = 0
//     LastTradePrice    = 0
//     LastTradeVolume   = 0
//     LastTradeDateTime = 0
// }
func (m MarketDataUpdateLastTradeSnapshotBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateLastTradeSnapshotDefault)
}

// Clear
// {
//     Size              = MarketDataUpdateLastTradeSnapshotSize  (32)
//     Type              = MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT  (134)
//     SymbolID          = 0
//     LastTradePrice    = 0
//     LastTradeVolume   = 0
//     LastTradeDateTime = 0
// }
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateLastTradeSnapshotDefault)
}

// ToBuilder
func (m MarketDataUpdateLastTradeSnapshot) ToBuilder() MarketDataUpdateLastTradeSnapshotBuilder {
	return MarketDataUpdateLastTradeSnapshotBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateLastTradeSnapshotPointer) ToBuilder() MarketDataUpdateLastTradeSnapshotPointerBuilder {
	return MarketDataUpdateLastTradeSnapshotPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateLastTradeSnapshotBuilder) Finish() MarketDataUpdateLastTradeSnapshot {
	return MarketDataUpdateLastTradeSnapshot{m.Fixed}
}

// Finish
func (m *MarketDataUpdateLastTradeSnapshotPointerBuilder) Finish() MarketDataUpdateLastTradeSnapshotPointer {
	return MarketDataUpdateLastTradeSnapshotPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateLastTradeSnapshot) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateLastTradeSnapshotBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateLastTradeSnapshotPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// LastTradePrice The most recent last trade price.
func (m MarketDataUpdateLastTradeSnapshot) LastTradePrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// LastTradePrice The most recent last trade price.
func (m MarketDataUpdateLastTradeSnapshotBuilder) LastTradePrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// LastTradePrice The most recent last trade price.
func (m MarketDataUpdateLastTradeSnapshotPointer) LastTradePrice() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// LastTradePrice The most recent last trade price.
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) LastTradePrice() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// LastTradeVolume The quantity/volume of the most recent last trade.
func (m MarketDataUpdateLastTradeSnapshot) LastTradeVolume() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// LastTradeVolume The quantity/volume of the most recent last trade.
func (m MarketDataUpdateLastTradeSnapshotBuilder) LastTradeVolume() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// LastTradeVolume The quantity/volume of the most recent last trade.
func (m MarketDataUpdateLastTradeSnapshotPointer) LastTradeVolume() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// LastTradeVolume The quantity/volume of the most recent last trade.
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) LastTradeVolume() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// LastTradeDateTime The Date-Time of the last trade.
func (m MarketDataUpdateLastTradeSnapshot) LastTradeDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 32, 24))
}

// LastTradeDateTime The Date-Time of the last trade.
func (m MarketDataUpdateLastTradeSnapshotBuilder) LastTradeDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 32, 24))
}

// LastTradeDateTime The Date-Time of the last trade.
func (m MarketDataUpdateLastTradeSnapshotPointer) LastTradeDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 32, 24))
}

// LastTradeDateTime The Date-Time of the last trade.
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) LastTradeDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 32, 24))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateLastTradeSnapshotBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetLastTradePrice The most recent last trade price.
func (m MarketDataUpdateLastTradeSnapshotBuilder) SetLastTradePrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, value)
}

// SetLastTradePrice The most recent last trade price.
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) SetLastTradePrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, value)
}

// SetLastTradeVolume The quantity/volume of the most recent last trade.
func (m MarketDataUpdateLastTradeSnapshotBuilder) SetLastTradeVolume(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetLastTradeVolume The quantity/volume of the most recent last trade.
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) SetLastTradeVolume(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetLastTradeDateTime The Date-Time of the last trade.
func (m MarketDataUpdateLastTradeSnapshotBuilder) SetLastTradeDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 32, 24, float64(value))
}

// SetLastTradeDateTime The Date-Time of the last trade.
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) SetLastTradeDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, float64(value))
}

// Copy
func (m MarketDataUpdateLastTradeSnapshot) Copy(to MarketDataUpdateLastTradeSnapshotBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
}

// Copy
func (m MarketDataUpdateLastTradeSnapshotBuilder) Copy(to MarketDataUpdateLastTradeSnapshotBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
}

// CopyPointer
func (m MarketDataUpdateLastTradeSnapshot) CopyPointer(to MarketDataUpdateLastTradeSnapshotPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
}

// CopyPointer
func (m MarketDataUpdateLastTradeSnapshotBuilder) CopyPointer(to MarketDataUpdateLastTradeSnapshotPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
}

// Copy
func (m MarketDataUpdateLastTradeSnapshotPointer) Copy(to MarketDataUpdateLastTradeSnapshotBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
}

// Copy
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) Copy(to MarketDataUpdateLastTradeSnapshotBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
}

// CopyPointer
func (m MarketDataUpdateLastTradeSnapshotPointer) CopyPointer(to MarketDataUpdateLastTradeSnapshotPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
}

// CopyPointer
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) CopyPointer(to MarketDataUpdateLastTradeSnapshotPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
}
