package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                      = MarketDataSnapshot_IntSize  (96)
//     Type                      = MARKET_DATA_SNAPSHOT_INT  (125)
//     SymbolID                  = 0
//     SessionSettlementPrice    = 32767
//     SessionOpenPrice          = 32767
//     SessionHighPrice          = 32767
//     SessionLowPrice           = 32767
//     SessionVolume             = 32767
//     SessionNumTrades          = 65535
//     OpenInterest              = 65535
//     BidPrice                  = 32767
//     AskPrice                  = 32767
//     AskQuantity               = 32767
//     BidQuantity               = 32767
//     LastTradePrice            = 32767
//     LastTradeVolume           = 32767
//     LastTradeDateTime         = 0
//     BidAskDateTime            = 0
//     SessionSettlementDateTime = 0
//     TradingSessionDate        = 0
//     TradingStatus             = 0
// }
var _MarketDataSnapshot_IntDefault = []byte{96, 0, 125, 0, 0, 0, 0, 0, 255, 127, 0, 0, 255, 127, 0, 0, 255, 127, 0, 0, 255, 127, 0, 0, 255, 127, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 127, 0, 0, 255, 127, 0, 0, 255, 127, 0, 0, 255, 127, 0, 0, 255, 127, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataSnapshot_IntSize = 96

type MarketDataSnapshot_Int struct {
	message.Fixed
}

type MarketDataSnapshot_IntBuilder struct {
	message.Fixed
}

func NewMarketDataSnapshot_IntFrom(b []byte) MarketDataSnapshot_Int {
	return MarketDataSnapshot_Int{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataSnapshot_Int(b []byte) MarketDataSnapshot_Int {
	return MarketDataSnapshot_Int{Fixed: message.WrapFixed(b)}
}

func NewMarketDataSnapshot_Int() MarketDataSnapshot_IntBuilder {
	a := MarketDataSnapshot_IntBuilder{message.NewFixed(96)}
	a.Unsafe().SetBytes(0, _MarketDataSnapshot_IntDefault)
	return a
}

// Clear
// {
//     Size                      = MarketDataSnapshot_IntSize  (96)
//     Type                      = MARKET_DATA_SNAPSHOT_INT  (125)
//     SymbolID                  = 0
//     SessionSettlementPrice    = 32767
//     SessionOpenPrice          = 32767
//     SessionHighPrice          = 32767
//     SessionLowPrice           = 32767
//     SessionVolume             = 32767
//     SessionNumTrades          = 65535
//     OpenInterest              = 65535
//     BidPrice                  = 32767
//     AskPrice                  = 32767
//     AskQuantity               = 32767
//     BidQuantity               = 32767
//     LastTradePrice            = 32767
//     LastTradeVolume           = 32767
//     LastTradeDateTime         = 0
//     BidAskDateTime            = 0
//     SessionSettlementDateTime = 0
//     TradingSessionDate        = 0
//     TradingStatus             = 0
// }
func (m MarketDataSnapshot_IntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataSnapshot_IntDefault)
}

// ToBuilder
func (m MarketDataSnapshot_Int) ToBuilder() MarketDataSnapshot_IntBuilder {
	return MarketDataSnapshot_IntBuilder{m.Fixed}
}

// Finish
func (m MarketDataSnapshot_IntBuilder) Finish() MarketDataSnapshot_Int {
	return MarketDataSnapshot_Int{m.Fixed}
}

// SymbolID
func (m MarketDataSnapshot_Int) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m MarketDataSnapshot_IntBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SessionSettlementPrice
func (m MarketDataSnapshot_Int) SessionSettlementPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// SessionSettlementPrice
func (m MarketDataSnapshot_IntBuilder) SessionSettlementPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// SessionOpenPrice
func (m MarketDataSnapshot_Int) SessionOpenPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// SessionOpenPrice
func (m MarketDataSnapshot_IntBuilder) SessionOpenPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// SessionHighPrice
func (m MarketDataSnapshot_Int) SessionHighPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// SessionHighPrice
func (m MarketDataSnapshot_IntBuilder) SessionHighPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 20, 16)
}

// SessionLowPrice
func (m MarketDataSnapshot_Int) SessionLowPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 24, 20)
}

// SessionLowPrice
func (m MarketDataSnapshot_IntBuilder) SessionLowPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 24, 20)
}

// SessionVolume
func (m MarketDataSnapshot_Int) SessionVolume() int32 {
	return message.Int32Fixed(m.Unsafe(), 28, 24)
}

// SessionVolume
func (m MarketDataSnapshot_IntBuilder) SessionVolume() int32 {
	return message.Int32Fixed(m.Unsafe(), 28, 24)
}

// SessionNumTrades
func (m MarketDataSnapshot_Int) SessionNumTrades() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 32, 28)
}

