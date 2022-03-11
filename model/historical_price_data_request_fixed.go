package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                             = HistoricalPriceDataRequestFixedSize  (120)
//     Type                             = HISTORICAL_PRICE_DATA_REQUEST  (800)
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
var _HistoricalPriceDataRequestFixedDefault = []byte{120, 0, 32, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalPriceDataRequestFixedSize = 120

// HistoricalPriceDataRequestFixed This is a message from the Client to the Server for requesting historical
// price data.
//
// This request can be on the same or a separate network socket connection
// compared to the streaming market data. This is going to be specified by
// the Server.
type HistoricalPriceDataRequestFixed struct {
	message.Fixed
}

// HistoricalPriceDataRequestFixedBuilder This is a message from the Client to the Server for requesting historical
// price data.
//
// This request can be on the same or a separate network socket connection
// compared to the streaming market data. This is going to be specified by
// the Server.
type HistoricalPriceDataRequestFixedBuilder struct {
	message.Fixed
}

func NewHistoricalPriceDataRequestFixedFrom(b []byte) HistoricalPriceDataRequestFixed {
	return HistoricalPriceDataRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalPriceDataRequestFixed(b []byte) HistoricalPriceDataRequestFixed {
	return HistoricalPriceDataRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalPriceDataRequestFixed() HistoricalPriceDataRequestFixedBuilder {
	a := HistoricalPriceDataRequestFixedBuilder{message.NewFixed(120)}
	a.Unsafe().SetBytes(0, _HistoricalPriceDataRequestFixedDefault)
	return a
}

// Clear
// {
//     Size                             = HistoricalPriceDataRequestFixedSize  (120)
//     Type                             = HISTORICAL_PRICE_DATA_REQUEST  (800)
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
func (m HistoricalPriceDataRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataRequestFixedDefault)
}

// ToBuilder
func (m HistoricalPriceDataRequestFixed) ToBuilder() HistoricalPriceDataRequestFixedBuilder {
	return HistoricalPriceDataRequestFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalPriceDataRequestFixedBuilder) Finish() HistoricalPriceDataRequestFixed {
	return HistoricalPriceDataRequestFixed{m.Fixed}
}

// RequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// Symbol The Symbol historical price data is requested for.
func (m HistoricalPriceDataRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol The Symbol historical price data is requested for.
func (m HistoricalPriceDataRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Exchange Optional: The exchange for the Symbol.
func (m HistoricalPriceDataRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// Exchange Optional: The exchange for the Symbol.
func (m HistoricalPriceDataRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// RecordInterval The interval/timeframe of each record for the time range of the historical
// data requested. Can be any of:
//
// INTERVAL_TICK = 0
// INTERVAL_1_SECOND = 1
// INTERVAL_1_MINUTE = 60
// INTERVAL_1_DAY = 86400
// INTERVAL_1_WEEK = 604800
func (m HistoricalPriceDataRequestFixed) RecordInterval() HistoricalDataIntervalEnum {
	return HistoricalDataIntervalEnum(message.Int32Fixed(m.Unsafe(), 92, 88))
}

// RecordInterval The interval/timeframe of each record for the time range of the historical
// data requested. Can be any of:
//
// INTERVAL_TICK = 0
// INTERVAL_1_SECOND = 1
// INTERVAL_1_MINUTE = 60
// INTERVAL_1_DAY = 86400
// INTERVAL_1_WEEK = 604800
func (m HistoricalPriceDataRequestFixedBuilder) RecordInterval() HistoricalDataIntervalEnum {
	return HistoricalDataIntervalEnum(message.Int32Fixed(m.Unsafe(), 92, 88))
}

// StartDateTime The starting Date-Time for the historical price data returned, if available
// for the specified Symbol.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data starting at the earliest data available for the Symbol.
func (m HistoricalPriceDataRequestFixed) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 104, 96))
}

// StartDateTime The starting Date-Time for the historical price data returned, if available
// for the specified Symbol.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data starting at the earliest data available for the Symbol.
func (m HistoricalPriceDataRequestFixedBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 104, 96))
}

// EndDateTime The ending Date-Time for the historical price data returned.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data ending at the very latest data available for the symbol.
func (m HistoricalPriceDataRequestFixed) EndDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 112, 104))
}

// EndDateTime The ending Date-Time for the historical price data returned.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data ending at the very latest data available for the symbol.
func (m HistoricalPriceDataRequestFixedBuilder) EndDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 112, 104))
}

// MaxDaysToReturn specifies the maximum number of days of data the Server
// needs to return, if available, counting back from the latest Date-Time
// of data available for the symbol, or counting back from EndDateTime if
// it is set to a value other than 0.
//
// If MaxDaysToReturn is set to 0, then it is ignored by the Server.
func (m HistoricalPriceDataRequestFixed) MaxDaysToReturn() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 116, 112)
}

