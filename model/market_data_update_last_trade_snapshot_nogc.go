package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateLastTradeSnapshotDefault)
}

// ToBuilder
func (m MarketDataUpdateLastTradeSnapshotPointer) ToBuilder() MarketDataUpdateLastTradeSnapshotPointerBuilder {
	return MarketDataUpdateLastTradeSnapshotPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataUpdateLastTradeSnapshotPointerBuilder) Finish() MarketDataUpdateLastTradeSnapshotPointer {
	return MarketDataUpdateLastTradeSnapshotPointer{m.FixedPointer}
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
func (m MarketDataUpdateLastTradeSnapshotPointer) LastTradePrice() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// LastTradePrice The most recent last trade price.
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) LastTradePrice() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
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
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetLastTradePrice The most recent last trade price.
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) SetLastTradePrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, value)
}

// SetLastTradeVolume The quantity/volume of the most recent last trade.
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) SetLastTradeVolume(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetLastTradeDateTime The Date-Time of the last trade.
func (m MarketDataUpdateLastTradeSnapshotPointerBuilder) SetLastTradeDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, float64(value))
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
