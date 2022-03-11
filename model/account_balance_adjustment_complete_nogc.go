package model

import (
	"github.com/moontrade/dtc-go/message"
)

type AccountBalanceAdjustmentCompletePointer struct {
	message.FixedPointer
}

type AccountBalanceAdjustmentCompletePointerBuilder struct {
	message.FixedPointer
}

func AllocAccountBalanceAdjustmentComplete() AccountBalanceAdjustmentCompletePointerBuilder {
	a := AccountBalanceAdjustmentCompletePointerBuilder{message.AllocFixed(16)}
	a.Ptr.SetBytes(0, _AccountBalanceAdjustmentCompleteDefault)
	return a
}

func AllocAccountBalanceAdjustmentCompleteFrom(b []byte) AccountBalanceAdjustmentCompletePointer {
	return AccountBalanceAdjustmentCompletePointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = AccountBalanceAdjustmentCompleteSize  (16)
//     Type          = ACCOUNT_BALANCE_ADJUSTMENT_COMPLETE  (609)
//     RequestID     = 0
//     TransactionID = 0
// }
func (m AccountBalanceAdjustmentCompletePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceAdjustmentCompleteDefault)
}

// ToBuilder
func (m AccountBalanceAdjustmentCompletePointer) ToBuilder() AccountBalanceAdjustmentCompletePointerBuilder {
	return AccountBalanceAdjustmentCompletePointerBuilder{m.FixedPointer}
}

// Finish
func (m *AccountBalanceAdjustmentCompletePointerBuilder) Finish() AccountBalanceAdjustmentCompletePointer {
	return AccountBalanceAdjustmentCompletePointer{m.FixedPointer}
}

// RequestID
func (m AccountBalanceAdjustmentCompletePointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m AccountBalanceAdjustmentCompletePointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// TransactionID
func (m AccountBalanceAdjustmentCompletePointer) TransactionID() int64 {
	return message.Int64Fixed(m.Ptr, 16, 8)
}

// TransactionID
func (m AccountBalanceAdjustmentCompletePointerBuilder) TransactionID() int64 {
	return message.Int64Fixed(m.Ptr, 16, 8)
}

// SetRequestID
func (m AccountBalanceAdjustmentCompletePointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTransactionID
func (m AccountBalanceAdjustmentCompletePointerBuilder) SetTransactionID(value int64) {
	message.SetInt64Fixed(m.Ptr, 16, 8, value)
}

// Copy
func (m AccountBalanceAdjustmentCompletePointer) Copy(to AccountBalanceAdjustmentCompleteBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTransactionID(m.TransactionID())
}

// Copy
func (m AccountBalanceAdjustmentCompletePointerBuilder) Copy(to AccountBalanceAdjustmentCompleteBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTransactionID(m.TransactionID())
}

// CopyPointer
func (m AccountBalanceAdjustmentCompletePointer) CopyPointer(to AccountBalanceAdjustmentCompletePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTransactionID(m.TransactionID())
}

// CopyPointer
func (m AccountBalanceAdjustmentCompletePointerBuilder) CopyPointer(to AccountBalanceAdjustmentCompletePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTransactionID(m.TransactionID())
}
