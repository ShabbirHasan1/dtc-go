package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                    = HistoricalPriceDataResponseTrailerSize  (16)
//     Type                    = HISTORICAL_PRICE_DATA_RESPONSE_TRAILER  (807)
//     RequestID               = 0
//     FinalRecordLastDateTime = 0
// }
var _HistoricalPriceDataResponseTrailerDefault = []byte{16, 0, 39, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalPriceDataResponseTrailerSize = 16

type HistoricalPriceDataResponseTrailer struct {
	message.Fixed
}

type HistoricalPriceDataResponseTrailerBuilder struct {
	message.Fixed
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

// ToBuilder
func (m HistoricalPriceDataResponseTrailer) ToBuilder() HistoricalPriceDataResponseTrailerBuilder {
	return HistoricalPriceDataResponseTrailerBuilder{m.Fixed}
}

// Finish
func (m HistoricalPriceDataResponseTrailerBuilder) Finish() HistoricalPriceDataResponseTrailer {
	return HistoricalPriceDataResponseTrailer{m.Fixed}
}

// RequestID
func (m HistoricalPriceDataResponseTrailer) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalPriceDataResponseTrailerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// FinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailer) FinalRecordLastDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// FinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailerBuilder) FinalRecordLastDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// SetRequestID
func (m HistoricalPriceDataResponseTrailerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetFinalRecordLastDateTime
func (m HistoricalPriceDataResponseTrailerBuilder) SetFinalRecordLastDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, int64(value))
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
