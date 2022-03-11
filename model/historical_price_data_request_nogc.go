package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HistoricalPriceDataRequestPointer This is a message from the Client to the Server for requesting historical
// price data.
//
// This request can be on the same or a separate network socket connection
// compared to the streaming market data. This is going to be specified by
// the Server.
type HistoricalPriceDataRequestPointer struct {
	message.VLSPointer
}

// HistoricalPriceDataRequestPointerBuilder This is a message from the Client to the Server for requesting historical
// price data.
//
// This request can be on the same or a separate network socket connection
// compared to the streaming market data. This is going to be specified by
// the Server.
type HistoricalPriceDataRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocHistoricalPriceDataRequest() HistoricalPriceDataRequestPointerBuilder {
	a := HistoricalPriceDataRequestPointerBuilder{message.AllocVLSBuilder(48)}
	a.Ptr.SetBytes(0, _HistoricalPriceDataRequestDefault)
	return a
}

func AllocHistoricalPriceDataRequestFrom(b []byte) HistoricalPriceDataRequestPointer {
	return HistoricalPriceDataRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                             = HistoricalPriceDataRequestSize  (48)
//     Type                             = HISTORICAL_PRICE_DATA_REQUEST  (800)
//     BaseSize                         = HistoricalPriceDataRequestSize  (48)
//     RequestID                        = 0
//     Symbol                           = ""
//     Exchange                         = ""
//     RecordInterval                   = 0
//     StartDateTime                    = 0
//     EndDateTime                      = 0
//     MaxDaysToReturn                  = 0
//     UseZLibCompression               = 0
//     RequestDividendAdjustedStockData = 0
//     Integer_1                        = 0
// }
func (m HistoricalPriceDataRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataRequestDefault)
}

// ToBuilder
func (m HistoricalPriceDataRequestPointer) ToBuilder() HistoricalPriceDataRequestPointerBuilder {
	return HistoricalPriceDataRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalPriceDataRequestPointerBuilder) Finish() HistoricalPriceDataRequestPointer {
	return HistoricalPriceDataRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// Symbol The Symbol historical price data is requested for.
func (m HistoricalPriceDataRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Symbol The Symbol historical price data is requested for.
func (m HistoricalPriceDataRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Exchange Optional: The exchange for the Symbol.
func (m HistoricalPriceDataRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Exchange Optional: The exchange for the Symbol.
func (m HistoricalPriceDataRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// RecordInterval The interval/timeframe of each record for the time range of the historical
// data requested. Can be any of:
//
// INTERVAL_TICK = 0
// INTERVAL_1_SECOND = 1
// INTERVAL_1_MINUTE = 60
// INTERVAL_1_DAY = 86400
// INTERVAL_1_WEEK = 604800
func (m HistoricalPriceDataRequestPointer) RecordInterval() HistoricalDataIntervalEnum {
	return HistoricalDataIntervalEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// RecordInterval The interval/timeframe of each record for the time range of the historical
// data requested. Can be any of:
//
// INTERVAL_TICK = 0
// INTERVAL_1_SECOND = 1
// INTERVAL_1_MINUTE = 60
// INTERVAL_1_DAY = 86400
// INTERVAL_1_WEEK = 604800
func (m HistoricalPriceDataRequestPointerBuilder) RecordInterval() HistoricalDataIntervalEnum {
	return HistoricalDataIntervalEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// StartDateTime The starting Date-Time for the historical price data returned, if available
// for the specified Symbol.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data starting at the earliest data available for the Symbol.
func (m HistoricalPriceDataRequestPointer) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 32, 24))
}

// StartDateTime The starting Date-Time for the historical price data returned, if available
// for the specified Symbol.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data starting at the earliest data available for the Symbol.
func (m HistoricalPriceDataRequestPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 32, 24))
}

// EndDateTime The ending Date-Time for the historical price data returned.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data ending at the very latest data available for the symbol.
func (m HistoricalPriceDataRequestPointer) EndDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 40, 32))
}

// EndDateTime The ending Date-Time for the historical price data returned.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data ending at the very latest data available for the symbol.
func (m HistoricalPriceDataRequestPointerBuilder) EndDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 40, 32))
}

// MaxDaysToReturn specifies the maximum number of days of data the Server
// needs to return, if available, counting back from the latest Date-Time
// of data available for the symbol, or counting back from EndDateTime if
// it is set to a value other than 0.
//
// If MaxDaysToReturn is set to 0, then it is ignored by the Server.
func (m HistoricalPriceDataRequestPointer) MaxDaysToReturn() uint32 {
	return message.Uint32VLS(m.Ptr, 44, 40)
}

