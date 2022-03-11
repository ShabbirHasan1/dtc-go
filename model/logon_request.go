package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                           = LogonRequestSize  (56)
//     Type                           = LOGON_REQUEST  (1)
//     BaseSize                       = LogonRequestSize  (56)
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
var _LogonRequestDefault = []byte{56, 0, 1, 0, 56, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const LogonRequestSize = 56

// LogonRequest The LogonRequest message is sent from the Client to the Server requesting
// to logon to the Server.
//
// This is the very first message the Client sends to the Server before being
// allowed to send any other message other than the EncodingRequest.
type LogonRequest struct {
	message.VLS
}

// LogonRequestBuilder The LogonRequest message is sent from the Client to the Server requesting
// to logon to the Server.
//
// This is the very first message the Client sends to the Server before being
// allowed to send any other message other than the EncodingRequest.
type LogonRequestBuilder struct {
	message.VLSBuilder
}

func NewLogonRequestFrom(b []byte) LogonRequest {
	return LogonRequest{VLS: message.NewVLSFrom(b)}
}

func WrapLogonRequest(b []byte) LogonRequest {
	return LogonRequest{VLS: message.WrapVLS(b)}
}

func NewLogonRequest() LogonRequestBuilder {
	a := LogonRequestBuilder{message.NewVLSBuilder(56)}
	a.Unsafe().SetBytes(0, _LogonRequestDefault)
	return a
}

// Clear
// {
//     Size                           = LogonRequestSize  (56)
//     Type                           = LOGON_REQUEST  (1)
//     BaseSize                       = LogonRequestSize  (56)
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
func (m LogonRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _LogonRequestDefault)
}

// ToBuilder
func (m LogonRequest) ToBuilder() LogonRequestBuilder {
	return LogonRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m LogonRequestBuilder) Finish() LogonRequest {
	return LogonRequest{m.VLSBuilder.Finish()}
}

// ProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequest) ProtocolVersion() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// ProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestBuilder) ProtocolVersion() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequest) Username() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequestBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequest) Password() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequestBuilder) Password() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// GeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequest) GeneralTextData() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// GeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestBuilder) GeneralTextData() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequest) Integer_1() int32 {
	return message.Int32VLS(m.Unsafe(), 28, 24)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequestBuilder) Integer_1() int32 {
	return message.Int32VLS(m.Unsafe(), 28, 24)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequest) Integer_2() int32 {
	return message.Int32VLS(m.Unsafe(), 32, 28)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequestBuilder) Integer_2() int32 {
	return message.Int32VLS(m.Unsafe(), 32, 28)
}

// HeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequest) HeartbeatIntervalInSeconds() int32 {
	return message.Int32VLS(m.Unsafe(), 36, 32)
}

// HeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestBuilder) HeartbeatIntervalInSeconds() int32 {
	return message.Int32VLS(m.Unsafe(), 36, 32)
}

// TradeMode
func (m LogonRequest) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32VLS(m.Unsafe(), 40, 36))
}

// TradeMode
func (m LogonRequestBuilder) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32VLS(m.Unsafe(), 40, 36))
}

// TradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 44, 40)
}

// TradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 44, 40)
}

// HardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequest) HardwareIdentifier() string {
	return message.StringVLS(m.Unsafe(), 48, 44)
}

// HardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequestBuilder) HardwareIdentifier() string {
	return message.StringVLS(m.Unsafe(), 48, 44)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequest) ClientName() string {
	return message.StringVLS(m.Unsafe(), 52, 48)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequestBuilder) ClientName() string {
	return message.StringVLS(m.Unsafe(), 52, 48)
}

// MarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequest) MarketDataTransmissionInterval() int32 {
	return message.Int32VLS(m.Unsafe(), 56, 52)
}

// MarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestBuilder) MarketDataTransmissionInterval() int32 {
	return message.Int32VLS(m.Unsafe(), 56, 52)
}

// SetProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestBuilder) SetProtocolVersion(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetUsername Optional username for the server to authenticate the Client.
func (m *LogonRequestBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetPassword Optional password for the server to authenticate the Client.
func (m *LogonRequestBuilder) SetPassword(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetGeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m *LogonRequestBuilder) SetGeneralTextData(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetInteger_1 Optional. General-purpose integer.
func (m LogonRequestBuilder) SetInteger_1(value int32) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, value)
}

// SetInteger_2 Optional. General-purpose integer.
func (m LogonRequestBuilder) SetInteger_2(value int32) {
	message.SetInt32VLS(m.Unsafe(), 32, 28, value)
}

// SetHeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestBuilder) SetHeartbeatIntervalInSeconds(value int32) {
	message.SetInt32VLS(m.Unsafe(), 36, 32, value)
}

// SetTradeMode
func (m LogonRequestBuilder) SetTradeMode(value TradeModeEnum) {
	message.SetInt32VLS(m.Unsafe(), 40, 36, int32(value))
}

// SetTradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m *LogonRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 44, 40, value)
}

// SetHardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m *LogonRequestBuilder) SetHardwareIdentifier(value string) {
	message.SetStringVLS(&m.VLSBuilder, 48, 44, value)
}

// SetClientName The Client name. This is a free-form text string.
func (m *LogonRequestBuilder) SetClientName(value string) {
	message.SetStringVLS(&m.VLSBuilder, 52, 48, value)
}

// SetMarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestBuilder) SetMarketDataTransmissionInterval(value int32) {
	message.SetInt32VLS(m.Unsafe(), 56, 52, value)
}

// Copy
func (m LogonRequest) Copy(to LogonRequestBuilder) {
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
func (m LogonRequestBuilder) Copy(to LogonRequestBuilder) {
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
func (m LogonRequest) CopyPointer(to LogonRequestPointerBuilder) {
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
func (m LogonRequestBuilder) CopyPointer(to LogonRequestPointerBuilder) {
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
func (m LogonRequest) CopyToPointer(to LogonRequestFixedPointerBuilder) {
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
func (m LogonRequestBuilder) CopyToPointer(to LogonRequestFixedPointerBuilder) {
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
func (m LogonRequest) CopyTo(to LogonRequestFixedBuilder) {
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
func (m LogonRequestBuilder) CopyTo(to LogonRequestFixedBuilder) {
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
