package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size               = HistoricalMarketDepthDataResponseHeaderSize  (12)
//     Type               = HISTORICAL_MARKET_DEPTH_DATA_RESPONSE_HEADER  (901)
//     RequestID          = 0
//     UseZLibCompression = 0
//     NoRecordsToReturn  = 0
// }
var _HistoricalMarketDepthDataResponseHeaderDefault = []byte{12, 0, 133, 3, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalMarketDepthDataResponseHeaderSize = 12

type HistoricalMarketDepthDataResponseHeader struct {
	message.Fixed
}

type HistoricalMarketDepthDataResponseHeaderBuilder struct {
	message.Fixed
}

func NewHistoricalMarketDepthDataResponseHeaderFrom(b []byte) HistoricalMarketDepthDataResponseHeader {
	return HistoricalMarketDepthDataResponseHeader{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalMarketDepthDataResponseHeader(b []byte) HistoricalMarketDepthDataResponseHeader {
	return HistoricalMarketDepthDataResponseHeader{Fixed: message.WrapFixed(b)}
}

func NewHistoricalMarketDepthDataResponseHeader() HistoricalMarketDepthDataResponseHeaderBuilder {
	a := HistoricalMarketDepthDataResponseHeaderBuilder{message.NewFixed(12)}
	a.Unsafe().SetBytes(0, _HistoricalMarketDepthDataResponseHeaderDefault)
	return a
}

// Clear
// {
//     Size               = HistoricalMarketDepthDataResponseHeaderSize  (12)
//     Type               = HISTORICAL_MARKET_DEPTH_DATA_RESPONSE_HEADER  (901)
//     RequestID          = 0
//     UseZLibCompression = 0
//     NoRecordsToReturn  = 0
// }
func (m HistoricalMarketDepthDataResponseHeaderBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalMarketDepthDataResponseHeaderDefault)
}

// ToBuilder
func (m HistoricalMarketDepthDataResponseHeader) ToBuilder() HistoricalMarketDepthDataResponseHeaderBuilder {
	return HistoricalMarketDepthDataResponseHeaderBuilder{m.Fixed}
}

// Finish
func (m HistoricalMarketDepthDataResponseHeaderBuilder) Finish() HistoricalMarketDepthDataResponseHeader {
	return HistoricalMarketDepthDataResponseHeader{m.Fixed}
}

// RequestID
func (m HistoricalMarketDepthDataResponseHeader) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalMarketDepthDataResponseHeaderBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// UseZLibCompression
func (m HistoricalMarketDepthDataResponseHeader) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 9, 8)
}

// UseZLibCompression
func (m HistoricalMarketDepthDataResponseHeaderBuilder) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 9, 8)
}

// NoRecordsToReturn
func (m HistoricalMarketDepthDataResponseHeader) NoRecordsToReturn() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 10, 9)
}

// NoRecordsToReturn
func (m HistoricalMarketDepthDataResponseHeaderBuilder) NoRecordsToReturn() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 10, 9)
}

// SetRequestID
func (m HistoricalMarketDepthDataResponseHeaderBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetUseZLibCompression
func (m HistoricalMarketDepthDataResponseHeaderBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 9, 8, value)
}

// SetNoRecordsToReturn
func (m HistoricalMarketDepthDataResponseHeaderBuilder) SetNoRecordsToReturn(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 10, 9, value)
}

// Copy
func (m HistoricalMarketDepthDataResponseHeader) Copy(to HistoricalMarketDepthDataResponseHeaderBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetNoRecordsToReturn(m.NoRecordsToReturn())
}

// Copy
func (m HistoricalMarketDepthDataResponseHeaderBuilder) Copy(to HistoricalMarketDepthDataResponseHeaderBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetNoRecordsToReturn(m.NoRecordsToReturn())
}

// CopyPointer
func (m HistoricalMarketDepthDataResponseHeader) CopyPointer(to HistoricalMarketDepthDataResponseHeaderPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetNoRecordsToReturn(m.NoRecordsToReturn())
}

// CopyPointer
func (m HistoricalMarketDepthDataResponseHeaderBuilder) CopyPointer(to HistoricalMarketDepthDataResponseHeaderPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetNoRecordsToReturn(m.NoRecordsToReturn())
}
