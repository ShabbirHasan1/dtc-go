package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size               = HistoricalPriceDataRejectSize  (20)
//     Type               = HISTORICAL_PRICE_DATA_REJECT  (802)
//     BaseSize           = HistoricalPriceDataRejectSize  (20)
//     RequestID          = 0
//     RejectText         = ""
//     RejectReasonCode   = 0
//     RetryTimeInSeconds = 0
// }
var _HistoricalPriceDataRejectDefault = []byte{20, 0, 34, 3, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalPriceDataRejectSize = 20

// HistoricalPriceDataReject When the Server rejects a historical price data request from the Client,
// a HistoricalPriceDataReject message will be sent.
//
// This message is never compressed.
type HistoricalPriceDataReject struct {
	message.VLS
}

// HistoricalPriceDataRejectBuilder When the Server rejects a historical price data request from the Client,
// a HistoricalPriceDataReject message will be sent.
//
// This message is never compressed.
type HistoricalPriceDataRejectBuilder struct {
	message.VLSBuilder
}

func NewHistoricalPriceDataRejectFrom(b []byte) HistoricalPriceDataReject {
	return HistoricalPriceDataReject{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalPriceDataReject(b []byte) HistoricalPriceDataReject {
	return HistoricalPriceDataReject{VLS: message.WrapVLS(b)}
}

func NewHistoricalPriceDataReject() HistoricalPriceDataRejectBuilder {
	a := HistoricalPriceDataRejectBuilder{message.NewVLSBuilder(20)}
	a.Unsafe().SetBytes(0, _HistoricalPriceDataRejectDefault)
	return a
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
func (m HistoricalPriceDataRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataRejectDefault)
}

// ToBuilder
func (m HistoricalPriceDataReject) ToBuilder() HistoricalPriceDataRejectBuilder {
	return HistoricalPriceDataRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m HistoricalPriceDataRejectBuilder) Finish() HistoricalPriceDataReject {
	return HistoricalPriceDataReject{m.VLSBuilder.Finish()}
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RejectText Text reason for rejection.
func (m HistoricalPriceDataReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Text reason for rejection.
func (m HistoricalPriceDataRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataReject) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Unsafe(), 18, 16))
}

// RejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Unsafe(), 18, 16))
}

// RetryTimeInSeconds This is an optional field from the Server. This field will normally be
// zero.
//
// If a retry is intended to be performed, the server may give an indication
// of how long to wait in seconds. This field indicates that.
//
// This field is not recommended to be used. If it is used, it is really
// an indication of a substandard Server.
func (m HistoricalPriceDataReject) RetryTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Unsafe(), 20, 18)
}

// RetryTimeInSeconds This is an optional field from the Server. This field will normally be
// zero.
//
// If a retry is intended to be performed, the server may give an indication
// of how long to wait in seconds. This field indicates that.
//
// This field is not recommended to be used. If it is used, it is really
// an indication of a substandard Server.
func (m HistoricalPriceDataRejectBuilder) RetryTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Unsafe(), 20, 18)
}

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText Text reason for rejection.
func (m *HistoricalPriceDataRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16VLS(m.Unsafe(), 18, 16, int16(value))
}

// SetRetryTimeInSeconds This is an optional field from the Server. This field will normally be
// zero.
//
// If a retry is intended to be performed, the server may give an indication
// of how long to wait in seconds. This field indicates that.
//
// This field is not recommended to be used. If it is used, it is really
// an indication of a substandard Server.
func (m HistoricalPriceDataRejectBuilder) SetRetryTimeInSeconds(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 20, 18, value)
}

// Copy
func (m HistoricalPriceDataReject) Copy(to HistoricalPriceDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// Copy
func (m HistoricalPriceDataRejectBuilder) Copy(to HistoricalPriceDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyPointer
func (m HistoricalPriceDataReject) CopyPointer(to HistoricalPriceDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyPointer
func (m HistoricalPriceDataRejectBuilder) CopyPointer(to HistoricalPriceDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyToPointer
func (m HistoricalPriceDataReject) CopyToPointer(to HistoricalPriceDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyToPointer
func (m HistoricalPriceDataRejectBuilder) CopyToPointer(to HistoricalPriceDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyTo
func (m HistoricalPriceDataReject) CopyTo(to HistoricalPriceDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyTo
func (m HistoricalPriceDataRejectBuilder) CopyTo(to HistoricalPriceDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}
