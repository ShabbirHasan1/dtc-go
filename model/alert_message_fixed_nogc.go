package model

import (
	"github.com/moontrade/dtc-go/message"
)

type AlertMessageFixedPointer struct {
	message.FixedPointer
}

type AlertMessageFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocAlertMessageFixed() AlertMessageFixedPointerBuilder {
	a := AlertMessageFixedPointerBuilder{message.AllocFixed(164)}
	a.Ptr.SetBytes(0, _AlertMessageFixedDefault)
	return a
}

func AllocAlertMessageFixedFrom(b []byte) AlertMessageFixedPointer {
	return AlertMessageFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size         = AlertMessageFixedSize  (164)
//     Type         = ALERT_MESSAGE  (702)
//     MessageText  = ""
//     TradeAccount = ""
// }
func (m AlertMessageFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AlertMessageFixedDefault)
}

// ToBuilder
func (m AlertMessageFixedPointer) ToBuilder() AlertMessageFixedPointerBuilder {
	return AlertMessageFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *AlertMessageFixedPointerBuilder) Finish() AlertMessageFixedPointer {
	return AlertMessageFixedPointer{m.FixedPointer}
}

// MessageText
func (m AlertMessageFixedPointer) MessageText() string {
	return message.StringFixed(m.Ptr, 132, 4)
}

// MessageText
func (m AlertMessageFixedPointerBuilder) MessageText() string {
	return message.StringFixed(m.Ptr, 132, 4)
}

// TradeAccount
func (m AlertMessageFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 164, 132)
}

// TradeAccount
func (m AlertMessageFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 164, 132)
}

// SetMessageText
func (m AlertMessageFixedPointerBuilder) SetMessageText(value string) {
	message.SetStringFixed(m.Ptr, 132, 4, value)
}

// SetTradeAccount
func (m AlertMessageFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 164, 132, value)
}

// Copy
func (m AlertMessageFixedPointer) Copy(to AlertMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AlertMessageFixedPointerBuilder) Copy(to AlertMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AlertMessageFixedPointer) CopyPointer(to AlertMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AlertMessageFixedPointerBuilder) CopyPointer(to AlertMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AlertMessageFixedPointer) CopyTo(to AlertMessageBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AlertMessageFixedPointerBuilder) CopyTo(to AlertMessageBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AlertMessageFixedPointer) CopyToPointer(to AlertMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AlertMessageFixedPointerBuilder) CopyToPointer(to AlertMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}
