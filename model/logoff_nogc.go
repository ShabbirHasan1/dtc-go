package model

import (
	"github.com/moontrade/dtc-go/message"
)

// LogoffPointer A Logoff is a message which can be sent either by the Client or the Server
// to the other side. It indicates that the Client or the Server is logging
// off and going to be closing the connection.
//
// When one side receives this message, it should expect the connection will
// be closed. It should not be expected that any messages will follow the
// Logoff message, and it should close the network connection and consider
// it finished. The side receiving this message can send a Logoff message
// to the other side before closing the connection.
type LogoffPointer struct {
	message.VLSPointer
}

// LogoffPointerBuilder A Logoff is a message which can be sent either by the Client or the Server
// to the other side. It indicates that the Client or the Server is logging
// off and going to be closing the connection.
//
// When one side receives this message, it should expect the connection will
// be closed. It should not be expected that any messages will follow the
// Logoff message, and it should close the network connection and consider
// it finished. The side receiving this message can send a Logoff message
// to the other side before closing the connection.
type LogoffPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocLogoff() LogoffPointerBuilder {
	a := LogoffPointerBuilder{message.AllocVLSBuilder(12)}
	a.Ptr.SetBytes(0, _LogoffDefault)
	return a
}

func AllocLogoffFrom(b []byte) LogoffPointer {
	return LogoffPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size           = LogoffSize  (12)
//     Type           = LOGOFF  (5)
//     BaseSize       = LogoffSize  (12)
//     Reason         = ""
//     DoNotReconnect = 0
// }
func (m LogoffPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _LogoffDefault)
}

// ToBuilder
func (m LogoffPointer) ToBuilder() LogoffPointerBuilder {
	return LogoffPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *LogoffPointerBuilder) Finish() LogoffPointer {
	return LogoffPointer{m.VLSPointerBuilder.Finish()}
}

// Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m LogoffPointer) Reason() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m LogoffPointerBuilder) Reason() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// DoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffPointer) DoNotReconnect() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// DoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffPointerBuilder) DoNotReconnect() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// SetReason Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m *LogoffPointerBuilder) SetReason(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetDoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffPointerBuilder) SetDoNotReconnect(value uint8) {
	message.SetUint8VLS(m.Ptr, 11, 10, value)
}

// Copy
func (m LogoffPointer) Copy(to LogoffBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// Copy
func (m LogoffPointerBuilder) Copy(to LogoffBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyPointer
func (m LogoffPointer) CopyPointer(to LogoffPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyPointer
func (m LogoffPointerBuilder) CopyPointer(to LogoffPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyTo
func (m LogoffPointer) CopyTo(to LogoffFixedBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyTo
func (m LogoffPointerBuilder) CopyTo(to LogoffFixedBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyToPointer
func (m LogoffPointer) CopyToPointer(to LogoffFixedPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyToPointer
func (m LogoffPointerBuilder) CopyToPointer(to LogoffFixedPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}
