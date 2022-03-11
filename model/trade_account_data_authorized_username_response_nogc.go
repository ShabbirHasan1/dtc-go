package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataAuthorizedUsernameResponsePointer struct {
	message.VLSPointer
}

type TradeAccountDataAuthorizedUsernameResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeAccountDataAuthorizedUsernameResponse() TradeAccountDataAuthorizedUsernameResponsePointerBuilder {
	a := TradeAccountDataAuthorizedUsernameResponsePointerBuilder{message.AllocVLSBuilder(25)}
	a.Ptr.SetBytes(0, _TradeAccountDataAuthorizedUsernameResponseDefault)
	return a
}

func AllocTradeAccountDataAuthorizedUsernameResponseFrom(b []byte) TradeAccountDataAuthorizedUsernameResponsePointer {
	return TradeAccountDataAuthorizedUsernameResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                       = TradeAccountDataAuthorizedUsernameResponseSize  (25)
//     Type                       = TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_RESPONSE  (10124)
//     BaseSize                   = TradeAccountDataAuthorizedUsernameResponseSize  (25)
//     RequestID                  = 0
//     TradeAccount               = ""
//     Username                   = ""
//     UpdateOperationComplete    = 0
//     UpdateOperationMessageType = 0
//     UpdateOperationErrorText   = ""
// }
func (m TradeAccountDataAuthorizedUsernameResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataAuthorizedUsernameResponseDefault)
}

// ToBuilder
func (m TradeAccountDataAuthorizedUsernameResponsePointer) ToBuilder() TradeAccountDataAuthorizedUsernameResponsePointerBuilder {
	return TradeAccountDataAuthorizedUsernameResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataAuthorizedUsernameResponsePointerBuilder) Finish() TradeAccountDataAuthorizedUsernameResponsePointer {
	return TradeAccountDataAuthorizedUsernameResponsePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameResponsePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameResponsePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameResponsePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameResponsePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m TradeAccountDataAuthorizedUsernameResponsePointer) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Username
func (m TradeAccountDataAuthorizedUsernameResponsePointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// UpdateOperationComplete
func (m TradeAccountDataAuthorizedUsernameResponsePointer) UpdateOperationComplete() uint8 {
	return message.Uint8VLS(m.Ptr, 19, 18)
}

// UpdateOperationComplete
func (m TradeAccountDataAuthorizedUsernameResponsePointerBuilder) UpdateOperationComplete() uint8 {
	return message.Uint8VLS(m.Ptr, 19, 18)
}

// UpdateOperationMessageType
func (m TradeAccountDataAuthorizedUsernameResponsePointer) UpdateOperationMessageType() uint16 {
	return message.Uint16VLS(m.Ptr, 21, 19)
}

// UpdateOperationMessageType
func (m TradeAccountDataAuthorizedUsernameResponsePointerBuilder) UpdateOperationMessageType() uint16 {
	return message.Uint16VLS(m.Ptr, 21, 19)
}

// UpdateOperationErrorText
func (m TradeAccountDataAuthorizedUsernameResponsePointer) UpdateOperationErrorText() string {
	return message.StringVLS(m.Ptr, 25, 21)
}

// UpdateOperationErrorText
func (m TradeAccountDataAuthorizedUsernameResponsePointerBuilder) UpdateOperationErrorText() string {
	return message.StringVLS(m.Ptr, 25, 21)
}

// SetRequestID
func (m TradeAccountDataAuthorizedUsernameResponsePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataAuthorizedUsernameResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataAuthorizedUsernameResponsePointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetUpdateOperationComplete
func (m TradeAccountDataAuthorizedUsernameResponsePointerBuilder) SetUpdateOperationComplete(value uint8) {
	message.SetUint8VLS(m.Ptr, 19, 18, value)
}

// SetUpdateOperationMessageType
func (m TradeAccountDataAuthorizedUsernameResponsePointerBuilder) SetUpdateOperationMessageType(value uint16) {
	message.SetUint16VLS(m.Ptr, 21, 19, value)
}

// SetUpdateOperationErrorText
func (m *TradeAccountDataAuthorizedUsernameResponsePointerBuilder) SetUpdateOperationErrorText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 25, 21, value)
}

// Copy
func (m TradeAccountDataAuthorizedUsernameResponsePointer) Copy(to TradeAccountDataAuthorizedUsernameResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}

// Copy
func (m TradeAccountDataAuthorizedUsernameResponsePointerBuilder) Copy(to TradeAccountDataAuthorizedUsernameResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameResponsePointer) CopyPointer(to TradeAccountDataAuthorizedUsernameResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameResponsePointerBuilder) CopyPointer(to TradeAccountDataAuthorizedUsernameResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}
