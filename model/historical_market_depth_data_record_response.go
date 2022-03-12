package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalMarketDepthDataRecordResponseSize = 32

// {
//     Size          = HistoricalMarketDepthDataRecordResponseSize  (32)
//     Type          = HISTORICAL_MARKET_DEPTH_DATA_RECORD_RESPONSE  (903)
//     RequestID     = 0
//     StartDateTime = 0
//     Command       = 0
//     Flags         = 0
//     NumOrders     = 0
//     Price         = 0
//     Quantity      = 0
//     IsFinalRecord = false
// }
var _HistoricalMarketDepthDataRecordResponseDefault = []byte{32, 0, 135, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HistoricalMarketDepthDataRecordResponse struct {
	message.Fixed
}

type HistoricalMarketDepthDataRecordResponseBuilder struct {
	message.Fixed
}

type HistoricalMarketDepthDataRecordResponsePointer struct {
	message.FixedPointer
}

type HistoricalMarketDepthDataRecordResponsePointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalMarketDepthDataRecordResponseFrom(b []byte) HistoricalMarketDepthDataRecordResponse {
	return HistoricalMarketDepthDataRecordResponse{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalMarketDepthDataRecordResponse(b []byte) HistoricalMarketDepthDataRecordResponse {
	return HistoricalMarketDepthDataRecordResponse{Fixed: message.WrapFixed(b)}
}

func NewHistoricalMarketDepthDataRecordResponse() HistoricalMarketDepthDataRecordResponseBuilder {
	a := HistoricalMarketDepthDataRecordResponseBuilder{message.NewFixed(32)}
	a.Unsafe().SetBytes(0, _HistoricalMarketDepthDataRecordResponseDefault)
	return a
}

func AllocHistoricalMarketDepthDataRecordResponse() HistoricalMarketDepthDataRecordResponsePointerBuilder {
	a := HistoricalMarketDepthDataRecordResponsePointerBuilder{message.AllocFixed(32)}
	a.Ptr.SetBytes(0, _HistoricalMarketDepthDataRecordResponseDefault)
	return a
}

func AllocHistoricalMarketDepthDataRecordResponseFrom(b []byte) HistoricalMarketDepthDataRecordResponsePointer {
	return HistoricalMarketDepthDataRecordResponsePointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = HistoricalMarketDepthDataRecordResponseSize  (32)
//     Type          = HISTORICAL_MARKET_DEPTH_DATA_RECORD_RESPONSE  (903)
//     RequestID     = 0
//     StartDateTime = 0
//     Command       = 0
//     Flags         = 0
//     NumOrders     = 0
//     Price         = 0
//     Quantity      = 0
//     IsFinalRecord = false
// }
func (m HistoricalMarketDepthDataRecordResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalMarketDepthDataRecordResponseDefault)
}

// Clear
// {
//     Size          = HistoricalMarketDepthDataRecordResponseSize  (32)
//     Type          = HISTORICAL_MARKET_DEPTH_DATA_RECORD_RESPONSE  (903)
//     RequestID     = 0
//     StartDateTime = 0
//     Command       = 0
//     Flags         = 0
//     NumOrders     = 0
//     Price         = 0
//     Quantity      = 0
//     IsFinalRecord = false
// }
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalMarketDepthDataRecordResponseDefault)
}

// ToBuilder
func (m HistoricalMarketDepthDataRecordResponse) ToBuilder() HistoricalMarketDepthDataRecordResponseBuilder {
	return HistoricalMarketDepthDataRecordResponseBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalMarketDepthDataRecordResponsePointer) ToBuilder() HistoricalMarketDepthDataRecordResponsePointerBuilder {
	return HistoricalMarketDepthDataRecordResponsePointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalMarketDepthDataRecordResponseBuilder) Finish() HistoricalMarketDepthDataRecordResponse {
	return HistoricalMarketDepthDataRecordResponse{m.Fixed}
}

// Finish
func (m *HistoricalMarketDepthDataRecordResponsePointerBuilder) Finish() HistoricalMarketDepthDataRecordResponsePointer {
	return HistoricalMarketDepthDataRecordResponsePointer{m.FixedPointer}
}

// RequestID
func (m HistoricalMarketDepthDataRecordResponse) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalMarketDepthDataRecordResponseBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalMarketDepthDataRecordResponsePointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// StartDateTime
func (m HistoricalMarketDepthDataRecordResponse) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// StartDateTime
func (m HistoricalMarketDepthDataRecordResponseBuilder) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 16, 8))
}

