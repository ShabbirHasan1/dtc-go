package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDepthRequestSize = 28

const MarketDepthRequestFixedSize = 96

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

type MarketDepthRequest struct {
	message.VLS
}

type MarketDepthRequestBuilder struct {
	message.VLSBuilder
}

type MarketDepthRequestFixed struct {
	message.Fixed
}

type MarketDepthRequestFixedBuilder struct {
	message.Fixed
}

type MarketDepthRequestPointer struct {
	message.VLSPointer
}

type MarketDepthRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

type MarketDepthRequestFixedPointer struct {
	message.FixedPointer
}

type MarketDepthRequestFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocMarketDepthRequest() MarketDepthRequestPointerBuilder {
	a := MarketDepthRequestPointerBuilder{message.AllocVLSBuilder(28)}
	a.Ptr.SetBytes(0, _MarketDepthRequestDefault)
	return a
}

func AllocMarketDepthRequestFrom(b []byte) MarketDepthRequestPointer {
	return MarketDepthRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m MarketDepthRequest) ToBuilder() MarketDepthRequestBuilder {
	return MarketDepthRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m MarketDepthRequestFixed) ToBuilder() MarketDepthRequestFixedBuilder {
	return MarketDepthRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDepthRequestPointer) ToBuilder() MarketDepthRequestPointerBuilder {
	return MarketDepthRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m MarketDepthRequestFixedPointer) ToBuilder() MarketDepthRequestFixedPointerBuilder {
	return MarketDepthRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDepthRequestBuilder) Finish() MarketDepthRequest {
	return MarketDepthRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m MarketDepthRequestFixedBuilder) Finish() MarketDepthRequestFixed {
	return MarketDepthRequestFixed{m.Fixed}
}

// Finish
func (m *MarketDepthRequestPointerBuilder) Finish() MarketDepthRequestPointer {
	return MarketDepthRequestPointer{m.VLSPointerBuilder.Finish()}
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
func (m MarketDepthRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Symbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
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
func (m MarketDepthRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Exchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
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
func (m MarketDepthRequest) NumLevels() int32 {
	return message.Int32VLS(m.Unsafe(), 28, 24)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestBuilder) NumLevels() int32 {
	return message.Int32VLS(m.Unsafe(), 28, 24)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestPointer) NumLevels() int32 {
	return message.Int32VLS(m.Ptr, 28, 24)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestPointerBuilder) NumLevels() int32 {
	return message.Int32VLS(m.Ptr, 28, 24)
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
func (m MarketDepthRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 76, 12)
}

// Symbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 76, 12)
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
func (m MarketDepthRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 92, 76)
}

// Exchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 92, 76)
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
func (m MarketDepthRequestFixed) NumLevels() int32 {
	return message.Int32Fixed(m.Unsafe(), 96, 92)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestFixedBuilder) NumLevels() int32 {
	return message.Int32Fixed(m.Unsafe(), 96, 92)
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
func (m MarketDepthRequestBuilder) SetRequestAction(value RequestActionEnum) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, int32(value))
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
func (m MarketDepthRequestBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 16, 12, value)
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
func (m *MarketDepthRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetSymbol The symbol for the market depth request. Not set when unsubscribing.
func (m *MarketDepthRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetExchange The optional exchange for the symbol. Not set when unsubscribing.
func (m *MarketDepthRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetExchange The optional exchange for the symbol. Not set when unsubscribing.
func (m *MarketDepthRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetNumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestBuilder) SetNumLevels(value int32) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, value)
}

// SetNumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestPointerBuilder) SetNumLevels(value int32) {
	message.SetInt32VLS(m.Ptr, 28, 24, value)
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
func (m MarketDepthRequestFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 12, 8, value)
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
func (m MarketDepthRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 76, 12, value)
}

// SetSymbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 76, 12, value)
}

// SetExchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 92, 76, value)
}

// SetExchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 92, 76, value)
}

// SetNumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestFixedBuilder) SetNumLevels(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 96, 92, value)
}

// SetNumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestFixedPointerBuilder) SetNumLevels(value int32) {
	message.SetInt32Fixed(m.Ptr, 96, 92, value)
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
