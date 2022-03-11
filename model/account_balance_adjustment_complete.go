package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = AccountBalanceAdjustmentCompleteSize  (16)
//     Type          = ACCOUNT_BALANCE_ADJUSTMENT_COMPLETE  (609)
//     RequestID     = 0
//     TransactionID = 0
// }
var _AccountBalanceAdjustmentCompleteDefault = []byte{16, 0, 97, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AccountBalanceAdjustmentCompleteSize = 16

type AccountBalanceAdjustmentComplete struct {
	message.Fixed
}

type AccountBalanceAdjustmentCompleteBuilder struct {
	message.Fixed
}

func NewAccountBalanceAdjustmentCompleteFrom(b []byte) AccountBalanceAdjustmentComplete {
	return AccountBalanceAdjustmentComplete{Fixed: message.NewFixedFrom(b)}
}

func WrapAccountBalanceAdjustmentComplete(b []byte) AccountBalanceAdjustmentComplete {
	return AccountBalanceAdjustmentComplete{Fixed: message.WrapFixed(b)}
}

func NewAccountBalanceAdjustmentComplete() AccountBalanceAdjustmentCompleteBuilder {
	a := AccountBalanceAdjustmentCompleteBuilder{message.NewFixed(16)}
	a.Unsafe().SetBytes(0, _AccountBalanceAdjustmentCompleteDefault)
	return a
}

// Clear
// {
//     Size          = AccountBalanceAdjustmentCompleteSize  (16)
//     Type          = ACCOUNT_BALANCE_ADJUSTMENT_COMPLETE  (609)
//     RequestID     = 0
//     TransactionID = 0
// }
func (m AccountBalanceAdjustmentCompleteBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceAdjustmentCompleteDefault)
}

// ToBuilder
func (m AccountBalanceAdjustmentComplete) ToBuilder() AccountBalanceAdjustmentCompleteBuilder {
	return AccountBalanceAdjustmentCompleteBuilder{m.Fixed}
}

// Finish
func (m AccountBalanceAdjustmentCompleteBuilder) Finish() AccountBalanceAdjustmentComplete {
	return AccountBalanceAdjustmentComplete{m.Fixed}
}

// RequestID
func (m AccountBalanceAdjustmentComplete) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m AccountBalanceAdjustmentCompleteBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// TransactionID
func (m AccountBalanceAdjustmentComplete) TransactionID() int64 {
	return message.Int64Fixed(m.Unsafe(), 16, 8)
}

// TransactionID
func (m AccountBalanceAdjustmentCompleteBuilder) TransactionID() int64 {
	return message.Int64Fixed(m.Unsafe(), 16, 8)
}

// SetRequestID
func (m AccountBalanceAdjustmentCompleteBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetTransactionID
func (m AccountBalanceAdjustmentCompleteBuilder) SetTransactionID(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, value)
}

// Copy
func (m AccountBalanceAdjustmentComplete) Copy(to AccountBalanceAdjustmentCompleteBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTransactionID(m.TransactionID())
}

// Copy
func (m AccountBalanceAdjustmentCompleteBuilder) Copy(to AccountBalanceAdjustmentCompleteBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTransactionID(m.TransactionID())
}

// CopyPointer
func (m AccountBalanceAdjustmentComplete) CopyPointer(to AccountBalanceAdjustmentCompletePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTransactionID(m.TransactionID())
}

// CopyPointer
func (m AccountBalanceAdjustmentCompleteBuilder) CopyPointer(to AccountBalanceAdjustmentCompletePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTransactionID(m.TransactionID())
}
