package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalMarketDepthDataResponseHeaderPointer struct {
	message.FixedPointer
}

type HistoricalMarketDepthDataResponseHeaderPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalMarketDepthDataResponseHeader() HistoricalMarketDepthDataResponseHeaderPointerBuilder {
	a := HistoricalMarketDepthDataResponseHeaderPointerBuilder{message.AllocFixed(12)}
	a.Ptr.SetBytes(0, _HistoricalMarketDepthDataResponseHeaderDefault)
	return a
}

func AllocHistoricalMarketDepthDataResponseHeaderFrom(b []byte) HistoricalMarketDepthDataResponseHeaderPointer {
	return HistoricalMarketDepthDataResponseHeaderPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size               = HistoricalMarketDepthDataResponseHeaderSize  (12)
//     Type               = HISTORICAL_MARKET_DEPTH_DATA_RESPONSE_HEADER  (901)
//     RequestID          = 0
//     UseZLibCompression = 0
//     NoRecordsToReturn  = 0
// }
func (m HistoricalMarketDepthDataResponseHeaderPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalMarketDepthDataResponseHeaderDefault)
}

// ToBuilder
func (m HistoricalMarketDepthDataResponseHeaderPointer) ToBuilder() HistoricalMarketDepthDataResponseHeaderPointerBuilder {
	return HistoricalMarketDepthDataResponseHeaderPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalMarketDepthDataResponseHeaderPointerBuilder) Finish() HistoricalMarketDepthDataResponseHeaderPointer {
	return HistoricalMarketDepthDataResponseHeaderPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalMarketDepthDataResponseHeaderPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalMarketDepthDataResponseHeaderPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// UseZLibCompression
func (m HistoricalMarketDepthDataResponseHeaderPointer) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Ptr, 9, 8)
}

// UseZLibCompression
func (m HistoricalMarketDepthDataResponseHeaderPointerBuilder) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Ptr, 9, 8)
}

// NoRecordsToReturn
func (m HistoricalMarketDepthDataResponseHeaderPointer) NoRecordsToReturn() uint8 {
	return message.Uint8Fixed(m.Ptr, 10, 9)
}

// NoRecordsToReturn
func (m HistoricalMarketDepthDataResponseHeaderPointerBuilder) NoRecordsToReturn() uint8 {
	return message.Uint8Fixed(m.Ptr, 10, 9)
}

// SetRequestID
func (m HistoricalMarketDepthDataResponseHeaderPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetUseZLibCompression
func (m HistoricalMarketDepthDataResponseHeaderPointerBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8Fixed(m.Ptr, 9, 8, value)
}

// SetNoRecordsToReturn
func (m HistoricalMarketDepthDataResponseHeaderPointerBuilder) SetNoRecordsToReturn(value uint8) {
	message.SetUint8Fixed(m.Ptr, 10, 9, value)
}

// Copy
func (m HistoricalMarketDepthDataResponseHeaderPointer) Copy(to HistoricalMarketDepthDataResponseHeaderBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetNoRecordsToReturn(m.NoRecordsToReturn())
}

// Copy
func (m HistoricalMarketDepthDataResponseHeaderPointerBuilder) Copy(to HistoricalMarketDepthDataResponseHeaderBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetNoRecordsToReturn(m.NoRecordsToReturn())
}

// CopyPointer
func (m HistoricalMarketDepthDataResponseHeaderPointer) CopyPointer(to HistoricalMarketDepthDataResponseHeaderPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetNoRecordsToReturn(m.NoRecordsToReturn())
}

// CopyPointer
func (m HistoricalMarketDepthDataResponseHeaderPointerBuilder) CopyPointer(to HistoricalMarketDepthDataResponseHeaderPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetNoRecordsToReturn(m.NoRecordsToReturn())
}
