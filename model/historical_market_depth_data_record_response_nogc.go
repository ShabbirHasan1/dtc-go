package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalMarketDepthDataRecordResponsePointer struct {
	message.FixedPointer
}

type HistoricalMarketDepthDataRecordResponsePointerBuilder struct {
	message.FixedPointer
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
//     IsFinalRecord = 0
// }
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalMarketDepthDataRecordResponseDefault)
}

// ToBuilder
func (m HistoricalMarketDepthDataRecordResponsePointer) ToBuilder() HistoricalMarketDepthDataRecordResponsePointerBuilder {
	return HistoricalMarketDepthDataRecordResponsePointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalMarketDepthDataRecordResponsePointerBuilder) Finish() HistoricalMarketDepthDataRecordResponsePointer {
	return HistoricalMarketDepthDataRecordResponsePointer{m.FixedPointer}
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
func (m HistoricalMarketDepthDataRecordResponsePointer) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 16, 8))
}

// StartDateTime
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 16, 8))
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
func (m HistoricalMarketDepthDataRecordResponsePointer) Flags() uint8 {
	return message.Uint8Fixed(m.Ptr, 18, 17)
}

// Flags
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) Flags() uint8 {
	return message.Uint8Fixed(m.Ptr, 18, 17)
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
func (m HistoricalMarketDepthDataRecordResponsePointer) Price() float32 {
	return message.Float32Fixed(m.Ptr, 24, 20)
}

// Price
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) Price() float32 {
	return message.Float32Fixed(m.Ptr, 24, 20)
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
func (m HistoricalMarketDepthDataRecordResponsePointer) IsFinalRecord() uint8 {
	return message.Uint8Fixed(m.Ptr, 29, 28)
}

// IsFinalRecord
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) IsFinalRecord() uint8 {
	return message.Uint8Fixed(m.Ptr, 29, 28)
}

// SetRequestID
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetStartDateTime
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetStartDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 16, 8, int64(value))
}

// SetCommand
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetCommand(value uint8) {
	message.SetUint8Fixed(m.Ptr, 17, 16, value)
}

// SetFlags
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetFlags(value uint8) {
	message.SetUint8Fixed(m.Ptr, 18, 17, value)
}

// SetNumOrders
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetNumOrders(value uint16) {
	message.SetUint16Fixed(m.Ptr, 20, 18, value)
}

// SetPrice
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetPrice(value float32) {
	message.SetFloat32Fixed(m.Ptr, 24, 20, value)
}

// SetQuantity
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetQuantity(value uint32) {
	message.SetUint32Fixed(m.Ptr, 28, 24, value)
}

// SetIsFinalRecord
func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) SetIsFinalRecord(value uint8) {
	message.SetUint8Fixed(m.Ptr, 29, 28, value)
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
