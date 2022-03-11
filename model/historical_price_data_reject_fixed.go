package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size               = HistoricalPriceDataRejectFixedSize  (108)
//     Type               = HISTORICAL_PRICE_DATA_REJECT  (802)
//     RequestID          = 0
//     RejectText         = ""
//     RejectReasonCode   = 0
//     RetryTimeInSeconds = 0
// }
var _HistoricalPriceDataRejectFixedDefault = []byte{108, 0, 34, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalPriceDataRejectFixedSize = 108

// HistoricalPriceDataRejectFixed When the Server rejects a historical price data request from the Client,
// a HistoricalPriceDataReject message will be sent.
//
// This message is never compressed.
type HistoricalPriceDataRejectFixed struct {
	message.Fixed
}

// HistoricalPriceDataRejectFixedBuilder When the Server rejects a historical price data request from the Client,
// a HistoricalPriceDataReject message will be sent.
//
// This message is never compressed.
type HistoricalPriceDataRejectFixedBuilder struct {
	message.Fixed
}

func NewHistoricalPriceDataRejectFixedFrom(b []byte) HistoricalPriceDataRejectFixed {
	return HistoricalPriceDataRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalPriceDataRejectFixed(b []byte) HistoricalPriceDataRejectFixed {
	return HistoricalPriceDataRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalPriceDataRejectFixed() HistoricalPriceDataRejectFixedBuilder {
	a := HistoricalPriceDataRejectFixedBuilder{message.NewFixed(108)}
	a.Unsafe().SetBytes(0, _HistoricalPriceDataRejectFixedDefault)
	return a
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
func (m HistoricalPriceDataRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataRejectFixedDefault)
}

// ToBuilder
func (m HistoricalPriceDataRejectFixed) ToBuilder() HistoricalPriceDataRejectFixedBuilder {
	return HistoricalPriceDataRejectFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalPriceDataRejectFixedBuilder) Finish() HistoricalPriceDataRejectFixed {
	return HistoricalPriceDataRejectFixed{m.Fixed}
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RejectText Text reason for rejection.
func (m HistoricalPriceDataRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Text reason for rejection.
func (m HistoricalPriceDataRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectFixed) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16Fixed(m.Unsafe(), 106, 104))
}

// RejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectFixedBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16Fixed(m.Unsafe(), 106, 104))
}

// RetryTimeInSeconds This is an optional field from the Server. This field will normally be
// zero.
//
// If a retry is intended to be performed, the server may give an indication
// of how long to wait in seconds. This field indicates that.
//
// This field is not recommended to be used. If it is used, it is really
// an indication of a substandard Server.
func (m HistoricalPriceDataRejectFixed) RetryTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 108, 106)
}

// RetryTimeInSeconds This is an optional field from the Server. This field will normally be
// zero.
//
// If a retry is intended to be performed, the server may give an indication
// of how long to wait in seconds. This field indicates that.
//
// This field is not recommended to be used. If it is used, it is really
// an indication of a substandard Server.
func (m HistoricalPriceDataRejectFixedBuilder) RetryTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 108, 106)
}

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText Text reason for rejection.
func (m HistoricalPriceDataRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectFixedBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16Fixed(m.Unsafe(), 106, 104, int16(value))
}

// SetRetryTimeInSeconds This is an optional field from the Server. This field will normally be
// zero.
//
// If a retry is intended to be performed, the server may give an indication
// of how long to wait in seconds. This field indicates that.
//
// This field is not recommended to be used. If it is used, it is really
// an indication of a substandard Server.
func (m HistoricalPriceDataRejectFixedBuilder) SetRetryTimeInSeconds(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 108, 106, value)
}

// Copy
func (m HistoricalPriceDataRejectFixed) Copy(to HistoricalPriceDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// Copy
func (m HistoricalPriceDataRejectFixedBuilder) Copy(to HistoricalPriceDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyPointer
func (m HistoricalPriceDataRejectFixed) CopyPointer(to HistoricalPriceDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyPointer
func (m HistoricalPriceDataRejectFixedBuilder) CopyPointer(to HistoricalPriceDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyToPointer
func (m HistoricalPriceDataRejectFixed) CopyToPointer(to HistoricalPriceDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyToPointer
func (m HistoricalPriceDataRejectFixedBuilder) CopyToPointer(to HistoricalPriceDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyTo
func (m HistoricalPriceDataRejectFixed) CopyTo(to HistoricalPriceDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}

// CopyTo
func (m HistoricalPriceDataRejectFixedBuilder) CopyTo(to HistoricalPriceDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
	to.SetRetryTimeInSeconds(m.RetryTimeInSeconds())
}
