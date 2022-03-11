package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
//     TradingIsSupported                            = 0
//     OCOOrdersSupported                            = 0
//     OrderCancelReplaceSupported                   = 1
//     SymbolExchangeDelimiter                       = ""
//     SecurityDefinitionsSupported                  = 0
//     HistoricalPriceDataSupported                  = 0
//     ResubscribeWhenMarketDataFeedAvailable        = 0
//     MarketDepthIsSupported                        = 1
//     OneHistoricalPriceDataRequestPerConnection    = 0
//     BracketOrdersSupported                        = 0
//     UseIntegerPriceOrderMessages                  = 0
//     UsesMultiplePositionsPerSymbolAndTradeAccount = 0
//     MarketDataSupported                           = 1
// }
var _LogonResponseFixedDefault = []byte{0, 1, 2, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0}

const LogonResponseFixedSize = 256

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
//     TradingIsSupported                            = 0
//     OCOOrdersSupported                            = 0
//     OrderCancelReplaceSupported                   = 1
//     SymbolExchangeDelimiter                       = ""
//     SecurityDefinitionsSupported                  = 0
//     HistoricalPriceDataSupported                  = 0
//     ResubscribeWhenMarketDataFeedAvailable        = 0
//     MarketDepthIsSupported                        = 1
//     OneHistoricalPriceDataRequestPerConnection    = 0
//     BracketOrdersSupported                        = 0
//     UseIntegerPriceOrderMessages                  = 0
//     UsesMultiplePositionsPerSymbolAndTradeAccount = 0
//     MarketDataSupported                           = 1
// }
func (m LogonResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _LogonResponseFixedDefault)
}

// ToBuilder
func (m LogonResponseFixed) ToBuilder() LogonResponseFixedBuilder {
	return LogonResponseFixedBuilder{m.Fixed}
}

// Finish
func (m LogonResponseFixedBuilder) Finish() LogonResponseFixed {
	return LogonResponseFixed{m.Fixed}
}

// ProtocolVersion This is automatically set by the constructor.
func (m LogonResponseFixed) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// ProtocolVersion This is automatically set by the constructor.
func (m LogonResponseFixedBuilder) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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

// TradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseFixed) TradingIsSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 238, 237)
}

// TradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseFixedBuilder) TradingIsSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 238, 237)
}

// OCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseFixed) OCOOrdersSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 239, 238)
}

// OCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseFixedBuilder) OCOOrdersSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 239, 238)
}

// OrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseFixed) OrderCancelReplaceSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 240, 239)
}

// OrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseFixedBuilder) OrderCancelReplaceSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 240, 239)
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

// SecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseFixed) SecurityDefinitionsSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 245, 244)
}

// SecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseFixedBuilder) SecurityDefinitionsSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 245, 244)
}

// HistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseFixed) HistoricalPriceDataSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 246, 245)
}

// HistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseFixedBuilder) HistoricalPriceDataSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 246, 245)
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

// MarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseFixed) MarketDepthIsSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 248, 247)
}

// MarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseFixedBuilder) MarketDepthIsSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 248, 247)
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

// BracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseFixed) BracketOrdersSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 250, 249)
}

// BracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseFixedBuilder) BracketOrdersSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 250, 249)
}

// UseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseFixed) UseIntegerPriceOrderMessages() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 251, 250)
}

// UseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseFixedBuilder) UseIntegerPriceOrderMessages() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 251, 250)
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

// MarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseFixed) MarketDataSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 253, 252)
}

// MarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseFixedBuilder) MarketDataSupported() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 253, 252)
}

// SetProtocolVersion This is automatically set by the constructor.
func (m LogonResponseFixedBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
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

// SetResultText Optional freeform text to provide information related to a successful
// or unsuccessful logon. The Client will display this text to the user.
func (m LogonResponseFixedBuilder) SetResultText(value string) {
	message.SetStringFixed(m.Unsafe(), 108, 12, value)
}

// SetReconnectAddress Server address/IP number and optional port number to reconnect to. Format:
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
// [Server Address:Port Number]. Only used if Result is set to LOGON_RECONNECT_NEW_ADDRESS.
func (m LogonResponseFixedBuilder) SetReconnectAddress(value string) {
	message.SetStringFixed(m.Unsafe(), 172, 108, value)
}

// SetInteger_1 Optional. General-purpose integer for the Server to communicate to the
// Client an integer value on logon.
func (m LogonResponseFixedBuilder) SetInteger_1(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 176, 172, value)
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

// SetMarketDepthUpdatesBestBidAndAsk Set this to 1 to indicate that the Server will only be sending market
// depth updates and not best bid and ask updates. The Client will use depth
// at level 1 to update the best bid and ask prices.
//
// Some Clients will maintain separate best bid and ask prices from market
// depth data.
func (m LogonResponseFixedBuilder) SetMarketDepthUpdatesBestBidAndAsk(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 237, 236, value)
}

// SetTradingIsSupported Set this to 1 to indicate the Server supports trading. Otherwise, the
// Client will not send through any trading messages.
func (m LogonResponseFixedBuilder) SetTradingIsSupported(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 238, 237, value)
}

// SetOCOOrdersSupported Set this to 1 to indicate the Server supports OCO orders.
func (m LogonResponseFixedBuilder) SetOCOOrdersSupported(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 239, 238, value)
}

// SetOrderCancelReplaceSupported Set this to 0 if Server does not support the CancelReplaceOrder message.
// Set this to 0 if Server does not support the CancelReplaceOrder message.
func (m LogonResponseFixedBuilder) SetOrderCancelReplaceSupported(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 240, 239, value)
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

// SetSecurityDefinitionsSupported Set to 1 if the Server supports Security Definition messages.
func (m LogonResponseFixedBuilder) SetSecurityDefinitionsSupported(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 245, 244, value)
}

// SetHistoricalPriceDataSupported Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
// Set this to 1 if the Server supports the HistoricalPriceDataRequest message.
func (m LogonResponseFixedBuilder) SetHistoricalPriceDataSupported(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 246, 245, value)
}

// SetResubscribeWhenMarketDataFeedAvailable Set this to 1, so that when the Client receives a MarketDataFeedStatus
// indicating the market data feed is restored, it will resubscribe to market
// data and market depth for all of the symbols it was previously tracking.
// data and market depth for all of the symbols it was previously tracking.
func (m LogonResponseFixedBuilder) SetResubscribeWhenMarketDataFeedAvailable(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 247, 246, value)
}

// SetMarketDepthIsSupported Set this to 1, if the Server supports the MarketDepthRequest message.
//
// The default is 0.
func (m LogonResponseFixedBuilder) SetMarketDepthIsSupported(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 248, 247, value)
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

// SetBracketOrdersSupported Set this to 1 to indicate the Server supports bracket orders.
func (m LogonResponseFixedBuilder) SetBracketOrdersSupported(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 250, 249, value)
}

// SetUseIntegerPriceOrderMessages With the integer trading messages discontinued as of August 2020, this
// field is no longer relevant.
func (m LogonResponseFixedBuilder) SetUseIntegerPriceOrderMessages(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 251, 250, value)
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

// SetMarketDataSupported Set this to 1, if the Server supports the MarketDataRequest message.
//
// The default is 1.
func (m LogonResponseFixedBuilder) SetMarketDataSupported(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 253, 252, value)
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
