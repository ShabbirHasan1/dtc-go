package model

import (
	"github.com/moontrade/dtc-go/message"
)

// AccountBalanceRequestPointer This is a message from the Client to the Server to request Trade Account
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
type AccountBalanceRequestPointer struct {
	message.VLSPointer
}

// AccountBalanceRequestPointerBuilder This is a message from the Client to the Server to request Trade Account
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
type AccountBalanceRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocAccountBalanceRequest() AccountBalanceRequestPointerBuilder {
	a := AccountBalanceRequestPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _AccountBalanceRequestDefault)
	return a
}

func AllocAccountBalanceRequestFrom(b []byte) AccountBalanceRequestPointer {
	return AccountBalanceRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = AccountBalanceRequestSize  (16)
//     Type         = ACCOUNT_BALANCE_REQUEST  (601)
//     BaseSize     = AccountBalanceRequestSize  (16)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m AccountBalanceRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceRequestDefault)
}

// ToBuilder
func (m AccountBalanceRequestPointer) ToBuilder() AccountBalanceRequestPointerBuilder {
	return AccountBalanceRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *AccountBalanceRequestPointerBuilder) Finish() AccountBalanceRequestPointer {
	return AccountBalanceRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetRequestID A unique request identifier for this request.
func (m AccountBalanceRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m *AccountBalanceRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// Copy
func (m AccountBalanceRequestPointer) Copy(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AccountBalanceRequestPointerBuilder) Copy(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestPointer) CopyPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestPointerBuilder) CopyPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestPointer) CopyTo(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestPointerBuilder) CopyTo(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestPointer) CopyToPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestPointerBuilder) CopyToPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}
