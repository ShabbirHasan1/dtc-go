package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                           = LogonRequestFixedSize  (284)
//     Type                           = LOGON_REQUEST  (1)
//     ProtocolVersion                = CURRENT_VERSION  (8)
//     Username                       = ""
//     Password                       = ""
//     GeneralTextData                = ""
//     Integer_1                      = 0
//     Integer_2                      = 0
//     HeartbeatIntervalInSeconds     = 0
//     TradeMode                      = 0
//     TradeAccount                   = ""
//     HardwareIdentifier             = ""
//     ClientName                     = ""
//     MarketDataTransmissionInterval = 0
// }
var _LogonRequestFixedDefault = []byte{28, 1, 1, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const LogonRequestFixedSize = 284

// LogonRequestFixed The LogonRequest message is sent from the Client to the Server requesting
// to logon to the Server.
//
// This is the very first message the Client sends to the Server before being
// allowed to send any other message other than the EncodingRequest.
type LogonRequestFixed struct {
	message.Fixed
}

// LogonRequestFixedBuilder The LogonRequest message is sent from the Client to the Server requesting
// to logon to the Server.
//
// This is the very first message the Client sends to the Server before being
// allowed to send any other message other than the EncodingRequest.
type LogonRequestFixedBuilder struct {
	message.Fixed
}

func NewLogonRequestFixedFrom(b []byte) LogonRequestFixed {
	return LogonRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapLogonRequestFixed(b []byte) LogonRequestFixed {
	return LogonRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewLogonRequestFixed() LogonRequestFixedBuilder {
	a := LogonRequestFixedBuilder{message.NewFixed(284)}
	a.Unsafe().SetBytes(0, _LogonRequestFixedDefault)
	return a
}

// Clear
// {
//     Size                           = LogonRequestFixedSize  (284)
//     Type                           = LOGON_REQUEST  (1)
//     ProtocolVersion                = CURRENT_VERSION  (8)
//     Username                       = ""
//     Password                       = ""
//     GeneralTextData                = ""
//     Integer_1                      = 0
//     Integer_2                      = 0
//     HeartbeatIntervalInSeconds     = 0
//     TradeMode                      = 0
//     TradeAccount                   = ""
//     HardwareIdentifier             = ""
//     ClientName                     = ""
//     MarketDataTransmissionInterval = 0
// }
func (m LogonRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _LogonRequestFixedDefault)
}

// ToBuilder
func (m LogonRequestFixed) ToBuilder() LogonRequestFixedBuilder {
	return LogonRequestFixedBuilder{m.Fixed}
}

// Finish
func (m LogonRequestFixedBuilder) Finish() LogonRequestFixed {
	return LogonRequestFixed{m.Fixed}
}

// ProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestFixed) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// ProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestFixedBuilder) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequestFixed) Username() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequestFixedBuilder) Username() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequestFixed) Password() string {
	return message.StringFixed(m.Unsafe(), 72, 40)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequestFixedBuilder) Password() string {
	return message.StringFixed(m.Unsafe(), 72, 40)
}

// GeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestFixed) GeneralTextData() string {
	return message.StringFixed(m.Unsafe(), 136, 72)
}

// GeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestFixedBuilder) GeneralTextData() string {
	return message.StringFixed(m.Unsafe(), 136, 72)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequestFixed) Integer_1() int32 {
	return message.Int32Fixed(m.Unsafe(), 140, 136)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequestFixedBuilder) Integer_1() int32 {
	return message.Int32Fixed(m.Unsafe(), 140, 136)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequestFixed) Integer_2() int32 {
	return message.Int32Fixed(m.Unsafe(), 144, 140)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequestFixedBuilder) Integer_2() int32 {
	return message.Int32Fixed(m.Unsafe(), 144, 140)
}

// HeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestFixed) HeartbeatIntervalInSeconds() int32 {
	return message.Int32Fixed(m.Unsafe(), 148, 144)
}

// HeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestFixedBuilder) HeartbeatIntervalInSeconds() int32 {
	return message.Int32Fixed(m.Unsafe(), 148, 144)
}

