package model

import (
	"github.com/moontrade/dtc-go/message"
)

const LogonRequestSize = 56

const LogonRequestFixedSize = 284

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
//     TradeMode                      = TRADE_MODE_UNSET  (0)
//     TradeAccount                   = ""
//     HardwareIdentifier             = ""
//     ClientName                     = ""
//     MarketDataTransmissionInterval = 0
// }
var _LogonRequestDefault = []byte{56, 0, 1, 0, 56, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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
//     TradeMode                      = TRADE_MODE_UNSET  (0)
//     TradeAccount                   = ""
//     HardwareIdentifier             = ""
//     ClientName                     = ""
//     MarketDataTransmissionInterval = 0
// }
var _LogonRequestFixedDefault = []byte{28, 1, 1, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

// LogonRequestFixedPointer The LogonRequest message is sent from the Client to the Server requesting
// to logon to the Server.
//
// This is the very first message the Client sends to the Server before being
// allowed to send any other message other than the EncodingRequest.
type LogonRequestFixedPointer struct {
	message.FixedPointer
}

// LogonRequestFixedPointerBuilder The LogonRequest message is sent from the Client to the Server requesting
// to logon to the Server.
//
// This is the very first message the Client sends to the Server before being
// allowed to send any other message other than the EncodingRequest.
type LogonRequestFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocLogonRequest() LogonRequestPointerBuilder {
	a := LogonRequestPointerBuilder{message.AllocVLSBuilder(56)}
	a.Ptr.SetBytes(0, _LogonRequestDefault)
	return a
}

func AllocLogonRequestFrom(b []byte) LogonRequestPointer {
	return LogonRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocLogonRequestFixed() LogonRequestFixedPointerBuilder {
	a := LogonRequestFixedPointerBuilder{message.AllocFixed(284)}
	a.Ptr.SetBytes(0, _LogonRequestFixedDefault)
	return a
}

func AllocLogonRequestFixedFrom(b []byte) LogonRequestFixedPointer {
	return LogonRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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
//     TradeMode                      = TRADE_MODE_UNSET  (0)
//     TradeAccount                   = ""
//     HardwareIdentifier             = ""
//     ClientName                     = ""
//     MarketDataTransmissionInterval = 0
// }
func (m LogonRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _LogonRequestDefault)
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
//     TradeMode                      = TRADE_MODE_UNSET  (0)
//     TradeAccount                   = ""
//     HardwareIdentifier             = ""
//     ClientName                     = ""
//     MarketDataTransmissionInterval = 0
// }
func (m LogonRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _LogonRequestFixedDefault)
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
//     TradeMode                      = TRADE_MODE_UNSET  (0)
//     TradeAccount                   = ""
//     HardwareIdentifier             = ""
//     ClientName                     = ""
//     MarketDataTransmissionInterval = 0
// }
func (m LogonRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _LogonRequestDefault)
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
//     TradeMode                      = TRADE_MODE_UNSET  (0)
//     TradeAccount                   = ""
//     HardwareIdentifier             = ""
//     ClientName                     = ""
//     MarketDataTransmissionInterval = 0
// }
func (m LogonRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _LogonRequestFixedDefault)
}

// ToBuilder
func (m LogonRequest) ToBuilder() LogonRequestBuilder {
	return LogonRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m LogonRequestFixed) ToBuilder() LogonRequestFixedBuilder {
	return LogonRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m LogonRequestPointer) ToBuilder() LogonRequestPointerBuilder {
	return LogonRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m LogonRequestFixedPointer) ToBuilder() LogonRequestFixedPointerBuilder {
	return LogonRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m LogonRequestBuilder) Finish() LogonRequest {
	return LogonRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m LogonRequestFixedBuilder) Finish() LogonRequestFixed {
	return LogonRequestFixed{m.Fixed}
}

// Finish
func (m *LogonRequestPointerBuilder) Finish() LogonRequestPointer {
	return LogonRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *LogonRequestFixedPointerBuilder) Finish() LogonRequestFixedPointer {
	return LogonRequestFixedPointer{m.FixedPointer}
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
func (m LogonRequest) Username() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequestBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
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
func (m LogonRequest) Password() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequestBuilder) Password() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
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
func (m LogonRequest) GeneralTextData() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// GeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestBuilder) GeneralTextData() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
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
func (m LogonRequest) Integer_1() int32 {
	return message.Int32VLS(m.Unsafe(), 28, 24)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequestBuilder) Integer_1() int32 {
	return message.Int32VLS(m.Unsafe(), 28, 24)
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
func (m LogonRequest) Integer_2() int32 {
	return message.Int32VLS(m.Unsafe(), 32, 28)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequestBuilder) Integer_2() int32 {
	return message.Int32VLS(m.Unsafe(), 32, 28)
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
func (m LogonRequest) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32VLS(m.Unsafe(), 40, 36))
}

// TradeMode
func (m LogonRequestBuilder) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32VLS(m.Unsafe(), 40, 36))
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
func (m LogonRequest) ClientName() string {
	return message.StringVLS(m.Unsafe(), 52, 48)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequestBuilder) ClientName() string {
	return message.StringVLS(m.Unsafe(), 52, 48)
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

// ProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestFixedPointer) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// ProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestFixedPointerBuilder) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequestFixed) Username() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequestFixedBuilder) Username() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequestFixedPointer) Username() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// Username Optional username for the server to authenticate the Client.
func (m LogonRequestFixedPointerBuilder) Username() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequestFixed) Password() string {
	return message.StringFixed(m.Unsafe(), 72, 40)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequestFixedBuilder) Password() string {
	return message.StringFixed(m.Unsafe(), 72, 40)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequestFixedPointer) Password() string {
	return message.StringFixed(m.Ptr, 72, 40)
}

// Password Optional password for the server to authenticate the Client.
func (m LogonRequestFixedPointerBuilder) Password() string {
	return message.StringFixed(m.Ptr, 72, 40)
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

// GeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestFixedPointer) GeneralTextData() string {
	return message.StringFixed(m.Ptr, 136, 72)
}

// GeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestFixedPointerBuilder) GeneralTextData() string {
	return message.StringFixed(m.Ptr, 136, 72)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequestFixed) Integer_1() int32 {
	return message.Int32Fixed(m.Unsafe(), 140, 136)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequestFixedBuilder) Integer_1() int32 {
	return message.Int32Fixed(m.Unsafe(), 140, 136)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequestFixedPointer) Integer_1() int32 {
	return message.Int32Fixed(m.Ptr, 140, 136)
}

// Integer_1 Optional. General-purpose integer.
func (m LogonRequestFixedPointerBuilder) Integer_1() int32 {
	return message.Int32Fixed(m.Ptr, 140, 136)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequestFixed) Integer_2() int32 {
	return message.Int32Fixed(m.Unsafe(), 144, 140)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequestFixedBuilder) Integer_2() int32 {
	return message.Int32Fixed(m.Unsafe(), 144, 140)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequestFixedPointer) Integer_2() int32 {
	return message.Int32Fixed(m.Ptr, 144, 140)
}

// Integer_2 Optional. General-purpose integer.
func (m LogonRequestFixedPointerBuilder) Integer_2() int32 {
	return message.Int32Fixed(m.Ptr, 144, 140)
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

// HeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestFixedPointer) HeartbeatIntervalInSeconds() int32 {
	return message.Int32Fixed(m.Ptr, 148, 144)
}

// HeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestFixedPointerBuilder) HeartbeatIntervalInSeconds() int32 {
	return message.Int32Fixed(m.Ptr, 148, 144)
}

