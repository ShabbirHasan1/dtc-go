package model

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDataUpdateSessionVolumeSize = 24

// {
//     Size                 = MarketDataUpdateSessionVolumeSize  (24)
//     Type                 = MARKET_DATA_UPDATE_SESSION_VOLUME  (113)
//     SymbolID             = 0
//     Volume               = 0
//     TradingSessionDate   = 0
//     IsFinalSessionVolume = false
// }
var _MarketDataUpdateSessionVolumeDefault = []byte{24, 0, 113, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateSessionVolume Sent by the Server to the Client when the session trade Volume needs to
// be updated.
//
// The recommended rule for the Server to notify the Client of a change with
// the session trade volume to maintain bandwidth efficiency, is as follows:
// When a trade occurs for a symbol subscribed to, the Server will send a
// MarketDataUpdateTrade message to the Client. The Client should then increment
// its session trade volume value for the symbol by the value in the Volume
// field in this message.
//
// The Server will assume the Client is doing this. Therefore, when a trade
// occurs and the session trade volume does not equal the prior session trade
// volume plus the Volume for the most recent trade sent to the Client, then
// the Server must send out a MarketDataUpdateSessionVolume message to the
// client since the client calculation of the session trade volume is no
// longer correct.
//
// It is assumed that the reason for this inconsistency is due to trades
// included within the session trade volume which have not been sent out
// as normal trades.
//
// The Server should also send this message out at the frequency that the
// Server determines, such as every minute if there also has been a trade
// at that time.
type MarketDataUpdateSessionVolume struct {
	message.Fixed
}

// MarketDataUpdateSessionVolumeBuilder Sent by the Server to the Client when the session trade Volume needs to
// be updated.
//
// The recommended rule for the Server to notify the Client of a change with
// the session trade volume to maintain bandwidth efficiency, is as follows:
// When a trade occurs for a symbol subscribed to, the Server will send a
// MarketDataUpdateTrade message to the Client. The Client should then increment
// its session trade volume value for the symbol by the value in the Volume
// field in this message.
//
// The Server will assume the Client is doing this. Therefore, when a trade
// occurs and the session trade volume does not equal the prior session trade
// volume plus the Volume for the most recent trade sent to the Client, then
// the Server must send out a MarketDataUpdateSessionVolume message to the
// client since the client calculation of the session trade volume is no
// longer correct.
//
// It is assumed that the reason for this inconsistency is due to trades
// included within the session trade volume which have not been sent out
// as normal trades.
//
// The Server should also send this message out at the frequency that the
// Server determines, such as every minute if there also has been a trade
// at that time.
type MarketDataUpdateSessionVolumeBuilder struct {
	message.Fixed
}

// MarketDataUpdateSessionVolumePointer Sent by the Server to the Client when the session trade Volume needs to
// be updated.
//
// The recommended rule for the Server to notify the Client of a change with
// the session trade volume to maintain bandwidth efficiency, is as follows:
// When a trade occurs for a symbol subscribed to, the Server will send a
// MarketDataUpdateTrade message to the Client. The Client should then increment
// its session trade volume value for the symbol by the value in the Volume
// field in this message.
//
// The Server will assume the Client is doing this. Therefore, when a trade
// occurs and the session trade volume does not equal the prior session trade
// volume plus the Volume for the most recent trade sent to the Client, then
// the Server must send out a MarketDataUpdateSessionVolume message to the
// client since the client calculation of the session trade volume is no
// longer correct.
//
// It is assumed that the reason for this inconsistency is due to trades
// included within the session trade volume which have not been sent out
// as normal trades.
//
// The Server should also send this message out at the frequency that the
// Server determines, such as every minute if there also has been a trade
// at that time.
type MarketDataUpdateSessionVolumePointer struct {
	message.FixedPointer
}

// MarketDataUpdateSessionVolumePointerBuilder Sent by the Server to the Client when the session trade Volume needs to
// be updated.
//
// The recommended rule for the Server to notify the Client of a change with
// the session trade volume to maintain bandwidth efficiency, is as follows:
// When a trade occurs for a symbol subscribed to, the Server will send a
// MarketDataUpdateTrade message to the Client. The Client should then increment
// its session trade volume value for the symbol by the value in the Volume
// field in this message.
//
// The Server will assume the Client is doing this. Therefore, when a trade
// occurs and the session trade volume does not equal the prior session trade
// volume plus the Volume for the most recent trade sent to the Client, then
// the Server must send out a MarketDataUpdateSessionVolume message to the
// client since the client calculation of the session trade volume is no
// longer correct.
//
// It is assumed that the reason for this inconsistency is due to trades
// included within the session trade volume which have not been sent out
// as normal trades.
//
// The Server should also send this message out at the frequency that the
// Server determines, such as every minute if there also has been a trade
// at that time.
type MarketDataUpdateSessionVolumePointerBuilder struct {
	message.FixedPointer
}

func NewMarketDataUpdateSessionVolumeFrom(b []byte) MarketDataUpdateSessionVolume {
	return MarketDataUpdateSessionVolume{Fixed: message.NewFixedFrom(b)}
}

func WrapMarketDataUpdateSessionVolume(b []byte) MarketDataUpdateSessionVolume {
	return MarketDataUpdateSessionVolume{Fixed: message.WrapFixed(b)}
}

func NewMarketDataUpdateSessionVolume() MarketDataUpdateSessionVolumeBuilder {
	a := MarketDataUpdateSessionVolumeBuilder{message.NewFixed(24)}
	a.Unsafe().SetBytes(0, _MarketDataUpdateSessionVolumeDefault)
	return a
}

func AllocMarketDataUpdateSessionVolume() MarketDataUpdateSessionVolumePointerBuilder {
	a := MarketDataUpdateSessionVolumePointerBuilder{message.AllocFixed(24)}
	a.Ptr.SetBytes(0, _MarketDataUpdateSessionVolumeDefault)
	return a
}

func AllocMarketDataUpdateSessionVolumeFrom(b []byte) MarketDataUpdateSessionVolumePointer {
	return MarketDataUpdateSessionVolumePointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                 = MarketDataUpdateSessionVolumeSize  (24)
//     Type                 = MARKET_DATA_UPDATE_SESSION_VOLUME  (113)
//     SymbolID             = 0
//     Volume               = 0
//     TradingSessionDate   = 0
//     IsFinalSessionVolume = false
// }
func (m MarketDataUpdateSessionVolumeBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDataUpdateSessionVolumeDefault)
}

// Clear
// {
//     Size                 = MarketDataUpdateSessionVolumeSize  (24)
//     Type                 = MARKET_DATA_UPDATE_SESSION_VOLUME  (113)
//     SymbolID             = 0
//     Volume               = 0
//     TradingSessionDate   = 0
//     IsFinalSessionVolume = false
// }
func (m MarketDataUpdateSessionVolumePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDataUpdateSessionVolumeDefault)
}

