package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDataRequestFixedPointer The MarketDataRequest message will subscribe to market data for a particular
// Symbol or request a market data snapshot.
//
// The Server can also send market depth data in response to this message
// and not require a MarketDepthRequest.
type MarketDataRequestFixedPointer struct {
	message.FixedPointer
}

// MarketDataRequestFixedPointerBuilder The MarketDataRequest message will subscribe to market data for a particular
// Symbol or request a market data snapshot.
//
// The Server can also send market depth data in response to this message
// and not require a MarketDepthRequest.
type MarketDataRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDataRequestFixed() MarketDataRequestFixedPointerBuilder {
	a := MarketDataRequestFixedPointerBuilder{message.AllocFixed(96)}
	a.Ptr.SetBytes(0, _MarketDataRequestFixedDefault)
	return a
}

func AllocMarketDataRequestFixedFrom(b []byte) MarketDataRequestFixedPointer {
	return MarketDataRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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
func (m MarketDataRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataRequestFixedDefault)
}

// ToBuilder
func (m MarketDataRequestFixedPointer) ToBuilder() MarketDataRequestFixedPointerBuilder {
	return MarketDataRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDataRequestFixedPointerBuilder) Finish() MarketDataRequestFixedPointer {
	return MarketDataRequestFixedPointer{m.FixedPointer}
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The Server will respond with an initial MarketDataSnapshot
// message and then provide MARKET_DATA_UPDATE_* updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
//
// To request only a MarketDataSnapshot message, set this to SNAPSHOT.
func (m MarketDataRequestFixedPointer) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Ptr, 8, 4))
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The Server will respond with an initial MarketDataSnapshot
// message and then provide MARKET_DATA_UPDATE_* updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
//
// To request only a MarketDataSnapshot message, set this to SNAPSHOT.
func (m MarketDataRequestFixedPointerBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Ptr, 8, 4))
}

// SymbolID This is the identifier which will be used in all of the market data response
// messages. This identifier is used so that the Symbol and Exchange do not
// have to be passed back in response messages from the Server. This makes
// the market data feed bandwidth efficient.
//
// If the Server receives a MarketDataRequest for a Symbol and Exchange to
// subscribe to data for, that is currently subscribed to and this SymbolID
// is different, then the Server needs to reject it.
func (m MarketDataRequestFixedPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// SymbolID This is the identifier which will be used in all of the market data response
// messages. This identifier is used so that the Symbol and Exchange do not
// have to be passed back in response messages from the Server. This makes
// the market data feed bandwidth efficient.
//
// If the Server receives a MarketDataRequest for a Symbol and Exchange to
// subscribe to data for, that is currently subscribed to and this SymbolID
// is different, then the Server needs to reject it.
func (m MarketDataRequestFixedPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// Symbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequestFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 76, 12)
}

// Symbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequestFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 76, 12)
}

// Exchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 92, 76)
}

// Exchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 92, 76)
}

// IntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestFixedPointer) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32Fixed(m.Ptr, 96, 92)
}

// IntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestFixedPointerBuilder) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32Fixed(m.Ptr, 96, 92)
}

// SetRequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The Server will respond with an initial MarketDataSnapshot
// message and then provide MARKET_DATA_UPDATE_* updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
//
// To request only a MarketDataSnapshot message, set this to SNAPSHOT.
func (m MarketDataRequestFixedPointerBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32Fixed(m.Ptr, 8, 4, int32(value))
}

// SetSymbolID This is the identifier which will be used in all of the market data response
// messages. This identifier is used so that the Symbol and Exchange do not
// have to be passed back in response messages from the Server. This makes
// the market data feed bandwidth efficient.
//
// If the Server receives a MarketDataRequest for a Symbol and Exchange to
// subscribe to data for, that is currently subscribed to and this SymbolID
// is different, then the Server needs to reject it.
func (m MarketDataRequestFixedPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 12, 8, value)
}

// SetSymbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 76, 12, value)
}

// SetExchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 92, 76, value)
}

// SetIntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestFixedPointerBuilder) SetIntervalForSnapshotUpdatesInMilliseconds(value uint32) {
	message.SetUint32Fixed(m.Ptr, 96, 92, value)
}

// Copy
func (m MarketDataRequestFixedPointer) Copy(to MarketDataRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// Copy
func (m MarketDataRequestFixedPointerBuilder) Copy(to MarketDataRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyPointer
func (m MarketDataRequestFixedPointer) CopyPointer(to MarketDataRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyPointer
func (m MarketDataRequestFixedPointerBuilder) CopyPointer(to MarketDataRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyTo
func (m MarketDataRequestFixedPointer) CopyTo(to MarketDataRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyTo
func (m MarketDataRequestFixedPointerBuilder) CopyTo(to MarketDataRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyToPointer
func (m MarketDataRequestFixedPointer) CopyToPointer(to MarketDataRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyToPointer
func (m MarketDataRequestFixedPointerBuilder) CopyToPointer(to MarketDataRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}
