package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalMarketDepthDataRejectPointer struct {
	message.VLSPointer
}

type HistoricalMarketDepthDataRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocHistoricalMarketDepthDataReject() HistoricalMarketDepthDataRejectPointerBuilder {
	a := HistoricalMarketDepthDataRejectPointerBuilder{message.AllocVLSBuilder(20)}
	a.Ptr.SetBytes(0, _HistoricalMarketDepthDataRejectDefault)
	return a
}

func AllocHistoricalMarketDepthDataRejectFrom(b []byte) HistoricalMarketDepthDataRejectPointer {
	return HistoricalMarketDepthDataRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
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

// ToBuilder
func (m HistoricalMarketDepthDataRejectPointer) ToBuilder() HistoricalMarketDepthDataRejectPointerBuilder {
	return HistoricalMarketDepthDataRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalMarketDepthDataRejectPointerBuilder) Finish() HistoricalMarketDepthDataRejectPointer {
	return HistoricalMarketDepthDataRejectPointer{m.VLSPointerBuilder.Finish()}
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
func (m HistoricalMarketDepthDataRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText
func (m HistoricalMarketDepthDataRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectReasonCode
func (m HistoricalMarketDepthDataRejectPointer) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Ptr, 18, 16))
}

// RejectReasonCode
func (m HistoricalMarketDepthDataRejectPointerBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16VLS(m.Ptr, 18, 16))
}

// SetRequestID
func (m HistoricalMarketDepthDataRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText
func (m *HistoricalMarketDepthDataRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRejectReasonCode
func (m HistoricalMarketDepthDataRejectPointerBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16VLS(m.Ptr, 18, 16, int16(value))
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
