package model

import (
	"github.com/moontrade/dtc-go/message"
)

// MarketDataRequestPointer The MarketDataRequest message will subscribe to market data for a particular
// Symbol or request a market data snapshot.
//
// The Server can also send market depth data in response to this message
// and not require a MarketDepthRequest.
type MarketDataRequestPointer struct {
	message.VLSPointer
}

// MarketDataRequestPointerBuilder The MarketDataRequest message will subscribe to market data for a particular
// Symbol or request a market data snapshot.
//
// The Server can also send market depth data in response to this message
// and not require a MarketDepthRequest.
type MarketDataRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocMarketDataRequest() MarketDataRequestPointerBuilder {
	a := MarketDataRequestPointerBuilder{message.AllocVLSBuilder(28)}
	a.Ptr.SetBytes(0, _MarketDataRequestDefault)
	return a
}

func AllocMarketDataRequestFrom(b []byte) MarketDataRequestPointer {
	return MarketDataRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                                     = MarketDataRequestSize  (28)
//     Type                                     = MARKET_DATA_REQUEST  (101)
//     BaseSize                                 = MarketDataRequestSize  (28)
//     RequestAction                            = SUBSCRIBE  (1)
//     SymbolID                                 = 0
//     Symbol                                   = ""
//     Exchange                                 = ""
//     IntervalForSnapshotUpdatesInMilliseconds = 0
// }
func (m MarketDataRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataRequestDefault)
}

// ToBuilder
func (m MarketDataRequestPointer) ToBuilder() MarketDataRequestPointerBuilder {
	return MarketDataRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *MarketDataRequestPointerBuilder) Finish() MarketDataRequestPointer {
	return MarketDataRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The Server will respond with an initial MarketDataSnapshot
// message and then provide MARKET_DATA_UPDATE_* updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
//
// To request only a MarketDataSnapshot message, set this to SNAPSHOT.
func (m MarketDataRequestPointer) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Ptr, 12, 8))
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The Server will respond with an initial MarketDataSnapshot
// message and then provide MARKET_DATA_UPDATE_* updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
//
// To request only a MarketDataSnapshot message, set this to SNAPSHOT.
func (m MarketDataRequestPointerBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Ptr, 12, 8))
}

// SymbolID This is the identifier which will be used in all of the market data response
// messages. This identifier is used so that the Symbol and Exchange do not
// have to be passed back in response messages from the Server. This makes
// the market data feed bandwidth efficient.
//
// If the Server receives a MarketDataRequest for a Symbol and Exchange to
// subscribe to data for, that is currently subscribed to and this SymbolID
// is different, then the Server needs to reject it.
func (m MarketDataRequestPointer) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 16, 12)
}

// SymbolID This is the identifier which will be used in all of the market data response
// messages. This identifier is used so that the Symbol and Exchange do not
// have to be passed back in response messages from the Server. This makes
// the market data feed bandwidth efficient.
//
// If the Server receives a MarketDataRequest for a Symbol and Exchange to
// subscribe to data for, that is currently subscribed to and this SymbolID
// is different, then the Server needs to reject it.
func (m MarketDataRequestPointerBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 16, 12)
}

// Symbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Symbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Exchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Exchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// IntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestPointer) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32VLS(m.Ptr, 28, 24)
}

// IntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestPointerBuilder) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32VLS(m.Ptr, 28, 24)
}

// SetRequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The Server will respond with an initial MarketDataSnapshot
// message and then provide MARKET_DATA_UPDATE_* updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
//
// To request only a MarketDataSnapshot message, set this to SNAPSHOT.
func (m MarketDataRequestPointerBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Ptr, 12, 8, int32(value))
}

// SetSymbolID This is the identifier which will be used in all of the market data response
// messages. This identifier is used so that the Symbol and Exchange do not
// have to be passed back in response messages from the Server. This makes
// the market data feed bandwidth efficient.
//
// If the Server receives a MarketDataRequest for a Symbol and Exchange to
// subscribe to data for, that is currently subscribed to and this SymbolID
// is different, then the Server needs to reject it.
func (m MarketDataRequestPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Ptr, 16, 12, value)
}

// SetSymbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m *MarketDataRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetExchange Optional Exchange. Not set when unsubscribing.
func (m *MarketDataRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetIntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestPointerBuilder) SetIntervalForSnapshotUpdatesInMilliseconds(value uint32) {
	message.SetUint32VLS(m.Ptr, 28, 24, value)
}

// Copy
func (m MarketDataRequestPointer) Copy(to MarketDataRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// Copy
func (m MarketDataRequestPointerBuilder) Copy(to MarketDataRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyPointer
func (m MarketDataRequestPointer) CopyPointer(to MarketDataRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyPointer
func (m MarketDataRequestPointerBuilder) CopyPointer(to MarketDataRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyTo
func (m MarketDataRequestPointer) CopyTo(to MarketDataRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyTo
func (m MarketDataRequestPointerBuilder) CopyTo(to MarketDataRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyToPointer
func (m MarketDataRequestPointer) CopyToPointer(to MarketDataRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyToPointer
func (m MarketDataRequestPointerBuilder) CopyToPointer(to MarketDataRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}
