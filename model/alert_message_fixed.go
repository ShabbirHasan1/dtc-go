package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = AlertMessageFixedSize  (164)
//     Type         = ALERT_MESSAGE  (702)
//     MessageText  = ""
//     TradeAccount = ""
// }
var _AlertMessageFixedDefault = []byte{164, 0, 190, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AlertMessageFixedSize = 164

type AlertMessageFixed struct {
	message.Fixed
}

type AlertMessageFixedBuilder struct {
	message.Fixed
}

func NewAlertMessageFixedFrom(b []byte) AlertMessageFixed {
	return AlertMessageFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapAlertMessageFixed(b []byte) AlertMessageFixed {
	return AlertMessageFixed{Fixed: message.WrapFixed(b)}
}

func NewAlertMessageFixed() AlertMessageFixedBuilder {
	a := AlertMessageFixedBuilder{message.NewFixed(164)}
	a.Unsafe().SetBytes(0, _AlertMessageFixedDefault)
	return a
}

// Clear
// {
//     Size         = AlertMessageFixedSize  (164)
//     Type         = ALERT_MESSAGE  (702)
//     MessageText  = ""
//     TradeAccount = ""
// }
func (m AlertMessageFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AlertMessageFixedDefault)
}

// ToBuilder
func (m AlertMessageFixed) ToBuilder() AlertMessageFixedBuilder {
	return AlertMessageFixedBuilder{m.Fixed}
}

// Finish
func (m AlertMessageFixedBuilder) Finish() AlertMessageFixed {
	return AlertMessageFixed{m.Fixed}
}

// MessageText
func (m AlertMessageFixed) MessageText() string {
	return message.StringFixed(m.Unsafe(), 132, 4)
}

// MessageText
func (m AlertMessageFixedBuilder) MessageText() string {
	return message.StringFixed(m.Unsafe(), 132, 4)
}

// TradeAccount
func (m AlertMessageFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 164, 132)
}

// TradeAccount
func (m AlertMessageFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 164, 132)
}

// SetMessageText
func (m AlertMessageFixedBuilder) SetMessageText(value string) {
	message.SetStringFixed(m.Unsafe(), 132, 4, value)
}

// SetTradeAccount
func (m AlertMessageFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 164, 132, value)
}

// Copy
func (m AlertMessageFixed) Copy(to AlertMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AlertMessageFixedBuilder) Copy(to AlertMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AlertMessageFixed) CopyPointer(to AlertMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AlertMessageFixedBuilder) CopyPointer(to AlertMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AlertMessageFixed) CopyToPointer(to AlertMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AlertMessageFixedBuilder) CopyToPointer(to AlertMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AlertMessageFixed) CopyTo(to AlertMessageBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AlertMessageFixedBuilder) CopyTo(to AlertMessageBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}
