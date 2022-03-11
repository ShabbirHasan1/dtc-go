package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size        = ReplayChartDataPerformActionFixedSize  (18)
//     Type        = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     BaseSize    = ReplayChartDataPerformActionFixedSize  (18)
//     RequestID   = 0
//     Action      = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed = 0
// }
var _ReplayChartDataPerformActionFixedDefault = []byte{18, 0, 121, 39, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const ReplayChartDataPerformActionFixedSize = 18

type ReplayChartDataPerformActionFixed struct {
	message.Fixed
}

type ReplayChartDataPerformActionFixedBuilder struct {
	message.Fixed
}

func NewReplayChartDataPerformActionFixedFrom(b []byte) ReplayChartDataPerformActionFixed {
	return ReplayChartDataPerformActionFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapReplayChartDataPerformActionFixed(b []byte) ReplayChartDataPerformActionFixed {
	return ReplayChartDataPerformActionFixed{Fixed: message.WrapFixed(b)}
}

func NewReplayChartDataPerformActionFixed() ReplayChartDataPerformActionFixedBuilder {
	a := ReplayChartDataPerformActionFixedBuilder{message.NewFixed(18)}
	a.Unsafe().SetBytes(0, _ReplayChartDataPerformActionFixedDefault)
	return a
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
func (m ReplayChartDataPerformActionFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ReplayChartDataPerformActionFixedDefault)
}

// ToBuilder
func (m ReplayChartDataPerformActionFixed) ToBuilder() ReplayChartDataPerformActionFixedBuilder {
	return ReplayChartDataPerformActionFixedBuilder{m.Fixed}
}

// Finish
func (m ReplayChartDataPerformActionFixedBuilder) Finish() ReplayChartDataPerformActionFixed {
	return ReplayChartDataPerformActionFixed{m.Fixed}
}

// RequestID
func (m ReplayChartDataPerformActionFixed) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 10, 6)
}

// RequestID
func (m ReplayChartDataPerformActionFixedBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 10, 6)
}

// Action
func (m ReplayChartDataPerformActionFixed) Action() ReplayChartDataActionEnum {
	return ReplayChartDataActionEnum(message.Int32Fixed(m.Unsafe(), 14, 10))
}

// Action
func (m ReplayChartDataPerformActionFixedBuilder) Action() ReplayChartDataActionEnum {
	return ReplayChartDataActionEnum(message.Int32Fixed(m.Unsafe(), 14, 10))
}

// ReplaySpeed
func (m ReplayChartDataPerformActionFixed) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Unsafe(), 18, 14)
}

// ReplaySpeed
func (m ReplayChartDataPerformActionFixedBuilder) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Unsafe(), 18, 14)
}

// SetRequestID
func (m ReplayChartDataPerformActionFixedBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 10, 6, value)
}

// SetAction
func (m ReplayChartDataPerformActionFixedBuilder) SetAction(value ReplayChartDataActionEnum) {
	message.SetInt32Fixed(m.Unsafe(), 14, 10, int32(value))
}

// SetReplaySpeed
func (m ReplayChartDataPerformActionFixedBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 18, 14, value)
}

// Copy
func (m ReplayChartDataPerformActionFixed) Copy(to ReplayChartDataPerformActionFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// Copy
func (m ReplayChartDataPerformActionFixedBuilder) Copy(to ReplayChartDataPerformActionFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyPointer
func (m ReplayChartDataPerformActionFixed) CopyPointer(to ReplayChartDataPerformActionFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyPointer
func (m ReplayChartDataPerformActionFixedBuilder) CopyPointer(to ReplayChartDataPerformActionFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyToPointer
func (m ReplayChartDataPerformActionFixed) CopyToPointer(to ReplayChartDataPerformActionFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyToPointer
func (m ReplayChartDataPerformActionFixedBuilder) CopyToPointer(to ReplayChartDataPerformActionFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyTo
func (m ReplayChartDataPerformActionFixed) CopyTo(to ReplayChartDataPerformActionFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyTo
func (m ReplayChartDataPerformActionFixedBuilder) CopyTo(to ReplayChartDataPerformActionFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}
