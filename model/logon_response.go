package model

import (
	"github.com/moontrade/dtc-go/message"
)

const LogonResponseSize = 52

const LogonResponseFixedSize = 256

// {
//     Size                                          = LogonResponseSize  (52)
//     Type                                          = LOGON_RESPONSE  (2)
//     BaseSize                                      = LogonResponseSize  (52)
//     ProtocolVersion                               = CURRENT_VERSION  (8)
//     Result                                        = 0
//     ResultText                                    = ""
//     ReconnectAddress                              = ""
//     Integer_1                                     = 0
//     ServerName                                    = ""
//     MarketDepthUpdatesBestBidAndAsk               = 0
//     TradingIsSupported                            = false
//     OCOOrdersSupported                            = false
//     OrderCancelReplaceSupported                   = true
//     SymbolExchangeDelimiter                       = ""
//     SecurityDefinitionsSupported                  = false
//     HistoricalPriceDataSupported                  = false
//     ResubscribeWhenMarketDataFeedAvailable        = 0
//     MarketDepthIsSupported                        = true
//     OneHistoricalPriceDataRequestPerConnection    = 0
//     BracketOrdersSupported                        = false
//     UseIntegerPriceOrderMessages                  = false
//     UsesMultiplePositionsPerSymbolAndTradeAccount = 0
//     MarketDataSupported                           = false
// }
var _LogonResponseDefault = []byte{52, 0, 2, 0, 52, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size                                          = LogonResponseFixedSize  (256)
//     Type                                          = LOGON_RESPONSE  (2)
//     ProtocolVersion                               = CURRENT_VERSION  (8)
//     Result                                        = 0
//     ResultText                                    = ""
//     ReconnectAddress                              = ""
//     Integer_1                                     = 0
//     ServerName                                    = ""
//     MarketDepthUpdatesBestBidAndAsk               = 0
//     TradingIsSupported                            = false
//     OCOOrdersSupported                            = false
//     OrderCancelReplaceSupported                   = true
//     SymbolExchangeDelimiter                       = ""
//     SecurityDefinitionsSupported                  = false
//     HistoricalPriceDataSupported                  = false
//     ResubscribeWhenMarketDataFeedAvailable        = 0
//     MarketDepthIsSupported                        = true
//     OneHistoricalPriceDataRequestPerConnection    = 0
//     BracketOrdersSupported                        = false
//     UseIntegerPriceOrderMessages                  = false
//     UsesMultiplePositionsPerSymbolAndTradeAccount = 0
//     MarketDataSupported                           = true
// }
var _LogonResponseFixedDefault = []byte{0, 1, 2, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// LogonResponse This is a response message indicating either success or an error logging
// on to the Server.
type LogonResponse struct {
	message.VLS
}

// LogonResponseBuilder This is a response message indicating either success or an error logging
// on to the Server.
type LogonResponseBuilder struct {
	message.VLSBuilder
}

// LogonResponseFixed This is a response message indicating either success or an error logging
// on to the Server.
type LogonResponseFixed struct {
	message.Fixed
}

// LogonResponseFixedBuilder This is a response message indicating either success or an error logging
// on to the Server.
type LogonResponseFixedBuilder struct {
	message.Fixed
}

// LogonResponsePointer This is a response message indicating either success or an error logging
// on to the Server.
type LogonResponsePointer struct {
	message.VLSPointer
}

// LogonResponsePointerBuilder This is a response message indicating either success or an error logging
// on to the Server.
type LogonResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

// LogonResponseFixedPointer This is a response message indicating either success or an error logging
// on to the Server.
type LogonResponseFixedPointer struct {
	message.FixedPointer
}

// LogonResponseFixedPointerBuilder This is a response message indicating either success or an error logging
// on to the Server.
type LogonResponseFixedPointerBuilder struct {
	message.FixedPointer
}

func NewLogonResponseFrom(b []byte) LogonResponse {
	return LogonResponse{VLS: message.NewVLSFrom(b)}
}

func WrapLogonResponse(b []byte) LogonResponse {
	return LogonResponse{VLS: message.WrapVLS(b)}
}

func NewLogonResponse() LogonResponseBuilder {
	a := LogonResponseBuilder{message.NewVLSBuilder(52)}
	a.Unsafe().SetBytes(0, _LogonResponseDefault)
	return a
}

func NewLogonResponseFixedFrom(b []byte) LogonResponseFixed {
	return LogonResponseFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapLogonResponseFixed(b []byte) LogonResponseFixed {
	return LogonResponseFixed{Fixed: message.WrapFixed(b)}
}

func NewLogonResponseFixed() LogonResponseFixedBuilder {
	a := LogonResponseFixedBuilder{message.NewFixed(256)}
	a.Unsafe().SetBytes(0, _LogonResponseFixedDefault)
	return a
}

func AllocLogonResponse() LogonResponsePointerBuilder {
	a := LogonResponsePointerBuilder{message.AllocVLSBuilder(52)}
	a.Ptr.SetBytes(0, _LogonResponseDefault)
	return a
}

func AllocLogonResponseFrom(b []byte) LogonResponsePointer {
	return LogonResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocLogonResponseFixed() LogonResponseFixedPointerBuilder {
	a := LogonResponseFixedPointerBuilder{message.AllocFixed(256)}
	a.Ptr.SetBytes(0, _LogonResponseFixedDefault)
	return a
}

func AllocLogonResponseFixedFrom(b []byte) LogonResponseFixedPointer {
	return LogonResponseFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                                          = LogonResponseSize  (52)
//     Type                                          = LOGON_RESPONSE  (2)
//     BaseSize                                      = LogonResponseSize  (52)
//     ProtocolVersion                               = CURRENT_VERSION  (8)
//     Result                                        = 0
//     ResultText                                    = ""
//     ReconnectAddress                              = ""
//     Integer_1                                     = 0
//     ServerName                                    = ""
//     MarketDepthUpdatesBestBidAndAsk               = 0
//     TradingIsSupported                            = false
//     OCOOrdersSupported                            = false
//     OrderCancelReplaceSupported                   = true
//     SymbolExchangeDelimiter                       = ""
//     SecurityDefinitionsSupported                  = false
//     HistoricalPriceDataSupported                  = false
//     ResubscribeWhenMarketDataFeedAvailable        = 0
//     MarketDepthIsSupported                        = true
//     OneHistoricalPriceDataRequestPerConnection    = 0
//     BracketOrdersSupported                        = false
//     UseIntegerPriceOrderMessages                  = false
//     UsesMultiplePositionsPerSymbolAndTradeAccount = 0
//     MarketDataSupported                           = false
// }
func (m LogonResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _LogonResponseDefault)
}

// Clear
// {
//     Size                                          = LogonResponseFixedSize  (256)
//     Type                                          = LOGON_RESPONSE  (2)
//     ProtocolVersion                               = CURRENT_VERSION  (8)
//     Result                                        = 0
//     ResultText                                    = ""
//     ReconnectAddress                              = ""
//     Integer_1                                     = 0
//     ServerName                                    = ""
//     MarketDepthUpdatesBestBidAndAsk               = 0
//     TradingIsSupported                            = false
//     OCOOrdersSupported                            = false
//     OrderCancelReplaceSupported                   = true
//     SymbolExchangeDelimiter                       = ""
//     SecurityDefinitionsSupported                  = false
//     HistoricalPriceDataSupported                  = false
//     ResubscribeWhenMarketDataFeedAvailable        = 0
//     MarketDepthIsSupported                        = true
//     OneHistoricalPriceDataRequestPerConnection    = 0
//     BracketOrdersSupported                        = false
//     UseIntegerPriceOrderMessages                  = false
//     UsesMultiplePositionsPerSymbolAndTradeAccount = 0
//     MarketDataSupported                           = true
// }
func (m LogonResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _LogonResponseFixedDefault)
}

// Clear
// {
//     Size                                          = LogonResponseSize  (52)
//     Type                                          = LOGON_RESPONSE  (2)
//     BaseSize                                      = LogonResponseSize  (52)
//     ProtocolVersion                               = CURRENT_VERSION  (8)
//     Result                                        = 0
//     ResultText                                    = ""
//     ReconnectAddress                              = ""
//     Integer_1                                     = 0
//     ServerName                                    = ""
//     MarketDepthUpdatesBestBidAndAsk               = 0
//     TradingIsSupported                            = false
//     OCOOrdersSupported                            = false
//     OrderCancelReplaceSupported                   = true
//     SymbolExchangeDelimiter                       = ""
//     SecurityDefinitionsSupported                  = false
//     HistoricalPriceDataSupported                  = false
//     ResubscribeWhenMarketDataFeedAvailable        = 0
//     MarketDepthIsSupported                        = true
//     OneHistoricalPriceDataRequestPerConnection    = 0
//     BracketOrdersSupported                        = false
//     UseIntegerPriceOrderMessages                  = false
//     UsesMultiplePositionsPerSymbolAndTradeAccount = 0
//     MarketDataSupported                           = false
// }
func (m LogonResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _LogonResponseDefault)
}

// Clear
// {
//     Size                                          = LogonResponseFixedSize  (256)
//     Type                                          = LOGON_RESPONSE  (2)
//     ProtocolVersion                               = CURRENT_VERSION  (8)
//     Result                                        = 0
//     ResultText                                    = ""
//     ReconnectAddress                              = ""
//     Integer_1                                     = 0
//     ServerName                                    = ""
//     MarketDepthUpdatesBestBidAndAsk               = 0
//     TradingIsSupported                            = false
//     OCOOrdersSupported                            = false
//     OrderCancelReplaceSupported                   = true
//     SymbolExchangeDelimiter                       = ""
//     SecurityDefinitionsSupported                  = false
//     HistoricalPriceDataSupported                  = false
//     ResubscribeWhenMarketDataFeedAvailable        = 0
//     MarketDepthIsSupported                        = true
//     OneHistoricalPriceDataRequestPerConnection    = 0
//     BracketOrdersSupported                        = false
//     UseIntegerPriceOrderMessages                  = false
//     UsesMultiplePositionsPerSymbolAndTradeAccount = 0
//     MarketDataSupported                           = true
// }
func (m LogonResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _LogonResponseFixedDefault)
}

// ToBuilder
func (m LogonResponse) ToBuilder() LogonResponseBuilder {
	return LogonResponseBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m LogonResponseFixed) ToBuilder() LogonResponseFixedBuilder {
	return LogonResponseFixedBuilder{m.Fixed}
}

// ToBuilder
func (m LogonResponsePointer) ToBuilder() LogonResponsePointerBuilder {
	return LogonResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m LogonResponseFixedPointer) ToBuilder() LogonResponseFixedPointerBuilder {
	return LogonResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m LogonResponseBuilder) Finish() LogonResponse {
	return LogonResponse{m.VLSBuilder.Finish()}
}

// Finish
func (m LogonResponseFixedBuilder) Finish() LogonResponseFixed {
	return LogonResponseFixed{m.Fixed}
}

// Finish
func (m *LogonResponsePointerBuilder) Finish() LogonResponsePointer {
	return LogonResponsePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *LogonResponseFixedPointerBuilder) Finish() LogonResponseFixedPointer {
	return LogonResponseFixedPointer{m.FixedPointer}
}

// ProtocolVersion This is automatically set by the constructor.
func (m LogonResponse) ProtocolVersion() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// ProtocolVersion This is automatically set by the constructor.
func (m LogonResponseBuilder) ProtocolVersion() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// ProtocolVersion This is automatically set by the constructor.
func (m LogonResponsePointer) ProtocolVersion() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// ProtocolVersion This is automatically set by the constructor.
func (m LogonResponsePointerBuilder) ProtocolVersion() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// Result This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponse) Result() LogonStatusEnum {
	return LogonStatusEnum(message.Int32VLS(m.Unsafe(), 16, 12))
}

