package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                  = InterprocessSynchronizationTradeActivityRequestSize  (17)
//     Type                  = INTERPROCESS_SYNCHRONIZATION_TRADE_ACTIVITY_REQUEST  (10137)
//     RequestID             = 0
//     StartDateTimeUTC      = 0
//     SendOrderActivityOnly = 0
// }
var _InterprocessSynchronizationTradeActivityRequestDefault = []byte{17, 0, 153, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const InterprocessSynchronizationTradeActivityRequestSize = 17

type InterprocessSynchronizationTradeActivityRequest struct {
	message.Fixed
}

type InterprocessSynchronizationTradeActivityRequestBuilder struct {
	message.Fixed
}

func NewInterprocessSynchronizationTradeActivityRequestFrom(b []byte) InterprocessSynchronizationTradeActivityRequest {
	return InterprocessSynchronizationTradeActivityRequest{Fixed: message.NewFixedFrom(b)}
}

func WrapInterprocessSynchronizationTradeActivityRequest(b []byte) InterprocessSynchronizationTradeActivityRequest {
	return InterprocessSynchronizationTradeActivityRequest{Fixed: message.WrapFixed(b)}
}

func NewInterprocessSynchronizationTradeActivityRequest() InterprocessSynchronizationTradeActivityRequestBuilder {
	a := InterprocessSynchronizationTradeActivityRequestBuilder{message.NewFixed(17)}
	a.Unsafe().SetBytes(0, _InterprocessSynchronizationTradeActivityRequestDefault)
	return a
}

// Clear
// {
//     Size                  = InterprocessSynchronizationTradeActivityRequestSize  (17)
//     Type                  = INTERPROCESS_SYNCHRONIZATION_TRADE_ACTIVITY_REQUEST  (10137)
//     RequestID             = 0
//     StartDateTimeUTC      = 0
//     SendOrderActivityOnly = 0
// }
func (m InterprocessSynchronizationTradeActivityRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _InterprocessSynchronizationTradeActivityRequestDefault)
}

// ToBuilder
func (m InterprocessSynchronizationTradeActivityRequest) ToBuilder() InterprocessSynchronizationTradeActivityRequestBuilder {
	return InterprocessSynchronizationTradeActivityRequestBuilder{m.Fixed}
}

// Finish
func (m InterprocessSynchronizationTradeActivityRequestBuilder) Finish() InterprocessSynchronizationTradeActivityRequest {
	return InterprocessSynchronizationTradeActivityRequest{m.Fixed}
}

// RequestID
func (m InterprocessSynchronizationTradeActivityRequest) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m InterprocessSynchronizationTradeActivityRequestBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// StartDateTimeUTC
func (m InterprocessSynchronizationTradeActivityRequest) StartDateTimeUTC() int64 {
	return message.Int64Fixed(m.Unsafe(), 16, 8)
}

// StartDateTimeUTC
func (m InterprocessSynchronizationTradeActivityRequestBuilder) StartDateTimeUTC() int64 {
	return message.Int64Fixed(m.Unsafe(), 16, 8)
}

// SendOrderActivityOnly
func (m InterprocessSynchronizationTradeActivityRequest) SendOrderActivityOnly() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 17, 16)
}

// SendOrderActivityOnly
func (m InterprocessSynchronizationTradeActivityRequestBuilder) SendOrderActivityOnly() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 17, 16)
}

// SetRequestID
func (m InterprocessSynchronizationTradeActivityRequestBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetStartDateTimeUTC
func (m InterprocessSynchronizationTradeActivityRequestBuilder) SetStartDateTimeUTC(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, value)
}

// SetSendOrderActivityOnly
func (m InterprocessSynchronizationTradeActivityRequestBuilder) SetSendOrderActivityOnly(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 17, 16, value)
}

// Copy
func (m InterprocessSynchronizationTradeActivityRequest) Copy(to InterprocessSynchronizationTradeActivityRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTimeUTC(m.StartDateTimeUTC())
	to.SetSendOrderActivityOnly(m.SendOrderActivityOnly())
}

// Copy
func (m InterprocessSynchronizationTradeActivityRequestBuilder) Copy(to InterprocessSynchronizationTradeActivityRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTimeUTC(m.StartDateTimeUTC())
	to.SetSendOrderActivityOnly(m.SendOrderActivityOnly())
}

// CopyPointer
func (m InterprocessSynchronizationTradeActivityRequest) CopyPointer(to InterprocessSynchronizationTradeActivityRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTimeUTC(m.StartDateTimeUTC())
	to.SetSendOrderActivityOnly(m.SendOrderActivityOnly())
}

// CopyPointer
func (m InterprocessSynchronizationTradeActivityRequestBuilder) CopyPointer(to InterprocessSynchronizationTradeActivityRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTimeUTC(m.StartDateTimeUTC())
	to.SetSendOrderActivityOnly(m.SendOrderActivityOnly())
}
