package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalPriceDataRequestSize = 48

const HistoricalPriceDataRequestFixedSize = 120

// {
//     Size                             = HistoricalPriceDataRequestSize  (48)
//     Type                             = HISTORICAL_PRICE_DATA_REQUEST  (800)
//     BaseSize                         = HistoricalPriceDataRequestSize  (48)
//     RequestID                        = 0
//     Symbol                           = ""
//     Exchange                         = ""
//     RecordInterval                   = INTERVAL_TICK  (0)
//     StartDateTime                    = 0
//     EndDateTime                      = 0
//     MaxDaysToReturn                  = 0
//     UseZLibCompression               = 0
//     RequestDividendAdjustedStockData = 0
//     Integer_1                        = 0
// }
var _HistoricalPriceDataRequestDefault = []byte{48, 0, 32, 3, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size                             = HistoricalPriceDataRequestFixedSize  (120)
//     Type                             = HISTORICAL_PRICE_DATA_REQUEST  (800)
//     RequestID                        = 0
//     Symbol                           = ""
//     Exchange                         = ""
//     RecordInterval                   = INTERVAL_TICK  (0)
//     StartDateTime                    = 0
//     EndDateTime                      = 0
//     MaxDaysToReturn                  = 0
//     UseZLibCompression               = 0
//     RequestDividendAdjustedStockData = 0
//     Integer_1                        = 0
// }
var _HistoricalPriceDataRequestFixedDefault = []byte{120, 0, 32, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// HistoricalPriceDataRequest This is a message from the Client to the Server for requesting historical
// price data.
//
// This request can be on the same or a separate network socket connection
// compared to the streaming market data. This is going to be specified by
// the Server.
type HistoricalPriceDataRequest struct {
	message.VLS
}

// HistoricalPriceDataRequestBuilder This is a message from the Client to the Server for requesting historical
// price data.
//
// This request can be on the same or a separate network socket connection
// compared to the streaming market data. This is going to be specified by
// the Server.
type HistoricalPriceDataRequestBuilder struct {
	message.VLSBuilder
}

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

// HistoricalPriceDataRequestFixedPointer This is a message from the Client to the Server for requesting historical
// price data.
//
// This request can be on the same or a separate network socket connection
// compared to the streaming market data. This is going to be specified by
// the Server.
type HistoricalPriceDataRequestFixedPointer struct {
	message.FixedPointer
}

// HistoricalPriceDataRequestFixedPointerBuilder This is a message from the Client to the Server for requesting historical
// price data.
//
// This request can be on the same or a separate network socket connection
// compared to the streaming market data. This is going to be specified by
// the Server.
type HistoricalPriceDataRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalPriceDataRequestFrom(b []byte) HistoricalPriceDataRequest {
	return HistoricalPriceDataRequest{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalPriceDataRequest(b []byte) HistoricalPriceDataRequest {
	return HistoricalPriceDataRequest{VLS: message.WrapVLS(b)}
}

func NewHistoricalPriceDataRequest() HistoricalPriceDataRequestBuilder {
	a := HistoricalPriceDataRequestBuilder{message.NewVLSBuilder(48)}
	a.Unsafe().SetBytes(0, _HistoricalPriceDataRequestDefault)
	return a
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

func AllocHistoricalPriceDataRequest() HistoricalPriceDataRequestPointerBuilder {
	a := HistoricalPriceDataRequestPointerBuilder{message.AllocVLSBuilder(48)}
	a.Ptr.SetBytes(0, _HistoricalPriceDataRequestDefault)
	return a
}

func AllocHistoricalPriceDataRequestFrom(b []byte) HistoricalPriceDataRequestPointer {
	return HistoricalPriceDataRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocHistoricalPriceDataRequestFixed() HistoricalPriceDataRequestFixedPointerBuilder {
	a := HistoricalPriceDataRequestFixedPointerBuilder{message.AllocFixed(120)}
	a.Ptr.SetBytes(0, _HistoricalPriceDataRequestFixedDefault)
	return a
}

func AllocHistoricalPriceDataRequestFixedFrom(b []byte) HistoricalPriceDataRequestFixedPointer {
	return HistoricalPriceDataRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                             = HistoricalPriceDataRequestSize  (48)
//     Type                             = HISTORICAL_PRICE_DATA_REQUEST  (800)
//     BaseSize                         = HistoricalPriceDataRequestSize  (48)
//     RequestID                        = 0
//     Symbol                           = ""
//     Exchange                         = ""
//     RecordInterval                   = INTERVAL_TICK  (0)
//     StartDateTime                    = 0
//     EndDateTime                      = 0
//     MaxDaysToReturn                  = 0
//     UseZLibCompression               = 0
//     RequestDividendAdjustedStockData = 0
//     Integer_1                        = 0
// }
func (m HistoricalPriceDataRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalPriceDataRequestDefault)
}

// Clear
// {
//     Size                             = HistoricalPriceDataRequestFixedSize  (120)
//     Type                             = HISTORICAL_PRICE_DATA_REQUEST  (800)
//     RequestID                        = 0
//     Symbol                           = ""
//     Exchange                         = ""
//     RecordInterval                   = INTERVAL_TICK  (0)
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

// Clear
// {
//     Size                             = HistoricalPriceDataRequestSize  (48)
//     Type                             = HISTORICAL_PRICE_DATA_REQUEST  (800)
//     BaseSize                         = HistoricalPriceDataRequestSize  (48)
//     RequestID                        = 0
//     Symbol                           = ""
//     Exchange                         = ""
//     RecordInterval                   = INTERVAL_TICK  (0)
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

// Clear
// {
//     Size                             = HistoricalPriceDataRequestFixedSize  (120)
//     Type                             = HISTORICAL_PRICE_DATA_REQUEST  (800)
//     RequestID                        = 0
//     Symbol                           = ""
//     Exchange                         = ""
//     RecordInterval                   = INTERVAL_TICK  (0)
//     StartDateTime                    = 0
//     EndDateTime                      = 0
//     MaxDaysToReturn                  = 0
//     UseZLibCompression               = 0
//     RequestDividendAdjustedStockData = 0
//     Integer_1                        = 0
// }
func (m HistoricalPriceDataRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalPriceDataRequestFixedDefault)
}

// ToBuilder
func (m HistoricalPriceDataRequest) ToBuilder() HistoricalPriceDataRequestBuilder {
	return HistoricalPriceDataRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalPriceDataRequestFixed) ToBuilder() HistoricalPriceDataRequestFixedBuilder {
	return HistoricalPriceDataRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalPriceDataRequestPointer) ToBuilder() HistoricalPriceDataRequestPointerBuilder {
	return HistoricalPriceDataRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalPriceDataRequestFixedPointer) ToBuilder() HistoricalPriceDataRequestFixedPointerBuilder {
	return HistoricalPriceDataRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalPriceDataRequestBuilder) Finish() HistoricalPriceDataRequest {
	return HistoricalPriceDataRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalPriceDataRequestFixedBuilder) Finish() HistoricalPriceDataRequestFixed {
	return HistoricalPriceDataRequestFixed{m.Fixed}
}

// Finish
func (m *HistoricalPriceDataRequestPointerBuilder) Finish() HistoricalPriceDataRequestPointer {
	return HistoricalPriceDataRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalPriceDataRequestFixedPointerBuilder) Finish() HistoricalPriceDataRequestFixedPointer {
	return HistoricalPriceDataRequestFixedPointer{m.FixedPointer}
}

// RequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
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
func (m HistoricalPriceDataRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Symbol The Symbol historical price data is requested for.
func (m HistoricalPriceDataRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
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
func (m HistoricalPriceDataRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange Optional: The exchange for the Symbol.
func (m HistoricalPriceDataRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
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
func (m HistoricalPriceDataRequest) RecordInterval() HistoricalDataIntervalEnum {
	return HistoricalDataIntervalEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// RecordInterval The interval/timeframe of each record for the time range of the historical
// data requested. Can be any of:
//
// INTERVAL_TICK = 0
// INTERVAL_1_SECOND = 1
// INTERVAL_1_MINUTE = 60
// INTERVAL_1_DAY = 86400
// INTERVAL_1_WEEK = 604800
func (m HistoricalPriceDataRequestBuilder) RecordInterval() HistoricalDataIntervalEnum {
	return HistoricalDataIntervalEnum(message.Int32VLS(m.Unsafe(), 24, 20))
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
func (m HistoricalPriceDataRequest) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 32, 24))
}

// StartDateTime The starting Date-Time for the historical price data returned, if available
// for the specified Symbol.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data starting at the earliest data available for the Symbol.
func (m HistoricalPriceDataRequestBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 32, 24))
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
func (m HistoricalPriceDataRequest) EndDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 40, 32))
}

// EndDateTime The ending Date-Time for the historical price data returned.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data ending at the very latest data available for the symbol.
func (m HistoricalPriceDataRequestBuilder) EndDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 40, 32))
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
func (m HistoricalPriceDataRequest) MaxDaysToReturn() uint32 {
	return message.Uint32VLS(m.Unsafe(), 44, 40)
}

// MaxDaysToReturn specifies the maximum number of days of data the Server
// needs to return, if available, counting back from the latest Date-Time
// of data available for the symbol, or counting back from EndDateTime if
// it is set to a value other than 0.
//
// If MaxDaysToReturn is set to 0, then it is ignored by the Server.
func (m HistoricalPriceDataRequestBuilder) MaxDaysToReturn() uint32 {
	return message.Uint32VLS(m.Unsafe(), 44, 40)
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
func (m HistoricalPriceDataRequest) UseZLibCompression() uint8 {
	return message.Uint8VLS(m.Unsafe(), 45, 44)
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
func (m HistoricalPriceDataRequestBuilder) UseZLibCompression() uint8 {
	return message.Uint8VLS(m.Unsafe(), 45, 44)
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
func (m HistoricalPriceDataRequest) RequestDividendAdjustedStockData() uint8 {
	return message.Uint8VLS(m.Unsafe(), 46, 45)
}

// RequestDividendAdjustedStockData In the case of a stock symbol, setting this to a value of 1 will request
// dividend adjusted data from the Server, if available. It is optional for
// the Server to support this.
func (m HistoricalPriceDataRequestBuilder) RequestDividendAdjustedStockData() uint8 {
	return message.Uint8VLS(m.Unsafe(), 46, 45)
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
func (m HistoricalPriceDataRequest) Integer_1() uint16 {
	return message.Uint16VLS(m.Unsafe(), 48, 46)
}

// Integer_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestBuilder) Integer_1() uint16 {
	return message.Uint16VLS(m.Unsafe(), 48, 46)
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

// RequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Symbol The Symbol historical price data is requested for.
func (m HistoricalPriceDataRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol The Symbol historical price data is requested for.
func (m HistoricalPriceDataRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol The Symbol historical price data is requested for.
func (m HistoricalPriceDataRequestFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Symbol The Symbol historical price data is requested for.
func (m HistoricalPriceDataRequestFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Exchange Optional: The exchange for the Symbol.
func (m HistoricalPriceDataRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// Exchange Optional: The exchange for the Symbol.
func (m HistoricalPriceDataRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// Exchange Optional: The exchange for the Symbol.
func (m HistoricalPriceDataRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 88, 72)
}

// Exchange Optional: The exchange for the Symbol.
func (m HistoricalPriceDataRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 88, 72)
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

// RecordInterval The interval/timeframe of each record for the time range of the historical
// data requested. Can be any of:
//
// INTERVAL_TICK = 0
// INTERVAL_1_SECOND = 1
// INTERVAL_1_MINUTE = 60
// INTERVAL_1_DAY = 86400
// INTERVAL_1_WEEK = 604800
func (m HistoricalPriceDataRequestFixedPointer) RecordInterval() HistoricalDataIntervalEnum {
	return HistoricalDataIntervalEnum(message.Int32Fixed(m.Ptr, 92, 88))
}

// RecordInterval The interval/timeframe of each record for the time range of the historical
// data requested. Can be any of:
//
// INTERVAL_TICK = 0
// INTERVAL_1_SECOND = 1
// INTERVAL_1_MINUTE = 60
// INTERVAL_1_DAY = 86400
// INTERVAL_1_WEEK = 604800
func (m HistoricalPriceDataRequestFixedPointerBuilder) RecordInterval() HistoricalDataIntervalEnum {
	return HistoricalDataIntervalEnum(message.Int32Fixed(m.Ptr, 92, 88))
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

// StartDateTime The starting Date-Time for the historical price data returned, if available
// for the specified Symbol.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data starting at the earliest data available for the Symbol.
func (m HistoricalPriceDataRequestFixedPointer) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 104, 96))
}

// StartDateTime The starting Date-Time for the historical price data returned, if available
// for the specified Symbol.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data starting at the earliest data available for the Symbol.
func (m HistoricalPriceDataRequestFixedPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 104, 96))
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

// EndDateTime The ending Date-Time for the historical price data returned.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data ending at the very latest data available for the symbol.
func (m HistoricalPriceDataRequestFixedPointer) EndDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 112, 104))
}

// EndDateTime The ending Date-Time for the historical price data returned.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data ending at the very latest data available for the symbol.
func (m HistoricalPriceDataRequestFixedPointerBuilder) EndDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 112, 104))
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

// MaxDaysToReturn specifies the maximum number of days of data the Server
// needs to return, if available, counting back from the latest Date-Time
// of data available for the symbol, or counting back from EndDateTime if
// it is set to a value other than 0.
//
// If MaxDaysToReturn is set to 0, then it is ignored by the Server.
func (m HistoricalPriceDataRequestFixedPointer) MaxDaysToReturn() uint32 {
	return message.Uint32Fixed(m.Ptr, 116, 112)
}

// MaxDaysToReturn specifies the maximum number of days of data the Server
// needs to return, if available, counting back from the latest Date-Time
// of data available for the symbol, or counting back from EndDateTime if
// it is set to a value other than 0.
//
// If MaxDaysToReturn is set to 0, then it is ignored by the Server.
func (m HistoricalPriceDataRequestFixedPointerBuilder) MaxDaysToReturn() uint32 {
	return message.Uint32Fixed(m.Ptr, 116, 112)
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
func (m HistoricalPriceDataRequestFixedPointer) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Ptr, 117, 116)
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
func (m HistoricalPriceDataRequestFixedPointerBuilder) UseZLibCompression() uint8 {
	return message.Uint8Fixed(m.Ptr, 117, 116)
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

// RequestDividendAdjustedStockData In the case of a stock symbol, setting this to a value of 1 will request
// dividend adjusted data from the Server, if available. It is optional for
// the Server to support this.
func (m HistoricalPriceDataRequestFixedPointer) RequestDividendAdjustedStockData() uint8 {
	return message.Uint8Fixed(m.Ptr, 118, 117)
}

// RequestDividendAdjustedStockData In the case of a stock symbol, setting this to a value of 1 will request
// dividend adjusted data from the Server, if available. It is optional for
// the Server to support this.
func (m HistoricalPriceDataRequestFixedPointerBuilder) RequestDividendAdjustedStockData() uint8 {
	return message.Uint8Fixed(m.Ptr, 118, 117)
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

// Integer_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestFixedPointer) Integer_1() uint16 {
	return message.Uint16Fixed(m.Ptr, 120, 118)
}

// Integer_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestFixedPointerBuilder) Integer_1() uint16 {
	return message.Uint16Fixed(m.Ptr, 120, 118)
}

// SetRequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
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
func (m *HistoricalPriceDataRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetSymbol The Symbol historical price data is requested for.
func (m *HistoricalPriceDataRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetExchange Optional: The exchange for the Symbol.
func (m *HistoricalPriceDataRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
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
func (m HistoricalPriceDataRequestBuilder) SetRecordInterval(value HistoricalDataIntervalEnum) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, int32(value))
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
func (m HistoricalPriceDataRequestBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 32, 24, int64(value))
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
func (m HistoricalPriceDataRequestBuilder) SetEndDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 40, 32, int64(value))
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
func (m HistoricalPriceDataRequestBuilder) SetMaxDaysToReturn(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 44, 40, value)
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
func (m HistoricalPriceDataRequestBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 45, 44, value)
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
func (m HistoricalPriceDataRequestBuilder) SetRequestDividendAdjustedStockData(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 46, 45, value)
}

// SetRequestDividendAdjustedStockData In the case of a stock symbol, setting this to a value of 1 will request
// dividend adjusted data from the Server, if available. It is optional for
// the Server to support this.
func (m HistoricalPriceDataRequestPointerBuilder) SetRequestDividendAdjustedStockData(value uint8) {
	message.SetUint8VLS(m.Ptr, 46, 45, value)
}

// SetInteger_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestBuilder) SetInteger_1(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 48, 46, value)
}

// SetInteger_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestPointerBuilder) SetInteger_1(value uint16) {
	message.SetUint16VLS(m.Ptr, 48, 46, value)
}

// SetRequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID Unique integer identifier to identify this request. The historical price
// data response messages from the Server will contain this identifier so
// they can be matched up with the request from the Client. This identifier
// only needs to be unique to the historical price data messages. It can
// conflict with identifiers used with other classes of messages.
func (m HistoricalPriceDataRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetSymbol The Symbol historical price data is requested for.
func (m HistoricalPriceDataRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetSymbol The Symbol historical price data is requested for.
func (m HistoricalPriceDataRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 72, 8, value)
}

// SetExchange Optional: The exchange for the Symbol.
func (m HistoricalPriceDataRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 88, 72, value)
}

// SetExchange Optional: The exchange for the Symbol.
func (m HistoricalPriceDataRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 88, 72, value)
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

// SetRecordInterval The interval/timeframe of each record for the time range of the historical
// data requested. Can be any of:
//
// INTERVAL_TICK = 0
// INTERVAL_1_SECOND = 1
// INTERVAL_1_MINUTE = 60
// INTERVAL_1_DAY = 86400
// INTERVAL_1_WEEK = 604800
func (m HistoricalPriceDataRequestFixedPointerBuilder) SetRecordInterval(value HistoricalDataIntervalEnum) {
	message.SetInt32Fixed(m.Ptr, 92, 88, int32(value))
}

// SetStartDateTime The starting Date-Time for the historical price data returned, if available
// for the specified Symbol.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data starting at the earliest data available for the Symbol.
func (m HistoricalPriceDataRequestFixedBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 104, 96, int64(value))
}