// Result This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponseBuilder) Result() LogonStatusEnum {
	return LogonStatusEnum(message.Int32VLS(m.Unsafe(), 16, 12))
}

// Result This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponsePointer) Result() LogonStatusEnum {
	return LogonStatusEnum(message.Int32VLS(m.Ptr, 16, 12))
}

// Result This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponsePointerBuilder) Result() LogonStatusEnum {
	return LogonStatusEnum(message.Int32VLS(m.Ptr, 16, 12))
}

// ResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponse) ResultText() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// ResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponseBuilder) ResultText() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// ResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponsePointer) ResultText() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// ResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponsePointerBuilder) ResultText() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// ReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponse) ReconnectAddress() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// ReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponseBuilder) ReconnectAddress() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// ReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponsePointer) ReconnectAddress() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// ReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponsePointerBuilder) ReconnectAddress() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Integer_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponse) Integer_1() int32 {
	return message.Int32VLS(m.Unsafe(), 28, 24)
}

// Integer_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponseBuilder) Integer_1() int32 {
	return message.Int32VLS(m.Unsafe(), 28, 24)
}

// Integer_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponsePointer) Integer_1() int32 {
	return message.Int32VLS(m.Ptr, 28, 24)
}

// Integer_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponsePointerBuilder) Integer_1() int32 {
	return message.Int32VLS(m.Ptr, 28, 24)
}

// ServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m LogonResponse) ServerName() string {
	return message.StringVLS(m.Unsafe(), 32, 28)
}

// ServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m LogonResponseBuilder) ServerName() string {
	return message.StringVLS(m.Unsafe(), 32, 28)
}

// ServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m LogonResponsePointer) ServerName() string {
	return message.StringVLS(m.Ptr, 32, 28)
}

// ServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m LogonResponsePointerBuilder) ServerName() string {
	return message.StringVLS(m.Ptr, 32, 28)
}

// MarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponse) MarketDepthUpdatesBestBidAndAsk() uint8 {
	return message.Uint8VLS(m.Unsafe(), 33, 32)
}

// MarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponseBuilder) MarketDepthUpdatesBestBidAndAsk() uint8 {
	return message.Uint8VLS(m.Unsafe(), 33, 32)
}

// MarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponsePointer) MarketDepthUpdatesBestBidAndAsk() uint8 {
	return message.Uint8VLS(m.Ptr, 33, 32)
}

// MarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponsePointerBuilder) MarketDepthUpdatesBestBidAndAsk() uint8 {
	return message.Uint8VLS(m.Ptr, 33, 32)
}

// TradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponse) TradingIsSupported() bool {
	return message.BoolVLS(m.Unsafe(), 34, 33)
}

// TradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseBuilder) TradingIsSupported() bool {
	return message.BoolVLS(m.Unsafe(), 34, 33)
}

// TradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponsePointer) TradingIsSupported() bool {
	return message.BoolVLS(m.Ptr, 34, 33)
}

// TradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponsePointerBuilder) TradingIsSupported() bool {
	return message.BoolVLS(m.Ptr, 34, 33)
}

// OCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponse) OCOOrdersSupported() bool {
	return message.BoolVLS(m.Unsafe(), 35, 34)
}

// OCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseBuilder) OCOOrdersSupported() bool {
	return message.BoolVLS(m.Unsafe(), 35, 34)
}

// OCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponsePointer) OCOOrdersSupported() bool {
	return message.BoolVLS(m.Ptr, 35, 34)
}

// OCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponsePointerBuilder) OCOOrdersSupported() bool {
	return message.BoolVLS(m.Ptr, 35, 34)
}

// OrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponse) OrderCancelReplaceSupported() bool {
	return message.BoolVLS(m.Unsafe(), 36, 35)
}

// OrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseBuilder) OrderCancelReplaceSupported() bool {
	return message.BoolVLS(m.Unsafe(), 36, 35)
}

// OrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponsePointer) OrderCancelReplaceSupported() bool {
	return message.BoolVLS(m.Ptr, 36, 35)
}

// OrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponsePointerBuilder) OrderCancelReplaceSupported() bool {
	return message.BoolVLS(m.Ptr, 36, 35)
}

// SymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m LogonResponse) SymbolExchangeDelimiter() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// SymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m LogonResponseBuilder) SymbolExchangeDelimiter() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// SymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m LogonResponsePointer) SymbolExchangeDelimiter() string {
	return message.StringVLS(m.Ptr, 40, 36)
}

// SymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m LogonResponsePointerBuilder) SymbolExchangeDelimiter() string {
	return message.StringVLS(m.Ptr, 40, 36)
}

// SecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponse) SecurityDefinitionsSupported() bool {
	return message.BoolVLS(m.Unsafe(), 41, 40)
}

// SecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseBuilder) SecurityDefinitionsSupported() bool {
	return message.BoolVLS(m.Unsafe(), 41, 40)
}

// SecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponsePointer) SecurityDefinitionsSupported() bool {
	return message.BoolVLS(m.Ptr, 41, 40)
}

// SecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponsePointerBuilder) SecurityDefinitionsSupported() bool {
	return message.BoolVLS(m.Ptr, 41, 40)
}

// HistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponse) HistoricalPriceDataSupported() bool {
	return message.BoolVLS(m.Unsafe(), 42, 41)
}

// HistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseBuilder) HistoricalPriceDataSupported() bool {
	return message.BoolVLS(m.Unsafe(), 42, 41)
}

// HistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponsePointer) HistoricalPriceDataSupported() bool {
	return message.BoolVLS(m.Ptr, 42, 41)
}

// HistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponsePointerBuilder) HistoricalPriceDataSupported() bool {
	return message.BoolVLS(m.Ptr, 42, 41)
}

// ResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponse) ResubscribeWhenMarketDataFeedAvailable() uint8 {
	return message.Uint8VLS(m.Unsafe(), 43, 42)
}

// ResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponseBuilder) ResubscribeWhenMarketDataFeedAvailable() uint8 {
	return message.Uint8VLS(m.Unsafe(), 43, 42)
}

// ResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponsePointer) ResubscribeWhenMarketDataFeedAvailable() uint8 {
	return message.Uint8VLS(m.Ptr, 43, 42)
}

// ResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponsePointerBuilder) ResubscribeWhenMarketDataFeedAvailable() uint8 {
	return message.Uint8VLS(m.Ptr, 43, 42)
}

// MarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponse) MarketDepthIsSupported() bool {
	return message.BoolVLS(m.Unsafe(), 44, 43)
}

// MarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseBuilder) MarketDepthIsSupported() bool {
	return message.BoolVLS(m.Unsafe(), 44, 43)
}

// MarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponsePointer) MarketDepthIsSupported() bool {
	return message.BoolVLS(m.Ptr, 44, 43)
}

// MarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponsePointerBuilder) MarketDepthIsSupported() bool {
	return message.BoolVLS(m.Ptr, 44, 43)
}

// OneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponse) OneHistoricalPriceDataRequestPerConnection() uint8 {
	return message.Uint8VLS(m.Unsafe(), 45, 44)
}

// OneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponseBuilder) OneHistoricalPriceDataRequestPerConnection() uint8 {
	return message.Uint8VLS(m.Unsafe(), 45, 44)
}

// OneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponsePointer) OneHistoricalPriceDataRequestPerConnection() uint8 {
	return message.Uint8VLS(m.Ptr, 45, 44)
}

// OneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponsePointerBuilder) OneHistoricalPriceDataRequestPerConnection() uint8 {
	return message.Uint8VLS(m.Ptr, 45, 44)
}

// BracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponse) BracketOrdersSupported() bool {
	return message.BoolVLS(m.Unsafe(), 46, 45)
}

// BracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseBuilder) BracketOrdersSupported() bool {
	return message.BoolVLS(m.Unsafe(), 46, 45)
}

// BracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponsePointer) BracketOrdersSupported() bool {
	return message.BoolVLS(m.Ptr, 46, 45)
}

// BracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponsePointerBuilder) BracketOrdersSupported() bool {
	return message.BoolVLS(m.Ptr, 46, 45)
}

// UseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponse) UseIntegerPriceOrderMessages() bool {
	return message.BoolVLS(m.Unsafe(), 47, 46)
}

// UseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseBuilder) UseIntegerPriceOrderMessages() bool {
	return message.BoolVLS(m.Unsafe(), 47, 46)
}

// UseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponsePointer) UseIntegerPriceOrderMessages() bool {
	return message.BoolVLS(m.Ptr, 47, 46)
}

// UseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponsePointerBuilder) UseIntegerPriceOrderMessages() bool {
	return message.BoolVLS(m.Ptr, 47, 46)
}

// UsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponse) UsesMultiplePositionsPerSymbolAndTradeAccount() uint8 {
	return message.Uint8VLS(m.Unsafe(), 48, 47)
}

// UsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponseBuilder) UsesMultiplePositionsPerSymbolAndTradeAccount() uint8 {
	return message.Uint8VLS(m.Unsafe(), 48, 47)
}

// UsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponsePointer) UsesMultiplePositionsPerSymbolAndTradeAccount() uint8 {
	return message.Uint8VLS(m.Ptr, 48, 47)
}

// UsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponsePointerBuilder) UsesMultiplePositionsPerSymbolAndTradeAccount() uint8 {
	return message.Uint8VLS(m.Ptr, 48, 47)
}

// MarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponse) MarketDataSupported() bool {
	return message.BoolVLS(m.Unsafe(), 49, 48)
}

// MarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseBuilder) MarketDataSupported() bool {
	return message.BoolVLS(m.Unsafe(), 49, 48)
}

// MarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponsePointer) MarketDataSupported() bool {
	return message.BoolVLS(m.Ptr, 49, 48)
}

// MarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponsePointerBuilder) MarketDataSupported() bool {
	return message.BoolVLS(m.Ptr, 49, 48)
}

// ProtocolVersion This is automatically set by the constructor.
func (m LogonResponseFixed) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// ProtocolVersion This is automatically set by the constructor.
func (m LogonResponseFixedBuilder) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// ProtocolVersion This is automatically set by the constructor.
func (m LogonResponseFixedPointer) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// ProtocolVersion This is automatically set by the constructor.
func (m LogonResponseFixedPointerBuilder) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Result This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponseFixed) Result() LogonStatusEnum {
	return LogonStatusEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Result This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponseFixedBuilder) Result() LogonStatusEnum {
	return LogonStatusEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Result This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponseFixedPointer) Result() LogonStatusEnum {
	return LogonStatusEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// Result This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponseFixedPointerBuilder) Result() LogonStatusEnum {
	return LogonStatusEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// ResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponseFixed) ResultText() string {
	return message.StringFixed(m.Unsafe(), 108, 12)
}

// ResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponseFixedBuilder) ResultText() string {
	return message.StringFixed(m.Unsafe(), 108, 12)
}

// ResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponseFixedPointer) ResultText() string {
	return message.StringFixed(m.Ptr, 108, 12)
}

// ResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponseFixedPointerBuilder) ResultText() string {
	return message.StringFixed(m.Ptr, 108, 12)
}

// ReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponseFixed) ReconnectAddress() string {
	return message.StringFixed(m.Unsafe(), 172, 108)
}

// ReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponseFixedBuilder) ReconnectAddress() string {
	return message.StringFixed(m.Unsafe(), 172, 108)
}

// ReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponseFixedPointer) ReconnectAddress() string {
	return message.StringFixed(m.Ptr, 172, 108)
}

// ReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponseFixedPointerBuilder) ReconnectAddress() string {
	return message.StringFixed(m.Ptr, 172, 108)
}

// Integer_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponseFixed) Integer_1() int32 {
	return message.Int32Fixed(m.Unsafe(), 176, 172)
}

// Integer_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponseFixedBuilder) Integer_1() int32 {
	return message.Int32Fixed(m.Unsafe(), 176, 172)
}

// Integer_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponseFixedPointer) Integer_1() int32 {
	return message.Int32Fixed(m.Ptr, 176, 172)
}

// Integer_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponseFixedPointerBuilder) Integer_1() int32 {
	return message.Int32Fixed(m.Ptr, 176, 172)
}

// ServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m LogonResponseFixed) ServerName() string {
	return message.StringFixed(m.Unsafe(), 236, 176)
}

// ServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m LogonResponseFixedBuilder) ServerName() string {
	return message.StringFixed(m.Unsafe(), 236, 176)
}

// ServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m LogonResponseFixedPointer) ServerName() string {
	return message.StringFixed(m.Ptr, 236, 176)
}

// ServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m LogonResponseFixedPointerBuilder) ServerName() string {
	return message.StringFixed(m.Ptr, 236, 176)
}

// MarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponseFixed) MarketDepthUpdatesBestBidAndAsk() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 237, 236)
}

// MarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponseFixedBuilder) MarketDepthUpdatesBestBidAndAsk() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 237, 236)
}

// MarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponseFixedPointer) MarketDepthUpdatesBestBidAndAsk() uint8 {
	return message.Uint8Fixed(m.Ptr, 237, 236)
}

// MarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponseFixedPointerBuilder) MarketDepthUpdatesBestBidAndAsk() uint8 {
	return message.Uint8Fixed(m.Ptr, 237, 236)
}

// TradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseFixed) TradingIsSupported() bool {
	return message.BoolFixed(m.Unsafe(), 238, 237)
}

// TradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseFixedBuilder) TradingIsSupported() bool {
	return message.BoolFixed(m.Unsafe(), 238, 237)
}

// TradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseFixedPointer) TradingIsSupported() bool {
	return message.BoolFixed(m.Ptr, 238, 237)
}

// TradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseFixedPointerBuilder) TradingIsSupported() bool {
	return message.BoolFixed(m.Ptr, 238, 237)
}

// OCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseFixed) OCOOrdersSupported() bool {
	return message.BoolFixed(m.Unsafe(), 239, 238)
}

// OCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseFixedBuilder) OCOOrdersSupported() bool {
	return message.BoolFixed(m.Unsafe(), 239, 238)
}

// OCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseFixedPointer) OCOOrdersSupported() bool {
	return message.BoolFixed(m.Ptr, 239, 238)
}

// OCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseFixedPointerBuilder) OCOOrdersSupported() bool {
	return message.BoolFixed(m.Ptr, 239, 238)
}

// OrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseFixed) OrderCancelReplaceSupported() bool {
	return message.BoolFixed(m.Unsafe(), 240, 239)
}

// OrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseFixedBuilder) OrderCancelReplaceSupported() bool {
	return message.BoolFixed(m.Unsafe(), 240, 239)
}

// OrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseFixedPointer) OrderCancelReplaceSupported() bool {
	return message.BoolFixed(m.Ptr, 240, 239)
}

// OrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseFixedPointerBuilder) OrderCancelReplaceSupported() bool {
	return message.BoolFixed(m.Ptr, 240, 239)
}

// SymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m LogonResponseFixed) SymbolExchangeDelimiter() string {
	return message.StringFixed(m.Unsafe(), 244, 240)
}

// SymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m LogonResponseFixedBuilder) SymbolExchangeDelimiter() string {
	return message.StringFixed(m.Unsafe(), 244, 240)
}

// SymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m LogonResponseFixedPointer) SymbolExchangeDelimiter() string {
	return message.StringFixed(m.Ptr, 244, 240)
}

// SymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m LogonResponseFixedPointerBuilder) SymbolExchangeDelimiter() string {
	return message.StringFixed(m.Ptr, 244, 240)
}

// SecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseFixed) SecurityDefinitionsSupported() bool {
	return message.BoolFixed(m.Unsafe(), 245, 244)
}

// SecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseFixedBuilder) SecurityDefinitionsSupported() bool {
	return message.BoolFixed(m.Unsafe(), 245, 244)
}

// SecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseFixedPointer) SecurityDefinitionsSupported() bool {
	return message.BoolFixed(m.Ptr, 245, 244)
}

// SecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseFixedPointerBuilder) SecurityDefinitionsSupported() bool {
	return message.BoolFixed(m.Ptr, 245, 244)
}

// HistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseFixed) HistoricalPriceDataSupported() bool {
	return message.BoolFixed(m.Unsafe(), 246, 245)
}

// HistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseFixedBuilder) HistoricalPriceDataSupported() bool {
	return message.BoolFixed(m.Unsafe(), 246, 245)
}

// HistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseFixedPointer) HistoricalPriceDataSupported() bool {
	return message.BoolFixed(m.Ptr, 246, 245)
}

// HistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseFixedPointerBuilder) HistoricalPriceDataSupported() bool {
	return message.BoolFixed(m.Ptr, 246, 245)
}

// ResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponseFixed) ResubscribeWhenMarketDataFeedAvailable() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 247, 246)
}

// ResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponseFixedBuilder) ResubscribeWhenMarketDataFeedAvailable() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 247, 246)
}

// ResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponseFixedPointer) ResubscribeWhenMarketDataFeedAvailable() uint8 {
	return message.Uint8Fixed(m.Ptr, 247, 246)
}

// ResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponseFixedPointerBuilder) ResubscribeWhenMarketDataFeedAvailable() uint8 {
	return message.Uint8Fixed(m.Ptr, 247, 246)
}

// MarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseFixed) MarketDepthIsSupported() bool {
	return message.BoolFixed(m.Unsafe(), 248, 247)
}

// MarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseFixedBuilder) MarketDepthIsSupported() bool {
	return message.BoolFixed(m.Unsafe(), 248, 247)
}

// MarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseFixedPointer) MarketDepthIsSupported() bool {
	return message.BoolFixed(m.Ptr, 248, 247)
}

// MarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseFixedPointerBuilder) MarketDepthIsSupported() bool {
	return message.BoolFixed(m.Ptr, 248, 247)
}

// OneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponseFixed) OneHistoricalPriceDataRequestPerConnection() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 249, 248)
}

// OneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponseFixedBuilder) OneHistoricalPriceDataRequestPerConnection() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 249, 248)
}

// OneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponseFixedPointer) OneHistoricalPriceDataRequestPerConnection() uint8 {
	return message.Uint8Fixed(m.Ptr, 249, 248)
}

// OneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponseFixedPointerBuilder) OneHistoricalPriceDataRequestPerConnection() uint8 {
	return message.Uint8Fixed(m.Ptr, 249, 248)
}

// BracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseFixed) BracketOrdersSupported() bool {
	return message.BoolFixed(m.Unsafe(), 250, 249)
}

// BracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseFixedBuilder) BracketOrdersSupported() bool {
	return message.BoolFixed(m.Unsafe(), 250, 249)
}

// BracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseFixedPointer) BracketOrdersSupported() bool {
	return message.BoolFixed(m.Ptr, 250, 249)
}

// BracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseFixedPointerBuilder) BracketOrdersSupported() bool {
	return message.BoolFixed(m.Ptr, 250, 249)
}

// UseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseFixed) UseIntegerPriceOrderMessages() bool {
	return message.BoolFixed(m.Unsafe(), 251, 250)
}

// UseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseFixedBuilder) UseIntegerPriceOrderMessages() bool {
	return message.BoolFixed(m.Unsafe(), 251, 250)
}

// UseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseFixedPointer) UseIntegerPriceOrderMessages() bool {
	return message.BoolFixed(m.Ptr, 251, 250)
}

// UseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseFixedPointerBuilder) UseIntegerPriceOrderMessages() bool {
	return message.BoolFixed(m.Ptr, 251, 250)
}

// UsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponseFixed) UsesMultiplePositionsPerSymbolAndTradeAccount() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 252, 251)
}

// UsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponseFixedBuilder) UsesMultiplePositionsPerSymbolAndTradeAccount() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 252, 251)
}

// UsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponseFixedPointer) UsesMultiplePositionsPerSymbolAndTradeAccount() uint8 {
	return message.Uint8Fixed(m.Ptr, 252, 251)
}

// UsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponseFixedPointerBuilder) UsesMultiplePositionsPerSymbolAndTradeAccount() uint8 {
	return message.Uint8Fixed(m.Ptr, 252, 251)
}

// MarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseFixed) MarketDataSupported() bool {
	return message.BoolFixed(m.Unsafe(), 253, 252)
}

// MarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseFixedBuilder) MarketDataSupported() bool {
	return message.BoolFixed(m.Unsafe(), 253, 252)
}

// MarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseFixedPointer) MarketDataSupported() bool {
	return message.BoolFixed(m.Ptr, 253, 252)
}

// MarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseFixedPointerBuilder) MarketDataSupported() bool {
	return message.BoolFixed(m.Ptr, 253, 252)
}

// SetProtocolVersion This is automatically set by the constructor.
func (m LogonResponseBuilder) SetProtocolVersion(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetProtocolVersion This is automatically set by the constructor.
func (m LogonResponsePointerBuilder) SetProtocolVersion(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetResult This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponseBuilder) SetResult(value LogonStatusEnum) {
	message.SetInt32VLS(m.Unsafe(), 16, 12, int32(value))
}

// SetResult This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponsePointerBuilder) SetResult(value LogonStatusEnum) {
	message.SetInt32VLS(m.Ptr, 16, 12, int32(value))
}

// SetResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m *LogonResponseBuilder) SetResultText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m *LogonResponsePointerBuilder) SetResultText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m *LogonResponseBuilder) SetReconnectAddress(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m *LogonResponsePointerBuilder) SetReconnectAddress(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetInteger_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponseBuilder) SetInteger_1(value int32) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, value)
}

// SetInteger_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponsePointerBuilder) SetInteger_1(value int32) {
	message.SetInt32VLS(m.Ptr, 28, 24, value)
}

// SetServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m *LogonResponseBuilder) SetServerName(value string) {
	message.SetStringVLS(&m.VLSBuilder, 32, 28, value)
}

// SetServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m *LogonResponsePointerBuilder) SetServerName(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 32, 28, value)
}

// SetMarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponseBuilder) SetMarketDepthUpdatesBestBidAndAsk(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 33, 32, value)
}

// SetMarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponsePointerBuilder) SetMarketDepthUpdatesBestBidAndAsk(value uint8) {
	message.SetUint8VLS(m.Ptr, 33, 32, value)
}

// SetTradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseBuilder) SetTradingIsSupported(value bool) {
	message.SetBoolVLS(m.Unsafe(), 34, 33, value)
}

// SetTradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponsePointerBuilder) SetTradingIsSupported(value bool) {
	message.SetBoolVLS(m.Ptr, 34, 33, value)
}

// SetOCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseBuilder) SetOCOOrdersSupported(value bool) {
	message.SetBoolVLS(m.Unsafe(), 35, 34, value)
}

// SetOCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponsePointerBuilder) SetOCOOrdersSupported(value bool) {
	message.SetBoolVLS(m.Ptr, 35, 34, value)
}

// SetOrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseBuilder) SetOrderCancelReplaceSupported(value bool) {
	message.SetBoolVLS(m.Unsafe(), 36, 35, value)
}

// SetOrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponsePointerBuilder) SetOrderCancelReplaceSupported(value bool) {
	message.SetBoolVLS(m.Ptr, 36, 35, value)
}

// SetSymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m *LogonResponseBuilder) SetSymbolExchangeDelimiter(value string) {
	message.SetStringVLS(&m.VLSBuilder, 40, 36, value)
}

// SetSymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m *LogonResponsePointerBuilder) SetSymbolExchangeDelimiter(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 40, 36, value)
}

// SetSecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseBuilder) SetSecurityDefinitionsSupported(value bool) {
	message.SetBoolVLS(m.Unsafe(), 41, 40, value)
}

// SetSecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponsePointerBuilder) SetSecurityDefinitionsSupported(value bool) {
	message.SetBoolVLS(m.Ptr, 41, 40, value)
}

// SetHistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseBuilder) SetHistoricalPriceDataSupported(value bool) {
	message.SetBoolVLS(m.Unsafe(), 42, 41, value)
}

// SetHistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponsePointerBuilder) SetHistoricalPriceDataSupported(value bool) {
	message.SetBoolVLS(m.Ptr, 42, 41, value)
}

// SetResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponseBuilder) SetResubscribeWhenMarketDataFeedAvailable(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 43, 42, value)
}

// SetResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponsePointerBuilder) SetResubscribeWhenMarketDataFeedAvailable(value uint8) {
	message.SetUint8VLS(m.Ptr, 43, 42, value)
}

// SetMarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseBuilder) SetMarketDepthIsSupported(value bool) {
	message.SetBoolVLS(m.Unsafe(), 44, 43, value)
}

// SetMarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponsePointerBuilder) SetMarketDepthIsSupported(value bool) {
	message.SetBoolVLS(m.Ptr, 44, 43, value)
}

// SetOneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponseBuilder) SetOneHistoricalPriceDataRequestPerConnection(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 45, 44, value)
}

// SetOneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponsePointerBuilder) SetOneHistoricalPriceDataRequestPerConnection(value uint8) {
	message.SetUint8VLS(m.Ptr, 45, 44, value)
}

// SetBracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseBuilder) SetBracketOrdersSupported(value bool) {
	message.SetBoolVLS(m.Unsafe(), 46, 45, value)
}

// SetBracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponsePointerBuilder) SetBracketOrdersSupported(value bool) {
	message.SetBoolVLS(m.Ptr, 46, 45, value)
}

// SetUseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseBuilder) SetUseIntegerPriceOrderMessages(value bool) {
	message.SetBoolVLS(m.Unsafe(), 47, 46, value)
}

// SetUseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponsePointerBuilder) SetUseIntegerPriceOrderMessages(value bool) {
	message.SetBoolVLS(m.Ptr, 47, 46, value)
}

// SetUsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponseBuilder) SetUsesMultiplePositionsPerSymbolAndTradeAccount(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 48, 47, value)
}

// SetUsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponsePointerBuilder) SetUsesMultiplePositionsPerSymbolAndTradeAccount(value uint8) {
	message.SetUint8VLS(m.Ptr, 48, 47, value)
}

// SetMarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseBuilder) SetMarketDataSupported(value bool) {
	message.SetBoolVLS(m.Unsafe(), 49, 48, value)
}

// SetMarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponsePointerBuilder) SetMarketDataSupported(value bool) {
	message.SetBoolVLS(m.Ptr, 49, 48, value)
}

// SetProtocolVersion This is automatically set by the constructor.
func (m LogonResponseFixedBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetProtocolVersion This is automatically set by the constructor.
func (m LogonResponseFixedPointerBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetResult This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponseFixedBuilder) SetResult(value LogonStatusEnum) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, int32(value))
}

