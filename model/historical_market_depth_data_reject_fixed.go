package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size             = HistoricalMarketDepthDataRejectFixedSize  (108)
//     Type             = HISTORICAL_MARKET_DEPTH_DATA_REJECT  (902)
//     RequestID        = 0
//     RejectText       = ""
//     RejectReasonCode = HPDR_UNSET  (0)
// }
var _HistoricalMarketDepthDataRejectFixedDefault = []byte{108, 0, 134, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalMarketDepthDataRejectFixedSize = 108

type HistoricalMarketDepthDataRejectFixed struct {
	message.Fixed
}

type HistoricalMarketDepthDataRejectFixedBuilder struct {
	message.Fixed
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

// ToBuilder
func (m HistoricalMarketDepthDataRejectFixed) ToBuilder() HistoricalMarketDepthDataRejectFixedBuilder {
	return HistoricalMarketDepthDataRejectFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalMarketDepthDataRejectFixedBuilder) Finish() HistoricalMarketDepthDataRejectFixed {
	return HistoricalMarketDepthDataRejectFixed{m.Fixed}
}

// RequestID
func (m HistoricalMarketDepthDataRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalMarketDepthDataRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RejectText
func (m HistoricalMarketDepthDataRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText
func (m HistoricalMarketDepthDataRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectReasonCode
func (m HistoricalMarketDepthDataRejectFixed) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16Fixed(m.Unsafe(), 106, 104))
}

// RejectReasonCode
func (m HistoricalMarketDepthDataRejectFixedBuilder) RejectReasonCode() HistoricalPriceDataRejectReasonCodeEnum {
	return HistoricalPriceDataRejectReasonCodeEnum(message.Int16Fixed(m.Unsafe(), 106, 104))
}

// SetRequestID
func (m HistoricalMarketDepthDataRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRejectText
func (m HistoricalMarketDepthDataRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectReasonCode
func (m HistoricalMarketDepthDataRejectFixedBuilder) SetRejectReasonCode(value HistoricalPriceDataRejectReasonCodeEnum) {
	message.SetInt16Fixed(m.Unsafe(), 106, 104, int16(value))
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
