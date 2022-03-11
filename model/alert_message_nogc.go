package model

import (
	"github.com/moontrade/dtc-go/message"
)

type AlertMessagePointer struct {
	message.VLSPointer
}

type AlertMessagePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocAlertMessage() AlertMessagePointerBuilder {
	a := AlertMessagePointerBuilder{message.AllocVLSBuilder(14)}
	a.Ptr.SetBytes(0, _AlertMessageDefault)
	return a
}

func AllocAlertMessageFrom(b []byte) AlertMessagePointer {
	return AlertMessagePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = AlertMessageSize  (14)
//     Type         = ALERT_MESSAGE  (702)
//     BaseSize     = AlertMessageSize  (14)
//     MessageText  = ""
//     TradeAccount = ""
// }
func (m AlertMessagePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AlertMessageDefault)
}

// ToBuilder
func (m AlertMessagePointer) ToBuilder() AlertMessagePointerBuilder {
	return AlertMessagePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *AlertMessagePointerBuilder) Finish() AlertMessagePointer {
	return AlertMessagePointer{m.VLSPointerBuilder.Finish()}
}

// MessageText
func (m AlertMessagePointer) MessageText() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// MessageText
func (m AlertMessagePointerBuilder) MessageText() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m AlertMessagePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m AlertMessagePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// SetMessageText
func (m *AlertMessagePointerBuilder) SetMessageText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetTradeAccount
func (m *AlertMessagePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// Copy
func (m AlertMessagePointer) Copy(to AlertMessageBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m AlertMessagePointerBuilder) Copy(to AlertMessageBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AlertMessagePointer) CopyPointer(to AlertMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m AlertMessagePointerBuilder) CopyPointer(to AlertMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AlertMessagePointer) CopyTo(to AlertMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m AlertMessagePointerBuilder) CopyTo(to AlertMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AlertMessagePointer) CopyToPointer(to AlertMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m AlertMessagePointerBuilder) CopyToPointer(to AlertMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
	to.SetTradeAccount(m.TradeAccount())
}
