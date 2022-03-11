package model

import (
	"github.com/moontrade/dtc-go/message"
)

const GeneralLogMessageSize = 10

const GeneralLogMessageFixedSize = 132

// {
//     Size        = GeneralLogMessageSize  (10)
//     Type        = GENERAL_LOG_MESSAGE  (701)
//     BaseSize    = GeneralLogMessageSize  (10)
//     MessageText = ""
// }
var _GeneralLogMessageDefault = []byte{10, 0, 189, 2, 10, 0, 0, 0, 0, 0}

// {
//     Size        = GeneralLogMessageFixedSize  (132)
//     Type        = GENERAL_LOG_MESSAGE  (701)
//     MessageText = ""
// }
var _GeneralLogMessageFixedDefault = []byte{132, 0, 189, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// GeneralLogMessage This message from the Server to the Client is a message which is to be
// added to a log file indicating information from the server. For example,
// if there are informational messages to provide during the process of a
// logon, this can be used to send those messages to a Client. A Client should
// never implement this message as a pop-up type message. Instead, it should
// be treated as a lower-level log type message.
//
// This message can be sent even before a LogonResponse is given.
type GeneralLogMessage struct {
	message.VLS
}

// GeneralLogMessageBuilder This message from the Server to the Client is a message which is to be
// added to a log file indicating information from the server. For example,
// if there are informational messages to provide during the process of a
// logon, this can be used to send those messages to a Client. A Client should
// never implement this message as a pop-up type message. Instead, it should
// be treated as a lower-level log type message.
//
// This message can be sent even before a LogonResponse is given.
type GeneralLogMessageBuilder struct {
	message.VLSBuilder
}

// GeneralLogMessageFixed This message from the Server to the Client is a message which is to be
// added to a log file indicating information from the server. For example,
// if there are informational messages to provide during the process of a
// logon, this can be used to send those messages to a Client. A Client should
// never implement this message as a pop-up type message. Instead, it should
// be treated as a lower-level log type message.
//
// This message can be sent even before a LogonResponse is given.
type GeneralLogMessageFixed struct {
	message.Fixed
}

// GeneralLogMessageFixedBuilder This message from the Server to the Client is a message which is to be
// added to a log file indicating information from the server. For example,
// if there are informational messages to provide during the process of a
// logon, this can be used to send those messages to a Client. A Client should
// never implement this message as a pop-up type message. Instead, it should
// be treated as a lower-level log type message.
//
// This message can be sent even before a LogonResponse is given.
type GeneralLogMessageFixedBuilder struct {
	message.Fixed
}

// GeneralLogMessagePointer This message from the Server to the Client is a message which is to be
// added to a log file indicating information from the server. For example,
// if there are informational messages to provide during the process of a
// logon, this can be used to send those messages to a Client. A Client should
// never implement this message as a pop-up type message. Instead, it should
// be treated as a lower-level log type message.
//
// This message can be sent even before a LogonResponse is given.
type GeneralLogMessagePointer struct {
	message.VLSPointer
}

// GeneralLogMessagePointerBuilder This message from the Server to the Client is a message which is to be
// added to a log file indicating information from the server. For example,
// if there are informational messages to provide during the process of a
// logon, this can be used to send those messages to a Client. A Client should
// never implement this message as a pop-up type message. Instead, it should
// be treated as a lower-level log type message.
//
// This message can be sent even before a LogonResponse is given.
type GeneralLogMessagePointerBuilder struct {
	message.VLSPointerBuilder
}

// GeneralLogMessageFixedPointer This message from the Server to the Client is a message which is to be
// added to a log file indicating information from the server. For example,
// if there are informational messages to provide during the process of a
// logon, this can be used to send those messages to a Client. A Client should
// never implement this message as a pop-up type message. Instead, it should
// be treated as a lower-level log type message.
//
// This message can be sent even before a LogonResponse is given.
type GeneralLogMessageFixedPointer struct {
	message.FixedPointer
}

// GeneralLogMessageFixedPointerBuilder This message from the Server to the Client is a message which is to be
// added to a log file indicating information from the server. For example,
// if there are informational messages to provide during the process of a
// logon, this can be used to send those messages to a Client. A Client should
// never implement this message as a pop-up type message. Instead, it should
// be treated as a lower-level log type message.
//
// This message can be sent even before a LogonResponse is given.
type GeneralLogMessageFixedPointerBuilder struct {
	message.FixedPointer
}

func NewGeneralLogMessageFrom(b []byte) GeneralLogMessage {
	return GeneralLogMessage{VLS: message.NewVLSFrom(b)}
}

func WrapGeneralLogMessage(b []byte) GeneralLogMessage {
	return GeneralLogMessage{VLS: message.WrapVLS(b)}
}

func NewGeneralLogMessage() GeneralLogMessageBuilder {
	a := GeneralLogMessageBuilder{message.NewVLSBuilder(10)}
	a.Unsafe().SetBytes(0, _GeneralLogMessageDefault)
	return a
}

func NewGeneralLogMessageFixedFrom(b []byte) GeneralLogMessageFixed {
	return GeneralLogMessageFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapGeneralLogMessageFixed(b []byte) GeneralLogMessageFixed {
	return GeneralLogMessageFixed{Fixed: message.WrapFixed(b)}
}

func NewGeneralLogMessageFixed() GeneralLogMessageFixedBuilder {
	a := GeneralLogMessageFixedBuilder{message.NewFixed(132)}
	a.Unsafe().SetBytes(0, _GeneralLogMessageFixedDefault)
	return a
}

func AllocGeneralLogMessage() GeneralLogMessagePointerBuilder {
	a := GeneralLogMessagePointerBuilder{message.AllocVLSBuilder(10)}
	a.Ptr.SetBytes(0, _GeneralLogMessageDefault)
	return a
}

func AllocGeneralLogMessageFrom(b []byte) GeneralLogMessagePointer {
	return GeneralLogMessagePointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocGeneralLogMessageFixed() GeneralLogMessageFixedPointerBuilder {
	a := GeneralLogMessageFixedPointerBuilder{message.AllocFixed(132)}
	a.Ptr.SetBytes(0, _GeneralLogMessageFixedDefault)
	return a
}

func AllocGeneralLogMessageFixedFrom(b []byte) GeneralLogMessageFixedPointer {
	return GeneralLogMessageFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size        = GeneralLogMessageSize  (10)
//     Type        = GENERAL_LOG_MESSAGE  (701)
//     BaseSize    = GeneralLogMessageSize  (10)
//     MessageText = ""
// }
func (m GeneralLogMessageBuilder) Clear() {
	m.Unsafe().SetBytes(0, _GeneralLogMessageDefault)
}

// Clear
// {
//     Size        = GeneralLogMessageFixedSize  (132)
//     Type        = GENERAL_LOG_MESSAGE  (701)
//     MessageText = ""
// }
func (m GeneralLogMessageFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _GeneralLogMessageFixedDefault)
}

// Clear
// {
//     Size        = GeneralLogMessageSize  (10)
//     Type        = GENERAL_LOG_MESSAGE  (701)
//     BaseSize    = GeneralLogMessageSize  (10)
//     MessageText = ""
// }
func (m GeneralLogMessagePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _GeneralLogMessageDefault)
}

// Clear
// {
//     Size        = GeneralLogMessageFixedSize  (132)
//     Type        = GENERAL_LOG_MESSAGE  (701)
//     MessageText = ""
// }
func (m GeneralLogMessageFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _GeneralLogMessageFixedDefault)
}

// ToBuilder
func (m GeneralLogMessage) ToBuilder() GeneralLogMessageBuilder {
	return GeneralLogMessageBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m GeneralLogMessageFixed) ToBuilder() GeneralLogMessageFixedBuilder {
	return GeneralLogMessageFixedBuilder{m.Fixed}
}

// ToBuilder
func (m GeneralLogMessagePointer) ToBuilder() GeneralLogMessagePointerBuilder {
	return GeneralLogMessagePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m GeneralLogMessageFixedPointer) ToBuilder() GeneralLogMessageFixedPointerBuilder {
	return GeneralLogMessageFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m GeneralLogMessageBuilder) Finish() GeneralLogMessage {
	return GeneralLogMessage{m.VLSBuilder.Finish()}
}

// Finish
func (m GeneralLogMessageFixedBuilder) Finish() GeneralLogMessageFixed {
	return GeneralLogMessageFixed{m.Fixed}
}

// Finish
func (m *GeneralLogMessagePointerBuilder) Finish() GeneralLogMessagePointer {
	return GeneralLogMessagePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *GeneralLogMessageFixedPointerBuilder) Finish() GeneralLogMessageFixedPointer {
	return GeneralLogMessageFixedPointer{m.FixedPointer}
}

// MessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessage) MessageText() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// MessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessageBuilder) MessageText() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// MessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessagePointer) MessageText() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// MessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessagePointerBuilder) MessageText() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// MessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessageFixed) MessageText() string {
	return message.StringFixed(m.Unsafe(), 132, 4)
}

// MessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessageFixedBuilder) MessageText() string {
	return message.StringFixed(m.Unsafe(), 132, 4)
}

// MessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessageFixedPointer) MessageText() string {
	return message.StringFixed(m.Ptr, 132, 4)
}

// MessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessageFixedPointerBuilder) MessageText() string {
	return message.StringFixed(m.Ptr, 132, 4)
}

// SetMessageText The message text to the Client which it should add to its log.
//
func (m *GeneralLogMessageBuilder) SetMessageText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetMessageText The message text to the Client which it should add to its log.
//
func (m *GeneralLogMessagePointerBuilder) SetMessageText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetMessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessageFixedBuilder) SetMessageText(value string) {
	message.SetStringFixed(m.Unsafe(), 132, 4, value)
}

// SetMessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessageFixedPointerBuilder) SetMessageText(value string) {
	message.SetStringFixed(m.Ptr, 132, 4, value)
}

// Copy
func (m GeneralLogMessage) Copy(to GeneralLogMessageBuilder) {
	to.SetMessageText(m.MessageText())
}

// Copy
func (m GeneralLogMessageBuilder) Copy(to GeneralLogMessageBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyPointer
func (m GeneralLogMessage) CopyPointer(to GeneralLogMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyPointer
func (m GeneralLogMessageBuilder) CopyPointer(to GeneralLogMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyToPointer
func (m GeneralLogMessage) CopyToPointer(to GeneralLogMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyToPointer
func (m GeneralLogMessageBuilder) CopyToPointer(to GeneralLogMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyTo
func (m GeneralLogMessage) CopyTo(to GeneralLogMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyTo
func (m GeneralLogMessageBuilder) CopyTo(to GeneralLogMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
}

// Copy
func (m GeneralLogMessageFixed) Copy(to GeneralLogMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
}

// Copy
func (m GeneralLogMessageFixedBuilder) Copy(to GeneralLogMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyPointer
func (m GeneralLogMessageFixed) CopyPointer(to GeneralLogMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyPointer
func (m GeneralLogMessageFixedBuilder) CopyPointer(to GeneralLogMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyToPointer
func (m GeneralLogMessageFixed) CopyToPointer(to GeneralLogMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyToPointer
func (m GeneralLogMessageFixedBuilder) CopyToPointer(to GeneralLogMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyTo
func (m GeneralLogMessageFixed) CopyTo(to GeneralLogMessageBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyTo
func (m GeneralLogMessageFixedBuilder) CopyTo(to GeneralLogMessageBuilder) {
	to.SetMessageText(m.MessageText())
}

// Copy
func (m GeneralLogMessagePointer) Copy(to GeneralLogMessageBuilder) {
	to.SetMessageText(m.MessageText())
}

// Copy
func (m GeneralLogMessagePointerBuilder) Copy(to GeneralLogMessageBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyPointer
func (m GeneralLogMessagePointer) CopyPointer(to GeneralLogMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyPointer
func (m GeneralLogMessagePointerBuilder) CopyPointer(to GeneralLogMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyTo
func (m GeneralLogMessagePointer) CopyTo(to GeneralLogMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyTo
func (m GeneralLogMessagePointerBuilder) CopyTo(to GeneralLogMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyToPointer
func (m GeneralLogMessagePointer) CopyToPointer(to GeneralLogMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyToPointer
func (m GeneralLogMessagePointerBuilder) CopyToPointer(to GeneralLogMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// Copy
func (m GeneralLogMessageFixedPointer) Copy(to GeneralLogMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
}

// Copy
func (m GeneralLogMessageFixedPointerBuilder) Copy(to GeneralLogMessageFixedBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyPointer
func (m GeneralLogMessageFixedPointer) CopyPointer(to GeneralLogMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyPointer
func (m GeneralLogMessageFixedPointerBuilder) CopyPointer(to GeneralLogMessageFixedPointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyTo
func (m GeneralLogMessageFixedPointer) CopyTo(to GeneralLogMessageBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyTo
func (m GeneralLogMessageFixedPointerBuilder) CopyTo(to GeneralLogMessageBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyToPointer
func (m GeneralLogMessageFixedPointer) CopyToPointer(to GeneralLogMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
}

// CopyToPointer
func (m GeneralLogMessageFixedPointerBuilder) CopyToPointer(to GeneralLogMessagePointerBuilder) {
	to.SetMessageText(m.MessageText())
}
