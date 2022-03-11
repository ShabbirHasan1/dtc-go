package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                = TradeAccountResponseFixedSize  (48)
//     Type                = TRADE_ACCOUNT_RESPONSE  (401)
//     TotalNumberMessages = 0
//     MessageNumber       = 0
//     TradeAccount        = ""
//     RequestID           = 0
// }
var _TradeAccountResponseFixedDefault = []byte{48, 0, 145, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountResponseFixedSize = 48

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

// ToBuilder
func (m TradeAccountResponseFixed) ToBuilder() TradeAccountResponseFixedBuilder {
	return TradeAccountResponseFixedBuilder{m.Fixed}
}

// Finish
func (m TradeAccountResponseFixedBuilder) Finish() TradeAccountResponseFixed {
	return TradeAccountResponseFixed{m.Fixed}
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

// SetTotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixedBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetMessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixedBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetTradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponseFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 44, 12, value)
}

// SetRequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponseFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 48, 44, value)
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
