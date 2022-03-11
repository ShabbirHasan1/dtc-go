package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const ReplayChartDataStatusSize = 22

const ReplayChartDataStatusFixedSize = 272

// {
//     Size                 = ReplayChartDataStatusSize  (22)
//     Type                 = REPLAY_CHART_DATA_STATUS  (10106)
//     BaseSize             = ReplayChartDataStatusSize  (22)
//     RequestID            = 0
//     ErrorMessage         = ""
//     Status               = REPLAY_CHART_DATA_STATUS_UNSET  (0)
//     SubAccountIdentifier = 0
// }
var _ReplayChartDataStatusDefault = []byte{22, 0, 122, 39, 22, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size                 = ReplayChartDataStatusFixedSize  (272)
//     Type                 = REPLAY_CHART_DATA_STATUS  (10106)
//     RequestID            = 0
//     ErrorMessage         = ""
//     Status               = REPLAY_CHART_DATA_STATUS_UNSET  (0)
//     SubAccountIdentifier = 0
// }
var _ReplayChartDataStatusFixedDefault = []byte{16, 1, 122, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type ReplayChartDataStatus struct {
	message.VLS
}

type ReplayChartDataStatusBuilder struct {
	message.VLSBuilder
}

type ReplayChartDataStatusFixed struct {
	message.Fixed
}

type ReplayChartDataStatusFixedBuilder struct {
	message.Fixed
}

type ReplayChartDataStatusPointer struct {
	message.VLSPointer
}

type ReplayChartDataStatusPointerBuilder struct {
	message.VLSPointerBuilder
}

type ReplayChartDataStatusFixedPointer struct {
	message.FixedPointer
}

type ReplayChartDataStatusFixedPointerBuilder struct {
	message.FixedPointer
}

func NewReplayChartDataStatusFrom(b []byte) ReplayChartDataStatus {
	return ReplayChartDataStatus{VLS: message.NewVLSFrom(b)}
}

func WrapReplayChartDataStatus(b []byte) ReplayChartDataStatus {
	return ReplayChartDataStatus{VLS: message.WrapVLS(b)}
}

func NewReplayChartDataStatus() ReplayChartDataStatusBuilder {
	a := ReplayChartDataStatusBuilder{message.NewVLSBuilder(22)}
	a.Unsafe().SetBytes(0, _ReplayChartDataStatusDefault)
	return a
}

func NewReplayChartDataStatusFixedFrom(b []byte) ReplayChartDataStatusFixed {
	return ReplayChartDataStatusFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapReplayChartDataStatusFixed(b []byte) ReplayChartDataStatusFixed {
	return ReplayChartDataStatusFixed{Fixed: message.WrapFixed(b)}
}

func NewReplayChartDataStatusFixed() ReplayChartDataStatusFixedBuilder {
	a := ReplayChartDataStatusFixedBuilder{message.NewFixed(272)}
	a.Unsafe().SetBytes(0, _ReplayChartDataStatusFixedDefault)
	return a
}

func AllocReplayChartDataStatus() ReplayChartDataStatusPointerBuilder {
	a := ReplayChartDataStatusPointerBuilder{message.AllocVLSBuilder(22)}
	a.Ptr.SetBytes(0, _ReplayChartDataStatusDefault)
	return a
}

func AllocReplayChartDataStatusFrom(b []byte) ReplayChartDataStatusPointer {
	return ReplayChartDataStatusPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocReplayChartDataStatusFixed() ReplayChartDataStatusFixedPointerBuilder {
	a := ReplayChartDataStatusFixedPointerBuilder{message.AllocFixed(272)}
	a.Ptr.SetBytes(0, _ReplayChartDataStatusFixedDefault)
	return a
}

func AllocReplayChartDataStatusFixedFrom(b []byte) ReplayChartDataStatusFixedPointer {
	return ReplayChartDataStatusFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                 = ReplayChartDataStatusSize  (22)
//     Type                 = REPLAY_CHART_DATA_STATUS  (10106)
//     BaseSize             = ReplayChartDataStatusSize  (22)
//     RequestID            = 0
//     ErrorMessage         = ""
//     Status               = REPLAY_CHART_DATA_STATUS_UNSET  (0)
//     SubAccountIdentifier = 0
// }
func (m ReplayChartDataStatusBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ReplayChartDataStatusDefault)
}

// Clear
// {
//     Size                 = ReplayChartDataStatusFixedSize  (272)
//     Type                 = REPLAY_CHART_DATA_STATUS  (10106)
//     RequestID            = 0
//     ErrorMessage         = ""
//     Status               = REPLAY_CHART_DATA_STATUS_UNSET  (0)
//     SubAccountIdentifier = 0
// }
func (m ReplayChartDataStatusFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ReplayChartDataStatusFixedDefault)
}

// Clear
// {
//     Size                 = ReplayChartDataStatusSize  (22)
//     Type                 = REPLAY_CHART_DATA_STATUS  (10106)
//     BaseSize             = ReplayChartDataStatusSize  (22)
//     RequestID            = 0
//     ErrorMessage         = ""
//     Status               = REPLAY_CHART_DATA_STATUS_UNSET  (0)
//     SubAccountIdentifier = 0
// }
func (m ReplayChartDataStatusPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ReplayChartDataStatusDefault)
}

// Clear
// {
//     Size                 = ReplayChartDataStatusFixedSize  (272)
//     Type                 = REPLAY_CHART_DATA_STATUS  (10106)
//     RequestID            = 0
//     ErrorMessage         = ""
//     Status               = REPLAY_CHART_DATA_STATUS_UNSET  (0)
//     SubAccountIdentifier = 0
// }
func (m ReplayChartDataStatusFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ReplayChartDataStatusFixedDefault)
}

// ToBuilder
func (m ReplayChartDataStatus) ToBuilder() ReplayChartDataStatusBuilder {
	return ReplayChartDataStatusBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m ReplayChartDataStatusFixed) ToBuilder() ReplayChartDataStatusFixedBuilder {
	return ReplayChartDataStatusFixedBuilder{m.Fixed}
}

// ToBuilder
func (m ReplayChartDataStatusPointer) ToBuilder() ReplayChartDataStatusPointerBuilder {
	return ReplayChartDataStatusPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m ReplayChartDataStatusFixedPointer) ToBuilder() ReplayChartDataStatusFixedPointerBuilder {
	return ReplayChartDataStatusFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m ReplayChartDataStatusBuilder) Finish() ReplayChartDataStatus {
	return ReplayChartDataStatus{m.VLSBuilder.Finish()}
}

// Finish
func (m ReplayChartDataStatusFixedBuilder) Finish() ReplayChartDataStatusFixed {
	return ReplayChartDataStatusFixed{m.Fixed}
}

// Finish
func (m *ReplayChartDataStatusPointerBuilder) Finish() ReplayChartDataStatusPointer {
	return ReplayChartDataStatusPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *ReplayChartDataStatusFixedPointerBuilder) Finish() ReplayChartDataStatusFixedPointer {
	return ReplayChartDataStatusFixedPointer{m.FixedPointer}
}

// RequestID
func (m ReplayChartDataStatus) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m ReplayChartDataStatusBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m ReplayChartDataStatusPointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m ReplayChartDataStatusPointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// ErrorMessage
func (m ReplayChartDataStatus) ErrorMessage() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// ErrorMessage
func (m ReplayChartDataStatusBuilder) ErrorMessage() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// ErrorMessage
func (m ReplayChartDataStatusPointer) ErrorMessage() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// ErrorMessage
func (m ReplayChartDataStatusPointerBuilder) ErrorMessage() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Status
func (m ReplayChartDataStatus) Status() ReplayChartDataStatusEnum {
	return ReplayChartDataStatusEnum(message.Int32VLS(m.Unsafe(), 18, 14))
}

// Status
func (m ReplayChartDataStatusBuilder) Status() ReplayChartDataStatusEnum {
	return ReplayChartDataStatusEnum(message.Int32VLS(m.Unsafe(), 18, 14))
}

// Status
func (m ReplayChartDataStatusPointer) Status() ReplayChartDataStatusEnum {
	return ReplayChartDataStatusEnum(message.Int32VLS(m.Ptr, 18, 14))
}

// Status
func (m ReplayChartDataStatusPointerBuilder) Status() ReplayChartDataStatusEnum {
	return ReplayChartDataStatusEnum(message.Int32VLS(m.Ptr, 18, 14))
}

// SubAccountIdentifier
func (m ReplayChartDataStatus) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 22, 18)
}

