package model

import (
	"github.com/moontrade/dtc-go/message"
)

type SubmitFlattenPositionOrderPointer struct {
	message.VLSPointer
}

type SubmitFlattenPositionOrderPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocSubmitFlattenPositionOrder() SubmitFlattenPositionOrderPointerBuilder {
	a := SubmitFlattenPositionOrderPointerBuilder{message.AllocVLSBuilder(28)}
	a.Ptr.SetBytes(0, _SubmitFlattenPositionOrderDefault)
	return a
}

func AllocSubmitFlattenPositionOrderFrom(b []byte) SubmitFlattenPositionOrderPointer {
	return SubmitFlattenPositionOrderPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m SubmitFlattenPositionOrderPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitFlattenPositionOrderDefault)
}

// ToBuilder
func (m SubmitFlattenPositionOrderPointer) ToBuilder() SubmitFlattenPositionOrderPointerBuilder {
	return SubmitFlattenPositionOrderPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *SubmitFlattenPositionOrderPointerBuilder) Finish() SubmitFlattenPositionOrderPointer {
	return SubmitFlattenPositionOrderPointer{m.VLSPointerBuilder.Finish()}
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderPointer) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderPointerBuilder) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderPointer) FreeFormText() string {
	return message.StringVLS(m.Ptr, 26, 22)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderPointerBuilder) FreeFormText() string {
	return message.StringVLS(m.Ptr, 26, 22)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderPointer) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 27, 26)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderPointerBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 27, 26)
}

// SetSymbol The symbol of the Trade Position to flatten.
func (m *SubmitFlattenPositionOrderPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetExchange The optional exchange for the Symbol.
func (m *SubmitFlattenPositionOrderPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetTradeAccount The trade account as a text string of the Trade Position to flatten.
func (m *SubmitFlattenPositionOrderPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m *SubmitFlattenPositionOrderPointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 22, 18, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m *SubmitFlattenPositionOrderPointerBuilder) SetFreeFormText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 26, 22, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderPointerBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8VLS(m.Ptr, 27, 26, value)
}

// Copy
func (m SubmitFlattenPositionOrderPointer) Copy(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m SubmitFlattenPositionOrderPointerBuilder) Copy(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderPointer) CopyPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderPointerBuilder) CopyPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderPointer) CopyTo(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderPointerBuilder) CopyTo(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderPointer) CopyToPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderPointerBuilder) CopyToPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}
