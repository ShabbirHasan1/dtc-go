package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size        = GeneralLogMessageFixedSize  (132)
//     Type        = GENERAL_LOG_MESSAGE  (701)
//     MessageText = ""
// }
var _GeneralLogMessageFixedDefault = []byte{132, 0, 189, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const GeneralLogMessageFixedSize = 132

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

// Clear
// {
//     Size        = GeneralLogMessageFixedSize  (132)
//     Type        = GENERAL_LOG_MESSAGE  (701)
//     MessageText = ""
// }
func (m GeneralLogMessageFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _GeneralLogMessageFixedDefault)
}

// ToBuilder
func (m GeneralLogMessageFixed) ToBuilder() GeneralLogMessageFixedBuilder {
	return GeneralLogMessageFixedBuilder{m.Fixed}
}

// Finish
func (m GeneralLogMessageFixedBuilder) Finish() GeneralLogMessageFixed {
	return GeneralLogMessageFixed{m.Fixed}
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

// SetMessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessageFixedBuilder) SetMessageText(value string) {
	message.SetStringFixed(m.Unsafe(), 132, 4, value)
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
