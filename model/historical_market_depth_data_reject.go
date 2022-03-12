package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalMarketDepthDataRejectSize = 20

const HistoricalMarketDepthDataRejectFixedSize = 108

// {
//     Size             = HistoricalMarketDepthDataRejectSize  (20)
//     Type             = HISTORICAL_MARKET_DEPTH_DATA_REJECT  (902)
//     BaseSize         = HistoricalMarketDepthDataRejectSize  (20)
//     RequestID        = 0
//     RejectText       = ""
//     RejectReasonCode = HPDR_UNSET  (0)
// }
var _HistoricalMarketDepthDataRejectDefault = []byte{20, 0, 134, 3, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size             = HistoricalMarketDepthDataRejectFixedSize  (108)
//     Type             = HISTORICAL_MARKET_DEPTH_DATA_REJECT  (902)
//     RequestID        = 0
//     RejectText       = ""
//     RejectReasonCode = HPDR_UNSET  (0)
// }
var _HistoricalMarketDepthDataRejectFixedDefault = []byte{108, 0, 134, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HistoricalMarketDepthDataReject struct {
	message.VLS
}

type HistoricalMarketDepthDataRejectBuilder struct {
	message.VLSBuilder
}

type HistoricalMarketDepthDataRejectFixed struct {
	message.Fixed
}

type HistoricalMarketDepthDataRejectFixedBuilder struct {
	message.Fixed
}

type HistoricalMarketDepthDataRejectPointer struct {
	message.VLSPointer
}

type HistoricalMarketDepthDataRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

type HistoricalMarketDepthDataRejectFixedPointer struct {
	message.FixedPointer
}

type HistoricalMarketDepthDataRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalMarketDepthDataRejectFrom(b []byte) HistoricalMarketDepthDataReject {
	return HistoricalMarketDepthDataReject{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalMarketDepthDataReject(b []byte) HistoricalMarketDepthDataReject {
	return HistoricalMarketDepthDataReject{VLS: message.WrapVLS(b)}
}

func NewHistoricalMarketDepthDataReject() HistoricalMarketDepthDataRejectBuilder {
	a := HistoricalMarketDepthDataRejectBuilder{message.NewVLSBuilder(20)}
	a.Unsafe().SetBytes(0, _HistoricalMarketDepthDataRejectDefault)
	return a
}

func NewHistoricalMarketDepthDataRejectFixedFrom(b []byte) HistoricalMarketDepthDataRejectFixed {
	return HistoricalMarketDepthDataRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalMarketDepthDataRejectFixed(b []byte) HistoricalMarketDepthDataRejectFixed {
	return HistoricalMarketDepthDataRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalMarketDepthDataRejectFixed() HistoricalMarketDepthDataRejectFixedBuilder {
	a := HistoricalMarketDepthDataRejectFixedBuilder{message.NewFixed(108)}
	a.Unsafe().SetBytes(0, _HistoricalMarketDepthDataRejectFixedDefault)
	return a
}

func AllocHistoricalMarketDepthDataReject() HistoricalMarketDepthDataRejectPointerBuilder {
	a := HistoricalMarketDepthDataRejectPointerBuilder{message.AllocVLSBuilder(20)}
	a.Ptr.SetBytes(0, _HistoricalMarketDepthDataRejectDefault)
	return a
}

func AllocHistoricalMarketDepthDataRejectFrom(b []byte) HistoricalMarketDepthDataRejectPointer {
	return HistoricalMarketDepthDataRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocHistoricalMarketDepthDataRejectFixed() HistoricalMarketDepthDataRejectFixedPointerBuilder {
	a := HistoricalMarketDepthDataRejectFixedPointerBuilder{message.AllocFixed(108)}
	a.Ptr.SetBytes(0, _HistoricalMarketDepthDataRejectFixedDefault)
	return a
}

func AllocHistoricalMarketDepthDataRejectFixedFrom(b []byte) HistoricalMarketDepthDataRejectFixedPointer {
	return HistoricalMarketDepthDataRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size             = HistoricalMarketDepthDataRejectSize  (20)
//     Type             = HISTORICAL_MARKET_DEPTH_DATA_REJECT  (902)
//     BaseSize         = HistoricalMarketDepthDataRejectSize  (20)
//     RequestID        = 0
//     RejectText       = ""
//     RejectReasonCode = HPDR_UNSET  (0)
// }
func (m HistoricalMarketDepthDataRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalMarketDepthDataRejectDefault)
}

// Clear
// {
//     Size             = HistoricalMarketDepthDataRejectFixedSize  (108)
//     Type             = HISTORICAL_MARKET_DEPTH_DATA_REJECT  (902)
//     RequestID        = 0
//     RejectText       = ""
//     RejectReasonCode = HPDR_UNSET  (0)
// }
func (m HistoricalMarketDepthDataRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalMarketDepthDataRejectFixedDefault)
}

// Clear
// {
//     Size             = HistoricalMarketDepthDataRejectSize  (20)
//     Type             = HISTORICAL_MARKET_DEPTH_DATA_REJECT  (902)
//     BaseSize         = HistoricalMarketDepthDataRejectSize  (20)
//     RequestID        = 0
//     RejectText       = ""
//     RejectReasonCode = HPDR_UNSET  (0)
// }
func (m HistoricalMarketDepthDataRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalMarketDepthDataRejectDefault)
}

// Clear
// {
//     Size             = HistoricalMarketDepthDataRejectFixedSize  (108)
//     Type             = HISTORICAL_MARKET_DEPTH_DATA_REJECT  (902)
//     RequestID        = 0
//     RejectText       = ""
//     RejectReasonCode = HPDR_UNSET  (0)
// }
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalMarketDepthDataRejectFixedDefault)
}

// ToBuilder
func (m HistoricalMarketDepthDataReject) ToBuilder() HistoricalMarketDepthDataRejectBuilder {
	return HistoricalMarketDepthDataRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalMarketDepthDataRejectFixed) ToBuilder() HistoricalMarketDepthDataRejectFixedBuilder {
	return HistoricalMarketDepthDataRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalMarketDepthDataRejectPointer) ToBuilder() HistoricalMarketDepthDataRejectPointerBuilder {
	return HistoricalMarketDepthDataRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalMarketDepthDataRejectFixedPointer) ToBuilder() HistoricalMarketDepthDataRejectFixedPointerBuilder {
	return HistoricalMarketDepthDataRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalMarketDepthDataRejectBuilder) Finish() HistoricalMarketDepthDataReject {
	return HistoricalMarketDepthDataReject{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalMarketDepthDataRejectFixedBuilder) Finish() HistoricalMarketDepthDataRejectFixed {
	return HistoricalMarketDepthDataRejectFixed{m.Fixed}
}

// Finish
func (m *HistoricalMarketDepthDataRejectPointerBuilder) Finish() HistoricalMarketDepthDataRejectPointer {
	return HistoricalMarketDepthDataRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalMarketDepthDataRejectFixedPointerBuilder) Finish() HistoricalMarketDepthDataRejectFixedPointer {
	return HistoricalMarketDepthDataRejectFixedPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalMarketDepthDataReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m HistoricalMarketDepthDataRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m HistoricalMarketDepthDataRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID
func (m HistoricalMarketDepthDataRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText
func (m HistoricalMarketDepthDataReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText
func (m HistoricalMarketDepthDataRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText
func (m HistoricalMarketDepthDataRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText
func (m HistoricalMarketDepthDataRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectReasonCode
func (m HistoricalMarketDepthDataReject) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Unsafe(), 18, 16))
}

// RejectReasonCode
func (m HistoricalMarketDepthDataRejectBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Unsafe(), 18, 16))
}

// RejectReasonCode
func (m HistoricalMarketDepthDataRejectPointer) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Ptr, 18, 16))
}

// RejectReasonCode
func (m HistoricalMarketDepthDataRejectPointerBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Ptr, 18, 16))
}