// ToBuilder
func (m MarketDataUpdateSessionVolume) ToBuilder() MarketDataUpdateSessionVolumeBuilder {
	return MarketDataUpdateSessionVolumeBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDataUpdateSessionVolumePointer) ToBuilder() MarketDataUpdateSessionVolumePointerBuilder {
	return MarketDataUpdateSessionVolumePointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDataUpdateSessionVolumeBuilder) Finish() MarketDataUpdateSessionVolume {
	return MarketDataUpdateSessionVolume{m.Fixed}
}

// Finish
func (m *MarketDataUpdateSessionVolumePointerBuilder) Finish() MarketDataUpdateSessionVolumePointer {
	return MarketDataUpdateSessionVolumePointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionVolume) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionVolumeBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionVolumePointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionVolumePointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Volume The total Volume for the session.
func (m MarketDataUpdateSessionVolume) Volume() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// Volume The total Volume for the session.
func (m MarketDataUpdateSessionVolumeBuilder) Volume() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// Volume The total Volume for the session.
func (m MarketDataUpdateSessionVolumePointer) Volume() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// Volume The total Volume for the session.
func (m MarketDataUpdateSessionVolumePointerBuilder) Volume() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// TradingSessionDate
func (m MarketDataUpdateSessionVolume) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionVolumeBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionVolumePointer) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 20, 16))
}

// TradingSessionDate
func (m MarketDataUpdateSessionVolumePointerBuilder) TradingSessionDate() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 20, 16))
}

