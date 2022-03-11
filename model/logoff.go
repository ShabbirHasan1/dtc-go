package model

import (
	"github.com/moontrade/dtc-go/message"
)

const LogoffSize = 12

const LogoffFixedSize = 102

// {
//     Size           = LogoffSize  (12)
//     Type           = LOGOFF  (5)
//     BaseSize       = LogoffSize  (12)
//     Reason         = ""
//     DoNotReconnect = 0
// }
var _LogoffDefault = []byte{12, 0, 5, 0, 12, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size           = LogoffFixedSize  (102)
//     Type           = LOGOFF  (5)
//     Reason         = ""
//     DoNotReconnect = 0
// }
var _LogoffFixedDefault = []byte{102, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocLogoff() LogoffPointerBuilder {
	a := LogoffPointerBuilder{message.AllocVLSBuilder(12)}
	a.Ptr.SetBytes(0, _LogoffDefault)
	return a
}

func AllocLogoffFrom(b []byte) LogoffPointer {
	return LogoffPointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size           = LogoffSize  (12)
//     Type           = LOGOFF  (5)
//     BaseSize       = LogoffSize  (12)
//     Reason         = ""
//     DoNotReconnect = 0
// }
func (m LogoffBuilder) Clear() {
	m.Unsafe().SetBytes(0, _LogoffDefault)
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
func (m Logoff) ToBuilder() LogoffBuilder {
	return LogoffBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m LogoffFixed) ToBuilder() LogoffFixedBuilder {
	return LogoffFixedBuilder{m.Fixed}
}

// ToBuilder
func (m LogoffPointer) ToBuilder() LogoffPointerBuilder {
	return LogoffPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m LogoffFixedPointer) ToBuilder() LogoffFixedPointerBuilder {
	return LogoffFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m LogoffBuilder) Finish() Logoff {
	return Logoff{m.VLSBuilder.Finish()}
}

// Finish
func (m LogoffFixedBuilder) Finish() LogoffFixed {
	return LogoffFixed{m.Fixed}
}

// Finish
func (m *LogoffPointerBuilder) Finish() LogoffPointer {
	return LogoffPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *LogoffFixedPointerBuilder) Finish() LogoffFixedPointer {
	return LogoffFixedPointer{m.FixedPointer}
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
func (m Logoff) DoNotReconnect() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// DoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffBuilder) DoNotReconnect() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
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
func (m LogoffFixed) DoNotReconnect() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 101, 100)
}

// DoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffFixedBuilder) DoNotReconnect() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 101, 100)
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
func (m *LogoffBuilder) SetReason(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetReason Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m *LogoffPointerBuilder) SetReason(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetDoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffBuilder) SetDoNotReconnect(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 11, 10, value)
}

// SetDoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffPointerBuilder) SetDoNotReconnect(value uint8) {
	message.SetUint8VLS(m.Ptr, 11, 10, value)
}

// SetReason Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m LogoffFixedBuilder) SetReason(value string) {
	message.SetStringFixed(m.Unsafe(), 100, 4, value)
}

// SetReason Reason is a character string indicating the reason for the log off from
// either the Client or the Server.
func (m LogoffFixedPointerBuilder) SetReason(value string) {
	message.SetStringFixed(m.Ptr, 100, 4, value)
}

// SetDoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffFixedBuilder) SetDoNotReconnect(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 101, 100, value)
}

// SetDoNotReconnect When DoNotReconnect is set to a 1, this indicates to the other side that
// a reconnect to the opposite side should not occur automatically.
func (m LogoffFixedPointerBuilder) SetDoNotReconnect(value uint8) {
	message.SetUint8Fixed(m.Ptr, 101, 100, value)
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
