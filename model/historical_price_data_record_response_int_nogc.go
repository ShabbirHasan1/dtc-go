package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalPriceDataRecordResponse_IntPointer struct {
	message.FixedPointer
}

type HistoricalPriceDataRecordResponse_IntPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalPriceDataRecordResponse_Int() HistoricalPriceDataRecordResponse_IntPointerBuilder {
	a := HistoricalPriceDataRecordResponse_IntPointerBuilder{message.AllocFixed(56)}
	a.Ptr.SetBytes(0, _HistoricalPriceDataRecordResponse_IntDefault)
	return a
}

func AllocHistoricalPriceDataRecordResponse_IntFrom(b []byte) HistoricalPriceDataRecordResponse_IntPointer {
	return HistoricalPriceDataRecordResponse_IntPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = HistoricalPriceDataRecordResponse_IntSize  (56)
//     Type          = HISTORICAL_PRICE_DATA_RECORD_RESPONSE_INT  (805)
//     RequestID     = 0
//     StartDateTime = 0
//     OpenPrice     = 0
//     HighPrice     = 0
//     LowPrice      = 0
//     LastPrice     = 0
//     Volume        = 0
//     BidVolume     = 0
//     AskVolume     = 0
//     IsFinalRecord = 0
// }
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataRecordResponse_IntDefault)
}

// ToBuilder
func (m HistoricalPriceDataRecordResponse_IntPointer) ToBuilder() HistoricalPriceDataRecordResponse_IntPointerBuilder {
	return HistoricalPriceDataRecordResponse_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalPriceDataRecordResponse_IntPointerBuilder) Finish() HistoricalPriceDataRecordResponse_IntPointer {
	return HistoricalPriceDataRecordResponse_IntPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalPriceDataRecordResponse_IntPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// StartDateTime
func (m HistoricalPriceDataRecordResponse_IntPointer) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 16, 8))
}

// StartDateTime
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 16, 8))
}

// OpenPrice
func (m HistoricalPriceDataRecordResponse_IntPointer) OpenPrice() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// OpenPrice
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) OpenPrice() int32 {
	return message.Int32Fixed(m.Ptr, 20, 16)
}

// HighPrice
func (m HistoricalPriceDataRecordResponse_IntPointer) HighPrice() int32 {
	return message.Int32Fixed(m.Ptr, 24, 20)
}

// HighPrice
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) HighPrice() int32 {
	return message.Int32Fixed(m.Ptr, 24, 20)
}

// LowPrice
func (m HistoricalPriceDataRecordResponse_IntPointer) LowPrice() int32 {
	return message.Int32Fixed(m.Ptr, 28, 24)
}

// LowPrice
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) LowPrice() int32 {
	return message.Int32Fixed(m.Ptr, 28, 24)
}

// LastPrice
func (m HistoricalPriceDataRecordResponse_IntPointer) LastPrice() int32 {
	return message.Int32Fixed(m.Ptr, 32, 28)
}

// LastPrice
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) LastPrice() int32 {
	return message.Int32Fixed(m.Ptr, 32, 28)
}

// Volume
func (m HistoricalPriceDataRecordResponse_IntPointer) Volume() int32 {
	return message.Int32Fixed(m.Ptr, 36, 32)
}

// Volume
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) Volume() int32 {
	return message.Int32Fixed(m.Ptr, 36, 32)
}

// OpenInterest
func (m HistoricalPriceDataRecordResponse_IntPointer) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Ptr, 40, 36)
}

// OpenInterest
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Ptr, 40, 36)
}

// NumTrades
func (m HistoricalPriceDataRecordResponse_IntPointer) NumTrades() uint32 {
	return message.Uint32Fixed(m.Ptr, 40, 36)
}

// NumTrades
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) NumTrades() uint32 {
	return message.Uint32Fixed(m.Ptr, 40, 36)
}

// BidVolume
func (m HistoricalPriceDataRecordResponse_IntPointer) BidVolume() int32 {
	return message.Int32Fixed(m.Ptr, 44, 40)
}

// BidVolume
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) BidVolume() int32 {
	return message.Int32Fixed(m.Ptr, 44, 40)
}

