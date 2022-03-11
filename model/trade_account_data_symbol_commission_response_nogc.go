package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataSymbolCommissionResponsePointer struct {
	message.VLSPointer
}

type TradeAccountDataSymbolCommissionResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeAccountDataSymbolCommissionResponse() TradeAccountDataSymbolCommissionResponsePointerBuilder {
	a := TradeAccountDataSymbolCommissionResponsePointerBuilder{message.AllocVLSBuilder(26)}
	a.Ptr.SetBytes(0, _TradeAccountDataSymbolCommissionResponseDefault)
	return a
}

func AllocTradeAccountDataSymbolCommissionResponseFrom(b []byte) TradeAccountDataSymbolCommissionResponsePointer {
	return TradeAccountDataSymbolCommissionResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m TradeAccountDataSymbolCommissionResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataSymbolCommissionResponseDefault)
}

// ToBuilder
func (m TradeAccountDataSymbolCommissionResponsePointer) ToBuilder() TradeAccountDataSymbolCommissionResponsePointerBuilder {
	return TradeAccountDataSymbolCommissionResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataSymbolCommissionResponsePointerBuilder) Finish() TradeAccountDataSymbolCommissionResponsePointer {
	return TradeAccountDataSymbolCommissionResponsePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataSymbolCommissionResponsePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataSymbolCommissionResponsePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataSymbolCommissionResponsePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataSymbolCommissionResponsePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Symbol
func (m TradeAccountDataSymbolCommissionResponsePointer) Symbol() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Symbol
func (m TradeAccountDataSymbolCommissionResponsePointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TradeFeePerContract
func (m TradeAccountDataSymbolCommissionResponsePointer) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Ptr, 26, 18)
}

// TradeFeePerContract
func (m TradeAccountDataSymbolCommissionResponsePointerBuilder) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Ptr, 26, 18)
}

// SetRequestID
func (m TradeAccountDataSymbolCommissionResponsePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataSymbolCommissionResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetSymbol
func (m *TradeAccountDataSymbolCommissionResponsePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataSymbolCommissionResponsePointerBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Ptr, 26, 18, value)
}

// Copy
func (m TradeAccountDataSymbolCommissionResponsePointer) Copy(to TradeAccountDataSymbolCommissionResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

// Copy
func (m TradeAccountDataSymbolCommissionResponsePointerBuilder) Copy(to TradeAccountDataSymbolCommissionResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

// CopyPointer
func (m TradeAccountDataSymbolCommissionResponsePointer) CopyPointer(to TradeAccountDataSymbolCommissionResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

// CopyPointer
func (m TradeAccountDataSymbolCommissionResponsePointerBuilder) CopyPointer(to TradeAccountDataSymbolCommissionResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}
