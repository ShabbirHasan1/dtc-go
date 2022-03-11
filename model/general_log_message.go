package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size        = GeneralLogMessageSize  (10)
//     Type        = GENERAL_LOG_MESSAGE  (701)
//     BaseSize    = GeneralLogMessageSize  (10)
//     MessageText = ""
// }
var _GeneralLogMessageDefault = []byte{10, 0, 189, 2, 10, 0, 0, 0, 0, 0}

const GeneralLogMessageSize = 10

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

// ToBuilder
func (m GeneralLogMessage) ToBuilder() GeneralLogMessageBuilder {
	return GeneralLogMessageBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m GeneralLogMessageBuilder) Finish() GeneralLogMessage {
	return GeneralLogMessage{m.VLSBuilder.Finish()}
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

// SetMessageText The message text to the Client which it should add to its log.
//
func (m *GeneralLogMessageBuilder) SetMessageText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
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
