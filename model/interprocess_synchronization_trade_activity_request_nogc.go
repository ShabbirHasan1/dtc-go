package model

import (
	"github.com/moontrade/dtc-go/message"
)

type InterprocessSynchronizationTradeActivityRequestPointer struct {
	message.FixedPointer
}

type InterprocessSynchronizationTradeActivityRequestPointerBuilder struct {
	message.FixedPointer
}

func AllocInterprocessSynchronizationTradeActivityRequest() InterprocessSynchronizationTradeActivityRequestPointerBuilder {
	a := InterprocessSynchronizationTradeActivityRequestPointerBuilder{message.AllocFixed(17)}
	a.Ptr.SetBytes(0, _InterprocessSynchronizationTradeActivityRequestDefault)
	return a
}

func AllocInterprocessSynchronizationTradeActivityRequestFrom(b []byte) InterprocessSynchronizationTradeActivityRequestPointer {
	return InterprocessSynchronizationTradeActivityRequestPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                  = InterprocessSynchronizationTradeActivityRequestSize  (17)
//     Type                  = INTERPROCESS_SYNCHRONIZATION_TRADE_ACTIVITY_REQUEST  (10137)
//     RequestID             = 0
//     StartDateTimeUTC      = 0
//     SendOrderActivityOnly = 0
// }
func (m InterprocessSynchronizationTradeActivityRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _InterprocessSynchronizationTradeActivityRequestDefault)
}

// ToBuilder
func (m InterprocessSynchronizationTradeActivityRequestPointer) ToBuilder() InterprocessSynchronizationTradeActivityRequestPointerBuilder {
	return InterprocessSynchronizationTradeActivityRequestPointerBuilder{m.FixedPointer}
}

// Finish
func (m *InterprocessSynchronizationTradeActivityRequestPointerBuilder) Finish() InterprocessSynchronizationTradeActivityRequestPointer {
	return InterprocessSynchronizationTradeActivityRequestPointer{m.FixedPointer}
}

// RequestID
func (m InterprocessSynchronizationTradeActivityRequestPointer) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m InterprocessSynchronizationTradeActivityRequestPointerBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// StartDateTimeUTC
func (m InterprocessSynchronizationTradeActivityRequestPointer) StartDateTimeUTC() int64 {
	return message.Int64Fixed(m.Ptr, 16, 8)
}

// StartDateTimeUTC
func (m InterprocessSynchronizationTradeActivityRequestPointerBuilder) StartDateTimeUTC() int64 {
	return message.Int64Fixed(m.Ptr, 16, 8)
}

// SendOrderActivityOnly
func (m InterprocessSynchronizationTradeActivityRequestPointer) SendOrderActivityOnly() uint8 {
	return message.Uint8Fixed(m.Ptr, 17, 16)
}

// SendOrderActivityOnly
func (m InterprocessSynchronizationTradeActivityRequestPointerBuilder) SendOrderActivityOnly() uint8 {
	return message.Uint8Fixed(m.Ptr, 17, 16)
}

// SetRequestID
func (m InterprocessSynchronizationTradeActivityRequestPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetStartDateTimeUTC
func (m InterprocessSynchronizationTradeActivityRequestPointerBuilder) SetStartDateTimeUTC(value int64) {
	message.SetInt64Fixed(m.Ptr, 16, 8, value)
}

// SetSendOrderActivityOnly
func (m InterprocessSynchronizationTradeActivityRequestPointerBuilder) SetSendOrderActivityOnly(value uint8) {
	message.SetUint8Fixed(m.Ptr, 17, 16, value)
}

// Copy
func (m InterprocessSynchronizationTradeActivityRequestPointer) Copy(to InterprocessSynchronizationTradeActivityRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTimeUTC(m.StartDateTimeUTC())
	to.SetSendOrderActivityOnly(m.SendOrderActivityOnly())
}

// Copy
func (m InterprocessSynchronizationTradeActivityRequestPointerBuilder) Copy(to InterprocessSynchronizationTradeActivityRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTimeUTC(m.StartDateTimeUTC())
	to.SetSendOrderActivityOnly(m.SendOrderActivityOnly())
}

// CopyPointer
func (m InterprocessSynchronizationTradeActivityRequestPointer) CopyPointer(to InterprocessSynchronizationTradeActivityRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTimeUTC(m.StartDateTimeUTC())
	to.SetSendOrderActivityOnly(m.SendOrderActivityOnly())
}

// CopyPointer
func (m InterprocessSynchronizationTradeActivityRequestPointerBuilder) CopyPointer(to InterprocessSynchronizationTradeActivityRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTimeUTC(m.StartDateTimeUTC())
	to.SetSendOrderActivityOnly(m.SendOrderActivityOnly())
}