// StartDateTime
func (m HistoricalMarketDepthDataRecordResponsePointer) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 16, 8))
}

// StartDateTime
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 16, 8))
}

// Command
func (m HistoricalMarketDepthDataRecordResponse) Command() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 17, 16)
}

// Command
func (m HistoricalMarketDepthDataRecordResponseBuilder) Command() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 17, 16)
}

// Command
func (m HistoricalMarketDepthDataRecordResponsePointer) Command() uint8 {
	return message.Uint8Fixed(m.Ptr, 17, 16)
}

// Command
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) Command() uint8 {
	return message.Uint8Fixed(m.Ptr, 17, 16)
}

// Flags
func (m HistoricalMarketDepthDataRecordResponse) Flags() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 18, 17)
}

// Flags
func (m HistoricalMarketDepthDataRecordResponseBuilder) Flags() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 18, 17)
}

// Flags
func (m HistoricalMarketDepthDataRecordResponsePointer) Flags() uint8 {
	return message.Uint8Fixed(m.Ptr, 18, 17)
}

// Flags
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) Flags() uint8 {
	return message.Uint8Fixed(m.Ptr, 18, 17)
}

// NumOrders
func (m HistoricalMarketDepthDataRecordResponse) NumOrders() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 20, 18)
}

// NumOrders
func (m HistoricalMarketDepthDataRecordResponseBuilder) NumOrders() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 20, 18)
}

// NumOrders
func (m HistoricalMarketDepthDataRecordResponsePointer) NumOrders() uint16 {
	return message.Uint16Fixed(m.Ptr, 20, 18)
}

// NumOrders
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) NumOrders() uint16 {
	return message.Uint16Fixed(m.Ptr, 20, 18)
}

// Price
func (m HistoricalMarketDepthDataRecordResponse) Price() float32 {
	return message.Float32Fixed(m.Unsafe(), 24, 20)
}

// Price
func (m HistoricalMarketDepthDataRecordResponseBuilder) Price() float32 {
	return message.Float32Fixed(m.Unsafe(), 24, 20)
}

// Price
func (m HistoricalMarketDepthDataRecordResponsePointer) Price() float32 {
	return message.Float32Fixed(m.Ptr, 24, 20)
}

// Price
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) Price() float32 {
	return message.Float32Fixed(m.Ptr, 24, 20)
}

// Quantity
func (m HistoricalMarketDepthDataRecordResponse) Quantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 28, 24)
}

// Quantity
func (m HistoricalMarketDepthDataRecordResponseBuilder) Quantity() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 28, 24)
}

// Quantity
func (m HistoricalMarketDepthDataRecordResponsePointer) Quantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 28, 24)
}

// Quantity
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) Quantity() uint32 {
	return message.Uint32Fixed(m.Ptr, 28, 24)
}

// IsFinalRecord
func (m HistoricalMarketDepthDataRecordResponse) IsFinalRecord() bool {
	return message.BoolFixed(m.Unsafe(), 29, 28)
}

// IsFinalRecord
func (m HistoricalMarketDepthDataRecordResponseBuilder) IsFinalRecord() bool {
	return message.BoolFixed(m.Unsafe(), 29, 28)
}

// IsFinalRecord
func (m HistoricalMarketDepthDataRecordResponsePointer) IsFinalRecord() bool {
	return message.BoolFixed(m.Ptr, 29, 28)
}

