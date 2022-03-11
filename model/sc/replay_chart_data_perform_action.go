package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const ReplayChartDataPerformActionSize = 18

const ReplayChartDataPerformActionFixedSize = 16

// {
//     Size        = ReplayChartDataPerformActionSize  (18)
//     Type        = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     BaseSize    = ReplayChartDataPerformActionSize  (18)
//     RequestID   = 0
//     Action      = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed = 0
// }
var _ReplayChartDataPerformActionDefault = []byte{18, 0, 121, 39, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size        = ReplayChartDataPerformActionFixedSize  (16)
//     Type        = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     RequestID   = 0
//     Action      = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed = 0
// }
var _ReplayChartDataPerformActionFixedDefault = []byte{16, 0, 121, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type ReplayChartDataPerformAction struct {
	message.VLS
}

type ReplayChartDataPerformActionBuilder struct {
	message.VLSBuilder
}

type ReplayChartDataPerformActionFixed struct {
	message.Fixed
}

type ReplayChartDataPerformActionFixedBuilder struct {
	message.Fixed
}

type ReplayChartDataPerformActionPointer struct {
	message.VLSPointer
}

type ReplayChartDataPerformActionPointerBuilder struct {
	message.VLSPointerBuilder
}

type ReplayChartDataPerformActionFixedPointer struct {
	message.FixedPointer
}

type ReplayChartDataPerformActionFixedPointerBuilder struct {
	message.FixedPointer
}

func NewReplayChartDataPerformActionFrom(b []byte) ReplayChartDataPerformAction {
	return ReplayChartDataPerformAction{VLS: message.NewVLSFrom(b)}
}

func WrapReplayChartDataPerformAction(b []byte) ReplayChartDataPerformAction {
	return ReplayChartDataPerformAction{VLS: message.WrapVLS(b)}
}

func NewReplayChartDataPerformAction() ReplayChartDataPerformActionBuilder {
	a := ReplayChartDataPerformActionBuilder{message.NewVLSBuilder(18)}
	a.Unsafe().SetBytes(0, _ReplayChartDataPerformActionDefault)
	return a
}

func NewReplayChartDataPerformActionFixedFrom(b []byte) ReplayChartDataPerformActionFixed {
	return ReplayChartDataPerformActionFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapReplayChartDataPerformActionFixed(b []byte) ReplayChartDataPerformActionFixed {
	return ReplayChartDataPerformActionFixed{Fixed: message.WrapFixed(b)}
}

func NewReplayChartDataPerformActionFixed() ReplayChartDataPerformActionFixedBuilder {
	a := ReplayChartDataPerformActionFixedBuilder{message.NewFixed(16)}
	a.Unsafe().SetBytes(0, _ReplayChartDataPerformActionFixedDefault)
	return a
}

func AllocReplayChartDataPerformAction() ReplayChartDataPerformActionPointerBuilder {
	a := ReplayChartDataPerformActionPointerBuilder{message.AllocVLSBuilder(18)}
	a.Ptr.SetBytes(0, _ReplayChartDataPerformActionDefault)
	return a
}

func AllocReplayChartDataPerformActionFrom(b []byte) ReplayChartDataPerformActionPointer {
	return ReplayChartDataPerformActionPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocReplayChartDataPerformActionFixed() ReplayChartDataPerformActionFixedPointerBuilder {
	a := ReplayChartDataPerformActionFixedPointerBuilder{message.AllocFixed(16)}
	a.Ptr.SetBytes(0, _ReplayChartDataPerformActionFixedDefault)
	return a
}

func AllocReplayChartDataPerformActionFixedFrom(b []byte) ReplayChartDataPerformActionFixedPointer {
	return ReplayChartDataPerformActionFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size        = ReplayChartDataPerformActionSize  (18)
//     Type        = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     BaseSize    = ReplayChartDataPerformActionSize  (18)
//     RequestID   = 0
//     Action      = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed = 0
// }
func (m ReplayChartDataPerformActionBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ReplayChartDataPerformActionDefault)
}

// Clear
// {
//     Size        = ReplayChartDataPerformActionFixedSize  (16)
//     Type        = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     RequestID   = 0
//     Action      = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed = 0
// }
func (m ReplayChartDataPerformActionFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ReplayChartDataPerformActionFixedDefault)
}

// Clear
// {
//     Size        = ReplayChartDataPerformActionSize  (18)
//     Type        = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     BaseSize    = ReplayChartDataPerformActionSize  (18)
//     RequestID   = 0
//     Action      = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed = 0
// }
func (m ReplayChartDataPerformActionPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ReplayChartDataPerformActionDefault)
}

// Clear
// {
//     Size        = ReplayChartDataPerformActionFixedSize  (16)
//     Type        = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     RequestID   = 0
//     Action      = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed = 0
// }
func (m ReplayChartDataPerformActionFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ReplayChartDataPerformActionFixedDefault)
}

// ToBuilder
func (m ReplayChartDataPerformAction) ToBuilder() ReplayChartDataPerformActionBuilder {
	return ReplayChartDataPerformActionBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m ReplayChartDataPerformActionFixed) ToBuilder() ReplayChartDataPerformActionFixedBuilder {
	return ReplayChartDataPerformActionFixedBuilder{m.Fixed}
}

// ToBuilder
func (m ReplayChartDataPerformActionPointer) ToBuilder() ReplayChartDataPerformActionPointerBuilder {
	return ReplayChartDataPerformActionPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m ReplayChartDataPerformActionFixedPointer) ToBuilder() ReplayChartDataPerformActionFixedPointerBuilder {
	return ReplayChartDataPerformActionFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m ReplayChartDataPerformActionBuilder) Finish() ReplayChartDataPerformAction {
	return ReplayChartDataPerformAction{m.VLSBuilder.Finish()}
}

// Finish
func (m ReplayChartDataPerformActionFixedBuilder) Finish() ReplayChartDataPerformActionFixed {
	return ReplayChartDataPerformActionFixed{m.Fixed}
}

// Finish
func (m *ReplayChartDataPerformActionPointerBuilder) Finish() ReplayChartDataPerformActionPointer {
	return ReplayChartDataPerformActionPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *ReplayChartDataPerformActionFixedPointerBuilder) Finish() ReplayChartDataPerformActionFixedPointer {
	return ReplayChartDataPerformActionFixedPointer{m.FixedPointer}
}

// RequestID
func (m ReplayChartDataPerformAction) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m ReplayChartDataPerformActionBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m ReplayChartDataPerformActionPointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m ReplayChartDataPerformActionPointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// Action
func (m ReplayChartDataPerformAction) Action() ReplayChartDataActionEnum {
	return ReplayChartDataActionEnum(message.Int32VLS(m.Unsafe(), 14, 10))
}

// Action
func (m ReplayChartDataPerformActionBuilder) Action() ReplayChartDataActionEnum {
	return ReplayChartDataActionEnum(message.Int32VLS(m.Unsafe(), 14, 10))
}

// Action
func (m ReplayChartDataPerformActionPointer) Action() ReplayChartDataActionEnum {
	return ReplayChartDataActionEnum(message.Int32VLS(m.Ptr, 14, 10))
}

// Action
func (m ReplayChartDataPerformActionPointerBuilder) Action() ReplayChartDataActionEnum {
	return ReplayChartDataActionEnum(message.Int32VLS(m.Ptr, 14, 10))
}

// ReplaySpeed
func (m ReplayChartDataPerformAction) ReplaySpeed() float32 {
	return message.Float32VLS(m.Unsafe(), 18, 14)
}

// ReplaySpeed
func (m ReplayChartDataPerformActionBuilder) ReplaySpeed() float32 {
	return message.Float32VLS(m.Unsafe(), 18, 14)
}

// ReplaySpeed
func (m ReplayChartDataPerformActionPointer) ReplaySpeed() float32 {
	return message.Float32VLS(m.Ptr, 18, 14)
}

// ReplaySpeed
func (m ReplayChartDataPerformActionPointerBuilder) ReplaySpeed() float32 {
	return message.Float32VLS(m.Ptr, 18, 14)
}

// RequestID
func (m ReplayChartDataPerformActionFixed) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m ReplayChartDataPerformActionFixedBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m ReplayChartDataPerformActionFixedPointer) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m ReplayChartDataPerformActionFixedPointerBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Action
func (m ReplayChartDataPerformActionFixed) Action() ReplayChartDataActionEnum {
	return ReplayChartDataActionEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Action
func (m ReplayChartDataPerformActionFixedBuilder) Action() ReplayChartDataActionEnum {
	return ReplayChartDataActionEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Action
func (m ReplayChartDataPerformActionFixedPointer) Action() ReplayChartDataActionEnum {
	return ReplayChartDataActionEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// Action
func (m ReplayChartDataPerformActionFixedPointerBuilder) Action() ReplayChartDataActionEnum {
	return ReplayChartDataActionEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// ReplaySpeed
func (m ReplayChartDataPerformActionFixed) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Unsafe(), 16, 12)
}

// ReplaySpeed
func (m ReplayChartDataPerformActionFixedBuilder) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Unsafe(), 16, 12)
}

