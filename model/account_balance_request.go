package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = AccountBalanceRequestSize  (16)
//     Type         = ACCOUNT_BALANCE_REQUEST  (601)
//     BaseSize     = AccountBalanceRequestSize  (16)
//     RequestID    = 0
//     TradeAccount = ""
// }
var _AccountBalanceRequestDefault = []byte{16, 0, 89, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AccountBalanceRequestSize = 16

// AccountBalanceRequest This is a message from the Client to the Server to request Trade Account
// Balance data.
//
// The Server will respond with an AccountBalanceUpdate or reject the request.
// The Server will respond with an AccountBalanceUpdate or reject the request.
//
// The Server will set the RequestID in the AccountBalanceUpdate message
// to match the RequestID in the AccountBalanceRequest.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When unsolicited AccountBalanceUpdate messages are sent by the Server,
// the RequestID will be 0.
type AccountBalanceRequest struct {
	message.VLS
}

// AccountBalanceRequestBuilder This is a message from the Client to the Server to request Trade Account
// Balance data.
//
// The Server will respond with an AccountBalanceUpdate or reject the request.
// The Server will respond with an AccountBalanceUpdate or reject the request.
//
// The Server will set the RequestID in the AccountBalanceUpdate message
// to match the RequestID in the AccountBalanceRequest.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When unsolicited AccountBalanceUpdate messages are sent by the Server,
// the RequestID will be 0.
type AccountBalanceRequestBuilder struct {
	message.VLSBuilder
}

func NewAccountBalanceRequestFrom(b []byte) AccountBalanceRequest {
	return AccountBalanceRequest{VLS: message.NewVLSFrom(b)}
}

func WrapAccountBalanceRequest(b []byte) AccountBalanceRequest {
	return AccountBalanceRequest{VLS: message.WrapVLS(b)}
}

func NewAccountBalanceRequest() AccountBalanceRequestBuilder {
	a := AccountBalanceRequestBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _AccountBalanceRequestDefault)
	return a
}

// Clear
// {
//     Size         = AccountBalanceRequestSize  (16)
//     Type         = ACCOUNT_BALANCE_REQUEST  (601)
//     BaseSize     = AccountBalanceRequestSize  (16)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m AccountBalanceRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceRequestDefault)
}

// ToBuilder
func (m AccountBalanceRequest) ToBuilder() AccountBalanceRequestBuilder {
	return AccountBalanceRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m AccountBalanceRequestBuilder) Finish() AccountBalanceRequest {
	return AccountBalanceRequest{m.VLSBuilder.Finish()}
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetRequestID A unique request identifier for this request.
func (m AccountBalanceRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m *AccountBalanceRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// Copy
func (m AccountBalanceRequest) Copy(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AccountBalanceRequestBuilder) Copy(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequest) CopyPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestBuilder) CopyPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequest) CopyToPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestBuilder) CopyToPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequest) CopyTo(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestBuilder) CopyTo(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}
