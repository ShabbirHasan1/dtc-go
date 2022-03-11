package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = AccountBalanceRequestFixedSize  (40)
//     Type         = ACCOUNT_BALANCE_REQUEST  (601)
//     RequestID    = 0
//     TradeAccount = ""
// }
var _AccountBalanceRequestFixedDefault = []byte{40, 0, 89, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AccountBalanceRequestFixedSize = 40

// AccountBalanceRequestFixed This is a message from the Client to the Server to request Trade Account
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
type AccountBalanceRequestFixed struct {
	message.Fixed
}

// AccountBalanceRequestFixedBuilder This is a message from the Client to the Server to request Trade Account
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
type AccountBalanceRequestFixedBuilder struct {
	message.Fixed
}

func NewAccountBalanceRequestFixedFrom(b []byte) AccountBalanceRequestFixed {
	return AccountBalanceRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapAccountBalanceRequestFixed(b []byte) AccountBalanceRequestFixed {
	return AccountBalanceRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewAccountBalanceRequestFixed() AccountBalanceRequestFixedBuilder {
	a := AccountBalanceRequestFixedBuilder{message.NewFixed(40)}
	a.Unsafe().SetBytes(0, _AccountBalanceRequestFixedDefault)
	return a
}

// Clear
// {
//     Size         = AccountBalanceRequestFixedSize  (40)
//     Type         = ACCOUNT_BALANCE_REQUEST  (601)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m AccountBalanceRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceRequestFixedDefault)
}

// ToBuilder
func (m AccountBalanceRequestFixed) ToBuilder() AccountBalanceRequestFixedBuilder {
	return AccountBalanceRequestFixedBuilder{m.Fixed}
}

// Finish
func (m AccountBalanceRequestFixedBuilder) Finish() AccountBalanceRequestFixed {
	return AccountBalanceRequestFixed{m.Fixed}
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID A unique request identifier for this request.
func (m AccountBalanceRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// SetRequestID A unique request identifier for this request.
func (m AccountBalanceRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// Account Balance data for all Trade Accounts on the logged in Username.
// Otherwise, specify a particular Trade Account to request Account Balance
// data for.
func (m AccountBalanceRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// Copy
func (m AccountBalanceRequestFixed) Copy(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AccountBalanceRequestFixedBuilder) Copy(to AccountBalanceRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestFixed) CopyPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AccountBalanceRequestFixedBuilder) CopyPointer(to AccountBalanceRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestFixed) CopyToPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AccountBalanceRequestFixedBuilder) CopyToPointer(to AccountBalanceRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestFixed) CopyTo(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AccountBalanceRequestFixedBuilder) CopyTo(to AccountBalanceRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}
