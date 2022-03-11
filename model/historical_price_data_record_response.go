package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalPriceDataRecordResponseSize = 88

// {
//     Size          = HistoricalPriceDataRecordResponseSize  (88)
//     Type          = HISTORICAL_PRICE_DATA_RECORD_RESPONSE  (803)
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
var _HistoricalPriceDataRecordResponseDefault = []byte{88, 0, 35, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// HistoricalPriceDataRecordResponse The HistoricalPriceDataTickRecordResponse message is used when the RecordInterval
// field in a historical data request message is set to a value greater than
// INTERVAL_TICK. For example, if the RecordInterval is INTERVAL_1_MINUTE,
// then a message of this type will contain data for a 1 minute timeframe
// with a start time specified by the StartDateTime field.
//
// Even when RecordInterval is INTERVAL_TICK, the HistoricalPriceDataTickRecordResponse
// message can still be used instead of HistoricalPriceDataTickRecordResponse.
// message can still be used instead of HistoricalPriceDataTickRecordResponse.
//
// This message can be part of a compressed series of messages of this same
// type, if the Client requested compression be used.
type HistoricalPriceDataRecordResponse struct {
	message.Fixed
}

// HistoricalPriceDataRecordResponseBuilder The HistoricalPriceDataTickRecordResponse message is used when the RecordInterval
// field in a historical data request message is set to a value greater than
// INTERVAL_TICK. For example, if the RecordInterval is INTERVAL_1_MINUTE,
// then a message of this type will contain data for a 1 minute timeframe
// with a start time specified by the StartDateTime field.
//
// Even when RecordInterval is INTERVAL_TICK, the HistoricalPriceDataTickRecordResponse
// message can still be used instead of HistoricalPriceDataTickRecordResponse.
// message can still be used instead of HistoricalPriceDataTickRecordResponse.
//
// This message can be part of a compressed series of messages of this same
// type, if the Client requested compression be used.
type HistoricalPriceDataRecordResponseBuilder struct {
	message.Fixed
}

// HistoricalPriceDataRecordResponsePointer The HistoricalPriceDataTickRecordResponse message is used when the RecordInterval
// field in a historical data request message is set to a value greater than
// INTERVAL_TICK. For example, if the RecordInterval is INTERVAL_1_MINUTE,
// then a message of this type will contain data for a 1 minute timeframe
// with a start time specified by the StartDateTime field.
//
// Even when RecordInterval is INTERVAL_TICK, the HistoricalPriceDataTickRecordResponse
// message can still be used instead of HistoricalPriceDataTickRecordResponse.
// message can still be used instead of HistoricalPriceDataTickRecordResponse.
//
// This message can be part of a compressed series of messages of this same
// type, if the Client requested compression be used.
type HistoricalPriceDataRecordResponsePointer struct {
	message.FixedPointer
}

// HistoricalPriceDataRecordResponsePointerBuilder The HistoricalPriceDataTickRecordResponse message is used when the RecordInterval
// field in a historical data request message is set to a value greater than
// INTERVAL_TICK. For example, if the RecordInterval is INTERVAL_1_MINUTE,
// then a message of this type will contain data for a 1 minute timeframe
// with a start time specified by the StartDateTime field.
//
// Even when RecordInterval is INTERVAL_TICK, the HistoricalPriceDataTickRecordResponse
// message can still be used instead of HistoricalPriceDataTickRecordResponse.
// message can still be used instead of HistoricalPriceDataTickRecordResponse.
//
// This message can be part of a compressed series of messages of this same
// type, if the Client requested compression be used.
type HistoricalPriceDataRecordResponsePointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalPriceDataRecordResponseFrom(b []byte) HistoricalPriceDataRecordResponse {
	return HistoricalPriceDataRecordResponse{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalPriceDataRecordResponse(b []byte) HistoricalPriceDataRecordResponse {
	return HistoricalPriceDataRecordResponse{Fixed: message.WrapFixed(b)}
}

func NewHistoricalPriceDataRecordResponse() HistoricalPriceDataRecordResponseBuilder {
	a := HistoricalPriceDataRecordResponseBuilder{message.NewFixed(88)}
	a.Unsafe().SetBytes(0, _HistoricalPriceDataRecordResponseDefault)
	return a
}

func AllocHistoricalPriceDataRecordResponse() HistoricalPriceDataRecordResponsePointerBuilder {
	a := HistoricalPriceDataRecordResponsePointerBuilder{message.AllocFixed(88)}
	a.Ptr.SetBytes(0, _HistoricalPriceDataRecordResponseDefault)
	return a
}

func AllocHistoricalPriceDataRecordResponseFrom(b []byte) HistoricalPriceDataRecordResponsePointer {
	return HistoricalPriceDataRecordResponsePointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = HistoricalPriceDataRecordResponseSize  (88)
//     Type          = HISTORICAL_PRICE_DATA_RECORD_RESPONSE  (803)
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
func (m HistoricalPriceDataRecordResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataRecordResponseDefault)
}

// Clear
// {
//     Size          = HistoricalPriceDataRecordResponseSize  (88)
//     Type          = HISTORICAL_PRICE_DATA_RECORD_RESPONSE  (803)
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
func (m HistoricalPriceDataRecordResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataRecordResponseDefault)
}

// ToBuilder
func (m HistoricalPriceDataRecordResponse) ToBuilder() HistoricalPriceDataRecordResponseBuilder {
	return HistoricalPriceDataRecordResponseBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalPriceDataRecordResponsePointer) ToBuilder() HistoricalPriceDataRecordResponsePointerBuilder {
	return HistoricalPriceDataRecordResponsePointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalPriceDataRecordResponseBuilder) Finish() HistoricalPriceDataRecordResponse {
	return HistoricalPriceDataRecordResponse{m.Fixed}
}

// Finish
func (m *HistoricalPriceDataRecordResponsePointerBuilder) Finish() HistoricalPriceDataRecordResponsePointer {
	return HistoricalPriceDataRecordResponsePointer{m.FixedPointer}
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRecordResponse) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRecordResponseBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRecordResponsePointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRecordResponsePointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// StartDateTime The starting Date-Time in UTC of the data record in this message.
//
// It is part of the DTC Protocol specification that this must be the starting
// Date-Time of the data record.
func (m HistoricalPriceDataRecordResponse) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// StartDateTime The starting Date-Time in UTC of the data record in this message.
//
// It is part of the DTC Protocol specification that this must be the starting
// Date-Time of the data record.
func (m HistoricalPriceDataRecordResponseBuilder) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// StartDateTime The starting Date-Time in UTC of the data record in this message.
//
// It is part of the DTC Protocol specification that this must be the starting
// Date-Time of the data record.
func (m HistoricalPriceDataRecordResponsePointer) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 16, 8))
}

