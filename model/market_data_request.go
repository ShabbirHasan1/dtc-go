package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataRequestSize = 28

const MarketDataRequestFixedSize = 96

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

func AllocMarketDataRequest() MarketDataRequestPointerBuilder {
	a := MarketDataRequestPointerBuilder{message.AllocVLSBuilder(28)}
	a.Ptr.SetBytes(0, _MarketDataRequestDefault)
	return a
}

func AllocMarketDataRequestFrom(b []byte) MarketDataRequestPointer {
	return MarketDataRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m MarketDataRequest) ToBuilder() MarketDataRequestBuilder {
	return MarketDataRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m MarketDataRequestFixed) ToBuilder() MarketDataRequestFixedBuilder {
	return MarketDataRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataRequestPointer) ToBuilder() MarketDataRequestPointerBuilder {
	return MarketDataRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m MarketDataRequestFixedPointer) ToBuilder() MarketDataRequestFixedPointerBuilder {
	return MarketDataRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataRequestBuilder) Finish() MarketDataRequest {
	return MarketDataRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m MarketDataRequestFixedBuilder) Finish() MarketDataRequestFixed {
	return MarketDataRequestFixed{m.Fixed}
}

// Finish
func (m *MarketDataRequestPointerBuilder) Finish() MarketDataRequestPointer {
	return MarketDataRequestPointer{m.VLSPointerBuilder.Finish()}
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
func (m MarketDataRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Symbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
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
func (m MarketDataRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Exchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
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
func (m MarketDataRequest) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32VLS(m.Unsafe(), 28, 24)
}

// IntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestBuilder) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32VLS(m.Unsafe(), 28, 24)
}

// IntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestPointer) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32VLS(m.Ptr, 28, 24)
}

// IntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestPointerBuilder) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32VLS(m.Ptr, 28, 24)
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
func (m MarketDataRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 76, 12)
}

// Symbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 76, 12)
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
func (m MarketDataRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 92, 76)
}

// Exchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 92, 76)
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
func (m MarketDataRequestFixed) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 96, 92)
}

// IntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestFixedBuilder) IntervalForSnapshotUpdatesInMilliseconds() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 96, 92)
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
func (m MarketDataRequestBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, int32(value))
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
func (m MarketDataRequestBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 16, 12, value)
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
func (m *MarketDataRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetSymbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m *MarketDataRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetExchange Optional Exchange. Not set when unsubscribing.
func (m *MarketDataRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetExchange Optional Exchange. Not set when unsubscribing.
func (m *MarketDataRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetIntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestBuilder) SetIntervalForSnapshotUpdatesInMilliseconds(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 28, 24, value)
}

// SetIntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestPointerBuilder) SetIntervalForSnapshotUpdatesInMilliseconds(value uint32) {
	message.SetUint32VLS(m.Ptr, 28, 24, value)
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
func (m MarketDataRequestFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 12, 8, value)
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
func (m MarketDataRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 76, 12, value)
}

// SetSymbol The Symbol that market data is requested for. Not set when unsubscribing.
// The Symbol that market data is requested for. Not set when unsubscribing.
func (m MarketDataRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 76, 12, value)
}

// SetExchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 92, 76, value)
}

// SetExchange Optional Exchange. Not set when unsubscribing.
func (m MarketDataRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 92, 76, value)
}

// SetIntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestFixedBuilder) SetIntervalForSnapshotUpdatesInMilliseconds(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 96, 92, value)
}

// SetIntervalForSnapshotUpdatesInMilliseconds
func (m MarketDataRequestFixedPointerBuilder) SetIntervalForSnapshotUpdatesInMilliseconds(value uint32) {
	message.SetUint32Fixed(m.Ptr, 96, 92, value)
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