// SetResult This can be set to one of the following constants:
//
// LOGON_SUCCESS
// LOGON_ERROR
// LOGON_ERROR_NO_RECONNECT
// LOGON_RECONNECT_NEW_ADDRESS
// LOGON_ERROR_NO_RECONNECT means that there has been an error logging on
// and the Client should not try to reconnect.
//
// The Server can set this field to LOGON_RECONNECT_NEW_ADDRESS to instruct
// the Client to reconnect to the Server at a different address. The new
// address is specified through the ReconnectAddress field. This supports
// dynamic connections to a server farm.
func (m LogonResponseFixedPointerBuilder) SetResult(value LogonStatusEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
}

// SetResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponseFixedBuilder) SetResultText(value string) {
	message.SetStringFixed(m.Unsafe(), 108, 12, value)
}

// SetResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponseFixedPointerBuilder) SetResultText(value string) {
	message.SetStringFixed(m.Ptr, 108, 12, value)
}

// SetReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponseFixedBuilder) SetReconnectAddress(value string) {
	message.SetStringFixed(m.Unsafe(), 172, 108, value)
}

// SetReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponseFixedPointerBuilder) SetReconnectAddress(value string) {
	message.SetStringFixed(m.Ptr, 172, 108, value)
}

// SetInteger_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponseFixedBuilder) SetInteger_1(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 176, 172, value)
}

// SetInteger_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponseFixedPointerBuilder) SetInteger_1(value int32) {
	message.SetInt32Fixed(m.Ptr, 176, 172, value)
}

// SetServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m LogonResponseFixedBuilder) SetServerName(value string) {
	message.SetStringFixed(m.Unsafe(), 236, 176, value)
}

// SetServerName Optional free-form text for the Server to fill out.
//
// It is recommended that the Server fill this in with descriptive text identifying
// itself to the Client.
//
// The length of this text string is 60 characters when fixed length strings
// are used.
func (m LogonResponseFixedPointerBuilder) SetServerName(value string) {
	message.SetStringFixed(m.Ptr, 236, 176, value)
}

// SetMarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponseFixedBuilder) SetMarketDepthUpdatesBestBidAndAsk(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 237, 236, value)
}

// SetMarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponseFixedPointerBuilder) SetMarketDepthUpdatesBestBidAndAsk(value uint8) {
	message.SetUint8Fixed(m.Ptr, 237, 236, value)
}

// SetTradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseFixedBuilder) SetTradingIsSupported(value bool) {
	message.SetBoolFixed(m.Unsafe(), 238, 237, value)
}

// SetTradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseFixedPointerBuilder) SetTradingIsSupported(value bool) {
	message.SetBoolFixed(m.Ptr, 238, 237, value)
}

// SetOCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseFixedBuilder) SetOCOOrdersSupported(value bool) {
	message.SetBoolFixed(m.Unsafe(), 239, 238, value)
}

// SetOCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseFixedPointerBuilder) SetOCOOrdersSupported(value bool) {
	message.SetBoolFixed(m.Ptr, 239, 238, value)
}

// SetOrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseFixedBuilder) SetOrderCancelReplaceSupported(value bool) {
	message.SetBoolFixed(m.Unsafe(), 240, 239, value)
}

// SetOrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseFixedPointerBuilder) SetOrderCancelReplaceSupported(value bool) {
	message.SetBoolFixed(m.Ptr, 240, 239, value)
}

// SetSymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m LogonResponseFixedBuilder) SetSymbolExchangeDelimiter(value string) {
	message.SetStringFixed(m.Unsafe(), 244, 240, value)
}

// SetSymbolExchangeDelimiter Some Clients will usually consider the Symbol and Exchange fields as a
// single text string. If the Server will be using the Exchange field in
// DTC messages that have a Symbol and Exchange fields, it must specify the
// SymbolExchangeDelimiter field to provide a standard delimiter for the
// Client to use to combine the Symbol and the Exchange into a single text
// string.
//
// It is recommended to use a "-" or ".". Examples of how the Client will
// then combine the Symbol and exchange.
//
// Symbol-Exchange
// Symbol.Exchange
// If this field is unset, then this is an indication to the Client that
// the Exchange field in DTC Protocol messages are not used.
//
// Even if the symbols supported by a Server have an Exchange text string,
// does not mean the Server has to use the Exchange field in DTC messages.
// The Server can combine the Symbol and the Exchange in Security Definition
// responses into the Symbol field only.
//
// When a Client sees that the SymbolExchangeDelimiter field is set, then
// it can use this delimiter to combine the Symbol and Exchange into a single
// text string. When the Client is setting the Symbol and Exchange in DTC
// messages, it needs to separate out the Symbol and Exchange from the larger
// text string and set those fields separately.
func (m LogonResponseFixedPointerBuilder) SetSymbolExchangeDelimiter(value string) {
	message.SetStringFixed(m.Ptr, 244, 240, value)
}

// SetSecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseFixedBuilder) SetSecurityDefinitionsSupported(value bool) {
	message.SetBoolFixed(m.Unsafe(), 245, 244, value)
}

// SetSecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseFixedPointerBuilder) SetSecurityDefinitionsSupported(value bool) {
	message.SetBoolFixed(m.Ptr, 245, 244, value)
}

// SetHistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseFixedBuilder) SetHistoricalPriceDataSupported(value bool) {
	message.SetBoolFixed(m.Unsafe(), 246, 245, value)
}

// SetHistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseFixedPointerBuilder) SetHistoricalPriceDataSupported(value bool) {
	message.SetBoolFixed(m.Ptr, 246, 245, value)
}

// SetResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponseFixedBuilder) SetResubscribeWhenMarketDataFeedAvailable(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 247, 246, value)
}

// SetResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponseFixedPointerBuilder) SetResubscribeWhenMarketDataFeedAvailable(value uint8) {
	message.SetUint8Fixed(m.Ptr, 247, 246, value)
}

// SetMarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseFixedBuilder) SetMarketDepthIsSupported(value bool) {
	message.SetBoolFixed(m.Unsafe(), 248, 247, value)
}

// SetMarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseFixedPointerBuilder) SetMarketDepthIsSupported(value bool) {
	message.SetBoolFixed(m.Ptr, 248, 247, value)
}

// SetOneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponseFixedBuilder) SetOneHistoricalPriceDataRequestPerConnection(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 249, 248, value)
}

