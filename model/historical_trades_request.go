package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _HistoricalTradesRequestDefault = []byte{31, 0, 116, 39, 31, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalTradesRequestSize = 31

type HistoricalTradesRequest struct {
	message.VLS
}

type HistoricalTradesRequestBuilder struct {
	message.VLSBuilder
}

func NewHistoricalTradesRequestFrom(b []byte) HistoricalTradesRequest {
	return HistoricalTradesRequest{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalTradesRequest(b []byte) HistoricalTradesRequest {
	return HistoricalTradesRequest{VLS: message.WrapVLS(b)}
}

func NewHistoricalTradesRequest() HistoricalTradesRequestBuilder {
	a := HistoricalTradesRequestBuilder{message.NewVLSBuilder(31)}
	a.Unsafe().SetBytes(0, _HistoricalTradesRequestDefault)
	return a
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
func (m HistoricalTradesRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalTradesRequestDefault)
}

// ToBuilder
func (m HistoricalTradesRequest) ToBuilder() HistoricalTradesRequestBuilder {
	return HistoricalTradesRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m HistoricalTradesRequestBuilder) Finish() HistoricalTradesRequest {
	return HistoricalTradesRequest{m.VLSBuilder.Finish()}
}

// RequestID
func (m HistoricalTradesRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m HistoricalTradesRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// Symbol
func (m HistoricalTradesRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Symbol
func (m HistoricalTradesRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m HistoricalTradesRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount
func (m HistoricalTradesRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// StartDateTime
func (m HistoricalTradesRequest) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 26, 18))
}

// StartDateTime
func (m HistoricalTradesRequestBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 26, 18))
}

// SubAccountIdentifier
func (m HistoricalTradesRequest) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 30, 26)
}

// SubAccountIdentifier
func (m HistoricalTradesRequestBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 30, 26)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequest) CreateFlatToFlatTrades() uint8 {
	return message.Uint8VLS(m.Unsafe(), 31, 30)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestBuilder) CreateFlatToFlatTrades() uint8 {
	return message.Uint8VLS(m.Unsafe(), 31, 30)
}

// SetRequestID
func (m HistoricalTradesRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 10, 6, value)
}

// SetSymbol
func (m *HistoricalTradesRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *HistoricalTradesRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetStartDateTime
func (m HistoricalTradesRequestBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 26, 18, int64(value))
}

// SetSubAccountIdentifier
func (m HistoricalTradesRequestBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 30, 26, value)
}

// SetCreateFlatToFlatTrades
func (m HistoricalTradesRequestBuilder) SetCreateFlatToFlatTrades(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 31, 30, value)
}

// Copy
func (m HistoricalTradesRequest) Copy(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// Copy
func (m HistoricalTradesRequestBuilder) Copy(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequest) CopyPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestBuilder) CopyPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequest) CopyToPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestBuilder) CopyToPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequest) CopyTo(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestBuilder) CopyTo(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}
