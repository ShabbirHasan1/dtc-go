package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size             = FlattenPositionsForTradeAccountSize  (19)
//     Type             = FLATTEN_POSITIONS_FOR_TRADE_ACCOUNT  (210)
//     BaseSize         = FlattenPositionsForTradeAccountSize  (19)
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = 0
// }
var _FlattenPositionsForTradeAccountDefault = []byte{19, 0, 210, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const FlattenPositionsForTradeAccountSize = 19

type FlattenPositionsForTradeAccount struct {
	message.VLS
}

type FlattenPositionsForTradeAccountBuilder struct {
	message.VLSBuilder
}

func NewFlattenPositionsForTradeAccountFrom(b []byte) FlattenPositionsForTradeAccount {
	return FlattenPositionsForTradeAccount{VLS: message.NewVLSFrom(b)}
}

func WrapFlattenPositionsForTradeAccount(b []byte) FlattenPositionsForTradeAccount {
	return FlattenPositionsForTradeAccount{VLS: message.WrapVLS(b)}
}

func NewFlattenPositionsForTradeAccount() FlattenPositionsForTradeAccountBuilder {
	a := FlattenPositionsForTradeAccountBuilder{message.NewVLSBuilder(19)}
	a.Unsafe().SetBytes(0, _FlattenPositionsForTradeAccountDefault)
	return a
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
func (m FlattenPositionsForTradeAccountBuilder) Clear() {
	m.Unsafe().SetBytes(0, _FlattenPositionsForTradeAccountDefault)
}

// ToBuilder
func (m FlattenPositionsForTradeAccount) ToBuilder() FlattenPositionsForTradeAccountBuilder {
	return FlattenPositionsForTradeAccountBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m FlattenPositionsForTradeAccountBuilder) Finish() FlattenPositionsForTradeAccount {
	return FlattenPositionsForTradeAccount{m.VLSBuilder.Finish()}
}

// TradeAccount
func (m FlattenPositionsForTradeAccount) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m FlattenPositionsForTradeAccountBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// ClientOrderID
func (m FlattenPositionsForTradeAccount) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// ClientOrderID
func (m FlattenPositionsForTradeAccountBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// FreeFormText
func (m FlattenPositionsForTradeAccount) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// FreeFormText
func (m FlattenPositionsForTradeAccountBuilder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// IsAutomatedOrder
func (m FlattenPositionsForTradeAccount) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 19, 18)
}

// IsAutomatedOrder
func (m FlattenPositionsForTradeAccountBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 19, 18)
}

// SetTradeAccount
func (m *FlattenPositionsForTradeAccountBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetClientOrderID
func (m *FlattenPositionsForTradeAccountBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetFreeFormText
func (m *FlattenPositionsForTradeAccountBuilder) SetFreeFormText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetIsAutomatedOrder
func (m FlattenPositionsForTradeAccountBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 19, 18, value)
}

// Copy
func (m FlattenPositionsForTradeAccount) Copy(to FlattenPositionsForTradeAccountBuilder) {
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m FlattenPositionsForTradeAccountBuilder) Copy(to FlattenPositionsForTradeAccountBuilder) {
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m FlattenPositionsForTradeAccount) CopyPointer(to FlattenPositionsForTradeAccountPointerBuilder) {
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyPointer
func (m FlattenPositionsForTradeAccountBuilder) CopyPointer(to FlattenPositionsForTradeAccountPointerBuilder) {
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}
