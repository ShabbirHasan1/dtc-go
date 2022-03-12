package model

import (
	"github.com/moontrade/dtc-go/message"
)

const AccountBalanceAdjustmentCompleteSize = 16

// {
//     Size          = AccountBalanceAdjustmentCompleteSize  (16)
//     Type          = ACCOUNT_BALANCE_ADJUSTMENT_COMPLETE  (609)
//     RequestID     = 0
//     TransactionID = 0
// }
var _AccountBalanceAdjustmentCompleteDefault = []byte{16, 0, 97, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type AccountBalanceAdjustmentComplete struct {
	message.Fixed
}

type AccountBalanceAdjustmentCompleteBuilder struct {
	message.Fixed
}

type AccountBalanceAdjustmentCompletePointer struct {
	message.FixedPointer
}

type AccountBalanceAdjustmentCompletePointerBuilder struct {
	message.FixedPointer
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
func (m AccountBalanceAdjustmentCompleteBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceAdjustmentCompleteDefault)
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
func (m AccountBalanceAdjustmentComplete) ToBuilder() AccountBalanceAdjustmentCompleteBuilder {
	return AccountBalanceAdjustmentCompleteBuilder{m.Fixed}
}

// ToBuilder
func (m AccountBalanceAdjustmentCompletePointer) ToBuilder() AccountBalanceAdjustmentCompletePointerBuilder {
	return AccountBalanceAdjustmentCompletePointerBuilder{m.FixedPointer}
}

// Finish
func (m AccountBalanceAdjustmentCompleteBuilder) Finish() AccountBalanceAdjustmentComplete {
	return AccountBalanceAdjustmentComplete{m.Fixed}
}

// Finish
func (m *AccountBalanceAdjustmentCompletePointerBuilder) Finish() AccountBalanceAdjustmentCompletePointer {
	return AccountBalanceAdjustmentCompletePointer{m.FixedPointer}
}

// RequestID
func (m AccountBalanceAdjustmentComplete) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m AccountBalanceAdjustmentCompleteBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m AccountBalanceAdjustmentComplete) TransactionID() int64 {
	return message.Int64Fixed(m.Unsafe(), 16, 8)
}

// TransactionID
func (m AccountBalanceAdjustmentCompleteBuilder) TransactionID() int64 {
	return message.Int64Fixed(m.Unsafe(), 16, 8)
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
func (m AccountBalanceAdjustmentCompleteBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m AccountBalanceAdjustmentCompletePointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTransactionID
func (m AccountBalanceAdjustmentCompleteBuilder) SetTransactionID(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 16, 8, value)
}

// SetTransactionID
func (m AccountBalanceAdjustmentCompletePointerBuilder) SetTransactionID(value int64) {
	message.SetInt64Fixed(m.Ptr, 16, 8, value)
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
