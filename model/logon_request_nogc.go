package model

import (
	"github.com/moontrade/dtc-go/message"
)

// LogonRequestPointer The LogonRequest message is sent from the Client to the Server requesting
// to logon to the Server.
//
// This is the very first message the Client sends to the Server before being
// allowed to send any other message other than the EncodingRequest.
type LogonRequestPointer struct {
	message.VLSPointer
}

// LogonRequestPointerBuilder The LogonRequest message is sent from the Client to the Server requesting
// to logon to the Server.
//
// This is the very first message the Client sends to the Server before being
// allowed to send any other message other than the EncodingRequest.
type LogonRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocLogonRequest() LogonRequestPointerBuilder {
	a := LogonRequestPointerBuilder{message.AllocVLSBuilder(56)}
	a.Ptr.SetBytes(0, _LogonRequestDefault)
	return a
}

func AllocLogonRequestFrom(b []byte) LogonRequestPointer {
	return LogonRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m LogonRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _LogonRequestDefault)
}

// ToBuilder
func (m LogonRequestPointer) ToBuilder() LogonRequestPointerBuilder {
	return LogonRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *LogonRequestPointerBuilder) Finish() LogonRequestPointer {
	return LogonRequestPointer{m.VLSPointerBuilder.Finish()}
}

// ProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestPointer) ProtocolVersion() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// ProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestPointerBuilder) ProtocolVersion() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequestPointer) Username() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequestPointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequestPointer) Password() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequestPointerBuilder) Password() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// GeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestPointer) GeneralTextData() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// GeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestPointerBuilder) GeneralTextData() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequestPointer) Integer_1() int32 {
	return message.Int32VLS(m.Ptr, 28, 24)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequestPointerBuilder) Integer_1() int32 {
	return message.Int32VLS(m.Ptr, 28, 24)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequestPointer) Integer_2() int32 {
	return message.Int32VLS(m.Ptr, 32, 28)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequestPointerBuilder) Integer_2() int32 {
	return message.Int32VLS(m.Ptr, 32, 28)
}

// HeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestPointer) HeartbeatIntervalInSeconds() int32 {
	return message.Int32VLS(m.Ptr, 36, 32)
}

// HeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestPointerBuilder) HeartbeatIntervalInSeconds() int32 {
	return message.Int32VLS(m.Ptr, 36, 32)
}

// TradeMode
func (m LogonRequestPointer) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32VLS(m.Ptr, 40, 36))
}

// TradeMode
func (m LogonRequestPointerBuilder) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32VLS(m.Ptr, 40, 36))
}

// TradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 44, 40)
}

// TradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 44, 40)
}

// HardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequestPointer) HardwareIdentifier() string {
	return message.StringVLS(m.Ptr, 48, 44)
}

// HardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequestPointerBuilder) HardwareIdentifier() string {
	return message.StringVLS(m.Ptr, 48, 44)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequestPointer) ClientName() string {
	return message.StringVLS(m.Ptr, 52, 48)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequestPointerBuilder) ClientName() string {
	return message.StringVLS(m.Ptr, 52, 48)
}

// MarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestPointer) MarketDataTransmissionInterval() int32 {
	return message.Int32VLS(m.Ptr, 56, 52)
}

// MarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestPointerBuilder) MarketDataTransmissionInterval() int32 {
	return message.Int32VLS(m.Ptr, 56, 52)
}

// SetProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestPointerBuilder) SetProtocolVersion(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetUsername Optional username for the server to authenticate the Client.
func (m *LogonRequestPointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetPassword Optional password for the server to authenticate the Client.
func (m *LogonRequestPointerBuilder) SetPassword(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetGeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m *LogonRequestPointerBuilder) SetGeneralTextData(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetInteger_1 Optional. General-purpose integer.
func (m LogonRequestPointerBuilder) SetInteger_1(value int32) {
	message.SetInt32VLS(m.Ptr, 28, 24, value)
}

// SetInteger_2 Optional. General-purpose integer.
func (m LogonRequestPointerBuilder) SetInteger_2(value int32) {
	message.SetInt32VLS(m.Ptr, 32, 28, value)
}

// SetHeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestPointerBuilder) SetHeartbeatIntervalInSeconds(value int32) {
	message.SetInt32VLS(m.Ptr, 36, 32, value)
}

// SetTradeMode
func (m LogonRequestPointerBuilder) SetTradeMode(value TradeModeEnum) {
	message.SetInt32VLS(m.Ptr, 40, 36, int32(value))
}

// SetTradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m *LogonRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 44, 40, value)
}

// SetHardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m *LogonRequestPointerBuilder) SetHardwareIdentifier(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 48, 44, value)
}

// SetClientName The Client name. This is a free-form text string.
func (m *LogonRequestPointerBuilder) SetClientName(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 52, 48, value)
}

// SetMarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestPointerBuilder) SetMarketDataTransmissionInterval(value int32) {
	message.SetInt32VLS(m.Ptr, 56, 52, value)
}

// Copy
func (m LogonRequestPointer) Copy(to LogonRequestBuilder) {
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
func (m LogonRequestPointerBuilder) Copy(to LogonRequestBuilder) {
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
func (m LogonRequestPointer) CopyPointer(to LogonRequestPointerBuilder) {
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
func (m LogonRequestPointerBuilder) CopyPointer(to LogonRequestPointerBuilder) {
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
func (m LogonRequestPointer) CopyTo(to LogonRequestFixedBuilder) {
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
func (m LogonRequestPointerBuilder) CopyTo(to LogonRequestFixedBuilder) {
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
func (m LogonRequestPointer) CopyToPointer(to LogonRequestFixedPointerBuilder) {
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
func (m LogonRequestPointerBuilder) CopyToPointer(to LogonRequestFixedPointerBuilder) {
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
