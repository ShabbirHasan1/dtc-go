package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalPriceDataTickRecordResponseSize = 48

// {
//     Size          = HistoricalPriceDataTickRecordResponseSize  (48)
//     Type          = HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE  (804)
//     RequestID     = 0
//     DateTime      = 0
//     AtBidOrAsk    = BID_ASK_UNSET  (0)
//     Price         = 0
//     Volume        = 0
//     IsFinalRecord = false
// }
var _HistoricalPriceDataTickRecordResponseDefault = []byte{48, 0, 36, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// HistoricalPriceDataTickRecordResponse This is the response message when the RecordInterval field in a historical
// data request message is set to INTERVAL_TICK.
//
// If the Server does not support 1 Tick historical data or does not have
// 1 Tick historical data for the specified time period, it can respond with
// HistoricalPriceDataRecordResponse messages instead. The Server must only
// respond with messages of one type in response to a particular historical
// price data request.
//
// This message can be part of a compressed series of messages of this same
// type, if the Client requested compression be used.
type HistoricalPriceDataTickRecordResponse struct {
	message.Fixed
}

// HistoricalPriceDataTickRecordResponseBuilder This is the response message when the RecordInterval field in a historical
// data request message is set to INTERVAL_TICK.
//
// If the Server does not support 1 Tick historical data or does not have
// 1 Tick historical data for the specified time period, it can respond with
// HistoricalPriceDataRecordResponse messages instead. The Server must only
// respond with messages of one type in response to a particular historical
// price data request.
//
// This message can be part of a compressed series of messages of this same
// type, if the Client requested compression be used.
type HistoricalPriceDataTickRecordResponseBuilder struct {
	message.Fixed
}

// HistoricalPriceDataTickRecordResponsePointer This is the response message when the RecordInterval field in a historical
// data request message is set to INTERVAL_TICK.
//
// If the Server does not support 1 Tick historical data or does not have
// 1 Tick historical data for the specified time period, it can respond with
// HistoricalPriceDataRecordResponse messages instead. The Server must only
// respond with messages of one type in response to a particular historical
// price data request.
//
// This message can be part of a compressed series of messages of this same
// type, if the Client requested compression be used.
type HistoricalPriceDataTickRecordResponsePointer struct {
	message.FixedPointer
}

// HistoricalPriceDataTickRecordResponsePointerBuilder This is the response message when the RecordInterval field in a historical
// data request message is set to INTERVAL_TICK.
//
// If the Server does not support 1 Tick historical data or does not have
// 1 Tick historical data for the specified time period, it can respond with
// HistoricalPriceDataRecordResponse messages instead. The Server must only
// respond with messages of one type in response to a particular historical
// price data request.
//
// This message can be part of a compressed series of messages of this same
// type, if the Client requested compression be used.
type HistoricalPriceDataTickRecordResponsePointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalPriceDataTickRecordResponseFrom(b []byte) HistoricalPriceDataTickRecordResponse {
	return HistoricalPriceDataTickRecordResponse{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalPriceDataTickRecordResponse(b []byte) HistoricalPriceDataTickRecordResponse {
	return HistoricalPriceDataTickRecordResponse{Fixed: message.WrapFixed(b)}
}

func NewHistoricalPriceDataTickRecordResponse() HistoricalPriceDataTickRecordResponseBuilder {
	a := HistoricalPriceDataTickRecordResponseBuilder{message.NewFixed(48)}
	a.Unsafe().SetBytes(0, _HistoricalPriceDataTickRecordResponseDefault)
	return a
}

func AllocHistoricalPriceDataTickRecordResponse() HistoricalPriceDataTickRecordResponsePointerBuilder {
	a := HistoricalPriceDataTickRecordResponsePointerBuilder{message.AllocFixed(48)}
	a.Ptr.SetBytes(0, _HistoricalPriceDataTickRecordResponseDefault)
	return a
}

func AllocHistoricalPriceDataTickRecordResponseFrom(b []byte) HistoricalPriceDataTickRecordResponsePointer {
	return HistoricalPriceDataTickRecordResponsePointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = HistoricalPriceDataTickRecordResponseSize  (48)
//     Type          = HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE  (804)
//     RequestID     = 0
//     DateTime      = 0
//     AtBidOrAsk    = BID_ASK_UNSET  (0)
//     Price         = 0
//     Volume        = 0
//     IsFinalRecord = false
// }
func (m HistoricalPriceDataTickRecordResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataTickRecordResponseDefault)
}

// Clear
// {
//     Size          = HistoricalPriceDataTickRecordResponseSize  (48)
//     Type          = HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE  (804)
//     RequestID     = 0
//     DateTime      = 0
//     AtBidOrAsk    = BID_ASK_UNSET  (0)
//     Price         = 0
//     Volume        = 0
//     IsFinalRecord = false
// }
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataTickRecordResponseDefault)
}

// ToBuilder
func (m HistoricalPriceDataTickRecordResponse) ToBuilder() HistoricalPriceDataTickRecordResponseBuilder {
	return HistoricalPriceDataTickRecordResponseBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalPriceDataTickRecordResponsePointer) ToBuilder() HistoricalPriceDataTickRecordResponsePointerBuilder {
	return HistoricalPriceDataTickRecordResponsePointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalPriceDataTickRecordResponseBuilder) Finish() HistoricalPriceDataTickRecordResponse {
	return HistoricalPriceDataTickRecordResponse{m.Fixed}
}

// Finish
func (m *HistoricalPriceDataTickRecordResponsePointerBuilder) Finish() HistoricalPriceDataTickRecordResponsePointer {
	return HistoricalPriceDataTickRecordResponsePointer{m.FixedPointer}
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataTickRecordResponse) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataTickRecordResponseBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataTickRecordResponsePointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// DateTime The Date-Time of the trade.
func (m HistoricalPriceDataTickRecordResponse) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 16, 8))
}