// SessionNumTrades
func (m MarketDataSnapshot_IntBuilder) SessionNumTrades() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 32, 28)
}

// OpenInterest
func (m MarketDataSnapshot_Int) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 36, 32)
}

// OpenInterest
func (m MarketDataSnapshot_IntBuilder) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 36, 32)
}

// BidPrice
func (m MarketDataSnapshot_Int) BidPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 40, 36)
}

// BidPrice
func (m MarketDataSnapshot_IntBuilder) BidPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 40, 36)
}

// AskPrice
func (m MarketDataSnapshot_Int) AskPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 44, 40)
}

// AskPrice
func (m MarketDataSnapshot_IntBuilder) AskPrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 44, 40)
}

// AskQuantity
func (m MarketDataSnapshot_Int) AskQuantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 48, 44)
}

// AskQuantity
func (m MarketDataSnapshot_IntBuilder) AskQuantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 48, 44)
}

// BidQuantity
func (m MarketDataSnapshot_Int) BidQuantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 52, 48)
}

// BidQuantity
func (m MarketDataSnapshot_IntBuilder) BidQuantity() int32 {
	return message.Int32Fixed(m.Unsafe(), 52, 48)
}

// LastTradePrice
func (m MarketDataSnapshot_Int) LastTradePrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 56, 52)
}

// LastTradePrice
func (m MarketDataSnapshot_IntBuilder) LastTradePrice() int32 {
	return message.Int32Fixed(m.Unsafe(), 56, 52)
}

// LastTradeVolume
func (m MarketDataSnapshot_Int) LastTradeVolume() int32 {
	return message.Int32Fixed(m.Unsafe(), 60, 56)
}

// LastTradeVolume
func (m MarketDataSnapshot_IntBuilder) LastTradeVolume() int32 {
	return message.Int32Fixed(m.Unsafe(), 60, 56)
}

// LastTradeDateTime
func (m MarketDataSnapshot_Int) LastTradeDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 72, 64))
}

// LastTradeDateTime
func (m MarketDataSnapshot_IntBuilder) LastTradeDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 72, 64))
}

// BidAskDateTime
func (m MarketDataSnapshot_Int) BidAskDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 80, 72))
}

// BidAskDateTime
func (m MarketDataSnapshot_IntBuilder) BidAskDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 80, 72))
}

// SessionSettlementDateTime
func (m MarketDataSnapshot_Int) SessionSettlementDateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 84, 80))
}

// SessionSettlementDateTime
func (m MarketDataSnapshot_IntBuilder) SessionSettlementDateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 84, 80))
}

// TradingSessionDate
func (m MarketDataSnapshot_Int) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 88, 84))
}

// TradingSessionDate
func (m MarketDataSnapshot_IntBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 88, 84))
}

// TradingStatus
func (m MarketDataSnapshot_Int) TradingStatus() TradingStatusEnum {
	return TradingStatusEnum(message.Int8Fixed(m.Unsafe(), 89, 88))
}

// TradingStatus
func (m MarketDataSnapshot_IntBuilder) TradingStatus() TradingStatusEnum {
	return TradingStatusEnum(message.Int8Fixed(m.Unsafe(), 89, 88))
}

// SetSymbolID
func (m MarketDataSnapshot_IntBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSessionSettlementPrice
func (m MarketDataSnapshot_IntBuilder) SetSessionSettlementPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetSessionOpenPrice
func (m MarketDataSnapshot_IntBuilder) SetSessionOpenPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 16, 12, value)
}

// SetSessionHighPrice
func (m MarketDataSnapshot_IntBuilder) SetSessionHighPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 20, 16, value)
}

// SetSessionLowPrice
func (m MarketDataSnapshot_IntBuilder) SetSessionLowPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 24, 20, value)
}

// SetSessionVolume
func (m MarketDataSnapshot_IntBuilder) SetSessionVolume(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 28, 24, value)
}

// SetSessionNumTrades
func (m MarketDataSnapshot_IntBuilder) SetSessionNumTrades(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 32, 28, value)
}

// SetOpenInterest
func (m MarketDataSnapshot_IntBuilder) SetOpenInterest(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 36, 32, value)
}

// SetBidPrice
func (m MarketDataSnapshot_IntBuilder) SetBidPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 40, 36, value)
}

// SetAskPrice
func (m MarketDataSnapshot_IntBuilder) SetAskPrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 44, 40, value)
}

// SetAskQuantity
func (m MarketDataSnapshot_IntBuilder) SetAskQuantity(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 48, 44, value)
}

// SetBidQuantity
func (m MarketDataSnapshot_IntBuilder) SetBidQuantity(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 52, 48, value)
}

// SetLastTradePrice
func (m MarketDataSnapshot_IntBuilder) SetLastTradePrice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 56, 52, value)
}

