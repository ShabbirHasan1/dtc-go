package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _AddCorrectingOrderFillFixedDefault = []byte{216, 0, 53, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AddCorrectingOrderFillFixedSize = 216

type AddCorrectingOrderFillFixed struct {
	message.Fixed
}

type AddCorrectingOrderFillFixedBuilder struct {
	message.Fixed
}

func NewAddCorrectingOrderFillFixedFrom(b []byte) AddCorrectingOrderFillFixed {
	return AddCorrectingOrderFillFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapAddCorrectingOrderFillFixed(b []byte) AddCorrectingOrderFillFixed {
	return AddCorrectingOrderFillFixed{Fixed: message.WrapFixed(b)}
}

func NewAddCorrectingOrderFillFixed() AddCorrectingOrderFillFixedBuilder {
	a := AddCorrectingOrderFillFixedBuilder{message.NewFixed(216)}
	a.Unsafe().SetBytes(0, _AddCorrectingOrderFillFixedDefault)
	return a
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
func (m AddCorrectingOrderFillFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AddCorrectingOrderFillFixedDefault)
}

// ToBuilder
func (m AddCorrectingOrderFillFixed) ToBuilder() AddCorrectingOrderFillFixedBuilder {
	return AddCorrectingOrderFillFixedBuilder{m.Fixed}
}

// Finish
func (m AddCorrectingOrderFillFixedBuilder) Finish() AddCorrectingOrderFillFixed {
	return AddCorrectingOrderFillFixed{m.Fixed}
}

// Symbol
func (m AddCorrectingOrderFillFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Symbol
func (m AddCorrectingOrderFillFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Exchange
func (m AddCorrectingOrderFillFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// Exchange
func (m AddCorrectingOrderFillFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// TradeAccount
func (m AddCorrectingOrderFillFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// TradeAccount
func (m AddCorrectingOrderFillFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// ClientOrderID
func (m AddCorrectingOrderFillFixed) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 148, 116)
}

// ClientOrderID
func (m AddCorrectingOrderFillFixedBuilder) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 148, 116)
}

// BuySell
func (m AddCorrectingOrderFillFixed) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 152, 148))
}

// BuySell
func (m AddCorrectingOrderFillFixedBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 152, 148))
}

// FillPrice
func (m AddCorrectingOrderFillFixed) FillPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 160, 152)
}

// FillPrice
func (m AddCorrectingOrderFillFixedBuilder) FillPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 160, 152)
}

// FillQuantity
func (m AddCorrectingOrderFillFixed) FillQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 168, 160)
}

// FillQuantity
func (m AddCorrectingOrderFillFixedBuilder) FillQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 168, 160)
}

// FreeFormText
func (m AddCorrectingOrderFillFixed) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 216, 168)
}

// FreeFormText
func (m AddCorrectingOrderFillFixedBuilder) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 216, 168)
}

// SetSymbol
func (m AddCorrectingOrderFillFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 68, 4, value)
}

// SetExchange
func (m AddCorrectingOrderFillFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 84, 68, value)
}

// SetTradeAccount
func (m AddCorrectingOrderFillFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 116, 84, value)
}

// SetClientOrderID
func (m AddCorrectingOrderFillFixedBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 148, 116, value)
}

// SetBuySell
func (m AddCorrectingOrderFillFixedBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 152, 148, int32(value))
}

// SetFillPrice
func (m AddCorrectingOrderFillFixedBuilder) SetFillPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 160, 152, value)
}

// SetFillQuantity
func (m AddCorrectingOrderFillFixedBuilder) SetFillQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 168, 160, value)
}

// SetFreeFormText
func (m AddCorrectingOrderFillFixedBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Unsafe(), 216, 168, value)
}

// Copy
func (m AddCorrectingOrderFillFixed) Copy(to AddCorrectingOrderFillFixedBuilder) {
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
func (m AddCorrectingOrderFillFixedBuilder) Copy(to AddCorrectingOrderFillFixedBuilder) {
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
func (m AddCorrectingOrderFillFixed) CopyPointer(to AddCorrectingOrderFillFixedPointerBuilder) {
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
func (m AddCorrectingOrderFillFixedBuilder) CopyPointer(to AddCorrectingOrderFillFixedPointerBuilder) {
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
func (m AddCorrectingOrderFillFixed) CopyToPointer(to AddCorrectingOrderFillPointerBuilder) {
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
func (m AddCorrectingOrderFillFixedBuilder) CopyToPointer(to AddCorrectingOrderFillPointerBuilder) {
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
func (m AddCorrectingOrderFillFixed) CopyTo(to AddCorrectingOrderFillBuilder) {
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
func (m AddCorrectingOrderFillFixedBuilder) CopyTo(to AddCorrectingOrderFillBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetBuySell(m.BuySell())
	to.SetFillPrice(m.FillPrice())
	to.SetFillQuantity(m.FillQuantity())
	to.SetFreeFormText(m.FreeFormText())
}
