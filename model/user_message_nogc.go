package model

import (
	"github.com/moontrade/dtc-go/message"
)

// UserMessagePointer This message from the Server to the Client is for providing a message
// to the user.
//
// This message can be sent even before a LogonResponse.
type UserMessagePointer struct {
	message.VLSPointer
}

// UserMessagePointerBuilder This message from the Server to the Client is for providing a message
// to the user.
//
// This message can be sent even before a LogonResponse.
type UserMessagePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocUserMessage() UserMessagePointerBuilder {
	a := UserMessagePointerBuilder{message.AllocVLSBuilder(12)}
	a.Ptr.SetBytes(0, _UserMessageDefault)
	return a
}

func AllocUserMessageFrom(b []byte) UserMessagePointer {
	return UserMessagePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size           = UserMessageSize  (12)
//     Type           = USER_MESSAGE  (700)
//     BaseSize       = UserMessageSize  (12)
//     UserMessage    = ""
//     IsPopupMessage = 1
// }
func (m UserMessagePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _UserMessageDefault)
}

// ToBuilder
func (m UserMessagePointer) ToBuilder() UserMessagePointerBuilder {
	return UserMessagePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *UserMessagePointerBuilder) Finish() UserMessagePointer {
	return UserMessagePointer{m.VLSPointerBuilder.Finish()}
}

// UserMessage General message to present to user in the Client.
func (m UserMessagePointer) UserMessage() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// UserMessage General message to present to user in the Client.
func (m UserMessagePointerBuilder) UserMessage() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessagePointer) IsPopupMessage() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessagePointerBuilder) IsPopupMessage() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// SetUserMessage General message to present to user in the Client.
func (m *UserMessagePointerBuilder) SetUserMessage(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetIsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessagePointerBuilder) SetIsPopupMessage(value uint8) {
	message.SetUint8VLS(m.Ptr, 11, 10, value)
}

// Copy
func (m UserMessagePointer) Copy(to UserMessageBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// Copy
func (m UserMessagePointerBuilder) Copy(to UserMessageBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyPointer
func (m UserMessagePointer) CopyPointer(to UserMessagePointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyPointer
func (m UserMessagePointerBuilder) CopyPointer(to UserMessagePointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyTo
func (m UserMessagePointer) CopyTo(to UserMessageFixedBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyTo
func (m UserMessagePointerBuilder) CopyTo(to UserMessageFixedBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyToPointer
func (m UserMessagePointer) CopyToPointer(to UserMessageFixedPointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyToPointer
func (m UserMessagePointerBuilder) CopyToPointer(to UserMessageFixedPointerBuilder) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}
