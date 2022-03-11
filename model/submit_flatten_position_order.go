package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size             = SubmitFlattenPositionOrderSize  (28)
//     Type             = SUBMIT_FLATTEN_POSITION_ORDER  (209)
//     BaseSize         = SubmitFlattenPositionOrderSize  (28)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = 0
// }
var _SubmitFlattenPositionOrderDefault = []byte{28, 0, 209, 0, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SubmitFlattenPositionOrderSize = 28

type SubmitFlattenPositionOrder struct {
	message.VLS
}

type SubmitFlattenPositionOrderBuilder struct {
	message.VLSBuilder
}

func NewSubmitFlattenPositionOrderFrom(b []byte) SubmitFlattenPositionOrder {
	return SubmitFlattenPositionOrder{VLS: message.NewVLSFrom(b)}
}

func WrapSubmitFlattenPositionOrder(b []byte) SubmitFlattenPositionOrder {
	return SubmitFlattenPositionOrder{VLS: message.WrapVLS(b)}
}

func NewSubmitFlattenPositionOrder() SubmitFlattenPositionOrderBuilder {
	a := SubmitFlattenPositionOrderBuilder{message.NewVLSBuilder(28)}
	a.Unsafe().SetBytes(0, _SubmitFlattenPositionOrderDefault)
	return a
}

// Clear
// {
//     Size             = SubmitFlattenPositionOrderSize  (28)
//     Type             = SUBMIT_FLATTEN_POSITION_ORDER  (209)
//     BaseSize         = SubmitFlattenPositionOrderSize  (28)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = 0
// }
func (m SubmitFlattenPositionOrderBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitFlattenPositionOrderDefault)
}

// ToBuilder
func (m SubmitFlattenPositionOrder) ToBuilder() SubmitFlattenPositionOrderBuilder {
	return SubmitFlattenPositionOrderBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m SubmitFlattenPositionOrderBuilder) Finish() SubmitFlattenPositionOrder {
	return SubmitFlattenPositionOrder{m.VLSBuilder.Finish()}
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 26, 22)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderBuilder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 26, 22)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 27, 26)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 27, 26)
}

// SetSymbol The symbol of the Trade Position to flatten.
func (m *SubmitFlattenPositionOrderBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetExchange The optional exchange for the Symbol.
func (m *SubmitFlattenPositionOrderBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount The trade account as a text string of the Trade Position to flatten.
func (m *SubmitFlattenPositionOrderBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m *SubmitFlattenPositionOrderBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 22, 18, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m *SubmitFlattenPositionOrderBuilder) SetFreeFormText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 26, 22, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 27, 26, value)
}

// Copy
func (m SubmitFlattenPositionOrder) Copy(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m SubmitFlattenPositionOrderBuilder) Copy(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrder) CopyPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderBuilder) CopyPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrder) CopyToPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderBuilder) CopyToPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrder) CopyTo(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderBuilder) CopyTo(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}
