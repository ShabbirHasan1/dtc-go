package model

import (
	"github.com/moontrade/dtc-go/message"
)

type ReplayChartDataStatusFixedPointer struct {
	message.FixedPointer
}

type ReplayChartDataStatusFixedPointerBuilder struct {
	message.FixedPointer
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
func (m ReplayChartDataStatusFixedPointer) ToBuilder() ReplayChartDataStatusFixedPointerBuilder {
	return ReplayChartDataStatusFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *ReplayChartDataStatusFixedPointerBuilder) Finish() ReplayChartDataStatusFixedPointer {
	return ReplayChartDataStatusFixedPointer{m.FixedPointer}
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
func (m ReplayChartDataStatusFixedPointer) ErrorMessage() string {
	return message.StringFixed(m.Ptr, 264, 8)
}

// ErrorMessage
func (m ReplayChartDataStatusFixedPointerBuilder) ErrorMessage() string {
	return message.StringFixed(m.Ptr, 264, 8)
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
func (m ReplayChartDataStatusFixedPointer) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Ptr, 272, 268)
}

// SubAccountIdentifier
func (m ReplayChartDataStatusFixedPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Ptr, 272, 268)
}

// SetRequestID
func (m ReplayChartDataStatusFixedPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetErrorMessage
func (m ReplayChartDataStatusFixedPointerBuilder) SetErrorMessage(value string) {
	message.SetStringFixed(m.Ptr, 264, 8, value)
}

// SetStatus
func (m ReplayChartDataStatusFixedPointerBuilder) SetStatus(value ReplayChartDataStatusEnum) {
	message.SetInt32Fixed(m.Ptr, 268, 264, int32(value))
}

// SetSubAccountIdentifier
func (m ReplayChartDataStatusFixedPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Ptr, 272, 268, value)
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
