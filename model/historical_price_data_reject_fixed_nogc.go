package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HistoricalPriceDataRejectFixedPointer When the Server rejects a historical price data request from the Client,
// a HistoricalPriceDataReject message will be sent.
//
// This message is never compressed.
type HistoricalPriceDataRejectFixedPointer struct {
	message.FixedPointer
}

// HistoricalPriceDataRejectFixedPointerBuilder When the Server rejects a historical price data request from the Client,
// a HistoricalPriceDataReject message will be sent.
//
// This message is never compressed.
type HistoricalPriceDataRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalPriceDataRejectFixed() HistoricalPriceDataRejectFixedPointerBuilder {
	a := HistoricalPriceDataRejectFixedPointerBuilder{message.AllocFixed(108)}
	a.Ptr.SetBytes(0, _HistoricalPriceDataRejectFixedDefault)
	return a
}

func AllocHistoricalPriceDataRejectFixedFrom(b []byte) HistoricalPriceDataRejectFixedPointer {
	return HistoricalPriceDataRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size               = HistoricalPriceDataRejectFixedSize  (108)
//     Type               = HISTORICAL_PRICE_DATA_REJECT  (802)
//     RequestID          = 0
//     RejectText         = ""
//     RejectReasonCode   = 0
//     RetryTimeInSeconds = 0
// }
func (m HistoricalPriceDataRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataRejectFixedDefault)
}

// ToBuilder
func (m HistoricalPriceDataRejectFixedPointer) ToBuilder() HistoricalPriceDataRejectFixedPointerBuilder {
	return HistoricalPriceDataRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalPriceDataRejectFixedPointerBuilder) Finish() HistoricalPriceDataRejectFixedPointer {
	return HistoricalPriceDataRejectFixedPointer{m.FixedPointer}
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText Text reason for rejection.
func (m HistoricalPriceDataRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText Text reason for rejection.
func (m HistoricalPriceDataRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectFixedPointer) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16Fixed(m.Ptr, 106, 104))
}

// RejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectFixedPointerBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16Fixed(m.Ptr, 106, 104))
}

// RetryTimeInSeconds This is an optional field from the Server. This field will normally be
// zero.
//
// If a retry is intended to be performed, the server may give an indication
// of how long to wait in seconds. This field indicates that.
//
// This field is not recommended to be used. If it is used, it is really
// an indication of a substandard Server.
func (m HistoricalPriceDataRejectFixedPointer) RetryTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Ptr, 108, 106)
}

// RetryTimeInSeconds This is an optional field from the Server. This field will normally be
// zero.
//
// If a retry is intended to be performed, the server may give an indication
// of how long to wait in seconds. This field indicates that.
//
// This field is not recommended to be used. If it is used, it is really
// an indication of a substandard Server.
func (m HistoricalPriceDataRejectFixedPointerBuilder) RetryTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Ptr, 108, 106)
}

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Text reason for rejection.
func (m HistoricalPriceDataRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// SetRejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectFixedPointerBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16Fixed(m.Ptr, 106, 104, int16(value))
}

// SetRetryTimeInSeconds This is an optional field from the Server. This field will normally be
// zero.
//
// If a retry is intended to be performed, the server may give an indication
// of how long to wait in seconds. This field indicates that.
//
// This field is not recommended to be used. If it is used, it is really
// an indication of a substandard Server.
func (m HistoricalPriceDataRejectFixedPointerBuilder) SetRetryTimeInSeconds(value uint16) {
	message.SetUint16Fixed(m.Ptr, 108, 106, value)
}

// Copy
func (m HistoricalPriceDataRejectFixedPointer) Copy(to HistoricalPriceDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// Copy
func (m HistoricalPriceDataRejectFixedPointerBuilder) Copy(to HistoricalPriceDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyPointer
func (m HistoricalPriceDataRejectFixedPointer) CopyPointer(to HistoricalPriceDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyPointer
func (m HistoricalPriceDataRejectFixedPointerBuilder) CopyPointer(to HistoricalPriceDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyTo
func (m HistoricalPriceDataRejectFixedPointer) CopyTo(to HistoricalPriceDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyTo
func (m HistoricalPriceDataRejectFixedPointerBuilder) CopyTo(to HistoricalPriceDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyToPointer
func (m HistoricalPriceDataRejectFixedPointer) CopyToPointer(to HistoricalPriceDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyToPointer
func (m HistoricalPriceDataRejectFixedPointerBuilder) CopyToPointer(to HistoricalPriceDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}