// TradeMode
func (m LogonRequestFixed) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32Fixed(m.Unsafe(), 152, 148))
}

// TradeMode
func (m LogonRequestFixedBuilder) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32Fixed(m.Unsafe(), 152, 148))
}

// TradeMode
func (m LogonRequestFixedPointer) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32Fixed(m.Ptr, 152, 148))
}

// TradeMode
func (m LogonRequestFixedPointerBuilder) TradeMode() TradeModeEnum {
	return TradeModeEnum(message.Int32Fixed(m.Ptr, 152, 148))
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

// TradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequestFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 184, 152)
}

// TradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequestFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 184, 152)
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

// HardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequestFixedPointer) HardwareIdentifier() string {
	return message.StringFixed(m.Ptr, 248, 184)
}

// HardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequestFixedPointerBuilder) HardwareIdentifier() string {
	return message.StringFixed(m.Ptr, 248, 184)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequestFixed) ClientName() string {
	return message.StringFixed(m.Unsafe(), 280, 248)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequestFixedBuilder) ClientName() string {
	return message.StringFixed(m.Unsafe(), 280, 248)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequestFixedPointer) ClientName() string {
	return message.StringFixed(m.Ptr, 280, 248)
}

// ClientName The Client name. This is a free-form text string.
func (m LogonRequestFixedPointerBuilder) ClientName() string {
	return message.StringFixed(m.Ptr, 280, 248)
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

// MarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestFixedPointer) MarketDataTransmissionInterval() int32 {
	return message.Int32Fixed(m.Ptr, 284, 280)
}

// MarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestFixedPointerBuilder) MarketDataTransmissionInterval() int32 {
	return message.Int32Fixed(m.Ptr, 284, 280)
}

// SetProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestBuilder) SetProtocolVersion(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestPointerBuilder) SetProtocolVersion(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetUsername Optional username for the server to authenticate the Client.
func (m *LogonRequestBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetUsername Optional username for the server to authenticate the Client.
func (m *LogonRequestPointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetPassword Optional password for the server to authenticate the Client.
func (m *LogonRequestBuilder) SetPassword(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetPassword Optional password for the server to authenticate the Client.
func (m *LogonRequestPointerBuilder) SetPassword(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetGeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m *LogonRequestBuilder) SetGeneralTextData(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetGeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m *LogonRequestPointerBuilder) SetGeneralTextData(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetInteger_1 Optional. General-purpose integer.
func (m LogonRequestBuilder) SetInteger_1(value int32) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, value)
}

// SetInteger_1 Optional. General-purpose integer.
func (m LogonRequestPointerBuilder) SetInteger_1(value int32) {
	message.SetInt32VLS(m.Ptr, 28, 24, value)
}

// SetInteger_2 Optional. General-purpose integer.
func (m LogonRequestBuilder) SetInteger_2(value int32) {
	message.SetInt32VLS(m.Unsafe(), 32, 28, value)
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
func (m LogonRequestBuilder) SetHeartbeatIntervalInSeconds(value int32) {
	message.SetInt32VLS(m.Unsafe(), 36, 32, value)
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
func (m LogonRequestBuilder) SetTradeMode(value TradeModeEnum) {
	message.SetInt32VLS(m.Unsafe(), 40, 36, int32(value))
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
func (m *LogonRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 44, 40, value)
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
func (m *LogonRequestBuilder) SetHardwareIdentifier(value string) {
	message.SetStringVLS(&m.VLSBuilder, 48, 44, value)
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
func (m *LogonRequestBuilder) SetClientName(value string) {
	message.SetStringVLS(&m.VLSBuilder, 52, 48, value)
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
func (m LogonRequestBuilder) SetMarketDataTransmissionInterval(value int32) {
	message.SetInt32VLS(m.Unsafe(), 56, 52, value)
}

// SetMarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestPointerBuilder) SetMarketDataTransmissionInterval(value int32) {
	message.SetInt32VLS(m.Ptr, 56, 52, value)
}

// SetProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestFixedBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m LogonRequestFixedPointerBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetUsername Optional username for the server to authenticate the Client.
func (m LogonRequestFixedBuilder) SetUsername(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetUsername Optional username for the server to authenticate the Client.
func (m LogonRequestFixedPointerBuilder) SetUsername(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetPassword Optional password for the server to authenticate the Client.
func (m LogonRequestFixedBuilder) SetPassword(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 40, value)
}

// SetPassword Optional password for the server to authenticate the Client.
func (m LogonRequestFixedPointerBuilder) SetPassword(value string) {
	message.SetStringFixed(m.Ptr, 72, 40, value)
}

// SetGeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestFixedBuilder) SetGeneralTextData(value string) {
	message.SetStringFixed(m.Unsafe(), 136, 72, value)
}

// SetGeneralTextData Optional general-purpose text string. For example, this could be used
// to pass a license key that the Server may require.
func (m LogonRequestFixedPointerBuilder) SetGeneralTextData(value string) {
	message.SetStringFixed(m.Ptr, 136, 72, value)
}

// SetInteger_1 Optional. General-purpose integer.
func (m LogonRequestFixedBuilder) SetInteger_1(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 140, 136, value)
}

// SetInteger_1 Optional. General-purpose integer.
func (m LogonRequestFixedPointerBuilder) SetInteger_1(value int32) {
	message.SetInt32Fixed(m.Ptr, 140, 136, value)
}

// SetInteger_2 Optional. General-purpose integer.
func (m LogonRequestFixedBuilder) SetInteger_2(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 144, 140, value)
}

// SetInteger_2 Optional. General-purpose integer.
func (m LogonRequestFixedPointerBuilder) SetInteger_2(value int32) {
	message.SetInt32Fixed(m.Ptr, 144, 140, value)
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

// SetHeartbeatIntervalInSeconds The interval in seconds that each side, the Client and the Server, needs
// to use to send Heartbeat messages to the other side.
//
// This should be a value from anywhere from 5 to 60 seconds.
//
// This field is required.
func (m LogonRequestFixedPointerBuilder) SetHeartbeatIntervalInSeconds(value int32) {
	message.SetInt32Fixed(m.Ptr, 148, 144, value)
}

// SetTradeMode
func (m LogonRequestFixedBuilder) SetTradeMode(value TradeModeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 152, 148, int32(value))
}

// SetTradeMode
func (m LogonRequestFixedPointerBuilder) SetTradeMode(value TradeModeEnum) {
	message.SetInt32Fixed(m.Ptr, 152, 148, int32(value))
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

// SetTradeAccount This is an optional field and this should only be set to a Trade Account
// identifier if that is required to logon by the Server. this would only
// be implemented in rare cases. Usually this would be the case if the logon
// is bound to a particular Trade Account and not changeable after the log
// in.
//
// The server is still required to implement the TradeAccountsRequest and
// TradeAccountResponsemessages.
func (m LogonRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 184, 152, value)
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

// SetHardwareIdentifier Optional: This is the computer hardware identifier. The intention of this
// is that this will be implemented by the Client program developer on a
// case-by-case basis for specific Data/Trading service providers. It will
// be a reasonable implementation to uniquely identify a system and will
// not be publicly disclosed. It will never contain personally identifiable
// information.
func (m LogonRequestFixedPointerBuilder) SetHardwareIdentifier(value string) {
	message.SetStringFixed(m.Ptr, 248, 184, value)
}

// SetClientName The Client name. This is a free-form text string.
func (m LogonRequestFixedBuilder) SetClientName(value string) {
	message.SetStringFixed(m.Unsafe(), 280, 248, value)
}

// SetClientName The Client name. This is a free-form text string.
func (m LogonRequestFixedPointerBuilder) SetClientName(value string) {
	message.SetStringFixed(m.Ptr, 280, 248, value)
}

// SetMarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestFixedBuilder) SetMarketDataTransmissionInterval(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 284, 280, value)
}

// SetMarketDataTransmissionInterval This is an optional field to be used by the Server which specifies in
// milliseconds, the delay with transmitting market data to the Client.
//
// For reasons of efficiency, the server may buffer data over this timeframe,
// and send data after this time frame expires.
func (m LogonRequestFixedPointerBuilder) SetMarketDataTransmissionInterval(value int32) {
	message.SetInt32Fixed(m.Ptr, 284, 280, value)
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

// Copy
func (m LogonRequestFixedPointer) Copy(to LogonRequestFixedBuilder) {
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
func (m LogonRequestFixedPointerBuilder) Copy(to LogonRequestFixedBuilder) {
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
func (m LogonRequestFixedPointer) CopyPointer(to LogonRequestFixedPointerBuilder) {
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
func (m LogonRequestFixedPointerBuilder) CopyPointer(to LogonRequestFixedPointerBuilder) {
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
func (m LogonRequestFixedPointer) CopyTo(to LogonRequestBuilder) {
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
func (m LogonRequestFixedPointerBuilder) CopyTo(to LogonRequestBuilder) {
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
func (m LogonRequestFixedPointer) CopyToPointer(to LogonRequestPointerBuilder) {
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
func (m LogonRequestFixedPointerBuilder) CopyToPointer(to LogonRequestPointerBuilder) {
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
