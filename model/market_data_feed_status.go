package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataFeedStatusSize = 8

// {
//     Size   = MarketDataFeedStatusSize  (8)
//     Type   = MARKET_DATA_FEED_STATUS  (100)
//     Status = MARKET_DATA_FEED_STATUS_UNSET  (0)
// }
var _MarketDataFeedStatusDefault = []byte{8, 0, 100, 0, 0, 0, 0, 0}

// MarketDataFeedStatus The s_MarketDataFeed_STATUS message is an optional message sent by the
// Server to indicate the overall status of the market data feed. This status
// applies to all symbols that have been subscribed to for market data.
type MarketDataFeedStatus struct {
	message.Fixed
}

// MarketDataFeedStatusBuilder The s_MarketDataFeed_STATUS message is an optional message sent by the
// Server to indicate the overall status of the market data feed. This status
// applies to all symbols that have been subscribed to for market data.
type MarketDataFeedStatusBuilder struct {
	message.Fixed
}

// MarketDataFeedStatusPointer The s_MarketDataFeed_STATUS message is an optional message sent by the
// Server to indicate the overall status of the market data feed. This status
// applies to all symbols that have been subscribed to for market data.
type MarketDataFeedStatusPointer struct {
	message.FixedPointer
}

// MarketDataFeedStatusPointerBuilder The s_MarketDataFeed_STATUS message is an optional message sent by the
// Server to indicate the overall status of the market data feed. This status
// applies to all symbols that have been subscribed to for market data.
type MarketDataFeedStatusPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataFeedStatusFrom(b []byte) MarketDataFeedStatus {
	return MarketDataFeedStatus{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataFeedStatus(b []byte) MarketDataFeedStatus {
	return MarketDataFeedStatus{Fixed: message.WrapFixed(b)}
}

func NewMarketDataFeedStatus() MarketDataFeedStatusBuilder {
	a := MarketDataFeedStatusBuilder{message.NewFixed(8)}
	a.Unsafe().SetBytes(0, _MarketDataFeedStatusDefault)
	return a
}

func AllocMarketDataFeedStatus() MarketDataFeedStatusPointerBuilder {
	a := MarketDataFeedStatusPointerBuilder{message.AllocFixed(8)}
	a.Ptr.SetBytes(0, _MarketDataFeedStatusDefault)
	return a
}

func AllocMarketDataFeedStatusFrom(b []byte) MarketDataFeedStatusPointer {
	return MarketDataFeedStatusPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size   = MarketDataFeedStatusSize  (8)
//     Type   = MARKET_DATA_FEED_STATUS  (100)
//     Status = MARKET_DATA_FEED_STATUS_UNSET  (0)
// }
func (m MarketDataFeedStatusBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataFeedStatusDefault)
}

// Clear
// {
//     Size   = MarketDataFeedStatusSize  (8)
//     Type   = MARKET_DATA_FEED_STATUS  (100)
//     Status = MARKET_DATA_FEED_STATUS_UNSET  (0)
// }
func (m MarketDataFeedStatusPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataFeedStatusDefault)
}

// ToBuilder
func (m MarketDataFeedStatus) ToBuilder() MarketDataFeedStatusBuilder {
	return MarketDataFeedStatusBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataFeedStatusPointer) ToBuilder() MarketDataFeedStatusPointerBuilder {
	return MarketDataFeedStatusPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataFeedStatusBuilder) Finish() MarketDataFeedStatus {
	return MarketDataFeedStatus{m.Fixed}
}

// Finish
func (m *MarketDataFeedStatusPointerBuilder) Finish() MarketDataFeedStatusPointer {
	return MarketDataFeedStatusPointer{m.FixedPointer}
}

// Status This can be set to s_MarketDataFeed_UNAVAILABLE, to indicate the market
// data feed is presently not available. Or it can be set to s_MarketDataFeed_AVAILABLE,
// to indicate the market data feed has been restored.
//
// Upon a connection to the server, s_MarketDataFeed_AVAILABLE is assumed
// to be the status. It is not until there has been expressly given s_MarketDataFeed_UNAVAILABLE,
// will the data feed be considered lost.
func (m MarketDataFeedStatus) Status() MarketDataFeedStatusEnum {
	return MarketDataFeedStatusEnum(message.Int32Fixed(m.Unsafe(), 8, 4))
}

// Status This can be set to s_MarketDataFeed_UNAVAILABLE, to indicate the market
// data feed is presently not available. Or it can be set to s_MarketDataFeed_AVAILABLE,
// to indicate the market data feed has been restored.
//
// Upon a connection to the server, s_MarketDataFeed_AVAILABLE is assumed
// to be the status. It is not until there has been expressly given s_MarketDataFeed_UNAVAILABLE,
// will the data feed be considered lost.
func (m MarketDataFeedStatusBuilder) Status() MarketDataFeedStatusEnum {
	return MarketDataFeedStatusEnum(message.Int32Fixed(m.Unsafe(), 8, 4))
}

// Status This can be set to s_MarketDataFeed_UNAVAILABLE, to indicate the market
// data feed is presently not available. Or it can be set to s_MarketDataFeed_AVAILABLE,
// to indicate the market data feed has been restored.
//
// Upon a connection to the server, s_MarketDataFeed_AVAILABLE is assumed
// to be the status. It is not until there has been expressly given s_MarketDataFeed_UNAVAILABLE,
// will the data feed be considered lost.
func (m MarketDataFeedStatusPointer) Status() MarketDataFeedStatusEnum {
	return MarketDataFeedStatusEnum(message.Int32Fixed(m.Ptr, 8, 4))
}

// Status This can be set to s_MarketDataFeed_UNAVAILABLE, to indicate the market
// data feed is presently not available. Or it can be set to s_MarketDataFeed_AVAILABLE,
// to indicate the market data feed has been restored.
//
// Upon a connection to the server, s_MarketDataFeed_AVAILABLE is assumed
// to be the status. It is not until there has been expressly given s_MarketDataFeed_UNAVAILABLE,
// will the data feed be considered lost.
func (m MarketDataFeedStatusPointerBuilder) Status() MarketDataFeedStatusEnum {
	return MarketDataFeedStatusEnum(message.Int32Fixed(m.Ptr, 8, 4))
}

// SetStatus This can be set to s_MarketDataFeed_UNAVAILABLE, to indicate the market
// data feed is presently not available. Or it can be set to s_MarketDataFeed_AVAILABLE,
// to indicate the market data feed has been restored.
//
// Upon a connection to the server, s_MarketDataFeed_AVAILABLE is assumed
// to be the status. It is not until there has been expressly given s_MarketDataFeed_UNAVAILABLE,
// will the data feed be considered lost.
func (m MarketDataFeedStatusBuilder) SetStatus(value MarketDataFeedStatusEnum) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, int32(value))
}

