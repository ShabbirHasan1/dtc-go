package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _MarketDataRequestDefault = []byte{28, 0, 101, 0, 28, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataRequestSize = 28

// MarketDataRequest The MarketDataRequest message will subscribe to market data for a particular
// Symbol or request a market data snapshot.
//
// The Server can also send market depth data in response to this message
// and not require a MarketDepthRequest.
type MarketDataRequest struct {
	message.VLS
}

// MarketDataRequestBuilder The MarketDataRequest message will subscribe to market data for a particular
// Symbol or request a market data snapshot.
//
// The Server can also send market depth data in response to this message
// and not require a MarketDepthRequest.
type MarketDataRequestBuilder struct {
	message.VLSBuilder
}

func NewMarketDataRequestFrom(b []byte) MarketDataRequest {
	return MarketDataRequest{VLS: message.NewVLSFrom(b)}
}

func WrapMarketDataRequest(b []byte) MarketDataRequest {
	return MarketDataRequest{VLS: message.WrapVLS(b)}
}

func NewMarketDataRequest() MarketDataRequestBuilder {
	a := MarketDataRequestBuilder{message.NewVLSBuilder(28)}
	a.Unsafe().SetBytes(0, _MarketDataRequestDefault)
	return a
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
func (m MarketDataRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataRequestDefault)
}

// ToBuilder
func (m MarketDataRequest) ToBuilder() MarketDataRequestBuilder {
	return MarketDataRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m MarketDataRequestBuilder) Finish() MarketDataRequest {
	return MarketDataRequest{m.VLSBuilder.Finish()}
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The Server will respond with an initial MarketDataSnapshot
// message and then provide MARKET_DATA_UPDATE_* updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
//
// To request only a MarketDataSnapshot message, set this to SNAPSHOT.
func (m MarketDataRequest) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Unsafe(), 12, 8))
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The Server will respond with an initial MarketDataSnapshot
// message and then provide MARKET_DATA_UPDATE_* updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
//
// To request only a MarketDataSnapshot message, set this to SNAPSHOT.
func (m MarketDataRequestBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Unsafe(), 12, 8))
}

// SymbolID This is the identifier which will be used in all of the market data response
// messages. This identifier is used so that the Symbol and Exchange do not
// have to be passed back in response messages from the Server. This makes
// the market data feed bandwidth efficient.
//
// If the Server receives a MarketDataRequest for a Symbol and Exchange to
// subscribe to data for, that is currently subscribed to and this SymbolID
// is different, then the Server needs to reject it.
func (m MarketDataRequest) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 16, 12)
}

// SymbolID This is the identifier which will be used in all of the market data response
// messages. This identifier is used so that the Symbol and Exchange do not
// have to be passed back in response messages from the Server. This makes
// the market data feed bandwidth efficient.
//
// If the Server receives a MarketDataRequest for a Symbol and Exchange to
// subscribe to data for, that is currently subscribed to and this SymbolID
// is different, then the Server needs to reject it.
func (m MarketDataRequestBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 16, 12)
}

// Symbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Symbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Exchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// IntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequest) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32VLS(m.Unsafe(), 28, 24)
}

// IntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestBuilder) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32VLS(m.Unsafe(), 28, 24)
}

// SetRequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The Server will respond with an initial MarketDataSnapshot
// message and then provide MARKET_DATA_UPDATE_* updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
//
// To request only a MarketDataSnapshot message, set this to SNAPSHOT.
func (m MarketDataRequestBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, int32(value))
}

// SetSymbolID This is the identifier which will be used in all of the market data response
// messages. This identifier is used so that the Symbol and Exchange do not
// have to be passed back in response messages from the Server. This makes
// the market data feed bandwidth efficient.
//
// If the Server receives a MarketDataRequest for a Symbol and Exchange to
// subscribe to data for, that is currently subscribed to and this SymbolID
// is different, then the Server needs to reject it.
func (m MarketDataRequestBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 16, 12, value)
}

// SetSymbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m *MarketDataRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetExchange Optional Exchange. Not set when unsubscribing.
func (m *MarketDataRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetIntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestBuilder) SetIntervalForSnapshotUpdatesInMilliseconds(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 28, 24, value)
}

// Copy
func (m MarketDataRequest) Copy(to MarketDataRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// Copy
func (m MarketDataRequestBuilder) Copy(to MarketDataRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyPointer
func (m MarketDataRequest) CopyPointer(to MarketDataRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyPointer
func (m MarketDataRequestBuilder) CopyPointer(to MarketDataRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyToPointer
func (m MarketDataRequest) CopyToPointer(to MarketDataRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyToPointer
func (m MarketDataRequestBuilder) CopyToPointer(to MarketDataRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyTo
func (m MarketDataRequest) CopyTo(to MarketDataRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}

// CopyTo
func (m MarketDataRequestBuilder) CopyTo(to MarketDataRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetIntervalForSnapshotUpdatesInMilliseconds(m.IntervalForSnapshotUpdatesInMilliseconds())
}
