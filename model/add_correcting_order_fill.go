package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = AddCorrectingOrderFillSize  (56)
//     Type          = ADD_CORRECTING_ORDER_FILL  (309)
//     BaseSize      = AddCorrectingOrderFillSize  (56)
//     Symbol        = ""
//     Exchange      = ""
//     TradeAccount  = ""
//     ClientOrderID = ""
//     BuySell       = BUY_SELL_UNSET  (0)
//     FillPrice     = 0
//     FillQuantity  = 0
//     FreeFormText  = ""
// }
var _AddCorrectingOrderFillDefault = []byte{56, 0, 53, 1, 56, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AddCorrectingOrderFillSize = 56

type AddCorrectingOrderFill struct {
	message.VLS
}

type AddCorrectingOrderFillBuilder struct {
	message.VLSBuilder
}

func NewAddCorrectingOrderFillFrom(b []byte) AddCorrectingOrderFill {
	return AddCorrectingOrderFill{VLS: message.NewVLSFrom(b)}
}

func WrapAddCorrectingOrderFill(b []byte) AddCorrectingOrderFill {
	return AddCorrectingOrderFill{VLS: message.WrapVLS(b)}
}

func NewAddCorrectingOrderFill() AddCorrectingOrderFillBuilder {
	a := AddCorrectingOrderFillBuilder{message.NewVLSBuilder(56)}
	a.Unsafe().SetBytes(0, _AddCorrectingOrderFillDefault)
	return a
}

// Clear
// {
//     Size          = AddCorrectingOrderFillSize  (56)
//     Type          = ADD_CORRECTING_ORDER_FILL  (309)
//     BaseSize      = AddCorrectingOrderFillSize  (56)
//     Symbol        = ""
//     Exchange      = ""
//     TradeAccount  = ""
//     ClientOrderID = ""
//     BuySell       = BUY_SELL_UNSET  (0)
//     FillPrice     = 0
//     FillQuantity  = 0
//     FreeFormText  = ""
// }
func (m AddCorrectingOrderFillBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AddCorrectingOrderFillDefault)
}

// ToBuilder
func (m AddCorrectingOrderFill) ToBuilder() AddCorrectingOrderFillBuilder {
	return AddCorrectingOrderFillBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m AddCorrectingOrderFillBuilder) Finish() AddCorrectingOrderFill {
	return AddCorrectingOrderFill{m.VLSBuilder.Finish()}
}

// Symbol
func (m AddCorrectingOrderFill) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol
func (m AddCorrectingOrderFillBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Exchange
func (m AddCorrectingOrderFill) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Exchange
func (m AddCorrectingOrderFillBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m AddCorrectingOrderFill) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount
func (m AddCorrectingOrderFillBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// ClientOrderID
func (m AddCorrectingOrderFill) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// ClientOrderID
func (m AddCorrectingOrderFillBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// BuySell
func (m AddCorrectingOrderFill) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// BuySell
func (m AddCorrectingOrderFillBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// FillPrice
func (m AddCorrectingOrderFill) FillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 40, 32)
}

// FillPrice
func (m AddCorrectingOrderFillBuilder) FillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 40, 32)
}

// FillQuantity
func (m AddCorrectingOrderFill) FillQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// FillQuantity
func (m AddCorrectingOrderFillBuilder) FillQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// FreeFormText
func (m AddCorrectingOrderFill) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 52, 48)
}

// FreeFormText
func (m AddCorrectingOrderFillBuilder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 52, 48)
}

// SetSymbol
func (m *AddCorrectingOrderFillBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetExchange
func (m *AddCorrectingOrderFillBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *AddCorrectingOrderFillBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetClientOrderID
func (m *AddCorrectingOrderFillBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 22, 18, value)
}

// SetBuySell
func (m AddCorrectingOrderFillBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, int32(value))
}

// SetFillPrice
func (m AddCorrectingOrderFillBuilder) SetFillPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 40, 32, value)
}

// SetFillQuantity
func (m AddCorrectingOrderFillBuilder) SetFillQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 48, 40, value)
}

// SetFreeFormText
func (m *AddCorrectingOrderFillBuilder) SetFreeFormText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 52, 48, value)
}

// Copy
func (m AddCorrectingOrderFill) Copy(to AddCorrectingOrderFillBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetBuySell(m.BuySell())
	to.SetFillPrice(m.FillPrice())
	to.SetFillQuantity(m.FillQuantity())
	to.SetFreeFormText(m.FreeFormText())
}

// Copy
func (m AddCorrectingOrderFillBuilder) Copy(to AddCorrectingOrderFillBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetBuySell(m.BuySell())
	to.SetFillPrice(m.FillPrice())
	to.SetFillQuantity(m.FillQuantity())
	to.SetFreeFormText(m.FreeFormText())
}

// CopyPointer
func (m AddCorrectingOrderFill) CopyPointer(to AddCorrectingOrderFillPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetBuySell(m.BuySell())
	to.SetFillPrice(m.FillPrice())
	to.SetFillQuantity(m.FillQuantity())
	to.SetFreeFormText(m.FreeFormText())
}

// CopyPointer
func (m AddCorrectingOrderFillBuilder) CopyPointer(to AddCorrectingOrderFillPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetBuySell(m.BuySell())
	to.SetFillPrice(m.FillPrice())
	to.SetFillQuantity(m.FillQuantity())
	to.SetFreeFormText(m.FreeFormText())
}

// CopyToPointer
func (m AddCorrectingOrderFill) CopyToPointer(to AddCorrectingOrderFillFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetBuySell(m.BuySell())
	to.SetFillPrice(m.FillPrice())
	to.SetFillQuantity(m.FillQuantity())
	to.SetFreeFormText(m.FreeFormText())
}

// CopyToPointer
func (m AddCorrectingOrderFillBuilder) CopyToPointer(to AddCorrectingOrderFillFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetBuySell(m.BuySell())
	to.SetFillPrice(m.FillPrice())
	to.SetFillQuantity(m.FillQuantity())
	to.SetFreeFormText(m.FreeFormText())
}

// CopyTo
func (m AddCorrectingOrderFill) CopyTo(to AddCorrectingOrderFillFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetBuySell(m.BuySell())
	to.SetFillPrice(m.FillPrice())
	to.SetFillQuantity(m.FillQuantity())
	to.SetFreeFormText(m.FreeFormText())
}

// CopyTo
func (m AddCorrectingOrderFillBuilder) CopyTo(to AddCorrectingOrderFillFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetBuySell(m.BuySell())
	to.SetFillPrice(m.FillPrice())
	to.SetFillQuantity(m.FillQuantity())
	to.SetFreeFormText(m.FreeFormText())
}
