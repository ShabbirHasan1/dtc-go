package model

import (
	"github.com/moontrade/dtc-go/message"
)

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

const ReplayChartDataStatusSize = 22

type ReplayChartDataStatus struct {
	message.VLS
}

type ReplayChartDataStatusBuilder struct {
	message.VLSBuilder
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

// ToBuilder
func (m ReplayChartDataStatus) ToBuilder() ReplayChartDataStatusBuilder {
	return ReplayChartDataStatusBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m ReplayChartDataStatusBuilder) Finish() ReplayChartDataStatus {
	return ReplayChartDataStatus{m.VLSBuilder.Finish()}
}

// RequestID
func (m ReplayChartDataStatus) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m ReplayChartDataStatusBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// ErrorMessage
func (m ReplayChartDataStatus) ErrorMessage() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// ErrorMessage
func (m ReplayChartDataStatusBuilder) ErrorMessage() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Status
func (m ReplayChartDataStatus) Status() ReplayChartDataStatusEnum {
	return ReplayChartDataStatusEnum(message.Int32VLS(m.Unsafe(), 18, 14))
}

// Status
func (m ReplayChartDataStatusBuilder) Status() ReplayChartDataStatusEnum {
	return ReplayChartDataStatusEnum(message.Int32VLS(m.Unsafe(), 18, 14))
}

// SubAccountIdentifier
func (m ReplayChartDataStatus) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 22, 18)
}

// SubAccountIdentifier
func (m ReplayChartDataStatusBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 22, 18)
}

// SetRequestID
func (m ReplayChartDataStatusBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetErrorMessage
func (m *ReplayChartDataStatusBuilder) SetErrorMessage(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetStatus
func (m ReplayChartDataStatusBuilder) SetStatus(value ReplayChartDataStatusEnum) {
	message.SetInt32VLS(m.Unsafe(), 18, 14, int32(value))
}

// SetSubAccountIdentifier
func (m ReplayChartDataStatusBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 22, 18, value)
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
