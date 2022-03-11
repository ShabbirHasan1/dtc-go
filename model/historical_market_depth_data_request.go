package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalMarketDepthDataRequestSize = 48

const HistoricalMarketDepthDataRequestFixedSize = 112

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

type HistoricalMarketDepthDataRequest struct {
	message.VLS
}

type HistoricalMarketDepthDataRequestBuilder struct {
	message.VLSBuilder
}

type HistoricalMarketDepthDataRequestFixed struct {
	message.Fixed
}

type HistoricalMarketDepthDataRequestFixedBuilder struct {
	message.Fixed
}

type HistoricalMarketDepthDataRequestPointer struct {
	message.VLSPointer
}

type HistoricalMarketDepthDataRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

type HistoricalMarketDepthDataRequestFixedPointer struct {
	message.FixedPointer
}

type HistoricalMarketDepthDataRequestFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocHistoricalMarketDepthDataRequest() HistoricalMarketDepthDataRequestPointerBuilder {
	a := HistoricalMarketDepthDataRequestPointerBuilder{message.AllocVLSBuilder(48)}
	a.Ptr.SetBytes(0, _HistoricalMarketDepthDataRequestDefault)
	return a
}

func AllocHistoricalMarketDepthDataRequestFrom(b []byte) HistoricalMarketDepthDataRequestPointer {
	return HistoricalMarketDepthDataRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m HistoricalMarketDepthDataRequest) ToBuilder() HistoricalMarketDepthDataRequestBuilder {
	return HistoricalMarketDepthDataRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalMarketDepthDataRequestFixed) ToBuilder() HistoricalMarketDepthDataRequestFixedBuilder {
	return HistoricalMarketDepthDataRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalMarketDepthDataRequestPointer) ToBuilder() HistoricalMarketDepthDataRequestPointerBuilder {
	return HistoricalMarketDepthDataRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalMarketDepthDataRequestFixedPointer) ToBuilder() HistoricalMarketDepthDataRequestFixedPointerBuilder {
	return HistoricalMarketDepthDataRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalMarketDepthDataRequestBuilder) Finish() HistoricalMarketDepthDataRequest {
	return HistoricalMarketDepthDataRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalMarketDepthDataRequestFixedBuilder) Finish() HistoricalMarketDepthDataRequestFixed {
	return HistoricalMarketDepthDataRequestFixed{m.Fixed}
}

// Finish
func (m *HistoricalMarketDepthDataRequestPointerBuilder) Finish() HistoricalMarketDepthDataRequestPointer {
	return HistoricalMarketDepthDataRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalMarketDepthDataRequestFixedPointerBuilder) Finish() HistoricalMarketDepthDataRequestFixedPointer {
	return HistoricalMarketDepthDataRequestFixedPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalMarketDepthDataRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID
func (m HistoricalMarketDepthDataRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
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
func (m HistoricalMarketDepthDataRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Symbol
func (m HistoricalMarketDepthDataRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
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
func (m HistoricalMarketDepthDataRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange
func (m HistoricalMarketDepthDataRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
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
func (m HistoricalMarketDepthDataRequest) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Unsafe(), 32, 24))
}

// StartDateTime
func (m HistoricalMarketDepthDataRequestBuilder) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Unsafe(), 32, 24))
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
func (m HistoricalMarketDepthDataRequest) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Unsafe(), 40, 32))
}

// EndDateTime
func (m HistoricalMarketDepthDataRequestBuilder) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Unsafe(), 40, 32))
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
func (m HistoricalMarketDepthDataRequest) UseZLibCompression() uint8 {
	return message.Uint8VLS(m.Unsafe(), 41, 40)
}

