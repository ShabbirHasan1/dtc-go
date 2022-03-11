package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                                     = MarketDataRequestFixedSize  (96)
//     Type                                     = MARKET_DATA_REQUEST  (101)
//     RequestAction                            = SUBSCRIBE  (1)
//     SymbolID                                 = 0
//     Symbol                                   = ""
//     Exchange                                 = ""
//     IntervalForSnapshotUpdatesInMilliseconds = 0
// }
var _MarketDataRequestFixedDefault = []byte{96, 0, 101, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataRequestFixedSize = 96

// MarketDataRequestFixed The MarketDataRequest message will subscribe to market data for a particular
// Symbol or request a market data snapshot.
//
// The Server can also send market depth data in response to this message
// and not require a MarketDepthRequest.
type MarketDataRequestFixed struct {
	message.Fixed
}

// MarketDataRequestFixedBuilder The MarketDataRequest message will subscribe to market data for a particular
// Symbol or request a market data snapshot.
//
// The Server can also send market depth data in response to this message
// and not require a MarketDepthRequest.
type MarketDataRequestFixedBuilder struct {
	message.Fixed
}

func NewMarketDataRequestFixedFrom(b []byte) MarketDataRequestFixed {
	return MarketDataRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataRequestFixed(b []byte) MarketDataRequestFixed {
	return MarketDataRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewMarketDataRequestFixed() MarketDataRequestFixedBuilder {
	a := MarketDataRequestFixedBuilder{message.NewFixed(96)}
	a.Unsafe().SetBytes(0, _MarketDataRequestFixedDefault)
	return a
}

// Clear
// {
//     Size                                     = MarketDataRequestFixedSize  (96)
//     Type                                     = MARKET_DATA_REQUEST  (101)
//     RequestAction                            = SUBSCRIBE  (1)
//     SymbolID                                 = 0
//     Symbol                                   = ""
//     Exchange                                 = ""
//     IntervalForSnapshotUpdatesInMilliseconds = 0
// }
func (m MarketDataRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataRequestFixedDefault)
}

// ToBuilder
func (m MarketDataRequestFixed) ToBuilder() MarketDataRequestFixedBuilder {
	return MarketDataRequestFixedBuilder{m.Fixed}
}

// Finish
func (m MarketDataRequestFixedBuilder) Finish() MarketDataRequestFixed {
	return MarketDataRequestFixed{m.Fixed}
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The Server will respond with an initial MarketDataSnapshot
// message and then provide MARKET_DATA_UPDATE_* updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
//
// To request only a MarketDataSnapshot message, set this to SNAPSHOT.
func (m MarketDataRequestFixed) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Unsafe(), 8, 4))
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The Server will respond with an initial MarketDataSnapshot
// message and then provide MARKET_DATA_UPDATE_* updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
//
// To request only a MarketDataSnapshot message, set this to SNAPSHOT.
func (m MarketDataRequestFixedBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Unsafe(), 8, 4))
}

// SymbolID This is the identifier which will be used in all of the market data response
// messages. This identifier is used so that the Symbol and Exchange do not
// have to be passed back in response messages from the Server. This makes
// the market data feed bandwidth efficient.
//
// If the Server receives a MarketDataRequest for a Symbol and Exchange to
// subscribe to data for, that is currently subscribed to and this SymbolID
// is different, then the Server needs to reject it.
func (m MarketDataRequestFixed) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// SymbolID This is the identifier which will be used in all of the market data response
// messages. This identifier is used so that the Symbol and Exchange do not
// have to be passed back in response messages from the Server. This makes
// the market data feed bandwidth efficient.
//
// If the Server receives a MarketDataRequest for a Symbol and Exchange to
// subscribe to data for, that is currently subscribed to and this SymbolID
// is different, then the Server needs to reject it.
func (m MarketDataRequestFixedBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// Symbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 76, 12)
}

// Symbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 76, 12)
}

// Exchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 92, 76)
}

// Exchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 92, 76)
}

// IntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestFixed) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 96, 92)
}

// IntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestFixedBuilder) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 96, 92)
}

// SetRequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The Server will respond with an initial MarketDataSnapshot
// message and then provide MARKET_DATA_UPDATE_* updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
//
// To request only a MarketDataSnapshot message, set this to SNAPSHOT.
func (m MarketDataRequestFixedBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, int32(value))
}

// SetSymbolID This is the identifier which will be used in all of the market data response
// messages. This identifier is used so that the Symbol and Exchange do not
// have to be passed back in response messages from the Server. This makes
// the market data feed bandwidth efficient.
//
// If the Server receives a MarketDataRequest for a Symbol and Exchange to
// subscribe to data for, that is currently subscribed to and this SymbolID
// is different, then the Server needs to reject it.
func (m MarketDataRequestFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 12, 8, value)
}

// SetSymbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 76, 12, value)
}

// SetExchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 92, 76, value)
}

// SetIntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestFixedBuilder) SetIntervalForSnapshotUpdatesInMilliseconds(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 96, 92, value)
}

// Copy
func (m MarketDataRequestFixed) Copy(to MarketDataRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// Copy
func (m MarketDataRequestFixedBuilder) Copy(to MarketDataRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyPointer
func (m MarketDataRequestFixed) CopyPointer(to MarketDataRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyPointer
func (m MarketDataRequestFixedBuilder) CopyPointer(to MarketDataRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyToPointer
func (m MarketDataRequestFixed) CopyToPointer(to MarketDataRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyToPointer
func (m MarketDataRequestFixedBuilder) CopyToPointer(to MarketDataRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyTo
func (m MarketDataRequestFixed) CopyTo(to MarketDataRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyTo
func (m MarketDataRequestFixedBuilder) CopyTo(to MarketDataRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}
