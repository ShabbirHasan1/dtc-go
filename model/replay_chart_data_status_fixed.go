package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                 = ReplayChartDataStatusFixedSize  (272)
//     Type                 = REPLAY_CHART_DATA_STATUS  (10106)
//     RequestID            = 0
//     ErrorMessage         = ""
//     Status               = REPLAY_CHART_DATA_STATUS_UNSET  (0)
//     SubAccountIdentifier = 0
// }
var _ReplayChartDataStatusFixedDefault = []byte{16, 1, 122, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const ReplayChartDataStatusFixedSize = 272

type ReplayChartDataStatusFixed struct {
	message.Fixed
}

type ReplayChartDataStatusFixedBuilder struct {
	message.Fixed
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

// ToBuilder
func (m ReplayChartDataStatusFixed) ToBuilder() ReplayChartDataStatusFixedBuilder {
	return ReplayChartDataStatusFixedBuilder{m.Fixed}
}

// Finish
func (m ReplayChartDataStatusFixedBuilder) Finish() ReplayChartDataStatusFixed {
	return ReplayChartDataStatusFixed{m.Fixed}
}

// RequestID
func (m ReplayChartDataStatusFixed) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m ReplayChartDataStatusFixedBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// ErrorMessage
func (m ReplayChartDataStatusFixed) ErrorMessage() string {
	return message.StringFixed(m.Unsafe(), 264, 8)
}

// ErrorMessage
func (m ReplayChartDataStatusFixedBuilder) ErrorMessage() string {
	return message.StringFixed(m.Unsafe(), 264, 8)
}

// Status
func (m ReplayChartDataStatusFixed) Status() ReplayChartDataStatusEnum {
	return ReplayChartDataStatusEnum(message.Int32Fixed(m.Unsafe(), 268, 264))
}

// Status
func (m ReplayChartDataStatusFixedBuilder) Status() ReplayChartDataStatusEnum {
	return ReplayChartDataStatusEnum(message.Int32Fixed(m.Unsafe(), 268, 264))
}

// SubAccountIdentifier
func (m ReplayChartDataStatusFixed) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 272, 268)
}

// SubAccountIdentifier
func (m ReplayChartDataStatusFixedBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 272, 268)
}

// SetRequestID
func (m ReplayChartDataStatusFixedBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetErrorMessage
func (m ReplayChartDataStatusFixedBuilder) SetErrorMessage(value string) {
	message.SetStringFixed(m.Unsafe(), 264, 8, value)
}

// SetStatus
func (m ReplayChartDataStatusFixedBuilder) SetStatus(value ReplayChartDataStatusEnum) {
	message.SetInt32Fixed(m.Unsafe(), 268, 264, int32(value))
}

// SetSubAccountIdentifier
func (m ReplayChartDataStatusFixedBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 272, 268, value)
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
