package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalPriceDataRejectSize = 20

const HistoricalPriceDataRejectFixedSize = 108

// {
//     Size               = HistoricalPriceDataRejectSize  (20)
//     Type               = HISTORICAL_PRICE_DATA_REJECT  (802)
//     BaseSize           = HistoricalPriceDataRejectSize  (20)
//     RequestID          = 0
//     RejectText         = ""
//     RejectReasonCode   = HPDR_UNSET  (0)
//     RetryTimeInSeconds = 0
// }
var _HistoricalPriceDataRejectDefault = []byte{20, 0, 34, 3, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size               = HistoricalPriceDataRejectFixedSize  (108)
//     Type               = HISTORICAL_PRICE_DATA_REJECT  (802)
//     RequestID          = 0
//     RejectText         = ""
//     RejectReasonCode   = HPDR_UNSET  (0)
//     RetryTimeInSeconds = 0
// }
var _HistoricalPriceDataRejectFixedDefault = []byte{108, 0, 34, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocHistoricalPriceDataReject() HistoricalPriceDataRejectPointerBuilder {
	a := HistoricalPriceDataRejectPointerBuilder{message.AllocVLSBuilder(20)}
	a.Ptr.SetBytes(0, _HistoricalPriceDataRejectDefault)
	return a
}

func AllocHistoricalPriceDataRejectFrom(b []byte) HistoricalPriceDataRejectPointer {
	return HistoricalPriceDataRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size               = HistoricalPriceDataRejectSize  (20)
//     Type               = HISTORICAL_PRICE_DATA_REJECT  (802)
//     BaseSize           = HistoricalPriceDataRejectSize  (20)
//     RequestID          = 0
//     RejectText         = ""
//     RejectReasonCode   = HPDR_UNSET  (0)
//     RetryTimeInSeconds = 0
// }
func (m HistoricalPriceDataRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataRejectDefault)
}

// Clear
// {
//     Size               = HistoricalPriceDataRejectFixedSize  (108)
//     Type               = HISTORICAL_PRICE_DATA_REJECT  (802)
//     RequestID          = 0
//     RejectText         = ""
//     RejectReasonCode   = HPDR_UNSET  (0)
//     RetryTimeInSeconds = 0
// }
func (m HistoricalPriceDataRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataRejectFixedDefault)
}

// Clear
// {
//     Size               = HistoricalPriceDataRejectSize  (20)
//     Type               = HISTORICAL_PRICE_DATA_REJECT  (802)
//     BaseSize           = HistoricalPriceDataRejectSize  (20)
//     RequestID          = 0
//     RejectText         = ""
//     RejectReasonCode   = HPDR_UNSET  (0)
//     RetryTimeInSeconds = 0
// }
func (m HistoricalPriceDataRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataRejectDefault)
}

// Clear
// {
//     Size               = HistoricalPriceDataRejectFixedSize  (108)
//     Type               = HISTORICAL_PRICE_DATA_REJECT  (802)
//     RequestID          = 0
//     RejectText         = ""
//     RejectReasonCode   = HPDR_UNSET  (0)
//     RetryTimeInSeconds = 0
// }
func (m HistoricalPriceDataRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataRejectFixedDefault)
}

// ToBuilder
func (m HistoricalPriceDataReject) ToBuilder() HistoricalPriceDataRejectBuilder {
	return HistoricalPriceDataRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalPriceDataRejectFixed) ToBuilder() HistoricalPriceDataRejectFixedBuilder {
	return HistoricalPriceDataRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalPriceDataRejectPointer) ToBuilder() HistoricalPriceDataRejectPointerBuilder {
	return HistoricalPriceDataRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalPriceDataRejectFixedPointer) ToBuilder() HistoricalPriceDataRejectFixedPointerBuilder {
	return HistoricalPriceDataRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalPriceDataRejectBuilder) Finish() HistoricalPriceDataReject {
	return HistoricalPriceDataReject{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalPriceDataRejectFixedBuilder) Finish() HistoricalPriceDataRejectFixed {
	return HistoricalPriceDataRejectFixed{m.Fixed}
}

// Finish
func (m *HistoricalPriceDataRejectPointerBuilder) Finish() HistoricalPriceDataRejectPointer {
	return HistoricalPriceDataRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalPriceDataRejectFixedPointerBuilder) Finish() HistoricalPriceDataRejectFixedPointer {
	return HistoricalPriceDataRejectFixedPointer{m.FixedPointer}
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
func (m HistoricalPriceDataReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Text reason for rejection.
func (m HistoricalPriceDataRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
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
func (m HistoricalPriceDataReject) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Unsafe(), 18, 16))
}

// RejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Unsafe(), 18, 16))
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
func (m HistoricalPriceDataRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Text reason for rejection.
func (m HistoricalPriceDataRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
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
func (m HistoricalPriceDataRejectFixed) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16Fixed(m.Unsafe(), 106, 104))
}

// RejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectFixedBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16Fixed(m.Unsafe(), 106, 104))
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
func (m HistoricalPriceDataRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Text reason for rejection.
func (m *HistoricalPriceDataRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText Text reason for rejection.
func (m *HistoricalPriceDataRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16VLS(m.Unsafe(), 18, 16, int16(value))
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
func (m HistoricalPriceDataRejectBuilder) SetRetryTimeInSeconds(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 20, 18, value)
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

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Text reason for rejection.
func (m HistoricalPriceDataRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText Text reason for rejection.
func (m HistoricalPriceDataRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// SetRejectReasonCode Integer identifier identifying the reason for the rejection. For the text
// reason, refer to the RejectText field.
func (m HistoricalPriceDataRejectFixedBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16Fixed(m.Unsafe(), 106, 104, int16(value))
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
func (m HistoricalPriceDataRejectFixedBuilder) SetRetryTimeInSeconds(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 108, 106, value)
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
