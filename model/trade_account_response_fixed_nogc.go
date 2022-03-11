package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
func (m TradeAccountResponseFixedPointer) ToBuilder() TradeAccountResponseFixedPointerBuilder {
	return TradeAccountResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *TradeAccountResponseFixedPointerBuilder) Finish() TradeAccountResponseFixedPointer {
	return TradeAccountResponseFixedPointer{m.FixedPointer}
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
func (m TradeAccountResponseFixedPointerBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
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
func (m TradeAccountResponseFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 44, 12, value)
}

// SetRequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponseFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 48, 44, value)
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