// MaxDaysToReturn specifies the maximum number of days of data the Server
// needs to return, if available, counting back from the latest Date-Time
// of data available for the symbol, or counting back from EndDateTime if
// it is set to a value other than 0.
//
// If MaxDaysToReturn is set to 0, then it is ignored by the Server.
func (m HistoricalPriceDataRequestPointerBuilder) MaxDaysToReturn() uint32 {
	return message.Uint32VLS(m.Ptr, 44, 40)
}

// UseZLibCompression Set this to 1, to request the Server use ZLib compression in the response
// when returning data. The Server can optionally ignore this if it does
// not support compression or does not want to use compression for any reason.
// not support compression or does not want to use compression for any reason.
//
// The HistoricalPriceDataResponseHeader will not be compressed. Only the
// historical price data records themselves.
//
// When receiving a batch of data on the network socket, give it to ZLib.
// It will respond with some uncompressed data and you put that into a buffer
// and then process as many complete historical data messages out of it that
// you can. Continue with this process until finished.
func (m HistoricalPriceDataRequestPointer) UseZLibCompression() uint8 {
	return message.Uint8VLS(m.Ptr, 45, 44)
}

// UseZLibCompression Set this to 1, to request the Server use ZLib compression in the response
// when returning data. The Server can optionally ignore this if it does
// not support compression or does not want to use compression for any reason.
// not support compression or does not want to use compression for any reason.
//
// The HistoricalPriceDataResponseHeader will not be compressed. Only the
// historical price data records themselves.
//
// When receiving a batch of data on the network socket, give it to ZLib.
// It will respond with some uncompressed data and you put that into a buffer
// and then process as many complete historical data messages out of it that
// you can. Continue with this process until finished.
func (m HistoricalPriceDataRequestPointerBuilder) UseZLibCompression() uint8 {
	return message.Uint8VLS(m.Ptr, 45, 44)
}

// RequestDividendAdjustedStockData In the case of a stock symbol, setting this to a value of 1 will request
// dividend adjusted data from the Server, if available. It is optional for
// the Server to support this.
func (m HistoricalPriceDataRequestPointer) RequestDividendAdjustedStockData() uint8 {
	return message.Uint8VLS(m.Ptr, 46, 45)
}

// RequestDividendAdjustedStockData In the case of a stock symbol, setting this to a value of 1 will request
// dividend adjusted data from the Server, if available. It is optional for
// the Server to support this.
func (m HistoricalPriceDataRequestPointerBuilder) RequestDividendAdjustedStockData() uint8 {
	return message.Uint8VLS(m.Ptr, 46, 45)
}

// Integer_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestPointer) Integer_1() uint16 {
	return message.Uint16VLS(m.Ptr, 48, 46)
}

// Integer_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestPointerBuilder) Integer_1() uint16 {
	return message.Uint16VLS(m.Ptr, 48, 46)
}

// SetRequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetSymbol The Symbol historical price data is requested for.
func (m *HistoricalPriceDataRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetExchange Optional: The exchange for the Symbol.
func (m *HistoricalPriceDataRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetRecordInterval The interval/timeframe of each record for the time range of the historical
// data requested. Can be any of:
//
// INTERVAL_TICK = 0
// INTERVAL_1_SECOND = 1
// INTERVAL_1_MINUTE = 60
// INTERVAL_1_DAY = 86400
// INTERVAL_1_WEEK = 604800
func (m HistoricalPriceDataRequestPointerBuilder) SetRecordInterval(value HistoricalDataIntervalEnum) {
	message.SetInt32VLS(m.Ptr, 24, 20, int32(value))
}

// SetStartDateTime The starting Date-Time for the historical price data returned, if available
// for the specified Symbol.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data starting at the earliest data available for the Symbol.
func (m HistoricalPriceDataRequestPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 32, 24, int64(value))
}

// SetEndDateTime The ending Date-Time for the historical price data returned.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data ending at the very latest data available for the symbol.
func (m HistoricalPriceDataRequestPointerBuilder) SetEndDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 40, 32, int64(value))
}