// UseZLibCompression
func (m HistoricalMarketDepthDataRequestBuilder) UseZLibCompression() uint8 {
	return message.Uint8VLS(m.Unsafe(), 41, 40)
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
func (m HistoricalMarketDepthDataRequest) Integer_1() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// Integer_1
func (m HistoricalMarketDepthDataRequestBuilder) Integer_1() uint8 {
	return message.Uint8VLS(m.Unsafe(), 42, 41)
}

// Integer_1
func (m HistoricalMarketDepthDataRequestPointer) Integer_1() uint8 {
	return message.Uint8VLS(m.Ptr, 42, 41)
}

// Integer_1
func (m HistoricalMarketDepthDataRequestPointerBuilder) Integer_1() uint8 {
	return message.Uint8VLS(m.Ptr, 42, 41)
}

// RequestID
func (m HistoricalMarketDepthDataRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalMarketDepthDataRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m HistoricalMarketDepthDataRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol
func (m HistoricalMarketDepthDataRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
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
func (m HistoricalMarketDepthDataRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// Exchange
func (m HistoricalMarketDepthDataRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
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
func (m HistoricalMarketDepthDataRequestFixed) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 96, 88))
}

// StartDateTime
func (m HistoricalMarketDepthDataRequestFixedBuilder) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 96, 88))
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
func (m HistoricalMarketDepthDataRequestFixed) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 104, 96))
}

// EndDateTime
func (m HistoricalMarketDepthDataRequestFixedBuilder) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 104, 96))
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
func (m HistoricalMarketDepthDataRequestFixed) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 105, 104)
}

// UseZLibCompression
func (m HistoricalMarketDepthDataRequestFixedBuilder) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 105, 104)
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
func (m HistoricalMarketDepthDataRequestFixed) Integer_1() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 106, 105)
}

// Integer_1
func (m HistoricalMarketDepthDataRequestFixedBuilder) Integer_1() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 106, 105)
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
func (m HistoricalMarketDepthDataRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID
func (m HistoricalMarketDepthDataRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetSymbol
func (m *HistoricalMarketDepthDataRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetSymbol
func (m *HistoricalMarketDepthDataRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetExchange
func (m *HistoricalMarketDepthDataRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetExchange
func (m *HistoricalMarketDepthDataRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetStartDateTime
func (m HistoricalMarketDepthDataRequestBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 32, 24, int64(value))
}

// SetStartDateTime
func (m HistoricalMarketDepthDataRequestPointerBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64VLS(m.Ptr, 32, 24, int64(value))
}

// SetEndDateTime
func (m HistoricalMarketDepthDataRequestBuilder) SetEndDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 40, 32, int64(value))
}

// SetEndDateTime
func (m HistoricalMarketDepthDataRequestPointerBuilder) SetEndDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64VLS(m.Ptr, 40, 32, int64(value))
}

// SetUseZLibCompression
func (m HistoricalMarketDepthDataRequestBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 41, 40, value)
}

// SetUseZLibCompression
func (m HistoricalMarketDepthDataRequestPointerBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8VLS(m.Ptr, 41, 40, value)
}

// SetInteger_1
func (m HistoricalMarketDepthDataRequestBuilder) SetInteger_1(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 42, 41, value)
}

// SetInteger_1
func (m HistoricalMarketDepthDataRequestPointerBuilder) SetInteger_1(value uint8) {
	message.SetUint8VLS(m.Ptr, 42, 41, value)
}

// SetRequestID
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetSymbol
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetSymbol
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 72, 8, value)
}

// SetExchange
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 88, 72, value)
}

// SetExchange
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 88, 72, value)
}

// SetStartDateTime
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 96, 88, int64(value))
}

// SetStartDateTime
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 96, 88, int64(value))
}

// SetEndDateTime
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetEndDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 104, 96, int64(value))
}

// SetEndDateTime
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetEndDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 104, 96, int64(value))
}

// SetUseZLibCompression
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 105, 104, value)
}

// SetUseZLibCompression
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8Fixed(m.Ptr, 105, 104, value)
}

// SetInteger_1
func (m HistoricalMarketDepthDataRequestFixedBuilder) SetInteger_1(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 106, 105, value)
}

// SetInteger_1
func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) SetInteger_1(value uint8) {
	message.SetUint8Fixed(m.Ptr, 106, 105, value)
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
