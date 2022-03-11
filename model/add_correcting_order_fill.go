package model

import (
	"github.com/moontrade/dtc-go/message"
)

const AddCorrectingOrderFillSize = 56

const AddCorrectingOrderFillFixedSize = 216

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

// {
//     Size          = AddCorrectingOrderFillFixedSize  (216)
//     Type          = ADD_CORRECTING_ORDER_FILL  (309)
//     Symbol        = ""
//     Exchange      = ""
//     TradeAccount  = ""
//     ClientOrderID = ""
//     BuySell       = BUY_SELL_UNSET  (0)
//     FillPrice     = 0
//     FillQuantity  = 0
//     FreeFormText  = ""
// }
var _AddCorrectingOrderFillFixedDefault = []byte{216, 0, 53, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type AddCorrectingOrderFill struct {
	message.VLS
}

type AddCorrectingOrderFillBuilder struct {
	message.VLSBuilder
}

type AddCorrectingOrderFillFixed struct {
	message.Fixed
}

type AddCorrectingOrderFillFixedBuilder struct {
	message.Fixed
}

type AddCorrectingOrderFillPointer struct {
	message.VLSPointer
}

type AddCorrectingOrderFillPointerBuilder struct {
	message.VLSPointerBuilder
}

type AddCorrectingOrderFillFixedPointer struct {
	message.FixedPointer
}

type AddCorrectingOrderFillFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocAddCorrectingOrderFill() AddCorrectingOrderFillPointerBuilder {
	a := AddCorrectingOrderFillPointerBuilder{message.AllocVLSBuilder(56)}
	a.Ptr.SetBytes(0, _AddCorrectingOrderFillDefault)
	return a
}

func AllocAddCorrectingOrderFillFrom(b []byte) AddCorrectingOrderFillPointer {
	return AddCorrectingOrderFillPointer{VLSPointer: message.AllocVLSFrom(b)}
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

// Clear
// {
//     Size          = AddCorrectingOrderFillFixedSize  (216)
//     Type          = ADD_CORRECTING_ORDER_FILL  (309)
//     Symbol        = ""
//     Exchange      = ""
//     TradeAccount  = ""
//     ClientOrderID = ""
//     BuySell       = BUY_SELL_UNSET  (0)
//     FillPrice     = 0
//     FillQuantity  = 0
//     FreeFormText  = ""
// }
func (m AddCorrectingOrderFillFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AddCorrectingOrderFillFixedDefault)
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

// Clear
// {
//     Size          = AddCorrectingOrderFillFixedSize  (216)
//     Type          = ADD_CORRECTING_ORDER_FILL  (309)
//     Symbol        = ""
//     Exchange      = ""
//     TradeAccount  = ""
//     ClientOrderID = ""
//     BuySell       = BUY_SELL_UNSET  (0)
//     FillPrice     = 0
//     FillQuantity  = 0
//     FreeFormText  = ""
// }
func (m AddCorrectingOrderFillFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AddCorrectingOrderFillFixedDefault)
}

// ToBuilder
func (m AddCorrectingOrderFill) ToBuilder() AddCorrectingOrderFillBuilder {
	return AddCorrectingOrderFillBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m AddCorrectingOrderFillFixed) ToBuilder() AddCorrectingOrderFillFixedBuilder {
	return AddCorrectingOrderFillFixedBuilder{m.Fixed}
}

// ToBuilder
func (m AddCorrectingOrderFillPointer) ToBuilder() AddCorrectingOrderFillPointerBuilder {
	return AddCorrectingOrderFillPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m AddCorrectingOrderFillFixedPointer) ToBuilder() AddCorrectingOrderFillFixedPointerBuilder {
	return AddCorrectingOrderFillFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m AddCorrectingOrderFillBuilder) Finish() AddCorrectingOrderFill {
	return AddCorrectingOrderFill{m.VLSBuilder.Finish()}
}

// Finish
func (m AddCorrectingOrderFillFixedBuilder) Finish() AddCorrectingOrderFillFixed {
	return AddCorrectingOrderFillFixed{m.Fixed}
}

// Finish
func (m *AddCorrectingOrderFillPointerBuilder) Finish() AddCorrectingOrderFillPointer {
	return AddCorrectingOrderFillPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *AddCorrectingOrderFillFixedPointerBuilder) Finish() AddCorrectingOrderFillFixedPointer {
	return AddCorrectingOrderFillFixedPointer{m.FixedPointer}
}

// Symbol
func (m AddCorrectingOrderFill) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol
func (m AddCorrectingOrderFillBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
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
func (m AddCorrectingOrderFill) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Exchange
func (m AddCorrectingOrderFillBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
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
func (m AddCorrectingOrderFill) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount
func (m AddCorrectingOrderFillBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
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
func (m AddCorrectingOrderFill) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// ClientOrderID
func (m AddCorrectingOrderFillBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
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
func (m AddCorrectingOrderFill) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// BuySell
func (m AddCorrectingOrderFillBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 28, 24))
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
func (m AddCorrectingOrderFill) FillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 40, 32)
}

// FillPrice
func (m AddCorrectingOrderFillBuilder) FillPrice() float64 {
	return message.Float64VLS(m.Unsafe(), 40, 32)
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
func (m AddCorrectingOrderFill) FillQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// FillQuantity
func (m AddCorrectingOrderFillBuilder) FillQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
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
func (m AddCorrectingOrderFill) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 52, 48)
}

// FreeFormText
func (m AddCorrectingOrderFillBuilder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 52, 48)
}