// IsFinalRecord
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) IsFinalRecord() bool {
	return message.BoolFixed(m.Ptr, 29, 28)
}

// SetRequestID
func (m HistoricalMarketDepthDataRecordResponseBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetStartDateTime
func (m HistoricalMarketDepthDataRecordResponseBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, int64(value))
}

// SetStartDateTime
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
}

// SetCommand
func (m HistoricalMarketDepthDataRecordResponseBuilder) SetCommand(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 17, 16, value)
}

// SetCommand
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetCommand(value uint8) {
	message.SetUint8Fixed(m.Ptr, 17, 16, value)
}

// SetFlags
func (m HistoricalMarketDepthDataRecordResponseBuilder) SetFlags(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 18, 17, value)
}

// SetFlags
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetFlags(value uint8) {
	message.SetUint8Fixed(m.Ptr, 18, 17, value)
}

// SetNumOrders
func (m HistoricalMarketDepthDataRecordResponseBuilder) SetNumOrders(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 20, 18, value)
}

// SetNumOrders
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetNumOrders(value uint16) {
	message.SetUint16Fixed(m.Ptr, 20, 18, value)
}

// SetPrice
func (m HistoricalMarketDepthDataRecordResponseBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 24, 20, value)
}

// SetPrice
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 24, 20, value)
}

// SetQuantity
func (m HistoricalMarketDepthDataRecordResponseBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 28, 24, value)
}

// SetQuantity
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 28, 24, value)
}

// SetIsFinalRecord
func (m HistoricalMarketDepthDataRecordResponseBuilder) SetIsFinalRecord(value bool) {
	message.SetBoolFixed(m.Unsafe(), 29, 28, value)
}

// SetIsFinalRecord
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetIsFinalRecord(value bool) {
	message.SetBoolFixed(m.Ptr, 29, 28, value)
}

// Copy
func (m HistoricalMarketDepthDataRecordResponse) Copy(to HistoricalMarketDepthDataRecordResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
	to.SetCommand(m.Command())
	to.SetFlags(m.Flags())
	to.SetNumOrders(m.NumOrders())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// Copy
func (m HistoricalMarketDepthDataRecordResponseBuilder) Copy(to HistoricalMarketDepthDataRecordResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
	to.SetCommand(m.Command())
	to.SetFlags(m.Flags())
	to.SetNumOrders(m.NumOrders())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalMarketDepthDataRecordResponse) CopyPointer(to HistoricalMarketDepthDataRecordResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
	to.SetCommand(m.Command())
	to.SetFlags(m.Flags())
	to.SetNumOrders(m.NumOrders())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalMarketDepthDataRecordResponseBuilder) CopyPointer(to HistoricalMarketDepthDataRecordResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
	to.SetCommand(m.Command())
	to.SetFlags(m.Flags())
	to.SetNumOrders(m.NumOrders())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// Copy
func (m HistoricalMarketDepthDataRecordResponsePointer) Copy(to HistoricalMarketDepthDataRecordResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
	to.SetCommand(m.Command())
	to.SetFlags(m.Flags())
	to.SetNumOrders(m.NumOrders())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// Copy
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) Copy(to HistoricalMarketDepthDataRecordResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
	to.SetCommand(m.Command())
	to.SetFlags(m.Flags())
	to.SetNumOrders(m.NumOrders())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalMarketDepthDataRecordResponsePointer) CopyPointer(to HistoricalMarketDepthDataRecordResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
	to.SetCommand(m.Command())
	to.SetFlags(m.Flags())
	to.SetNumOrders(m.NumOrders())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetIsFinalRecord(m.IsFinalRecord())
}

// CopyPointer
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) CopyPointer(to HistoricalMarketDepthDataRecordResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetStartDateTime(m.StartDateTime())
	to.SetCommand(m.Command())
	to.SetFlags(m.Flags())
	to.SetNumOrders(m.NumOrders())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetIsFinalRecord(m.IsFinalRecord())
}