// SubAccountIdentifier
func (m ReplayChartDataStatusBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 22, 18)
}

// SubAccountIdentifier
func (m ReplayChartDataStatusPointer) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 22, 18)
}

// SubAccountIdentifier
func (m ReplayChartDataStatusPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 22, 18)
}

// RequestID
func (m ReplayChartDataStatusFixed) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m ReplayChartDataStatusFixedBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m ReplayChartDataStatusFixedPointer) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m ReplayChartDataStatusFixedPointerBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// ErrorMessage
func (m ReplayChartDataStatusFixed) ErrorMessage() string {
	return message.StringFixed(m.Unsafe(), 264, 8)
}

// ErrorMessage
func (m ReplayChartDataStatusFixedBuilder) ErrorMessage() string {
	return message.StringFixed(m.Unsafe(), 264, 8)
}

// ErrorMessage
func (m ReplayChartDataStatusFixedPointer) ErrorMessage() string {
	return message.StringFixed(m.Ptr, 264, 8)
}

// ErrorMessage
func (m ReplayChartDataStatusFixedPointerBuilder) ErrorMessage() string {
	return message.StringFixed(m.Ptr, 264, 8)
}

// Status
func (m ReplayChartDataStatusFixed) Status() ReplayChartDataStatusEnum {
	return ReplayChartDataStatusEnum(message.Int32Fixed(m.Unsafe(), 268, 264))
}

