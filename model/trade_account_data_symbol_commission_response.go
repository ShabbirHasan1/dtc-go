package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                = TradeAccountDataSymbolCommissionResponseSize  (26)
//     Type                = TRADE_ACCOUNT_DATA_SYMBOL_COMMISSION_RESPONSE  (10119)
//     BaseSize            = TradeAccountDataSymbolCommissionResponseSize  (26)
//     RequestID           = 0
//     TradeAccount        = ""
//     Symbol              = ""
//     TradeFeePerContract = 0
// }
var _TradeAccountDataSymbolCommissionResponseDefault = []byte{26, 0, 135, 39, 26, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataSymbolCommissionResponseSize = 26

type TradeAccountDataSymbolCommissionResponse struct {
	message.VLS
}

type TradeAccountDataSymbolCommissionResponseBuilder struct {
	message.VLSBuilder
}

func NewTradeAccountDataSymbolCommissionResponseFrom(b []byte) TradeAccountDataSymbolCommissionResponse {
	return TradeAccountDataSymbolCommissionResponse{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataSymbolCommissionResponse(b []byte) TradeAccountDataSymbolCommissionResponse {
	return TradeAccountDataSymbolCommissionResponse{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataSymbolCommissionResponse() TradeAccountDataSymbolCommissionResponseBuilder {
	a := TradeAccountDataSymbolCommissionResponseBuilder{message.NewVLSBuilder(26)}
	a.Unsafe().SetBytes(0, _TradeAccountDataSymbolCommissionResponseDefault)
	return a
}

// Clear
// {
//     Size                = TradeAccountDataSymbolCommissionResponseSize  (26)
//     Type                = TRADE_ACCOUNT_DATA_SYMBOL_COMMISSION_RESPONSE  (10119)
//     BaseSize            = TradeAccountDataSymbolCommissionResponseSize  (26)
//     RequestID           = 0
//     TradeAccount        = ""
//     Symbol              = ""
//     TradeFeePerContract = 0
// }
func (m TradeAccountDataSymbolCommissionResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataSymbolCommissionResponseDefault)
}

// ToBuilder
func (m TradeAccountDataSymbolCommissionResponse) ToBuilder() TradeAccountDataSymbolCommissionResponseBuilder {
	return TradeAccountDataSymbolCommissionResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataSymbolCommissionResponseBuilder) Finish() TradeAccountDataSymbolCommissionResponse {
	return TradeAccountDataSymbolCommissionResponse{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataSymbolCommissionResponse) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataSymbolCommissionResponseBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m TradeAccountDataSymbolCommissionResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataSymbolCommissionResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Symbol
func (m TradeAccountDataSymbolCommissionResponse) Symbol() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Symbol
func (m TradeAccountDataSymbolCommissionResponseBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeFeePerContract
func (m TradeAccountDataSymbolCommissionResponse) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Unsafe(), 26, 18)
}

// TradeFeePerContract
func (m TradeAccountDataSymbolCommissionResponseBuilder) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Unsafe(), 26, 18)
}

// SetRequestID
func (m TradeAccountDataSymbolCommissionResponseBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataSymbolCommissionResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetSymbol
func (m *TradeAccountDataSymbolCommissionResponseBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataSymbolCommissionResponseBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 26, 18, value)
}

// Copy
func (m TradeAccountDataSymbolCommissionResponse) Copy(to TradeAccountDataSymbolCommissionResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

// Copy
func (m TradeAccountDataSymbolCommissionResponseBuilder) Copy(to TradeAccountDataSymbolCommissionResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

// CopyPointer
func (m TradeAccountDataSymbolCommissionResponse) CopyPointer(to TradeAccountDataSymbolCommissionResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

// CopyPointer
func (m TradeAccountDataSymbolCommissionResponseBuilder) CopyPointer(to TradeAccountDataSymbolCommissionResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}
