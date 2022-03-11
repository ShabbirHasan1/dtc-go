package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = MarketDepthRequestSize  (28)
//     Type          = MARKET_DEPTH_REQUEST  (102)
//     BaseSize      = MarketDepthRequestSize  (28)
//     RequestAction = SUBSCRIBE  (1)
//     SymbolID      = 0
//     Symbol        = ""
//     Exchange      = ""
//     NumLevels     = 10
// }
var _MarketDepthRequestDefault = []byte{28, 0, 102, 0, 28, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0}

const MarketDepthRequestSize = 28

type MarketDepthRequest struct {
	message.VLS
}

type MarketDepthRequestBuilder struct {
	message.VLSBuilder
}

func NewMarketDepthRequestFrom(b []byte) MarketDepthRequest {
	return MarketDepthRequest{VLS: message.NewVLSFrom(b)}
}

func WrapMarketDepthRequest(b []byte) MarketDepthRequest {
	return MarketDepthRequest{VLS: message.WrapVLS(b)}
}

func NewMarketDepthRequest() MarketDepthRequestBuilder {
	a := MarketDepthRequestBuilder{message.NewVLSBuilder(28)}
	a.Unsafe().SetBytes(0, _MarketDepthRequestDefault)
	return a
}

// Clear
// {
//     Size          = MarketDepthRequestSize  (28)
//     Type          = MARKET_DEPTH_REQUEST  (102)
//     BaseSize      = MarketDepthRequestSize  (28)
//     RequestAction = SUBSCRIBE  (1)
//     SymbolID      = 0
//     Symbol        = ""
//     Exchange      = ""
//     NumLevels     = 10
// }
func (m MarketDepthRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthRequestDefault)
}

// ToBuilder
func (m MarketDepthRequest) ToBuilder() MarketDepthRequestBuilder {
	return MarketDepthRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m MarketDepthRequestBuilder) Finish() MarketDepthRequest {
	return MarketDepthRequest{m.VLSBuilder.Finish()}
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequest) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Unsafe(), 12, 8))
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequestBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Unsafe(), 12, 8))
}

// SymbolID This is the identifier which will be used in all of the market depth data
// response messages.
//
// This SymbolID can be the same as the one used in the MARKET_DATA_REQUEST
// message for the same Symbol and Exchange.
//
// This identifier is used so that the Symbol does not have to be passed
// back in response messages from the Server. If the Server receives a MARKET_DEPTH_REQUEST
// for a Symbol and Exchange to subscribe to market depth data for, that
// is currently subscribed to and this SymbolID is different, then the Server
// should reject it.
func (m MarketDepthRequest) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 16, 12)
}

// SymbolID This is the identifier which will be used in all of the market depth data
// response messages.
//
// This SymbolID can be the same as the one used in the MARKET_DATA_REQUEST
// message for the same Symbol and Exchange.
//
// This identifier is used so that the Symbol does not have to be passed
// back in response messages from the Server. If the Server receives a MARKET_DEPTH_REQUEST
// for a Symbol and Exchange to subscribe to market depth data for, that
// is currently subscribed to and this SymbolID is different, then the Server
// should reject it.
func (m MarketDepthRequestBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 16, 12)
}

// Symbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Symbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Exchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequest) NumLevels() int32 {
	return message.Int32VLS(m.Unsafe(), 28, 24)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestBuilder) NumLevels() int32 {
	return message.Int32VLS(m.Unsafe(), 28, 24)
}

// SetRequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequestBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, int32(value))
}

// SetSymbolID This is the identifier which will be used in all of the market depth data
// response messages.
//
// This SymbolID can be the same as the one used in the MARKET_DATA_REQUEST
// message for the same Symbol and Exchange.
//
// This identifier is used so that the Symbol does not have to be passed
// back in response messages from the Server. If the Server receives a MARKET_DEPTH_REQUEST
// for a Symbol and Exchange to subscribe to market depth data for, that
// is currently subscribed to and this SymbolID is different, then the Server
// should reject it.
func (m MarketDepthRequestBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 16, 12, value)
}

// SetSymbol The symbol for the market depth request. Not set when unsubscribing.
func (m *MarketDepthRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetExchange The optional exchange for the symbol. Not set when unsubscribing.
func (m *MarketDepthRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetNumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestBuilder) SetNumLevels(value int32) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, value)
}

// Copy
func (m MarketDepthRequest) Copy(to MarketDepthRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// Copy
func (m MarketDepthRequestBuilder) Copy(to MarketDepthRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyPointer
func (m MarketDepthRequest) CopyPointer(to MarketDepthRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyPointer
func (m MarketDepthRequestBuilder) CopyPointer(to MarketDepthRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyToPointer
func (m MarketDepthRequest) CopyToPointer(to MarketDepthRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyToPointer
func (m MarketDepthRequestBuilder) CopyToPointer(to MarketDepthRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyTo
func (m MarketDepthRequest) CopyTo(to MarketDepthRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyTo
func (m MarketDepthRequestBuilder) CopyTo(to MarketDepthRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}
