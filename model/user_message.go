package model

import (
	"github.com/moontrade/dtc-go/message"
)

const UserMessageSize = 12

const UserMessageFixedSize = 262

// {
//     Size           = UserMessageSize  (12)
//     Type           = USER_MESSAGE  (700)
//     BaseSize       = UserMessageSize  (12)
//     UserMessage    = ""
//     IsPopupMessage = true
// }
var _UserMessageDefault = []byte{12, 0, 188, 2, 12, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size           = UserMessageFixedSize  (262)
//     Type           = USER_MESSAGE  (700)
//     UserMessage    = ""
//     IsPopupMessage = true
// }
var _UserMessageFixedDefault = []byte{6, 1, 188, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocUserMessage() UserMessagePointerBuilder {
	a := UserMessagePointerBuilder{message.AllocVLSBuilder(12)}
	a.Ptr.SetBytes(0, _UserMessageDefault)
	return a
}

func AllocUserMessageFrom(b []byte) UserMessagePointer {
	return UserMessagePointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size           = UserMessageSize  (12)
//     Type           = USER_MESSAGE  (700)
//     BaseSize       = UserMessageSize  (12)
//     UserMessage    = ""
//     IsPopupMessage = true
// }
func (m UserMessageBuilder) Clear() {
	m.Unsafe().SetBytes(0, _UserMessageDefault)
}

// Clear
// {
//     Size           = UserMessageFixedSize  (262)
//     Type           = USER_MESSAGE  (700)
//     UserMessage    = ""
//     IsPopupMessage = true
// }
func (m UserMessageFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _UserMessageFixedDefault)
}

// Clear
// {
//     Size           = UserMessageSize  (12)
//     Type           = USER_MESSAGE  (700)
//     BaseSize       = UserMessageSize  (12)
//     UserMessage    = ""
//     IsPopupMessage = true
// }
func (m UserMessagePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _UserMessageDefault)
}

// Clear
// {
//     Size           = UserMessageFixedSize  (262)
//     Type           = USER_MESSAGE  (700)
//     UserMessage    = ""
//     IsPopupMessage = true
// }
func (m UserMessageFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _UserMessageFixedDefault)
}

// ToBuilder
func (m UserMessage) ToBuilder() UserMessageBuilder {
	return UserMessageBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m UserMessageFixed) ToBuilder() UserMessageFixedBuilder {
	return UserMessageFixedBuilder{m.Fixed}
}

// ToBuilder
func (m UserMessagePointer) ToBuilder() UserMessagePointerBuilder {
	return UserMessagePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m UserMessageFixedPointer) ToBuilder() UserMessageFixedPointerBuilder {
	return UserMessageFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m UserMessageBuilder) Finish() UserMessage {
	return UserMessage{m.VLSBuilder.Finish()}
}

// Finish
func (m UserMessageFixedBuilder) Finish() UserMessageFixed {
	return UserMessageFixed{m.Fixed}
}

// Finish
func (m *UserMessagePointerBuilder) Finish() UserMessagePointer {
	return UserMessagePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *UserMessageFixedPointerBuilder) Finish() UserMessageFixedPointer {
	return UserMessageFixedPointer{m.FixedPointer}
}

// UserMessage General message to present to user in the Client.
func (m UserMessage) UserMessage() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// UserMessage General message to present to user in the Client.
func (m UserMessageBuilder) UserMessage() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
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
func (m UserMessage) IsPopupMessage() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageBuilder) IsPopupMessage() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessagePointer) IsPopupMessage() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessagePointerBuilder) IsPopupMessage() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// UserMessage General message to present to user in the Client.
func (m UserMessageFixed) UserMessage() string {
	return message.StringFixed(m.Unsafe(), 260, 4)
}

// UserMessage General message to present to user in the Client.
func (m UserMessageFixedBuilder) UserMessage() string {
	return message.StringFixed(m.Unsafe(), 260, 4)
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
func (m UserMessageFixed) IsPopupMessage() bool {
	return message.BoolFixed(m.Unsafe(), 261, 260)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageFixedBuilder) IsPopupMessage() bool {
	return message.BoolFixed(m.Unsafe(), 261, 260)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageFixedPointer) IsPopupMessage() bool {
	return message.BoolFixed(m.Ptr, 261, 260)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageFixedPointerBuilder) IsPopupMessage() bool {
	return message.BoolFixed(m.Ptr, 261, 260)
}

// SetUserMessage General message to present to user in the Client.
func (m *UserMessageBuilder) SetUserMessage(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetUserMessage General message to present to user in the Client.
func (m *UserMessagePointerBuilder) SetUserMessage(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetIsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageBuilder) SetIsPopupMessage(value bool) {
	message.SetBoolVLS(m.Unsafe(), 11, 10, value)
}

// SetIsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessagePointerBuilder) SetIsPopupMessage(value bool) {
	message.SetBoolVLS(m.Ptr, 11, 10, value)
}

// SetUserMessage General message to present to user in the Client.
func (m UserMessageFixedBuilder) SetUserMessage(value string) {
	message.SetStringFixed(m.Unsafe(), 260, 4, value)
}

// SetUserMessage General message to present to user in the Client.
func (m UserMessageFixedPointerBuilder) SetUserMessage(value string) {
	message.SetStringFixed(m.Ptr, 260, 4, value)
}

// SetIsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageFixedBuilder) SetIsPopupMessage(value bool) {
	message.SetBoolFixed(m.Unsafe(), 261, 260, value)
}

// SetIsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageFixedPointerBuilder) SetIsPopupMessage(value bool) {
	message.SetBoolFixed(m.Ptr, 261, 260, value)
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