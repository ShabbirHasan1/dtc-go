package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
//     Size        = GeneralLogMessageFixedSize  (132)
//     Type        = GENERAL_LOG_MESSAGE  (701)
//     MessageText = ""
// }
func (m GeneralLogMessageFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _GeneralLogMessageFixedDefault)
}

// ToBuilder
func (m GeneralLogMessageFixedPointer) ToBuilder() GeneralLogMessageFixedPointerBuilder {
	return GeneralLogMessageFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *GeneralLogMessageFixedPointerBuilder) Finish() GeneralLogMessageFixedPointer {
	return GeneralLogMessageFixedPointer{m.FixedPointer}
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
func (m GeneralLogMessageFixedPointerBuilder) SetMessageText(value string) {
	message.SetStringFixed(m.Ptr, 132, 4, value)
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