// SetMaxDaysToReturn MaxDaysToReturn specifies the maximum number of days of data the Server
// needs to return, if available, counting back from the latest Date-Time
// of data available for the symbol, or counting back from EndDateTime if
// it is set to a value other than 0.
//
// If MaxDaysToReturn is set to 0, then it is ignored by the Server.
func (m HistoricalPriceDataRequestPointerBuilder) SetMaxDaysToReturn(value uint32) {
	message.SetUint32VLS(m.Ptr, 44, 40, value)
}

// SetUseZLibCompression Set this to 1, to request the Server use ZLib compression in the response
// when returning data. The Server can optionally ignore this if it does
// not support compression or does not want to use compression for any reason.
// not support compression or does not want to use compression for any reason.
//
// The HistoricalPriceDataResponseHeader will not be compressed. Only the
// historical price data records themselves.
//
// When receiving a batch of data on the network socket, give it to ZLib.
// It will respond with some uncompressed data and you put that into a buffer
// and then process as many complete historical data messages out of it that
// you can. Continue with this process until finished.
func (m HistoricalPriceDataRequestPointerBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8VLS(m.Ptr, 45, 44, value)
}

// SetRequestDividendAdjustedStockData In the case of a stock symbol, setting this to a value of 1 will request
// dividend adjusted data from the Server, if available. It is optional for
// the Server to support this.
func (m HistoricalPriceDataRequestPointerBuilder) SetRequestDividendAdjustedStockData(value uint8) {
	message.SetUint8VLS(m.Ptr, 46, 45, value)
}

// SetInteger_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestPointerBuilder) SetInteger_1(value uint16) {
	message.SetUint16VLS(m.Ptr, 48, 46, value)
}

// Copy
func (m HistoricalPriceDataRequestPointer) Copy(to HistoricalPriceDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetRecordInterval(m.RecordInterval())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetMaxDaysToReturn(m.MaxDaysToReturn())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetRequestDividendAdjustedStockData(m.RequestDividendAdjustedStockData())
	to.SetInteger_1(m.Integer_1())
}

// Copy
func (m HistoricalPriceDataRequestPointerBuilder) Copy(to HistoricalPriceDataRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetRecordInterval(m.RecordInterval())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetMaxDaysToReturn(m.MaxDaysToReturn())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetRequestDividendAdjustedStockData(m.RequestDividendAdjustedStockData())
	to.SetInteger_1(m.Integer_1())
}

// CopyPointer
func (m HistoricalPriceDataRequestPointer) CopyPointer(to HistoricalPriceDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetRecordInterval(m.RecordInterval())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetMaxDaysToReturn(m.MaxDaysToReturn())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetRequestDividendAdjustedStockData(m.RequestDividendAdjustedStockData())
	to.SetInteger_1(m.Integer_1())
}

// CopyPointer
func (m HistoricalPriceDataRequestPointerBuilder) CopyPointer(to HistoricalPriceDataRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetRecordInterval(m.RecordInterval())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetMaxDaysToReturn(m.MaxDaysToReturn())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetRequestDividendAdjustedStockData(m.RequestDividendAdjustedStockData())
	to.SetInteger_1(m.Integer_1())
}

// CopyTo
func (m HistoricalPriceDataRequestPointer) CopyTo(to HistoricalPriceDataRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetRecordInterval(m.RecordInterval())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetMaxDaysToReturn(m.MaxDaysToReturn())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetRequestDividendAdjustedStockData(m.RequestDividendAdjustedStockData())
	to.SetInteger_1(m.Integer_1())
}

// CopyTo
func (m HistoricalPriceDataRequestPointerBuilder) CopyTo(to HistoricalPriceDataRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetRecordInterval(m.RecordInterval())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetMaxDaysToReturn(m.MaxDaysToReturn())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetRequestDividendAdjustedStockData(m.RequestDividendAdjustedStockData())
	to.SetInteger_1(m.Integer_1())
}

// CopyToPointer
func (m HistoricalPriceDataRequestPointer) CopyToPointer(to HistoricalPriceDataRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetRecordInterval(m.RecordInterval())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetMaxDaysToReturn(m.MaxDaysToReturn())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetRequestDividendAdjustedStockData(m.RequestDividendAdjustedStockData())
	to.SetInteger_1(m.Integer_1())
}

// CopyToPointer
func (m HistoricalPriceDataRequestPointerBuilder) CopyToPointer(to HistoricalPriceDataRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetRecordInterval(m.RecordInterval())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetMaxDaysToReturn(m.MaxDaysToReturn())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetRequestDividendAdjustedStockData(m.RequestDividendAdjustedStockData())
	to.SetInteger_1(m.Integer_1())
}
