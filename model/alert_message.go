package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = AlertMessageSize  (14)
//     Type         = ALERT_MESSAGE  (702)
//     BaseSize     = AlertMessageSize  (14)
//     MessageText  = ""
//     TradeAccount = ""
// }
var _AlertMessageDefault = []byte{14, 0, 190, 2, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AlertMessageSize = 14

type AlertMessage struct {
	message.VLS
}

type AlertMessageBuilder struct {
	message.VLSBuilder
}

func NewAlertMessageFrom(b []byte) AlertMessage {
	return AlertMessage{VLS: message.NewVLSFrom(b)}
}

func WrapAlertMessage(b []byte) AlertMessage {
	return AlertMessage{VLS: message.WrapVLS(b)}
}

func NewAlertMessage() AlertMessageBuilder {
	a := AlertMessageBuilder{message.NewVLSBuilder(14)}
	a.Unsafe().SetBytes(0, _AlertMessageDefault)
	return a
}

// Clear
// {
//     Size         = AlertMessageSize  (14)
//     Type         = ALERT_MESSAGE  (702)
//     BaseSize     = AlertMessageSize  (14)
//     MessageText  = ""
//     TradeAccount = ""
// }
func (m AlertMessageBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AlertMessageDefault)
}

// ToBuilder
func (m AlertMessage) ToBuilder() AlertMessageBuilder {
	return AlertMessageBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m AlertMessageBuilder) Finish() AlertMessage {
	return AlertMessage{m.VLSBuilder.Finish()}
}

// MessageText
func (m AlertMessage) MessageText() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// MessageText
func (m AlertMessageBuilder) MessageText() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m AlertMessage) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m AlertMessageBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// SetMessageText
func (m *AlertMessageBuilder) SetMessageText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetTradeAccount
func (m *AlertMessageBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// Copy
func (m AlertMessage) Copy(to AlertMessageBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AlertMessageBuilder) Copy(to AlertMessageBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AlertMessage) CopyPointer(to AlertMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AlertMessageBuilder) CopyPointer(to AlertMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AlertMessage) CopyToPointer(to AlertMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AlertMessageBuilder) CopyToPointer(to AlertMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AlertMessage) CopyTo(to AlertMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AlertMessageBuilder) CopyTo(to AlertMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}
