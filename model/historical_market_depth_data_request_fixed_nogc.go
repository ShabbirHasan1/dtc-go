package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalMarketDepthDataRequestFixedPointer struct {
	message.FixedPointer
}

type HistoricalMarketDepthDataRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalMarketDepthDataRequestFixed() HistoricalMarketDepthDataRequestFixedPointerBuilder {
	a := HistoricalMarketDepthDataRequestFixedPointerBuilder{message.AllocFixed(112)}
	a.Ptr.SetBytes(0, _HistoricalMarketDepthDataRequestFixedDefault)
	return a
}

func AllocHistoricalMarketDepthDataRequestFixedFrom(b []byte) HistoricalMarketDepthDataRequestFixedPointer {
	return HistoricalMarketDepthDataRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size               = HistoricalMarketDepthDataRequestFixedSize  (112)
//     Type               = HISTORICAL_MARKET_DEPTH_DATA_REQUEST  (900)
//     RequestID          = 0
//     Symbol             = ""
//     Exchange           = ""
//     StartDateTime      = 0
//     EndDateTime        = 0
//     UseZLibCompression = 0
//     Integer_1          = 0
// }
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalMarketDepthDataRequestFixedDefault)
}

// ToBuilder
func (m HistoricalMarketDepthDataRequestFixedPointer) ToBuilder() HistoricalMarketDepthDataRequestFixedPointerBuilder {
	return HistoricalMarketDepthDataRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalMarketDepthDataRequestFixedPointerBuilder) Finish() HistoricalMarketDepthDataRequestFixedPointer {
	return HistoricalMarketDepthDataRequestFixedPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalMarketDepthDataRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Symbol
func (m HistoricalMarketDepthDataRequestFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Symbol
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Exchange
func (m HistoricalMarketDepthDataRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 88, 72)
}

// Exchange
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 88, 72)
}

// StartDateTime
func (m HistoricalMarketDepthDataRequestFixedPointer) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 96, 88))
}

// StartDateTime
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 96, 88))
}

// EndDateTime
func (m HistoricalMarketDepthDataRequestFixedPointer) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 104, 96))
}

// EndDateTime
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 104, 96))
}

// UseZLibCompression
func (m HistoricalMarketDepthDataRequestFixedPointer) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Ptr, 105, 104)
}

// UseZLibCompression
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Ptr, 105, 104)
}

// Integer_1
func (m HistoricalMarketDepthDataRequestFixedPointer) Integer_1() uint8 {
	return message.Uint8Fixed(m.Ptr, 106, 105)
}

// Integer_1
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) Integer_1() uint8 {
	return message.Uint8Fixed(m.Ptr, 106, 105)
}

// SetRequestID
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetSymbol
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 72, 8, value)
}

// SetExchange
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 88, 72, value)
}

// SetStartDateTime
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 96, 88, int64(value))
}

// SetEndDateTime
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetEndDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 104, 96, int64(value))
}

// SetUseZLibCompression
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8Fixed(m.Ptr, 105, 104, value)
}

// SetInteger_1
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetInteger_1(value uint8) {
	message.SetUint8Fixed(m.Ptr, 106, 105, value)
}

// Copy
func (m HistoricalMarketDepthDataRequestFixedPointer) Copy(to HistoricalMarketDepthDataRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// Copy
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) Copy(to HistoricalMarketDepthDataRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyPointer
func (m HistoricalMarketDepthDataRequestFixedPointer) CopyPointer(to HistoricalMarketDepthDataRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyPointer
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) CopyPointer(to HistoricalMarketDepthDataRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyTo
func (m HistoricalMarketDepthDataRequestFixedPointer) CopyTo(to HistoricalMarketDepthDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyTo
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) CopyTo(to HistoricalMarketDepthDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRequestFixedPointer) CopyToPointer(to HistoricalMarketDepthDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) CopyToPointer(to HistoricalMarketDepthDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}
