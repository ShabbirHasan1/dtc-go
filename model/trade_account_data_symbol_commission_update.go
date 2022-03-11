package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _TradeAccountDataSymbolCommissionUpdateDefault = []byte{27, 0, 136, 39, 27, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataSymbolCommissionUpdateSize = 27

type TradeAccountDataSymbolCommissionUpdate struct {
	message.VLS
}

type TradeAccountDataSymbolCommissionUpdateBuilder struct {
	message.VLSBuilder
}

func NewTradeAccountDataSymbolCommissionUpdateFrom(b []byte) TradeAccountDataSymbolCommissionUpdate {
	return TradeAccountDataSymbolCommissionUpdate{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataSymbolCommissionUpdate(b []byte) TradeAccountDataSymbolCommissionUpdate {
	return TradeAccountDataSymbolCommissionUpdate{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataSymbolCommissionUpdate() TradeAccountDataSymbolCommissionUpdateBuilder {
	a := TradeAccountDataSymbolCommissionUpdateBuilder{message.NewVLSBuilder(27)}
	a.Unsafe().SetBytes(0, _TradeAccountDataSymbolCommissionUpdateDefault)
	return a
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
func (m TradeAccountDataSymbolCommissionUpdateBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataSymbolCommissionUpdateDefault)
}

// ToBuilder
func (m TradeAccountDataSymbolCommissionUpdate) ToBuilder() TradeAccountDataSymbolCommissionUpdateBuilder {
	return TradeAccountDataSymbolCommissionUpdateBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataSymbolCommissionUpdateBuilder) Finish() TradeAccountDataSymbolCommissionUpdate {
	return TradeAccountDataSymbolCommissionUpdate{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataSymbolCommissionUpdate) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataSymbolCommissionUpdateBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// IsDeleted
func (m TradeAccountDataSymbolCommissionUpdate) IsDeleted() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// IsDeleted
func (m TradeAccountDataSymbolCommissionUpdateBuilder) IsDeleted() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// TradeAccount
func (m TradeAccountDataSymbolCommissionUpdate) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// TradeAccount
func (m TradeAccountDataSymbolCommissionUpdateBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// Symbol
func (m TradeAccountDataSymbolCommissionUpdate) Symbol() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
}

// Symbol
func (m TradeAccountDataSymbolCommissionUpdateBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
}

// TradeFeePerContract
func (m TradeAccountDataSymbolCommissionUpdate) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Unsafe(), 27, 19)
}

// TradeFeePerContract
func (m TradeAccountDataSymbolCommissionUpdateBuilder) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Unsafe(), 27, 19)
}

// SetRequestID
func (m TradeAccountDataSymbolCommissionUpdateBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetIsDeleted
func (m TradeAccountDataSymbolCommissionUpdateBuilder) SetIsDeleted(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 11, 10, value)
}

// SetTradeAccount
func (m *TradeAccountDataSymbolCommissionUpdateBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 15, 11, value)
}

// SetSymbol
func (m *TradeAccountDataSymbolCommissionUpdateBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 19, 15, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataSymbolCommissionUpdateBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 27, 19, value)
}

// Copy
func (m TradeAccountDataSymbolCommissionUpdate) Copy(to TradeAccountDataSymbolCommissionUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

// Copy
func (m TradeAccountDataSymbolCommissionUpdateBuilder) Copy(to TradeAccountDataSymbolCommissionUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

// CopyPointer
func (m TradeAccountDataSymbolCommissionUpdate) CopyPointer(to TradeAccountDataSymbolCommissionUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}

// CopyPointer
func (m TradeAccountDataSymbolCommissionUpdateBuilder) CopyPointer(to TradeAccountDataSymbolCommissionUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsDeleted(m.IsDeleted())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSymbol(m.Symbol())
	to.SetTradeFeePerContract(m.TradeFeePerContract())
}
