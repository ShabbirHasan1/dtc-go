package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size             = SubmitFlattenPositionOrderFixedSize  (198)
//     Type             = SUBMIT_FLATTEN_POSITION_ORDER  (209)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = 0
// }
var _SubmitFlattenPositionOrderFixedDefault = []byte{198, 0, 209, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SubmitFlattenPositionOrderFixedSize = 198

type SubmitFlattenPositionOrderFixed struct {
	message.Fixed
}

type SubmitFlattenPositionOrderFixedBuilder struct {
	message.Fixed
}

func NewSubmitFlattenPositionOrderFixedFrom(b []byte) SubmitFlattenPositionOrderFixed {
	return SubmitFlattenPositionOrderFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapSubmitFlattenPositionOrderFixed(b []byte) SubmitFlattenPositionOrderFixed {
	return SubmitFlattenPositionOrderFixed{Fixed: message.WrapFixed(b)}
}

func NewSubmitFlattenPositionOrderFixed() SubmitFlattenPositionOrderFixedBuilder {
	a := SubmitFlattenPositionOrderFixedBuilder{message.NewFixed(198)}
	a.Unsafe().SetBytes(0, _SubmitFlattenPositionOrderFixedDefault)
	return a
}

// Clear
// {
//     Size             = SubmitFlattenPositionOrderFixedSize  (198)
//     Type             = SUBMIT_FLATTEN_POSITION_ORDER  (209)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = 0
// }
func (m SubmitFlattenPositionOrderFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitFlattenPositionOrderFixedDefault)
}

// ToBuilder
func (m SubmitFlattenPositionOrderFixed) ToBuilder() SubmitFlattenPositionOrderFixedBuilder {
	return SubmitFlattenPositionOrderFixedBuilder{m.Fixed}
}

// Finish
func (m SubmitFlattenPositionOrderFixedBuilder) Finish() SubmitFlattenPositionOrderFixed {
	return SubmitFlattenPositionOrderFixed{m.Fixed}
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixed) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 148, 116)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixedBuilder) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 148, 116)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixed) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 196, 148)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixedBuilder) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 196, 148)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixed) IsAutomatedOrder() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 197, 196)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixedBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 197, 196)
}

// SetSymbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 68, 4, value)
}

// SetExchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 84, 68, value)
}

// SetTradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 116, 84, value)
}

// SetClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixedBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 148, 116, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixedBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Unsafe(), 196, 148, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixedBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 197, 196, value)
}

// Copy
func (m SubmitFlattenPositionOrderFixed) Copy(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m SubmitFlattenPositionOrderFixedBuilder) Copy(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderFixed) CopyPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderFixedBuilder) CopyPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderFixed) CopyToPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderFixedBuilder) CopyToPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderFixed) CopyTo(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderFixedBuilder) CopyTo(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}