// StartDateTime The starting Date-Time in UTC of the data record in this message.
//
// It is part of the DTC Protocol specification that this must be the starting
// Date-Time of the data record.
func (m HistoricalPriceDataRecordResponsePointerBuilder) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 16, 8))
}

// OpenPrice The Open price of the data record in this message.
func (m HistoricalPriceDataRecordResponse) OpenPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// OpenPrice The Open price of the data record in this message.
func (m HistoricalPriceDataRecordResponseBuilder) OpenPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// OpenPrice The Open price of the data record in this message.
func (m HistoricalPriceDataRecordResponsePointer) OpenPrice() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// OpenPrice The Open price of the data record in this message.
func (m HistoricalPriceDataRecordResponsePointerBuilder) OpenPrice() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// HighPrice The High price of the data record in this message.
//
// In the case where NumTrades is 1, the HighPrice field can be the Ask/Offer
// price at the time of the trade. In this case the OpenPrice field needs
// to be 0 in this case.
func (m HistoricalPriceDataRecordResponse) HighPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// HighPrice The High price of the data record in this message.
//
// In the case where NumTrades is 1, the HighPrice field can be the Ask/Offer
// price at the time of the trade. In this case the OpenPrice field needs
// to be 0 in this case.
func (m HistoricalPriceDataRecordResponseBuilder) HighPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// HighPrice The High price of the data record in this message.
//
// In the case where NumTrades is 1, the HighPrice field can be the Ask/Offer
// price at the time of the trade. In this case the OpenPrice field needs
// to be 0 in this case.
func (m HistoricalPriceDataRecordResponsePointer) HighPrice() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
}

