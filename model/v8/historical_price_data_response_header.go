// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-16 08:00:54.519198 +0800 WITA m=+0.055446543

package v8

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalPriceDataResponseHeaderSize = 20

//     Size                    uint16                      = HistoricalPriceDataResponseHeaderSize  (20)
//     Type                    uint16                      = HISTORICAL_PRICE_DATA_RESPONSE_HEADER  (801)
//     RequestID               int32                       = 0
//     RecordInterval          INTERVAL_TICK= HistoricalDataIntervalEnum    (0)
//     UseZLibCompression      bool                        = false
//     NoRecordsToReturn       uint8                       = 0
//     IntToFloatPriceDivisor  float32                     = 0
var _HistoricalPriceDataResponseHeaderDefault = []byte{20, 0, 33, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// HistoricalPriceDataResponseHeader When a historical price data request is not rejected, this message header
// will begin the historical price data response from the Server. There will
// be one HistoricalPriceDataResponseHeader message sent ahead of the HistoricalPriceDataRecordResponse
// / HistoricalPriceDataTickRecordResponse messages. If the NoRecordsToReturn
// field is nonzero, then there are no further records that will be sent
// by the Server in response to the request by the Client.
//
// This message is never compressed.
type HistoricalPriceDataResponseHeader struct {
	message.Fixed
}

// HistoricalPriceDataResponseHeaderBuilder When a historical price data request is not rejected, this message header
// will begin the historical price data response from the Server. There will
// be one HistoricalPriceDataResponseHeader message sent ahead of the HistoricalPriceDataRecordResponse
// / HistoricalPriceDataTickRecordResponse messages. If the NoRecordsToReturn
// field is nonzero, then there are no further records that will be sent
// by the Server in response to the request by the Client.
//
// This message is never compressed.
type HistoricalPriceDataResponseHeaderBuilder struct {
	message.Fixed
}

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

func NewHistoricalPriceDataResponseHeaderFrom(b []byte) HistoricalPriceDataResponseHeader {
	return HistoricalPriceDataResponseHeader{Fixed: message.NewFixed(b)}
}

func WrapHistoricalPriceDataResponseHeader(b []byte) HistoricalPriceDataResponseHeader {
	return HistoricalPriceDataResponseHeader{Fixed: message.WrapFixed(b)}
}

func NewHistoricalPriceDataResponseHeader() HistoricalPriceDataResponseHeaderBuilder {
	return HistoricalPriceDataResponseHeaderBuilder{message.NewFixed(_HistoricalPriceDataResponseHeaderDefault)}
}

func AllocHistoricalPriceDataResponseHeader() HistoricalPriceDataResponseHeaderPointerBuilder {
	return HistoricalPriceDataResponseHeaderPointerBuilder{message.AllocFixed(_HistoricalPriceDataResponseHeaderDefault)}
}

func AllocHistoricalPriceDataResponseHeaderFrom(b []byte) HistoricalPriceDataResponseHeaderPointer {
	return HistoricalPriceDataResponseHeaderPointer{FixedPointer: message.AllocFixed(b)}
}

// Clear
//     Size                    uint16                      = HistoricalPriceDataResponseHeaderSize  (20)
//     Type                    uint16                      = HISTORICAL_PRICE_DATA_RESPONSE_HEADER  (801)
//     RequestID               int32                       = 0
//     RecordInterval          INTERVAL_TICK= HistoricalDataIntervalEnum    (0)
//     UseZLibCompression      bool                        = false
//     NoRecordsToReturn       uint8                       = 0
//     IntToFloatPriceDivisor  float32                     = 0
func (m HistoricalPriceDataResponseHeaderBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataResponseHeaderDefault)
}

// Clear
//     Size                    uint16                      = HistoricalPriceDataResponseHeaderSize  (20)
//     Type                    uint16                      = HISTORICAL_PRICE_DATA_RESPONSE_HEADER  (801)
//     RequestID               int32                       = 0
//     RecordInterval          INTERVAL_TICK= HistoricalDataIntervalEnum    (0)
//     UseZLibCompression      bool                        = false
//     NoRecordsToReturn       uint8                       = 0
//     IntToFloatPriceDivisor  float32                     = 0
func (m HistoricalPriceDataResponseHeaderPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataResponseHeaderDefault)
}

