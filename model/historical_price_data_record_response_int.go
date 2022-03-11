package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalPriceDataRecordResponse_IntSize = 56

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
//     IsFinalRecord = false
// }
var _HistoricalPriceDataRecordResponse_IntDefault = []byte{56, 0, 37, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HistoricalPriceDataRecordResponse_Int struct {
	message.Fixed
}

type HistoricalPriceDataRecordResponse_IntBuilder struct {
	message.Fixed
}

type HistoricalPriceDataRecordResponse_IntPointer struct {
	message.FixedPointer
}

type HistoricalPriceDataRecordResponse_IntPointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalPriceDataRecordResponse_IntFrom(b []byte) HistoricalPriceDataRecordResponse_Int {
	return HistoricalPriceDataRecordResponse_Int{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalPriceDataRecordResponse_Int(b []byte) HistoricalPriceDataRecordResponse_Int {
	return HistoricalPriceDataRecordResponse_Int{Fixed: message.WrapFixed(b)}
}

func NewHistoricalPriceDataRecordResponse_Int() HistoricalPriceDataRecordResponse_IntBuilder {
	a := HistoricalPriceDataRecordResponse_IntBuilder{message.NewFixed(56)}
	a.Unsafe().SetBytes(0, _HistoricalPriceDataRecordResponse_IntDefault)
	return a
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
//     IsFinalRecord = false
// }
func (m HistoricalPriceDataRecordResponse_IntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataRecordResponse_IntDefault)
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
//     IsFinalRecord = false
// }
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataRecordResponse_IntDefault)
}

// ToBuilder
func (m HistoricalPriceDataRecordResponse_Int) ToBuilder() HistoricalPriceDataRecordResponse_IntBuilder {
	return HistoricalPriceDataRecordResponse_IntBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalPriceDataRecordResponse_IntPointer) ToBuilder() HistoricalPriceDataRecordResponse_IntPointerBuilder {
	return HistoricalPriceDataRecordResponse_IntPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalPriceDataRecordResponse_IntBuilder) Finish() HistoricalPriceDataRecordResponse_Int {
	return HistoricalPriceDataRecordResponse_Int{m.Fixed}
}

// Finish
func (m *HistoricalPriceDataRecordResponse_IntPointerBuilder) Finish() HistoricalPriceDataRecordResponse_IntPointer {
	return HistoricalPriceDataRecordResponse_IntPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalPriceDataRecordResponse_Int) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalPriceDataRecordResponse_IntBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m HistoricalPriceDataRecordResponse_Int) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// StartDateTime
func (m HistoricalPriceDataRecordResponse_IntBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 16, 8))
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
func (m HistoricalPriceDataRecordResponse_Int) OpenPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// OpenPrice
func (m HistoricalPriceDataRecordResponse_IntBuilder) OpenPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
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
func (m HistoricalPriceDataRecordResponse_Int) HighPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 24, 20)
}

// HighPrice
func (m HistoricalPriceDataRecordResponse_IntBuilder) HighPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 24, 20)
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
func (m HistoricalPriceDataRecordResponse_Int) LowPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 28, 24)
}

// LowPrice
func (m HistoricalPriceDataRecordResponse_IntBuilder) LowPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 28, 24)
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
func (m HistoricalPriceDataRecordResponse_Int) LastPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 32, 28)
}

// LastPrice
func (m HistoricalPriceDataRecordResponse_IntBuilder) LastPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 32, 28)
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
func (m HistoricalPriceDataRecordResponse_Int) Volume() int32 {
	return message.Int32Fixed(m.Unsafe(), 36, 32)
}

// Volume
func (m HistoricalPriceDataRecordResponse_IntBuilder) Volume() int32 {
	return message.Int32Fixed(m.Unsafe(), 36, 32)
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
func (m HistoricalPriceDataRecordResponse_Int) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 40, 36)
}

// OpenInterest
func (m HistoricalPriceDataRecordResponse_IntBuilder) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 40, 36)
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
func (m HistoricalPriceDataRecordResponse_Int) NumTrades() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 40, 36)
}

// NumTrades
func (m HistoricalPriceDataRecordResponse_IntBuilder) NumTrades() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 40, 36)
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
func (m HistoricalPriceDataRecordResponse_Int) BidVolume() int32 {
	return message.Int32Fixed(m.Unsafe(), 44, 40)
}