// HighPrice The High price of the data record in this message.
//
// In the case where NumTrades is 1, the HighPrice field can be the Ask/Offer
// price at the time of the trade. In this case the OpenPrice field needs
// to be 0 in this case.
func (m HistoricalPriceDataRecordResponsePointerBuilder) HighPrice() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
}

// LowPrice The Low price of the data record in this message.
//
// In the case where NumTrades is 1, the LowPrice field can be the Bid price
// at the time of the trade. In this case the OpenPrice field needs to be
// 0 in this case.
func (m HistoricalPriceDataRecordResponse) LowPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 40, 32)
}

// LowPrice The Low price of the data record in this message.
//
// In the case where NumTrades is 1, the LowPrice field can be the Bid price
// at the time of the trade. In this case the OpenPrice field needs to be
// 0 in this case.
func (m HistoricalPriceDataRecordResponseBuilder) LowPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 40, 32)
}

// LowPrice The Low price of the data record in this message.
//
// In the case where NumTrades is 1, the LowPrice field can be the Bid price
// at the time of the trade. In this case the OpenPrice field needs to be
// 0 in this case.
func (m HistoricalPriceDataRecordResponsePointer) LowPrice() float64 {
	return message.Float64Fixed(m.Ptr, 40, 32)
}

// LowPrice The Low price of the data record in this message.
//
// In the case where NumTrades is 1, the LowPrice field can be the Bid price
// at the time of the trade. In this case the OpenPrice field needs to be
// 0 in this case.
func (m HistoricalPriceDataRecordResponsePointerBuilder) LowPrice() float64 {
	return message.Float64Fixed(m.Ptr, 40, 32)
}

// LastPrice The Last price of the data record in this message.
func (m HistoricalPriceDataRecordResponse) LastPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 48, 40)
}

// LastPrice The Last price of the data record in this message.
func (m HistoricalPriceDataRecordResponseBuilder) LastPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 48, 40)
}

// LastPrice The Last price of the data record in this message.
func (m HistoricalPriceDataRecordResponsePointer) LastPrice() float64 {
	return message.Float64Fixed(m.Ptr, 48, 40)
}

// LastPrice The Last price of the data record in this message.
func (m HistoricalPriceDataRecordResponsePointerBuilder) LastPrice() float64 {
	return message.Float64Fixed(m.Ptr, 48, 40)
}

// Volume The Volume of this data record of this message.
//
// (union)
func (m HistoricalPriceDataRecordResponse) Volume() float64 {
	return message.Float64Fixed(m.Unsafe(), 56, 48)
}

// Volume The Volume of this data record of this message.
//
// (union)
func (m HistoricalPriceDataRecordResponseBuilder) Volume() float64 {
	return message.Float64Fixed(m.Unsafe(), 56, 48)
}

// Volume The Volume of this data record of this message.
//
// (union)
func (m HistoricalPriceDataRecordResponsePointer) Volume() float64 {
	return message.Float64Fixed(m.Ptr, 56, 48)
}

// Volume The Volume of this data record of this message.
//
// (union)
func (m HistoricalPriceDataRecordResponsePointerBuilder) Volume() float64 {
	return message.Float64Fixed(m.Ptr, 56, 48)
}

func (m HistoricalPriceDataRecordResponse) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 60, 56)
}

func (m HistoricalPriceDataRecordResponseBuilder) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 60, 56)
}

func (m HistoricalPriceDataRecordResponsePointer) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Ptr, 60, 56)
}

