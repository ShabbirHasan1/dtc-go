package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _HistoricalMarketDepthDataRequestFixedDefault = []byte{112, 0, 132, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalMarketDepthDataRequestFixedSize = 112

type HistoricalMarketDepthDataRequestFixed struct {
	message.Fixed
}

type HistoricalMarketDepthDataRequestFixedBuilder struct {
	message.Fixed
}

func NewHistoricalMarketDepthDataRequestFixedFrom(b []byte) HistoricalMarketDepthDataRequestFixed {
	return HistoricalMarketDepthDataRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalMarketDepthDataRequestFixed(b []byte) HistoricalMarketDepthDataRequestFixed {
	return HistoricalMarketDepthDataRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalMarketDepthDataRequestFixed() HistoricalMarketDepthDataRequestFixedBuilder {
	a := HistoricalMarketDepthDataRequestFixedBuilder{message.NewFixed(112)}
	a.Unsafe().SetBytes(0, _HistoricalMarketDepthDataRequestFixedDefault)
	return a
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
func (m HistoricalMarketDepthDataRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalMarketDepthDataRequestFixedDefault)
}

// ToBuilder
func (m HistoricalMarketDepthDataRequestFixed) ToBuilder() HistoricalMarketDepthDataRequestFixedBuilder {
	return HistoricalMarketDepthDataRequestFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalMarketDepthDataRequestFixedBuilder) Finish() HistoricalMarketDepthDataRequestFixed {
	return HistoricalMarketDepthDataRequestFixed{m.Fixed}
}

// RequestID
func (m HistoricalMarketDepthDataRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalMarketDepthDataRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// Symbol
func (m HistoricalMarketDepthDataRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol
func (m HistoricalMarketDepthDataRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Exchange
func (m HistoricalMarketDepthDataRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// Exchange
func (m HistoricalMarketDepthDataRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// StartDateTime
func (m HistoricalMarketDepthDataRequestFixed) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 96, 88))
}

// StartDateTime
func (m HistoricalMarketDepthDataRequestFixedBuilder) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 96, 88))
}

// EndDateTime
func (m HistoricalMarketDepthDataRequestFixed) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 104, 96))
}

// EndDateTime
func (m HistoricalMarketDepthDataRequestFixedBuilder) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 104, 96))
}

// UseZLibCompression
func (m HistoricalMarketDepthDataRequestFixed) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 105, 104)
}

// UseZLibCompression
func (m HistoricalMarketDepthDataRequestFixedBuilder) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 105, 104)
}

// Integer_1
func (m HistoricalMarketDepthDataRequestFixed) Integer_1() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 106, 105)
}

// Integer_1
func (m HistoricalMarketDepthDataRequestFixedBuilder) Integer_1() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 106, 105)
}

// SetRequestID
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbol
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetExchange
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 88, 72, value)
}

// SetStartDateTime
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 96, 88, int64(value))
}

// SetEndDateTime
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetEndDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 104, 96, int64(value))
}

// SetUseZLibCompression
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 105, 104, value)
}

// SetInteger_1
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetInteger_1(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 106, 105, value)
}

// Copy
func (m HistoricalMarketDepthDataRequestFixed) Copy(to HistoricalMarketDepthDataRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// Copy
func (m HistoricalMarketDepthDataRequestFixedBuilder) Copy(to HistoricalMarketDepthDataRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyPointer
func (m HistoricalMarketDepthDataRequestFixed) CopyPointer(to HistoricalMarketDepthDataRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyPointer
func (m HistoricalMarketDepthDataRequestFixedBuilder) CopyPointer(to HistoricalMarketDepthDataRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRequestFixed) CopyToPointer(to HistoricalMarketDepthDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyToPointer
func (m HistoricalMarketDepthDataRequestFixedBuilder) CopyToPointer(to HistoricalMarketDepthDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyTo
func (m HistoricalMarketDepthDataRequestFixed) CopyTo(to HistoricalMarketDepthDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyTo
func (m HistoricalMarketDepthDataRequestFixedBuilder) CopyTo(to HistoricalMarketDepthDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}
