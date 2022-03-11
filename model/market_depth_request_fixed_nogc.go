package model

import (
	"github.com/moontrade/dtc-go/message"
)

type MarketDepthRequestFixedPointer struct {
	message.FixedPointer
}

type MarketDepthRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocMarketDepthRequestFixed() MarketDepthRequestFixedPointerBuilder {
	a := MarketDepthRequestFixedPointerBuilder{message.AllocFixed(96)}
	a.Ptr.SetBytes(0, _MarketDepthRequestFixedDefault)
	return a
}

func AllocMarketDepthRequestFixedFrom(b []byte) MarketDepthRequestFixedPointer {
	return MarketDepthRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = MarketDepthRequestFixedSize  (96)
//     Type          = MARKET_DEPTH_REQUEST  (102)
//     RequestAction = SUBSCRIBE  (1)
//     SymbolID      = 0
//     Symbol        = ""
//     Exchange      = ""
//     NumLevels     = 0
// }
func (m MarketDepthRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthRequestFixedDefault)
}

// ToBuilder
func (m MarketDepthRequestFixedPointer) ToBuilder() MarketDepthRequestFixedPointerBuilder {
	return MarketDepthRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *MarketDepthRequestFixedPointerBuilder) Finish() MarketDepthRequestFixedPointer {
	return MarketDepthRequestFixedPointer{m.FixedPointer}
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequestFixedPointer) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Ptr, 8, 4))
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequestFixedPointerBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Ptr, 8, 4))
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
func (m MarketDepthRequestFixedPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
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
func (m MarketDepthRequestFixedPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 12, 8)
}

// Symbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 76, 12)
}

// Symbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 76, 12)
}

// Exchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 92, 76)
}

// Exchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 92, 76)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestFixedPointer) NumLevels() int32 {
	return message.Int32Fixed(m.Ptr, 96, 92)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestFixedPointerBuilder) NumLevels() int32 {
	return message.Int32Fixed(m.Ptr, 96, 92)
}

// SetRequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequestFixedPointerBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32Fixed(m.Ptr, 8, 4, int32(value))
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
func (m MarketDepthRequestFixedPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 12, 8, value)
}

// SetSymbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 76, 12, value)
}

// SetExchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 92, 76, value)
}

// SetNumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestFixedPointerBuilder) SetNumLevels(value int32) {
	message.SetInt32Fixed(m.Ptr, 96, 92, value)
}

// Copy
func (m MarketDepthRequestFixedPointer) Copy(to MarketDepthRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// Copy
func (m MarketDepthRequestFixedPointerBuilder) Copy(to MarketDepthRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyPointer
func (m MarketDepthRequestFixedPointer) CopyPointer(to MarketDepthRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyPointer
func (m MarketDepthRequestFixedPointerBuilder) CopyPointer(to MarketDepthRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyTo
func (m MarketDepthRequestFixedPointer) CopyTo(to MarketDepthRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyTo
func (m MarketDepthRequestFixedPointerBuilder) CopyTo(to MarketDepthRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyToPointer
func (m MarketDepthRequestFixedPointer) CopyToPointer(to MarketDepthRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyToPointer
func (m MarketDepthRequestFixedPointerBuilder) CopyToPointer(to MarketDepthRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}
