package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size           = LogoffFixedSize  (102)
//     Type           = LOGOFF  (5)
//     Reason         = ""
//     DoNotReconnect = 0
// }
var _LogoffFixedDefault = []byte{102, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const LogoffFixedSize = 102

// LogoffFixed A Logoff is a message which can be sent either by the Client or the Server
// to the other side. It indicates that the Client or the Server is logging
// off and going to be closing the connection.
//
// When one side receives this message, it should expect the connection will
// be closed. It should not be expected that any messages will follow the
// Logoff message, and it should close the network connection and consider
// it finished. The side receiving this message can send a Logoff message
// to the other side before closing the connection.
type LogoffFixed struct {
	message.Fixed
}

// LogoffFixedBuilder A Logoff is a message which can be sent either by the Client or the Server
// to the other side. It indicates that the Client or the Server is logging
// off and going to be closing the connection.
//
// When one side receives this message, it should expect the connection will
// be closed. It should not be expected that any messages will follow the
// Logoff message, and it should close the network connection and consider
// it finished. The side receiving this message can send a Logoff message
// to the other side before closing the connection.
type LogoffFixedBuilder struct {
	message.Fixed
}

func NewLogoffFixedFrom(b []byte) LogoffFixed {
	return LogoffFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapLogoffFixed(b []byte) LogoffFixed {
	return LogoffFixed{Fixed: message.WrapFixed(b)}
}

func NewLogoffFixed() LogoffFixedBuilder {
	a := LogoffFixedBuilder{message.NewFixed(102)}
	a.Unsafe().SetBytes(0, _LogoffFixedDefault)
	return a
}

// Clear
// {
//     Size           = LogoffFixedSize  (102)
//     Type           = LOGOFF  (5)
//     Reason         = ""
//     DoNotReconnect = 0
// }
func (m LogoffFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _LogoffFixedDefault)
}

// ToBuilder
func (m LogoffFixed) ToBuilder() LogoffFixedBuilder {
	return LogoffFixedBuilder{m.Fixed}
}

// Finish
func (m LogoffFixedBuilder) Finish() LogoffFixed {
	return LogoffFixed{m.Fixed}
}

// Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m LogoffFixed) Reason() string {
	return message.StringFixed(m.Unsafe(), 100, 4)
}

// Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m LogoffFixedBuilder) Reason() string {
	return message.StringFixed(m.Unsafe(), 100, 4)
}

// DoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffFixed) DoNotReconnect() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 101, 100)
}

// DoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffFixedBuilder) DoNotReconnect() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 101, 100)
}

// SetReason Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m LogoffFixedBuilder) SetReason(value string) {
	message.SetStringFixed(m.Unsafe(), 100, 4, value)
}

// SetDoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffFixedBuilder) SetDoNotReconnect(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 101, 100, value)
}

// Copy
func (m LogoffFixed) Copy(to LogoffFixedBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// Copy
func (m LogoffFixedBuilder) Copy(to LogoffFixedBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyPointer
func (m LogoffFixed) CopyPointer(to LogoffFixedPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyPointer
func (m LogoffFixedBuilder) CopyPointer(to LogoffFixedPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyToPointer
func (m LogoffFixed) CopyToPointer(to LogoffPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyToPointer
func (m LogoffFixedBuilder) CopyToPointer(to LogoffPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyTo
func (m LogoffFixed) CopyTo(to LogoffBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyTo
func (m LogoffFixedBuilder) CopyTo(to LogoffBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}
