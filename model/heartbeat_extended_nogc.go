package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HeartbeatExtendedPointer struct {
	message.FixedPointer
}

type HeartbeatExtendedPointerBuilder struct {
	message.FixedPointer
}

func AllocHeartbeatExtended() HeartbeatExtendedPointerBuilder {
	a := HeartbeatExtendedPointerBuilder{message.AllocFixed(62)}
	a.Ptr.SetBytes(0, _HeartbeatExtendedDefault)
	return a
}

func AllocHeartbeatExtendedFrom(b []byte) HeartbeatExtendedPointer {
	return HeartbeatExtendedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                                   = HeartbeatExtendedSize  (62)
//     Type                                   = HEARTBEAT  (3)
//     NumDroppedMessages                     = 0
//     CurrentDateTime                        = 0
//     SecondsSinceLastReceivedHeartbeat      = 0
//     NumberOfOutstandingSentBuffers         = 0
//     PendingTransmissionDelayInMilliseconds = 0
//     CurrentSendBufferSizeInBytes           = 0
//     SendingDateTimeMicroseconds            = 0
//     DataCompressionRatio                   = 0
//     TotalUncompressedBytes                 = 0
//     TotalCompressionTime                   = 0
//     NumberOfCompressions                   = 0
//     SourceIPAddress                        = 0
// }
func (m HeartbeatExtendedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HeartbeatExtendedDefault)
}

// ToBuilder
func (m HeartbeatExtendedPointer) ToBuilder() HeartbeatExtendedPointerBuilder {
	return HeartbeatExtendedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HeartbeatExtendedPointerBuilder) Finish() HeartbeatExtendedPointer {
	return HeartbeatExtendedPointer{m.FixedPointer}
}

// NumDroppedMessages
func (m HeartbeatExtendedPointer) NumDroppedMessages() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// NumDroppedMessages
func (m HeartbeatExtendedPointerBuilder) NumDroppedMessages() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// CurrentDateTime
func (m HeartbeatExtendedPointer) CurrentDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 16, 8))
}

// CurrentDateTime
func (m HeartbeatExtendedPointerBuilder) CurrentDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 16, 8))
}

// SecondsSinceLastReceivedHeartbeat
func (m HeartbeatExtendedPointer) SecondsSinceLastReceivedHeartbeat() uint16 {
	return message.Uint16Fixed(m.Ptr, 18, 16)
}

// SecondsSinceLastReceivedHeartbeat
func (m HeartbeatExtendedPointerBuilder) SecondsSinceLastReceivedHeartbeat() uint16 {
	return message.Uint16Fixed(m.Ptr, 18, 16)
}

// NumberOfOutstandingSentBuffers
func (m HeartbeatExtendedPointer) NumberOfOutstandingSentBuffers() uint16 {
	return message.Uint16Fixed(m.Ptr, 20, 18)
}

// NumberOfOutstandingSentBuffers
func (m HeartbeatExtendedPointerBuilder) NumberOfOutstandingSentBuffers() uint16 {
	return message.Uint16Fixed(m.Ptr, 20, 18)
}

// PendingTransmissionDelayInMilliseconds
func (m HeartbeatExtendedPointer) PendingTransmissionDelayInMilliseconds() uint16 {
	return message.Uint16Fixed(m.Ptr, 22, 20)
}

// PendingTransmissionDelayInMilliseconds
func (m HeartbeatExtendedPointerBuilder) PendingTransmissionDelayInMilliseconds() uint16 {
	return message.Uint16Fixed(m.Ptr, 22, 20)
}

// CurrentSendBufferSizeInBytes
func (m HeartbeatExtendedPointer) CurrentSendBufferSizeInBytes() uint32 {
	return message.Uint32Fixed(m.Ptr, 26, 22)
}

// CurrentSendBufferSizeInBytes
func (m HeartbeatExtendedPointerBuilder) CurrentSendBufferSizeInBytes() uint32 {
	return message.Uint32Fixed(m.Ptr, 26, 22)
}

// SendingDateTimeMicroseconds
func (m HeartbeatExtendedPointer) SendingDateTimeMicroseconds() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 34, 26))
}

// SendingDateTimeMicroseconds
func (m HeartbeatExtendedPointerBuilder) SendingDateTimeMicroseconds() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 34, 26))
}

