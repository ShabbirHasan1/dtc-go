package model

import (
	"github.com/moontrade/dtc-go/message"
)

type HistoricalTradesRequestPointer struct {
	message.VLSPointer
}

type HistoricalTradesRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocHistoricalTradesRequest() HistoricalTradesRequestPointerBuilder {
	a := HistoricalTradesRequestPointerBuilder{message.AllocVLSBuilder(31)}
	a.Ptr.SetBytes(0, _HistoricalTradesRequestDefault)
	return a
}

func AllocHistoricalTradesRequestFrom(b []byte) HistoricalTradesRequestPointer {
	return HistoricalTradesRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                   = HistoricalTradesRequestSize  (31)
//     Type                   = HISTORICAL_TRADES_REQUEST  (10100)
//     BaseSize               = HistoricalTradesRequestSize  (31)
//     RequestID              = 0
//     Symbol                 = ""
//     TradeAccount           = ""
//     StartDateTime          = 0
//     SubAccountIdentifier   = 0
//     CreateFlatToFlatTrades = 0
// }
func (m HistoricalTradesRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalTradesRequestDefault)
}

// ToBuilder
func (m HistoricalTradesRequestPointer) ToBuilder() HistoricalTradesRequestPointerBuilder {
	return HistoricalTradesRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalTradesRequestPointerBuilder) Finish() HistoricalTradesRequestPointer {
	return HistoricalTradesRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m HistoricalTradesRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m HistoricalTradesRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// Symbol
func (m HistoricalTradesRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Symbol
func (m HistoricalTradesRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m HistoricalTradesRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TradeAccount
func (m HistoricalTradesRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// StartDateTime
func (m HistoricalTradesRequestPointer) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 26, 18))
}

// StartDateTime
func (m HistoricalTradesRequestPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 26, 18))
}

// SubAccountIdentifier
func (m HistoricalTradesRequestPointer) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 30, 26)
}

// SubAccountIdentifier
func (m HistoricalTradesRequestPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 30, 26)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestPointer) CreateFlatToFlatTrades() uint8 {
	return message.Uint8VLS(m.Ptr, 31, 30)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestPointerBuilder) CreateFlatToFlatTrades() uint8 {
	return message.Uint8VLS(m.Ptr, 31, 30)
}

// SetRequestID
func (m HistoricalTradesRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 10, 6, value)
}

// SetSymbol
func (m *HistoricalTradesRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *HistoricalTradesRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetStartDateTime
func (m HistoricalTradesRequestPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 26, 18, int64(value))
}

// SetSubAccountIdentifier
func (m HistoricalTradesRequestPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Ptr, 30, 26, value)
}

// SetCreateFlatToFlatTrades
func (m HistoricalTradesRequestPointerBuilder) SetCreateFlatToFlatTrades(value uint8) {
	message.SetUint8VLS(m.Ptr, 31, 30, value)
}

// Copy
func (m HistoricalTradesRequestPointer) Copy(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// Copy
func (m HistoricalTradesRequestPointerBuilder) Copy(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestPointer) CopyPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestPointerBuilder) CopyPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestPointer) CopyTo(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestPointerBuilder) CopyTo(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestPointer) CopyToPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestPointerBuilder) CopyToPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}
