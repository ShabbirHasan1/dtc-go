package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalMarketDepthDataRejectFixedPointer struct {
	message.FixedPointer
}

type HistoricalMarketDepthDataRejectFixedPointerBuilder struct {
	message.FixedPointer
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
func (m HistoricalMarketDepthDataRejectFixedPointer) ToBuilder() HistoricalMarketDepthDataRejectFixedPointerBuilder {
	return HistoricalMarketDepthDataRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalMarketDepthDataRejectFixedPointerBuilder) Finish() HistoricalMarketDepthDataRejectFixedPointer {
	return HistoricalMarketDepthDataRejectFixedPointer{m.FixedPointer}
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
func (m HistoricalMarketDepthDataRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
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
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// SetRejectReasonCode
func (m HistoricalMarketDepthDataRejectFixedPointerBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16Fixed(m.Ptr, 106, 104, int16(value))
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
