package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size               = HeartbeatSize  (16)
//     Type               = HEARTBEAT  (3)
//     NumDroppedMessages = 0
//     CurrentDateTime    = 0
// }
var _HeartbeatDefault = []byte{16, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HeartbeatSize = 16

// Heartbeat Both the Client and the Server need to send to the other side a heartbeat
// at the interval specified by the HeartbeatIntervalInSeconds member in
// the LogonRequest.
//
// There are no required member fields to set in this message. The purpose
// of the Heartbeat message is so that the Client or the Server can determine
// whether the other side is still connected.
//
// It is recommended that if there is a loss of Heartbeat messages from the
// other side, for twice the amount of the HeartbeatIntervalInSeconds time
// that it is safe to assume that the other side is no longer present and
// the network socket should be then gracefully closed.
//
// The Server may choose to send a heartbeat message every second to the
// Client. In this particular case, it is recommended the Client use a minimum
// time of about 5 to 10 seconds without a heartbeat to determine the loss
// of the connection rather than the standard of twice the amount of the
// heartbeat time interval.
type Heartbeat struct {
	message.Fixed
}

// HeartbeatBuilder Both the Client and the Server need to send to the other side a heartbeat
// at the interval specified by the HeartbeatIntervalInSeconds member in
// the LogonRequest.
//
// There are no required member fields to set in this message. The purpose
// of the Heartbeat message is so that the Client or the Server can determine
// whether the other side is still connected.
//
// It is recommended that if there is a loss of Heartbeat messages from the
// other side, for twice the amount of the HeartbeatIntervalInSeconds time
// that it is safe to assume that the other side is no longer present and
// the network socket should be then gracefully closed.
//
// The Server may choose to send a heartbeat message every second to the
// Client. In this particular case, it is recommended the Client use a minimum
// time of about 5 to 10 seconds without a heartbeat to determine the loss
// of the connection rather than the standard of twice the amount of the
// heartbeat time interval.
type HeartbeatBuilder struct {
	message.Fixed
}

func NewHeartbeatFrom(b []byte) Heartbeat {
	return Heartbeat{Fixed: message.NewFixedFrom(b)}
}

func WrapHeartbeat(b []byte) Heartbeat {
	return Heartbeat{Fixed: message.WrapFixed(b)}
}

func NewHeartbeat() HeartbeatBuilder {
	a := HeartbeatBuilder{message.NewFixed(16)}
	a.Unsafe().SetBytes(0, _HeartbeatDefault)
	return a
}

// Clear
// {
//     Size               = HeartbeatSize  (16)
//     Type               = HEARTBEAT  (3)
//     NumDroppedMessages = 0
//     CurrentDateTime    = 0
// }
func (m HeartbeatBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HeartbeatDefault)
}

// ToBuilder
func (m Heartbeat) ToBuilder() HeartbeatBuilder {
	return HeartbeatBuilder{m.Fixed}
}

// Finish
func (m HeartbeatBuilder) Finish() Heartbeat {
	return Heartbeat{m.Fixed}
}

// NumDroppedMessages
func (m Heartbeat) NumDroppedMessages() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// NumDroppedMessages
func (m HeartbeatBuilder) NumDroppedMessages() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// CurrentDateTime
func (m Heartbeat) CurrentDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// CurrentDateTime
func (m HeartbeatBuilder) CurrentDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// SetNumDroppedMessages
func (m HeartbeatBuilder) SetNumDroppedMessages(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetCurrentDateTime
func (m HeartbeatBuilder) SetCurrentDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, int64(value))
}

// Copy
func (m Heartbeat) Copy(to HeartbeatBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
}

// Copy
func (m HeartbeatBuilder) Copy(to HeartbeatBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
}

// CopyPointer
func (m Heartbeat) CopyPointer(to HeartbeatPointerBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
}

// CopyPointer
func (m HeartbeatBuilder) CopyPointer(to HeartbeatPointerBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
}
