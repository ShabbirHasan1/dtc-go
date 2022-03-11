package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = HistoricalPriceDataTickRecordResponse_IntSize  (32)
//     Type          = HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE_INT  (806)
//     RequestID     = 0
//     DateTime      = 0
//     Price         = 0
//     Volume        = 0
//     AtBidOrAsk    = 0
//     IsFinalRecord = 0
// }
var _HistoricalPriceDataTickRecordResponse_IntDefault = []byte{32, 0, 38, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalPriceDataTickRecordResponse_IntSize = 32

type HistoricalPriceDataTickRecordResponse_Int struct {
	message.Fixed
}

type HistoricalPriceDataTickRecordResponse_IntBuilder struct {
	message.Fixed
}

func NewHistoricalPriceDataTickRecordResponse_IntFrom(b []byte) HistoricalPriceDataTickRecordResponse_Int {
	return HistoricalPriceDataTickRecordResponse_Int{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalPriceDataTickRecordResponse_Int(b []byte) HistoricalPriceDataTickRecordResponse_Int {
	return HistoricalPriceDataTickRecordResponse_Int{Fixed: message.WrapFixed(b)}
}

func NewHistoricalPriceDataTickRecordResponse_Int() HistoricalPriceDataTickRecordResponse_IntBuilder {
	a := HistoricalPriceDataTickRecordResponse_IntBuilder{message.NewFixed(32)}
	a.Unsafe().SetBytes(0, _HistoricalPriceDataTickRecordResponse_IntDefault)
	return a
}

// Clear
// {
//     Size          = HistoricalPriceDataTickRecordResponse_IntSize  (32)
//     Type          = HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE_INT  (806)
//     RequestID     = 0
//     DateTime      = 0
//     Price         = 0
//     Volume        = 0
//     AtBidOrAsk    = 0
//     IsFinalRecord = 0
// }
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataTickRecordResponse_IntDefault)
}

// ToBuilder
func (m HistoricalPriceDataTickRecordResponse_Int) ToBuilder() HistoricalPriceDataTickRecordResponse_IntBuilder {
	return HistoricalPriceDataTickRecordResponse_IntBuilder{m.Fixed}
}

// Finish
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) Finish() HistoricalPriceDataTickRecordResponse_Int {
	return HistoricalPriceDataTickRecordResponse_Int{m.Fixed}
}

// RequestID
func (m HistoricalPriceDataTickRecordResponse_Int) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// DateTime
func (m HistoricalPriceDataTickRecordResponse_Int) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 16, 8))
}

// DateTime
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 16, 8))
}

// Price
func (m HistoricalPriceDataTickRecordResponse_Int) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// Price
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) Price() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// Volume
func (m HistoricalPriceDataTickRecordResponse_Int) Volume() int32 {
	return message.Int32Fixed(m.Unsafe(), 24, 20)
}

// Volume
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) Volume() int32 {
	return message.Int32Fixed(m.Unsafe(), 24, 20)
}

// AtBidOrAsk
func (m HistoricalPriceDataTickRecordResponse_Int) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 26, 24))
}

// AtBidOrAsk
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 26, 24))
}

// IsFinalRecord
func (m HistoricalPriceDataTickRecordResponse_Int) IsFinalRecord() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 27, 26)
}

// IsFinalRecord
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) IsFinalRecord() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 27, 26)
}

// SetRequestID
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetDateTime
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, float64(value))
}

// SetPrice
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) SetPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 20, 16, value)
}

// SetVolume
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) SetVolume(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 24, 20, value)
}

// SetAtBidOrAsk
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) SetAtBidOrAsk(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Unsafe(), 26, 24, uint16(value))
}

// SetIsFinalRecord
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) SetIsFinalRecord(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 27, 26, value)
}

// Copy
func (m HistoricalPriceDataTickRecordResponse_Int) Copy(to HistoricalPriceDataTickRecordResponse_IntBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// Copy
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) Copy(to HistoricalPriceDataTickRecordResponse_IntBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalPriceDataTickRecordResponse_Int) CopyPointer(to HistoricalPriceDataTickRecordResponse_IntPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalPriceDataTickRecordResponse_IntBuilder) CopyPointer(to HistoricalPriceDataTickRecordResponse_IntPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetIsFinalRecord(m.IsFinalRecord())
}