// IsFinalSessionVolume
func (m MarketDataUpdateSessionVolume) IsFinalSessionVolume() bool {
	return message.BoolFixed(m.Unsafe(), 21, 20)
}

// IsFinalSessionVolume
func (m MarketDataUpdateSessionVolumeBuilder) IsFinalSessionVolume() bool {
	return message.BoolFixed(m.Unsafe(), 21, 20)
}

// IsFinalSessionVolume
func (m MarketDataUpdateSessionVolumePointer) IsFinalSessionVolume() bool {
	return message.BoolFixed(m.Ptr, 21, 20)
}

// IsFinalSessionVolume
func (m MarketDataUpdateSessionVolumePointerBuilder) IsFinalSessionVolume() bool {
	return message.BoolFixed(m.Ptr, 21, 20)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionVolumeBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateSessionVolumePointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetVolume The total Volume for the session.
func (m MarketDataUpdateSessionVolumeBuilder) SetVolume(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, value)
}

// SetVolume The total Volume for the session.
func (m MarketDataUpdateSessionVolumePointerBuilder) SetVolume(value float64) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, value)
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionVolumeBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 20, 16, uint32(value))
}

// SetTradingSessionDate
func (m MarketDataUpdateSessionVolumePointerBuilder) SetTradingSessionDate(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 20, 16, uint32(value))
}

// SetIsFinalSessionVolume
func (m MarketDataUpdateSessionVolumeBuilder) SetIsFinalSessionVolume(value bool) {
	message.SetBoolFixed(m.Unsafe(), 21, 20, value)
}

// SetIsFinalSessionVolume
func (m MarketDataUpdateSessionVolumePointerBuilder) SetIsFinalSessionVolume(value bool) {
	message.SetBoolFixed(m.Ptr, 21, 20, value)
}

// Copy
func (m MarketDataUpdateSessionVolume) Copy(to MarketDataUpdateSessionVolumeBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetVolume(m.Volume())
	to.SetTradingSessionDate(m.TradingSessionDate())
	to.SetIsFinalSessionVolume(m.IsFinalSessionVolume())
}

// Copy
func (m MarketDataUpdateSessionVolumeBuilder) Copy(to MarketDataUpdateSessionVolumeBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetVolume(m.Volume())
	to.SetTradingSessionDate(m.TradingSessionDate())
	to.SetIsFinalSessionVolume(m.IsFinalSessionVolume())
}

// CopyPointer
func (m MarketDataUpdateSessionVolume) CopyPointer(to MarketDataUpdateSessionVolumePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetVolume(m.Volume())
	to.SetTradingSessionDate(m.TradingSessionDate())
	to.SetIsFinalSessionVolume(m.IsFinalSessionVolume())
}

// CopyPointer
func (m MarketDataUpdateSessionVolumeBuilder) CopyPointer(to MarketDataUpdateSessionVolumePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetVolume(m.Volume())
	to.SetTradingSessionDate(m.TradingSessionDate())
	to.SetIsFinalSessionVolume(m.IsFinalSessionVolume())
}

// Copy
func (m MarketDataUpdateSessionVolumePointer) Copy(to MarketDataUpdateSessionVolumeBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetVolume(m.Volume())
	to.SetTradingSessionDate(m.TradingSessionDate())
	to.SetIsFinalSessionVolume(m.IsFinalSessionVolume())
}

// Copy
func (m MarketDataUpdateSessionVolumePointerBuilder) Copy(to MarketDataUpdateSessionVolumeBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetVolume(m.Volume())
	to.SetTradingSessionDate(m.TradingSessionDate())
	to.SetIsFinalSessionVolume(m.IsFinalSessionVolume())
}

// CopyPointer
func (m MarketDataUpdateSessionVolumePointer) CopyPointer(to MarketDataUpdateSessionVolumePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetVolume(m.Volume())
	to.SetTradingSessionDate(m.TradingSessionDate())
	to.SetIsFinalSessionVolume(m.IsFinalSessionVolume())
}

// CopyPointer
func (m MarketDataUpdateSessionVolumePointerBuilder) CopyPointer(to MarketDataUpdateSessionVolumePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetVolume(m.Volume())
	to.SetTradingSessionDate(m.TradingSessionDate())
	to.SetIsFinalSessionVolume(m.IsFinalSessionVolume())
}
