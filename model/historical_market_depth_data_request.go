package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _HistoricalMarketDepthDataRequestDefault = []byte{48, 0, 132, 3, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalMarketDepthDataRequestSize = 48

type HistoricalMarketDepthDataRequest struct {
	message.VLS
}

type HistoricalMarketDepthDataRequestBuilder struct {
	message.VLSBuilder
}

func NewHistoricalMarketDepthDataRequestFrom(b []byte) HistoricalMarketDepthDataRequest {
	return HistoricalMarketDepthDataRequest{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalMarketDepthDataRequest(b []byte) HistoricalMarketDepthDataRequest {
	return HistoricalMarketDepthDataRequest{VLS: message.WrapVLS(b)}
}

func NewHistoricalMarketDepthDataRequest() HistoricalMarketDepthDataRequestBuilder {
	a := HistoricalMarketDepthDataRequestBuilder{message.NewVLSBuilder(48)}
	a.Unsafe().SetBytes(0, _HistoricalMarketDepthDataRequestDefault)
	return a
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
func (m HistoricalMarketDepthDataRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalMarketDepthDataRequestDefault)
}

// ToBuilder
func (m HistoricalMarketDepthDataRequest) ToBuilder() HistoricalMarketDepthDataRequestBuilder {
	return HistoricalMarketDepthDataRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m HistoricalMarketDepthDataRequestBuilder) Finish() HistoricalMarketDepthDataRequest {
	return HistoricalMarketDepthDataRequest{m.VLSBuilder.Finish()}
}

// RequestID
func (m HistoricalMarketDepthDataRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m HistoricalMarketDepthDataRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// Symbol
func (m HistoricalMarketDepthDataRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Symbol
func (m HistoricalMarketDepthDataRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Exchange
func (m HistoricalMarketDepthDataRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange
func (m HistoricalMarketDepthDataRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// StartDateTime
func (m HistoricalMarketDepthDataRequest) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Unsafe(), 32, 24))
}

// StartDateTime
func (m HistoricalMarketDepthDataRequestBuilder) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Unsafe(), 32, 24))
}

// EndDateTime
func (m HistoricalMarketDepthDataRequest) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Unsafe(), 40, 32))
}

// EndDateTime
func (m HistoricalMarketDepthDataRequestBuilder) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Unsafe(), 40, 32))
}

// UseZLibCompression
func (m HistoricalMarketDepthDataRequest) UseZLibCompression() uint8 {
	return message.Uint8VLS(m.Unsafe(), 41, 40)
}

// UseZLibCompression
func (m HistoricalMarketDepthDataRequestBuilder) UseZLibCompression() uint8 {
	return message.Uint8VLS(m.Unsafe(), 41, 40)
}

// Integer_1
func (m HistoricalMarketDepthDataRequest) Integer_1() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// Integer_1
func (m HistoricalMarketDepthDataRequestBuilder) Integer_1() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// SetRequestID
func (m HistoricalMarketDepthDataRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetSymbol
func (m *HistoricalMarketDepthDataRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetExchange
func (m *HistoricalMarketDepthDataRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetStartDateTime
func (m HistoricalMarketDepthDataRequestBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 32, 24, int64(value))
}

// SetEndDateTime
func (m HistoricalMarketDepthDataRequestBuilder) SetEndDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 40, 32, int64(value))
}

// SetUseZLibCompression
func (m HistoricalMarketDepthDataRequestBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 41, 40, value)
}

// SetInteger_1
func (m HistoricalMarketDepthDataRequestBuilder) SetInteger_1(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 42, 41, value)
}

// Copy
func (m HistoricalMarketDepthDataRequest) Copy(to HistoricalMarketDepthDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// Copy
func (m HistoricalMarketDepthDataRequestBuilder) Copy(to HistoricalMarketDepthDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyPointer
func (m HistoricalMarketDepthDataRequest) CopyPointer(to HistoricalMarketDepthDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyPointer
func (m HistoricalMarketDepthDataRequestBuilder) CopyPointer(to HistoricalMarketDepthDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRequest) CopyToPointer(to HistoricalMarketDepthDataRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRequestBuilder) CopyToPointer(to HistoricalMarketDepthDataRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyTo
func (m HistoricalMarketDepthDataRequest) CopyTo(to HistoricalMarketDepthDataRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyTo
func (m HistoricalMarketDepthDataRequestBuilder) CopyTo(to HistoricalMarketDepthDataRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}
