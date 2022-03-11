package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
//     Status = 0
// }
func (m MarketDataFeedStatusPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataFeedStatusDefault)
}

// ToBuilder
func (m MarketDataFeedStatusPointer) ToBuilder() MarketDataFeedStatusPointerBuilder {
	return MarketDataFeedStatusPointerBuilder{m.FixedPointer}
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
func (m MarketDataFeedStatusPointerBuilder) SetStatus(value MarketDataFeedStatusEnum) {
	message.SetInt32Fixed(m.Ptr, 8, 4, int32(value))
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