// SetOneHistoricalPriceDataRequestPerConnection The server can optionally set the OneHistoricalPriceDataRequestPerConnection
// field to 1 in the LogonResponse message to indicate that it only will
// accept one historical price data request per network connection.
//
// After the first request is served or rejected, the network connection
// will be gracefully closed at the appropriate time by the Server. This
// method simplifies the serving of historical price data on the Server side
// and the implementation on the Client side when data compression is used.
// and the implementation on the Client side when data compression is used.
func (m LogonResponseFixedPointerBuilder) SetOneHistoricalPriceDataRequestPerConnection(value uint8) {
	message.SetUint8Fixed(m.Ptr, 249, 248, value)
}

// SetBracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseFixedBuilder) SetBracketOrdersSupported(value bool) {
	message.SetBoolFixed(m.Unsafe(), 250, 249, value)
}

// SetBracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseFixedPointerBuilder) SetBracketOrdersSupported(value bool) {
	message.SetBoolFixed(m.Ptr, 250, 249, value)
}

// SetUseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseFixedBuilder) SetUseIntegerPriceOrderMessages(value bool) {
	message.SetBoolFixed(m.Unsafe(), 251, 250, value)
}

// SetUseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseFixedPointerBuilder) SetUseIntegerPriceOrderMessages(value bool) {
	message.SetBoolFixed(m.Ptr, 251, 250, value)
}

// SetUsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponseFixedBuilder) SetUsesMultiplePositionsPerSymbolAndTradeAccount(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 252, 251, value)
}

// SetUsesMultiplePositionsPerSymbolAndTradeAccount If the Server can report more than one Trade Position for a specific Symbol
// and Trade Account, then it needs to set UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1.
//
// When the server has set to 1, it must always set PositionIdentifier in
// the PositionUpdate message to the identifier of the Trade Position.
//
// When the Client checks that this is set to 1, then it knows that it can
// expect there potentially can be more than one Trade Position for a specific
// Symbol and Trade Account being reported by the PositionUpdate messages.
// The Client can then handle this appropriately.
func (m LogonResponseFixedPointerBuilder) SetUsesMultiplePositionsPerSymbolAndTradeAccount(value uint8) {
	message.SetUint8Fixed(m.Ptr, 252, 251, value)
}

// SetMarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseFixedBuilder) SetMarketDataSupported(value bool) {
	message.SetBoolFixed(m.Unsafe(), 253, 252, value)
}

// SetMarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseFixedPointerBuilder) SetMarketDataSupported(value bool) {
	message.SetBoolFixed(m.Ptr, 253, 252, value)
}

// Copy
func (m LogonResponse) Copy(to LogonResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// Copy
func (m LogonResponseBuilder) Copy(to LogonResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyPointer
func (m LogonResponse) CopyPointer(to LogonResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyPointer
func (m LogonResponseBuilder) CopyPointer(to LogonResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyToPointer
func (m LogonResponse) CopyToPointer(to LogonResponseFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyToPointer
func (m LogonResponseBuilder) CopyToPointer(to LogonResponseFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyTo
func (m LogonResponse) CopyTo(to LogonResponseFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyTo
func (m LogonResponseBuilder) CopyTo(to LogonResponseFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// Copy
func (m LogonResponseFixed) Copy(to LogonResponseFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// Copy
func (m LogonResponseFixedBuilder) Copy(to LogonResponseFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyPointer
func (m LogonResponseFixed) CopyPointer(to LogonResponseFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyPointer
func (m LogonResponseFixedBuilder) CopyPointer(to LogonResponseFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyToPointer
func (m LogonResponseFixed) CopyToPointer(to LogonResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyToPointer
func (m LogonResponseFixedBuilder) CopyToPointer(to LogonResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyTo
func (m LogonResponseFixed) CopyTo(to LogonResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyTo
func (m LogonResponseFixedBuilder) CopyTo(to LogonResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// Copy
func (m LogonResponsePointer) Copy(to LogonResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// Copy
func (m LogonResponsePointerBuilder) Copy(to LogonResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyPointer
func (m LogonResponsePointer) CopyPointer(to LogonResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyPointer
func (m LogonResponsePointerBuilder) CopyPointer(to LogonResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyTo
func (m LogonResponsePointer) CopyTo(to LogonResponseFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyTo
func (m LogonResponsePointerBuilder) CopyTo(to LogonResponseFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyToPointer
func (m LogonResponsePointer) CopyToPointer(to LogonResponseFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyToPointer
func (m LogonResponsePointerBuilder) CopyToPointer(to LogonResponseFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// Copy
func (m LogonResponseFixedPointer) Copy(to LogonResponseFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// Copy
func (m LogonResponseFixedPointerBuilder) Copy(to LogonResponseFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyPointer
func (m LogonResponseFixedPointer) CopyPointer(to LogonResponseFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyPointer
func (m LogonResponseFixedPointerBuilder) CopyPointer(to LogonResponseFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyTo
func (m LogonResponseFixedPointer) CopyTo(to LogonResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyTo
func (m LogonResponseFixedPointerBuilder) CopyTo(to LogonResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyToPointer
func (m LogonResponseFixedPointer) CopyToPointer(to LogonResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}

// CopyToPointer
func (m LogonResponseFixedPointerBuilder) CopyToPointer(to LogonResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetResult(m.Result())
	to.SetResultText(m.ResultText())
	to.SetReconnectAddress(m.ReconnectAddress())
	to.SetInteger_1(m.Integer_1())
	to.SetServerName(m.ServerName())
	to.SetMarketDepthUpdatesBestBidAndAsk(m.MarketDepthUpdatesBestBidAndAsk())
	to.SetTradingIsSupported(m.TradingIsSupported())
	to.SetOCOOrdersSupported(m.OCOOrdersSupported())
	to.SetOrderCancelReplaceSupported(m.OrderCancelReplaceSupported())
	to.SetSymbolExchangeDelimiter(m.SymbolExchangeDelimiter())
	to.SetSecurityDefinitionsSupported(m.SecurityDefinitionsSupported())
	to.SetHistoricalPriceDataSupported(m.HistoricalPriceDataSupported())
	to.SetResubscribeWhenMarketDataFeedAvailable(m.ResubscribeWhenMarketDataFeedAvailable())
	to.SetMarketDepthIsSupported(m.MarketDepthIsSupported())
	to.SetOneHistoricalPriceDataRequestPerConnection(m.OneHistoricalPriceDataRequestPerConnection())
	to.SetBracketOrdersSupported(m.BracketOrdersSupported())
	to.SetUseIntegerPriceOrderMessages(m.UseIntegerPriceOrderMessages())
	to.SetUsesMultiplePositionsPerSymbolAndTradeAccount(m.UsesMultiplePositionsPerSymbolAndTradeAccount())
	to.SetMarketDataSupported(m.MarketDataSupported())
}
