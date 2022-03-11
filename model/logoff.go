package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size           = LogoffSize  (12)
//     Type           = LOGOFF  (5)
//     BaseSize       = LogoffSize  (12)
//     Reason         = ""
//     DoNotReconnect = 0
// }
var _LogoffDefault = []byte{12, 0, 5, 0, 12, 0, 0, 0, 0, 0, 0, 0}

const LogoffSize = 12

// Logoff A Logoff is a message which can be sent either by the Client or the Server
// to the other side. It indicates that the Client or the Server is logging
// off and going to be closing the connection.
//
// When one side receives this message, it should expect the connection will
// be closed. It should not be expected that any messages will follow the
// Logoff message, and it should close the network connection and consider
// it finished. The side receiving this message can send a Logoff message
// to the other side before closing the connection.
type Logoff struct {
	message.VLS
}

// LogoffBuilder A Logoff is a message which can be sent either by the Client or the Server
// to the other side. It indicates that the Client or the Server is logging
// off and going to be closing the connection.
//
// When one side receives this message, it should expect the connection will
// be closed. It should not be expected that any messages will follow the
// Logoff message, and it should close the network connection and consider
// it finished. The side receiving this message can send a Logoff message
// to the other side before closing the connection.
type LogoffBuilder struct {
	message.VLSBuilder
}

func NewLogoffFrom(b []byte) Logoff {
	return Logoff{VLS: message.NewVLSFrom(b)}
}

func WrapLogoff(b []byte) Logoff {
	return Logoff{VLS: message.WrapVLS(b)}
}

func NewLogoff() LogoffBuilder {
	a := LogoffBuilder{message.NewVLSBuilder(12)}
	a.Unsafe().SetBytes(0, _LogoffDefault)
	return a
}

// Clear
// {
//     Size           = LogoffSize  (12)
//     Type           = LOGOFF  (5)
//     BaseSize       = LogoffSize  (12)
//     Reason         = ""
//     DoNotReconnect = 0
// }
func (m LogoffBuilder) Clear() {
	m.Unsafe().SetBytes(0, _LogoffDefault)
}

// ToBuilder
func (m Logoff) ToBuilder() LogoffBuilder {
	return LogoffBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m LogoffBuilder) Finish() Logoff {
	return Logoff{m.VLSBuilder.Finish()}
}

// Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m Logoff) Reason() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m LogoffBuilder) Reason() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// DoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m Logoff) DoNotReconnect() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// DoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffBuilder) DoNotReconnect() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// SetReason Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m *LogoffBuilder) SetReason(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetDoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffBuilder) SetDoNotReconnect(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 11, 10, value)
}

// Copy
func (m Logoff) Copy(to LogoffBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// Copy
func (m LogoffBuilder) Copy(to LogoffBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyPointer
func (m Logoff) CopyPointer(to LogoffPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyPointer
func (m LogoffBuilder) CopyPointer(to LogoffPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyToPointer
func (m Logoff) CopyToPointer(to LogoffFixedPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyToPointer
func (m LogoffBuilder) CopyToPointer(to LogoffFixedPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyTo
func (m Logoff) CopyTo(to LogoffFixedBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyTo
func (m LogoffBuilder) CopyTo(to LogoffFixedBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}
