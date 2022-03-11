package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size             = HistoricalMarketDepthDataRejectSize  (20)
//     Type             = HISTORICAL_MARKET_DEPTH_DATA_REJECT  (902)
//     BaseSize         = HistoricalMarketDepthDataRejectSize  (20)
//     RequestID        = 0
//     RejectText       = ""
//     RejectReasonCode = HPDR_UNSET  (0)
// }
var _HistoricalMarketDepthDataRejectDefault = []byte{20, 0, 134, 3, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalMarketDepthDataRejectSize = 20

type HistoricalMarketDepthDataReject struct {
	message.VLS
}

type HistoricalMarketDepthDataRejectBuilder struct {
	message.VLSBuilder
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

// ToBuilder
func (m HistoricalMarketDepthDataReject) ToBuilder() HistoricalMarketDepthDataRejectBuilder {
	return HistoricalMarketDepthDataRejectBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m HistoricalMarketDepthDataRejectBuilder) Finish() HistoricalMarketDepthDataReject {
	return HistoricalMarketDepthDataReject{m.VLSBuilder.Finish()}
}

// RequestID
func (m HistoricalMarketDepthDataReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m HistoricalMarketDepthDataRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RejectText
func (m HistoricalMarketDepthDataReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText
func (m HistoricalMarketDepthDataRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectReasonCode
func (m HistoricalMarketDepthDataReject) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Unsafe(), 18, 16))
}

// RejectReasonCode
func (m HistoricalMarketDepthDataRejectBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Unsafe(), 18, 16))
}

// SetRequestID
func (m HistoricalMarketDepthDataRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRejectText
func (m *HistoricalMarketDepthDataRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectReasonCode
func (m HistoricalMarketDepthDataRejectBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16VLS(m.Unsafe(), 18, 16, int16(value))
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