// ReplaySpeed
func (m ReplayChartDataPerformActionFixedPointer) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Ptr, 16, 12)
}

// ReplaySpeed
func (m ReplayChartDataPerformActionFixedPointerBuilder) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Ptr, 16, 12)
}

// SetRequestID
func (m ReplayChartDataPerformActionBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m ReplayChartDataPerformActionPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetAction
func (m ReplayChartDataPerformActionBuilder) SetAction(value ReplayChartDataActionEnum) {
	message.SetInt32VLS(m.Unsafe(), 14, 10, int32(value))
}

// SetAction
func (m ReplayChartDataPerformActionPointerBuilder) SetAction(value ReplayChartDataActionEnum) {
	message.SetInt32VLS(m.Ptr, 14, 10, int32(value))
}

// SetReplaySpeed
func (m ReplayChartDataPerformActionBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 18, 14, value)
}

// SetReplaySpeed
func (m ReplayChartDataPerformActionPointerBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32VLS(m.Ptr, 18, 14, value)
}

// SetRequestID
func (m ReplayChartDataPerformActionFixedBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m ReplayChartDataPerformActionFixedPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetAction
func (m ReplayChartDataPerformActionFixedBuilder) SetAction(value ReplayChartDataActionEnum) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, int32(value))
}

// SetAction
func (m ReplayChartDataPerformActionFixedPointerBuilder) SetAction(value ReplayChartDataActionEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
}