func (m HistoricalPriceDataRecordResponsePointerBuilder) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Ptr, 60, 56)
}

// NumTrades The Open Interest or Number of Trades of this data record in this message.
// The Open Interest or Number of Trades of this data record in this message.
func (m HistoricalPriceDataRecordResponse) NumTrades() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 60, 56)
}

// NumTrades The Open Interest or Number of Trades of this data record in this message.
// The Open Interest or Number of Trades of this data record in this message.
func (m HistoricalPriceDataRecordResponseBuilder) NumTrades() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 60, 56)
}

// NumTrades The Open Interest or Number of Trades of this data record in this message.
// The Open Interest or Number of Trades of this data record in this message.
func (m HistoricalPriceDataRecordResponsePointer) NumTrades() uint32 {
	return message.Uint32Fixed(m.Ptr, 60, 56)
}

// NumTrades The Open Interest or Number of Trades of this data record in this message.
// The Open Interest or Number of Trades of this data record in this message.
func (m HistoricalPriceDataRecordResponsePointerBuilder) NumTrades() uint32 {
	return message.Uint32Fixed(m.Ptr, 60, 56)
}

// BidVolume The volume of trades at the bid price or lower of the data record in this
// message.
//
// In the case where this message consists of a single trade, if the trade
// was at the Ask, then BidVolume must be zero.
func (m HistoricalPriceDataRecordResponse) BidVolume() float64 {
	return message.Float64Fixed(m.Unsafe(), 72, 64)
}

// BidVolume The volume of trades at the bid price or lower of the data record in this
// message.
//
// In the case where this message consists of a single trade, if the trade
// was at the Ask, then BidVolume must be zero.
func (m HistoricalPriceDataRecordResponseBuilder) BidVolume() float64 {
	return message.Float64Fixed(m.Unsafe(), 72, 64)
}

// BidVolume The volume of trades at the bid price or lower of the data record in this
// message.
//
// In the case where this message consists of a single trade, if the trade
// was at the Ask, then BidVolume must be zero.
func (m HistoricalPriceDataRecordResponsePointer) BidVolume() float64 {
	return message.Float64Fixed(m.Ptr, 72, 64)
}

// BidVolume The volume of trades at the bid price or lower of the data record in this
// message.
//
// In the case where this message consists of a single trade, if the trade
// was at the Ask, then BidVolume must be zero.
func (m HistoricalPriceDataRecordResponsePointerBuilder) BidVolume() float64 {
	return message.Float64Fixed(m.Ptr, 72, 64)
}

// AskVolume The volume of trades at the ask price or higher of the data record in
// this message.
//
// In the case where this message consists of a single trade, if the trade
// was at the Bid, then AskVolume must be zero.
func (m HistoricalPriceDataRecordResponse) AskVolume() float64 {
	return message.Float64Fixed(m.Unsafe(), 80, 72)
}

// AskVolume The volume of trades at the ask price or higher of the data record in
// this message.
//
// In the case where this message consists of a single trade, if the trade
// was at the Bid, then AskVolume must be zero.
func (m HistoricalPriceDataRecordResponseBuilder) AskVolume() float64 {
	return message.Float64Fixed(m.Unsafe(), 80, 72)
}

// AskVolume The volume of trades at the ask price or higher of the data record in
// this message.
//
// In the case where this message consists of a single trade, if the trade
// was at the Bid, then AskVolume must be zero.
func (m HistoricalPriceDataRecordResponsePointer) AskVolume() float64 {
	return message.Float64Fixed(m.Ptr, 80, 72)
}

// AskVolume The volume of trades at the ask price or higher of the data record in
// this message.
//
// In the case where this message consists of a single trade, if the trade
// was at the Bid, then AskVolume must be zero.
func (m HistoricalPriceDataRecordResponsePointerBuilder) AskVolume() float64 {
	return message.Float64Fixed(m.Ptr, 80, 72)
}

