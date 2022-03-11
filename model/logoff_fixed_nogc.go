package model

import (
	"github.com/moontrade/dtc-go/message"
)

// LogoffFixedPointer A Logoff is a message which can be sent either by the Client or the Server
// to the other side. It indicates that the Client or the Server is logging
// off and going to be closing the connection.
//
// When one side receives this message, it should expect the connection will
// be closed. It should not be expected that any messages will follow the
// Logoff message, and it should close the network connection and consider
// it finished. The side receiving this message can send a Logoff message
// to the other side before closing the connection.
type LogoffFixedPointer struct {
	message.FixedPointer
}

// LogoffFixedPointerBuilder A Logoff is a message which can be sent either by the Client or the Server
// to the other side. It indicates that the Client or the Server is logging
// off and going to be closing the connection.
//
// When one side receives this message, it should expect the connection will
// be closed. It should not be expected that any messages will follow the
// Logoff message, and it should close the network connection and consider
// it finished. The side receiving this message can send a Logoff message
// to the other side before closing the connection.
type LogoffFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocLogoffFixed() LogoffFixedPointerBuilder {
	a := LogoffFixedPointerBuilder{message.AllocFixed(102)}
	a.Ptr.SetBytes(0, _LogoffFixedDefault)
	return a
}

func AllocLogoffFixedFrom(b []byte) LogoffFixedPointer {
	return LogoffFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size           = LogoffFixedSize  (102)
//     Type           = LOGOFF  (5)
//     Reason         = ""
//     DoNotReconnect = 0
// }
func (m LogoffFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _LogoffFixedDefault)
}

// ToBuilder
func (m LogoffFixedPointer) ToBuilder() LogoffFixedPointerBuilder {
	return LogoffFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *LogoffFixedPointerBuilder) Finish() LogoffFixedPointer {
	return LogoffFixedPointer{m.FixedPointer}
}

// Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m LogoffFixedPointer) Reason() string {
	return message.StringFixed(m.Ptr, 100, 4)
}

// Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m LogoffFixedPointerBuilder) Reason() string {
	return message.StringFixed(m.Ptr, 100, 4)
}

// DoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffFixedPointer) DoNotReconnect() uint8 {
	return message.Uint8Fixed(m.Ptr, 101, 100)
}

// DoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffFixedPointerBuilder) DoNotReconnect() uint8 {
	return message.Uint8Fixed(m.Ptr, 101, 100)
}

// SetReason Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m LogoffFixedPointerBuilder) SetReason(value string) {
	message.SetStringFixed(m.Ptr, 100, 4, value)
}

// SetDoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffFixedPointerBuilder) SetDoNotReconnect(value uint8) {
	message.SetUint8Fixed(m.Ptr, 101, 100, value)
}

// Copy
func (m LogoffFixedPointer) Copy(to LogoffFixedBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// Copy
func (m LogoffFixedPointerBuilder) Copy(to LogoffFixedBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyPointer
func (m LogoffFixedPointer) CopyPointer(to LogoffFixedPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyPointer
func (m LogoffFixedPointerBuilder) CopyPointer(to LogoffFixedPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyTo
func (m LogoffFixedPointer) CopyTo(to LogoffBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyTo
func (m LogoffFixedPointerBuilder) CopyTo(to LogoffBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyToPointer
func (m LogoffFixedPointer) CopyToPointer(to LogoffPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}

// CopyToPointer
func (m LogoffFixedPointerBuilder) CopyToPointer(to LogoffPointerBuilder) {
	to.SetReason(m.Reason())
	to.SetDoNotReconnect(m.DoNotReconnect())
}