// SetReplaySpeed
func (m ReplayChartDataPerformActionFixedBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 16, 12, value)
}

// SetReplaySpeed
func (m ReplayChartDataPerformActionFixedPointerBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32Fixed(m.Ptr, 16, 12, value)
}

// Copy
func (m ReplayChartDataPerformAction) Copy(to ReplayChartDataPerformActionBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// Copy
func (m ReplayChartDataPerformActionBuilder) Copy(to ReplayChartDataPerformActionBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyPointer
func (m ReplayChartDataPerformAction) CopyPointer(to ReplayChartDataPerformActionPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyPointer
func (m ReplayChartDataPerformActionBuilder) CopyPointer(to ReplayChartDataPerformActionPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyToPointer
func (m ReplayChartDataPerformAction) CopyToPointer(to ReplayChartDataPerformActionFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyToPointer
func (m ReplayChartDataPerformActionBuilder) CopyToPointer(to ReplayChartDataPerformActionFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyTo
func (m ReplayChartDataPerformAction) CopyTo(to ReplayChartDataPerformActionFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyTo
func (m ReplayChartDataPerformActionBuilder) CopyTo(to ReplayChartDataPerformActionFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
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
func (m ReplayChartDataPerformActionFixed) CopyToPointer(to ReplayChartDataPerformActionPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyToPointer
func (m ReplayChartDataPerformActionFixedBuilder) CopyToPointer(to ReplayChartDataPerformActionPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyTo
func (m ReplayChartDataPerformActionFixed) CopyTo(to ReplayChartDataPerformActionBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyTo
func (m ReplayChartDataPerformActionFixedBuilder) CopyTo(to ReplayChartDataPerformActionBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// Copy
func (m ReplayChartDataPerformActionPointer) Copy(to ReplayChartDataPerformActionBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// Copy
func (m ReplayChartDataPerformActionPointerBuilder) Copy(to ReplayChartDataPerformActionBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyPointer
func (m ReplayChartDataPerformActionPointer) CopyPointer(to ReplayChartDataPerformActionPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyPointer
func (m ReplayChartDataPerformActionPointerBuilder) CopyPointer(to ReplayChartDataPerformActionPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyTo
func (m ReplayChartDataPerformActionPointer) CopyTo(to ReplayChartDataPerformActionFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyTo
func (m ReplayChartDataPerformActionPointerBuilder) CopyTo(to ReplayChartDataPerformActionFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyToPointer
func (m ReplayChartDataPerformActionPointer) CopyToPointer(to ReplayChartDataPerformActionFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyToPointer
func (m ReplayChartDataPerformActionPointerBuilder) CopyToPointer(to ReplayChartDataPerformActionFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
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
func (m ReplayChartDataPerformActionFixedPointer) CopyTo(to ReplayChartDataPerformActionBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyTo
func (m ReplayChartDataPerformActionFixedPointerBuilder) CopyTo(to ReplayChartDataPerformActionBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyToPointer
func (m ReplayChartDataPerformActionFixedPointer) CopyToPointer(to ReplayChartDataPerformActionPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}

// CopyToPointer
func (m ReplayChartDataPerformActionFixedPointerBuilder) CopyToPointer(to ReplayChartDataPerformActionPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetAction(m.Action())
	to.SetReplaySpeed(m.ReplaySpeed())
}