// RequestID
func (m HistoricalMarketDepthDataRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalMarketDepthDataRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalMarketDepthDataRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText
func (m HistoricalMarketDepthDataRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m HistoricalMarketDepthDataRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m HistoricalMarketDepthDataRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectReasonCode
func (m HistoricalMarketDepthDataRejectFixed) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16Fixed(m.Unsafe(), 106, 104))
}

// RejectReasonCode
func (m HistoricalMarketDepthDataRejectFixedBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16Fixed(m.Unsafe(), 106, 104))
}

// RejectReasonCode
func (m HistoricalMarketDepthDataRejectFixedPointer) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16Fixed(m.Ptr, 106, 104))
}

// RejectReasonCode
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16Fixed(m.Ptr, 106, 104))
}

// SetRequestID
func (m HistoricalMarketDepthDataRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID
func (m HistoricalMarketDepthDataRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText
func (m *HistoricalMarketDepthDataRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText
func (m *HistoricalMarketDepthDataRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRejectReasonCode
func (m HistoricalMarketDepthDataRejectBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16VLS(m.Unsafe(), 18, 16, int16(value))
}

// SetRejectReasonCode
func (m HistoricalMarketDepthDataRejectPointerBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16VLS(m.Ptr, 18, 16, int16(value))
}

// SetRequestID
func (m HistoricalMarketDepthDataRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText
func (m HistoricalMarketDepthDataRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// SetRejectReasonCode
func (m HistoricalMarketDepthDataRejectFixedBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16Fixed(m.Unsafe(), 106, 104, int16(value))
}

// SetRejectReasonCode
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16Fixed(m.Ptr, 106, 104, int16(value))
}

// Copy
func (m HistoricalMarketDepthDataReject) Copy(to HistoricalMarketDepthDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// Copy
func (m HistoricalMarketDepthDataRejectBuilder) Copy(to HistoricalMarketDepthDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyPointer
func (m HistoricalMarketDepthDataReject) CopyPointer(to HistoricalMarketDepthDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyPointer
func (m HistoricalMarketDepthDataRejectBuilder) CopyPointer(to HistoricalMarketDepthDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyToPointer
func (m HistoricalMarketDepthDataReject) CopyToPointer(to HistoricalMarketDepthDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRejectBuilder) CopyToPointer(to HistoricalMarketDepthDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyTo
func (m HistoricalMarketDepthDataReject) CopyTo(to HistoricalMarketDepthDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyTo
func (m HistoricalMarketDepthDataRejectBuilder) CopyTo(to HistoricalMarketDepthDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// Copy
func (m HistoricalMarketDepthDataRejectFixed) Copy(to HistoricalMarketDepthDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// Copy
func (m HistoricalMarketDepthDataRejectFixedBuilder) Copy(to HistoricalMarketDepthDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyPointer
func (m HistoricalMarketDepthDataRejectFixed) CopyPointer(to HistoricalMarketDepthDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyPointer
func (m HistoricalMarketDepthDataRejectFixedBuilder) CopyPointer(to HistoricalMarketDepthDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRejectFixed) CopyToPointer(to HistoricalMarketDepthDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRejectFixedBuilder) CopyToPointer(to HistoricalMarketDepthDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyTo
func (m HistoricalMarketDepthDataRejectFixed) CopyTo(to HistoricalMarketDepthDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyTo
func (m HistoricalMarketDepthDataRejectFixedBuilder) CopyTo(to HistoricalMarketDepthDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// Copy
func (m HistoricalMarketDepthDataRejectPointer) Copy(to HistoricalMarketDepthDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// Copy
func (m HistoricalMarketDepthDataRejectPointerBuilder) Copy(to HistoricalMarketDepthDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyPointer
func (m HistoricalMarketDepthDataRejectPointer) CopyPointer(to HistoricalMarketDepthDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyPointer
func (m HistoricalMarketDepthDataRejectPointerBuilder) CopyPointer(to HistoricalMarketDepthDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyTo
func (m HistoricalMarketDepthDataRejectPointer) CopyTo(to HistoricalMarketDepthDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyTo
func (m HistoricalMarketDepthDataRejectPointerBuilder) CopyTo(to HistoricalMarketDepthDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRejectPointer) CopyToPointer(to HistoricalMarketDepthDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRejectPointerBuilder) CopyToPointer(to HistoricalMarketDepthDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// Copy
func (m HistoricalMarketDepthDataRejectFixedPointer) Copy(to HistoricalMarketDepthDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// Copy
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) Copy(to HistoricalMarketDepthDataRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyPointer
func (m HistoricalMarketDepthDataRejectFixedPointer) CopyPointer(to HistoricalMarketDepthDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyPointer
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) CopyPointer(to HistoricalMarketDepthDataRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyTo
func (m HistoricalMarketDepthDataRejectFixedPointer) CopyTo(to HistoricalMarketDepthDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyTo
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) CopyTo(to HistoricalMarketDepthDataRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRejectFixedPointer) CopyToPointer(to HistoricalMarketDepthDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) CopyToPointer(to HistoricalMarketDepthDataRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
	to.SetRejectReasonCode(m.RejectReasonCode())
}
