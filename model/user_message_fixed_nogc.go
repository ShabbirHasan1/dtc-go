package model

import (
	"github.com/moontrade/dtc-go/message"
)

// UserMessageFixedPointer This message from the Server to the Client is for providing a message
// to the user.
//
// This message can be sent even before a LogonResponse.
type UserMessageFixedPointer struct {
	message.FixedPointer
}

// UserMessageFixedPointerBuilder This message from the Server to the Client is for providing a message
// to the user.
//
// This message can be sent even before a LogonResponse.
type UserMessageFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocUserMessageFixed() UserMessageFixedPointerBuilder {
	a := UserMessageFixedPointerBuilder{message.AllocFixed(262)}
	a.Ptr.SetBytes(0, _UserMessageFixedDefault)
	return a
}

func AllocUserMessageFixedFrom(b []byte) UserMessageFixedPointer {
	return UserMessageFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size           = UserMessageFixedSize  (262)
//     Type           = USER_MESSAGE  (700)
//     UserMessage    = ""
//     IsPopupMessage = 1
// }
func (m UserMessageFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _UserMessageFixedDefault)
}

// ToBuilder
func (m UserMessageFixedPointer) ToBuilder() UserMessageFixedPointerBuilder {
	return UserMessageFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *UserMessageFixedPointerBuilder) Finish() UserMessageFixedPointer {
	return UserMessageFixedPointer{m.FixedPointer}
}

// UserMessage General message to present to user in the Client.
func (m UserMessageFixedPointer) UserMessage() string {
	return message.StringFixed(m.Ptr, 260, 4)
}

// UserMessage General message to present to user in the Client.
func (m UserMessageFixedPointerBuilder) UserMessage() string {
	return message.StringFixed(m.Ptr, 260, 4)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageFixedPointer) IsPopupMessage() uint8 {
	return message.Uint8Fixed(m.Ptr, 261, 260)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageFixedPointerBuilder) IsPopupMessage() uint8 {
	return message.Uint8Fixed(m.Ptr, 261, 260)
}

// SetUserMessage General message to present to user in the Client.
func (m UserMessageFixedPointerBuilder) SetUserMessage(value string) {
	message.SetStringFixed(m.Ptr, 260, 4, value)
}

// SetIsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageFixedPointerBuilder) SetIsPopupMessage(value uint8) {
	message.SetUint8Fixed(m.Ptr, 261, 260, value)
}

// Copy
func (m UserMessageFixedPointer) Copy(to UserMessageFixedBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// Copy
func (m UserMessageFixedPointerBuilder) Copy(to UserMessageFixedBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyPointer
func (m UserMessageFixedPointer) CopyPointer(to UserMessageFixedPointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyPointer
func (m UserMessageFixedPointerBuilder) CopyPointer(to UserMessageFixedPointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyTo
func (m UserMessageFixedPointer) CopyTo(to UserMessageBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyTo
func (m UserMessageFixedPointerBuilder) CopyTo(to UserMessageBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyToPointer
func (m UserMessageFixedPointer) CopyToPointer(to UserMessagePointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyToPointer
func (m UserMessageFixedPointerBuilder) CopyToPointer(to UserMessagePointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}