// SetStartDateTime The starting Date-Time for the historical price data returned, if available
// for the specified Symbol.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data starting at the earliest data available for the Symbol.
func (m HistoricalPriceDataRequestFixedPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 104, 96, int64(value))
}

// SetEndDateTime The ending Date-Time for the historical price data returned.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data ending at the very latest data available for the symbol.
func (m HistoricalPriceDataRequestFixedBuilder) SetEndDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 112, 104, int64(value))
}

// SetEndDateTime The ending Date-Time for the historical price data returned.
//
// If it is not set or set to 0, then this is a request to the Server to
// return data ending at the very latest data available for the symbol.
func (m HistoricalPriceDataRequestFixedPointerBuilder) SetEndDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 112, 104, int64(value))
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

// SetMaxDaysToReturn MaxDaysToReturn specifies the maximum number of days of data the Server
// needs to return, if available, counting back from the latest Date-Time
// of data available for the symbol, or counting back from EndDateTime if
// it is set to a value other than 0.
//
// If MaxDaysToReturn is set to 0, then it is ignored by the Server.
func (m HistoricalPriceDataRequestFixedPointerBuilder) SetMaxDaysToReturn(value uint32) {
	message.SetUint32Fixed(m.Ptr, 116, 112, value)
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
func (m HistoricalPriceDataRequestFixedPointerBuilder) SetUseZLibCompression(value uint8) {
	message.SetUint8Fixed(m.Ptr, 117, 116, value)
}

// SetRequestDividendAdjustedStockData In the case of a stock symbol, setting this to a value of 1 will request
// dividend adjusted data from the Server, if available. It is optional for
// the Server to support this.
func (m HistoricalPriceDataRequestFixedBuilder) SetRequestDividendAdjustedStockData(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 118, 117, value)
}

