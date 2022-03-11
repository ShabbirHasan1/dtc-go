package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                   = HistoricalTradesRequestFixedSize  (117)
//     Type                   = HISTORICAL_TRADES_REQUEST  (10100)
//     RequestID              = 0
//     Symbol                 = ""
//     TradeAccount           = ""
//     StartDateTime          = 0
//     SubAccountIdentifier   = 0
//     CreateFlatToFlatTrades = 0
// }
var _HistoricalTradesRequestFixedDefault = []byte{117, 0, 116, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalTradesRequestFixedSize = 117

type HistoricalTradesRequestFixed struct {
	message.Fixed
}

type HistoricalTradesRequestFixedBuilder struct {
	message.Fixed
}

func NewHistoricalTradesRequestFixedFrom(b []byte) HistoricalTradesRequestFixed {
	return HistoricalTradesRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalTradesRequestFixed(b []byte) HistoricalTradesRequestFixed {
	return HistoricalTradesRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalTradesRequestFixed() HistoricalTradesRequestFixedBuilder {
	a := HistoricalTradesRequestFixedBuilder{message.NewFixed(117)}
	a.Unsafe().SetBytes(0, _HistoricalTradesRequestFixedDefault)
	return a
}

// Clear
// {
//     Size                   = HistoricalTradesRequestFixedSize  (117)
//     Type                   = HISTORICAL_TRADES_REQUEST  (10100)
//     RequestID              = 0
//     Symbol                 = ""
//     TradeAccount           = ""
//     StartDateTime          = 0
//     SubAccountIdentifier   = 0
//     CreateFlatToFlatTrades = 0
// }
func (m HistoricalTradesRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalTradesRequestFixedDefault)
}

// ToBuilder
func (m HistoricalTradesRequestFixed) ToBuilder() HistoricalTradesRequestFixedBuilder {
	return HistoricalTradesRequestFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalTradesRequestFixedBuilder) Finish() HistoricalTradesRequestFixed {
	return HistoricalTradesRequestFixed{m.Fixed}
}

// RequestID
func (m HistoricalTradesRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalTradesRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// Symbol
func (m HistoricalTradesRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol
func (m HistoricalTradesRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// TradeAccount
func (m HistoricalTradesRequestFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 104, 72)
}

// TradeAccount
func (m HistoricalTradesRequestFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 104, 72)
}

// StartDateTime
func (m HistoricalTradesRequestFixed) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 112, 104))
}

// StartDateTime
func (m HistoricalTradesRequestFixedBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 112, 104))
}

// SubAccountIdentifier
func (m HistoricalTradesRequestFixed) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 116, 112)
}

// SubAccountIdentifier
func (m HistoricalTradesRequestFixedBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 116, 112)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestFixed) CreateFlatToFlatTrades() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 117, 116)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestFixedBuilder) CreateFlatToFlatTrades() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 117, 116)
}

// SetRequestID
func (m HistoricalTradesRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbol
func (m HistoricalTradesRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetTradeAccount
func (m HistoricalTradesRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 72, value)
}

// SetStartDateTime
func (m HistoricalTradesRequestFixedBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 112, 104, int64(value))
}

// SetSubAccountIdentifier
func (m HistoricalTradesRequestFixedBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 116, 112, value)
}

// SetCreateFlatToFlatTrades
func (m HistoricalTradesRequestFixedBuilder) SetCreateFlatToFlatTrades(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 117, 116, value)
}

// Copy
func (m HistoricalTradesRequestFixed) Copy(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// Copy
func (m HistoricalTradesRequestFixedBuilder) Copy(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestFixed) CopyPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestFixedBuilder) CopyPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestFixed) CopyToPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestFixedBuilder) CopyToPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestFixed) CopyTo(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestFixedBuilder) CopyTo(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}
