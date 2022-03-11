package model

import (
	"github.com/moontrade/dtc-go/message"
)

type ReplayChartDataPerformActionFixedPointer struct {
	message.FixedPointer
}

type ReplayChartDataPerformActionFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocReplayChartDataPerformActionFixed() ReplayChartDataPerformActionFixedPointerBuilder {
	a := ReplayChartDataPerformActionFixedPointerBuilder{message.AllocFixed(18)}
	a.Ptr.SetBytes(0, _ReplayChartDataPerformActionFixedDefault)
	return a
}

func AllocReplayChartDataPerformActionFixedFrom(b []byte) ReplayChartDataPerformActionFixedPointer {
	return ReplayChartDataPerformActionFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size        = ReplayChartDataPerformActionFixedSize  (18)
//     Type        = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     BaseSize    = ReplayChartDataPerformActionFixedSize  (18)
//     RequestID   = 0
//     Action      = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed = 0
// }
func (m ReplayChartDataPerformActionFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ReplayChartDataPerformActionFixedDefault)
}

// ToBuilder
func (m ReplayChartDataPerformActionFixedPointer) ToBuilder() ReplayChartDataPerformActionFixedPointerBuilder {
	return ReplayChartDataPerformActionFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *ReplayChartDataPerformActionFixedPointerBuilder) Finish() ReplayChartDataPerformActionFixedPointer {
	return ReplayChartDataPerformActionFixedPointer{m.FixedPointer}
}

// RequestID
func (m ReplayChartDataPerformActionFixedPointer) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 10, 6)
}

// RequestID
func (m ReplayChartDataPerformActionFixedPointerBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 10, 6)
}

// Action
func (m ReplayChartDataPerformActionFixedPointer) Action() ReplayChartDataActionEnum {
	return ReplayChartDataActionEnum(message.Int32Fixed(m.Ptr, 14, 10))
}

// Action
func (m ReplayChartDataPerformActionFixedPointerBuilder) Action() ReplayChartDataActionEnum {
	return ReplayChartDataActionEnum(message.Int32Fixed(m.Ptr, 14, 10))
}

// ReplaySpeed
func (m ReplayChartDataPerformActionFixedPointer) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Ptr, 18, 14)
}

// ReplaySpeed
func (m ReplayChartDataPerformActionFixedPointerBuilder) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Ptr, 18, 14)
}

// SetRequestID
func (m ReplayChartDataPerformActionFixedPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 10, 6, value)
}

// SetAction
func (m ReplayChartDataPerformActionFixedPointerBuilder) SetAction(value ReplayChartDataActionEnum) {
	message.SetInt32Fixed(m.Ptr, 14, 10, int32(value))
}

// SetReplaySpeed
func (m ReplayChartDataPerformActionFixedPointerBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32Fixed(m.Ptr, 18, 14, value)
}

// Copy
func (m ReplayChartDataPerformActionFixedPointer) Copy(to ReplayChartDataPerformActionFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// Copy
func (m ReplayChartDataPerformActionFixedPointerBuilder) Copy(to ReplayChartDataPerformActionFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyPointer
func (m ReplayChartDataPerformActionFixedPointer) CopyPointer(to ReplayChartDataPerformActionFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyPointer
func (m ReplayChartDataPerformActionFixedPointerBuilder) CopyPointer(to ReplayChartDataPerformActionFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyTo
func (m ReplayChartDataPerformActionFixedPointer) CopyTo(to ReplayChartDataPerformActionFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyTo
func (m ReplayChartDataPerformActionFixedPointerBuilder) CopyTo(to ReplayChartDataPerformActionFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyToPointer
func (m ReplayChartDataPerformActionFixedPointer) CopyToPointer(to ReplayChartDataPerformActionFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyToPointer
func (m ReplayChartDataPerformActionFixedPointerBuilder) CopyToPointer(to ReplayChartDataPerformActionFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}