// Status
func (m ReplayChartDataStatusFixedBuilder) Status() ReplayChartDataStatusEnum {
	return ReplayChartDataStatusEnum(message.Int32Fixed(m.Unsafe(), 268, 264))
}

// Status
func (m ReplayChartDataStatusFixedPointer) Status() ReplayChartDataStatusEnum {
	return ReplayChartDataStatusEnum(message.Int32Fixed(m.Ptr, 268, 264))
}

// Status
func (m ReplayChartDataStatusFixedPointerBuilder) Status() ReplayChartDataStatusEnum {
	return ReplayChartDataStatusEnum(message.Int32Fixed(m.Ptr, 268, 264))
}

// SubAccountIdentifier
func (m ReplayChartDataStatusFixed) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 272, 268)
}

// SubAccountIdentifier
func (m ReplayChartDataStatusFixedBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 272, 268)
}

// SubAccountIdentifier
func (m ReplayChartDataStatusFixedPointer) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Ptr, 272, 268)
}

// SubAccountIdentifier
func (m ReplayChartDataStatusFixedPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Ptr, 272, 268)
}

// SetRequestID
func (m ReplayChartDataStatusBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m ReplayChartDataStatusPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetErrorMessage
func (m *ReplayChartDataStatusBuilder) SetErrorMessage(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetErrorMessage
func (m *ReplayChartDataStatusPointerBuilder) SetErrorMessage(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetStatus
func (m ReplayChartDataStatusBuilder) SetStatus(value ReplayChartDataStatusEnum) {
	message.SetInt32VLS(m.Unsafe(), 18, 14, int32(value))
}

// SetStatus
func (m ReplayChartDataStatusPointerBuilder) SetStatus(value ReplayChartDataStatusEnum) {
	message.SetInt32VLS(m.Ptr, 18, 14, int32(value))
}

// SetSubAccountIdentifier
func (m ReplayChartDataStatusBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 22, 18, value)
}

// SetSubAccountIdentifier
func (m ReplayChartDataStatusPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Ptr, 22, 18, value)
}

// SetRequestID
func (m ReplayChartDataStatusFixedBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m ReplayChartDataStatusFixedPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetErrorMessage
func (m ReplayChartDataStatusFixedBuilder) SetErrorMessage(value string) {
	message.SetStringFixed(m.Unsafe(), 264, 8, value)
}

// SetErrorMessage
func (m ReplayChartDataStatusFixedPointerBuilder) SetErrorMessage(value string) {
	message.SetStringFixed(m.Ptr, 264, 8, value)
}

// SetStatus
func (m ReplayChartDataStatusFixedBuilder) SetStatus(value ReplayChartDataStatusEnum) {
	message.SetInt32Fixed(m.Unsafe(), 268, 264, int32(value))
}

// SetStatus
func (m ReplayChartDataStatusFixedPointerBuilder) SetStatus(value ReplayChartDataStatusEnum) {
	message.SetInt32Fixed(m.Ptr, 268, 264, int32(value))
}

// SetSubAccountIdentifier
func (m ReplayChartDataStatusFixedBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 272, 268, value)
}

// SetSubAccountIdentifier
func (m ReplayChartDataStatusFixedPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Ptr, 272, 268, value)
}

