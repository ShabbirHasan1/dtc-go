package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                      = MarketDataSnapshotSize  (144)
//     Type                      = MARKET_DATA_SNAPSHOT  (104)
//     SymbolID                  = 0
//     SessionSettlementPrice    = math.MaxFloat64
//     SessionOpenPrice          = math.MaxFloat64
//     SessionHighPrice          = math.MaxFloat64
//     SessionLowPrice           = math.MaxFloat64
//     SessionVolume             = math.MaxFloat64
//     SessionNumTrades          = 65535
//     OpenInterest              = 65535
//     BidPrice                  = math.MaxFloat64
//     AskPrice                  = math.MaxFloat64
//     AskQuantity               = math.MaxFloat64
//     BidQuantity               = math.MaxFloat64
//     LastTradePrice            = math.MaxFloat64
//     LastTradeVolume           = math.MaxFloat64
//     LastTradeDateTime         = 0
//     BidAskDateTime            = 0
//     SessionSettlementDateTime = 0
//     TradingSessionDate        = 0
//     TradingStatus             = 0
//     MarketDepthUpdateDateTime = 0
// }
var _MarketDataSnapshotDefault = []byte{144, 0, 104, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 255, 255, 255, 255, 255, 255, 239, 127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const MarketDataSnapshotSize = 144

// MarketDataSnapshot The Server sends the MarketDataSnapshot message to the Client immediately
// after a successful MarketDataRequest message has been received from the
// Client and it has indicated to subscribe to the symbol or requested the
// snapshot of data.
//
// Any changes to the data fields within the MarketDataSnapshot message during
// the trading session will be sent by the Server to the Client through the
// corresponding MARKET_DATA_UPDATE_* messages.
//
// It is recommended that the MarketDataSnapshot be sent by the Server at
// the start of a new trading session.
//
// This message can be sent more often, however it is not intended to be
// sent frequently.
//
// This message type does not signify a trade has occurred. It should never
// be interpreted by the Client in that way.
//
// There is no need to send this when there is a new High or Low during the
// trading session. The Server should use the MarketDataUpdateSessionHigh
// or MarketDataUpdateSessionLow messages instead.
type MarketDataSnapshot struct {
	message.Fixed
}

// MarketDataSnapshotBuilder The Server sends the MarketDataSnapshot message to the Client immediately
// after a successful MarketDataRequest message has been received from the
// Client and it has indicated to subscribe to the symbol or requested the
// snapshot of data.
//
// Any changes to the data fields within the MarketDataSnapshot message during
// the trading session will be sent by the Server to the Client through the
// corresponding MARKET_DATA_UPDATE_* messages.
//
// It is recommended that the MarketDataSnapshot be sent by the Server at
// the start of a new trading session.
//
// This message can be sent more often, however it is not intended to be
// sent frequently.
//
// This message type does not signify a trade has occurred. It should never
// be interpreted by the Client in that way.
//
// There is no need to send this when there is a new High or Low during the
// trading session. The Server should use the MarketDataUpdateSessionHigh
// or MarketDataUpdateSessionLow messages instead.
type MarketDataSnapshotBuilder struct {
	message.Fixed
}

func NewMarketDataSnapshotFrom(b []byte) MarketDataSnapshot {
	return MarketDataSnapshot{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataSnapshot(b []byte) MarketDataSnapshot {
	return MarketDataSnapshot{Fixed: message.WrapFixed(b)}
}

func NewMarketDataSnapshot() MarketDataSnapshotBuilder {
	a := MarketDataSnapshotBuilder{message.NewFixed(144)}
	a.Unsafe().SetBytes(0, _MarketDataSnapshotDefault)
	return a
}

// Clear
// {
//     Size                      = MarketDataSnapshotSize  (144)
//     Type                      = MARKET_DATA_SNAPSHOT  (104)
//     SymbolID                  = 0
//     SessionSettlementPrice    = math.MaxFloat64
//     SessionOpenPrice          = math.MaxFloat64
//     SessionHighPrice          = math.MaxFloat64
//     SessionLowPrice           = math.MaxFloat64
//     SessionVolume             = math.MaxFloat64
//     SessionNumTrades          = 65535
//     OpenInterest              = 65535
//     BidPrice                  = math.MaxFloat64
//     AskPrice                  = math.MaxFloat64
//     AskQuantity               = math.MaxFloat64
//     BidQuantity               = math.MaxFloat64
//     LastTradePrice            = math.MaxFloat64
//     LastTradeVolume           = math.MaxFloat64
//     LastTradeDateTime         = 0
//     BidAskDateTime            = 0
//     SessionSettlementDateTime = 0
//     TradingSessionDate        = 0
//     TradingStatus             = 0
//     MarketDepthUpdateDateTime = 0
// }
func (m MarketDataSnapshotBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataSnapshotDefault)
}

// ToBuilder
func (m MarketDataSnapshot) ToBuilder() MarketDataSnapshotBuilder {
	return MarketDataSnapshotBuilder{m.Fixed}
}

// Finish
func (m MarketDataSnapshotBuilder) Finish() MarketDataSnapshot {
	return MarketDataSnapshot{m.Fixed}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataSnapshot) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataSnapshotBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SessionSettlementPrice The previous Settlement price when this message is sent before the market
// closes for the trading session. After the market has closed, this is the
// most recent Settlement price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) SessionSettlementPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// SessionSettlementPrice The previous Settlement price when this message is sent before the market
// closes for the trading session. After the market has closed, this is the
// most recent Settlement price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SessionSettlementPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// SessionOpenPrice The Opening price for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) SessionOpenPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// SessionOpenPrice The Opening price for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SessionOpenPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// SessionHighPrice The the High price for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) SessionHighPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// SessionHighPrice The the High price for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SessionHighPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 32, 24)
}

// SessionLowPrice The Low price for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) SessionLowPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 40, 32)
}

