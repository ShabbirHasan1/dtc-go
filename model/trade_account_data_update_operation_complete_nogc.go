package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataUpdateOperationCompletePointer struct {
	message.VLSPointer
}

type TradeAccountDataUpdateOperationCompletePointerBuilder struct {
	message.VLSPointerBuilder
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
//     IsError                           = 0
//     ErrorText                         = ""
//     IsDeleteOperation                 = 0
//     IsSymbolLimitsUpdateOperation     = 0
//     IsSymbolCommissionUpdateOperation = 0
//     TradeAccount                      = ""
// }
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataUpdateOperationCompleteDefault)
}

// ToBuilder
func (m TradeAccountDataUpdateOperationCompletePointer) ToBuilder() TradeAccountDataUpdateOperationCompletePointerBuilder {
	return TradeAccountDataUpdateOperationCompletePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataUpdateOperationCompletePointerBuilder) Finish() TradeAccountDataUpdateOperationCompletePointer {
	return TradeAccountDataUpdateOperationCompletePointer{m.VLSPointerBuilder.Finish()}
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
func (m TradeAccountDataUpdateOperationCompletePointer) IsError() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// IsError
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) IsError() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
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
func (m TradeAccountDataUpdateOperationCompletePointer) IsDeleteOperation() uint8 {
	return message.Uint8VLS(m.Ptr, 16, 15)
}

// IsDeleteOperation
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) IsDeleteOperation() uint8 {
	return message.Uint8VLS(m.Ptr, 16, 15)
}

// IsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationCompletePointer) IsSymbolLimitsUpdateOperation() uint8 {
	return message.Uint8VLS(m.Ptr, 17, 16)
}

// IsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) IsSymbolLimitsUpdateOperation() uint8 {
	return message.Uint8VLS(m.Ptr, 17, 16)
}

// IsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationCompletePointer) IsSymbolCommissionUpdateOperation() uint8 {
	return message.Uint8VLS(m.Ptr, 18, 17)
}

// IsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) IsSymbolCommissionUpdateOperation() uint8 {
	return message.Uint8VLS(m.Ptr, 18, 17)
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
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetIsError
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) SetIsError(value uint8) {
	message.SetUint8VLS(m.Ptr, 11, 10, value)
}

// SetErrorText
func (m *TradeAccountDataUpdateOperationCompletePointerBuilder) SetErrorText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 15, 11, value)
}

// SetIsDeleteOperation
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) SetIsDeleteOperation(value uint8) {
	message.SetUint8VLS(m.Ptr, 16, 15, value)
}

// SetIsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) SetIsSymbolLimitsUpdateOperation(value uint8) {
	message.SetUint8VLS(m.Ptr, 17, 16, value)
}

// SetIsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationCompletePointerBuilder) SetIsSymbolCommissionUpdateOperation(value uint8) {
	message.SetUint8VLS(m.Ptr, 18, 17, value)
}

// SetTradeAccount
func (m *TradeAccountDataUpdateOperationCompletePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 22, 18, value)
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
