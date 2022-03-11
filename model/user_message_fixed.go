package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size           = UserMessageFixedSize  (262)
//     Type           = USER_MESSAGE  (700)
//     UserMessage    = ""
//     IsPopupMessage = 1
// }
var _UserMessageFixedDefault = []byte{6, 1, 188, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}

const UserMessageFixedSize = 262

// UserMessageFixed This message from the Server to the Client is for providing a message
// to the user.
//
// This message can be sent even before a LogonResponse.
type UserMessageFixed struct {
	message.Fixed
}

// UserMessageFixedBuilder This message from the Server to the Client is for providing a message
// to the user.
//
// This message can be sent even before a LogonResponse.
type UserMessageFixedBuilder struct {
	message.Fixed
}

func NewUserMessageFixedFrom(b []byte) UserMessageFixed {
	return UserMessageFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapUserMessageFixed(b []byte) UserMessageFixed {
	return UserMessageFixed{Fixed: message.WrapFixed(b)}
}

func NewUserMessageFixed() UserMessageFixedBuilder {
	a := UserMessageFixedBuilder{message.NewFixed(262)}
	a.Unsafe().SetBytes(0, _UserMessageFixedDefault)
	return a
}

// Clear
// {
//     Size           = UserMessageFixedSize  (262)
//     Type           = USER_MESSAGE  (700)
//     UserMessage    = ""
//     IsPopupMessage = 1
// }
func (m UserMessageFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _UserMessageFixedDefault)
}

// ToBuilder
func (m UserMessageFixed) ToBuilder() UserMessageFixedBuilder {
	return UserMessageFixedBuilder{m.Fixed}
}

// Finish
func (m UserMessageFixedBuilder) Finish() UserMessageFixed {
	return UserMessageFixed{m.Fixed}
}

// UserMessage General message to present to user in the Client.
func (m UserMessageFixed) UserMessage() string {
	return message.StringFixed(m.Unsafe(), 260, 4)
}

// UserMessage General message to present to user in the Client.
func (m UserMessageFixedBuilder) UserMessage() string {
	return message.StringFixed(m.Unsafe(), 260, 4)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageFixed) IsPopupMessage() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 261, 260)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageFixedBuilder) IsPopupMessage() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 261, 260)
}

// SetUserMessage General message to present to user in the Client.
func (m UserMessageFixedBuilder) SetUserMessage(value string) {
	message.SetStringFixed(m.Unsafe(), 260, 4, value)
}

// SetIsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageFixedBuilder) SetIsPopupMessage(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 261, 260, value)
}

// Copy
func (m UserMessageFixed) Copy(to UserMessageFixedBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// Copy
func (m UserMessageFixedBuilder) Copy(to UserMessageFixedBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyPointer
func (m UserMessageFixed) CopyPointer(to UserMessageFixedPointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyPointer
func (m UserMessageFixedBuilder) CopyPointer(to UserMessageFixedPointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyToPointer
func (m UserMessageFixed) CopyToPointer(to UserMessagePointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyToPointer
func (m UserMessageFixedBuilder) CopyToPointer(to UserMessagePointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyTo
func (m UserMessageFixed) CopyTo(to UserMessageBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyTo
func (m UserMessageFixedBuilder) CopyTo(to UserMessageBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}