// SessionLowPrice The Low price for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SessionLowPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 40, 32)
}

// SessionVolume The total Volume for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) SessionVolume() float64 {
	return message.Float64Fixed(m.Unsafe(), 48, 40)
}

// SessionVolume The total Volume for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SessionVolume() float64 {
	return message.Float64Fixed(m.Unsafe(), 48, 40)
}

// SessionNumTrades The number of Trades for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to UINT_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) SessionNumTrades() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 52, 48)
}

// SessionNumTrades The number of Trades for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to UINT_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SessionNumTrades() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 52, 48)
}

// OpenInterest Contains the Open Interest for futures and options.
//
// For binary encoding, if this field is not set it needs to be set to UINT_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 56, 52)
}

// OpenInterest Contains the Open Interest for futures and options.
//
// For binary encoding, if this field is not set it needs to be set to UINT_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) OpenInterest() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 56, 52)
}

// BidPrice The latest best Bid price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) BidPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 64, 56)
}

// BidPrice The latest best Bid price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) BidPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 64, 56)
}

// AskPrice The latest best Ask price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) AskPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 72, 64)
}

// AskPrice The latest best Ask price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) AskPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 72, 64)
}

// AskQuantity The quantity of the orders at the Ask price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) AskQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 80, 72)
}

// AskQuantity The quantity of the orders at the Ask price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) AskQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 80, 72)
}

// BidQuantity The quantity of the orders at the Bid price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) BidQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 88, 80)
}

// BidQuantity The quantity of the orders at the Bid price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) BidQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 88, 80)
}

// LastTradePrice The most recent last trade price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) LastTradePrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 96, 88)
}

// LastTradePrice The most recent last trade price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) LastTradePrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 96, 88)
}

// LastTradeVolume The quantity/volume of the most recent last trade.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshot) LastTradeVolume() float64 {
	return message.Float64Fixed(m.Unsafe(), 104, 96)
}

// LastTradeVolume The quantity/volume of the most recent last trade.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) LastTradeVolume() float64 {
	return message.Float64Fixed(m.Unsafe(), 104, 96)
}

// LastTradeDateTime The Date-Time of the last trade.
func (m MarketDataSnapshot) LastTradeDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 112, 104))
}

// LastTradeDateTime The Date-Time of the last trade.
func (m MarketDataSnapshotBuilder) LastTradeDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 112, 104))
}

// BidAskDateTime The Date-Time of the last Bid and Ask quote data update.
func (m MarketDataSnapshot) BidAskDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 120, 112))
}

// BidAskDateTime The Date-Time of the last Bid and Ask quote data update.
func (m MarketDataSnapshotBuilder) BidAskDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 120, 112))
}

// SessionSettlementDateTime The trading date the Settlement price is for. The time component is not
// normally considered relevant for this field.
//
// This field will be 0 if this field is not available from the data feed.
// This field will be 0 if this field is not available from the data feed.
func (m MarketDataSnapshot) SessionSettlementDateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 124, 120))
}

// SessionSettlementDateTime The trading date the Settlement price is for. The time component is not
// normally considered relevant for this field.
//
// This field will be 0 if this field is not available from the data feed.
// This field will be 0 if this field is not available from the data feed.
func (m MarketDataSnapshotBuilder) SessionSettlementDateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 124, 120))
}

// TradingSessionDate This is the Date of the trading session that the data contained in this
// snapshot message is for.
//
// The time component is not normally considered relevant for this field.
func (m MarketDataSnapshot) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 128, 124))
}

// TradingSessionDate This is the Date of the trading session that the data contained in this
// snapshot message is for.
//
// The time component is not normally considered relevant for this field.
func (m MarketDataSnapshotBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 128, 124))
}

