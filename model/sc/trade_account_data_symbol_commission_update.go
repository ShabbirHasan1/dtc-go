package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeAccountDataSymbolCommissionUpdateSize = 27

// {
//     Size                = TradeAccountDataSymbolCommissionUpdateSize  (27)
//     Type                = TRADE_ACCOUNT_DATA_SYMBOL_COMMISSION_UPDATE  (10120)
//     BaseSize            = TradeAccountDataSymbolCommissionUpdateSize  (27)
//     RequestID           = 0
//     IsDeleted           = false
//     TradeAccount        = ""
//     Symbol              = ""
//     TradeFeePerContract = 0
// }
var _TradeAccountDataSymbolCommissionUpdateDefault = []byte{27, 0, 136, 39, 27, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataSymbolCommissionUpdate struct {
	message.VLS
}

type TradeAccountDataSymbolCommissionUpdateBuilder struct {
	message.VLSBuilder
}

type TradeAccountDataSymbolCommissionUpdatePointer struct {
	message.VLSPointer
}

type TradeAccountDataSymbolCommissionUpdatePointerBuilder struct {
	message.VLSPointerBuilder
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
//     IsDeleted           = false
//     TradeAccount        = ""
//     Symbol              = ""
//     TradeFeePerContract = 0
// }
func (m TradeAccountDataSymbolCommissionUpdateBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataSymbolCommissionUpdateDefault)
}

// Clear
// {
//     Size                = TradeAccountDataSymbolCommissionUpdateSize  (27)
//     Type                = TRADE_ACCOUNT_DATA_SYMBOL_COMMISSION_UPDATE  (10120)
//     BaseSize            = TradeAccountDataSymbolCommissionUpdateSize  (27)
//     RequestID           = 0
//     IsDeleted           = false
//     TradeAccount        = ""
//     Symbol              = ""
//     TradeFeePerContract = 0
// }
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataSymbolCommissionUpdateDefault)
}

// ToBuilder
func (m TradeAccountDataSymbolCommissionUpdate) ToBuilder() TradeAccountDataSymbolCommissionUpdateBuilder {
	return TradeAccountDataSymbolCommissionUpdateBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradeAccountDataSymbolCommissionUpdatePointer) ToBuilder() TradeAccountDataSymbolCommissionUpdatePointerBuilder {
	return TradeAccountDataSymbolCommissionUpdatePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m TradeAccountDataSymbolCommissionUpdateBuilder) Finish() TradeAccountDataSymbolCommissionUpdate {
	return TradeAccountDataSymbolCommissionUpdate{m.VLSBuilder.Finish()}
}

// Finish
func (m *TradeAccountDataSymbolCommissionUpdatePointerBuilder) Finish() TradeAccountDataSymbolCommissionUpdatePointer {
	return TradeAccountDataSymbolCommissionUpdatePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataSymbolCommissionUpdate) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataSymbolCommissionUpdateBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
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
func (m TradeAccountDataSymbolCommissionUpdate) IsDeleted() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsDeleted
func (m TradeAccountDataSymbolCommissionUpdateBuilder) IsDeleted() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsDeleted
func (m TradeAccountDataSymbolCommissionUpdatePointer) IsDeleted() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// IsDeleted
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) IsDeleted() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// TradeAccount
func (m TradeAccountDataSymbolCommissionUpdate) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// TradeAccount
func (m TradeAccountDataSymbolCommissionUpdateBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
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
func (m TradeAccountDataSymbolCommissionUpdate) Symbol() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
}

// Symbol
func (m TradeAccountDataSymbolCommissionUpdateBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 19, 15)
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
func (m TradeAccountDataSymbolCommissionUpdate) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Unsafe(), 27, 19)
}

// TradeFeePerContract
func (m TradeAccountDataSymbolCommissionUpdateBuilder) TradeFeePerContract() float64 {
	return message.Float64VLS(m.Unsafe(), 27, 19)
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
func (m TradeAccountDataSymbolCommissionUpdateBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetIsDeleted
func (m TradeAccountDataSymbolCommissionUpdateBuilder) SetIsDeleted(value bool) {
	message.SetBoolVLS(m.Unsafe(), 11, 10, value)
}

// SetIsDeleted
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) SetIsDeleted(value bool) {
	message.SetBoolVLS(m.Ptr, 11, 10, value)
}

// SetTradeAccount
func (m *TradeAccountDataSymbolCommissionUpdateBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 15, 11, value)
}

// SetTradeAccount
func (m *TradeAccountDataSymbolCommissionUpdatePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetSymbol
func (m *TradeAccountDataSymbolCommissionUpdateBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 19, 15, value)
}

// SetSymbol
func (m *TradeAccountDataSymbolCommissionUpdatePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 19, 15, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataSymbolCommissionUpdateBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 27, 19, value)
}

// SetTradeFeePerContract
func (m TradeAccountDataSymbolCommissionUpdatePointerBuilder) SetTradeFeePerContract(value float64) {
	message.SetFloat64VLS(m.Ptr, 27, 19, value)
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
