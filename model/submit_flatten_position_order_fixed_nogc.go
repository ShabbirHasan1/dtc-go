package model

import (
	"github.com/moontrade/dtc-go/message"
)

type SubmitFlattenPositionOrderFixedPointer struct {
	message.FixedPointer
}

type SubmitFlattenPositionOrderFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocSubmitFlattenPositionOrderFixed() SubmitFlattenPositionOrderFixedPointerBuilder {
	a := SubmitFlattenPositionOrderFixedPointerBuilder{message.AllocFixed(198)}
	a.Ptr.SetBytes(0, _SubmitFlattenPositionOrderFixedDefault)
	return a
}

func AllocSubmitFlattenPositionOrderFixedFrom(b []byte) SubmitFlattenPositionOrderFixedPointer {
	return SubmitFlattenPositionOrderFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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
func (m SubmitFlattenPositionOrderFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitFlattenPositionOrderFixedDefault)
}

// ToBuilder
func (m SubmitFlattenPositionOrderFixedPointer) ToBuilder() SubmitFlattenPositionOrderFixedPointerBuilder {
	return SubmitFlattenPositionOrderFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *SubmitFlattenPositionOrderFixedPointerBuilder) Finish() SubmitFlattenPositionOrderFixedPointer {
	return SubmitFlattenPositionOrderFixedPointer{m.FixedPointer}
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 116, 84)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 116, 84)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixedPointer) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 148, 116)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 148, 116)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixedPointer) FreeFormText() string {
	return message.StringFixed(m.Ptr, 196, 148)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) FreeFormText() string {
	return message.StringFixed(m.Ptr, 196, 148)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixedPointer) IsAutomatedOrder() uint8 {
	return message.Uint8Fixed(m.Ptr, 197, 196)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8Fixed(m.Ptr, 197, 196)
}

// SetSymbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 68, 4, value)
}

// SetExchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 84, 68, value)
}

// SetTradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 116, 84, value)
}

// SetClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 148, 116, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Ptr, 196, 148, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixedPointerBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8Fixed(m.Ptr, 197, 196, value)
}

// Copy
func (m SubmitFlattenPositionOrderFixedPointer) Copy(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m SubmitFlattenPositionOrderFixedPointerBuilder) Copy(to SubmitFlattenPositionOrderFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderFixedPointer) CopyPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m SubmitFlattenPositionOrderFixedPointerBuilder) CopyPointer(to SubmitFlattenPositionOrderFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderFixedPointer) CopyTo(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderFixedPointerBuilder) CopyTo(to SubmitFlattenPositionOrderBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderFixedPointer) CopyToPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyToPointer
func (m SubmitFlattenPositionOrderFixedPointerBuilder) CopyToPointer(to SubmitFlattenPositionOrderPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}