// IsFinalRecord Set to 1 to indicate final record in response to the historical price
// data request.
//
// The default is 0 meaning there are more records to follow.
func (m HistoricalPriceDataRecordResponse) IsFinalRecord() bool {
	return message.BoolFixed(m.Unsafe(), 81, 80)
}

// IsFinalRecord Set to 1 to indicate final record in response to the historical price
// data request.
//
// The default is 0 meaning there are more records to follow.
func (m HistoricalPriceDataRecordResponseBuilder) IsFinalRecord() bool {
	return message.BoolFixed(m.Unsafe(), 81, 80)
}

// IsFinalRecord Set to 1 to indicate final record in response to the historical price
// data request.
//
// The default is 0 meaning there are more records to follow.
func (m HistoricalPriceDataRecordResponsePointer) IsFinalRecord() bool {
	return message.BoolFixed(m.Ptr, 81, 80)
}

// IsFinalRecord Set to 1 to indicate final record in response to the historical price
// data request.
//
// The default is 0 meaning there are more records to follow.
func (m HistoricalPriceDataRecordResponsePointerBuilder) IsFinalRecord() bool {
	return message.BoolFixed(m.Ptr, 81, 80)
}

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRecordResponseBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataRecordResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetStartDateTime The starting Date-Time in UTC of the data record in this message.
//
// It is part of the DTC Protocol specification that this must be the starting
// Date-Time of the data record.
func (m HistoricalPriceDataRecordResponseBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, int64(value))
}

// SetStartDateTime The starting Date-Time in UTC of the data record in this message.
//
// It is part of the DTC Protocol specification that this must be the starting
// Date-Time of the data record.
func (m HistoricalPriceDataRecordResponsePointerBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
}

// SetOpenPrice The Open price of the data record in this message.
func (m HistoricalPriceDataRecordResponseBuilder) SetOpenPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetOpenPrice The Open price of the data record in this message.
func (m HistoricalPriceDataRecordResponsePointerBuilder) SetOpenPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetHighPrice The High price of the data record in this message.
//
// In the case where NumTrades is 1, the HighPrice field can be the Ask/Offer
// price at the time of the trade. In this case the OpenPrice field needs
// to be 0 in this case.
func (m HistoricalPriceDataRecordResponseBuilder) SetHighPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 32, 24, value)
}

// SetHighPrice The High price of the data record in this message.
//
// In the case where NumTrades is 1, the HighPrice field can be the Ask/Offer
// price at the time of the trade. In this case the OpenPrice field needs
// to be 0 in this case.
func (m HistoricalPriceDataRecordResponsePointerBuilder) SetHighPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, value)
}

// SetLowPrice The Low price of the data record in this message.
//
// In the case where NumTrades is 1, the LowPrice field can be the Bid price
// at the time of the trade. In this case the OpenPrice field needs to be
// 0 in this case.
func (m HistoricalPriceDataRecordResponseBuilder) SetLowPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 40, 32, value)
}

// SetLowPrice The Low price of the data record in this message.
//
// In the case where NumTrades is 1, the LowPrice field can be the Bid price
// at the time of the trade. In this case the OpenPrice field needs to be
// 0 in this case.
func (m HistoricalPriceDataRecordResponsePointerBuilder) SetLowPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 40, 32, value)
}

// SetLastPrice The Last price of the data record in this message.
func (m HistoricalPriceDataRecordResponseBuilder) SetLastPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 48, 40, value)
}

// SetLastPrice The Last price of the data record in this message.
func (m HistoricalPriceDataRecordResponsePointerBuilder) SetLastPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 48, 40, value)
}

// SetVolume The Volume of this data record of this message.
//
// (union)
func (m HistoricalPriceDataRecordResponseBuilder) SetVolume(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 56, 48, value)
}

// SetVolume The Volume of this data record of this message.
//
// (union)
func (m HistoricalPriceDataRecordResponsePointerBuilder) SetVolume(value float64) {
	message.SetFloat64Fixed(m.Ptr, 56, 48, value)
}

