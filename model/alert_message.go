package model

import (
	"github.com/moontrade/dtc-go/message"
)

const AlertMessageSize = 14

const AlertMessageFixedSize = 164

// {
//     Size         = AlertMessageSize  (14)
//     Type         = ALERT_MESSAGE  (702)
//     BaseSize     = AlertMessageSize  (14)
//     MessageText  = ""
//     TradeAccount = ""
// }
var _AlertMessageDefault = []byte{14, 0, 190, 2, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size         = AlertMessageFixedSize  (164)
//     Type         = ALERT_MESSAGE  (702)
//     MessageText  = ""
//     TradeAccount = ""
// }
var _AlertMessageFixedDefault = []byte{164, 0, 190, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type AlertMessage struct {
	message.VLS
}

type AlertMessageBuilder struct {
	message.VLSBuilder
}

type AlertMessageFixed struct {
	message.Fixed
}

type AlertMessageFixedBuilder struct {
	message.Fixed
}

type AlertMessagePointer struct {
	message.VLSPointer
}

type AlertMessagePointerBuilder struct {
	message.VLSPointerBuilder
}

type AlertMessageFixedPointer struct {
	message.FixedPointer
}

type AlertMessageFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocAlertMessage() AlertMessagePointerBuilder {
	a := AlertMessagePointerBuilder{message.AllocVLSBuilder(14)}
	a.Ptr.SetBytes(0, _AlertMessageDefault)
	return a
}

func AllocAlertMessageFrom(b []byte) AlertMessagePointer {
	return AlertMessagePointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size         = AlertMessageSize  (14)
//     Type         = ALERT_MESSAGE  (702)
//     BaseSize     = AlertMessageSize  (14)
//     MessageText  = ""
//     TradeAccount = ""
// }
func (m AlertMessageBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AlertMessageDefault)
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
func (m AlertMessage) ToBuilder() AlertMessageBuilder {
	return AlertMessageBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m AlertMessageFixed) ToBuilder() AlertMessageFixedBuilder {
	return AlertMessageFixedBuilder{m.Fixed}
}

// ToBuilder
func (m AlertMessagePointer) ToBuilder() AlertMessagePointerBuilder {
	return AlertMessagePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m AlertMessageFixedPointer) ToBuilder() AlertMessageFixedPointerBuilder {
	return AlertMessageFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m AlertMessageBuilder) Finish() AlertMessage {
	return AlertMessage{m.VLSBuilder.Finish()}
}

// Finish
func (m AlertMessageFixedBuilder) Finish() AlertMessageFixed {
	return AlertMessageFixed{m.Fixed}
}

// Finish
func (m *AlertMessagePointerBuilder) Finish() AlertMessagePointer {
	return AlertMessagePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *AlertMessageFixedPointerBuilder) Finish() AlertMessageFixedPointer {
	return AlertMessageFixedPointer{m.FixedPointer}
}

// MessageText
func (m AlertMessage) MessageText() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// MessageText
func (m AlertMessageBuilder) MessageText() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
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
func (m AlertMessage) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m AlertMessageBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m AlertMessagePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m AlertMessagePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// MessageText
func (m AlertMessageFixed) MessageText() string {
	return message.StringFixed(m.Unsafe(), 132, 4)
}

// MessageText
func (m AlertMessageFixedBuilder) MessageText() string {
	return message.StringFixed(m.Unsafe(), 132, 4)
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
func (m AlertMessageFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 164, 132)
}

// TradeAccount
func (m AlertMessageFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 164, 132)
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
func (m *AlertMessageBuilder) SetMessageText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetMessageText
func (m *AlertMessagePointerBuilder) SetMessageText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetTradeAccount
func (m *AlertMessageBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *AlertMessagePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetMessageText
func (m AlertMessageFixedBuilder) SetMessageText(value string) {
	message.SetStringFixed(m.Unsafe(), 132, 4, value)
}

// SetMessageText
func (m AlertMessageFixedPointerBuilder) SetMessageText(value string) {
	message.SetStringFixed(m.Ptr, 132, 4, value)
}

// SetTradeAccount
func (m AlertMessageFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 164, 132, value)
}

// SetTradeAccount
func (m AlertMessageFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 164, 132, value)
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
