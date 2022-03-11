package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _TradeAccountDataUpdateOperationCompleteDefault = []byte{22, 0, 147, 39, 22, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataUpdateOperationCompleteSize = 22

type TradeAccountDataUpdateOperationComplete struct {
	message.VLS
}

type TradeAccountDataUpdateOperationCompleteBuilder struct {
	message.VLSBuilder
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
func (m TradeAccountDataUpdateOperationCompleteBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataUpdateOperationCompleteDefault)
}

// ToBuilder
func (m TradeAccountDataUpdateOperationComplete) ToBuilder() TradeAccountDataUpdateOperationCompleteBuilder {
	return TradeAccountDataUpdateOperationCompleteBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataUpdateOperationCompleteBuilder) Finish() TradeAccountDataUpdateOperationComplete {
	return TradeAccountDataUpdateOperationComplete{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataUpdateOperationComplete) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataUpdateOperationCompleteBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// IsError
func (m TradeAccountDataUpdateOperationComplete) IsError() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// IsError
func (m TradeAccountDataUpdateOperationCompleteBuilder) IsError() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// ErrorText
func (m TradeAccountDataUpdateOperationComplete) ErrorText() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// ErrorText
func (m TradeAccountDataUpdateOperationCompleteBuilder) ErrorText() string {
	return message.StringVLS(m.Unsafe(), 15, 11)
}

// IsDeleteOperation
func (m TradeAccountDataUpdateOperationComplete) IsDeleteOperation() uint8 {
	return message.Uint8VLS(m.Unsafe(), 16, 15)
}

// IsDeleteOperation
func (m TradeAccountDataUpdateOperationCompleteBuilder) IsDeleteOperation() uint8 {
	return message.Uint8VLS(m.Unsafe(), 16, 15)
}

// IsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationComplete) IsSymbolLimitsUpdateOperation() uint8 {
	return message.Uint8VLS(m.Unsafe(), 17, 16)
}

// IsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationCompleteBuilder) IsSymbolLimitsUpdateOperation() uint8 {
	return message.Uint8VLS(m.Unsafe(), 17, 16)
}

// IsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationComplete) IsSymbolCommissionUpdateOperation() uint8 {
	return message.Uint8VLS(m.Unsafe(), 18, 17)
}

// IsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationCompleteBuilder) IsSymbolCommissionUpdateOperation() uint8 {
	return message.Uint8VLS(m.Unsafe(), 18, 17)
}

// TradeAccount
func (m TradeAccountDataUpdateOperationComplete) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// TradeAccount
func (m TradeAccountDataUpdateOperationCompleteBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// SetRequestID
func (m TradeAccountDataUpdateOperationCompleteBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetIsError
func (m TradeAccountDataUpdateOperationCompleteBuilder) SetIsError(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 11, 10, value)
}

// SetErrorText
func (m *TradeAccountDataUpdateOperationCompleteBuilder) SetErrorText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 15, 11, value)
}

// SetIsDeleteOperation
func (m TradeAccountDataUpdateOperationCompleteBuilder) SetIsDeleteOperation(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 16, 15, value)
}

// SetIsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationCompleteBuilder) SetIsSymbolLimitsUpdateOperation(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 17, 16, value)
}

// SetIsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationCompleteBuilder) SetIsSymbolCommissionUpdateOperation(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 18, 17, value)
}

// SetTradeAccount
func (m *TradeAccountDataUpdateOperationCompleteBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 22, 18, value)
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