// DateTime The Date-Time of the trade.
func (m HistoricalPriceDataTickRecordResponseBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 16, 8))
}

// DateTime The Date-Time of the trade.
func (m HistoricalPriceDataTickRecordResponsePointer) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 16, 8))
}

// DateTime The Date-Time of the trade.
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Ptr, 16, 8))
}

// AtBidOrAsk This indicates whether the trade occurred at the Bid price or lower or
// at the Ask price or higher.
func (m HistoricalPriceDataTickRecordResponse) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 18, 16))
}

// AtBidOrAsk This indicates whether the trade occurred at the Bid price or lower or
// at the Ask price or higher.
func (m HistoricalPriceDataTickRecordResponseBuilder) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Unsafe(), 18, 16))
}

// AtBidOrAsk This indicates whether the trade occurred at the Bid price or lower or
// at the Ask price or higher.
func (m HistoricalPriceDataTickRecordResponsePointer) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 18, 16))
}

// AtBidOrAsk This indicates whether the trade occurred at the Bid price or lower or
// at the Ask price or higher.
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(message.Uint16Fixed(m.Ptr, 18, 16))
}

// Price The price of the trade.
func (m HistoricalPriceDataTickRecordResponse) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// Price The price of the trade.
func (m HistoricalPriceDataTickRecordResponseBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// Price The price of the trade.
func (m HistoricalPriceDataTickRecordResponsePointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
}

// Price The price of the trade.
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 32, 24)
}

// Volume The volume of the trade.
func (m HistoricalPriceDataTickRecordResponse) Volume() float64 {
	return message.Float64Fixed(m.Unsafe(), 40, 32)
}

// Volume The volume of the trade.
func (m HistoricalPriceDataTickRecordResponseBuilder) Volume() float64 {
	return message.Float64Fixed(m.Unsafe(), 40, 32)
}

// Volume The volume of the trade.
func (m HistoricalPriceDataTickRecordResponsePointer) Volume() float64 {
	return message.Float64Fixed(m.Ptr, 40, 32)
}

// Volume The volume of the trade.
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) Volume() float64 {
	return message.Float64Fixed(m.Ptr, 40, 32)
}

