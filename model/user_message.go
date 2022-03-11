package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size           = UserMessageSize  (12)
//     Type           = USER_MESSAGE  (700)
//     BaseSize       = UserMessageSize  (12)
//     UserMessage    = ""
//     IsPopupMessage = 1
// }
var _UserMessageDefault = []byte{12, 0, 188, 2, 12, 0, 0, 0, 0, 0, 1, 0}

const UserMessageSize = 12

// UserMessage This message from the Server to the Client is for providing a message
// to the user.
//
// This message can be sent even before a LogonResponse.
type UserMessage struct {
	message.VLS
}

// UserMessageBuilder This message from the Server to the Client is for providing a message
// to the user.
//
// This message can be sent even before a LogonResponse.
type UserMessageBuilder struct {
	message.VLSBuilder
}

func NewUserMessageFrom(b []byte) UserMessage {
	return UserMessage{VLS: message.NewVLSFrom(b)}
}

func WrapUserMessage(b []byte) UserMessage {
	return UserMessage{VLS: message.WrapVLS(b)}
}

func NewUserMessage() UserMessageBuilder {
	a := UserMessageBuilder{message.NewVLSBuilder(12)}
	a.Unsafe().SetBytes(0, _UserMessageDefault)
	return a
}

// Clear
// {
//     Size           = UserMessageSize  (12)
//     Type           = USER_MESSAGE  (700)
//     BaseSize       = UserMessageSize  (12)
//     UserMessage    = ""
//     IsPopupMessage = 1
// }
func (m UserMessageBuilder) Clear() {
	m.Unsafe().SetBytes(0, _UserMessageDefault)
}

// ToBuilder
func (m UserMessage) ToBuilder() UserMessageBuilder {
	return UserMessageBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m UserMessageBuilder) Finish() UserMessage {
	return UserMessage{m.VLSBuilder.Finish()}
}

// UserMessage General message to present to user in the Client.
func (m UserMessage) UserMessage() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// UserMessage General message to present to user in the Client.
func (m UserMessageBuilder) UserMessage() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessage) IsPopupMessage() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageBuilder) IsPopupMessage() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// SetUserMessage General message to present to user in the Client.
func (m *UserMessageBuilder) SetUserMessage(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetIsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageBuilder) SetIsPopupMessage(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 11, 10, value)
}

// Copy
func (m UserMessage) Copy(to UserMessageBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// Copy
func (m UserMessageBuilder) Copy(to UserMessageBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyPointer
func (m UserMessage) CopyPointer(to UserMessagePointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyPointer
func (m UserMessageBuilder) CopyPointer(to UserMessagePointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyToPointer
func (m UserMessage) CopyToPointer(to UserMessageFixedPointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyToPointer
func (m UserMessageBuilder) CopyToPointer(to UserMessageFixedPointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyTo
func (m UserMessage) CopyTo(to UserMessageFixedBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyTo
func (m UserMessageBuilder) CopyTo(to UserMessageFixedBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}