// BidVolume
func (m HistoricalPriceDataRecordResponse_IntBuilder) BidVolume() int32 {
	return message.Int32Fixed(m.Unsafe(), 44, 40)
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
func (m HistoricalPriceDataRecordResponse_Int) AskVolume() int32 {
	return message.Int32Fixed(m.Unsafe(), 48, 44)
}

// AskVolume
func (m HistoricalPriceDataRecordResponse_IntBuilder) AskVolume() int32 {
	return message.Int32Fixed(m.Unsafe(), 48, 44)
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
func (m HistoricalPriceDataRecordResponse_Int) IsFinalRecord() bool {
	return message.BoolFixed(m.Unsafe(), 49, 48)
}

// IsFinalRecord
func (m HistoricalPriceDataRecordResponse_IntBuilder) IsFinalRecord() bool {
	return message.BoolFixed(m.Unsafe(), 49, 48)
}

// IsFinalRecord
func (m HistoricalPriceDataRecordResponse_IntPointer) IsFinalRecord() bool {
	return message.BoolFixed(m.Ptr, 49, 48)
}

// IsFinalRecord
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) IsFinalRecord() bool {
	return message.BoolFixed(m.Ptr, 49, 48)
}

// SetRequestID
func (m HistoricalPriceDataRecordResponse_IntBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetStartDateTime
func (m HistoricalPriceDataRecordResponse_IntBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, int64(value))
}

// SetStartDateTime
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
}

// SetOpenPrice
func (m HistoricalPriceDataRecordResponse_IntBuilder) SetOpenPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 20, 16, value)
}

// SetOpenPrice
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetOpenPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 20, 16, value)
}

// SetHighPrice
func (m HistoricalPriceDataRecordResponse_IntBuilder) SetHighPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 24, 20, value)
}

// SetHighPrice
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetHighPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 24, 20, value)
}

// SetLowPrice
func (m HistoricalPriceDataRecordResponse_IntBuilder) SetLowPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 28, 24, value)
}

// SetLowPrice
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetLowPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 28, 24, value)
}

// SetLastPrice
func (m HistoricalPriceDataRecordResponse_IntBuilder) SetLastPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 32, 28, value)
}

// SetLastPrice
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetLastPrice(value int32) {
	message.SetInt32Fixed(m.Ptr, 32, 28, value)
}

// SetVolume
func (m HistoricalPriceDataRecordResponse_IntBuilder) SetVolume(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 36, 32, value)
}

// SetVolume
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetVolume(value int32) {
	message.SetInt32Fixed(m.Ptr, 36, 32, value)
}

// SetOpenInterest
func (m HistoricalPriceDataRecordResponse_IntBuilder) SetOpenInterest(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 40, 36, value)
}

// SetOpenInterest
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetOpenInterest(value uint32) {
	message.SetUint32Fixed(m.Ptr, 40, 36, value)
}

// SetNumTrades
func (m HistoricalPriceDataRecordResponse_IntBuilder) SetNumTrades(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 40, 36, value)
}

// SetNumTrades
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetNumTrades(value uint32) {
	message.SetUint32Fixed(m.Ptr, 40, 36, value)
}

// SetBidVolume
func (m HistoricalPriceDataRecordResponse_IntBuilder) SetBidVolume(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 44, 40, value)
}

// SetBidVolume
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetBidVolume(value int32) {
	message.SetInt32Fixed(m.Ptr, 44, 40, value)
}

// SetAskVolume
func (m HistoricalPriceDataRecordResponse_IntBuilder) SetAskVolume(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 48, 44, value)
}

// SetAskVolume
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetAskVolume(value int32) {
	message.SetInt32Fixed(m.Ptr, 48, 44, value)
}

// SetIsFinalRecord
func (m HistoricalPriceDataRecordResponse_IntBuilder) SetIsFinalRecord(value bool) {
	message.SetBoolFixed(m.Unsafe(), 49, 48, value)
}

// SetIsFinalRecord
func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) SetIsFinalRecord(value bool) {
	message.SetBoolFixed(m.Ptr, 49, 48, value)
}

// Copy
func (m HistoricalPriceDataRecordResponse_Int) Copy(to HistoricalPriceDataRecordResponse_IntBuilder) {
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
func (m HistoricalPriceDataRecordResponse_IntBuilder) Copy(to HistoricalPriceDataRecordResponse_IntBuilder) {
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
func (m HistoricalPriceDataRecordResponse_Int) CopyPointer(to HistoricalPriceDataRecordResponse_IntPointerBuilder) {
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
func (m HistoricalPriceDataRecordResponse_IntBuilder) CopyPointer(to HistoricalPriceDataRecordResponse_IntPointerBuilder) {
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