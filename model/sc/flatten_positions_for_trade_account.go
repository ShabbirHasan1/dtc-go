package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const FlattenPositionsForTradeAccountSize = 19

// {
//     Size             = FlattenPositionsForTradeAccountSize  (19)
//     Type             = FLATTEN_POSITIONS_FOR_TRADE_ACCOUNT  (210)
//     BaseSize         = FlattenPositionsForTradeAccountSize  (19)
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = false
// }
var _FlattenPositionsForTradeAccountDefault = []byte{19, 0, 210, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type FlattenPositionsForTradeAccount struct {
	message.VLS
}

type FlattenPositionsForTradeAccountBuilder struct {
	message.VLSBuilder
}

type FlattenPositionsForTradeAccountPointer struct {
	message.VLSPointer
}

type FlattenPositionsForTradeAccountPointerBuilder struct {
	message.VLSPointerBuilder
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
//     IsAutomatedOrder = false
// }
func (m FlattenPositionsForTradeAccountBuilder) Clear() {
	m.Unsafe().SetBytes(0, _FlattenPositionsForTradeAccountDefault)
}

// Clear
// {
//     Size             = FlattenPositionsForTradeAccountSize  (19)
//     Type             = FLATTEN_POSITIONS_FOR_TRADE_ACCOUNT  (210)
//     BaseSize         = FlattenPositionsForTradeAccountSize  (19)
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     FreeFormText     = ""
//     IsAutomatedOrder = false
// }
func (m FlattenPositionsForTradeAccountPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _FlattenPositionsForTradeAccountDefault)
}

// ToBuilder
func (m FlattenPositionsForTradeAccount) ToBuilder() FlattenPositionsForTradeAccountBuilder {
	return FlattenPositionsForTradeAccountBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m FlattenPositionsForTradeAccountPointer) ToBuilder() FlattenPositionsForTradeAccountPointerBuilder {
	return FlattenPositionsForTradeAccountPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m FlattenPositionsForTradeAccountBuilder) Finish() FlattenPositionsForTradeAccount {
	return FlattenPositionsForTradeAccount{m.VLSBuilder.Finish()}
}

// Finish
func (m *FlattenPositionsForTradeAccountPointerBuilder) Finish() FlattenPositionsForTradeAccountPointer {
	return FlattenPositionsForTradeAccountPointer{m.VLSPointerBuilder.Finish()}
}

// TradeAccount
func (m FlattenPositionsForTradeAccount) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m FlattenPositionsForTradeAccountBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
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
func (m FlattenPositionsForTradeAccount) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// ClientOrderID
func (m FlattenPositionsForTradeAccountBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
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
func (m FlattenPositionsForTradeAccount) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// FreeFormText
func (m FlattenPositionsForTradeAccountBuilder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
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
func (m FlattenPositionsForTradeAccount) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Unsafe(), 19, 18)
}

// IsAutomatedOrder
func (m FlattenPositionsForTradeAccountBuilder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Unsafe(), 19, 18)
}

// IsAutomatedOrder
func (m FlattenPositionsForTradeAccountPointer) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Ptr, 19, 18)
}

// IsAutomatedOrder
func (m FlattenPositionsForTradeAccountPointerBuilder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Ptr, 19, 18)
}

// SetTradeAccount
func (m *FlattenPositionsForTradeAccountBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetTradeAccount
func (m *FlattenPositionsForTradeAccountPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetClientOrderID
func (m *FlattenPositionsForTradeAccountBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetClientOrderID
func (m *FlattenPositionsForTradeAccountPointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetFreeFormText
func (m *FlattenPositionsForTradeAccountBuilder) SetFreeFormText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetFreeFormText
func (m *FlattenPositionsForTradeAccountPointerBuilder) SetFreeFormText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetIsAutomatedOrder
func (m FlattenPositionsForTradeAccountBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolVLS(m.Unsafe(), 19, 18, value)
}

// SetIsAutomatedOrder
func (m FlattenPositionsForTradeAccountPointerBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolVLS(m.Ptr, 19, 18, value)
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
