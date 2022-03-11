package model

import (
	"github.com/moontrade/dtc-go/message"
)

type AddCorrectingOrderFillPointer struct {
	message.VLSPointer
}

type AddCorrectingOrderFillPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocAddCorrectingOrderFill() AddCorrectingOrderFillPointerBuilder {
	a := AddCorrectingOrderFillPointerBuilder{message.AllocVLSBuilder(56)}
	a.Ptr.SetBytes(0, _AddCorrectingOrderFillDefault)
	return a
}

func AllocAddCorrectingOrderFillFrom(b []byte) AddCorrectingOrderFillPointer {
	return AddCorrectingOrderFillPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m AddCorrectingOrderFillPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AddCorrectingOrderFillDefault)
}

// ToBuilder
func (m AddCorrectingOrderFillPointer) ToBuilder() AddCorrectingOrderFillPointerBuilder {
	return AddCorrectingOrderFillPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *AddCorrectingOrderFillPointerBuilder) Finish() AddCorrectingOrderFillPointer {
	return AddCorrectingOrderFillPointer{m.VLSPointerBuilder.Finish()}
}

// Symbol
func (m AddCorrectingOrderFillPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Symbol
func (m AddCorrectingOrderFillPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Exchange
func (m AddCorrectingOrderFillPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Exchange
func (m AddCorrectingOrderFillPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m AddCorrectingOrderFillPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TradeAccount
func (m AddCorrectingOrderFillPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// ClientOrderID
func (m AddCorrectingOrderFillPointer) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// ClientOrderID
func (m AddCorrectingOrderFillPointerBuilder) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// BuySell
func (m AddCorrectingOrderFillPointer) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 28, 24))
}

// BuySell
func (m AddCorrectingOrderFillPointerBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 28, 24))
}

// FillPrice
func (m AddCorrectingOrderFillPointer) FillPrice() float64 {
	return message.Float64VLS(m.Ptr, 40, 32)
}

// FillPrice
func (m AddCorrectingOrderFillPointerBuilder) FillPrice() float64 {
	return message.Float64VLS(m.Ptr, 40, 32)
}

// FillQuantity
func (m AddCorrectingOrderFillPointer) FillQuantity() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
}

// FillQuantity
func (m AddCorrectingOrderFillPointerBuilder) FillQuantity() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
}

// FreeFormText
func (m AddCorrectingOrderFillPointer) FreeFormText() string {
	return message.StringVLS(m.Ptr, 52, 48)
}

// FreeFormText
func (m AddCorrectingOrderFillPointerBuilder) FreeFormText() string {
	return message.StringVLS(m.Ptr, 52, 48)
}

// SetSymbol
func (m *AddCorrectingOrderFillPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetExchange
func (m *AddCorrectingOrderFillPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *AddCorrectingOrderFillPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetClientOrderID
func (m *AddCorrectingOrderFillPointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 22, 18, value)
}

// SetBuySell
func (m AddCorrectingOrderFillPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 28, 24, int32(value))
}

// SetFillPrice
func (m AddCorrectingOrderFillPointerBuilder) SetFillPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 40, 32, value)
}

// SetFillQuantity
func (m AddCorrectingOrderFillPointerBuilder) SetFillQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 48, 40, value)
}

// SetFreeFormText
func (m *AddCorrectingOrderFillPointerBuilder) SetFreeFormText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 52, 48, value)
}

// Copy
func (m AddCorrectingOrderFillPointer) Copy(to AddCorrectingOrderFillBuilder) {
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
func (m AddCorrectingOrderFillPointerBuilder) Copy(to AddCorrectingOrderFillBuilder) {
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
func (m AddCorrectingOrderFillPointer) CopyPointer(to AddCorrectingOrderFillPointerBuilder) {
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
func (m AddCorrectingOrderFillPointerBuilder) CopyPointer(to AddCorrectingOrderFillPointerBuilder) {
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
func (m AddCorrectingOrderFillPointer) CopyTo(to AddCorrectingOrderFillFixedBuilder) {
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
func (m AddCorrectingOrderFillPointerBuilder) CopyTo(to AddCorrectingOrderFillFixedBuilder) {
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
func (m AddCorrectingOrderFillPointer) CopyToPointer(to AddCorrectingOrderFillFixedPointerBuilder) {
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
func (m AddCorrectingOrderFillPointerBuilder) CopyToPointer(to AddCorrectingOrderFillFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetBuySell(m.BuySell())
	to.SetFillPrice(m.FillPrice())
	to.SetFillQuantity(m.FillQuantity())
	to.SetFreeFormText(m.FreeFormText())
}
