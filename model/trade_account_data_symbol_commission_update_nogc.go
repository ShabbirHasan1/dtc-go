package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataSymbolCommissionUpdatePointer struct {
	message.VLSPointer
}

type TradeAccountDataSymbolCommissionUpdatePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeAccountDataSymbolCommissionUpdate() TradeAccountDataSymbolCommissionUpdatePointerBuilder {
	a := TradeAccountDataSymbolCommissionUpdatePointerBuilder{message.AllocVLSBuilder(27)}
	a.Ptr.SetBytes(0, _TradeAccountDataSymbolCommissionUpdateDefault)
	return a
}

func AllocTradeAccountDataSymbolCommissionUpdateFrom(b []byte) TradeAccountDataSymbolCommissionUpdatePointer {
	return TradeAccountDataSymbolCommissionUpdatePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                = TradeAccountDataSymbolCommissionUpdateSize  (27)
//     Type                = TRADE_ACCOUNT_DATA_SYMBOL_COMMISSION_UPDATE  (10120)
//     BaseSize            = TradeAccountDataSymbolCommissionUpdateSize  (27)
//     RequestID           = 0
//     IsDeleted           = 0
//     TradeAccount        = ""
//     Symbol              = ""
//     TradeFeePerContract = 0
// }
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataSymbolCommissionUpdateDefault)
}

// ToBuilder
func (m TradeAccountDataSymbolCommissionUpdatePointer) ToBuilder() TradeAccountDataSymbolCommissionUpdatePointerBuilder {
	return TradeAccountDataSymbolCommissionUpdatePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataSymbolCommissionUpdatePointerBuilder) Finish() TradeAccountDataSymbolCommissionUpdatePointer {
	return TradeAccountDataSymbolCommissionUpdatePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataSymbolCommissionUpdatePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// IsDeleted
func (m TradeAccountDataSymbolCommissionUpdatePointer) IsDeleted() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// IsDeleted
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) IsDeleted() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// TradeAccount
func (m TradeAccountDataSymbolCommissionUpdatePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// TradeAccount
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// Symbol
func (m TradeAccountDataSymbolCommissionUpdatePointer) Symbol() string {
	return message.StringVLS(m.Ptr, 19, 15)
}

// Symbol
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 19, 15)
}

// TradeFeePerContract
func (m TradeAccountDataSymbolCommissionUpdatePointer) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Ptr, 27, 19)
}

// TradeFeePerContract
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Ptr, 27, 19)
}

// SetRequestID
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetIsDeleted
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) SetIsDeleted(value uint8) {
	message.SetUint8VLS(m.Ptr, 11, 10, value)
}

// SetTradeAccount
func (m *TradeAccountDataSymbolCommissionUpdatePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetSymbol
func (m *TradeAccountDataSymbolCommissionUpdatePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 19, 15, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Ptr, 27, 19, value)
}

// Copy
func (m TradeAccountDataSymbolCommissionUpdatePointer) Copy(to TradeAccountDataSymbolCommissionUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

// Copy
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) Copy(to TradeAccountDataSymbolCommissionUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

// CopyPointer
func (m TradeAccountDataSymbolCommissionUpdatePointer) CopyPointer(to TradeAccountDataSymbolCommissionUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

// CopyPointer
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) CopyPointer(to TradeAccountDataSymbolCommissionUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}
