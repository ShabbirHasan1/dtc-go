// generated by github.com/moontrade/dtc-go/codegen/go at 2022-03-14 08:04:12.929995 +0800 WITA m=+0.027582959

package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const ReplayChartDataPerformActionSize = 18

const ReplayChartDataPerformActionFixedSize = 16

//     Size         uint16                     = ReplayChartDataPerformActionSize  (18)
//     Type         uint16                     = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     BaseSize     uint16                     = ReplayChartDataPerformActionSize  (18)
//     RequestID    uint32                     = 0
//     Action       ReplayChartDataActionEnum  = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed  float32                    = 0
var _ReplayChartDataPerformActionDefault = []byte{18, 0, 121, 39, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size         uint16                     = ReplayChartDataPerformActionFixedSize  (16)
//     Type         uint16                     = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     RequestID    uint32                     = 0
//     Action       ReplayChartDataActionEnum  = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed  float32                    = 0
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
	return ReplayChartDataPerformAction{VLS: message.NewVLS(b)}
}

func WrapReplayChartDataPerformAction(b []byte) ReplayChartDataPerformAction {
	return ReplayChartDataPerformAction{VLS: message.WrapVLS(b)}
}

func NewReplayChartDataPerformAction() ReplayChartDataPerformActionBuilder {
	return ReplayChartDataPerformActionBuilder{message.NewVLSBuilder(_ReplayChartDataPerformActionDefault)}
}

func NewReplayChartDataPerformActionFixedFrom(b []byte) ReplayChartDataPerformActionFixed {
	return ReplayChartDataPerformActionFixed{Fixed: message.NewFixed(b)}
}

func WrapReplayChartDataPerformActionFixed(b []byte) ReplayChartDataPerformActionFixed {
	return ReplayChartDataPerformActionFixed{Fixed: message.WrapFixed(b)}
}

func NewReplayChartDataPerformActionFixed() ReplayChartDataPerformActionFixedBuilder {
	return ReplayChartDataPerformActionFixedBuilder{message.NewFixed(_ReplayChartDataPerformActionFixedDefault)}
}

func AllocReplayChartDataPerformAction() ReplayChartDataPerformActionPointerBuilder {
	return ReplayChartDataPerformActionPointerBuilder{message.AllocVLSBuilder(_ReplayChartDataPerformActionDefault)}
}

func AllocReplayChartDataPerformActionFrom(b []byte) ReplayChartDataPerformActionPointer {
	return ReplayChartDataPerformActionPointer{VLSPointer: message.AllocVLS(b)}
}

func AllocReplayChartDataPerformActionFixed() ReplayChartDataPerformActionFixedPointerBuilder {
	return ReplayChartDataPerformActionFixedPointerBuilder{message.AllocFixed(_ReplayChartDataPerformActionFixedDefault)}
}

func AllocReplayChartDataPerformActionFixedFrom(b []byte) ReplayChartDataPerformActionFixedPointer {
	return ReplayChartDataPerformActionFixedPointer{FixedPointer: message.AllocFixed(b)}
}

// Clear
//     Size         uint16                     = ReplayChartDataPerformActionSize  (18)
//     Type         uint16                     = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     BaseSize     uint16                     = ReplayChartDataPerformActionSize  (18)
//     RequestID    uint32                     = 0
//     Action       ReplayChartDataActionEnum  = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed  float32                    = 0
func (m ReplayChartDataPerformActionBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ReplayChartDataPerformActionDefault)
}

// Clear
//     Size         uint16                     = ReplayChartDataPerformActionFixedSize  (16)
//     Type         uint16                     = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     RequestID    uint32                     = 0
//     Action       ReplayChartDataActionEnum  = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed  float32                    = 0
func (m ReplayChartDataPerformActionFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ReplayChartDataPerformActionFixedDefault)
}

// Clear
//     Size         uint16                     = ReplayChartDataPerformActionSize  (18)
//     Type         uint16                     = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     BaseSize     uint16                     = ReplayChartDataPerformActionSize  (18)
//     RequestID    uint32                     = 0
//     Action       ReplayChartDataActionEnum  = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed  float32                    = 0
func (m ReplayChartDataPerformActionPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ReplayChartDataPerformActionDefault)
}

// Clear
//     Size         uint16                     = ReplayChartDataPerformActionFixedSize  (16)
//     Type         uint16                     = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
//     RequestID    uint32                     = 0
//     Action       ReplayChartDataActionEnum  = REPLAY_CHART_DATA_ACTION_NONE  (0)
//     ReplaySpeed  float32                    = 0
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