// DataCompressionRatio
func (m HeartbeatExtendedPointer) DataCompressionRatio() float32 {
	return message.Float32Fixed(m.Ptr, 38, 34)
}

// DataCompressionRatio
func (m HeartbeatExtendedPointerBuilder) DataCompressionRatio() float32 {
	return message.Float32Fixed(m.Ptr, 38, 34)
}

// TotalUncompressedBytes
func (m HeartbeatExtendedPointer) TotalUncompressedBytes() uint64 {
	return message.Uint64Fixed(m.Ptr, 46, 38)
}

// TotalUncompressedBytes
func (m HeartbeatExtendedPointerBuilder) TotalUncompressedBytes() uint64 {
	return message.Uint64Fixed(m.Ptr, 46, 38)
}

// TotalCompressionTime
func (m HeartbeatExtendedPointer) TotalCompressionTime() float64 {
	return message.Float64Fixed(m.Ptr, 54, 46)
}

// TotalCompressionTime
func (m HeartbeatExtendedPointerBuilder) TotalCompressionTime() float64 {
	return message.Float64Fixed(m.Ptr, 54, 46)
}

// NumberOfCompressions
func (m HeartbeatExtendedPointer) NumberOfCompressions() uint32 {
	return message.Uint32Fixed(m.Ptr, 58, 54)
}

// NumberOfCompressions
func (m HeartbeatExtendedPointerBuilder) NumberOfCompressions() uint32 {
	return message.Uint32Fixed(m.Ptr, 58, 54)
}

// SourceIPAddress
func (m HeartbeatExtendedPointer) SourceIPAddress() uint32 {
	return message.Uint32Fixed(m.Ptr, 62, 58)
}

// SourceIPAddress
func (m HeartbeatExtendedPointerBuilder) SourceIPAddress() uint32 {
	return message.Uint32Fixed(m.Ptr, 62, 58)
}

// SetNumDroppedMessages
func (m HeartbeatExtendedPointerBuilder) SetNumDroppedMessages(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetCurrentDateTime
func (m HeartbeatExtendedPointerBuilder) SetCurrentDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
}

// SetSecondsSinceLastReceivedHeartbeat
func (m HeartbeatExtendedPointerBuilder) SetSecondsSinceLastReceivedHeartbeat(value uint16) {
	message.SetUint16Fixed(m.Ptr, 18, 16, value)
}

// SetNumberOfOutstandingSentBuffers
func (m HeartbeatExtendedPointerBuilder) SetNumberOfOutstandingSentBuffers(value uint16) {
	message.SetUint16Fixed(m.Ptr, 20, 18, value)
}

// SetPendingTransmissionDelayInMilliseconds
func (m HeartbeatExtendedPointerBuilder) SetPendingTransmissionDelayInMilliseconds(value uint16) {
	message.SetUint16Fixed(m.Ptr, 22, 20, value)
}

// SetCurrentSendBufferSizeInBytes
func (m HeartbeatExtendedPointerBuilder) SetCurrentSendBufferSizeInBytes(value uint32) {
	message.SetUint32Fixed(m.Ptr, 26, 22, value)
}

// SetSendingDateTimeMicroseconds
func (m HeartbeatExtendedPointerBuilder) SetSendingDateTimeMicroseconds(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 34, 26, int64(value))
}

// SetDataCompressionRatio
func (m HeartbeatExtendedPointerBuilder) SetDataCompressionRatio(value float32) {
	message.SetFloat32Fixed(m.Ptr, 38, 34, value)
}

// SetTotalUncompressedBytes
func (m HeartbeatExtendedPointerBuilder) SetTotalUncompressedBytes(value uint64) {
	message.SetUint64Fixed(m.Ptr, 46, 38, value)
}

// SetTotalCompressionTime
func (m HeartbeatExtendedPointerBuilder) SetTotalCompressionTime(value float64) {
	message.SetFloat64Fixed(m.Ptr, 54, 46, value)
}

// SetNumberOfCompressions
func (m HeartbeatExtendedPointerBuilder) SetNumberOfCompressions(value uint32) {
	message.SetUint32Fixed(m.Ptr, 58, 54, value)
}

