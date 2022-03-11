package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HeartbeatPointer Both the Client and the Server need to send to the other side a heartbeat
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
type HeartbeatPointer struct {
	message.FixedPointer
}

// HeartbeatPointerBuilder Both the Client and the Server need to send to the other side a heartbeat
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
type HeartbeatPointerBuilder struct {
	message.FixedPointer
}

func AllocHeartbeat() HeartbeatPointerBuilder {
	a := HeartbeatPointerBuilder{message.AllocFixed(16)}
	a.Ptr.SetBytes(0, _HeartbeatDefault)
	return a
}

func AllocHeartbeatFrom(b []byte) HeartbeatPointer {
	return HeartbeatPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size               = HeartbeatSize  (16)
//     Type               = HEARTBEAT  (3)
//     NumDroppedMessages = 0
//     CurrentDateTime    = 0
// }
func (m HeartbeatPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HeartbeatDefault)
}

// ToBuilder
func (m HeartbeatPointer) ToBuilder() HeartbeatPointerBuilder {
	return HeartbeatPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HeartbeatPointerBuilder) Finish() HeartbeatPointer {
	return HeartbeatPointer{m.FixedPointer}
}

// NumDroppedMessages
func (m HeartbeatPointer) NumDroppedMessages() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// NumDroppedMessages
func (m HeartbeatPointerBuilder) NumDroppedMessages() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// CurrentDateTime
func (m HeartbeatPointer) CurrentDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 16, 8))
}

// CurrentDateTime
func (m HeartbeatPointerBuilder) CurrentDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 16, 8))
}

// SetNumDroppedMessages
func (m HeartbeatPointerBuilder) SetNumDroppedMessages(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetCurrentDateTime
func (m HeartbeatPointerBuilder) SetCurrentDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
}

// Copy
func (m HeartbeatPointer) Copy(to HeartbeatBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
}

// Copy
func (m HeartbeatPointerBuilder) Copy(to HeartbeatBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
}

// CopyPointer
func (m HeartbeatPointer) CopyPointer(to HeartbeatPointerBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
}

// CopyPointer
func (m HeartbeatPointerBuilder) CopyPointer(to HeartbeatPointerBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
}
