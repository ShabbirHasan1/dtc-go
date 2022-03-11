package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketDepthRequestPointer struct {
	message.VLSPointer
}

type MarketDepthRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocMarketDepthRequest() MarketDepthRequestPointerBuilder {
	a := MarketDepthRequestPointerBuilder{message.AllocVLSBuilder(28)}
	a.Ptr.SetBytes(0, _MarketDepthRequestDefault)
	return a
}

func AllocMarketDepthRequestFrom(b []byte) MarketDepthRequestPointer {
	return MarketDepthRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m MarketDepthRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthRequestDefault)
}

// ToBuilder
func (m MarketDepthRequestPointer) ToBuilder() MarketDepthRequestPointerBuilder {
	return MarketDepthRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *MarketDepthRequestPointerBuilder) Finish() MarketDepthRequestPointer {
	return MarketDepthRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequestPointer) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Ptr, 12, 8))
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequestPointerBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32VLS(m.Ptr, 12, 8))
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
func (m MarketDepthRequestPointer) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 16, 12)
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
func (m MarketDepthRequestPointerBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 16, 12)
}

// Symbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Symbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Exchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Exchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestPointer) NumLevels() int32 {
	return message.Int32VLS(m.Ptr, 28, 24)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestPointerBuilder) NumLevels() int32 {
	return message.Int32VLS(m.Ptr, 28, 24)
}

// SetRequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequestPointerBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Ptr, 12, 8, int32(value))
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
func (m MarketDepthRequestPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Ptr, 16, 12, value)
}

// SetSymbol The symbol for the market depth request. Not set when unsubscribing.
func (m *MarketDepthRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetExchange The optional exchange for the symbol. Not set when unsubscribing.
func (m *MarketDepthRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetNumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestPointerBuilder) SetNumLevels(value int32) {
	message.SetInt32VLS(m.Ptr, 28, 24, value)
}

// Copy
func (m MarketDepthRequestPointer) Copy(to MarketDepthRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// Copy
func (m MarketDepthRequestPointerBuilder) Copy(to MarketDepthRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyPointer
func (m MarketDepthRequestPointer) CopyPointer(to MarketDepthRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyPointer
func (m MarketDepthRequestPointerBuilder) CopyPointer(to MarketDepthRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyTo
func (m MarketDepthRequestPointer) CopyTo(to MarketDepthRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyTo
func (m MarketDepthRequestPointerBuilder) CopyTo(to MarketDepthRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyToPointer
func (m MarketDepthRequestPointer) CopyToPointer(to MarketDepthRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyToPointer
func (m MarketDepthRequestPointerBuilder) CopyToPointer(to MarketDepthRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}