// SetLastTradeVolume
func (m MarketDataSnapshot_IntBuilder) SetLastTradeVolume(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 60, 56, value)
}

// SetLastTradeDateTime
func (m MarketDataSnapshot_IntBuilder) SetLastTradeDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 72, 64, float64(value))
}

// SetBidAskDateTime
func (m MarketDataSnapshot_IntBuilder) SetBidAskDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 80, 72, float64(value))
}

// SetSessionSettlementDateTime
func (m MarketDataSnapshot_IntBuilder) SetSessionSettlementDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 84, 80, uint32(value))
}

// SetTradingSessionDate
func (m MarketDataSnapshot_IntBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 88, 84, uint32(value))
}

// SetTradingStatus
func (m MarketDataSnapshot_IntBuilder) SetTradingStatus(value TradingStatusEnum) {
	message.SetInt8Fixed(m.Unsafe(), 89, 88, int8(value))
}

// Copy
func (m MarketDataSnapshot_Int) Copy(to MarketDataSnapshot_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSessionSettlementPrice(m.SessionSettlementPrice())
	to.SetSessionOpenPrice(m.SessionOpenPrice())
	to.SetSessionHighPrice(m.SessionHighPrice())
	to.SetSessionLowPrice(m.SessionLowPrice())
	to.SetSessionVolume(m.SessionVolume())
	to.SetSessionNumTrades(m.SessionNumTrades())
	to.SetOpenInterest(m.OpenInterest())
	to.SetBidPrice(m.BidPrice())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetBidQuantity(m.BidQuantity())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
	to.SetBidAskDateTime(m.BidAskDateTime())
	to.SetSessionSettlementDateTime(m.SessionSettlementDateTime())
	to.SetTradingSessionDate(m.TradingSessionDate())
	to.SetTradingStatus(m.TradingStatus())
}

// Copy
func (m MarketDataSnapshot_IntBuilder) Copy(to MarketDataSnapshot_IntBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSessionSettlementPrice(m.SessionSettlementPrice())
	to.SetSessionOpenPrice(m.SessionOpenPrice())
	to.SetSessionHighPrice(m.SessionHighPrice())
	to.SetSessionLowPrice(m.SessionLowPrice())
	to.SetSessionVolume(m.SessionVolume())
	to.SetSessionNumTrades(m.SessionNumTrades())
	to.SetOpenInterest(m.OpenInterest())
	to.SetBidPrice(m.BidPrice())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetBidQuantity(m.BidQuantity())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
	to.SetBidAskDateTime(m.BidAskDateTime())
	to.SetSessionSettlementDateTime(m.SessionSettlementDateTime())
	to.SetTradingSessionDate(m.TradingSessionDate())
	to.SetTradingStatus(m.TradingStatus())
}

// CopyPointer
func (m MarketDataSnapshot_Int) CopyPointer(to MarketDataSnapshot_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSessionSettlementPrice(m.SessionSettlementPrice())
	to.SetSessionOpenPrice(m.SessionOpenPrice())
	to.SetSessionHighPrice(m.SessionHighPrice())
	to.SetSessionLowPrice(m.SessionLowPrice())
	to.SetSessionVolume(m.SessionVolume())
	to.SetSessionNumTrades(m.SessionNumTrades())
	to.SetOpenInterest(m.OpenInterest())
	to.SetBidPrice(m.BidPrice())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetBidQuantity(m.BidQuantity())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
	to.SetBidAskDateTime(m.BidAskDateTime())
	to.SetSessionSettlementDateTime(m.SessionSettlementDateTime())
	to.SetTradingSessionDate(m.TradingSessionDate())
	to.SetTradingStatus(m.TradingStatus())
}

// CopyPointer
func (m MarketDataSnapshot_IntBuilder) CopyPointer(to MarketDataSnapshot_IntPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetSessionSettlementPrice(m.SessionSettlementPrice())
	to.SetSessionOpenPrice(m.SessionOpenPrice())
	to.SetSessionHighPrice(m.SessionHighPrice())
	to.SetSessionLowPrice(m.SessionLowPrice())
	to.SetSessionVolume(m.SessionVolume())
	to.SetSessionNumTrades(m.SessionNumTrades())
	to.SetOpenInterest(m.OpenInterest())
	to.SetBidPrice(m.BidPrice())
	to.SetAskPrice(m.AskPrice())
	to.SetAskQuantity(m.AskQuantity())
	to.SetBidQuantity(m.BidQuantity())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
	to.SetBidAskDateTime(m.BidAskDateTime())
	to.SetSessionSettlementDateTime(m.SessionSettlementDateTime())
	to.SetTradingSessionDate(m.TradingSessionDate())
	to.SetTradingStatus(m.TradingStatus())
}
