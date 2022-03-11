package model

import (
	"github.com/moontrade/dtc-go/message"
)

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

func AllocGeneralLogMessage() GeneralLogMessagePointerBuilder {
	a := GeneralLogMessagePointerBuilder{message.AllocVLSBuilder(10)}
	a.Ptr.SetBytes(0, _GeneralLogMessageDefault)
	return a
}

func AllocGeneralLogMessageFrom(b []byte) GeneralLogMessagePointer {
	return GeneralLogMessagePointer{VLSPointer: message.AllocVLSFrom(b)}
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

// ToBuilder
func (m GeneralLogMessagePointer) ToBuilder() GeneralLogMessagePointerBuilder {
	return GeneralLogMessagePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *GeneralLogMessagePointerBuilder) Finish() GeneralLogMessagePointer {
	return GeneralLogMessagePointer{m.VLSPointerBuilder.Finish()}
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

// SetMessageText The message text to the Client which it should add to its log.
//
func (m *GeneralLogMessagePointerBuilder) SetMessageText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
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