// TradeMode
func (m LogonRequestFixed) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32Fixed(m.Unsafe(), 152, 148))
}

// TradeMode
func (m LogonRequestFixedBuilder) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32Fixed(m.Unsafe(), 152, 148))
}

// TradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequestFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 184, 152)
}

// TradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequestFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 184, 152)
}

// HardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequestFixed) HardwareIdentifier() string {
	return message.StringFixed(m.Unsafe(), 248, 184)
}

// HardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequestFixedBuilder) HardwareIdentifier() string {
	return message.StringFixed(m.Unsafe(), 248, 184)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequestFixed) ClientName() string {
	return message.StringFixed(m.Unsafe(), 280, 248)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequestFixedBuilder) ClientName() string {
	return message.StringFixed(m.Unsafe(), 280, 248)
}

// MarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestFixed) MarketDataTransmissionInterval() int32 {
	return message.Int32Fixed(m.Unsafe(), 284, 280)
}

// MarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestFixedBuilder) MarketDataTransmissionInterval() int32 {
	return message.Int32Fixed(m.Unsafe(), 284, 280)
}

// SetProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestFixedBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetUsername Optional username for the server to authenticate the Client.
func (m LogonRequestFixedBuilder) SetUsername(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetPassword Optional password for the server to authenticate the Client.
func (m LogonRequestFixedBuilder) SetPassword(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 40, value)
}

// SetGeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestFixedBuilder) SetGeneralTextData(value string) {
	message.SetStringFixed(m.Unsafe(), 136, 72, value)
}

// SetInteger_1 Optional. General-purpose integer.
func (m LogonRequestFixedBuilder) SetInteger_1(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 140, 136, value)
}

// SetInteger_2 Optional. General-purpose integer.
func (m LogonRequestFixedBuilder) SetInteger_2(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 144, 140, value)
}

// SetHeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestFixedBuilder) SetHeartbeatIntervalInSeconds(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 148, 144, value)
}

// SetTradeMode
func (m LogonRequestFixedBuilder) SetTradeMode(value TradeModeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 152, 148, int32(value))
}

// SetTradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 184, 152, value)
}

// SetHardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequestFixedBuilder) SetHardwareIdentifier(value string) {
	message.SetStringFixed(m.Unsafe(), 248, 184, value)
}

// SetClientName The Client name. This is a free-form text string.
func (m LogonRequestFixedBuilder) SetClientName(value string) {
	message.SetStringFixed(m.Unsafe(), 280, 248, value)
}

// SetMarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestFixedBuilder) SetMarketDataTransmissionInterval(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 284, 280, value)
}

// Copy
func (m LogonRequestFixed) Copy(to LogonRequestFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// Copy
func (m LogonRequestFixedBuilder) Copy(to LogonRequestFixedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// CopyPointer
func (m LogonRequestFixed) CopyPointer(to LogonRequestFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// CopyPointer
func (m LogonRequestFixedBuilder) CopyPointer(to LogonRequestFixedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// CopyToPointer
func (m LogonRequestFixed) CopyToPointer(to LogonRequestPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// CopyToPointer
func (m LogonRequestFixedBuilder) CopyToPointer(to LogonRequestPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// CopyTo
func (m LogonRequestFixed) CopyTo(to LogonRequestBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}

// CopyTo
func (m LogonRequestFixedBuilder) CopyTo(to LogonRequestBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetUsername(m.Username())
	to.SetPassword(m.Password())
	to.SetGeneralTextData(m.GeneralTextData())
	to.SetInteger_1(m.Integer_1())
	to.SetInteger_2(m.Integer_2())
	to.SetHeartbeatIntervalInSeconds(m.HeartbeatIntervalInSeconds())
	to.SetTradeMode(m.TradeMode())
	to.SetTradeAccount(m.TradeAccount())
	to.SetHardwareIdentifier(m.HardwareIdentifier())
	to.SetClientName(m.ClientName())
	to.SetMarketDataTransmissionInterval(m.MarketDataTransmissionInterval())
}