// SetRequestDividendAdjustedStockData In the case of a stock symbol, setting this to a value of 1 will request
// dividend adjusted data from the Server, if available. It is optional for
// the Server to support this.
func (m HistoricalPriceDataRequestFixedPointerBuilder) SetRequestDividendAdjustedStockData(value uint8) {
	message.SetUint8Fixed(m.Ptr, 118, 117, value)
}

// SetInteger_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestFixedBuilder) SetInteger_1(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 120, 118, value)
}

// SetInteger_1 A general purpose 2 byte flag field from the Client to the Server which
// can be used for any special purpose the Client and Server require.
func (m HistoricalPriceDataRequestFixedPointerBuilder) SetInteger_1(value uint16) {
	message.SetUint16Fixed(m.Ptr, 120, 118, value)
}

// Copy
func (m HistoricalPriceDataRequest) Copy(to HistoricalPriceDataRequestBuilder) {
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
func (m HistoricalPriceDataRequestBuilder) Copy(to HistoricalPriceDataRequestBuilder) {
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
func (m HistoricalPriceDataRequest) CopyPointer(to HistoricalPriceDataRequestPointerBuilder) {
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
func (m HistoricalPriceDataRequestBuilder) CopyPointer(to HistoricalPriceDataRequestPointerBuilder) {
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
func (m HistoricalPriceDataRequest) CopyToPointer(to HistoricalPriceDataRequestFixedPointerBuilder) {
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
func (m HistoricalPriceDataRequestBuilder) CopyToPointer(to HistoricalPriceDataRequestFixedPointerBuilder) {
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
func (m HistoricalPriceDataRequest) CopyTo(to HistoricalPriceDataRequestFixedBuilder) {
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
func (m HistoricalPriceDataRequestBuilder) CopyTo(to HistoricalPriceDataRequestFixedBuilder) {
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

// Copy
func (m HistoricalPriceDataRequestFixedPointer) Copy(to HistoricalPriceDataRequestFixedBuilder) {
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
func (m HistoricalPriceDataRequestFixedPointerBuilder) Copy(to HistoricalPriceDataRequestFixedBuilder) {
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
func (m HistoricalPriceDataRequestFixedPointer) CopyPointer(to HistoricalPriceDataRequestFixedPointerBuilder) {
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
func (m HistoricalPriceDataRequestFixedPointerBuilder) CopyPointer(to HistoricalPriceDataRequestFixedPointerBuilder) {
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
func (m HistoricalPriceDataRequestFixedPointer) CopyTo(to HistoricalPriceDataRequestBuilder) {
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
func (m HistoricalPriceDataRequestFixedPointerBuilder) CopyTo(to HistoricalPriceDataRequestBuilder) {
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
func (m HistoricalPriceDataRequestFixedPointer) CopyToPointer(to HistoricalPriceDataRequestPointerBuilder) {
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
func (m HistoricalPriceDataRequestFixedPointerBuilder) CopyToPointer(to HistoricalPriceDataRequestPointerBuilder) {
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
