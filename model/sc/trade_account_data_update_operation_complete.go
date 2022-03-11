package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeAccountDataUpdateOperationCompleteSize = 22

// {
//     Size                              = TradeAccountDataUpdateOperationCompleteSize  (22)
//     Type                              = TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE  (10131)
//     BaseSize                          = TradeAccountDataUpdateOperationCompleteSize  (22)
//     RequestID                         = 0
//     IsError                           = false
//     ErrorText                         = ""
//     IsDeleteOperation                 = false
//     IsSymbolLimitsUpdateOperation     = false
//     IsSymbolCommissionUpdateOperation = false
//     TradeAccount                      = ""
// }
var _TradeAccountDataUpdateOperationCompleteDefault = []byte{22, 0, 147, 39, 22, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataUpdateOperationComplete struct {
	message.VLS
}

type TradeAccountDataUpdateOperationCompleteBuilder struct {
	message.VLSBuilder
}

type TradeAccountDataUpdateOperationCompletePointer struct {
	message.VLSPointer
}

type TradeAccountDataUpdateOperationCompletePointerBuilder struct {
	message.VLSPointerBuilder
}

func NewTradeAccountDataUpdateOperationCompleteFrom(b []byte) TradeAccountDataUpdateOperationComplete {
	return TradeAccountDataUpdateOperationComplete{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataUpdateOperationComplete(b []byte) TradeAccountDataUpdateOperationComplete {
	return TradeAccountDataUpdateOperationComplete{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataUpdateOperationComplete() TradeAccountDataUpdateOperationCompleteBuilder {
	a := TradeAccountDataUpdateOperationCompleteBuilder{message.NewVLSBuilder(22)}
	a.Unsafe().SetBytes(0, _TradeAccountDataUpdateOperationCompleteDefault)
	return a
}

func AllocTradeAccountDataUpdateOperationComplete() TradeAccountDataUpdateOperationCompletePointerBuilder {
	a := TradeAccountDataUpdateOperationCompletePointerBuilder{message.AllocVLSBuilder(22)}
	a.Ptr.SetBytes(0, _TradeAccountDataUpdateOperationCompleteDefault)
	return a
}

func AllocTradeAccountDataUpdateOperationCompleteFrom(b []byte) TradeAccountDataUpdateOperationCompletePointer {
	return TradeAccountDataUpdateOperationCompletePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                              = TradeAccountDataUpdateOperationCompleteSize  (22)
//     Type                              = TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE  (10131)
//     BaseSize                          = TradeAccountDataUpdateOperationCompleteSize  (22)
//     RequestID                         = 0
//     IsError                           = false
//     ErrorText                         = ""
//     IsDeleteOperation                 = false
//     IsSymbolLimitsUpdateOperation     = false
//     IsSymbolCommissionUpdateOperation = false
//     TradeAccount                      = ""
// }
func (m TradeAccountDataUpdateOperationCompleteBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataUpdateOperationCompleteDefault)
}

// Clear
// {
//     Size                              = TradeAccountDataUpdateOperationCompleteSize  (22)
//     Type                              = TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE  (10131)
//     BaseSize                          = TradeAccountDataUpdateOperationCompleteSize  (22)
//     RequestID                         = 0
//     IsError                           = false
//     ErrorText                         = ""
//     IsDeleteOperation                 = false
//     IsSymbolLimitsUpdateOperation     = false
//     IsSymbolCommissionUpdateOperation = false
//     TradeAccount                      = ""
// }
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataUpdateOperationCompleteDefault)
}

// ToBuilder
func (m TradeAccountDataUpdateOperationComplete) ToBuilder() TradeAccountDataUpdateOperationCompleteBuilder {
	return TradeAccountDataUpdateOperationCompleteBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradeAccountDataUpdateOperationCompletePointer) ToBuilder() TradeAccountDataUpdateOperationCompletePointerBuilder {
	return TradeAccountDataUpdateOperationCompletePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m TradeAccountDataUpdateOperationCompleteBuilder) Finish() TradeAccountDataUpdateOperationComplete {
	return TradeAccountDataUpdateOperationComplete{m.VLSBuilder.Finish()}
}

// Finish
func (m *TradeAccountDataUpdateOperationCompletePointerBuilder) Finish() TradeAccountDataUpdateOperationCompletePointer {
	return TradeAccountDataUpdateOperationCompletePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataUpdateOperationComplete) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataUpdateOperationCompleteBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataUpdateOperationCompletePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// IsError
func (m TradeAccountDataUpdateOperationComplete) IsError() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsError
func (m TradeAccountDataUpdateOperationCompleteBuilder) IsError() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsError
func (m TradeAccountDataUpdateOperationCompletePointer) IsError() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// IsError
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) IsError() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// ErrorText
func (m TradeAccountDataUpdateOperationComplete) ErrorText() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// ErrorText
func (m TradeAccountDataUpdateOperationCompleteBuilder) ErrorText() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// ErrorText
func (m TradeAccountDataUpdateOperationCompletePointer) ErrorText() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// ErrorText
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) ErrorText() string {
	return message.StringVLS(m.Ptr, 15, 11)
}

// IsDeleteOperation
func (m TradeAccountDataUpdateOperationComplete) IsDeleteOperation() bool {
	return message.BoolVLS(m.Unsafe(), 16, 15)
}

// IsDeleteOperation
func (m TradeAccountDataUpdateOperationCompleteBuilder) IsDeleteOperation() bool {
	return message.BoolVLS(m.Unsafe(), 16, 15)
}

// IsDeleteOperation
func (m TradeAccountDataUpdateOperationCompletePointer) IsDeleteOperation() bool {
	return message.BoolVLS(m.Ptr, 16, 15)
}

// IsDeleteOperation
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) IsDeleteOperation() bool {
	return message.BoolVLS(m.Ptr, 16, 15)
}

// IsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationComplete) IsSymbolLimitsUpdateOperation() bool {
	return message.BoolVLS(m.Unsafe(), 17, 16)
}

// IsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationCompleteBuilder) IsSymbolLimitsUpdateOperation() bool {
	return message.BoolVLS(m.Unsafe(), 17, 16)
}

// IsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationCompletePointer) IsSymbolLimitsUpdateOperation() bool {
	return message.BoolVLS(m.Ptr, 17, 16)
}

// IsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) IsSymbolLimitsUpdateOperation() bool {
	return message.BoolVLS(m.Ptr, 17, 16)
}

// IsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationComplete) IsSymbolCommissionUpdateOperation() bool {
	return message.BoolVLS(m.Unsafe(), 18, 17)
}

// IsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationCompleteBuilder) IsSymbolCommissionUpdateOperation() bool {
	return message.BoolVLS(m.Unsafe(), 18, 17)
}

// IsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationCompletePointer) IsSymbolCommissionUpdateOperation() bool {
	return message.BoolVLS(m.Ptr, 18, 17)
}

// IsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) IsSymbolCommissionUpdateOperation() bool {
	return message.BoolVLS(m.Ptr, 18, 17)
}

// TradeAccount
func (m TradeAccountDataUpdateOperationComplete) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// TradeAccount
func (m TradeAccountDataUpdateOperationCompleteBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// TradeAccount
func (m TradeAccountDataUpdateOperationCompletePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// TradeAccount
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// SetRequestID
func (m TradeAccountDataUpdateOperationCompleteBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetIsError
func (m TradeAccountDataUpdateOperationCompleteBuilder) SetIsError(value bool) {
	message.SetBoolVLS(m.Unsafe(), 11, 10, value)
}

// SetIsError
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) SetIsError(value bool) {
	message.SetBoolVLS(m.Ptr, 11, 10, value)
}

// SetErrorText
func (m *TradeAccountDataUpdateOperationCompleteBuilder) SetErrorText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 15, 11, value)
}

// SetErrorText
func (m *TradeAccountDataUpdateOperationCompletePointerBuilder) SetErrorText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetIsDeleteOperation
func (m TradeAccountDataUpdateOperationCompleteBuilder) SetIsDeleteOperation(value bool) {
	message.SetBoolVLS(m.Unsafe(), 16, 15, value)
}

// SetIsDeleteOperation
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) SetIsDeleteOperation(value bool) {
	message.SetBoolVLS(m.Ptr, 16, 15, value)
}

// SetIsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationCompleteBuilder) SetIsSymbolLimitsUpdateOperation(value bool) {
	message.SetBoolVLS(m.Unsafe(), 17, 16, value)
}

// SetIsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) SetIsSymbolLimitsUpdateOperation(value bool) {
	message.SetBoolVLS(m.Ptr, 17, 16, value)
}

// SetIsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationCompleteBuilder) SetIsSymbolCommissionUpdateOperation(value bool) {
	message.SetBoolVLS(m.Unsafe(), 18, 17, value)
}

// SetIsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) SetIsSymbolCommissionUpdateOperation(value bool) {
	message.SetBoolVLS(m.Ptr, 18, 17, value)
}

// SetTradeAccount
func (m *TradeAccountDataUpdateOperationCompleteBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 22, 18, value)
}

// SetTradeAccount
func (m *TradeAccountDataUpdateOperationCompletePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 22, 18, value)
}

// Copy
func (m TradeAccountDataUpdateOperationComplete) Copy(to TradeAccountDataUpdateOperationCompleteBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsError(m.IsError())
	to.SetErrorText(m.ErrorText())
	to.SetIsDeleteOperation(m.IsDeleteOperation())
	to.SetIsSymbolLimitsUpdateOperation(m.IsSymbolLimitsUpdateOperation())
	to.SetIsSymbolCommissionUpdateOperation(m.IsSymbolCommissionUpdateOperation())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m TradeAccountDataUpdateOperationCompleteBuilder) Copy(to TradeAccountDataUpdateOperationCompleteBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsError(m.IsError())
	to.SetErrorText(m.ErrorText())
	to.SetIsDeleteOperation(m.IsDeleteOperation())
	to.SetIsSymbolLimitsUpdateOperation(m.IsSymbolLimitsUpdateOperation())
	to.SetIsSymbolCommissionUpdateOperation(m.IsSymbolCommissionUpdateOperation())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataUpdateOperationComplete) CopyPointer(to TradeAccountDataUpdateOperationCompletePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsError(m.IsError())
	to.SetErrorText(m.ErrorText())
	to.SetIsDeleteOperation(m.IsDeleteOperation())
	to.SetIsSymbolLimitsUpdateOperation(m.IsSymbolLimitsUpdateOperation())
	to.SetIsSymbolCommissionUpdateOperation(m.IsSymbolCommissionUpdateOperation())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataUpdateOperationCompleteBuilder) CopyPointer(to TradeAccountDataUpdateOperationCompletePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsError(m.IsError())
	to.SetErrorText(m.ErrorText())
	to.SetIsDeleteOperation(m.IsDeleteOperation())
	to.SetIsSymbolLimitsUpdateOperation(m.IsSymbolLimitsUpdateOperation())
	to.SetIsSymbolCommissionUpdateOperation(m.IsSymbolCommissionUpdateOperation())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m TradeAccountDataUpdateOperationCompletePointer) Copy(to TradeAccountDataUpdateOperationCompleteBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsError(m.IsError())
	to.SetErrorText(m.ErrorText())
	to.SetIsDeleteOperation(m.IsDeleteOperation())
	to.SetIsSymbolLimitsUpdateOperation(m.IsSymbolLimitsUpdateOperation())
	to.SetIsSymbolCommissionUpdateOperation(m.IsSymbolCommissionUpdateOperation())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) Copy(to TradeAccountDataUpdateOperationCompleteBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsError(m.IsError())
	to.SetErrorText(m.ErrorText())
	to.SetIsDeleteOperation(m.IsDeleteOperation())
	to.SetIsSymbolLimitsUpdateOperation(m.IsSymbolLimitsUpdateOperation())
	to.SetIsSymbolCommissionUpdateOperation(m.IsSymbolCommissionUpdateOperation())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataUpdateOperationCompletePointer) CopyPointer(to TradeAccountDataUpdateOperationCompletePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsError(m.IsError())
	to.SetErrorText(m.ErrorText())
	to.SetIsDeleteOperation(m.IsDeleteOperation())
	to.SetIsSymbolLimitsUpdateOperation(m.IsSymbolLimitsUpdateOperation())
	to.SetIsSymbolCommissionUpdateOperation(m.IsSymbolCommissionUpdateOperation())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) CopyPointer(to TradeAccountDataUpdateOperationCompletePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsError(m.IsError())
	to.SetErrorText(m.ErrorText())
	to.SetIsDeleteOperation(m.IsDeleteOperation())
	to.SetIsSymbolLimitsUpdateOperation(m.IsSymbolLimitsUpdateOperation())
	to.SetIsSymbolCommissionUpdateOperation(m.IsSymbolCommissionUpdateOperation())
	to.SetTradeAccount(m.TradeAccount())
}
