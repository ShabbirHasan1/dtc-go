package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataFeedSymbolStatusSize = 12

// {
//     Size     = MarketDataFeedSymbolStatusSize  (12)
//     Type     = MARKET_DATA_FEED_SYMBOL_STATUS  (116)
//     SymbolID = 0
//     Status   = MARKET_DATA_FEED_STATUS_UNSET  (0)
// }
var _MarketDataFeedSymbolStatusDefault = []byte{12, 0, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketDataFeedSymbolStatus struct {
	message.Fixed
}

type MarketDataFeedSymbolStatusBuilder struct {
	message.Fixed
}

type MarketDataFeedSymbolStatusPointer struct {
	message.FixedPointer
}

type MarketDataFeedSymbolStatusPointerBuilder struct {
	message.FixedPointer
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
//     Status   = MARKET_DATA_FEED_STATUS_UNSET  (0)
// }
func (m MarketDataFeedSymbolStatusBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataFeedSymbolStatusDefault)
}

// Clear
// {
//     Size     = MarketDataFeedSymbolStatusSize  (12)
//     Type     = MARKET_DATA_FEED_SYMBOL_STATUS  (116)
//     SymbolID = 0
//     Status   = MARKET_DATA_FEED_STATUS_UNSET  (0)
// }
func (m MarketDataFeedSymbolStatusPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataFeedSymbolStatusDefault)
}

// ToBuilder
func (m MarketDataFeedSymbolStatus) ToBuilder() MarketDataFeedSymbolStatusBuilder {
	return MarketDataFeedSymbolStatusBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataFeedSymbolStatusPointer) ToBuilder() MarketDataFeedSymbolStatusPointerBuilder {
	return MarketDataFeedSymbolStatusPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataFeedSymbolStatusBuilder) Finish() MarketDataFeedSymbolStatus {
	return MarketDataFeedSymbolStatus{m.Fixed}
}

// Finish
func (m *MarketDataFeedSymbolStatusPointerBuilder) Finish() MarketDataFeedSymbolStatusPointer {
	return MarketDataFeedSymbolStatusPointer{m.FixedPointer}
}

// SymbolID
func (m MarketDataFeedSymbolStatus) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDataFeedSymbolStatusBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
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
func (m MarketDataFeedSymbolStatus) Status() MarketDataFeedStatusEnum {
	return MarketDataFeedStatusEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Status
func (m MarketDataFeedSymbolStatusBuilder) Status() MarketDataFeedStatusEnum {
	return MarketDataFeedStatusEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
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
func (m MarketDataFeedSymbolStatusBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID
func (m MarketDataFeedSymbolStatusPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetStatus
func (m MarketDataFeedSymbolStatusBuilder) SetStatus(value MarketDataFeedStatusEnum) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, int32(value))
}

// SetStatus
func (m MarketDataFeedSymbolStatusPointerBuilder) SetStatus(value MarketDataFeedStatusEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
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
