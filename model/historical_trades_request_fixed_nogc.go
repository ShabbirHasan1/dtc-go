package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalTradesRequestFixedPointer struct {
	message.FixedPointer
}

type HistoricalTradesRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalTradesRequestFixed() HistoricalTradesRequestFixedPointerBuilder {
	a := HistoricalTradesRequestFixedPointerBuilder{message.AllocFixed(117)}
	a.Ptr.SetBytes(0, _HistoricalTradesRequestFixedDefault)
	return a
}

func AllocHistoricalTradesRequestFixedFrom(b []byte) HistoricalTradesRequestFixedPointer {
	return HistoricalTradesRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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
func (m HistoricalTradesRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalTradesRequestFixedDefault)
}

// ToBuilder
func (m HistoricalTradesRequestFixedPointer) ToBuilder() HistoricalTradesRequestFixedPointerBuilder {
	return HistoricalTradesRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalTradesRequestFixedPointerBuilder) Finish() HistoricalTradesRequestFixedPointer {
	return HistoricalTradesRequestFixedPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalTradesRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalTradesRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Symbol
func (m HistoricalTradesRequestFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Symbol
func (m HistoricalTradesRequestFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// TradeAccount
func (m HistoricalTradesRequestFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 104, 72)
}

// TradeAccount
func (m HistoricalTradesRequestFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 104, 72)
}

// StartDateTime
func (m HistoricalTradesRequestFixedPointer) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 112, 104))
}

// StartDateTime
func (m HistoricalTradesRequestFixedPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 112, 104))
}

// SubAccountIdentifier
func (m HistoricalTradesRequestFixedPointer) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Ptr, 116, 112)
}

// SubAccountIdentifier
func (m HistoricalTradesRequestFixedPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Ptr, 116, 112)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestFixedPointer) CreateFlatToFlatTrades() uint8 {
	return message.Uint8Fixed(m.Ptr, 117, 116)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestFixedPointerBuilder) CreateFlatToFlatTrades() uint8 {
	return message.Uint8Fixed(m.Ptr, 117, 116)
}

// SetRequestID
func (m HistoricalTradesRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetSymbol
func (m HistoricalTradesRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 72, 8, value)
}

// SetTradeAccount
func (m HistoricalTradesRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 104, 72, value)
}

// SetStartDateTime
func (m HistoricalTradesRequestFixedPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 112, 104, int64(value))
}

// SetSubAccountIdentifier
func (m HistoricalTradesRequestFixedPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Ptr, 116, 112, value)
}

// SetCreateFlatToFlatTrades
func (m HistoricalTradesRequestFixedPointerBuilder) SetCreateFlatToFlatTrades(value uint8) {
	message.SetUint8Fixed(m.Ptr, 117, 116, value)
}

// Copy
func (m HistoricalTradesRequestFixedPointer) Copy(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// Copy
func (m HistoricalTradesRequestFixedPointerBuilder) Copy(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestFixedPointer) CopyPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestFixedPointerBuilder) CopyPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestFixedPointer) CopyTo(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestFixedPointerBuilder) CopyTo(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestFixedPointer) CopyToPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestFixedPointerBuilder) CopyToPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}