// FreeFormText
func (m AddCorrectingOrderFillPointer) FreeFormText() string {
	return message.StringVLS(m.Ptr, 52, 48)
}

// FreeFormText
func (m AddCorrectingOrderFillPointerBuilder) FreeFormText() string {
	return message.StringVLS(m.Ptr, 52, 48)
}

// Symbol
func (m AddCorrectingOrderFillFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Symbol
func (m AddCorrectingOrderFillFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
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
func (m AddCorrectingOrderFillFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// Exchange
func (m AddCorrectingOrderFillFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
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
func (m AddCorrectingOrderFillFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// TradeAccount
func (m AddCorrectingOrderFillFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
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
func (m AddCorrectingOrderFillFixed) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 148, 116)
}

// ClientOrderID
func (m AddCorrectingOrderFillFixedBuilder) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 148, 116)
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
func (m AddCorrectingOrderFillFixed) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 152, 148))
}

// BuySell
func (m AddCorrectingOrderFillFixedBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 152, 148))
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
func (m AddCorrectingOrderFillFixed) FillPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 160, 152)
}

// FillPrice
func (m AddCorrectingOrderFillFixedBuilder) FillPrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 160, 152)
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
func (m AddCorrectingOrderFillFixed) FillQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 168, 160)
}

// FillQuantity
func (m AddCorrectingOrderFillFixedBuilder) FillQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 168, 160)
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
func (m AddCorrectingOrderFillFixed) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 216, 168)
}

// FreeFormText
func (m AddCorrectingOrderFillFixedBuilder) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 216, 168)
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
func (m *AddCorrectingOrderFillBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetSymbol
func (m *AddCorrectingOrderFillPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetExchange
func (m *AddCorrectingOrderFillBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetExchange
func (m *AddCorrectingOrderFillPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *AddCorrectingOrderFillBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetTradeAccount
func (m *AddCorrectingOrderFillPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetClientOrderID
func (m *AddCorrectingOrderFillBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 22, 18, value)
}

// SetClientOrderID
func (m *AddCorrectingOrderFillPointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 22, 18, value)
}

// SetBuySell
func (m AddCorrectingOrderFillBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, int32(value))
}

// SetBuySell
func (m AddCorrectingOrderFillPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 28, 24, int32(value))
}

// SetFillPrice
func (m AddCorrectingOrderFillBuilder) SetFillPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 40, 32, value)
}

// SetFillPrice
func (m AddCorrectingOrderFillPointerBuilder) SetFillPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 40, 32, value)
}

// SetFillQuantity
func (m AddCorrectingOrderFillBuilder) SetFillQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 48, 40, value)
}

// SetFillQuantity
func (m AddCorrectingOrderFillPointerBuilder) SetFillQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 48, 40, value)
}

// SetFreeFormText
func (m *AddCorrectingOrderFillBuilder) SetFreeFormText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 52, 48, value)
}

// SetFreeFormText
func (m *AddCorrectingOrderFillPointerBuilder) SetFreeFormText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 52, 48, value)
}

// SetSymbol
func (m AddCorrectingOrderFillFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 68, 4, value)
}

// SetSymbol
func (m AddCorrectingOrderFillFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 68, 4, value)
}

// SetExchange
func (m AddCorrectingOrderFillFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 84, 68, value)
}

// SetExchange
func (m AddCorrectingOrderFillFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 84, 68, value)
}

// SetTradeAccount
func (m AddCorrectingOrderFillFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 116, 84, value)
}

// SetTradeAccount
func (m AddCorrectingOrderFillFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 116, 84, value)
}

// SetClientOrderID
func (m AddCorrectingOrderFillFixedBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 148, 116, value)
}

// SetClientOrderID
func (m AddCorrectingOrderFillFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 148, 116, value)
}

// SetBuySell
func (m AddCorrectingOrderFillFixedBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 152, 148, int32(value))
}

// SetBuySell
func (m AddCorrectingOrderFillFixedPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 152, 148, int32(value))
}

// SetFillPrice
func (m AddCorrectingOrderFillFixedBuilder) SetFillPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 160, 152, value)
}

// SetFillPrice
func (m AddCorrectingOrderFillFixedPointerBuilder) SetFillPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 160, 152, value)
}

// SetFillQuantity
func (m AddCorrectingOrderFillFixedBuilder) SetFillQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 168, 160, value)
}

// SetFillQuantity
func (m AddCorrectingOrderFillFixedPointerBuilder) SetFillQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 168, 160, value)
}

// SetFreeFormText
func (m AddCorrectingOrderFillFixedBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Unsafe(), 216, 168, value)
}

// SetFreeFormText
func (m AddCorrectingOrderFillFixedPointerBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Ptr, 216, 168, value)
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