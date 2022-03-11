package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = MarketDepthRequestFixedSize  (96)
//     Type          = MARKET_DEPTH_REQUEST  (102)
//     RequestAction = SUBSCRIBE  (1)
//     SymbolID      = 0
//     Symbol        = ""
//     Exchange      = ""
//     NumLevels     = 0
// }
var _MarketDepthRequestFixedDefault = []byte{96, 0, 102, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDepthRequestFixedSize = 96

type MarketDepthRequestFixed struct {
	message.Fixed
}

type MarketDepthRequestFixedBuilder struct {
	message.Fixed
}

func NewMarketDepthRequestFixedFrom(b []byte) MarketDepthRequestFixed {
	return MarketDepthRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDepthRequestFixed(b []byte) MarketDepthRequestFixed {
	return MarketDepthRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewMarketDepthRequestFixed() MarketDepthRequestFixedBuilder {
	a := MarketDepthRequestFixedBuilder{message.NewFixed(96)}
	a.Unsafe().SetBytes(0, _MarketDepthRequestFixedDefault)
	return a
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
func (m MarketDepthRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthRequestFixedDefault)
}

// ToBuilder
func (m MarketDepthRequestFixed) ToBuilder() MarketDepthRequestFixedBuilder {
	return MarketDepthRequestFixedBuilder{m.Fixed}
}

// Finish
func (m MarketDepthRequestFixedBuilder) Finish() MarketDepthRequestFixed {
	return MarketDepthRequestFixed{m.Fixed}
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequestFixed) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Unsafe(), 8, 4))
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequestFixedBuilder) RequestAction() RequestActionEnum {
	return RequestActionEnum(message.Int32Fixed(m.Unsafe(), 8, 4))
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
func (m MarketDepthRequestFixed) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
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
func (m MarketDepthRequestFixedBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 12, 8)
}

// Symbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 76, 12)
}

// Symbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 76, 12)
}

// Exchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 92, 76)
}

// Exchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 92, 76)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestFixed) NumLevels() int32 {
	return message.Int32Fixed(m.Unsafe(), 96, 92)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestFixedBuilder) NumLevels() int32 {
	return message.Int32Fixed(m.Unsafe(), 96, 92)
}

// SetRequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequestFixedBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, int32(value))
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
func (m MarketDepthRequestFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 12, 8, value)
}

// SetSymbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 76, 12, value)
}

// SetExchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 92, 76, value)
}

// SetNumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestFixedBuilder) SetNumLevels(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 96, 92, value)
}

// Copy
func (m MarketDepthRequestFixed) Copy(to MarketDepthRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// Copy
func (m MarketDepthRequestFixedBuilder) Copy(to MarketDepthRequestFixedBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyPointer
func (m MarketDepthRequestFixed) CopyPointer(to MarketDepthRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyPointer
func (m MarketDepthRequestFixedBuilder) CopyPointer(to MarketDepthRequestFixedPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyToPointer
func (m MarketDepthRequestFixed) CopyToPointer(to MarketDepthRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyToPointer
func (m MarketDepthRequestFixedBuilder) CopyToPointer(to MarketDepthRequestPointerBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyTo
func (m MarketDepthRequestFixed) CopyTo(to MarketDepthRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyTo
func (m MarketDepthRequestFixedBuilder) CopyTo(to MarketDepthRequestBuilder) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}