// SetStatus This can be set to s_MarketDataFeed_UNAVAILABLE, to indicate the market
// data feed is presently not available. Or it can be set to s_MarketDataFeed_AVAILABLE,
// to indicate the market data feed has been restored.
//
// Upon a connection to the server, s_MarketDataFeed_AVAILABLE is assumed
// to be the status. It is not until there has been expressly given s_MarketDataFeed_UNAVAILABLE,
// will the data feed be considered lost.
func (m MarketDataFeedStatusPointerBuilder) SetStatus(value MarketDataFeedStatusEnum) {
	message.SetInt32Fixed(m.Ptr, 8, 4, int32(value))
}

// Copy
func (m MarketDataFeedStatus) Copy(to MarketDataFeedStatusBuilder) {
	to.SetStatus(m.Status())
}

// Copy
func (m MarketDataFeedStatusBuilder) Copy(to MarketDataFeedStatusBuilder) {
	to.SetStatus(m.Status())
}

// CopyPointer
func (m MarketDataFeedStatus) CopyPointer(to MarketDataFeedStatusPointerBuilder) {
	to.SetStatus(m.Status())
}

// CopyPointer
func (m MarketDataFeedStatusBuilder) CopyPointer(to MarketDataFeedStatusPointerBuilder) {
	to.SetStatus(m.Status())
}

// Copy
func (m MarketDataFeedStatusPointer) Copy(to MarketDataFeedStatusBuilder) {
	to.SetStatus(m.Status())
}

// Copy
func (m MarketDataFeedStatusPointerBuilder) Copy(to MarketDataFeedStatusBuilder) {
	to.SetStatus(m.Status())
}

// CopyPointer
func (m MarketDataFeedStatusPointer) CopyPointer(to MarketDataFeedStatusPointerBuilder) {
	to.SetStatus(m.Status())
}

// CopyPointer
func (m MarketDataFeedStatusPointerBuilder) CopyPointer(to MarketDataFeedStatusPointerBuilder) {
	to.SetStatus(m.Status())
}