// Copy
func (m ReplayChartDataStatus) Copy(to ReplayChartDataStatusBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// Copy
func (m ReplayChartDataStatusBuilder) Copy(to ReplayChartDataStatusBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyPointer
func (m ReplayChartDataStatus) CopyPointer(to ReplayChartDataStatusPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyPointer
func (m ReplayChartDataStatusBuilder) CopyPointer(to ReplayChartDataStatusPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyToPointer
func (m ReplayChartDataStatus) CopyToPointer(to ReplayChartDataStatusFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyToPointer
func (m ReplayChartDataStatusBuilder) CopyToPointer(to ReplayChartDataStatusFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyTo
func (m ReplayChartDataStatus) CopyTo(to ReplayChartDataStatusFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyTo
func (m ReplayChartDataStatusBuilder) CopyTo(to ReplayChartDataStatusFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// Copy
func (m ReplayChartDataStatusFixed) Copy(to ReplayChartDataStatusFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// Copy
func (m ReplayChartDataStatusFixedBuilder) Copy(to ReplayChartDataStatusFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyPointer
func (m ReplayChartDataStatusFixed) CopyPointer(to ReplayChartDataStatusFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyPointer
func (m ReplayChartDataStatusFixedBuilder) CopyPointer(to ReplayChartDataStatusFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyToPointer
func (m ReplayChartDataStatusFixed) CopyToPointer(to ReplayChartDataStatusPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyToPointer
func (m ReplayChartDataStatusFixedBuilder) CopyToPointer(to ReplayChartDataStatusPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyTo
func (m ReplayChartDataStatusFixed) CopyTo(to ReplayChartDataStatusBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyTo
func (m ReplayChartDataStatusFixedBuilder) CopyTo(to ReplayChartDataStatusBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// Copy
func (m ReplayChartDataStatusPointer) Copy(to ReplayChartDataStatusBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// Copy
func (m ReplayChartDataStatusPointerBuilder) Copy(to ReplayChartDataStatusBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyPointer
func (m ReplayChartDataStatusPointer) CopyPointer(to ReplayChartDataStatusPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyPointer
func (m ReplayChartDataStatusPointerBuilder) CopyPointer(to ReplayChartDataStatusPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyTo
func (m ReplayChartDataStatusPointer) CopyTo(to ReplayChartDataStatusFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyTo
func (m ReplayChartDataStatusPointerBuilder) CopyTo(to ReplayChartDataStatusFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyToPointer
func (m ReplayChartDataStatusPointer) CopyToPointer(to ReplayChartDataStatusFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyToPointer
func (m ReplayChartDataStatusPointerBuilder) CopyToPointer(to ReplayChartDataStatusFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// Copy
func (m ReplayChartDataStatusFixedPointer) Copy(to ReplayChartDataStatusFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// Copy
func (m ReplayChartDataStatusFixedPointerBuilder) Copy(to ReplayChartDataStatusFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyPointer
func (m ReplayChartDataStatusFixedPointer) CopyPointer(to ReplayChartDataStatusFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyPointer
func (m ReplayChartDataStatusFixedPointerBuilder) CopyPointer(to ReplayChartDataStatusFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyTo
func (m ReplayChartDataStatusFixedPointer) CopyTo(to ReplayChartDataStatusBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyTo
func (m ReplayChartDataStatusFixedPointerBuilder) CopyTo(to ReplayChartDataStatusBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyToPointer
func (m ReplayChartDataStatusFixedPointer) CopyToPointer(to ReplayChartDataStatusPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}

// CopyToPointer
func (m ReplayChartDataStatusFixedPointerBuilder) CopyToPointer(to ReplayChartDataStatusPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetErrorMessage(m.ErrorMessage())
	to.SetStatus(m.Status())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
}
