package model

import (
	"github.com/moontrade/dtc-go/message"
)

type FlattenPositionsForTradeAccountPointer struct {
	message.VLSPointer
}

type FlattenPositionsForTradeAccountPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocFlattenPositionsForTradeAccount() FlattenPositionsForTradeAccountPointerBuilder {
	a := FlattenPositionsForTradeAccountPointerBuilder{message.AllocVLSBuilder(19)}
	a.Ptr.SetBytes(0, _FlattenPositionsForTradeAccountDefault)
	return a
}

func AllocFlattenPositionsForTradeAccountFrom(b []byte) FlattenPositionsForTradeAccountPointer {
	return FlattenPositionsForTradeAccountPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size             = FlattenPositionsForTradeAccountSize  (19)
//     Type             = FLATTEN_POSITIONS_FOR_TRADE_ACCOUNT  (210)
//     BaseSize         = FlattenPositionsForTradeAccountSize  (19)
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = 0
// }
func (m FlattenPositionsForTradeAccountPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _FlattenPositionsForTradeAccountDefault)
}

// ToBuilder
func (m FlattenPositionsForTradeAccountPointer) ToBuilder() FlattenPositionsForTradeAccountPointerBuilder {
	return FlattenPositionsForTradeAccountPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *FlattenPositionsForTradeAccountPointerBuilder) Finish() FlattenPositionsForTradeAccountPointer {
	return FlattenPositionsForTradeAccountPointer{m.VLSPointerBuilder.Finish()}
}

// TradeAccount
func (m FlattenPositionsForTradeAccountPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m FlattenPositionsForTradeAccountPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// ClientOrderID
func (m FlattenPositionsForTradeAccountPointer) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// ClientOrderID
func (m FlattenPositionsForTradeAccountPointerBuilder) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// FreeFormText
func (m FlattenPositionsForTradeAccountPointer) FreeFormText() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// FreeFormText
func (m FlattenPositionsForTradeAccountPointerBuilder) FreeFormText() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// IsAutomatedOrder
func (m FlattenPositionsForTradeAccountPointer) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 19, 18)
}

// IsAutomatedOrder
func (m FlattenPositionsForTradeAccountPointerBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 19, 18)
}

// SetTradeAccount
func (m *FlattenPositionsForTradeAccountPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetClientOrderID
func (m *FlattenPositionsForTradeAccountPointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetFreeFormText
func (m *FlattenPositionsForTradeAccountPointerBuilder) SetFreeFormText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetIsAutomatedOrder
func (m FlattenPositionsForTradeAccountPointerBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8VLS(m.Ptr, 19, 18, value)
}

// Copy
func (m FlattenPositionsForTradeAccountPointer) Copy(to FlattenPositionsForTradeAccountBuilder) {
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m FlattenPositionsForTradeAccountPointerBuilder) Copy(to FlattenPositionsForTradeAccountBuilder) {
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m FlattenPositionsForTradeAccountPointer) CopyPointer(to FlattenPositionsForTradeAccountPointerBuilder) {
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m FlattenPositionsForTradeAccountPointerBuilder) CopyPointer(to FlattenPositionsForTradeAccountPointerBuilder) {
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}
