package model

import (
	"github.com/moontrade/dtc-go/message"
)

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

func AllocTradeAccountResponse() TradeAccountResponsePointerBuilder {
	a := TradeAccountResponsePointerBuilder{message.AllocVLSBuilder(24)}
	a.Ptr.SetBytes(0, _TradeAccountResponseDefault)
	return a
}

func AllocTradeAccountResponseFrom(b []byte) TradeAccountResponsePointer {
	return TradeAccountResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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

// ToBuilder
func (m TradeAccountResponsePointer) ToBuilder() TradeAccountResponsePointerBuilder {
	return TradeAccountResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountResponsePointerBuilder) Finish() TradeAccountResponsePointer {
	return TradeAccountResponsePointer{m.VLSPointerBuilder.Finish()}
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
func (m TradeAccountResponsePointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 24, 20)
}

// RequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponsePointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 24, 20)
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
func (m TradeAccountResponsePointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Ptr, 16, 12, value)
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
func (m TradeAccountResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 24, 20, value)
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
