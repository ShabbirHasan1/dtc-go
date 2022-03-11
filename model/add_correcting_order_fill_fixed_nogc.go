package model

import (
	"github.com/moontrade/dtc-go/message"
)

type AddCorrectingOrderFillFixedPointer struct {
	message.FixedPointer
}

type AddCorrectingOrderFillFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocAddCorrectingOrderFillFixed() AddCorrectingOrderFillFixedPointerBuilder {
	a := AddCorrectingOrderFillFixedPointerBuilder{message.AllocFixed(216)}
	a.Ptr.SetBytes(0, _AddCorrectingOrderFillFixedDefault)
	return a
}

func AllocAddCorrectingOrderFillFixedFrom(b []byte) AddCorrectingOrderFillFixedPointer {
	return AddCorrectingOrderFillFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = AddCorrectingOrderFillFixedSize  (216)
//     Type          = ADD_CORRECTING_ORDER_FILL  (309)
//     Symbol        = ""
//     Exchange      = ""
//     TradeAccount  = ""
//     ClientOrderID = ""
//     BuySell       = 0
//     FillPrice     = 0
//     FillQuantity  = 0
//     FreeFormText  = ""
// }
func (m AddCorrectingOrderFillFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AddCorrectingOrderFillFixedDefault)
}

// ToBuilder
func (m AddCorrectingOrderFillFixedPointer) ToBuilder() AddCorrectingOrderFillFixedPointerBuilder {
	return AddCorrectingOrderFillFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *AddCorrectingOrderFillFixedPointerBuilder) Finish() AddCorrectingOrderFillFixedPointer {
	return AddCorrectingOrderFillFixedPointer{m.FixedPointer}
}

// Symbol
func (m AddCorrectingOrderFillFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Symbol
func (m AddCorrectingOrderFillFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Exchange
func (m AddCorrectingOrderFillFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// Exchange
func (m AddCorrectingOrderFillFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// TradeAccount
func (m AddCorrectingOrderFillFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 116, 84)
}

// TradeAccount
func (m AddCorrectingOrderFillFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 116, 84)
}

// ClientOrderID
func (m AddCorrectingOrderFillFixedPointer) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 148, 116)
}

// ClientOrderID
func (m AddCorrectingOrderFillFixedPointerBuilder) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 148, 116)
}

// BuySell
func (m AddCorrectingOrderFillFixedPointer) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 152, 148))
}

// BuySell
func (m AddCorrectingOrderFillFixedPointerBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 152, 148))
}

// FillPrice
func (m AddCorrectingOrderFillFixedPointer) FillPrice() float64 {
	return message.Float64Fixed(m.Ptr, 160, 152)
}

// FillPrice
func (m AddCorrectingOrderFillFixedPointerBuilder) FillPrice() float64 {
	return message.Float64Fixed(m.Ptr, 160, 152)
}

// FillQuantity
func (m AddCorrectingOrderFillFixedPointer) FillQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 168, 160)
}

// FillQuantity
func (m AddCorrectingOrderFillFixedPointerBuilder) FillQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 168, 160)
}

// FreeFormText
func (m AddCorrectingOrderFillFixedPointer) FreeFormText() string {
	return message.StringFixed(m.Ptr, 216, 168)
}

// FreeFormText
func (m AddCorrectingOrderFillFixedPointerBuilder) FreeFormText() string {
	return message.StringFixed(m.Ptr, 216, 168)
}

// SetSymbol
func (m AddCorrectingOrderFillFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 68, 4, value)
}

// SetExchange
func (m AddCorrectingOrderFillFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 84, 68, value)
}

// SetTradeAccount
func (m AddCorrectingOrderFillFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 116, 84, value)
}

// SetClientOrderID
func (m AddCorrectingOrderFillFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 148, 116, value)
}

// SetBuySell
func (m AddCorrectingOrderFillFixedPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 152, 148, int32(value))
}

// SetFillPrice
func (m AddCorrectingOrderFillFixedPointerBuilder) SetFillPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 160, 152, value)
}

// SetFillQuantity
func (m AddCorrectingOrderFillFixedPointerBuilder) SetFillQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 168, 160, value)
}

// SetFreeFormText
func (m AddCorrectingOrderFillFixedPointerBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Ptr, 216, 168, value)
}

// Copy
func (m AddCorrectingOrderFillFixedPointer) Copy(to AddCorrectingOrderFillFixedBuilder) {
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
func (m AddCorrectingOrderFillFixedPointerBuilder) Copy(to AddCorrectingOrderFillFixedBuilder) {
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
func (m AddCorrectingOrderFillFixedPointer) CopyPointer(to AddCorrectingOrderFillFixedPointerBuilder) {
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
func (m AddCorrectingOrderFillFixedPointerBuilder) CopyPointer(to AddCorrectingOrderFillFixedPointerBuilder) {
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
func (m AddCorrectingOrderFillFixedPointer) CopyTo(to AddCorrectingOrderFillBuilder) {
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
func (m AddCorrectingOrderFillFixedPointerBuilder) CopyTo(to AddCorrectingOrderFillBuilder) {
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
func (m AddCorrectingOrderFillFixedPointer) CopyToPointer(to AddCorrectingOrderFillPointerBuilder) {
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
func (m AddCorrectingOrderFillFixedPointerBuilder) CopyToPointer(to AddCorrectingOrderFillPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetBuySell(m.BuySell())
	to.SetFillPrice(m.FillPrice())
	to.SetFillQuantity(m.FillQuantity())
	to.SetFreeFormText(m.FreeFormText())
}