// AskVolume
func (m HistoricalPriceDataRecordResponse_IntPointer) AskVolume() int32 {
	return message.Int32Fixed(m.Ptr, 48, 44)
}

// AskVolume
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) AskVolume() int32 {
	return message.Int32Fixed(m.Ptr, 48, 44)
}

// IsFinalRecord
func (m HistoricalPriceDataRecordResponse_IntPointer) IsFinalRecord() uint8 {
	return message.Uint8Fixed(m.Ptr, 49, 48)
}

// IsFinalRecord
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) IsFinalRecord() uint8 {
	return message.Uint8Fixed(m.Ptr, 49, 48)
}

// SetRequestID
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetStartDateTime
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
}

// SetOpenPrice
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetOpenPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 20, 16, value)
}

// SetHighPrice
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetHighPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 24, 20, value)
}

// SetLowPrice
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetLowPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 28, 24, value)
}

// SetLastPrice
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetLastPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 32, 28, value)
}

// SetVolume
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetVolume(value int32) {
	message.SetInt32Fixed(m.Ptr, 36, 32, value)
}

// SetOpenInterest
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetOpenInterest(value uint32) {
	message.SetUint32Fixed(m.Ptr, 40, 36, value)
}

// SetNumTrades
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetNumTrades(value uint32) {
	message.SetUint32Fixed(m.Ptr, 40, 36, value)
}

// SetBidVolume
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetBidVolume(value int32) {
	message.SetInt32Fixed(m.Ptr, 44, 40, value)
}

// SetAskVolume
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetAskVolume(value int32) {
	message.SetInt32Fixed(m.Ptr, 48, 44, value)
}

// SetIsFinalRecord
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetIsFinalRecord(value uint8) {
	message.SetUint8Fixed(m.Ptr, 49, 48, value)
}

// Copy
func (m HistoricalPriceDataRecordResponse_IntPointer) Copy(to HistoricalPriceDataRecordResponse_IntBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
	to.SetOpenPrice(m.OpenPrice())
	to.SetHighPrice(m.HighPrice())
	to.SetLowPrice(m.LowPrice())
	to.SetLastPrice(m.LastPrice())
	to.SetVolume(m.Volume())
	to.SetOpenInterest(m.OpenInterest())
	to.SetNumTrades(m.NumTrades())
	to.SetBidVolume(m.BidVolume())
	to.SetAskVolume(m.AskVolume())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// Copy
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) Copy(to HistoricalPriceDataRecordResponse_IntBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
	to.SetOpenPrice(m.OpenPrice())
	to.SetHighPrice(m.HighPrice())
	to.SetLowPrice(m.LowPrice())
	to.SetLastPrice(m.LastPrice())
	to.SetVolume(m.Volume())
	to.SetOpenInterest(m.OpenInterest())
	to.SetNumTrades(m.NumTrades())
	to.SetBidVolume(m.BidVolume())
	to.SetAskVolume(m.AskVolume())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalPriceDataRecordResponse_IntPointer) CopyPointer(to HistoricalPriceDataRecordResponse_IntPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
	to.SetOpenPrice(m.OpenPrice())
	to.SetHighPrice(m.HighPrice())
	to.SetLowPrice(m.LowPrice())
	to.SetLastPrice(m.LastPrice())
	to.SetVolume(m.Volume())
	to.SetOpenInterest(m.OpenInterest())
	to.SetNumTrades(m.NumTrades())
	to.SetBidVolume(m.BidVolume())
	to.SetAskVolume(m.AskVolume())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) CopyPointer(to HistoricalPriceDataRecordResponse_IntPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
	to.SetOpenPrice(m.OpenPrice())
	to.SetHighPrice(m.HighPrice())
	to.SetLowPrice(m.LowPrice())
	to.SetLastPrice(m.LastPrice())
	to.SetVolume(m.Volume())
	to.SetOpenInterest(m.OpenInterest())
	to.SetNumTrades(m.NumTrades())
	to.SetBidVolume(m.BidVolume())
	to.SetAskVolume(m.AskVolume())
	to.SetIsFinalRecord(m.IsFinalRecord())
}