func (m HistoricalPriceDataRecordResponseBuilder) SetOpenInterest(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 60, 56, value)
}

func (m HistoricalPriceDataRecordResponsePointerBuilder) SetOpenInterest(value uint32) {
	message.SetUint32Fixed(m.Ptr, 60, 56, value)
}

// SetNumTrades The Open Interest or Number of Trades of this data record in this message.
// The Open Interest or Number of Trades of this data record in this message.
func (m HistoricalPriceDataRecordResponseBuilder) SetNumTrades(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 60, 56, value)
}

// SetNumTrades The Open Interest or Number of Trades of this data record in this message.
// The Open Interest or Number of Trades of this data record in this message.
func (m HistoricalPriceDataRecordResponsePointerBuilder) SetNumTrades(value uint32) {
	message.SetUint32Fixed(m.Ptr, 60, 56, value)
}

// SetBidVolume The volume of trades at the bid price or lower of the data record in this
// message.
//
// In the case where this message consists of a single trade, if the trade
// was at the Ask, then BidVolume must be zero.
func (m HistoricalPriceDataRecordResponseBuilder) SetBidVolume(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 72, 64, value)
}

// SetBidVolume The volume of trades at the bid price or lower of the data record in this
// message.
//
// In the case where this message consists of a single trade, if the trade
// was at the Ask, then BidVolume must be zero.
func (m HistoricalPriceDataRecordResponsePointerBuilder) SetBidVolume(value float64) {
	message.SetFloat64Fixed(m.Ptr, 72, 64, value)
}

// SetAskVolume The volume of trades at the ask price or higher of the data record in
// this message.
//
// In the case where this message consists of a single trade, if the trade
// was at the Bid, then AskVolume must be zero.
func (m HistoricalPriceDataRecordResponseBuilder) SetAskVolume(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 80, 72, value)
}

// SetAskVolume The volume of trades at the ask price or higher of the data record in
// this message.
//
// In the case where this message consists of a single trade, if the trade
// was at the Bid, then AskVolume must be zero.
func (m HistoricalPriceDataRecordResponsePointerBuilder) SetAskVolume(value float64) {
	message.SetFloat64Fixed(m.Ptr, 80, 72, value)
}

// SetIsFinalRecord Set to 1 to indicate final record in response to the historical price
// data request.
//
// The default is 0 meaning there are more records to follow.
func (m HistoricalPriceDataRecordResponseBuilder) SetIsFinalRecord(value bool) {
	message.SetBoolFixed(m.Unsafe(), 81, 80, value)
}

// SetIsFinalRecord Set to 1 to indicate final record in response to the historical price
// data request.
//
// The default is 0 meaning there are more records to follow.
func (m HistoricalPriceDataRecordResponsePointerBuilder) SetIsFinalRecord(value bool) {
	message.SetBoolFixed(m.Ptr, 81, 80, value)
}

// Copy
func (m HistoricalPriceDataRecordResponse) Copy(to HistoricalPriceDataRecordResponseBuilder) {
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
func (m HistoricalPriceDataRecordResponseBuilder) Copy(to HistoricalPriceDataRecordResponseBuilder) {
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
func (m HistoricalPriceDataRecordResponse) CopyPointer(to HistoricalPriceDataRecordResponsePointerBuilder) {
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
func (m HistoricalPriceDataRecordResponseBuilder) CopyPointer(to HistoricalPriceDataRecordResponsePointerBuilder) {
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
func (m HistoricalPriceDataRecordResponsePointer) Copy(to HistoricalPriceDataRecordResponseBuilder) {
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
func (m HistoricalPriceDataRecordResponsePointerBuilder) Copy(to HistoricalPriceDataRecordResponseBuilder) {
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
func (m HistoricalPriceDataRecordResponsePointer) CopyPointer(to HistoricalPriceDataRecordResponsePointerBuilder) {
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
func (m HistoricalPriceDataRecordResponsePointerBuilder) CopyPointer(to HistoricalPriceDataRecordResponsePointerBuilder) {
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