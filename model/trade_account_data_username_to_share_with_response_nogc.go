package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataUsernameToShareWithResponsePointer struct {
	message.VLSPointer
}

type TradeAccountDataUsernameToShareWithResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocTradeAccountDataUsernameToShareWithResponse() TradeAccountDataUsernameToShareWithResponsePointerBuilder {
	a := TradeAccountDataUsernameToShareWithResponsePointerBuilder{message.AllocVLSBuilder(26)}
	a.Ptr.SetBytes(0, _TradeAccountDataUsernameToShareWithResponseDefault)
	return a
}

func AllocTradeAccountDataUsernameToShareWithResponseFrom(b []byte) TradeAccountDataUsernameToShareWithResponsePointer {
	return TradeAccountDataUsernameToShareWithResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                       = TradeAccountDataUsernameToShareWithResponseSize  (26)
//     Type                       = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_RESPONSE  (10127)
//     BaseSize                   = TradeAccountDataUsernameToShareWithResponseSize  (26)
//     RequestID                  = 0
//     TradeAccount               = ""
//     Username                   = ""
//     IsReadOnly                 = 0
//     UpdateOperationComplete    = 0
//     UpdateOperationMessageType = 0
//     UpdateOperationErrorText   = ""
// }
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataUsernameToShareWithResponseDefault)
}

// ToBuilder
func (m TradeAccountDataUsernameToShareWithResponsePointer) ToBuilder() TradeAccountDataUsernameToShareWithResponsePointerBuilder {
	return TradeAccountDataUsernameToShareWithResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataUsernameToShareWithResponsePointerBuilder) Finish() TradeAccountDataUsernameToShareWithResponsePointer {
	return TradeAccountDataUsernameToShareWithResponsePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataUsernameToShareWithResponsePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithResponsePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m TradeAccountDataUsernameToShareWithResponsePointer) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Username
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithResponsePointer) IsReadOnly() uint8 {
	return message.Uint8VLS(m.Ptr, 19, 18)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) IsReadOnly() uint8 {
	return message.Uint8VLS(m.Ptr, 19, 18)
}

// UpdateOperationComplete
func (m TradeAccountDataUsernameToShareWithResponsePointer) UpdateOperationComplete() uint8 {
	return message.Uint8VLS(m.Ptr, 20, 19)
}

// UpdateOperationComplete
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) UpdateOperationComplete() uint8 {
	return message.Uint8VLS(m.Ptr, 20, 19)
}

// UpdateOperationMessageType
func (m TradeAccountDataUsernameToShareWithResponsePointer) UpdateOperationMessageType() uint16 {
	return message.Uint16VLS(m.Ptr, 22, 20)
}

// UpdateOperationMessageType
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) UpdateOperationMessageType() uint16 {
	return message.Uint16VLS(m.Ptr, 22, 20)
}

// UpdateOperationErrorText
func (m TradeAccountDataUsernameToShareWithResponsePointer) UpdateOperationErrorText() string {
	return message.StringVLS(m.Ptr, 26, 22)
}

// UpdateOperationErrorText
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) UpdateOperationErrorText() string {
	return message.StringVLS(m.Ptr, 26, 22)
}

// SetRequestID
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataUsernameToShareWithResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataUsernameToShareWithResponsePointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetIsReadOnly
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) SetIsReadOnly(value uint8) {
	message.SetUint8VLS(m.Ptr, 19, 18, value)
}

// SetUpdateOperationComplete
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) SetUpdateOperationComplete(value uint8) {
	message.SetUint8VLS(m.Ptr, 20, 19, value)
}

// SetUpdateOperationMessageType
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) SetUpdateOperationMessageType(value uint16) {
	message.SetUint16VLS(m.Ptr, 22, 20, value)
}

// SetUpdateOperationErrorText
func (m *TradeAccountDataUsernameToShareWithResponsePointerBuilder) SetUpdateOperationErrorText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 26, 22, value)
}

// Copy
func (m TradeAccountDataUsernameToShareWithResponsePointer) Copy(to TradeAccountDataUsernameToShareWithResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}

// Copy
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) Copy(to TradeAccountDataUsernameToShareWithResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithResponsePointer) CopyPointer(to TradeAccountDataUsernameToShareWithResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithResponsePointerBuilder) CopyPointer(to TradeAccountDataUsernameToShareWithResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}