// MaxDaysToReturn specifies the maximum number of days of data the Server
// needs to return, if available, counting back from the latest Date-Time
// of data available for the symbol, or counting back from EndDateTime if
// it is set to a value other than 0.
//
// If MaxDaysToReturn is set to 0, then it is ignored by the Server.
func (m HistoricalPriceDataRequestFixedBuilder) MaxDaysToReturn() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 116, 112)
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
func (m HistoricalPriceDataRequestFixed) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 117, 116)
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
func (m HistoricalPriceDataRequestFixedBuilder) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 117, 116)
}

// RequestDividendAdjustedStockData In the case of a stock symbol, setting this to a value of 1 will request
// dividend adjusted data from the Server, if available. It is optional for
// the Server to support this.
func (m HistoricalPriceDataRequestFixed) RequestDividendAdjustedStockData() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 118, 117)
}

// RequestDividendAdjustedStockData In the case of a stock symbol, setting this to a value of 1 will request
// dividend adjusted data from the Server, if available. It is optional for
// the Server to support this.
func (m HistoricalPriceDataRequestFixedBuilder) RequestDividendAdjustedStockData() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 118, 117)
}

// Integer_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestFixed) Integer_1() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 120, 118)
}

// Integer_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestFixedBuilder) Integer_1() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 120, 118)
}

// SetRequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbol The Symbol historical price data is requested for.
func (m HistoricalPriceDataRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetExchange Optional: The exchange for the Symbol.
func (m HistoricalPriceDataRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 88, 72, value)
}

// SetRecordInterval The interval/timeframe of each record for the time range of the historical
// data requested. Can be any of:
//
// INTERVAL_TICK = 0
// INTERVAL_1_SECOND = 1
// INTERVAL_1_MINUTE = 60
// INTERVAL_1_DAY = 86400
// INTERVAL_1_WEEK = 604800
func (m HistoricalPriceDataRequestFixedBuilder) SetRecordInterval(value HistoricalDataIntervalEnum) {
	message.SetInt32Fixed(m.Unsafe(), 92, 88, int32(value))
}

// SetStartDateTime The starting Date-Time for the historical price data returned, if available
// for the specified Symbol.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data starting at the earliest data available for the Symbol.
func (m HistoricalPriceDataRequestFixedBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 104, 96, int64(value))
}

// SetEndDateTime The ending Date-Time for the historical price data returned.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data ending at the very latest data available for the symbol.
func (m HistoricalPriceDataRequestFixedBuilder) SetEndDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 112, 104, int64(value))
}

// SetMaxDaysToReturn MaxDaysToReturn specifies the maximum number of days of data the Server
// needs to return, if available, counting back from the latest Date-Time
// of data available for the symbol, or counting back from EndDateTime if
// it is set to a value other than 0.
//
// If MaxDaysToReturn is set to 0, then it is ignored by the Server.
func (m HistoricalPriceDataRequestFixedBuilder) SetMaxDaysToReturn(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 116, 112, value)
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
func (m HistoricalPriceDataRequestFixedBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 117, 116, value)
}

// SetRequestDividendAdjustedStockData In the case of a stock symbol, setting this to a value of 1 will request
// dividend adjusted data from the Server, if available. It is optional for
// the Server to support this.
func (m HistoricalPriceDataRequestFixedBuilder) SetRequestDividendAdjustedStockData(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 118, 117, value)
}

// SetInteger_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestFixedBuilder) SetInteger_1(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 120, 118, value)
}

// Copy
func (m HistoricalPriceDataRequestFixed) Copy(to HistoricalPriceDataRequestFixedBuilder) {
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
func (m HistoricalPriceDataRequestFixedBuilder) Copy(to HistoricalPriceDataRequestFixedBuilder) {
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
func (m HistoricalPriceDataRequestFixed) CopyPointer(to HistoricalPriceDataRequestFixedPointerBuilder) {
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
func (m HistoricalPriceDataRequestFixedBuilder) CopyPointer(to HistoricalPriceDataRequestFixedPointerBuilder) {
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
func (m HistoricalPriceDataRequestFixed) CopyToPointer(to HistoricalPriceDataRequestPointerBuilder) {
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
func (m HistoricalPriceDataRequestFixedBuilder) CopyToPointer(to HistoricalPriceDataRequestPointerBuilder) {
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
func (m HistoricalPriceDataRequestFixed) CopyTo(to HistoricalPriceDataRequestBuilder) {
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
func (m HistoricalPriceDataRequestFixedBuilder) CopyTo(to HistoricalPriceDataRequestBuilder) {
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