// ToBuilder
func (m HistoricalPriceDataResponseHeader) ToBuilder() HistoricalPriceDataResponseHeaderBuilder {
	return HistoricalPriceDataResponseHeaderBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalPriceDataResponseHeaderPointer) ToBuilder() HistoricalPriceDataResponseHeaderPointerBuilder {
	return HistoricalPriceDataResponseHeaderPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalPriceDataResponseHeaderBuilder) Finish() HistoricalPriceDataResponseHeader {
	return HistoricalPriceDataResponseHeader{m.Fixed}
}

// Finish
func (m *HistoricalPriceDataResponseHeaderPointerBuilder) Finish() HistoricalPriceDataResponseHeaderPointer {
	return HistoricalPriceDataResponseHeaderPointer{m.FixedPointer}
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataResponseHeader) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataResponseHeaderBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m HistoricalPriceDataResponseHeader) RecordInterval() HistoricalDataIntervalEnum {
	return HistoricalDataIntervalEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// RecordInterval The data interval of type HistoricalDataIntervalEnum requested by the
// Client.
func (m HistoricalPriceDataResponseHeaderBuilder) RecordInterval() HistoricalDataIntervalEnum {
	return HistoricalDataIntervalEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
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
func (m HistoricalPriceDataResponseHeader) UseZLibCompression() bool {
	return message.BoolFixed(m.Unsafe(), 13, 12)
}

// UseZLibCompression 1 = All subsequent messages are using standard ZLib compression. 0 = no
// compression.
func (m HistoricalPriceDataResponseHeaderBuilder) UseZLibCompression() bool {
	return message.BoolFixed(m.Unsafe(), 13, 12)
}

// UseZLibCompression 1 = All subsequent messages are using standard ZLib compression. 0 = no
// compression.
func (m HistoricalPriceDataResponseHeaderPointer) UseZLibCompression() bool {
	return message.BoolFixed(m.Ptr, 13, 12)
}

// UseZLibCompression 1 = All subsequent messages are using standard ZLib compression. 0 = no
// compression.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) UseZLibCompression() bool {
	return message.BoolFixed(m.Ptr, 13, 12)
}

// NoRecordsToReturn If there are no records to return in response to the request and there
// was no error, this will be set to 1.
func (m HistoricalPriceDataResponseHeader) NoRecordsToReturn() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 14, 13)
}

// NoRecordsToReturn If there are no records to return in response to the request and there
// was no error, this will be set to 1.
func (m HistoricalPriceDataResponseHeaderBuilder) NoRecordsToReturn() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 14, 13)
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
func (m HistoricalPriceDataResponseHeader) IntToFloatPriceDivisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 20, 16)
}

// IntToFloatPriceDivisor This field is no longer used.
func (m HistoricalPriceDataResponseHeaderBuilder) IntToFloatPriceDivisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 20, 16)
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
func (m HistoricalPriceDataResponseHeaderBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The numeric identifier from the historical price data request that this
// response is in response to.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRecordInterval The data interval of type HistoricalDataIntervalEnum requested by the
// Client.
func (m HistoricalPriceDataResponseHeaderBuilder) SetRecordInterval(value HistoricalDataIntervalEnum) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, int32(value))
}

// SetRecordInterval The data interval of type HistoricalDataIntervalEnum requested by the
// Client.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) SetRecordInterval(value HistoricalDataIntervalEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
}

// SetUseZLibCompression 1 = All subsequent messages are using standard ZLib compression. 0 = no
// compression.
func (m HistoricalPriceDataResponseHeaderBuilder) SetUseZLibCompression(value bool) {
	message.SetBoolFixed(m.Unsafe(), 13, 12, value)
}

// SetUseZLibCompression 1 = All subsequent messages are using standard ZLib compression. 0 = no
// compression.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) SetUseZLibCompression(value bool) {
	message.SetBoolFixed(m.Ptr, 13, 12, value)
}

// SetNoRecordsToReturn If there are no records to return in response to the request and there
// was no error, this will be set to 1.
func (m HistoricalPriceDataResponseHeaderBuilder) SetNoRecordsToReturn(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 14, 13, value)
}

// SetNoRecordsToReturn If there are no records to return in response to the request and there
// was no error, this will be set to 1.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) SetNoRecordsToReturn(value uint8) {
	message.SetUint8Fixed(m.Ptr, 14, 13, value)
}

// SetIntToFloatPriceDivisor This field is no longer used.
func (m HistoricalPriceDataResponseHeaderBuilder) SetIntToFloatPriceDivisor(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 20, 16, value)
}

// SetIntToFloatPriceDivisor This field is no longer used.
func (m HistoricalPriceDataResponseHeaderPointerBuilder) SetIntToFloatPriceDivisor(value float32) {
	message.SetFloat32Fixed(m.Ptr, 20, 16, value)
}
