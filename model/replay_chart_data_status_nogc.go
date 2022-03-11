package model

import (
	"github.com/moontrade/dtc-go/message"
)

type ReplayChartDataStatusPointer struct {
	message.VLSPointer
}

type ReplayChartDataStatusPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocReplayChartDataStatus() ReplayChartDataStatusPointerBuilder {
	a := ReplayChartDataStatusPointerBuilder{message.AllocVLSBuilder(22)}
	a.Ptr.SetBytes(0, _ReplayChartDataStatusDefault)
	return a
}

func AllocReplayChartDataStatusFrom(b []byte) ReplayChartDataStatusPointer {
	return ReplayChartDataStatusPointer{VLSPointer: message.AllocVLSFrom(b)}
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

// ToBuilder
func (m ReplayChartDataStatusPointer) ToBuilder() ReplayChartDataStatusPointerBuilder {
	return ReplayChartDataStatusPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *ReplayChartDataStatusPointerBuilder) Finish() ReplayChartDataStatusPointer {
	return ReplayChartDataStatusPointer{m.VLSPointerBuilder.Finish()}
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
func (m ReplayChartDataStatusPointer) ErrorMessage() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// ErrorMessage
func (m ReplayChartDataStatusPointerBuilder) ErrorMessage() string {
	return message.StringVLS(m.Ptr, 14, 10)
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
func (m ReplayChartDataStatusPointer) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 22, 18)
}

// SubAccountIdentifier
func (m ReplayChartDataStatusPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 22, 18)
}

// SetRequestID
func (m ReplayChartDataStatusPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetErrorMessage
func (m *ReplayChartDataStatusPointerBuilder) SetErrorMessage(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetStatus
func (m ReplayChartDataStatusPointerBuilder) SetStatus(value ReplayChartDataStatusEnum) {
	message.SetInt32VLS(m.Ptr, 18, 14, int32(value))
}

// SetSubAccountIdentifier
func (m ReplayChartDataStatusPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Ptr, 22, 18, value)
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
