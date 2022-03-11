package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _TradeAccountDataAuthorizedUsernameResponseDefault = []byte{25, 0, 140, 39, 25, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataAuthorizedUsernameResponseSize = 25

type TradeAccountDataAuthorizedUsernameResponse struct {
	message.VLS
}

type TradeAccountDataAuthorizedUsernameResponseBuilder struct {
	message.VLSBuilder
}

func NewTradeAccountDataAuthorizedUsernameResponseFrom(b []byte) TradeAccountDataAuthorizedUsernameResponse {
	return TradeAccountDataAuthorizedUsernameResponse{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataAuthorizedUsernameResponse(b []byte) TradeAccountDataAuthorizedUsernameResponse {
	return TradeAccountDataAuthorizedUsernameResponse{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataAuthorizedUsernameResponse() TradeAccountDataAuthorizedUsernameResponseBuilder {
	a := TradeAccountDataAuthorizedUsernameResponseBuilder{message.NewVLSBuilder(25)}
	a.Unsafe().SetBytes(0, _TradeAccountDataAuthorizedUsernameResponseDefault)
	return a
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
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataAuthorizedUsernameResponseDefault)
}

// ToBuilder
func (m TradeAccountDataAuthorizedUsernameResponse) ToBuilder() TradeAccountDataAuthorizedUsernameResponseBuilder {
	return TradeAccountDataAuthorizedUsernameResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) Finish() TradeAccountDataAuthorizedUsernameResponse {
	return TradeAccountDataAuthorizedUsernameResponse{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameResponse) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Username
func (m TradeAccountDataAuthorizedUsernameResponse) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Username
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// UpdateOperationComplete
func (m TradeAccountDataAuthorizedUsernameResponse) UpdateOperationComplete() uint8 {
	return message.Uint8VLS(m.Unsafe(), 19, 18)
}

// UpdateOperationComplete
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) UpdateOperationComplete() uint8 {
	return message.Uint8VLS(m.Unsafe(), 19, 18)
}

// UpdateOperationMessageType
func (m TradeAccountDataAuthorizedUsernameResponse) UpdateOperationMessageType() uint16 {
	return message.Uint16VLS(m.Unsafe(), 21, 19)
}

// UpdateOperationMessageType
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) UpdateOperationMessageType() uint16 {
	return message.Uint16VLS(m.Unsafe(), 21, 19)
}

// UpdateOperationErrorText
func (m TradeAccountDataAuthorizedUsernameResponse) UpdateOperationErrorText() string {
	return message.StringVLS(m.Unsafe(), 25, 21)
}

// UpdateOperationErrorText
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) UpdateOperationErrorText() string {
	return message.StringVLS(m.Unsafe(), 25, 21)
}

// SetRequestID
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataAuthorizedUsernameResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataAuthorizedUsernameResponseBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetUpdateOperationComplete
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) SetUpdateOperationComplete(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 19, 18, value)
}

// SetUpdateOperationMessageType
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) SetUpdateOperationMessageType(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 21, 19, value)
}

// SetUpdateOperationErrorText
func (m *TradeAccountDataAuthorizedUsernameResponseBuilder) SetUpdateOperationErrorText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 25, 21, value)
}

// Copy
func (m TradeAccountDataAuthorizedUsernameResponse) Copy(to TradeAccountDataAuthorizedUsernameResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}

// Copy
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) Copy(to TradeAccountDataAuthorizedUsernameResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameResponse) CopyPointer(to TradeAccountDataAuthorizedUsernameResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameResponseBuilder) CopyPointer(to TradeAccountDataAuthorizedUsernameResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}