// SetSourceIPAddress
func (m HeartbeatExtendedPointerBuilder) SetSourceIPAddress(value uint32) {
	message.SetUint32Fixed(m.Ptr, 62, 58, value)
}

// Copy
func (m HeartbeatExtendedPointer) Copy(to HeartbeatExtendedBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
	to.SetSecondsSinceLastReceivedHeartbeat(m.SecondsSinceLastReceivedHeartbeat())
	to.SetNumberOfOutstandingSentBuffers(m.NumberOfOutstandingSentBuffers())
	to.SetPendingTransmissionDelayInMilliseconds(m.PendingTransmissionDelayInMilliseconds())
	to.SetCurrentSendBufferSizeInBytes(m.CurrentSendBufferSizeInBytes())
	to.SetSendingDateTimeMicroseconds(m.SendingDateTimeMicroseconds())
	to.SetDataCompressionRatio(m.DataCompressionRatio())
	to.SetTotalUncompressedBytes(m.TotalUncompressedBytes())
	to.SetTotalCompressionTime(m.TotalCompressionTime())
	to.SetNumberOfCompressions(m.NumberOfCompressions())
	to.SetSourceIPAddress(m.SourceIPAddress())
}

// Copy
func (m HeartbeatExtendedPointerBuilder) Copy(to HeartbeatExtendedBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
	to.SetSecondsSinceLastReceivedHeartbeat(m.SecondsSinceLastReceivedHeartbeat())
	to.SetNumberOfOutstandingSentBuffers(m.NumberOfOutstandingSentBuffers())
	to.SetPendingTransmissionDelayInMilliseconds(m.PendingTransmissionDelayInMilliseconds())
	to.SetCurrentSendBufferSizeInBytes(m.CurrentSendBufferSizeInBytes())
	to.SetSendingDateTimeMicroseconds(m.SendingDateTimeMicroseconds())
	to.SetDataCompressionRatio(m.DataCompressionRatio())
	to.SetTotalUncompressedBytes(m.TotalUncompressedBytes())
	to.SetTotalCompressionTime(m.TotalCompressionTime())
	to.SetNumberOfCompressions(m.NumberOfCompressions())
	to.SetSourceIPAddress(m.SourceIPAddress())
}

// CopyPointer
func (m HeartbeatExtendedPointer) CopyPointer(to HeartbeatExtendedPointerBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
	to.SetSecondsSinceLastReceivedHeartbeat(m.SecondsSinceLastReceivedHeartbeat())
	to.SetNumberOfOutstandingSentBuffers(m.NumberOfOutstandingSentBuffers())
	to.SetPendingTransmissionDelayInMilliseconds(m.PendingTransmissionDelayInMilliseconds())
	to.SetCurrentSendBufferSizeInBytes(m.CurrentSendBufferSizeInBytes())
	to.SetSendingDateTimeMicroseconds(m.SendingDateTimeMicroseconds())
	to.SetDataCompressionRatio(m.DataCompressionRatio())
	to.SetTotalUncompressedBytes(m.TotalUncompressedBytes())
	to.SetTotalCompressionTime(m.TotalCompressionTime())
	to.SetNumberOfCompressions(m.NumberOfCompressions())
	to.SetSourceIPAddress(m.SourceIPAddress())
}

// CopyPointer
func (m HeartbeatExtendedPointerBuilder) CopyPointer(to HeartbeatExtendedPointerBuilder) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
	to.SetSecondsSinceLastReceivedHeartbeat(m.SecondsSinceLastReceivedHeartbeat())
	to.SetNumberOfOutstandingSentBuffers(m.NumberOfOutstandingSentBuffers())
	to.SetPendingTransmissionDelayInMilliseconds(m.PendingTransmissionDelayInMilliseconds())
	to.SetCurrentSendBufferSizeInBytes(m.CurrentSendBufferSizeInBytes())
	to.SetSendingDateTimeMicroseconds(m.SendingDateTimeMicroseconds())
	to.SetDataCompressionRatio(m.DataCompressionRatio())
	to.SetTotalUncompressedBytes(m.TotalUncompressedBytes())
	to.SetTotalCompressionTime(m.TotalCompressionTime())
	to.SetNumberOfCompressions(m.NumberOfCompressions())
	to.SetSourceIPAddress(m.SourceIPAddress())
}
