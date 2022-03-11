package model

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeAccountResponseSize = 24

const TradeAccountResponseFixedSize = 48

// {
//     Size                = TradeAccountResponseSize  (24)
//     Type                = TRADE_ACCOUNT_RESPONSE  (401)
//     BaseSize            = TradeAccountResponseSize  (24)
//     TotalNumberMessages = 0
//     MessageNumber       = 0
//     TradeAccount        = ""
//     RequestID           = 0
// }
var _TradeAccountResponseDefault = []byte{24, 0, 145, 1, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size                = TradeAccountResponseFixedSize  (48)
//     Type                = TRADE_ACCOUNT_RESPONSE  (401)
//     TotalNumberMessages = 0
//     MessageNumber       = 0
//     TradeAccount        = ""
//     RequestID           = 0
// }
var _TradeAccountResponseFixedDefault = []byte{48, 0, 145, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// TradeAccountResponse This is a message from the Server to the Client in response to a TradeAccountsRequest
// message, providing a single trade account. There is one message for each
// trade account.
type TradeAccountResponse struct {
	message.VLS
}

// TradeAccountResponseBuilder This is a message from the Server to the Client in response to a TradeAccountsRequest
// message, providing a single trade account. There is one message for each
// trade account.
type TradeAccountResponseBuilder struct {
	message.VLSBuilder
}

// TradeAccountResponseFixed This is a message from the Server to the Client in response to a TradeAccountsRequest
// message, providing a single trade account. There is one message for each
// trade account.
type TradeAccountResponseFixed struct {
	message.Fixed
}

// TradeAccountResponseFixedBuilder This is a message from the Server to the Client in response to a TradeAccountsRequest
// message, providing a single trade account. There is one message for each
// trade account.
type TradeAccountResponseFixedBuilder struct {
	message.Fixed
}

// TradeAccountResponsePointer This is a message from the Server to the Client in response to a TradeAccountsRequest
// message, providing a single trade account. There is one message for each
// trade account.
type TradeAccountResponsePointer struct {
	message.VLSPointer
}

// TradeAccountResponsePointerBuilder This is a message from the Server to the Client in response to a TradeAccountsRequest
// message, providing a single trade account. There is one message for each
// trade account.
type TradeAccountResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

// TradeAccountResponseFixedPointer This is a message from the Server to the Client in response to a TradeAccountsRequest
// message, providing a single trade account. There is one message for each
// trade account.
type TradeAccountResponseFixedPointer struct {
	message.FixedPointer
}

// TradeAccountResponseFixedPointerBuilder This is a message from the Server to the Client in response to a TradeAccountsRequest
// message, providing a single trade account. There is one message for each
// trade account.
type TradeAccountResponseFixedPointerBuilder struct {
	message.FixedPointer
}

func NewTradeAccountResponseFrom(b []byte) TradeAccountResponse {
	return TradeAccountResponse{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountResponse(b []byte) TradeAccountResponse {
	return TradeAccountResponse{VLS: message.WrapVLS(b)}
}

func NewTradeAccountResponse() TradeAccountResponseBuilder {
	a := TradeAccountResponseBuilder{message.NewVLSBuilder(24)}
	a.Unsafe().SetBytes(0, _TradeAccountResponseDefault)
	return a
}

func NewTradeAccountResponseFixedFrom(b []byte) TradeAccountResponseFixed {
	return TradeAccountResponseFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapTradeAccountResponseFixed(b []byte) TradeAccountResponseFixed {
	return TradeAccountResponseFixed{Fixed: message.WrapFixed(b)}
}

func NewTradeAccountResponseFixed() TradeAccountResponseFixedBuilder {
	a := TradeAccountResponseFixedBuilder{message.NewFixed(48)}
	a.Unsafe().SetBytes(0, _TradeAccountResponseFixedDefault)
	return a
}

func AllocTradeAccountResponse() TradeAccountResponsePointerBuilder {
	a := TradeAccountResponsePointerBuilder{message.AllocVLSBuilder(24)}
	a.Ptr.SetBytes(0, _TradeAccountResponseDefault)
	return a
}

func AllocTradeAccountResponseFrom(b []byte) TradeAccountResponsePointer {
	return TradeAccountResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocTradeAccountResponseFixed() TradeAccountResponseFixedPointerBuilder {
	a := TradeAccountResponseFixedPointerBuilder{message.AllocFixed(48)}
	a.Ptr.SetBytes(0, _TradeAccountResponseFixedDefault)
	return a
}

func AllocTradeAccountResponseFixedFrom(b []byte) TradeAccountResponseFixedPointer {
	return TradeAccountResponseFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                = TradeAccountResponseSize  (24)
//     Type                = TRADE_ACCOUNT_RESPONSE  (401)
//     BaseSize            = TradeAccountResponseSize  (24)
//     TotalNumberMessages = 0
//     MessageNumber       = 0
//     TradeAccount        = ""
//     RequestID           = 0
// }
func (m TradeAccountResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountResponseDefault)
}

// Clear
// {
//     Size                = TradeAccountResponseFixedSize  (48)
//     Type                = TRADE_ACCOUNT_RESPONSE  (401)
//     TotalNumberMessages = 0
//     MessageNumber       = 0
//     TradeAccount        = ""
//     RequestID           = 0
// }
func (m TradeAccountResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountResponseFixedDefault)
}

// Clear
// {
//     Size                = TradeAccountResponseSize  (24)
//     Type                = TRADE_ACCOUNT_RESPONSE  (401)
//     BaseSize            = TradeAccountResponseSize  (24)
//     TotalNumberMessages = 0
//     MessageNumber       = 0
//     TradeAccount        = ""
//     RequestID           = 0
// }
func (m TradeAccountResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountResponseDefault)
}

// Clear
// {
//     Size                = TradeAccountResponseFixedSize  (48)
//     Type                = TRADE_ACCOUNT_RESPONSE  (401)
//     TotalNumberMessages = 0
//     MessageNumber       = 0
//     TradeAccount        = ""
//     RequestID           = 0
// }
func (m TradeAccountResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountResponseFixedDefault)
}

// ToBuilder
func (m TradeAccountResponse) ToBuilder() TradeAccountResponseBuilder {
	return TradeAccountResponseBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradeAccountResponseFixed) ToBuilder() TradeAccountResponseFixedBuilder {
	return TradeAccountResponseFixedBuilder{m.Fixed}
}

// ToBuilder
func (m TradeAccountResponsePointer) ToBuilder() TradeAccountResponsePointerBuilder {
	return TradeAccountResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m TradeAccountResponseFixedPointer) ToBuilder() TradeAccountResponseFixedPointerBuilder {
	return TradeAccountResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m TradeAccountResponseBuilder) Finish() TradeAccountResponse {
	return TradeAccountResponse{m.VLSBuilder.Finish()}
}

// Finish
func (m TradeAccountResponseFixedBuilder) Finish() TradeAccountResponseFixed {
	return TradeAccountResponseFixed{m.Fixed}
}

// Finish
func (m *TradeAccountResponsePointerBuilder) Finish() TradeAccountResponsePointer {
	return TradeAccountResponsePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *TradeAccountResponseFixedPointerBuilder) Finish() TradeAccountResponseFixedPointer {
	return TradeAccountResponseFixedPointer{m.FixedPointer}
}

// TotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponse) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// TotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseBuilder) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// TotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponsePointer) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// TotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponsePointerBuilder) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// MessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponse) MessageNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 16, 12)
}

// MessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseBuilder) MessageNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 16, 12)
}

// MessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponsePointer) MessageNumber() int32 {
	return message.Int32VLS(m.Ptr, 16, 12)
}

// MessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponsePointerBuilder) MessageNumber() int32 {
	return message.Int32VLS(m.Ptr, 16, 12)
}

// TradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// TradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// TradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponsePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// TradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponsePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// RequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponse) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 24, 20)
}

// RequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponseBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 24, 20)
}

// RequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponsePointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 24, 20)
}

// RequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponsePointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 24, 20)
}

// TotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixed) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// TotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixedBuilder) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// TotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixedPointer) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// TotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixedPointerBuilder) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// MessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixed) MessageNumber() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// MessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixedBuilder) MessageNumber() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// MessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixedPointer) MessageNumber() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// MessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixedPointerBuilder) MessageNumber() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// TradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponseFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 44, 12)
}

// TradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponseFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 44, 12)
}

// TradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponseFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 44, 12)
}

// TradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponseFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 44, 12)
}

// RequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponseFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 48, 44)
}

// RequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponseFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 48, 44)
}

// RequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponseFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 48, 44)
}

// RequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponseFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 48, 44)
}

// SetTotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetTotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponsePointerBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetMessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 16, 12, value)
}

// SetMessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponsePointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Ptr, 16, 12, value)
}

// SetTradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m *TradeAccountResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetTradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m *TradeAccountResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetRequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, value)
}

// SetRequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 24, 20, value)
}

// SetTotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixedBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetTotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixedPointerBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetMessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixedBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetMessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixedPointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetTradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponseFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 44, 12, value)
}

// SetTradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponseFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 44, 12, value)
}

// SetRequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponseFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 48, 44, value)
}

// SetRequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponseFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 48, 44, value)
}

// Copy
func (m TradeAccountResponse) Copy(to TradeAccountResponseBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// Copy
func (m TradeAccountResponseBuilder) Copy(to TradeAccountResponseBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m TradeAccountResponse) CopyPointer(to TradeAccountResponsePointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m TradeAccountResponseBuilder) CopyPointer(to TradeAccountResponsePointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyToPointer
func (m TradeAccountResponse) CopyToPointer(to TradeAccountResponseFixedPointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyToPointer
func (m TradeAccountResponseBuilder) CopyToPointer(to TradeAccountResponseFixedPointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyTo
func (m TradeAccountResponse) CopyTo(to TradeAccountResponseFixedBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyTo
func (m TradeAccountResponseBuilder) CopyTo(to TradeAccountResponseFixedBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// Copy
func (m TradeAccountResponseFixed) Copy(to TradeAccountResponseFixedBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// Copy
func (m TradeAccountResponseFixedBuilder) Copy(to TradeAccountResponseFixedBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m TradeAccountResponseFixed) CopyPointer(to TradeAccountResponseFixedPointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m TradeAccountResponseFixedBuilder) CopyPointer(to TradeAccountResponseFixedPointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyToPointer
func (m TradeAccountResponseFixed) CopyToPointer(to TradeAccountResponsePointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyToPointer
func (m TradeAccountResponseFixedBuilder) CopyToPointer(to TradeAccountResponsePointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyTo
func (m TradeAccountResponseFixed) CopyTo(to TradeAccountResponseBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyTo
func (m TradeAccountResponseFixedBuilder) CopyTo(to TradeAccountResponseBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// Copy
func (m TradeAccountResponsePointer) Copy(to TradeAccountResponseBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// Copy
func (m TradeAccountResponsePointerBuilder) Copy(to TradeAccountResponseBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m TradeAccountResponsePointer) CopyPointer(to TradeAccountResponsePointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m TradeAccountResponsePointerBuilder) CopyPointer(to TradeAccountResponsePointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyTo
func (m TradeAccountResponsePointer) CopyTo(to TradeAccountResponseFixedBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyTo
func (m TradeAccountResponsePointerBuilder) CopyTo(to TradeAccountResponseFixedBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyToPointer
func (m TradeAccountResponsePointer) CopyToPointer(to TradeAccountResponseFixedPointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyToPointer
func (m TradeAccountResponsePointerBuilder) CopyToPointer(to TradeAccountResponseFixedPointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// Copy
func (m TradeAccountResponseFixedPointer) Copy(to TradeAccountResponseFixedBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// Copy
func (m TradeAccountResponseFixedPointerBuilder) Copy(to TradeAccountResponseFixedBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m TradeAccountResponseFixedPointer) CopyPointer(to TradeAccountResponseFixedPointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m TradeAccountResponseFixedPointerBuilder) CopyPointer(to TradeAccountResponseFixedPointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyTo
func (m TradeAccountResponseFixedPointer) CopyTo(to TradeAccountResponseBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyTo
func (m TradeAccountResponseFixedPointerBuilder) CopyTo(to TradeAccountResponseBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyToPointer
func (m TradeAccountResponseFixedPointer) CopyToPointer(to TradeAccountResponsePointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyToPointer
func (m TradeAccountResponseFixedPointerBuilder) CopyToPointer(to TradeAccountResponsePointerBuilder) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}
