package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HistoricalPriceDataRejectPointer When the Server rejects a historical price data request from the Client,
// a HistoricalPriceDataReject message will be sent.
//
// This message is never compressed.
type HistoricalPriceDataRejectPointer struct {
	message.VLSPointer
}

// HistoricalPriceDataRejectPointerBuilder When the Server rejects a historical price data request from the Client,
// a HistoricalPriceDataReject message will be sent.
//
// This message is never compressed.
type HistoricalPriceDataRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocHistoricalPriceDataReject() HistoricalPriceDataRejectPointerBuilder {
	a := HistoricalPriceDataRejectPointerBuilder{message.AllocVLSBuilder(20)}
	a.Ptr.SetBytes(0, _HistoricalPriceDataRejectDefault)
	return a
}

func AllocHistoricalPriceDataRejectFrom(b []byte) HistoricalPriceDataRejectPointer {
	return HistoricalPriceDataRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size               = HistoricalPriceDataRejectSize  (20)
//     Type               = HISTORICAL_PRICE_DATA_REJECT  (802)
//     BaseSize           = HistoricalPriceDataRejectSize  (20)
//     RequestID          = 0
//     RejectText         = ""
//     RejectReasonCode   = 0
//     RetryTimeInSeconds = 0
// }
func (m HistoricalPriceDataRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataRejectDefault)
}

// ToBuilder
func (m HistoricalPriceDataRejectPointer) ToBuilder() HistoricalPriceDataRejectPointerBuilder {
	return HistoricalPriceDataRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalPriceDataRejectPointerBuilder) Finish() HistoricalPriceDataRejectPointer {
	return HistoricalPriceDataRejectPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText Text reason for rejection.
func (m HistoricalPriceDataRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Text reason for rejection.
func (m HistoricalPriceDataRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectPointer) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Ptr, 18, 16))
}

// RejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectPointerBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Ptr, 18, 16))
}

// RetryTimeInSeconds This is an optional field from the Server. This field will normally be
// zero.
//
// If a retry is intended to be performed, the server may give an indication
// of how long to wait in seconds. This field indicates that.
//
// This field is not recommended to be used. If it is used, it is really
// an indication of a substandard Server.
func (m HistoricalPriceDataRejectPointer) RetryTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Ptr, 20, 18)
}

// RetryTimeInSeconds This is an optional field from the Server. This field will normally be
// zero.
//
// If a retry is intended to be performed, the server may give an indication
// of how long to wait in seconds. This field indicates that.
//
// This field is not recommended to be used. If it is used, it is really
// an indication of a substandard Server.
func (m HistoricalPriceDataRejectPointerBuilder) RetryTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Ptr, 20, 18)
}

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Text reason for rejection.
func (m *HistoricalPriceDataRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectPointerBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16VLS(m.Ptr, 18, 16, int16(value))
}

// SetRetryTimeInSeconds This is an optional field from the Server. This field will normally be
// zero.
//
// If a retry is intended to be performed, the server may give an indication
// of how long to wait in seconds. This field indicates that.
//
// This field is not recommended to be used. If it is used, it is really
// an indication of a substandard Server.
func (m HistoricalPriceDataRejectPointerBuilder) SetRetryTimeInSeconds(value uint16) {
	message.SetUint16VLS(m.Ptr, 20, 18, value)
}

// Copy
func (m HistoricalPriceDataRejectPointer) Copy(to HistoricalPriceDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// Copy
func (m HistoricalPriceDataRejectPointerBuilder) Copy(to HistoricalPriceDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyPointer
func (m HistoricalPriceDataRejectPointer) CopyPointer(to HistoricalPriceDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyPointer
func (m HistoricalPriceDataRejectPointerBuilder) CopyPointer(to HistoricalPriceDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyTo
func (m HistoricalPriceDataRejectPointer) CopyTo(to HistoricalPriceDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyTo
func (m HistoricalPriceDataRejectPointerBuilder) CopyTo(to HistoricalPriceDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyToPointer
func (m HistoricalPriceDataRejectPointer) CopyToPointer(to HistoricalPriceDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyToPointer
func (m HistoricalPriceDataRejectPointerBuilder) CopyToPointer(to HistoricalPriceDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}
