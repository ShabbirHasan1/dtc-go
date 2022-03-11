package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HistoricalPriceDataResponseHeaderPointer When a historical price data request is not rejected, this message header
// will begin the historical price data response from the Server. There will
// be one HistoricalPriceDataResponseHeader message sent ahead of the HistoricalPriceDataRecordResponse
// / HistoricalPriceDataTickRecordResponse messages. If the NoRecordsToReturn
// field is nonzero, then there are no further records that will be sent
// by the Server in response to the request by the Client.
//
// This message is never compressed.
type HistoricalPriceDataResponseHeaderPointer struct {
	message.FixedPointer
}

// HistoricalPriceDataResponseHeaderPointerBuilder When a historical price data request is not rejected, this message header
// will begin the historical price data response from the Server. There will
// be one HistoricalPriceDataResponseHeader message sent ahead of the HistoricalPriceDataRecordResponse
// / HistoricalPriceDataTickRecordResponse messages. If the NoRecordsToReturn
// field is nonzero, then there are no further records that will be sent
// by the Server in response to the request by the Client.
//
// This message is never compressed.
type HistoricalPriceDataResponseHeaderPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalPriceDataResponseHeader() HistoricalPriceDataResponseHeaderPointerBuilder {
	a := HistoricalPriceDataResponseHeaderPointerBuilder{message.AllocFixed(20)}
	a.Ptr.SetBytes(0, _HistoricalPriceDataResponseHeaderDefault)
	return a
}

func AllocHistoricalPriceDataResponseHeaderFrom(b []byte) HistoricalPriceDataResponseHeaderPointer {
	return HistoricalPriceDataResponseHeaderPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                   = HistoricalPriceDataResponseHeaderSize  (20)
//     Type                   = HISTORICAL_PRICE_DATA_RESPONSE_HEADER  (801)
//     RequestID              = 0
//     RecordInterval         = 0
//     UseZLibCompression     = 0
//     NoRecordsToReturn      = 0
//     IntToFloatPriceDivisor = 0
// }
func (m HistoricalPriceDataResponseHeaderPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataResponseHeaderDefault)
}

// ToBuilder
func (m HistoricalPriceDataResponseHeaderPointer) ToBuilder() HistoricalPriceDataResponseHeaderPointerBuilder {
	return HistoricalPriceDataResponseHeaderPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalPriceDataResponseHeaderPointerBuilder) Finish() HistoricalPriceDataResponseHeaderPointer {
	return HistoricalPriceDataResponseHeaderPointer{m.FixedPointer}
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataResponseHeaderPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RecordInterval The data interval of type HistoricalDataIntervalEnum requested by the
// Client.
func (m HistoricalPriceDataResponseHeaderPointer) RecordInterval() HistoricalDataIntervalEnum {
	return HistoricalDataIntervalEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// RecordInterval The data interval of type HistoricalDataIntervalEnum requested by the
// Client.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) RecordInterval() HistoricalDataIntervalEnum {
	return HistoricalDataIntervalEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// UseZLibCompression 1 = All subsequent messages are using standard ZLib compression. 0 = no
// compression.
func (m HistoricalPriceDataResponseHeaderPointer) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Ptr, 13, 12)
}

// UseZLibCompression 1 = All subsequent messages are using standard ZLib compression. 0 = no
// compression.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Ptr, 13, 12)
}

// NoRecordsToReturn If there are no records to return in response to the request and there
// was no error, this will be set to 1.
func (m HistoricalPriceDataResponseHeaderPointer) NoRecordsToReturn() uint8 {
	return message.Uint8Fixed(m.Ptr, 14, 13)
}

// NoRecordsToReturn If there are no records to return in response to the request and there
// was no error, this will be set to 1.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) NoRecordsToReturn() uint8 {
	return message.Uint8Fixed(m.Ptr, 14, 13)
}

// IntToFloatPriceDivisor This field is no longer used.
func (m HistoricalPriceDataResponseHeaderPointer) IntToFloatPriceDivisor() float32 {
	return message.Float32Fixed(m.Ptr, 20, 16)
}

// IntToFloatPriceDivisor This field is no longer used.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) IntToFloatPriceDivisor() float32 {
	return message.Float32Fixed(m.Ptr, 20, 16)
}

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRecordInterval The data interval of type HistoricalDataIntervalEnum requested by the
// Client.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) SetRecordInterval(value HistoricalDataIntervalEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
}

// SetUseZLibCompression 1 = All subsequent messages are using standard ZLib compression. 0 = no
// compression.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8Fixed(m.Ptr, 13, 12, value)
}

// SetNoRecordsToReturn If there are no records to return in response to the request and there
// was no error, this will be set to 1.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) SetNoRecordsToReturn(value uint8) {
	message.SetUint8Fixed(m.Ptr, 14, 13, value)
}

// SetIntToFloatPriceDivisor This field is no longer used.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) SetIntToFloatPriceDivisor(value float32) {
	message.SetFloat32Fixed(m.Ptr, 20, 16, value)
}

// Copy
func (m HistoricalPriceDataResponseHeaderPointer) Copy(to HistoricalPriceDataResponseHeaderBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRecordInterval(m.RecordInterval())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetNoRecordsToReturn(m.NoRecordsToReturn())
	to.SetIntToFloatPriceDivisor(m.IntToFloatPriceDivisor())
}

// Copy
func (m HistoricalPriceDataResponseHeaderPointerBuilder) Copy(to HistoricalPriceDataResponseHeaderBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRecordInterval(m.RecordInterval())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetNoRecordsToReturn(m.NoRecordsToReturn())
	to.SetIntToFloatPriceDivisor(m.IntToFloatPriceDivisor())
}

// CopyPointer
func (m HistoricalPriceDataResponseHeaderPointer) CopyPointer(to HistoricalPriceDataResponseHeaderPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRecordInterval(m.RecordInterval())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetNoRecordsToReturn(m.NoRecordsToReturn())
	to.SetIntToFloatPriceDivisor(m.IntToFloatPriceDivisor())
}

// CopyPointer
func (m HistoricalPriceDataResponseHeaderPointerBuilder) CopyPointer(to HistoricalPriceDataResponseHeaderPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRecordInterval(m.RecordInterval())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetNoRecordsToReturn(m.NoRecordsToReturn())
	to.SetIntToFloatPriceDivisor(m.IntToFloatPriceDivisor())
}
