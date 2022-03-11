package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalMarketDepthDataRequestPointer struct {
	message.VLSPointer
}

type HistoricalMarketDepthDataRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocHistoricalMarketDepthDataRequest() HistoricalMarketDepthDataRequestPointerBuilder {
	a := HistoricalMarketDepthDataRequestPointerBuilder{message.AllocVLSBuilder(48)}
	a.Ptr.SetBytes(0, _HistoricalMarketDepthDataRequestDefault)
	return a
}

func AllocHistoricalMarketDepthDataRequestFrom(b []byte) HistoricalMarketDepthDataRequestPointer {
	return HistoricalMarketDepthDataRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size               = HistoricalMarketDepthDataRequestSize  (48)
//     Type               = HISTORICAL_MARKET_DEPTH_DATA_REQUEST  (900)
//     BaseSize           = HistoricalMarketDepthDataRequestSize  (48)
//     RequestID          = 0
//     Symbol             = ""
//     Exchange           = ""
//     StartDateTime      = 0
//     EndDateTime        = 0
//     UseZLibCompression = 0
//     Integer_1          = 0
// }
func (m HistoricalMarketDepthDataRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalMarketDepthDataRequestDefault)
}

// ToBuilder
func (m HistoricalMarketDepthDataRequestPointer) ToBuilder() HistoricalMarketDepthDataRequestPointerBuilder {
	return HistoricalMarketDepthDataRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalMarketDepthDataRequestPointerBuilder) Finish() HistoricalMarketDepthDataRequestPointer {
	return HistoricalMarketDepthDataRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m HistoricalMarketDepthDataRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID
func (m HistoricalMarketDepthDataRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// Symbol
func (m HistoricalMarketDepthDataRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Symbol
func (m HistoricalMarketDepthDataRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Exchange
func (m HistoricalMarketDepthDataRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Exchange
func (m HistoricalMarketDepthDataRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// StartDateTime
func (m HistoricalMarketDepthDataRequestPointer) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Ptr, 32, 24))
}

// StartDateTime
func (m HistoricalMarketDepthDataRequestPointerBuilder) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Ptr, 32, 24))
}

// EndDateTime
func (m HistoricalMarketDepthDataRequestPointer) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Ptr, 40, 32))
}

// EndDateTime
func (m HistoricalMarketDepthDataRequestPointerBuilder) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Ptr, 40, 32))
}

// UseZLibCompression
func (m HistoricalMarketDepthDataRequestPointer) UseZLibCompression() uint8 {
	return message.Uint8VLS(m.Ptr, 41, 40)
}

// UseZLibCompression
func (m HistoricalMarketDepthDataRequestPointerBuilder) UseZLibCompression() uint8 {
	return message.Uint8VLS(m.Ptr, 41, 40)
}

// Integer_1
func (m HistoricalMarketDepthDataRequestPointer) Integer_1() uint8 {
	return message.Uint8VLS(m.Ptr, 42, 41)
}

// Integer_1
func (m HistoricalMarketDepthDataRequestPointerBuilder) Integer_1() uint8 {
	return message.Uint8VLS(m.Ptr, 42, 41)
}

// SetRequestID
func (m HistoricalMarketDepthDataRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetSymbol
func (m *HistoricalMarketDepthDataRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetExchange
func (m *HistoricalMarketDepthDataRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetStartDateTime
func (m HistoricalMarketDepthDataRequestPointerBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64VLS(m.Ptr, 32, 24, int64(value))
}

// SetEndDateTime
func (m HistoricalMarketDepthDataRequestPointerBuilder) SetEndDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64VLS(m.Ptr, 40, 32, int64(value))
}

// SetUseZLibCompression
func (m HistoricalMarketDepthDataRequestPointerBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8VLS(m.Ptr, 41, 40, value)
}

// SetInteger_1
func (m HistoricalMarketDepthDataRequestPointerBuilder) SetInteger_1(value uint8) {
	message.SetUint8VLS(m.Ptr, 42, 41, value)
}

// Copy
func (m HistoricalMarketDepthDataRequestPointer) Copy(to HistoricalMarketDepthDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// Copy
func (m HistoricalMarketDepthDataRequestPointerBuilder) Copy(to HistoricalMarketDepthDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyPointer
func (m HistoricalMarketDepthDataRequestPointer) CopyPointer(to HistoricalMarketDepthDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyPointer
func (m HistoricalMarketDepthDataRequestPointerBuilder) CopyPointer(to HistoricalMarketDepthDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyTo
func (m HistoricalMarketDepthDataRequestPointer) CopyTo(to HistoricalMarketDepthDataRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyTo
func (m HistoricalMarketDepthDataRequestPointerBuilder) CopyTo(to HistoricalMarketDepthDataRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRequestPointer) CopyToPointer(to HistoricalMarketDepthDataRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRequestPointerBuilder) CopyToPointer(to HistoricalMarketDepthDataRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}
