package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size     = MarketDataFeedSymbolStatusSize  (12)
//     Type     = MARKET_DATA_FEED_SYMBOL_STATUS  (116)
//     SymbolID = 0
//     Status   = 0
// }
var _MarketDataFeedSymbolStatusDefault = []byte{12, 0, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataFeedSymbolStatusSize = 12

type MarketDataFeedSymbolStatus struct {
	message.Fixed
}

type MarketDataFeedSymbolStatusBuilder struct {
	message.Fixed
}

func NewMarketDataFeedSymbolStatusFrom(b []byte) MarketDataFeedSymbolStatus {
	return MarketDataFeedSymbolStatus{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataFeedSymbolStatus(b []byte) MarketDataFeedSymbolStatus {
	return MarketDataFeedSymbolStatus{Fixed: message.WrapFixed(b)}
}

func NewMarketDataFeedSymbolStatus() MarketDataFeedSymbolStatusBuilder {
	a := MarketDataFeedSymbolStatusBuilder{message.NewFixed(12)}
	a.Unsafe().SetBytes(0, _MarketDataFeedSymbolStatusDefault)
	return a
}

// Clear
// {
//     Size     = MarketDataFeedSymbolStatusSize  (12)
//     Type     = MARKET_DATA_FEED_SYMBOL_STATUS  (116)
//     SymbolID = 0
//     Status   = 0
// }
func (m MarketDataFeedSymbolStatusBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataFeedSymbolStatusDefault)
}

// ToBuilder
func (m MarketDataFeedSymbolStatus) ToBuilder() MarketDataFeedSymbolStatusBuilder {
	return MarketDataFeedSymbolStatusBuilder{m.Fixed}
}

// Finish
func (m MarketDataFeedSymbolStatusBuilder) Finish() MarketDataFeedSymbolStatus {
	return MarketDataFeedSymbolStatus{m.Fixed}
}

// SymbolID
func (m MarketDataFeedSymbolStatus) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDataFeedSymbolStatusBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Status
func (m MarketDataFeedSymbolStatus) Status() MarketDataFeedStatusEnum {
	return MarketDataFeedStatusEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Status
func (m MarketDataFeedSymbolStatusBuilder) Status() MarketDataFeedStatusEnum {
	return MarketDataFeedStatusEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// SetSymbolID
func (m MarketDataFeedSymbolStatusBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetStatus
func (m MarketDataFeedSymbolStatusBuilder) SetStatus(value MarketDataFeedStatusEnum) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, int32(value))
}

// Copy
func (m MarketDataFeedSymbolStatus) Copy(to MarketDataFeedSymbolStatusBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// Copy
func (m MarketDataFeedSymbolStatusBuilder) Copy(to MarketDataFeedSymbolStatusBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// CopyPointer
func (m MarketDataFeedSymbolStatus) CopyPointer(to MarketDataFeedSymbolStatusPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// CopyPointer
func (m MarketDataFeedSymbolStatusBuilder) CopyPointer(to MarketDataFeedSymbolStatusPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}
