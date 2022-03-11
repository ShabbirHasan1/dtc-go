package model

import (
	"github.com/moontrade/dtc-go/message"
)

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

const TradeAccountResponseSize = 24

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

// ToBuilder
func (m TradeAccountResponse) ToBuilder() TradeAccountResponseBuilder {
	return TradeAccountResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountResponseBuilder) Finish() TradeAccountResponse {
	return TradeAccountResponse{m.VLSBuilder.Finish()}
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

// SetTotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetMessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 16, 12, value)
}

// SetTradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m *TradeAccountResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetRequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, value)
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
