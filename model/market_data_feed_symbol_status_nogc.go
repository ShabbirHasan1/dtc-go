package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketDataFeedSymbolStatusPointer struct {
	message.FixedPointer
}

type MarketDataFeedSymbolStatusPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataFeedSymbolStatus() MarketDataFeedSymbolStatusPointerBuilder {
	a := MarketDataFeedSymbolStatusPointerBuilder{message.AllocFixed(12)}
	a.Ptr.SetBytes(0, _MarketDataFeedSymbolStatusDefault)
	return a
}

func AllocMarketDataFeedSymbolStatusFrom(b []byte) MarketDataFeedSymbolStatusPointer {
	return MarketDataFeedSymbolStatusPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size     = MarketDataFeedSymbolStatusSize  (12)
//     Type     = MARKET_DATA_FEED_SYMBOL_STATUS  (116)
//     SymbolID = 0
//     Status   = 0
// }
func (m MarketDataFeedSymbolStatusPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataFeedSymbolStatusDefault)
}

// ToBuilder
func (m MarketDataFeedSymbolStatusPointer) ToBuilder() MarketDataFeedSymbolStatusPointerBuilder {
	return MarketDataFeedSymbolStatusPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataFeedSymbolStatusPointerBuilder) Finish() MarketDataFeedSymbolStatusPointer {
	return MarketDataFeedSymbolStatusPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDataFeedSymbolStatusPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m MarketDataFeedSymbolStatusPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Status
func (m MarketDataFeedSymbolStatusPointer) Status() MarketDataFeedStatusEnum {
	return MarketDataFeedStatusEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// Status
func (m MarketDataFeedSymbolStatusPointerBuilder) Status() MarketDataFeedStatusEnum {
	return MarketDataFeedStatusEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// SetSymbolID
func (m MarketDataFeedSymbolStatusPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetStatus
func (m MarketDataFeedSymbolStatusPointerBuilder) SetStatus(value MarketDataFeedStatusEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
}

// Copy
func (m MarketDataFeedSymbolStatusPointer) Copy(to MarketDataFeedSymbolStatusBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// Copy
func (m MarketDataFeedSymbolStatusPointerBuilder) Copy(to MarketDataFeedSymbolStatusBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// CopyPointer
func (m MarketDataFeedSymbolStatusPointer) CopyPointer(to MarketDataFeedSymbolStatusPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// CopyPointer
func (m MarketDataFeedSymbolStatusPointerBuilder) CopyPointer(to MarketDataFeedSymbolStatusPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}