// TradingStatus
func (m MarketDataSnapshot) TradingStatus() TradingStatusEnum {
	return TradingStatusEnum(message.Int8Fixed(m.Unsafe(), 129, 128))
}

// TradingStatus
func (m MarketDataSnapshotBuilder) TradingStatus() TradingStatusEnum {
	return TradingStatusEnum(message.Int8Fixed(m.Unsafe(), 129, 128))
}

// MarketDepthUpdateDateTime
func (m MarketDataSnapshot) MarketDepthUpdateDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 144, 136))
}

// MarketDepthUpdateDateTime
func (m MarketDataSnapshotBuilder) MarketDepthUpdateDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(message.Float64Fixed(m.Unsafe(), 144, 136))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataSnapshotBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSessionSettlementPrice The previous Settlement price when this message is sent before the market
// closes for the trading session. After the market has closed, this is the
// most recent Settlement price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetSessionSettlementPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, value)
}

// SetSessionOpenPrice The Opening price for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetSessionOpenPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetSessionHighPrice The the High price for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetSessionHighPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 32, 24, value)
}

// SetSessionLowPrice The Low price for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetSessionLowPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 40, 32, value)
}

// SetSessionVolume The total Volume for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetSessionVolume(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 48, 40, value)
}

// SetSessionNumTrades The number of Trades for the trading session.
//
// For binary encoding, if this field is not set it needs to be set to UINT_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetSessionNumTrades(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 52, 48, value)
}

// SetOpenInterest Contains the Open Interest for futures and options.
//
// For binary encoding, if this field is not set it needs to be set to UINT_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetOpenInterest(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 56, 52, value)
}

// SetBidPrice The latest best Bid price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetBidPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 64, 56, value)
}

// SetAskPrice The latest best Ask price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetAskPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 72, 64, value)
}

// SetAskQuantity The quantity of the orders at the Ask price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetAskQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 80, 72, value)
}

// SetBidQuantity The quantity of the orders at the Bid price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetBidQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 88, 80, value)
}

// SetLastTradePrice The most recent last trade price.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetLastTradePrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 96, 88, value)
}

// SetLastTradeVolume The quantity/volume of the most recent last trade.
//
// For binary encoding, if this field is not set it needs to be set to DBL_MAX.
// Refer to Unset Message Fields.
func (m MarketDataSnapshotBuilder) SetLastTradeVolume(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 104, 96, value)
}

// SetLastTradeDateTime The Date-Time of the last trade.
func (m MarketDataSnapshotBuilder) SetLastTradeDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 112, 104, float64(value))
}

// SetBidAskDateTime The Date-Time of the last Bid and Ask quote data update.
func (m MarketDataSnapshotBuilder) SetBidAskDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 120, 112, float64(value))
}

// SetSessionSettlementDateTime The trading date the Settlement price is for. The time component is not
// normally considered relevant for this field.
//
// This field will be 0 if this field is not available from the data feed.
// This field will be 0 if this field is not available from the data feed.
func (m MarketDataSnapshotBuilder) SetSessionSettlementDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 124, 120, uint32(value))
}

// SetTradingSessionDate This is the Date of the trading session that the data contained in this
// snapshot message is for.
//
// The time component is not normally considered relevant for this field.
func (m MarketDataSnapshotBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 128, 124, uint32(value))
}

// SetTradingStatus
func (m MarketDataSnapshotBuilder) SetTradingStatus(value TradingStatusEnum) {
	message.SetInt8Fixed(m.Unsafe(), 129, 128, int8(value))
}

// SetMarketDepthUpdateDateTime
func (m MarketDataSnapshotBuilder) SetMarketDepthUpdateDateTime(value DateTimeWithMilliseconds) {
	message.SetFloat64Fixed(m.Unsafe(), 144, 136, float64(value))
}

// Copy
func (m MarketDataSnapshot) Copy(to MarketDataSnapshotBuilder) {
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
	to.SetMarketDepthUpdateDateTime(m.MarketDepthUpdateDateTime())
}

// Copy
func (m MarketDataSnapshotBuilder) Copy(to MarketDataSnapshotBuilder) {
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
	to.SetMarketDepthUpdateDateTime(m.MarketDepthUpdateDateTime())
}

// CopyPointer
func (m MarketDataSnapshot) CopyPointer(to MarketDataSnapshotPointerBuilder) {
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
	to.SetMarketDepthUpdateDateTime(m.MarketDepthUpdateDateTime())
}

// CopyPointer
func (m MarketDataSnapshotBuilder) CopyPointer(to MarketDataSnapshotPointerBuilder) {
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
	to.SetMarketDepthUpdateDateTime(m.MarketDepthUpdateDateTime())
}
