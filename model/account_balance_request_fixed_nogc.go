package model

import (
	"github.com/moontrade/dtc-go/message"
)

// AccountBalanceRequestFixedPointer This is a message from the Client to the Server to request Trade Account
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
type AccountBalanceRequestFixedPointer struct {
	message.FixedPointer
}

// AccountBalanceRequestFixedPointerBuilder This is a message from the Client to the Server to request Trade Account
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
type AccountBalanceRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocAccountBalanceRequestFixed() AccountBalanceRequestFixedPointerBuilder {
	a := AccountBalanceRequestFixedPointerBuilder{message.AllocFixed(40)}
	a.Ptr.SetBytes(0, _AccountBalanceRequestFixedDefault)
	return a
}

func AllocAccountBalanceRequestFixedFrom(b []byte) AccountBalanceRequestFixedPointer {
	return AccountBalanceRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size         = AccountBalanceRequestFixedSize  (40)
//     Type         = ACCOUNT_BALANCE_REQUEST  (601)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m AccountBalanceRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceRequestFixedDefault)
}

// ToBuilder
func (m AccountBalanceRequestFixedPointer) ToBuilder() AccountBalanceRequestFixedPointerBuilder {
	return AccountBalanceRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *AccountBalanceRequestFixedPointerBuilder) Finish() AccountBalanceRequestFixedPointer {
	return AccountBalanceRequestFixedPointer{m.FixedPointer}
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// SetRequestID A unique request identifier for this request.
func (m AccountBalanceRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// Copy
func (m AccountBalanceRequestFixedPointer) Copy(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AccountBalanceRequestFixedPointerBuilder) Copy(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestFixedPointer) CopyPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestFixedPointerBuilder) CopyPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestFixedPointer) CopyTo(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestFixedPointerBuilder) CopyTo(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestFixedPointer) CopyToPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestFixedPointerBuilder) CopyToPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}
