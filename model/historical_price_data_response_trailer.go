package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalPriceDataResponseTrailerSize = 16

// {
//     Size                    = HistoricalPriceDataResponseTrailerSize  (16)
//     Type                    = HISTORICAL_PRICE_DATA_RESPONSE_TRAILER  (807)
//     RequestID               = 0
//     FinalRecordLastDateTime = 0
// }
var _HistoricalPriceDataResponseTrailerDefault = []byte{16, 0, 39, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HistoricalPriceDataResponseTrailer struct {
	message.Fixed
}

type HistoricalPriceDataResponseTrailerBuilder struct {
	message.Fixed
}

type HistoricalPriceDataResponseTrailerPointer struct {
	message.FixedPointer
}

type HistoricalPriceDataResponseTrailerPointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalPriceDataResponseTrailerFrom(b []byte) HistoricalPriceDataResponseTrailer {
	return HistoricalPriceDataResponseTrailer{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalPriceDataResponseTrailer(b []byte) HistoricalPriceDataResponseTrailer {
	return HistoricalPriceDataResponseTrailer{Fixed: message.WrapFixed(b)}
}

func NewHistoricalPriceDataResponseTrailer() HistoricalPriceDataResponseTrailerBuilder {
	a := HistoricalPriceDataResponseTrailerBuilder{message.NewFixed(16)}
	a.Unsafe().SetBytes(0, _HistoricalPriceDataResponseTrailerDefault)
	return a
}

func AllocHistoricalPriceDataResponseTrailer() HistoricalPriceDataResponseTrailerPointerBuilder {
	a := HistoricalPriceDataResponseTrailerPointerBuilder{message.AllocFixed(16)}
	a.Ptr.SetBytes(0, _HistoricalPriceDataResponseTrailerDefault)
	return a
}

func AllocHistoricalPriceDataResponseTrailerFrom(b []byte) HistoricalPriceDataResponseTrailerPointer {
	return HistoricalPriceDataResponseTrailerPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                    = HistoricalPriceDataResponseTrailerSize  (16)
//     Type                    = HISTORICAL_PRICE_DATA_RESPONSE_TRAILER  (807)
//     RequestID               = 0
//     FinalRecordLastDateTime = 0
// }
func (m HistoricalPriceDataResponseTrailerBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataResponseTrailerDefault)
}

// Clear
// {
//     Size                    = HistoricalPriceDataResponseTrailerSize  (16)
//     Type                    = HISTORICAL_PRICE_DATA_RESPONSE_TRAILER  (807)
//     RequestID               = 0
//     FinalRecordLastDateTime = 0
// }
func (m HistoricalPriceDataResponseTrailerPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataResponseTrailerDefault)
}

// ToBuilder
func (m HistoricalPriceDataResponseTrailer) ToBuilder() HistoricalPriceDataResponseTrailerBuilder {
	return HistoricalPriceDataResponseTrailerBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalPriceDataResponseTrailerPointer) ToBuilder() HistoricalPriceDataResponseTrailerPointerBuilder {
	return HistoricalPriceDataResponseTrailerPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalPriceDataResponseTrailerBuilder) Finish() HistoricalPriceDataResponseTrailer {
	return HistoricalPriceDataResponseTrailer{m.Fixed}
}

// Finish
func (m *HistoricalPriceDataResponseTrailerPointerBuilder) Finish() HistoricalPriceDataResponseTrailerPointer {
	return HistoricalPriceDataResponseTrailerPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalPriceDataResponseTrailer) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalPriceDataResponseTrailerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalPriceDataResponseTrailerPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalPriceDataResponseTrailerPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// FinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailer) FinalRecordLastDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// FinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailerBuilder) FinalRecordLastDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// FinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailerPointer) FinalRecordLastDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 16, 8))
}

// FinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailerPointerBuilder) FinalRecordLastDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 16, 8))
}

// SetRequestID
func (m HistoricalPriceDataResponseTrailerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m HistoricalPriceDataResponseTrailerPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetFinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailerBuilder) SetFinalRecordLastDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, int64(value))
}

// SetFinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailerPointerBuilder) SetFinalRecordLastDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
}

// Copy
func (m HistoricalPriceDataResponseTrailer) Copy(to HistoricalPriceDataResponseTrailerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetFinalRecordLastDateTime(m.FinalRecordLastDateTime())
}

// Copy
func (m HistoricalPriceDataResponseTrailerBuilder) Copy(to HistoricalPriceDataResponseTrailerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetFinalRecordLastDateTime(m.FinalRecordLastDateTime())
}

// CopyPointer
func (m HistoricalPriceDataResponseTrailer) CopyPointer(to HistoricalPriceDataResponseTrailerPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetFinalRecordLastDateTime(m.FinalRecordLastDateTime())
}

// CopyPointer
func (m HistoricalPriceDataResponseTrailerBuilder) CopyPointer(to HistoricalPriceDataResponseTrailerPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetFinalRecordLastDateTime(m.FinalRecordLastDateTime())
}

// Copy
func (m HistoricalPriceDataResponseTrailerPointer) Copy(to HistoricalPriceDataResponseTrailerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetFinalRecordLastDateTime(m.FinalRecordLastDateTime())
}

// Copy
func (m HistoricalPriceDataResponseTrailerPointerBuilder) Copy(to HistoricalPriceDataResponseTrailerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetFinalRecordLastDateTime(m.FinalRecordLastDateTime())
}

// CopyPointer
func (m HistoricalPriceDataResponseTrailerPointer) CopyPointer(to HistoricalPriceDataResponseTrailerPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetFinalRecordLastDateTime(m.FinalRecordLastDateTime())
}

// CopyPointer
func (m HistoricalPriceDataResponseTrailerPointerBuilder) CopyPointer(to HistoricalPriceDataResponseTrailerPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetFinalRecordLastDateTime(m.FinalRecordLastDateTime())
}