// IsFinalRecord Set to 1 to indicate final record in response to the historical price
// data request.
//
// The default is 0 meaning there are more records to follow.
func (m HistoricalPriceDataTickRecordResponse) IsFinalRecord() bool {
	return message.BoolFixed(m.Unsafe(), 41, 40)
}

// IsFinalRecord Set to 1 to indicate final record in response to the historical price
// data request.
//
// The default is 0 meaning there are more records to follow.
func (m HistoricalPriceDataTickRecordResponseBuilder) IsFinalRecord() bool {
	return message.BoolFixed(m.Unsafe(), 41, 40)
}

// IsFinalRecord Set to 1 to indicate final record in response to the historical price
// data request.
//
// The default is 0 meaning there are more records to follow.
func (m HistoricalPriceDataTickRecordResponsePointer) IsFinalRecord() bool {
	return message.BoolFixed(m.Ptr, 41, 40)
}

// IsFinalRecord Set to 1 to indicate final record in response to the historical price
// data request.
//
// The default is 0 meaning there are more records to follow.
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) IsFinalRecord() bool {
	return message.BoolFixed(m.Ptr, 41, 40)
}

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataTickRecordResponseBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetDateTime The Date-Time of the trade.
func (m HistoricalPriceDataTickRecordResponseBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, float64(value))
}

// SetDateTime The Date-Time of the trade.
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) SetDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, float64(value))
}

// SetAtBidOrAsk This indicates whether the trade occurred at the Bid price or lower or
// at the Ask price or higher.
func (m HistoricalPriceDataTickRecordResponseBuilder) SetAtBidOrAsk(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Unsafe(), 18, 16, uint16(value))
}

// SetAtBidOrAsk This indicates whether the trade occurred at the Bid price or lower or
// at the Ask price or higher.
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) SetAtBidOrAsk(value AtBidOrAskEnum) {
	message.SetUint16Fixed(m.Ptr, 18, 16, uint16(value))
}

// SetPrice The price of the trade.
func (m HistoricalPriceDataTickRecordResponseBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 32, 24, value)
}

// SetPrice The price of the trade.
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 32, 24, value)
}

// SetVolume The volume of the trade.
func (m HistoricalPriceDataTickRecordResponseBuilder) SetVolume(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 40, 32, value)
}

// SetVolume The volume of the trade.
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) SetVolume(value float64) {
	message.SetFloat64Fixed(m.Ptr, 40, 32, value)
}

// SetIsFinalRecord Set to 1 to indicate final record in response to the historical price
// data request.
//
// The default is 0 meaning there are more records to follow.
func (m HistoricalPriceDataTickRecordResponseBuilder) SetIsFinalRecord(value bool) {
	message.SetBoolFixed(m.Unsafe(), 41, 40, value)
}

// SetIsFinalRecord Set to 1 to indicate final record in response to the historical price
// data request.
//
// The default is 0 meaning there are more records to follow.
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) SetIsFinalRecord(value bool) {
	message.SetBoolFixed(m.Ptr, 41, 40, value)
}

// Copy
func (m HistoricalPriceDataTickRecordResponse) Copy(to HistoricalPriceDataTickRecordResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// Copy
func (m HistoricalPriceDataTickRecordResponseBuilder) Copy(to HistoricalPriceDataTickRecordResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalPriceDataTickRecordResponse) CopyPointer(to HistoricalPriceDataTickRecordResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalPriceDataTickRecordResponseBuilder) CopyPointer(to HistoricalPriceDataTickRecordResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// Copy
func (m HistoricalPriceDataTickRecordResponsePointer) Copy(to HistoricalPriceDataTickRecordResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// Copy
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) Copy(to HistoricalPriceDataTickRecordResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalPriceDataTickRecordResponsePointer) CopyPointer(to HistoricalPriceDataTickRecordResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalPriceDataTickRecordResponsePointerBuilder) CopyPointer(to HistoricalPriceDataTickRecordResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetDateTime(m.DateTime())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetIsFinalRecord(m.IsFinalRecord())
}
