package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _TradeAccountDataUsernameToShareWithResponseDefault = []byte{26, 0, 143, 39, 26, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataUsernameToShareWithResponseSize = 26

type TradeAccountDataUsernameToShareWithResponse struct {
	message.VLS
}

type TradeAccountDataUsernameToShareWithResponseBuilder struct {
	message.VLSBuilder
}

func NewTradeAccountDataUsernameToShareWithResponseFrom(b []byte) TradeAccountDataUsernameToShareWithResponse {
	return TradeAccountDataUsernameToShareWithResponse{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataUsernameToShareWithResponse(b []byte) TradeAccountDataUsernameToShareWithResponse {
	return TradeAccountDataUsernameToShareWithResponse{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataUsernameToShareWithResponse() TradeAccountDataUsernameToShareWithResponseBuilder {
	a := TradeAccountDataUsernameToShareWithResponseBuilder{message.NewVLSBuilder(26)}
	a.Unsafe().SetBytes(0, _TradeAccountDataUsernameToShareWithResponseDefault)
	return a
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
func (m TradeAccountDataUsernameToShareWithResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataUsernameToShareWithResponseDefault)
}

// ToBuilder
func (m TradeAccountDataUsernameToShareWithResponse) ToBuilder() TradeAccountDataUsernameToShareWithResponseBuilder {
	return TradeAccountDataUsernameToShareWithResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataUsernameToShareWithResponseBuilder) Finish() TradeAccountDataUsernameToShareWithResponse {
	return TradeAccountDataUsernameToShareWithResponse{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataUsernameToShareWithResponse) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithResponseBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Username
func (m TradeAccountDataUsernameToShareWithResponse) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Username
func (m TradeAccountDataUsernameToShareWithResponseBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithResponse) IsReadOnly() uint8 {
	return message.Uint8VLS(m.Unsafe(), 19, 18)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithResponseBuilder) IsReadOnly() uint8 {
	return message.Uint8VLS(m.Unsafe(), 19, 18)
}

// UpdateOperationComplete
func (m TradeAccountDataUsernameToShareWithResponse) UpdateOperationComplete() uint8 {
	return message.Uint8VLS(m.Unsafe(), 20, 19)
}

// UpdateOperationComplete
func (m TradeAccountDataUsernameToShareWithResponseBuilder) UpdateOperationComplete() uint8 {
	return message.Uint8VLS(m.Unsafe(), 20, 19)
}

// UpdateOperationMessageType
func (m TradeAccountDataUsernameToShareWithResponse) UpdateOperationMessageType() uint16 {
	return message.Uint16VLS(m.Unsafe(), 22, 20)
}

// UpdateOperationMessageType
func (m TradeAccountDataUsernameToShareWithResponseBuilder) UpdateOperationMessageType() uint16 {
	return message.Uint16VLS(m.Unsafe(), 22, 20)
}

// UpdateOperationErrorText
func (m TradeAccountDataUsernameToShareWithResponse) UpdateOperationErrorText() string {
	return message.StringVLS(m.Unsafe(), 26, 22)
}

// UpdateOperationErrorText
func (m TradeAccountDataUsernameToShareWithResponseBuilder) UpdateOperationErrorText() string {
	return message.StringVLS(m.Unsafe(), 26, 22)
}

// SetRequestID
func (m TradeAccountDataUsernameToShareWithResponseBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataUsernameToShareWithResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataUsernameToShareWithResponseBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetIsReadOnly
func (m TradeAccountDataUsernameToShareWithResponseBuilder) SetIsReadOnly(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 19, 18, value)
}

// SetUpdateOperationComplete
func (m TradeAccountDataUsernameToShareWithResponseBuilder) SetUpdateOperationComplete(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 20, 19, value)
}

// SetUpdateOperationMessageType
func (m TradeAccountDataUsernameToShareWithResponseBuilder) SetUpdateOperationMessageType(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 22, 20, value)
}

// SetUpdateOperationErrorText
func (m *TradeAccountDataUsernameToShareWithResponseBuilder) SetUpdateOperationErrorText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 26, 22, value)
}

// Copy
func (m TradeAccountDataUsernameToShareWithResponse) Copy(to TradeAccountDataUsernameToShareWithResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}

// Copy
func (m TradeAccountDataUsernameToShareWithResponseBuilder) Copy(to TradeAccountDataUsernameToShareWithResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithResponse) CopyPointer(to TradeAccountDataUsernameToShareWithResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithResponseBuilder) CopyPointer(to TradeAccountDataUsernameToShareWithResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
	to.SetUpdateOperationComplete(m.UpdateOperationComplete())
	to.SetUpdateOperationMessageType(m.UpdateOperationMessageType())
	to.SetUpdateOperationErrorText(m.UpdateOperationErrorText())
}
